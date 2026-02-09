package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

/*
GO CAPSTONE TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/73-go-capstone-tests-1-10_test.go -run TestLesson -v
2) Why this command is file-specific: lesson files are standalone by design
*/

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskRepo interface {
	Add(title string) Task
	List() []Task
	UpdateStatus(id int, status string)
}

type InMemoryTaskRepo struct {
	mu    sync.Mutex
	items []Task
}

func (r *InMemoryTaskRepo) Add(title string) Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	nextID := 1
	if len(r.items) > 0 {
		nextID = r.items[len(r.items)-1].ID + 1
	}
	t := Task{ID: nextID, Title: title, Status: "queued", CreatedAt: time.Now()}
	r.items = append(r.items, t)
	return t
}

func (r *InMemoryTaskRepo) List() []Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]Task, len(r.items))
	copy(out, r.items)
	return out
}

func (r *InMemoryTaskRepo) UpdateStatus(id int, status string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.items {
		if r.items[i].ID == id {
			r.items[i].Status = status
			return
		}
	}
}

type TaskService struct {
	repo      TaskRepo
	workQueue chan Task
}

func NewTaskService(repo TaskRepo, queueSize int) *TaskService {
	return &TaskService{repo: repo, workQueue: make(chan Task, queueSize)}
}

func (s *TaskService) CreateTask(title string) (Task, error) {
	clean := strings.TrimSpace(title)
	if clean == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	t := s.repo.Add(clean)
	s.workQueue <- t
	return t, nil
}

func (s *TaskService) ListTasks() []Task {
	return s.repo.List()
}

func (s *TaskService) StartWorker() {
	go func() {
		for task := range s.workQueue {
			s.repo.UpdateStatus(task.ID, "processing")
			time.Sleep(10 * time.Millisecond)
			s.repo.UpdateStatus(task.ID, "done")
		}
	}()
}

type createTaskRequest struct {
	Title string `json:"title"`
}

func writeJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func buildMux(service *TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, service.ListTasks())
		case http.MethodPost:
			var req createTaskRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
				return
			}
			task, err := service.CreateTask(req.Title)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid title"})
				return
			}
			writeJSON(w, http.StatusCreated, task)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})
	return mux
}

func TestLesson1CreateTask(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	_, err := service.CreateTask("demo")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestLesson2CreateTaskEmpty(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	_, err := service.CreateTask("")
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestLesson3HTTPPostCreatesTask(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	mux := buildMux(service)

	body := bytes.NewBufferString(`{"title":"task-a"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d", w.Code)
	}
}

func TestLesson4HTTPPostInvalidJSON(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	mux := buildMux(service)

	body := bytes.NewBufferString(`{bad`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}

func TestLesson5HTTPGetList(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	_, _ = service.CreateTask("task-a")
	mux := buildMux(service)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson6WorkerMarksDone(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	service.StartWorker()
	_, _ = service.CreateTask("task-a")
	time.Sleep(25 * time.Millisecond)
	items := service.ListTasks()
	if len(items) != 1 || items[0].Status != "done" {
		t.Fatalf("expected status done, got %+v", items)
	}
}

func TestLesson7MethodNotAllowed(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	mux := buildMux(service)

	req := httptest.NewRequest(http.MethodDelete, "/tasks", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("want 405, got %d", w.Code)
	}
}

func TestLesson8ListCopySafety(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	_ = repo.Add("task-a")
	items := repo.List()
	items[0].Title = "mutated"
	fresh := repo.List()
	if fresh[0].Title == "mutated" {
		t.Fatalf("repo list should return copy")
	}
}

func TestLesson9JSONShape(t *testing.T) {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 10)
	mux := buildMux(service)
	body := bytes.NewBufferString(`{"title":"task-a"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	var payload map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &payload); err != nil {
		t.Fatalf("invalid json response")
	}
	if _, ok := payload["id"]; !ok {
		t.Fatalf("expected id field")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Capstone Tests 1-10
