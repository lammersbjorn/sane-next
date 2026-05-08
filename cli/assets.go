package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func defaultSourceRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return "."
	}
	if filepath.Base(wd) == "cli" {
		return filepath.Clean(filepath.Join(wd, ".."))
	}
	return wd
}

func syncOwnedAssets(root, sourceRoot string) error {
	sourceRoot = filepath.Clean(sourceRoot)
	cfgPath := filepath.Join(sourceRoot, "pi-plugin", "config-schema.toml")
	cfg, err := loadConfig(cfgPath)
	if err != nil {
		return err
	}
	existingCfg, hasExistingCfg := loadExistingInstalledConfig(root)

	for _, dir := range []string{"packs", "extensions", "exports", "themes", "user-packs"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			return fmt.Errorf("create %s: %w", dir, err)
		}
	}

	if err := os.RemoveAll(filepath.Join(root, "pi-plugin")); err != nil {
		return fmt.Errorf("remove legacy pi-plugin dir: %w", err)
	}
	if err := replaceDir(filepath.Join(sourceRoot, "pi-plugin"), filepath.Join(root, "extensions", "sane-next")); err != nil {
		return err
	}
	if err := replaceDir(filepath.Join(sourceRoot, "themes"), filepath.Join(root, "themes")); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, "packs"), 0o755); err != nil {
		return err
	}
	if err := os.RemoveAll(filepath.Join(root, "skills")); err != nil {
		return fmt.Errorf("remove legacy skills dir: %w", err)
	}
	for _, p := range cfg.Packs {
		if !p.Enabled {
			continue
		}
		src := filepath.Clean(filepath.Join(filepath.Dir(cfgPath), p.Source))
		dst, err := safeJoinUnder(filepath.Join(root, "packs"), p.ID)
		if err != nil {
			return err
		}
		if err := replaceDir(src, dst); err != nil {
			return err
		}
	}
	if err := writePackageJSON(root); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, "packs", "VERSION"), []byte(version+"\n"), 0o644); err != nil {
		return fmt.Errorf("write pack version: %w", err)
	}
	if hasExistingCfg && cfg.Ownership.PreserveUserConfig {
		cfg = preserveUserConfig(cfg, existingCfg)
	}
	if err := writeInstalledConfig(root, cfg); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, ".sane-next-owned"), []byte("sane-next\n"), 0o644); err != nil {
		return fmt.Errorf("write ownership marker: %w", err)
	}
	return nil
}

func writePackageJSON(root string) error {
	content := fmt.Sprintf(`{
  "name": "sane-next-overlay",
  "version": %q,
  "private": true,
  "keywords": ["pi-package"],
  "peerDependencies": {
    "@mariozechner/pi-coding-agent": "*"
  },
  "pi": {
    "extensions": ["./extensions/sane-next/index.ts"],
    "themes": ["./themes"]
  }
}
`, version)
	return os.WriteFile(filepath.Join(root, "package.json"), []byte(content), 0o644)
}

func loadExistingInstalledConfig(root string) (saneConfig, bool) {
	cfg, err := loadConfig(filepath.Join(root, "extensions", "sane-next", "config-schema.toml"))
	return cfg, err == nil
}

func preserveUserConfig(sourceCfg, existingCfg saneConfig) saneConfig {
	sourceCfg.Defaults = existingCfg.Defaults
	sourceCfg.Rtk = existingCfg.Rtk
	sourceCfg.Pretty = existingCfg.Pretty
	preservePackChoices(sourceCfg.Packs, existingCfg.Packs)
	preservePackChoices(sourceCfg.UserPacks, existingCfg.UserPacks)
	return sourceCfg
}

func preservePackChoices(sourcePacks, existingPacks []pack) {
	existingByID := map[string]pack{}
	for _, p := range existingPacks {
		existingByID[p.ID] = p
	}
	for i := range sourcePacks {
		if existing, ok := existingByID[sourcePacks[i].ID]; ok {
			sourcePacks[i].Enabled = existing.Enabled
			sourcePacks[i].Targets = existing.Targets
		}
	}
}

func writeInstalledConfig(root string, cfg saneConfig) error {
	extensionCfg := installedConfigWithPackBase(cfg, filepath.Join("..", ".."))
	if err := os.WriteFile(filepath.Join(root, "extensions", "sane-next", "config-schema.toml"), []byte(renderConfig(extensionCfg)), 0o644); err != nil {
		return err
	}
	return writeInstalledManifest(root)
}

func installedConfigWithPackBase(cfg saneConfig, base string) saneConfig {
	for i := range cfg.Packs {
		cfg.Packs[i].Source = filepath.ToSlash(filepath.Join(base, "packs", cfg.Packs[i].ID))
	}
	for i := range cfg.UserPacks {
		cfg.UserPacks[i].Source = filepath.ToSlash(filepath.Join(base, "user-packs", cfg.UserPacks[i].ID))
	}
	return cfg
}

func writeInstalledManifest(root string) error {
	content := fmt.Sprintf(`id = "sane-next"
name = "Sane Next"
version = %q
description = "Pi-first overlay for Sane shared workflow packs and companion CLI integration."
entrypoint = "index.ts"

[runtime]
host = "pi"
minimum_version = "0.0.0"

[config]
schema = "config-schema.toml"

# Skills are discovered dynamically from config-schema.toml by the extension.
`, version)
	return os.WriteFile(filepath.Join(root, "extensions", "sane-next", "manifest.toml"), []byte(content), 0o644)
}

func tomlStringArray(values []string) string {
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = fmt.Sprintf("%q", v)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func replaceDir(src, dst string) error {
	if err := os.RemoveAll(dst); err != nil {
		return fmt.Errorf("remove %s: %w", dst, err)
	}
	return copyDir(src, dst)
}

func safeJoinUnder(root string, elem ...string) (string, error) {
	cleanRoot, err := filepath.Abs(filepath.Clean(root))
	if err != nil {
		return "", fmt.Errorf("resolve root %s: %w", root, err)
	}
	parts := append([]string{cleanRoot}, elem...)
	joined := filepath.Clean(filepath.Join(parts...))
	rel, err := filepath.Rel(cleanRoot, joined)
	if err != nil {
		return "", fmt.Errorf("resolve path %s: %w", joined, err)
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel) {
		return "", fmt.Errorf("path %s escapes root %s", joined, cleanRoot)
	}
	return joined, nil
}

func copyDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("stat source %s: %w", src, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("source %s is not a directory", src)
	}
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if d.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("refusing to copy symlink %s", path)
		}
		return copyFile(path, target)
	})
}

func copyFile(source string, target string) error {
	in, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open source %s: %w", source, err)
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return fmt.Errorf("create target dir: %w", err)
	}
	out, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open target %s: %w", target, err)
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s to %s: %w", source, target, err)
	}
	return nil
}
