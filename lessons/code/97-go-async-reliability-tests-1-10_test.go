package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
GO ASYNC RELIABILITY TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/97-go-async-reliability-tests-1-10_test.go -run TestLesson -v
2) Focus on retry behavior and DLQ movement
*/

type JobStatus string

const (
	StatusQueued     JobStatus = "queued"
	StatusProcessing JobStatus = "processing"
	StatusDone       JobStatus = "done"
	StatusFailed     JobStatus = "failed"
)

type Job struct {
	ID        int64
	Type      string
	Payload   string
	Status    JobStatus
	Attempts  int
	LastError string
}

type DLQEntry struct {
	JobID    int64
	JobType  string
	Payload  string
	Attempts int
	Reason   string
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

func (s *JobStore) Get(id int64) (Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	j, ok := s.items[id]
	return j, ok
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
				JobID:    job.ID,
				JobType:  job.Type,
				Payload:  job.Payload,
				Attempts: job.Attempts,
				Reason:   job.LastError,
			})
		}
	}()
}

func eventually(t *testing.T, timeout time.Duration, fn func() bool) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if fn() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("condition not met within %s", timeout)
}

func TestLesson1QueuedToDoneForSuccessfulJob(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond}
	StartRetryingWorker(store, dlq, queue, func(job Job) error { return nil }, policy)

	job := store.Add("ok", "payload")
	queue.Enqueue(job)

	eventually(t, 300*time.Millisecond, func() bool {
		j, _ := store.Get(job.ID)
		return j.Status == StatusDone && j.Attempts == 1
	})
}

func TestLesson2TemporaryFailureThenSuccess(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond}

	var mu sync.Mutex
	attemptMap := map[int64]int{}
	handler := func(job Job) error {
		mu.Lock()
		attemptMap[job.ID]++
		n := attemptMap[job.ID]
		mu.Unlock()
		if n < 2 {
			return fmt.Errorf("temporary failure")
		}
		return nil
	}

	StartRetryingWorker(store, dlq, queue, handler, policy)
	job := store.Add("unstable", "payload")
	queue.Enqueue(job)

	eventually(t, 500*time.Millisecond, func() bool {
		j, _ := store.Get(job.ID)
		return j.Status == StatusDone && j.Attempts == 2
	})
}

func TestLesson3PermanentFailureMovesToDLQ(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond}

	StartRetryingWorker(store, dlq, queue, func(job Job) error { return fmt.Errorf("permanent failure") }, policy)
	job := store.Add("always_fail", "payload")
	queue.Enqueue(job)

	eventually(t, 700*time.Millisecond, func() bool {
		j, _ := store.Get(job.ID)
		return j.Status == StatusFailed && j.Attempts == 3
	})

	entries := dlq.List()
	if len(entries) != 1 {
		t.Fatalf("expected 1 DLQ entry, got %d", len(entries))
	}
	if entries[0].JobID != job.ID {
		t.Fatalf("expected DLQ job id %d, got %d", job.ID, entries[0].JobID)
	}
}

func TestLesson4NoDLQForSuccessfulJob(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond}
	StartRetryingWorker(store, dlq, queue, func(job Job) error { return nil }, policy)

	job := store.Add("ok", "payload")
	queue.Enqueue(job)

	eventually(t, 300*time.Millisecond, func() bool {
		j, _ := store.Get(job.ID)
		return j.Status == StatusDone
	})
	if len(dlq.List()) != 0 {
		t.Fatalf("DLQ should be empty for successful job")
	}
}

func TestLesson5RetryPolicyDelayIncreases(t *testing.T) {
	p := RetryPolicy{BaseDelay: 10 * time.Millisecond}
	d1 := p.NextDelay(1)
	d2 := p.NextDelay(2)
	if d2 <= d1 {
		t.Fatalf("expected delay to increase, got %s then %s", d1, d2)
	}
}

func TestLesson6StoreGetUnknownJob(t *testing.T) {
	store := NewJobStore()
	_, ok := store.Get(999)
	if ok {
		t.Fatalf("expected unknown job to be absent")
	}
}

func TestLesson7DLQIncludesReason(t *testing.T) {
	dlq := NewDLQStore()
	dlq.Add(DLQEntry{JobID: 1, Reason: "timeout"})
	entries := dlq.List()
	if entries[0].Reason != "timeout" {
		t.Fatalf("expected reason timeout, got %q", entries[0].Reason)
	}
}

func TestLesson8DeterministicSuccessFlow(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 2, BaseDelay: 1 * time.Millisecond}
	StartRetryingWorker(store, dlq, queue, func(job Job) error { return nil }, policy)

	job1 := store.Add("ok", "a")
	queue.Enqueue(job1)
	job2 := store.Add("ok", "b")
	queue.Enqueue(job2)

	eventually(t, 500*time.Millisecond, func() bool {
		j1, _ := store.Get(job1.ID)
		j2, _ := store.Get(job2.ID)
		return j1.Status == StatusDone && j2.Status == StatusDone
	})
}

func TestLesson9AttemptsCappedAtMax(t *testing.T) {
	store := NewJobStore()
	dlq := NewDLQStore()
	queue := NewJobQueue(10)
	policy := RetryPolicy{MaxAttempts: 2, BaseDelay: 1 * time.Millisecond}
	StartRetryingWorker(store, dlq, queue, func(job Job) error { return fmt.Errorf("fail") }, policy)

	job := store.Add("always_fail", "payload")
	queue.Enqueue(job)

	eventually(t, 500*time.Millisecond, func() bool {
		j, _ := store.Get(job.ID)
		return j.Status == StatusFailed
	})
	j, _ := store.Get(job.ID)
	if j.Attempts != 2 {
		t.Fatalf("expected attempts 2, got %d", j.Attempts)
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Async Reliability Tests 1-10
