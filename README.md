# sane-next

`sane-next` is a **Pi-first overlay/distribution** for high-leverage coding-agent workflows, with shared packs that also export to Codex-native skill paths.

This repo is intentionally small, but it is **not** a single-summary-doc repo. The implementation agent should be able to work from the repo alone without chat history, so durable context is split by purpose.

## Read in this order

1. `TRACK.toml` — the current active slice
2. `docs/reference/OLD-SANE-POSTMORTEM.md` — what old Sane taught us and what not to copy
3. core ADRs:
   - `docs/adr/0002-use-track-toml-plus-adrs-for-tracking.md`
   - `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`
   - `docs/adr/0007-use-go-for-companion-cli.md`
   - `docs/adr/0008-use-small-tooling-policy-and-optional-mcps.md`
   - `docs/adr/0009-use-annotated-semver-and-linux-only-ci-while-private.md`
4. standards:
   - `docs/standards/TRACK-STRUCTURE-STANDARD.md`
   - `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`
   - `docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md`
5. `AGENTS.md` — small startup rules
6. `CONTRIBUTING.md` — commit and hook rules

## Repo structure

- `TRACK.toml` — bounded active execution window only
- `docs/reference/` — reference context and postmortems
- `docs/adr/` — durable decisions
- `docs/standards/` — durable standards and execution protocols
- `AGENTS.md` — tiny always-on instruction surface
- `.githooks/` — committed repo hooks

Do not add random planning, TODO, memory, or research markdown files outside this structure.
