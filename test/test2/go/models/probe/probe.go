// generated code - do not edit
package probe

import (
	"embed"
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"time"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
)

type Notification = gongprobe.Notification

const NbNotificationMax = gongprobe.NbNotificationMax

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

	diagramEditor *split.AsSplitArea
	docStager     *doc_models.Stager

	fileName string

	dataEditor *split.AsSplit

	notification []*Notification

	maxElementsNbPerGongStructNode int
	commitMode                     bool
	bulkDeleteMode                 bool

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
	stageOfInterest *models.Stage,
) (probe *Probe) {

	splitStage, _ := split_fullstack.NewStackInstance(r, stageOfInterest.GetProbeSplitStageName())
	splitStage.Commit()

	stageOfInterest.MetaPackageImportPath = "github.com/fullstack-lang/gong/test/test2/go/models"

	stage := gong_models.NewStage(stageOfInterest.GetName())
	gong_models.LoadEmbedded(stage, goModelsDir)

	treeStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTreeSidebarStageName())
	treeNavigationStage, _ := tree_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNavigationTreeSidebarStageName())

	tableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeTableStageName())
	tableStage.Commit()

	notificationTableStage, _ := table_fullstack.NewStackInstance(r, stageOfInterest.GetProbeNotificationTableStageName())
	notificationTableStage.Commit()

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

	probe.diagramEditor = &split.AsSplitArea{
		Name:             "Bottom",
		ShowNameInHeader: false,
		Size:             50,
	}

	metaPackageImports := []*doc_models.MetaPackageImport{
		{Alias: "ref_models", Path: "\"github.com/fullstack-lang/gong/test/test2/go/models\""},
		{Alias: "ref_x", Path: "\"github.com/fullstack-lang/gong/test/test2/go/models/x\""},
		{Alias: "ref_y", Path: "\"github.com/fullstack-lang/gong/test/test2/go/models/y\""},
		{Alias: "ref_model", Path: "\"github.com/fullstack-lang/gong/test/test2/go/models/x/models\""},
	}

	probe.docStager = prepare.PrepareStageSet(
		r,
		embeddedDiagrams,
		"github.com/fullstack-lang/gong/test/test2/go:"+stageOfInterest.GetName(),
		metaPackageImports,
		goModelsDir,
		goDiagramsDir,
		probe.diagramEditor,
		probe.ComputeInstancesNb(),
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

	return probe
}

func (probe *Probe) GetDocStager() *doc_models.Stager {
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
	if probe.stageOfInterest != nil {
		for k, v := range probe.stageOfInterest.Map_GongStructName_InstancesNb {
			res[k] = v
			res["models."+k] = v
		}
	}
	if probe.stageOfInterest.XStage != nil {
		for k, v := range probe.stageOfInterest.XStage.Map_GongStructName_InstancesNb {
			res[k] = v
			res["x."+k] = v
		}
	}
	if probe.stageOfInterest.YStage != nil {
		for k, v := range probe.stageOfInterest.YStage.Map_GongStructName_InstancesNb {
			res[k] = v
			res["y."+k] = v
		}
	}
	if probe.stageOfInterest.ModelStage != nil {
		for k, v := range probe.stageOfInterest.ModelStage.Map_GongStructName_InstancesNb {
			res[k] = v
			res["model."+k] = v
		}
	}
	return res
}

func (probe *Probe) Refresh() {
	if probe.docStager != nil {
		probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
		probe.docStager.Svg()
	}
	probe.ux_tree()
	probe.ux_table()
	probe.ux_form()
}

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
		"test2",
		probe.stageOfInterest.GetName(),
		probe.loadStage,
		probe.initLoadStage,
		probe.AddNotification,
		probe.CommitNotificationTable,
	)
}

func (probe *Probe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	if probe.fileName == "" {
		probe.fileName = "test2-stage.go"
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
