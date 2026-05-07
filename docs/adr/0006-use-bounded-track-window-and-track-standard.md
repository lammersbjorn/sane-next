# ADR 0006: Use a bounded active-window TRACK.toml and a fixed schema

## Status

Accepted

## Context

ADR 0002 correctly chose `TRACK.toml` plus ADRs as the tracking model, but it did not define the structure tightly enough.

That gap showed up immediately: the file could drift toward a long mixed ledger of completed tasks, vague future work, and floating hypotheses with weak context. If that happens, `TRACK.toml` becomes another bloated planning surface instead of the small continuation file it was meant to be.

The repo needs a stricter rule:

- `TRACK.toml` is the active execution window
- history belongs in git and ADRs
- structured task context must be mandatory, not optional

## Decision

Refine the tracking system from ADR 0002:

1. **`TRACK.toml` is a bounded active-window file, not a backlog or history log.**
   - It stores only the current phase, current focus, current exit criteria, and a small set of active or near-term work items.
   - Completed work is removed once its resulting state is committed.
   - Hypotheses, research summaries, and historical notes do not live in `TRACK.toml`.

2. **The `TRACK.toml` schema is owned by `docs/standards/TRACK-STRUCTURE-STANDARD.md`.**
   - Schema detail is a standard, not scattered convention.
   - Agents should update `TRACK.toml` to match that standard rather than improvise new fields.

3. **Each work item must be structured.**
   - Required fields: `id`, `status`, `title`, `why_now`, `write_scope`, `inputs`, `done_when`
   - The file must stay small enough to load quickly at the start of a run.

4. **Repo hooks should validate the shape.**
   - Parsing alone is not enough.
   - The committed pre-commit hook should reject legacy sections and malformed work items.

## Rejected alternatives

### Keep `TRACK.toml` loose and rely on contributor discipline

Rejected. The repo already has evidence that prose-only discipline drifts.

### Move backlog, hypotheses, and history into `TRACK.toml` with more sections

Rejected. That recreates the same sprawl problem in a structured file.

### Create a second repo-local planning file for richer task context

Rejected. That would split current-state truth again.

## Consequences

Positive:

- `TRACK.toml` stays fast to read and easy to trust
- tasks carry enough context to resume without turning into essays
- continuation quality improves without creating a second planning product

Negative:

- broader backlog and risk material must live somewhere else or stay out of the repo until needed
- agents must trim the file as milestones land instead of treating it like append-only storage
