"""
PYTHON SOFTWARE DESIGN CAPSTONE TESTS (Lessons 1-10)

Suggested use:
1) Run: python3 lessons/code/60-python-design-capstone-tests-1-10.py
2) Read PASS/FAIL output
3) Change one capstone rule and verify failing tests
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

    def replace_subscriber(self, sub: Subscriber) -> None:
        ...


class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_subscribers(self) -> list[Subscriber]:
        return list(self._items)

    def save_subscriber(self, sub: Subscriber) -> None:
        self._items.append(sub)

    def replace_subscriber(self, sub: Subscriber) -> None:
        self._items = [sub if x.id == sub.id else x for x in self._items]


@dataclass
class Result:
    ok: bool
    message: str


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
# LESSON 1: valid email parsing
run_test("Lesson 1 parse valid", Email.parse("mia@example.com") is not None)

# LESSON 2: invalid email parsing
run_test("Lesson 2 parse invalid", Email.parse("invalid") is None)

repo = InMemorySubscriberRepo()
service = SubscriptionService(repo)

# LESSON 3: add success
r1 = service.add("mia@example.com")
run_test("Lesson 3 add success", r1.ok)

# LESSON 4: duplicate blocked
r2 = service.add("mia@example.com")
run_test("Lesson 4 duplicate blocked", not r2.ok)

# LESSON 5: second add increments id
r3 = service.add("ana@example.com")
run_test("Lesson 5 second add", r3.ok and [s.id for s in repo.list_subscribers()] == [1, 2])

# LESSON 6: deactivate success
r4 = service.deactivate("mia@example.com")
run_test("Lesson 6 deactivate success", r4.ok)

# LESSON 7: deactivate already inactive
r5 = service.deactivate("mia@example.com")
run_test("Lesson 7 already inactive", not r5.ok)

# LESSON 8: deactivate not found
r6 = service.deactivate("nobody@example.com")
run_test("Lesson 8 deactivate missing", not r6.ok)

# LESSON 9: list copy safety
items = repo.list_subscribers()
items.clear()
run_test("Lesson 9 copy safety", len(repo.list_subscribers()) == 2)

# LESSON 10: completion marker
print("Lesson 10: design capstone tests complete")

# End of Python Software Design Capstone Tests 1-10
