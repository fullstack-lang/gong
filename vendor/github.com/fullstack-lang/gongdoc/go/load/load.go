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
// of the gong stack [stackName]
// then parse the diagram.
// the diagram  can be embedded if [embeddedDiagrams] is true (possible if
// no edit is wished and if the binary need to be shipped as a standalone item)
// Number of instances if the working models can be computed to be
// displayed. This is stored in the [map_StructName_InstanceNb] parameter
func Load(
	stackName string,
	goSourceDirectories embed.FS,
	r *gin.Engine,
	embeddedDiagrams bool,
	map_StructName_InstanceNb *map[string]int) {

	gong_fullstack.Init(r)
	gongdoc_fullstack.Init(r)
	modelPackage, _ := gong_models.LoadEmbedded(goSourceDirectories)

	// create the diagrams
	// prepare the model views
	var diagramPackage *gongdoc_models.DiagramPackage

	gongStage := gong_models.Stage
	_ = gongStage

	// first, get all gong struct in the model
	for gongStruct := range gong_models.Stage.GongStructs {
		// let create the gong struct in the gongdoc models
		// and put the numbre of instances
		reference := (&gongdoc_models.Reference{Name: gongStruct.Name}).Stage()
		reference.Type = gongdoc_models.REFERENCE_GONG_STRUCT

		nbInstances, ok := (*map_StructName_InstanceNb)[gongStruct.Name]
		if ok {
			reference.NbInstances = nbInstances
		}
	}
	if embeddedDiagrams {
		diagramPackage, _ = gongdoc_models.LoadEmbeddedDiagramPackage(goSourceDirectories, modelPackage)
	} else {
		diagramPackage, _ = gongdoc_models.LoadDiagramPackage(filepath.Join("../../diagrams"), modelPackage, true)
	}
	diagramPackage.GongModelPath = stackName + "/go/models"
}
