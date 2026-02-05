# Searching in the terminal (macOS)

Goal: find files or text when you don’t know where they are.

Why do we care?
- Projects grow fast
- Searching saves time
- You can locate problems quickly

First principles
- Search tools scan files and folders
- You can search by name or by text

Core commands
- `find . -name "pattern"` finds files by name
- `grep -R "text" .` finds text inside files

What to do (slow and simple)
- From repo root, run: `find lessons -name "*git*"`
- See which lesson files mention git

Checkpoint
- If you can say “I can search by name or text,” you are ready to move on

Reflection
- How would you find a file if you forgot its folder?
- Why is text search useful in code?
