// generated code - do not edit
package probe

import (
	"fmt"
	"log"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	gongtable_models "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

func NewOnSortingEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	probe *Probe,
) (onSortingEdition *OnSortingEditon[InstanceType, FieldType]) {

	onSortingEdition = new(OnSortingEditon[InstanceType, FieldType])

	onSortingEdition.probe = probe
	onSortingEdition.instance = instance
	onSortingEdition.field = field

	return
}

type OnSortingEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	instance InstanceType
	field    *[]FieldType
	probe    *Probe
}

func (onSortingEditon *OnSortingEditon[InstanceType, FieldType]) OnButtonPressed() {

	tableStackName := onSortingEditon.probe.formStage.GetName() +
		string(form.StackNamePostFixForTableForAssociationSorting)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onSortingEditon.probe.r, tableStackName)

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
			value := models.GetFieldStringValueFromPointer(instance, fieldName)
			cellString.Name = value.GetValueString()
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	table.Impl = NewTableSortSaver[InstanceType, FieldType](
		onSortingEditon.instance,
		onSortingEditon.field,
		onSortingEditon.probe,
		&map_RowID_instance)
	tableStageForSelection.Commit()
}

func NewTableSortSaver[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	probe *Probe,
	map_RowID_instance *map[*gongtable_models.Row]FieldType,
) (tableSortSaver *TableSortSaver[InstanceType, FieldType]) {

	tableSortSaver = new(TableSortSaver[InstanceType, FieldType])
	tableSortSaver.instance = instance
	tableSortSaver.field = field
	tableSortSaver.probe = probe
	tableSortSaver.map_RowID_instance = map_RowID_instance

	return
}

type TableSortSaver[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	instance   InstanceType
	field      *[]FieldType
	probe *Probe

	// map giving the relation between the row ID and the instance
	map_RowID_instance *map[*gongtable_models.Row]FieldType
}

func (tableSortSaver *TableSortSaver[InstanceType, FieldType]) TableUpdated(stage *form.Stage, table, updatedTable *form.Table) {
	log.Println("TableSortSaver: TableUpdated")

	// checkout to the stage to get the rows that have been checked and not
	stage.Checkout()

	*tableSortSaver.field = make([]FieldType, 0)

	for _, row := range table.Rows {
		instance := (*tableSortSaver.map_RowID_instance)[row]
		*tableSortSaver.field = append(*tableSortSaver.field, instance)
	}
	tableSortSaver.probe.stageOfInterest.Commit()

	// see the result
	fillUpTablePointerToGongstruct[InstanceType](
		tableSortSaver.probe,
	)
	tableSortSaver.probe.tableStage.Commit()
}
