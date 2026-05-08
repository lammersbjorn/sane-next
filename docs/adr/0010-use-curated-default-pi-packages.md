# ADR 0010: Use curated default Pi packages for high-value runtime gaps

## Status

Accepted

## Context

ADR 0008 kept the default tool surface small to avoid old-Sane-style sprawl. That remains correct for broad MCP/plugin bundles, but it left an important gap: Sane's `agent-lanes` pack described subagent discipline while Pi itself intentionally ships without built-in subagents.

Current Pi package search shows multiple third-party subagent packages. The best default fit for Sane's lane model is `pi-subagents` because it is a Pi package, is purpose-built for delegation, publishes recent versions, and exposes parallel/chained subagent behavior that matches the `agent-lanes` pack.

## Decision

Allow Sane to auto-install a **curated allowlist** of high-value Pi packages during the companion CLI `install` flow.

Initial curated defaults:

- `npm:pi-subagents@0.24.0`
  - purpose: runtime-backed Pi subagent delegation for Sane agent lanes
  - provenance: npm package `pi-subagents`, repository `github.com/nicobailon/pi-subagents`
  - verification: `PI_CODING_AGENT_DIR=<tmp> pi install npm:pi-subagents@0.24.0` succeeded and `pi list` showed the package
  - security note: third-party Pi packages run with full extension privileges and must stay pinned/reviewed
- `npm:pi-rewind@0.5.0`
  - purpose: git-backed checkpoint and rewind support for safer long-running Pi sessions
  - provenance: npm package `pi-rewind`, repository `github.com/arpagon/pi-rewind`
  - verification: `PI_CODING_AGENT_DIR=<tmp> pi install npm:pi-rewind@0.5.0` succeeded and `pi list` showed the package
  - selection note: preferred over `pi-rewind-hook@1.8.4` after both installed cleanly because `pi-rewind` documents a dedicated `/rewind` command, diff preview, redo stack, branch safety, safe restore filters, and auto-pruning
  - security note: creates git-backed checkpoints and must remain pinned/reviewed

Initial curated opt-in:

- `npm:pi-prompt-template-model@0.9.3`
  - purpose: optional prompt-template model, reasoning, skill, and subagent routing for power workflows
  - provenance: npm package `pi-prompt-template-model`, repository `github.com/nicobailon/pi-prompt-template-model`
  - verification: npm metadata confirms a Pi package manifest with extension and skills
  - selection note: not default because Sane already owns model/reasoning defaults and this adds mode-routing surface area

Rules:

1. Curated packages must be explicit in `pi-plugin/config-schema.toml`.
2. Packages must be pinned by version or immutable source reference.
3. Auto-install happens through Pi's package manager, not by vendoring third-party code into this repo.
4. Fixture installs do not mutate a user's global Pi config unless explicitly forced.
5. Users can opt out with the companion CLI install flag.
6. Additional default packages require a new ADR update or successor ADR with provenance and purpose.

## Rejected alternatives

### Keep subagents advisory only

Rejected. It makes `agent-lanes` weaker than the product promise and misses a high-value Pi extension point.

### Install broad plugin/MCP bundles

Rejected. That repeats old Sane's tool-surface sprawl and weakens trust.

### Vendor or fork a subagent implementation

Rejected. Sane is a Pi overlay; Pi packages should provide runtime extensions where possible.

## Consequences

Positive:

- Sane's default Pi install can actually use subagents instead of only describing lanes.
- The runtime gap is closed without deep-forking Pi.
- The default remains auditable because packages are pinned and listed in config.

Negative:

- Default install now depends on a third-party package and network/package-manager availability.
- The curated allowlist must be actively reviewed over time.
- Package behavior may change when the pinned version is intentionally updated.
