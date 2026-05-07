package main

import (
	"bytes"
	"fmt"
)

func runInstall(args []string) (commandResult, error) {
	var flagOutput bytes.Buffer
	fs := stringFlagSet("install", &flagOutput)
	root := fs.String("root", defaultRoot(), "installation root")
	sourceRoot := fs.String("source-root", defaultSourceRoot(), "sane-next source root")
	if err := fs.Parse(args); err != nil {
		return commandResult{}, fmt.Errorf("%s", flagOutput.String())
	}
	if fs.NArg() != 0 {
		return commandResult{}, fmt.Errorf("install does not accept positional arguments")
	}

	if err := syncOwnedAssets(*root, *sourceRoot); err != nil {
		return commandResult{}, err
	}

	return commandResult{Message: fmt.Sprintf("installed sane-next overlay at %s", *root)}, nil
}
