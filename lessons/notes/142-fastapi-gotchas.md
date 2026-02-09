# FastAPI gotchas

Goal: avoid common pitfalls when building FastAPI services.

Why do we care?
- Small mistakes in API layers can cause unstable behavior

Common gotchas
- Putting business logic directly in route handlers
- Mixing sync and async code incorrectly
- Returning inconsistent error shapes
- Skipping request/response models
- Using global mutable state without care

Rule of thumb
- Validate at the boundary (Pydantic models)
- Keep service logic framework-agnostic
- Handle errors with clear HTTP status + message contracts

If all you remember is one thing
- Keep handlers thin and contracts explicit
