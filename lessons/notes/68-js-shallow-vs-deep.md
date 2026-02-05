# JS shallow vs deep copy (gotchas)

Goal: understand why some “copies” still change together.

Why do we care?
- Beginners often expect a full copy
- Bugs happen when nested data stays linked

First principles
- A shallow copy copies the top level only
- Nested objects still point to the same inner object

Example
- `const a = { info: { level: 1 } }`
- `const b = { ...a }`
- Changing `b.info.level` also changes `a.info.level`

Rule of thumb
- Shallow copies are fine for flat data
- For nested data, you need a deep copy (more advanced)

If all you remember is one thing
- Spread makes a shallow copy, not a deep one
