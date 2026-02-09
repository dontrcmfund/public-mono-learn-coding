# SQL joins (first principles)

Goal: understand how to combine related data from multiple tables.

Why do we care?
- Real data is split across tables to avoid duplication
- Joins let you ask questions across those tables

First principles
- A join connects rows using related keys
- `INNER JOIN` keeps only matches
- `LEFT JOIN` keeps all left rows and optional right matches

Design tip
- Choose clear key columns (`id`, `student_id`, etc.)
- Index join keys when data grows

If all you remember is one thing
- Joins let separate tables answer one combined question
