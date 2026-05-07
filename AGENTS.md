# sane-next

## Product frame

- `sane-next` is a **Pi-first overlay/distribution**, not a fresh runtime and not a deep Pi fork.
- Pi is the primary runtime. Codex-native skill export is a secondary target.

## Startup rules

- Read `TRACK.toml`, `docs/adr/0002-use-track-toml-plus-adrs-for-tracking.md`, `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`, and `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`.
- Use current repo files and fresh external research. Do not trust old Sane claims or stale chat memory.
- Keep always-on instructions small. Prefer trigger-based skills and on-demand references.
- For stack, MCP, language, Pi, or Codex behavior choices, research current evidence before implementing.
- Update `TRACK.toml` for current state and ADRs for durable decisions.

## Hard boundaries

- Do not rebuild Pi's runtime in this repo.
- Do not add a TUI-heavy standalone surface.
- Do not create parallel `TODO`, `plan`, `memory`, or research-dump files in the repo.
- Do not add duplicated policy or generated repo summaries to always-on surfaces.

## Done

- Current files reflect current truth.
- Superseded files are removed.
- Instruction surfaces stay small and non-duplicated.
