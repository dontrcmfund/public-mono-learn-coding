"""
PYTHON + SQLITE CLI CAPSTONE TESTS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/48-python-sqlite-cli-capstone-tests-1-10.py
2) Read PASS/FAIL output
3) Break one helper in 47 file and verify failures
"""

import csv
import sqlite3
import tempfile
from pathlib import Path


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


def setup_schema(conn: sqlite3.Connection) -> None:
    cur = conn.cursor()
    cur.execute(
        """
        CREATE TABLE students (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            class_name TEXT NOT NULL,
            score INTEGER NOT NULL
        )
        """
    )
    conn.commit()


def seed_data(conn: sqlite3.Connection) -> None:
    cur = conn.cursor()
    cur.executemany(
        "INSERT INTO students (name, class_name, score) VALUES (?, ?, ?)",
        [
            ("Mia", "Math", 92),
            ("Leo", "Math", 78),
            ("Ana", "Science", 88),
            ("Noa", "Science", 95),
        ],
    )
    conn.commit()


def query_students(conn: sqlite3.Connection, min_score: int):
    cur = conn.cursor()
    cur.execute(
        "SELECT id, name, class_name, score FROM students WHERE score >= ? ORDER BY score DESC, name ASC",
        (min_score,),
    )
    return cur.fetchall()


def render_text_report(rows):
    lines = ["Student Report", f"Rows: {len(rows)}", ""]
    for _id, name, class_name, score in rows:
        lines.append(f"- {name} ({class_name}): {score}")
    return "\n".join(lines)


def export_csv(path: Path, rows):
    with path.open("w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=["id", "name", "class_name", "score"])
        writer.writeheader()
        for r in rows:
            writer.writerow({"id": r[0], "name": r[1], "class_name": r[2], "score": r[3]})


# -----------------------------------------------------------------------------
# LESSON 1: setup + seed count
with tempfile.TemporaryDirectory() as tmp:
    db = Path(tmp) / "capstone.db"
    conn = sqlite3.connect(db)
    setup_schema(conn)
    seed_data(conn)
    cur = conn.cursor()
    cur.execute("SELECT COUNT(*) FROM students")
    run_test("Lesson 1 row count", cur.fetchone()[0] == 4)
    conn.close()

# LESSON 2: filtered query
with tempfile.TemporaryDirectory() as tmp:
    conn = sqlite3.connect(Path(tmp) / "capstone.db")
    setup_schema(conn)
    seed_data(conn)
    rows = query_students(conn, 90)
    run_test("Lesson 2 min_score filter", [r[1] for r in rows] == ["Noa", "Mia"])
    conn.close()

# LESSON 3: ordering contract
with tempfile.TemporaryDirectory() as tmp:
    conn = sqlite3.connect(Path(tmp) / "capstone.db")
    setup_schema(conn)
    seed_data(conn)
    rows = query_students(conn, 0)
    run_test("Lesson 3 order desc", rows[0][3] >= rows[1][3] >= rows[2][3])
    conn.close()

# LESSON 4: report header
sample_rows = [(1, "Mia", "Math", 92)]
text = render_text_report(sample_rows)
run_test("Lesson 4 report header", "Student Report" in text)

# LESSON 5: report row text
run_test("Lesson 5 report row", "- Mia (Math): 92" in text)

# LESSON 6: csv export creates file
with tempfile.TemporaryDirectory() as tmp:
    path = Path(tmp) / "out.csv"
    export_csv(path, sample_rows)
    run_test("Lesson 6 csv exists", path.exists())

# LESSON 7: csv export content
with tempfile.TemporaryDirectory() as tmp:
    path = Path(tmp) / "out.csv"
    export_csv(path, sample_rows)
    raw = path.read_text(encoding="utf-8")
    run_test("Lesson 7 csv content", "name,class_name,score" in raw and "Mia,Math,92" in raw)

# LESSON 8: empty result report
empty_report = render_text_report([])
run_test("Lesson 8 empty rows", "Rows: 0" in empty_report)

# LESSON 9: min-score edge case
with tempfile.TemporaryDirectory() as tmp:
    conn = sqlite3.connect(Path(tmp) / "capstone.db")
    setup_schema(conn)
    seed_data(conn)
    rows = query_students(conn, 200)
    run_test("Lesson 9 high threshold", rows == [])
    conn.close()

# LESSON 10: completion marker
print("Lesson 10: sqlite CLI capstone tests complete")

# End of Python + SQLite CLI Capstone Tests 1-10
