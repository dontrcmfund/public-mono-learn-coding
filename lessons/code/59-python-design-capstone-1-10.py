"""
PYTHON SOFTWARE DESIGN CAPSTONE (Lessons 1-10)

Project: Newsletter subscription service

Suggested use:
1) Run: python3 lessons/code/59-python-design-capstone-1-10.py
2) Read service outputs and report
3) Swap adapter implementation later without changing use-case logic

Extra context:
- lessons/notes/134-design-capstone-plan.md
- lessons/notes/131-clean-architecture-first-principles.md
"""

from dataclasses import dataclass
from typing import Protocol


# -----------------------------------------------------------------------------
# LESSON 1: domain value object
# Why this matters: validate core value once.
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
# LESSON 2: domain entity
# Why this matters: entity represents business concept and state.
@dataclass
class Subscriber:
    id: int
    email: Email
    active: bool = True


# -----------------------------------------------------------------------------
# LESSON 3: repository interface
# Why this matters: use-case stays storage-agnostic.
class SubscriberRepo(Protocol):
    def list_subscribers(self) -> list[Subscriber]:
        ...

    def save_subscriber(self, sub: Subscriber) -> None:
        ...

    def replace_subscriber(self, sub: Subscriber) -> None:
        ...


# -----------------------------------------------------------------------------
# LESSON 4: in-memory adapter
# Why this matters: fast local adapter for development/testing.
class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_subscribers(self) -> list[Subscriber]:
        return list(self._items)

    def save_subscriber(self, sub: Subscriber) -> None:
        self._items.append(sub)

    def replace_subscriber(self, sub: Subscriber) -> None:
        self._items = [sub if x.id == sub.id else x for x in self._items]


# -----------------------------------------------------------------------------
# LESSON 5: result object
# Why this matters: explicit success/failure outcomes simplify handlers.
@dataclass
class Result:
    ok: bool
    message: str


# -----------------------------------------------------------------------------
# LESSON 6: use-case service
# Why this matters: business workflow belongs here, not in adapters.
class SubscriptionService:
    def __init__(self, repo: SubscriberRepo) -> None:
        self.repo = repo

    def add(self, raw_email: str) -> Result:
        email = Email.parse(raw_email)
        if email is None:
            return Result(False, "Invalid email")

        existing = self.repo.list_subscribers()
        if any(s.email.value == email.value for s in existing):
            return Result(False, "Duplicate email")

        next_id = 1 if not existing else max(s.id for s in existing) + 1
        self.repo.save_subscriber(Subscriber(id=next_id, email=email))
        return Result(True, f"Added #{next_id}")

    def deactivate(self, email_raw: str) -> Result:
        email = Email.parse(email_raw)
        if email is None:
            return Result(False, "Invalid email")

        for sub in self.repo.list_subscribers():
            if sub.email.value == email.value:
                if not sub.active:
                    return Result(False, "Already inactive")
                self.repo.replace_subscriber(Subscriber(id=sub.id, email=sub.email, active=False))
                return Result(True, "Deactivated")
        return Result(False, "Not found")


# -----------------------------------------------------------------------------
# LESSON 7: presenter
# Why this matters: output formatting remains separate.
def present(subscribers: list[Subscriber]) -> str:
    lines = ["Subscribers"]
    for s in subscribers:
        state = "active" if s.active else "inactive"
        lines.append(f"- #{s.id} {s.email.value} [{state}]")
    return "\n".join(lines)


# -----------------------------------------------------------------------------
# LESSON 8: controller-like handlers
# Why this matters: boundary layer adapts request payloads.
def handle_add(service: SubscriptionService, payload: dict[str, str]) -> str:
    return service.add(payload.get("email", "")).message


def handle_deactivate(service: SubscriptionService, payload: dict[str, str]) -> str:
    return service.deactivate(payload.get("email", "")).message


# -----------------------------------------------------------------------------
# LESSON 9: composition root
# Why this matters: wiring should be centralized.
repo = InMemorySubscriberRepo()
service = SubscriptionService(repo)
print("Lesson 9:", handle_add(service, {"email": "mia@example.com"}))
print("Lesson 9:", handle_add(service, {"email": "ana@example.com"}))
print("Lesson 9:", handle_deactivate(service, {"email": "mia@example.com"}))


# -----------------------------------------------------------------------------
# LESSON 10: final report
# Why this matters: final output should reflect all workflow changes.
print("Lesson 10 report:\n" + present(repo.list_subscribers()))

# End of Python Software Design Capstone 1-10
