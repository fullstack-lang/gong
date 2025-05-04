// generated code - do not edit
package probe

import (
	"embed"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/lib/doc2/go/prepare"
	gongsplit_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"

	slider_go "github.com/fullstack-lang/gong/lib/slider/go"
)

type Probe struct {
	r               *gin.Engine
	stageOfInterest *models.Stage
	gongStage       *gong_models.Stage
	treeStage       *tree.Stage
	formStage       *form.Stage
	tableStage      *form.Stage
	splitStage      *split.Stage
}

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.Stage,
	pathToDiagramFile string) (probe *Probe) {

	// split stage for the whole probe
	splitStage, _ := gongsplit_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	// load the gong
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

	probe = new(Probe)
	probe.r = r
	probe.stageOfInterest = stageOfInterest
	probe.gongStage = gongStage
	probe.treeStage = treeStage
	probe.formStage = formStage
	probe.tableStage = tableStage
	probe.splitStage = splitStage

	// prepare the receiving AsSplitArea
	receivingAsSplitArea := &split.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}

	prepare.Prepare(
		r,
		embeddedDiagrams,
		pathToDiagramFile,
		stageOfInterest.GetName(),
		slider_go.GoModelsDir,
		slider_go.GoDiagramsDir,
		receivingAsSplitArea,
	)

	split.StageBranch(probe.splitStage, &split.View{
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
							}),
						}),
						(&split.AsSplitArea{
							Name: "table",
							Size: 50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: probe.tableStage.GetName(),
								TableName: TableName,
							}),
						}),
						(&split.AsSplitArea{
							Name: "form",
							Size: 30,
							Form: (&split.Form{
								Name:      "Form",
								StackName: probe.formStage.GetName(),
								FormName:  FormName,
							}),
						}),
					},
				}),
			}),
			receivingAsSplitArea,
		},
	})
	probe.splitStage.Commit()

	updateAndCommitTree(probe)

	return
}

func (probe *Probe) Refresh() {
	updateAndCommitTree(probe)
}

func (probe *Probe) GetFormStage() *form.Stage {
	return probe.formStage
}
