package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
GO AUTH TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/91-go-auth-tests-1-10_test.go -run TestLesson -v
2) Focus on 401 vs 403 behavior, not only happy paths
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
	found, ok := s.tokens[token]
	if !ok {
		return "", "", false
	}
	return found.UserID, found.Role, true
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

func userFromContext(ctx context.Context) (string, string, bool) {
	userID, ok1 := ctx.Value(ctxUserIDKey).(string)
	role, ok2 := ctx.Value(ctxRoleKey).(string)
	if !ok1 || !ok2 {
		return "", "", false
	}
	return userID, role, true
}

func requireRole(required string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, role, ok := userFromContext(r.Context())
		if !ok {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "missing auth context"})
			return
		}
		if role != required {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID, role, _ := userFromContext(r.Context())
	writeJSON(w, http.StatusOK, map[string]string{"user_id": userID, "role": role})
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"status": "ok"})
}

func buildMux(store TokenStore) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/profile", authMiddleware(store, http.HandlerFunc(profileHandler)))
	mux.Handle("/admin/action", authMiddleware(store, requireRole("admin", http.HandlerFunc(adminHandler))))
	return mux
}

func newStore() *InMemoryTokenStore {
	return &InMemoryTokenStore{
		tokens: map[string]struct {
			UserID string
			Role   string
		}{
			"student-token": {UserID: "student-1", Role: "student"},
			"admin-token":   {UserID: "admin-1", Role: "admin"},
		},
	}
}

func TestLesson1PublicProtectedWithoutTokenIs401(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", w.Code)
	}
}

func TestLesson2InvalidTokenIs401(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer bad-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", w.Code)
	}
}

func TestLesson3ValidTokenGetsProfile(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer student-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson4ProfileResponseContainsRole(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer student-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	var payload map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	if payload["role"] != "student" {
		t.Fatalf("want role student, got %q", payload["role"])
	}
}

func TestLesson5StudentCannotAccessAdminEndpoint(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodPost, "/admin/action", nil)
	req.Header.Set("Authorization", "Bearer student-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusForbidden {
		t.Fatalf("want 403, got %d", w.Code)
	}
}

func TestLesson6AdminCanAccessAdminEndpoint(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodPost, "/admin/action", nil)
	req.Header.Set("Authorization", "Bearer admin-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d", w.Code)
	}
}

func TestLesson7AdminWrongMethodGets405(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/admin/action", nil)
	req.Header.Set("Authorization", "Bearer admin-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("want 405, got %d", w.Code)
	}
}

func TestLesson8MalformedAuthorizationHeaderIs401(t *testing.T) {
	mux := buildMux(newStore())
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Token student-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", w.Code)
	}
}

func TestLesson9DeterministicProfileResponse(t *testing.T) {
	mux := buildMux(newStore())
	req1 := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req1.Header.Set("Authorization", "Bearer student-token")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req2.Header.Set("Authorization", "Bearer student-token")
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w1.Body.String() != w2.Body.String() {
		t.Fatalf("profile response should be deterministic")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Auth Tests 1-10
