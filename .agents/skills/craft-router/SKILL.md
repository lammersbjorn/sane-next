---
name: craft-router
description: Use this skill when user-facing frontend, rendered UI review, accessibility, documentation, or UX-copy work needs routing to one focused Sane craft skill. Route backend-only, packaging-only, typo-only, and file-path-only matches to the main session.
license: MIT
compatibility: Pi, Codex
---

# Craft Router

## Goal

Classify the user-facing craft surface and load the smallest matching subordinate skill.

## Use When

- frontend UI, components, pages, layout, styling, design systems, or visual polish are the work
- screenshots, responsive behavior, design fidelity, or rendered UI review is requested
- accessibility work touches forms, dialogs, navigation, custom controls, keyboard/focus, labels, semantics, or contrast
- README, guides, reference docs, changelog, release notes, migration notes, support docs, or docs review is requested
- UX copy covers labels, empty states, errors, onboarding, tooltips, or product microcopy

## Use Main Session When

- the task is backend-only, CI-only, packaging-only, release-mechanics-only, or security-only
- the change is a mechanical version bump, import cleanup, formatter pass, or typo-only edit
- a README or UI path is mentioned but the deliverable is code, config, or file movement rather than docs or user-facing UI

## Inputs

- user request and changed-file intent
- available repo scripts and tests
- any explicit screenshot, design, accessibility, or docs requirements

## Outputs

- exact skill id(s): `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, `ux-copy`, or `none`
- one selected subordinate skill by default
- two selected skills only for real boundary work, such as frontend plus accessibility, frontend plus UX copy, or feature plus docs
- a short reason for the dispatch when ambiguity matters

## How To Run

1. Return exact Sane skill ids: `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, `ux-copy`, or `none`.
2. Choose `frontend-craft` for implementing or polishing UI.
3. Choose `frontend-review` for reviewing rendered UI quality, design fidelity, responsive behavior, or visual regressions.
4. Choose `frontend-accessibility` for accessibility audits or fixes.
5. Choose `docs-writing` for source-grounded docs, changelogs, release notes, guides, or docs review.
6. Choose `ux-copy` for product microcopy and user-facing strings.
7. Choose `none` for backend-only, packaging-only, typo-only, or false-positive path matches.
8. If two surfaces are genuinely inseparable, load exactly two and state the boundary.

## Verification

- every selected value is exactly one of: `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, `ux-copy`, `none`
- selected skill matches the actual deliverable, not just filenames
- this router only dispatches; subordinate skills own craft doctrine
- one task selects one skill by default and two skills at most

## Gotchas / Safety

- keep this skill dispatch-only
- put MCP, package, browser, design-system, docs-style, and copywriting doctrine in the subordinate skill that owns it
- choose `none` for false-positive craft matches

## Examples

### Positive

- "Fix the dialog focus trap" → `frontend-accessibility`.
- "Polish this pricing page from the screenshot" → `frontend-craft`, `frontend-review`.
- "Write migration notes for the new CLI flag" → `docs-writing`.
- "Rewrite the failed-import empty state" → `ux-copy`.
- "Fix API pagination" → `none`.

### Negative

- Output only exact skill ids, such as `frontend-accessibility`, instead of aliases like `ui`, `a11y`, `docs`, `copy`, or `api`.
- "Bump README version from 1.2 to 1.3" → `none` unless docs content changes are requested.
