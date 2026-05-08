# sane-next

**Sane is a Pi-first workflow overlay for getting more reliable software work out of your Codex subscription.**

Codex is powerful, but real projects still fall apart when agents trust stale chat, overstuff prompts, skip verification, or leave changes too messy to review. Sane gives Codex a better local operating environment: Pi as the fast interactive runtime, shared workflow skills, bounded task discipline, agent lanes, safe lifecycle commands, and optional Codex-native export.

- Use your Codex subscription inside a stronger local workflow.
- Turn prompt habits into shared skills and repeatable workflows.
- Keep changes grounded in repo files, checks, and reviewable diffs.

```text
Developer -> Pi runtime -> Sane skills / lanes / adapters -> verified diff
                         \-> optional Codex-native skills and hooks
```

> [!NOTE]
> `sane-next` is pre-stable. The current release line is `0.3.0-beta.4`. GitHub Releases are the only packaged install channel; there is no Homebrew or npm channel yet.

## The problem Sane solves

Raw model access is not enough. Good agent work needs an environment:

- repo truth before chat memory
- compact skills instead of giant always-on prompts
- one main thread accountable for final edits
- lanes for research, review, and verification when work is broad
- durable docs and ADRs instead of chat-only decisions
- concrete checks before calling work done

Sane packages those habits for Codex-powered work. Pi is the primary runtime because it gives Sane the best local overlay, package, TUI, and extension surface today. Codex support is native but thin: the same shared skills can be exported to Codex, with optional managed instructions and hooks.

## What Sane is

Sane is:

- a Pi-first overlay installed under `~/.sane-next`
- a shared skill pack for workflow, routing, docs, frontend craft, accessibility, review, and UX copy
- a companion CLI for install, update, repair, doctor, export, package recommendations, theme config, and uninstall
- a small Codex Core layer for native skill export, managed `AGENTS.md` guidance, and optional shell-policy hooks

Sane is not a new agent runtime, a deep Pi fork, a Codex replacement, or a broad model/sandbox/MCP settings manager.

## Quick start

Download the archive for your OS/CPU from the latest GitHub Release.

### macOS/Linux

```bash
tar -xzf sane-next_0.3.0-beta.4_darwin_arm64.tar.gz
cd sane-next_0.3.0-beta.4_darwin_arm64
chmod +x sane-next
./sane-next version
./sane-next install
pi install ~/.sane-next
./sane-next doctor --root ~/.sane-next
```

### Windows PowerShell

```powershell
Expand-Archive sane-next_0.3.0-beta.4_windows_amd64.zip -DestinationPath sane-next
cd sane-next\sane-next_0.3.0-beta.4_windows_amd64
.\sane-next.exe version
.\sane-next.exe install
pi install $env:USERPROFILE\.sane-next
.\sane-next.exe doctor --root $env:USERPROFILE\.sane-next
```

Skip default Pi package recommendations if you want Sane to install only its own files:

```bash
./sane-next install --recommended-pi-packages=false
pi install ~/.sane-next
```

Undo the Pi overlay if you dislike it:

```bash
sane-next uninstall --root ~/.sane-next --dry-run
sane-next uninstall --root ~/.sane-next
```

## Example workflow

A safe feature change with Sane should look boring:

1. Start from repo instructions and the active tracker instead of chat memory.
2. Load the matching Sane skill instead of stuffing the prompt with repeated rules.
3. Split research, review, or verification into lanes when the task is broad.
4. Make the smallest useful edit.
5. Run the project check or name the concrete reason it could not run.
6. Review the diff and update durable docs when behavior changed.

Sane does not guarantee the patch is right. It improves the run shape so a skilled developer can supervise the work.

## Use with Codex

Codex support is optional. Use it when you want the same Sane skills available in Codex-native sessions.

Install skills plus the Sane-managed global `AGENTS.md` block:

```bash
sane-next codex install
sane-next codex doctor
```

Hooks are off by default. Add shell-policy hooks only if you want stronger command hygiene:

```bash
sane-next codex install --hooks warn
sane-next codex install --hooks enforce
```

Hook modes:

- `off` — default; exports skills and manages `AGENTS.md` only
- `warn` — warns on matched raw shell discovery/test/diff commands
- `enforce` — uses Codex `PreToolUse` denial for matched shell commands

Undo the Codex layer:

```bash
sane-next codex uninstall
```

The Codex layer reuses the shared Sane skill source. It does not change model, sandbox, approval, MCP, slash-command, or user-owned prompt settings by default.

## What gets installed

| Component | Path | Ownership model |
| --- | --- | --- |
| Pi overlay | `~/.sane-next` | Sane-owned overlay root |
| Shared packs | `~/.sane-next/packs` | Sane-owned generated pack files |
| User packs | `~/.sane-next/user-packs` | User-owned; preserved by uninstall |
| Codex skills | `~/.codex/skills` | exported shared skills; user-owned conflicts are protected |
| Codex instructions | `~/.codex/AGENTS.md` | Sane-managed block |
| Codex hooks | `~/.codex/hooks.json`, `~/.sane-next/codex/hooks` | Sane markers and Sane script files; user files are preserved |

Use `--dry-run` where available to preview writes.

## Build from source

```bash
cd cli
go build -o sane-next .
./sane-next install --source-root ..
pi install ~/.sane-next
```

Preview writes before changing your real overlay:

```bash
./sane-next install --root /tmp/sane-next-overlay --source-root .. --dry-run
./sane-next install --root /tmp/sane-next-overlay --source-root ..
./sane-next doctor --root /tmp/sane-next-overlay
```

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

## Requirements

| Need | Used for |
| --- | --- |
| Pi CLI/runtime | Primary local runtime for the Sane overlay. |
| Go 1.22+ | Build and test the companion CLI from source. |
| Node.js | Run Pi plugin tests. |
| Network/npm access | Install recommended Pi packages automatically. |
| Codex / ChatGPT plan | Optional Codex-native skill export and hooks. |

## Troubleshooting

| Problem | Try |
| --- | --- |
| I want to see writes before they happen | Use `--dry-run` where listed, inspect `~/.sane-next`, and run uninstall dry-runs before removing files. |
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
