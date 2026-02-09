# Contract testing gotchas

Goal: avoid false confidence when testing API contracts.

Why do we care?
- Passing unit tests can still hide integration contract breaks
- Consumers fail when response shapes drift unexpectedly

Common gotchas
- Testing status code only, not response schema
- Accepting extra/missing fields silently
- Inconsistent error response format across endpoints
- Forgetting backward compatibility checks

Rule of thumb
- Test happy path and error path contracts
- Validate key fields and types in responses
- Treat contract failures as release blockers

If all you remember is one thing
- Contract tests protect consumers from accidental API drift
