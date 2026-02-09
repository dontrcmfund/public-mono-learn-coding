"""
PYTHON SOFTWARE DESIGN BASICS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/56-python-design-basics-1-10.py
2) Identify which lines are domain logic vs IO
3) Change one layer and verify others stay stable

Extra context:
- lessons/notes/130-what-is-software-design.md
- lessons/notes/131-clean-architecture-first-principles.md
- lessons/notes/132-separation-of-concerns-gotchas.md
"""

from dataclasses import dataclass
from typing import Protocol


# -----------------------------------------------------------------------------
# LESSON 1: Domain entity
# Why this matters: entities model business concepts independent of tools.
@dataclass
class Task:
    id: int
    title: str
    done: bool = False


# -----------------------------------------------------------------------------
# LESSON 2: Domain rule method
# Why this matters: business rules belong with domain concepts.
def complete_task(task: Task) -> Task:
    return Task(id=task.id, title=task.title, done=True)


# -----------------------------------------------------------------------------
# LESSON 3: Repository interface (port)
# Why this matters: use cases depend on abstraction, not storage details.
class TaskRepository(Protocol):
    def list_tasks(self) -> list[Task]:
        ...

    def save_task(self, task: Task) -> None:
        ...


# -----------------------------------------------------------------------------
# LESSON 4: In-memory adapter
# Why this matters: adapters implement interfaces and isolate infrastructure.
class InMemoryTaskRepo:
    def __init__(self) -> None:
        self._tasks: list[Task] = []

    def list_tasks(self) -> list[Task]:
        return list(self._tasks)

    def save_task(self, task: Task) -> None:
        self._tasks.append(task)


# -----------------------------------------------------------------------------
# LESSON 5: Use case service
# Why this matters: use cases orchestrate domain + repository behavior.
class TaskService:
    def __init__(self, repo: TaskRepository) -> None:
        self.repo = repo

    def add_task(self, title: str) -> Task:
        existing = self.repo.list_tasks()
        next_id = 1 if not existing else max(t.id for t in existing) + 1
        task = Task(id=next_id, title=title.strip())
        self.repo.save_task(task)
        return task


# -----------------------------------------------------------------------------
# LESSON 6: Separate presenter
# Why this matters: output formatting should not pollute business logic.
def present_tasks(tasks: list[Task]) -> str:
    lines = ["Tasks"]
    for t in tasks:
        status = "done" if t.done else "open"
        lines.append(f"- #{t.id} {t.title} [{status}]")
    return "\n".join(lines)


# -----------------------------------------------------------------------------
# LESSON 7: Command-like handler
# Why this matters: handlers map external requests to use case calls.
def handle_add_task(service: TaskService, title: str) -> str:
    task = service.add_task(title)
    return f"Added task #{task.id}: {task.title}"


# -----------------------------------------------------------------------------
# LESSON 8: Pure vs side effect example
# Why this matters: pure functions are easier to test.
def count_open_tasks(tasks: list[Task]) -> int:
    return sum(1 for t in tasks if not t.done)


# -----------------------------------------------------------------------------
# LESSON 9: Composition root
# Why this matters: object wiring belongs in one top-level place.
repo = InMemoryTaskRepo()
service = TaskService(repo)
print("Lesson 9:", handle_add_task(service, "Write design notes"))
print("Lesson 9:", handle_add_task(service, "Review architecture"))


# -----------------------------------------------------------------------------
# LESSON 10: End-to-end flow
# Why this matters: demonstrate layers working together without coupling.
all_tasks = repo.list_tasks()
all_tasks[0] = complete_task(all_tasks[0])
print("Lesson 10 open count:", count_open_tasks(all_tasks))
print("Lesson 10 report:\n" + present_tasks(all_tasks))

# End of Python Software Design Basics 1-10
