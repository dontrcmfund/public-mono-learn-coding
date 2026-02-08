# What is Docker? (first principles)

Goal: understand what Docker is, why it exists, and why teams use it.

Why do we care?
- It helps software run consistently across different computers
- It reduces setup mismatch and dependency drift
- It makes local testing closer to production behavior

Why this matters to you (past, present, future)
- Past: if tutorials broke due to machine differences, containers reduce that pain
- Present: you can learn in isolated environments with lower risk
- Future: container workflows are standard in modern deployment pipelines

First principles
- A container is an isolated runtime environment
- An image is the blueprint used to create containers
- Docker is the toolchain that builds and runs those containers

Short history (why this exists)
- Before containers, teams used virtual machines for isolation
- VMs were heavier and slower for many developer workflows
- Docker (2013) made container workflows accessible and repeatable for everyday development

Etymology
- `Docker` comes from the metaphor of shipping containers and docks
- The idea is standardized packaging and movement across environments

What to do (slow and simple)
- Read this lesson fully before running commands
- Focus on the concept of "same environment, different machine"

If all you remember is one thing
- Docker gives you a consistent software environment you can run anywhere

Checkpoint
- If you can explain image vs container in one sentence, you are ready to move on

Reflection
- Where in your experience would consistency have saved time?
- What feels risky about setup today that containers might reduce?
