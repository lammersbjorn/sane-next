# ADR 0002: Use TRACK.toml plus ADRs for tracking

## Status

Accepted; refined by ADR 0006, which makes `TRACK.toml` a bounded active-window file and forbids hypotheses/history there.

## Context

The first tracking decision favored GitHub Issues plus ADRs, but that is no longer the right fit for this project. The repo needs a canonical task ledger that lives locally, remains easy for both humans and agents to load, and does not depend on issue-based workflow as the primary interface.

The design constraints are unchanged:

- small and low-bloat
- auditable progress
- traceable decisions
- strong multi-session continuation
- Git-friendly review
- machine-usable structure where it helps

The lesson from current Sane still stands: too many overlapping planning files, research notes, TODO surfaces, and runtime state create drift. The replacement system must therefore be repo-local, but also narrow.

## Decision

Use a two-surface repo-local system:

1. **`TRACK.toml` is the canonical task and current-state ledger.**
   - It stores only active and near-term work.
   - It may include concise hypothesis status, next focus, and milestone references.
   - It is current-state oriented, not a historical log.
   - It replaces repo-local `TODO.md`, `plan.md`, task ledgers, and agent memory files.

2. **ADRs are the canonical decision history.**
   - Record only significant product, workflow, and architecture decisions.
   - Keep ADRs short and explicit about context, decision, rejected alternatives, and consequences.
   - Do not use ADRs as a task board or research dump.

Operational rule:

- **Commits and pull requests are the progress evidence layer.**
- Optional external mirrors are allowed later, but they must be generated from or reconciled back to `TRACK.toml`.
- External mirrors are never the source of truth.

## Rejected alternatives

### GitHub Issues as the primary tracker

Rejected. They are useful, but they push the canonical task ledger outside the repo and force an issue-based workflow that is not wanted here.

### Markdown-only tracking

Rejected as the primary system. It reads well, but it tends to sprawl into overlapping files and weakly structured state.

### JSON or YAML as the primary tracker

Rejected. Both are workable, but TOML is a better fit for this repo's need for small, readable, comment-free structured state without JSON noise or YAML footguns.

### Hybrid repo-local plan files plus structured metadata

Rejected. That recreates the dual-source planning problem that already hurt Sane.

### Commit history as the only tracker

Rejected. It is evidence, not a live task ledger.

## Consequences

Positive:

- canonical work state lives in the repo
- one obvious file for continuation
- structured enough for agent use without building a planner product
- lower risk of uncontrolled planning sprawl

Negative:

- `TRACK.toml` must stay disciplined or it can still grow badly
- external collaboration views are weaker until a mirror is added
- there is less built-in workflow automation than GitHub Projects or Linear

Follow-up rule:

- If `TRACK.toml` becomes crowded, split by rule rather than by habit: keep one canonical current-state file until there is concrete evidence that one file can no longer stay readable.
