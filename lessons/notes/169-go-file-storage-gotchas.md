# Go file storage gotchas

Goal: avoid beginner mistakes when storing app data in files.

Why do we care?
- File bugs cause silent data loss or corrupt data
- Persistence must be trustworthy before adding complexity

Common gotchas
- Forgetting to handle read/write errors
- Assuming file exists on first run
- Overwriting data accidentally instead of append/merge
- Invalid JSON shape when schema evolves
- Ignoring concurrent writes from multiple goroutines

Rule of thumb
- Treat every file operation as fallible
- Use clear load/save functions and deterministic formats
- Add tests for missing file, malformed file, and round-trip behavior

If all you remember is one thing
- File storage is simple, but only reliable with explicit error handling and tests
