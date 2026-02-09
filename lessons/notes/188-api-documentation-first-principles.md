# API documentation (first principles)

Goal: understand why clear API documentation is part of correctness, not optional polish.

Why do we care?
- APIs are used by humans and machines that cannot read your internal code
- Undocumented behavior causes integration bugs and support churn

Past, present, future value
- Past: implicit API assumptions caused frequent breaking integrations
- Present: docs reduce onboarding time and production mistakes
- Future: well-documented APIs scale across teams and external partners

Core idea
- Document request shape, response shape, status codes, and error formats
- Keep docs versioned with code changes

Rule of thumb
- If behavior matters to consumers, it belongs in the contract docs

If all you remember is one thing
- API docs are the user manual for your service contract
