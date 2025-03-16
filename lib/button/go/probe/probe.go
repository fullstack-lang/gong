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

	"github.com/fullstack-lang/gong/lib/button/go/models"
	"github.com/fullstack-lang/gong/lib/button/go/orm"
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
const ProbeSplitSuffix = "-probe" // for simplicity sake

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


func updateSplitStage(probe *Probe) {

	probe.splitStage.Reset()

	mainView := (&split.View{
		Name: "Main view",
	}).Stage(probe.splitStage)

	topSplitArea := (&split.AsSplitArea{
		Name: "Top",
		Size: 50,
	}).Stage(probe.splitStage)
	mainView.RootAsSplitAreas = append(mainView.RootAsSplitAreas, topSplitArea)

	horizontalSplit := (&split.AsSplit{
		Name:      "Top, sidebar, table & form",
		Direction: split.Horizontal,
	}).Stage(probe.splitStage)
	topSplitArea.AsSplits = append(topSplitArea.AsSplits, horizontalSplit)

	sidebarArea := (&split.AsSplitArea{
		Name: "sidebar tree",
		Size: 20,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, sidebarArea)

	tree := (&split.Tree{
		Name:      "Sidebar",
		StackName: probe.stackPath + ProbeTreeSidebarSuffix,
		TreeName:  SideBarTreeName,
	}).Stage(probe.splitStage)
	sidebarArea.Tree = tree

	tableArea := (&split.AsSplitArea{
		Name: "table",
		Size: 50,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, tableArea)

	table := (&split.Table{
		Name:      "Table",
		StackName: probe.stackPath + ProbeTableSuffix,
		TableName: TableName,
	}).Stage(probe.splitStage)
	tableArea.Table = table

	formArea := (&split.AsSplitArea{
		Name: "form",
		Size: 30,
	}).Stage(probe.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, formArea)

	form := (&split.Form{
		Name:      "Form",
		StackName: probe.stackPath + ProbeFormSuffix,
		FormName:  FormName,
	}).Stage(probe.splitStage)
	formArea.Form = form

	bottomSplitArea := (&split.AsSplitArea{
		Name: "Bottom",
		Size: 50,
	}).Stage(probe.splitStage)
	mainView.RootAsSplitAreas = append(mainView.RootAsSplitAreas, bottomSplitArea)

	doc := (&split.Doc{
		Name:      "Doc",
		StackName: probe.stageOfInterest.GetType(),
	}).Stage(probe.splitStage)
	bottomSplitArea.Doc = doc

	probe.splitStage.Commit()
}

