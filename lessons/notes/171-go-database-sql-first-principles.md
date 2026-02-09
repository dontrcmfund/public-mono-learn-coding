# Go `database/sql` (first principles)

Goal: understand why Go uses a generic DB package and driver model.

Why do we care?
- Business logic should not be tied to one database vendor
- `database/sql` gives a stable API for querying and transactions

History context
- Different SQL databases have different wire protocols and features
- Go separated common DB behavior (`database/sql`) from drivers
- This design keeps most application code portable

Core ideas
- `sql.DB` is a pooled connection manager, not a single connection
- Use parameterized queries (`?` / placeholders) to avoid SQL injection
- Always check and return query errors
- Use transactions for multi-step write consistency

Rule of thumb
- Keep SQL in repository adapters
- Keep service layer independent of SQL details

If all you remember is one thing
- `database/sql` is the portability layer; drivers are replaceable adapters
