# Dependency injection in Go (first principles)

Goal: understand why dependencies should be passed in, not hidden inside functions.

Why do we care?
- Hidden dependencies create surprises and hard-to-test code
- Injected dependencies make behavior explicit and predictable

History context
- Go emphasized explicitness and simple construction patterns
- Constructor functions and struct fields became a common DI style
- This fits Go's "clear over clever" design philosophy

Core idea
- Dependency injection means a component receives collaborators from outside
- The component owns workflow logic, not creation of collaborators

Common Go pattern
- Define service struct with interface fields
- Build it with a constructor like `NewService(dep DepInterface)`
- Wire concrete dependencies at the application edge (`main`)

Past, present, future value
- Past: easier debugging because dependencies are visible
- Present: faster tests with in-memory fakes
- Future: safer refactors when infrastructure changes

If all you remember is one thing
- Inject dependencies so business logic stays testable and replaceable
