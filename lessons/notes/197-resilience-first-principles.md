# Resilience (first principles)

Goal: understand how systems keep serving users when dependencies are slow or failing.

Why do we care?
- Real systems always face partial failures
- One failing dependency can cascade into full-service outages

Core ideas
- Timeouts bound waiting time
- Retries recover transient failures
- Isolation patterns prevent one failure from consuming all resources
- Fallbacks keep degraded but useful behavior

Past, present, future value
- Past: cascading failures caused large outages from small incidents
- Present: resilience patterns reduce blast radius and downtime
- Future: resilient systems scale safely under uncertainty

If all you remember is one thing
- Resilience is designing for failure as a normal condition
