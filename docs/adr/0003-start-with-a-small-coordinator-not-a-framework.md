# ADR 0003: Start with a small coordinator, not a framework

## Status

Accepted

## Context

`sane-next` is meant to be a clean-sheet successor to Sane, not a port of the current repo.

The strongest lessons from the current Sane repo and the external workflow research are consistent:

- the vision is useful, but the current implementation surface grew faster than the evidence
- large always-on instruction surfaces and tracking sprawl create cost, drift, and reload pain
- long-running coding agents need resumability, stop conditions, and recovery signals
- small helper models are useful, but complex routing systems can become a product of their own
- parallel work is only reliable when ownership and isolation are explicit

Current Sane is a warning against overbuilding. The repo now has 382 tracked files, 87 tracked Markdown files, 16 research notes, 7 specs, 3 plans, 12 repo skill files, a live `TODO.md`, and a `.sane` runtime surface. Some of that structure is thoughtful, but together it is too much system for the amount of proven user value.

## Decision

Build `sane-next` as a small coordination system for Codex-centric coding workflows.

### Product shape

- Start **without** a TUI-first control center.
- Start **without** a plugin framework.
- Start **without** a broad pack/export system.
- Prefer a small config-first tool with strong defaults and explicit workflow discipline.

### Runtime shape

- Use **one coordinator** as the default operating model.
- Add bounded worker lanes only when the work clearly decomposes.
- Keep one verifier/reviewer authority instead of multiple peers debating.
- Use worktrees only for parallel write lanes or risky isolation, not by default for every task.

### Model policy

- Start with **one strong default model** for most work.
- Add at most one cheap/fast helper lane for read-heavy fanout or fast iteration once tests show clear benefit.
- Do **not** begin with a heavy routing matrix across many model roles.
- Escalate model complexity only where evidence is strong, such as high-risk verification or visual/UI review.

### State and recovery

- Keep committed state minimal: `TRACK.toml` plus ADRs.
- Keep execution state ephemeral or tool-managed until there is proven need for durable committed runtime state.
- Design the implementation around explicit checkpoints, idempotent actions, stall detection, and resume-from-known-state behavior.
- Human checkpoints should be built in at milestone boundaries, not bolted on later.

### Delivery discipline

- The main value is workflow quality, not orchestration novelty.
- Prefer clear stop conditions, narrow scopes, and verification over autonomous sprawl.
- If a workflow problem can be fixed with clearer boundaries, tests, or better defaults, do that before adding more framework code.

## Rejected alternatives

### TUI-first successor

Rejected for the first implementation. Current evidence does not show that a setup-heavy control surface is the highest-leverage starting point for Codex subscription optimization.

### Framework-first successor

Rejected. A large abstraction stack for packs, exports, plugins, and multi-surface management would repeat the main failure mode of the current repo.

### Many-role model router from day one

Rejected. The evidence supports task-shaped escalation, but not a large early routing framework. Start with one strong default and earn every extra lane.

### Repo-local durable memory system from day one

Rejected. Durable execution matters, but a committed memory/runtime format is premature before the actual failure modes of `sane-next` are observed.

## Hypothesis classification

| Hypothesis | Status | Reason |
| --- | --- | --- |
| Sane’s vision is good, but the current version became too bloated. | **Confirmed** | Repo evidence shows useful goals under too many surfaces, abstractions, and planning artifacts. |
| The new project should be smaller and more focused. | **Confirmed** | This is the clearest lesson from the postmortem and public workflow signal. |
| A small config file may be better than a TUI-heavy config direction. | **Likely** | The problem is workflow discipline and defaults more than a rich setup surface. |
| GPT-5.5 low may be enough for many tasks. | **Likely** | Current OpenAI and workflow commentary point toward low-effort defaults for ordinary work. |
| GPT-5.5 low may reduce tool use and cost. | **Plausible** | Fewer tokens and less overthinking are plausible, but this should be measured in `sane-next` itself. |
| Model routing may be less useful than one strong default. | **Likely** | Early routing complexity is easier to overbuild than to justify. |
| Pi Agent-style architecture may be efficient. | **Plausible** | A small coordinator plus bounded workers fits the evidence, but still needs direct testing. |
| The most important feature may be workflow discipline, not code complexity. | **Confirmed** | Both the Sane postmortem and external research point here. |
| The tracking/planning system may make or break the project. | **Confirmed** | The current repo’s drift strongly supports this. |

## Consequences

Positive:

- much smaller starting surface
- less architectural lock-in
- easier to test what actually matters
- lower risk of building another self-referential framework

Negative:

- some conveniences from current Sane are delayed or dropped
- later expansion must be earned with evidence instead of front-loaded
- the first implementation run must stay disciplined and resist feature creep

## Notes

This ADR sets the starting direction, not the final product scope forever. If later evidence shows that a richer control surface, committed runtime state, or broader model policy is truly needed, that should be a new ADR rather than an implicit drift.
