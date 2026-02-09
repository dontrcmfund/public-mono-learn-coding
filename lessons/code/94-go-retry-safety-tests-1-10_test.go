package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

/*
GO RETRY SAFETY TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/94-go-retry-safety-tests-1-10_test.go -run TestLesson -v
2) Focus on duplicate-request behavior and conflict handling
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

type PaymentService struct {
	mu     sync.Mutex
	nextID int64
}

func NewPaymentService() *PaymentService {
	return &PaymentService{nextID: 1}
}

func (s *PaymentService) Create(req createPaymentRequest) map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	return map[string]any{
		"id":           id,
		"amount_cents": req.AmountCents,
		"currency":     strings.ToUpper(strings.TrimSpace(req.Currency)),
		"note":         strings.TrimSpace(req.Note),
		"status":       "accepted",
	}
}

func validatePayment(req createPaymentRequest) error {
	if req.AmountCents <= 0 {
		return http.ErrContentLength
	}
	if len(strings.TrimSpace(req.Currency)) != 3 {
		return http.ErrNotSupported
	}
	return nil
}

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
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payment"})
			return
		}
		fp, _ := fingerprintJSON(req)
		scopedKey := userID + ":" + key
		if rec, ok := store.Get(scopedKey); ok {
			if rec.Fingerprint != fp {
				writeJSON(w, http.StatusConflict, map[string]string{"error": "idempotency key reused with different payload"})
				return
			}
			writeJSON(w, rec.Status, rec.Body)
			return
		}
		payload := service.Create(req)
		store.Set(scopedKey, idempotencyRecord{Fingerprint: fp, Status: http.StatusCreated, Body: payload})
		writeJSON(w, http.StatusCreated, payload)
	}
}

func buildMux(store TokenStore, ids *IdempotencyStore, service *PaymentService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/payments", authMiddleware(store, http.HandlerFunc(createPaymentHandler(service, ids))))
	return mux
}

func newStore() *InMemoryTokenStore {
	return &InMemoryTokenStore{
		tokens: map[string]struct {
			UserID string
			Role   string
		}{
			"student-token": {UserID: "student-1", Role: "student"},
		},
	}
}

func makePaymentBody(amount int64, currency string, note string) *bytes.Buffer {
	payload := map[string]any{"amount_cents": amount, "currency": currency, "note": note}
	data, _ := json.Marshal(payload)
	return bytes.NewBuffer(data)
}

func TestLesson1MissingTokenIs401(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req.Header.Set("Idempotency-Key", "k1")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", w.Code)
	}
}

func TestLesson2MissingIdempotencyKeyIs400(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req.Header.Set("Authorization", "Bearer student-token")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}

func TestLesson3CreatePaymentReturns201(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req.Header.Set("Authorization", "Bearer student-token")
	req.Header.Set("Idempotency-Key", "k1")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d", w.Code)
	}
}

func TestLesson4SameKeySamePayloadReturnsSameResult(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req1 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req1.Header.Set("Authorization", "Bearer student-token")
	req1.Header.Set("Idempotency-Key", "k1")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req2.Header.Set("Authorization", "Bearer student-token")
	req2.Header.Set("Idempotency-Key", "k1")
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w1.Body.String() != w2.Body.String() {
		t.Fatalf("expected same response body for same key/payload")
	}
}

func TestLesson5SameKeyDifferentPayloadIs409(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req1 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req1.Header.Set("Authorization", "Bearer student-token")
	req1.Header.Set("Idempotency-Key", "k1")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(2000, "USD", "b"))
	req2.Header.Set("Authorization", "Bearer student-token")
	req2.Header.Set("Idempotency-Key", "k1")
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Fatalf("want 409, got %d", w2.Code)
	}
}

func TestLesson6DifferentKeysCreateDifferentRecords(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req1 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req1.Header.Set("Authorization", "Bearer student-token")
	req1.Header.Set("Idempotency-Key", "k1")
	w1 := httptest.NewRecorder()
	mux.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(1000, "USD", "a"))
	req2.Header.Set("Authorization", "Bearer student-token")
	req2.Header.Set("Idempotency-Key", "k2")
	w2 := httptest.NewRecorder()
	mux.ServeHTTP(w2, req2)

	if w1.Body.String() == w2.Body.String() {
		t.Fatalf("different keys should create different IDs")
	}
}

func TestLesson7InvalidJSONIs400(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewBufferString("{bad"))
	req.Header.Set("Authorization", "Bearer student-token")
	req.Header.Set("Idempotency-Key", "k1")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}

func TestLesson8InvalidPaymentIs400(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodPost, "/payments", makePaymentBody(0, "US", "bad"))
	req.Header.Set("Authorization", "Bearer student-token")
	req.Header.Set("Idempotency-Key", "k1")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", w.Code)
	}
}

func TestLesson9MethodNotAllowedIs405(t *testing.T) {
	mux := buildMux(newStore(), NewIdempotencyStore(), NewPaymentService())
	req := httptest.NewRequest(http.MethodGet, "/payments", nil)
	req.Header.Set("Authorization", "Bearer student-token")
	req.Header.Set("Idempotency-Key", "k1")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("want 405, got %d", w.Code)
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Retry Safety Tests 1-10
