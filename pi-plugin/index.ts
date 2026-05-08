import { existsSync } from "node:fs";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { createRequire } from "node:module";
import type { ExtensionAPI } from "@mariozechner/pi-coding-agent";
import { isToolCallEventType } from "@mariozechner/pi-coding-agent";

const require = createRequire(import.meta.url);
const pluginPath = require.resolve("./plugin.js");
delete require.cache[pluginPath];
const plugin = require(pluginPath) as {
  buildRelevantLedgerContext: (entries: unknown[], prompt: string) => string;
  buildQuietStatusSummary: (config: any, state?: any) => string;
  buildSubagentRoutingHint: (config: any, prompt: string) => string;
  buildWebResearchHint: (config: any, prompt: string) => string;
  applyPrettyEnvironmentDefaults: (config: any, env?: any) => any;
  commandRequiresRtk: (command: string) => boolean;
  extractAssistantProgress: (messages: unknown[]) => string | null;
  getLedgerEntries: (sessionManager: unknown) => unknown[];
  isRtkRoutingEnabled: (config: any) => boolean;
  getRtkRoutingMode: (config: any) => string;
  isRtkRoutingEnforced: (config: any) => boolean;
  LEDGER_ENTRY_TYPE: string;
  loadSaneConfig: (configPath: string) => any;
  makeLedgerEntry: (kind: string, data?: Record<string, unknown>) => unknown;
  parseGoalCommand: (args: string) => { action: string; value: string };
  summarizeGoalState: (entries: unknown[]) => { activeGoal?: { text?: string } };
};
const { applyPrettyEnvironmentDefaults, buildRelevantLedgerContext, buildQuietStatusSummary, buildSubagentRoutingHint, buildWebResearchHint, commandRequiresRtk, extractAssistantProgress, getLedgerEntries, getRtkRoutingMode, isRtkRoutingEnabled, isRtkRoutingEnforced, LEDGER_ENTRY_TYPE, loadSaneConfig, makeLedgerEntry, parseGoalCommand, summarizeGoalState } = plugin;

const baseDir = dirname(fileURLToPath(import.meta.url));
const configPath = join(baseDir, "config-schema.toml");

function enabledPiSkillPaths(cwd?: string) {
  const cfg = loadSaneConfig(configPath);
  const packs = [...cfg.packs, ...cfg.userPacks]
    .filter((pack) => pack.enabled && pack.targets.includes("pi"));
  return packs
    .filter((pack) => !projectProvidesSkill(cwd, pack.id))
    .map((pack) => resolve(baseDir, pack.source ?? join("..", "packs", pack.id)));
}

function projectProvidesSkill(cwd: string | undefined, id: string) {
  if (!cwd) return false;
  return existsSync(resolve(cwd, ".agents", "skills", id, "SKILL.md"));
}

export default function (pi: ExtensionAPI) {
  applyPrettyEnvironmentDefaults(loadSaneConfig(configPath));

  pi.on("resources_discover", (event) => ({
    skillPaths: enabledPiSkillPaths(event.cwd),
  }));

  pi.on("before_agent_start", async (event, ctx) => {
    const cfg = loadSaneConfig(configPath);
    const additions: string[] = [];
    if (cfg.defaults.responseStyle === "caveman") {
      additions.push("Sane response style: answer in caveman speak. Use short primitive phrasing, simple words, and terse sentences unless exact code, file paths, commands, or quoted text require normal spelling.");
    }

    const webHint = buildWebResearchHint(cfg, event.prompt ?? "");
    if (webHint) additions.push(webHint);

    const subagentHint = buildSubagentRoutingHint(cfg, event.prompt ?? "");
    if (subagentHint) additions.push(subagentHint);

    const ledgerContext = buildRelevantLedgerContext(getLedgerEntries(ctx.sessionManager), event.prompt ?? "");
    if (ledgerContext) additions.push(ledgerContext);
    if (additions.length === 0) return;

    return {
      systemPrompt: `${event.systemPrompt}\n\n${additions.join("\n\n")}`,
    };
  });

  pi.on("session_start", async (_event, ctx) => {
    const cfg = loadSaneConfig(configPath);
    if (getRtkRoutingMode(cfg) !== "off") {
      const rtk = await pi.exec("sh", ["-lc", "command -v rtk"], { timeout: 2000 });
      if (rtk.exitCode !== 0) {
        ctx.ui.notify("Sane RTK routing is enabled, but rtk is not on PATH", "warning");
      }
    }

    const state = summarizeGoalState(getLedgerEntries(ctx.sessionManager));
    ctx.ui.setStatus?.("sane-next", buildQuietStatusSummary(cfg, state));
  });

  pi.on("tool_call", async (event) => {
    const cfg = loadSaneConfig(configPath);
    if (!isRtkRoutingEnforced(cfg)) return;
    if (!isToolCallEventType("bash", event)) return;

    const command = event.input.command ?? "";
    if (commandRequiresRtk(command)) {
      return {
        block: true,
        reason: rtkBlockReason(command),
      };
    }
  });

  pi.on("user_bash", (event) => {
    const cfg = loadSaneConfig(configPath);
    if (!isRtkRoutingEnforced(cfg) || !commandRequiresRtk(event.command)) return;

    return {
      result: {
        output: rtkBlockReason(event.command),
        exitCode: 1,
        cancelled: false,
        truncated: false,
      },
    };
  });

  pi.on("agent_end", async (event, ctx) => {
    const progress = extractAssistantProgress(event.messages ?? []);
    if (!progress) return;
    try {
      pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("progress", {
        text: progress,
        source: "assistant-turn",
        confidence: "observed",
        evidence: [`session:${ctx.sessionManager.getSessionFile?.() ?? "current"}`],
      }));
    } catch {
      // Pi can invalidate extension actions while print-mode sessions exit. Progress capture is best effort.
    }
  });

  pi.registerCommand("sane-status", {
    description: "Show Sane overlay status, defaults, discovered pack count, and active goal",
    handler: async (_args, ctx) => {
      const cfg = loadSaneConfig(configPath);
      const packCount = [...cfg.packs, ...cfg.userPacks].filter((pack) => pack.enabled).length;
      const rtkStatus = `, rtk=${getRtkRoutingMode(cfg)}`;
      const styleStatus = cfg.defaults.responseStyle ? `, style=${cfg.defaults.responseStyle}` : "";
      const { activeGoal } = summarizeGoalState(getLedgerEntries(ctx.sessionManager));
      const goalStatus = activeGoal ? `, goal=${activeGoal.text}` : ", goal=none";
      const subagentStatus = buildQuietStatusSummary(cfg, { activeGoal }).match(/subagents=[^,]+/)?.[0] ?? "subagents=off";
      ctx.ui.notify(`Sane overlay loaded ${packCount} pack(s), model=${cfg.defaults.model}, reasoning=${cfg.defaults.reasoning}${styleStatus}${rtkStatus}, ${subagentStatus}${goalStatus}. For broad independent work, use agent-lanes with pi-subagents (/subagents-status, /subagents-doctor, or /parallel when available).`, "info");
    },
  });

  pi.registerCommand("sane-lanes", {
    description: "Ask Sane to split broad independent work into pi-subagents agent lanes when useful",
    handler: async (args, _ctx) => {
      const objective = args.trim() || "the current task";
      await pi.sendUserMessage(`Use Sane agent-lanes for ${objective}: if the work has independent research, review, verification, or disjoint implementation lanes and pi-subagents is available, launch focused subagents with clear boundaries; otherwise explain why main-thread work is better.`);
    },
  });

  pi.registerCommand("sane-goal", {
    description: "Manage Sane goal/ledger state: set <goal>, decide <decision>, block <reason>, status, run, handoff, clear",
    handler: async (args, ctx) => {
      const parsed = parseGoalCommand(args);
      if (parsed.action === "set") {
        if (!parsed.value) return ctx.ui.notify("Usage: /sane-goal set <goal>", "warning");
        pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("goal", { text: parsed.value, source: "user-command", confidence: "explicit" }));
        return ctx.ui.notify(`Sane goal set: ${parsed.value}`, "success");
      }
      if (parsed.action === "decide") {
        if (!parsed.value) return ctx.ui.notify("Usage: /sane-goal decide <decision>", "warning");
        pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("decision", { text: parsed.value, source: "user-command", confidence: "explicit", status: "accepted" }));
        return ctx.ui.notify(`Sane decision recorded: ${parsed.value}`, "success");
      }
      if (parsed.action === "block") {
        if (!parsed.value) return ctx.ui.notify("Usage: /sane-goal block <reason>", "warning");
        pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("blocker", { text: parsed.value, source: "user-command", confidence: "explicit" }));
        return ctx.ui.notify(`Sane blocker recorded: ${parsed.value}`, "warning");
      }
      if (parsed.action === "clear") {
        pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("goal", { text: "Goal cleared by user.", status: "done", source: "user-command", confidence: "explicit" }));
        return ctx.ui.notify("Sane goal cleared", "info");
      }

      const state = summarizeGoalState(getLedgerEntries(ctx.sessionManager));
      if (parsed.action === "run") {
        if (!state.activeGoal) return ctx.ui.notify("No active Sane goal. Use /sane-goal set <goal> first.", "warning");
        await pi.sendUserMessage(`Continue working toward this Sane goal until it is done, blocked, unsafe, or needs user approval. Goal: ${state.activeGoal.text}`);
        return;
      }
      if (parsed.action === "handoff") {
        const handoff = buildGoalStatusText(state);
        pi.appendEntry(LEDGER_ENTRY_TYPE, makeLedgerEntry("handoff", { text: handoff, source: "sane-next", confidence: "observed" }));
        return ctx.ui.notify(handoff, "info");
      }
      ctx.ui.notify(buildGoalStatusText(state), "info");
    },
  });
}

function buildGoalStatusText(state: { activeGoal?: { text?: string }, decisions?: Array<{ text?: string }>, progress?: Array<{ text?: string }>, blockers?: Array<{ text?: string }> }) {
  const lines = ["Sane goal status:"];
  lines.push(`Goal: ${state.activeGoal?.text ?? "none"}`);
  if (state.blockers?.length) lines.push(`Blockers: ${state.blockers.map((entry) => entry.text).join("; ")}`);
  if (state.decisions?.length) lines.push(`Recent decisions: ${state.decisions.slice(-3).map((entry) => entry.text).join("; ")}`);
  if (state.progress?.length) lines.push(`Recent progress: ${state.progress.slice(-3).map((entry) => entry.text).join("; ")}`);
  return lines.join("\n");
}

function rtkBlockReason(command: string) {
  return `Sane RTK policy blocked raw shell command: ${command}. Use an RTK route such as rtk grep, rtk read, rtk find, rtk ls, rtk git, rtk test, rtk lint, rtk diff, or rtk run when exact shell semantics are required.`;
}
