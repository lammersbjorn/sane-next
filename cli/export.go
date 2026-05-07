package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func runExport(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("export", &flagOutput)
	configPath := fs.String("config", "../pi-plugin/config-schema.toml", "Sane config path")
	targetID := fs.String("target", "codex", "export target id")
	targetRoot := fs.String("target-root", "", "root directory for exported artifacts")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("export does not accept positional arguments")
	}

	cfg, err := loadConfig(*configPath)
	if err != nil {
		return commandResult{}, err
	}
	target, ok := cfg.target(*targetID)
	if !ok {
		return commandResult{}, fmt.Errorf("unknown export target %q", *targetID)
	}

	root := *targetRoot
	if root == "" {
		root = filepath.Join(defaultRoot(), "exports")
	}

	count := 0
	configDir := filepath.Dir(*configPath)
	for _, p := range cfg.Packs {
		if !p.Enabled || !supportsTarget(p, *targetID) {
			continue
		}
		sourceDir := filepath.Clean(filepath.Join(configDir, p.Source))
		sourceSkill := filepath.Join(sourceDir, "SKILL.md")
		targetSkill := filepath.Join(root, target.Path, p.ID, "SKILL.md")
		if err := copyFile(sourceSkill, targetSkill); err != nil {
			return commandResult{}, err
		}
		count++
	}

	return commandResult{Message: fmt.Sprintf("exported %d pack(s) to %s target at %s", count, target.ID, root)}, nil
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
