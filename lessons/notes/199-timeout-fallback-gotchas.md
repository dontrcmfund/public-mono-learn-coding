# Timeout and fallback gotchas

Goal: avoid common mistakes when adding timeout and fallback behavior.

Why do we care?
- Wrong timeout/fallback settings can hide failures or harm user experience

Common gotchas
- Timeouts set too high (slow failure detection)
- Timeouts set too low (false failures under normal load)
- Fallback responses that violate contract format
- Fallback data without clear freshness or degradation signal
- Missing observability around timeout/fallback usage

Rule of thumb
- Keep timeout explicit per dependency
- Ensure fallback responses are contract-safe
- Track timeout and fallback rates

If all you remember is one thing
- Fallback is controlled degradation, not silent correctness loss
