# Dockerfile basics (first principles)

Goal: understand what a Dockerfile is.

Why do we care?
- Dockerfiles let you build your own images
- They make setup repeatable
- They help you share your setup with others

Why this matters to you (real life)
- You can rebuild your environment any time
- A teammate (or future you) can get the same setup easily

First principles
- A Dockerfile is a recipe written as steps
- Each line adds to the image
- Docker builds images from Dockerfiles

What to do (slow and simple)
- Read this tiny example:
  - `FROM node:20`
  - `WORKDIR /app`
  - `COPY . .`
  - `CMD ["node", "index.js"]`
- You do not need to run it yet

If all you remember is one thing
- A Dockerfile is a repeatable recipe for your environment

Checkpoint
- If you can say “a Dockerfile is a recipe,” you are ready to move on

Reflection
- Why would a repeatable recipe be useful?
- What could go wrong without one?
