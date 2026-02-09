# Go fakes and mocks gotchas

Goal: avoid common testing mistakes when isolating dependencies.

Why do we care?
- Tests that touch real infrastructure become slow and flaky
- Fakes speed up feedback and improve confidence

Terms
- Fake: working lightweight implementation for tests (often in-memory)
- Mock: test double that verifies interactions/expectations

Common gotchas
- Over-mocking every dependency
- Testing private internals instead of behavior
- Shared mutable fake state leaking between tests
- Forgetting deterministic time/input control

Rule of thumb
- Prefer simple fakes first
- Mock only when interaction verification is the actual requirement
- Keep each test independent and explicit

If all you remember is one thing
- Good tests verify behavior with simple, deterministic doubles
