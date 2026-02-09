# Outbox pattern (first principles)

Goal: understand how to avoid lost events when writing data and publishing events.

Why do we care?
- If DB write succeeds but publish fails, system state and events diverge
- Reliability requires handling this boundary explicitly

Core idea
- Write domain change and outbox event record in same transaction
- Separate publisher process reads outbox and publishes events
- Mark outbox records as delivered after successful publish

Benefits
- No lost events from transient broker failures
- Retry publication safely
- Clear audit trail of pending/sent events

Rule of thumb
- For important integrations, prefer outbox over direct "fire-and-forget" publish

If all you remember is one thing
- Outbox turns fragile publish timing into a reliable, retryable workflow
