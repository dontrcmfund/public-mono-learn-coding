package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
GO HTTP API BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/69-go-http-api-basics-1-10.go
2) Call endpoints:
   - GET  http://localhost:8080/health
   - GET  http://localhost:8080/tasks
   - POST http://localhost:8080/tasks  {"title":"Write Go API"}

Extra context:
- lessons/notes/156-go-api-principles.md
- lessons/notes/155-go-projects-principles.md
*/

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskRepo interface {
	List() []Task
	Add(title string) Task
}

type InMemoryTaskRepo struct {
	items []Task
}

func (r *InMemoryTaskRepo) List() []Task {
	out := make([]Task, len(r.items))
	copy(out, r.items)
	return out
}

func (r *InMemoryTaskRepo) Add(title string) Task {
	nextID := 1
	if len(r.items) > 0 {
		nextID = r.items[len(r.items)-1].ID + 1
	}
	t := Task{ID: nextID, Title: title, Done: false}
	r.items = append(r.items, t)
	return t
}

type TaskService struct {
	repo TaskRepo
}

func (s *TaskService) CreateTask(title string) (Task, error) {
	clean := strings.TrimSpace(title)
	if clean == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	return s.repo.Add(clean), nil
}

func (s *TaskService) ListTasks(minID int) []Task {
	all := s.repo.List()
	if minID <= 0 {
		return all
	}
	filtered := make([]Task, 0)
	for _, t := range all {
		if t.ID >= minID {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

type createTaskRequest struct {
	Title string `json:"title"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 1: Health handler
// Why this matters: fast readiness signal for operations.
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// LESSON 2-9: Task handlers with layering
// Why this matters: thin HTTP handlers + reusable service logic.
func buildMux(service *TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			minID := 0
			if raw := r.URL.Query().Get("min_id"); raw != "" {
				parsed, err := strconv.Atoi(raw)
				if err != nil {
					writeJSON(w, http.StatusBadRequest, map[string]string{"error": "min_id must be integer"})
					return
				}
				minID = parsed
			}
			writeJSON(w, http.StatusOK, service.ListTasks(minID))
			return

		case http.MethodPost:
			var payload createTaskRequest
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON payload"})
				return
			}

			task, err := service.CreateTask(payload.Title)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, http.StatusCreated, task)
			return

		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	return mux
}

// LESSON 10: Server bootstrap
// Why this matters: composition root wires service and transport.
func main() {
	repo := &InMemoryTaskRepo{}
	service := &TaskService{repo: repo}
	mux := buildMux(service)

	addr := ":8080"
	fmt.Println("Lesson 10: server listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go HTTP API Basics 1-10
