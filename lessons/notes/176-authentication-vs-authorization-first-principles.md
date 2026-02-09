# Authentication vs authorization (first principles)

Goal: understand the security difference between "who are you?" and "what can you do?"

Why do we care?
- Many security bugs come from mixing these two concepts
- Clear separation prevents privilege mistakes

Core definitions
- Authentication (authn): verify identity
- Authorization (authz): verify permissions for an action

Past, present, future value
- Past: systems often trusted users too broadly after login
- Present: APIs need explicit permission checks per endpoint
- Future: role and policy systems grow from this foundation

Rule of thumb
- Authenticate first
- Authorize per action/resource
- Deny by default

If all you remember is one thing
- Identity is not permission; both checks are required
