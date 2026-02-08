# TypeScript tooling and runtime flow

Goal: understand what happens between writing `.ts` and running code.

Why do we care?
- Many beginners confuse TypeScript checks with runtime execution
- Knowing the flow removes setup confusion

First principles
- You write `.ts` source files
- TypeScript compiler checks and transforms them
- Runtime executes resulting JavaScript

Short history (why this exists)
- JavaScript runtimes execute JavaScript, not TypeScript syntax
- TypeScript kept compatibility by compiling to standard JavaScript
- This allowed gradual adoption in existing JS ecosystems

Etymology
- `compile` means transform source into executable target form
- `transpile` often means compile between high-level languages (TS to JS)

Common confusion
- Node usually cannot run raw `.ts` directly without extra tooling
- Type checking can pass while runtime logic can still fail

Rule of thumb
- Treat TypeScript as design-time safety
- Still test runtime behavior with real inputs

If all you remember is one thing
- TypeScript improves development safety, then JavaScript runs in production
