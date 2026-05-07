# sane-next

## Product frame

- `sane-next` is a **Pi-first overlay/distribution**, not a fresh runtime and not a deep Pi fork.
- **Primary runtime:** Pi.
- **Secondary export target:** Codex-native skill paths.
- Sane owns **packs, defaults, workflow discipline, export/install/update/doctor flows, and companion CLI behavior**.
- Pi owns the **agent loop, runtime, TUI, MCP bridge, and native layer**.

## Read first

- `TRACK.toml`
- `docs/adr/0002-use-track-toml-plus-adrs-for-tracking.md`
- `docs/adr/0004-use-pi-overlay-with-codex-skill-export.md`
- `.github/copilot-instructions.md`

Ignore deleted or superseded history. Use current repo files and current external research, not stale chat memory.

## Hard boundaries

- Do **not** rebuild Pi's runtime in this repo.
- Do **not** deep-fork Pi.
- Do **not** add a TUI-heavy standalone product surface.
- Do **not** create parallel `TODO`, `plan`, `memory`, or research-dump files in the repo.
- Do **not** copy old Sane code or trust old Sane claims without fresh evidence.
- Do **not** add broad always-on instruction files, generated repo summaries, or file-by-file overviews.

## Instruction surface rules

- Keep always-on context **small, durable, and high-signal**.
- Prefer **trigger-based skills** and on-demand references over large root instructions.
- One skill should do **one job**.
- Skill format should stay close to the shared ecosystem standard:
  - `name`
  - `description`
  - short body
  - optional `references/`
  - optional `scripts/`
  - optional `agents/`
- Keep the main skill body short and action-oriented. Put depth in `references/` and load it on demand.
- Prefer **negative constraints** for guardrails over broad positive style doctrine.
- Avoid verbosity constraints that risk lowering quality.
- Avoid repeated policy across `AGENTS.md`, Copilot instructions, skills, and ADRs. Put each rule in the smallest surface that can own it.

## Research rules

- For any implementation choice that depends on the current ecosystem, **research latest evidence first**.
- This applies especially to:
  - language/runtime choices
  - CLI stack
  - MCPs
  - Pi integrations
  - Codex/OpenAI behavior
  - tool routing and token optimization
  - external skills or packs
- Prefer primary sources, source repos, official docs, release notes, and directly inspectable code. Label mirrors or summaries as lower confidence.

## Work protocol

- Read `TRACK.toml` before broad work.
- Treat `TRACK.toml` as the only current-state ledger and ADRs as the only durable decision history.
- Keep ownership boundaries explicit and avoid overlapping edits.
- Use the narrowest real validation that matches the change. Until implementation exists, validate source consistency and decision alignment.
- When product direction changes materially, update `TRACK.toml` and add or update an ADR.

## Done

- Current files reflect current truth.
- Dead or superseded files are removed.
- Rules are not duplicated across multiple surfaces.
- New prompt/skill surfaces follow the minimal, trigger-first structure above.
