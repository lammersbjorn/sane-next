"use strict";

const fs = require("fs");

function loadSaneConfig(configPath, io = fs) {
  const raw = io.readFileSync(configPath, "utf8");
  const parsed = parseSaneToml(raw);
  return validateSaneConfig(parsed);
}

function parseSaneToml(raw) {
  const config = { defaults: {}, packs: [], exportTargets: [] };
  let current = null;

  for (const originalLine of raw.split(/\r?\n/)) {
    const line = originalLine.trim();
    if (!line || line.startsWith("#")) continue;

    if (line === "[defaults]") {
      current = config.defaults;
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
    if (line === "[[export_targets]]") {
      current = {};
      config.exportTargets.push(current);
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
  if (!Array.isArray(config.packs)) {
    throw new Error("Sane config requires packs");
  }
  if (!Array.isArray(config.exportTargets)) {
    throw new Error("Sane config requires exportTargets");
  }

  return {
    defaults: {
      model: config.defaults.model,
      reasoning: config.defaults.reasoning,
    },
    packs: config.packs.map(validatePack),
    exportTargets: config.exportTargets.map(validateExportTarget),
  };
}

function validatePack(pack) {
  if (!pack || typeof pack.id !== "string" || typeof pack.enabled !== "boolean" || !Array.isArray(pack.targets)) {
    throw new Error("Sane pack requires id, enabled, and targets");
  }
  return {
    id: pack.id,
    enabled: pack.enabled,
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

module.exports = {
  loadSaneConfig,
  parseSaneToml,
  validateSaneConfig,
};
