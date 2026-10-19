// generated code - do not edit
package probe

import (
	"embed"
	"encoding/base64"
	"fmt"
	"regexp"
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

	"github.com/fullstack-lang/gong/app/reqif/go/models"

	embeddedgo "github.com/fullstack-lang/gong/app/reqif/go"
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

	stageOfInterest.MetaPackageImportPath = "github.com/fullstack-lang/gong/app/reqif/go/models"

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
		"github.com/fullstack-lang/gong/app/reqif/go"+":"+stageOfInterest.GetName(),
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

type stageLoadProxy struct {
	probe *Probe
}

func (proxy *stageLoadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	proxy.probe.fileName = uploadedFile.GetName()
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}
	proxy.probe.stageOfInterest.Reset()
	err = proxy.probe.stageOfInterest.ParseAstString(string(decodedBytes), true)
	if err != nil {
		return err
	}
	proxy.probe.stageOfInterest.Commit()
	proxy.probe.Refresh()
	return nil
}

func (probe *Probe) initLoadStage() {
	probe.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &stageLoadProxy{
			probe: probe,
		},
	}

	probe.loadStage.StageBranch(
		fileToUpload,
	)

	message := &load.Message{
		Name: "Drop your stage.go file here or ",
	}

	message.Stage(probe.loadStage)

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

func (probe *Probe) GetDocStager() *doc.Stager {
	return probe.docStager
}

func (probe *Probe) GetStage() *models.Stage {
	return probe.stageOfInterest
}

func (probe *Probe) GetSplitStage() *split.Stage {
	return probe.splitStage
}

func (probe *Probe) GetTreeStage() *tree.Stage {
	return probe.treeStage
}

func (probe *Probe) GetFormStage() *form.Stage {
	return probe.formStage
}

func (probe *Probe) GetTableStage() *table.Stage {
	return probe.tableStage
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

func (probe *Probe) ComputeInstancesNb() map[string]int {
	probe.stageOfInterest.ComputeInstancesNb()
	res := make(map[string]int)
	for k, v := range probe.stageOfInterest.Map_GongStructName_InstancesNb {
		res[k] = v
		res["models."+k] = v
	}
	return res
}

func (probe *Probe) FillUpFormFromGongstruct(instance any, formName string) {
	FillUpFormFromGongstruct(instance, probe)
}

func (probe *Probe) DownloadNotificationsCSV() {
	gongprobe.DownloadNotificationsCSV(probe.notification, probe.loadStage, probe.initLoadStage)
}

func (probe *Probe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	if probe.fileName == "" {
		probe.fileName = "reqif-stage.go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	stageString, err := probe.stageOfInterest.MarshallToString("main")
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))
	probe.loadStage.StageBranch(fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second)
	probe.initLoadStage()
}

func (probe *Probe) ExportStageExcel() {
	excelBytes, err := probe.stageOfInterest.SerializeStageAsBytes(false)
	gongprobe.ExportStageExcel(
		excelBytes,
		err,
		probe.fileName,
		"reqif",
		probe.stageOfInterest.GetName(),
		probe.loadStage,
		probe.initLoadStage,
		probe.AddNotification,
		probe.CommitNotificationTable,
	)
}
