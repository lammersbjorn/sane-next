package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func runPack(args []string) (commandResult, error) {
	if len(args) == 0 {
		return commandResult{}, fmt.Errorf("pack requires subcommand: list, validate, enable, disable")
	}
	var flagOutput bytes.Buffer
	fs := stringFlagSet("pack "+args[0], &flagOutput)
	configPath := fs.String("config", "../pi-plugin/config-schema.toml", "Sane config path")
	if err := fs.Parse(args[1:]); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}

	switch args[0] {
	case "list":
		if fs.NArg() != 0 {
			return commandResult{}, fmt.Errorf("pack list does not accept positional arguments")
		}
		cfg, err := loadConfig(*configPath)
		if err != nil {
			return commandResult{}, err
		}
		var lines []string
		for _, p := range cfg.allPacks() {
			state := "disabled"
			if p.Enabled {
				state = "enabled"
			}
			lines = append(lines, fmt.Sprintf("%s\t%s\t%s", p.ID, state, strings.Join(p.Targets, ",")))
		}
		return commandResult{Message: strings.Join(lines, "\n")}, nil
	case "validate":
		if fs.NArg() != 0 {
			return commandResult{}, fmt.Errorf("pack validate does not accept positional arguments")
		}
		cfg, err := loadConfig(*configPath)
		if err != nil {
			return commandResult{}, err
		}
		if err := validatePackFiles(cfg, *configPath); err != nil {
			return commandResult{}, err
		}
		return commandResult{Message: fmt.Sprintf("validated %d configured pack(s) in %s", len(cfg.allPacks()), *configPath)}, nil
	case "enable", "disable":
		if fs.NArg() != 1 {
			return commandResult{}, fmt.Errorf("pack %s requires exactly one pack id", args[0])
		}
		id := fs.Arg(0)
		enabled := args[0] == "enable"
		if err := setPackEnabled(*configPath, id, enabled); err != nil {
			return commandResult{}, err
		}
		state := "disabled"
		if enabled {
			state = "enabled"
		}
		return commandResult{Message: fmt.Sprintf("%s %s in %s", state, id, *configPath)}, nil
	default:
		return commandResult{}, fmt.Errorf("unknown pack subcommand %q", args[0])
	}
}

func validatePackFiles(cfg saneConfig, configPath string) error {
	seen := map[string]bool{}
	base := filepath.Dir(configPath)
	for _, p := range cfg.allPacks() {
		if seen[p.ID] {
			return fmt.Errorf("duplicate pack id %q", p.ID)
		}
		seen[p.ID] = true
		skillPath := filepath.Join(base, p.Source, "SKILL.md")
		data, err := os.ReadFile(skillPath)
		if err != nil {
			return fmt.Errorf("pack %s missing SKILL.md at %s: %w", p.ID, skillPath, err)
		}
		body := string(data)
		for _, section := range []string{"# ", "## Goal", "## Use When", "## Inputs", "## Outputs", "## How To Run", "## Verification", "## Gotchas / Safety"} {
			if !strings.Contains(body, section) {
				return fmt.Errorf("pack %s SKILL.md missing required section marker %q", p.ID, section)
			}
		}
	}
	return nil
}

func setPackEnabled(path, id string, enabled bool) error {
	cfg, err := loadConfig(path)
	if err != nil {
		return err
	}
	found := false
	for i := range cfg.Packs {
		if cfg.Packs[i].ID == id {
			cfg.Packs[i].Enabled = enabled
			found = true
		}
	}
	for i := range cfg.UserPacks {
		if cfg.UserPacks[i].ID == id {
			cfg.UserPacks[i].Enabled = enabled
			found = true
		}
	}
	if !found {
		return fmt.Errorf("unknown pack %q", id)
	}
	return os.WriteFile(path, []byte(renderConfig(cfg)), 0o644)
}
