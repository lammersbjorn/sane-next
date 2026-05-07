# ADR 0009: Use annotated SemVer and Linux-only CI while private

## Status

Accepted

## Context

The repo needs a release/versioning policy before implementation starts, but it is still private and pre-stable.

The goal is to keep release discipline from the start without paying for premature automation or expensive runner choices.

Research favored:

- Conventional Commits
- annotated tags
- pre-stable `v0.y.z`
- Linux-only CI while private
- delayed release automation until it clearly pays for itself

## Decision

Use the following release policy:

1. **Versioning**
   - Use SemVer.
   - While pre-stable, use `v0.y.z`.
   - Use **annotated** git tags only.

2. **CI while private**
   - Keep CI Linux-only while the repo is private.
   - Do not add macOS or Windows runners yet.

3. **Release automation**
   - Keep release automation minimal at first.
   - Do not add full binary release automation until the companion CLI is real enough to ship.
   - Revisit **GoReleaser** once external binary distribution becomes a real need.

## Rejected alternatives

### Cross-platform CI from day one

Rejected. It costs more and does not match current product maturity.

### Unstructured tags or no tags until later

Rejected. That throws away release discipline exactly when it is easiest to establish it.

### Full release automation immediately

Rejected. It is premature before there is a real shipped artifact set.

## Consequences

Positive:

- release history stays clean early
- cost stays low while the repo is private
- the path to later automation stays open

Negative:

- non-Linux issues may be found later
- early releases remain slightly more manual
