# sane-next

**Sane** aims to be the best way to use your Codex subscription for real software work.

It is a Pi-first workflow overlay: Pi stays the fast interactive runtime, while Sane adds the skills, defaults, and habits that make Codex sessions more reliable. The goal is not another agent runtime. The goal is to turn a strong model subscription into a better coding environment: repo-aware, bounded, parallel when useful, and honest about verification.

Sane installs a small set of Pi skills, package recommendations, and CLI lifecycle commands that help agents start from repo truth, use focused craft skills, and leave recoverable handoffs instead of long, unreviewable chat drift.

> [!NOTE]
> `sane-next` is pre-stable. The current release line is `0.3.0-beta.4`, with GitHub Releases as the first distribution channel. There is no Homebrew or npm install channel yet.

## Contents

- [Philosophy](#philosophy)
- [What you get](#what-you-get)
- [Requirements](#requirements)
- [Install](#install)
- [Common tasks](#common-tasks)
- [Included packages and packs](#included-packages-and-packs)
- [Troubleshooting](#troubleshooting)
- [Project docs](#project-docs)

## Philosophy

Sane is built around a few product beliefs:

- **Codex is strongest with a good operating environment.** Sane supplies workflow structure, not a replacement brain.
- **Repo truth beats chat memory.** Agents should read current files and project rules before making broad claims.
- **Big work needs lanes.** Research, implementation, review, and verification can run in parallel, but one main agent should stay accountable.
- **Craft matters.** Docs, frontend, accessibility, review, and UX copy deserve focused skills instead of generic prompting.
- **Verification is part of the answer.** A session should end with concrete checks or an explicit handoff, not vibes.
- **User control comes first.** Sane can recommend packages and themes, but personal Pi preferences change only through explicit commands.

## What you get

- A Sane Pi overlay installed to `~/.sane-next` by default.
- Focused skill packs for core workflow discipline, agent lanes, docs, frontend craft, accessibility, review, UX copy, and Sane routing.
- Curated default Pi packages for subagents, checkpoints, and web research.
- Optional package recommendations for themes, previews, prompt routing, clarification, edit review, and sandboxing.
- Codex-native skill export for enabled packs.
- A thin Sane Core for Codex layer: managed `AGENTS.md` guidance plus optional warn/enforce shell-policy hooks.
- Safe lifecycle commands that track Sane-owned files and preserve user-owned config.

## Requirements

| Need | Used for |
| --- | --- |
| Go 1.22+ | Build and test the companion CLI from source. |
| Pi CLI/runtime | Load the overlay with `pi install`. |
| Node.js | Run Pi plugin tests. |
| Network/npm access | Install recommended Pi packages automatically. |

## Install

### Option 1: GitHub Release archive

Download the archive for your OS/CPU from the latest GitHub Release, extract it, then run the included binary.

<details open>
<summary>macOS/Linux</summary>

```bash
tar -xzf sane-next_0.3.0-beta.4_darwin_arm64.tar.gz
cd sane-next_0.3.0-beta.4_darwin_arm64
chmod +x sane-next
./sane-next version
./sane-next install
pi install ~/.sane-next
```

</details>

<details>
<summary>Windows PowerShell</summary>

```powershell
Expand-Archive sane-next_0.3.0-beta.4_windows_amd64.zip -DestinationPath sane-next
cd sane-next\sane-next_0.3.0-beta.4_windows_amd64
.\sane-next.exe version
.\sane-next.exe install
pi install $env:USERPROFILE\.sane-next
```

</details>

### Option 2: Build from source

```bash
cd cli
go build -o sane-next .
./sane-next install --source-root ..
pi install ~/.sane-next
```

Preview an install without writing to your home directory:

```bash
./sane-next install --root /tmp/sane-next-overlay --source-root .. --dry-run
./sane-next install --root /tmp/sane-next-overlay --source-root ..
./sane-next doctor --root /tmp/sane-next-overlay
```

> [!TIP]
> Use `--recommended-pi-packages=false` if you want to install the overlay without letting Sane call `pi install` for default package recommendations.

## Common tasks

| Task | Command |
| --- | --- |
| Check install health | `sane-next doctor --root ~/.sane-next` |
| Update Sane-owned files | `sane-next update --root ~/.sane-next` |
| Repair missing Sane-owned files | `sane-next repair --root ~/.sane-next` |
| Preview uninstall | `sane-next uninstall --root ~/.sane-next --dry-run` |
| Uninstall Sane-owned files | `sane-next uninstall --root ~/.sane-next` |
| List pack config | `sane-next pack list --config pi-plugin/config-schema.toml` |
| Validate pack config | `sane-next pack validate --config pi-plugin/config-schema.toml` |
| Export enabled packs to Codex | `sane-next export --target codex --config pi-plugin/config-schema.toml` |
| Install Sane Core for Codex | `sane-next codex install` |
| Install Codex warn-mode hooks | `sane-next codex install --hooks warn` |
| Install Codex enforce-mode shell hooks | `sane-next codex install --hooks enforce` |
| Check Codex layer health | `sane-next codex doctor` |
| Remove Codex managed blocks and Sane-owned hook config | `sane-next codex uninstall` |
| List Pi package recommendations | `sane-next package list --config pi-plugin/config-schema.toml` |
| Install one package recommendation | `sane-next package install --config pi-plugin/config-schema.toml pi-subagents` |
| Apply the bundled Pi theme | `sane-next configure --theme github-dark-pro` |

Use `--dry-run` where available to preview writes first.

## Included packages and packs

### Default Pi packages

These packages are installed by `sane-next install` unless disabled:

| ID | Package |
| --- | --- |
| `pi-subagents` | `npm:pi-subagents@0.24.0` |
| `pi-rewind` | `npm:pi-rewind@0.5.0` |
| `pi-web-providers` | `npm:pi-web-providers@3.0.0` |

### Optional Pi packages

Optional packages are listed in config but installed only when requested:

`pi-prompt-template-model`, `pi-curated-themes`, `pi-markdown-preview`, `pi-pretty`, `pi-ask-user`, `pi-pledit`, and `pi-container-sandbox`.

### Workflow packs

Enabled packs load in Pi and can export to Codex: `core-workflow`, `rtk-routing`, `agent-lanes`, `sane-router`, `craft-router`, `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, and `ux-copy`.

`Sane Core for Codex` keeps one source of truth: it exports those same packs, adds a tiny managed block to `~/.codex/AGENTS.md`, and can opt into warn- or enforce-mode hooks from `~/.sane-next/codex/hooks`. It does not change Codex model, sandbox, approval, MCP, or prompt settings by default.

Disabled example packs are kept as fixtures: `caveman-speak` and `example-user-pack`.

## Troubleshooting

| Problem | Try |
| --- | --- |
| `pi` is not found | Install/configure Pi first, or rerun install with `--recommended-pi-packages=false`. |
| Package installation fails | Check network/npm access, then install the package ID manually with `sane-next package install`. |
| `doctor` reports missing Sane-owned assets | Run `sane-next repair --root PATH`, then rerun `doctor`. |
| Codex export refuses to overwrite a directory | Move the user-owned directory or export to a different `--target-root`. |
| Codex hooks behave unexpectedly | Run `sane-next codex uninstall`, then reinstall without `--hooks warn` or `--hooks enforce`. |
| Theme config is risky to apply directly | Preview with `sane-next configure --theme github-dark-pro --agent-dir PATH --dry-run`. |

## Project docs

- [`docs/README.md`](docs/README.md) — documentation map.
- [`docs/roadmap/ROADMAP.md`](docs/roadmap/ROADMAP.md) — current and future product direction.
- [`docs/adr/`](docs/adr/) — accepted decisions.
- [`docs/standards/`](docs/standards/) — repo documentation and implementation rules.
