# Workspaces basics

Goal: understand what a workspace is and why it matters in real projects.

Why do we care?
- Workspaces connect related packages without publishing to npm
- Shared code is easier and safer when managed in one repo
- Tooling can treat many packages as one system

Why this matters to you (past, present, future)
- Past: copying code between folders creates drift and confusion
- Present: local linking helps you learn shared code quickly
- Future: many production monorepos depend on this exact pattern

First principles
- A workspace is a package inside a larger repo
- The workspace tool links packages locally
- One root can manage shared dependencies and scripts

Short history (why this exists)
- Teams used to keep many repos and duplicate shared code
- Monorepos became popular to reduce version mismatch and duplication
- Workspace tools (`pnpm`, `npm`, `yarn`) standardized local linking

Etymology
- `workspace` means a defined area for related work
- In JS tooling, it means a package that participates in root-level dependency management

What to do
- Open `pnpm-workspace.yaml`
- Confirm it includes `apps/*` and `packages/*`
- Open `packages/shared/package.json` and note the package name

Mini exercise
- Create `packages/shared/src/math.ts`
- Add `add(a: number, b: number)`
- Export it from `packages/shared/src/index.ts`

If all you remember is one thing
- Workspaces let many packages behave like one coordinated project

Checkpoint
- If you can explain why local linking beats copy-paste, you are ready to move on

Reflection
- What should live in shared code?
- What should stay app-specific?
