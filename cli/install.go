package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runInstall(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("install", &flagOutput)
	root := fs.String("root", defaultRoot(), "installation root")
	sourceRoot := fs.String("source-root", defaultSourceRoot(), "sane-next source root")
	installRecommended := fs.Bool("recommended-pi-packages", true, "install curated recommended Pi packages with pi install")
	dryRun := fs.Bool("dry-run", false, "preview Sane-owned paths without writing")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("install does not accept positional arguments")
	}

	if *dryRun {
		return commandResult{Message: lifecyclePreview("install", *root, *sourceRoot)}, nil
	}
	if err := syncOwnedAssets(*root, *sourceRoot); err != nil {
		return commandResult{}, err
	}

	installed, err := installRecommendedPiPackages(*root, *sourceRoot, *installRecommended)
	if err != nil {
		return commandResult{}, err
	}
	message := fmt.Sprintf("installed sane-next overlay at %s", *root)
	if len(installed) > 0 {
		message = fmt.Sprintf("%s; installed recommended Pi packages: %s", message, strings.Join(installed, ", "))
	}
	return commandResult{Message: message}, nil
}

func installRecommendedPiPackages(root, sourceRoot string, enabled bool) ([]string, error) {
	if !enabled || (filepath.Clean(root) != filepath.Clean(defaultRoot()) && os.Getenv("SANE_NEXT_INSTALL_RECOMMENDED_IN_FIXTURES") != "1") {
		return nil, nil
	}
	if _, err := exec.LookPath("pi"); err != nil {
		return nil, nil
	}
	cfg, err := loadConfig(filepath.Join(sourceRoot, "pi-plugin", "config-schema.toml"))
	if err != nil {
		return nil, err
	}
	installed := []string{}
	for _, pkg := range cfg.RecommendedPiPackages {
		if !pkg.Enabled || !pkg.DefaultInstall {
			continue
		}
		cmd := exec.Command("pi", "install", pkg.Package)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return installed, fmt.Errorf("install recommended Pi package %s: %w\n%s", pkg.Package, err, string(out))
		}
		installed = append(installed, pkg.Package)
	}
	return installed, nil
}
