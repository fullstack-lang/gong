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

	"github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Probe struct {
	r                  *gin.Engine
	stageOfInterest    *models.StageStruct
	gongStage          *gong_models.StageStruct
	treeStage          *tree.StageStruct
	formStage          *form.StageStruct
	tableStage         *form.StageStruct
	splitStage         *split.StageStruct
}

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.StageStruct) (probe *Probe) {

	gongStage, _ := gong_fullstack.NewStackInstance(r, stageOfInterest.GetName())

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName())

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName())
	tableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeFormStageName())
	formStage.Commit()

	splitStage, _ := gongsplit_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	probe = new(Probe)
	probe.r = r
	probe.stageOfInterest = stageOfInterest
	probe.gongStage = gongStage
	probe.treeStage = treeStage
	probe.formStage = formStage
	probe.tableStage = tableStage
	probe.splitStage = splitStage

	updateSplitStage(probe)
	fillUpTree(probe)

	gongdoc_load.Load(
		"",
		probe.stageOfInterest.GetProbeSplitStageName(),
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

	(&split.View{
		Name: "Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "Top",
				Size: 50,
				AsSplit: (&split.AsSplit{
					Name:      "Top, sidebar, table & form",
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "sidebar tree",
							Size: 20,
							Tree: (&split.Tree{
								Name:      "Sidebar",
								StackName: probe.treeStage.GetName(),
								TreeName:  SideBarTreeName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name: "table",
							Size: 50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: probe.tableStage.GetName(),
								TableName: TableName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
						(&split.AsSplitArea{
							Name: "form",
							Size: 30,
							Form: (&split.Form{
								Name:      "Form",
								StackName: probe.formStage.GetName(),
								FormName:  FormName,
							}).Stage(probe.splitStage),
						}).Stage(probe.splitStage),
					},
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
			(&split.AsSplitArea{
				Name: "Bottom",
				Size: 50,
				Doc: (&split.Doc{
					Name:      "Doc",
					StackName: probe.splitStage.GetName(),
				}).Stage(probe.splitStage),
			}).Stage(probe.splitStage),
		},
	}).Stage(probe.splitStage)

	probe.splitStage.Commit()
}
