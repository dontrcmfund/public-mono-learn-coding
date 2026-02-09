# Async jobs (first principles)

Goal: understand why some work should happen outside the request/response path.

Why do we care?
- Some tasks are slow (emails, reports, webhooks, file processing)
- Blocking HTTP requests for slow work hurts user experience and reliability

Past, present, future value
- Past: synchronous-only systems timed out under load
- Present: async jobs keep APIs fast and responsive
- Future: job systems enable scalable pipelines and automation

Core idea
- API accepts request quickly and enqueues background work
- Worker processes the job independently
- Client can poll status or receive callback later

Rule of thumb
- Keep request path small
- Move long-running/fragile tasks to workers

If all you remember is one thing
- Async jobs separate user responsiveness from heavy processing time
