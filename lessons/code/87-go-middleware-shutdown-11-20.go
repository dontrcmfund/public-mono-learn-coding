package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
GO MIDDLEWARE + SHUTDOWN (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/87-go-middleware-shutdown-11-20.go
2) Call: curl http://localhost:8086/health
3) Stop with Ctrl+C and observe graceful shutdown logs

Extra context:
- lessons/notes/174-go-observability-first-principles.md
- lessons/notes/175-go-graceful-shutdown-gotchas.md
*/

// LESSON 11: JSON helper
// Why this matters: keep response format consistent across handlers.
func writeJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 12: Request logging middleware
// Why this matters: every request should leave an observability trail.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("level=INFO msg=%q method=%s path=%s duration_ms=%d",
			"request.completed",
			r.Method,
			r.URL.Path,
			time.Since(start).Milliseconds(),
		)
	})
}

// LESSON 13: Recovery middleware
// Why this matters: one handler panic should not crash whole server.
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("level=ERROR msg=%q panic=%v path=%s", "handler.panic_recovered", rec, r.URL.Path)
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// LESSON 14: Middleware composition helper
// Why this matters: explicit order avoids confusing behavior.
func chain(mw ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		h := final
		for i := len(mw) - 1; i >= 0; i-- {
			h = mw[i](h)
		}
		return h
	}
}

// LESSON 15: Health handler
// Why this matters: baseline endpoint for runtime checks.
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// LESSON 16: Example panic handler
// Why this matters: demonstrates recovery middleware value.
func panicHandler(w http.ResponseWriter, _ *http.Request) {
	panic("lesson panic")
}

// LESSON 17: Build server with timeouts
// Why this matters: timeouts protect resources under slow clients.
func buildServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
}

// LESSON 18: Graceful shutdown signal handling
// Why this matters: restart/deploy events should not drop requests abruptly.
func waitForShutdownSignal() os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	return <-signals
}

// LESSON 19: Graceful shutdown with timeout context
// Why this matters: bounded shutdown avoids hanging forever.
func gracefulShutdown(server *http.Server, timeoutSeconds int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}

// LESSON 20: End-to-end operational flow
// Why this matters: combines middleware + timeouts + controlled shutdown.
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/panic", panicHandler)

	wrapped := chain(loggingMiddleware, recoveryMiddleware)(mux)
	server := buildServer(":8086", wrapped)

	go func() {
		log.Printf("level=INFO msg=%q addr=%s", "server.starting", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("level=ERROR msg=%q error=%v", "server.listen_failed", err)
		}
	}()

	sig := waitForShutdownSignal()
	log.Printf("level=INFO msg=%q signal=%s", "server.shutdown_requested", sig.String())

	timeoutSeconds := 8
	if raw := os.Getenv("SHUTDOWN_TIMEOUT_S"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			timeoutSeconds = parsed
		}
	}

	if err := gracefulShutdown(server, timeoutSeconds); err != nil {
		log.Printf("level=ERROR msg=%q error=%v", "server.shutdown_failed", err)
		return
	}
	fmt.Println("Lesson 20: graceful shutdown complete")
}

// End of Go Middleware + Shutdown 11-20
