package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runExport(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("export", &flagOutput)
	configPath := fs.String("config", "../pi-plugin/config-schema.toml", "Sane config path")
	sourceRoot := fs.String("source-root", "", "base directory for relative pack sources")
	targetID := fs.String("target", "codex", "export target id")
	targetRoot := fs.String("target-root", "", "root directory for exported artifacts")
	dryRun := fs.Bool("dry-run", false, "preview export paths without writing")
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
		root = defaultExportRoot(target.ID)
	}

	count := 0
	var preview []string
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
		destDir, err := safeJoinUnder(root, target.Path, p.ID)
		if err != nil {
			return commandResult{}, err
		}
		if *dryRun {
			preview = append(preview, fmt.Sprintf("%s -> %s", sourceDir, destDir))
			count++
			continue
		}
		if err := requireExportWritable(destDir); err != nil {
			return commandResult{}, err
		}
		if err := replaceDir(sourceDir, destDir); err != nil {
			return commandResult{}, err
		}
		if err := os.WriteFile(filepath.Join(destDir, ".sane-next-exported"), []byte("sane-next\n"), 0o644); err != nil {
			return commandResult{}, fmt.Errorf("write export marker: %w", err)
		}
		count++
	}

	if *dryRun {
		return commandResult{Message: fmt.Sprintf("dry-run export %d pack(s) to %s target at %s\n%s", count, target.ID, root, strings.Join(preview, "\n"))}, nil
	}
	return commandResult{Message: fmt.Sprintf("exported %d pack(s) to %s target at %s", count, target.ID, root)}, nil
}

func requireExportWritable(destDir string) error {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(destDir, ".sane-next-exported")); err != nil {
		return fmt.Errorf("refusing to overwrite non-Sane export at %s", destDir)
	}
	return nil
}
