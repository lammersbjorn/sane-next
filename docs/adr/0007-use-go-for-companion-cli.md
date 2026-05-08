# ADR 0007: Use Go for the companion CLI

## Status

Accepted

## Context

`sane-next` needs a small companion CLI for lifecycle, export, package recommendation, and explicit configuration flows around a Pi-first overlay product.

The CLI is not the runtime. It should stay boring, fast to start, easy to distribute, and easy to keep small.

Past attempts around Rust and TypeScript did not produce the right balance for this product. The current research direction favored:

- small static binaries
- simple cross-platform distribution
- low maintenance overhead
- good enough ergonomics without a heavy framework

## Decision

Use **Go** for the companion CLI.

Scope of this decision:

- applies to the standalone Sane companion CLI only
- does not change the Pi runtime choice
- does not force Go for every future helper or test utility

Implementation rules:

1. Keep the CLI small and command-focused.
2. Prefer the standard library unless a dependency clearly pays for itself.
3. Keep command growth tied to overlay lifecycle, package recommendation, export, and explicit user configuration needs.
4. Do not build a TUI-heavy CLI surface.

## Rejected alternatives

### Rust

Rejected for the companion CLI. It can produce excellent binaries, but the ergonomics and maintenance cost are not justified for this narrow utility layer.

### TypeScript / Node.js

Rejected for the companion CLI. Runtime and packaging friction are a worse fit for a small boring utility that should be easy to ship as a binary.

### Shell-only scripts

Rejected as the primary interface. They are useful for small helpers, but they are not a strong long-term CLI foundation.

## Consequences

Positive:

- small binaries
- straightforward cross-platform distribution path
- simple implementation path for overlay lifecycle and export flows

Negative:

- another language in the repo
- some future contributors may prefer a different language for CLI work
