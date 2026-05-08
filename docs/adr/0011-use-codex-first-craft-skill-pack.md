# ADR 0011: Use a Codex-first craft skill pack with a small router

## Status

Accepted

## Context

`sane-next` currently exports compact workflow packs, but it does not include the frontend and documentation craft packs that existed in old Sane planning and prior worktrees.

The target runtime is **GPT-5.5 in Codex and Pi**. The pack must therefore follow Codex Agent Skills behavior and Pi export constraints, not Claude-specific runtime assumptions. Current research points to a consistent pattern:

- Codex skills should use concise `description` metadata for discovery and load full instructions only after selection.
- GPT-5.5 / GPT-5-Codex guidance favors short tactical constraints over large prompt handbooks.
- Always-on context files can increase cost and reduce success when they become broad or stale.
- Frontend craft sources such as OpenAI's curated frontend skill, Taste Skill, Impeccable, Make Interfaces Feel Better, and Microsoft frontend review skills are useful upstream inspiration, but wholesale vendoring would add license, maintenance, and prompt-surface risk.
- Documentation craft has no single dominant upstream skill; durable guidance comes from Codex skills, Diátaxis, Google and Microsoft writing guides, Keep a Changelog, SemVer, and executable docs checks.

## Decision

Add a **Codex-first craft pack** made of one small router skill and narrow subordinate skills:

- `craft-router`
- `frontend-craft`
- `frontend-review`
- `frontend-accessibility`
- `docs-writing`
- `ux-copy`

The router is a classifier and dispatcher only. It must not contain detailed craft doctrine. It should load one subordinate skill by default and load two only for real boundary tasks, such as frontend plus accessibility, frontend plus UX copy, or feature work plus docs.

Subordinate skills must be original Sane-authored skills optimized for GPT-5.5/Codex/Pi:

- short bodies following `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`
- progressive disclosure through optional `references/` files
- source/provenance notes instead of large copied upstream text
- verification ladders that prefer local repo checks and Playwright CLI when browser verification is available

The pack is enabled through `pi-plugin/config-schema.toml` as normal Sane packs targeting both Pi and Codex. It must not add new default Pi packages, MCPs, or runtime bridges.

## Invocation rules

`craft-router` should trigger for explicit work involving:

- frontend UI, components, pages, layout, styling, design systems, or visual polish
- screenshots, design fidelity, responsive behavior, or rendered UI review
- accessibility work on forms, dialogs, navigation, custom controls, keyboard/focus, labels, semantics, or contrast
- README, guides, reference docs, changelog, release notes, migration notes, support docs, or source-grounded docs review
- UX copy such as labels, empty states, errors, onboarding, tooltips, and product microcopy

It should not trigger for:

- backend-only work
- CI, packaging, release mechanics, or security tasks unless docs/UI are the actual deliverable
- mechanical version bumps, import cleanup, formatting, or typo-only edits
- generic refactors with no user-facing or documentation surface
- file-path false positives, such as changing a README version string only

## Upstream policy

Use upstream material as inspiration in this priority order:

1. OpenAI/Codex Agent Skills and OpenAI curated frontend skill
2. OpenAI GPT-5/GPT-5-Codex frontend and prompting guidance
3. Taste Skill for anti-slop frontend direction
4. Impeccable for visual/design review vocabulary and anti-patterns
5. Make Interfaces Feel Better for micro-polish details
6. Microsoft frontend-design-review and accessibility guidance for review posture
7. Diátaxis, Google developer docs style, Microsoft Writing Style Guide, Keep a Changelog, and SemVer for docs-writing behavior

Do not copy large upstream text into Sane. If vendoring or direct adaptation is later desired, first verify license/provenance and record the reason.

## Rejected alternatives

### One monolithic frontend/docs skill

Rejected. It would increase token overhead, blur triggers, and conflict with the instruction-surface standard.

### Always-on AGENTS.md craft guidance

Rejected. `AGENTS.md` must stay small and durable; task-specific craft belongs in skills.

### Claude-first skill compatibility

Rejected. Claude ecosystem material may be useful prior art, but Sane Next's product target is Codex/Pi/GPT-5.5.

### New MCP or package dependencies for craft packs

Rejected. ADR 0008 and ADR 0010 keep tooling small and curated. Browser/Figma/MCP integrations may be documented as optional external workflows, not required defaults.

## Consequences

Positive:

- restores the missing frontend/docs craft scope without bloating always-on instructions
- gives GPT-5.5 targeted constraints where it benefits most
- keeps Pi/Codex export simple and testable
- preserves future upstream collaboration options

Negative:

- skill descriptions need careful tuning to avoid false positives
- upstream source freshness and licenses require periodic review
- frontend quality still depends on available project verification tools and human visual judgment
