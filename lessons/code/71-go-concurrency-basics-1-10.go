package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
GO CONCURRENCY BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/71-go-concurrency-basics-1-10.go
2) Observe order and timing of outputs
3) Modify worker count and task count

Extra context:
- lessons/notes/157-go-concurrency-first-principles.md
- lessons/notes/159-go-concurrency-gotchas.md
*/

// LESSON 1: Start a goroutine
// Why this matters: goroutines are lightweight concurrent units.

// LESSON 2: WaitGroup synchronization
// Why this matters: wait for concurrent work completion.

// LESSON 3: Channels for communication
// Why this matters: safe data transfer between goroutines.

// LESSON 4: Buffered channels
// Why this matters: small queues reduce blocking in simple pipelines.

// LESSON 5: Worker pool pattern
// Why this matters: bounded concurrency controls resource usage.

// LESSON 6: Fan-out/fan-in
// Why this matters: split work and combine results.

// LESSON 7: Select with timeout
// Why this matters: prevent indefinite waits.

// LESSON 8: Context cancellation
// Why this matters: stop work when request is no longer needed.

// LESSON 9: Safe shared state with mutex
// Why this matters: avoid race conditions.

// LESSON 10: End-to-end concurrent processing demo
// Why this matters: combine primitives into practical flow.

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Lesson 1: goroutine started")
	}()
	wg.Wait()

	tasks := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		tasks <- i
	}
	close(tasks)

	worker := func(id int) {
		for n := range tasks {
			time.Sleep(20 * time.Millisecond)
			results <- n * n
			fmt.Printf("Lesson 5 worker %d processed %d\n", id, n)
		}
	}

	wg = sync.WaitGroup{}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("Lesson 6 fan-in sum:", sum)

	// Timeout select
	slow := make(chan string)
	go func() {
		time.Sleep(100 * time.Millisecond)
		slow <- "done"
	}()

	select {
	case msg := <-slow:
		fmt.Println("Lesson 7:", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Lesson 7: timeout")
	}

	// Context cancellation
	ctx, cancel := context.WithCancel(context.Background())
	cancelled := make(chan struct{})
	go func() {
		defer close(cancelled)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Lesson 8: context cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	cancel()
	<-cancelled

	// Mutex example
	var mu sync.Mutex
	counter := 0
	wg = sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Lesson 9 counter:", counter)

	fmt.Println("Lesson 10: concurrency flow complete")
}

// End of Go Concurrency Basics 1-10
