# Go testing principles

Goal: use tests to validate behavior and protect refactoring.

Why do we care?
- Go tooling makes tests fast and easy to run
- Tests document expected behavior as executable examples

First principles
- Keep tests deterministic and focused
- Test exported behavior, not implementation details
- Use table-driven tests for multiple similar cases

Go workflow
- Run tests with `go test`
- In this repo's lesson layout, run file-scoped tests like `go test lessons/code/67-go-testing-1-10_test.go -v`
- Keep logic in small functions for easy testing
- Prefer clear failure messages over clever assertions

If all you remember is one thing
- In Go, tests are a normal part of writing maintainable code
