package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
GO HTTP OPERATIONAL TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/88-go-http-operational-tests-1-10_test.go -run TestLesson -v
2) Focus on why middleware behavior is tested, not just endpoint body
*/

func writeJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func chain(mw ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		h := final
		for i := len(mw) - 1; i >= 0; i-- {
			h = mw[i](h)
		}
		return h
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func panicHandler(w http.ResponseWriter, _ *http.Request) {
	panic("lesson panic")
}

func delayedHandler(delay time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(delay)
		writeJSON(w, http.StatusOK, map[string]string{"status": "done"})
	}
}

func buildTestMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/panic", panicHandler)
	mux.Handle("/delay", delayedHandler(5*time.Millisecond))
	return chain(loggingMiddleware, recoveryMiddleware)(mux)
}

func TestLesson1HealthStatusCode(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson2HealthJSONShape(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	var payload map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &payload); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("want status ok, got %q", payload["status"])
	}
}

func TestLesson3RecoveryMiddlewareReturns500(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("want 500, got %d", w.Code)
	}
}

func TestLesson4RecoveryMiddlewareJSONError(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/panic", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	var payload map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &payload)
	if payload["error"] != "internal server error" {
		t.Fatalf("unexpected error message: %q", payload["error"])
	}
}

func TestLesson5UnknownPathIs404(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("want 404, got %d", w.Code)
	}
}

func TestLesson6DelayedHandlerCompletes(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/delay", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson7ContentTypeSet(t *testing.T) {
	handler := buildTestMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if got := w.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("want application/json, got %q", got)
	}
}

func TestLesson8MiddlewareChainOrderSafe(t *testing.T) {
	base := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	handler := chain(loggingMiddleware, recoveryMiddleware)(base)
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestLesson9DeterministicHealth(t *testing.T) {
	handler := buildTestMux()
	req1 := httptest.NewRequest(http.MethodGet, "/health", nil)
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	req2 := httptest.NewRequest(http.MethodGet, "/health", nil)
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	if w1.Body.String() != w2.Body.String() {
		t.Fatalf("health response should be deterministic")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go HTTP Operational Tests 1-10
