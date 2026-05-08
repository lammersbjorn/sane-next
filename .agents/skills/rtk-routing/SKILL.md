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

## Don't Use When

- the repo does not use RTK and a direct command is clearer
- an interactive program needs raw terminal behavior RTK cannot provide
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
   - `off`: do not route solely for RTK policy.
   - `advise`: prefer RTK when it is clearly helpful, but raw shell is allowed.
   - `warn`: prefer RTK for direct equivalents and expect nudges, but raw shell is still allowed.
   - `enforce`: use RTK routes for commands covered by policy; raw covered commands may be blocked.
2. Prefer RTK-native commands for common work: `rtk grep`, `rtk read`, `rtk ls`, `rtk tree`, `rtk find`, `rtk diff`, `rtk git`, `rtk test`, `rtk lint`, and package-specific RTK wrappers.
3. Use `rtk run` only when exact shell semantics matter or no native RTK command exists.
4. Keep searches targeted before broadening.
5. Preserve exact paths, flags, and error text when reporting failures.
6. If RTK is missing, report the missing dependency and do not silently bypass a repo policy that requires enforcement.

## Verification

- command output answers the intended question or proves the intended behavior
- failures include enough exact output to choose the next command
- broad verification commands are matched to the changed behavior

## Gotchas / Safety

- do not use RTK as a reason to run destructive commands casually
- do not wrap an RTK-native command in another shell layer unless the wrapper is required
- do not treat compact output as complete evidence when full logs are needed for a bug

## Examples

### Positive

- Use `rtk grep "handler" cli` to inspect a code path.
- Use `rtk git diff -- cli` before committing CLI changes.

### Negative

- Run a raw `grep -R` when `[rtk].mode = "enforce"` or repo instructions require RTK.
- Ignore an RTK hook suggestion and retry the same rejected command.
