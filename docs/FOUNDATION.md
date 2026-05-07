# Foundation

This file holds the durable product and repo decisions for `sane-next`.

It is meant to replace the earlier ADR sprawl without losing the decision trail. If a future implementation agent needs the "why", it should be able to get it here without chat history.

## Research inputs

This repo direction came from combining:

- the postmortem of the current `sane` repo as a reference / proof of concept
- Pi architecture and extension-surface research
- Pi skills ecosystem research
- public workflow/prompt research around Anthropic agent guidance and ETH-style context findings
- public workflow takes from Theo, Ben Davis, Nerdsnipe, and adjacent agentic-coding research
- release/tooling/runtime research for a small companion CLI and low-bloat repo setup

`sane` is **reference material only**. It is not a source of truth for claims about the new project.

## Product direction

`sane-next` is a **Pi-first overlay/distribution**.

- **Primary runtime:** Pi
- **Primary integration shape:** Pi plugin(s) plus skill/command/config distribution
- **Secondary export target:** Codex-native skill paths
- **Standalone Sane surface:** a small companion CLI for install, export, update, and doctor flows only

`sane-next` is **not**:

- a fresh runtime
- a deep Pi fork
- a TUI-heavy standalone product

## Why this direction

Earlier research pointed to a tighter direction than old Sane:

- Pi already provides a strong runtime, extension API, plugin system, skills, commands, MCP loading, and performant core.
- Pi's extension points are strong enough for an overlay/distribution model.
- Pi's internals are a poor deep-fork target because the coupling cost is too high.
- Skills are the cross-tool substrate, so shared pack content can export to both Pi and Codex-native paths.
- The most valuable thing to preserve from old Sane is workflow discipline, skill-first guidance, and the pack idea.

## Old Sane reference analysis

The current `sane` repo was treated as a postmortem source, not as a base to copy.

### What was worth preserving

- outcome-first workflow discipline
- minimal always-on guidance
- skill-first extensibility
- explicit verification and recovery
- RTK-style task-aware tool routing when a target repo uses RTK

### What became bloated or drift-prone

- package-per-concern architecture
- TUI-heavy product direction
- overlapping planning files and repo-local tracking surfaces
- prompt/policy duplication across many files
- premature model-routing complexity
- deep runtime ownership where the host runtime already solved the problem

### Transfer candidates

- the **pack** concept, but kept small and skill-first
- RTK-style routing patterns when the target repo already uses RTK
- strong acceptance/verification discipline

### What not to import directly

- the old Sane codebase itself
- TUI architecture
- multi-surface planning/docs patterns
- complex package structure as a starting point

## Pack and config model

- Packs are **skill-first bundles**.
- One pack should stay close to one concern.
- Shared pack content should be authored once and exported to:
  - Pi skills paths
  - Codex-native skills paths
- Pi-only behavior belongs in the Pi plugin layer, not in the shared skill format.

Config rules:

- keep config small and explicit
- use one unified Sane config schema
- let config control enabled packs, default model/reasoning preferences, export targets, and minimal CLI behavior
- do not build a TUI-heavy settings product

## Tooling and MCP policy

Use a small three-tier policy:

1. **Default core**
   - no third-party MCP is required by default
   - the base product must work with host runtime + local repo tools + standard shell/file/git capabilities

2. **Recommended helpers**
   - RTK when the target repo uses RTK
   - Playwright CLI over a browser MCP for frontend verification
   - Context7 and grep.app only when external docs or broad code search are genuinely needed

3. **Optional curated integrations**
   - explicit and opt-in only
   - documented with provenance and purpose
   - Sane configures/document integrations; it does not rebuild the runtime bridge

## Companion CLI

Use **Go** for the companion CLI.

Rules:

- keep it small and command-focused
- prefer the standard library unless a dependency clearly pays for itself
- start with `install`, then add `export`, `update`, and `doctor` only when justified
- do not build a TUI-heavy CLI surface

## Release and repo hygiene

- use Conventional Commits
- keep committed hooks in `.githooks/`
- use **annotated** tags only
- use `v0.y.z` while pre-stable
- keep CI Linux-only while the repo is private
- delay full binary release automation until there is a real artifact set to ship

## Tracking decision

The project does **not** use GitHub Issues as the canonical tracker.

Instead:

- `TRACK.toml` is the canonical active-window task ledger
- commits and pull requests are the progress evidence layer
- broader decision/history material belongs in durable repo docs, not in `TRACK.toml`

Why:

- the canonical tracker needs to live in the repo
- markdown-only planning tended to sprawl
- GitHub Issues as the source of truth was explicitly rejected
- the current-state ledger must stay small, resumable, and machine-usable

## Decision register

These are the durable research-backed decisions that should guide implementation unless explicitly changed.

### 1. Architecture

- use a **Pi-first overlay/distribution**
- export the same shared pack content to Codex-native skill paths
- do not build a fresh runtime
- do not deep-fork Pi

### 2. Product shape

- keep a small config-first product
- do not build a TUI-heavy standalone settings surface
- keep the standalone Sane surface to a small companion CLI

### 3. Tracking

- keep `TRACK.toml` as a bounded active window
- keep broader reasoning out of TRACK
- keep repo-local truth small and explicit

### 4. Instruction surfaces

- keep always-on context tiny
- use progressive disclosure
- avoid duplicated policy across repo surfaces
- keep skills as the reusable trigger-loaded unit

### 5. Tooling and MCPs

- no third-party MCP is required by default
- keep integrations small, explicit, and opt-in
- prefer Playwright CLI over a browser MCP
- recommend RTK only when the target repo already uses RTK

### 6. CLI language

- use **Go** for the companion CLI

### 7. Release policy

- use Conventional Commits
- use annotated tags only
- use `v0.y.z` while pre-stable
- keep CI Linux-only while the repo is private

## Rejected alternatives

- GitHub Issues as the primary tracker
- markdown-only tracking as the primary system
- a fresh Sane runtime
- a deep Pi fork
- a Codex-only primary product
- a TUI-heavy standalone app
- MCP-heavy default setup
- early cross-platform CI and heavy release automation

## Open hypotheses and cautions

These are still important, but they are **not** locked facts:

- `gpt-5.5` with low reasoning is likely a strong default, but it should stay configurable rather than hard-coded
- model routing may be less useful than one strong default, but this should be proven with real use
- Pi plugin API stability may force adjustments later
- export-path details may change once real Pi/Codex paths are tested

## Repo truth ownership

Read in this order:

1. `TRACK.toml`
2. `docs/FOUNDATION.md`
3. `docs/STANDARDS.md`
4. `docs/IMPLEMENTATION.md`
5. `AGENTS.md`

Canonical ownership:

- `TRACK.toml` — active slice only
- `docs/FOUNDATION.md` — durable product/repo decisions
- `docs/STANDARDS.md` — durable tracking/prompt/skill/repo standards
- `docs/IMPLEMENTATION.md` — exact build order and acceptance path
- `AGENTS.md` — tiny startup rules
- `CONTRIBUTING.md` — commit and hook discipline
