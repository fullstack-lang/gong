// generated code - do not edit
package probe

import (
	"fmt"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	gongtable_models "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	table.SaveButtonLabel = "Close form"
	table.CanDragDropRows = true

	for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](onSortingEditon.probe.stageOfInterest)
	instanceSlice := make([]FieldType, 0)
	for instance := range instanceSet {
		instanceSlice = append(instanceSlice, instance)
	}
	map_RowID_instance := make(map[*gongtable_models.Row]FieldType)
	for instance := range instanceSet {
		row := new(gongtable_models.Row).Stage(tableStageForSelection)
		row.Name = instance.GetName()
		map_RowID_instance[row] = instance
		table.Rows = append(table.Rows, row)

		cell := (&gongtable_models.Cell{
			Name: "ID",
		}).Stage(tableStageForSelection)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable_models.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				onSortingEditon.probe.stageOfInterest,
				instance,
			)),
		}).Stage(tableStageForSelection)
		cell.CellInt = cellInt

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

	tableStageForSelection.Commit()
}
