# ADR 0005: Use native git hooks and Conventional Commits

## Status

Accepted

## Context

`sane-next` needs a strict, low-bloat contribution protocol from the start so the repo stays clean while the product is still in heavy research and early implementation.

The repo is intentionally small and pre-implementation. This is the right time to lock lightweight discipline around commits and local validation before habits drift.

## Decision

Use a native repo-hook setup plus Conventional Commits:

1. Commit hooks live in the repository under `.githooks/`.
2. Local clones should set `core.hooksPath` to `.githooks`.
3. Commit messages must follow Conventional Commits.
4. The repo allows a `research:` type because early project work includes substantial evidence gathering before implementation.

## Consequences

Positive:

- commit history stays structured from day one
- changelog and release tooling can be automated later without rewriting history
- repo hygiene checks run before commits without adding package-manager tooling

Negative:

- contributors must opt into the committed hooks locally
- the hook checks must stay small and boring to avoid becoming a workflow tax

## Notes

This ADR does not yet define the full release automation policy. That remains a separate decision after the current release/versioning research is synthesized.
