"""
FASTAPI CAPSTONE (Lessons 1-10)

Project: subscription API with clean architecture layering

Suggested use:
1) Read structure in this single file (acts like multi-module layout)
2) Install dependencies later: fastapi, uvicorn
3) Run later: uvicorn lessons.code.64-fastapi-capstone-1-10:app --reload

Extra context:
- lessons/notes/141-fastapi-layering-and-dependencies.md
- lessons/notes/142-fastapi-gotchas.md
"""

from dataclasses import dataclass
from typing import Protocol

from fastapi import Depends, FastAPI, HTTPException, status
from pydantic import BaseModel, Field

app = FastAPI(title="FastAPI Capstone", version="0.1.0")


# -----------------------------------------------------------------------------
# LESSON 1: Domain rules
# Why this matters: core validation must not depend on HTTP framework.
def normalize_email(raw: str) -> str:
    return raw.strip().lower()


def is_valid_email(email: str) -> bool:
    return "@" in email and not email.startswith("@") and not email.endswith("@")


# -----------------------------------------------------------------------------
# LESSON 2: Domain entity
# Why this matters: business object shape stays explicit.
@dataclass
class Subscriber:
    id: int
    email: str
    active: bool = True


# -----------------------------------------------------------------------------
# LESSON 3: Repository port
# Why this matters: service works with abstraction, not concrete persistence.
class SubscriberRepo(Protocol):
    def list_all(self) -> list[Subscriber]:
        ...

    def save(self, subscriber: Subscriber) -> None:
        ...

    def replace(self, subscriber: Subscriber) -> None:
        ...


# -----------------------------------------------------------------------------
# LESSON 4: In-memory adapter
# Why this matters: fast local adapter enables simple startup/testing.
class InMemorySubscriberRepo:
    def __init__(self) -> None:
        self._items: list[Subscriber] = []

    def list_all(self) -> list[Subscriber]:
        return list(self._items)

    def save(self, subscriber: Subscriber) -> None:
        self._items.append(subscriber)

    def replace(self, subscriber: Subscriber) -> None:
        self._items = [subscriber if x.id == subscriber.id else x for x in self._items]


# -----------------------------------------------------------------------------
# LESSON 5: Service result
# Why this matters: explicit outcomes simplify API error mapping.
@dataclass
class Result:
    ok: bool
    message: str
    subscriber: Subscriber | None = None


# -----------------------------------------------------------------------------
# LESSON 6: Use-case service
# Why this matters: workflow logic belongs here, not in route handlers.
class SubscriberService:
    def __init__(self, repo: SubscriberRepo) -> None:
        self.repo = repo

    def create(self, raw_email: str) -> Result:
        email = normalize_email(raw_email)
        if not is_valid_email(email):
            return Result(False, "Invalid email")

        existing = self.repo.list_all()
        if any(s.email == email for s in existing):
            return Result(False, "Duplicate email")

        next_id = 1 if not existing else max(s.id for s in existing) + 1
        sub = Subscriber(id=next_id, email=email)
        self.repo.save(sub)
        return Result(True, "Created", sub)

    def deactivate(self, raw_email: str) -> Result:
        email = normalize_email(raw_email)
        if not is_valid_email(email):
            return Result(False, "Invalid email")

        for sub in self.repo.list_all():
            if sub.email == email:
                if not sub.active:
                    return Result(False, "Already inactive")
                updated = Subscriber(id=sub.id, email=sub.email, active=False)
                self.repo.replace(updated)
                return Result(True, "Deactivated", updated)
        return Result(False, "Not found")


# -----------------------------------------------------------------------------
# LESSON 7: API models
# Why this matters: request/response contracts remain clear and versionable.
class SubscriberCreateRequest(BaseModel):
    email: str = Field(min_length=3, max_length=200)


class SubscriberActionRequest(BaseModel):
    email: str = Field(min_length=3, max_length=200)


class SubscriberResponse(BaseModel):
    id: int
    email: str
    active: bool


# -----------------------------------------------------------------------------
# LESSON 8: Dependency injection
# Why this matters: central wiring keeps handlers thin.
repo_singleton = InMemorySubscriberRepo()


def get_service() -> SubscriberService:
    return SubscriberService(repo_singleton)


# -----------------------------------------------------------------------------
# LESSON 9: Routes with error mapping
# Why this matters: stable HTTP behavior is part of API contract.
@app.get("/health")
def health() -> dict[str, str]:
    return {"status": "ok"}


@app.post("/subscribers", response_model=SubscriberResponse, status_code=status.HTTP_201_CREATED)
def create_subscriber(
    payload: SubscriberCreateRequest,
    service: SubscriberService = Depends(get_service),
) -> SubscriberResponse:
    result = service.create(payload.email)
    if not result.ok:
        code = status.HTTP_400_BAD_REQUEST if result.message != "Duplicate email" else status.HTTP_409_CONFLICT
        raise HTTPException(status_code=code, detail=result.message)

    sub = result.subscriber
    assert sub is not None
    return SubscriberResponse(id=sub.id, email=sub.email, active=sub.active)


@app.post("/subscribers/deactivate", response_model=SubscriberResponse)
def deactivate_subscriber(
    payload: SubscriberActionRequest,
    service: SubscriberService = Depends(get_service),
) -> SubscriberResponse:
    result = service.deactivate(payload.email)
    if not result.ok:
        mapping = {
            "Invalid email": status.HTTP_400_BAD_REQUEST,
            "Not found": status.HTTP_404_NOT_FOUND,
            "Already inactive": status.HTTP_409_CONFLICT,
        }
        raise HTTPException(status_code=mapping.get(result.message, status.HTTP_400_BAD_REQUEST), detail=result.message)

    sub = result.subscriber
    assert sub is not None
    return SubscriberResponse(id=sub.id, email=sub.email, active=sub.active)


@app.get("/subscribers", response_model=list[SubscriberResponse])
def list_subscribers(service: SubscriberService = Depends(get_service)) -> list[SubscriberResponse]:
    return [SubscriberResponse(id=s.id, email=s.email, active=s.active) for s in service.repo.list_all()]


# -----------------------------------------------------------------------------
# LESSON 10: Architecture marker endpoint
# Why this matters: makes layer responsibilities explicit to learners.
@app.get("/architecture")
def architecture() -> dict[str, str]:
    return {
        "domain": "email rules + entity",
        "service": "create/deactivate workflows",
        "repo": "in-memory adapter",
        "routes": "HTTP mapping only",
    }


# End of FastAPI Capstone 1-10
