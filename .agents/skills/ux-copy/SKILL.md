---
name: ux-copy
description: Use this skill when writing or reviewing user-facing product microcopy such as labels, buttons, empty states, errors, onboarding, tooltips, confirmations, settings text, or in-app help. Do not use for long-form docs.
license: MIT
compatibility: Pi, Codex
---

# UX Copy

## Goal

Make product text clear, actionable, consistent, and kind at the exact moment the user needs it.

## Use When

- labels, buttons, menus, empty states, errors, confirmations, onboarding, tooltips, or settings text change
- a flow needs clearer next steps, state explanation, or recovery guidance
- user-facing strings need review for tone, consistency, or localization risk

## Don't Use When

- the task is long-form README, guide, changelog, or API reference work; use `docs-writing`
- the task is visual layout or component implementation only
- copy is internal logs, test fixtures, or developer-only comments

## Inputs

- target user, context, and user goal
- current UI state and surrounding strings
- product terminology and localization constraints

## Outputs

- concise replacement copy with location and intended state
- alternatives only when tradeoffs are meaningful
- notes for validation, localization, or product/legal review when needed

## How To Run

1. Identify what the user is trying to do and what state the product is in.
2. Prefer specific nouns and verbs over vague encouragement.
3. Make errors explain what happened, why it matters, and what to do next when known.
4. Make empty states orient the user and offer a useful next action.
5. Keep button and label text consistent with nearby product terminology.
6. Remove blame, jokes in stressful states, and unnecessary exclamation.

## Verification

- check surrounding UI for duplicate or conflicting terms
- verify strings fit likely component space and responsive layouts
- check placeholders, pluralization, and variable interpolation when relevant
- run local tests or snapshots if strings are covered

## Gotchas / Safety

- do not promise capabilities the product does not have
- do not hide destructive consequences behind cute copy
- do not use color, icon, or position as the only explanation of state
- flag legal, billing, privacy, or safety-sensitive copy for human review

## Examples

### Positive

- Empty state: "No projects yet. Create a project to invite teammates and track work."
- Error: "We couldn't save changes. Check your connection and try again."

### Negative

- "Oopsie! Something went wrong!!!"
- Button text that says "OK" for a destructive confirmation.
