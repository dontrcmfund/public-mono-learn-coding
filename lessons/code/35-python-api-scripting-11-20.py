"""
PYTHON API SCRIPTING (Lessons 11-20)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/35-python-api-scripting-11-20.py
3) Change one rule and predict output before running

Extra context:
- lessons/notes/95-python-api-scripting-principles.md
- lessons/notes/96-python-api-gotchas.md
"""

import csv
import json
import time
from typing import Any
from urllib import parse


# -----------------------------------------------------------------------------
# LESSON 11: Header builder helper
# Why this matters: APIs often require consistent headers.
def build_headers(token: str | None = None) -> dict[str, str]:
    headers = {"Accept": "application/json", "User-Agent": "lesson-script/1.0"}
    if token:
        headers["Authorization"] = f"Bearer {token}"
    return headers

print("Lesson 11:", build_headers("demo-token"))


# -----------------------------------------------------------------------------
# LESSON 12: Query param helper
# Why this matters: pagination and filters depend on query strings.
def with_query(base_url: str, params: dict[str, str | int]) -> str:
    normalized = {k: str(v) for k, v in params.items()}
    return f"{base_url}?{parse.urlencode(normalized)}"

print("Lesson 12:", with_query("https://api.example.com/items", {"page": 2, "limit": 50}))


# -----------------------------------------------------------------------------
# LESSON 13: Parse list payload safely
# Why this matters: scripts should reject malformed data early.
def parse_items(payload: dict[str, Any]) -> list[dict[str, Any]]:
    raw_items = payload.get("items")
    if not isinstance(raw_items, list):
        return []
    return [item for item in raw_items if isinstance(item, dict)]

print("Lesson 13:", parse_items({"items": [{"id": 1}, "bad"]}))


# -----------------------------------------------------------------------------
# LESSON 14: Pagination loop (simulated pages)
# Why this matters: many APIs return data across multiple pages.
SIMULATED_PAGES = {
    1: {"items": [{"id": 1}, {"id": 2}], "next_page": 2},
    2: {"items": [{"id": 3}], "next_page": None},
}


def fetch_page(page: int) -> dict[str, Any]:
    return SIMULATED_PAGES.get(page, {"items": [], "next_page": None})


def fetch_all_items() -> list[dict[str, Any]]:
    all_items: list[dict[str, Any]] = []
    page = 1
    while page is not None:
        payload = fetch_page(page)
        all_items.extend(parse_items(payload))
        next_page = payload.get("next_page")
        page = next_page if isinstance(next_page, int) else None
    return all_items

print("Lesson 14:", fetch_all_items())


# -----------------------------------------------------------------------------
# LESSON 15: Rate-limit friendly sleep
# Why this matters: slower requests can prevent API blocking.
def wait_between_requests(delay_seconds: float) -> None:
    time.sleep(max(0.0, delay_seconds))

wait_between_requests(0.01)
print("Lesson 15: waited safely")


# -----------------------------------------------------------------------------
# LESSON 16: Retry policy decision
# Why this matters: retries should be explicit and bounded.
def should_retry(status_code: int, attempt: int, max_attempts: int) -> bool:
    if attempt >= max_attempts:
        return False
    return status_code in {429, 500, 502, 503, 504}

print("Lesson 16:", should_retry(503, 1, 3), should_retry(404, 1, 3))


# -----------------------------------------------------------------------------
# LESSON 17: Backoff calculator
# Why this matters: exponential delays reduce pressure on services.
def backoff_seconds(attempt: int) -> float:
    return 0.2 * (2 ** max(0, attempt - 1))

print("Lesson 17:", [backoff_seconds(i) for i in [1, 2, 3]])


# -----------------------------------------------------------------------------
# LESSON 18: Simple in-memory cache
# Why this matters: cache avoids repeated work and duplicate requests.
CACHE: dict[str, dict[str, Any]] = {}


def get_or_set_cache(key: str, factory) -> dict[str, Any]:
    if key not in CACHE:
        CACHE[key] = factory()
    return CACHE[key]

cached = get_or_set_cache("health", lambda: {"ok": True, "service": "demo"})
print("Lesson 18:", cached)


# -----------------------------------------------------------------------------
# LESSON 19: Export rows to CSV
# Why this matters: API scripts often produce spreadsheet-friendly output.
def export_csv(path: str, rows: list[dict[str, Any]], fieldnames: list[str]) -> None:
    with open(path, "w", encoding="utf-8", newline="") as f:
        writer = csv.DictWriter(f, fieldnames=fieldnames)
        writer.writeheader()
        for row in rows:
            writer.writerow({k: row.get(k, "") for k in fieldnames})

sample_rows = [{"id": 1, "name": "web"}, {"id": 2, "name": "db"}]
export_csv("lessons/code/tmp-services.csv", sample_rows, ["id", "name"])
print("Lesson 19: exported CSV")


# -----------------------------------------------------------------------------
# LESSON 20: End-to-end API report pipeline (simulated)
# Why this matters: real scripts chain fetch -> validate -> transform -> export.
def api_report_pipeline() -> str:
    items = fetch_all_items()
    report = {
        "total_items": len(items),
        "item_ids": [item.get("id") for item in items],
    }
    text = json.dumps(report, indent=2)
    with open("lessons/code/tmp-api-summary.json", "w", encoding="utf-8") as f:
        f.write(text + "\n")
    return text

print("Lesson 20:\n" + api_report_pipeline())

# End of Python API Scripting 11-20
