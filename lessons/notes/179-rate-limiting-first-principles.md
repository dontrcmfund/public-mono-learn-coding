# Rate limiting (first principles)

Goal: understand why APIs must control request volume.

Why do we care?
- Without limits, one client can overwhelm shared resources
- Spikes and abuse can make service unavailable for everyone

Past, present, future value
- Past: unbounded APIs caused outages during traffic bursts
- Present: limits protect reliability and fairness
- Future: limits enable predictable scaling and cost control

Core idea
- Set a maximum request allowance per client in a time window
- Return `429 Too Many Requests` when limit is exceeded

Rule of thumb
- Start simple and explicit
- Make limits observable (logs/metrics)
- Communicate retry guidance with headers when possible

If all you remember is one thing
- Rate limiting protects system stability and user fairness
