// generated code - do not edit
package probe

import (
	"embed"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go/parser"
	"go/token"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"

	gong_models "github.com/fullstack-lang/gong/go/models"

	doc "github.com/fullstack-lang/gong/lib/doc/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"

	embeddedgo "github.com/fullstack-lang/gong/lib/ssg/go"
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

	loadStage, _ := load_fullstack.NewStackInstance(r, stageOfInterest.GetName()+"-probe")

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
		"github.com/fullstack-lang/gong/lib/ssg/go"+":"+stageOfInterest.GetName(),
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
						{
							Name: "load",
							Size: 83,
							Load: &split.Load{
								Name:      "Table",
								StackName: probe.loadStage.GetName(),
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

	probe.initLoadStage()

	probe.ux_tree()

	return
}

type loadProxy struct {
	probe *Probe
}

func (proxy *loadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	fmt.Println("OnFileUpload: start")
	proxy.probe.fileName = uploadedFile.GetName()

	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}

	// if the user loads a second file, we don't want the previous file to be committed
	proxy.probe.stageOfInterest.OnInitCommitCallback = nil

	proxy.probe.stageOfInterest.Reset()
	fmt.Println("OnFileUpload: after reset")

	fset := token.NewFileSet()
	inFile, errParser := parser.ParseFile(fset, "", decodedBytes, parser.ParseComments)
	if errParser != nil {
		return fmt.Errorf("Unable to parse: %w", errParser)
	}
	errParse := models.ParseAstFileFromAst(proxy.probe.stageOfInterest, inFile, fset, false)
	if errParse != nil {
		return errParse
	}

	fmt.Println("OnFileUpload: after parse")
	proxy.probe.stageOfInterest.Commit()
	fmt.Println("OnFileUpload: after commit")

	return nil
}

func (probe *Probe) initLoadStage() {
	probe.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &loadProxy{
			probe: probe,
		},
	}

	load.StageBranch(probe.loadStage,
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

func (probe *Probe) DownloadNotificationsCSV() {
	var csvContent string
	csvContent += "Date,Message\n"
	for _, notification := range probe.notification {
		// Escape quotes in message
		escapedMessage := strings.ReplaceAll(notification.Message, "\"", "\"\"")
		csvContent += fmt.Sprintf("\"%s\",\"%s\"\n", notification.Date.Format(time.StampMicro), escapedMessage)
	}

	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	fileToDownload.Name = "notifications.csv"
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(csvContent))

	load.StageBranch(probe.loadStage, fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	probe.initLoadStage()
}

func (probe *Probe) ExportStageExcel() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)

	if probe.fileName == "" {
		probe.fileName = probe.stageOfInterest.GetName() + ".go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")
	cleanFileName = strings.TrimSuffix(cleanFileName, ".go") + ".xlsx"

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	excelBytes, err := models.SerializeStageAsBytes(probe.stageOfInterest, false)
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(excelBytes)

	load.StageBranch(probe.loadStage, fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	probe.initLoadStage()
}

func (probe *Probe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)

	if probe.fileName == "" {
		probe.fileName = probe.stageOfInterest.GetName() + ".go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	stageString, err := probe.stageOfInterest.MarshallToString(probe.stageOfInterest.MetaPackageImportPath, "main")
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))

	load.StageBranch(probe.loadStage, fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second) // Sleep to ensure the client has time to start the download before we reset the stage.
	probe.initLoadStage()
}
