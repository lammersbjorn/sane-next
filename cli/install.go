package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func runInstall(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("install", &flagOutput)
	root := fs.String("root", defaultRoot(), "installation root")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("install does not accept positional arguments")
	}

	dirs := []string{
		filepath.Join(*root, "packs"),
		filepath.Join(*root, "pi-plugin"),
		filepath.Join(*root, "exports"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return commandResult{}, fmt.Errorf("create %s: %w", dir, err)
		}
	}

	marker := filepath.Join(*root, ".sane-next-owned")
	if err := os.WriteFile(marker, []byte("sane-next\n"), 0o644); err != nil {
		return commandResult{}, fmt.Errorf("write ownership marker: %w", err)
	}

	return commandResult{Message: fmt.Sprintf("installed sane-next overlay at %s", *root)}, nil
}
