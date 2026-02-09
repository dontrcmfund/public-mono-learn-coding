# FastAPI layering and dependencies

Goal: apply domain/service/repo layering inside FastAPI apps.

Why do we care?
- Framework code should stay thin
- Business rules should stay testable without HTTP

Layer model
- Domain: entities and rules
- Service: use-case workflows
- Repository: data persistence adapters
- FastAPI router: request parsing and response mapping

Dependency injection basics
- `Depends(...)` provides dependencies to route handlers
- Keep providers simple and explicit
- Inject services, not raw DB logic, into endpoints

Rule of thumb
- Route handlers orchestrate; services decide; repos persist

If all you remember is one thing
- FastAPI should call your architecture, not replace it
