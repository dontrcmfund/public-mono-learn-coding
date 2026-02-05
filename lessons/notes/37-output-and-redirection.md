# Output and redirection (macOS)

Goal: understand where command output goes and how to redirect it.

Why do we care?
- You can save output to a file
- You can chain commands together
- It makes the terminal more powerful

First principles
- Commands write output to the screen
- You can redirect output to a file
- You can pipe output into another command

Core symbols
- `>` write output to a file (overwrite)
- `>>` append output to a file
- `|` pipe output to another command

What to do (slow and simple)
- Run `ls > files.txt`
- Open `files.txt` with `cat files.txt`
- Run `ls | grep lessons`

Checkpoint
- If you can say “I can save or pass output,” you are ready to move on

Reflection
- Why might you want to save output?
- What could you do with piped output later?
