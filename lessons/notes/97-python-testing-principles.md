# Python testing principles

Goal: understand why testing is part of learning, not an advanced extra.

Why do we care?
- Tests protect against regressions
- Tests clarify expected behavior
- Tests speed up refactoring and confidence

First principles
- A test checks one behavior with a clear expected result
- Good tests are small, deterministic, and readable
- Failing tests are useful feedback, not failure of the learner

Minimal testing workflow
- Write function
- Add test for normal case
- Add test for edge case
- Fix code until both pass

If all you remember is one thing
- Tests are executable explanations of what your code should do
