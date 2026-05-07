# sane-next repository instructions

Read `TRACK.toml`, `docs/adr/0002-use-track-toml-plus-adrs-for-tracking.md`, `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`, and `AGENTS.md` before broad work.

## Current architecture

- This repo is **pre-implementation**.
- `sane-next` is a **Pi-first overlay/distribution** with **Codex-native skill export** as a secondary target.
- Do not assume old Sane behavior, structure, prompts, or claims are correct.

## Working rules

- Keep repo guidance minimal and durable.
- Do not create extra planning files, TODO files, memory files, or research dumps in the repo.
- Use `TRACK.toml` for current state and ADRs for durable decisions.
- Prefer trigger-based skills and short instruction bodies with on-demand references.
- For stack/tool/MCP/language choices, research current evidence before implementing.

## Validation

- Do not invent build/test/run commands before they exist.
- For research/doc/rules changes, validate by internal consistency, file references, and decision alignment.
- When implementation begins, document real bootstrap/build/test/lint commands here only after they are verified.
