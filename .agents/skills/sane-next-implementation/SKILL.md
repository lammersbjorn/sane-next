---
name: sane-next-implementation
description: Use this skill when implementing sane-next itself, especially for the future /goal build run or near-one-shot implementation sessions. Use it for phase-based execution, pack/export work, Pi plugin work, companion CLI work, and bounded continuation. Don't use it for old-Sane maintenance, broad pre-implementation research, or unrelated repos.
license: MIT
compatibility: Pi, Codex
---

# Sane Next Implementation

## Goal

Implement `sane-next` as a full remake with bounded phases, not as an unbounded wandering build and not as a tiny accidental MVP.

## Use When

- working on `sane-next` implementation itself
- running the future `/goal` or near-one-shot build run
- building phase 1 foundations, pack export, extensibility, or CLI lifecycle flows
- resuming implementation after a pause, failure, or budget hit

## Don't Use When

- maintaining the old `sane` repo
- doing pre-implementation landscape research
- writing generic docs for another repo
- doing frontend-only work unrelated to `sane-next`

## Inputs

- `TRACK.toml`
- `docs/roadmap/ROADMAP.md`
- `docs/reference/OLD-SANE-POSTMORTEM.md`
- core ADRs in `docs/adr/`
- standards in `docs/standards/`
- current repo state and git history

## Outputs

- changes only inside the current phase write boundary
- explicit verification results for the phase
- a milestone commit when the phase or sub-phase lands
- updated `TRACK.toml` if the active phase changes

## How To Run

1. Read `TRACK.toml` first and identify the active phase plus write boundaries.
2. Read `docs/roadmap/ROADMAP.md` to understand the full remake scope and the verified checklist.
3. Read the old-Sane postmortem and the relevant ADRs/standards before choosing the approach.
4. Treat the product scope as the full remake: packs, extensibility, Pi integration, Codex export, and CLI lifecycle all matter.
5. Treat the active phase as bounded execution, not as the full product scope.
6. If running with `/goal`, use the concrete `/goal` text from `docs/roadmap/ROADMAP.md`.
7. Keep `AGENTS.md` small; put recurring procedure detail here or in referenced docs instead of expanding always-on instructions.
8. Prefer one clear write boundary per lane. If parallelizing, keep file ownership disjoint.
9. Commit after meaningful milestones so git becomes the primary resume anchor.
10. If the run pauses or hits budget, resume from git state first, then continue the remaining unchecked verified work from the roadmap.

## Verification

- use the exact verification required by the active phase in `TRACK.toml` and `docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md`
- do not claim a phase is done without the named build/test/acceptance checks for that phase
- review `/diff` and git history before closing a phase

## Gotchas / Safety

- do not copy old Sane code into `sane-next`
- do not let the run narrow the full remake into "just the first slice"
- do not let packs/extensibility disappear from the plan
- do not check roadmap boxes without matching verification
- do not enable broad MCP/tool surfaces unless they are clearly justified
- do not treat the old `sane` repo as a writable workspace
- do not grow `TRACK.toml` into a backlog or research dump

## Examples

### Positive

- "Implement phase 1 of sane-next, then stop with a milestone commit and updated TRACK."
- "Resume the sane-next /goal run after a budget limit and continue phase 2 only."
- "Build the pack export path without reaching into Pi internals."

### Negative

- "Refactor the old sane repo to match sane-next ideas."
- "Add five overlapping implementation skills."
- "Expand AGENTS.md with the full execution procedure."
