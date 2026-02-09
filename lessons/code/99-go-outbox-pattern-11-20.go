package main

import (
	"fmt"
	"sync"
	"time"
)

/*
GO OUTBOX PATTERN (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/99-go-outbox-pattern-11-20.go
2) Observe: write domain record + outbox, then publish pending events

Extra context:
- lessons/notes/186-event-contracts-and-versioning.md
- lessons/notes/187-outbox-pattern-first-principles.md
*/

type Task struct {
	ID    int64
	Title string
}

type OutboxEvent struct {
	ID         string
	EventType  string
	Version    int
	OccurredAt string
	Payload    map[string]any
	Delivered  bool
}

// LESSON 11: Repository that stores domain + outbox
// Why this matters: preserves event intent even when publish is unavailable.
type TaskRepository struct {
	mu        sync.Mutex
	nextID    int64
	tasks     []Task
	outbox    []OutboxEvent
	nextEvent int64
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		nextID:    1,
		nextEvent: 1,
		tasks:     []Task{},
		outbox:    []OutboxEvent{},
	}
}

// LESSON 12: Save task + outbox in one operation
// Why this matters: avoids "task saved but event lost" gap.
func (r *TaskRepository) CreateTaskWithOutbox(title string) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	task := Task{ID: r.nextID, Title: title}
	r.nextID++
	r.tasks = append(r.tasks, task)

	event := OutboxEvent{
		ID:         fmt.Sprintf("outbox-%d", r.nextEvent),
		EventType:  "task.created",
		Version:    1,
		OccurredAt: time.Now().Format(time.RFC3339),
		Payload: map[string]any{
			"task_id": task.ID,
			"title":   task.Title,
		},
		Delivered: false,
	}
	r.nextEvent++
	r.outbox = append(r.outbox, event)

	return task
}

// LESSON 13: Read pending outbox events
// Why this matters: publisher can retry safely from durable list.
func (r *TaskRepository) PendingOutbox() []OutboxEvent {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := []OutboxEvent{}
	for _, e := range r.outbox {
		if !e.Delivered {
			out = append(out, e)
		}
	}
	return out
}

// LESSON 14: Mark delivered after successful publish
// Why this matters: idempotent progress tracking.
func (r *TaskRepository) MarkDelivered(eventID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.outbox {
		if r.outbox[i].ID == eventID {
			r.outbox[i].Delivered = true
			return
		}
	}
}

// LESSON 15: Publisher abstraction
// Why this matters: outbox flow should not depend on one broker technology.
type EventPublisher interface {
	Publish(event OutboxEvent) error
}

// LESSON 16: In-memory publisher simulation
// Why this matters: deterministic learning for outbox process.
type MemoryPublisher struct {
	mu   sync.Mutex
	Sent []string
}

func (p *MemoryPublisher) Publish(event OutboxEvent) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Sent = append(p.Sent, event.ID)
	return nil
}

// LESSON 17: Outbox dispatcher
// Why this matters: isolated component for publish + mark-delivered logic.
type OutboxDispatcher struct {
	repo      *TaskRepository
	publisher EventPublisher
}

func NewOutboxDispatcher(repo *TaskRepository, publisher EventPublisher) *OutboxDispatcher {
	return &OutboxDispatcher{repo: repo, publisher: publisher}
}

func (d *OutboxDispatcher) FlushPending() int {
	pending := d.repo.PendingOutbox()
	sentCount := 0
	for _, e := range pending {
		if err := d.publisher.Publish(e); err != nil {
			continue
		}
		d.repo.MarkDelivered(e.ID)
		sentCount++
	}
	return sentCount
}

// LESSON 18: Service layer
// Why this matters: business action triggers domain write, not direct publish.
type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(repo *TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) Task {
	return s.repo.CreateTaskWithOutbox(title)
}

// LESSON 19: Composition root
// Why this matters: separate write path from publish path.

// LESSON 20: End-to-end outbox demo
// Why this matters: demonstrates reliable publish workflow.
func main() {
	repo := NewTaskRepository()
	service := NewTaskService(repo)
	publisher := &MemoryPublisher{}
	dispatcher := NewOutboxDispatcher(repo, publisher)

	task := service.CreateTask("learn outbox")
	fmt.Println("created task:", task)
	fmt.Println("pending before flush:", len(repo.PendingOutbox()))

	sent := dispatcher.FlushPending()
	fmt.Println("sent events:", sent)
	fmt.Println("pending after flush:", len(repo.PendingOutbox()))
}

// End of Go Outbox Pattern 11-20
