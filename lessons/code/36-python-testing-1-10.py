"""
PYTHON TESTING (Lessons 1-10)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/36-python-testing-1-10.py
3) Read the test output and change one case at a time

Extra context:
- lessons/notes/97-python-testing-principles.md
"""


def add(a: int, b: int) -> int:
    return a + b


def safe_divide(a: float, b: float) -> float | None:
    if b == 0:
        return None
    return a / b


def normalize_name(name: str) -> str:
    return " ".join(name.strip().split()).title()


def run_test(name: str, condition: bool) -> None:
    if condition:
        print(f"PASS: {name}")
    else:
        print(f"FAIL: {name}")


# -----------------------------------------------------------------------------
# LESSON 1: Arrange-Act-Assert pattern
# Why this matters: predictable test structure improves readability.
result = add(2, 3)
run_test("Lesson 1 add(2,3)==5", result == 5)


# -----------------------------------------------------------------------------
# LESSON 2: Normal case tests
# Why this matters: test expected everyday behavior first.
run_test("Lesson 2 add(0,0)==0", add(0, 0) == 0)


# -----------------------------------------------------------------------------
# LESSON 3: Edge case tests
# Why this matters: edge cases are where bugs hide.
run_test("Lesson 3 safe_divide by zero", safe_divide(5, 0) is None)


# -----------------------------------------------------------------------------
# LESSON 4: String normalization test
# Why this matters: text cleaning logic should be explicit.
run_test("Lesson 4 normalize_name", normalize_name("  aLEx   kim ") == "Alex Kim")


# -----------------------------------------------------------------------------
# LESSON 5: Group related assertions
# Why this matters: one function can have multiple expected properties.
value = safe_divide(9, 3)
run_test("Lesson 5 safe_divide returns value", value is not None)
run_test("Lesson 5 safe_divide value correct", value == 3)


# -----------------------------------------------------------------------------
# LESSON 6: Regression test mindset
# Why this matters: tests protect fixed bugs from returning.
run_test("Lesson 6 regression add negatives", add(-2, -3) == -5)


# -----------------------------------------------------------------------------
# LESSON 7: Deterministic tests
# Why this matters: stable tests avoid random failures.
run_test("Lesson 7 deterministic", normalize_name("mia") == "Mia")


# -----------------------------------------------------------------------------
# LESSON 8: Failure message quality
# Why this matters: clear failures speed up debugging.
expected = "Ana Lee"
actual = normalize_name("ana   lee")
run_test(f"Lesson 8 expected '{expected}' got '{actual}'", actual == expected)


# -----------------------------------------------------------------------------
# LESSON 9: Test data tables
# Why this matters: compactly test many similar cases.
cases = [
    ("bob", "Bob"),
    ("  sAra ", "Sara"),
    ("mike jones", "Mike Jones"),
]
for i, (raw, wanted) in enumerate(cases, start=1):
    run_test(f"Lesson 9 case {i}", normalize_name(raw) == wanted)


# -----------------------------------------------------------------------------
# LESSON 10: Tiny test summary
# Why this matters: feedback loops should be easy to scan.
print("Lesson 10: tests complete")

# End of Python Testing 1-10
