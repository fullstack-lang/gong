package load

import (
	"embed"
	"go/token"
	"log"
	"path/filepath"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"

	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func LoadEmbeddedDiagramPackage(
	gongdocStage *gongdoc_models.Stage,
	gongtreeStage *gongtree_models.Stage,
	goModelsDir embed.FS,
	modelPkg *gong_models.ModelPkg,
) (diagramPackage *gongdoc_models.DiagramPackage, err error) {

	diagramPackage = (&gongdoc_models.DiagramPackage{}).Stage(gongdocStage)
	diagramPackage.Map_Identifier_NbInstances = make(map[string]int)
	diagramPackage.IsEditable = false
	diagramPackage.ModelPkg = modelPkg
	diagramPackage.Stage_ = gongdocStage

	diagramPkgPath := filepath.Join(modelPkg.PkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage = "go/models"
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	fset := new(token.FileSet)
	pkgsParser := gong_models.ParseEmbedModel(goModelsDir, "diagrams")
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		gongdocStage.Commit()
		return diagramPackage, nil
	}

	diagramPackage.Ast = diagramPackageAst
	diagramPackage.Fset = fset
	// load all diagram files
	for diagramName, inFile := range diagramPackageAst.Files {

		diagramName := strings.TrimSuffix(filepath.Base(diagramName), ".go")
		diagramPackage.UnmarshallOneDiagram(gongdocStage, diagramName, inFile, fset)
	}

	return diagramPackage, nil
}
