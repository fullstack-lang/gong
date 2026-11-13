package angular

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// AngularMajorVersion is the Angular major version targeted by gong stack-height=4 generation.
const AngularMajorVersion = "22"

func parseSemverDir(name string) (major, minor, patch int, ok bool) {
	trimmed := strings.TrimPrefix(name, "v")
	parts := strings.SplitN(trimmed, ".", 3)
	if len(parts) != 3 {
		return 0, 0, 0, false
	}
	var err error
	major, err = strconv.Atoi(parts[0])
	if err != nil || major < 0 {
		return 0, 0, 0, false
	}
	minor, err = strconv.Atoi(parts[1])
	if err != nil || minor < 0 {
		return 0, 0, 0, false
	}
	// Strip any pre-release suffix on patch if present (e.g. "2-rc1")
	patchStr := parts[2]
	if idx := strings.IndexAny(patchStr, "-+"); idx != -1 {
		patchStr = patchStr[:idx]
	}
	patch, err = strconv.Atoi(patchStr)
	if err != nil || patch < 0 {
		return 0, 0, 0, false
	}
	return major, minor, patch, true
}

func selectLatestNodeVersionDir(entries []os.DirEntry) string {
	var bestSemver string
	bestMaj, bestMin, bestPat := -1, -1, -1
	var fallback string

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		fallback = name
		maj, min, pat, ok := parseSemverDir(name)
		if !ok {
			continue
		}
		if maj > bestMaj ||
			(maj == bestMaj && min > bestMin) ||
			(maj == bestMaj && min == bestMin && pat > bestPat) {
			bestMaj, bestMin, bestPat = maj, min, pat
			bestSemver = name
		}
	}
	if bestSemver != "" {
		return bestSemver
	}
	return fallback
}

func findLatestNvmBin() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	nvmDir := filepath.Join(home, ".nvm/versions/node")
	entries, err := os.ReadDir(nvmDir)
	if err != nil || len(entries) == 0 {
		return ""
	}

	latest := selectLatestNodeVersionDir(entries)
	if latest == "" {
		return ""
	}

	nvmBin := filepath.Join(nvmDir, latest, "bin")
	if info, err := os.Stat(nvmBin); err == nil && info.IsDir() {
		return nvmBin
	}
	return ""
}

func isExecutableFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		return false
	}
	if runtime.GOOS != "windows" {
		return info.Mode()&0o111 != 0
	}
	return true
}

func resolveExecutableInDirs(name string, dirs []string) string {
	var suffixes []string
	if runtime.GOOS == "windows" {
		suffixes = []string{"", ".cmd", ".exe", ".bat"}
	} else {
		suffixes = []string{""}
	}

	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		for _, suffix := range suffixes {
			candidate := filepath.Join(dir, name+suffix)
			if isExecutableFile(candidate) {
				return candidate
			}
		}
	}
	return ""
}

// isNgBinaryMajorVersion inspects an `ng` executable path to check whether it belongs
// to @angular/cli of the requested major version without spawning a subprocess when possible.
func isNgBinaryMajorVersion(ngBinPath string, wantMajor string) bool {
	if ngBinPath == "" {
		return false
	}

	var candidatePkgJsons []string
	if realPath, err := filepath.EvalSymlinks(ngBinPath); err == nil {
		// Standard npm global/local layout: .../@angular/cli/bin/ng.js -> ../package.json
		candidatePkgJsons = append(candidatePkgJsons,
			filepath.Join(filepath.Dir(realPath), "..", "package.json"),
		)
	}
	binDir := filepath.Dir(ngBinPath)
	candidatePkgJsons = append(candidatePkgJsons,
		filepath.Join(binDir, "..", "lib", "node_modules", "@angular", "cli", "package.json"),
		filepath.Join(binDir, "..", "@angular", "cli", "package.json"),
		filepath.Join(binDir, "node_modules", "@angular", "cli", "package.json"),
	)

	type pkgMeta struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	for _, pkgJsonPath := range candidatePkgJsons {
		data, err := os.ReadFile(pkgJsonPath)
		if err != nil {
			continue
		}
		var meta pkgMeta
		if err := json.Unmarshal(data, &meta); err != nil {
			continue
		}
		if meta.Name == "@angular/cli" && meta.Version != "" {
			return strings.HasPrefix(meta.Version, wantMajor+".")
		}
	}

	return false
}

// ConfigureNodeEnv ensures that if an NVM Node installation is present,
// its bin directory is prepended to the command's PATH environment variable,
// and re-resolves cmd.Path (clearing cmd.Err) against workspace-local node_modules/.bin
// and the updated PATH so Go's earlier exec.Command LookPath does not lock in a stale
// system binary (such as an older /usr/local/bin/ng).
func ConfigureNodeEnv(cmd *exec.Cmd) {
	nvmBin := findLatestNvmBin()

	baseEnv := cmd.Env
	if len(baseEnv) == 0 {
		baseEnv = os.Environ()
	}

	effectivePath := os.Getenv("PATH")
	var newEnv []string
	pathFound := false
	for _, e := range baseEnv {
		if strings.HasPrefix(e, "PATH=") {
			existingPath := e[5:]
			if nvmBin != "" {
				effectivePath = fmt.Sprintf("%s%c%s", nvmBin, os.PathListSeparator, existingPath)
			} else {
				effectivePath = existingPath
			}
			newEnv = append(newEnv, "PATH="+effectivePath)
			pathFound = true
		} else {
			newEnv = append(newEnv, e)
		}
	}

	if !pathFound {
		if nvmBin != "" {
			effectivePath = fmt.Sprintf("%s%c%s", nvmBin, os.PathListSeparator, effectivePath)
		}
		newEnv = append(newEnv, "PATH="+effectivePath)
	}

	cmd.Env = newEnv

	// Re-resolve bare executable names (e.g. "ng", "npm", "npx") against workspace-local
	// node_modules/.bin and the updated PATH, since exec.Command already ran LookPath
	// against the parent process's unmodified PATH.
	if len(cmd.Args) > 0 && filepath.Base(cmd.Args[0]) == cmd.Args[0] {
		name := cmd.Args[0]
		var searchDirs []string

		if cmd.Dir != "" {
			// Check workspace-local and parent npm-workspace-root node_modules/.bin
			localBinDirs := []string{
				filepath.Join(cmd.Dir, "node_modules", ".bin"),
				filepath.Join(cmd.Dir, "..", "node_modules", ".bin"),
			}
			for _, binDir := range localBinDirs {
				if name == "ng" {
					if candidate := resolveExecutableInDirs(name, []string{binDir}); candidate != "" {
						if isNgBinaryMajorVersion(candidate, AngularMajorVersion) {
							searchDirs = append(searchDirs, binDir)
						}
					}
				} else {
					searchDirs = append(searchDirs, binDir)
				}
			}
		}

		searchDirs = append(searchDirs, filepath.SplitList(effectivePath)...)
		if resolved := resolveExecutableInDirs(name, searchDirs); resolved != "" {
			cmd.Path = resolved
			cmd.Err = nil
		}
	}
}

// NewNgNewCommand builds the command to scaffold a new Angular workspace targeting
// AngularMajorVersion (with --zoneless=false so zone.js matches NgFileAppConfig).
// If a local/NVM ng binary of AngularMajorVersion is available, it is invoked directly;
// otherwise it falls back to `npx --yes @angular/cli@<AngularMajorVersion>`.
func NewNgNewCommand(dir string, workspaceName string, skipNpmInstall bool) *exec.Cmd {
	newArgs := []string{"new", workspaceName, "--defaults=true", "--minimal=true", "--zoneless=false"}
	if skipNpmInstall {
		newArgs = append(newArgs, "--skip-install")
	}

	probe := exec.Command("ng", "--version")
	probe.Dir = dir
	ConfigureNodeEnv(probe)

	if probe.Err == nil && probe.Path != "" && isNgBinaryMajorVersion(probe.Path, AngularMajorVersion) {
		cmd := exec.Command("ng", newArgs...)
		cmd.Dir = dir
		ConfigureNodeEnv(cmd)
		return cmd
	}

	npxArgs := append([]string{"--yes", "@angular/cli@" + AngularMajorVersion}, newArgs...)
	cmd := exec.Command("npx", npxArgs...)
	cmd.Dir = dir
	ConfigureNodeEnv(cmd)
	return cmd
}
