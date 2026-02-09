# Python DSA gotchas

Goal: avoid common errors when implementing algorithms under time pressure.

Why do we care?
- Most algorithm bugs are off-by-one, boundary, or mutation mistakes
- Catching these patterns early saves hours of debugging

Common gotchas
- Off-by-one loop boundaries (`<` vs `<=`)
- Not handling empty inputs
- Mutating input list when caller expects original unchanged
- Forgetting to sort before using binary search or two-pointer patterns
- Confusing stack vs queue behavior

Rule of thumb
- Write edge-case tests first (empty, one item, duplicates)
- Document assumptions (sorted input, non-empty input, unique values)
- Prefer clear variable names (`left`, `right`, `mid`)

If all you remember is one thing
- Correct boundary handling is usually the difference between pass and fail
