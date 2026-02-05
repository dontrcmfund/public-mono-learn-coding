# Volumes and data (first principles)

Goal: understand where container data lives.

Why do we care?
- Containers are temporary by default
- Volumes keep data safe between runs
- You do not want to lose work

Why this matters to you (real life)
- Databases can lose data without volumes
- You can edit files on your computer and see them in the container

First principles
- Container files can disappear when it stops
- A volume is a safe place on your computer
- You can connect a volume to a container

Simple analogy
- Volume = storage closet outside the lunchbox

What to do (slow and simple)
- Read this example: `docker run -v /my/data:/data app`
- Know it means “connect a folder to the container”

If all you remember is one thing
- Volumes keep data safe when containers stop

Checkpoint
- If you can say “volumes keep data safe,” you are ready to move on

Reflection
- Why might you want data to survive restarts?
- What data would you store in a volume?
