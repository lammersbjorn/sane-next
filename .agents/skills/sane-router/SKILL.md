---
name: sane-router
description: "Use this skill when a Sane-enabled session needs to choose between direct work, RTK routing, agent lanes, pack export, or lifecycle commands."
license: MIT
compatibility: Pi, Codex
---

# Sane Router

## Goal

Pick the smallest Sane surface that fits the task: direct main-thread work, RTK routing, subagent lanes, pack export, or lifecycle management.

## Use When

- a task is broad enough that routing matters
- active packs expose multiple workflow skills
- the user asks for subagents, RTK, lifecycle commands, or pack export
- a Pi or Codex session needs Sane's workflow defaults without loading every pack body

## Don't Use When

- a specific skill already clearly matches
- the task is a one-line answer or one local command
- repo instructions define a stricter route

## Inputs

- active Sane config
- enabled packs and export targets
- user request and repo instructions
- available runtime capabilities in Pi or Codex

## Outputs

- one routing choice and the reason for it
- optional skill names to load next
- optional lifecycle command to run
- clear fallback when a runtime capability is unavailable

## How To Run

1. If the task is tiny, keep it in the main thread.
2. If shell/search/test/log work is involved and RTK is enabled or required, load `rtk-routing`.
3. If work can run in parallel with clean boundaries, load `agent-lanes`; when `pi-subagents` is available, actually delegate focused lanes instead of only planning them.
4. If the task is an ongoing resume or implementation run, load `core-workflow`.
5. If the task is install/export/update/doctor/repair/uninstall, use the Sane companion CLI.
6. Keep only the selected skill bodies in context.

## Verification

- the chosen route matches the task shape and enabled pack config
- lifecycle commands are run in fixture or real targets as appropriate
- subagent lane results are integrated and verified by the coordinator

## Gotchas / Safety

- do not load every pack just because it exists
- do not make subagents mandatory for small tasks
- do not bypass repo-local shell policy
- do not use Sane as a separate runtime when Pi owns the agent loop

## Examples

### Positive

- A broad refactor with independent modules loads `agent-lanes`.
- A repo requiring RTK loads `rtk-routing` before shell work.

### Negative

- Loading all packs for a small typo fix.
- Starting a standalone Sane workflow when Pi or Codex can run the selected skill directly.
