package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	// gongtable_data "github.com/fullstack-lang/gong/lib/table/go/data"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	"github.com/fullstack-lang/gong/lib/table/go/models"

	gongtable_models "github.com/fullstack-lang/gong/lib/table/go/models"
	gongtable_orm "github.com/fullstack-lang/gong/lib/table/go/orm"
	gongtable_static "github.com/fullstack-lang/gong/lib/table/go/static"

	gongtable_probe "github.com/fullstack-lang/gong/lib/table/go/probe"

	gongtable_go "github.com/fullstack-lang/gong/lib/table/go"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplForManualyEditedTableStage struct {
}

func (impl *BeforeCommitImplForManualyEditedTableStage) BeforeCommit(stage *gongtable_models.Stage) {
	file, err := os.Create("./table_stage.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/lib/table/go/models", "main")
}

type BeforeCommitImplForManualyEditedFormStage struct {
}

func (impl *BeforeCommitImplForManualyEditedFormStage) BeforeCommit(stage *gongtable_models.Stage) {
	file, err := os.Create("./form_stage_issue30.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/lib/table/go/models", "main")
}

func main() {

	log.SetPrefix("gongtable: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongtable_static.ServeStaticFiles(*logGINFlag)

	// stageForManualyEditedTable is the stage from the code
	var stageForManualyEditedTable *gongtable_models.Stage
	var backRepoForManualyEditedTable *gongtable_orm.BackRepoStruct
	_ = backRepoForManualyEditedTable

	var stageForManualyEditedForm *gongtable_models.Stage
	var backRepoForManualyEditedForm *gongtable_orm.BackRepoStruct
	_ = backRepoForManualyEditedForm

	// persistence in a SQLite file on disk in memory
	stageForManualyEditedTable, backRepoForManualyEditedTable =
		gongtable_fullstack.NewStackInstance(r, string(models.ManualyEditedTableStackName))
	stageForManualyEditedForm, backRepoForManualyEditedForm =
		gongtable_fullstack.NewStackInstance(r, string(models.ManualyEditedFormStackName))

	// generated stage is the stage generated by filling data with code
	// to test performances
	stageForGeneratedTable, _ :=
		gongtable_fullstack.NewStackInstance(r, models.GeneratedTableStackName.ToString())

	fillUpTableWithDummyStuff(stageForGeneratedTable, gongtable_models.TableDefaultName.ToString())

	stageForGeneratedTable.Commit()

	{
		stageForManualyEditedTable.Checkout()
		stageForManualyEditedTable.Reset()
		stageForManualyEditedTable.Commit()
		err := gongtable_models.ParseAstFile(stageForManualyEditedTable, "table_stage.go")

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stageForManualyEditedTable.Commit()
	}
	{
		stageForManualyEditedForm.Checkout()
		stageForManualyEditedForm.Reset()
		stageForManualyEditedForm.Commit()
		err := gongtable_models.ParseAstFile(stageForManualyEditedForm, "form_stage.go")

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stageForManualyEditedForm.Commit()

		// hook table stack creation to on pressed button
		for formEditAssocButton := range *gongtable_models.GetGongstructInstancesSet[gongtable_models.FormEditAssocButton](
			stageForManualyEditedForm) {
			_ = formEditAssocButton

			// create the struct for callabck
			onAssocEditon := new(OnAssocEditon)
			onAssocEditon.r = r
			onAssocEditon.sourceStack = stageForManualyEditedForm
			formEditAssocButton.OnAssocEditon = onAssocEditon
		}
		for formSortAssocButton := range *gongtable_models.GetGongstructInstancesSet[gongtable_models.FormSortAssocButton](
			stageForManualyEditedForm) {
			_ = formSortAssocButton

			// create the struct for callabck
			onSortEdition := new(OnSortEditon)
			onSortEdition.r = r
			onSortEdition.sourceStack = stageForManualyEditedForm
			formSortAssocButton.OnSortEdition = onSortEdition
		}

	}

	// hook automatic marshall to go code at every commit
	{
		hook := new(BeforeCommitImplForManualyEditedTableStage)
		stageForManualyEditedTable.OnInitCommitFromFrontCallback = hook
	}
	{
		hook := new(BeforeCommitImplForManualyEditedFormStage)
		stageForManualyEditedForm.OnInitCommitFromFrontCallback = hook // callback when the front is used
		stageForManualyEditedForm.OnInitCommitFromBackCallback = hook  // callback when the probe is used
	}

	gongtable_probe.NewProbe(r,
		gongtable_go.GoModelsDir,
		gongtable_go.GoDiagramsDir,
		*embeddedDiagrams,
		stageForManualyEditedTable)

	gongtable_probe.NewProbe(r,
		gongtable_go.GoModelsDir,
		gongtable_go.GoDiagramsDir,
		*embeddedDiagrams,
		stageForManualyEditedForm)

	gongtable_probe.NewProbe(r,
		gongtable_go.GoModelsDir,
		gongtable_go.GoDiagramsDir,
		*embeddedDiagrams,
		stageForGeneratedTable)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stageForManualyEditedTable.GetName(),
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "table",
							Size: 50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: stageForManualyEditedTable.GetName(),
								TableName: "Table with 5 types",
							}),
						}),
						(&split.AsSplitArea{
							Name: "table",
							Size: 50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: stageForManualyEditedTable.GetName(),
								TableName: "EmptyTable",
							}),
						}),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stageForManualyEditedTable.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stageForManualyEditedForm.GetName(),
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "Form",
							Size: 50,
							Form: (&split.Form{
								Name:      "Form",
								StackName: stageForManualyEditedForm.GetName(),
								FormName:  "Form 1",
							}),
						}),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stageForManualyEditedForm.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	// one for the probe of the
	split.StageBranch(splitStage, &split.View{
		Name: stageForGeneratedTable.GetName(),
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "table",
							Size: 50,
							Table: (&split.Table{
								Name:      "Table",
								StackName: stageForGeneratedTable.GetName(),
								TableName: "Table",
							}),
						}),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stageForGeneratedTable.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	splitStage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type OnAssocEditon struct {
	r           *gin.Engine
	sourceStack *gongtable_models.Stage
}

// OnSave implements models.FormEditAssocButtonInterface.
func (onAssocEditon *OnAssocEditon) OnSave() {
	panic("unimplemented")
}

func (onAssocEditon *OnAssocEditon) OnButtonPressed() {
	// create a new with the name append "-table"
	newStage, _ := gongtable_fullstack.NewStackInstance(
		onAssocEditon.r,
		onAssocEditon.sourceStack.GetName()+
			string(gongtable_models.StackNamePostFixForTableForAssociation))

	fillUpSelectTableDummyStuff(newStage, string(gongtable_models.TableSelectExtraName))
	newStage.Commit()
}

type OnSortEditon struct {
	r           *gin.Engine
	sourceStack *gongtable_models.Stage
}

func (onSortEditon *OnSortEditon) OnButtonPressed() {

	newStage, _ := gongtable_fullstack.NewStackInstance(
		onSortEditon.r,
		onSortEditon.sourceStack.GetName()+
			string(gongtable_models.StackNamePostFixForTableForAssociationSorting))

	fillUpSortTableDummyStuff(newStage, string(gongtable_models.TableSortExtraName))
	newStage.Commit()
}
func fillUpSortTableDummyStuff(stage *gongtable_models.Stage, tableName string) {
	nbRows := 10
	nbColumns := 1
	table := new(gongtable_models.Table).Stage(stage)
	table.Name = tableName
	table.CanDragDropRows = true
	table.HasCloseButton = true

	for j := 0; j < nbColumns; j++ {
		column := new(gongtable_models.DisplayedColumn).Stage(stage)
		column.Name = fmt.Sprintf("Column %d", j)

		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(gongtable_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		for j := 0; j < nbColumns; j++ {
			cell := new(gongtable_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(gongtable_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}

func fillUpSelectTableDummyStuff(stage *gongtable_models.Stage, tableName string) {
	nbRows := 10
	nbColumns := 1
	table := new(gongtable_models.Table).Stage(stage)
	table.Name = tableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = false
	table.HasCheckableRows = true
	table.HasSaveButton = true

	for j := 0; j < nbColumns; j++ {
		column := new(gongtable_models.DisplayedColumn).Stage(stage)
		column.Name = fmt.Sprintf("Column %d", j)
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(gongtable_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		if i%2 == 0 {
			row.IsChecked = true
		}

		for j := 0; j < nbColumns; j++ {
			cell := new(gongtable_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(gongtable_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}

func fillUpTableWithDummyStuff(stage *gongtable_models.Stage, tableName string) {
	nbRows := 200
	nbColumns := 10
	table := new(gongtable_models.Table).Stage(stage)
	table.Name = tableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	table.NbOfStickyColumns = 2

	for j := 0; j < nbColumns; j++ {
		column := new(gongtable_models.DisplayedColumn).Stage(stage)
		column.Name = fmt.Sprintf("Column %d", j)
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(gongtable_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		for j := 0; j < nbColumns; j++ {
			cell := new(gongtable_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(gongtable_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}
