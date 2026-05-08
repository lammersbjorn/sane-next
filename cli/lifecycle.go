package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runDoctor(args []string) (commandResult, error) {
	root, _, _, err := lifecycleRoot("doctor", args)
	if err != nil {
		return commandResult{}, err
	}
	if err := requireHealthyInstall(root); err != nil {
		return commandResult{}, fmt.Errorf("sane-next install at %s is unhealthy: %w\nNext: run `sane-next repair --root %s` if this is a Sane-owned install, or `sane-next install --root %s` for a fresh overlay", root, err, root, root)
	}
	return commandResult{Message: fmt.Sprintf("sane-next install at %s is healthy; owned assets and Pi extension files are present", root)}, nil
}

func runRepair(args []string) (commandResult, error) {
	root, sourceRoot, dryRun, err := lifecycleRoot("repair", args)
	if err != nil {
		return commandResult{}, err
	}
	if err := requireOwnedInstall(root); err != nil {
		return commandResult{}, err
	}
	if dryRun {
		return commandResult{Message: lifecyclePreview("repair", root, sourceRoot)}, nil
	}
	if err := syncOwnedAssets(root, sourceRoot); err != nil {
		return commandResult{}, err
	}
	return commandResult{Message: fmt.Sprintf("repaired sane-next overlay at %s", root)}, nil
}

func runUpdate(args []string) (commandResult, error) {
	root, sourceRoot, dryRun, err := lifecycleRoot("update", args)
	if err != nil {
		return commandResult{}, err
	}
	if err := requireOwnedInstall(root); err != nil {
		return commandResult{}, err
	}
	if dryRun {
		return commandResult{Message: lifecyclePreview("update", root, sourceRoot)}, nil
	}
	if err := syncOwnedAssets(root, sourceRoot); err != nil {
		return commandResult{}, err
	}
	return commandResult{Message: fmt.Sprintf("updated sane-next owned packs at %s", root)}, nil
}

func runUninstall(args []string) (commandResult, error) {
	root, sourceRoot, dryRun, err := lifecycleRoot("uninstall", args)
	if err != nil {
		return commandResult{}, err
	}
	if err := requireOwnedInstall(root); err != nil {
		return commandResult{}, err
	}
	if dryRun {
		return commandResult{Message: lifecyclePreview("uninstall", root, sourceRoot)}, nil
	}
	for _, path := range []string{"packs", "skills", "extensions", "pi-plugin", "exports", "themes", "package.json", ".sane-next-owned"} {
		if err := os.RemoveAll(filepath.Join(root, path)); err != nil {
			return commandResult{}, fmt.Errorf("remove %s: %w", path, err)
		}
	}
	return commandResult{Message: fmt.Sprintf("uninstalled sane-next owned material from %s", root)}, nil
}

func lifecycleRoot(name string, args []string) (string, string, bool, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet(name, &flagOutput)
	root := fs.String("root", defaultRoot(), "installation root")
	sourceRoot := fs.String("source-root", defaultSourceRoot(), "sane-next source root")
	dryRun := fs.Bool("dry-run", false, "preview Sane-owned paths without writing")
	if err := fs.Parse(args); err != nil {
		return "", "", false, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return "", "", false, fmt.Errorf("%s does not accept positional arguments", name)
	}
	return *root, *sourceRoot, *dryRun, nil
}

func lifecyclePreview(action, root, sourceRoot string) string {
	owned := []string{"packs", "extensions", "pi-plugin", "exports", "themes", "package.json", ".sane-next-owned"}
	return fmt.Sprintf("dry-run %s: root=%s source=%s Sane-owned paths=%s; user config and user-packs are preserved", action, root, sourceRoot, strings.Join(owned, ", "))
}

func requireOwnedInstall(root string) error {
	marker := filepath.Join(root, ".sane-next-owned")
	if _, err := os.Stat(marker); err != nil {
		return fmt.Errorf("missing Sane ownership marker at %s", marker)
	}
	return nil
}

func requireHealthyInstall(root string) error {
	if err := requireOwnedInstall(root); err != nil {
		return err
	}
	for _, path := range []string{
		"packs/core-workflow/SKILL.md",
		"packs/rtk-routing/SKILL.md",
		"packs/agent-lanes/SKILL.md",
		"packs/sane-router/SKILL.md",
		"packs/VERSION",
		"extensions/sane-next/index.ts",
		"themes/github-dark-pro.json",
		"package.json",
		"pi-plugin/index.ts",
		"pi-plugin/manifest.toml",
		"pi-plugin/config-schema.toml",
	} {
		full := filepath.Join(root, path)
		if _, err := os.Stat(full); err != nil {
			return fmt.Errorf("missing or broken Sane-owned asset %s: %w", full, err)
		}
	}
	return nil
}
