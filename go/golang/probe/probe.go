package probe

const ProbeTemplate = `// generated code - do not edit
package probe

import (
	"embed"

	"github.com/gin-gonic/gin"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	form "github.com/fullstack-lang/gongtable/go/models"
	tree "github.com/fullstack-lang/gongtree/go/models"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

type Probe struct {
	r                  *gin.Engine
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	gongStage          *gong_models.StageStruct
	treeStage          *tree.StageStruct
	formStage          *form.StageStruct
	tableStage         *form.StageStruct
}

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stackPath string,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct) (probe *Probe) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stackPath+"-sidebar")

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+"-table")
	tableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+"-form")
	formStage.Commit()

	probe = new(Probe)
	probe.r = r
	probe.stageOfInterest = stageOfInterest
	probe.backRepoOfInterest = backRepoOfInterest
	probe.gongStage = gongStage
	probe.treeStage = treeStage
	probe.formStage = formStage
	probe.tableStage = tableStage

	fillUpTree(probe)

	gongdoc_load.Load(
		"",
		probe.stageOfInterest.GetType(),
		goModelsDir,
		goDiagramsDir,
		r,
		embeddedDiagrams,
		&stageOfInterest.Map_GongStructName_InstancesNb)

	return
}

func (probe *Probe) Refresh() {
	fillUpTree(probe)
}

func (probe *Probe) GetFormStage() *form.StageStruct {
	return probe.formStage
}
`
