package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

/*
GO PROJECTS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/68-go-projects-1-10.go
2) Change one business rule and predict output before running

Extra context:
- lessons/notes/155-go-projects-principles.md
- lessons/notes/151-go-first-principles.md
*/

type Expense struct {
	Category string
	Amount   float64
}

// LESSON 1: Aggregate totals by category
// Why this matters: report aggregation is common in real tools.
func TotalsByCategory(items []Expense) map[string]float64 {
	out := make(map[string]float64)
	for _, it := range items {
		out[it.Category] += it.Amount
	}
	return out
}

// LESSON 2: Input validation helper
// Why this matters: fail fast with clear errors.
func ValidateExpense(e Expense) error {
	if strings.TrimSpace(e.Category) == "" {
		return errors.New("category is required")
	}
	if e.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	return nil
}

// LESSON 3: Normalize category naming
// Why this matters: normalization prevents duplicate semantic keys.
func NormalizeCategory(raw string) string {
	return strings.ToLower(strings.TrimSpace(raw))
}

// LESSON 4: Build report lines
// Why this matters: separate formatting from computation.
func BuildReportLines(totals map[string]float64) []string {
	keys := make([]string, 0, len(totals))
	for k := range totals {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for _, k := range keys {
		lines = append(lines, fmt.Sprintf("- %s: %.2f", k, totals[k]))
	}
	return lines
}

// LESSON 5: Service-style processor
// Why this matters: encapsulates workflow logic.
func ProcessExpenses(items []Expense) ([]string, error) {
	normalized := make([]Expense, 0, len(items))
	for _, e := range items {
		e.Category = NormalizeCategory(e.Category)
		if err := ValidateExpense(e); err != nil {
			return nil, err
		}
		normalized = append(normalized, e)
	}
	totals := TotalsByCategory(normalized)
	return BuildReportLines(totals), nil
}

// LESSON 6: Basic command-like runner
// Why this matters: thin orchestration at the edge.
func RunReport(items []Expense) {
	lines, err := ProcessExpenses(items)
	if err != nil {
		fmt.Println("Lesson 6 error:", err)
		return
	}
	fmt.Println("Lesson 6 report:")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// LESSON 7: Search helper
// Why this matters: simple filtering utility for future CLI/API.
func FilterByMin(items []Expense, min float64) []Expense {
	out := make([]Expense, 0)
	for _, e := range items {
		if e.Amount >= min {
			out = append(out, e)
		}
	}
	return out
}

// LESSON 8: Budget status helper
// Why this matters: domain-oriented output can drive UI/API.
func BudgetStatus(limit float64, spent float64) string {
	if spent < limit {
		return "under"
	}
	if spent > limit {
		return "over"
	}
	return "exact"
}

// LESSON 9: Sum helper
// Why this matters: reusable single-purpose utility.
func SumAmounts(items []Expense) float64 {
	total := 0.0
	for _, e := range items {
		total += e.Amount
	}
	return total
}

// LESSON 10: End-to-end sample
// Why this matters: composes helpers into practical output.
func main() {
	items := []Expense{
		{Category: " Food ", Amount: 20.5},
		{Category: "food", Amount: 15.0},
		{Category: "transport", Amount: 12.25},
	}

	RunReport(items)

	filtered := FilterByMin(items, 15)
	spent := SumAmounts(filtered)
	fmt.Println("Lesson 10 filtered spent:", spent)
	fmt.Println("Lesson 10 status:", BudgetStatus(40, spent))
}

// End of Go Projects 1-10
