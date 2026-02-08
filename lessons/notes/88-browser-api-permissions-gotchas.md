# Browser API permissions gotchas

Goal: avoid confusion when APIs fail due to browser security rules.

Why do we care?
- Some APIs require user gestures or secure contexts
- Behavior can differ by browser and environment

Common examples
- Clipboard writes may fail without a direct user action
- Some APIs are limited on `http` and work better on `https`
- Cross-origin requests can fail due to CORS policy

Rule of thumb
- Always handle failure paths for permission-sensitive APIs
- Check capability before use and provide fallback UI

If all you remember is one thing
- Browser API failures are often security rules, not broken code
