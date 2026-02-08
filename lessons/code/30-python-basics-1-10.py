"""
PYTHON BASICS (Lessons 1-10)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/30-python-basics-1-10.py
3) Change one line and predict output before running

Extra context:
- lessons/notes/90-what-is-python.md
- lessons/notes/91-python-common-gotchas.md
- lessons/notes/78-how-to-study-this-curriculum.md
"""

# -----------------------------------------------------------------------------
# LESSON 1: Print output
# Why this matters: feedback is how you see program behavior.
print("Lesson 1: Hello, Python")

# -----------------------------------------------------------------------------
# LESSON 2: Variables and values
# Why this matters: named values are the base of all programs.
name = "Alex"
attempts = 2
is_learning = True
print("Lesson 2:", name, attempts, is_learning)

# -----------------------------------------------------------------------------
# LESSON 3: Basic data types
# Why this matters: each type supports different operations.
text = "hello"
number = 42
ratio = 3.14
print("Lesson 3:", type(text), type(number), type(ratio))

# -----------------------------------------------------------------------------
# LESSON 4: Basic math
# Why this matters: many programs are math + decisions.
a = 10
b = 3
print("Lesson 4:", a + b, a - b, a * b, a / b, a // b, a % b)

# -----------------------------------------------------------------------------
# LESSON 5: Comparisons and booleans
# Why this matters: decisions need True/False checks.
print("Lesson 5:", a > b, a == b, a != b)

# -----------------------------------------------------------------------------
# LESSON 6: If/elif/else
# Why this matters: programs react to conditions.
score = 82
if score >= 90:
    print("Lesson 6: grade A")
elif score >= 80:
    print("Lesson 6: grade B")
else:
    print("Lesson 6: grade C")

# -----------------------------------------------------------------------------
# LESSON 7: Lists
# Why this matters: lists hold related values.
colors = ["red", "green", "blue"]
colors.append("yellow")
print("Lesson 7:", colors, colors[0], len(colors))

# -----------------------------------------------------------------------------
# LESSON 8: Dictionaries
# Why this matters: dicts store labeled data.
profile = {"name": "Mia", "level": 1}
profile["level"] = 2
print("Lesson 8:", profile)

# -----------------------------------------------------------------------------
# LESSON 9: Loops
# Why this matters: loops reduce repetition.
for color in colors:
    print("Lesson 9 color:", color)

# -----------------------------------------------------------------------------
# LESSON 10: Functions
# Why this matters: functions package reusable logic.
def greet(person_name: str) -> str:
    return f"Hello, {person_name}!"

print("Lesson 10:", greet("Avery"))

# End of Python Basics 1-10
