package main

import (
	"encoding/json"
	"os"
	"path/filepath"
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
	plugin := filepath.Join(root, "pi-plugin", "index.ts")
	if _, err := os.Stat(plugin); err != nil {
		t.Fatalf("expected plugin copied: %v", err)
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
	piBin := filepath.Join(root, "pi")
	log := filepath.Join(root, "pi.log")
	if err := os.WriteFile(piBin, []byte("#!/bin/sh\necho \"$@\" >> "+log+"\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := runPackage([]string{"install", "--config", filepath.Join("..", "pi-plugin", "config-schema.toml"), "--pi-bin", piBin, "pi-markdown-preview"}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(log)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(data); got != "install npm:pi-markdown-preview@0.9.7\n" {
		t.Fatalf("unexpected pi invocation %q", got)
	}
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
