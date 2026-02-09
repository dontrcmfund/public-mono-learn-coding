# Go errors and context (first principles)

Goal: learn why Go uses explicit errors and `context.Context`.

Why do we care?
- Most production failures are timeout, cancellation, or dependency failures
- Good error and context handling makes systems observable and resilient

History context
- Go avoided exception-heavy control flow to keep behavior explicit
- As cloud services grew, cancellation and timeouts became essential
- `context` became standard so services can stop work when requests end

Core ideas
- Errors are values: return and handle them directly
- Wrap errors with extra meaning so future debugging is faster
- Use context for deadlines and cancellation signals

Rule of thumb
- Every external call should be cancellable
- Every returned error should carry useful context

If all you remember is one thing
- Explicit errors explain failures; context controls work lifetime
