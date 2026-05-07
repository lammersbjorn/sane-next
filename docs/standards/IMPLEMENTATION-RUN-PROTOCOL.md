# Implementation Run Protocol

This is the strict execution standard for the first `sane-next` implementation run.

It is intentionally small: one authoritative surface per concern, one clear first slice, and exact write boundaries for parallel work.

## Surface ownership

| Concern | Owner |
| --- | --- |
| Current execution window | `TRACK.toml` |
| Tracking shape | `docs/standards/TRACK-STRUCTURE-STANDARD.md` |
| Durable decisions | `docs/adr/` |
| Instruction/skill/agent authoring | `docs/standards/INSTRUCTION-SURFACE-STANDARD.md` |
| Implementation run shape | `docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md` |
| Commit hygiene | `.githooks/` + `CONTRIBUTING.md` |

## One-shot prompt contract

A one-shot implementation prompt must include:

1. the exact outcome for the run
2. the current truth files to read first
3. the implementation phase being attempted
4. a lane table with exact write boundaries
5. the acceptance path that defines done
6. stop conditions or approval points for risky choices
7. the expected handoff format for subagents

It must not include:

- a giant repo summary
- duplicate policy copied from other surfaces
- broad future backlog
- speculative framework ideas unrelated to the current slice

## Parallel agent rules

1. One writing lane owns one path tree.
2. `TRACK.toml`, `docs/adr/`, and `docs/standards/` are coordinator-owned surfaces unless the run is explicitly about those files.
3. If two writing lanes run at once, use separate worktrees when practical.
4. Reading lanes may inspect any repo file they need, but they do not edit outside their boundary.
5. Every lane must return:
   - changed files
   - what was verified
   - unresolved risk or follow-up needed
6. No lane claims success without the verification named in its prompt.

## Initial file structure

### Create first

```text
.agents/skills/core-workflow/
  SKILL.md

pi-plugin/
  manifest.toml
  config-schema.toml

cli/
  go.mod
  main.go
  cmd/install.go
```

### Create later

```text
cli/
  cmd/export.go
  cmd/update.go
  cmd/doctor.go

.github/workflows/
  ci.yml

.agents/skills/
  <additional packs>
```

Do **not** create repo-local `TODO`, `plan`, `memory`, or research-dump files.

## Tool and MCP policy

Apply ADR 0008:

- no third-party MCP is required by default
- RTK is recommended when the target repo uses RTK
- Playwright CLI is preferred over a browser MCP
- Context7 and grep.app are optional recommended external helpers
- additional MCPs must stay explicit and opt-in

## Release policy

Apply ADR 0009:

- `v0.y.z`
- annotated tags only
- Linux-only CI while private
- no full release automation yet

## First implementation slice

The first slice proves the product shape without rebuilding Pi:

1. **Core workflow skill**
   - create `.agents/skills/core-workflow/SKILL.md`
   - keep it Pi/Codex compatible
   - follow the instruction-surface standard exactly

2. **Pi plugin scaffold**
   - create `pi-plugin/manifest.toml`
   - create `pi-plugin/config-schema.toml`
   - keep the schema limited to packs, model/reasoning defaults, and export targets

3. **Companion CLI scaffold**
   - use Go per ADR 0007
   - create `cli/go.mod`, `cli/main.go`, and `cli/cmd/install.go`
   - start with `install` only

4. **Export/load follow-up**
   - after the above exists, wire the first export path and config loading path
   - do not reach into Pi internals when the extension API is enough

## Lane sequence

| Lane | Scope | Write boundary | Can run in parallel |
| --- | --- | --- | --- |
| `skill-pack` | first shared skill | `.agents/skills/core-workflow/` | yes |
| `pi-plugin` | plugin manifest + config schema | `pi-plugin/` | yes |
| `cli-install` | Go CLI install scaffold | `cli/` | yes |
| `export-load` | export path + config loading | `cli/`, `pi-plugin/` | after first three |
| `acceptance` | end-to-end proof + CI if needed | `cli/`, `pi-plugin/`, `.github/workflows/` | after export/load |

## First-slice acceptance path

Done means all of the following are true:

1. `core-workflow/SKILL.md` exists and follows the required skill structure.
2. `pi-plugin/manifest.toml` and `pi-plugin/config-schema.toml` parse cleanly.
3. `go build ./...` for the CLI succeeds.
4. the install command exits cleanly.
5. the first export/load path is wired without rebuilding Pi internals.
6. any CI added remains Linux-only while the repo is private.

## Open risks that remain outside TRACK

- default model should remain configurable rather than hard-coded
- Pi plugin API stability may force adjustment later
- export-path details may need revision once real Pi/Codex paths are tested
