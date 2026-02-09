package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

/*
GO ERRORS + CONTEXT TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/76-go-errors-context-tests-1-10_test.go -run TestLesson -v
2) Compare each test with the matching concept in 75-go-errors-context-11-20.go
*/

var ErrEmptyName = errors.New("empty name")

func normalizeName(raw string) (string, error) {
	clean := strings.TrimSpace(raw)
	if clean == "" {
		return "", ErrEmptyName
	}
	return strings.ToLower(clean), nil
}

func parseUser(raw string) (string, error) {
	name, err := normalizeName(raw)
	if err != nil {
		return "", fmt.Errorf("parse user: %w", err)
	}
	return name, nil
}

func fetchProfile(ctx context.Context, user string, delay time.Duration) (string, error) {
	select {
	case <-time.After(delay):
		return "profile:" + user, nil
	case <-ctx.Done():
		return "", fmt.Errorf("fetch profile timeout/cancel: %w", ctx.Err())
	}
}

func TestLesson1NormalizeName(t *testing.T) {
	got, err := normalizeName("  Mia  ")
	if err != nil || got != "mia" {
		t.Fatalf("want mia with nil err, got %q and %v", got, err)
	}
}

func TestLesson2NormalizeNameEmpty(t *testing.T) {
	_, err := normalizeName("   ")
	if !errors.Is(err, ErrEmptyName) {
		t.Fatalf("want ErrEmptyName, got %v", err)
	}
}

func TestLesson3ParseWrapsError(t *testing.T) {
	_, err := parseUser("  ")
	if err == nil || !strings.Contains(err.Error(), "parse user:") {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestLesson4ErrorsIsThroughWrap(t *testing.T) {
	_, err := parseUser("  ")
	if !errors.Is(err, ErrEmptyName) {
		t.Fatalf("errors.Is should match wrapped sentinel")
	}
}

func TestLesson5FetchProfileSuccess(t *testing.T) {
	got, err := fetchProfile(context.Background(), "leo", 1*time.Millisecond)
	if err != nil || got != "profile:leo" {
		t.Fatalf("want profile:leo with nil err, got %q and %v", got, err)
	}
}

func TestLesson6FetchProfileTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	_, err := fetchProfile(ctx, "leo", 15*time.Millisecond)
	if err == nil {
		t.Fatalf("expected timeout/cancel error")
	}
}

func TestLesson7FetchProfileContextDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	_, err := fetchProfile(ctx, "leo", 20*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("want context.DeadlineExceeded in wrapped error, got %v", err)
	}
}

func TestLesson8DeterministicNormalize(t *testing.T) {
	a, _ := normalizeName("  ALEX ")
	b, _ := normalizeName("  ALEX ")
	if a != b {
		t.Fatalf("normalize should be deterministic")
	}
}

func TestLesson9WhitespaceKinds(t *testing.T) {
	got, err := normalizeName("\n\t Mia \t")
	if err != nil || got != "mia" {
		t.Fatalf("want mia with nil err, got %q and %v", got, err)
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Errors + Context Tests 1-10
