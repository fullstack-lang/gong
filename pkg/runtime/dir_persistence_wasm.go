//go:build js && wasm

package runtime

func findLocalModuleInfo(modPath string) (localModDir string, goVersion string, replaces []replaceDirective) {
	return "", "", nil
}

func finalizeDataGoModPlatform(dirPath string, goModChanged bool) {
}
