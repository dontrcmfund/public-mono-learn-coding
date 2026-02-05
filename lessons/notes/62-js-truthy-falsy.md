# JS truthy and falsy (gotchas)

Goal: understand why some values act like true or false.

Why do we care?
- Conditions often use non‑boolean values
- This can make code behave unexpectedly

First principles
- JavaScript treats some values as “falsy” in `if` statements
- Everything else is “truthy”

Common falsy values
- `false`
- `0`
- `""` (empty string)
- `null`
- `undefined`
- `NaN`

Why this matters in real life
- An empty string might skip a block of code
- A `0` might be mistaken for “no value”

Rule of thumb
- When in doubt, compare explicitly

If all you remember is one thing
- Empty values can act like false in conditions
