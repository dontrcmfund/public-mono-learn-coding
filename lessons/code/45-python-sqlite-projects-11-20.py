"""
PYTHON + SQLITE PROJECTS (Lessons 11-20)

Suggested use:
1) Run: python3 lessons/code/45-python-sqlite-projects-11-20.py
2) Read outputs and DB rows
3) Modify one query and predict output

Extra context:
- lessons/notes/112-sqlite-gotchas.md
- lessons/notes/113-sql-joins-first-principles.md
"""

import sqlite3
from pathlib import Path

DB_PATH = Path("lessons/code/tmp-school.db")
conn = sqlite3.connect(DB_PATH)
cursor = conn.cursor()


# -----------------------------------------------------------------------------
# LESSON 11: Create related tables
# Why this matters: normalized schemas reduce duplication.
cursor.execute(
    """
    CREATE TABLE IF NOT EXISTS classes (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL
    )
    """
)
cursor.execute(
    """
    CREATE TABLE IF NOT EXISTS students (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        class_id INTEGER NOT NULL,
        score INTEGER NOT NULL,
        FOREIGN KEY (class_id) REFERENCES classes(id)
    )
    """
)
conn.commit()
print("Lesson 11: schema ready")


# -----------------------------------------------------------------------------
# LESSON 12: Seed parent table data
# Why this matters: parent rows are needed before child references.
cursor.executemany("INSERT INTO classes (name) VALUES (?)", [("Math",), ("Science",)])
conn.commit()
print("Lesson 12: classes inserted")


# -----------------------------------------------------------------------------
# LESSON 13: Seed child table with foreign keys
# Why this matters: relational links connect records.
students = [
    ("Mia", 1, 91),
    ("Leo", 1, 78),
    ("Ana", 2, 88),
    ("Noa", 2, 94),
]
cursor.executemany("INSERT INTO students (name, class_id, score) VALUES (?, ?, ?)", students)
conn.commit()
print("Lesson 13: students inserted")


# -----------------------------------------------------------------------------
# LESSON 14: INNER JOIN query
# Why this matters: combine names and class labels in one report.
cursor.execute(
    """
    SELECT s.name, c.name, s.score
    FROM students s
    INNER JOIN classes c ON s.class_id = c.id
    ORDER BY s.score DESC
    """
)
print("Lesson 14:", cursor.fetchall())


# -----------------------------------------------------------------------------
# LESSON 15: Aggregate by class
# Why this matters: grouped summaries are core analytics.
cursor.execute(
    """
    SELECT c.name, ROUND(AVG(s.score), 2) AS avg_score, COUNT(*) AS student_count
    FROM students s
    INNER JOIN classes c ON s.class_id = c.id
    GROUP BY c.name
    ORDER BY avg_score DESC
    """
)
print("Lesson 15:", cursor.fetchall())


# -----------------------------------------------------------------------------
# LESSON 16: Parameterized class report
# Why this matters: reusable reports need dynamic filters.
def class_report(class_name: str):
    cursor.execute(
        """
        SELECT s.name, s.score
        FROM students s
        INNER JOIN classes c ON s.class_id = c.id
        WHERE c.name = ?
        ORDER BY s.score DESC
        """,
        (class_name,),
    )
    return cursor.fetchall()


print("Lesson 16:", class_report("Math"))


# -----------------------------------------------------------------------------
# LESSON 17: Transaction block
# Why this matters: all-or-nothing updates protect consistency.
try:
    conn.execute("BEGIN")
    cursor.execute("UPDATE students SET score = score + 1 WHERE class_id = ?", (1,))
    cursor.execute("UPDATE students SET score = score + 1 WHERE class_id = ?", (2,))
    conn.commit()
    print("Lesson 17: transaction committed")
except sqlite3.DatabaseError:
    conn.rollback()
    print("Lesson 17: transaction rolled back")


# -----------------------------------------------------------------------------
# LESSON 18: LEFT JOIN example
# Why this matters: keep parent rows even when child rows are missing.
cursor.execute("INSERT INTO classes (name) VALUES (?)", ("History",))
conn.commit()
cursor.execute(
    """
    SELECT c.name, COUNT(s.id)
    FROM classes c
    LEFT JOIN students s ON s.class_id = c.id
    GROUP BY c.id
    ORDER BY c.name
    """
)
print("Lesson 18:", cursor.fetchall())


# -----------------------------------------------------------------------------
# LESSON 19: Safe delete with filter
# Why this matters: deletion should be explicit and narrow.
cursor.execute("DELETE FROM students WHERE name = ?", ("Leo",))
conn.commit()
print("Lesson 19: deleted rows", cursor.rowcount)


# -----------------------------------------------------------------------------
# LESSON 20: Final state snapshot
# Why this matters: always verify final data state.
cursor.execute("SELECT id, name, class_id, score FROM students ORDER BY id")
print("Lesson 20:", cursor.fetchall())

conn.close()
print("Lesson 20: connection closed")

# End of Python + SQLite Projects 11-20
