package load

import (
	"embed"
	"go/token"
	"log"
	"path/filepath"
	"strings"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_node2gongdoc "github.com/fullstack-lang/gongdoc/go/node2gongdoc"
)

func LoadEmbeddedDiagramPackage(fs embed.FS, modelPkg *gong_models.ModelPkg) (diagramPackage *gongdoc_models.DiagramPackage, err error) {

	diagramPackage = (&gongdoc_models.DiagramPackage{}).Stage()
	diagramPackage.Map_Identifier_NbInstances = make(map[string]int)
	diagramPackage.IsEditable = false
	diagramPackage.ModelPkg = modelPkg

	diagramPkgPath := filepath.Join(modelPkg.PkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage = "go/models"
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	fset := new(token.FileSet)
	pkgsParser := gong_models.ParseEmbedModel(fs, "go/diagrams")
	if len(pkgsParser) != 1 {
		log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		gongdoc_node2gongdoc.FillUpNodeTree(diagramPackage)
		gongdoc_models.Stage.Commit()
		return diagramPackage, nil
	}

	diagramPackage.Ast = diagramPackageAst
	diagramPackage.Fset = fset
	// load all diagram files
	for diagramName, inFile := range diagramPackageAst.Files {

		diagramName := strings.TrimSuffix(filepath.Base(diagramName), ".go")
		diagramPackage.UnmarshallOneDiagram(diagramName, inFile, fset)
	}

	gongdoc_node2gongdoc.FillUpNodeTree(diagramPackage)
	return diagramPackage, nil
}
