# Event contracts and versioning

Goal: understand why event payloads need stable contracts over time.

Why do we care?
- Consumers depend on event fields and semantics
- Breaking changes can silently fail downstream systems

Core ideas
- Event contract = event name + required fields + field meanings
- Version events when making incompatible changes
- Additive changes are safer than destructive changes

Practical guidance
- Include `event_type`, `version`, `occurred_at`, and `id`
- Keep payloads focused on business facts, not internal implementation details
- Document contracts in code and notes

Rule of thumb
- Treat event schema as public API

If all you remember is one thing
- Stable event contracts protect consumers from producer changes
