---
name: agent-lanes
description: "Use this skill when broad coding work can be split into independent subagent lanes for faster research, implementation, review, or verification."
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
- Pi or Codex exposes subagent execution for the current session, such as Sane's recommended `pi-subagents` Pi package

## Use Main Session When

- the task is a small direct edit
- one blocking fact determines the next step and the main session can inspect it faster
- write boundaries overlap or need one writer
- the user has paused or disallowed delegation

## Inputs

- user objective and stop conditions
- repo instructions and active tracking files
- candidate lane write boundaries
- verification commands for each lane
- available subagent models, skills, and permissions

## Outputs

- a compact lane table with owner, scope, write boundary, and verification
- launched subagent tasks through the available runtime package when they can run independently
- returned summaries with changed files, evidence, and unresolved risk
- coordinator integration and final verification

## How To Run

1. Check whether the runtime exposes a subagent package/tool first. In Sane's default Pi install, prefer the curated `pi-subagents` package when available; `/sane-status`, `/subagents-status`, or `/subagents-doctor` can confirm availability. Otherwise use a lane table for manual/tmux delegation.
2. Decide the main thread's immediate blocking task before delegating.
3. For broad work with independent research, review, verification, or disjoint implementation slices, launch focused `pi-subagents` lanes.
4. Split work only when each lane can run independently while the main thread continues.
5. Give every lane one clear owner and one write boundary.
6. Use read-only explorer lanes for broad discovery and verifier lanes for fresh review.
7. Use implementation lanes with non-overlapping write sets.
8. Ask subagents for summaries and evidence, rather than full logs or broad file dumps.
9. Integrate results in the main thread and run final verification before claiming completion.

## Verification

- each lane reports its own verification or explicit limits
- the coordinator reviews changed files and resolves conflicts
- final verification runs from the integrated repo state
- roadmap or tracking boxes are checked after integrated verification passes

## Gotchas / Safety

- coordinator owns final decisions, conflict resolution, and completion claims
- parallel work helps when dependencies are real and write sets are clean
- fresh-review lanes start from repo truth and inspect independently
- use persistent memory only when the repo explicitly owns it

## Examples

### Positive

- Launch three read-only explorers for API, database, and frontend discovery while the main thread prepares a fix.
- Use a verifier lane to review ownership-boundary behavior before committing install/uninstall changes.

### Negative

- Spawn a worker for a change whose result is needed before the next local command.
- Let two workers edit the same config file without one clear owner.
