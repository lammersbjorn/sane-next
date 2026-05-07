# sane-next

## Product frame

- `sane-next` is a **Pi-first overlay/distribution**, not a fresh runtime and not a deep Pi fork.
- Pi is the primary runtime. Codex-native skill export is a secondary target.

## Startup rules

- Read `TRACK.toml`, `docs/reference/OLD-SANE-POSTMORTEM.md`, `docs/adr/0002-use-track-toml-plus-adrs-for-tracking.md`, `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`, `docs/adr/0007-use-go-for-companion-cli.md`, `docs/adr/0008-use-small-tooling-policy-and-optional-mcps.md`, `docs/adr/0009-use-annotated-semver-and-linux-only-ci-while-private.md`, `docs/standards/TRACK-STRUCTURE-STANDARD.md`, `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`, and `docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md` before broad work.
- For the future `/goal` build run, also read `docs/standards/GOAL-RUN-STANDARD.md` and load `.agents/skills/sane-next-implementation/SKILL.md`.
- Use current repo files and fresh external research. Do not trust old Sane claims or stale chat memory.
- Keep always-on instructions small.
- Research current evidence before changing stack, Pi, MCP, or model assumptions.
- Update `TRACK.toml` only for the active slice.

## Hard boundaries

- Do not rebuild Pi's runtime in this repo.
- Do not add a TUI-heavy standalone surface.
- Do not create parallel `TODO`, `plan`, `memory`, ADR, or research-dump files in the repo.
- Do not add markdown docs outside the fixed repo structure without intentionally changing the structure first.

## Done

- Current files reflect current truth.
- Repo docs stay small and structured.
- The implementation path stays clear without chat context.
