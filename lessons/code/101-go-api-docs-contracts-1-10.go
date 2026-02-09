package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
GO API DOCS + CONTRACTS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/101-go-api-docs-contracts-1-10.go
2) Open:
   - GET /docs/contract
   - POST /users {"email":"mia@example.com","name":"Mia"}
   - GET /users

Extra context:
- lessons/notes/188-api-documentation-first-principles.md
- lessons/notes/189-openapi-first-principles.md
- lessons/notes/190-contract-testing-gotchas.md
*/

// LESSON 1: Response contract types
// Why this matters: explicit structs make response schema stable.
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

// LESSON 2: In-memory repository
// Why this matters: deterministic lesson behavior without setup friction.
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

// LESSON 3: Shared error writer
// Why this matters: one error schema across endpoints.
func writeAPIError(w http.ResponseWriter, status int, code string, msg string) {
	writeJSON(w, status, APIError{Code: code, Message: msg})
}

// LESSON 4: Contract document endpoint
// Why this matters: docs should be available from the service itself.
func contractDocHandler(w http.ResponseWriter, _ *http.Request) {
	doc := map[string]any{
		"title":   "User API Contract",
		"version": "1.0.0",
		"endpoints": []map[string]any{
			{
				"method":      "GET",
				"path":        "/users",
				"description": "List users",
				"responses": map[string]any{
					"200": "[]User",
				},
			},
			{
				"method":      "POST",
				"path":        "/users",
				"description": "Create user",
				"request":     "CreateUserRequest",
				"responses": map[string]any{
					"201": "User",
					"400": "APIError",
				},
			},
		},
	}
	writeJSON(w, http.StatusOK, doc)
}

// LESSON 5: Validation helper
// Why this matters: contract errors should be deterministic.
func validateCreateUser(req CreateUserRequest) (string, string, bool) {
	email := strings.TrimSpace(strings.ToLower(req.Email))
	name := strings.TrimSpace(req.Name)
	if !strings.Contains(email, "@") {
		return "", "", false
	}
	if name == "" {
		return "", "", false
	}
	return email, name, true
}

// LESSON 6: Create endpoint with documented schema
// Why this matters: request and response shapes stay predictable.
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
		created := repo.Add(email, name)
		writeJSON(w, http.StatusCreated, created)
	}
}

// LESSON 7: List endpoint with stable array schema
// Why this matters: clients rely on consistent list responses.
func listUsersHandler(repo *UserRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeAPIError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
			return
		}
		writeJSON(w, http.StatusOK, repo.List())
	}
}

// LESSON 8: Router wiring
// Why this matters: endpoint map is part of API contract clarity.
func buildMux(repo *UserRepo) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/docs/contract", contractDocHandler)
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
	return mux
}

// LESSON 9: Startup logging for contract version
// Why this matters: operators and clients need version visibility.

// LESSON 10: End-to-end API contract demo
// Why this matters: docs and handlers evolve together in one service.
func main() {
	repo := NewUserRepo()
	mux := buildMux(repo)
	addr := ":8093"
	fmt.Println("User API contract version 1.0.0")
	fmt.Println("Go API docs + contracts lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go API Docs + Contracts 1-10
