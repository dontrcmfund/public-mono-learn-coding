# Node fetch gotchas

Goal: avoid common confusion when using fetch.

Why do we care?
- fetch does not throw on HTTP errors by default
- Network issues are common

First principles
- fetch resolves for 404/500 responses
- You must check res.ok or res.status
- Only network failures throw an error

Rule of thumb
- Always check res.ok before reading data

If all you remember is one thing
- fetch only throws on network failure, not on HTTP errors
