package load

import (
	"embed"
	"path/filepath"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	// for diagrams
	"github.com/fullstack-lang/gong/lib/doc/go/adapter"
	"github.com/fullstack-lang/gong/lib/doc/go/diagrammer"

	gongdoc_fullstack "github.com/fullstack-lang/gong/lib/doc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"

	gongsvg_fullstack "github.com/fullstack-lang/gong/lib/svg/go/fullstack"

	gongtree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/gin-gonic/gin"
)

// Load have gongdoc init itself and the gong stack as well
// then parse the model source code in [goSourceDirectories]
// [stackName], for instance "gongsvg"
// of the gong stack [pathPath], for instance "github.com/fullstack-lang/gong/lib/svg/go/models"
// then parse the diagram.
// the diagram  can be embedded if [embeddedDiagrams] is true (possible if
// no edit is wished and if the binary need to be shipped as a standalone item)
// Number of instances if the working models can be computed to be
// displayed. This is stored in the [map_StructName_InstanceNb] parameter
func Load(
	stackName string,
	pkgPath string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	r *gin.Engine,
	embeddedDiagrams bool,
	map_StructName_InstanceNb map[string]int) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, pkgPath)
	gongdocStage, _ := gongdoc_fullstack.NewStackInstance(r, pkgPath)
	gongsvgStage, _ := gongsvg_fullstack.NewStackInstance(r, pkgPath)
	gongtreeStage, _ := gongtree_fullstack.NewStackInstance(r, pkgPath)

	diagramPackageCallbackSingloton := new(DiagramPackageCallbacksSingloton)
	diagramPackageCallbackSingloton.gongtreeStage = gongtreeStage
	diagramPackageCallbackSingloton.gongStage = gongStage
	gongdocStage.OnAfterDiagramPackageUpdateCallback = diagramPackageCallbackSingloton

	modelPackage, _ := gong_models.LoadEmbedded(gongStage, goModelsDir)

	modelPackage.Name = stackName
	modelPackage.PkgPath = pkgPath

	// create the diagrams
	// prepare the model views
	var diagramPackage *gongdoc_models.DiagramPackage

	gongdocStage.MetaPackageImportAlias = stackName
	gongdocStage.MetaPackageImportPath = pkgPath

	if embeddedDiagrams {
		diagramPackage, _ = LoadEmbeddedDiagramPackage(gongdocStage, gongtreeStage, goDiagramsDir, modelPackage)
	} else {
		diagramPackage, _ = LoadDiagramPackage(gongdocStage, gongtreeStage, filepath.Join("../../diagrams"), modelPackage, true)
	}
	diagramPackage.GongModelPath = pkgPath

	portfolioAdapter := adapter.NewPortfolioAdapter(gongStage, gongdocStage, gongsvgStage)
	modelAdapter := adapter.NewModelAdapter(portfolioAdapter)

	diagrammer := diagrammer.NewDiagrammer(modelAdapter, portfolioAdapter, gongtreeStage)
	portfolioAdapter.SetDiagrammer(diagrammer)

	treeOfModelObjects := (&gongtree_models.Tree{Name: gongdoc_models.Model.ToString()}).Stage(gongtreeStage)
	diagrammer.FillUpModelTree(treeOfModelObjects)

	treeOfPortfolioObjects := (&gongtree_models.Tree{Name: gongdoc_models.Portfolio.ToString()}).Stage(gongtreeStage)
	diagrammer.FillUpPortfolioUITree(treeOfPortfolioObjects)
	diagrammer.CommitTreeStage()

	// first, get all gong struct in the model
	for gongStruct := range gongStage.GongStructs {
		nbInstances, ok := map_StructName_InstanceNb[gongStruct.Name]
		if ok {
			diagramPackage.Map_Identifier_NbInstances[gongStruct.Name] = nbInstances
		}
	}

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming(gongStage, diagramPackage.Stage_)
	// end of the be removed

	for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](modelPackage.GetStage()) {
		diagramPackage.Map_Identifier_NbInstances[gongStruct.Name] =
			map_StructName_InstanceNb[gongStruct.Name]

	}

	for _, classdiagram := range diagramPackage.Classdiagrams {
		for _, classshape := range classdiagram.GongStructShapes {

			gongStructName := gongdoc_models.IdentifierToGongObjectName(classshape.Identifier)
			nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongStructName]

			if ok {
				classshape.ShowNbInstances = false
				classshape.NbInstances = nbInstances
			}
		}
	}
}
