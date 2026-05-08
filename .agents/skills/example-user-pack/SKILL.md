---
name: example-user-pack
description: "Example disabled user pack fixture used to verify sane-next user-pack discovery and export behavior."
license: MIT
compatibility: Pi, Codex
---

# Example User Pack

## Goal

Provide a small fixture pack for testing user-added pack discovery without enabling it by default.

## Use When

- verifying user pack export behavior
- testing config-driven enablement

## Use Real Workflow Packs When

- doing real project work

## Inputs

- Sane config with this pack enabled

## Outputs

- exported fixture skill artifact

## How To Run

1. Enable this pack in a fixture Sane config.
2. Run `sane-next export` against the fixture target.
3. Inspect the exported skill path.

## Verification

- the exported target includes `example-user-pack/SKILL.md` only when enabled

## Gotchas / Safety

- Treat this as a fixture pack for discovery and export tests.

## Examples

### Positive

- A test enables this pack and confirms the export appears.
