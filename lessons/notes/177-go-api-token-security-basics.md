# Go API token security basics

Goal: learn why APIs use tokens and how middleware enforces security consistently.

Why do we care?
- APIs are called by scripts, apps, and services, not just browsers
- Token-based auth is simple to automate and validate

History context
- Session cookies dominated web apps with browser state
- APIs moved toward stateless tokens for service-to-service communication
- Middleware became standard to centralize repeated security checks

Core ideas
- Read token from `Authorization: Bearer <token>`
- Validate token before protected handlers run
- Attach identity/role context to request for downstream checks

Security basics
- Never hardcode production secrets in code
- Prefer short-lived credentials and rotation
- Log denied requests safely without leaking secrets

If all you remember is one thing
- Centralized token validation middleware reduces duplicated security mistakes
