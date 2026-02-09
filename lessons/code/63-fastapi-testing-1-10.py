"""
FASTAPI TESTING (Lessons 1-10)

Suggested use:
1) Read this file first
2) Install dependencies later when ready: fastapi, pytest
3) Run tests later: pytest lessons/code/63-fastapi-testing-1-10.py

Extra context:
- lessons/notes/142-fastapi-gotchas.md
- lessons/notes/133-design-testing-principles.md
"""

from dataclasses import dataclass

from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.testclient import TestClient
from pydantic import BaseModel, Field

app = FastAPI(title="FastAPI Testing Lessons")


# -----------------------------------------------------------------------------
# LESSON 1: Domain + service kept framework-agnostic
# Why this matters: service tests should run without HTTP stack.
@dataclass
class Item:
    id: int
    name: str


class ItemService:
    def __init__(self) -> None:
        self._items: list[Item] = []

    def list_items(self) -> list[Item]:
        return list(self._items)

    def add_item(self, name: str) -> Item:
        clean = name.strip()
        if not clean:
            raise ValueError("Name cannot be empty")
        if any(i.name == clean for i in self._items):
            raise ValueError("Duplicate name")
        next_id = 1 if not self._items else max(i.id for i in self._items) + 1
        item = Item(id=next_id, name=clean)
        self._items.append(item)
        return item


# -----------------------------------------------------------------------------
# LESSON 2: API models
# Why this matters: request/response contracts should be testable.
class ItemCreateRequest(BaseModel):
    name: str = Field(min_length=1, max_length=120)


class ItemResponse(BaseModel):
    id: int
    name: str


# -----------------------------------------------------------------------------
# LESSON 3: Dependency provider
# Why this matters: tests can override dependencies if needed.
service_singleton = ItemService()


def get_service() -> ItemService:
    return service_singleton


# -----------------------------------------------------------------------------
# LESSON 4: Routes
# Why this matters: thin route handlers map service errors to HTTP.
@app.get("/items", response_model=list[ItemResponse])
def list_items(service: ItemService = Depends(get_service)) -> list[ItemResponse]:
    return [ItemResponse(id=i.id, name=i.name) for i in service.list_items()]


@app.post("/items", response_model=ItemResponse, status_code=status.HTTP_201_CREATED)
def create_item(payload: ItemCreateRequest, service: ItemService = Depends(get_service)) -> ItemResponse:
    try:
        item = service.add_item(payload.name)
    except ValueError as exc:
        code = status.HTTP_409_CONFLICT if "Duplicate" in str(exc) else status.HTTP_400_BAD_REQUEST
        raise HTTPException(status_code=code, detail=str(exc)) from exc
    return ItemResponse(id=item.id, name=item.name)


# -----------------------------------------------------------------------------
# LESSON 5: Test client
# Why this matters: integration-like HTTP checks are easy to run.
client = TestClient(app)


# -----------------------------------------------------------------------------
# LESSON 6: Service unit test
# Why this matters: verify business rule without API layer.
def test_service_add_success() -> None:
    service = ItemService()
    item = service.add_item("Notebook")
    assert item.id == 1
    assert item.name == "Notebook"


# -----------------------------------------------------------------------------
# LESSON 7: Service duplicate test
# Why this matters: protects duplicate-prevention rule.
def test_service_add_duplicate() -> None:
    service = ItemService()
    service.add_item("Notebook")
    try:
        service.add_item("Notebook")
        assert False, "Expected ValueError"
    except ValueError as exc:
        assert "Duplicate" in str(exc)


# -----------------------------------------------------------------------------
# LESSON 8: API create test
# Why this matters: validates HTTP status + payload contract.
def test_api_create_item() -> None:
    global service_singleton
    service_singleton = ItemService()
    response = client.post("/items", json={"name": "Pen"})
    assert response.status_code == 201
    data = response.json()
    assert data["id"] == 1
    assert data["name"] == "Pen"


# -----------------------------------------------------------------------------
# LESSON 9: API duplicate conflict test
# Why this matters: error mapping must remain stable.
def test_api_duplicate_item_conflict() -> None:
    global service_singleton
    service_singleton = ItemService()
    client.post("/items", json={"name": "Pen"})
    response = client.post("/items", json={"name": "Pen"})
    assert response.status_code == 409


# -----------------------------------------------------------------------------
# LESSON 10: API list test
# Why this matters: read endpoint should reflect created state.
def test_api_list_items() -> None:
    global service_singleton
    service_singleton = ItemService()
    client.post("/items", json={"name": "Pen"})
    client.post("/items", json={"name": "Pencil"})
    response = client.get("/items")
    assert response.status_code == 200
    data = response.json()
    assert len(data) == 2


# End of FastAPI Testing 1-10
