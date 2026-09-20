package models

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// StageSetModel represents a StageSet struct in a models package coordinating
// multiple stages across packages for unified persistence.
type StageSetModel struct {
	Name     string
	Fields   []*StageSetField
	IsManual bool // true if defined manually by developer in a Go file
}

// StageSetField represents a field of the StageSet struct that points to a Stage.
type StageSetField struct {
	Name        string // Field name on StageSet, e.g. "Stage", "XStage", "YStage"
	PackageName string // Package name, e.g. "models", "x", "y"
	PackagePath string // Full package import path, e.g. "github.com/fullstack-lang/gong/test/test2/go/models/x"
	IsLocal     bool   // true if this stage belongs to the current package being generated
	ImportAlias string // deterministic synthetic import alias, e.g. "__stage_0__"
}

// SynthesizeStageSetFromDependencies automatically constructs a StageSetModel
// for modelPkg if it has internal subpackage dependencies discovered under pkgPath.
func (modelPkg *ModelPkg) SynthesizeStageSetFromDependencies(pkgPath string) error {
	depPkgPaths, err := DiscoverModelDependencies(pkgPath)
	if err != nil {
		return err
	}

	stageSet := &StageSetModel{
		Name:     "StageSet",
		IsManual: false,
	}

	rootPkgGoName := modelPkg.PkgGoName
	if rootPkgGoName == "" {
		rootPkgGoName = "models"
	}

	// Field 0: Local root package Stage
	localField := &StageSetField{
		Name:        "Stage",
		PackageName: rootPkgGoName,
		PackagePath: modelPkg.PkgPath,
		IsLocal:     true,
	}
	stageSet.Fields = append(stageSet.Fields, localField)

	// depPkgPaths are in topological order (leaves first, e.g. [y, x]).
	// Reverse topological order gives [x, y] (dependents first).
	caser := cases.Title(language.English)
	goModDir, modPath := FindGoMod(pkgPath)

	for i := len(depPkgPaths) - 1; i >= 0; i-- {
		depPath := depPkgPaths[i]

		// Verify that the dependency package is actually a Gong model package
		depStage := NewStage("")
		depModelPkg, errLoad := LoadSource(depStage, depPath)
		if errLoad != nil || (len(depModelPkg.GongStructs) == 0 && len(depModelPkg.GongEnums) == 0) {
			continue
		}

		// Determine package name from package declaration in source files
		fset := token.NewFileSet()
		pkgs, err := parser.ParseDir(fset, depPath, nil, parser.PackageClauseOnly)
		var depPkgName string
		if err == nil {
			for name := range pkgs {
				depPkgName = name
				break
			}
		}
		if depPkgName == "" {
			depPkgName = filepath.Base(depPath)
		}

		// Determine full package import path
		var depFullPkgPath string
		if goModDir != "" && modPath != "" {
			depAbs, _ := filepath.Abs(depPath)
			rel, errRel := filepath.Rel(goModDir, depAbs)
			if errRel == nil {
				depFullPkgPath = filepath.ToSlash(filepath.Join(modPath, rel))
			}
		}
		if depFullPkgPath == "" {
			_, depFullPkgPath = ComputePkgPathFromGoModFile(depPath)
		}

		depField := &StageSetField{
			Name:        caser.String(depPkgName) + "Stage",
			PackageName: depPkgName,
			PackagePath: depFullPkgPath,
			IsLocal:     false,
		}
		stageSet.Fields = append(stageSet.Fields, depField)
	}


	for idx, f := range stageSet.Fields {
		f.ImportAlias = fmt.Sprintf("__stage_%d__", idx)
	}

	modelPkg.StageSet = stageSet
	return nil
}

