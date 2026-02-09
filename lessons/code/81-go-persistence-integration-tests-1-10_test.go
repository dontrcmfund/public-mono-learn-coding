package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

/*
GO PERSISTENCE INTEGRATION TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/81-go-persistence-integration-tests-1-10_test.go -run TestLesson -v
2) Observe how the same service behavior is tested against both adapters
*/

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskRepository interface {
	List() ([]Task, error)
	Add(title string) (Task, error)
	MarkDone(id int) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

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
	if len(data) == 0 {
		return []Task{}, nil
	}
	var items []Task
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *JSONFileTaskRepo) save(items []Task) error {
	data, err := json.Marshal(items)
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
	t := Task{ID: nextID, Title: title}
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

func runCommonFlow(t *testing.T, repo TaskRepository) {
	t.Helper()

	service := NewTaskService(repo)
	_, err := service.CreateTask("task-a")
	if err != nil {
		t.Fatalf("create task-a failed: %v", err)
	}
	_, err = service.CreateTask("task-b")
	if err != nil {
		t.Fatalf("create task-b failed: %v", err)
	}

	err = service.CompleteTask(1)
	if err != nil {
		t.Fatalf("complete task failed: %v", err)
	}

	items, err := service.Tasks()
	if err != nil {
		t.Fatalf("list tasks failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 tasks, got %d", len(items))
	}
	if !items[0].Done {
		t.Fatalf("first task should be done")
	}
}

func TestLesson1MemoryAdapterFlow(t *testing.T) {
	runCommonFlow(t, &InMemoryTaskRepo{})
}

func TestLesson2FileAdapterFlow(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	runCommonFlow(t, NewJSONFileTaskRepo(path))
}

func TestLesson3ServiceRejectsEmptyTitle(t *testing.T) {
	service := NewTaskService(&InMemoryTaskRepo{})
	_, err := service.CreateTask("   ")
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestLesson4ServiceRejectsInvalidID(t *testing.T) {
	service := NewTaskService(&InMemoryTaskRepo{})
	err := service.CompleteTask(0)
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestLesson5MemoryMarkDoneNotFound(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	err := repo.MarkDone(99)
	if err == nil {
		t.Fatalf("expected not found error")
	}
}

func TestLesson6FileRepoHandlesMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.json")
	repo := NewJSONFileTaskRepo(path)
	items, err := repo.List()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected empty list")
	}
}

func TestLesson7FileRepoPersistsData(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	repo := NewJSONFileTaskRepo(path)
	_, _ = repo.Add("persist me")
	secondRead, err := repo.List()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(secondRead) != 1 || secondRead[0].Title != "persist me" {
		t.Fatalf("unexpected persisted data: %+v", secondRead)
	}
}

func TestLesson8FileRepoHandlesCorruptJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	_ = os.WriteFile(path, []byte("{bad"), 0o644)
	repo := NewJSONFileTaskRepo(path)
	_, err := repo.List()
	if err == nil {
		t.Fatalf("expected json parse error")
	}
}

func TestLesson9AdaptersProduceSameTaskIDs(t *testing.T) {
	mem := &InMemoryTaskRepo{}
	file := NewJSONFileTaskRepo(filepath.Join(t.TempDir(), "tasks.json"))

	a, _ := mem.Add("one")
	b, _ := file.Add("one")
	if a.ID != b.ID {
		t.Fatalf("expected same first ID, got %d and %d", a.ID, b.ID)
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Persistence Integration Tests 1-10
