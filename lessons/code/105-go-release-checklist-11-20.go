package main

import (
	"fmt"
	"strings"
)

/*
GO RELEASE CHECKLIST (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/105-go-release-checklist-11-20.go
2) Change checklist values and observe go/no-go output

Extra context:
- lessons/notes/192-go-release-versioning-first-principles.md
- lessons/notes/193-deployment-safety-gotchas.md
*/

// LESSON 11: Release input model
// Why this matters: release decisions need explicit, reviewable criteria.
type ReleaseCandidate struct {
	Version            string
	Branch             string
	ChecksPassed       bool
	MigrationReviewed  bool
	RollbackPlanReady  bool
	PostDeployChecks   bool
}

// LESSON 12: Simple semantic version validator
// Why this matters: malformed versions confuse consumers and tooling.
func isSemverLike(v string) bool {
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return false
	}
	for _, p := range parts {
		if p == "" {
			return false
		}
		for _, r := range p {
			if r < '0' || r > '9' {
				return false
			}
		}
	}
	return true
}

// LESSON 13: Branch policy check
// Why this matters: protect production releases from unreviewed branches.
func branchAllowed(branch string) bool {
	return branch == "main"
}

// LESSON 14: Release gate decision
// Why this matters: centralized decision avoids ad-hoc risky overrides.
func canRelease(c ReleaseCandidate) (bool, []string) {
	reasons := []string{}
	if !isSemverLike(c.Version) {
		reasons = append(reasons, "invalid version format")
	}
	if !branchAllowed(c.Branch) {
		reasons = append(reasons, "release branch must be main")
	}
	if !c.ChecksPassed {
		reasons = append(reasons, "pipeline checks failed")
	}
	if !c.MigrationReviewed {
		reasons = append(reasons, "migration review missing")
	}
	if !c.RollbackPlanReady {
		reasons = append(reasons, "rollback plan missing")
	}
	if !c.PostDeployChecks {
		reasons = append(reasons, "post-deploy checks not defined")
	}
	return len(reasons) == 0, reasons
}

// LESSON 15: Checklist printer
// Why this matters: humans need clear release readiness feedback.
func printDecision(c ReleaseCandidate) {
	ok, reasons := canRelease(c)
	fmt.Printf("version=%s branch=%s\n", c.Version, c.Branch)
	if ok {
		fmt.Println("release_decision=APPROVED")
		return
	}
	fmt.Println("release_decision=BLOCKED")
	for _, r := range reasons {
		fmt.Println("- reason:", r)
	}
}

// LESSON 16: Safe defaults
// Why this matters: default-deny is safer than default-allow.

// LESSON 17: Version as compatibility signal
// Why this matters: clients rely on version meaning, not just label.

// LESSON 18: Rollback planning habit
// Why this matters: incident recovery speed depends on preparation.

// LESSON 19: Post-deploy verification habit
// Why this matters: successful deploy command is not successful user outcome.

// LESSON 20: End-to-end release checklist demo
// Why this matters: combines policy, quality gates, and safety checks.
func main() {
	candidate := ReleaseCandidate{
		Version:           "1.4.2",
		Branch:            "main",
		ChecksPassed:      true,
		MigrationReviewed: true,
		RollbackPlanReady: true,
		PostDeployChecks:  true,
	}
	printDecision(candidate)
}

// End of Go Release Checklist 11-20
