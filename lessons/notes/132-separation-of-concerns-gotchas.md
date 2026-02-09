# Separation of concerns gotchas

Goal: avoid design mistakes that create tightly coupled code.

Why do we care?
- Tight coupling makes testing and changes expensive
- Mixed responsibilities hide bugs and intent

Common gotchas
- Business rules mixed with IO operations
- Database queries scattered across unrelated modules
- Validation logic duplicated in many places
- Functions that both compute and print/save data

Rule of thumb
- One function should have one reason to change
- Keep pure logic separate from side effects

If all you remember is one thing
- Separate compute logic from IO logic to keep code flexible
