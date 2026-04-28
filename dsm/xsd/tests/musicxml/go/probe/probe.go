// generated code - do not edit
package probe

import (
	"embed"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong_models "github.com/fullstack-lang/gong/go/models"

	doc "github.com/fullstack-lang/gong/lib/doc/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go/models"

	embeddedgo "github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go"
)

type Probe struct {
	r                      *gin.Engine
	stageOfInterest        *models.Stage
	gongStage              *gong_models.Stage
	treeStage              *tree.Stage
	treeNavigationStage    *tree.Stage
	formStage              *form.Stage
	tableStage             *table.Stage
	notificationTableStage *table.Stage
	splitStage             *split.Stage

	// AsSplit to be used if one need only the data editor
	dataEditor *split.AsSplit

	// AsSplitArea for the diagram editor
	diagramEditor *split.AsSplitArea

	docStager *doc.Stager

	notification []*Notification

	// to limit the  number of elements per gong struct node in the tree
	maxElementsNbPerGongStructNode int

	// commit mode is used to control if the commit button is enabled or not.
	// It is set to false when the probe is in a state where the commit is not possible (for example when the stage is dirty and the commit would fail)
	commitMode bool

	// bulkDeleteMode is used to control if the bulk delete button has been clicked.
	bulkDeleteMode bool
}

func (probe *Probe) SetCommitMode(commitMode bool) {
	probe.commitMode = commitMode
}

func (probe *Probe) GetCommitMode() bool {
	return probe.commitMode
}

func (probe *Probe) SetMaxElementsNbPerGongStructNode(nb int) {
	probe.maxElementsNbPerGongStructNode = nb
}

func (probe *Probe) GetMaxElementsNbPerGongStructNode() int {
	return probe.maxElementsNbPerGongStructNode
}

func (probe *Probe) RefreshNavigationTree() {
	probe.ux_navigation_tree()
}

func NewProbe(
	r *gin.Engine,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.Stage) (probe *Probe) {

	// split stage for the whole probe
	splitStage, _ := split_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	// load the gong
	stage := gong_models.NewStage(stageOfInterest.GetName())
	gong_models.LoadEmbedded(stage, goModelsDir)

	// treeForSelectingDate that is on the sidebar
	treeStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName())
	treeNavigationStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNavigationTreeSidebarStageName())

	// stage for main table
	tableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName())
	tableStage.Commit()

	notificationTableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNotificationTableStageName())
	notificationTableStage.Commit()

	// stage for reusable form
	formStage, _ := form_fullstack.NewStackInstance(r, stageOfInterest.GetProbeFormStageName())
	formStage.Commit()

	probe = &Probe{
		r:                              r,
		stageOfInterest:                stageOfInterest,
		gongStage:                      stage,
		treeStage:                      treeStage,
		treeNavigationStage:            treeNavigationStage,
		formStage:                      formStage,
		tableStage:                     tableStage,
		notificationTableStage:         notificationTableStage,
		splitStage:                     splitStage,
		maxElementsNbPerGongStructNode: 10,
		commitMode:                     true,
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
		"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go"+":"+stageOfInterest.GetName(),
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		probe.diagramEditor,
		stageOfInterest.Map_GongStructName_InstancesNb,
	)

	probe.dataEditor = &split.AsSplit{
		Name:      "Top, sidebar, table & form",
		Direction: split.Horizontal,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name: "sidebar",
				Size: 20,
				AsSplit: &split.AsSplit{
					Direction:              split.Vertical,
					IsSizeInPixel:          true,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "sidebar tree",
							Size: 53, // to align on the top of the table
							Tree: &split.Tree{
								Name:      "Sidebar",
								StackName: probe.treeNavigationStage.GetName(),
							},
						},
						{
							Name:  "sidebar tree",
							IsAny: true,
							Tree: &split.Tree{
								Name:      "Sidebar",
								StackName: probe.treeStage.GetName(),
							},
						},
					},
				},
			},

			{
				Name: "both tables",
				Size: 50,
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

	probe.ux_tree()

	return
}

func (probe *Probe) Refresh() {
	probe.ux_tree()
	probe.ux_table()
	probe.ux_form()
	probe.docStager.Svg()
}

const NbNotificationMax = 100

func (probe *Probe) AddNotification(date time.Time, message string) {
	notification := Notification{
		Date:    date,
		Message: message,
	}
	probe.notification = append(probe.notification, &notification)

	if len(probe.notification) > NbNotificationMax {
		probe.notification = probe.notification[1:] // Drop the first element (index 0)
	}
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

func (probe *Probe) GetNavigationTreeStage() *tree.Stage {
	return probe.treeNavigationStage
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
