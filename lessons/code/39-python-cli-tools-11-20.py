"""
PYTHON CLI TOOLS (Lessons 11-20)

Suggested use:
1) Open this file in VS Code
2) Run examples:
   - python3 lessons/code/39-python-cli-tools-11-20.py stats --numbers 2 4 6 8
   - python3 lessons/code/39-python-cli-tools-11-20.py normalize --text "  hello   world  "
   - python3 lessons/code/39-python-cli-tools-11-20.py export-json --out lessons/code/tmp-cli-data.json

Extra context:
- lessons/notes/98-python-cli-tools-principles.md
- lessons/notes/100-python-capstone-setup.md
"""

import argparse
import json
from pathlib import Path


# -----------------------------------------------------------------------------
# LESSON 11: add more subcommands
# Why this matters: expanding tools without chaos needs structure.
def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Advanced lesson CLI tool")
    sub = parser.add_subparsers(dest="command", required=True)

    stats = sub.add_parser("stats", help="Show stats for numbers")
    stats.add_argument("--numbers", nargs="+", type=float, required=True)

    normalize = sub.add_parser("normalize", help="Normalize text")
    normalize.add_argument("--text", required=True)

    export_json = sub.add_parser("export-json", help="Export sample JSON report")
    export_json.add_argument("--out", required=True)

    return parser


# -----------------------------------------------------------------------------
# LESSON 12: reusable statistics helper
# Why this matters: testable helpers reduce CLI complexity.
def compute_stats(values: list[float]) -> dict[str, float]:
    total = sum(values)
    count = len(values)
    avg = total / count if count else 0.0
    return {"count": float(count), "sum": total, "avg": avg}


# -----------------------------------------------------------------------------
# LESSON 13: text normalization helper
# Why this matters: normalization is common in automation scripts.
def normalize_text(text: str) -> str:
    return " ".join(text.strip().split())


# -----------------------------------------------------------------------------
# LESSON 14: export helper
# Why this matters: writing structured output enables downstream tooling.
def export_sample_json(path_text: str) -> Path:
    path = Path(path_text)
    path.parent.mkdir(parents=True, exist_ok=True)
    payload = {
        "service": "demo",
        "ok": True,
        "items": ["web", "db", "worker"],
    }
    path.write_text(json.dumps(payload, indent=2) + "\n", encoding="utf-8")
    return path


# -----------------------------------------------------------------------------
# LESSON 15: explicit command handlers
# Why this matters: easier maintenance and testing.
def handle_stats(numbers: list[float]) -> int:
    if not numbers:
        print("Lesson 15: no numbers provided")
        return 2
    stats = compute_stats(numbers)
    print("Lesson 15:", stats)
    return 0


def handle_normalize(text: str) -> int:
    print("Lesson 15:", normalize_text(text))
    return 0


def handle_export(path_text: str) -> int:
    path = export_sample_json(path_text)
    print(f"Lesson 15: wrote {path}")
    return 0


# -----------------------------------------------------------------------------
# LESSON 16: command dispatcher
# Why this matters: one dispatch point keeps flow predictable.
def dispatch(args: argparse.Namespace) -> int:
    if args.command == "stats":
        return handle_stats(args.numbers)
    if args.command == "normalize":
        return handle_normalize(args.text)
    if args.command == "export-json":
        return handle_export(args.out)

    print("Lesson 16: unknown command")
    return 1


# -----------------------------------------------------------------------------
# LESSON 17: argument-level validation (example)
# Why this matters: fail fast before doing work.
def validate_args(args: argparse.Namespace) -> int | None:
    if args.command == "stats" and len(args.numbers) > 1000:
        print("Lesson 17: too many numbers (max 1000)")
        return 2
    return None


# -----------------------------------------------------------------------------
# LESSON 18: main function with validation + dispatch
# Why this matters: clear top-level flow is easier to reason about.
def main() -> int:
    parser = build_parser()
    args = parser.parse_args()

    validation_error = validate_args(args)
    if validation_error is not None:
        return validation_error

    return dispatch(args)


# -----------------------------------------------------------------------------
# LESSON 19: entrypoint pattern
# Why this matters: module stays import-safe and CLI-ready.
if __name__ == "__main__":
    raise SystemExit(main())


# LESSON 20: architecture recap
# - parse args
# - validate inputs
# - run focused handler
# - return explicit exit code

# End of Python CLI Tools 11-20
