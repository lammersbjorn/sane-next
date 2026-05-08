# ADR 0013: Use GitHub Releases as the first distribution channel

## Status

Accepted

## Context

`sane-next` is pre-stable and was source-built when this decision was made. The companion CLI is implemented in Go, the repo already uses annotated pre-stable `v0.y.z` tags, and the release-readiness acceptance path verifies install, export, doctor, repair, update, and uninstall behavior.

The next release channel should make source-built users safer and more repeatable without adding a premature package-maintenance burden.

## Decision

Use **GitHub Releases** as the first distribution channel for pre-stable `sane-next` users.

Initial release expectations:

1. Publish releases from annotated `v0.y.z` tags.
2. Attach the source archive and built companion CLI artifact(s) that match README install instructions.
3. Keep the README honest about the repo being pre-stable and about the currently supported install channels.
4. Do not add Homebrew, npm, or broad release automation until there is evidence that GitHub Releases are insufficient.

## Rejected alternatives

### Homebrew first

Rejected for the first channel. Homebrew would improve macOS install ergonomics, but it adds formula or tap maintenance before the product has proven release cadence.

### npm first

Rejected for the first channel. The companion CLI is Go, so npm distribution would add wrapper and packaging complexity that does not fit the current binary-first CLI choice.

### Source-built only

Rejected as the release channel decision. The source-built flow remains supported, but a pre-stable project still benefits from tagged release artifacts and repeatable downloads.

## Consequences

Positive:

- aligns with the existing annotated tag policy
- keeps pre-stable release operations simple
- avoids premature package-manager maintenance
- leaves Homebrew, npm, and GoReleaser paths open for later

Negative:

- users still need to download or place binaries manually
- install UX is less polished than Homebrew or npm
- release artifact creation may remain manual until automation is justified
