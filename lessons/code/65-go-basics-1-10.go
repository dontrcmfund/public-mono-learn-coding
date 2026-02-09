package main

import "fmt"

/*
GO BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/65-go-basics-1-10.go
2) Change one line and predict output before running

Extra context:
- lessons/notes/150-what-is-go.md
- lessons/notes/151-go-first-principles.md
- lessons/notes/152-go-gotchas.md
*/

// LESSON 1: Basic output
// Why this matters: observable output closes the learning loop.

// LESSON 2: Variables and types
// Why this matters: explicit types reduce ambiguity.

// LESSON 3: Short declaration vs assignment
// Why this matters: helps avoid shadowing mistakes.

// LESSON 4: Constants
// Why this matters: immutable values protect intent.

// LESSON 5: Basic math
// Why this matters: many workflows are transform + calculate.

// LESSON 6: If/else
// Why this matters: programs branch on conditions.

// LESSON 7: For loop
// Why this matters: controlled repetition is core logic.

// LESSON 8: Arrays and slices
// Why this matters: slices are the primary sequence type in Go.

// LESSON 9: Maps
// Why this matters: key-value lookup is common in services.

// LESSON 10: Functions
// Why this matters: reusable, testable logic units.

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Lesson 1: Hello, Go")

	var name string = "Mia"
	attempts := 2
	fmt.Println("Lesson 2:", name, attempts)

	attempts = attempts + 1
	fmt.Println("Lesson 3:", attempts)

	const serviceName string = "go-lessons"
	fmt.Println("Lesson 4:", serviceName)

	a, b := 10, 3
	fmt.Println("Lesson 5:", a+b, a-b, a*b, a/b)

	score := 82
	if score >= 90 {
		fmt.Println("Lesson 6: grade A")
	} else if score >= 80 {
		fmt.Println("Lesson 6: grade B")
	} else {
		fmt.Println("Lesson 6: grade C")
	}

	sum := 0
	for i := 1; i <= 3; i++ {
		sum += i
	}
	fmt.Println("Lesson 7:", sum)

	arr := [3]int{1, 2, 3}
	slice := arr[:]
	slice = append(slice, 4)
	fmt.Println("Lesson 8:", arr, slice)

	scores := map[string]int{"mia": 92, "leo": 78}
	scores["ana"] = 88
	fmt.Println("Lesson 9:", scores["ana"], len(scores))

	fmt.Println("Lesson 10:", add(5, 7))
}

// End of Go Basics 1-10
