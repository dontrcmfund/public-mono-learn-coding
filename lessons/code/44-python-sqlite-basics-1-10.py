"""
PYTHON + SQLITE BASICS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/44-python-sqlite-basics-1-10.py
2) Observe created DB file in lessons/code/
3) Modify one query and predict output before running

Extra context:
- lessons/notes/110-what-is-a-database.md
- lessons/notes/111-sql-first-principles.md
"""

import sqlite3
from pathlib import Path

DB_PATH = Path("lessons/code/tmp-lessons.db")


# -----------------------------------------------------------------------------
# LESSON 1: Connect to SQLite database
# Why this matters: connection is the entrypoint for all DB work.
conn = sqlite3.connect(DB_PATH)
cursor = conn.cursor()
print("Lesson 1: connected", DB_PATH)


# -----------------------------------------------------------------------------
# LESSON 2: Create table
# Why this matters: schema defines structure and constraints.
cursor.execute(
    """
    CREATE TABLE IF NOT EXISTS students (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        score INTEGER NOT NULL
    )
    """
)
conn.commit()
print("Lesson 2: table ready")


# -----------------------------------------------------------------------------
# LESSON 3: Insert rows safely with placeholders
# Why this matters: placeholders prevent SQL injection and formatting bugs.
students = [("Mia", 92), ("Leo", 78), ("Ana", 85)]
cursor.executemany("INSERT INTO students (name, score) VALUES (?, ?)", students)
conn.commit()
print("Lesson 3: inserted", len(students), "rows")


# -----------------------------------------------------------------------------
# LESSON 4: Basic SELECT
# Why this matters: read queries are core DB usage.
cursor.execute("SELECT id, name, score FROM students")
all_rows = cursor.fetchall()
print("Lesson 4:", all_rows)


# -----------------------------------------------------------------------------
# LESSON 5: WHERE filter
# Why this matters: filtering avoids scanning irrelevant data.
cursor.execute("SELECT name, score FROM students WHERE score >= ?", (85,))
high_scores = cursor.fetchall()
print("Lesson 5:", high_scores)


# -----------------------------------------------------------------------------
# LESSON 6: ORDER BY
# Why this matters: sorted output improves reports.
cursor.execute("SELECT name, score FROM students ORDER BY score DESC")
ranked = cursor.fetchall()
print("Lesson 6:", ranked)


# -----------------------------------------------------------------------------
# LESSON 7: UPDATE with WHERE
# Why this matters: targeted updates prevent accidental bulk changes.
cursor.execute("UPDATE students SET score = ? WHERE name = ?", (80, "Leo"))
conn.commit()
print("Lesson 7: updated rows", cursor.rowcount)


# -----------------------------------------------------------------------------
# LESSON 8: COUNT aggregate
# Why this matters: summary metrics are common in dashboards.
cursor.execute("SELECT COUNT(*) FROM students")
count = cursor.fetchone()[0]
print("Lesson 8: count", count)


# -----------------------------------------------------------------------------
# LESSON 9: DELETE with WHERE
# Why this matters: safe deletes require explicit filters.
cursor.execute("DELETE FROM students WHERE name = ?", ("Ana",))
conn.commit()
print("Lesson 9: deleted rows", cursor.rowcount)


# -----------------------------------------------------------------------------
# LESSON 10: Final read + close connection
# Why this matters: confirm state and release resources.
cursor.execute("SELECT id, name, score FROM students ORDER BY id")
print("Lesson 10:", cursor.fetchall())
conn.close()
print("Lesson 10: connection closed")

# End of Python + SQLite Basics 1-10
