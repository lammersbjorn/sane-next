# Implementation Run Protocol

This is the strict execution standard for the first `sane-next` implementation run.

The **active phase** should be intentionally bounded, but the **product scope** is the broader remake: packs, extensibility, Pi integration, Codex export, and the Sane companion CLI lifecycle.

The repo should not confuse "phase 1" with "the whole product."

## Surface ownership

| Concern | Owner |
| --- | --- |
| Current execution window | `TRACK.toml` |
| Tracking shape | `docs/standards/TRACK-STRUCTURE-STANDARD.md` |
| `/goal` implementation-run setup | `docs/standards/GOAL-RUN-STANDARD.md` |
| Durable decisions | `docs/adr/` |
| Instruction/skill/agent authoring | `docs/standards/INSTRUCTION-SURFACE-STANDARD.md` |
| Implementation run shape | `docs/standards/IMPLEMENTATION-RUN-PROTOCOL.md` |
| Commit hygiene | `.githooks/` + `CONTRIBUTING.md` |

## Full remake scope

The later implementation run is for a **full remake of Sane that is better**, not a tiny end-state MVP.

That remake includes:

- shared pack authoring
- Codex export of shared pack content
- Pi plugin/config integration
- extensibility and user-added packs
- the small Sane companion CLI for install/export/update/doctor/repair flows
- explicit verification and recovery

`TRACK.toml` only holds the current phase because the execution window should be bounded. The product scope still includes the larger remake.

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
  cmd/repair.go

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

## Phase roadmap

### Phase 1 — foundations

Create the foundation surfaces without pretending phase 1 is the whole product:

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

### Phase 2 — packs and export/load

After the phase 1 foundations exist:

1. wire the first shared-pack export path to Codex-native locations
2. make the Pi side load Sane config and pack state without reaching into Pi internals
3. prove the pack model is actually central, not decorative

### Phase 3 — extensibility and CLI lifecycle

After export/load works:

1. add the path for user-added packs and explicit extensibility
2. extend the companion CLI beyond `install` toward `export`, `update`, `doctor`, and `repair`
3. keep the config model small while supporting packs and exports cleanly

### Phase 4 — acceptance, recovery, and release discipline

After the main remake surfaces exist:

1. verify install/export/extensibility/recovery behavior end to end
2. add only the CI needed for the acceptance path
3. keep release discipline aligned with ADR 0009

## Lane sequence

| Lane | Scope | Write boundary | Can run in parallel |
| --- | --- | --- | --- |
| `skill-pack` | first shared skill | `.agents/skills/core-workflow/` | yes |
| `pi-plugin` | plugin manifest + config schema | `pi-plugin/` | yes |
| `cli-install` | Go CLI install scaffold | `cli/` | yes |
| `export-load` | shared pack export + Pi-side loading | `.agents/skills/`, `cli/`, `pi-plugin/` | after first three |
| `extensibility` | user packs + lifecycle commands | `.agents/skills/`, `cli/`, `pi-plugin/` | after export/load |
| `acceptance` | end-to-end proof + CI if needed | `cli/`, `pi-plugin/`, `.github/workflows/` | after extensibility |

## Acceptance path

Done means all of the following are true:

1. `core-workflow/SKILL.md` exists and follows the required skill structure.
2. `pi-plugin/manifest.toml` and `pi-plugin/config-schema.toml` parse cleanly.
3. `go build ./...` for the CLI succeeds.
4. the install command exits cleanly.
5. the first export/load path is wired without rebuilding Pi internals.
6. the repo has a clear path for packs/extensibility beyond the initial pack.
7. any CI added remains Linux-only while the repo is private.

## Open risks that remain outside TRACK

- default model should remain configurable rather than hard-coded
- Pi plugin API stability may force adjustment later
- export-path details may need revision once real Pi/Codex paths are tested
