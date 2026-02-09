"""
FASTAPI BASICS (Lessons 1-10)

Suggested use:
1) Read this file first
2) Install dependencies later when ready: fastapi, uvicorn
3) Run app (later): uvicorn lessons.code.61-fastapi-basics-1-10:app --reload

Extra context:
- lessons/notes/140-what-is-fastapi.md
- lessons/notes/141-fastapi-layering-and-dependencies.md
"""

from dataclasses import dataclass
from typing import Protocol

from fastapi import Depends, FastAPI, HTTPException
from pydantic import BaseModel, Field

app = FastAPI(title="FastAPI Lessons", version="0.1.0")


# -----------------------------------------------------------------------------
# LESSON 1: Domain entity
# Why this matters: domain should not depend on HTTP or framework.
@dataclass
class Task:
    id: int
    title: str
    done: bool = False


# -----------------------------------------------------------------------------
# LESSON 2: Repository interface
# Why this matters: service depends on abstraction.
class TaskRepo(Protocol):
    def list_tasks(self) -> list[Task]:
        ...

    def add_task(self, title: str) -> Task:
        ...


# -----------------------------------------------------------------------------
# LESSON 3: In-memory repository adapter
# Why this matters: simple adapter for local development and tests.
class InMemoryTaskRepo:
    def __init__(self) -> None:
        self._items: list[Task] = []

    def list_tasks(self) -> list[Task]:
        return list(self._items)

    def add_task(self, title: str) -> Task:
        next_id = 1 if not self._items else max(t.id for t in self._items) + 1
        task = Task(id=next_id, title=title.strip())
        self._items.append(task)
        return task


# -----------------------------------------------------------------------------
# LESSON 4: Service layer
# Why this matters: workflow rules live outside route handlers.
class TaskService:
    def __init__(self, repo: TaskRepo) -> None:
        self.repo = repo

    def create_task(self, title: str) -> Task:
        cleaned = title.strip()
        if not cleaned:
            raise ValueError("Title cannot be empty")
        return self.repo.add_task(cleaned)

    def list_tasks(self) -> list[Task]:
        return self.repo.list_tasks()


# -----------------------------------------------------------------------------
# LESSON 5: Request and response models
# Why this matters: typed API contracts reduce ambiguity.
class CreateTaskRequest(BaseModel):
    title: str = Field(min_length=1, max_length=120)


class TaskResponse(BaseModel):
    id: int
    title: str
    done: bool


# -----------------------------------------------------------------------------
# LESSON 6: Dependency provider
# Why this matters: route receives service dependency explicitly.
repo_singleton = InMemoryTaskRepo()


def get_service() -> TaskService:
    return TaskService(repo_singleton)


# -----------------------------------------------------------------------------
# LESSON 7: Health route
# Why this matters: simple operational check endpoint.
@app.get("/health")
def health() -> dict[str, str]:
    return {"status": "ok"}


# -----------------------------------------------------------------------------
# LESSON 8: List route
# Why this matters: read endpoint maps service output to response model.
@app.get("/tasks", response_model=list[TaskResponse])
def list_tasks(service: TaskService = Depends(get_service)) -> list[TaskResponse]:
    tasks = service.list_tasks()
    return [TaskResponse(id=t.id, title=t.title, done=t.done) for t in tasks]


# -----------------------------------------------------------------------------
# LESSON 9: Create route with error mapping
# Why this matters: convert domain errors into HTTP contract.
@app.post("/tasks", response_model=TaskResponse, status_code=201)
def create_task(payload: CreateTaskRequest, service: TaskService = Depends(get_service)) -> TaskResponse:
    try:
        task = service.create_task(payload.title)
    except ValueError as exc:
        raise HTTPException(status_code=400, detail=str(exc)) from exc
    return TaskResponse(id=task.id, title=task.title, done=task.done)


# -----------------------------------------------------------------------------
# LESSON 10: Root route with guidance
# Why this matters: lightweight onboarding for API consumers.
@app.get("/")
def root() -> dict[str, str]:
    return {"message": "Use /docs to explore the API"}


# End of FastAPI Basics 1-10
