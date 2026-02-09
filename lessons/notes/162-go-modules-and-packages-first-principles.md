# Go modules and packages (first principles)

Goal: understand how Go organizes code so projects stay maintainable.

Why do we care?
- Small files are easy to read, but real apps need structure across many files
- Clear package boundaries make testing and refactoring safer

History context
- Before Go modules, teams depended on `GOPATH` conventions that were harder to reproduce
- Go modules made dependency and version management explicit and portable
- This changed Go from "works on my machine" risk toward reproducible builds

Core ideas
- A `module` is the versioned project unit
- A `package` is a folder-level code unit with one package name
- `import` connects packages through explicit dependencies

Why the design looks this way
- Go favors simple, explicit build rules over complex magic
- Fewer hidden behaviors means fewer surprises for teams

If all you remember is one thing
- Modules manage project versions; packages manage code boundaries
