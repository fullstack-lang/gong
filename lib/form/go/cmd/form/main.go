package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	form_stack "github.com/fullstack-lang/gong/lib/form/go/stack"
	form_static "github.com/fullstack-lang/gong/lib/form/go/static"

	"github.com/gin-gonic/gin"

	// form_models "github.com/fullstack-lang/gong/lib/form/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	table_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

type StackName string

// values for TableTestNameEnum
const (
	ManualyEditedFormStackName StackName = "manualy edited form"
)

func main() {

	log.SetPrefix("form: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := form_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := form_stack.NewStack(r, "form", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	stageForManualyEditedForm := stack.Stage

	{

		// hook table stack creation to on pressed button
		for formEditAssocButton := range *form.GetGongstructInstancesSet[form.FormEditAssocButton](
			stageForManualyEditedForm) {
			_ = formEditAssocButton

			// create the struct for callabck
			onAssocEditon := new(OnAssocEditon)
			onAssocEditon.r = r
			onAssocEditon.sourceStackName = stageForManualyEditedForm.GetName()
			formEditAssocButton.OnAssocEditon = onAssocEditon
		}
		for formSortAssocButton := range *form.GetGongstructInstancesSet[form.FormSortAssocButton](
			stageForManualyEditedForm) {
			_ = formSortAssocButton

			// create the struct for callabck
			onSortEdition := new(OnSortEditon)
			onSortEdition.r = r
			onSortEdition.sourceStackNName = stageForManualyEditedForm.GetName()
			formSortAssocButton.OnSortEdition = onSortEdition
		}

	}

	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// form_models.NewStager(r, stack.Stage, stack.Probe)
	split.StageBranch(splitStage, &split.View{
		Name: "Split",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "Form",
							Size: 50,
							Form2: (&split.Form2{
								Name:      "Form2",
								StackName: stack.Stage.GetName(),
							}),
						}),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type OnAssocEditon struct {
	r               *gin.Engine
	sourceStackName string
}

// OnSave implements models.FormEditAssocButtonInterface.
func (onAssocEditon *OnAssocEditon) OnSave() {
	panic("unimplemented")
}

func (onAssocEditon *OnAssocEditon) OnButtonPressed() {
	// create a new with the name append "-table"
	newStage, _ := table_fullstack.NewStackInstance(
		onAssocEditon.r,
		onAssocEditon.sourceStackName+
			string(table.StackNamePostFixForTableForAssociation))

	fillUpSelectTableDummyStuff(newStage, string(table.TableSelectExtraName))
	newStage.Commit()
}

type OnSortEditon struct {
	r                *gin.Engine
	sourceStackNName string
}

func (onSortEditon *OnSortEditon) OnButtonPressed() {

	newStage, _ := table_fullstack.NewStackInstance(
		onSortEditon.r,
		onSortEditon.sourceStackNName+
			string(table.StackNamePostFixForTableForAssociationSorting))

	fillUpSortTableDummyStuff(newStage, string(table.TableSortExtraName))
	newStage.Commit()
}

func fillUpSortTableDummyStuff(stage *table_models.Stage, tableName string) {
	nbRows := 10
	nbColumns := 1
	table := new(table_models.Table).Stage(stage)
	table.Name = tableName
	table.CanDragDropRows = true
	table.HasCloseButton = true

	for j := 0; j < nbColumns; j++ {
		column := new(table_models.DisplayedColumn).Stage(stage)
		column.Name = fmt.Sprintf("Column %d", j)

		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(table_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		for j := 0; j < nbColumns; j++ {
			cell := new(table_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(table_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}

func fillUpSelectTableDummyStuff(stage *table_models.Stage, tableName string) {
	nbRows := 10
	nbColumns := 1
	table := new(table_models.Table).Stage(stage)
	table.Name = tableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = false
	table.HasCheckableRows = true
	table.HasSaveButton = true

	for j := 0; j < nbColumns; j++ {
		column := new(table_models.DisplayedColumn).Stage(stage)
		column.Name = fmt.Sprintf("Column %d", j)
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(table_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		if i%2 == 0 {
			row.IsChecked = true
		}

		for j := 0; j < nbColumns; j++ {
			cell := new(table_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(table_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}
