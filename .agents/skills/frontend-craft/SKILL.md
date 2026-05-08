---
name: frontend-craft
description: Use this skill when implementing or polishing user-facing frontend UI, components, layout, styling, responsive behavior, visual hierarchy, or interaction details. Use frontend-review for review-only work and backend skills for backend-only work.
license: MIT
compatibility: Pi, Codex
---

# Frontend Craft

## Goal

Ship frontend changes that are usable, coherent with the existing product, and locally verified.

## Use When

- building or changing components, pages, layouts, styling, design-system usage, or interaction states
- improving visual hierarchy, spacing, responsive behavior, motion, or perceived quality
- translating a screenshot, design note, or product requirement into code

## Route Elsewhere When

- use `frontend-review` for review-only UI feedback
- use `frontend-accessibility` when accessibility remediation is primary
- use an implementation skill for work with no user-facing UI surface

## Inputs

- existing UI patterns and nearby components
- product requirement, screenshot, or design description
- repo build, lint, test, and preview commands

## Outputs

- minimal UI code changes that fit current architecture
- updated states for loading, empty, error, disabled, hover/focus, and small/large viewport cases when affected
- concise notes on verification and any unverified visual assumptions

## How To Run

1. Inspect nearby UI before inventing new patterns.
2. Preserve the app's current design language unless the request explicitly changes it.
3. Implement the smallest coherent change across structure, styling, state, and copy placeholders.
4. Check edge states that the changed surface creates or exposes.
5. Prefer repo-native CSS/design tokens and components over one-off styling.
6. Use Playwright CLI or the project's preview/test workflow when browser verification is available.

## Verification

- run the closest repo lint/type/test/build command that covers the changed UI
- if available, inspect the rendered result with Playwright CLI or the project's browser workflow
- verify at least one narrow and one wide viewport for responsive layout changes
- report any visual checks that could not be run

## Gotchas / Safety

- add new UI dependencies only when the repo already expects them or the user approves them
- use decorative detail to support the user's task
- keep data, controls, focus outlines, and error states visible and understandable
- keep generated screenshots or artifacts out of source unless the repo already tracks them

## Examples

### Positive

- Reuse the existing card component and token spacing for a new dashboard panel.
- Add explicit empty and error states when a list layout changes.

### Negative

- Rewrite the design system for a single button request.
- Claim visual quality without rendering or naming the missing verification.
