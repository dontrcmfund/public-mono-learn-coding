package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

/*
GO DI + HTTP BRIDGE (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/79-go-di-http-bridge-11-20.go
2) Test:
   - POST http://localhost:8083/welcome {"name":"Mia","email":"mia@example.com"}

Extra context:
- lessons/notes/164-go-interfaces-first-principles.md
- lessons/notes/165-go-dependency-injection-first-principles.md
*/

// LESSON 11: Notification boundary interface
// Why this matters: HTTP layer should not care about delivery technology.
type Notifier interface {
	Send(recipient string, message string) error
}

// LESSON 12: Infrastructure adapter
// Why this matters: isolate external behavior at the edge.
type ConsoleNotifier struct{}

func (ConsoleNotifier) Send(recipient string, message string) error {
	fmt.Printf("SEND -> to=%s message=%q at=%s\n", recipient, message, time.Now().Format(time.RFC3339))
	return nil
}

// LESSON 13: Domain input model
// Why this matters: clear input contracts simplify validation.
type WelcomeInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// LESSON 14: Service with DI
// Why this matters: transport-independent business workflow.
type WelcomeService struct {
	notifier Notifier
}

func NewWelcomeService(notifier Notifier) *WelcomeService {
	return &WelcomeService{notifier: notifier}
}

func (s *WelcomeService) SendWelcome(input WelcomeInput) error {
	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(strings.ToLower(input.Email))

	if name == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(email, "@") {
		return errors.New("valid email is required")
	}

	message := "Welcome, " + name + "! You are all set."
	return s.notifier.Send(email, message)
}

// LESSON 15: HTTP request DTO
// Why this matters: explicit request shape at transport boundary.
type welcomeRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 16: Thin handler
// Why this matters: handler translates, service decides.
func makeWelcomeHandler(service *WelcomeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}

		var req welcomeRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		err := service.SendWelcome(WelcomeInput{Name: req.Name, Email: req.Email})
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]string{"status": "sent"})
	}
}

// LESSON 17: Health endpoint
// Why this matters: operational confidence and simple diagnostics.
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// LESSON 18: Routing composition
// Why this matters: one place to see transport wiring.
func buildMux(service *WelcomeService) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/welcome", makeWelcomeHandler(service))
	return mux
}

// LESSON 19: Composition root (`main`)
// Why this matters: keep object creation at system edge.

// LESSON 20: End-to-end server bootstrap
// Why this matters: completes boundary-to-service flow.
func main() {
	service := NewWelcomeService(ConsoleNotifier{})
	mux := buildMux(service)

	addr := ":8083"
	fmt.Println("Go DI + HTTP bridge lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go DI + HTTP Bridge 11-20
