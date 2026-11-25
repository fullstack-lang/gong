//go:build !(js && wasm)

package runtime

import (
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func findLocalModuleInfo(modPath string) (localModDir string, goVersion string, replaces []replaceDirective) {
	if modPath == "" {
		return "", "", nil
	}

	var roots []string
	if cwd, err := os.Getwd(); err == nil {
		roots = append(roots, cwd)
	}
	if exe, err := os.Executable(); err == nil {
		roots = append(roots, filepath.Dir(exe))
	}
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		roots = append(roots, filepath.Join(gopath, "src", filepath.FromSlash(modPath)))
	} else if home, err := os.UserHomeDir(); err == nil {
		roots = append(roots, filepath.Join(home, "go", "src", filepath.FromSlash(modPath)))
	}

	for _, start := range roots {
		curr, err := filepath.Abs(start)
		if err != nil {
			continue
		}
		for range 15 {
			candidate := filepath.Join(curr, "go.mod")
			if buf, err := os.ReadFile(candidate); err == nil {
				if mf, err := modfile.Parse(candidate, buf, nil); err == nil && mf.Module != nil && mf.Module.Mod.Path == modPath {
					var gv string
					if mf.Go != nil {
						gv = mf.Go.Version
					}
					var reps []replaceDirective
					for _, r := range mf.Replace {
						newPath := r.New.Path
						if r.New.Version == "" && !filepath.IsAbs(newPath) {
							newPath = filepath.Clean(filepath.Join(curr, newPath))
						}
						reps = append(reps, replaceDirective{
							oldPath:    r.Old.Path,
							oldVersion: r.Old.Version,
							newPath:    newPath,
							newVersion: r.New.Version,
						})
					}
					return curr, gv, reps
				}
			}
			parent := filepath.Dir(curr)
			if parent == curr {
				break
			}
			curr = parent
		}
	}
	return "", "", nil
}

func finalizeDataGoModPlatform(dirPath string, goModChanged bool) {
	goSumPath := filepath.Join(dirPath, "go.sum")
	if !goModChanged {
		if _, err := os.Stat(goSumPath); err == nil {
			return
		}
	}
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dirPath
	_ = cmd.Run()
}
