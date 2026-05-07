package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type saneConfig struct {
	Defaults      defaults       `toml:"defaults"`
	Packs         []pack         `toml:"packs"`
	ExportTargets []exportTarget `toml:"export_targets"`
	Ownership     ownership      `toml:"ownership"`
}

type defaults struct {
	Model     string `toml:"model"`
	Reasoning string `toml:"reasoning"`
}

type pack struct {
	ID      string   `toml:"id"`
	Enabled bool     `toml:"enabled"`
	Source  string   `toml:"source"`
	Targets []string `toml:"targets"`
}

type exportTarget struct {
	ID   string `toml:"id"`
	Kind string `toml:"kind"`
	Path string `toml:"path"`
}

type ownership struct {
	MarkerFile         string `toml:"marker_file"`
	PreserveUserConfig bool   `toml:"preserve_user_config"`
}

func loadConfig(path string) (saneConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return saneConfig{}, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg saneConfig
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return saneConfig{}, fmt.Errorf("parse config %s: %w", path, err)
	}
	if err := validateConfig(cfg); err != nil {
		return saneConfig{}, err
	}
	return cfg, nil
}

func validateConfig(cfg saneConfig) error {
	if cfg.Defaults.Model == "" {
		return fmt.Errorf("config defaults.model is required")
	}
	if cfg.Defaults.Reasoning == "" {
		return fmt.Errorf("config defaults.reasoning is required")
	}
	if len(cfg.Packs) == 0 {
		return fmt.Errorf("config must define at least one pack")
	}
	if len(cfg.ExportTargets) == 0 {
		return fmt.Errorf("config must define at least one export target")
	}
	for _, p := range cfg.Packs {
		if p.ID == "" || p.Source == "" {
			return fmt.Errorf("pack id and source are required")
		}
	}
	for _, target := range cfg.ExportTargets {
		if target.ID == "" || target.Kind == "" || target.Path == "" {
			return fmt.Errorf("export target id, kind, and path are required")
		}
	}
	return nil
}

func (cfg saneConfig) target(id string) (exportTarget, bool) {
	for _, target := range cfg.ExportTargets {
		if target.ID == id {
			return target, true
		}
	}
	return exportTarget{}, false
}

func supportsTarget(p pack, target string) bool {
	for _, t := range p.Targets {
		if t == target {
			return true
		}
	}
	return false
}
