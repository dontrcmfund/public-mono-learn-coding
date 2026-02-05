# JS timers (gotchas)

Goal: understand why timers can feel confusing.

Why do we care?
- Timers run later, not immediately
- That timing difference can confuse beginners

First principles
- `setTimeout` schedules work for the future
- Code below it still runs right away

Common confusion
- People expect the timer to “pause” the program
- It does not — it just schedules a future action

Rule of thumb
- Use timers when you truly want a delay

If all you remember is one thing
- `setTimeout` does not stop the program; it schedules a future action
