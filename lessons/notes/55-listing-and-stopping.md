# Listing and stopping containers

Goal: see what is running and how to stop it.

Why do we care?
- You should know what is running on your computer
- You need a safe way to stop containers
- Stopping containers frees memory and battery

Why this matters to you (real life)
- If your laptop gets slow, a container might be running
- You can stop things without uninstalling them

First principles
- `docker ps` shows running containers
- `docker ps -a` shows all containers
- `docker stop` stops a running container

What to do (slow and simple)
- Run `docker ps`
- Run `docker ps -a`
- Notice the difference

If all you remember is one thing
- `docker ps` shows what is running right now

Checkpoint
- If you can explain “running vs all,” you are ready to move on

Reflection
- Why is it useful to list all containers?
- When would you want to stop one?
