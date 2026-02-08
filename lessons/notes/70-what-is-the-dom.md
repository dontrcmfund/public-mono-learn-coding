# What is the DOM? (first principles)

Goal: understand what the DOM is, why browsers use it, and how it changes your web programming model.

Why do we care?
- The DOM is how JavaScript reads and updates a webpage
- It is the bridge between code and visible UI
- Most interactive web behavior depends on DOM operations

Why this matters to you (past, present, future)
- Past: if HTML felt static, DOM explains how pages become interactive
- Present: this is how your buttons, inputs, and messages actually work
- Future: frameworks still rely on DOM concepts under the hood

First principles
- HTML is source text for structure
- Browser parsing turns that structure into a live object tree
- That live tree is the Document Object Model (DOM)

Short history (why this exists)
- In the early web, pages were mostly static documents
- Browsers needed a common programmable model for page elements
- The DOM standardized element access and events across browsers

Etymology
- `document` means the web page
- `object` means programmable entities in memory
- `model` means an abstract representation you can operate on

What to do (slow and simple)
- Open one DOM lesson HTML file and inspect elements in DevTools
- Change text with JavaScript and observe instant UI changes

If all you remember is one thing
- The DOM is the browser's live, programmable model of your HTML page

Checkpoint
- If you can explain why changing `textContent` updates the screen, you are ready to move on

Reflection
- Which UI behavior felt "magical" before learning DOM?
- How does the tree model make that behavior less mysterious?
