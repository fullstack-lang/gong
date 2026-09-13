package models

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"math"
	"path/filepath"
	"slices"
	"strings"
)

// tolerantImporter wraps a standard types.Importer and returns a dummy package
// if the import fails, preventing missing external dependencies from crashing type analysis.
type tolerantImporter struct {
	underlying types.Importer
	packages   map[string]*types.Package
}

func newTolerantImporter(underlying types.Importer) *tolerantImporter {
	return &tolerantImporter{
		underlying: underlying,
		packages:   make(map[string]*types.Package),
	}
}

func (ti *tolerantImporter) Import(path string) (*types.Package, error) {
	if pkg, ok := ti.packages[path]; ok {
		return pkg, nil
	}
	if ti.underlying != nil {
		pkg, err := ti.underlying.Import(path)
		if err == nil {
			ti.packages[path] = pkg
			return pkg, nil
		}
	}
	// Fallback dummy package so missing/unbuilt imports never halt type analysis
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
	stubBuf.WriteString("type GongstructIF interface{}\n\n")

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

	conf := types.Config{
		IgnoreFuncBodies:         true,
		DisableUnusedImportCheck: true,
		Importer:                 newTolerantImporter(importer.Default()),
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
