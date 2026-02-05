# What is a dependency? (first principles)

Goal: understand what a dependency means and why it exists.

Why do we care?
- We often reuse code instead of rewriting it
- Dependencies save time and reduce bugs
- Understanding them helps you avoid confusion later

First principles
- A dependency is code your code needs
- You did not write it, but you use it
- Versions exist because code changes over time

What to do (slow and simple)
- Open `package.json`
- Notice there is no `dependencies` section yet
- That is okay; we will add one when needed

Checkpoint
- If you can say “a dependency is code I rely on,” you are ready to move on

Reflection
- Why might you want to use a library instead of writing everything yourself?
- What could go wrong if a dependency changes unexpectedly?
