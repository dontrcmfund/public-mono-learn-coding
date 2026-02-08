# TypeScript common gotchas

Goal: avoid early confusion when moving from JavaScript to TypeScript.

Why do we care?
- New learners often think TypeScript changes runtime behavior
- Most early frustration is misunderstanding what TypeScript does and does not do

First principles
- TypeScript checks types before running
- After compilation, runtime is plain JavaScript
- Type errors are development feedback, not runtime exceptions

Common gotchas
- Type annotations do not make values safe at runtime by themselves
- `any` disables type safety and should be rare
- Union types need narrowing before specific operations
- Optional properties can be undefined and need checks

Rule of thumb
- Prefer specific types over `any`
- When TypeScript complains, ask "what value shape is uncertain?"

If all you remember is one thing
- TypeScript is a planning and checking tool; runtime is still JavaScript
