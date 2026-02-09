package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	_ "modernc.org/sqlite"
)

/*
GO SQLITE + HTTP (Lessons 11-20)

Suggested use:
1) Install driver (if needed): go get modernc.org/sqlite
2) Run: go run lessons/code/84-go-sqlite-http-11-20.go
3) Call:
   - POST /tasks {"title":"study sql"}
   - POST /tasks/1/done
   - GET  /tasks

Extra context:
- lessons/notes/171-go-database-sql-first-principles.md
- lessons/notes/172-go-sqlite-gotchas.md
*/

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskRepository interface {
	List() ([]Task, error)
	Add(title string) (Task, error)
	MarkDone(id int64) error
}

type SQLiteTaskRepo struct {
	db *sql.DB
}

func NewSQLiteTaskRepo(db *sql.DB) *SQLiteTaskRepo {
	return &SQLiteTaskRepo{db: db}
}

func (r *SQLiteTaskRepo) Migrate() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  done INTEGER NOT NULL DEFAULT 0
);`)
	return err
}

func (r *SQLiteTaskRepo) List() ([]Task, error) {
	rows, err := r.db.Query(`SELECT id, title, done FROM tasks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Task{}
	for rows.Next() {
		var t Task
		var doneInt int
		if err := rows.Scan(&t.ID, &t.Title, &doneInt); err != nil {
			return nil, err
		}
		t.Done = doneInt == 1
		items = append(items, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *SQLiteTaskRepo) Add(title string) (Task, error) {
	result, err := r.db.Exec(`INSERT INTO tasks (title, done) VALUES (?, 0)`, title)
	if err != nil {
		return Task{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Task{}, err
	}
	return Task{ID: id, Title: title, Done: false}, nil
}

func (r *SQLiteTaskRepo) MarkDone(id int64) error {
	result, err := r.db.Exec(`UPDATE tasks SET done = 1 WHERE id = ?`, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("task id %d not found", id)
	}
	return nil
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

func (s *TaskService) CompleteTask(id int64) error {
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

	// LESSON 11-15: list + create task endpoint
	// Why this matters: transport delegates to service, not SQL directly.
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

	// LESSON 16-19: completion endpoint
	// Why this matters: stable API contract for state transitions.
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
		var id int64
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
	// Why this matters: operational checks are part of real API design.
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	return mux
}

func main() {
	db, err := sql.Open("sqlite", "lessons/code/tmp_tasks_api_sqlite.db")
	if err != nil {
		fmt.Println("open db error:", err)
		return
	}
	defer db.Close()

	repo := NewSQLiteTaskRepo(db)
	if err := repo.Migrate(); err != nil {
		fmt.Println("migrate error:", err)
		return
	}

	service := NewTaskService(repo)
	mux := buildMux(service)

	addr := ":8085"
	fmt.Println("Go SQLite HTTP lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go SQLite + HTTP 11-20
