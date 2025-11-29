// generated code - do not edit
package probe

import (
	"fmt"
	"sort"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
		string(table.StackNamePostFixForTableForAssociationSorting)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onSortingEditon.probe.r, tableStackName)

	selectedInstancesTable := new(table.Table).Stage(tableStageForSelection)
	selectedInstancesTable.Name = string(table.TableSortExtraName)
	selectedInstancesTable.HasColumnSorting = false
	selectedInstancesTable.HasFiltering = false
	selectedInstancesTable.HasPaginator = false
	selectedInstancesTable.HasCheckableRows = false
	selectedInstancesTable.HasSaveButton = true
	selectedInstancesTable.SaveButtonLabel = "Close form"
	selectedInstancesTable.CanDragDropRows = true

	// add a column for the ID
	selectedInstancesTable.DisplayedColumns = append(selectedInstancesTable.DisplayedColumns, (&table.DisplayedColumn{
		Name: "ID",
	}).Stage(tableStageForSelection))

	for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
		column := new(table.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fieldName.Name
		selectedInstancesTable.DisplayedColumns = append(selectedInstancesTable.DisplayedColumns, column)
	}

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](onSortingEditon.probe.stageOfInterest)
	instances := make([]FieldType, 0)
	for instance := range instanceSet {
		instances = append(instances, instance)
	}
	sort.Slice(instances[:], func(i, j int) bool {
		instancei := instances[i]
		instancej := instances[j]
		instancei_order := models.GetOrderPointerGongstruct(onSortingEditon.probe.stageOfInterest, instancei)
		instnacej_order := models.GetOrderPointerGongstruct(onSortingEditon.probe.stageOfInterest, instancej)
		return instancei_order < instnacej_order
	})
	map_RowID_instance := make(map[*table.Row]FieldType)
	for _, instance := range instances {
		row := new(table.Row).Stage(tableStageForSelection)
		row.Name = instance.GetName()
		map_RowID_instance[row] = instance
		selectedInstancesTable.Rows = append(selectedInstancesTable.Rows, row)

		cell := (&table.Cell{
			Name: "ID",
		}).Stage(tableStageForSelection)
		row.Cells = append(row.Cells, cell)
		cellInt := (&table.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				onSortingEditon.probe.stageOfInterest,
				instance,
			)),
		}).Stage(tableStageForSelection)
		cell.CellInt = cellInt

		for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
			cell := new(table.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %s - Column %s", instance.GetName(), fieldName)

			cellString := new(table.CellString).Stage(tableStageForSelection)
			value := models.GetFieldStringValueFromPointer(instance, fieldName.Name, onSortingEditon.probe.stageOfInterest)
			cellString.Name = value.GetValueString()
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	tableStageForSelection.Commit()
}
