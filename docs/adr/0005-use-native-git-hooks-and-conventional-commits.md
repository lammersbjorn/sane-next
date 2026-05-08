# ADR 0005: Use native git hooks and Conventional Commits

## Status

Accepted

## Context

`sane-next` needed a strict, low-bloat contribution protocol from the start so the repo would stay clean through research, implementation, and pre-stable source-built use.

When this ADR was accepted, the repo was intentionally small and pre-implementation. The same lightweight discipline still applies now that the CLI, packs, overlay, and acceptance flow exist.

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

Release/versioning policy is covered by ADR 0009, ADR 0013, and ADR 0014.
