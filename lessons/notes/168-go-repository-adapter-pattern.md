# Go repository adapter pattern

Goal: learn how repositories isolate storage details from domain logic.

Why do we care?
- Storage APIs change (file, SQLite, Postgres, remote service)
- Without boundaries, storage changes force rewrites everywhere

Pattern
- `Repository interface`: behavior contract used by service
- `Adapter implementations`: concrete storage choices
- `Service`: business rules using the interface
- `main`: composition root wiring service + adapter

History context
- Layered architecture became common because large systems need controlled change
- Go's interface model makes adapter swaps straightforward and explicit

Rule of thumb
- Define interface by what service needs, not by storage features
- Keep adapters boring and predictable

If all you remember is one thing
- Repositories protect business logic from storage churn
