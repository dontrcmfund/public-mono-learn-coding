"""
PYTHON API SCRIPTING (Lessons 1-10)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/34-python-api-scripting-1-10.py
3) Change one rule and predict output before running

Extra context:
- lessons/notes/95-python-api-scripting-principles.md
- lessons/notes/91-python-common-gotchas.md
"""

import json
import time
from typing import Any
from urllib import error, parse, request


# -----------------------------------------------------------------------------
# LESSON 1: Build URL with query params
# Why this matters: avoid fragile string concatenation.
def build_url(base: str, params: dict[str, str]) -> str:
    return f"{base}?{parse.urlencode(params)}"

print("Lesson 1:", build_url("https://example.com/search", {"q": "python"}))


# -----------------------------------------------------------------------------
# LESSON 2: Parse JSON text safely
# Why this matters: API payloads can be invalid.
def parse_json_safely(raw: str) -> dict[str, Any] | None:
    try:
        data = json.loads(raw)
        if isinstance(data, dict):
            return data
        return None
    except json.JSONDecodeError:
        return None

print("Lesson 2:", parse_json_safely('{"ok": true}'))


# -----------------------------------------------------------------------------
# LESSON 3: Validate response shape
# Why this matters: typed assumptions should be checked explicitly.
def parse_health(data: dict[str, Any]) -> tuple[bool, str] | None:
    ok = data.get("ok")
    service = data.get("service")
    if isinstance(ok, bool) and isinstance(service, str):
        return ok, service
    return None

print("Lesson 3:", parse_health({"ok": True, "service": "demo"}))


# -----------------------------------------------------------------------------
# LESSON 4: Basic GET request helper
# Why this matters: wrap networking in reusable functions.
def http_get_text(url: str, timeout: float = 3.0) -> str | None:
    req = request.Request(url, method="GET")
    try:
        with request.urlopen(req, timeout=timeout) as response:
            return response.read().decode("utf-8")
    except (error.URLError, TimeoutError):
        return None

print("Lesson 4:", http_get_text("https://example.com") is not None)


# -----------------------------------------------------------------------------
# LESSON 5: Retry with backoff
# Why this matters: networks fail; retries improve robustness.
def get_with_retries(url: str, attempts: int = 3) -> str | None:
    for attempt in range(1, attempts + 1):
        text = http_get_text(url)
        if text is not None:
            return text
        if attempt < attempts:
            time.sleep(0.2 * (2 ** (attempt - 1)))
    return None

print("Lesson 5:", get_with_retries("https://example.com") is not None)


# -----------------------------------------------------------------------------
# LESSON 6: Extract fields from JSON payload
# Why this matters: scripts usually only need part of API data.
def extract_title(payload: dict[str, Any]) -> str:
    title = payload.get("title")
    return title if isinstance(title, str) else "(missing title)"

print("Lesson 6:", extract_title({"title": "Daily Report"}))


# -----------------------------------------------------------------------------
# LESSON 7: Convert list payload to report lines
# Why this matters: reports are practical script output.
def items_to_report(items: list[dict[str, Any]]) -> list[str]:
    lines: list[str] = []
    for item in items:
        name = item.get("name")
        status = item.get("status")
        if isinstance(name, str) and isinstance(status, str):
            lines.append(f"- {name}: {status}")
    return lines

sample_items = [{"name": "API", "status": "ok"}, {"name": "DB", "status": "degraded"}]
print("Lesson 7:", items_to_report(sample_items))


# -----------------------------------------------------------------------------
# LESSON 8: Save API result to file
# Why this matters: automation often writes artifacts for later review.
def save_text(path: str, text: str) -> None:
    with open(path, "w", encoding="utf-8") as f:
        f.write(text)

save_text("lessons/code/tmp-api-report.txt", "Lesson 8: report saved\n")
print("Lesson 8: saved file")


# -----------------------------------------------------------------------------
# LESSON 9: Load API result from file
# Why this matters: scripts frequently read previous outputs.
def load_text(path: str) -> str | None:
    try:
        with open(path, "r", encoding="utf-8") as f:
            return f.read()
    except FileNotFoundError:
        return None

print("Lesson 9:", (load_text("lessons/code/tmp-api-report.txt") or "").strip())


# -----------------------------------------------------------------------------
# LESSON 10: End-to-end mini pipeline
# Why this matters: real scripts combine fetch, parse, validate, and report.
def mini_pipeline() -> list[str]:
    fake_raw = '{"services": [{"name": "web", "status": "ok"}, {"name": "db", "status": "warn"}]}'
    payload = parse_json_safely(fake_raw)
    if payload is None:
        return ["Pipeline error: invalid JSON"]

    services = payload.get("services")
    if not isinstance(services, list):
        return ["Pipeline error: missing services"]

    typed_services = [s for s in services if isinstance(s, dict)]
    return items_to_report(typed_services)

print("Lesson 10:", mini_pipeline())

# End of Python API Scripting 1-10
