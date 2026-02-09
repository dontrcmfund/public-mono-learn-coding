"""
PYTHON + SQLITE CLI CAPSTONE (Lessons 1-10)

Suggested use:
1) Run setup and report:
   python3 lessons/code/47-python-sqlite-cli-capstone-1-10.py setup
   python3 lessons/code/47-python-sqlite-cli-capstone-1-10.py report --min-score 85
2) Export report:
   python3 lessons/code/47-python-sqlite-cli-capstone-1-10.py export --min-score 80 --out lessons/code/tmp-sqlite-report.csv

Extra context:
- lessons/notes/114-sqlite-cli-capstone-plan.md
- lessons/notes/112-sqlite-gotchas.md
"""

import argparse
import csv
import sqlite3
from pathlib import Path

DB_PATH = Path("lessons/code/tmp-sqlite-capstone.db")


def connect_db() -> sqlite3.Connection:
    return sqlite3.connect(DB_PATH)


# -----------------------------------------------------------------------------
# LESSON 1: setup schema
# Why this matters: repeatable schema setup supports reliable tooling.
def setup_schema(conn: sqlite3.Connection) -> None:
    cur = conn.cursor()
    cur.execute(
        """
        CREATE TABLE IF NOT EXISTS students (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            class_name TEXT NOT NULL,
            score INTEGER NOT NULL
        )
        """
    )
    conn.commit()


# -----------------------------------------------------------------------------
# LESSON 2: seed sample data
# Why this matters: deterministic input helps testing and demos.
def seed_data(conn: sqlite3.Connection) -> None:
    cur = conn.cursor()
    cur.execute("DELETE FROM students")
    rows = [
        ("Mia", "Math", 92),
        ("Leo", "Math", 78),
        ("Ana", "Science", 88),
        ("Noa", "Science", 95),
    ]
    cur.executemany("INSERT INTO students (name, class_name, score) VALUES (?, ?, ?)", rows)
    conn.commit()


# -----------------------------------------------------------------------------
# LESSON 3: query helper with filter
# Why this matters: parameterized filters are safer and reusable.
def query_students(conn: sqlite3.Connection, min_score: int) -> list[tuple[int, str, str, int]]:
    cur = conn.cursor()
    cur.execute(
        """
        SELECT id, name, class_name, score
        FROM students
        WHERE score >= ?
        ORDER BY score DESC, name ASC
        """,
        (min_score,),
    )
    return cur.fetchall()


# -----------------------------------------------------------------------------
# LESSON 4: text report renderer
# Why this matters: readable output improves CLI usability.
def render_text_report(rows: list[tuple[int, str, str, int]]) -> str:
    lines = ["Student Report"]
    lines.append(f"Rows: {len(rows)}")
    lines.append("")
    for _id, name, class_name, score in rows:
        lines.append(f"- {name} ({class_name}): {score}")
    return "\n".join(lines)


# -----------------------------------------------------------------------------
# LESSON 5: csv export helper
# Why this matters: report exports are common workflow outputs.
def export_csv(path: Path, rows: list[tuple[int, str, str, int]]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=["id", "name", "class_name", "score"])
        writer.writeheader()
        for r in rows:
            writer.writerow({"id": r[0], "name": r[1], "class_name": r[2], "score": r[3]})


# -----------------------------------------------------------------------------
# LESSON 6: command handlers
# Why this matters: each command should have one clear responsibility.
def handle_setup() -> int:
    conn = connect_db()
    try:
        setup_schema(conn)
        seed_data(conn)
        print("Lesson 6: setup complete")
        return 0
    finally:
        conn.close()


def handle_report(min_score: int) -> int:
    conn = connect_db()
    try:
        rows = query_students(conn, min_score)
        print(render_text_report(rows))
        return 0
    finally:
        conn.close()


def handle_export(min_score: int, out_path: str) -> int:
    conn = connect_db()
    try:
        rows = query_students(conn, min_score)
        export_csv(Path(out_path), rows)
        print(f"Lesson 6: exported {len(rows)} rows to {out_path}")
        return 0
    finally:
        conn.close()


# -----------------------------------------------------------------------------
# LESSON 7: parser with subcommands
# Why this matters: structured CLI design scales better.
def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="SQLite capstone CLI")
    sub = parser.add_subparsers(dest="cmd", required=True)

    sub.add_parser("setup", help="Initialize and seed database")

    report = sub.add_parser("report", help="Print filtered report")
    report.add_argument("--min-score", type=int, default=0)

    export = sub.add_parser("export", help="Export filtered report to CSV")
    export.add_argument("--min-score", type=int, default=0)
    export.add_argument("--out", required=True)

    return parser


# -----------------------------------------------------------------------------
# LESSON 8: dispatch command
# Why this matters: central dispatch keeps entrypoint logic simple.
def dispatch(args: argparse.Namespace) -> int:
    if args.cmd == "setup":
        return handle_setup()
    if args.cmd == "report":
        return handle_report(args.min_score)
    if args.cmd == "export":
        return handle_export(args.min_score, args.out)
    print("Unknown command")
    return 1


# -----------------------------------------------------------------------------
# LESSON 9: input validation
# Why this matters: fail fast with clear user feedback.
def validate_args(args: argparse.Namespace) -> int | None:
    if hasattr(args, "min_score") and args.min_score < 0:
        print("Lesson 9: min-score must be >= 0")
        return 2
    return None


# -----------------------------------------------------------------------------
# LESSON 10: main entrypoint
# Why this matters: explicit top-level flow improves maintainability.
def main() -> int:
    parser = build_parser()
    args = parser.parse_args()

    validation_code = validate_args(args)
    if validation_code is not None:
        return validation_code

    return dispatch(args)


if __name__ == "__main__":
    raise SystemExit(main())

# End of Python + SQLite CLI Capstone 1-10
