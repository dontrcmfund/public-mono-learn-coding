# Go persistence (first principles)

Goal: understand why programs need persistence and how to design for it.

Why do we care?
- Memory is temporary; data disappears when the process stops
- Most useful software needs state to survive restarts

Past, present, future value
- Past: without persistence, users lose work
- Present: apps need stable records for reporting and debugging
- Future: persistent data enables analytics, automation, and integrations

Core idea
- Keep business rules separate from storage technology
- Business logic should say "save/load task" not "open file/sql statement"

Why this design matters
- You can start with in-memory storage for learning and tests
- Later, swap to file or database adapter with minimal service changes

If all you remember is one thing
- Persistence keeps data alive; boundaries keep architecture flexible
