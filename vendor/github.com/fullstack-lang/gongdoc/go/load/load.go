package load

import (
	"embed"
	"path/filepath"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	// for diagrams
	gongdoc_fullstack "github.com/fullstack-lang/gongdoc/go/fullstack"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

	"github.com/gin-gonic/gin"
)

// Load have gongdoc init itself and the gong stack as well
// then parse the model source code in [goSourceDirectories]
// [stackName], for instance "gongsvg"
// of the gong stack [pathPath], for instance "github.com/fullstack-lang/gongsvg/go/models"
// then parse the diagram.
// the diagram  can be embedded if [embeddedDiagrams] is true (possible if
// no edit is wished and if the binary need to be shipped as a standalone item)
// Number of instances if the working models can be computed to be
// displayed. This is stored in the [map_StructName_InstanceNb] parameter
func Load(
	stackName string,
	pkgPath string,
	goSourceDirectories embed.FS,
	r *gin.Engine,
	embeddedDiagrams bool,
	map_StructName_InstanceNb *map[string]int) {

	gongStage := gong_fullstack.NewStackInstance(r, "")
	gongdoc_fullstack.NewStackInstance(r, "")
	modelPackage, _ := gong_models.LoadEmbedded(gongStage, goSourceDirectories)
	modelPackage.Name = stackName
	modelPackage.PkgPath = pkgPath

	// create the diagrams
	// prepare the model views
	var diagramPackage *gongdoc_models.DiagramPackage

	gongdoc_models.GetDefaultStage().MetaPackageImportAlias = stackName
	gongdoc_models.GetDefaultStage().MetaPackageImportPath = pkgPath

	if embeddedDiagrams {
		diagramPackage, _ = LoadEmbeddedDiagramPackage(goSourceDirectories, modelPackage)
	} else {
		diagramPackage, _ = LoadDiagramPackage(filepath.Join("../../diagrams"), modelPackage, true)
	}
	diagramPackage.GongModelPath = pkgPath

	// first, get all gong struct in the model
	for gongStruct := range gongStage.GongStructs {
		nbInstances, ok := (*map_StructName_InstanceNb)[gongStruct.Name]
		if ok {
			diagramPackage.Map_Identifier_NbInstances[gongStruct.Name] = nbInstances
		}
	}

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming(diagramPackage.Stage_)
	// end of the be removed

	// set up the number of instance per classshape
	if map_StructName_InstanceNb != nil {

		for gongStruct := range *gong_models.GetGongstructInstancesSet[gong_models.GongStruct]() {
			diagramPackage.Map_Identifier_NbInstances[gongStruct.Name] =
				(*map_StructName_InstanceNb)[gongStruct.Name]

		}

		for _, classdiagram := range diagramPackage.Classdiagrams {
			for _, classshape := range classdiagram.GongStructShapes {

				gongStructName := gongdoc_models.IdentifierToGongObjectName(classshape.Identifier)
				nbInstances, ok := diagramPackage.Map_Identifier_NbInstances[gongStructName]

				if ok {
					classshape.ShowNbInstances = true
					classshape.NbInstances = nbInstances
				}
			}
		}
	}
}
