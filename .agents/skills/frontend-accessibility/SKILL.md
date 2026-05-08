---
name: frontend-accessibility
description: Use this skill when auditing or fixing frontend accessibility for forms, dialogs, navigation, custom controls, keyboard/focus behavior, labels, semantics, ARIA, or contrast. Do not use for generic visual polish.
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

## Don't Use When

- the task is only visual styling with no accessibility surface
- the task is product copy only; use `ux-copy`
- the task asks for broad compliance certification beyond local code review

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
2. Do not add ARIA roles until native elements and expected keyboard behavior are clear.
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
- do not remove focus indicators for aesthetics
- do not rely on color alone for state or errors
- avoid claiming WCAG compliance from automated checks alone

## Examples

### Positive

- Add `aria-describedby` from an input to its error text and verify keyboard submission.
- Replace a clickable `div` with a native `button` when it acts like a button.

### Negative

- Add broad ARIA roles without testing keyboard behavior.
- Mark decorative content as important to screen readers.
