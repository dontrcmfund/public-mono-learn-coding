# TypeScript small projects gotchas

Goal: avoid the mistakes that make beginner projects feel harder than they are.

Why do we care?
- Most project bugs come from state handling and unclear data contracts
- Catching these patterns early speeds up future project work

Common gotchas
- Mutating arrays/objects by accident instead of returning new values
- Using loose IDs or duplicate IDs
- Mixing validation and business logic in one large function
- Assuming optional fields always exist

Rule of thumb
- Keep functions small and single-purpose
- Validate input at boundaries
- Return typed results for success/failure states

If all you remember is one thing
- Most bugs shrink when data shapes are explicit and updates are predictable
