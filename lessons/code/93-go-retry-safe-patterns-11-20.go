package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

/*
GO RETRY-SAFE PATTERNS (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/93-go-retry-safe-patterns-11-20.go
2) Call POST /payments twice with same `Idempotency-Key`:
   - same body -> same response
   - different body -> 409 conflict

Extra context:
- lessons/notes/180-idempotency-first-principles.md
- lessons/notes/181-go-retry-safety-gotchas.md
*/

type contextKey string

const ctxUserIDKey contextKey = "user_id"

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

// LESSON 11: Auth middleware
// Why this matters: idempotency keys should be scoped to authenticated identity.
func authMiddleware(store TokenStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := parseBearerToken(r.Header.Get("Authorization"))
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing or invalid bearer token"})
			return
		}
		userID, _, ok := store.Lookup(token)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userIDFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(ctxUserIDKey).(string)
	return v, ok
}

type idempotencyRecord struct {
	Fingerprint string
	Status      int
	Body        map[string]any
}

// LESSON 12: Idempotency store with payload fingerprint
// Why this matters: same key + different payload must be rejected.
type IdempotencyStore struct {
	mu      sync.Mutex
	records map[string]idempotencyRecord
}

func NewIdempotencyStore() *IdempotencyStore {
	return &IdempotencyStore{records: map[string]idempotencyRecord{}}
}

func (s *IdempotencyStore) Get(key string) (idempotencyRecord, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.records[key]
	return v, ok
}

func (s *IdempotencyStore) Set(key string, rec idempotencyRecord) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[key] = rec
}

// LESSON 13: Canonical payload fingerprint
// Why this matters: semantically equal payloads map to same dedupe identity.
func fingerprintJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

type createPaymentRequest struct {
	AmountCents int64  `json:"amount_cents"`
	Currency    string `json:"currency"`
	Note        string `json:"note"`
}

// LESSON 14: Payment service
// Why this matters: side-effect simulation for retry-safe mutation endpoint.
type PaymentService struct {
	mu      sync.Mutex
	nextID  int64
	records []map[string]any
}

func NewPaymentService() *PaymentService {
	return &PaymentService{nextID: 1}
}

func (s *PaymentService) Create(req createPaymentRequest) map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	record := map[string]any{
		"id":           id,
		"amount_cents": req.AmountCents,
		"currency":     strings.ToUpper(strings.TrimSpace(req.Currency)),
		"note":         strings.TrimSpace(req.Note),
		"status":       "accepted",
	}
	s.records = append(s.records, record)
	return record
}

// LESSON 15: Validation helper
// Why this matters: reject invalid requests before mutation.
func validatePayment(req createPaymentRequest) error {
	if req.AmountCents <= 0 {
		return fmt.Errorf("amount_cents must be positive")
	}
	cur := strings.TrimSpace(strings.ToUpper(req.Currency))
	if len(cur) != 3 {
		return fmt.Errorf("currency must be 3-letter code")
	}
	return nil
}

// LESSON 16-19: Retry-safe create handler
// Why this matters: key-based dedupe + payload conflict protection.
func createPaymentHandler(service *PaymentService, store *IdempotencyStore) http.HandlerFunc {
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
		userID, ok := userIDFromContext(r.Context())
		if !ok {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "missing auth context"})
			return
		}

		var req createPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		if err := validatePayment(req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		fp, err := fingerprintJSON(req)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "fingerprint failed"})
			return
		}
		scopedKey := userID + ":" + key

		if existing, ok := store.Get(scopedKey); ok {
			if existing.Fingerprint != fp {
				writeJSON(w, http.StatusConflict, map[string]string{"error": "idempotency key reused with different payload"})
				return
			}
			writeJSON(w, existing.Status, existing.Body)
			return
		}

		payload := service.Create(req)
		store.Set(scopedKey, idempotencyRecord{
			Fingerprint: fp,
			Status:      http.StatusCreated,
			Body:        payload,
		})
		writeJSON(w, http.StatusCreated, payload)
	}
}

func buildMux(store TokenStore, ids *IdempotencyStore, service *PaymentService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/payments", authMiddleware(store, http.HandlerFunc(createPaymentHandler(service, ids))))
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	return mux
}

// LESSON 20: End-to-end retry-safe payment API
// Why this matters: practical pattern for high-value create endpoints.
func main() {
	store := &InMemoryTokenStore{
		tokens: map[string]struct {
			UserID string
			Role   string
		}{
			"student-token": {UserID: "student-1", Role: "student"},
		},
	}
	ids := NewIdempotencyStore()
	service := NewPaymentService()
	mux := buildMux(store, ids, service)

	addr := ":8090"
	fmt.Println("Go retry-safe patterns lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go Retry-Safe Patterns 11-20
