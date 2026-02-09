# What is SQLite in Go?

Goal: understand why SQLite is a useful next step after file-based persistence.

Why do we care?
- JSON files are simple, but querying and concurrency become harder as data grows
- SQLite gives structured storage with SQL, indexes, and transactions in one file

Past, present, future value
- Past: file persistence teaches storage basics
- Present: SQLite teaches relational thinking without server setup
- Future: the same SQL concepts transfer to Postgres/MySQL

History context
- SQLite was created to embed a full SQL database directly in applications
- Its design goal: reliable, zero-admin, portable database engine
- This made it common in mobile, desktop, and local tooling

Why it fits this course
- One file database is beginner friendly
- It still teaches real database practices: schema, queries, constraints

If all you remember is one thing
- SQLite is the bridge from simple files to real relational data modeling
