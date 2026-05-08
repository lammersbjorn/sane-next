package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const saneBlockStart = "# sane-next:start"
const saneBlockEnd = "# sane-next:end"

func runCodex(args []string) (commandResult, error) {
	if len(args) == 0 {
		return commandResult{}, fmt.Errorf("codex requires subcommand: install, export, doctor, uninstall")
	}
	switch args[0] {
	case "export":
		return runCodexExport(args[1:])
	case "install":
		return runCodexInstall(args[1:])
	case "doctor":
		return runCodexDoctor(args[1:])
	case "uninstall":
		return runCodexUninstall(args[1:])
	default:
		return commandResult{}, fmt.Errorf("unknown codex subcommand %q", args[0])
	}
}

func runCodexExport(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("codex export", &flagOutput)
	configPath := fs.String("config", filepath.Join(defaultSourceRoot(), "pi-plugin", "config-schema.toml"), "Sane config path")
	codexHome := fs.String("codex-home", defaultCodexHome(), "Codex home directory")
	dryRun := fs.Bool("dry-run", false, "preview export paths without writing")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("codex export does not accept positional arguments")
	}
	if filepath.Base(filepath.Clean(*codexHome)) != ".codex" {
		return commandResult{}, fmt.Errorf("--codex-home must end in .codex for current skill export target")
	}
	exportArgs := []string{"--target", "codex", "--target-root", filepath.Dir(filepath.Clean(*codexHome)), "--config", *configPath}
	if *dryRun {
		exportArgs = append(exportArgs, "--dry-run")
	}
	return runExport(exportArgs)
}

func runCodexInstall(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("codex install", &flagOutput)
	root := fs.String("root", defaultRoot(), "Sane installation root")
	sourceRoot := fs.String("source-root", defaultSourceRoot(), "sane-next source root")
	codexHome := fs.String("codex-home", defaultCodexHome(), "Codex home directory")
	hooks := fs.String("hooks", "off", "hook mode: off, warn, or enforce")
	dryRun := fs.Bool("dry-run", false, "preview Codex-owned paths without writing")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("codex install does not accept positional arguments")
	}
	if *hooks != "off" && *hooks != "warn" && *hooks != "enforce" {
		return commandResult{}, fmt.Errorf("--hooks must be off, warn, or enforce")
	}
	if *dryRun {
		return commandResult{Message: fmt.Sprintf("dry-run codex install: export skills to %s; manage AGENTS.md block; hooks=%s", filepath.Join(*codexHome, "skills"), *hooks)}, nil
	}
	if err := os.MkdirAll(*codexHome, 0o755); err != nil {
		return commandResult{}, err
	}
	if _, err := runCodexExport([]string{"--codex-home", *codexHome, "--config", filepath.Join(*sourceRoot, "pi-plugin", "config-schema.toml")}); err != nil {
		return commandResult{}, err
	}
	if err := upsertManagedBlock(filepath.Join(*codexHome, "AGENTS.md"), codexAgentsBlock()); err != nil {
		return commandResult{}, err
	}
	if *hooks == "warn" || *hooks == "enforce" {
		if err := installCodexHooks(*root); err != nil {
			return commandResult{}, err
		}
		configPath := filepath.Join(*codexHome, "config.toml")
		if err := ensureTomlFeature(configPath, "hooks"); err != nil {
			return commandResult{}, err
		}
		if err := writeCodexHooksJSON(*codexHome, *root, *hooks); err != nil {
			return commandResult{}, err
		}
	}
	return commandResult{Message: fmt.Sprintf("installed Sane Core for Codex at %s; hooks=%s", *codexHome, *hooks)}, nil
}

func runCodexDoctor(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("codex doctor", &flagOutput)
	root := fs.String("root", defaultRoot(), "Sane installation root")
	codexHome := fs.String("codex-home", defaultCodexHome(), "Codex home directory")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	checks := []string{filepath.Join(*codexHome, "skills", "core-workflow", "SKILL.md"), filepath.Join(*codexHome, "AGENTS.md")}
	for _, path := range checks {
		if _, err := os.Stat(path); err != nil {
			return commandResult{}, fmt.Errorf("missing Codex Sane asset %s: %w", path, err)
		}
	}
	if _, err := os.Stat(filepath.Join(*codexHome, "hooks.json.sane-next-owned")); err == nil {
		for _, path := range []string{filepath.Join(*codexHome, "hooks.json"), filepath.Join(*root, "codex", "hooks", "sane-user-prompt-submit"), filepath.Join(*root, "codex", "hooks", "sane-pre-tool-use")} {
			if _, err := os.Stat(path); err != nil {
				return commandResult{}, fmt.Errorf("missing Codex Sane hook asset %s: %w", path, err)
			}
		}
		data, err := os.ReadFile(filepath.Join(*codexHome, "config.toml"))
		if err != nil {
			return commandResult{}, err
		}
		if !tomlFeatureEnabled(string(data), "hooks") {
			return commandResult{}, fmt.Errorf("missing Codex hooks feature flag in %s", filepath.Join(*codexHome, "config.toml"))
		}
	} else if err != nil && !os.IsNotExist(err) {
		return commandResult{}, err
	}
	return commandResult{Message: fmt.Sprintf("Sane Core for Codex is healthy at %s", *codexHome)}, nil
}

func runCodexUninstall(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("codex uninstall", &flagOutput)
	root := fs.String("root", defaultRoot(), "Sane installation root")
	codexHome := fs.String("codex-home", defaultCodexHome(), "Codex home directory")
	dryRun := fs.Bool("dry-run", false, "preview Codex cleanup without writing")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if *dryRun {
		return commandResult{Message: fmt.Sprintf("dry-run codex uninstall: remove managed blocks from %s and hook scripts under %s", *codexHome, filepath.Join(*root, "codex", "hooks"))}, nil
	}
	if err := removeManagedBlock(filepath.Join(*codexHome, "AGENTS.md")); err != nil {
		return commandResult{}, err
	}
	if err := removeTomlFeature(filepath.Join(*codexHome, "config.toml"), "hooks"); err != nil {
		return commandResult{}, err
	}
	if err := removeSaneHooksJSON(filepath.Join(*codexHome, "hooks.json")); err != nil {
		return commandResult{}, err
	}
	if err := removeCodexHookScripts(*root); err != nil {
		return commandResult{}, err
	}
	return commandResult{Message: fmt.Sprintf("removed Sane Core for Codex managed blocks from %s", *codexHome)}, nil
}

func defaultCodexHome() string {
	if home := os.Getenv("CODEX_HOME"); home != "" {
		return home
	}
	return filepath.Join(userHomeOrDot(), ".codex")
}

func codexAgentsBlock() string {
	return strings.Join([]string{saneBlockStart, "# Sane Core for Codex", "", "Use installed Sane skills when the task matches their descriptions; read the matching SKILL.md before relying on it.", "Start from current repo files and instructions, not memory. Keep work bounded, verify with local checks, and state what passed or was not run.", "For broad work, split research, implementation, review, or verification into lanes when available; keep one main thread accountable for final edits.", "Prefer executable checks and managed config over large prompt-only rules.", saneBlockEnd, ""}, "\n")
}

func writeCodexHooksJSON(codexHome, root, mode string) error {
	path := filepath.Join(codexHome, "hooks.json")
	marker := filepath.Join(codexHome, "hooks.json.sane-next-owned")
	if _, err := os.Stat(path); err == nil {
		if _, markerErr := os.Stat(marker); markerErr != nil {
			return fmt.Errorf("refusing to overwrite user-owned Codex hooks file: %s", path)
		}
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	hooksDir := filepath.ToSlash(filepath.Join(root, "codex", "hooks"))
	sessionStartCommand := inlineNodeHookCommand(codexSessionStartContext(), "hook session-start")
	body := fmt.Sprintf(`{
  "hooks": {
    "SessionStart": [
      {
        "matcher": "startup|resume",
        "hooks": [
          {
            "type": "command",
            "command": %q,
            "statusMessage": "Loading Sane session defaults"
          }
        ]
      }
    ],
    "UserPromptSubmit": [
      {
        "hooks": [
          {
            "type": "command",
            "command": %q,
            "statusMessage": "Checking Sane routing",
            "timeout": 5
          }
        ]
      }
    ],
    "PreToolUse": [
      {
        "matcher": "^Bash$",
        "hooks": [
          {
            "type": "command",
            "command": %q,
            "statusMessage": "Checking Sane shell policy",
            "timeout": 5
          }
        ]
      }
    ]
  }
}
`, sessionStartCommand, filepath.ToSlash(filepath.Join(hooksDir, "sane-user-prompt-submit")), fmt.Sprintf("%s --mode %s", filepath.ToSlash(filepath.Join(hooksDir, "sane-pre-tool-use")), mode))
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		return err
	}
	return os.WriteFile(marker, []byte("sane-next\n"), 0o644)
}

func removeSaneHooksJSON(path string) error {
	marker := path + ".sane-next-owned"
	if _, err := os.Stat(marker); err != nil {
		return nil
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.Remove(marker)
}

func codexSessionStartContext() string {
	return strings.Join([]string{
		"Sane Core for Codex is active.",
		"Use repo instructions first; when a Sane skill matches, read its SKILL.md before acting.",
		"For broad work, use agent-lanes and the active Codex lane/subagent surface when available; keep one main thread accountable.",
		"End with checks run or a clear reason they were skipped.",
	}, " ")
}

func inlineNodeHookCommand(additionalContext, marker string) string {
	payload := fmt.Sprintf(`{"hookSpecificOutput":{"hookEventName":"SessionStart","additionalContext":%q}}`, additionalContext)
	script := fmt.Sprintf("process.stdout.write(%q)", payload)
	return fmt.Sprintf("node -e %s # %s", shellQuote(script), marker)
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

func installCodexHooks(root string) error {
	dir := filepath.Join(root, "codex", "hooks")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	scripts := map[string]string{
		"sane-user-prompt-submit": `#!/bin/sh
# Intentionally quiet. SessionStart carries Sane's Codex context; keeping this
# hook silent avoids repeated visible "Hook context" blocks in Codex Desktop.
cat >/dev/null
exit 0
`,

		"sane-pre-tool-use": `#!/usr/bin/env python3
import argparse
import json
import re
import sys

parser = argparse.ArgumentParser()
parser.add_argument("--mode", choices=["warn", "enforce"], default="warn")
args = parser.parse_args()
try:
    payload = json.load(sys.stdin)
except Exception:
    sys.exit(0)
command = (((payload or {}).get("tool_input") or {}).get("command") or "").strip()
if not command:
    sys.exit(0)
blocked = [
    (r"^(cat|sed|awk|grep|rg|find|ls)\b", "Use Codex read/search tools or an approved compact route instead of raw shell for file discovery/read work."),
    (r"^(go test|npm test|pnpm test|yarn test|pytest|cargo test)\b", "Use the project test route or a compact test wrapper so failures stay readable."),
    (r"^(git diff|diff)\b", "Use Codex diff view or a compact diff route instead of raw shell diff."),
]
for pattern, reason in blocked:
    if re.search(pattern, command):
        message = f"Sane shell policy: {reason} Command: {command}"
        if args.mode == "enforce":
            print(json.dumps({"hookSpecificOutput": {"hookEventName": "PreToolUse", "permissionDecision": "deny", "permissionDecisionReason": message}}))
        else:
            print(json.dumps({"systemMessage": message}))
        sys.exit(0)
sys.exit(0)
`,
	}
	for name, body := range scripts {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o755); err != nil {
			return err
		}
	}
	return nil
}

func removeCodexHookScripts(root string) error {
	dir := filepath.Join(root, "codex", "hooks")
	for _, name := range []string{"sane-user-prompt-submit", "sane-pre-tool-use"} {
		if err := os.Remove(filepath.Join(dir, name)); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func ensureTomlFeature(path, name string) error {
	old := ""
	if data, err := os.ReadFile(path); err == nil {
		old = string(data)
	} else if !os.IsNotExist(err) {
		return err
	}
	updated := setTomlFeature(old, name, true)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return err
	}
	return os.WriteFile(path+".sane-next-hooks-owned", []byte("sane-next\n"), 0o644)
}

func removeTomlFeature(path, name string) error {
	marker := path + ".sane-next-hooks-owned"
	if _, err := os.Stat(marker); err != nil {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		line := name + " = true"
		var kept []string
		for _, raw := range strings.Split(string(data), "\n") {
			if strings.TrimSpace(raw) == line {
				continue
			}
			kept = append(kept, raw)
		}
		updated := strings.TrimRight(strings.Join(kept, "\n"), "\n")
		if updated != "" {
			updated += "\n"
		}
		if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
			return err
		}
	}
	return os.Remove(marker)
}

func tomlFeatureEnabled(data, name string) bool {
	line := name + " = true"
	for _, raw := range strings.Split(data, "\n") {
		if strings.TrimSpace(raw) == line {
			return true
		}
	}
	return false
}

func setTomlFeature(data, name string, value bool) string {
	line := fmt.Sprintf("%s = %t", name, value)
	lines := strings.Split(data, "\n")
	inFeatures := false
	for i, raw := range lines {
		trimmed := strings.TrimSpace(raw)
		if trimmed == "[features]" {
			inFeatures = true
			continue
		}
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			inFeatures = false
		}
		if inFeatures && strings.HasPrefix(trimmed, name+" ") && strings.Contains(trimmed, "=") {
			lines[i] = line
			updated := strings.TrimRight(strings.Join(lines, "\n"), "\n")
			if updated != "" {
				updated += "\n"
			}
			return updated
		}
	}
	features := strings.Index(data, "[features]")
	if features == -1 {
		prefix := "[features]\n" + line + "\n\n"
		if strings.TrimSpace(data) == "" {
			return prefix
		}
		return prefix + data
	}
	insert := features + len("[features]")
	if insert < len(data) && data[insert] == '\n' {
		insert++
	}
	return data[:insert] + line + "\n" + data[insert:]
}

func upsertManagedBlock(path, block string) error {
	old := ""
	if data, err := os.ReadFile(path); err == nil {
		old = string(data)
	} else if !os.IsNotExist(err) {
		return err
	}
	updated, err := replaceManagedBlock(old, block)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(updated), 0o644)
}

func removeManagedBlock(path string) error {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	updated, err := replaceManagedBlock(string(data), "")
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(updated), 0o644)
}

func replaceManagedBlock(old, block string) (string, error) {
	start := strings.Index(old, saneBlockStart)
	end := strings.Index(old, saneBlockEnd)
	if start == -1 && end == -1 {
		if strings.TrimSpace(old) == "" {
			return block, nil
		}
		if block == "" {
			return old, nil
		}
		return strings.TrimRight(old, "\n") + "\n\n" + block, nil
	}
	if start == -1 || end == -1 || end < start {
		return "", fmt.Errorf("malformed Sane managed block")
	}
	end += len(saneBlockEnd)
	if end < len(old) && old[end] == '\n' {
		end++
	}
	updated := old[:start] + block + old[end:]
	return strings.TrimLeft(updated, "\n"), nil
}
