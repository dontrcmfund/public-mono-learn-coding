"""
PYTHON + SQLITE TESTING (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/46-python-sqlite-testing-1-10.py
2) Read PASS/FAIL output
3) Break one query and verify tests catch it
"""

import sqlite3
import tempfile
from pathlib import Path


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


def setup_db(path: Path) -> sqlite3.Connection:
    conn = sqlite3.connect(path)
    cur = conn.cursor()
    cur.execute("CREATE TABLE students (id INTEGER PRIMARY KEY, name TEXT, score INTEGER)")
    cur.executemany(
        "INSERT INTO students (name, score) VALUES (?, ?)",
        [("Mia", 90), ("Leo", 70), ("Ana", 85)],
    )
    conn.commit()
    return conn


# -----------------------------------------------------------------------------
# LESSON 1: Setup test DB in temp dir
# Why this matters: isolated DB avoids flaky tests.
with tempfile.TemporaryDirectory() as tmp:
    db_path = Path(tmp) / "test.db"
    conn = setup_db(db_path)
    run_test("Lesson 1 DB file exists", db_path.exists())
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 2: Count rows
# Why this matters: sanity check for seed setup.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("SELECT COUNT(*) FROM students")
    run_test("Lesson 2 row count", cur.fetchone()[0] == 3)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 3: WHERE filter behavior
# Why this matters: reports depend on filters.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("SELECT name FROM students WHERE score >= ? ORDER BY name", (85,))
    rows = [r[0] for r in cur.fetchall()]
    run_test("Lesson 3 high score names", rows == ["Ana", "Mia"])
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 4: Update with WHERE
# Why this matters: scoped update correctness is critical.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("UPDATE students SET score = ? WHERE name = ?", (75, "Leo"))
    conn.commit()
    cur.execute("SELECT score FROM students WHERE name = ?", ("Leo",))
    run_test("Lesson 4 update target row", cur.fetchone()[0] == 75)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 5: Aggregate test (AVG)
# Why this matters: summaries should be validated.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("SELECT ROUND(AVG(score), 2) FROM students")
    run_test("Lesson 5 average", cur.fetchone()[0] == 81.67)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 6: Delete with WHERE
# Why this matters: deletion logic needs proof.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("DELETE FROM students WHERE name = ?", ("Leo",))
    conn.commit()
    cur.execute("SELECT COUNT(*) FROM students")
    run_test("Lesson 6 delete one row", cur.fetchone()[0] == 2)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 7: Ordering guarantees
# Why this matters: deterministic order prevents flaky outputs.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    cur.execute("SELECT name FROM students ORDER BY score DESC")
    names = [r[0] for r in cur.fetchall()]
    run_test("Lesson 7 ordering", names == ["Mia", "Ana", "Leo"])
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 8: Parameter placeholder usage
# Why this matters: prevent SQL injection patterns.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    user_input = "Mia' OR 1=1 --"
    cur.execute("SELECT COUNT(*) FROM students WHERE name = ?", (user_input,))
    run_test("Lesson 8 placeholder safety", cur.fetchone()[0] == 0)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 9: Transaction rollback
# Why this matters: failed operations should not partially persist.
with tempfile.TemporaryDirectory() as tmp:
    conn = setup_db(Path(tmp) / "test.db")
    cur = conn.cursor()
    try:
        conn.execute("BEGIN")
        cur.execute("UPDATE students SET score = 100 WHERE name = 'Mia'")
        raise RuntimeError("simulate failure")
    except RuntimeError:
        conn.rollback()
    cur.execute("SELECT score FROM students WHERE name = 'Mia'")
    run_test("Lesson 9 rollback", cur.fetchone()[0] == 90)
    conn.close()


# -----------------------------------------------------------------------------
# LESSON 10: Completion marker
# Why this matters: clear suite end helps scan test output.
print("Lesson 10: sqlite testing suite complete")

# End of Python + SQLite Testing 1-10
