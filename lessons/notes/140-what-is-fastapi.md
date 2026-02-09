# What is FastAPI? (first principles)

Goal: understand what FastAPI is and why it fits modern Python API development.

Why do we care?
- It is fast to build and easy to document APIs
- It uses Python type hints for validation and tooling
- It aligns with clean architecture patterns (domain/service/repo)

Why this matters to you (past, present, future)
- Past: if backend setup felt heavy, FastAPI lowers startup friction
- Present: you can expose your Python logic as web endpoints quickly
- Future: typed APIs and OpenAPI docs are widely used in production systems

First principles
- An API is a contract for how clients interact with a service
- FastAPI maps HTTP routes to Python functions
- Request/response models define data contracts explicitly

Short history
- FastAPI (2018) built on Starlette + Pydantic
- It gained adoption through strong typing, async support, and built-in docs
- It emerged as a modern alternative to heavier frameworks for API-first work

If all you remember is one thing
- FastAPI turns typed Python functions into reliable, documented APIs
