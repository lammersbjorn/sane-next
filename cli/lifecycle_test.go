package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLifecyclePreservesUserPacksOnUninstall(t *testing.T) {
	root := t.TempDir()
	if _, err := runInstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	userPack := filepath.Join(root, "user-packs", "mine", "SKILL.md")
	if err := os.MkdirAll(filepath.Dir(userPack), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(userPack, []byte("user owned\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := runUpdate([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "packs", "VERSION")); err != nil {
		t.Fatalf("expected owned pack version: %v", err)
	}

	if _, err := runUninstall([]string{"--root", root}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(userPack); err != nil {
		t.Fatalf("expected user pack to remain: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "packs")); !os.IsNotExist(err) {
		t.Fatalf("expected owned packs removed, got %v", err)
	}
}

func TestDoctorRequiresOwnershipMarker(t *testing.T) {
	root := t.TempDir()
	if _, err := runDoctor([]string{"--root", root}); err == nil {
		t.Fatal("expected doctor to fail without ownership marker")
	}
}
