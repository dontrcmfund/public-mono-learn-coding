# Go API principles

Goal: build HTTP APIs with clear contracts and maintainable layers.

Why do we care?
- APIs are a common backend interface
- Stable contracts reduce integration errors

First principles
- Parse and validate input at the handler boundary
- Keep service logic independent of HTTP details
- Map errors to explicit HTTP status codes

Go-specific guidance
- Use `net/http` basics before adding frameworks
- Keep handlers thin and delegate to services
- Return JSON responses with consistent shapes

If all you remember is one thing
- Good Go APIs keep HTTP handling thin and business logic separate
