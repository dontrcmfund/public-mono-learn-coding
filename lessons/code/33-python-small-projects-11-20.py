"""
PYTHON SMALL PROJECTS (Lessons 11-20)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/33-python-small-projects-11-20.py
3) Change one rule and predict output before running

Extra context:
- lessons/notes/93-python-small-projects-principles.md
- lessons/notes/94-python-learning-roadmap.md
"""

# -----------------------------------------------------------------------------
# LESSON 11 PROJECT: Budget tracker summary
# Why this matters: financial summaries are common real tasks.
def budget_summary(limit: float, spending: list[float]) -> dict[str, float | bool]:
    spent = sum(spending)
    return {
        "limit": limit,
        "spent": spent,
        "remaining": limit - spent,
        "over": spent > limit,
    }

print("Lesson 11:", budget_summary(200, [20, 40, 60]))

# -----------------------------------------------------------------------------
# LESSON 12 PROJECT: Top-N items
# Why this matters: ranking appears in analytics and dashboards.
def top_n(values: list[int], n: int) -> list[int]:
    return sorted(values, reverse=True)[: max(0, n)]

print("Lesson 12:", top_n([8, 3, 9, 1, 7], 3))

# -----------------------------------------------------------------------------
# LESSON 13 PROJECT: Date label formatter
# Why this matters: clean date strings improve reports.
from datetime import date


def format_label(day: date, name: str) -> str:
    return f"{day.isoformat()} | {name.strip().title()}"

print("Lesson 13:", format_label(date(2026, 2, 8), "weekly sync"))

# -----------------------------------------------------------------------------
# LESSON 14 PROJECT: Input sanitizer
# Why this matters: normalize noisy user input early.
def sanitize_text(value: str) -> str:
    return " ".join(value.strip().split())

print("Lesson 14:", sanitize_text("  hello    world  "))

# -----------------------------------------------------------------------------
# LESSON 15 PROJECT: Duplicate detector
# Why this matters: detecting collisions prevents data errors.
def find_duplicates(items: list[str]) -> list[str]:
    seen: set[str] = set()
    dups: set[str] = set()
    for item in items:
        if item in seen:
            dups.add(item)
        seen.add(item)
    return sorted(dups)

print("Lesson 15:", find_duplicates(["a", "b", "a", "c", "b", "d"]))

# -----------------------------------------------------------------------------
# LESSON 16 PROJECT: Safe percentage helper
# Why this matters: division-by-zero bugs are common.
def safe_percent(part: float, whole: float) -> float | None:
    if whole == 0:
        return None
    return (part / whole) * 100

print("Lesson 16:", safe_percent(25, 200), safe_percent(10, 0))

# -----------------------------------------------------------------------------
# LESSON 17 PROJECT: Group records by key
# Why this matters: grouping is foundational for reports.
def group_by_role(records: list[dict[str, str]]) -> dict[str, list[dict[str, str]]]:
    groups: dict[str, list[dict[str, str]]] = {}
    for record in records:
        role = record.get("role", "unknown")
        groups.setdefault(role, []).append(record)
    return groups

sample_records = [
    {"name": "Mia", "role": "admin"},
    {"name": "Leo", "role": "viewer"},
    {"name": "Ana", "role": "admin"},
]
print("Lesson 17:", group_by_role(sample_records))

# -----------------------------------------------------------------------------
# LESSON 18 PROJECT: Basic command parser
# Why this matters: CLIs start from parsing simple commands.
def parse_command(text: str) -> tuple[str, list[str]]:
    parts = text.strip().split()
    if not parts:
        return "", []
    return parts[0], parts[1:]

print("Lesson 18:", parse_command("add task Buy milk"))

# -----------------------------------------------------------------------------
# LESSON 19 PROJECT: Mini log filter
# Why this matters: filtering logs helps debugging and monitoring.
def filter_logs(lines: list[str], keyword: str) -> list[str]:
    key = keyword.lower()
    return [line for line in lines if key in line.lower()]

logs = ["INFO start", "ERROR bad input", "INFO done"]
print("Lesson 19:", filter_logs(logs, "error"))

# -----------------------------------------------------------------------------
# LESSON 20 PROJECT: Combined report pipeline
# Why this matters: real scripts chain helper functions.
def report_pipeline(raw_names: list[str]) -> str:
    cleaned = [sanitize_text(name) for name in raw_names if sanitize_text(name)]
    unique = sorted(set(cleaned))
    return "\n".join(f"- {name.title()}" for name in unique)

print("Lesson 20:\n" + report_pipeline([" alice", "Bob", "alice", " "]))

# End of Python Small Projects 11-20
