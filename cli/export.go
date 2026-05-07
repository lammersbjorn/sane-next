package main

import (
	"bytes"
	"fmt"
	"path/filepath"
)

func runExport(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("export", &flagOutput)
	configPath := fs.String("config", "../pi-plugin/config-schema.toml", "Sane config path")
	sourceRoot := fs.String("source-root", "", "base directory for relative pack sources")
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
	baseDir := configDir
	if *sourceRoot != "" {
		baseDir = *sourceRoot
	}
	for _, p := range cfg.allPacks() {
		if !p.Enabled || !supportsTarget(p, *targetID) {
			continue
		}
		sourceDir := filepath.Clean(filepath.Join(baseDir, p.Source))
		targetSkill := filepath.Join(root, target.Path, p.ID, "SKILL.md")
		if err := copyDir(sourceDir, filepath.Dir(targetSkill)); err != nil {
			return commandResult{}, err
		}
		count++
	}

	return commandResult{Message: fmt.Sprintf("exported %d pack(s) to %s target at %s", count, target.ID, root)}, nil
}
