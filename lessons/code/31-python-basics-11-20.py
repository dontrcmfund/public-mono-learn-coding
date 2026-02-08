"""
PYTHON BASICS (Lessons 11-20)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/31-python-basics-11-20.py
3) Change one line and predict output before running

Extra context:
- lessons/notes/91-python-common-gotchas.md
- lessons/notes/92-python-design-and-indentation.md
"""

# -----------------------------------------------------------------------------
# LESSON 11: Strings and methods
# Why this matters: most user data starts as text.
message = "  python basics  "
print("Lesson 11:", message.strip().title())

# -----------------------------------------------------------------------------
# LESSON 12: f-strings
# Why this matters: clear output formatting improves debugging.
name = "Kai"
level = 2
print(f"Lesson 12: {name} is at level {level}")

# -----------------------------------------------------------------------------
# LESSON 13: Tuple basics
# Why this matters: tuples represent fixed grouped values.
point = (10, 20)
print("Lesson 13:", point, point[0])

# -----------------------------------------------------------------------------
# LESSON 14: Set basics
# Why this matters: sets help remove duplicates quickly.
nums = [1, 2, 2, 3, 3, 3]
unique_nums = set(nums)
print("Lesson 14:", unique_nums)

# -----------------------------------------------------------------------------
# LESSON 15: While loops
# Why this matters: repeat work until condition changes.
count = 0
while count < 3:
    print("Lesson 15 count:", count)
    count += 1

# -----------------------------------------------------------------------------
# LESSON 16: List comprehension
# Why this matters: concise transformation of list data.
squares = [n * n for n in range(1, 6)]
print("Lesson 16:", squares)

# -----------------------------------------------------------------------------
# LESSON 17: Function defaults
# Why this matters: safer functions with predictable fallbacks.
def greet(person: str, greeting: str = "Hello") -> str:
    return f"{greeting}, {person}!"

print("Lesson 17:", greet("Mia"), greet("Leo", "Hi"))

# -----------------------------------------------------------------------------
# LESSON 18: *args and **kwargs
# Why this matters: flexible function inputs for reusable helpers.
def show_values(*args, **kwargs):
    print("Lesson 18 args:", args)
    print("Lesson 18 kwargs:", kwargs)

show_values(1, 2, 3, role="student", active=True)

# -----------------------------------------------------------------------------
# LESSON 19: Basic exception handling
# Why this matters: handle bad input without crashing.
def safe_divide(a: float, b: float) -> float | None:
    try:
        return a / b
    except ZeroDivisionError:
        return None

print("Lesson 19:", safe_divide(10, 2), safe_divide(10, 0))

# -----------------------------------------------------------------------------
# LESSON 20: File read/write basics
# Why this matters: scripts often store and load data.
file_path = "lessons/code/tmp-python-note.txt"
with open(file_path, "w", encoding="utf-8") as f:
    f.write("Lesson 20: Python file IO works.\n")

with open(file_path, "r", encoding="utf-8") as f:
    print("Lesson 20:", f.read().strip())

# End of Python Basics 11-20
