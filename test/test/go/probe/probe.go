// generated code - do not edit
package probe

import (
	"embed"

	"github.com/gin-gonic/gin"

	gongsplit_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	gongdoc_load "github.com/fullstack-lang/gong/lib/doc/go/load"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/test/test/go/models"
	"github.com/fullstack-lang/gong/test/test/go/orm"
)

type Probe struct {
	r                  *gin.Engine
	stackPath          string
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	gongStage          *gong_models.StageStruct
	treeStage          *tree.StageStruct
	formStage          *form.StageStruct
	tableStage         *form.StageStruct
	splitStage         *split.StageStruct
}

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "" // for simplicity sake

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
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stackPath+ProbeTreeSidebarSuffix)

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+ProbeTableSuffix)
	tableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stackPath+ProbeFormSuffix)
	formStage.Commit()

	splitStage, _ := gongsplit_fullstack.NewStackInstance(r, stackPath+ProbeSplitSuffix)
	splitStage.Commit()

	probe = new(Probe)
	probe.r = r
	probe.stackPath = stackPath
	probe.stageOfInterest = stageOfInterest
	probe.backRepoOfInterest = backRepoOfInterest
	probe.gongStage = gongStage
	probe.treeStage = treeStage
	probe.formStage = formStage
	probe.tableStage = tableStage
	probe.splitStage = splitStage

	updateSplitStage(probe)
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
