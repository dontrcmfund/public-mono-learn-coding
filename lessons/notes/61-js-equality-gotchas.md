# JS equality gotchas (== vs ===)

Goal: understand why strict equality is safer.

Why do we care?
- `==` can convert types in ways you do not expect
- Bugs from loose equality are common for beginners

First principles
- `===` compares value AND type
- `==` tries to make the values the same type first

Simple examples
- `5 === "5"` is false (number vs string)
- `5 == "5"` is true (string becomes number)

Why this matters in real life
- A user types “5” into a form (string)
- You compare it to 5 (number)
- `==` says true even though types differ

Rule of thumb
- Use `===` unless you have a very good reason

If all you remember is one thing
- `===` is predictable; `==` can surprise you
