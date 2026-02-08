"""
PYTHON SMALL PROJECTS (Lessons 1-10)

Suggested use:
1) Open this file in VS Code
2) Run: python3 lessons/code/32-python-small-projects-1-10.py
3) For each lesson, change one rule and predict new output

Extra context:
- lessons/notes/93-python-small-projects-principles.md
- lessons/notes/91-python-common-gotchas.md
"""

# -----------------------------------------------------------------------------
# LESSON 1 PROJECT: Expense total
# Why this matters: summing values is common in personal automation.
expenses = [12.5, 30.0, 7.25]
print("Lesson 1:", sum(expenses))

# -----------------------------------------------------------------------------
# LESSON 2 PROJECT: Case-insensitive search
# Why this matters: user input rarely matches exact case.
def search_items(items: list[str], query: str) -> list[str]:
    q = query.strip().lower()
    return [item for item in items if q in item.lower()]

print("Lesson 2:", search_items(["Python", "Docker", "TypeScript"], "py"))

# -----------------------------------------------------------------------------
# LESSON 3 PROJECT: Grade classifier
# Why this matters: rule-based mapping appears in many systems.
def classify_grade(score: int) -> str:
    if score >= 90:
        return "A"
    if score >= 80:
        return "B"
    if score >= 70:
        return "C"
    return "D"

print("Lesson 3:", classify_grade(84))

# -----------------------------------------------------------------------------
# LESSON 4 PROJECT: Frequency counter
# Why this matters: counting categories is core analytics.
def count_words(words: list[str]) -> dict[str, int]:
    counts: dict[str, int] = {}
    for word in words:
        counts[word] = counts.get(word, 0) + 1
    return counts

print("Lesson 4:", count_words(["a", "b", "a", "c", "b", "a"]))

# -----------------------------------------------------------------------------
# LESSON 5 PROJECT: Unique sorted tags
# Why this matters: normalize and deduplicate noisy input.
def normalize_tags(tags: list[str]) -> list[str]:
    cleaned = {tag.strip().lower() for tag in tags if tag.strip()}
    return sorted(cleaned)

print("Lesson 5:", normalize_tags(["Python", " python ", "AI", "ai", ""]))

# -----------------------------------------------------------------------------
# LESSON 6 PROJECT: Basic CSV line parser
# Why this matters: quick data scripts often start from simple text rows.
def parse_csv_line(line: str) -> list[str]:
    return [part.strip() for part in line.split(",")]

print("Lesson 6:", parse_csv_line("name, age, city"))

# -----------------------------------------------------------------------------
# LESSON 7 PROJECT: Todo manager (in-memory)
# Why this matters: state + actions model real app behavior.
def add_todo(todos: list[dict[str, object]], title: str) -> list[dict[str, object]]:
    next_id = 1 if not todos else int(todos[-1]["id"]) + 1
    return todos + [{"id": next_id, "title": title, "done": False}]

def complete_todo(todos: list[dict[str, object]], todo_id: int) -> list[dict[str, object]]:
    updated: list[dict[str, object]] = []
    for todo in todos:
        if int(todo["id"]) == todo_id:
            updated.append({**todo, "done": True})
        else:
            updated.append(todo)
    return updated

todo_state: list[dict[str, object]] = []
todo_state = add_todo(todo_state, "Read Python lesson")
todo_state = add_todo(todo_state, "Practice project")
todo_state = complete_todo(todo_state, 1)
print("Lesson 7:", todo_state)

# -----------------------------------------------------------------------------
# LESSON 8 PROJECT: Safe integer parse
# Why this matters: input parsing failures are common.
def parse_int(value: str) -> int | None:
    try:
        return int(value)
    except ValueError:
        return None

print("Lesson 8:", parse_int("42"), parse_int("oops"))

# -----------------------------------------------------------------------------
# LESSON 9 PROJECT: Retry decision helper
# Why this matters: robust scripts need retry policies.
def retry_delay(attempt: int, max_attempts: int) -> int | None:
    if attempt >= max_attempts:
        return None
    return 2 ** (attempt - 1)

print("Lesson 9:", retry_delay(1, 4), retry_delay(4, 4))

# -----------------------------------------------------------------------------
# LESSON 10 PROJECT: File report generator
# Why this matters: transforming data into text reports is practical automation.
def build_report(items: list[str]) -> str:
    lines = [f"- {item}" for item in items]
    return "Report:\n" + "\n".join(lines)

report_text = build_report(["Task A done", "Task B pending"])
print("Lesson 10:\n" + report_text)

# End of Python Small Projects 1-10
