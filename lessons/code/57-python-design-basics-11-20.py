"""
PYTHON SOFTWARE DESIGN BASICS (Lessons 11-20)

Suggested use:
1) Run: python3 lessons/code/57-python-design-basics-11-20.py
2) Identify boundaries (validation, use case, adapter, presenter)
3) Modify one adapter and confirm use case code does not change

Extra context:
- lessons/notes/131-clean-architecture-first-principles.md
- lessons/notes/132-separation-of-concerns-gotchas.md
"""

from dataclasses import dataclass
from typing import Protocol


# -----------------------------------------------------------------------------
# LESSON 11: Value object for input validation
# Why this matters: validate once at boundary.
@dataclass(frozen=True)
class Email:
    value: str

    @staticmethod
    def parse(raw: str) -> "Email | None":
        text = raw.strip().lower()
        if "@" not in text or text.startswith("@") or text.endswith("@"):
            return None
        return Email(text)


# -----------------------------------------------------------------------------
# LESSON 12: Domain entity for subscriber
# Why this matters: core business data should be explicit.
@dataclass
class Subscriber:
    id: int
    email: Email
    active: bool = True


# -----------------------------------------------------------------------------
# LESSON 13: Repository protocol
# Why this matters: services depend on interface, not storage choice.
class SubscriberRepo(Protocol):
    def list_subscribers(self) -> list[Subscriber]:
        ...

    def save_subscriber(self, sub: Subscriber) -> None:
        ...


# -----------------------------------------------------------------------------
# LESSON 14: In-memory repo adapter
# Why this matters: simple adapter enables fast tests.
class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_subscribers(self) -> list[Subscriber]:
        return list(self._items)

    def save_subscriber(self, sub: Subscriber) -> None:
        self._items.append(sub)


# -----------------------------------------------------------------------------
# LESSON 15: Use case result object
# Why this matters: explicit success/failure states improve clarity.
@dataclass
class Result:
    ok: bool
    message: str


# -----------------------------------------------------------------------------
# LESSON 16: Use case service with guard clauses
# Why this matters: service encapsulates workflow and rules.
class SubscriberService:
    def __init__(self, repo: SubscriberRepo) -> None:
        self.repo = repo

    def add_subscriber(self, raw_email: str) -> Result:
        email = Email.parse(raw_email)
        if email is None:
            return Result(False, "Invalid email")

        existing = self.repo.list_subscribers()
        if any(s.email.value == email.value for s in existing):
            return Result(False, "Email already exists")

        next_id = 1 if not existing else max(s.id for s in existing) + 1
        self.repo.save_subscriber(Subscriber(id=next_id, email=email))
        return Result(True, f"Subscriber #{next_id} added")


# -----------------------------------------------------------------------------
# LESSON 17: Presenter function
# Why this matters: separate user-facing formatting from logic.
def present_subscribers(items: list[Subscriber]) -> str:
    lines = ["Subscribers"]
    for s in items:
        state = "active" if s.active else "inactive"
        lines.append(f"- #{s.id} {s.email.value} [{state}]")
    return "\n".join(lines)


# -----------------------------------------------------------------------------
# LESSON 18: Controller-like adapter
# Why this matters: map incoming request data to use case call.
def handle_add(service: SubscriberService, payload: dict[str, str]) -> str:
    raw_email = payload.get("email", "")
    result = service.add_subscriber(raw_email)
    return result.message


# -----------------------------------------------------------------------------
# LESSON 19: Replaceable adapter demonstration
# Why this matters: swapping adapters should not break service code.
repo = InMemorySubscriberRepo()
service = SubscriberService(repo)
print("Lesson 19:", handle_add(service, {"email": "mia@example.com"}))
print("Lesson 19:", handle_add(service, {"email": "mia@example.com"}))
print("Lesson 19:", handle_add(service, {"email": "invalid-email"}))


# -----------------------------------------------------------------------------
# LESSON 20: End-to-end report
# Why this matters: final output comes from composed layers.
print("Lesson 20 report:\n" + present_subscribers(repo.list_subscribers()))

# End of Python Software Design Basics 11-20
