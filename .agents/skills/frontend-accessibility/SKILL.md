---
name: frontend-accessibility
description: Use this skill when auditing or fixing frontend accessibility for forms, dialogs, navigation, custom controls, keyboard/focus behavior, labels, semantics, ARIA, or contrast. Use frontend-craft for generic visual polish.
license: MIT
compatibility: Pi, Codex
---

# Frontend Accessibility

## Goal

Make affected UI perceivable, operable, understandable, and robust without adding unnecessary complexity.

## Use When

- forms, dialogs, navigation, menus, tabs, custom controls, or focus behavior change
- labels, descriptions, errors, headings, landmarks, semantics, ARIA, or contrast are in scope
- an accessibility audit, regression, or bug is requested

## Route Elsewhere When

- use `frontend-craft` for visual styling with no accessibility surface
- use `ux-copy` for product copy only
- use a dedicated compliance review process for broad certification beyond local code review

## Inputs

- affected routes/components and user flows
- existing design-system accessibility patterns
- repo tests plus browser/Playwright CLI availability

## Outputs

- targeted accessibility fixes or findings
- keyboard path, focus order, semantic structure, name/role/value, and contrast notes for changed surfaces
- verification performed and remaining manual checks

## How To Run

1. Start with native HTML semantics before ARIA.
2. Add ARIA roles only after native elements and expected keyboard behavior are clear.
3. Preserve visible focus and logical keyboard order.
4. Ensure controls have accessible names and state where needed.
5. Connect form labels, helper text, validation errors, and summaries.
6. For dialogs/menus/popovers, check focus entry, containment if appropriate, Escape/close behavior, and return focus.
7. Use automated checks as a floor, then perform targeted manual keyboard review when possible.

## Verification

- run local lint/type/test/build checks that cover the component
- use Playwright CLI, browser devtools, axe integration, or repo-native accessibility tests when available
- manually verify keyboard navigation and focus for changed interactive flows when possible
- report unverified assistive-technology behavior honestly

## Gotchas / Safety

- ARIA does not fix incorrect interaction behavior
- keep focus indicators visible and easy to see
- pair color with text, icon shape, or another cue for state and errors
- treat automated checks as evidence, not as full WCAG compliance proof

## Examples

### Positive

- Add `aria-describedby` from an input to its error text and verify keyboard submission.
- Replace a clickable `div` with a native `button` when it acts like a button.

### Negative

- Add broad ARIA roles without testing keyboard behavior.
- Mark decorative content as important to screen readers.
