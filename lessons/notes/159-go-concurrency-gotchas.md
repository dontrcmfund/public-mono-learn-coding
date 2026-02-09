# Go concurrency gotchas

Goal: avoid common pitfalls when starting goroutines and channels.

Why do we care?
- Concurrency bugs are subtle and expensive

Common gotchas
- Goroutine leaks from blocked sends/receives
- Closing channels from the wrong side
- Forgetting to drain or stop workers on shutdown
- Capturing loop variables incorrectly in goroutines
- Missing timeouts for external calls

Rule of thumb
- Use `context.Context` for cancellation
- Close channels only from producer side
- Keep goroutine lifetimes explicit

If all you remember is one thing
- Most Go concurrency bugs are lifecycle bugs, not syntax bugs
