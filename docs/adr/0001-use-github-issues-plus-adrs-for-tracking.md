# ADR 0001: Use GitHub Issues plus ADRs for tracking

## Status

Superseded by ADR 0002

## Context

`sane-next` needs a tracking system that stays small, survives long Codex work, remains auditable, and does not recreate the document sprawl that hurt the current Sane repo.

The project needs all of these at once:

- clear current work state
- durable decision history
- low maintenance overhead
- good Git reviewability
- machine-usable structure where it actually helps
- resistance to bloated planning files

GitHub now provides issue hierarchies, dependencies, and project views. ADRs remain a proven way to keep design decisions short, reviewable, and traceable without turning the repo into a planning product.

Current Sane is a cautionary example: too many plans, specs, research notes, TODO surfaces, and state files make the repo harder to reload and easier to drift.

## Decision

Use a two-surface tracking system:

1. **GitHub Issues are the canonical task ledger.**
   - Use issues for active work, backlog, and milestones.
   - Use sub-issues and dependencies when work needs decomposition.
   - Use issue comments only for milestone checkpoints or resume handoffs, not turn-by-turn logs.
   - Do not add repo-local `TODO.md`, `plan.md`, task ledgers, or parallel tracking files.

2. **ADRs are the canonical decision history.**
   - Record only architecturally or workflow-significant decisions.
   - Keep each ADR short and explicit about context, decision, rejected alternatives, and consequences.
   - Do not use ADRs as a research dump or progress log.

Operational rule:

- **Commits and pull requests are the progress evidence layer.**
- Link commits and PRs to issues instead of maintaining a second handwritten changelog during early development.
- Do not adopt GitHub Projects, Linear, Plane, or a custom structured ledger unless issue volume or coordination pain proves that Issues alone are no longer enough.

## Rejected alternatives

### Markdown-only tracking

Rejected as the primary system. It is easy to read and review, but it tends to sprawl into multiple overlapping files, weak machine structure, stale checklists, and ambiguous source-of-truth problems.

### JSON or YAML as the primary tracker

Rejected as the primary system. It is machine-friendly, but it is worse for review, easy to overdesign, and encourages building a tracking product inside the repo before the product itself exists.

### Hybrid repo-local plan files plus structured metadata

Rejected for now. It can work, but it creates exactly the kind of dual-source planning surface that already drifted in Sane. The repo should not start with a local planning runtime.

### GitHub Projects as the primary tracker

Rejected as the default. Projects provides useful views and automation, but it adds extra metadata and workflow overhead before there is enough work to justify it. It remains an escalation path, not the starting point.

### Linear or Plane

Rejected. They are polished, but they add an external system and workflow dependency without giving enough benefit for a small Git-centric devtool project at this stage.

### Commit history as the only tracker

Rejected. Commits show what changed, but not open questions, blocked work, dependencies, or decision rationale.

### Agent memory files

Rejected as a repo surface. Long-running agents need resumable state, but that state should stay ephemeral or tool-managed unless there is proven need for a committed project memory format.

## Consequences

Positive:

- one clear current-work surface
- one clear decision-history surface
- low repo overhead
- good GitHub and API interoperability
- easy review and audit

Negative:

- issue context lives outside the repo checkout
- issue comments can still bloat if used carelessly
- some machine automation is weaker than a custom structured ledger

Follow-up rule:

- If active issue count, cross-cutting scheduling, or multi-lane coordination becomes hard to read in plain Issues, adopt a minimal GitHub Project view at that point instead of inventing a repo-local planner.
