# JS JSON and try/catch (gotchas)

Goal: understand why JSON parsing fails and how try/catch helps.

Why do we care?
- JSON must be valid text, or parsing will crash
- Real apps receive messy data

First principles
- JSON is just a string format for data
- `JSON.parse` throws an error if the string is invalid
- `try/catch` lets you handle that error safely

Common mistakes
- Single quotes (`'`) are not valid in JSON
- Missing quotes around keys
- Trailing commas

Rule of thumb
- Always wrap `JSON.parse` in `try/catch` when input is uncertain

If all you remember is one thing
- `JSON.parse` can crash, so protect it with `try/catch`
