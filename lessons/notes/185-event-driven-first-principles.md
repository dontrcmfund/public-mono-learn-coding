# Event-driven design (first principles)

Goal: understand why systems communicate with events instead of direct calls only.

Why do we care?
- Direct synchronous calls create tight coupling
- Events let producers and consumers evolve more independently

Past, present, future value
- Past: tightly coupled services were hard to change safely
- Present: events support async workflows and integration boundaries
- Future: event streams enable analytics, automation, and scalable architectures

Core idea
- Producer emits an event when business state changes
- Consumer reacts to event when ready
- Producer does not need to know all consumers in advance

Rule of thumb
- Use events for integration boundaries and asynchronous reactions
- Keep event meaning explicit and stable

If all you remember is one thing
- Events reduce coupling by broadcasting facts, not calling specific implementations
