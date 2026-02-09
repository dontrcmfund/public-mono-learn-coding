package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

/*
GO BENCHMARK TESTS (Lessons 1-10)

Suggested use:
1) Run tests:      go test lessons/code/109-go-benchmark-tests-1-10_test.go -run TestLesson -v
2) Run benchmarks: go test lessons/code/109-go-benchmark-tests-1-10_test.go -bench BenchmarkLesson -benchmem
*/

func generateData(size int) []int {
	out := make([]int, size)
	for i := range out {
		out[i] = rand.Intn(size * 10)
	}
	return out
}

func sortData(items []int) []int {
	cp := make([]int, len(items))
	copy(cp, items)
	sort.Ints(cp)
	return cp
}

func isSorted(items []int) bool {
	for i := 1; i < len(items); i++ {
		if items[i] < items[i-1] {
			return false
		}
	}
	return true
}

func TestLesson1SortProducesOrderedOutput(t *testing.T) {
	in := []int{5, 1, 4, 2, 3}
	out := sortData(in)
	if !isSorted(out) {
		t.Fatalf("expected sorted output")
	}
}

func TestLesson2SortDoesNotMutateInput(t *testing.T) {
	in := []int{3, 2, 1}
	_ = sortData(in)
	if in[0] != 3 || in[1] != 2 || in[2] != 1 {
		t.Fatalf("input mutated")
	}
}

func TestLesson3GenerateDataSize(t *testing.T) {
	d := generateData(100)
	if len(d) != 100 {
		t.Fatalf("want len=100, got %d", len(d))
	}
}

func TestLesson4SortedDeterministic(t *testing.T) {
	in := []int{9, 7, 8}
	a := sortData(in)
	b := sortData(in)
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("expected deterministic output")
		}
	}
}

func TestLesson5SortEmpty(t *testing.T) {
	out := sortData([]int{})
	if len(out) != 0 {
		t.Fatalf("expected empty output")
	}
}

func TestLesson6SortSingle(t *testing.T) {
	out := sortData([]int{42})
	if len(out) != 1 || out[0] != 42 {
		t.Fatalf("unexpected single-item sort result")
	}
}

func TestLesson7SortHandlesDuplicates(t *testing.T) {
	out := sortData([]int{3, 1, 2, 2, 1})
	if !isSorted(out) {
		t.Fatalf("expected sorted output with duplicates")
	}
}

func TestLesson8GenerateDeterministicWithSeed(t *testing.T) {
	rand.Seed(7)
	a := generateData(5)
	rand.Seed(7)
	b := generateData(5)
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("expected deterministic generation with fixed seed")
		}
	}
}

func TestLesson9LatencyMeasurementSanity(t *testing.T) {
	start := time.Now()
	_ = sortData(generateData(1000))
	if time.Since(start) <= 0 {
		t.Fatalf("expected positive duration")
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

func BenchmarkLesson1Sort1K(b *testing.B) {
	rand.Seed(42)
	data := generateData(1_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sortData(data)
	}
}

func BenchmarkLesson2Sort10K(b *testing.B) {
	rand.Seed(42)
	data := generateData(10_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sortData(data)
	}
}

func BenchmarkLesson3SortWithAllocationReport(b *testing.B) {
	rand.Seed(42)
	data := generateData(2_000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sortData(data)
	}
}

// End of Go Benchmark Tests 1-10
