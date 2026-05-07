# Foundation

This file is the fixed product and repo truth for `sane-next`.

## Product

- `sane-next` is a **Pi-first overlay/distribution**.
- Pi is the primary runtime.
- Codex-native skill export is a secondary target.
- `sane-next` is **not** a fresh runtime and **not** a deep Pi fork.

## Product boundaries

- Keep the system small.
- Prefer a small config-first product over a TUI-heavy surface.
- Preserve packs and extensibility, but keep them skill-first and low-bloat.
- Use a small companion CLI only for install/export/update/doctor flows.

## What to preserve from old Sane

- outcome-first workflow discipline
- minimal always-on guidance
- skill-first extensibility
- explicit verification and recovery
- RTK-style task-aware tool routing when a target repo uses RTK

## What not to carry over

- package-per-concern architecture
- TUI-heavy product direction
- overlapping planning files
- prompt/policy duplication across many files
- premature model-routing complexity

## Repo truth

Read in this order:

1. `TRACK.toml`
2. `docs/FOUNDATION.md`
3. `docs/IMPLEMENTATION.md`
4. `AGENTS.md`

Canonical ownership:

- `TRACK.toml` — active slice only
- `docs/FOUNDATION.md` — durable product/repo decisions
- `docs/IMPLEMENTATION.md` — exact build order and acceptance path
- `AGENTS.md` — tiny startup rules
- `CONTRIBUTING.md` — commit and hook discipline

## TRACK rules

`TRACK.toml` is a bounded active window, not a backlog.

It must:

- stay short
- hold only active or near-term work
- contain structured work items with `why_now`, `write_scope`, `inputs`, and `done_when`
- avoid completed work, hypotheses, and research notes

## Implementation rules

- One rule should live in one place only.
- Always-on instruction surfaces stay tiny.
- Research before changing stack, Pi, MCP, or model assumptions.
- Prefer executable enforcement over repeated prose.
- Do not add extra markdown planning surfaces without intentionally changing the repo structure first.

## Tooling and runtime policy

- No third-party MCP is required by default.
- RTK is recommended when the target repo already uses RTK.
- Playwright CLI is preferred over a browser MCP for frontend verification.
- Context7 and grep.app are optional helpers, not mandatory dependencies.
- Keep additional integrations explicit and opt-in.

## Companion CLI

- Use **Go** for the companion CLI.
- Keep it small and command-focused.
- Prefer the standard library unless a dependency clearly pays for itself.

## Release and repo hygiene

- Use Conventional Commits.
- Keep committed hooks in `.githooks/`.
- Use **annotated** tags only.
- Use `v0.y.z` while pre-stable.
- Keep CI Linux-only while the repo is private.

## Skill and prompt rules

- `AGENTS.md` stays tiny.
- Skills do one job each.
- Skill details are trigger-loaded, not always-on.
- Each skill should declare goal, triggers, inputs, outputs, steps, verification, and gotchas.

## Current file structure

Only these markdown files should exist unless the structure is intentionally changed:

- `README.md`
- `AGENTS.md`
- `CONTRIBUTING.md`
- `.github/copilot-instructions.md`
- `docs/FOUNDATION.md`
- `docs/IMPLEMENTATION.md`
