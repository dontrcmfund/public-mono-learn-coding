package main

import (
	"reflect"
	"testing"
)

/*
GO TESTING (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/67-go-testing-1-10_test.go -run TestLesson -v
2) Read failures as guidance, not as setbacks
3) Why this command looks specific: each lesson file here is intentionally standalone

Extra context:
- lessons/notes/154-go-testing-principles.md
- lessons/notes/152-go-gotchas.md
*/

func Add(a int, b int) int {
	return a + b
}

func Normalize(input string) string {
	// Simple trim-space replacement for lesson clarity.
	out := ""
	space := false
	for _, r := range input {
		if r == ' ' || r == '\t' || r == '\n' {
			if !space {
				out += " "
				space = true
			}
			continue
		}
		out += string(r)
		space = false
	}
	if len(out) > 0 && out[0] == ' ' {
		out = out[1:]
	}
	if len(out) > 0 && out[len(out)-1] == ' ' {
		out = out[:len(out)-1]
	}
	return out
}

func FilterGE(items []int, threshold int) []int {
	out := make([]int, 0)
	for _, n := range items {
		if n >= threshold {
			out = append(out, n)
		}
	}
	return out
}

// LESSON 1: Basic function test
func TestLesson1Add(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

// LESSON 2: Table-driven test for Add
func TestLesson2AddTable(t *testing.T) {
	cases := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}

	for _, c := range cases {
		got := Add(c.a, c.b)
		if got != c.want {
			t.Fatalf("Add(%d,%d): want %d, got %d", c.a, c.b, c.want, got)
		}
	}
}

// LESSON 3: String normalization test
func TestLesson3Normalize(t *testing.T) {
	got := Normalize("  hello   world  ")
	if got != "hello world" {
		t.Fatalf("want %q, got %q", "hello world", got)
	}
}

// LESSON 4: Edge case empty normalization
func TestLesson4NormalizeEmpty(t *testing.T) {
	if got := Normalize("   "); got != "" {
		t.Fatalf("want empty, got %q", got)
	}
}

// LESSON 5: Slice filter behavior
func TestLesson5FilterGE(t *testing.T) {
	got := FilterGE([]int{1, 5, 3, 8}, 4)
	want := []int{5, 8}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

// LESSON 6: Filter edge case none
func TestLesson6FilterNone(t *testing.T) {
	got := FilterGE([]int{1, 2, 3}, 10)
	if len(got) != 0 {
		t.Fatalf("want empty, got %v", got)
	}
}

// LESSON 7: Filter edge case all
func TestLesson7FilterAll(t *testing.T) {
	got := FilterGE([]int{5, 6, 7}, 5)
	want := []int{5, 6, 7}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

// LESSON 8: Deterministic output check
func TestLesson8Deterministic(t *testing.T) {
	in := []int{9, 1, 9, 2}
	got1 := FilterGE(in, 2)
	got2 := FilterGE(in, 2)
	if !reflect.DeepEqual(got1, got2) {
		t.Fatalf("expected deterministic output")
	}
}

// LESSON 9: Input immutability expectation
func TestLesson9NoInputMutation(t *testing.T) {
	in := []int{1, 2, 3}
	_ = FilterGE(in, 2)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(in, want) {
		t.Fatalf("input mutated: want %v got %v", want, in)
	}
}

// LESSON 10: Completion marker test
func TestLesson10Completion(t *testing.T) {
	if true != true {
		t.Fatalf("unreachable")
	}
}

// End of Go Testing 1-10
