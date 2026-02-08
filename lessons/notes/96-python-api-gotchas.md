# Python API gotchas

Goal: avoid common mistakes when scripts interact with real APIs.

Why do we care?
- API failures are normal, not exceptional
- Small mistakes can cause silent bad data

Common gotchas
- Assuming every response is JSON
- Ignoring non-200 status codes
- Forgetting pagination and reading only first page
- Not setting timeout values
- Retrying too aggressively and triggering rate limits

Rule of thumb
- Check status, content type, and required fields
- Use bounded retries with backoff
- Keep request and parsing logic separate

If all you remember is one thing
- Production-safe API scripts treat failures as expected states
