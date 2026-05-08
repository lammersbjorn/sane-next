---
name: frontend-review
description: Use this skill when reviewing rendered UI for design fidelity, visual polish, responsive behavior, screenshots, regressions, or frontend implementation quality. Do not use for implementation-only work unless review is requested.
license: MIT
compatibility: Pi, Codex
---

# Frontend Review

## Goal

Find concrete rendered-UI issues and recommend small fixes ranked by user impact.

## Use When

- reviewing screenshots, preview builds, diffs, or UI changes
- checking design fidelity, layout, spacing, hierarchy, responsive behavior, or visual regressions
- giving pre-merge frontend quality feedback

## Don't Use When

- the task is to implement a UI from scratch; use `frontend-craft`
- the review is primarily accessibility compliance; use `frontend-accessibility`
- there is no rendered or renderable UI surface

## Inputs

- screenshots, routes, components, or diffs under review
- design/reference material if provided
- available preview, test, screenshot, and Playwright CLI commands

## Outputs

- prioritized findings with exact location, impact, and suggested fix
- explicit pass notes for important areas checked with no issue
- verification commands or browser steps used

## How To Run

1. Identify the intended user task for the screen.
2. Review hierarchy, alignment, spacing, density, contrast, affordance, state coverage, and responsive behavior.
3. Compare against existing product patterns before suggesting new ones.
4. Prefer specific fix guidance over subjective taste statements.
5. Separate blockers from polish.
6. If possible, verify by rendering the screen locally rather than reading code only.

## Verification

- inspect the rendered UI when tooling is available
- run relevant screenshot, component, e2e, lint, type, or build checks
- include viewport(s), route(s), and fixture data used
- state if review was code-only or screenshot-only

## Gotchas / Safety

- do not invent design requirements not present in the product or request
- do not over-prioritize tiny polish over broken flows
- do not ask for wholesale rewrites when a targeted fix works
- treat screenshots as evidence, not as complete app state

## Examples

### Positive

- "The primary action loses hierarchy below 640px; stack it above the secondary action."
- "No issue found in the empty state at desktop and mobile widths."

### Negative

- "Make it pop" without a concrete location or fix.
- Blocking merge on a harmless difference from a non-binding mockup.
