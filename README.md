# sane-next

`sane-next` is a **Pi-first overlay/distribution** for high-leverage coding-agent workflows, with shared packs that also export to Codex-native skill paths.

This repo is intentionally small. The implementation agent should be able to work from the repo alone without chat history.

## Read in this order

1. `TRACK.toml` — the current active slice
2. `docs/FOUNDATION.md` — the fixed product and repo decisions
3. `docs/IMPLEMENTATION.md` — the exact build order and acceptance path
4. `AGENTS.md` — small startup rules
5. `CONTRIBUTING.md` — commit and hook rules

## Repo structure

- `TRACK.toml` — bounded active execution window only
- `docs/FOUNDATION.md` — current truth about product, boundaries, tooling, and repo rules
- `docs/IMPLEMENTATION.md` — step-by-step implementation sequence
- `AGENTS.md` — tiny always-on instruction surface
- `.githooks/` — committed repo hooks

Do not add more planning, TODO, memory, ADR, or research markdown files unless the repo structure is intentionally changed first.
