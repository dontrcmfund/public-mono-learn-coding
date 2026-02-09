package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
GO INTERFACES + DI (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/77-go-interfaces-di-1-10.go
2) Replace the notifier implementation and observe unchanged service logic

Extra context:
- lessons/notes/164-go-interfaces-first-principles.md
- lessons/notes/165-go-dependency-injection-first-principles.md
*/

// LESSON 1: Interface at a boundary
// Why this matters: service depends on behavior, not concrete technology.
type Notifier interface {
	Send(recipient string, message string) error
}

// LESSON 2: Concrete implementation #1
// Why this matters: one concrete adapter for production-like behavior.
type ConsoleNotifier struct{}

func (ConsoleNotifier) Send(recipient string, message string) error {
	fmt.Printf("console notifier -> to=%s message=%q\n", recipient, message)
	return nil
}

// LESSON 3: Concrete implementation #2
// Why this matters: swapping implementation should not change business logic.
type MemoryNotifier struct {
	Sent []string
}

func (m *MemoryNotifier) Send(recipient string, message string) error {
	m.Sent = append(m.Sent, recipient+": "+message)
	return nil
}

// LESSON 4: Domain model
// Why this matters: explicit domain types improve readability.
type WelcomeUser struct {
	Email string
	Name  string
}

// LESSON 5: Service with injected dependency
// Why this matters: constructor-based DI makes dependencies visible.
type WelcomeService struct {
	notifier Notifier
}

func NewWelcomeService(notifier Notifier) *WelcomeService {
	return &WelcomeService{notifier: notifier}
}

// LESSON 6: Input validation
// Why this matters: fail early with clear errors.
func validateUser(u WelcomeUser) error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name is required")
	}
	email := strings.TrimSpace(strings.ToLower(u.Email))
	if !strings.Contains(email, "@") {
		return errors.New("valid email is required")
	}
	return nil
}

// LESSON 7: Small message formatter
// Why this matters: separating formatting from workflow keeps code testable.
func formatWelcomeMessage(name string) string {
	return "Welcome, " + strings.TrimSpace(name) + "! Your account is ready."
}

// LESSON 8: Service workflow method
// Why this matters: workflow logic is stable even if dependencies change.
func (s *WelcomeService) SendWelcome(u WelcomeUser) error {
	if err := validateUser(u); err != nil {
		return err
	}
	msg := formatWelcomeMessage(u.Name)
	return s.notifier.Send(strings.TrimSpace(strings.ToLower(u.Email)), msg)
}

// LESSON 9: Wiring at the application edge
// Why this matters: `main` decides concrete dependencies, service stays clean.

// LESSON 10: End-to-end demo with implementation swap
// Why this matters: demonstrates flexibility from interfaces + DI.
func main() {
	user := WelcomeUser{Name: "Mia", Email: " MIA@example.com "}

	consoleService := NewWelcomeService(ConsoleNotifier{})
	if err := consoleService.SendWelcome(user); err != nil {
		fmt.Println("Lesson 10 console error:", err)
	}

	memoryAdapter := &MemoryNotifier{}
	memoryService := NewWelcomeService(memoryAdapter)
	if err := memoryService.SendWelcome(user); err != nil {
		fmt.Println("Lesson 10 memory error:", err)
	}
	fmt.Println("Lesson 10 memory adapter sent:", memoryAdapter.Sent)
}

// End of Go Interfaces + DI 1-10
