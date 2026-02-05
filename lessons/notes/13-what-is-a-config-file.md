# What is a config file? (first principles)

Goal: understand why tools need config files.

Why do we care?
- Tools need to know how you want them to behave
- Config files make behavior clear and repeatable
- It keeps settings in one place

First principles
- A config file is just a text file of settings
- Tools read it before they run
- Changing the file changes the tool’s behavior

What to do (slow and simple)
- Look at `pnpm-workspace.yaml`
- That file is a config file for the workspace tool
- Notice it is short and readable

Checkpoint
- If you can say “a config file is tool settings,” you are ready to move on

Reflection
- Why is it better to store settings in a file instead of in your head?
- How does a config file help a team?
