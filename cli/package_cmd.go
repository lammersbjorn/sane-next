package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func runPackage(args []string) (commandResult, error) {
	if len(args) == 0 {
		return commandResult{}, fmt.Errorf("package requires subcommand: list, install")
	}
	var flagOutput bytes.Buffer
	fs := stringFlagSet("package "+args[0], &flagOutput)
	configPath := fs.String("config", "../pi-plugin/config-schema.toml", "Sane config path")
	piBin := fs.String("pi-bin", "pi", "pi executable")
	dryRun := fs.Bool("dry-run", false, "preview package install without running pi install")
	if err := fs.Parse(args[1:]); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}

	cfg, err := loadConfig(*configPath)
	if err != nil {
		return commandResult{}, err
	}

	switch args[0] {
	case "list":
		if fs.NArg() != 0 {
			return commandResult{}, fmt.Errorf("package list does not accept positional arguments")
		}
		return commandResult{Message: renderPackageList(cfg)}, nil
	case "install":
		if fs.NArg() != 1 {
			return commandResult{}, fmt.Errorf("package install requires exactly one package id")
		}
		pkg, ok := findRecommendedPackage(cfg, fs.Arg(0))
		if !ok || !pkg.Enabled {
			return commandResult{}, fmt.Errorf("unknown enabled package recommendation %q", fs.Arg(0))
		}
		if *dryRun {
			return commandResult{Message: fmt.Sprintf("would run: %s install %s", *piBin, pkg.Package)}, nil
		}
		cmd := exec.Command(*piBin, "install", pkg.Package)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return commandResult{}, fmt.Errorf("install Pi package %s: %w\n%s", pkg.ID, err, string(out))
		}
		return commandResult{Message: fmt.Sprintf("installed Pi package %s (%s)", pkg.ID, pkg.Package)}, nil
	default:
		return commandResult{}, fmt.Errorf("unknown package subcommand %q", args[0])
	}
}

func renderPackageList(cfg saneConfig) string {
	lines := []string{}
	for _, pkg := range cfg.RecommendedPiPackages {
		if !pkg.Enabled {
			continue
		}
		install := "optional"
		if pkg.DefaultInstall {
			install = "default"
		}
		lines = append(lines, fmt.Sprintf("%s\t%s\t%s\t%s", pkg.ID, install, pkg.Package, pkg.Purpose))
	}
	return strings.Join(lines, "\n")
}

func findRecommendedPackage(cfg saneConfig, id string) (recommendedPiPackage, bool) {
	for _, pkg := range cfg.RecommendedPiPackages {
		if pkg.ID == id {
			return pkg, true
		}
	}
	return recommendedPiPackage{}, false
}
