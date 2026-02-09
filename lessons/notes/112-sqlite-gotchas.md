# SQLite and SQL gotchas

Goal: avoid the mistakes that cause data loss or confusing query results.

Why do we care?
- Database mistakes can silently produce wrong output
- Some SQL behaviors differ from what beginners expect

Common gotchas
- Forgetting `WHERE` in `UPDATE` or `DELETE`
- Treating `NULL` like an empty string
- Building SQL with string concatenation instead of placeholders
- Assuming row order without `ORDER BY`
- Forgetting to `COMMIT` after writes

Rule of thumb
- Use placeholders for values
- Use explicit `ORDER BY`
- Wrap write operations in transactions

If all you remember is one thing
- Safe SQL means explicit filters, explicit ordering, and parameterized values
