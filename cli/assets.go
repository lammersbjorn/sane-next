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

	for _, dir := range []string{"packs", "skills", "extensions", "pi-plugin", "exports", "user-packs"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			return fmt.Errorf("create %s: %w", dir, err)
		}
	}

	if err := replaceDir(filepath.Join(sourceRoot, "pi-plugin"), filepath.Join(root, "pi-plugin")); err != nil {
		return err
	}
	if err := replaceDir(filepath.Join(sourceRoot, "pi-plugin"), filepath.Join(root, "extensions", "sane-next")); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, "packs"), 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, "skills"), 0o755); err != nil {
		return err
	}
	for _, p := range cfg.Packs {
		if !p.Enabled {
			continue
		}
		src := filepath.Clean(filepath.Join(filepath.Dir(cfgPath), p.Source))
		dst := filepath.Join(root, "packs", p.ID)
		if err := replaceDir(src, dst); err != nil {
			return err
		}
		if err := replaceDir(src, filepath.Join(root, "skills", p.ID)); err != nil {
			return err
		}
	}
	if err := writePackageJSON(root); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, "packs", "VERSION"), []byte(version+"\n"), 0o644); err != nil {
		return fmt.Errorf("write pack version: %w", err)
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
    "skills": ["./skills"]
  }
}
`, version)
	return os.WriteFile(filepath.Join(root, "package.json"), []byte(content), 0o644)
}

func writeInstalledConfig(root string, cfg saneConfig) error {
	var b strings.Builder
	b.WriteString("version = 1\n\n[defaults]\n")
	b.WriteString(fmt.Sprintf("model = %q\nreasoning = %q\n\n", cfg.Defaults.Model, cfg.Defaults.Reasoning))
	for _, p := range cfg.Packs {
		b.WriteString("[[packs]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nenabled = %t\nsource = %q\ntargets = %s\n\n", p.ID, p.Enabled, filepath.ToSlash(filepath.Join("..", "packs", p.ID)), tomlStringArray(p.Targets)))
	}
	for _, p := range cfg.UserPacks {
		b.WriteString("[[user_packs]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nenabled = %t\nsource = %q\ntargets = %s\n\n", p.ID, p.Enabled, p.Source, tomlStringArray(p.Targets)))
	}
	for _, t := range cfg.ExportTargets {
		b.WriteString("[[export_targets]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nkind = %q\npath = %q\n\n", t.ID, t.Kind, t.Path))
	}
	b.WriteString("[ownership]\nmarker_file = \".sane-next-owned\"\npreserve_user_config = true\n")
	return os.WriteFile(filepath.Join(root, "pi-plugin", "config-schema.toml"), []byte(b.String()), 0o644)
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
