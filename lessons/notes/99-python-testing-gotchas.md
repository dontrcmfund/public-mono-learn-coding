# Python testing gotchas

Goal: avoid test patterns that create false confidence.

Why do we care?
- Poor tests can pass while real behavior is still broken
- Flaky tests slow learning and damage trust

Common gotchas
- Tests depend on external state (network, clock, file leftovers)
- One test checks too many behaviors at once
- Assertions are too vague to diagnose failures
- Shared mutable test data leaks across cases

Rule of thumb
- Keep tests isolated and deterministic
- Prefer many small focused assertions
- Use temporary files/directories for file tests

If all you remember is one thing
- Reliable tests isolate behavior and make failures obvious
