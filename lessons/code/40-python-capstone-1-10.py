"""
PYTHON CAPSTONE (Lessons 1-10)

Project: Service health reporter

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/40-python-capstone-1-10.py
3) Change one rule and predict output before running

Extra context:
- lessons/notes/100-python-capstone-setup.md
- lessons/notes/101-python-capstone-workflow.md
"""

import json
from typing import Any


# -----------------------------------------------------------------------------
# LESSON 1: Define input shape expectation
# Why this matters: explicit contracts reduce downstream bugs.
def is_valid_service(record: dict[str, Any]) -> bool:
    return (
        isinstance(record.get("name"), str)
        and isinstance(record.get("status"), str)
        and record.get("status") in {"ok", "warn", "down"}
    )


# -----------------------------------------------------------------------------
# LESSON 2: Parse JSON safely
# Why this matters: malformed input should fail clearly.
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


# -----------------------------------------------------------------------------
# LESSON 3: Validate records
# Why this matters: boundary validation prevents silent corruption.
def validate_services(records: list[dict[str, Any]]) -> list[dict[str, Any]]:
    return [record for record in records if is_valid_service(record)]


# -----------------------------------------------------------------------------
# LESSON 4: Compute summary counts
# Why this matters: summaries are core reporting output.
def summarize_services(records: list[dict[str, Any]]) -> dict[str, int]:
    summary = {"ok": 0, "warn": 0, "down": 0, "total": 0}
    for record in records:
        status = record["status"]
        summary[status] += 1
        summary["total"] += 1
    return summary


# -----------------------------------------------------------------------------
# LESSON 5: Build human-readable lines
# Why this matters: users need readable status output.
def to_report_lines(records: list[dict[str, Any]]) -> list[str]:
    return [f"- {r['name']}: {r['status']}" for r in records]


# -----------------------------------------------------------------------------
# LESSON 6: Build full report text
# Why this matters: combine details + summary into one artifact.
def build_report_text(records: list[dict[str, Any]]) -> str:
    summary = summarize_services(records)
    lines = [
        "Service Health Report",
        f"Total: {summary['total']}",
        f"OK: {summary['ok']} | WARN: {summary['warn']} | DOWN: {summary['down']}",
        "",
        "Services:",
        *to_report_lines(records),
    ]
    return "\n".join(lines)


# -----------------------------------------------------------------------------
# LESSON 7: Build exportable row objects
# Why this matters: structured rows are easy to write to CSV.
def to_rows(records: list[dict[str, Any]]) -> list[dict[str, str]]:
    return [{"name": r["name"], "status": r["status"]} for r in records]


# -----------------------------------------------------------------------------
# LESSON 8: End-to-end in-memory pipeline
# Why this matters: one function should describe core flow.
def process_payload(raw: str) -> tuple[list[dict[str, Any]], str]:
    parsed = parse_services_json(raw)
    valid = validate_services(parsed)
    report = build_report_text(valid)
    return valid, report


# -----------------------------------------------------------------------------
# LESSON 9: Demo payload
# Why this matters: realistic sample data improves transfer to real APIs.
SAMPLE_RAW = json.dumps(
    {
        "services": [
            {"name": "web", "status": "ok"},
            {"name": "db", "status": "warn"},
            {"name": "worker", "status": "down"},
            {"name": "bad-entry", "status": "unknown"},
        ]
    }
)


# -----------------------------------------------------------------------------
# LESSON 10: Run demo pipeline
# Why this matters: immediate output closes the learning loop.
valid_services, report_text = process_payload(SAMPLE_RAW)
print("Lesson 10 valid count:", len(valid_services))
print("Lesson 10 report:\n" + report_text)

# End of Python Capstone 1-10
