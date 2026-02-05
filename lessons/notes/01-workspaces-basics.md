# Workspaces basics

Goal: understand what a workspace is and why it matters.

Key ideas
- A workspace is a package inside the monorepo
- The workspace tool links packages locally (no publishing needed)
- Shared dependencies can be installed once at the repo root

What to do
- Open `pnpm-workspace.yaml`
- Confirm it includes `apps/*` and `packages/*`
- Open `packages/shared/package.json` and note the name

Mini exercise
- Create a new file at `packages/shared/src/math.ts`
- Add a function `add(a: number, b: number)` that returns the sum
- Export it from `packages/shared/src/index.ts`

Reflection
- What would you put in `shared`?
- What should remain inside an app instead?
