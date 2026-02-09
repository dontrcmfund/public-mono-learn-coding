# Go profiling (first principles)

Goal: understand how profiling reveals where time and memory are actually spent.

Why do we care?
- Intuition about bottlenecks is often wrong
- Profilers expose hot paths and allocation pressure

Core ideas
- CPU profile: where execution time is spent
- Memory/alloc profile: where allocations happen
- Blocking profile: where goroutines wait

Go relevance
- Go provides built-in tooling (`pprof`) for profiling
- Profiling integrates with real workloads and tests

Rule of thumb
- Profile representative workloads
- Fix the dominant hotspot first
- Verify improvements with repeated measurements

If all you remember is one thing
- Profiling turns performance debugging into concrete evidence
