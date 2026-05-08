# sane-next

`sane-next` is a small Pi-first overlay for coding-agent workflows. It installs a Sane Pi extension plus shared workflow packs, and can export those packs to Codex-native skill directories.

## What you get

- Pi overlay files that can be installed with `pi install`.
- Built-in workflow packs in `.agents/skills/`.
- Optional user packs registered from your own directories.
- A small Go companion CLI for install, export, update, doctor, repair, uninstall, pack management, package recommendations, and explicit Pi settings.
- Codex skill export without making Codex the primary runtime.

## Current install path

There is no published binary or Homebrew/npm channel yet. Build from source:

```bash
cd cli
go build -o sane-next .
```

Create a local overlay, inspect it, then install it into Pi:

```bash
./sane-next install --root /tmp/sane-next-overlay --source-root .. --dry-run
./sane-next install --root /tmp/sane-next-overlay --source-root ..
./sane-next doctor --root /tmp/sane-next-overlay
pi install /tmp/sane-next-overlay
```

By default, `install` writes to `~/.sane-next`. Use `--root PATH` for a preview, fixture, or alternate install location.

The install flow keeps the default Pi package allowlist small and pinned:

- `npm:pi-subagents@0.24.0` for Sane agent lanes.
- `npm:pi-rewind@0.5.0` for checkpoint/rewind safety.
- `npm:pi-web-providers@3.0.0` for freshness-sensitive web research.

Use `--recommended-pi-packages=false` to skip these package installs. Fixture or alternate-root installs do not mutate your global Pi package set unless explicitly enabled; if needed, install the lane runtime directly with `./sane-next package install pi-subagents`.

When `pi-subagents` is available, Sane surfaces it through `/sane-status`, compact startup hints for broad work, and `/sane-lanes <objective>`. You can also use the package commands directly, such as `/subagents-status`, `/subagents-doctor`, and `/parallel` when exposed by your Pi session.

## Export packs to Codex

Export enabled packs to Codex-native skills:

```bash
./sane-next export --target codex --config ../pi-plugin/config-schema.toml --dry-run
./sane-next export --target codex --config ../pi-plugin/config-schema.toml
```

The default Codex target is `~/.codex/skills`. Use `--target-root PATH` for a fixture or alternate export directory.

Exported pack directories include a `.sane-next-exported` marker. Existing unmarked directories are treated as user-owned and are not overwritten.

## Manage packs

List and validate configured packs:

```bash
./sane-next pack list --config ../pi-plugin/config-schema.toml
./sane-next pack validate --config ../pi-plugin/config-schema.toml
```

Enable or disable a configured pack ID:

```bash
./sane-next pack enable --config ../pi-plugin/config-schema.toml example-user-pack
./sane-next pack disable --config ../pi-plugin/config-schema.toml example-user-pack
```

## Configure Pi preferences and optional packages

Sane ships a local `github-dark-pro` theme asset, but it does not silently overwrite your active Pi theme. Apply it explicitly:

```bash
./sane-next configure --theme github-dark-pro
./sane-next configure --theme github-dark-pro --agent-dir /tmp/pi-agent-fixture
```

List optional package recommendations and install one by ID through `pi install`:

```bash
./sane-next package list --config ../pi-plugin/config-schema.toml
./sane-next package install --config ../pi-plugin/config-schema.toml pi-markdown-preview
./sane-next package install --config ../pi-plugin/config-schema.toml pi-curated-themes
```

Optional recommendations are disabled for default install and include curated themes, Markdown preview, compact tool rendering, ask-user clarification, plan/accept-edits mode, and advanced container sandboxing. When `pi-pretty` is installed, Sane defaults compact its previews, keeps Nerd Font icons, and disables Shiki read backgrounds that clash with the GitHub theme.

Package pins should be refreshed from live Pi package pages plus installable npm artifacts; public search snippets and repository default branches can lag the currently published package.

## Add a user pack

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

## Maintenance commands

```bash
./sane-next update --root ~/.sane-next --dry-run
./sane-next update --root ~/.sane-next
./sane-next repair --root ~/.sane-next
./sane-next uninstall --root ~/.sane-next --dry-run
./sane-next uninstall --root ~/.sane-next
```

`doctor` reports missing owned assets and suggests the next install or repair command.

## Development notes

- Acceptance: `./cli/acceptance.sh`
- CLI tests: `cd cli && go test ./...`
- Current work state: `TRACK.toml`
- Full roadmap: `docs/roadmap/ROADMAP.md`
- Durable decisions: `docs/adr/`
- Standards: `docs/standards/`

Release tags should remain annotated `v0.y.z` tags while this repo is private.
