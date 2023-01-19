package models

import (
	"embed"
	"go/token"
	"log"
	"path/filepath"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
)

func LoadEmbeddedDiagramPackage(fs embed.FS, modelPkg *gong_models.ModelPkg) (diagramPackage *DiagramPackage, err error) {

	diagramPackage = (&DiagramPackage{}).Stage()
	diagramPackage.IsEditable = false
	diagramPackage.ModelPkg = modelPkg

	diagramPkgPath := filepath.Join(modelPkg.PkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage, _ = filepath.Abs(diagramPkgPath)
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	fset := new(token.FileSet)
	pkgsParser := gong_models.ParseEmbedModel(fs, "go/diagrams")
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		FillUpNodeTree(diagramPackage)
		Stage.Commit()
		return diagramPackage, nil
	}

	diagramPackage.ast = diagramPackageAst
	diagramPackage.fset = fset
	// load all diagram files
	for fileName := range diagramPackageAst.Files {
		diagramName := strings.TrimSuffix(filepath.Base(fileName), ".go")
		diagramPackage.UnmarshallOneDiagram(diagramName)
	}

	FillUpNodeTree(diagramPackage)
	return diagramPackage, nil
}
