# CI/CD (first principles)

Goal: understand why software teams automate build, test, and delivery workflows.

Why do we care?
- Manual release steps are slow and error-prone
- Automation gives consistent quality gates before code reaches users

Core ideas
- CI (continuous integration): automatically build/test on each change
- CD (continuous delivery/deployment): automatically prepare or ship validated changes

Past, present, future value
- Past: "works on my machine" broke releases
- Present: CI catches regressions early and cheaply
- Future: reliable delivery enables faster iteration with less risk

Rule of thumb
- Automate repeatable checks
- Block releases when critical checks fail

If all you remember is one thing
- CI/CD turns quality and delivery into repeatable systems, not heroics
