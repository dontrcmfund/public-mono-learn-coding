package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/*
GO AUTH MIDDLEWARE (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/89-go-auth-middleware-1-10.go
2) Call:
   - GET /public
   - GET /profile with header: Authorization: Bearer student-token

Extra context:
- lessons/notes/176-authentication-vs-authorization-first-principles.md
- lessons/notes/177-go-api-token-security-basics.md
- lessons/notes/178-go-auth-gotchas.md
*/

type contextKey string

const (
	ctxUserIDKey contextKey = "user_id"
	ctxRoleKey   contextKey = "role"
)

// LESSON 1: Token store abstraction
// Why this matters: auth logic should depend on behavior, not storage source.
type TokenStore interface {
	Lookup(token string) (userID string, role string, ok bool)
}

// LESSON 2: In-memory token store adapter
// Why this matters: easy learning/test setup with deterministic behavior.
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

// LESSON 3: Authorization header parser
// Why this matters: strict parsing avoids accidental acceptance of bad formats.
func parseBearerToken(header string) (string, bool) {
	parts := strings.SplitN(strings.TrimSpace(header), " ", 2)
	if len(parts) != 2 {
		return "", false
	}
	if strings.ToLower(parts[0]) != "bearer" {
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

// LESSON 4: Authentication middleware
// Why this matters: one centralized check protects many endpoints.
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

// LESSON 5: Helpers for context extraction
// Why this matters: handlers stay clean and explicit.
func userFromContext(ctx context.Context) (string, string, bool) {
	userID, ok1 := ctx.Value(ctxUserIDKey).(string)
	role, ok2 := ctx.Value(ctxRoleKey).(string)
	if !ok1 || !ok2 {
		return "", "", false
	}
	return userID, role, true
}

// LESSON 6: Public endpoint
// Why this matters: not all endpoints require auth.
func publicHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "public endpoint"})
}

// LESSON 7: Authenticated profile endpoint
// Why this matters: shows identity extraction after middleware.
func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID, role, ok := userFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "missing auth context"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"user_id": userID,
		"role":    role,
	})
}

// LESSON 8: Build routes with selective protection
// Why this matters: explicit protection boundary prevents accidental exposure.
func buildMux(store TokenStore) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/public", publicHandler)
	mux.Handle("/profile", authMiddleware(store, http.HandlerFunc(profileHandler)))
	return mux
}

// LESSON 9: Composition root
// Why this matters: wire tokens/roles at startup edge, not deep in handlers.

// LESSON 10: End-to-end server
// Why this matters: practical authn flow from header -> middleware -> handler.
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

	addr := ":8087"
	fmt.Println("Go auth middleware lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Auth Middleware 1-10
