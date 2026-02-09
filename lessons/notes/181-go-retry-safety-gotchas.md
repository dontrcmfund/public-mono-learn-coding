# Go retry-safety gotchas

Goal: avoid subtle bugs when clients retry failed requests.

Why do we care?
- "At least once" delivery and retry logic are common in real systems
- Poor retry handling causes duplicate writes and inconsistent state

Common gotchas
- No idempotency key on create endpoints
- Applying rate limits but not returning clear error semantics
- Retrying non-idempotent actions blindly
- Treating timeouts as guaranteed failures (operation may still have succeeded)
- Storing idempotency keys without expiration strategy

Status code guidance
- `429 Too Many Requests`: caller should back off
- `409 Conflict`: semantic conflict in request state
- `503 Service Unavailable`: temporary failure, retry may help

Rule of thumb
- Combine: idempotency + clear status codes + retry/backoff guidance

If all you remember is one thing
- Reliable APIs are designed for retries, not surprised by retries
