# Go + SQLite gotchas

Goal: avoid common pitfalls when learning SQL persistence in Go.

Why do we care?
- DB bugs are often silent and discovered late
- Good habits now prevent data integrity issues later

Common gotchas
- Forgetting schema migration before queries
- Ignoring SQL errors and `rows.Err()`
- Building SQL with string concatenation (injection risk)
- Assuming writes are instantly safe without transaction boundaries
- Not closing rows/statements when required

SQLite-specific gotchas
- `:memory:` databases are per connection unless configured carefully
- Concurrent writes can block; retry/backoff strategies matter

Rule of thumb
- Migrate first, parameterize always, handle every error

If all you remember is one thing
- Treat database operations as failure-prone boundaries that require explicit care
