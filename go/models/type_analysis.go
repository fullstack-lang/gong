package models

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"golang.org/x/mod/modfile"
)

// findGoMod traverses upwards to locate the nearest go.mod and extracts the module path.
func findGoMod(startDir string) (goModDir string, modPath string) {
	curr, err := filepath.Abs(startDir)
	if err != nil {
		return "", ""
	}
	for range 15 {
		goModFile := filepath.Join(curr, "go.mod")
		if buf, err := os.ReadFile(goModFile); err == nil {
			if mf, err := modfile.Parse(goModFile, buf, nil); err == nil && mf.Module != nil {
				return curr, mf.Module.Mod.Path
			}
			return curr, ""
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return "", ""
}

// tolerantImporter wraps a standard types.Importer, attempts to resolve internal
// workspace packages from source if the underlying importer fails, and returns a dummy package
// if the import fails, preventing missing external dependencies from crashing type analysis.
type tolerantImporter struct {
	underlying types.Importer
	packages   map[string]*types.Package
	fset       *token.FileSet
	goModDir   string
	modPath    string
}

func newTolerantImporter(underlying types.Importer, fset *token.FileSet, goModDir string, modPath string) *tolerantImporter {
	return &tolerantImporter{
		underlying: underlying,
		packages:   make(map[string]*types.Package),
		fset:       fset,
		goModDir:   goModDir,
		modPath:    modPath,
	}
}

func (ti *tolerantImporter) Import(path string) (*types.Package, error) {
	if pkg, ok := ti.packages[path]; ok {
		return pkg, nil
	}

	// 1. Try underlying importer (standard library, pre-built packages)
	if ti.underlying != nil {
		pkg, err := ti.underlying.Import(path)
		if err == nil {
			ti.packages[path] = pkg
			return pkg, nil
		}
	}

	// 2. Try loading internal module package from source
	if ti.goModDir != "" && ti.modPath != "" && strings.HasPrefix(path, ti.modPath) {
		rel := strings.TrimPrefix(path, ti.modPath)
		rel = strings.TrimPrefix(rel, "/")
		pkgDir := filepath.Join(ti.goModDir, filepath.FromSlash(rel))

		if fi, err := os.Stat(pkgDir); err == nil && fi.IsDir() {
			entries, err := os.ReadDir(pkgDir)
			if err == nil {
				var astFiles []*ast.File
				for _, entry := range entries {
					name := entry.Name()
					if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
						continue
					}
					// Skip generated gong files and documentation files
					if slices.Contains(GeneratedModelFiles, name) || strings.HasPrefix(name, "zzz_gong") || name == "docs.go" {
						continue
					}
					filePath := filepath.Join(pkgDir, name)
					f, err := parser.ParseFile(ti.fset, filePath, nil, parser.ParseComments)
					if err == nil {
						astFiles = append(astFiles, f)
					}
				}

				if len(astFiles) > 0 {
					subConf := types.Config{
						IgnoreFuncBodies:         true,
						DisableUnusedImportCheck: true,
						Importer:                 ti, // recursive for transitive imports
						Error:                    func(err error) {},
					}
					// Pre-cache package to prevent infinite recursion on import cycles
					dummy := types.NewPackage(path, astFiles[0].Name.Name)
					ti.packages[path] = dummy

					subPkg, _ := subConf.Check(path, ti.fset, astFiles, nil)
					if subPkg != nil {
						ti.packages[path] = subPkg
						return subPkg, nil
					}
				}
			}
		}
	}

	// 3. Fallback dummy package so missing/unbuilt imports never halt type analysis
	pkg := types.NewPackage(path, filepath.Base(path))
	pkg.MarkComplete()
	ti.packages[path] = pkg
	return pkg, nil
}

// RunTypeAnalysis performs resilient type checking on the parsed AST files.
// It injects synthetic stubs for missing framework symbols (Stage, StageStruct, .Stage() methods)
// and uses an error-tolerant configuration so incomplete code does not cause failures.
func RunTypeAnalysis(modelPkg *ModelPkg, astPackage *ast.Package) {
	if astPackage == nil || len(astPackage.Files) == 0 {
		return
	}

	if modelPkg.Fset == nil {
		modelPkg.Fset = token.NewFileSet()
	}

	// Determine minFilePathLength to exclude files in subdirectories
	minFilePathLength := math.MaxInt
	for filePath := range astPackage.Files {
		directories := make([]string, 0)
		workingFilePath := filePath
		for {
			dir := filepath.Dir(workingFilePath)
			if dir == workingFilePath {
				break
			}
			directories = append(directories, dir)
			workingFilePath = dir
		}
		if len(directories) < minFilePathLength {
			minFilePathLength = len(directories)
		}
	}

	// Filter user AST files (exclude generated files and subdirectories)
	var userAstFiles []*ast.File
	for filePath, file := range astPackage.Files {
		if slices.Contains(GeneratedModelFiles, filepath.Base(filePath)) {
			continue
		}

		directories := make([]string, 0)
		workingFilePath := filePath
		for {
			dir := filepath.Dir(workingFilePath)
			if dir == workingFilePath {
				break
			}
			directories = append(directories, dir)
			workingFilePath = dir
		}
		if len(directories) > minFilePathLength {
			continue
		}

		userAstFiles = append(userAstFiles, file)
	}

	// Synthesize in-memory stubs for Stage, StageStruct, GongstructIF, and .Stage() methods
	var stubBuf strings.Builder
	pkgName := astPackage.Name
	if pkgName == "" {
		pkgName = "models"
	}
	stubBuf.WriteString("package " + pkgName + "\n\n")
	stubBuf.WriteString("type Stage struct{}\n")
	stubBuf.WriteString("type StageStruct = Stage\n")
	stubBuf.WriteString("type GongstructIF any\n\n")

	for _, gongStruct := range modelPkg.GongStructs {
		stubBuf.WriteString(fmt.Sprintf("func (*%s) Stage(...*Stage) *%s { return nil }\n", gongStruct.Name, gongStruct.Name))
		stubBuf.WriteString(fmt.Sprintf("func (*%s) Unstage(...*Stage) *%s { return nil }\n", gongStruct.Name, gongStruct.Name))
	}

	stubFile, err := parser.ParseFile(modelPkg.Fset, "synthetic_gong_stubs.go", stubBuf.String(), 0)
	if err == nil {
		userAstFiles = append(userAstFiles, stubFile)
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}

	var sampleDir string
	for filePath := range astPackage.Files {
		sampleDir = filepath.Dir(filePath)
		break
	}
	goModDir, modPath := findGoMod(sampleDir)

	conf := types.Config{
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
		Importer:                 newTolerantImporter(importer.Default(), modelPkg.Fset, goModDir, modPath),
		Error: func(err error) {
			modelPkg.TypeErrors = append(modelPkg.TypeErrors, err)
		},
	}

	checkPkgPath := modelPkg.PkgPath
	if checkPkgPath == "" {
		checkPkgPath = pkgName
	}

	typesPkg, _ := conf.Check(checkPkgPath, modelPkg.Fset, userAstFiles, info)
	modelPkg.TypesPkg = typesPkg
	modelPkg.TypesInfo = info

	// Link GongStructs back to modelPkg and populate ImplementedInterfaces
	for _, gongStruct := range modelPkg.GongStructs {
		gongStruct.ModelPkg = modelPkg
		gongStruct.ImplementedInterfaces = modelPkg.GetImplementedInterfaces(gongStruct.Name)
	}
}
