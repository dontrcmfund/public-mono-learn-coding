# Clean architecture (first principles)

Goal: understand why layered architecture helps maintain large systems.

Why do we care?
- It separates business rules from frameworks and databases
- It makes core logic easier to test
- It reduces lock-in to specific tools

First principles
- Inner layers contain core business logic
- Outer layers contain implementation details (DB, HTTP, CLI)
- Dependencies point inward toward stable business rules

Typical layers
- Domain: entities and core rules
- Use cases/services: application workflows
- Interfaces/adapters: repos, gateways, presenters
- Frameworks/drivers: DB clients, web servers, CLIs

Rule of thumb
- Keep business logic independent of delivery mechanism

If all you remember is one thing
- Clean architecture protects core logic from tool and framework churn
