package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
GO CIRCUIT BREAKER + BULKHEAD (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/111-go-circuit-breaker-bulkhead-11-20.go
2) Observe breaker opening after failures and bulkhead limiting concurrency

Extra context:
- lessons/notes/198-circuit-breaker-and-bulkhead-basics.md
*/

type DependencyCall func(ctx context.Context) (string, error)

type BreakerState string

const (
	StateClosed   BreakerState = "closed"
	StateOpen     BreakerState = "open"
	StateHalfOpen BreakerState = "half_open"
)

// LESSON 11: Circuit breaker model
// Why this matters: fail-fast behavior prevents repeated costly failures.
type CircuitBreaker struct {
	mu            sync.Mutex
	state         BreakerState
	failureCount  int
	failureThresh int
	cooldownUntil time.Time
	cooldown      time.Duration
}

func NewCircuitBreaker(threshold int, cooldown time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:         StateClosed,
		failureThresh: threshold,
		cooldown:      cooldown,
	}
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

func (b *CircuitBreaker) State() BreakerState {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.state
}

// LESSON 12: Breaker execution wrapper
// Why this matters: consistent policy around dependency calls.
func (b *CircuitBreaker) Execute(ctx context.Context, call DependencyCall) (string, error) {
	now := time.Now()
	if !b.allow(now) {
		return "", fmt.Errorf("circuit open")
	}
	result, err := call(ctx)
	if err != nil {
		b.onFailure(now)
		return "", err
	}
	b.onSuccess()
	return result, nil
}

// LESSON 13: Bulkhead limiter
// Why this matters: cap concurrent calls to protect shared resources.
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

// LESSON 14: Simulated unstable dependency
// Why this matters: demonstrates breaker transitions.
func unstableDependencyFactory(failUntil int) DependencyCall {
	var mu sync.Mutex
	attempt := 0
	return func(ctx context.Context) (string, error) {
		mu.Lock()
		attempt++
		current := attempt
		mu.Unlock()
		select {
		case <-time.After(15 * time.Millisecond):
			if current <= failUntil {
				return "", fmt.Errorf("dependency temporary failure")
			}
			return "service-ok", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// LESSON 15: Combined resilience execution
// Why this matters: real systems layer protections.
func executeProtected(ctx context.Context, bulkhead *Bulkhead, breaker *CircuitBreaker, dep DependencyCall) (string, error) {
	return bulkhead.Execute(func() (string, error) {
		return breaker.Execute(ctx, dep)
	})
}

// LESSON 16: Breaker demo
// Why this matters: visible state changes teach fail-fast behavior.
func demoBreaker() {
	breaker := NewCircuitBreaker(2, 50*time.Millisecond)
	bulkhead := NewBulkhead(2)
	dep := unstableDependencyFactory(3)
	for i := 0; i < 5; i++ {
		res, err := executeProtected(context.Background(), bulkhead, breaker, dep)
		fmt.Printf("breaker-demo call=%d state=%s result=%q err=%v\n", i+1, breaker.State(), res, err)
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(60 * time.Millisecond)
	res, err := executeProtected(context.Background(), bulkhead, breaker, dep)
	fmt.Printf("breaker-demo after-cooldown state=%s result=%q err=%v\n", breaker.State(), res, err)
}

// LESSON 17: Bulkhead demo
// Why this matters: one lane can reject overload while others keep working.
func demoBulkhead() {
	bulkhead := NewBulkhead(1)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			res, err := bulkhead.Execute(func() (string, error) {
				time.Sleep(30 * time.Millisecond)
				return fmt.Sprintf("worker-%d-done", id), nil
			})
			fmt.Printf("bulkhead-demo worker=%d result=%q err=%v\n", id, res, err)
		}(i)
	}
	wg.Wait()
}

// LESSON 18: Fallback integration point
// Why this matters: breaker open/saturation can route to safe degraded response.

// LESSON 19: Observability integration point
// Why this matters: state transitions and saturation should be visible.

// LESSON 20: End-to-end resilience pattern demo
// Why this matters: combines isolation and fail-fast recovery flow.
func main() {
	demoBreaker()
	demoBulkhead()
}

// End of Go Circuit Breaker + Bulkhead 11-20
