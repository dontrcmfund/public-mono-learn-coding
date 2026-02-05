# Ports and access (first principles)

Goal: understand what ports are in Docker context.

Why do we care?
- Many containers run servers
- Ports let your browser talk to the container
- Without ports, you cannot see the app

Why this matters to you (real life)
- You will run a web app in Docker and want to open it
- Port mapping makes “localhost” work in your browser

First principles
- A port is a numbered door
- Containers have internal ports
- You can map them to your computer’s ports

What to do (slow and simple)
- Read this example: `docker run -p 8080:80 nginx`
- Know it means “map my port 8080 to container port 80”

If all you remember is one thing
- Ports are doors; mapping opens the door to your browser

Checkpoint
- If you can say “ports are doors,” you are ready to move on

Reflection
- Why do we need to map ports?
- What would happen if two apps used the same port?
