# ADR 0008: Use a small tooling policy and optional MCPs

## Status

Accepted

## Context

Research consistently pointed to the same failure mode: tool and MCP sprawl damages reliability, increases prompt surface area, and creates more overlap than value.

`sane-next` should optimize the coding workflow, not turn into a marketplace of loosely trusted integrations.

At the same time, some tools clearly help:

- RTK-style task-aware shell routing
- targeted external documentation/search when the repo is not enough
- browser verification for UI work

The policy needs to keep the default system small while leaving room for curated opt-in integrations.

## Decision

Use a three-tier policy:

1. **Default core**
   - No third-party MCP is required for `sane-next` to work.
   - The default system must work with the host runtime, local repo tools, and standard shell/file/git capabilities.

2. **Recommended helpers**
   - **RTK** is a first-class recommended integration when the target repo uses RTK.
   - **Playwright CLI** is preferred over a browser MCP for frontend verification.
   - **Context7** and **grep.app** are recommended only when current external docs or broader code search are genuinely needed.

3. **Optional curated integrations**
   - Any additional MCP must be explicit, opt-in, pack-scoped where possible, and documented with provenance and purpose.
   - Pi owns the runtime bridge; `sane-next` should configure or document integrations, not rebuild the bridge.

Hard boundaries:

- do not require marketplace-style MCP bundles by default
- do not enable overlapping MCPs without a clear reason
- do not make networked external tools mandatory for the base product

## Rejected alternatives

### MCP-heavy default setup

Rejected. It adds fragility and cognitive load before the user has proven they need it.

### No tooling policy at all

Rejected. Without a policy, tool surfaces bloat immediately.

### Browser MCP as the default verification path

Rejected. Playwright CLI is the better default for coding-agent workflows.

## Consequences

Positive:

- small and reliable default setup
- lower security and poisoning risk
- easier to reason about what the product actually depends on

Negative:

- power users may need to opt into tools they expected by default
- some integrations remain a manual step until better packaging exists
