# Implementation Slice Protocol

This standard defines how agents and contributors should run bounded implementation work in `sane-next` after the initial product baseline.

## Surface ownership

| Concern | Owner |
| --- | --- |
| Active execution window | `TRACK.toml` |
| Product direction and release readiness | `docs/roadmap/ROADMAP.md` |
| Tracking shape | `docs/standards/TRACK-STRUCTURE-STANDARD.md` |
| Documentation placement | `docs/standards/DOCS-STRUCTURE-STANDARD.md` |
| Durable decisions | `docs/adr/` |
| Instruction/skill/agent authoring | `docs/standards/INSTRUCTION-SURFACE-STANDARD.md` |
| Commit hygiene | `.githooks/` and `CONTRIBUTING.md` |

## Before broad work

1. Read `AGENTS.md` and `TRACK.toml`.
2. Follow `TRACK.toml` refs to only the roadmap, ADRs, standards, or source files needed for the active slice.
3. Check `git status --short` and preserve unrelated changes.
4. Inspect current source, tests, CLI help, or generated artifacts before documenting or changing behavior.

## Slice contract

Every active implementation slice should have:

- one concrete outcome
- one primary write boundary, or explicit non-overlapping lane boundaries
- source inputs and decision refs
- exact verification commands or inspection evidence
- stop conditions for risky, irreversible, or ambiguous decisions

Put that contract in `TRACK.toml`; do not create separate plan/TODO files.

## Parallel lane rules

1. One writing lane owns one path tree.
2. `TRACK.toml`, `docs/adr/`, and `docs/standards/` are coordinator-owned unless the slice is explicitly about those files.
3. Use read-only lanes for research or review when write boundaries would overlap.
4. Every lane returns changed files, verification evidence, and unresolved risk.
5. The coordinator integrates results and runs final verification from the combined repo state.

## Verification

Use the narrowest meaningful check first, then broaden when the change crosses boundaries.

Common checks:

```bash
cd cli && go test ./...
node --test pi-plugin/plugin.test.js
cd cli && ./acceptance.sh
```

Docs-only changes should still run the relevant validation path, such as hooks for tracking shape and markdown placement. CLI acceptance should stay focused on executable product behavior and generated artifacts, not roadmap prose.

## Commit and handoff

- Commit each meaningful verified milestone before switching task categories.
- If work pauses, hand off the current git state, changed files, verification run, and remaining blocker.
- Do not mark roadmap or tracking items complete until matching behavior exists and the named evidence passes.

## Boundaries

- Do not rebuild Pi's runtime or deep-fork Pi.
- Do not add a TUI-heavy standalone surface.
- Do not add markdown outside the fixed docs/skill/repo instruction structure unless `docs/standards/DOCS-STRUCTURE-STANDARD.md` is intentionally changed first.
- Do not store completed work, hypotheses, research dumps, or historical narrative in `TRACK.toml`.
