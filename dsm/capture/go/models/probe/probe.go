// generated code - do not edit
package probe

import (
	"embed"
	"time"

	"net/http"

	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	doc "github.com/fullstack-lang/gong/lib/doc/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/dsm/capture/go/models"

	embeddedgo "github.com/fullstack-lang/gong/dsm/capture/go"
)

type Probe struct {
	r                      *http.ServeMux
	stageOfInterest        *models.Stage
	gongStage              *gong_models.Stage
	treeStage              *tree.Stage
	treeNavigationStage    *tree.Stage
	formStage              *form.Stage
	tableStage             *table.Stage
	notificationTableStage *table.Stage
	splitStage             *split.Stage
	loadStage              *load.Stage

	fileName               string

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

	// updateSliceOfPointersCallback is called after a SliceOfPointers field is updated in the probe
	updateSliceOfPointersCallback func(instance any, fieldName string, slicePtr any)
}

func (probe *Probe) UpdateSliceOfPointersCallback(instance any, fieldName string, slicePtr any) {
	if probe.updateSliceOfPointersCallback != nil {
		probe.updateSliceOfPointersCallback(instance, fieldName, slicePtr)
	}
}

func (probe *Probe) SetUpdateSliceOfPointersCallback(cb func(instance any, fieldName string, slicePtr any)) {
	probe.updateSliceOfPointersCallback = cb
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

func (probe *Probe) GetProbeLoadStageName() string {
	return probe.stageOfInterest.GetProbeLoadStageName()
}

func NewProbe(
	r *http.ServeMux,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	embeddedDiagrams bool,
	stageOfInterest *models.Stage) (probe *Probe) {

	// split stage for the whole probe
	splitStage, _ := split_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	stageOfInterest.MetaPackageImportPath = "github.com/fullstack-lang/gong/dsm/capture/go/models"

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

	loadStage, _ := load_fullstack.NewStackInstance(r, stageOfInterest.GetProbeLoadStageName())

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
		loadStage:                      loadStage,
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
		"github.com/fullstack-lang/gong/dsm/capture/go"+":"+stageOfInterest.GetName(),
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		probe.diagramEditor,
		stageOfInterest.Map_GongStructName_InstancesNb,
	)

	probe.dataEditor = gongprobe.CreateDataEditorLayout(
		probe.treeNavigationStage.GetName(),
		probe.treeStage.GetName(),
		probe.loadStage.GetName(),
		probe.tableStage.GetName(),
		probe.notificationTableStage.GetName(),
		probe.formStage.GetName(),
	)

	probe.splitStage.StageBranch(&split.View{
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

	probe.initLoadStage()

	probe.ux_tree()

	return
}

func (probe *Probe) initLoadStage() {
	probe.loadStage.Reset()
	probe.loadStage.Commit()
}

func (probe *Probe) Refresh() {
	probe.ux_tree()
	probe.ux_table()
	probe.ux_form()
	probe.docStager.Svg()
}

type Notification = gongprobe.Notification

const NbNotificationMax = gongprobe.NbNotificationMax

func (probe *Probe) AddNotification(date time.Time, message string) {
	gongprobe.AddNotification(&probe.notification, date, message)
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

func (probe *Probe) DownloadNotificationsCSV() {
	gongprobe.DownloadNotificationsCSV(probe.notification, probe.loadStage, probe.initLoadStage)
}

func (probe *Probe) ExportStageExcel() {
	excelBytes, err := probe.stageOfInterest.SerializeStageAsBytes(false)
	gongprobe.ExportStageExcel(
		excelBytes,
		err,
		probe.fileName,
		"capture",
		probe.stageOfInterest.GetName(),
		probe.loadStage,
		probe.initLoadStage,
		probe.AddNotification,
		probe.CommitNotificationTable,
	)
}
