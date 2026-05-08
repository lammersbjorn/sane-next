# Contributing

This repo is still pre-implementation. Keep process small and strict.

## Current truth

- Current state: `TRACK.toml`
- Full remake roadmap: `docs/roadmap/ROADMAP.md`
- Old Sane reference context: `docs/reference/OLD-SANE-POSTMORTEM.md`
- Durable decisions: `docs/adr/`
- Tracking/prompt/skill standards: `docs/standards/`
- Repo-local implementation skill: `.agents/skills/sane-next-implementation/SKILL.md`
- Agent rules: `AGENTS.md`
- Copilot repo instructions: `.github/copilot-instructions.md`

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
research(pi): capture Pi extension boundaries
docs(rules): tighten prompt authoring rules
chore(repo): enable committed git hooks
feat(cli): add install command skeleton
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
- `commit-msg`
  - enforces the Conventional Commit format

## Versioning

- Use **annotated** tags only.
- While pre-stable, use `v0.y.z`.
- Keep CI Linux-only while the repo is private.
