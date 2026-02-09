package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
GO CONTRACT TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/103-go-contract-tests-1-10_test.go -run TestLesson -v
2) Focus on schema checks for both success and error responses
*/

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserRepo struct {
	nextID int64
	items  []User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{nextID: 1, items: []User{}}
}

func (r *UserRepo) List() []User {
	out := make([]User, len(r.items))
	copy(out, r.items)
	return out
}

func (r *UserRepo) Add(email string, name string) User {
	u := User{ID: r.nextID, Email: email, Name: name}
	r.nextID++
	r.items = append(r.items, u)
	return u
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeAPIError(w http.ResponseWriter, status int, code string, msg string) {
	writeJSON(w, status, APIError{Code: code, Message: msg})
}

func validateCreateUser(req CreateUserRequest) (string, string, bool) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	name := strings.TrimSpace(req.Name)
	if !strings.Contains(email, "@") || name == "" {
		return "", "", false
	}
	return email, name, true
}

func createUserHandler(repo *UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeAPIError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
			return
		}
		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeAPIError(w, http.StatusBadRequest, "invalid_json", "invalid JSON body")
			return
		}
		email, name, ok := validateCreateUser(req)
		if !ok {
			writeAPIError(w, http.StatusBadRequest, "validation_error", "email and name are required")
			return
		}
		writeJSON(w, http.StatusCreated, repo.Add(email, name))
	}
}

func listUsersHandler(repo *UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeAPIError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
			return
		}
		writeJSON(w, http.StatusOK, repo.List())
	}
}

func buildOpenAPIDoc() map[string]any {
	return map[string]any{
		"openapi": "3.0.3",
		"info":    map[string]any{"title": "User API", "version": "1.0.0"},
		"paths":   map[string]any{"/users": map[string]any{}},
		"components": map[string]any{
			"schemas": map[string]any{
				"User":     map[string]any{"type": "object"},
				"APIError": map[string]any{"type": "object"},
			},
		},
	}
}

func openAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, buildOpenAPIDoc())
}

func buildMux(repo *UserRepo) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createUserHandler(repo)(w, r)
			return
		}
		if r.Method == http.MethodGet {
			listUsersHandler(repo)(w, r)
			return
		}
		writeAPIError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
	})
	mux.HandleFunc("/openapi.json", openAPIHandler)
	return mux
}

func TestLesson1CreateUserReturns201(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"email":"mia@example.com","name":"Mia"}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d", w.Code)
	}
}

func TestLesson2CreateUserSchemaFields(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"email":"mia@example.com","name":"Mia"}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	var payload map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	if _, ok := payload["id"]; !ok {
		t.Fatalf("missing id")
	}
	if _, ok := payload["email"]; !ok {
		t.Fatalf("missing email")
	}
	if _, ok := payload["name"]; !ok {
		t.Fatalf("missing name")
	}
}

func TestLesson3CreateUserValidationErrorSchema(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"email":"bad","name":""}`))
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
	var payload map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	if _, ok := payload["code"]; !ok {
		t.Fatalf("missing error code")
	}
	if _, ok := payload["message"]; !ok {
		t.Fatalf("missing error message")
	}
}

func TestLesson4ListUsersReturnsArray(t *testing.T) {
	repo := NewUserRepo()
	_ = repo.Add("mia@example.com", "Mia")
	mux := buildMux(repo)
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
	var payload []map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected array response")
	}
}

func TestLesson5MethodNotAllowedUsesErrorSchema(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("want 405, got %d", w.Code)
	}
	var payload map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	if _, ok := payload["code"]; !ok {
		t.Fatalf("missing error code")
	}
	if _, ok := payload["message"]; !ok {
		t.Fatalf("missing error message")
	}
}

func TestLesson6OpenAPIReturns200(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson7OpenAPIHasCoreKeys(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	var payload map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	for _, key := range []string{"openapi", "info", "paths", "components"} {
		if _, ok := payload[key]; !ok {
			t.Fatalf("missing openapi key %s", key)
		}
	}
}

func TestLesson8ContentTypeJSON(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("want application/json, got %q", got)
	}
}

func TestLesson9DeterministicErrorShape(t *testing.T) {
	mux := buildMux(NewUserRepo())
	req1 := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"email":"bad","name":""}`))
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(`{"email":"bad","name":""}`))
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	var p1, p2 map[string]any
	_ = json.Unmarshal(w1.Body.Bytes(), &p1)
	_ = json.Unmarshal(w2.Body.Bytes(), &p2)
	if p1["code"] != p2["code"] {
		t.Fatalf("expected deterministic error code")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Contract Tests 1-10
