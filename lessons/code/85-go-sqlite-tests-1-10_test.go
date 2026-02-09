package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

/*
GO SQLITE TESTS (Lessons 1-10)

Suggested use:
1) Install driver (if needed): go get modernc.org/sqlite
2) Run: go test lessons/code/85-go-sqlite-tests-1-10_test.go -run TestLesson -v

Extra context:
- lessons/notes/172-go-sqlite-gotchas.md
*/

type Task struct {
	ID    int64
	Title string
	Done  bool
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
	repo *SQLiteTaskRepo
}

func NewTaskService(repo *SQLiteTaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (Task, error) {
	clean := strings.TrimSpace(title)
	if clean == "" {
		return Task{}, errors.New("title is required")
	}
	return s.repo.Add(clean)
}

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open test db error: %v", err)
	}
	return db
}

func TestLesson1MigrateWorks(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	if err := repo.Migrate(); err != nil {
		t.Fatalf("migrate failed: %v", err)
	}
}

func TestLesson2AddTask(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	task, err := repo.Add("task-a")
	if err != nil {
		t.Fatalf("add failed: %v", err)
	}
	if task.ID <= 0 {
		t.Fatalf("expected positive id")
	}
}

func TestLesson3ListTasks(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	_, _ = repo.Add("task-a")
	_, _ = repo.Add("task-b")
	items, err := repo.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d", len(items))
	}
}

func TestLesson4MarkDone(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	task, _ := repo.Add("task-a")
	if err := repo.MarkDone(task.ID); err != nil {
		t.Fatalf("mark done failed: %v", err)
	}
	items, _ := repo.List()
	if !items[0].Done {
		t.Fatalf("expected task done")
	}
}

func TestLesson5MarkDoneNotFound(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	err := repo.MarkDone(999)
	if err == nil {
		t.Fatalf("expected not found error")
	}
}

func TestLesson6ServiceValidation(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	service := NewTaskService(repo)
	_, err := service.CreateTask("   ")
	if err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestLesson7ServiceTrimsTitle(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	service := NewTaskService(repo)
	_, _ = service.CreateTask("  task-a ")
	items, _ := repo.List()
	if items[0].Title != "task-a" {
		t.Fatalf("expected trimmed title, got %q", items[0].Title)
	}
}

func TestLesson8DeterministicIDSequence(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	first, _ := repo.Add("a")
	second, _ := repo.Add("b")
	if second.ID != first.ID+1 {
		t.Fatalf("expected monotonic ids, got %d and %d", first.ID, second.ID)
	}
}

func TestLesson9ListEmptyOnFreshDB(t *testing.T) {
	db := openTestDB(t)
	defer db.Close()
	repo := NewSQLiteTaskRepo(db)
	_ = repo.Migrate()
	items, err := repo.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected empty list")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go SQLite Tests 1-10
