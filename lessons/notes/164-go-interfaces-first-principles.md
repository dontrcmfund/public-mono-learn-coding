# Go interfaces (first principles)

Goal: understand why interfaces exist and when they are useful.

Why do we care?
- Real systems depend on things that can change: database, API, filesystem, queue
- Interfaces let core logic depend on behavior, not concrete implementations

History context
- Go favors composition over inheritance
- Interfaces were designed to keep coupling low without complex class hierarchies
- This makes refactoring easier as codebases grow

Core idea
- An interface describes required behavior
- Any type that provides that behavior satisfies the interface
- You can swap implementations without rewriting service logic

When to use interfaces
- At boundaries: repositories, external clients, message publishers
- In tests: replace external dependencies with fakes

When not to use interfaces
- Not every struct needs an interface
- Start concrete, then extract interface at boundary or testing seam

If all you remember is one thing
- Interfaces reduce coupling by separating "what we need" from "how it is done"
