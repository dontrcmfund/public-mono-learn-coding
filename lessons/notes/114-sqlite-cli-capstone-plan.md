# SQLite CLI capstone plan

Goal: connect SQL querying, reporting, and CLI interaction in one practical tool.

Capstone scope
- Read student data from SQLite
- Run filterable reports
- Export report rows to CSV
- Return clear status messages and exit semantics

Why this matters
- Mirrors real internal tooling patterns
- Reinforces safe SQL, deterministic output, and reproducible workflows

Build order
- Parse CLI args
- Validate user inputs
- Execute parameterized query
- Format and export output
- Add focused tests for each step

If all you remember is one thing
- A good data CLI is query-safe, filterable, and testable
