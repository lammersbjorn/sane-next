# sane-next Roadmap

This is the living product roadmap for `sane-next`. It is not the active task ledger and not a historical implementation diary.

- Active execution window: `TRACK.toml`
- Durable decisions: `docs/adr/`
- Docs and tracking rules: `docs/standards/`
- Completion evidence: tests, acceptance runs, commits, pull requests, and release tags

## Current state

The initial Pi-first overlay is implemented and verified: companion CLI lifecycle flows, shared packs, Codex export, Pi plugin integration, goal/ledger commands, curated default packages, optional package recommendations, explicit theme configuration, and acceptance coverage exist in source.

The repo is pre-stable and source-built. The next work should focus on release readiness, package publishing decisions, and keeping docs/tracking surfaces small enough for agents to use.

## Now

- [x] Clean up repo docs and tracking structure after the initial implementation run.
  Evidence: `TRACK.toml` has a valid active window, docs directories have short ownership READMEs, stale run-specific roadmap/protocol text is removed or rewritten, and `bash .githooks/pre-commit`, `cd cli && go test ./...`, `node --test pi-plugin/plugin.test.js`, and `cd cli && ./acceptance.sh` pass.

## Next

- [x] Decide the first distribution channel for source-built users.
  Evidence: `docs/adr/0013-use-github-releases-as-first-distribution-channel.md` selects GitHub Releases as the first pre-stable distribution channel.
- [x] Prepare a pre-stable release checklist.
  Evidence: `bash .githooks/pre-commit`, `cd cli && go test ./...`, `node --test pi-plugin/plugin.test.js`, and `cd cli && ./acceptance.sh` pass; top-level CLI help now lists the source-built install/update/repair/package flags used by README and acceptance.
- [x] Refresh package recommendation pins from live Pi package/npm state.
  Evidence: `npm view <package> version dist-tags --json` confirmed every configured recommendation is already at the current `latest` dist-tag, and `cd cli && go test ./...` plus dry-run `sane-next package install` checks passed for every package ID.

## Later

- [ ] Add external mirrors only if there is evidence repo-local tracking is not enough.
  Evidence: a new ADR defines source of truth, sync direction, and failure mode.
- [ ] Revisit Codex-native export after real user feedback.
  Evidence: exported skill format still matches current Codex expectations and user-owned directories remain protected.
- [ ] Consider publishing or automating release channels.
  Evidence: release automation preserves Sane-owned/user-owned boundaries and does not require a deep Pi fork.

## Release discipline

- Use annotated tags only.
- While pre-stable, use `v0.y.z` tags.
- Keep CI Linux-only while the repo is private.
- Do not mark a release item done without concrete command output, fixture inspection, generated artifact inspection, or a commit/PR reference.

## Completed implementation baseline

The old full-remake checklist is complete and should not remain in the main reading path. Git history and acceptance runs are the evidence layer for completed work. Durable rationale remains in ADRs, especially:

- `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`
- `docs/adr/0006-use-bounded-track-window-and-track-standard.md`
- `docs/adr/0010-use-curated-default-pi-packages.md`
- `docs/adr/0011-use-codex-first-craft-skill-pack.md`
- `docs/adr/0012-use-layered-pi-customization-profile.md`
