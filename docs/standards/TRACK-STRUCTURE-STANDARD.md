# TRACK Structure Standard

This file defines the only allowed structure for `TRACK.toml`.

`TRACK.toml` is the repo's **active execution window**, not a backlog, not a history file, and not a hypothesis dump.

Broader product direction can live in `docs/roadmap/ROADMAP.md`. `TRACK.toml` should mirror only the current active phase from that roadmap.

## Hard rules

1. Keep `TRACK.toml` short enough to read at the start of a run.
2. Keep only active and near-term work in the file.
3. Remove completed work once the resulting repo state is committed.
4. Do not store hypotheses, research summaries, meeting notes, or historical narrative here.
5. Use ADRs for durable decisions and standards docs for durable structure/protocol rules.
6. Use `docs/roadmap/ROADMAP.md` for broader product direction and evidence when the active slice needs roadmap context.

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

- `docs/roadmap/ROADMAP.md` = broader product direction, release discipline, and evidence
- `TRACK.toml` = current active phase only

The roadmap may list broader priorities. `TRACK.toml` should list only the next bounded slice.

## Example

```toml
version = 2
updated = "2026-05-08"
phase = "release-readiness-complete"
goal = "Keep the pre-stable release-readiness baseline verified while choosing any later release automation deliberately."
next_focus = "Wait for the next explicit release or implementation slice."

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
summary = "No active release-readiness implementation slice is open."
exit_criteria = [
  "The active phase builds and validates cleanly."
]
refs = [
  "docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md"
]

[[work_items]]
id = "await-next-release-slice"
status = "queued"
title = "Await next bounded release or implementation slice"
why_now = "Further work should be driven by an explicit release artifact, automation, or product objective."
write_scope = ["TRACK.toml", "docs/roadmap/ROADMAP.md", "docs/adr/", "README.md"]
inputs = ["docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md"]
done_when = ["A new active slice is chosen only when there is an explicit next objective."]
```
