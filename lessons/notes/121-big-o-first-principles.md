# Big-O first principles

Goal: understand Big-O as a growth model, not a scary formula.

Why do we care?
- Input sizes grow in real systems
- Some solutions scale smoothly; others degrade quickly

First principles
- Big-O describes how runtime or memory grows with input size `n`
- It compares growth patterns, not exact milliseconds
- Lower growth usually means better scalability

Common classes
- `O(1)` constant
- `O(log n)` logarithmic
- `O(n)` linear
- `O(n log n)` efficient sorting class
- `O(n^2)` quadratic (often slow at scale)

Rule of thumb
- First, write a correct solution
- Then improve structure/algorithm where growth is poor

If all you remember is one thing
- Big-O helps predict scaling behavior before production pain appears
