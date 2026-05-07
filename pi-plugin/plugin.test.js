"use strict";

const assert = require("assert");
const { loadSaneConfig, parseSaneToml, validateSaneConfig } = require("./plugin");

const loaded = validateSaneConfig({
  defaults: { model: "gpt-5.5", reasoning: "low" },
  packs: [{ id: "core-workflow", enabled: true, targets: ["pi", "codex"] }],
  userPacks: [],
  exportTargets: [{ id: "codex", kind: "codex-skill", path: ".codex/skills" }],
});

assert.equal(loaded.defaults.reasoning, "low");
assert.equal(loaded.packs.length, 1);
assert.equal(loaded.packs[0].id, "core-workflow");
assert.throws(() => validateSaneConfig({}), /defaults/);

const parsed = parseSaneToml(`
[defaults]
model = "gpt-5.5"
reasoning = "low"

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
`);

assert.equal(parsed.packs[0].id, "agent-lanes");
assert.equal(parsed.userPacks[0].id, "custom-review");
assert.equal(parsed.packs[0].targets[1], "codex");

const loadedFromToml = loadSaneConfig("config-schema.toml", {
  readFileSync: () => `
[defaults]
model = "gpt-5.5"
reasoning = "low"

[[packs]]
id = "rtk-routing"
enabled = true
targets = ["pi"]

[[export_targets]]
id = "pi"
kind = "pi-skill"
path = ".pi/skills"
`,
});

assert.equal(loadedFromToml.packs[0].id, "rtk-routing");
