# innerHTML safety (gotchas)

Goal: understand when to avoid innerHTML.

Why do we care?
- innerHTML can insert raw HTML
- If the HTML comes from user input, it can be unsafe

First principles
- textContent inserts text only (safe)
- innerHTML inserts HTML (powerful but risky)

Rule of thumb
- Prefer createElement + textContent for user content
- Use innerHTML only with trusted strings

If all you remember is one thing
- innerHTML is powerful, but use it carefully
