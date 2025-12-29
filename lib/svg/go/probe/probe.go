// generated code - do not edit
package probe

import (
	"embed"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	gongsplit_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	gongtree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong_models "github.com/fullstack-lang/gong/go/models"

	doc "github.com/fullstack-lang/gong/lib/doc/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/svg/go/models"

	svg_go "github.com/fullstack-lang/gong/lib/svg/go"
)

type Probe struct {
	r                      *gin.Engine
	stageOfInterest        *models.Stage
	gongStage              *gong_models.Stage
	treeStage              *tree.Stage
	formStage              *form.Stage
	tableStage             *form.Stage
	notificationTableStage *form.Stage
	splitStage             *split.Stage

	// AsSplit to be used if one need only the data editor
	dataEditor *split.AsSplit

	// AsSplitArea for the diagram editor
	diagramEditor *split.AsSplitArea

	docStager *doc.Stager

	notification []*Notification
}

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.Stage) (probe *Probe) {

	// split stage for the whole probe
	splitStage, _ := gongsplit_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	// load the gong
	stage := gong_models.NewStage(stageOfInterest.GetName())
	gong_models.LoadEmbedded(stage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := gongtree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName())

	// stage for main table
	tableStage, _ := gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName())
	tableStage.Commit()

	notificationTableStage, _ := gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNotificationTableStageName())
	notificationTableStage.Commit()

	// stage for reusable form
	formStage, _ := gongtable_fullstack.NewStackInstance(r, stageOfInterest.GetProbeFormStageName())
	formStage.Commit()

	probe = &Probe{
		r:                      r,
		stageOfInterest:        stageOfInterest,
		gongStage:              stage,
		treeStage:              treeStage,
		formStage:              formStage,
		tableStage:             tableStage,
		notificationTableStage: notificationTableStage,
		splitStage:             splitStage,
	}

	// prepare the receiving AsSplitArea
	probe.diagramEditor = &split.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}

	probe.docStager = prepare.Prepare(
		r,
		embeddedDiagrams,

		// this is the prefix of the names of the stages svg and tree that will be created
		// by doc. Using a combination of the package name and the stage of interest name
		// might prevent name collisions if more that one probe is being instancied
		"github.com/fullstack-lang/gong/lib/svg/go"+":"+stageOfInterest.GetName(),
		svg_go.GoModelsDir,
		svg_go.GoDiagramsDir,
		probe.diagramEditor,
		stageOfInterest.Map_GongStructName_InstancesNb,
	)

	probe.dataEditor = &split.AsSplit{
		Name:      "Top, sidebar, table & form",
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name: "sidebar tree",
				Size: 20,
				Tree: &split.Tree{
					Name:      "Sidebar",
					StackName: probe.treeStage.GetName(),
				},
			},
			{
				Name: "both tables",
				AsSplit: &split.AsSplit{
					Direction: split.Vertical,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "table",
							Size: 50,
							Table: &split.Table{
								Name:      "Table",
								StackName: probe.tableStage.GetName(),
							},
						},
						{
							Name: "notification table",
							Size: 50,
							Table: &split.Table{
								Name:      "Table",
								StackName: probe.notificationTableStage.GetName(),
							},
						},
					},
				},
			},
			{
				Name: "form",
				Size: 30,
				Form: &split.Form{
					Name:      "Form",
					StackName: probe.formStage.GetName(),
				},
			},
		},
	}

	split.StageBranch(probe.splitStage, &split.View{
		Name: "Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:    "Top",
				Size:    50,
				AsSplit: probe.dataEditor,
			},
			probe.diagramEditor,
		},
	})
	probe.splitStage.Commit()

	updateAndCommitTree(probe)

	return
}

func (probe *Probe) Refresh() {
	updateAndCommitTree(probe)
	probe.docStager.UpdateAndCommitSVGStage()
}

func (probe *Probe) AddNotification(date time.Time, message string) {
	notification := Notification{
		Date:    date,
		Message: message,
	}
	probe.notification = append(probe.notification, &notification)
}

func (probe *Probe) CommitNotificationTable() {
	probe.UpdateAndCommitNotificationTable()
}

func (probe *Probe) ResetNotifications() {
	probe.notification = make([]*Notification, 0)
	probe.UpdateAndCommitNotificationTable()
}

func (probe *Probe) GetFormStage() *form.Stage {
	return probe.formStage
}

func (probe *Probe) GetDataEditor() *split.AsSplit {
	return probe.dataEditor
}

func (probe *Probe) GetDiagramEditor() *split.AsSplitArea {
	return probe.diagramEditor
}

func (probe *Probe) FillUpFormFromGongstruct(instance any, formName string) {
	FillUpFormFromGongstruct(instance, probe)
}

type Notification struct {
	Date    time.Time
	Message string
}
