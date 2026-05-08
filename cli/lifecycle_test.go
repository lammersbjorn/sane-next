package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLifecyclePreservesUserPacksOnUninstall(t *testing.T) {
	root := t.TempDir()
	if _, err := runInstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	userPack := filepath.Join(root, "user-packs", "mine", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(userPack), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(userPack, []byte("user owned\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := runUpdate([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "packs", "VERSION")); err != nil {
		t.Fatalf("expected owned pack version: %v", err)
	}

	if _, err := runUninstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(userPack); err != nil {
		t.Fatalf("expected user pack to remain: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "packs")); !os.IsNotExist(err) {
		t.Fatalf("expected owned packs removed, got %v", err)
	}
}

func TestCodexInstallDoctorAndUninstall(t *testing.T) {
	root := filepath.Join(t.TempDir(), "sane")
	codexHome := filepath.Join(t.TempDir(), ".codex")
	if _, err := runCodex([]string{"install", "--root", root, "--codex-home", codexHome, "--source-root", "..", "--hooks", "enforce"}); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{
		filepath.Join(codexHome, "skills", "core-workflow", "SKILL.md"),
		filepath.Join(codexHome, "AGENTS.md"),
		filepath.Join(codexHome, "config.toml"),
		filepath.Join(codexHome, "config.toml.sane-next-hooks-owned"),
		filepath.Join(root, "codex", "hooks", "sane-pre-tool-use"),
		filepath.Join(root, "codex", "hooks", "sane-user-prompt-submit"),
	} {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected Codex asset %s: %v", path, err)
		}
	}
	data, err := os.ReadFile(filepath.Join(codexHome, "config.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "hooks = true") || strings.Contains(string(data), "codex_hooks") {
		t.Fatalf("expected current hook feature flag only, got:\n%s", data)
	}
	hooksJSON, err := os.ReadFile(filepath.Join(codexHome, "hooks.json"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(hooksJSON), "sane-next") || !strings.Contains(string(hooksJSON), "PreToolUse") {
		t.Fatalf("expected schema-clean Sane hooks.json, got:\n%s", hooksJSON)
	}
	if _, err := os.Stat(filepath.Join(codexHome, "hooks.json.sane-next-owned")); err != nil {
		t.Fatalf("expected hooks ownership marker: %v", err)
	}
	if _, err := runCodex([]string{"doctor", "--root", root, "--codex-home", codexHome}); err != nil {
		t.Fatal(err)
	}
	if _, err := runCodex([]string{"uninstall", "--root", root, "--codex-home", codexHome}); err != nil {
		t.Fatal(err)
	}
	data, err = os.ReadFile(filepath.Join(codexHome, "config.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), saneBlockStart) || strings.Contains(string(data), "hooks = true") {
		t.Fatalf("expected managed Codex config removed, got:\n%s", data)
	}
	for _, path := range []string{filepath.Join(root, "codex", "hooks", "sane-pre-tool-use"), filepath.Join(root, "codex", "hooks", "sane-user-prompt-submit")} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("expected Sane hook script removed %s, got: %v", path, err)
		}
	}
	if _, err := os.Stat(filepath.Join(codexHome, "config.toml.sane-next-hooks-owned")); !os.IsNotExist(err) {
		t.Fatalf("expected config ownership marker removed, got: %v", err)
	}
}

func TestCodexUninstallPreservesUserHookFiles(t *testing.T) {
	root := filepath.Join(t.TempDir(), "sane")
	codexHome := filepath.Join(t.TempDir(), ".codex")
	if _, err := runCodex([]string{"install", "--root", root, "--codex-home", codexHome, "--source-root", "..", "--hooks", "warn"}); err != nil {
		t.Fatal(err)
	}
	userHook := filepath.Join(root, "codex", "hooks", "user-hook")
	if err := os.WriteFile(userHook, []byte("user owned\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := runCodex([]string{"uninstall", "--root", root, "--codex-home", codexHome}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(userHook); err != nil {
		t.Fatalf("expected user hook preserved: %v", err)
	}
}

func TestCodexInstallReplacesExistingHooksFeatureFlag(t *testing.T) {
	root := filepath.Join(t.TempDir(), "sane")
	codexHome := filepath.Join(t.TempDir(), ".codex")
	if err := os.MkdirAll(codexHome, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(codexHome, "config.toml"), []byte("[features]\nhooks = false\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runCodex([]string{"install", "--root", root, "--codex-home", codexHome, "--source-root", "..", "--hooks", "warn"}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(codexHome, "config.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "hooks = false") || strings.Count(string(data), "hooks = true") != 1 {
		t.Fatalf("expected one enabled hooks feature flag, got:\n%s", data)
	}
}

func TestCodexInstallWithoutHooksDoesNotWriteHookAssets(t *testing.T) {
	root := filepath.Join(t.TempDir(), "sane")
	codexHome := filepath.Join(t.TempDir(), ".codex")
	if _, err := runCodex([]string{"install", "--root", root, "--codex-home", codexHome, "--source-root", ".."}); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{filepath.Join(codexHome, "hooks.json"), filepath.Join(codexHome, "config.toml.sane-next-hooks-owned"), filepath.Join(root, "codex", "hooks")} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("expected no hook asset %s, got %v", path, err)
		}
	}
	if _, err := runCodex([]string{"doctor", "--root", root, "--codex-home", codexHome}); err != nil {
		t.Fatal(err)
	}
}

func TestCodexDoctorReportsBrokenManagedHooks(t *testing.T) {
	root := filepath.Join(t.TempDir(), "sane")
	codexHome := filepath.Join(t.TempDir(), ".codex")
	if _, err := runCodex([]string{"install", "--root", root, "--codex-home", codexHome, "--source-root", "..", "--hooks", "warn"}); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(codexHome, "hooks.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := runCodex([]string{"doctor", "--root", root, "--codex-home", codexHome}); err == nil {
		t.Fatal("expected Codex doctor to fail with missing managed hooks.json")
	}
}

func TestDoctorRequiresOwnershipMarker(t *testing.T) {
	root := t.TempDir()
	if _, err := runDoctor([]string{"--root", root}); err == nil {
		t.Fatal("expected doctor to fail without ownership marker")
	}
}

func TestInstallCopiesAndRepairRestoresOwnedAssets(t *testing.T) {
	root := t.TempDir()
	if _, err := runInstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "pi-plugin")); !os.IsNotExist(err) {
		t.Fatalf("expected legacy pi-plugin copy to be absent, got %v", err)
	}
	plugin := filepath.Join(root, "extensions", "sane-next", "index.ts")
	if _, err := os.Stat(plugin); err != nil {
		t.Fatalf("expected extension plugin copied: %v", err)
	}
	if err := os.Remove(plugin); err != nil {
		t.Fatal(err)
	}
	if _, err := runDoctor([]string{"--root", root}); err == nil {
		t.Fatal("expected doctor to fail with missing plugin")
	}
	if _, err := runRepair([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "pi-plugin")); !os.IsNotExist(err) {
		t.Fatalf("expected repair to keep legacy pi-plugin absent, got %v", err)
	}
	if _, err := os.Stat(plugin); err != nil {
		t.Fatalf("expected repair to restore plugin: %v", err)
	}
}

func TestDryRunPreviewsDoNotWrite(t *testing.T) {
	root := t.TempDir()
	if result, err := runInstall([]string{"--root", root, "--dry-run"}); err != nil {
		t.Fatal(err)
	} else if result.Message == "" {
		t.Fatal("expected dry-run message")
	}
	if _, err := os.Stat(filepath.Join(root, ".sane-next-owned")); !os.IsNotExist(err) {
		t.Fatalf("expected dry-run not to write ownership marker, got %v", err)
	}
}

func TestInstallRefusesNonEmptyUnownedRoot(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "user.txt"), []byte("keep me\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runInstall([]string{"--root", root}); err == nil {
		t.Fatal("expected install to refuse non-empty unowned root")
	}
	if _, err := os.Stat(filepath.Join(root, "user.txt")); err != nil {
		t.Fatalf("expected user file to remain: %v", err)
	}
}

func TestPackValidateChecksConfiguredSkills(t *testing.T) {
	if _, err := runPack([]string{"validate", "--config", filepath.Join("..", "pi-plugin", "config-schema.toml")}); err != nil {
		t.Fatal(err)
	}
}

func TestPackEnableDisableEditsConfig(t *testing.T) {
	root := t.TempDir()
	config := filepath.Join(root, "config.toml")
	data, err := os.ReadFile(filepath.Join("..", "pi-plugin", "config-schema.toml"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(config, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runPack([]string{"disable", "--config", config, "core-workflow"}); err != nil {
		t.Fatal(err)
	}
	cfg, err := loadConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Packs[0].Enabled {
		t.Fatal("expected core-workflow disabled")
	}
	if _, err := runPack([]string{"enable", "--config", config, "core-workflow"}); err != nil {
		t.Fatal(err)
	}
	cfg, err = loadConfig(config)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Packs[0].Enabled {
		t.Fatal("expected core-workflow enabled")
	}
}

func TestFullCLIFlowIsCrossPlatform(t *testing.T) {
	root := t.TempDir()
	if _, err := runInstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := runDoctor([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := runExport([]string{"--config", filepath.Join("..", "pi-plugin", "config-schema.toml"), "--target", "codex", "--target-root", filepath.Join(root, "exports-fixture")}); err != nil {
		t.Fatal(err)
	}
	if _, err := runPack([]string{"validate", "--config", filepath.Join("..", "pi-plugin", "config-schema.toml")}); err != nil {
		t.Fatal(err)
	}
	if _, err := runRepair([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := runUpdate([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := runUninstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "extensions")); !os.IsNotExist(err) {
		t.Fatalf("expected uninstall to remove extension assets, got %v", err)
	}
}

func TestPackageListAndInstallAreFixtureSafe(t *testing.T) {
	result, err := runPackage([]string{"list", "--config", filepath.Join("..", "pi-plugin", "config-schema.toml")})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"pi-subagents\tdefault", "pi-curated-themes\toptional", "pi-markdown-preview\toptional", "pi-pretty\toptional", "pi-pledit\toptional", "pi-container-sandbox\toptional"} {
		if !strings.Contains(result.Message, want) {
			t.Fatalf("package list missing %q in:\n%s", want, result.Message)
		}
	}

	root := t.TempDir()
	piBin, log := writeFakePi(t, root)
	if _, err := runPackage([]string{"install", "--config", filepath.Join("..", "pi-plugin", "config-schema.toml"), "--pi-bin", piBin, "pi-markdown-preview"}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(log)
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.ReplaceAll(string(data), "\r\n", "\n"); got != "install npm:pi-markdown-preview@0.9.7\n" {
		t.Fatalf("unexpected pi invocation %q", got)
	}
}

func writeFakePi(t *testing.T, root string) (string, string) {
	t.Helper()
	log := filepath.Join(root, "pi.log")
	if runtime.GOOS == "windows" {
		piBin := filepath.Join(root, "pi.bat")
		content := "@echo off\r\necho %*>>\"" + log + "\"\r\n"
		if err := os.WriteFile(piBin, []byte(content), 0o755); err != nil {
			t.Fatal(err)
		}
		return piBin, log
	}
	piBin := filepath.Join(root, "pi")
	content := "#!/bin/sh\necho \"$@\" >> \"" + log + "\"\n"
	if err := os.WriteFile(piBin, []byte(content), 0o755); err != nil {
		t.Fatal(err)
	}
	return piBin, log
}

func TestConfigureThemePreservesUnrelatedSettings(t *testing.T) {
	agentDir := t.TempDir()
	settingsPath := filepath.Join(agentDir, "settings.json")
	if err := os.WriteFile(settingsPath, []byte(`{"model":"gpt-5.5","nested":{"keep":true}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runConfigure([]string{"--agent-dir", agentDir, "--theme", "github-dark-pro"}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(settingsPath)
	if err != nil {
		t.Fatal(err)
	}
	var settings map[string]any
	if err := json.Unmarshal(data, &settings); err != nil {
		t.Fatal(err)
	}
	if settings["theme"] != "github-dark-pro" || settings["model"] != "gpt-5.5" {
		t.Fatalf("settings not merged correctly: %#v", settings)
	}
	if _, ok := settings["nested"].(map[string]any); !ok {
		t.Fatalf("nested setting was not preserved: %#v", settings)
	}
}

func TestExportRefusesNonSaneOwnedDestination(t *testing.T) {
	root := t.TempDir()
	dest := filepath.Join(root, ".codex", "skills", "core-workflow")
	if err := os.MkdirAll(dest, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dest, "SKILL.md"), []byte("user owned\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := runExport([]string{"--config", filepath.Join("..", "pi-plugin", "config-schema.toml"), "--target-root", root})
	if err == nil {
		t.Fatal("expected export to refuse non-Sane destination")
	}
}

func TestCopyDirRejectsSymlinks(t *testing.T) {
	root := t.TempDir()
	src := filepath.Join(root, "src")
	if err := os.MkdirAll(src, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "secret.txt"), []byte("secret\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "secret.txt"), filepath.Join(src, "linked.txt")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	if err := copyDir(src, filepath.Join(root, "dst")); err == nil {
		t.Fatal("expected symlink copy to be rejected")
	}
}

func TestConfigRejectsUnsafeIDsAndExportPaths(t *testing.T) {
	base, err := os.ReadFile(filepath.Join("..", "pi-plugin", "config-schema.toml"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name string
		old  string
		new  string
	}{
		{name: "pack traversal", old: `id = "core-workflow"`, new: `id = "../core-workflow"`},
		{name: "absolute export path", old: `path = ".codex/skills"`, new: `path = "/tmp/sane-next-skills"`},
		{name: "escaping export path", old: `path = ".codex/skills"`, new: `path = "../skills"`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			config := filepath.Join(t.TempDir(), "config.toml")
			data := strings.Replace(string(base), tc.old, tc.new, 1)
			if err := os.WriteFile(config, []byte(data), 0o644); err != nil {
				t.Fatal(err)
			}
			if _, err := loadConfig(config); err == nil {
				t.Fatal("expected invalid config to be rejected")
			}
		})
	}
}
