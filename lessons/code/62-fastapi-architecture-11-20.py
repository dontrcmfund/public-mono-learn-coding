"""
FASTAPI ARCHITECTURE (Lessons 11-20)

Suggested use:
1) Read this file to see layered patterns
2) Run later with uvicorn when dependencies are installed

Extra context:
- lessons/notes/141-fastapi-layering-and-dependencies.md
- lessons/notes/142-fastapi-gotchas.md
"""

from dataclasses import dataclass
from typing import Protocol

from fastapi import Depends, FastAPI, HTTPException, status
from pydantic import BaseModel, Field

app = FastAPI(title="FastAPI Architecture Lessons", version="0.1.0")


# -----------------------------------------------------------------------------
# LESSON 11: Domain model
# Why this matters: domain remains framework-agnostic.
@dataclass
class Subscriber:
    id: int
    email: str
    active: bool = True


# -----------------------------------------------------------------------------
# LESSON 12: Domain validation helper
# Why this matters: reusable rules should be central.
def is_valid_email(value: str) -> bool:
    text = value.strip().lower()
    return "@" in text and not text.startswith("@") and not text.endswith("@")


# -----------------------------------------------------------------------------
# LESSON 13: Repository port
# Why this matters: service logic should not care about storage details.
class SubscriberRepo(Protocol):
    def list_subscribers(self) -> list[Subscriber]:
        ...

    def save_subscriber(self, email: str) -> Subscriber:
        ...

    def deactivate_by_email(self, email: str) -> Subscriber | None:
        ...


# -----------------------------------------------------------------------------
# LESSON 14: In-memory adapter
# Why this matters: adapter can be swapped later (DB, API, etc.).
class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_subscribers(self) -> list[Subscriber]:
        return list(self._items)

    def save_subscriber(self, email: str) -> Subscriber:
        next_id = 1 if not self._items else max(s.id for s in self._items) + 1
        sub = Subscriber(id=next_id, email=email)
        self._items.append(sub)
        return sub

    def deactivate_by_email(self, email: str) -> Subscriber | None:
        for i, sub in enumerate(self._items):
            if sub.email == email:
                if not sub.active:
                    return sub
                updated = Subscriber(id=sub.id, email=sub.email, active=False)
                self._items[i] = updated
                return updated
        return None


# -----------------------------------------------------------------------------
# LESSON 15: Service result object
# Why this matters: explicit outcomes simplify route mapping.
@dataclass
class ServiceResult:
    ok: bool
    message: str
    subscriber: Subscriber | None = None


# -----------------------------------------------------------------------------
# LESSON 16: Service workflow
# Why this matters: workflows belong in service layer.
class SubscriberService:
    def __init__(self, repo: SubscriberRepo) -> None:
        self.repo = repo

    def add(self, email_raw: str) -> ServiceResult:
        email = email_raw.strip().lower()
        if not is_valid_email(email):
            return ServiceResult(False, "Invalid email")

        existing = self.repo.list_subscribers()
        if any(s.email == email for s in existing):
            return ServiceResult(False, "Duplicate email")

        sub = self.repo.save_subscriber(email)
        return ServiceResult(True, "Created", sub)

    def deactivate(self, email_raw: str) -> ServiceResult:
        email = email_raw.strip().lower()
        if not is_valid_email(email):
            return ServiceResult(False, "Invalid email")

        sub = self.repo.deactivate_by_email(email)
        if sub is None:
            return ServiceResult(False, "Not found")
        if not sub.active:
            return ServiceResult(False, "Already inactive")
        return ServiceResult(True, "Deactivated", sub)


# -----------------------------------------------------------------------------
# LESSON 17: API models
# Why this matters: request/response contracts stay explicit.
class SubscriberCreateRequest(BaseModel):
    email: str = Field(min_length=3, max_length=200)


class SubscriberResponse(BaseModel):
    id: int
    email: str
    active: bool


# -----------------------------------------------------------------------------
# LESSON 18: Dependency wiring
# Why this matters: dependency injection keeps handlers thin.
repo_singleton = InMemorySubscriberRepo()


def get_subscriber_service() -> SubscriberService:
    return SubscriberService(repo_singleton)


# -----------------------------------------------------------------------------
# LESSON 19: Route handlers with status mapping
# Why this matters: translate service outcomes into stable HTTP behavior.
@app.post("/subscribers", response_model=SubscriberResponse, status_code=status.HTTP_201_CREATED)
def create_subscriber(
    payload: SubscriberCreateRequest,
    service: SubscriberService = Depends(get_subscriber_service),
) -> SubscriberResponse:
    result = service.add(payload.email)
    if not result.ok:
        code = status.HTTP_400_BAD_REQUEST if result.message != "Duplicate email" else status.HTTP_409_CONFLICT
        raise HTTPException(status_code=code, detail=result.message)

    sub = result.subscriber
    assert sub is not None
    return SubscriberResponse(id=sub.id, email=sub.email, active=sub.active)


@app.get("/subscribers", response_model=list[SubscriberResponse])
def list_subscribers(service: SubscriberService = Depends(get_subscriber_service)) -> list[SubscriberResponse]:
    return [SubscriberResponse(id=s.id, email=s.email, active=s.active) for s in service.repo.list_subscribers()]


# -----------------------------------------------------------------------------
# LESSON 20: Architecture checklist endpoint
# Why this matters: explicit reminders reinforce design boundaries.
@app.get("/architecture-check")
def architecture_check() -> dict[str, str]:
    return {
        "domain": "separate",
        "service": "separate",
        "adapter": "replaceable",
        "routes": "thin",
    }


# End of FastAPI Architecture 11-20
