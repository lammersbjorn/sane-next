# Contributing

`sane-next` is implemented but still pre-stable. Keep changes small, verified, and grounded in the repo truth files.

## Current truth

- Current state: `TRACK.toml`
- Product roadmap and release discipline: `docs/roadmap/ROADMAP.md`
- Docs map and placement rules: `docs/README.md`, `docs/standards/DOCS-STRUCTURE-STANDARD.md`
- Old Sane reference context: `docs/reference/OLD-SANE-POSTMORTEM.md`
- Durable decisions: `docs/adr/`
- Tracking/prompt/skill standards: `docs/standards/`
- Repo-local implementation skill: `.agents/skills/sane-next-implementation/SKILL.md`
- Agent rules: `AGENTS.md`
- Copilot repo instructions: `.github/copilot-instructions.md`

Before broad implementation work, read `TRACK.toml` first, then follow only the docs and source files referenced by the active slice. Inspect current CLI behavior rather than relying on old chat context.

## Local verification

Use the narrowest check that proves your change, then run broader checks when behavior crosses boundaries.

```bash
cd cli && go test ./...
node --test pi-plugin/plugin.test.js
cd cli && ./acceptance.sh
(cd cli && SANE_NEXT_LIVE_PI=1 ./acceptance.sh) # optional live Pi install check
```

## Commit convention

Use **Conventional Commits**:

```text
type(scope): subject
```

Scope is optional.

Allowed types:

- `feat`
- `fix`
- `refactor`
- `docs`
- `research`
- `build`
- `ci`
- `test`
- `chore`
- `revert`

Recommended scopes:

- `repo`
- `rules`
- `tracking`
- `pi`
- `packs`
- `skills`
- `cli`
- `mcp`
- `tooling`
- `release`

Examples:

```text
docs(readme): clarify install side effects
fix(cli): preserve settings during theme configure
feat(packs): add accessibility skill export
```

Use lowercase type and scope. Keep the subject short, imperative, and specific.

## Agent commit cadence

When an agent is doing implementation work in this repo, git is the resume ledger:

- check `git status --short` before starting a new task lane
- commit each meaningful verified milestone before moving to another category
- split accumulated work by task/category before continuing if the tree gets mixed
- do not leave completed verified work only in chat context

## Git hooks

Committed hooks live in `.githooks/`.

Enable them locally:

```bash
git config core.hooksPath .githooks
```

Current hooks:

- `pre-commit`
  - fails on whitespace/conflict-marker issues
  - parses `TRACK.toml`
  - verifies required repo truth files exist
  - enforces the bounded TRACK shape
  - blocks markdown sprawl outside the fixed repo structure
  - requires the docs structure standard
- `commit-msg`
  - enforces the Conventional Commit format

## Versioning and CI

- Use SemVer and **annotated** tags only.
- While pre-stable, use `v0.y.z` tags.
- Run compatibility CI for Linux, macOS, and Windows; keep Bash acceptance on Linux unless a native acceptance slice is justified.
