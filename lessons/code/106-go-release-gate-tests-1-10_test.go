package main

import (
	"strings"
	"testing"
	"time"
)

/*
GO RELEASE GATE TESTS (Lessons 1-10)

Suggested use:
1) Run: go test lessons/code/106-go-release-gate-tests-1-10_test.go -run TestLesson -v
2) Focus on why releases are blocked, not only pass/fail
*/

type CheckResult struct {
	Name   string
	Passed bool
	Detail string
}

type PipelineRun struct {
	CommitSHA string
	Branch    string
	Checks    []CheckResult
	StartedAt string
}

func (p *PipelineRun) AddCheck(name string, passed bool, detail string) {
	p.Checks = append(p.Checks, CheckResult{Name: name, Passed: passed, Detail: detail})
}

func (p PipelineRun) AllCriticalChecksPassed() bool {
	for _, c := range p.Checks {
		if !c.Passed {
			return false
		}
	}
	return true
}

type ReleaseCandidate struct {
	Version           string
	Branch            string
	ChecksPassed      bool
	MigrationReviewed bool
	RollbackPlanReady bool
	PostDeployChecks  bool
}

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

func branchAllowed(branch string) bool {
	return branch == "main"
}

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

func TestLesson1AllChecksPassedIsTrue(t *testing.T) {
	run := PipelineRun{CommitSHA: "abc", Branch: "main", StartedAt: time.Now().Format(time.RFC3339)}
	run.AddCheck("build", true, "ok")
	run.AddCheck("test", true, "ok")
	if !run.AllCriticalChecksPassed() {
		t.Fatalf("expected all checks passed")
	}
}

func TestLesson2AnyFailBlocksPipeline(t *testing.T) {
	run := PipelineRun{}
	run.AddCheck("build", true, "ok")
	run.AddCheck("test", false, "failing test")
	if run.AllCriticalChecksPassed() {
		t.Fatalf("expected failed pipeline gate")
	}
}

func TestLesson3SemverValidatorAcceptsBasicVersion(t *testing.T) {
	if !isSemverLike("1.2.3") {
		t.Fatalf("expected valid semver-like version")
	}
}

func TestLesson4SemverValidatorRejectsInvalidVersion(t *testing.T) {
	if isSemverLike("1.2") {
		t.Fatalf("expected invalid version")
	}
}

func TestLesson5BranchPolicyAllowsMain(t *testing.T) {
	if !branchAllowed("main") {
		t.Fatalf("main should be allowed")
	}
}

func TestLesson6BranchPolicyBlocksFeatureBranch(t *testing.T) {
	if branchAllowed("feature/x") {
		t.Fatalf("feature branch should be blocked")
	}
}

func TestLesson7CanReleaseTrueWhenAllSafetyChecksPass(t *testing.T) {
	ok, reasons := canRelease(ReleaseCandidate{
		Version:           "1.0.0",
		Branch:            "main",
		ChecksPassed:      true,
		MigrationReviewed: true,
		RollbackPlanReady: true,
		PostDeployChecks:  true,
	})
	if !ok || len(reasons) != 0 {
		t.Fatalf("expected releasable candidate")
	}
}

func TestLesson8CanReleaseFalseOnMissingRollback(t *testing.T) {
	ok, reasons := canRelease(ReleaseCandidate{
		Version:           "1.0.0",
		Branch:            "main",
		ChecksPassed:      true,
		MigrationReviewed: true,
		RollbackPlanReady: false,
		PostDeployChecks:  true,
	})
	if ok {
		t.Fatalf("expected blocked release")
	}
	found := false
	for _, r := range reasons {
		if r == "rollback plan missing" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected rollback reason in block list")
	}
}

func TestLesson9CanReleaseReturnsMultipleReasons(t *testing.T) {
	ok, reasons := canRelease(ReleaseCandidate{
		Version:           "1.x.0",
		Branch:            "feature/x",
		ChecksPassed:      false,
		MigrationReviewed: false,
		RollbackPlanReady: false,
		PostDeployChecks:  false,
	})
	if ok {
		t.Fatalf("expected blocked release")
	}
	if len(reasons) < 3 {
		t.Fatalf("expected multiple blocking reasons, got %d", len(reasons))
	}
}

func TestLesson10Completion(t *testing.T) {
	if false {
		t.Fatalf("unreachable")
	}
}

// End of Go Release Gate Tests 1-10
