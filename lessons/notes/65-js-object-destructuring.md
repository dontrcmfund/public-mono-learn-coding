# JS object destructuring (gotchas)

Goal: understand what destructuring is doing under the hood.

Why do we care?
- Destructuring is a shortcut, but it can feel magical
- Understanding it prevents confusion

First principles
- Destructuring copies values from an object into new variables
- The variable names must match the object’s keys

Example
- `const user = { name: "Kai", role: "admin" }`
- `const { name, role } = user` creates two new variables

Common confusion
- If the key does not exist, the variable becomes `undefined`
- The original object does not change

Rule of thumb
- Use destructuring when it makes code clearer
- Avoid it if it feels confusing — clarity comes first

If all you remember is one thing
- Destructuring is just a shortcut for `const name = user.name`
