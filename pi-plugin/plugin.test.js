"use strict";

const assert = require("assert");
const { applyPrettyEnvironmentDefaults, buildGoalRunPrompt, buildQuietStatusSummary, buildRelevantLedgerContext, buildRtkRoutingHint, buildSubagentRoutingHint, buildWebResearchHint, commandRequiresRtk, extractAssistantProgress, getLedgerEntries, getRtkRoutingMode, hasLedgerConflict, isRtkRoutingEnabled, isRtkRoutingEnforced, isStaleLedgerEntry, isSubagentsConfigured, loadSaneConfig, makeLedgerEntry, parseGoalCommand, parseSaneToml, summarizeGoalState, validateSaneConfig } = require("./plugin");

const loaded = validateSaneConfig({
  defaults: { model: "gpt-5.5", reasoning: "low", responseStyle: "caveman" },
  packs: [{ id: "core-workflow", enabled: true, targets: ["pi", "codex"] }],
  userPacks: [],
  exportTargets: [{ id: "codex", kind: "codex-skill", path: ".codex/skills" }],
});

assert.equal(loaded.defaults.reasoning, "low");
assert.equal(loaded.defaults.responseStyle, "caveman");
assert.equal(loaded.packs.length, 1);
assert.equal(loaded.packs[0].id, "core-workflow");
assert.throws(() => validateSaneConfig({}), /defaults/);

const parsed = parseSaneToml(`
[defaults]
model = "gpt-5.5"
reasoning = "low"
response_style = "caveman"

[[packs]]
id = "agent-lanes"
enabled = true
targets = ["pi", "codex"]

[[user_packs]]
id = "custom-review"
enabled = true
targets = ["codex"]

[[export_targets]]
id = "codex"
kind = "codex-skill"
path = ".codex/skills"

[[recommended_pi_packages]]
id = "pi-subagents"
enabled = true
default_install = true
package = "npm:pi-subagents@0.24.0"
purpose = "Runtime-backed Pi subagent delegation for Sane agent lanes."

[[recommended_pi_packages]]
id = "pi-web-providers"
enabled = true
default_install = true
package = "npm:pi-web-providers@3.0.0"
purpose = "Configurable web search, content, answer, and research tools for freshness-sensitive work."
`);

assert.equal(parsed.defaults.responseStyle, "caveman");
assert.equal(parsed.packs[0].id, "agent-lanes");
assert.equal(parsed.userPacks[0].id, "custom-review");
assert.equal(parsed.packs[0].source, undefined);
assert.equal(parsed.packs[0].targets[1], "codex");
assert.equal(parsed.recommendedPiPackages.length, 2);
assert.equal(parsed.recommendedPiPackages[1].id, "pi-web-providers");
assert.equal(parsed.recommendedPiPackages[1].defaultInstall, true);

const prettyEnv = {};
assert.deepEqual(applyPrettyEnvironmentDefaults({ pretty: { maxPreviewLines: 24, maxHlChars: 1, icons: "nerd" } }, prettyEnv), {
  PRETTY_MAX_PREVIEW_LINES: "24",
  PRETTY_MAX_HL_CHARS: "1",
  PRETTY_ICONS: "nerd",
});

const realConfig = loadSaneConfig(require("node:path").join(__dirname, "config-schema.toml"));
assert.equal(getRtkRoutingMode(realConfig), "warn");
assert.match(buildRtkRoutingHint(realConfig), /mode=warn/);
assert.deepEqual(realConfig.pretty, { maxPreviewLines: 24, maxHlChars: 1, icons: "nerd" });
assert.equal(isRtkRoutingEnforced(realConfig), false);
const craftPackIds = ["craft-router", "frontend-craft", "frontend-review", "frontend-accessibility", "docs-writing", "ux-copy"];
for (const id of craftPackIds) {
  const pack = realConfig.packs.find((candidate) => candidate.id === id);
  assert.ok(pack, `real config missing ${id}`);
  assert.equal(pack.enabled, true, `${id} should be enabled`);
  assert.deepEqual(pack.targets, ["pi", "codex"], `${id} should target Pi and Codex`);
}
const subagentsPackage = realConfig.recommendedPiPackages.find((candidate) => candidate.id === "pi-subagents");
assert.ok(subagentsPackage, "real config missing pi-subagents package");
assert.equal(subagentsPackage.enabled, true);
assert.equal(subagentsPackage.defaultInstall, true);
assert.equal(isSubagentsConfigured(realConfig), true);
assert.match(buildQuietStatusSummary(realConfig, {}), /subagents=configured/);
assert.match(buildQuietStatusSummary(realConfig, { subagentsAvailable: true }), /subagents=ready/);
assert.match(buildSubagentRoutingHint(realConfig, "Do research and fix this properly"), /Sane subagent routing/);
assert.match(buildSubagentRoutingHint(realConfig, "Audit the whole codebase and verify CI"), /research, review\/verification, parallel\/disjoint/);
assert.match(buildSubagentRoutingHint(realConfig, "Rewrite all prompts in parallel"), /parallel\/disjoint/);
assert.match(buildSubagentRoutingHint(realConfig, "Review and verify the release workflow"), /review\/verification/);
assert.doesNotMatch(buildSubagentRoutingHint(realConfig, "Do research and fix this properly"), /choose not to delegate/);
assert.equal(buildSubagentRoutingHint(realConfig, "answer one tiny question"), "");

const loadedFromToml = loadSaneConfig("config-schema.toml", {
  readFileSync: () => `
[defaults]
model = "gpt-5.5"
reasoning = "low"

[rtk]
mode = "enforce"

[[packs]]
id = "rtk-routing"
enabled = true
targets = ["pi"]

[[export_targets]]
id = "pi"
kind = "pi-skill"
path = ".pi/skills"

[[recommended_pi_packages]]
id = "pi-web-providers"
enabled = true
default_install = true
package = "npm:pi-web-providers@3.0.0"
purpose = "Configurable web search, content, answer, and research tools for freshness-sensitive work."
`,
});

assert.equal(loadedFromToml.packs[0].id, "rtk-routing");
assert.equal(loadedFromToml.recommendedPiPackages[0].package, "npm:pi-web-providers@3.0.0");
assert.match(buildWebResearchHint(loadedFromToml, "research latest package status"), /Sane web hint/);
assert.equal(buildWebResearchHint(loadedFromToml, "edit local file"), "");
assert.equal(buildQuietStatusSummary(loadedFromToml, { activeGoal: { text: "ship" } }), "Sane: packs=1, subagents=off, web=ready, goal=active");
assert.equal(isRtkRoutingEnabled(loadedFromToml), true);
assert.equal(getRtkRoutingMode(loadedFromToml), "enforce");
assert.match(buildRtkRoutingHint(loadedFromToml), /mode=enforce/);
assert.equal(isRtkRoutingEnforced(loadedFromToml), true);
assert.equal(isRtkRoutingEnabled(loaded), false);
assert.equal(getRtkRoutingMode(loaded), "off");
assert.equal(buildRtkRoutingHint(loaded), "");
assert.equal(isRtkRoutingEnforced(loaded), false);

assert.equal(commandRequiresRtk("grep -R TODO ."), true);
assert.equal(commandRequiresRtk("cd repo && rg TODO"), true);
assert.equal(commandRequiresRtk("mkdir -p .tmp && ls -l .tmp"), true);
assert.equal(commandRequiresRtk("go build ./... && find . -maxdepth 1 -type f"), true);
assert.equal(commandRequiresRtk("go build ./... | grep ok"), true);
assert.equal(commandRequiresRtk("git diff -- cli"), true);
assert.equal(commandRequiresRtk("rtk grep TODO"), false);
assert.equal(commandRequiresRtk("command -v rtk"), false);
assert.equal(commandRequiresRtk("go build ./..."), false);

assert.deepEqual(parseGoalCommand("set ship the goal runner"), { action: "set", value: "ship the goal runner" });
assert.deepEqual(parseGoalCommand(""), { action: "status", value: "" });
assert.match(buildGoalRunPrompt("Audit everything"), /explicit goal as the current user objective/);
assert.match(buildGoalRunPrompt("Audit everything"), /TRACK\.toml are planning context, not a replacement/);

const goal = makeLedgerEntry("goal", { text: "Build goal runner", source: "user-command" });
const decision = makeLedgerEntry("decision", { text: "Use Pi extensions for goal state", status: "accepted" });
const progress = makeLedgerEntry("progress", { text: "Implemented command parsing" });
const sessionManager = {
  getEntries: () => [
    { type: "custom", customType: "sane-ledger", data: goal },
    { type: "custom", customType: "sane-ledger", data: decision },
    { type: "custom", customType: "sane-ledger", data: progress },
    { type: "custom", customType: "other", data: makeLedgerEntry("goal", { text: "ignore" }) },
  ],
};
const ledger = getLedgerEntries(sessionManager);
assert.equal(ledger.length, 3);
assert.equal(summarizeGoalState(ledger).activeGoal.text, "Build goal runner");
const context = buildRelevantLedgerContext(ledger, "continue the Pi extensions decision work");
assert.match(context, /Use as sourced context/);
assert.match(context, /Build goal runner/);
assert.match(context, /Use Pi extensions/);
const stale = makeLedgerEntry("decision", { text: "Use stale docs", timestamp: "2020-01-01T00:00:00.000Z" });
assert.equal(isStaleLedgerEntry(stale, new Date("2026-01-01T00:00:00.000Z"), 30), true);
assert.doesNotMatch(buildRelevantLedgerContext([stale], "stale docs", { now: "2026-01-01T00:00:00.000Z", maxAgeDays: 30 }), /Use stale docs/);
const d1 = makeLedgerEntry("decision", { text: "Use Pi overlay" });
const d2 = makeLedgerEntry("decision", { text: "Do not use Pi overlay" });
assert.equal(hasLedgerConflict(d1, [d1, d2]), true);
assert.match(buildRelevantLedgerContext([d1, d2], "Pi overlay"), /conflict/);

const assistantProgress = extractAssistantProgress([
  { role: "assistant", content: [{ type: "text", text: "Implemented tests.\nNext line." }] },
]);
assert.equal(assistantProgress, "Implemented tests.");
