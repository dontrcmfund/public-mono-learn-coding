# How to run these Go lessons in this repo

Goal: remove command confusion so you can focus on concepts.

Why this matters
- Go usually expects one package per folder
- This curriculum keeps many standalone lesson files in one folder to reduce context switching
- That layout is easier for scrolling and reading, but it changes how you run commands

How to run code lesson files
- Use file-specific run commands like:
- `go run lessons/code/65-go-basics-1-10.go`
- `go run lessons/code/69-go-http-api-basics-1-10.go`

How to run test lesson files
- Use file-specific test commands like:
- `go test lessons/code/67-go-testing-1-10_test.go -v`
- `go test lessons/code/73-go-capstone-tests-1-10_test.go -v`

First-principles explanation
- Tools have assumptions
- When we intentionally choose a different learning layout, we adjust commands to match
- This is normal software engineering: understand assumptions, then adapt safely

If all you remember is one thing
- In this course, use file-specific `go run` and `go test` commands
