"""
PYTHON CAPSTONE TESTS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/42-python-capstone-tests-1-10.py
2) Read PASS/FAIL output
3) Break one helper intentionally and rerun
"""

import json
from pathlib import Path
import tempfile
from typing import Any


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


def is_valid_service(record: dict[str, Any]) -> bool:
    return (
        isinstance(record.get("name"), str)
        and isinstance(record.get("status"), str)
        and record.get("status") in {"ok", "warn", "down"}
    )


def parse_services_json(raw: str) -> list[dict[str, Any]]:
    try:
        payload = json.loads(raw)
    except json.JSONDecodeError:
        return []
    if not isinstance(payload, dict):
        return []
    services = payload.get("services")
    if not isinstance(services, list):
        return []
    return [s for s in services if isinstance(s, dict)]


def summarize_services(records: list[dict[str, Any]]) -> dict[str, int]:
    summary = {"ok": 0, "warn": 0, "down": 0, "total": 0}
    for record in records:
        status = record["status"]
        summary[status] += 1
        summary["total"] += 1
    return summary


def write_text(path: Path, text: str) -> None:
    path.write_text(text, encoding="utf-8")


# -----------------------------------------------------------------------------
# LESSON 1: Validate good record
run_test("Lesson 1 valid record", is_valid_service({"name": "web", "status": "ok"}))

# LESSON 2: Validate bad record
run_test("Lesson 2 invalid status", not is_valid_service({"name": "web", "status": "unknown"}))

# LESSON 3: Parse good JSON
parsed = parse_services_json('{"services": [{"name": "db", "status": "warn"}]}')
run_test("Lesson 3 parse JSON", len(parsed) == 1)

# LESSON 4: Parse invalid JSON
run_test("Lesson 4 invalid JSON", parse_services_json("{bad") == [])

# LESSON 5: Summary totals
summary = summarize_services([
    {"name": "web", "status": "ok"},
    {"name": "db", "status": "warn"},
    {"name": "worker", "status": "down"},
])
run_test("Lesson 5 total", summary["total"] == 3)

# LESSON 6: Summary per status
run_test("Lesson 6 per status", summary["ok"] == 1 and summary["warn"] == 1 and summary["down"] == 1)

# LESSON 7: File write test
with tempfile.TemporaryDirectory() as tmp:
    p = Path(tmp) / "out.txt"
    write_text(p, "hello")
    run_test("Lesson 7 file write", p.exists())

# LESSON 8: File readback test
with tempfile.TemporaryDirectory() as tmp:
    p = Path(tmp) / "out.txt"
    write_text(p, "hello")
    run_test("Lesson 8 file readback", p.read_text(encoding="utf-8") == "hello")

# LESSON 9: Empty services handling
run_test("Lesson 9 empty services", summarize_services([])["total"] == 0)

# LESSON 10: Completion marker
print("Lesson 10: capstone test suite complete")

# End of Python Capstone Tests 1-10
