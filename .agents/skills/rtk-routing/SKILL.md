---
name: rtk-routing
description: "Use this skill when a repo requires RTK or shell, search, test, diff, and log work should be routed through compact RTK commands."
license: MIT
compatibility: Pi, Codex
---

# RTK Routing

## Goal

Keep command execution compact, policy-aware, and repo-compatible by routing shell work through RTK according to the configured RTK mode.

## Use When

- repo instructions require `rtk`
- Sane config sets `[rtk].mode` to `advise`, `warn`, or `enforce`
- search, test, diff, log, or file inspection output could become noisy
- a command has a direct RTK equivalent
- a hook rejects raw shell commands and suggests an RTK route

## Use Direct Shell When

- the repo has no RTK policy and a direct command is clearer
- an interactive program needs raw terminal behavior outside RTK coverage
- the user explicitly asks for a non-RTK command and repo policy allows it

## Inputs

- repo shell policy from `AGENTS.md`, `RTK.md`, `[rtk].mode`, or equivalent
- the command the task requires
- RTK help or rewrite output when available
- current working directory

## Outputs

- the smallest RTK command that answers the task
- exact command failures when RTK cannot route the work
- a fallback only when no RTK-native command fits

## How To Run

1. Read the active RTK mode when available:
   - `off`: use direct shell based on normal task fit.
   - `advise`: prefer RTK when it is clearly helpful, direct shell remains allowed.
   - `warn`: prefer RTK for direct equivalents and expect nudges, treat nudges as routing hints.
   - `enforce`: use RTK routes for commands covered by policy; raw covered commands may be blocked.
2. Prefer RTK-native commands for common work: `rtk grep`, `rtk read`, `rtk ls`, `rtk tree`, `rtk find`, `rtk diff`, `rtk git`, `rtk test`, `rtk lint`, and package-specific RTK wrappers.
3. Use `rtk run` when exact shell semantics matter or no native RTK command exists.
4. Keep searches targeted before broadening.
5. Preserve exact paths, flags, and error text when reporting failures.
6. When RTK is missing under an enforcing repo policy, report the missing dependency and stop for guidance.

## Verification

- command output answers the intended question or proves the intended behavior
- failures include enough exact output to choose the next command
- broad verification commands are matched to the changed behavior

## Gotchas / Safety

- Use destructive commands only with the same approval and care required outside RTK.
- Run RTK-native commands directly unless a wrapper is required.
- Treat compact output as a summary; request full logs when a bug needs complete evidence.

## Examples

### Positive

- Use `rtk grep "handler" cli` to inspect a code path.
- Use `rtk git diff -- cli` before committing CLI changes.

### Route Correction

- In `[rtk].mode = "enforce"`, route recursive search through `rtk grep`.
- When an RTK hook suggests a route, retry with the suggested RTK command.
