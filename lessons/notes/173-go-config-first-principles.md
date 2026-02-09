# Go service configuration (first principles)

Goal: understand why apps should read behavior from config instead of hardcoded values.

Why do we care?
- Different environments need different settings (dev, test, prod)
- Hardcoded values make deployments risky and inflexible

Past, present, future value
- Past: changing code for each environment caused mistakes
- Present: config allows safe and fast environment changes
- Future: scaling teams need predictable, auditable runtime behavior

Core idea
- Code defines behavior rules
- Config provides runtime values (port, DB path, log level, feature flags)

Rule of thumb
- Keep defaults explicit
- Validate config at startup
- Fail fast with clear errors when config is invalid

If all you remember is one thing
- Configuration separates deployment choices from business logic
