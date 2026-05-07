package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const version = "0.1.0"

type commandResult struct {
	Message string
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		usage(stdout)
		return nil
	}

	switch args[0] {
	case "install":
		result, err := runInstall(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "version":
		fmt.Fprintln(stdout, version)
		return nil
	case "help", "-h", "--help":
		usage(stdout)
		return nil
	default:
		usage(stderr)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func usage(w io.Writer) {
	fmt.Fprintln(w, "sane-next manages the Sane Pi overlay and shared workflow packs.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  sane-next install [--root PATH]")
	fmt.Fprintln(w, "  sane-next version")
}

func defaultRoot() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return "."
	}
	return filepath.Join(home, ".sane-next")
}

func stringFlagSet(name string, output io.Writer) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(output)
	return fs
}
