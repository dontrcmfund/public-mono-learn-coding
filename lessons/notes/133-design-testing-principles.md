# Design testing principles

Goal: test architecture boundaries, not only individual functions.

Why do we care?
- Good architecture is proven by replaceability and stable behavior
- Tests should verify use-case logic independent of adapters

What to test
- Domain rules (pure logic)
- Use-case success and failure paths
- Adapter contracts (save/read behavior)
- Presentation formatting boundaries

Rule of thumb
- Most tests should target domain and use-case layers
- Fewer integration tests should check adapter wiring

If all you remember is one thing
- Architecture quality is visible when core logic stays testable without infrastructure
