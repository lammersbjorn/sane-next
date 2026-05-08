package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type saneConfig struct {
	Defaults              defaults               `toml:"defaults"`
	Rtk                   rtk                    `toml:"rtk"`
	Pretty                pretty                 `toml:"pretty"`
	Packs                 []pack                 `toml:"packs"`
	UserPacks             []pack                 `toml:"user_packs"`
	ExportTargets         []exportTarget         `toml:"export_targets"`
	RecommendedPiPackages []recommendedPiPackage `toml:"recommended_pi_packages"`
	Ownership             ownership              `toml:"ownership"`
}

type defaults struct {
	Model         string `toml:"model"`
	Reasoning     string `toml:"reasoning"`
	ResponseStyle string `toml:"response_style"`
}

type rtk struct {
	Mode string `toml:"mode"`
}

type pretty struct {
	MaxPreviewLines int    `toml:"max_preview_lines"`
	MaxHlChars      int    `toml:"max_hl_chars"`
	Icons           string `toml:"icons"`
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

type recommendedPiPackage struct {
	ID             string `toml:"id"`
	Enabled        bool   `toml:"enabled"`
	DefaultInstall bool   `toml:"default_install"`
	Package        string `toml:"package"`
	Purpose        string `toml:"purpose"`
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
	for _, p := range cfg.UserPacks {
		if p.ID == "" || p.Source == "" {
			return fmt.Errorf("user pack id and source are required")
		}
	}
	for _, target := range cfg.ExportTargets {
		if target.ID == "" || target.Kind == "" || target.Path == "" {
			return fmt.Errorf("export target id, kind, and path are required")
		}
	}
	for _, pkg := range cfg.RecommendedPiPackages {
		if pkg.ID == "" || pkg.Package == "" || pkg.Purpose == "" {
			return fmt.Errorf("recommended Pi package id, package, and purpose are required")
		}
	}
	return nil
}

func (cfg saneConfig) allPacks() []pack {
	packs := make([]pack, 0, len(cfg.Packs)+len(cfg.UserPacks))
	packs = append(packs, cfg.Packs...)
	packs = append(packs, cfg.UserPacks...)
	return packs
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

func renderConfig(cfg saneConfig) string {
	var b strings.Builder
	b.WriteString("# Sane Next Pi overlay config.\n")
	b.WriteString("# Global install copy: ~/.sane-next/pi-plugin/config-schema.toml\n")
	b.WriteString("# Repo source copy: pi-plugin/config-schema.toml\n")
	b.WriteString("# After manual edits to the repo source, run: cd cli && ./sane-next update --root ~/.sane-next\n\n")
	b.WriteString("version = 1\n\n")
	b.WriteString("# Defaults requested from Pi sessions that load Sane.\n")
	b.WriteString("# reasoning: low, medium, high, or another value supported by your active model/provider.\n")
	b.WriteString("# response_style is optional; current special value: caveman. Omit for normal responses.\n")
	b.WriteString("[defaults]\n")
	b.WriteString(fmt.Sprintf("model = %q\nreasoning = %q\n", cfg.Defaults.Model, cfg.Defaults.Reasoning))
	if cfg.Defaults.ResponseStyle != "" {
		b.WriteString(fmt.Sprintf("response_style = %q\n", cfg.Defaults.ResponseStyle))
	}
	b.WriteString("\n")
	rtkMode := cfg.Rtk.Mode
	if rtkMode == "" {
		rtkMode = "warn"
	}
	b.WriteString("# RTK routing controls how strongly Sane discourages raw shell for search/read/test/diff work.\n")
	b.WriteString("# mode values: off = no RTK policy; advise = prompt guidance only; warn = warn/status; enforce = block matching raw shell commands.\n")
	b.WriteString("[rtk]\n")
	b.WriteString(fmt.Sprintf("mode = %q\n\n", rtkMode))
	prettyCfg := cfg.Pretty
	if prettyCfg.MaxPreviewLines == 0 {
		prettyCfg.MaxPreviewLines = 24
	}
	if prettyCfg.MaxHlChars == 0 {
		prettyCfg.MaxHlChars = 1
	}
	if prettyCfg.Icons == "" {
		prettyCfg.Icons = "nerd"
	}
	b.WriteString("# Defaults applied before pi-pretty renders tool output. These are no-ops unless pi-pretty is installed.\n")
	b.WriteString("[pretty]\n")
	b.WriteString(fmt.Sprintf("max_preview_lines = %d\nmax_hl_chars = %d\nicons = %q\n\n", prettyCfg.MaxPreviewLines, prettyCfg.MaxHlChars, prettyCfg.Icons))
	b.WriteString("# Built-in Sane skill packs. Set enabled = false to stop loading/exporting one.\n")
	b.WriteString("# targets: pi loads the pack in Pi; codex includes it in Codex skill export.\n")
	for _, p := range cfg.Packs {
		b.WriteString("[[packs]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nenabled = %t\nsource = %q\ntargets = %s\n\n", p.ID, p.Enabled, filepath.ToSlash(p.Source), tomlStringArray(p.Targets)))
	}
	b.WriteString("# User packs are externally authored skills registered into Sane.\n")
	for _, p := range cfg.UserPacks {
		b.WriteString("[[user_packs]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nenabled = %t\nsource = %q\ntargets = %s\n\n", p.ID, p.Enabled, filepath.ToSlash(p.Source), tomlStringArray(p.Targets)))
	}
	b.WriteString("# Export targets used by sane-next export. path is relative to --target-root unless absolute.\n")
	for _, t := range cfg.ExportTargets {
		b.WriteString("[[export_targets]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nkind = %q\npath = %q\n\n", t.ID, t.Kind, t.Path))
	}
	b.WriteString("# Pi packages known to Sane. default_install=true packages are installed by sane-next install unless skipped.\n")
	b.WriteString("# Optional packages stay default_install=false and can be installed with: sane-next package install <id>\n")
	for _, pkg := range cfg.RecommendedPiPackages {
		b.WriteString("[[recommended_pi_packages]]\n")
		b.WriteString(fmt.Sprintf("id = %q\nenabled = %t\ndefault_install = %t\npackage = %q\npurpose = %q\n\n", pkg.ID, pkg.Enabled, pkg.DefaultInstall, pkg.Package, pkg.Purpose))
	}
	marker := cfg.Ownership.MarkerFile
	if marker == "" {
		marker = ".sane-next-owned"
	}
	b.WriteString("# Ownership controls Sane lifecycle safety. Do not change marker_file for an existing install unless repairing deliberately.\n")
	b.WriteString("[ownership]\n")
	b.WriteString(fmt.Sprintf("marker_file = %q\npreserve_user_config = %t\n", marker, cfg.Ownership.PreserveUserConfig))
	return b.String()
}
