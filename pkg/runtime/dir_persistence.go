package runtime

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"unicode"

	"golang.org/x/mod/modfile"
)

// IsDirectoryPath reports whether targetPath represents a directory persistence target.
// It returns true if targetPath already exists on disk as a directory, or if it ends
// with a path separator ('/' or filepath.Separator).
func IsDirectoryPath(targetPath string) bool {
	if targetPath == "" {
		return false
	}
	if fi, err := os.Stat(targetPath); err == nil {
		return fi.IsDir()
	}
	return strings.HasSuffix(targetPath, "/") || strings.HasSuffix(targetPath, string(filepath.Separator))
}

// SanitizeModuleName derives a valid Go module name from dirPath's base directory name.
func SanitizeModuleName(dirPath string) string {
	cleaned := filepath.Clean(strings.TrimRight(dirPath, "/"+string(filepath.Separator)))
	base := filepath.Base(cleaned)
	if base == "" || base == "." || base == "/" || base == string(filepath.Separator) {
		return "gong_data"
	}

	var b strings.Builder
	for i, r := range base {
		if unicode.IsLetter(r) || r == '_' || (i > 0 && (unicode.IsDigit(r) || r == '-' || r == '.')) {
			b.WriteRune(unicode.ToLower(r))
		} else if unicode.IsDigit(r) && i == 0 {
			b.WriteString("m_")
			b.WriteRune(r)
		} else {
			b.WriteByte('_')
		}
	}
	res := strings.Trim(b.String(), "_.-")
	if res == "" {
		return "gong_data"
	}
	return res
}

func detectRuntimeGoVersion() string {
	v := strings.TrimPrefix(runtime.Version(), "go")
	parts := strings.Split(v, ".")
	if len(parts) >= 2 {
		// Strip any prerelease suffix on minor version
		minor := parts[1]
		for i, r := range minor {
			if !unicode.IsDigit(r) {
				minor = minor[:i]
				break
			}
		}
		if minor != "" {
			return parts[0] + "." + minor
		}
	}
	return "1.24"
}

// EnsureDataGoMod ensures that dirPath contains a valid go.mod requiring modPath,
// preserving any existing module name, requirements, and replace directives.
// It returns true if go.mod was newly created or modified.
func EnsureDataGoMod(dirPath string, modPath string) (changed bool, err error) {
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return false, err
	}
	if modPath == "" {
		return false, nil
	}

	goModPath := filepath.Join(dirPath, "go.mod")
	var mf *modfile.File
	existingBytes, readErr := os.ReadFile(goModPath)
	if readErr == nil {
		mf, err = modfile.Parse(goModPath, existingBytes, nil)
		if err != nil {
			return false, err
		}
	} else {
		mf = new(modfile.File)
		if err = mf.AddModuleStmt(SanitizeModuleName(dirPath)); err != nil {
			return false, err
		}
	}

	// Determine version and optional local module directory for replace
	reqVersion := ""
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if buildInfo.Main.Path == modPath && buildInfo.Main.Version != "" && buildInfo.Main.Version != "(devel)" {
			reqVersion = strings.ReplaceAll(buildInfo.Main.Version, "+dirty", "")
		} else {
			for _, dep := range buildInfo.Deps {
				if dep.Path == modPath && dep.Version != "" && dep.Version != "(devel)" {
					reqVersion = dep.Version
					break
				}
			}
		}
	}

	localModDir, localGoVer, localReplaces := findLocalModuleInfo(modPath)

	if mf.Go == nil {
		goVer := localGoVer
		if goVer == "" {
			goVer = detectRuntimeGoVersion()
		}
		_ = mf.AddGoStmt(goVer)
	}

	if reqVersion == "" {
		reqVersion = "v0.0.0"
	}

	hasRequire := false
	for _, req := range mf.Require {
		if req.Mod.Path == modPath {
			hasRequire = true
			break
		}
	}
	if !hasRequire {
		mf.AddNewRequire(modPath, reqVersion, false)
	}

	if localModDir != "" && reqVersion == "v0.0.0" {
		hasReplace := false
		for _, rep := range mf.Replace {
			if rep.Old.Path == modPath {
				hasReplace = true
				break
			}
		}
		if !hasReplace {
			_ = mf.AddReplace(modPath, "", localModDir, "")
		}
		for _, lr := range localReplaces {
			already := false
			for _, rep := range mf.Replace {
				if rep.Old.Path == lr.oldPath && rep.Old.Version == lr.oldVersion {
					already = true
					break
				}
			}
			if !already {
				_ = mf.AddReplace(lr.oldPath, lr.oldVersion, lr.newPath, lr.newVersion)
			}
		}
	}

	mf.Cleanup()
	formatted, fmtErr := mf.Format()
	if fmtErr != nil {
		return false, fmtErr
	}

	if readErr == nil && string(existingBytes) == string(formatted) {
		return false, nil
	}

	if err = os.WriteFile(goModPath, formatted, 0644); err != nil {
		return false, err
	}
	return true, nil
}

type replaceDirective struct {
	oldPath    string
	oldVersion string
	newPath    string
	newVersion string
}

// FinalizeDataGoMod runs post-write module finalization (such as `go mod tidy` on native platforms)
// when go.mod changed or go.sum is not yet initialized.
func FinalizeDataGoMod(dirPath string, goModChanged bool) {
	finalizeDataGoModPlatform(dirPath, goModChanged)
}
