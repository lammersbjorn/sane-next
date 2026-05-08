---
name: craft-router
description: Use this skill when user-facing frontend, rendered UI review, accessibility, documentation, or UX-copy work needs routing to one focused Sane craft skill. Do not use for backend-only, packaging-only, typo-only, or file-path-only matches.
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

## Don't Use When

- the task is backend-only, CI-only, packaging-only, release-mechanics-only, or security-only
- the change is a mechanical version bump, import cleanup, formatter pass, or typo-only edit
- a README or UI path is mentioned but the deliverable is not docs or user-facing UI

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

1. Return exact Sane skill ids, not generic labels like `ui`, `docs`, `api`, or `copywriting`.
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
- no craft doctrine is applied directly from this router
- no more than two subordinate skills are selected for one task

## Gotchas / Safety

- keep this skill dispatch-only
- do not add MCP, package, browser, design-system, docs-style, or copywriting doctrine here
- prefer no craft skill over a false-positive dispatch

## Examples

### Positive

- "Fix the dialog focus trap" → `frontend-accessibility`.
- "Polish this pricing page from the screenshot" → `frontend-craft`, `frontend-review`.
- "Write migration notes for the new CLI flag" → `docs-writing`.
- "Rewrite the failed-import empty state" → `ux-copy`.
- "Fix API pagination" → `none`.

### Negative

- Do not output `ui`, `a11y`, `docs`, `copy`, `api`, or other aliases.
- "Bump README version from 1.2 to 1.3" → `none` unless docs content changes are requested.
