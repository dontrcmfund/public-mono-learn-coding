package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
GO EVENT RELIABILITY TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/100-go-event-reliability-tests-1-10_test.go -run TestLesson -v
2) Focus on outbox guarantees and subscriber behavior
*/

type Event struct {
	ID         string
	EventType  string
	Version    int
	OccurredAt string
	Payload    map[string]any
}

type Subscriber func(Event) error

type EventBus struct {
	mu          sync.Mutex
	subscribers map[string][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{subscribers: map[string][]Subscriber{}}
}

func (b *EventBus) Subscribe(eventType string, sub Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[eventType] = append(b.subscribers[eventType], sub)
}

func (b *EventBus) Publish(event Event) []error {
	b.mu.Lock()
	subs := append([]Subscriber{}, b.subscribers[event.EventType]...)
	b.mu.Unlock()
	errs := []error{}
	for _, sub := range subs {
		if err := sub(event); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

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

type TaskRepository struct {
	mu        sync.Mutex
	nextID    int64
	nextEvent int64
	tasks     []Task
	outbox    []OutboxEvent
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		nextID:    1,
		nextEvent: 1,
		tasks:     []Task{},
		outbox:    []OutboxEvent{},
	}
}

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
		Payload:    map[string]any{"task_id": task.ID, "title": task.Title},
		Delivered:  false,
	}
	r.nextEvent++
	r.outbox = append(r.outbox, event)
	return task
}

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

type EventPublisher interface {
	Publish(event OutboxEvent) error
}

type MemoryPublisher struct {
	mu        sync.Mutex
	Sent      []string
	FailFirst bool
	calls     int
}

func (p *MemoryPublisher) Publish(event OutboxEvent) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.calls++
	if p.FailFirst && p.calls == 1 {
		return fmt.Errorf("temporary publish failure")
	}
	p.Sent = append(p.Sent, event.ID)
	return nil
}

type OutboxDispatcher struct {
	repo      *TaskRepository
	publisher EventPublisher
}

func NewOutboxDispatcher(repo *TaskRepository, publisher EventPublisher) *OutboxDispatcher {
	return &OutboxDispatcher{repo: repo, publisher: publisher}
}

func (d *OutboxDispatcher) FlushPending() int {
	pending := d.repo.PendingOutbox()
	sent := 0
	for _, e := range pending {
		if err := d.publisher.Publish(e); err != nil {
			continue
		}
		d.repo.MarkDelivered(e.ID)
		sent++
	}
	return sent
}

func TestLesson1PublishCallsAllSubscribers(t *testing.T) {
	bus := NewEventBus()
	count := 0
	bus.Subscribe("task.created", func(e Event) error { count++; return nil })
	bus.Subscribe("task.created", func(e Event) error { count++; return nil })
	bus.Publish(Event{EventType: "task.created"})
	if count != 2 {
		t.Fatalf("want 2 subscriber calls, got %d", count)
	}
}

func TestLesson2PublishCollectsErrors(t *testing.T) {
	bus := NewEventBus()
	bus.Subscribe("task.created", func(e Event) error { return fmt.Errorf("fail-a") })
	bus.Subscribe("task.created", func(e Event) error { return nil })
	errs := bus.Publish(Event{EventType: "task.created"})
	if len(errs) != 1 {
		t.Fatalf("want 1 error, got %d", len(errs))
	}
}

func TestLesson3OutboxCreatedWithTask(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("task-a")
	pending := repo.PendingOutbox()
	if len(pending) != 1 {
		t.Fatalf("want 1 outbox event, got %d", len(pending))
	}
}

func TestLesson4FlushMarksDelivered(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("task-a")
	dispatcher := NewOutboxDispatcher(repo, &MemoryPublisher{})
	sent := dispatcher.FlushPending()
	if sent != 1 {
		t.Fatalf("want sent=1, got %d", sent)
	}
	if len(repo.PendingOutbox()) != 0 {
		t.Fatalf("pending should be 0 after successful flush")
	}
}

func TestLesson5FailedPublishRemainsPending(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("task-a")
	dispatcher := NewOutboxDispatcher(repo, &MemoryPublisher{FailFirst: true})
	_ = dispatcher.FlushPending()
	if len(repo.PendingOutbox()) != 1 {
		t.Fatalf("event should remain pending on publish failure")
	}
}

func TestLesson6SecondFlushSucceedsAfterTransientFailure(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("task-a")
	pub := &MemoryPublisher{FailFirst: true}
	dispatcher := NewOutboxDispatcher(repo, pub)
	_ = dispatcher.FlushPending()
	sent := dispatcher.FlushPending()
	if sent != 1 {
		t.Fatalf("want second flush to send 1, got %d", sent)
	}
}

func TestLesson7OutboxIDsIncrement(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("a")
	_ = repo.CreateTaskWithOutbox("b")
	pending := repo.PendingOutbox()
	if pending[0].ID == pending[1].ID {
		t.Fatalf("outbox IDs should be unique")
	}
}

func TestLesson8NoDuplicateSendAfterDelivered(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("a")
	pub := &MemoryPublisher{}
	dispatcher := NewOutboxDispatcher(repo, pub)
	_ = dispatcher.FlushPending()
	second := dispatcher.FlushPending()
	if second != 0 {
		t.Fatalf("expected no additional sends, got %d", second)
	}
}

func TestLesson9DeterministicPendingCount(t *testing.T) {
	repo := NewTaskRepository()
	_ = repo.CreateTaskWithOutbox("a")
	_ = repo.CreateTaskWithOutbox("b")
	if len(repo.PendingOutbox()) != 2 {
		t.Fatalf("want pending=2")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Event Reliability Tests 1-10
