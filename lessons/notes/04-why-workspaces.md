# Why workspaces? (first principles)

Goal: understand why we use a workspace file.

Why do we care?
- Without workspaces, each package is isolated
- With workspaces, tools treat the repo as one big project
- This makes shared code and installs easier

First principles
- A workspace is just a list of package folders
- Tools read the list and “link” packages together
- The list lives in `pnpm-workspace.yaml`

What to do (slow and simple)
- Open `pnpm-workspace.yaml`
- Read the two lines under `packages:`
- Say out loud what each pattern means:
  - `apps/*` means “every folder inside apps”
  - `packages/*` means “every folder inside packages”

Checkpoint
- If you can explain what the two patterns do, you are ready to move on

Reflection
- Why is it helpful to list packages in one place?
- How would this help if we add more packages later?
