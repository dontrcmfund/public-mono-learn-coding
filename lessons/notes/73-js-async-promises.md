# JS async, promises, and await (gotchas)

Goal: understand why async code feels different.

Why do we care?
- Many tasks take time (files, network)
- JavaScript keeps running while it waits

First principles
- A promise represents a future result
- async/await is just a nicer way to use promises
- While waiting, other code can still run

Common confusion
- Code after an async call can run first
- Errors from promises must be handled with catch or try/catch

Rule of thumb
- Use async/await for clarity
- Always handle errors

If all you remember is one thing
- Async code means "wait for later" without freezing the program
