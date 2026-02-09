# Deployment safety gotchas

Goal: avoid common release mistakes that cause avoidable incidents.

Why do we care?
- Deployment is a normal operational event, but mistakes can impact many users quickly

Common gotchas
- Deploying without passing tests/lint/security gates
- Releasing config changes without validation
- No rollback plan
- No health checks after deploy
- Ignoring migration order (app vs schema timing)

Safety habits
- Gate releases on automated checks
- Use staged rollouts where possible
- Verify health and core user flows after deploy

If all you remember is one thing
- Safe deployment is planned, gated, and observable
