package main

import (
	"fmt"
	"strings"
	"time"
)

/*
GO CI/CD BASICS (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/104-go-cicd-basics-1-10.go
2) Toggle check results and observe release decision changes

Extra context:
- lessons/notes/191-cicd-first-principles.md
- lessons/notes/193-deployment-safety-gotchas.md
*/

// LESSON 1: Pipeline check model
// Why this matters: explicit checks make quality gates transparent.
type CheckResult struct {
	Name   string
	Passed bool
	Detail string
}

// LESSON 2: Pipeline run summary
// Why this matters: one object captures release readiness state.
type PipelineRun struct {
	CommitSHA string
	Branch    string
	Checks    []CheckResult
	StartedAt string
}

// LESSON 3: Add check helper
// Why this matters: consistent check recording reduces ad-hoc logic.
func (p *PipelineRun) AddCheck(name string, passed bool, detail string) {
	p.Checks = append(p.Checks, CheckResult{Name: name, Passed: passed, Detail: detail})
}

// LESSON 4: Gate evaluation
// Why this matters: release decisions should be deterministic and explainable.
func (p PipelineRun) AllCriticalChecksPassed() bool {
	for _, c := range p.Checks {
		if !c.Passed {
			return false
		}
	}
	return true
}

// LESSON 5: Build step simulation
// Why this matters: compile/build failures should stop delivery early.
func runBuildStep() CheckResult {
	return CheckResult{Name: "build", Passed: true, Detail: "binary compiled"}
}

// LESSON 6: Test step simulation
// Why this matters: tests protect expected behavior from regressions.
func runTestStep() CheckResult {
	return CheckResult{Name: "tests", Passed: true, Detail: "all tests passed"}
}

// LESSON 7: Lint step simulation
// Why this matters: style and static issues are cheaper to fix pre-release.
func runLintStep() CheckResult {
	return CheckResult{Name: "lint", Passed: true, Detail: "no lint violations"}
}

// LESSON 8: Security scan step simulation
// Why this matters: catch risky dependencies and patterns before deploy.
func runSecurityScanStep() CheckResult {
	return CheckResult{Name: "security", Passed: true, Detail: "no critical findings"}
}

// LESSON 9: Decision + explanation helper
// Why this matters: developers need clear reasons for blocked releases.
func summarize(run PipelineRun) string {
	lines := []string{
		fmt.Sprintf("commit=%s branch=%s started_at=%s", run.CommitSHA, run.Branch, run.StartedAt),
	}
	for _, c := range run.Checks {
		state := "PASS"
		if !c.Passed {
			state = "FAIL"
		}
		lines = append(lines, fmt.Sprintf("- %s: %s (%s)", c.Name, state, c.Detail))
	}
	if run.AllCriticalChecksPassed() {
		lines = append(lines, "release_decision=APPROVED")
	} else {
		lines = append(lines, "release_decision=BLOCKED")
	}
	return strings.Join(lines, "\n")
}

// LESSON 10: End-to-end pipeline simulation
// Why this matters: demonstrates how CI gates control CD safety.
func main() {
	run := PipelineRun{
		CommitSHA: "abc1234",
		Branch:    "main",
		StartedAt: time.Now().Format(time.RFC3339),
	}

	for _, step := range []func() CheckResult{
		runBuildStep,
		runTestStep,
		runLintStep,
		runSecurityScanStep,
	} {
		check := step()
		run.AddCheck(check.Name, check.Passed, check.Detail)
	}

	fmt.Println(summarize(run))
}

// End of Go CI/CD Basics 1-10
