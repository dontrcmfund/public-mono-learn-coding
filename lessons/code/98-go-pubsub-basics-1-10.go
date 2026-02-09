package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

/*
GO PUB/SUB BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/98-go-pubsub-basics-1-10.go
2) Observe how one published event is handled by multiple subscribers

Extra context:
- lessons/notes/185-event-driven-first-principles.md
- lessons/notes/186-event-contracts-and-versioning.md
*/

// LESSON 1: Event contract type
// Why this matters: explicit contracts prevent ambiguity across producers/consumers.
type Event struct {
	ID         string         `json:"id"`
	EventType  string         `json:"event_type"`
	Version    int            `json:"version"`
	OccurredAt string         `json:"occurred_at"`
	Payload    map[string]any `json:"payload"`
}

// LESSON 2: Subscriber function boundary
// Why this matters: any consumer can implement this behavior contract.
type Subscriber func(Event) error

// LESSON 3: In-memory event bus
// Why this matters: teaches event flow without infrastructure setup.
type EventBus struct {
	mu          sync.Mutex
	subscribers map[string][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{subscribers: map[string][]Subscriber{}}
}

// LESSON 4: Subscribe by event type
// Why this matters: consumers opt into relevant business facts.
func (b *EventBus) Subscribe(eventType string, sub Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[eventType] = append(b.subscribers[eventType], sub)
}

// LESSON 5: Publish event to all subscribers
// Why this matters: producer remains decoupled from concrete consumers.
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

// LESSON 6: Producer example
// Why this matters: producers emit facts after domain state change.
func makeTaskCreatedEvent(taskID int64, title string) Event {
	return Event{
		ID:         fmt.Sprintf("evt-task-created-%d", taskID),
		EventType:  "task.created",
		Version:    1,
		OccurredAt: time.Now().Format(time.RFC3339),
		Payload: map[string]any{
			"task_id": taskID,
			"title":   title,
		},
	}
}

// LESSON 7: Subscriber examples
// Why this matters: one event can drive multiple side effects.
func auditSubscriber(event Event) error {
	data, _ := json.Marshal(event)
	fmt.Println("audit subscriber:", string(data))
	return nil
}

func notificationSubscriber(event Event) error {
	fmt.Printf("notification subscriber: send message for event %s\n", event.ID)
	return nil
}

// LESSON 8: Error-tolerant consumer demo
// Why this matters: one consumer failure should be visible.
func flakySubscriber(event Event) error {
	if event.Payload["title"] == "cause-error" {
		return fmt.Errorf("simulated subscriber failure")
	}
	return nil
}

// LESSON 9: Wiring subscribers
// Why this matters: composition root documents integration behavior.

// LESSON 10: End-to-end publish flow
// Why this matters: complete event-driven interaction in one file.
func main() {
	bus := NewEventBus()
	bus.Subscribe("task.created", auditSubscriber)
	bus.Subscribe("task.created", notificationSubscriber)
	bus.Subscribe("task.created", flakySubscriber)

	event := makeTaskCreatedEvent(101, "learn events")
	errs := bus.Publish(event)
	fmt.Println("Lesson 10 publish errors count:", len(errs))
}

// End of Go Pub/Sub Basics 1-10
