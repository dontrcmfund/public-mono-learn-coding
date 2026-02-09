"""
PYTHON SOFTWARE DESIGN TESTING (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/58-python-design-testing-1-10.py
2) Read PASS/FAIL output
3) Change one use-case rule and verify failing tests

Extra context:
- lessons/notes/133-design-testing-principles.md
"""

from dataclasses import dataclass
from typing import Protocol


def run_test(name: str, condition: bool) -> None:
    print(("PASS" if condition else "FAIL") + f": {name}")


@dataclass(frozen=True)
class Email:
    value: str

    @staticmethod
    def parse(raw: str) -> "Email | None":
        text = raw.strip().lower()
        if "@" not in text or text.startswith("@") or text.endswith("@"):
            return None
        return Email(text)


@dataclass
class Subscriber:
    id: int
    email: Email
    active: bool = True


class SubscriberRepo(Protocol):
    def list_subscribers(self) -> list[Subscriber]:
        ...

    def save_subscriber(self, sub: Subscriber) -> None:
        ...


class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_subscribers(self) -> list[Subscriber]:
        return list(self._items)

    def save_subscriber(self, sub: Subscriber) -> None:
        self._items.append(sub)


@dataclass
class Result:
    ok: bool
    message: str


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
# LESSON 1: email parse success
run_test("Lesson 1 email parse", Email.parse("mia@example.com") is not None)

# LESSON 2: email parse failure
run_test("Lesson 2 invalid email", Email.parse("invalid") is None)

# LESSON 3: add subscriber success
repo = InMemorySubscriberRepo()
service = SubscriberService(repo)
r1 = service.add_subscriber("mia@example.com")
run_test("Lesson 3 add success", r1.ok)

# LESSON 4: duplicate prevention
r2 = service.add_subscriber("mia@example.com")
run_test("Lesson 4 duplicate blocked", not r2.ok and "exists" in r2.message)

# LESSON 5: id increments
service.add_subscriber("ana@example.com")
ids = [s.id for s in repo.list_subscribers()]
run_test("Lesson 5 id increments", ids == [1, 2])

# LESSON 6: repo adapter contract list copy
listed = repo.list_subscribers()
listed.clear()
run_test("Lesson 6 adapter returns copy", len(repo.list_subscribers()) == 2)

# LESSON 7: boundary validation stays in service
bad = service.add_subscriber("@bad")
run_test("Lesson 7 boundary validation", not bad.ok)

# LESSON 8: layer independence signal
run_test("Lesson 8 service unaffected by presentation", hasattr(service, "add_subscriber"))

# LESSON 9: active default true
run_test("Lesson 9 active default", all(s.active for s in repo.list_subscribers()))

# LESSON 10: completion marker
print("Lesson 10: design testing suite complete")

# End of Python Software Design Testing 1-10
