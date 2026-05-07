---
name: core-workflow
description: "Use this skill when a coding task needs Sane's compact long-run workflow discipline: repo truth first, bounded execution, concrete verification, and recoverable handoff."
license: MIT
compatibility: Pi, Codex
---

# Core Workflow

## Goal

Keep long-running coding work grounded in current repo truth, bounded phases, and concrete verification.

## Use When

- implementing or repairing code over multiple steps
- resuming work after a pause, failure, or context transition
- the repo has its own tracking, roadmap, ADR, or instruction files
- verification and recovery matter more than a quick answer

## Don't Use When

- the user asks a small standalone question
- the task is a one-command lookup
- the repo has a stricter local skill or instruction surface for the same work

## Inputs

- repo-local agent instructions
- active tracking or roadmap files
- current git status and diffs
- relevant source files and tests
- user-provided objective and constraints

## Outputs

- narrowly scoped code or docs changes
- verification output tied to the requested behavior
- an explicit blocker or handoff when work cannot continue safely

## How To Run

1. Read the repo's always-on instructions and active tracking files before broad work.
2. Check `git status` and preserve unrelated user changes.
3. Convert the user request into concrete deliverables and stop conditions.
4. Inspect the current implementation before choosing an approach.
5. Edit within the smallest write boundary that can satisfy the request.
6. Run the narrowest meaningful verification first, then broaden only when risk requires it.
7. Commit or hand off only after the matching behavior is verified.

## Verification

- cite the command or inspection that proves the requested behavior
- make sure passing tests cover the behavior being claimed
- treat missing tests, skipped checks, or proxy signals as residual risk
- rerun verification after any fix that could change the outcome

## Gotchas / Safety

- do not replace repo truth with old chat memory
- do not expand always-on instructions to solve a one-off problem
- do not revert unrelated work in a dirty tree
- do not mark tracking or roadmap items complete before matching behavior exists

## Examples

### Positive

- "Read the current tracker, implement the next unchecked phase, verify it, and commit."
- "Resume from git state and the roadmap instead of restarting the plan."

### Negative

- "Invent a new planning file because the next step is unclear."
- "Declare completion because the scaffold exists but no command proves behavior."
