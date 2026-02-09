package main

import (
	"errors"
	"strings"
	"testing"
)

/*
GO FAKES + TESTING (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/78-go-fakes-testing-1-10_test.go -run TestLesson -v
2) Compare each test with interfaces + DI lessons

Extra context:
- lessons/notes/166-go-fakes-and-mocks-gotchas.md
- lessons/notes/165-go-dependency-injection-first-principles.md
*/

type Notifier interface {
	Send(recipient string, message string) error
}

type WelcomeUser struct {
	Email string
	Name  string
}

type WelcomeService struct {
	notifier Notifier
}

func NewWelcomeService(notifier Notifier) *WelcomeService {
	return &WelcomeService{notifier: notifier}
}

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

func formatWelcomeMessage(name string) string {
	return "Welcome, " + strings.TrimSpace(name) + "! Your account is ready."
}

func (s *WelcomeService) SendWelcome(u WelcomeUser) error {
	if err := validateUser(u); err != nil {
		return err
	}
	return s.notifier.Send(strings.TrimSpace(strings.ToLower(u.Email)), formatWelcomeMessage(u.Name))
}

// Fake notifier for deterministic tests.
type FakeNotifier struct {
	Calls      int
	LastTo     string
	LastMsg    string
	ShouldFail bool
}

func (f *FakeNotifier) Send(recipient string, message string) error {
	f.Calls++
	f.LastTo = recipient
	f.LastMsg = message
	if f.ShouldFail {
		return errors.New("notifier unavailable")
	}
	return nil
}

func TestLesson1ValidUserPassesValidation(t *testing.T) {
	err := validateUser(WelcomeUser{Name: "Mia", Email: "mia@example.com"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestLesson2EmptyNameFailsValidation(t *testing.T) {
	err := validateUser(WelcomeUser{Name: " ", Email: "mia@example.com"})
	if err == nil {
		t.Fatalf("expected error for empty name")
	}
}

func TestLesson3InvalidEmailFailsValidation(t *testing.T) {
	err := validateUser(WelcomeUser{Name: "Mia", Email: "mia.example.com"})
	if err == nil {
		t.Fatalf("expected error for invalid email")
	}
}

func TestLesson4FormatWelcomeMessage(t *testing.T) {
	got := formatWelcomeMessage("  Mia ")
	want := "Welcome, Mia! Your account is ready."
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestLesson5SendWelcomeCallsNotifier(t *testing.T) {
	fake := &FakeNotifier{}
	service := NewWelcomeService(fake)
	err := service.SendWelcome(WelcomeUser{Name: "Mia", Email: "MIA@example.com"})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if fake.Calls != 1 {
		t.Fatalf("expected 1 call, got %d", fake.Calls)
	}
}

func TestLesson6SendWelcomeNormalizesEmail(t *testing.T) {
	fake := &FakeNotifier{}
	service := NewWelcomeService(fake)
	_ = service.SendWelcome(WelcomeUser{Name: "Mia", Email: " MIA@example.com "})
	if fake.LastTo != "mia@example.com" {
		t.Fatalf("want normalized email, got %q", fake.LastTo)
	}
}

func TestLesson7SendWelcomePropagatesDependencyError(t *testing.T) {
	fake := &FakeNotifier{ShouldFail: true}
	service := NewWelcomeService(fake)
	err := service.SendWelcome(WelcomeUser{Name: "Mia", Email: "mia@example.com"})
	if err == nil {
		t.Fatalf("expected error from dependency")
	}
}

func TestLesson8InvalidInputDoesNotCallNotifier(t *testing.T) {
	fake := &FakeNotifier{}
	service := NewWelcomeService(fake)
	_ = service.SendWelcome(WelcomeUser{Name: "", Email: "mia@example.com"})
	if fake.Calls != 0 {
		t.Fatalf("notifier should not be called on invalid input")
	}
}

func TestLesson9DeterministicBehavior(t *testing.T) {
	fake := &FakeNotifier{}
	service := NewWelcomeService(fake)
	_ = service.SendWelcome(WelcomeUser{Name: "Mia", Email: "mia@example.com"})
	first := fake.LastMsg
	_ = service.SendWelcome(WelcomeUser{Name: "Mia", Email: "mia@example.com"})
	second := fake.LastMsg
	if first != second {
		t.Fatalf("expected deterministic message output")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Fakes + Testing 1-10
