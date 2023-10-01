// generated code - do not edit
package x

import (
	"fmt"
	"log"

	"github.com/fullstack-lang/gong/test2/go/models/x"
	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"
)

func NewOnSortingEditon[InstanceType x.PointerToGongstruct, FieldType x.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	playground *Playground,
) (onSortingEdition *OnSortingEditon[InstanceType, FieldType]) {

	onSortingEdition = new(OnSortingEditon[InstanceType, FieldType])

	onSortingEdition.playground = playground
	onSortingEdition.instance = instance
	onSortingEdition.field = field

	return
}

type OnSortingEditon[InstanceType x.PointerToGongstruct, FieldType x.PointerToGongstruct] struct {
	instance   InstanceType
	field      *[]FieldType
	playground *Playground
}

func (onSortingEditon *OnSortingEditon[InstanceType, FieldType]) OnButtonPressed() {

	tableStackName := onSortingEditon.playground.formStage.GetPath() +
		string(form.StackNamePostFixForTableForAssociationSorting)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onSortingEditon.playground.r, tableStackName)

	table := new(gongtable_models.Table).Stage(tableStageForSelection)
	table.Name = string(form.TableSortExtraName)
	table.HasColumnSorting = false
	table.HasFiltering = false
	table.HasPaginator = false
	table.HasCheckableRows = false
	table.HasSaveButton = true
	table.CanDragDropRows = true

	for _, fieldName := range x.GetFieldsFromPointer[FieldType]() {
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

		for _, fieldName := range x.GetFieldsFromPointer[FieldType]() {
			cell := new(gongtable_models.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %s - Column %s", instance.GetName(), fieldName)

			cellString := new(gongtable_models.CellString).Stage(tableStageForSelection)
			cellString.Name = x.GetFieldStringValueFromPointer(instance, fieldName)
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	table.Impl = NewTableSortSaver[InstanceType, FieldType](
		onSortingEditon.instance,
		onSortingEditon.field,
		onSortingEditon.playground,
		&map_RowID_instance)
	tableStageForSelection.Commit()
}

func NewTableSortSaver[InstanceType x.PointerToGongstruct, FieldType x.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	playground *Playground,
	map_RowID_instance *map[*gongtable_models.Row]FieldType,
) (tableSortSaver *TableSortSaver[InstanceType, FieldType]) {

	tableSortSaver = new(TableSortSaver[InstanceType, FieldType])
	tableSortSaver.instance = instance
	tableSortSaver.field = field
	tableSortSaver.playground = playground
	tableSortSaver.map_RowID_instance = map_RowID_instance

	return
}

type TableSortSaver[InstanceType x.PointerToGongstruct, FieldType x.PointerToGongstruct] struct {
	instance   InstanceType
	field      *[]FieldType
	playground *Playground

	// map giving the relation between the row ID and the instance
	map_RowID_instance *map[*gongtable_models.Row]FieldType
}

func (tableSortSaver *TableSortSaver[InstanceType, FieldType]) TableUpdated(stage *form.StageStruct, table, updatedTable *form.Table) {
	log.Println("TableSortSaver: TableUpdated")

	// checkout to the stage to get the rows that have been checked and not
	stage.Checkout()

	*tableSortSaver.field = make([]FieldType, 0)

	for _, row := range table.Rows {
		instance := (*tableSortSaver.map_RowID_instance)[row]
		*tableSortSaver.field = append(*tableSortSaver.field, instance)
	}
	tableSortSaver.playground.stageOfInterest.Commit()

	// see the result
	fillUpTablePointerToGongstruct[InstanceType](
		tableSortSaver.playground,
	)
	tableSortSaver.playground.tableStage.Commit()
}
