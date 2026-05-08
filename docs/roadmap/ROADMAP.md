# sane-next ROADMAP

This roadmap is the execution ledger for the future Codex `/goal` implementation run.

Checkboxes are not decoration: a box may be checked only after the matching behavior exists in real repo code or runnable product behavior and is verified by a command, fixture, generated artifact inspection, deploy/uninstall proof, runtime hook execution, or other concrete evidence.

`TRACK.toml` is the bounded active phase. This roadmap is the **full remake checklist**.

## Recommended launch sequence

1. Start Codex in the `sane-next` repo with:
   - `gpt-5.5` if available, else `gpt-5.4`
   - reasoning `low`
   - `approval_policy = "on-request"`
   - `sandbox_mode = "workspace-write"`
2. Load `.agents/skills/sane-next-implementation/SKILL.md`.
3. Read `TRACK.toml`, this roadmap, `docs/reference/OLD-SANE-POSTMORTEM.md`, the core ADRs, and the core standards.
4. Set the `/goal` text below.
5. Work phase by phase, checking boxes only after verification.

## Recommended `/goal` prompt

```text
Build sane-next to completion as a better full remake of Sane, not a tiny MVP.

Read TRACK.toml, docs/roadmap/ROADMAP.md, docs/reference/OLD-SANE-POSTMORTEM.md, the core ADRs, the core standards, AGENTS.md, and .agents/skills/sane-next-implementation/SKILL.md before broad work.

Treat docs/roadmap/ROADMAP.md as the full verified execution ledger and TRACK.toml as the bounded active phase only.

Implement the full product across the roadmap phases: Pi-first overlay/plugin, shared packs, Codex export, extensibility and user-added packs, companion CLI lifecycle commands, install/update/doctor/repair/uninstall behavior, and final acceptance/release discipline.

Keep the old sane repo as read-only reference only. Do not copy old Sane code. Do not deep-fork Pi. Do not build a TUI-heavy standalone app. Do not add random docs or unbounded MCP/tool sprawl.

Commit after each meaningful verified milestone. Update ROADMAP checkboxes only after the matching behavior is implemented and verified. Update TRACK only for the active phase.

Stop only when the ROADMAP completion rule is satisfied and the final acceptance command passes, or when a true blocker, approval boundary, or budget limit requires handoff. If paused, resume from git state, ROADMAP, and TRACK and continue the remaining unchecked verified work.
```

## Completion rule

- [x] The Pi overlay provides goal/ledger workflow machinery: explicit goal commands, decision capture, progress capture, safe source-labeled context injection, and a keep-going command that resumes work toward the active goal until done, blocked, unsafe, or awaiting approval.
- [x] The full sane-next product can author shared packs, export them to Codex-native paths, load them through Pi integration, enable user-added packs, manage companion CLI lifecycle flows (`install`, `export`, `update`, `doctor`, `repair`, `uninstall`), and verify that only Sane-owned material is changed during install/uninstall.
- [x] A single acceptance command exists and passes.
- [x] `TRACK.toml` and this roadmap are updated only after matching behavior is implemented and verified.
- [x] No checkbox is checked because of docs, placeholder files, scaffolding-only progress, shallow prompt cards, disconnected metadata, or unverifiable intent.

## Hard boundaries

- [x] The old `sane` repo remains reference material only.
- [x] No implementation code is copied from old Sane without an explicit evidence-backed reason.
- [x] Pi remains the runtime in v1; sane-next does not become a deep Pi fork.
- [x] The product remains config-first; no TUI-heavy standalone app is introduced.
- [x] Packs and extensibility remain first-class scope, not optional extras.
- [x] Generated output is disposable and is not treated as authored source.
- [x] Random planning/TODO/research markdown files are not introduced outside the fixed repo structure.

---

# Phase 7 — Codex-first craft skill pack

- [x] ADR 0011 defines the craft-pack architecture and upstream policy.
  Verify: `docs/adr/0011-use-codex-first-craft-skill-pack.md` exists and is consistent with ADR 0008, ADR 0010, and the instruction-surface standard.
- [x] A small router skill exists and does not contain craft doctrine.
  Verify: `.agents/skills/craft-router/SKILL.md` exists and only classifies/dispatches frontend, docs, accessibility, review, and UX-copy work.
- [x] Frontend implementation, visual review, accessibility, docs-writing, and UX-copy skills exist as narrow subordinate skills.
  Verify: `.agents/skills/frontend-craft/SKILL.md`, `.agents/skills/frontend-review/SKILL.md`, `.agents/skills/frontend-accessibility/SKILL.md`, `.agents/skills/docs-writing/SKILL.md`, and `.agents/skills/ux-copy/SKILL.md` exist and follow `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`.
- [x] Upstream inspiration is recorded without vendoring large third-party prompt text.
  Verify: provenance notes exist under the relevant `references/UPSTREAM.md` files and no new large copied upstream prompt files are added.
- [x] The six craft skills are configured as built-in packs for both Pi and Codex targets.
  Verify: `pi-plugin/config-schema.toml` includes enabled pack entries for `craft-router`, `frontend-craft`, `frontend-review`, `frontend-accessibility`, `docs-writing`, and `ux-copy`.
- [x] Tests and acceptance cover the new pack set.
  Verify: `cd cli && go test ./...`, `node pi-plugin/plugin.test.js`, and `bash cli/acceptance.sh` pass, and fixture exports include the new skills.

---

# Phase 8 — Layered Pi customization profile

Start this phase only after Phase 7 craft-router work is implemented and acceptance passes.

- [x] ADR 0012 records package choices and the layered customization policy.
  Verify: `docs/adr/0012-use-layered-pi-customization-profile.md` exists and selects one preferred package for each optional category with rationale.
- [x] Optional Pi package recommendations are listed but not default-installed.
  Verify: `pi-plugin/config-schema.toml` includes disabled-by-default recommendations for curated themes, markdown preview, tool rendering, ask-user, plan mode, and sandbox packages while preserving the default-installed allowlist from ADR 0010.
- [x] The companion CLI can list and install configured package recommendations by ID.
  Verify: fixture-safe tests cover `sane-next package list` and `sane-next package install <id>` without mutating the real global Pi config.
- [x] The companion CLI can explicitly apply the Sane theme preference without overwriting unrelated settings.
  Verify: fixture-safe tests cover `sane-next configure --theme github-dark-pro` against a temporary Pi agent dir/settings file.
- [x] Sane's Pi extension gives compact runtime hints without adding broad always-on instructions.
  Verify: tests or source inspection cover a short web-research trigger and a quiet Sane status indicator when the relevant package/state is available.
- [x] User-facing docs explain default packages, optional package recommendations, and explicit configuration commands.
  Verify: README documents the commands and acceptance still passes.
- [x] Layered customization behavior is verified.
  Verify: `cd cli && go test ./...`, `node --test pi-plugin/plugin.test.js`, and `cd cli && ./acceptance.sh` pass.

---

# Phase 6 — Goal runner and safe ledger

- [x] Pi extension exposes a Sane goal command surface.
  Verify: `pi-plugin/index.ts` registers `sane-goal` and acceptance checks for it.
- [x] Explicit decisions and goals are persisted outside normal LLM context.
  Verify: tests cover `sane-ledger` custom entries and `pi.appendEntry` is wired.
- [x] Progress is captured automatically at agent-turn boundaries.
  Verify: tests cover assistant progress extraction and extension wires `agent_end`.
- [x] Relevant prior context is injected conservatively with source/confidence labels and conflict guidance.
  Verify: tests cover ledger context construction and acceptance checks for `buildRelevantLedgerContext`.
- [x] A goal runner can continue from the active goal and stops through normal user/approval/blocker boundaries.
  Verify: `sane-goal run` sends a bounded continuation prompt for the active goal; `sane-goal block` records blockers.
- [x] Goal/ledger behavior is included in the single acceptance command.
  Verify: `cli/acceptance.sh` passes.

---

# Phase 0 — `/goal` preflight

These are operator/session checks for a future Codex `/goal` run, not repo implementation requirements. They intentionally remain unchecked unless verified inside that live runtime session.

- [ ] GPT-5.5 access is verified locally, or an explicit GPT-5.4 fallback is chosen.
  Verify: start Codex with the target model and confirm via `/status`.
- [ ] Reasoning default is `low`.
  Verify: `/status` shows the active reasoning level.
- [ ] Approval policy is `on-request` and sandbox is `workspace-write`.
  Verify: `/status` shows both values.
- [ ] Only the minimal MCP/tool set needed for the run is enabled.
  Verify: active config or runtime status shows only the intended MCPs/tools.
- [ ] The old `sane` repo is available only as read-only reference, not as a writable root.
  Verify: writable roots exclude the old repo path.
- [x] The repo-local implementation skill exists.
  Verify: `.agents/skills/sane-next-implementation/SKILL.md` exists in the repo.
- [x] This roadmap contains the concrete `/goal` prompt and verified phase checklist.
  Verify: `docs/roadmap/ROADMAP.md` exists and is referenced by `README.md` and `TRACK.toml`.
- [ ] The `/goal` objective is set and visible.
  Verify: `/goal` shows the current objective and active goal state.

---

# Phase 1 — Foundations

- [x] The first shared pack source exists as `.agents/skills/core-workflow/SKILL.md`.
  Verify: file exists and matches `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`.
- [x] The Pi plugin manifest exists.
  Verify: `pi-plugin/manifest.toml` parses cleanly.
- [x] The Pi config schema exists.
  Verify: `pi-plugin/config-schema.toml` parses cleanly and covers packs, model/reasoning defaults, and export targets.
- [x] The Go CLI module exists.
  Verify: `cli/go.mod` exists and `go build ./...` succeeds.
- [x] The CLI entrypoint exists.
  Verify: `cli/main.go` builds cleanly.
- [x] The CLI exposes `install`.
  Verify: the install command runs and exits cleanly.
- [x] Phase 1 behavior is verified and committed.
  Verify: milestone commit exists and the phase 1 checks above pass.

---

# Phase 2 — Shared packs, export, and Pi-side loading

- [x] Shared pack content can export to a Codex-native location.
  Verify: exported skill artifact exists in a fixture export target.
- [x] Pi-side config loading works without reaching into Pi internals.
  Verify: plugin/config path reads Sane config through the intended integration layer.
- [x] The first pack is not decorative: it is actually used by export/load behavior.
  Verify: pack enablement changes produced artifacts or runtime behavior.
- [x] The CLI exposes a justified export path.
  Verify: export command or script path works in a fixture target.
- [x] Phase 2 behavior is verified and committed.
  Verify: milestone commit exists and the phase 2 checks above pass.

---

# Phase 3 — Packs and extensibility

- [x] Multiple built-in packs can exist without overlapping ownership or behavior drift.
  Verify: at least two packs can coexist and render/load correctly.
- [x] User-added packs have a supported discovery path.
  Verify: a fixture custom pack can be discovered and loaded.
- [x] Packs can be enabled/disabled through the Sane config model.
  Verify: config changes alter exported or loaded pack behavior.
- [x] Shared pack format stays compatible with both Pi and Codex export targets.
  Verify: the same authored pack source can drive both targets.
- [x] Extensibility is verified and committed.
  Verify: milestone commit exists and the phase 3 checks above pass.

---

# Phase 4 — Companion CLI lifecycle

- [x] CLI `install` is stable.
  Verify: fixture install succeeds and writes only Sane-owned material.
- [x] CLI `export` is implemented.
  Verify: fixture export succeeds and outputs expected artifacts.
- [x] CLI `update` is implemented.
  Verify: update changes Sane-owned material while preserving user-owned config.
- [x] CLI `doctor` is implemented.
  Verify: doctor reports missing/broken install state meaningfully.
- [x] CLI `repair` is implemented.
  Verify: repair restores a broken fixture install.
- [x] CLI `uninstall` is implemented.
  Verify: uninstall removes only Sane-owned material and preserves user-owned config.
- [x] Lifecycle behavior is verified and committed.
  Verify: milestone commit exists and the phase 4 checks above pass.

---

# Phase 5 — Verification, acceptance, and release discipline

- [x] A single acceptance command exists.
  Verify: one documented repo command runs the full acceptance path.
- [x] Acceptance covers install/export/load/extensibility/uninstall/recovery behavior.
  Verify: acceptance fixtures cover those behaviors.
- [x] Acceptance fails on drift, placeholder-only implementations, and broken ownership boundaries.
  Verify: negative fixtures or tests prove the checks fail correctly.
- [x] CI is added only as needed and stays Linux-only while the repo is private.
  Verify: workflow config shows Linux-only runners.
- [x] Release discipline matches ADR 0009.
  Verify: commit/tag workflow matches annotated tags and `v0.y.z`.
- [x] Final product acceptance is verified and committed.
  Verify: final acceptance command passes and the completion rule can be checked honestly.
