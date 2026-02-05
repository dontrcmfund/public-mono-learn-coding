# Reading files in the terminal (macOS)

Goal: safely view file contents without editing them.

Why do we care?
- You often need to check a file quickly
- Reading is safer than editing
- It helps you inspect logs and configs

First principles
- The terminal can print a file’s text
- You can scroll without changing anything

Core commands
- `cat file` prints the whole file
- `less file` shows a scrollable view
- `head file` shows the top lines
- `tail file` shows the bottom lines

What to do (slow and simple)
- Open `README.md` with `less README.md`
- Press `q` to quit

Checkpoint
- If you can say “I can view files without editing,” you are ready to move on

Reflection
- When would you choose `less` instead of `cat`?
- Why is “read‑only” helpful?
