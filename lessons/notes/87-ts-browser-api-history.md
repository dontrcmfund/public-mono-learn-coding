# Browser API history and why behavior feels inconsistent

Goal: understand why browser APIs sometimes feel uneven.

Why do we care?
- Historical context explains odd API design choices
- Knowing this lowers frustration when APIs differ

Short history
- Early browsers shipped features quickly without shared standards
- Standards bodies later aligned behavior across browsers
- Some old API patterns remain for backward compatibility

Etymology and context
- `DOM` means Document Object Model, a programmable page model
- `navigator` originally exposed browser/client capabilities
- `localStorage` reflects persistent key-value storage local to one origin

Practical takeaway
- Prefer modern standardized APIs when possible
- Expect legacy quirks and guard your code with checks

If all you remember is one thing
- Many browser API quirks are historical compatibility tradeoffs
