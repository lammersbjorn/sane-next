---
name: docs-writing
description: Use this skill when writing or reviewing source-grounded README content, guides, reference docs, changelogs, release notes, migration notes, support docs, or docs tests. Route code-only or typo-only changes to direct implementation.
license: MIT
compatibility: Pi, Codex
---

# Docs Writing

## Goal

Produce accurate, source-grounded documentation that helps the intended reader complete the task.

## Use When

- adding or revising README, guides, reference docs, changelog, release notes, migration notes, or support docs
- reviewing docs for correctness against code or behavior
- documenting CLI/API/config behavior that can be verified locally

## Use Another Route When

- for code-only tasks where docs are unchanged, stay with implementation
- for typo-only edits with no content decision, make the direct local edit
- for marketing copy or product microcopy, use `ux-copy`

## Inputs

- source files, tests, schema, CLI help, examples, or product behavior being documented
- target reader and task
- existing docs structure and style

## Outputs

- concise docs in the repo's existing location and voice
- examples that compile/run or are clearly marked as illustrative
- verification notes linking docs claims to source or commands

## How To Run

1. Identify the reader's job: learn, decide, perform a task, or look up reference.
2. Inspect source behavior before writing claims.
3. Put content in the existing docs structure. Use `docs/README.md` and the docs structure standard to choose the owning location.
4. Prefer task steps, examples, constraints, and troubleshooting over vague descriptions.
5. Keep changelog/release notes user-facing and grouped by meaningful change.
6. Use exact command names only when verified from source, help output, or the user's prompt; otherwise name the behavior without inventing invocation syntax.
7. Remove stale or contradictory docs near the edit.

## Verification

- run the command, test, generated help, or schema inspection that proves changed claims when practical
- verify links and file paths touched by the change
- ensure examples match current names, flags, defaults, and outputs
- state any claims that remain unverified

## Gotchas / Safety

- Describe planned behavior as planned, and existing behavior as existing.
- Link to durable policy owned by ADRs, standards, or `AGENTS.md` instead of restating it.
- Match depth to the reader job: user workflows first, internals only when they help the task.
- preserve changelog and SemVer conventions when relevant

## Examples

### Positive

- Document a CLI flag by checking `--help` and adding one minimal example.
- Write migration notes with before/after behavior and known incompatibilities.

### Negative

- Add a new roadmap file for a temporary plan.
- Claim an export target exists without checking config or generated artifacts.
