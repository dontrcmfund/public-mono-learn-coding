# Docker Compose basics (first principles)

Goal: understand what Docker Compose is and why it exists.

Why do we care?
- Real apps often need multiple containers
- Compose starts them together with one command
- It keeps your setup organized

Why this matters to you (real life)
- Many tutorials use Compose (app + database)
- One command is easier than five commands

First principles
- Compose uses a file called `docker-compose.yml`
- It lists services (containers) and how they connect
- One command can start everything

What to do (slow and simple)
- Read this idea: “one file, many containers”
- No commands yet

If all you remember is one thing
- Compose starts multiple containers using one file

Checkpoint
- If you can say “Compose runs multiple containers,” you are ready to move on

Reflection
- Why is one command helpful when there are many parts?
- When might you need more than one container?
