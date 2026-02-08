"""
PYTHON CLI TOOLS (Lessons 1-10)

Suggested use:
1) Open this file in VS Code
2) Run examples:
   - python3 lessons/code/38-python-cli-tools-1-10.py greet --name Mia
   - python3 lessons/code/38-python-cli-tools-1-10.py sum --values 3 5 7
   - python3 lessons/code/38-python-cli-tools-1-10.py report --items web:ok db:warn

Extra context:
- lessons/notes/98-python-cli-tools-principles.md
"""

import argparse
from pathlib import Path


# -----------------------------------------------------------------------------
# LESSON 1: Build parser with subcommands
# Why this matters: subcommands keep tools organized.
def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="Lesson CLI tool")
    sub = parser.add_subparsers(dest="command", required=True)

    greet = sub.add_parser("greet", help="Greet a user")
    greet.add_argument("--name", required=True)

    summation = sub.add_parser("sum", help="Sum integer values")
    summation.add_argument("--values", nargs="+", type=int, required=True)

    report = sub.add_parser("report", help="Create simple service report")
    report.add_argument("--items", nargs="+", required=True, help="Format: name:status")
    report.add_argument("--out", default="lessons/code/tmp-cli-report.txt")

    return parser


# -----------------------------------------------------------------------------
# LESSON 2: Command handler pattern
# Why this matters: one dispatcher keeps behavior clear.
def handle_command(args: argparse.Namespace) -> int:
    if args.command == "greet":
        return handle_greet(args.name)
    if args.command == "sum":
        return handle_sum(args.values)
    if args.command == "report":
        return handle_report(args.items, args.out)

    print("Unknown command")
    return 1


# -----------------------------------------------------------------------------
# LESSON 3: Greet command
# Why this matters: start with predictable tiny behavior.
def handle_greet(name: str) -> int:
    print(f"Lesson 3: Hello, {name}!")
    return 0


# -----------------------------------------------------------------------------
# LESSON 4: Sum command
# Why this matters: parse -> compute -> output flow.
def handle_sum(values: list[int]) -> int:
    print("Lesson 4:", sum(values))
    return 0


# -----------------------------------------------------------------------------
# LESSON 5: Parse report items
# Why this matters: validate user input format early.
def parse_report_items(items: list[str]) -> list[tuple[str, str]] | None:
    parsed: list[tuple[str, str]] = []
    for raw in items:
        if ":" not in raw:
            return None
        name, status = raw.split(":", 1)
        name = name.strip()
        status = status.strip()
        if not name or not status:
            return None
        parsed.append((name, status))
    return parsed


# -----------------------------------------------------------------------------
# LESSON 6: Report command with file output
# Why this matters: many CLIs generate artifacts.
def handle_report(items: list[str], out_path: str) -> int:
    parsed = parse_report_items(items)
    if parsed is None:
        print("Lesson 6: invalid item format. Use name:status")
        return 2

    lines = [f"- {name}: {status}" for name, status in parsed]
    content = "Report\n" + "\n".join(lines) + "\n"

    path = Path(out_path)
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")

    print(f"Lesson 6: report written to {path}")
    return 0


# -----------------------------------------------------------------------------
# LESSON 7: Exit code semantics
# Why this matters: scripts and CI use exit codes to detect success/failure.
def main() -> int:
    parser = build_parser()
    args = parser.parse_args()
    return handle_command(args)


# -----------------------------------------------------------------------------
# LESSON 8: Guard main entrypoint
# Why this matters: makes file import-safe for reuse and tests.
if __name__ == "__main__":
    raise SystemExit(main())


# LESSON 9 and LESSON 10 are embedded in design:
# - Lesson 9: separate parsing from business logic
# - Lesson 10: return explicit status codes for automation safety

# End of Python CLI Tools 1-10
