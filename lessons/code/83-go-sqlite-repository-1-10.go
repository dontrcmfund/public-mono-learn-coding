package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

/*
GO SQLITE REPOSITORY (Lessons 1-10)

Suggested use:
1) Install driver (if needed): go get modernc.org/sqlite
2) Run: go run lessons/code/83-go-sqlite-repository-1-10.go

Extra context:
- lessons/notes/170-what-is-sqlite-in-go.md
- lessons/notes/171-go-database-sql-first-principles.md
- lessons/notes/172-go-sqlite-gotchas.md
*/

// LESSON 1: Domain model
// Why this matters: domain and storage are related but not identical concepts.
type Task struct {
	ID    int64
	Title string
	Done  bool
}

// LESSON 2: Repository boundary
// Why this matters: service stays independent from SQL details.
type TaskRepository interface {
	List() ([]Task, error)
	Add(title string) (Task, error)
	MarkDone(id int64) error
}

// LESSON 3: SQL adapter struct
// Why this matters: isolate DB concerns in one component.
type SQLiteTaskRepo struct {
	db *sql.DB
}

func NewSQLiteTaskRepo(db *sql.DB) *SQLiteTaskRepo {
	return &SQLiteTaskRepo{db: db}
}

// LESSON 4: Migration step
// Why this matters: queries fail if schema does not exist.
func (r *SQLiteTaskRepo) Migrate() error {
	query := `
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  done INTEGER NOT NULL DEFAULT 0
);`
	_, err := r.db.Exec(query)
	return err
}

// LESSON 5: Parameterized insert
// Why this matters: prevents SQL injection and handles escaping safely.
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

// LESSON 6: Query rows safely
// Why this matters: structured scan keeps data mapping explicit.
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

// LESSON 7: Update with not-found check
// Why this matters: callers need clear signal when id is invalid.
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

// LESSON 8: Service layer with DI
// Why this matters: domain rules remain testable outside SQL.
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

// LESSON 9: Composition root
// Why this matters: wire DB and adapters at application edge.

// LESSON 10: End-to-end demo
// Why this matters: shows SQL adapter behind stable service interface.
func main() {
	db, err := sql.Open("sqlite", "lessons/code/tmp_tasks.db")
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
	_, _ = service.CreateTask("learn sql boundary")
	_, _ = service.CreateTask("write tests")
	_ = service.CompleteTask(1)

	items, err := service.Tasks()
	fmt.Println("Lesson 10 tasks:", items, "error:", err)
}

// End of Go SQLite Repository 1-10
