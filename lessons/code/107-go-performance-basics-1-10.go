package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
GO PERFORMANCE BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/107-go-performance-basics-1-10.go
2) Change data size and compare latency/throughput output

Extra context:
- lessons/notes/194-performance-first-principles.md
*/

// LESSON 1: Workload generator
// Why this matters: performance must be measured on explicit workloads.
func generateData(size int) []int {
	out := make([]int, size)
	for i := range out {
		out[i] = rand.Intn(size * 10)
	}
	return out
}

// LESSON 2: Baseline operation
// Why this matters: sorting is a clear, measurable CPU task.
func sortData(items []int) []int {
	copyItems := make([]int, len(items))
	copy(copyItems, items)
	sort.Ints(copyItems)
	return copyItems
}

// LESSON 3: Latency measurement helper
// Why this matters: a single number makes behavior comparable.
func measureLatency(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// LESSON 4: Throughput measurement helper
// Why this matters: capacity planning needs operations-per-second view.
func measureThroughput(fn func(), runs int) float64 {
	start := time.Now()
	for i := 0; i < runs; i++ {
		fn()
	}
	elapsed := time.Since(start).Seconds()
	if elapsed == 0 {
		return 0
	}
	return float64(runs) / elapsed
}

// LESSON 5: Percentile helper
// Why this matters: p95 latency is often more useful than average.
func percentile(durations []time.Duration, pct float64) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	cp := make([]time.Duration, len(durations))
	copy(cp, durations)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	idx := int(float64(len(cp)-1) * pct)
	if idx < 0 {
		idx = 0
	}
	if idx >= len(cp) {
		idx = len(cp) - 1
	}
	return cp[idx]
}

// LESSON 6: Repeated latency sampling
// Why this matters: one run is noisy; repeated runs show distribution.
func sampleLatencies(fn func(), samples int) []time.Duration {
	out := make([]time.Duration, 0, samples)
	for i := 0; i < samples; i++ {
		out = append(out, measureLatency(fn))
	}
	return out
}

// LESSON 7: Simple workload variants
// Why this matters: performance changes with input size.
func runWorkload(size int) {
	data := generateData(size)
	latencies := sampleLatencies(func() { _ = sortData(data) }, 20)
	p50 := percentile(latencies, 0.50)
	p95 := percentile(latencies, 0.95)
	tp := measureThroughput(func() { _ = sortData(data) }, 50)
	fmt.Printf("size=%d p50=%s p95=%s throughput=%.2f ops/sec\n", size, p50, p95, tp)
}

// LESSON 8: Warm-up note
// Why this matters: first run can differ due to setup and caches.

// LESSON 9: Compare two sizes
// Why this matters: scaling curve is more informative than one point.

// LESSON 10: End-to-end measurement demo
// Why this matters: establishes baseline before optimization attempts.
func main() {
	rand.Seed(42)
	fmt.Println("Go performance baseline measurements")
	runWorkload(1_000)
	runWorkload(10_000)
}

// End of Go Performance Basics 1-10
