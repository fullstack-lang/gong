package angular

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConfigureNodeEnv ensures that if an NVM Node installation is present,
// its bin directory is prepended to the command's PATH environment variable.
// This guarantees that ng and npm run with the required Node version (>= 22.22.3).
func ConfigureNodeEnv(cmd *exec.Cmd) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	nvmDir := filepath.Join(home, ".nvm/versions/node")
	entries, err := os.ReadDir(nvmDir)
	if err != nil || len(entries) == 0 {
		return
	}

	var latest string
	for _, entry := range entries {
		if entry.IsDir() {
			latest = entry.Name()
		}
	}

	if latest == "" {
		return
	}

	nvmBin := filepath.Join(nvmDir, latest, "bin")

	// Get base environment
	baseEnv := cmd.Env
	if len(baseEnv) == 0 {
		baseEnv = os.Environ()
	}

	var newEnv []string
	pathFound := false
	for _, e := range baseEnv {
		if strings.HasPrefix(e, "PATH=") {
			newEnv = append(newEnv, fmt.Sprintf("PATH=%s:%s", nvmBin, e[5:]))
			pathFound = true
		} else {
			newEnv = append(newEnv, e)
		}
	}

	if !pathFound {
		newEnv = append(newEnv, fmt.Sprintf("PATH=%s:%s", nvmBin, os.Getenv("PATH")))
	}

	cmd.Env = newEnv
}
