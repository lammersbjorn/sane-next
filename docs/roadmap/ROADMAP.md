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

- [ ] The full sane-next product can author shared packs, export them to Codex-native paths, load them through Pi integration, enable user-added packs, manage companion CLI lifecycle flows (`install`, `export`, `update`, `doctor`, `repair`, `uninstall`), and verify that only Sane-owned material is changed during install/uninstall.
- [ ] A single acceptance command exists and passes.
- [ ] `TRACK.toml` and this roadmap are updated only after matching behavior is implemented and verified.
- [ ] No checkbox is checked because of docs, placeholder files, scaffolding-only progress, shallow prompt cards, disconnected metadata, or unverifiable intent.

## Hard boundaries

- [ ] The old `sane` repo remains reference material only.
- [ ] No implementation code is copied from old Sane without an explicit evidence-backed reason.
- [ ] Pi remains the runtime in v1; sane-next does not become a deep Pi fork.
- [ ] The product remains config-first; no TUI-heavy standalone app is introduced.
- [ ] Packs and extensibility remain first-class scope, not optional extras.
- [ ] Generated output is disposable and is not treated as authored source.
- [ ] Random planning/TODO/research markdown files are not introduced outside the fixed repo structure.

---

# Phase 0 — `/goal` preflight

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

- [ ] The first shared pack source exists as `.agents/skills/core-workflow/SKILL.md`.
  Verify: file exists and matches `docs/standards/INSTRUCTION-SURFACE-STANDARD.md`.
- [ ] The Pi plugin manifest exists.
  Verify: `pi-plugin/manifest.toml` parses cleanly.
- [ ] The Pi config schema exists.
  Verify: `pi-plugin/config-schema.toml` parses cleanly and covers packs, model/reasoning defaults, and export targets.
- [ ] The Go CLI module exists.
  Verify: `cli/go.mod` exists and `go build ./...` succeeds.
- [ ] The CLI entrypoint exists.
  Verify: `cli/main.go` builds cleanly.
- [ ] The CLI exposes `install`.
  Verify: the install command runs and exits cleanly.
- [ ] Phase 1 behavior is verified and committed.
  Verify: milestone commit exists and the phase 1 checks above pass.

---

# Phase 2 — Shared packs, export, and Pi-side loading

- [ ] Shared pack content can export to a Codex-native location.
  Verify: exported skill artifact exists in a fixture export target.
- [ ] Pi-side config loading works without reaching into Pi internals.
  Verify: plugin/config path reads Sane config through the intended integration layer.
- [ ] The first pack is not decorative: it is actually used by export/load behavior.
  Verify: pack enablement changes produced artifacts or runtime behavior.
- [ ] The CLI exposes a justified export path.
  Verify: export command or script path works in a fixture target.
- [ ] Phase 2 behavior is verified and committed.
  Verify: milestone commit exists and the phase 2 checks above pass.

---

# Phase 3 — Packs and extensibility

- [ ] Multiple built-in packs can exist without overlapping ownership or behavior drift.
  Verify: at least two packs can coexist and render/load correctly.
- [ ] User-added packs have a supported discovery path.
  Verify: a fixture custom pack can be discovered and loaded.
- [ ] Packs can be enabled/disabled through the Sane config model.
  Verify: config changes alter exported or loaded pack behavior.
- [ ] Shared pack format stays compatible with both Pi and Codex export targets.
  Verify: the same authored pack source can drive both targets.
- [ ] Extensibility is verified and committed.
  Verify: milestone commit exists and the phase 3 checks above pass.

---

# Phase 4 — Companion CLI lifecycle

- [ ] CLI `install` is stable.
  Verify: fixture install succeeds and writes only Sane-owned material.
- [ ] CLI `export` is implemented.
  Verify: fixture export succeeds and outputs expected artifacts.
- [ ] CLI `update` is implemented.
  Verify: update changes Sane-owned material while preserving user-owned config.
- [ ] CLI `doctor` is implemented.
  Verify: doctor reports missing/broken install state meaningfully.
- [ ] CLI `repair` is implemented.
  Verify: repair restores a broken fixture install.
- [ ] CLI `uninstall` is implemented.
  Verify: uninstall removes only Sane-owned material and preserves user-owned config.
- [ ] Lifecycle behavior is verified and committed.
  Verify: milestone commit exists and the phase 4 checks above pass.

---

# Phase 5 — Verification, acceptance, and release discipline

- [ ] A single acceptance command exists.
  Verify: one documented repo command runs the full acceptance path.
- [ ] Acceptance covers install/export/load/extensibility/uninstall/recovery behavior.
  Verify: acceptance fixtures cover those behaviors.
- [ ] Acceptance fails on drift, placeholder-only implementations, and broken ownership boundaries.
  Verify: negative fixtures or tests prove the checks fail correctly.
- [ ] CI is added only as needed and stays Linux-only while the repo is private.
  Verify: workflow config shows Linux-only runners.
- [ ] Release discipline matches ADR 0009.
  Verify: commit/tag workflow matches annotated tags and `v0.y.z`.
- [ ] Final product acceptance is verified and committed.
  Verify: final acceptance command passes and the completion rule can be checked honestly.
