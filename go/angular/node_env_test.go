package angular

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func createFakeNgSetup(t *testing.T, rootDir string, version string) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("symlink test on unix")
	}
	binDir := filepath.Join(rootDir, "bin")
	cliDir := filepath.Join(rootDir, "lib", "node_modules", "@angular", "cli", "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(cliDir, 0o755); err != nil {
		t.Fatal(err)
	}

	ngJs := filepath.Join(cliDir, "ng.js")
	if err := os.WriteFile(ngJs, []byte("#!/usr/bin/env node\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	pkgJson := filepath.Join(filepath.Dir(cliDir), "package.json")
	pkgContent := `{"name": "@angular/cli", "version": "` + version + `"}`
	if err := os.WriteFile(pkgJson, []byte(pkgContent), 0o644); err != nil {
		t.Fatal(err)
	}

	ngSymlink := filepath.Join(binDir, "ng")
	relTarget, err := filepath.Rel(binDir, ngJs)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(relTarget, ngSymlink); err != nil {
		t.Fatal(err)
	}
	return binDir
}

func TestSelectLatestNodeVersionDir(t *testing.T) {
	tmp := t.TempDir()
	// Note: lexicographically "v22.9.0" > "v22.23.2", so semver comparison must pick v22.23.2.
	for _, dir := range []string{"v20.18.0", "v22.9.0", "v22.23.2"} {
		if err := os.MkdirAll(filepath.Join(tmp, dir), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	entries, err := os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	got := selectLatestNodeVersionDir(entries)
	if got != "v22.23.2" {
		t.Fatalf("selectLatestNodeVersionDir() = %q, want %q", got, "v22.23.2")
	}
}

func TestConfigureNodeEnvOverridesStaleLookPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix path test")
	}
	tmp := t.TempDir()
	oldGlobal := createFakeNgSetup(t, filepath.Join(tmp, "usr-local"), "20.3.13")

	fakeHome := filepath.Join(tmp, "home")
	nvmNodeRoot := filepath.Join(fakeHome, ".nvm", "versions", "node", "v22.23.2")
	nvmBin := createFakeNgSetup(t, nvmNodeRoot, "22.1.6")

	t.Setenv("HOME", fakeHome)
	t.Setenv("PATH", oldGlobal)

	// exec.Command resolves against oldGlobal (/usr/local/bin/ng -> 20.3.13) BEFORE ConfigureNodeEnv runs.
	cmd := exec.Command("ng", "add", "@angular/material@22", "--skip-confirmation")
	if cmd.Path != filepath.Join(oldGlobal, "ng") {
		t.Fatalf("precondition: expected initial cmd.Path=%s, got %s", filepath.Join(oldGlobal, "ng"), cmd.Path)
	}

	ConfigureNodeEnv(cmd)

	wantPath := filepath.Join(nvmBin, "ng")
	if cmd.Path != wantPath {
		t.Fatalf("ConfigureNodeEnv did not update cmd.Path: got %q, want %q", cmd.Path, wantPath)
	}
	if cmd.Err != nil {
		t.Fatalf("expected cmd.Err == nil, got %v", cmd.Err)
	}
}

func TestConfigureNodeEnvClearsErrNotFound(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix path test")
	}
	tmp := t.TempDir()
	fakeHome := filepath.Join(tmp, "home")
	nvmNodeRoot := filepath.Join(fakeHome, ".nvm", "versions", "node", "v22.23.2")
	nvmBin := createFakeNgSetup(t, nvmNodeRoot, "22.1.6")

	t.Setenv("HOME", fakeHome)
	t.Setenv("PATH", filepath.Join(tmp, "empty-bin"))

	cmd := exec.Command("ng", "version")
	if cmd.Err == nil {
		t.Fatalf("precondition: expected LookPath ErrNotFound when PATH is empty")
	}

	ConfigureNodeEnv(cmd)

	if cmd.Err != nil {
		t.Fatalf("expected ConfigureNodeEnv to clear cmd.Err, got %v", cmd.Err)
	}
	if cmd.Path != filepath.Join(nvmBin, "ng") {
		t.Fatalf("got cmd.Path=%q, want %q", cmd.Path, filepath.Join(nvmBin, "ng"))
	}
}

func TestNewNgNewCommandSelectsNg22OrFallsBackToNpx(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix path test")
	}
	tmp := t.TempDir()
	oldGlobal := createFakeNgSetup(t, filepath.Join(tmp, "usr-local"), "20.3.13")
	if err := os.WriteFile(filepath.Join(oldGlobal, "npx"), []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatal(err)
	}

	// Case 1: Only Angular CLI 20 is available -> must fall back to npx @angular/cli@22 with --zoneless=false
	emptyHome := filepath.Join(tmp, "empty-home")
	if err := os.MkdirAll(emptyHome, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", emptyHome)
	t.Setenv("PATH", oldGlobal)

	cmdFallback := NewNgNewCommand(tmp, "ng-helloworld", false)
	joinedFallback := strings.Join(cmdFallback.Args, " ")
	if !strings.Contains(joinedFallback, "@angular/cli@22 new ng-helloworld") || !strings.Contains(joinedFallback, "--zoneless=false") {
		t.Fatalf("expected npx @angular/cli@22 fallback with --zoneless=false, got Args=%v", cmdFallback.Args)
	}

	// Case 2: NVM provides Angular CLI 22 -> must invoke nvm's ng directly with --zoneless=false
	fakeHome := filepath.Join(tmp, "home")
	nvmNodeRoot := filepath.Join(fakeHome, ".nvm", "versions", "node", "v22.23.2")
	nvmBin := createFakeNgSetup(t, nvmNodeRoot, "22.1.6")
	t.Setenv("HOME", fakeHome)

	cmdDirect := NewNgNewCommand(tmp, "ng-helloworld", false)
	if cmdDirect.Path != filepath.Join(nvmBin, "ng") {
		t.Fatalf("expected direct nvm ng path %q, got %q", filepath.Join(nvmBin, "ng"), cmdDirect.Path)
	}
	joinedDirect := strings.Join(cmdDirect.Args, " ")
	if !strings.Contains(joinedDirect, "new ng-helloworld") || !strings.Contains(joinedDirect, "--zoneless=false") {
		t.Fatalf("expected ng new args with --zoneless=false, got Args=%v", cmdDirect.Args)
	}
}
