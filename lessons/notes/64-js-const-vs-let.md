# JS const vs let (gotchas)

Goal: understand when to use `const` and `let`.

Why do we care?
- `const` prevents accidental reassignment
- It makes your code safer and easier to read

First principles
- `const` means “this name will not be reassigned”
- `let` means “this name may change later”

Common confusion
- `const` does NOT mean the value is frozen
- You can still change the contents of arrays/objects

Example
- `const nums = [1, 2]; nums.push(3);` is allowed
- `const nums = [1, 2]; nums = [3, 4];` is not allowed

Rule of thumb
- Use `const` by default
- Use `let` only when you know the value must change

If all you remember is one thing
- `const` protects the name, not the contents
