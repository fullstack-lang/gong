// generated code - do not edit
package probe

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"regexp"
	"time"

	form_fullstack "github.com/fullstack-lang/gong/lib/form/go/fullstack"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	split_fullstack "github.com/fullstack-lang/gong/lib/split/go/fullstack"
	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
)

type StageSetProbe struct {
	r                      *http.ServeMux
	stageSet               *models.StageSet
	treeStage              *tree.Stage
	treeNavigationStage    *tree.Stage
	formStage              *form.Stage
	tableStage             *table.Stage
	notificationTableStage *table.Stage
	splitStage             *split.Stage
	loadStage              *load.Stage

	fileName string

	dataEditor *split.AsSplit

	notification []*Notification

	maxElementsNbPerGongStructNode int
	commitMode                     bool
	bulkDeleteMode                 bool
}

func NewStageSetProbe(
	r *http.ServeMux,
	stageSet *models.StageSet,
) (probe *StageSetProbe) {

	splitStage, _ := split_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName())
	splitStage.Commit()

	treeStage, _ := tree_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-tree")
	treeNavigationStage, _ := tree_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-tree-nav")

	tableStage, _ := table_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-table")
	tableStage.Commit()

	notificationTableStage, _ := table_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-notif")
	notificationTableStage.Commit()

	formStage, _ := form_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-form")
	formStage.Commit()

	loadStage, _ := load_fullstack.NewStackInstance(r, stageSet.GetProbeSplitStageName()+"-load")

	probe = &StageSetProbe{
		r:                              r,
		stageSet:                       stageSet,
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

	probe.dataEditor = &split.AsSplit{
		Name:          "StageSet Top, sidebar, table & form",
		Direction:     split.Horizontal,
		IsSizeInPixel: true,
		AsSplitAreas: []*split.AsSplitArea{
			{
				Name: "sidebar",
				Size: 525,
				AsSplit: &split.AsSplit{
					Direction:              split.Vertical,
					IsSizeInPixel:          true,
					IsWithCustomGutterSize: true,
					GutterSize:             1,
					AsSplitAreas: []*split.AsSplitArea{
						{
							Name: "sidebar tree",
							Size: 53,
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
							Size: 70,
							Load: &split.Load{
								Name:      "Table",
								StackName: probe.loadStage.GetName(),
							},
						},
					},
				},
			},
			{
				Name:  "both tables",
				IsAny: true,
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
				Size: 525,
				Form: &split.Form{
					Name:      "Form",
					StackName: probe.formStage.GetName(),
				},
			},
		},
	}

	split.StageBranch(probe.splitStage, &split.View{
		Name: "StageSet Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Name:    "Top",
				Size:    100,
				AsSplit: probe.dataEditor,
			},
		},
	})
	probe.splitStage.Commit()

	probe.initLoadStage()
	probe.ux_tree()

	return probe
}

func (probe *StageSetProbe) GetStageSet() *models.StageSet {
	return probe.stageSet
}

func (probe *StageSetProbe) GetSplitStage() *split.Stage {
	return probe.splitStage
}

func (probe *StageSetProbe) GetTreeStage() *tree.Stage {
	return probe.treeStage
}

func (probe *StageSetProbe) GetFormStage() *form.Stage {
	return probe.formStage
}

func (probe *StageSetProbe) GetTableStage() *table.Stage {
	return probe.tableStage
}

func (probe *StageSetProbe) SetMaxElementsNbPerGongStructNode(nb int) {
	probe.maxElementsNbPerGongStructNode = nb
}

func (probe *StageSetProbe) GetMaxElementsNbPerGongStructNode() int {
	return probe.maxElementsNbPerGongStructNode
}

func (probe *StageSetProbe) Refresh() {
	probe.ux_tree()
	probe.ux_table()
}

func (probe *StageSetProbe) AddNotification(date time.Time, message string) {
	notification := &Notification{
		Date:    date,
		Message: message,
	}
	probe.notification = append(probe.notification, notification)
}

func (probe *StageSetProbe) CommitNotificationTable() {
	probe.notificationTableStage.Reset()
	tableInstance := new(table.Table)
	tableInstance.Name = "Notifications"
	tableInstance.HasFiltering = true

	column := new(table.DisplayedColumn)
	column.Name = "Date"
	tableInstance.DisplayedColumns = append(tableInstance.DisplayedColumns, column)

	column = new(table.DisplayedColumn)
	column.Name = "Message"
	tableInstance.DisplayedColumns = append(tableInstance.DisplayedColumns, column)

	for _, notification := range probe.notification {
		row := new(table.Row)
		row.Name = notification.Date.Format(time.StampMicro)

		cell := &table.Cell{Name: "Date"}
		cell.CellString = &table.CellString{Value: notification.Date.Format(time.StampMicro)}
		row.Cells = append(row.Cells, cell)

		cell = &table.Cell{Name: "Message"}
		cell.CellString = &table.CellString{Value: notification.Message}
		row.Cells = append(row.Cells, cell)

		tableInstance.Rows = append(tableInstance.Rows, row)
	}

	table.StageBranch(probe.notificationTableStage, tableInstance)
	probe.notificationTableStage.Commit()
}

func (probe *StageSetProbe) ExportStage() {
	probe.loadStage.Reset()

	fileToDownload := new(load.FileToDownload)
	if probe.fileName == "" {
		probe.fileName = "test2-stageset.go"
	}

	prefixRegex := regexp.MustCompile("^\\d{8} \\d{4} ")
	cleanFileName := prefixRegex.ReplaceAllString(probe.fileName, "")

	fileToDownload.Name = time.Now().Format("20060102 1504 ") + cleanFileName

	stageString, err := probe.stageSet.MarshallToString("main")
	if err != nil {
		probe.AddNotification(time.Now(), "Error serializing stage: "+err.Error())
		probe.CommitNotificationTable()
		return
	}

	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString([]byte(stageString))
	load.StageBranch(probe.loadStage, fileToDownload)
	probe.loadStage.Commit()

	time.Sleep(1 * time.Second)
	probe.initLoadStage()
}

type stageSetLoadProxy struct {
	probe *StageSetProbe
}

func (proxy *stageSetLoadProxy) OnFileUpload(uploadedFile *load.FileToUpload) error {
	proxy.probe.fileName = uploadedFile.GetName()
	decodedBytes, err := base64.StdEncoding.DecodeString(uploadedFile.Base64EncodedContent)
	if err != nil {
		return fmt.Errorf("base64.StdEncoding.DecodeString failed: %w", err)
	}
	proxy.probe.stageSet.Reset()
	err = proxy.probe.stageSet.ParseAstString(string(decodedBytes), true)
	if err != nil {
		return err
	}
	proxy.probe.stageSet.Commit()
	proxy.probe.Refresh()
	return nil
}

func (probe *StageSetProbe) initLoadStage() {
	probe.loadStage.Reset()

	fileToUpload := &load.FileToUpload{
		Name: "Name of file",
		FileToUploadProxy: &stageSetLoadProxy{
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
