"""
PYTHON CAPSTONE (Lessons 11-20)

Project: Service health reporter (outputs + CLI style execution)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/41-python-capstone-11-20.py
3) Inspect generated report files in lessons/code/

Extra context:
- lessons/notes/101-python-capstone-workflow.md
"""

import csv
import json
from pathlib import Path
from typing import Any


def summarize_services(records: list[dict[str, Any]]) -> dict[str, int]:
    summary = {"ok": 0, "warn": 0, "down": 0, "total": 0}
    for record in records:
        status = record["status"]
        summary[status] += 1
        summary["total"] += 1
    return summary


# -----------------------------------------------------------------------------
# LESSON 11: Write text report to file
# Why this matters: durable output is core script value.
def write_text_report(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text + "\n", encoding="utf-8")


# -----------------------------------------------------------------------------
# LESSON 12: Write CSV rows
# Why this matters: CSV output supports spreadsheet workflows.
def write_csv_report(path: Path, rows: list[dict[str, str]]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=["name", "status"])
        writer.writeheader()
        for row in rows:
            writer.writerow(row)


# -----------------------------------------------------------------------------
# LESSON 13: Read JSON payload from file
# Why this matters: real pipelines often start from files.
def read_json_file(path: Path) -> dict[str, Any] | None:
    try:
        raw = path.read_text(encoding="utf-8")
        data = json.loads(raw)
        return data if isinstance(data, dict) else None
    except (FileNotFoundError, json.JSONDecodeError):
        return None


# -----------------------------------------------------------------------------
# LESSON 14: Normalize payload records
# Why this matters: robust tools sanitize data before use.
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


# -----------------------------------------------------------------------------
# LESSON 15: Build final text report
# Why this matters: combine summary + details clearly.
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


# -----------------------------------------------------------------------------
# LESSON 16: Build output paths
# Why this matters: deterministic paths simplify usage and tests.
def output_paths(base_dir: Path) -> tuple[Path, Path]:
    return base_dir / "tmp-capstone-report.txt", base_dir / "tmp-capstone-report.csv"


# -----------------------------------------------------------------------------
# LESSON 17: Pipeline function
# Why this matters: one orchestrator makes behavior easy to reason about.
def run_pipeline(input_path: Path, output_dir: Path) -> tuple[int, str]:
    payload = read_json_file(input_path)
    if payload is None:
        return 1, "Input JSON missing or invalid"

    records = normalize_records(payload)
    text = build_text(records)

    txt_path, csv_path = output_paths(output_dir)
    write_text_report(txt_path, text)
    write_csv_report(csv_path, records)

    return 0, f"Wrote {txt_path} and {csv_path}"


# -----------------------------------------------------------------------------
# LESSON 18: Create demo input
# Why this matters: deterministic demo enables repeatable learning.
def ensure_demo_input(path: Path) -> None:
    payload = {
        "services": [
            {"name": "web", "status": "ok"},
            {"name": "db", "status": "warn"},
            {"name": "worker", "status": "down"},
        ]
    }
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(payload, indent=2), encoding="utf-8")


# -----------------------------------------------------------------------------
# LESSON 19: Run pipeline and print status
# Why this matters: visible status supports quick debugging.
def main() -> int:
    input_path = Path("lessons/code/tmp-capstone-input.json")
    output_dir = Path("lessons/code")

    ensure_demo_input(input_path)
    code, message = run_pipeline(input_path, output_dir)
    print("Lesson 19:", message)
    return code


# -----------------------------------------------------------------------------
# LESSON 20: Entrypoint with explicit exit code
# Why this matters: scripts should communicate success/failure to shells.
if __name__ == "__main__":
    raise SystemExit(main())

# End of Python Capstone 11-20
