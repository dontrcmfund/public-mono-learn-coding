package main

import (
	"context"
	"fmt"
	"time"
)

/*
GO RESILIENCE BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/110-go-resilience-basics-1-10.go
2) Change timeout values and observe fallback behavior

Extra context:
- lessons/notes/197-resilience-first-principles.md
- lessons/notes/199-timeout-fallback-gotchas.md
*/

// LESSON 1: Dependency function boundary
// Why this matters: resilience wrappers should work around dependency calls.
type DependencyCall func(ctx context.Context) (string, error)

// LESSON 2: Simulated dependency
// Why this matters: controlled delay/failure for repeatable learning.
func simulatedDependency(delay time.Duration, shouldFail bool) DependencyCall {
	return func(ctx context.Context) (string, error) {
		select {
		case <-time.After(delay):
			if shouldFail {
				return "", fmt.Errorf("dependency failure")
			}
			return "fresh-data", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// LESSON 3: Timeout wrapper
// Why this matters: bounded waiting prevents request pileups.
func withTimeout(parent context.Context, timeout time.Duration, call DependencyCall) (string, error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()
	return call(ctx)
}

// LESSON 4: Fallback wrapper
// Why this matters: degraded response can preserve user workflow.
func withFallback(value string, call func() (string, error)) (string, bool, error) {
	result, err := call()
	if err != nil {
		return value, true, nil
	}
	return result, false, nil
}

// LESSON 5: Resilient fetch helper
// Why this matters: combines timeout and fallback in one safe path.
func fetchDataResilient(timeout time.Duration, dep DependencyCall) (string, bool) {
	result, usedFallback, _ := withFallback("cached-data", func() (string, error) {
		return withTimeout(context.Background(), timeout, dep)
	})
	return result, usedFallback
}

// LESSON 6: Fast success path
// Why this matters: resilience should not hurt healthy path behavior.
func demoFastSuccess() {
	dep := simulatedDependency(20*time.Millisecond, false)
	value, fallback := fetchDataResilient(100*time.Millisecond, dep)
	fmt.Printf("fast-success value=%s fallback=%v\n", value, fallback)
}

// LESSON 7: Slow timeout path
// Why this matters: timeout prevents long blocking and triggers fallback.
func demoTimeoutFallback() {
	dep := simulatedDependency(120*time.Millisecond, false)
	value, fallback := fetchDataResilient(40*time.Millisecond, dep)
	fmt.Printf("timeout-fallback value=%s fallback=%v\n", value, fallback)
}

// LESSON 8: Explicit failure path
// Why this matters: failures should degrade predictably, not crash path.
func demoFailureFallback() {
	dep := simulatedDependency(20*time.Millisecond, true)
	value, fallback := fetchDataResilient(100*time.Millisecond, dep)
	fmt.Printf("failure-fallback value=%s fallback=%v\n", value, fallback)
}

// LESSON 9: Observability reminder
// Why this matters: timeout/fallback events should be tracked in real systems.

// LESSON 10: End-to-end resilience demo
// Why this matters: shows healthy and degraded behavior side by side.
func main() {
	demoFastSuccess()
	demoTimeoutFallback()
	demoFailureFallback()
}

// End of Go Resilience Basics 1-10
