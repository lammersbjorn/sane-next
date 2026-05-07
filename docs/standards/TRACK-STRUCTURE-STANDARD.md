# TRACK Structure Standard

This file defines the only allowed structure for `TRACK.toml`.

`TRACK.toml` is the repo's **active execution window**, not a backlog, not a history file, and not a hypothesis dump.

The whole-program verified checklist can live in `docs/roadmap/ROADMAP.md`. `TRACK.toml` should mirror only the current active phase from that roadmap.

## Hard rules

1. Keep `TRACK.toml` short enough to read at the start of a run.
2. Keep only active and near-term work in the file.
3. Remove completed work once the resulting repo state is committed.
4. Do not store hypotheses, research summaries, meeting notes, or historical narrative here.
5. Use ADRs for durable decisions and standards docs for durable structure/protocol rules.
6. Use `docs/roadmap/ROADMAP.md` for the broader verified roadmap when the project needs a whole-program execution ledger.

## Top-level schema

Required keys:

- `version`
- `updated`
- `phase`
- `goal`
- `next_focus`
- `[tracking]`
- `[current]`
- `[[work_items]]`

Forbidden legacy sections:

- `[[items]]`
- `[[hypotheses]]`

## `[tracking]`

Required keys:

- `schema`
- `canonical_current_state`
- `canonical_decision_history`
- `canonical_standards`
- `progress_evidence`
- `rules`

Purpose:

- define what `TRACK.toml` owns
- point to the other durable surfaces
- keep the operational rules visible without turning the file into doctrine

## `[current]`

Required keys:

- `summary`
- `exit_criteria`
- `refs`

Purpose:

- explain the current execution window in one short paragraph
- define the stop condition for the current phase
- point to the minimum durable files needed for the next run

## `[[work_items]]`

Each work item must contain:

- `id`
- `status`
- `title`
- `why_now`
- `write_scope`
- `inputs`
- `done_when`

Allowed `status` values:

- `ready`
- `in_progress`
- `blocked`
- `queued`

Rules:

1. At most **one** work item may be `in_progress`.
2. Keep at most **four** work items total so the active phase stays readable without pretending to hold the whole product backlog.
3. `write_scope`, `inputs`, and `done_when` must all be non-empty lists.
4. Each item should be resumable without reading unrelated chat history.
5. If an item needs more context than these fields can carry, the context belongs in a durable doc it can reference, not in a giant inline note.

## What does not belong in TRACK

- completed milestones
- rejected alternatives
- decision rationale
- open-ended hypotheses
- research digests
- duplicate task ledgers
- broad future backlog

## Relationship to ROADMAP

If the repo has a canonical roadmap file, use this split:

- `docs/roadmap/ROADMAP.md` = full program checklist with verification gates
- `TRACK.toml` = current active phase only

The roadmap may list the whole remake. `TRACK.toml` should list only the next bounded slice.

## Example

```toml
version = 2
updated = "2026-05-07"
phase = "implementation-ready"
goal = "Build the smallest reliable Codex workflow tool that improves long-running coding outcomes without recreating Sane's bloat."
next_focus = "Run phase 1 of the full remake."

[tracking]
schema = "active-window-v1"
canonical_current_state = "TRACK.toml"
canonical_decision_history = "docs/adr/"
canonical_standards = "docs/standards/"
progress_evidence = "git commits and pull requests"
rules = [
  "Keep only active-window state here.",
  "Do not store completed work or hypotheses here."
]

[current]
summary = "Planning is locked. The next run should execute phase 1 of the full remake."
exit_criteria = [
  "The active phase builds and validates cleanly."
]
refs = [
  "docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md"
]

[[work_items]]
id = "phase-1-foundations"
status = "ready"
title = "Create the first foundation surfaces"
why_now = "This is phase 1 of the broader remake."
write_scope = [".agents/skills/", "pi-plugin/", "cli/"]
inputs = ["docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md"]
done_when = ["The phase 1 acceptance path passes."]
```
