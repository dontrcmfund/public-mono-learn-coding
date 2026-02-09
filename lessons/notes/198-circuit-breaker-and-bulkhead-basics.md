# Circuit breaker and bulkhead basics

Goal: understand two core patterns that prevent cascading failures.

Why do we care?
- When a dependency fails repeatedly, repeated calls waste resources
- Shared worker pools can be exhausted by one problematic path

Circuit breaker
- Closed: calls flow normally
- Open: calls fail fast for cooldown period
- Half-open: allow limited probe to check recovery

Bulkhead
- Reserve separate capacity for different work categories
- Prevent one noisy path from starving others

Rule of thumb
- Use circuit breakers for unstable dependencies
- Use bulkheads where shared resources can saturate

If all you remember is one thing
- Circuit breakers fail fast; bulkheads contain failure impact
