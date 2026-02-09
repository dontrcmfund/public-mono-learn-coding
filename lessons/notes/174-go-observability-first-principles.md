# Go observability (first principles)

Goal: understand how logs and metrics help you operate software confidently.

Why do we care?
- Bugs in production are inevitable
- Without visibility, you cannot explain failures quickly

History context
- Early systems relied on ad-hoc print debugging
- Modern services adopted structured logging and metrics for faster incident response
- Reliable operations now require observability from day one

Core ideas
- Logs answer: "what happened?"
- Metrics answer: "how often/how long?"
- Traces answer: "where in the flow did it slow or fail?"

Beginner-friendly starting point
- Use structured logs with stable keys
- Log request method/path/status and error context
- Track basic request count and latency

If all you remember is one thing
- Observability turns unknown failures into diagnosable problems
