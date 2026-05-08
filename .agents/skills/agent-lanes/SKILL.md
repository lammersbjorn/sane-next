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
- the active agent runtime exposes a real delegation mechanism

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
- available delegation tools, models, permissions, and isolation options

## Outputs

- a compact lane table with owner, scope, write boundary, and verification
- launched lanes through the active runtime only when real delegation is available
- returned summaries with changed files, evidence, and unresolved risk
- coordinator integration and final verification

## How To Run

1. Identify the active runtime and its actual delegation tool, if one is installed and enabled.
2. If no runtime delegation tool is available, write a lane table and execute the work in the main session or hand it to the user as a manual/tmux plan.
3. Decide the main thread's immediate blocking task before delegating.
4. For broad research, broad review, broad verification, or disjoint implementation, launch focused lanes only when they can run independently.
5. Give every lane one clear owner and one write boundary.
6. Use read-only explorer lanes for broad discovery and verifier lanes for fresh review.
7. Use implementation lanes only when write sets do not overlap.
8. Ask delegated lanes for summaries and evidence, not full logs or broad file dumps.
9. Integrate results in the main thread and run final verification before claiming completion.

## Verification

- each lane reports its own verification, evidence, or explicit limits
- the coordinator reviews changed files and resolves conflicts
- final verification runs from the integrated repo state
- roadmap or tracking boxes are checked after integrated verification passes

## Gotchas / Safety

- coordinator owns final decisions, conflict resolution, and completion claims
- runtime-specific delegation commands belong in runtime docs or companion tooling, not in this shared skill body
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
