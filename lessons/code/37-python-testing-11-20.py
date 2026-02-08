"""
PYTHON TESTING (Lessons 11-20)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/37-python-testing-11-20.py
3) Change one behavior and watch which tests fail

Extra context:
- lessons/notes/97-python-testing-principles.md
- lessons/notes/99-python-testing-gotchas.md
"""

from pathlib import Path
import tempfile


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


# -----------------------------------------------------------------------------
# LESSON 11: Test pure transformation function
# Why this matters: pure functions are easiest and safest to test.
def normalize_tags(tags: list[str]) -> list[str]:
    return sorted({tag.strip().lower() for tag in tags if tag.strip()})

run_test("Lesson 11 normalize tags", normalize_tags([" Py ", "py", "AI"]) == ["ai", "py"])


# -----------------------------------------------------------------------------
# LESSON 12: Test edge cases explicitly
# Why this matters: empty input bugs are common.
run_test("Lesson 12 empty input", normalize_tags([]) == [])


# -----------------------------------------------------------------------------
# LESSON 13: File output test with temp directory
# Why this matters: temporary paths keep tests isolated.
def save_report(path: Path, text: str) -> None:
    path.write_text(text, encoding="utf-8")

with tempfile.TemporaryDirectory() as tmp:
    file_path = Path(tmp) / "report.txt"
    save_report(file_path, "ok")
    run_test("Lesson 13 file exists", file_path.exists())
    run_test("Lesson 13 file content", file_path.read_text(encoding="utf-8") == "ok")


# -----------------------------------------------------------------------------
# LESSON 14: Parse helper tests
# Why this matters: parsing failures should be explicit.
def parse_int(value: str) -> int | None:
    try:
        return int(value)
    except ValueError:
        return None

run_test("Lesson 14 parse good", parse_int("42") == 42)
run_test("Lesson 14 parse bad", parse_int("x") is None)


# -----------------------------------------------------------------------------
# LESSON 15: Test branching behavior
# Why this matters: each branch should have coverage.
def classify(score: int) -> str:
    if score >= 90:
        return "A"
    if score >= 80:
        return "B"
    return "C"

run_test("Lesson 15 branch A", classify(95) == "A")
run_test("Lesson 15 branch B", classify(84) == "B")
run_test("Lesson 15 branch C", classify(60) == "C")


# -----------------------------------------------------------------------------
# LESSON 16: Test stable sorting outcome
# Why this matters: deterministic ordering avoids flaky behavior.
def rank(values: list[int]) -> list[int]:
    return sorted(values, reverse=True)

run_test("Lesson 16 ranking", rank([2, 9, 3]) == [9, 3, 2])


# -----------------------------------------------------------------------------
# LESSON 17: Retry decision tests
# Why this matters: reliability rules should be proven.
def should_retry(status_code: int, attempt: int, max_attempts: int) -> bool:
    if attempt >= max_attempts:
        return False
    return status_code in {429, 500, 502, 503, 504}

run_test("Lesson 17 retry 503", should_retry(503, 1, 3) is True)
run_test("Lesson 17 no retry 404", should_retry(404, 1, 3) is False)
run_test("Lesson 17 max attempts", should_retry(503, 3, 3) is False)


# -----------------------------------------------------------------------------
# LESSON 18: Test report line generator
# Why this matters: output formatting needs contract checks.
def to_lines(items: list[dict[str, str]]) -> list[str]:
    return [f"- {item['name']}: {item['status']}" for item in items]

lines = to_lines([{"name": "web", "status": "ok"}])
run_test("Lesson 18 line format", lines == ["- web: ok"])


# -----------------------------------------------------------------------------
# LESSON 19: Negative test for missing keys
# Why this matters: tests should catch malformed data assumptions.
try:
    to_lines([{"name": "db"}])
    run_test("Lesson 19 malformed item", False)
except KeyError:
    run_test("Lesson 19 malformed item", True)


# -----------------------------------------------------------------------------
# LESSON 20: Mini suite completion marker
# Why this matters: clear run end helps scanning logs.
print("Lesson 20: testing suite 11-20 complete")

# End of Python Testing 11-20
