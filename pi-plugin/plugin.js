"use strict";

const fs = require("fs");

function loadSaneConfig(configPath, io = fs) {
  const raw = io.readFileSync(configPath, "utf8");
  const parsed = parseSaneToml(raw);
  return validateSaneConfig(parsed);
}

function parseSaneToml(raw) {
  const config = { defaults: {}, rtk: {}, pretty: {}, packs: [], userPacks: [], exportTargets: [], recommendedPiPackages: [] };
  let current = null;

  for (const originalLine of raw.split(/\r?\n/)) {
    const line = originalLine.trim();
    if (!line || line.startsWith("#")) continue;

    if (line === "[defaults]") {
      current = config.defaults;
      continue;
    }
    if (line === "[rtk]") {
      current = config.rtk;
      continue;
    }
    if (line === "[pretty]") {
      current = config.pretty;
      continue;
    }
    if (line === "[ownership]") {
      current = {};
      continue;
    }
    if (line === "[[packs]]") {
      current = {};
      config.packs.push(current);
      continue;
    }
    if (line === "[[user_packs]]") {
      current = {};
      config.userPacks.push(current);
      continue;
    }
    if (line === "[[export_targets]]") {
      current = {};
      config.exportTargets.push(current);
      continue;
    }
    if (line === "[[recommended_pi_packages]]") {
      current = {};
      config.recommendedPiPackages.push(current);
      continue;
    }
    if (!current) continue;

    const match = line.match(/^([A-Za-z0-9_-]+)\s*=\s*(.+)$/);
    if (!match) {
      throw new Error(`Unsupported Sane config line: ${line}`);
    }
    const key = match[1].replace(/_([a-z])/g, (_, char) => char.toUpperCase());
    current[key] = parseTomlValue(match[2]);
  }

  return config;
}

function parseTomlValue(value) {
  const trimmed = value.trim();
  if (trimmed === "true") return true;
  if (trimmed === "false") return false;
  if (trimmed.startsWith('"') && trimmed.endsWith('"')) {
    return trimmed.slice(1, -1);
  }
  if (trimmed.startsWith("[") && trimmed.endsWith("]")) {
    const body = trimmed.slice(1, -1).trim();
    if (!body) return [];
    return body.split(",").map((item) => parseTomlValue(item.trim()));
  }
  const numeric = Number(trimmed);
  if (Number.isFinite(numeric)) return numeric;
  throw new Error(`Unsupported Sane config value: ${value}`);
}

function validateSaneConfig(config) {
  if (!config || typeof config !== "object" || Array.isArray(config)) {
    throw new Error("Sane config must be an object");
  }
  if (!config.defaults || typeof config.defaults.model !== "string" || typeof config.defaults.reasoning !== "string") {
    throw new Error("Sane config requires defaults.model and defaults.reasoning");
  }
  if (config.defaults.responseStyle !== undefined && typeof config.defaults.responseStyle !== "string") {
    throw new Error("Sane config defaults.response_style must be a string when set");
  }
  if (!Array.isArray(config.packs)) {
    throw new Error("Sane config requires packs");
  }
  if (!Array.isArray(config.exportTargets)) {
    throw new Error("Sane config requires exportTargets");
  }

  const rtkMode = config.rtk?.mode ?? "off";
  if (!["off", "advise", "warn", "enforce"].includes(rtkMode)) {
    throw new Error("Sane config rtk.mode must be one of off, advise, warn, or enforce");
  }

  return {
    defaults: {
      model: config.defaults.model,
      reasoning: config.defaults.reasoning,
      responseStyle: config.defaults.responseStyle,
    },
    rtk: {
      mode: rtkMode,
    },
    pretty: validatePrettyConfig(config.pretty || {}),
    packs: config.packs.map(validatePack),
    userPacks: (config.userPacks || []).map(validatePack),
    exportTargets: config.exportTargets.map(validateExportTarget),
    recommendedPiPackages: (config.recommendedPiPackages || []).map(validateRecommendedPiPackage),
  };
}

function validatePrettyConfig(pretty) {
  const maxPreviewLines = pretty.maxPreviewLines ?? 24;
  const maxHlChars = pretty.maxHlChars ?? 1;
  const icons = pretty.icons ?? "nerd";
  if (!Number.isInteger(maxPreviewLines) || maxPreviewLines <= 0) {
    throw new Error("Sane config pretty.max_preview_lines must be a positive integer");
  }
  if (!Number.isInteger(maxHlChars) || maxHlChars <= 0) {
    throw new Error("Sane config pretty.max_hl_chars must be a positive integer");
  }
  if (typeof icons !== "string") {
    throw new Error("Sane config pretty.icons must be a string");
  }
  return { maxPreviewLines, maxHlChars, icons };
}

function applyPrettyEnvironmentDefaults(config, env = process.env) {
  const pretty = validatePrettyConfig(config?.pretty || {});
  if (env.PRETTY_MAX_PREVIEW_LINES === undefined) env.PRETTY_MAX_PREVIEW_LINES = String(pretty.maxPreviewLines);
  if (env.PRETTY_MAX_HL_CHARS === undefined) env.PRETTY_MAX_HL_CHARS = String(pretty.maxHlChars);
  if (env.PRETTY_ICONS === undefined) env.PRETTY_ICONS = pretty.icons;
  return {
    PRETTY_MAX_PREVIEW_LINES: env.PRETTY_MAX_PREVIEW_LINES,
    PRETTY_MAX_HL_CHARS: env.PRETTY_MAX_HL_CHARS,
    PRETTY_ICONS: env.PRETTY_ICONS,
  };
}

function validatePack(pack) {
  if (!pack || typeof pack.id !== "string" || typeof pack.enabled !== "boolean" || !Array.isArray(pack.targets)) {
    throw new Error("Sane pack requires id, enabled, and targets");
  }
  return {
    id: pack.id,
    enabled: pack.enabled,
    source: typeof pack.source === "string" ? pack.source : undefined,
    targets: [...pack.targets],
  };
}

function validateExportTarget(target) {
  if (!target || typeof target.id !== "string" || typeof target.kind !== "string" || typeof target.path !== "string") {
    throw new Error("Sane export target requires id, kind, and path");
  }
  return {
    id: target.id,
    kind: target.kind,
    path: target.path,
  };
}

function validateRecommendedPiPackage(pkg) {
  if (!pkg || typeof pkg.id !== "string" || typeof pkg.enabled !== "boolean" || typeof pkg.defaultInstall !== "boolean" || typeof pkg.package !== "string" || typeof pkg.purpose !== "string") {
    throw new Error("Sane recommended Pi package requires id, enabled, default_install, package, and purpose");
  }
  return {
    id: pkg.id,
    enabled: pkg.enabled,
    defaultInstall: pkg.defaultInstall,
    package: pkg.package,
    purpose: pkg.purpose,
  };
}

function isRtkRoutingEnabled(config) {
  return [...(config.packs || []), ...(config.userPacks || [])]
    .some((pack) => pack.id === "rtk-routing" && pack.enabled === true && (pack.targets || []).includes("pi"));
}

function getRtkRoutingMode(config) {
  if (!isRtkRoutingEnabled(config)) return "off";
  return config.rtk?.mode ?? "off";
}

function isRtkRoutingEnforced(config) {
  return getRtkRoutingMode(config) === "enforce";
}

function commandRequiresRtk(command) {
  const normalized = stripShellPrefixes(command || "");
  if (!normalized) return false;
  if (/^(?:rtk|command\s+-v\s+rtk|which\s+rtk)\b/.test(normalized)) return false;

  return /^(?:grep|rg|find|ls|tree|cat|sed|awk|git\s+(?:status|diff|log|show)|go\s+test|npm\s+test|pnpm\s+(?:test|lint)|yarn\s+(?:test|lint))\b/.test(normalized);
}

const LEDGER_ENTRY_TYPE = "sane-ledger";

function parseGoalCommand(args) {
  const text = (args || "").trim();
  if (!text) return { action: "status", value: "" };
  const [action, ...rest] = text.split(/\s+/);
  return { action: action.toLowerCase(), value: rest.join(" ").trim() };
}

function makeLedgerEntry(kind, data = {}) {
  return {
    schema: 1,
    kind,
    status: data.status || "active",
    text: data.text || "",
    source: data.source || "sane-next",
    scope: data.scope || "session",
    confidence: data.confidence || "explicit",
    evidence: data.evidence || [],
    timestamp: data.timestamp || new Date().toISOString(),
  };
}

function getLedgerEntries(sessionManager) {
  const entries = typeof sessionManager.getEntries === "function" ? sessionManager.getEntries() : [];
  return entries
    .filter((entry) => entry && entry.type === "custom" && entry.customType === LEDGER_ENTRY_TYPE && entry.data)
    .map((entry) => entry.data)
    .filter((data) => data.schema === 1 && typeof data.kind === "string");
}

function summarizeGoalState(entries) {
  const goals = entries.filter((entry) => entry.kind === "goal");
  const activeGoal = [...goals].reverse().find((entry) => entry.status === "active");
  const decisions = entries.filter((entry) => entry.kind === "decision" && entry.status !== "superseded");
  const progress = entries.filter((entry) => entry.kind === "progress").slice(-5);
  const blockers = entries.filter((entry) => entry.kind === "blocker" && entry.status !== "resolved");
  return { activeGoal, decisions, progress, blockers };
}

function buildRelevantLedgerContext(entries, prompt, options = {}) {
  const maxAgeDays = Number.isFinite(options.maxAgeDays) ? options.maxAgeDays : 30;
  const now = options.now ? new Date(options.now) : new Date();
  const freshEntries = entries.filter((entry) => !isStaleLedgerEntry(entry, now, maxAgeDays));
  const { activeGoal, decisions, progress, blockers } = summarizeGoalState(freshEntries);
  const selected = [];
  if (activeGoal) selected.push(activeGoal);
  selected.push(...blockers.slice(-2));
  selected.push(...decisions.filter((entry) => isRelevant(entry.text, prompt)).slice(-5));
  selected.push(...progress.slice(-3));

  if (selected.length === 0) return "";
  const lines = [
    "Relevant Sane ledger entries. Use as sourced context, not immutable truth. Prefer current repo files and direct user instructions if they conflict. Treat entries labeled low confidence or conflict as hints only.",
  ];
  for (const entry of selected) {
    const label = `${entry.kind}${entry.status ? `/${entry.status}` : ""}${entry.confidence ? `/${entry.confidence}` : ""}${hasLedgerConflict(entry, selected) ? "/conflict" : ""}`;
    const evidence = Array.isArray(entry.evidence) && entry.evidence.length > 0 ? ` Evidence: ${entry.evidence.join("; ")}.` : "";
    lines.push(`- ${label}: ${entry.text}${evidence}`);
  }
  return lines.join("\n");
}

function isStaleLedgerEntry(entry, now, maxAgeDays) {
  if (!entry || !entry.timestamp || entry.kind === "goal" || entry.kind === "blocker") return false;
  const timestamp = new Date(entry.timestamp);
  if (Number.isNaN(timestamp.getTime())) return false;
  return now.getTime() - timestamp.getTime() > maxAgeDays * 24 * 60 * 60 * 1000;
}

function hasLedgerConflict(entry, selected) {
  if (!entry || entry.kind !== "decision") return false;
  const normalized = normalizeDecisionText(entry.text);
  return selected.some((other) => {
    if (other === entry || other.kind !== "decision") return false;
    const otherText = String(other.text || "").toLowerCase();
    return /\b(?:do not|don't|not)\b/.test(otherText) && normalizeDecisionText(other.text) === normalized;
  });
}

function normalizeDecisionText(text) {
  return String(text || "")
    .toLowerCase()
    .replace(/\b(?:do not|don't|not)\s+/g, "")
    .replace(/\s+/g, " ")
    .trim();
}

function isRelevant(text, prompt) {
  if (!text || !prompt) return true;
  const words = new Set(String(prompt).toLowerCase().match(/[a-z0-9][a-z0-9-]{3,}/g) || []);
  return (String(text).toLowerCase().match(/[a-z0-9][a-z0-9-]{3,}/g) || []).some((word) => words.has(word));
}

function extractAssistantProgress(messages) {
  const assistantText = messages
    .filter((message) => message && message.role === "assistant")
    .flatMap((message) => extractTextBlocks(message.content))
    .join("\n")
    .trim();
  if (!assistantText) return null;
  const firstLine = assistantText.split(/\r?\n/).map((line) => line.trim()).find(Boolean) || "Assistant completed a turn.";
  return firstLine.length > 240 ? `${firstLine.slice(0, 237)}...` : firstLine;
}

function buildWebResearchHint(config, prompt) {
  const hasWebPackage = (config.recommendedPiPackages || []).some((pkg) => pkg.enabled && pkg.id === "pi-web-providers");
  if (!hasWebPackage || !/\b(?:web|research|current|latest|today|status|docs?)\b/i.test(prompt || "")) return "";
  return "Sane web hint: for current docs, package status, or freshness-sensitive claims, use available web research tools and cite sources briefly.";
}

function isSubagentLanesEnabled(config) {
  return [...(config.packs || []), ...(config.userPacks || [])]
    .some((pack) => pack.id === "agent-lanes" && pack.enabled === true && (pack.targets || []).includes("pi"));
}

function isSubagentsConfigured(config) {
  const hasPackage = (config.recommendedPiPackages || [])
    .some((pkg) => pkg.enabled && pkg.defaultInstall && pkg.id === "pi-subagents");
  return hasPackage && isSubagentLanesEnabled(config);
}

function buildSubagentRoutingHint(config, prompt) {
  if (!isSubagentsConfigured(config)) return "";
  if (!/\b(?:sub-?agents?|agents?|lanes?|parallel|delegate|delegation|broad|research|review|verify|verification|implement|refactor|fix\s+this\s+properly)\b/i.test(prompt || "")) return "";
  return "Sane subagent hint: pi-subagents is a default Sane runtime package and agent-lanes is enabled. For broad independent research, review, verification, or disjoint implementation work, delegate lanes with the available subagent tool/commands; keep small or tightly coupled work in the main thread.";
}

function buildQuietStatusSummary(config, state = {}) {
  const packs = [...(config.packs || []), ...(config.userPacks || [])].filter((pack) => pack.enabled).length;
  const subagents = state.subagentsAvailable ? "subagents=ready" : (isSubagentsConfigured(config) ? "subagents=configured" : "subagents=off");
  const web = (config.recommendedPiPackages || []).some((pkg) => pkg.enabled && pkg.id === "pi-web-providers") ? "web=ready" : "web=off";
  const goal = state.activeGoal ? "goal=active" : "goal=none";
  return `Sane: packs=${packs}, ${subagents}, ${web}, ${goal}`;
}

function extractTextBlocks(content) {
  if (typeof content === "string") return [content];
  if (!Array.isArray(content)) return [];
  return content
    .filter((block) => block && typeof block === "object" && block.type === "text" && typeof block.text === "string")
    .map((block) => block.text);
}

function stripShellPrefixes(command) {
  let normalized = command.trim();
  normalized = normalized.replace(/^(?:cd\s+[^;&|]+\s*&&\s*)+/, "").trim();
  normalized = normalized.replace(/^(?:env\s+(?:[A-Za-z_][A-Za-z0-9_]*=(?:'[^']*'|\"[^\"]*\"|\S+)\s*)+)/, "").trim();
  return normalized;
}

module.exports = {
  loadSaneConfig,
  parseSaneToml,
  validateSaneConfig,
  isRtkRoutingEnabled,
  getRtkRoutingMode,
  isRtkRoutingEnforced,
  commandRequiresRtk,
  LEDGER_ENTRY_TYPE,
  parseGoalCommand,
  makeLedgerEntry,
  getLedgerEntries,
  summarizeGoalState,
  buildRelevantLedgerContext,
  isStaleLedgerEntry,
  hasLedgerConflict,
  extractAssistantProgress,
  buildWebResearchHint,
  buildSubagentRoutingHint,
  buildQuietStatusSummary,
  isSubagentLanesEnabled,
  isSubagentsConfigured,
  applyPrettyEnvironmentDefaults,
};
