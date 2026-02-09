# Performance (first principles)

Goal: understand why performance work starts with measurement, not guessing.

Why do we care?
- Slow systems hurt user trust and operational cost
- Premature optimization can add complexity without real benefit

Core ideas
- Latency: how long one request takes
- Throughput: how many requests/work items per time unit
- Resource use: CPU, memory, IO, network

Rule of thumb
- Measure baseline first
- Optimize the largest bottleneck
- Re-measure after each change

If all you remember is one thing
- Good performance engineering is evidence-driven iteration
