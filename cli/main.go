package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const version = "0.3.0-beta.3"

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
	case "configure":
		result, err := runConfigure(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "doctor":
		result, err := runDoctor(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "export":
		result, err := runExport(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "install":
		result, err := runInstall(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "pack":
		result, err := runPack(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "package":
		result, err := runPackage(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "repair":
		result, err := runRepair(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "uninstall":
		result, err := runUninstall(args[1:])
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, result.Message)
		return nil
	case "update":
		result, err := runUpdate(args[1:])
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
	fmt.Fprintln(w, "  sane-next configure --theme github-dark-pro [--agent-dir PATH] [--dry-run]")
	fmt.Fprintln(w, "  sane-next doctor [--root PATH]")
	fmt.Fprintln(w, "  sane-next export [--config PATH] [--source-root PATH] [--target codex] [--target-root PATH] [--dry-run]")
	fmt.Fprintln(w, "  sane-next install [--root PATH] [--source-root PATH] [--recommended-pi-packages=false] [--dry-run]")
	fmt.Fprintln(w, "  sane-next pack list|validate|enable|disable [--config PATH] [PACK_ID]")
	fmt.Fprintln(w, "  sane-next package list|install [--config PATH] [--pi-bin PATH] [--dry-run] [PACKAGE_ID]")
	fmt.Fprintln(w, "  sane-next repair [--root PATH] [--source-root PATH] [--dry-run]")
	fmt.Fprintln(w, "  sane-next uninstall [--root PATH] [--dry-run]")
	fmt.Fprintln(w, "  sane-next update [--root PATH] [--source-root PATH] [--dry-run]")
	fmt.Fprintln(w, "  sane-next version")
}

func defaultRoot() string {
	home, err := userHomeDir()
	if err != nil || home == "" {
		return "."
	}
	return filepath.Join(home, ".sane-next")
}

func userHomeDir() (string, error) {
	return os.UserHomeDir()
}

func stringFlagSet(name string, output io.Writer) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(output)
	return fs
}
