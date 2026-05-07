# Goal Run Standard

This file defines the recommended setup for the future Codex implementation run that will try to one-shot or near-one-shot `sane-next` using the experimental `/goal` command.

This is about the **implementation run stack**, not the final `sane-next` product stack.

## Verdict

Do **not** rely on `/goal` alone.

Use a **near-one-shot** setup:

1. prepare repo-local context first
2. load one focused custom skill
3. start the session with a tight implementation prompt
4. set `/goal` immediately after startup as the run-level objective and budget monitor

`/goal` is useful as:

- a session-level objective anchor
- a token/time budget surface
- a simple stop-state machine

It is **not** a proven full planner/orchestrator by itself.

## Model and reasoning

### Model

- **Preferred:** `gpt-5.5`
- **Fallback:** `gpt-5.4`

Use `gpt-5.5` only if the environment actually has access to it through ChatGPT sign-in.

### Reasoning

- **Default:** `low`
- **Escalate to:** `medium` for tricky architecture/debug loops
- **Reserve `high` for:** isolated verification or hard reasoning turns, not the whole build run

The implementation run should optimize for steady progress and low drift, not maximum deliberation on every turn.

## Custom skill requirement

Create and use one focused repo-local skill before the run:

- `.agents/skills/sane-next-implementation/SKILL.md`

Why:

- keeps `AGENTS.md` small
- moves recurring implementation procedure out of the `/goal` prompt
- gives the run a reusable trigger-loaded execution contract

Do **not** create multiple overlapping implementation skills for the same run.

## Skill recommendations

### Enable

- `sane-next-implementation`
- `sane-outcome-continuation`
- `sane-agent-lanes`
- `sane-rtk` only when the target repo uses RTK

### Disable unless explicitly needed

- `sane-self-hosting`
- `sane-bootstrap-research`
- frontend-focused skills
- docs-writing skills
- any skill aimed at old-Sane maintenance rather than `sane-next` implementation

## MCP and tool policy for the run

### Enable only if useful

- **GitHub MCP** for repo/PR operations
- **Context7** for current package/framework docs
- **OpenAI docs MCP** or equivalent official doc lookup if available

### Keep disabled or tightly scoped

- browser/Playwright MCPs by default
- Sentry, Figma, or design MCPs
- extra write-capable filesystem MCPs
- broad web/search tooling that encourages wandering

### General rule

- prefer native file/git/shell tools first
- keep MCP count minimal
- keep networked tools opt-in and purpose-bound

## Old Sane repo context

The old `sane` repo should be:

- available as **read-only reference only**
- attached by specific files when needed
- **not** added as a writable workspace root

Do not let the implementation run treat the old repo as the base to copy.

## Instruction placement

Split by durability and scope:

| Content | Owner |
| --- | --- |
| run-level objective, stop clause, budget | `/goal` |
| tiny startup rules | `AGENTS.md` |
| implementation procedure | `.agents/skills/sane-next-implementation/SKILL.md` |
| durable decisions | `docs/adr/` |
| durable execution standards | `docs/standards/` |
| current active phase | `TRACK.toml` |

### What belongs in `/goal`

Only the compressed run objective:

- what to build now
- what not to do
- when to stop
- optional token budget

### What does **not** belong in `/goal`

- long architecture history
- detailed skill rules
- large file maps
- broad research dumps

## Progress checks during the run

Use this pattern:

1. `/status` at start and after major pivots
2. `/goal` to inspect budget/time state
3. `/diff` to check whether file changes are real and scoped
4. `git log --oneline` for milestone progress
5. commit after meaningful milestones
6. `/compact` only after a commit if context is getting too large

Git commits are the primary durable progress signal.

## Scope creep controls

Use all of these together:

1. explicit "do not" clauses in `/goal`
2. hard boundaries in `AGENTS.md`
3. `Don't Use When` / stop-boundary language in the implementation skill
4. `approval_policy = "on-request"`
5. bounded write scopes in `TRACK.toml` and the implementation protocol

## Stop conditions

The run should stop when all of these are satisfied:

1. the `/goal` stop criteria are met
2. the relevant `TRACK.toml` active phase is done
3. the required verification for that phase has passed
4. a milestone commit exists

If a token budget is set, `BudgetLimited` is also a valid automatic stop.

## Resume rules

Primary resume anchor: **git state**

Resume pattern:

1. inspect `git log --oneline`
2. reopen the saved session with `/resume` if useful
3. update `/goal` to the **remaining work**, not the original full brief
4. use `/compact` only after the last good milestone is committed

## Full remake vs. active phase

The implementation run is for a **full remake of Sane that is better**, including:

- shared packs
- extensibility / user-added packs
- Pi integration
- Codex export
- small companion CLI lifecycle flows

The fact that `TRACK.toml` starts with a bounded first phase does **not** mean the product scope is just a tiny MVP. It means the execution window is bounded while the full product scope remains explicit in repo docs.

## Recommended config shape

```toml
model = "gpt-5.5"
model_reasoning_effort = "low"
approval_policy = "on-request"
sandbox_mode = "workspace-write"
web_search = "cached"
```

Keep this run conservative. The point is not maximum tool availability; the point is a controlled implementation run with low drift.
