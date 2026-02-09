package main

import (
	"fmt"
	"strings"
	"time"
)

/*
GO PROFILING-FRIENDLY CODE (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/108-go-profiling-friendly-11-20.go
2) Inspect timing per stage to identify dominant hotspot

Extra context:
- lessons/notes/195-go-profiling-first-principles.md
*/

// LESSON 11: Pipeline stage type
// Why this matters: named stages make hotspots easier to attribute.
type StageResult struct {
	Name      string
	Duration  time.Duration
	WorkUnits int
}

// LESSON 12: Stage timer helper
// Why this matters: stage-level timing gives actionable profiling clues.
func timeStage(name string, fn func() int) StageResult {
	start := time.Now()
	units := fn()
	return StageResult{Name: name, Duration: time.Since(start), WorkUnits: units}
}

// LESSON 13: Parse-like stage
// Why this matters: demonstrates string processing cost.
func parseStage(input []string) int {
	count := 0
	for _, s := range input {
		_ = strings.TrimSpace(strings.ToLower(s))
		count++
	}
	return count
}

// LESSON 14: Validate-like stage
// Why this matters: condition-heavy logic may dominate in some workloads.
func validateStage(input []string) int {
	valid := 0
	for _, s := range input {
		trim := strings.TrimSpace(s)
		if trim != "" && strings.Contains(trim, "@") {
			valid++
		}
	}
	return valid
}

// LESSON 15: Simulated storage stage
// Why this matters: IO-like waits often shape end-to-end latency.
func persistStage(items int) int {
	time.Sleep(15 * time.Millisecond)
	return items
}

// LESSON 16: Pipeline execution
// Why this matters: profile should mirror real code flow.
func runPipeline(data []string) []StageResult {
	results := []StageResult{}
	results = append(results, timeStage("parse", func() int { return parseStage(data) }))
	results = append(results, timeStage("validate", func() int { return validateStage(data) }))
	results = append(results, timeStage("persist", func() int { return persistStage(len(data)) }))
	return results
}

// LESSON 17: Hotspot finder
// Why this matters: optimization should target dominant stage first.
func maxStage(results []StageResult) StageResult {
	if len(results) == 0 {
		return StageResult{}
	}
	max := results[0]
	for _, r := range results[1:] {
		if r.Duration > max.Duration {
			max = r
		}
	}
	return max
}

// LESSON 18: Reporting helper
// Why this matters: concise output supports fast diagnosis.
func printResults(results []StageResult) {
	total := time.Duration(0)
	for _, r := range results {
		total += r.Duration
		fmt.Printf("stage=%s duration=%s work_units=%d\n", r.Name, r.Duration, r.WorkUnits)
	}
	hot := maxStage(results)
	fmt.Printf("total=%s hotspot=%s hotspot_duration=%s\n", total, hot.Name, hot.Duration)
}

// LESSON 19: Repeatability note
// Why this matters: repeated runs reduce noise-driven decisions.

// LESSON 20: End-to-end stage timing demo
// Why this matters: profiling begins with structured measurement points.
func main() {
	data := []string{
		" mia@example.com ",
		"leo@example.com",
		" bad-email ",
		" ana@example.com ",
	}
	results := runPipeline(data)
	printResults(results)
}

// End of Go Profiling-Friendly Code 11-20
