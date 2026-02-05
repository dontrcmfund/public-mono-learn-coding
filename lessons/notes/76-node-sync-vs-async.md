# Node sync vs async file ops (gotchas)

Goal: understand the difference between sync and async in Node.

Why do we care?
- Sync operations block the program
- Async operations let other work continue

First principles
- Sync (writeFileSync) waits and blocks
- Async (writeFile) runs later and uses callbacks/promises

Rule of thumb
- Use sync in tiny scripts
- Use async for anything long or frequent

If all you remember is one thing
- Sync blocks; async does not
