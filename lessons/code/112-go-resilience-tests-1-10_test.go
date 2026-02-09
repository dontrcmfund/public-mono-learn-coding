package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
GO RESILIENCE TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/112-go-resilience-tests-1-10_test.go -run TestLesson -v
2) Focus on deterministic failure handling behavior
*/

type DependencyCall func(ctx context.Context) (string, error)

type BreakerState string

const (
	StateClosed   BreakerState = "closed"
	StateOpen     BreakerState = "open"
	StateHalfOpen BreakerState = "half_open"
)

type CircuitBreaker struct {
	mu            sync.Mutex
	state         BreakerState
	failureCount  int
	failureThresh int
	cooldownUntil time.Time
	cooldown      time.Duration
}

func NewCircuitBreaker(threshold int, cooldown time.Duration) *CircuitBreaker {
	return &CircuitBreaker{state: StateClosed, failureThresh: threshold, cooldown: cooldown}
}

func (b *CircuitBreaker) allow(now time.Time) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.state == StateOpen {
		if now.After(b.cooldownUntil) {
			b.state = StateHalfOpen
			return true
		}
		return false
	}
	return true
}

func (b *CircuitBreaker) onSuccess() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.state = StateClosed
	b.failureCount = 0
}

func (b *CircuitBreaker) onFailure(now time.Time) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failureCount++
	if b.failureCount >= b.failureThresh {
		b.state = StateOpen
		b.cooldownUntil = now.Add(b.cooldown)
	}
}

func (b *CircuitBreaker) Execute(ctx context.Context, call DependencyCall) (string, error) {
	now := time.Now()
	if !b.allow(now) {
		return "", fmt.Errorf("circuit open")
	}
	res, err := call(ctx)
	if err != nil {
		b.onFailure(now)
		return "", err
	}
	b.onSuccess()
	return res, nil
}

func (b *CircuitBreaker) State() BreakerState {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}

type Bulkhead struct {
	slots chan struct{}
}

func NewBulkhead(maxConcurrent int) *Bulkhead {
	return &Bulkhead{slots: make(chan struct{}, maxConcurrent)}
}

func (b *Bulkhead) Execute(call func() (string, error)) (string, error) {
	select {
	case b.slots <- struct{}{}:
		defer func() { <-b.slots }()
		return call()
	default:
		return "", fmt.Errorf("bulkhead saturated")
	}
}

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

func withTimeout(parent context.Context, timeout time.Duration, call DependencyCall) (string, error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()
	return call(ctx)
}

func withFallback(value string, call func() (string, error)) (string, bool, error) {
	res, err := call()
	if err != nil {
		return value, true, nil
	}
	return res, false, nil
}

func fetchDataResilient(timeout time.Duration, dep DependencyCall) (string, bool) {
	v, fb, _ := withFallback("cached-data", func() (string, error) {
		return withTimeout(context.Background(), timeout, dep)
	})
	return v, fb
}

func TestLesson1TimeoutTriggersFallback(t *testing.T) {
	dep := simulatedDependency(50*time.Millisecond, false)
	val, fallback := fetchDataResilient(5*time.Millisecond, dep)
	if !fallback || val != "cached-data" {
		t.Fatalf("expected fallback cached-data, got %q fallback=%v", val, fallback)
	}
}

func TestLesson2FastDependencyNoFallback(t *testing.T) {
	dep := simulatedDependency(5*time.Millisecond, false)
	val, fallback := fetchDataResilient(50*time.Millisecond, dep)
	if fallback || val != "fresh-data" {
		t.Fatalf("expected fresh-data without fallback, got %q fallback=%v", val, fallback)
	}
}

func TestLesson3BreakerOpensAfterThreshold(t *testing.T) {
	breaker := NewCircuitBreaker(2, 20*time.Millisecond)
	dep := simulatedDependency(1*time.Millisecond, true)
	_, _ = breaker.Execute(context.Background(), dep)
	_, _ = breaker.Execute(context.Background(), dep)
	if breaker.State() != StateOpen {
		t.Fatalf("expected breaker open, got %s", breaker.State())
	}
}

func TestLesson4BreakerFailsFastWhenOpen(t *testing.T) {
	breaker := NewCircuitBreaker(1, 100*time.Millisecond)
	dep := simulatedDependency(1*time.Millisecond, true)
	_, _ = breaker.Execute(context.Background(), dep)
	_, err := breaker.Execute(context.Background(), simulatedDependency(1*time.Millisecond, false))
	if err == nil || err.Error() != "circuit open" {
		t.Fatalf("expected circuit open error, got %v", err)
	}
}

func TestLesson5BreakerRecoversAfterCooldown(t *testing.T) {
	breaker := NewCircuitBreaker(1, 10*time.Millisecond)
	_, _ = breaker.Execute(context.Background(), simulatedDependency(1*time.Millisecond, true))
	time.Sleep(15 * time.Millisecond)
	_, err := breaker.Execute(context.Background(), simulatedDependency(1*time.Millisecond, false))
	if err != nil {
		t.Fatalf("expected recovery after cooldown, got %v", err)
	}
}

func TestLesson6BulkheadRejectsWhenSaturated(t *testing.T) {
	bulkhead := NewBulkhead(1)
	start := make(chan struct{})
	done := make(chan struct{})

	go func() {
		_, _ = bulkhead.Execute(func() (string, error) {
			close(start)
			time.Sleep(20 * time.Millisecond)
			close(done)
			return "ok", nil
		})
	}()

	<-start
	_, err := bulkhead.Execute(func() (string, error) {
		return "second", nil
	})
	if err == nil || err.Error() != "bulkhead saturated" {
		t.Fatalf("expected bulkhead saturated, got %v", err)
	}
	<-done
}

func TestLesson7BulkheadAllowsWithinLimit(t *testing.T) {
	bulkhead := NewBulkhead(2)
	_, err1 := bulkhead.Execute(func() (string, error) { return "a", nil })
	_, err2 := bulkhead.Execute(func() (string, error) { return "b", nil })
	if err1 != nil || err2 != nil {
		t.Fatalf("expected both calls allowed")
	}
}

func TestLesson8FallbackDeterministic(t *testing.T) {
	dep := simulatedDependency(40*time.Millisecond, false)
	a, _ := fetchDataResilient(1*time.Millisecond, dep)
	b, _ := fetchDataResilient(1*time.Millisecond, dep)
	if a != b {
		t.Fatalf("expected deterministic fallback output")
	}
}

func TestLesson9TimeoutErrorPathHandled(t *testing.T) {
	_, _, err := withFallback("cached-data", func() (string, error) {
		return withTimeout(context.Background(), 1*time.Millisecond, simulatedDependency(20*time.Millisecond, false))
	})
	if err != nil {
		t.Fatalf("withFallback should absorb timeout error, got %v", err)
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Resilience Tests 1-10
