# What is a database? (first principles)

Goal: understand what a database is and why software uses one.

Why do we care?
- Files are good for simple storage, but databases scale better for querying and updates
- Databases keep structured records consistent over time
- Most real applications need persistent, searchable data

Why this matters to you (past, present, future)
- Past: if text/JSON files became hard to search, databases solve that pain
- Present: you can store and query project data with simple SQL
- Future: database literacy is required in most software roles

First principles
- A database stores records in structured form
- A table is a collection of rows with defined columns
- Queries retrieve and transform data

Short history
- Early systems used flat files; query flexibility was limited
- Relational databases popularized SQL for structured querying
- SQLite made embedded databases simple for local tools and apps

Etymology
- `data` means recorded facts
- `base` means foundation/store
- `relational` means data linked through shared keys

If all you remember is one thing
- Databases exist to store structured data you can query reliably
