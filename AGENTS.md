# sane-next

## Product frame

- `sane-next` is a **Pi-first overlay/distribution**, not a fresh runtime or deep Pi fork.
- Pi is the primary runtime; Codex-native skill export is secondary.

## Startup rules

- Before broad work, read `TRACK.toml`; then follow only the roadmap, ADRs, standards, or source files referenced by the active slice unless the user has given an explicit current goal. If an explicit goal conflicts with the tracker, pause to reconcile it instead of silently switching to the tracker slice.
- For docs placement, use `docs/README.md` and `docs/standards/DOCS-STRUCTURE-STANDARD.md`; for implementation runs, load `.agents/skills/sane-next-implementation/SKILL.md`.
- Use current repo files and fresh external research; treat old Sane claims and stale chat memory as secondary hints.
- Keep always-on instructions small and update `TRACK.toml` only for the active slice.

## Hard boundaries

- Treat Pi as upstream runtime; build overlays, packs, adapters, and companion tools here.
- Keep UI work Pi-first and overlay-sized.
- Use `TRACK.toml`, `docs/roadmap/`, `docs/adr/`, `docs/standards/`, or `docs/reference/` as the owning surface for durable planning and facts.
- Place markdown docs inside the fixed repo structure; change the docs standard first when a new structure is needed.

## Done

- Current files reflect current truth.
- Repo docs stay small, user-facing where appropriate, and structured.
- The implementation path stays clear without chat context.
