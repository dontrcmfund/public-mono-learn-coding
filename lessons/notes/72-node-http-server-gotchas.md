# Node HTTP server gotchas

Goal: avoid the most common beginner confusion.

Why do we care?
- Servers “keep running,” which can feel like a freeze
- Ports can conflict with other apps

First principles
- A server listens on a port and waits for requests
- Your program does not exit while the server is listening

Common issues
- “Address already in use” means the port is busy
- You may need to stop the old server before starting a new one

Rule of thumb
- Use ports like 3000, 4000, 5000 for local testing

If all you remember is one thing
- A running server keeps your script alive on purpose
