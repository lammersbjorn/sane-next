# sane-next

## Product frame

- `sane-next` is a **Pi-first overlay/distribution**, not a fresh runtime or deep Pi fork.
- Pi is the primary runtime; Codex-native skill export is secondary.

## Startup rules

- Before broad work, read `TRACK.toml`, `README.md`, the roadmap, core ADRs, and standards listed there.
- For implementation runs, load `.agents/skills/sane-next-implementation/SKILL.md`.
- Use current repo files and fresh external research; do not trust old Sane claims or stale chat memory.
- Keep always-on instructions small and update `TRACK.toml` only for the active slice.

## Hard boundaries

- Do not rebuild Pi's runtime in this repo.
- Do not add a TUI-heavy standalone surface.
- Do not create parallel `TODO`, `plan`, `memory`, ADR, or research-dump files.
- Do not add markdown docs outside the fixed repo structure without intentionally changing it first.

## Done

- Current files reflect current truth.
- Repo docs stay small, user-facing where appropriate, and structured.
- The implementation path stays clear without chat context.
