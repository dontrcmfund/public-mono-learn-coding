# What is a script? (first principles)

Goal: understand what a “script” is and why it exists.

Why do we care?
- Scripts let us run repeatable tasks with one command
- This saves time and avoids mistakes

First principles
- A script is just a named command
- The name is stored in `package.json` under `scripts`
- Running a script is the same as typing the command

What to do (slow and simple)
- Open `package.json`
- Find the `scripts` section
- Read the names: `lint`, `format`, `typecheck`, `test`
- Notice the commands are currently placeholders

Checkpoint
- If you can say “a script is a named command,” you are ready to move on

Reflection
- Why might it be helpful to have a short name for a long command?
- What script do you think you will use most later?
