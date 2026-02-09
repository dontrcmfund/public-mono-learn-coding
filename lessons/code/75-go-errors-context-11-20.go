package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
GO ERRORS + CONTEXT (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/75-go-errors-context-11-20.go
2) Change timeout values and observe behavior

Extra context:
- lessons/notes/163-go-errors-and-context-first-principles.md
- lessons/notes/152-go-gotchas.md
*/

var ErrEmptyName = errors.New("empty name")

// LESSON 11: Sentinel errors
// Why this matters: stable error identity supports robust checks.

// LESSON 12: Input validation with explicit errors
// Why this matters: bad input should fail quickly and clearly.
func normalizeName(raw string) (string, error) {
	clean := strings.TrimSpace(raw)
	if clean == "" {
		return "", ErrEmptyName
	}
	return strings.ToLower(clean), nil
}

// LESSON 13: Error wrapping
// Why this matters: wrapping preserves root cause and adds context.
func parseUser(raw string) (string, error) {
	name, err := normalizeName(raw)
	if err != nil {
		return "", fmt.Errorf("parse user: %w", err)
	}
	return name, nil
}

// LESSON 14: errors.Is for sentinel matching
// Why this matters: behavior checks should survive wrapping layers.

// LESSON 15: Context timeout
// Why this matters: external work should not run forever.
func fetchProfile(ctx context.Context, user string, delay time.Duration) (string, error) {
	select {
	case <-time.After(delay):
		return "profile:" + user, nil
	case <-ctx.Done():
		return "", fmt.Errorf("fetch profile timeout/cancel: %w", ctx.Err())
	}
}

// LESSON 16: Context cancel propagation
// Why this matters: one cancellation signal can stop chained work.

// LESSON 17: Separate domain error from transport decision
// Why this matters: business logic stays reusable outside HTTP.
func statusFromError(err error) int {
	if err == nil {
		return 200
	}
	if errors.Is(err, ErrEmptyName) {
		return 400
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return 504
	}
	return 500
}

// LESSON 18: Return early pattern
// Why this matters: linear control flow is easier to reason about.
func buildGreeting(ctx context.Context, raw string, delay time.Duration) (string, error) {
	name, err := parseUser(raw)
	if err != nil {
		return "", err
	}
	profile, err := fetchProfile(ctx, name, delay)
	if err != nil {
		return "", err
	}
	return "hello " + name + " (" + profile + ")", nil
}

// LESSON 19: Observable failure messages
// Why this matters: logs become useful debugging tools.

// LESSON 20: End-to-end flow with success and failure paths
// Why this matters: students see how errors and context combine in practice.
func main() {
	msg, err := buildGreeting(context.Background(), "  Mia  ", 10*time.Millisecond)
	fmt.Println("Lesson 20 success:", msg, "error:", err, "status:", statusFromError(err))

	_, err = buildGreeting(context.Background(), "   ", 10*time.Millisecond)
	fmt.Println("Lesson 12/14 empty-name error:", err, "status:", statusFromError(err))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	_, err = buildGreeting(ctx, "Leo", 20*time.Millisecond)
	fmt.Println("Lesson 15/16 timeout error:", err, "status:", statusFromError(err))
}

// End of Go Errors + Context 11-20
