# Go gotchas

Goal: avoid common early mistakes in Go.

Why do we care?
- Small syntax and semantics surprises can slow learning

Common gotchas
- `:=` declares a new variable; `=` assigns existing variable
- Unused variables and imports fail compilation
- `nil` handling requires explicit checks
- Loop variable capture can cause unexpected behavior in closures
- Errors are values and must be handled directly

Rule of thumb
- Read compiler messages carefully; they are usually precise and helpful

If all you remember is one thing
- In Go, explicit code and explicit handling are the safe defaults
