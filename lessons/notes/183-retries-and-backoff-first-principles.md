# Retries and backoff (first principles)

Goal: understand how to recover from temporary failures without making outages worse.

Why do we care?
- Network calls and dependencies fail transiently
- Immediate repeated retries can amplify overload

Core idea
- Retry only retryable failures
- Wait between attempts (backoff)
- Stop after max attempts

Backoff pattern
- Attempt 1 fails -> short delay
- Attempt 2 fails -> longer delay
- Attempt 3 fails -> longer delay or give up

Rule of thumb
- Limit retries
- Add jitter in production systems
- Log attempts for diagnosis

If all you remember is one thing
- Retries help only when paced and bounded
