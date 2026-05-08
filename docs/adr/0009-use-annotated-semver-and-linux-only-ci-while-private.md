# ADR 0009: Use annotated SemVer; CI scope later superseded

## Status

Accepted; CI operating-system scope superseded by [ADR 0014](0014-add-cross-platform-compatibility-checks.md)

## Context

At the time of this decision, the repo needed a release/versioning policy before implementation started and was still private and pre-stable.

The goal was to keep release discipline from the start without paying for premature automation or expensive runner choices.

Research favored:

- Conventional Commits
- annotated tags
- pre-stable `v0.y.z`
- Linux-only CI while private, later superseded by ADR 0014
- delayed release automation until it clearly pays for itself

## Decision

Use the following release policy:

1. **Versioning**
   - Use SemVer.
   - While pre-stable, use `v0.y.z`.
   - Use **annotated** git tags only.

2. **CI operating-system scope**
   - Superseded by ADR 0014 after the project adopted Linux, macOS, and Windows compatibility checks.
   - Historical decision: keep CI Linux-only while the repo was private and do not add macOS or Windows runners yet.

3. **Release automation**
   - Keep release automation minimal until it pays for itself.
   - Do not add full release automation just because binaries exist.
   - Revisit **GoReleaser** or similar tooling only when repeatable artifact publishing becomes a real need.

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
- CI cost stayed low while the repo was private
- the path to later automation stays open

Negative:

- superseded CI scope means readers must follow ADR 0014 for current platform policy
- early releases remain slightly more manual
