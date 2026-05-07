package main

import "path/filepath"

func defaultExportRoot(targetID string) string {
	if targetID == "codex" {
		return userHomeOrDot()
	}
	return filepath.Join(defaultRoot(), "exports")
}

func userHomeOrDot() string {
	home, err := userHomeDir()
	if err != nil || home == "" {
		return "."
	}
	return home
}
