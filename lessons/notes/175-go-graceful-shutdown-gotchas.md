# Go graceful shutdown gotchas

Goal: avoid data loss and broken requests when stopping services.

Why do we care?
- Services are restarted for deploys, scaling, and failures
- Hard stops can drop in-flight requests and corrupt work

Common gotchas
- Ignoring OS signals (`SIGINT`, `SIGTERM`)
- Calling `os.Exit` without cleanup
- Closing dependencies before request handling stops
- Missing shutdown timeout
- Not propagating `context` cancellation

First-principles pattern
- Stop accepting new requests
- Let in-flight requests finish (within timeout)
- Close background workers and resources cleanly

Rule of thumb
- Treat shutdown as a normal execution path, not an edge case

If all you remember is one thing
- Graceful shutdown protects correctness during normal operational events
