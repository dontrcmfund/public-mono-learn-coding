"""
PYTHON CAPSTONE TESTS (Lessons 11-20)

Suggested use:
1) Run: python3 lessons/code/43-python-capstone-tests-11-20.py
2) Read PASS/FAIL output
3) Break one capstone helper and verify tests catch it
"""

import csv
import json
from pathlib import Path
import tempfile
from typing import Any


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


def normalize_records(payload: dict[str, Any]) -> list[dict[str, str]]:
    raw_services = payload.get("services")
    if not isinstance(raw_services, list):
        return []
    out: list[dict[str, str]] = []
    for item in raw_services:
        if not isinstance(item, dict):
            continue
        name = item.get("name")
        status = item.get("status")
        if isinstance(name, str) and status in {"ok", "warn", "down"}:
            out.append({"name": name.strip(), "status": status})
    return out


def summarize_services(records: list[dict[str, str]]) -> dict[str, int]:
    summary = {"ok": 0, "warn": 0, "down": 0, "total": 0}
    for record in records:
        status = record["status"]
        summary[status] += 1
        summary["total"] += 1
    return summary


def build_text(records: list[dict[str, str]]) -> str:
    summary = summarize_services(records)
    detail_lines = [f"- {r['name']}: {r['status']}" for r in records]
    lines = [
        "Service Health Report",
        f"Total: {summary['total']}",
        f"OK: {summary['ok']} | WARN: {summary['warn']} | DOWN: {summary['down']}",
        "",
        "Services:",
        *detail_lines,
    ]
    return "\n".join(lines)


def write_csv_report(path: Path, rows: list[dict[str, str]]) -> None:
    with path.open("w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=["name", "status"])
        writer.writeheader()
        for row in rows:
            writer.writerow(row)


# -----------------------------------------------------------------------------
# LESSON 11: normalize ignores malformed items
payload = {"services": [{"name": "web", "status": "ok"}, {"name": "db"}, "bad"]}
records = normalize_records(payload)
run_test("Lesson 11 normalize malformed", len(records) == 1)

# LESSON 12: normalize strips whitespace
payload2 = {"services": [{"name": "  web  ", "status": "ok"}]}
records2 = normalize_records(payload2)
run_test("Lesson 12 trim names", records2[0]["name"] == "web")

# LESSON 13: summary counts warn/down
summary = summarize_services([{"name": "db", "status": "warn"}, {"name": "w", "status": "down"}])
run_test("Lesson 13 summary counts", summary["warn"] == 1 and summary["down"] == 1)

# LESSON 14: build_text contains header
text = build_text([{"name": "web", "status": "ok"}])
run_test("Lesson 14 report header", "Service Health Report" in text)

# LESSON 15: build_text contains detail row
run_test("Lesson 15 report row", "- web: ok" in text)

# LESSON 16: CSV writer creates file
with tempfile.TemporaryDirectory() as tmp:
    p = Path(tmp) / "report.csv"
    write_csv_report(p, [{"name": "web", "status": "ok"}])
    run_test("Lesson 16 csv exists", p.exists())

# LESSON 17: CSV has header and row
with tempfile.TemporaryDirectory() as tmp:
    p = Path(tmp) / "report.csv"
    write_csv_report(p, [{"name": "web", "status": "ok"}])
    content = p.read_text(encoding="utf-8")
    run_test("Lesson 17 csv content", "name,status" in content and "web,ok" in content)

# LESSON 18: empty records summary
empty_summary = summarize_services([])
run_test("Lesson 18 empty summary", empty_summary == {"ok": 0, "warn": 0, "down": 0, "total": 0})

# LESSON 19: JSON roundtrip sanity for report metadata
meta = {"version": 1, "kind": "health-report"}
raw_meta = json.dumps(meta)
back = json.loads(raw_meta)
run_test("Lesson 19 json roundtrip", back.get("kind") == "health-report")

# LESSON 20: completion marker
print("Lesson 20: capstone test suite 11-20 complete")

# End of Python Capstone Tests 11-20
