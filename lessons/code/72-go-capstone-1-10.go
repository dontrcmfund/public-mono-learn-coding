package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

/*
GO CAPSTONE (Lessons 1-10)

Project: Task Processing API with worker goroutines

Suggested use:
1) Run: go run lessons/code/72-go-capstone-1-10.go
2) Create task: curl -X POST localhost:8082/tasks -d '{"title":"demo"}'
3) List tasks:  curl localhost:8082/tasks

Extra context:
- lessons/notes/160-go-capstone-plan.md
- lessons/notes/158-go-api-architecture-principles.md
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
		return Task{}, fmt.Errorf("title cannot be empty")
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
			time.Sleep(40 * time.Millisecond)
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

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

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
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, http.StatusCreated, task)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	return mux
}

func main() {
	repo := &InMemoryTaskRepo{}
	service := NewTaskService(repo, 100)
	service.StartWorker()

	mux := buildMux(service)
	addr := ":8082"
	fmt.Println("Go capstone server listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Capstone 1-10
