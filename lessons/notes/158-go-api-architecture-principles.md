# Go API architecture principles

Goal: build HTTP services with clear boundaries and stable contracts.

Why do we care?
- APIs evolve; architecture must absorb change without breaking core logic
- Thin handlers and service boundaries reduce regression risk

Layering model
- Domain: entities and core rules
- Service: workflows and business policies
- Repository: data adapters
- Transport: HTTP handlers and JSON mapping

Go guidance
- Keep handlers small and explicit
- Map errors to stable status codes
- Keep service package independent from `net/http`

If all you remember is one thing
- Handlers translate requests; services decide behavior
