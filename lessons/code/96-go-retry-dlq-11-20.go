package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
GO RETRY + DLQ (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/96-go-retry-dlq-11-20.go
2) POST /jobs with type:
   - "unstable" (succeeds after retries)
   - "always_fail" (moves to DLQ)
3) GET /jobs and GET /dlq

Extra context:
- lessons/notes/183-retries-and-backoff-first-principles.md
- lessons/notes/184-dead-letter-queue-gotchas.md
*/

type JobStatus string

const (
	StatusQueued     JobStatus = "queued"
	StatusProcessing JobStatus = "processing"
	StatusDone       JobStatus = "done"
	StatusFailed     JobStatus = "failed"
)

type Job struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    JobStatus `json:"status"`
	Attempts  int       `json:"attempts"`
	LastError string    `json:"last_error,omitempty"`
}

type DLQEntry struct {
	JobID      int64  `json:"job_id"`
	JobType    string `json:"job_type"`
	Payload    string `json:"payload"`
	Attempts   int    `json:"attempts"`
	Reason     string `json:"reason"`
	OccurredAt string `json:"occurred_at"`
}

type JobStore struct {
	mu     sync.Mutex
	nextID int64
	items  map[int64]Job
}

func NewJobStore() *JobStore {
	return &JobStore{nextID: 1, items: map[int64]Job{}}
}

func (s *JobStore) Add(t string, payload string) Job {
	s.mu.Lock()
	defer s.mu.Unlock()
	job := Job{ID: s.nextID, Type: t, Payload: payload, Status: StatusQueued}
	s.nextID++
	s.items[job.ID] = job
	return job
}

func (s *JobStore) Update(job Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[job.ID] = job
}

func (s *JobStore) List() []Job {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Job, 0, len(s.items))
	for _, j := range s.items {
		out = append(out, j)
	}
	return out
}

type DLQStore struct {
	mu      sync.Mutex
	entries []DLQEntry
}

func NewDLQStore() *DLQStore {
	return &DLQStore{entries: []DLQEntry{}}
}

func (d *DLQStore) Add(e DLQEntry) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries = append(d.entries, e)
}

func (d *DLQStore) List() []DLQEntry {
	d.mu.Lock()
	defer d.mu.Unlock()
	out := make([]DLQEntry, len(d.entries))
	copy(out, d.entries)
	return out
}

type JobQueue struct {
	ch chan Job
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{ch: make(chan Job, size)}
}

func (q *JobQueue) Enqueue(job Job) {
	q.ch <- job
}

type JobHandler func(job Job) error

// LESSON 11: Retry policy
// Why this matters: explicit policy avoids infinite retries.
type RetryPolicy struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

func (p RetryPolicy) NextDelay(attempt int) time.Duration {
	if attempt <= 0 {
		return p.BaseDelay
	}
	return time.Duration(attempt) * p.BaseDelay
}

// LESSON 12-18: Worker with retry and DLQ
// Why this matters: failed jobs become visible and recoverable work items.
func StartRetryingWorker(store *JobStore, dlq *DLQStore, q *JobQueue, handler JobHandler, policy RetryPolicy) {
	go func() {
		for job := range q.ch {
			job.Status = StatusProcessing
			job.Attempts++
			store.Update(job)

			err := handler(job)
			if err == nil {
				job.Status = StatusDone
				job.LastError = ""
				store.Update(job)
				continue
			}

			job.LastError = err.Error()
			if job.Attempts < policy.MaxAttempts {
				job.Status = StatusQueued
				store.Update(job)
				time.Sleep(policy.NextDelay(job.Attempts))
				q.Enqueue(job)
				continue
			}

			job.Status = StatusFailed
			store.Update(job)
			dlq.Add(DLQEntry{
				JobID:      job.ID,
				JobType:    job.Type,
				Payload:    job.Payload,
				Attempts:   job.Attempts,
				Reason:     job.LastError,
				OccurredAt: time.Now().Format(time.RFC3339),
			})
		}
	}()
}

// `unstable` succeeds on 2nd attempt; `always_fail` always fails.
func simulatedRetryHandler(job Job) error {
	if job.Type == "always_fail" {
		return fmt.Errorf("permanent failure")
	}
	if job.Type == "unstable" && job.Attempts < 2 {
		return fmt.Errorf("temporary failure")
	}
	return nil
}

type createJobRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func buildMux(store *JobStore, dlq *DLQStore, q *JobQueue) *http.ServeMux {
	mux := http.NewServeMux()

	// LESSON 19: enqueue + observe endpoints
	// Why this matters: operational visibility is part of async design.
	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req createJobRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
				return
			}
			if req.Type == "" {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "type is required"})
				return
			}
			job := store.Add(req.Type, req.Payload)
			q.Enqueue(job)
			writeJSON(w, http.StatusAccepted, job)
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.List())
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	mux.HandleFunc("/dlq", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, dlq.List())
	})

	// LESSON 20: health endpoint
	// Why this matters: standard operational contract.
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	return mux
}

func main() {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(100)
	policy := RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:   100 * time.Millisecond,
	}
	StartRetryingWorker(store, dlq, queue, simulatedRetryHandler, policy)

	mux := buildMux(store, dlq, queue)
	addr := ":8092"
	fmt.Println("Go retry + DLQ lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Retry + DLQ 11-20
