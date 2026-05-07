# ADR 0004: Use a Pi overlay with Codex skill export

## Status

Accepted

## Context

Deeper research changed the architecture picture.

Current Sane proved that the vision is good but the implementation became too large: too many packages, too many tracking surfaces, too much TUI and export machinery, too much prompt/routing complexity, and too many assumptions baked into the framework.

Research on Pi, Pi skills, Anthropic workflow guidance, and Ben Davis's public setup all point to a tighter direction:

- Pi already provides a strong runtime, extension API, plugin system, skills, commands, MCP loading, and performant core.
- Pi's extension points are strong enough for a distribution/overlay model.
- Pi's internals are not a good deep-fork target: the settings singleton, TUI internals, and Rust native layer create real coupling cost.
- Skills are the cross-tool substrate. The same pack content can be exported to both Pi and Codex-native skill paths.
- Ben Davis's current public setup is Pi-first with `openai-codex`, `gpt-5.5`, and low reasoning as the default.
- The best part of current Sane to preserve is not its control surface; it is its workflow discipline, skill-first guidance, and pack idea.

## Decision

Build `sane-next` as a **Pi-first overlay/distribution**, not as a fresh runtime and not as a deep fork.

### Product boundary

- **Primary runtime:** Pi
- **Primary integration shape:** Pi plugin(s) plus skill/command/config distribution
- **Secondary export target:** Codex-native skill paths
- **Standalone Sane surface:** a small companion CLI for install, export, update, and doctor flows only

### Pack model

- Packs are **skill-first bundles**
- One pack should stay close to one concern
- Shared pack content should be authored once and exported to:
  - Pi skills paths
  - Codex-native skills paths
- Pi-only behavior belongs in the Pi plugin layer, not in the shared skill format

### Config model

- Keep config small and explicit
- Prefer a unified Sane config schema that controls:
  - enabled packs
  - default reasoning/model preferences
  - export targets
  - minimal companion-CLI behavior
- Do not build a TUI-heavy settings product

### Runtime policy

- Pi remains the runtime that owns the agent loop, TUI, MCP bridge, and native layer
- Sane should not try to replace Pi's runtime internals in v1
- If deeper loop changes are later required, prefer a thin fork only after the overlay path is proven insufficient

## Rejected alternatives

### Full fresh runtime

Rejected. Pi already solves too much of the hard runtime problem to justify rebuilding it now.

### Deep Pi fork

Rejected. The maintenance burden is too high and the Rust/native/runtime internals would become Sane's problem immediately.

### Codex-only primary product

Rejected as the main direction. Research and your clarified goal point toward Pi as the most interesting optimization substrate, while still keeping Codex-native skill export as an important secondary target.

### TUI-heavy standalone app

Rejected. Research favors a small config-first system with a narrow companion CLI.

## Consequences

Positive:

- faster path to a powerful result
- keeps Sane focused on workflow quality, packs, and defaults
- reuses Pi where Pi is already strong
- preserves cross-tool value through shared skills

Negative:

- Sane now depends on Pi's extension stability and release cadence
- some behaviors remain constrained by Pi's runtime model
- Pi-specific and Codex-specific layers must stay clearly separated to avoid drift

## Notes

This does not mean "Sane becomes just a theme pack for Pi." The product is a curated Pi overlay with opinionated packs, strong defaults, workflow discipline, and a secondary Codex-native skill export path. If the overlay model later proves too restrictive, that should trigger a new ADR comparing a thin fork against continuing the overlay.
