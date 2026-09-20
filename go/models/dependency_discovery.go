package models

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
)

// DiscoverModelDependencies finds all internal direct and indirect model dependencies
// of rootPkgPath in topological order (leaves/dependencies first, root package last).
//
// Only packages within the same stack directory (e.g. under the parent of go/models)
// are considered for generation. External or vendored libraries (e.g. lib/split) are excluded.
func DiscoverModelDependencies(rootPkgPath string) ([]string, error) {
	rootAbs, err := filepath.Abs(rootPkgPath)
	if err != nil {
		return nil, err
	}

	// 1. Locate go.mod
	var goModDir string
	var modPath string
	curr := rootAbs
	for range 10 {
		goModFile := filepath.Join(curr, "go.mod")
		if buf, err := os.ReadFile(goModFile); err == nil {
			goModDir = curr
			if mf, err := modfile.Parse(goModFile, buf, nil); err == nil && mf.Module != nil {
				modPath = mf.Module.Mod.Path
			}
			break
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}

	if goModDir == "" || modPath == "" {
		return nil, nil
	}

	// 2. DFS for topological sort of subpackages strictly under rootAbs
	visited := make(map[string]bool)
	inStack := make(map[string]bool)
	var result []string

	var dfs func(pkgDir string) error
	dfs = func(pkgDir string) error {
		if inStack[pkgDir] {
			return fmt.Errorf("import cycle detected involving: %s", pkgDir)
		}
		if visited[pkgDir] {
			return nil
		}

		inStack[pkgDir] = true

		// Parse files in pkgDir to inspect imports
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, pkgDir, nil, parser.ImportsOnly)
		if err == nil {
			// Collect unique internal imports
			importSet := make(map[string]bool)
			for _, pkg := range pkgs {
				for filePath, file := range pkg.Files {
					base := filepath.Base(filePath)
					// Skip generated files and stager UI files
					if strings.HasPrefix(base, "zzz_gong") ||
						strings.HasPrefix(base, "gong") ||
						strings.HasPrefix(base, "stager") ||
						base == "docs.go" {
						continue
					}
					for _, imp := range file.Imports {
						importPath := strings.Trim(imp.Path.Value, "\"")
						if strings.HasPrefix(importPath, modPath) {
							importSet[importPath] = true
						}
					}
				}
			}

			// Sort dependencies deterministically
			var impPaths []string
			for impPath := range importSet {
				impPaths = append(impPaths, impPath)
			}
			sort.Strings(impPaths)

			// Traverse dependencies
			for _, impPath := range impPaths {
				relFromMod := strings.TrimPrefix(impPath, modPath)
				relFromMod = strings.TrimPrefix(relFromMod, "/")
				targetDir := filepath.Join(goModDir, filepath.FromSlash(relFromMod))

				targetAbs, _ := filepath.Abs(targetDir)
				if targetAbs == pkgDir {
					continue
				}

				// Must be a strict subdirectory under rootAbs (e.g. rootPkg/x, rootPkg/y)
				relToRoot, errRel := filepath.Rel(rootAbs, targetAbs)
				if errRel != nil || relToRoot == "." || strings.HasPrefix(relToRoot, "..") || filepath.IsAbs(relToRoot) {
					continue
				}

				// Exclude probe, diagrams, or other non-model subdirectories
				relSlash := filepath.ToSlash(relToRoot)
				if relSlash == "diagrams" || strings.HasPrefix(relSlash, "diagrams/") ||
					relSlash == "probe" || strings.HasPrefix(relSlash, "probe/") {
					continue
				}

				// Check that targetDir exists and contains model type declarations
				if hasModelDeclarations(targetAbs) {
					if err := dfs(targetAbs); err != nil {
						return err
					}
				}
			}
		}

		inStack[pkgDir] = false
		visited[pkgDir] = true

		// Append to result if it is not the root package itself
		if pkgDir != rootAbs {
			result = append(result, pkgDir)
		}
		return nil
	}

	if err := dfs(rootAbs); err != nil {
		return nil, err
	}

	return result, nil
}

// hasModelDeclarations checks whether dir contains docs.go and at least one non-generated Go file with a type declaration.
func hasModelDeclarations(dir string) bool {
	// A Gong model package must have docs.go
	if _, err := os.Stat(filepath.Join(dir, "docs.go")); err != nil {
		return false
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") || strings.HasPrefix(e.Name(), "zzz_gong") || strings.HasPrefix(e.Name(), "stager") || e.Name() == "docs.go" {
			continue
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, filepath.Join(dir, e.Name()), nil, 0)
		if err != nil {
			continue
		}
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok || genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				if _, ok := spec.(*ast.TypeSpec); ok {
					return true
				}
			}
		}
	}
	return false
}

