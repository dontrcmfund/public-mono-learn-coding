package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
GO PERSISTENCE + REPOSITORY (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/80-go-persistence-repo-1-10.go
2) Delete `lessons/code/tmp_tasks.json` and rerun to observe first-run behavior

Extra context:
- lessons/notes/167-go-persistence-first-principles.md
- lessons/notes/168-go-repository-adapter-pattern.md
- lessons/notes/169-go-file-storage-gotchas.md
*/

// LESSON 1: Domain model
// Why this matters: clear domain types make persistence format explicit.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// LESSON 2: Repository boundary
// Why this matters: service depends on behavior, not storage details.
type TaskRepository interface {
	List() ([]Task, error)
	Add(title string) (Task, error)
	MarkDone(id int) error
}

// LESSON 3: Service with DI
// Why this matters: business rules remain stable across adapters.
type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// LESSON 4: Service validation
// Why this matters: domain rules belong in service layer.
func (s *TaskService) CreateTask(title string) (Task, error) {
	clean := strings.TrimSpace(title)
	if clean == "" {
		return Task{}, errors.New("title is required")
	}
	return s.repo.Add(clean)
}

func (s *TaskService) CompleteTask(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	return s.repo.MarkDone(id)
}

func (s *TaskService) Tasks() ([]Task, error) {
	return s.repo.List()
}

// LESSON 5: In-memory adapter
// Why this matters: fastest adapter for learning and unit tests.
type InMemoryTaskRepo struct {
	items []Task
}

func (r *InMemoryTaskRepo) List() ([]Task, error) {
	out := make([]Task, len(r.items))
	copy(out, r.items)
	return out, nil
}

func (r *InMemoryTaskRepo) Add(title string) (Task, error) {
	nextID := 1
	if len(r.items) > 0 {
		nextID = r.items[len(r.items)-1].ID + 1
	}
	t := Task{ID: nextID, Title: title, Done: false}
	r.items = append(r.items, t)
	return t, nil
}

func (r *InMemoryTaskRepo) MarkDone(id int) error {
	for i := range r.items {
		if r.items[i].ID == id {
			r.items[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task id %d not found", id)
}

// LESSON 6: File adapter
// Why this matters: persistence survives process restart.
type JSONFileTaskRepo struct {
	path string
}

func NewJSONFileTaskRepo(path string) *JSONFileTaskRepo {
	return &JSONFileTaskRepo{path: path}
}

func (r *JSONFileTaskRepo) load() ([]Task, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var items []Task
	if len(data) == 0 {
		return []Task{}, nil
	}
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("invalid json file format: %w", err)
	}
	return items, nil
}

func (r *JSONFileTaskRepo) save(items []Task) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, data, 0o644)
}

func (r *JSONFileTaskRepo) List() ([]Task, error) {
	return r.load()
}

func (r *JSONFileTaskRepo) Add(title string) (Task, error) {
	items, err := r.load()
	if err != nil {
		return Task{}, err
	}
	nextID := 1
	if len(items) > 0 {
		nextID = items[len(items)-1].ID + 1
	}
	t := Task{ID: nextID, Title: title, Done: false}
	items = append(items, t)
	if err := r.save(items); err != nil {
		return Task{}, err
	}
	return t, nil
}

func (r *JSONFileTaskRepo) MarkDone(id int) error {
	items, err := r.load()
	if err != nil {
		return err
	}
	found := false
	for i := range items {
		if items[i].ID == id {
			items[i].Done = true
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task id %d not found", id)
	}
	return r.save(items)
}

// LESSON 7: One workflow, multiple adapters
// Why this matters: service code stays the same while storage changes.
func runWorkflow(label string, service *TaskService) {
	fmt.Println("==", label, "==")
	_, _ = service.CreateTask("write lesson")
	_, _ = service.CreateTask("review notes")
	_ = service.CompleteTask(1)
	items, err := service.Tasks()
	fmt.Println("tasks:", items, "error:", err)
}

// LESSON 8: Error visibility
// Why this matters: explicit errors make storage failures debuggable.

// LESSON 9: Composition root
// Why this matters: main wires adapters; service stays reusable.

// LESSON 10: End-to-end demo
// Why this matters: practical proof of boundary-driven design.
func main() {
	memService := NewTaskService(&InMemoryTaskRepo{})
	runWorkflow("in-memory adapter", memService)

	filePath := "lessons/code/tmp_tasks.json"
	fileService := NewTaskService(NewJSONFileTaskRepo(filePath))
	runWorkflow("json-file adapter", fileService)

	fmt.Println("Lesson 10 file written at:", filePath)
}

// End of Go Persistence + Repository 1-10
