# Architecture Decision Records

This directory contains durable decisions for `sane-next`.

## Conventions

- Filename format: `NNNN-short-kebab-title.md`.
- Status is explicit near the top, usually `Accepted` or `Superseded`.
- Accepted ADRs are historical records. Prefer a new ADR for changed decisions instead of rewriting old rationale.
- ADRs explain context, decision, rejected alternatives, and consequences; they do not track tasks.

## Current anchors

- `0002-use-track-toml-plus-adrs-for-tracking.md` introduced repo-local tracking.
- `0006-use-bounded-track-window-and-track-standard.md` tightened `TRACK.toml` into an active-window file and supersedes looser tracking language from ADR 0002.
- `0012-use-layered-pi-customization-profile.md` records the current layered Pi customization policy.

For docs placement rules, see `../standards/DOCS-STRUCTURE-STANDARD.md`.
