# JS NaN and typeof quirks (gotchas)

Goal: understand common confusion around `NaN` and `typeof`.

Why do we care?
- `NaN` shows up when math fails
- `typeof` can surprise beginners

First principles
- `NaN` means “Not a Number”
- It is still a number type
- `typeof NaN` is "number"

Simple example
- `Number("abc")` gives `NaN`
- `Number.isNaN(value)` checks for it safely

Why this matters in real life
- User input might not be a number
- You need a way to detect bad input

Rule of thumb
- Use `Number.isNaN()` instead of `value === NaN` (that never works)

If all you remember is one thing
- `NaN` is a number type that signals “bad math”
