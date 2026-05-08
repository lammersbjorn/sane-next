package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func runConfigure(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("configure", &flagOutput)
	agentDir := fs.String("agent-dir", defaultPiAgentDir(), "Pi agent config directory")
	theme := fs.String("theme", "", "active Pi theme to set")
	dryRun := fs.Bool("dry-run", false, "preview settings change without writing")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("configure does not accept positional arguments")
	}
	if *theme == "" {
		return commandResult{}, fmt.Errorf("configure requires at least one setting, for example --theme github-dark-pro")
	}
	if *theme != "github-dark-pro" {
		return commandResult{}, fmt.Errorf("unsupported Sane theme %q", *theme)
	}

	settingsPath := filepath.Join(*agentDir, "settings.json")
	if *dryRun {
		return commandResult{Message: fmt.Sprintf("would set theme = %q in %s", *theme, settingsPath)}, nil
	}
	if err := mergePiSettings(settingsPath, map[string]any{"theme": *theme}); err != nil {
		return commandResult{}, err
	}
	return commandResult{Message: fmt.Sprintf("set theme = %q in %s", *theme, settingsPath)}, nil
}

func defaultPiAgentDir() string {
	if dir := os.Getenv("PI_CODING_AGENT_DIR"); dir != "" {
		return dir
	}
	home, err := userHomeDir()
	if err != nil || home == "" {
		return filepath.Join(".pi", "agent")
	}
	return filepath.Join(home, ".pi", "agent")
}

func mergePiSettings(path string, updates map[string]any) error {
	settings := map[string]any{}
	data, err := os.ReadFile(path)
	if err == nil && len(bytes.TrimSpace(data)) > 0 {
		if err := json.Unmarshal(data, &settings); err != nil {
			return fmt.Errorf("parse Pi settings %s: %w", path, err)
		}
	} else if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("read Pi settings %s: %w", path, err)
	}
	for key, value := range updates {
		settings[key] = value
	}
	out, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}
	out = append(out, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}
