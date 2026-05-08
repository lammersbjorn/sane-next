# sane-next repository instructions

Use `AGENTS.md` as the primary repo instruction surface.

Additional repo-wide notes:

- This repo is implemented but still pre-stable.
- Before broad work, read `TRACK.toml` first; then follow only the roadmap, ADRs, standards, or source files referenced by the active slice.
- For implementation runs, load `.agents/skills/sane-next-implementation/SKILL.md` and verify behavior with current source, tests, and CLI help.
- Do not rebuild Pi's runtime, deep-fork Pi, or add a TUI-heavy standalone surface.
- Do not invent build, test, or run commands; inspect current source or `--help` first.
- Put docs in the existing structure described by `docs/README.md` and `docs/standards/DOCS-STRUCTURE-STANDARD.md`.
- If you create commits, use the repo hooks in `.githooks/` and the commit convention in `CONTRIBUTING.md`.
