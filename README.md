# sane-next

**Sane** is a Pi-first workflow overlay for using Codex on real software projects.

Pi stays the primary interactive runtime. Sane adds shared skills, a small companion CLI, curated Pi package recommendations, and optional Codex-native export so sessions start from repo truth, use focused workflow/craft guidance, and end with concrete verification instead of chat drift.

> [!NOTE]
> `sane-next` is pre-stable. The current release line is `0.3.0-beta.4`. GitHub Releases are the only packaged install channel; there is no Homebrew or npm channel yet.

## What Sane installs

- **Pi overlay files** under `~/.sane-next` by default.
- **Shared skill packs** for core workflow, RTK routing, agent lanes, Sane routing, docs, frontend craft/review/accessibility, and UX copy.
- **Curated Pi package recommendations** for subagents, checkpoints, and web research.
- **Companion CLI lifecycle commands** for install, update, repair, doctor, export, uninstall, package recommendations, theme config, and Codex setup.
- **Optional Codex Core layer** that exports the same skills to `~/.codex/skills`, manages a tiny block in `~/.codex/AGENTS.md`, and can install warn/enforce shell-policy hooks.

Sane preserves user-owned files. It tracks Sane-owned material with managed blocks, markers, or owned install directories and keeps model, sandbox, approval, MCP, and prompt settings unchanged unless a command explicitly says otherwise.

## Requirements

| Need | Used for |
| --- | --- |
| Go 1.22+ | Build and test the companion CLI from source. |
| Pi CLI/runtime | Load and run the Pi overlay. |
| Node.js | Run Pi plugin tests. |
| Network/npm access | Install recommended Pi packages automatically. |

## Install from a release archive

Download the archive for your OS/CPU from the latest GitHub Release.

### macOS/Linux

```bash
tar -xzf sane-next_0.3.0-beta.4_darwin_arm64.tar.gz
cd sane-next_0.3.0-beta.4_darwin_arm64
chmod +x sane-next
./sane-next version
./sane-next install
pi install ~/.sane-next
```

### Windows PowerShell

```powershell
Expand-Archive sane-next_0.3.0-beta.4_windows_amd64.zip -DestinationPath sane-next
cd sane-next\sane-next_0.3.0-beta.4_windows_amd64
.\sane-next.exe version
.\sane-next.exe install
pi install $env:USERPROFILE\.sane-next
```

Use this flag if you want Sane to install only its own files and skip default Pi package recommendations:

```bash
./sane-next install --recommended-pi-packages=false
```

## Build from source

```bash
cd cli
go build -o sane-next .
./sane-next install --source-root ..
pi install ~/.sane-next
```

Preview writes first:

```bash
./sane-next install --root /tmp/sane-next-overlay --source-root .. --dry-run
./sane-next install --root /tmp/sane-next-overlay --source-root ..
./sane-next doctor --root /tmp/sane-next-overlay
```

## Codex setup

Codex support is optional and secondary to the Pi overlay. It reuses the same enabled skill packs instead of maintaining separate Codex-only skill sources.

Install skills plus the managed global `AGENTS.md` block:

```bash
sane-next codex install
```

Install shell-policy hooks too:

```bash
sane-next codex install --hooks warn
sane-next codex install --hooks enforce
```

Hook modes:

- `off` — default; exports skills and manages `AGENTS.md` only.
- `warn` — adds Codex hook warnings for matched raw shell discovery/test/diff commands.
- `enforce` — uses Codex `PreToolUse` denial for matched shell commands.

Codex lifecycle commands:

```bash
sane-next codex export
sane-next codex doctor
sane-next codex uninstall
```

`codex uninstall` removes Sane-owned Codex managed blocks, hook config, hook markers, and Sane hook scripts. It preserves user-owned Codex config and user files in the Sane hook directory.

## Common commands

### Overlay lifecycle

```bash
sane-next doctor --root ~/.sane-next
sane-next update --root ~/.sane-next
sane-next repair --root ~/.sane-next
sane-next uninstall --root ~/.sane-next --dry-run
sane-next uninstall --root ~/.sane-next
```

### Packs and export

```bash
sane-next pack list --config pi-plugin/config-schema.toml
sane-next pack validate --config pi-plugin/config-schema.toml
sane-next export --target codex --config pi-plugin/config-schema.toml
```

### Pi packages and theme

```bash
sane-next package list --config pi-plugin/config-schema.toml
sane-next package install --config pi-plugin/config-schema.toml pi-subagents
sane-next configure --theme github-dark-pro
```

Use `--dry-run` where available to preview writes.

## Included packages and packs

### Default Pi packages

Installed by `sane-next install` unless disabled:

| ID | Package |
| --- | --- |
| `pi-subagents` | `npm:pi-subagents@0.24.0` |
| `pi-rewind` | `npm:pi-rewind@0.5.0` |
| `pi-web-providers` | `npm:pi-web-providers@3.0.0` |

### Optional Pi packages

Listed in config but installed only when requested:

`pi-prompt-template-model`, `pi-curated-themes`, `pi-markdown-preview`, `pi-pretty`, `pi-ask-user`, `pi-pledit`, and `pi-container-sandbox`.

### Enabled workflow packs

`core-workflow`, `rtk-routing`, `agent-lanes`, `sane-router`, `craft-router`, `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, and `ux-copy`.

Disabled example packs are fixtures: `caveman-speak` and `example-user-pack`.

## Troubleshooting

| Problem | Try |
| --- | --- |
| `pi` is not found | Install/configure Pi first, or rerun install with `--recommended-pi-packages=false`. |
| Package installation fails | Check network/npm access, then install the package ID manually with `sane-next package install`. |
| `doctor` reports missing Sane-owned assets | Run `sane-next repair --root PATH`, then rerun `doctor`. |
| Codex export refuses to overwrite a directory | Move the user-owned directory or export to a different `--target-root`. |
| Codex hooks behave unexpectedly | Run `sane-next codex doctor`; if needed, run `sane-next codex uninstall`, then reinstall with `--hooks off`, `--hooks warn`, or `--hooks enforce`. |
| Theme config is risky to apply directly | Preview with `sane-next configure --theme github-dark-pro --agent-dir PATH --dry-run`. |

## Project docs

- [`docs/README.md`](docs/README.md) — documentation map.
- [`docs/roadmap/ROADMAP.md`](docs/roadmap/ROADMAP.md) — current and future product direction.
- [`docs/adr/`](docs/adr/) — accepted decisions.
- [`docs/standards/`](docs/standards/) — repo documentation and implementation rules.
