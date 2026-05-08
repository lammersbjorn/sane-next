# sane-next

Sane is a better way to use Codex for serious coding work.

It gives Codex a stronger working environment: Pi as the fast interactive runtime, focused skills for the kinds of work you actually do, subagent lanes for broad tasks, web/research support when freshness matters, and checkpoint-friendly habits for long sessions.

Instead of asking you to remember a perfect prompt every time, Sane installs a small overlay that makes the good workflow the default:

- **Start from repo truth.** Sane nudges agents to read the right local files, respect existing project rules, and avoid stale chat memory.
- **Keep work bounded.** Long tasks are split into clear slices with explicit stop conditions instead of drifting into an unreviewable mega-run.
- **Use the right skill at the right time.** Core workflow, agent lanes, docs, frontend craft, accessibility, review, UX copy, and RTK-aware routing live as focused skill packs.
- **Parallelize when it helps.** Broad work can use subagent lanes for research, implementation, review, and verification while one main agent stays in control.
- **Verify before claiming done.** Sane emphasizes concrete checks, acceptance commands, and recoverable handoffs.
- **Stay portable.** Packs are authored once and can load in Pi or export to Codex-native skill directories.
- **Respect your setup.** Optional packages and themes are available, but personal preferences are changed only by explicit commands.

`sane-next` implements this as a Pi-first distribution: a Sane Pi extension, shared workflow packs, curated Pi package defaults, and a small companion CLI for install, export, update, repair, and configuration.

## Status

- Current version: `0.2.0` from source.
- No published binary, Homebrew formula, or npm channel yet.
- Build the CLI locally from `cli/`.

## Prerequisites

- Go 1.22+ to build and test the companion CLI.
- A working `pi` CLI/runtime for `pi install` and Pi overlay loading.
- Node.js only when running Pi plugin tests.
- Network/npm access if you let Sane install recommended Pi packages.

## Quick start from source

```bash
cd cli
go build -o sane-next .
```

Preview the overlay in a fixture directory, install it there, and inspect it:

```bash
./sane-next install --root /tmp/sane-next-overlay --source-root .. --dry-run
./sane-next install --root /tmp/sane-next-overlay --source-root ..
./sane-next doctor --root /tmp/sane-next-overlay
```

Install that overlay into Pi:

```bash
pi install /tmp/sane-next-overlay
```

By default, `install` writes to `~/.sane-next`. Use `--root PATH` for a preview, fixture, or alternate install location.

## What Sane installs by default

The default install keeps the Pi package allowlist small and pinned:

| ID | Package | Why it is default |
| --- | --- | --- |
| `pi-subagents` | `npm:pi-subagents@0.24.0` | Runtime-backed agent lanes for broad Sane work. |
| `pi-rewind` | `npm:pi-rewind@0.5.0` | Checkpoint/rewind safety for long sessions. |
| `pi-web-providers` | `npm:pi-web-providers@3.0.0` | Fresh web search/research tools. |

Default-root installs may call `pi install` for those packages. To skip package installation:

```bash
./sane-next install --recommended-pi-packages=false
```

Fixture or alternate-root installs do not mutate your global Pi package set unless package installation is explicitly enabled for that path. If you only need lanes, install the runtime directly:

```bash
./sane-next package install pi-subagents
```

## Optional Pi packages

Optional recommendations are listed but not installed by default:

| ID | Package | Use when you want... |
| --- | --- | --- |
| `pi-prompt-template-model` | `npm:pi-prompt-template-model@0.9.3` | Prompt-template routing for model/reasoning/skill workflows. |
| `pi-curated-themes` | `npm:@victor-software-house/pi-curated-themes@0.2.1` | More Pi themes. |
| `pi-markdown-preview` | `npm:pi-markdown-preview@0.9.7` | Markdown, code, diff, browser, terminal, and PDF previews. |
| `pi-pretty` | `npm:@heyhuynhgiabuu/pi-pretty@0.4.4` | Compact visual rendering for tool calls. |
| `pi-ask-user` | `npm:pi-ask-user@0.10.0` | A clarification tool and skill for important ambiguity. |
| `pi-pledit` | `npm:pi-pledit@1.0.1` | Plan mode and accept-edits mode. |
| `pi-container-sandbox` | `npm:pi-container-sandbox@0.2.1` | Higher-safety container sandboxing. |

List and install package recommendations by ID:

```bash
./sane-next package list --config ../pi-plugin/config-schema.toml
./sane-next package install --config ../pi-plugin/config-schema.toml pi-markdown-preview
```

When `pi-pretty` is installed, Sane supplies compact defaults: short previews, Nerd Font icons, and minimal highlighting to avoid theme clashes.

## Configure Pi preferences

Sane ships a local `github-dark-pro` theme asset, but it does not silently change your active Pi theme. Apply it explicitly:

```bash
./sane-next configure --theme github-dark-pro
```

For a fixture or alternate Pi agent settings directory:

```bash
./sane-next configure --theme github-dark-pro --agent-dir /tmp/pi-agent-fixture --dry-run
./sane-next configure --theme github-dark-pro --agent-dir /tmp/pi-agent-fixture
```

The configure command merges Pi settings and preserves unrelated keys.

## Workflow packs

Built-in packs are authored once under `.agents/skills/` and can be loaded by Pi or exported to Codex. Enabled packs include core workflow discipline, RTK routing, agent lanes, Sane routing, and focused craft skills for frontend, accessibility, docs, and UX copy. Disabled examples include `caveman-speak` and the fixture `example-user-pack`.

Manage configured packs:

```bash
./sane-next pack list --config ../pi-plugin/config-schema.toml
./sane-next pack validate --config ../pi-plugin/config-schema.toml
./sane-next pack enable --config ../pi-plugin/config-schema.toml example-user-pack
./sane-next pack disable --config ../pi-plugin/config-schema.toml example-user-pack
```

### Add a user pack

Create a directory containing `SKILL.md`, then register it in a Sane config file:

```toml
[[user_packs]]
id = "my-review-pack"
enabled = true
source = "/absolute/path/to/my-review-pack"
targets = ["pi", "codex"]
```

Then validate and export:

```bash
./sane-next pack validate --config ../pi-plugin/config-schema.toml
./sane-next export --config ../pi-plugin/config-schema.toml --target codex
```

## Export packs to Codex

Export enabled packs to Codex-native skills:

```bash
./sane-next export --target codex --config ../pi-plugin/config-schema.toml --dry-run
./sane-next export --target codex --config ../pi-plugin/config-schema.toml
```

The default Codex target is `~/.codex/skills`. Use `--target-root PATH` for a fixture or alternate export directory. Exported directories include a `.sane-next-exported` marker. Existing unmarked directories are treated as user-owned and are not overwritten.

## Lifecycle commands and ownership

Sane-owned install assets use the `.sane-next-owned` marker. Lifecycle commands are designed to update, repair, or remove Sane-owned material while preserving user-owned config.

```bash
./sane-next doctor --root ~/.sane-next
./sane-next update --root ~/.sane-next --dry-run
./sane-next update --root ~/.sane-next
./sane-next repair --root ~/.sane-next
./sane-next uninstall --root ~/.sane-next --dry-run
./sane-next uninstall --root ~/.sane-next
```

## Troubleshooting

- `pi` not found: install/configure Pi first, or use `--recommended-pi-packages=false` while building the overlay.
- Package install fails: retry with network/npm access, or install the package manually with the package spec shown by `package list`.
- `doctor` reports missing Sane-owned assets: run `repair --root PATH`, then rerun `doctor`.
- Export refuses to overwrite a skill directory: remove or move the user-owned directory, or export to a different `--target-root`.
- Theme configuration looks wrong: run with `--agent-dir PATH --dry-run` first and inspect the target Pi settings file.
