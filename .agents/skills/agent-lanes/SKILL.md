---
name: agent-lanes
description: Use this skill when broad coding work can be split into independent subagent lanes for faster research, implementation, review, or verification.
license: MIT
compatibility: Pi, Codex
---

# Agent Lanes

## Goal

Use subagents deliberately for speed, context isolation, and independent review without turning every task into orchestration overhead.

## Use When

- research spans independent areas that can be explored in parallel
- implementation has disjoint write boundaries
- verification needs a fresh perspective before a commit
- logs, tests, or generated output would flood the main context
- Pi or Codex exposes subagent execution for the current session

## Don't Use When

- the task is a small direct edit
- the next step depends on one blocking fact the main session can inspect faster
- write boundaries overlap and cannot be separated safely
- the user has paused or disallowed delegation

## Inputs

- user objective and stop conditions
- repo instructions and active tracking files
- candidate lane write boundaries
- verification commands for each lane
- available subagent models, skills, and permissions

## Outputs

- a compact lane table with owner, scope, write boundary, and verification
- launched subagent tasks only when they can run independently
- returned summaries with changed files, evidence, and unresolved risk
- coordinator integration and final verification

## How To Run

1. Decide the main thread's immediate blocking task before delegating.
2. Split only work that is independent enough to run without waiting on the main task.
3. Give every lane one clear owner and one write boundary.
4. Prefer read-only explorer lanes for broad discovery and verifier lanes for fresh review.
5. Use implementation lanes only when their write sets do not overlap.
6. Ask subagents for summaries and evidence, not full logs or broad file dumps.
7. Integrate results in the main thread and run final verification before claiming completion.

## Verification

- each lane reports its own verification or explicit limits
- the coordinator reviews changed files and resolves conflicts
- final verification runs from the integrated repo state
- roadmap or tracking boxes are checked only after integrated verification passes

## Gotchas / Safety

- subagents do not remove coordinator responsibility
- parallel work is useful only when dependencies are real and write sets are clean
- fresh-review lanes should not inherit the coordinator's assumptions
- avoid persistent memory by default unless the repo explicitly owns it

## Examples

### Positive

- Launch three read-only explorers for API, database, and frontend discovery while the main thread prepares a fix.
- Use a verifier lane to review ownership-boundary behavior before committing install/uninstall changes.

### Negative

- Spawn a worker for a change whose result is needed before the next local command.
- Let two workers edit the same config file without one clear owner.
