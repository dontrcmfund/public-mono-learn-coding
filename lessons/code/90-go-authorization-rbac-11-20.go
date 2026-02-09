package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
GO AUTHORIZATION RBAC (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/90-go-authorization-rbac-11-20.go
2) Call:
   - GET  /reports with `student-token` or `admin-token`
   - POST /admin/reports with `admin-token` only

Extra context:
- lessons/notes/176-authentication-vs-authorization-first-principles.md
- lessons/notes/178-go-auth-gotchas.md
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

// LESSON 11: Authorization middleware by role
// Why this matters: valid identity still needs explicit permission checks.
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

// LESSON 12: Authenticated endpoint for all valid users
// Why this matters: baseline protected read.
func reportsHandler(w http.ResponseWriter, r *http.Request) {
	userID, role, _ := userFromContext(r.Context())
	writeJSON(w, http.StatusOK, map[string]any{
		"user_id": userID,
		"role":    role,
		"reports": []string{"daily", "weekly"},
	})
}

// LESSON 13: Admin-only action
// Why this matters: demonstrates authorization boundary.
func adminCreateReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"status": "report created"})
}

// LESSON 14: Route wiring pattern
// Why this matters: explicit authn/authz composition is easier to audit.
func buildMux(store TokenStore) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/reports", authMiddleware(store, http.HandlerFunc(reportsHandler)))
	mux.Handle(
		"/admin/reports",
		authMiddleware(store, requireRole("admin", http.HandlerFunc(adminCreateReportHandler))),
	)
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	return mux
}

// LESSON 15: 401 vs 403 behavior
// Why this matters: correct semantics help clients and security tooling.

// LESSON 16: Deny by default principle
// Why this matters: safest default when permissions are unclear.

// LESSON 17: Role stays server-side
// Why this matters: never trust client-sent role claims directly.

// LESSON 18: Narrow permissions
// Why this matters: least privilege limits blast radius.

// LESSON 19: Endpoint-specific checks
// Why this matters: permissions belong close to protected action.

// LESSON 20: End-to-end RBAC demo server
// Why this matters: complete authn + authz flow in one service.
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
	mux := buildMux(store)

	addr := ":8088"
	fmt.Println("Go RBAC lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Authorization RBAC 11-20
