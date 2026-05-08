# ADR 0012: Use a layered Pi customization profile

## Status

Accepted

## Context

Sane Next is a Pi-first overlay, so it can improve the default Pi experience through three mechanisms:

- Sane-owned package assets such as extensions, skills, prompt templates, and themes.
- Curated third-party Pi packages installed through Pi's package manager.
- User settings such as theme, keybindings, model cycling, and UI preferences.

ADR 0010 allows a small pinned default package allowlist for high-value runtime gaps. The current defaults are subagents, rewind/checkpointing, and web providers. Recent package and documentation research shows additional useful Pi customization surfaces:

- Theme packages: `@victor-software-house/pi-curated-themes`, `@spences10/pi-themes`, and Sane-owned local themes.
- Tool rendering packages: `pi-claude-style-tools` and `@heyhuynhgiabuu/pi-pretty`.
- Markdown preview: `pi-markdown-preview`.
- Clarifying-question tools: `pi-ask-user` and `@juicesharp/rpiv-ask-user-question`.
- Plan-mode packages: `pi-pledit`, `@firstpick/pi-extension-plan-mode-toggle`, and related plan-mode packages.
- Sandbox/routing packages: `pi-container-sandbox`, `context-mode`, and similar higher-friction safety tools.
- Pi-native settings and extension APIs for themes, keybindings, status widgets, footer widgets, command registration, and prompt guidance.

The product constraint remains: Sane should make Pi more capable without becoming a heavy distribution, silently overwriting user preferences, or reintroducing old-Sane tool sprawl.

## Decision

Use a **layered Pi customization profile**.

### Layer 1: Sane-owned defaults

Ship only low-risk, Sane-owned defaults directly in the overlay:

- `github-dark-pro` theme under the Sane package `themes/` directory.
- Compact Sane extension behavior such as status hints, research-trigger guidance, and future configuration commands.

The install flow may make Sane-owned assets available by default, but it must not silently overwrite unrelated global Pi preferences. For example, the theme may be installed and discoverable by default, while setting it as the active global theme should be an explicit configuration action.

### Layer 2: Default-installed curated packages

Keep the default-installed curated packages narrowly scoped to runtime gaps that directly support Sane's workflow promise:

- `npm:pi-subagents@0.24.0` for agent lanes.
- `npm:pi-rewind@0.5.0` for checkpoint/rewind safety.
- `npm:pi-web-providers@3.0.0` for web search, contents, answers, and research.

These remain pinned and auditable under `pi-plugin/config-schema.toml`.

### Layer 3: Opt-in recommended packages

Add additional packages as **opt-in recommendations**, not default installs:

- `npm:@victor-software-house/pi-curated-themes@0.2.1`
  - best theme-pack choice because it is MIT, Pi-package-native, curated specifically for Pi's theme model, and includes many dark themes and GitHub dark variants.
  - selected over `@spences10/pi-themes` for breadth, curation detail, and explicit adaptation to Pi's 51-token theme model.
- `npm:pi-markdown-preview@0.9.7`
  - best docs/preview choice because it focuses on Markdown, LaTeX, code, diff, browser, terminal, and PDF preview without changing core coding behavior.
- `npm:@heyhuynhgiabuu/pi-pretty@0.4.4`
  - selected as the optional tool-rendering choice because it avoids the Claude Code-like presentation that proved undesirable in local use while still improving tool-call readability.
  - keep opt-in because it can affect command/output ergonomics and should not be part of Sane defaults.
  - enforce compact Sane defaults through the extension: `PRETTY_MAX_PREVIEW_LINES=24`, `PRETTY_MAX_HL_CHARS=1` to avoid Shiki read backgrounds clashing with the active Pi theme, and `PRETTY_ICONS=nerd` to keep icons.
- `npm:pi-ask-user@0.10.0`
  - best ask-user choice because it exposes one general `ask_user` tool plus a bundled skill for high-impact ambiguity and material assumptions.
  - selected over narrower question packages because Sane needs a general clarification primitive if it adopts one.
- `npm:pi-pledit@1.0.1`
  - best plan/accept-edits candidate for users who want external plan-mode behavior, because it provides both Plan Mode and Accept Edits Mode with a small permission-mode model.
  - keep opt-in and do not integrate into Sane defaults unless it is tested against Sane workflow and craft-router behavior.
- `npm:pi-container-sandbox@0.2.1`
  - best high-safety sandbox candidate for users who accept Docker/Apple-container setup cost.
  - keep opt-in and label as advanced because it changes file and shell execution semantics.

Do not default-install optional visual, plan, ask-user, preview, or sandbox packages.

### Layer 4: Explicit configuration commands

Add companion CLI support for applying Sane preferences explicitly. The implemented command surface is:

```bash
sane-next configure --theme github-dark-pro
sane-next package list
sane-next package install pi-markdown-preview
sane-next package install pi-ask-user
```

Configuration commands should preserve unrelated user settings, be dry-run capable where practical, and write only documented Pi settings files or call `pi install` for selected package IDs. Other preferences, such as quiet startup or keybindings, require separate implementation and tests before they should appear as executable docs.

## Implementation status

This layered profile is implemented in `pi-plugin/config-schema.toml`, the companion CLI package/configure commands, the Sane Pi extension, README user docs, and lifecycle/plugin tests. Optional package smoke installs should still require explicit user approval.

## Rejected alternatives

### Make every useful package a default install

Rejected. Visual renderers, previewers, plan modes, ask-user tools, and sandboxes are taste- or environment-sensitive and would raise startup/tool overhead.

### Use a third-party theme package instead of a Sane-owned default theme

Rejected for the default. A Sane-owned `github-dark-pro` theme gives a stable baseline without forcing a large theme pack on every user. Theme packs remain good opt-in recommendations.

### Choose `pi-claude-style-tools` as the preferred tool renderer

Rejected after local use. Its Claude Code-like tool-call presentation is too taste-specific for the recommended Sane path. `@heyhuynhgiabuu/pi-pretty` is the preferred optional renderer instead.

### Adopt plan mode by default

Rejected. Sane already has workflow discipline and craft-router behavior. Plan mode needs compatibility testing before becoming a default workflow.

### Adopt container sandboxing by default

Rejected. It has real safety value, but Docker/container requirements and changed filesystem semantics are too heavy for the default Sane path.

## Consequences

Positive:

- Sane can grow Pi customization without bloating default installs.
- Users get a clear path from minimal defaults to richer local setups.
- The selected package choices are documented and can be reviewed later.
- Theme and settings behavior respects user ownership.

Negative:

- The companion CLI grows beyond lifecycle/export into preference management.
- Optional package versions need periodic review against live Pi package pages, npm artifacts, and fixture installs because search indexes and repository default branches can lag published package versions.
- More documented choices can confuse users unless README wording stays compact.
