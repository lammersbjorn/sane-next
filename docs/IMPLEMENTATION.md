# Implementation

This file gives the implementation agent the exact build order for `sane-next`.

## Read first

1. `TRACK.toml`
2. `docs/FOUNDATION.md`
3. `docs/IMPLEMENTATION.md`
4. `AGENTS.md`

Do not rely on chat history.

## Build order

### Phase 1 — scaffold the core surfaces

Create exactly these paths first:

```text
.agents/skills/core-workflow/SKILL.md
pi-plugin/manifest.toml
pi-plugin/config-schema.toml
cli/go.mod
cli/main.go
cli/cmd/install.go
```

Requirements:

1. `core-workflow` is Pi/Codex compatible.
2. The skill follows the fixed skill contract:
   - Goal
   - Use When
   - Don't Use When or Use Main Session When
   - Inputs
   - Outputs
   - How To Run
   - Verification
   - Gotchas / Safety
3. The Pi plugin stays inside extension/plugin boundaries.
4. The CLI uses Go and starts with `install` only.

Done when:

1. the skill exists and is structured correctly
2. the TOML files parse cleanly
3. `go build ./...` succeeds
4. the install command exits cleanly

### Phase 2 — wire export and config loading

After phase 1 exists:

1. add the first export path from shared pack content to a Codex-native location
2. make the Pi plugin read Sane config without reaching into Pi internals

Done when:

1. a CLI or script path can export the first skill
2. the Pi side reads config successfully

### Phase 3 — prove the first slice end to end

After phase 2 works:

1. verify install, export, and plugin/config behavior together
2. add minimal Linux-only CI only if it is needed for the acceptance path

Done when:

1. the full first-slice acceptance path passes
2. any CI added remains Linux-only while the repo is private

## Lane boundaries

Use these write boundaries when parallelizing work:

| Lane | Write boundary |
| --- | --- |
| `skill-pack` | `.agents/skills/core-workflow/` |
| `pi-plugin` | `pi-plugin/` |
| `cli-install` | `cli/` |
| `export-load` | `cli/`, `pi-plugin/` |
| `acceptance` | `cli/`, `pi-plugin/`, `.github/workflows/` |

Rules:

1. one writing lane owns one path tree
2. `TRACK.toml`, `docs/FOUNDATION.md`, and `docs/IMPLEMENTATION.md` are coordinator-owned unless the task is explicitly changing repo truth
3. every lane must report changed files, verification, and unresolved risk

## What not to build yet

Do not add:

- a deep Pi fork
- a TUI-heavy product surface
- extra packs before the first pack works
- model-routing complexity before evidence exists
- mandatory third-party MCP dependencies
- extra planning or research files
