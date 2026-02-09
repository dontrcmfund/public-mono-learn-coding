# OpenAPI (first principles)

Goal: understand why teams use machine-readable API contracts.

Why do we care?
- Humans need readable docs, tools need structured contracts
- OpenAPI enables validation, code generation, mocks, and test automation

Core idea
- OpenAPI describes endpoints, parameters, request/response schemas, and errors
- One contract can power docs UI and automated checks

History context
- REST APIs grew faster than documentation consistency
- OpenAPI standardized contract format for tooling ecosystem support

Practical guidance
- Define stable schema names
- Reuse common error schemas
- Version contract changes intentionally

If all you remember is one thing
- OpenAPI turns API expectations into an enforceable spec
