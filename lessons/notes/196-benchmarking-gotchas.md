# Benchmarking gotchas

Goal: avoid misleading benchmark results.

Why do we care?
- Bad benchmarks can push teams toward worse code
- Reproducibility matters for trust in optimization decisions

Common gotchas
- Benchmarking unrealistic tiny inputs only
- Including setup/IO noise inside hot loop
- Ignoring variance across runs
- Comparing unoptimized and optimized code with different behavior

Rule of thumb
- Benchmark realistic workloads
- Keep benchmark loops focused on target code
- Compare correctness and performance together

If all you remember is one thing
- A benchmark is useful only when it measures the right work reliably
