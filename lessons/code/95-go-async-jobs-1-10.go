package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
GO ASYNC JOBS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/95-go-async-jobs-1-10.go
2) POST /jobs {"type":"send_email","payload":"welcome user"}
3) GET /jobs to watch status transitions

Extra context:
- lessons/notes/182-async-jobs-first-principles.md
*/

type JobStatus string

const (
	StatusQueued     JobStatus = "queued"
	StatusProcessing JobStatus = "processing"
	StatusDone       JobStatus = "done"
	StatusFailed     JobStatus = "failed"
)

// LESSON 1: Domain model for async work
// Why this matters: explicit status lifecycle makes behavior understandable.
type Job struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    JobStatus `json:"status"`
	Attempts  int       `json:"attempts"`
	LastError string    `json:"last_error,omitempty"`
}

// LESSON 2: In-memory job store
// Why this matters: single source of truth for status tracking.
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
	job := Job{
		ID:       s.nextID,
		Type:     t,
		Payload:  payload,
		Status:   StatusQueued,
		Attempts: 0,
	}
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

// LESSON 3: Queue abstraction
// Why this matters: separates enqueueing from processing details.
type JobQueue struct {
	ch chan Job
}

func NewJobQueue(buffer int) *JobQueue {
	return &JobQueue{ch: make(chan Job, buffer)}
}

func (q *JobQueue) Enqueue(job Job) {
	q.ch <- job
}

// LESSON 4: Worker handler function boundary
// Why this matters: work logic can evolve independently from queue mechanics.
type JobHandler func(job Job) error

// LESSON 5: Worker loop
// Why this matters: central place to enforce lifecycle transitions.
func StartWorker(store *JobStore, q *JobQueue, handler JobHandler) {
	go func() {
		for job := range q.ch {
			job.Status = StatusProcessing
			job.Attempts++
			store.Update(job)

			err := handler(job)
			if err != nil {
				job.Status = StatusFailed
				job.LastError = err.Error()
				store.Update(job)
				continue
			}
			job.Status = StatusDone
			job.LastError = ""
			store.Update(job)
		}
	}()
}

// LESSON 6: Simulated job handler
// Why this matters: deterministic learning behavior for status flow.
func simulatedHandler(job Job) error {
	time.Sleep(50 * time.Millisecond)
	if job.Type == "fail_once_demo" {
		return fmt.Errorf("simulated temporary failure")
	}
	return nil
}

type createJobRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func writeJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 7: Create endpoint returns quickly
// Why this matters: API responsiveness stays high under slow work.
func createJobHandler(store *JobStore, q *JobQueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
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
	}
}

// LESSON 8: List endpoint for status polling
// Why this matters: clients need visibility into async progress.
func listJobsHandler(store *JobStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, store.List())
	}
}

// LESSON 9: Route composition
// Why this matters: keep wiring explicit and easy to audit.
func buildMux(store *JobStore, q *JobQueue) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createJobHandler(store, q)(w, r)
			return
		}
		if r.Method == http.MethodGet {
			listJobsHandler(store)(w, r)
			return
		}
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	return mux
}

// LESSON 10: End-to-end async flow
// Why this matters: shows API + worker collaboration in one runnable example.
func main() {
	store := NewJobStore()
	queue := NewJobQueue(100)
	StartWorker(store, queue, simulatedHandler)

	mux := buildMux(store, queue)
	addr := ":8091"
	fmt.Println("Go async jobs lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Async Jobs 1-10
