# Go projects principles

Goal: structure small Go projects so they can scale cleanly.

Why do we care?
- Early structure prevents tangled code later
- Clear package boundaries improve readability and testing

First principles
- `main` wires dependencies and starts execution
- Core logic lives in reusable functions/types
- IO and side effects stay near the edges

Rule of thumb
- Keep functions small and explicit
- Return errors instead of hiding them
- Separate parsing, logic, and output

If all you remember is one thing
- A good Go project keeps wiring, logic, and side effects clearly separated
