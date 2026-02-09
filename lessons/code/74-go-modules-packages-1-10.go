package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
GO MODULES + PACKAGES (Lessons 1-10)

Suggested use:
1) Run: go run lessons/code/74-go-modules-packages-1-10.go
2) Read each lesson comment and map it to code output

Extra context:
- lessons/notes/161-go-how-to-run-these-lessons.md
- lessons/notes/162-go-modules-and-packages-first-principles.md
*/

// LESSON 1: Package purpose
// Why this matters: packages create clear boundaries between concerns.

// LESSON 2: Module purpose
// Why this matters: modules version a whole project for reproducible builds.

// LESSON 3: Import purpose
// Why this matters: explicit imports show real dependencies.

// LESSON 4: Naming clarity
// Why this matters: consistent names reduce cognitive load.
func normalizePackageName(raw string) string {
	clean := strings.TrimSpace(strings.ToLower(raw))
	return strings.ReplaceAll(clean, "-", "_")
}

// LESSON 5: Keep rules in small helpers
// Why this matters: small units are easier to test and change.
func isValidPackageName(name string) bool {
	if name == "" {
		return false
	}
	for _, r := range name {
		if !(r == '_' || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

// LESSON 6: Build canonical path
// Why this matters: predictable paths prevent import confusion.
func packagePath(module string, pkg string) string {
	return filepath.ToSlash(filepath.Join(module, pkg))
}

// LESSON 7: Group metadata in structs
// Why this matters: structs model real project concepts.
type PackageInfo struct {
	ModulePath  string
	PackageName string
}

// LESSON 8: Methods for behavior
// Why this matters: behavior near data improves readability.
func (p PackageInfo) ImportPath() string {
	return packagePath(p.ModulePath, p.PackageName)
}

// LESSON 9: Validate before use
// Why this matters: fail fast with clear messages.
func buildPackageInfo(modulePath string, rawPackageName string) (PackageInfo, error) {
	normalized := normalizePackageName(rawPackageName)
	if !isValidPackageName(normalized) {
		return PackageInfo{}, fmt.Errorf("invalid package name: %q", rawPackageName)
	}
	return PackageInfo{
		ModulePath:  strings.TrimSpace(modulePath),
		PackageName: normalized,
	}, nil
}

// LESSON 10: End-to-end example
// Why this matters: composition shows how small rules become workflow.
func main() {
	fmt.Println("Lesson 1/2: packages organize code, modules version projects")
	fmt.Println("Lesson 3: imports make dependencies explicit")

	raw := "  Billing-API  "
	normalized := normalizePackageName(raw)
	fmt.Println("Lesson 4:", normalized)

	fmt.Println("Lesson 5 valid?", isValidPackageName(normalized))
	fmt.Println("Lesson 6 path:", packagePath("github.com/acme/learning", normalized))

	info, err := buildPackageInfo("github.com/acme/learning", raw)
	if err != nil {
		fmt.Println("Lesson 9 error:", err)
		return
	}

	fmt.Printf("Lesson 7/8: %+v\n", info)
	fmt.Println("Lesson 8 import path:", info.ImportPath())
	fmt.Println("Lesson 10: module+package model complete")
}

// End of Go Modules + Packages 1-10
