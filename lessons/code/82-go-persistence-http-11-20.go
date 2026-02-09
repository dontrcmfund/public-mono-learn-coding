package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

/*
GO PERSISTENCE + HTTP (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/82-go-persistence-http-11-20.go
2) Call:
   - POST /tasks {"title":"learn go persistence"}
   - POST /tasks/1/done
   - GET  /tasks

Extra context:
- lessons/notes/167-go-persistence-first-principles.md
- lessons/notes/168-go-repository-adapter-pattern.md
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
		}
	}
	if !found {
		return fmt.Errorf("task id %d not found", id)
	}
	return r.save(items)
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

type createTaskRequest struct {
	Title string `json:"title"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func buildMux(service *TaskService) *http.ServeMux {
	mux := http.NewServeMux()

	// LESSON 11-14: create and list handlers using service boundary
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			items, err := service.Tasks()
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "list failure"})
				return
			}
			writeJSON(w, http.StatusOK, items)
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

	// LESSON 15-19: completion endpoint with path parsing
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		path := strings.TrimPrefix(r.URL.Path, "/tasks/")
		if !strings.HasSuffix(path, "/done") {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		idPart := strings.TrimSuffix(path, "/done")
		var id int
		if _, err := fmt.Sscanf(idPart, "%d", &id); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be integer"})
			return
		}
		if err := service.CompleteTask(id); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "done"})
	})

	// LESSON 20: health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	return mux
}

func main() {
	repo := NewJSONFileTaskRepo("lessons/code/tmp_tasks_api.json")
	service := NewTaskService(repo)
	mux := buildMux(service)

	addr := ":8084"
	fmt.Println("Go persistence HTTP lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Persistence + HTTP 11-20
