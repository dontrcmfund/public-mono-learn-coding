# What is a package? (first principles)

Goal: understand what a package is without code.

Why do we care?
- Packages let us group related code
- Shared code is easier when it lives in one package
- Monorepos are just many packages in one place

First principles
- A package is a folder with a `package.json`
- `package.json` is a label: name, version, and settings
- Tools use that label to know “this is a package”

What to do (slow and simple)
- Open `packages/shared/package.json`
- Read the `name` field out loud
- Notice there is no real code yet — that is okay

Checkpoint
- If you can explain “a package is a labeled folder,” you are ready to move on

Reflection
- Why might you want a shared package later?
- What would be confusing if every app copied the same code?
