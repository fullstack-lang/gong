// generated code - do not edit
package data

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func NewOnSortingEditon[FieldType models.PointerToGongstruct](
	r *gin.Engine,
	formStage *form.StageStruct,
	field *[]FieldType,
	stageOfInterest *models.StageStruct,
) (onSortingEdition *OnSortingEditon[FieldType]) {

	onSortingEdition = new(OnSortingEditon[FieldType])

	onSortingEdition.r = r
	onSortingEdition.formStage = formStage
	onSortingEdition.field = field
	onSortingEdition.stageOfInterest = stageOfInterest

	return
}

type OnSortingEditon[FieldType models.PointerToGongstruct] struct {
	r               *gin.Engine
	formStage       *form.StageStruct
	field           *[]FieldType
	stageOfInterest *models.StageStruct
}

func (onSortingEditon *OnSortingEditon[FieldType]) OnButtonPressed() {

	tableStackName := onSortingEditon.formStage.GetPath() + string(form.StackNamePostFixForTableForAssociationSorting)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onSortingEditon.r, tableStackName)

	table := new(gongtable_models.Table).Stage(tableStageForSelection)
	table.Name = string(form.TableSortExtraName)
	table.HasColumnSorting = false
	table.HasFiltering = false
	table.HasPaginator = false
	table.HasCheckableRows = false
	table.HasSaveButton = true
	table.CanDragDropRows = true

	for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	map_RowID_instance := make(map[*gongtable_models.Row]FieldType)
	for _, instance := range *onSortingEditon.field {
		row := new(gongtable_models.Row).Stage(tableStageForSelection)
		row.Name = instance.GetName()
		map_RowID_instance[row] = instance
		table.Rows = append(table.Rows, row)

		for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
			cell := new(gongtable_models.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %s - Column %s", instance.GetName(), fieldName)

			cellString := new(gongtable_models.CellString).Stage(tableStageForSelection)
			cellString.Name = models.GetFieldStringValueFromPointer(instance, fieldName)
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	table.Impl = NewTableSortSaver[FieldType](onSortingEditon.field, onSortingEditon.stageOfInterest, &map_RowID_instance)
	tableStageForSelection.Commit()
}

func NewTableSortSaver[FieldType models.PointerToGongstruct](
	field *[]FieldType,
	stageOfInterest *models.StageStruct,
	map_RowID_instance *map[*gongtable_models.Row]FieldType,
) (tableSortSaver *TableSortSaver[FieldType]) {

	tableSortSaver = new(TableSortSaver[FieldType])
	tableSortSaver.field = field
	tableSortSaver.stageOfInterest = stageOfInterest
	tableSortSaver.map_RowID_instance = map_RowID_instance

	return
}

type TableSortSaver[FieldType models.PointerToGongstruct] struct {
	field           *[]FieldType
	stageOfInterest *models.StageStruct

	// map giving the relation between the row ID and the instance
	map_RowID_instance *map[*gongtable_models.Row]FieldType
}

func (tableSortSaver *TableSortSaver[FieldType]) TableUpdated(stage *form.StageStruct, table, updatedTable *form.Table) {
	log.Println("TableSortSaver: TableUpdated")

	// checkout to the stage to get the rows that have been checked and not
	stage.Checkout()

	*tableSortSaver.field = make([]FieldType, 0)

	for _, row := range table.Rows {
		instance := (*tableSortSaver.map_RowID_instance)[row]
		*tableSortSaver.field = append(*tableSortSaver.field, instance)
	}
	tableSortSaver.stageOfInterest.Commit()
}
