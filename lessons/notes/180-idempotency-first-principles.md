# Idempotency (first principles)

Goal: understand why repeating the same request should not create duplicate side effects.

Why do we care?
- Networks fail and clients retry
- Retries should be safe, not duplicate orders/payments/tasks

History context
- Distributed systems made retries normal, not exceptional
- Idempotency keys emerged to deduplicate repeated create requests

Core idea
- Client sends an idempotency key for mutation requests
- Server stores first successful result for that key
- Repeated requests with same key return the same result

Rule of thumb
- Use idempotency on create/charge/order endpoints
- Define key scope clearly (per user/per endpoint/per time window)
- Persist keys in real systems; in-memory is for learning

If all you remember is one thing
- Idempotency makes retries safe by turning duplicates into one effect
