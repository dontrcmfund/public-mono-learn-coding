# Go release versioning (first principles)

Goal: understand how versioning communicates change risk to consumers.

Why do we care?
- Consumers need to know whether upgrades are safe or breaking
- Clear versions prevent accidental incompatibility in production

Core idea
- Semantic versioning pattern: `MAJOR.MINOR.PATCH`
- `MAJOR`: breaking changes
- `MINOR`: backward-compatible features
- `PATCH`: backward-compatible fixes

Go-specific relevance
- Go modules and tags are version-aware
- Reliable versioning improves dependency upgrade confidence

Rule of thumb
- Tie version bumps to contract impact, not team mood

If all you remember is one thing
- Version numbers are promises about compatibility
