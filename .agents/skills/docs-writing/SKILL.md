---
name: docs-writing
description: Use this skill when writing or reviewing source-grounded README content, guides, reference docs, changelogs, release notes, migration notes, support docs, or docs tests. Do not use for code-only or typo-only changes.
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

## Don't Use When

- the task is code-only and docs are not requested or affected
- the change is typo-only with no content decision
- marketing copy or product microcopy is the main deliverable; use `ux-copy`

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
3. Put content in the existing docs structure; do not create random new markdown files.
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

- do not document planned behavior as existing behavior
- do not duplicate durable policy already owned by ADRs, standards, or AGENTS.md
- do not over-explain internals when the reader needs a user workflow
- preserve changelog and SemVer conventions when relevant

## Examples

### Positive

- Document a CLI flag by checking `--help` and adding one minimal example.
- Write migration notes with before/after behavior and known incompatibilities.

### Negative

- Add a new roadmap file for a temporary plan.
- Claim an export target exists without checking config or generated artifacts.
