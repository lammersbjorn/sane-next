# ADR 0015: Add a thin Codex core layer

## Status

Accepted

## Context

Sane Next already exports shared skills to Codex, but that left the Codex side as skills only. Current Codex documentation exposes a few native surfaces that are useful for a local distribution without becoming a Codex fork:

- global and project `AGENTS.md` custom instructions
- Agent Skills under `CODEX_HOME` / `~/.codex/skills`
- hooks configured from Codex config
- `config.toml` and MCP configuration

Custom prompts are deprecated, and slash commands are mainly built-ins. Codex internals and app-server behavior are moving quickly, so Sane should use only documented surfaces and keep all Codex-specific behavior small.

## Decision

Add **Sane Core for Codex** as a thin companion CLI layer:

```bash
sane-next codex install
sane-next codex export
sane-next codex doctor
sane-next codex uninstall
```

The layer uses one source of truth wherever possible:

1. Skills remain authored once under `.agents/skills/*` and keep targeting Pi and Codex through `pi-plugin/config-schema.toml`.
2. Codex skill installation reuses the existing `export --target codex` path.
3. Codex always-on instructions are a tiny managed block in `CODEX_HOME/AGENTS.md`.
4. Codex hook configuration is installed as schema-clean `CODEX_HOME/hooks.json`, with Sane ownership tracked outside the hook schema.
5. Hook scripts live under the Sane install root at `~/.sane-next/codex/hooks` when a script is needed; tiny lifecycle context may be emitted through an inline managed command to match Codex's proven hook behavior.
6. Codex config is edited only for feature flags and Sane-managed material; user-owned config outside that surface is preserved.

Default hook mode is `off`. Users can opt into shell-policy hooks with:

```bash
sane-next codex install --hooks warn
sane-next codex install --hooks enforce
```

`warn` adds developer context. `enforce` uses Codex `PreToolUse` denial for matched shell commands. MCP, model, sandbox, approval, and slash-command behavior remain out of the default Codex layer.

## Consequences

Positive:

- Codex gets native Sane entry points without duplicating the Pi overlay.
- Skills remain the shared substrate across Pi and Codex.
- Managed blocks make install, update, and uninstall reversible.
- Hooks can grow as executable policy instead of prompt bloat.

Negative:

- Codex hook config shape must be watched against upstream changes.
- `CODEX_HOME` support depends on the current skill export target path ending in `.codex`.
- Hook enforcement is intentionally narrow because Codex `PreToolUse` is a guardrail, not a complete sandbox boundary.

## Future rules

- Do not add Codex-only prompt copies unless a shared skill cannot express the behavior.
- Keep shared skills runtime-neutral except where a section explicitly distinguishes Pi and Codex behavior.
- Prefer generated artifacts, schema-clean hook files, or managed blocks over hand-maintained duplicate files.
- Keep Codex hooks small, inspectable, and configurable; use SessionStart for compact runtime context and avoid repeated UserPromptSubmit context unless a concrete decision/block is needed.
- Add MCP and model/sandbox config only as opt-in commands with explicit docs and tests.
