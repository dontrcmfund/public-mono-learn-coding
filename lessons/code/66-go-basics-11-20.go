package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
GO BASICS (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/66-go-basics-11-20.go
2) Modify one lesson block at a time and rerun

Extra context:
- lessons/notes/151-go-first-principles.md
- lessons/notes/152-go-gotchas.md
*/

// LESSON 11: Multiple return values
// Why this matters: Go returns value + error explicitly.
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// LESSON 12: Error handling pattern
// Why this matters: explicit error checks are a core Go habit.

// LESSON 13: Structs
// Why this matters: structs model domain data.
type Task struct {
	ID    int
	Title string
	Done  bool
}

// LESSON 14: Methods on structs
// Why this matters: behavior can live with domain data.
func (t Task) MarkDone() Task {
	t.Done = true
	return t
}

// LESSON 15: Pointers
// Why this matters: mutate shared state intentionally.
func renameTask(t *Task, newTitle string) {
	t.Title = strings.TrimSpace(newTitle)
}

// LESSON 16: Interfaces
// Why this matters: use abstractions for replaceable behavior.
type TaskRepository interface {
	Add(task Task)
	List() []Task
}

// LESSON 17: In-memory adapter
// Why this matters: simple adapter supports local testing.
type InMemoryRepo struct {
	items []Task
}

func (r *InMemoryRepo) Add(task Task) {
	r.items = append(r.items, task)
}

func (r *InMemoryRepo) List() []Task {
	copyItems := make([]Task, len(r.items))
	copy(copyItems, r.items)
	return copyItems
}

// LESSON 18: Slice iteration with range
// Why this matters: idiomatic sequence processing.

// LESSON 19: String utilities
// Why this matters: input normalization prevents messy data.

// LESSON 20: Composition mindset
// Why this matters: combine small functions into workflow.
func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Lesson 11/12 error:", err)
	} else {
		fmt.Println("Lesson 11/12:", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Lesson 12 expected error:", err)
	}

	task := Task{ID: 1, Title: "Write Go lesson", Done: false}
	fmt.Println("Lesson 13:", task)

	taskDone := task.MarkDone()
	fmt.Println("Lesson 14:", taskDone)

	renameTask(&task, "  Review Go architecture  ")
	fmt.Println("Lesson 15:", task.Title)

	repo := &InMemoryRepo{}
	repo.Add(task)
	repo.Add(Task{ID: 2, Title: "Test API", Done: false})

	listed := repo.List()
	fmt.Println("Lesson 16/17:", listed)

	for _, t := range listed {
		fmt.Println("Lesson 18 task:", t.ID, t.Title, t.Done)
	}

	normalized := strings.ToLower(strings.TrimSpace("  GO Lessons  "))
	fmt.Println("Lesson 19:", normalized)

	workflow := fmt.Sprintf("Task #%d ready: %s", task.ID, task.Title)
	fmt.Println("Lesson 20:", workflow)
}

// End of Go Basics 11-20
