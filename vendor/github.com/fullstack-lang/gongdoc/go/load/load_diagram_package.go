package load

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_node2gongdoc "github.com/fullstack-lang/gongdoc/go/node2gongdoc"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// LoadDiagramPackage fill up the stage with the diagrams elements
func LoadDiagramPackage(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	pkgPath string,
	modelPkg *gong_models.ModelPkg,
	editable bool,
) (diagramPackage *gongdoc_models.DiagramPackage, err error) {

	diagramPackage = (&gongdoc_models.DiagramPackage{
		Stage_: gongdocStage,
	}).Stage(gongdocStage)
	diagramPackage.Map_Identifier_NbInstances = make(map[string]int)
	diagramPackage.IsEditable = editable
	diagramPackage.ModelPkg = modelPkg
	diagramPackage.Name = modelPkg.Name + "_diagrams"

	// parse the diagram package
	diagramPkgPath := filepath.Join(pkgPath, "../diagrams")
	diagramPackage.AbsolutePathToDiagramPackage, _ = filepath.Abs(diagramPkgPath)
	diagramPackage.Path = diagramPkgPath
	diagramPackage.GongModelPath = modelPkg.PkgPath

	// diagram package, when marshalled, will reference identifiers in the
	// model package. Both of the variable need to be set up for the
	// generic marshalling/unmarshalling to work
	gongdocStage.MetaPackageImportAlias = "ref_" + filepath.Base(diagramPackage.GongModelPath)
	gongdocStage.MetaPackageImportPath = `"` + diagramPackage.GongModelPath + `"`
	if gongdocStage.Map_DocLink_Renaming == nil {
		gongdocStage.Map_DocLink_Renaming = make(map[string]gongdoc_models.GONG__Identifier)
	}

	// if diagrams directory does not exist create it
	_, errd := os.Stat(diagramPkgPath)
	if os.IsNotExist(errd) {
		log.Printf(diagramPkgPath, " does not exist, gongdoc creates it")

		errd := os.MkdirAll(diagramPkgPath, os.ModePerm)
		if os.IsNotExist(errd) {
			log.Println("creating directory : " + diagramPkgPath)
		}
		if os.IsExist(errd) {
			log.Println("directory " + diagramPkgPath + " allready exists")
		}
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	pkgsParser, errParser := parser.ParseDir(fset, diagramPkgPath, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ", errParser.Error())
	}
	diagramPackageAst, ok := pkgsParser["diagrams"]
	if !ok {
		diagramPackage.IsEditable = editable
		gongdoc_node2gongdoc.FillUpNodeTree(gongdocStage, gongtreeStage, diagramPackage)
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

	diagramPackage.IsEditable = editable
	gongdoc_node2gongdoc.FillUpNodeTree(gongdocStage, gongtreeStage, diagramPackage)
	gongdocStage.Commit()
	return diagramPackage, nil
}
