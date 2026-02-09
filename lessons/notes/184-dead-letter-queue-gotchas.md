# Dead-letter queue gotchas

Goal: avoid silent job loss when retries keep failing.

Why do we care?
- Some jobs never succeed due to bad payloads or permanent errors
- Without dead-letter handling, failures disappear and users trust breaks

Core idea
- After max retries, move job to dead-letter queue (DLQ)
- DLQ records failure reason for triage and replay decisions

Common gotchas
- Infinite retries with no upper bound
- No structured reason for DLQ entries
- No monitoring on DLQ growth
- Mixing retryable and permanent errors without classification

Rule of thumb
- Retries for transient failures
- DLQ for exhausted or non-retryable failures

If all you remember is one thing
- DLQ is a safety net that turns silent loss into visible work
