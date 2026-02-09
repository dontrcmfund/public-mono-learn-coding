package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

/*
GO RATE LIMIT + IDEMPOTENCY (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/92-go-rate-limit-idempotency-1-10.go
2) Call quickly several times:
   - POST /tasks with headers:
     Authorization: Bearer student-token
     Idempotency-Key: create-task-001
   - body: {"title":"learn retry safety"}

Extra context:
- lessons/notes/179-rate-limiting-first-principles.md
- lessons/notes/180-idempotency-first-principles.md
- lessons/notes/181-go-retry-safety-gotchas.md
*/

type contextKey string

const (
	ctxUserIDKey contextKey = "user_id"
	ctxRoleKey   contextKey = "role"
)

type TokenStore interface {
	Lookup(token string) (userID string, role string, ok bool)
}

type InMemoryTokenStore struct {
	tokens map[string]struct {
		UserID string
		Role   string
	}
}

func (s *InMemoryTokenStore) Lookup(token string) (string, string, bool) {
	t, ok := s.tokens[token]
	if !ok {
		return "", "", false
	}
	return t.UserID, t.Role, true
}

func parseBearerToken(header string) (string, bool) {
	parts := strings.SplitN(strings.TrimSpace(header), " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", false
	}
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", false
	}
	return token, true
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 1: Auth middleware
// Why this matters: limits and idempotency should apply per authenticated identity.
func authMiddleware(store TokenStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := parseBearerToken(r.Header.Get("Authorization"))
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing or invalid bearer token"})
			return
		}
		userID, role, ok := store.Lookup(token)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserIDKey, userID)
		ctx = context.WithValue(ctx, ctxRoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userIDFromRequest(r *http.Request) (string, bool) {
	v := r.Context().Value(ctxUserIDKey)
	s, ok := v.(string)
	return s, ok
}

// LESSON 2: Fixed-window per-user rate limiter
// Why this matters: simple fairness guard for shared API capacity.
type RateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	clients map[string]*clientWindow
}

type clientWindow struct {
	start time.Time
	count int
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:   limit,
		window:  window,
		clients: map[string]*clientWindow{},
	}
}

func (rl *RateLimiter) Allow(clientID string, now time.Time) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cw, ok := rl.clients[clientID]
	if !ok || now.Sub(cw.start) >= rl.window {
		rl.clients[clientID] = &clientWindow{start: now, count: 1}
		return true
	}
	if cw.count >= rl.limit {
		return false
	}
	cw.count++
	return true
}

// LESSON 3: Rate limit middleware
// Why this matters: centralized protection avoids per-handler duplication.
func rateLimitMiddleware(rl *RateLimiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := userIDFromRequest(r)
		if !ok {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "missing auth context"})
			return
		}
		if !rl.Allow(userID, time.Now()) {
			w.Header().Set("Retry-After", "2")
			writeJSON(w, http.StatusTooManyRequests, map[string]string{"error": "rate limit exceeded"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LESSON 4: Idempotency response record
// Why this matters: repeat requests should return consistent result.
type cachedResponse struct {
	Status int
	Body   map[string]any
}

// LESSON 5: In-memory idempotency store
// Why this matters: simple model for deduplication behavior.
type IdempotencyStore struct {
	mu      sync.Mutex
	records map[string]cachedResponse
}

func NewIdempotencyStore() *IdempotencyStore {
	return &IdempotencyStore{records: map[string]cachedResponse{}}
}

func (s *IdempotencyStore) Get(key string) (cachedResponse, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.records[key]
	return v, ok
}

func (s *IdempotencyStore) Set(key string, val cachedResponse) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[key] = val
}

// LESSON 6: Domain service (task create)
// Why this matters: idempotency protects mutation endpoints.
type TaskService struct {
	mu     sync.Mutex
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{nextID: 1}
}

func (s *TaskService) Create(title string) map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	return map[string]any{
		"id":    id,
		"title": title,
		"done":  false,
	}
}

type createTaskRequest struct {
	Title string `json:"title"`
}

// LESSON 7: Protected create endpoint with idempotency
// Why this matters: retrying same request key avoids duplicate side effects.
func createTaskHandler(service *TaskService, ids *IdempotencyStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		key := strings.TrimSpace(r.Header.Get("Idempotency-Key"))
		if key == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing Idempotency-Key"})
			return
		}

		userID, ok := userIDFromRequest(r)
		if !ok {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "missing auth context"})
			return
		}
		scopedKey := userID + ":" + key

		if cached, ok := ids.Get(scopedKey); ok {
			writeJSON(w, cached.Status, cached.Body)
			return
		}

		var req createTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		title := strings.TrimSpace(req.Title)
		if title == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title is required"})
			return
		}

		payload := service.Create(title)
		ids.Set(scopedKey, cachedResponse{Status: http.StatusCreated, Body: payload})
		writeJSON(w, http.StatusCreated, payload)
	}
}

// LESSON 8: Public health endpoint
// Why this matters: keep operational endpoint simple and always available.
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// LESSON 9: Middleware composition
// Why this matters: explicit order keeps behavior predictable.
func buildMux(store TokenStore, rl *RateLimiter, ids *IdempotencyStore, service *TaskService) *http.ServeMux {
	mux := http.NewServeMux()
	protectedCreate := authMiddleware(store, rateLimitMiddleware(rl, http.HandlerFunc(createTaskHandler(service, ids))))
	mux.Handle("/tasks", protectedCreate)
	mux.HandleFunc("/health", healthHandler)
	return mux
}

// LESSON 10: End-to-end retry-safe API
// Why this matters: combines auth + fairness + deduplication in one flow.
func main() {
	store := &InMemoryTokenStore{
		tokens: map[string]struct {
			UserID string
			Role   string
		}{
			"student-token": {UserID: "student-1", Role: "student"},
			"admin-token":   {UserID: "admin-1", Role: "admin"},
		},
	}
	rl := NewRateLimiter(3, 2*time.Second)
	ids := NewIdempotencyStore()
	service := NewTaskService()
	mux := buildMux(store, rl, ids, service)

	addr := ":8089"
	fmt.Println("Go rate limit + idempotency lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Rate Limit + Idempotency 1-10
