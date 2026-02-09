# Go auth gotchas

Goal: avoid common mistakes when adding auth to HTTP handlers.

Why do we care?
- Small auth mistakes can expose protected data
- Security regressions are expensive and high-impact

Common gotchas
- Checking token format but not token validity
- Authenticating but forgetting authorization checks
- Returning unclear status codes (`401` vs `403`)
- Trusting client-provided role fields directly
- Logging raw tokens in plaintext

Status code guide
- `401 Unauthorized`: identity missing/invalid
- `403 Forbidden`: identity valid, permission denied

Rule of thumb
- Use middleware for authn
- Use endpoint-level checks for authz
- Keep security logic explicit and test-covered

If all you remember is one thing
- Secure APIs require both correct checks and correct failure behavior
