# DOM event bubbling and delegation (gotchas)

Goal: understand why events seem to fire on parents.

Why do we care?
- Clicks and key events "bubble up" the DOM tree
- It explains why one listener can handle many items

First principles
- Events start on the element you clicked
- Then they bubble up to parents (unless stopped)

Common confusion
- You click a child, but the parent handler runs too
- That is normal and called "bubbling"

Rule of thumb
- Use event delegation on a parent for many items
- Use stopPropagation only when you must

If all you remember is one thing
- Events bubble from child to parent by default
