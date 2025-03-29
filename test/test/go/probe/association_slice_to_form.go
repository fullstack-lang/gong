// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	gongtable_models "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/test/test/go/models"
)

// AssociationSliceToForm add a form div with 2 buttons
// - one for selection
// - one for sorting
func AssociationSliceToForm[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	fieldName string,
	instance InstanceType,
	field *[]FieldType,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	formEditAssocButton := (&form.FormEditAssocButton{
		Name:  fieldName,
		Label: fieldName,
	}).Stage(probe.formStage)
	formDiv.FormEditAssocButton = formEditAssocButton
	onAssocEditon := NewOnAssocEditon(instance, field, fieldName, probe)
	formEditAssocButton.OnAssocEditon = onAssocEditon

	formSortAssocButton := (&form.FormSortAssocButton{
		Name:  fieldName,
		Label: fieldName,
	}).Stage(probe.formStage)
	formDiv.FormSortAssocButton = formSortAssocButton
	onSortingEditon := NewOnSortingEditon(instance, field, probe)
	formSortAssocButton.OnSortEdition = onSortingEditon

}

type OnAssocEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	instance  InstanceType
	field     *[]FieldType
	fieldName string
	probe     *Probe
}

func NewOnAssocEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	fieldName string,
	probe *Probe,
) (onAssocEdition *OnAssocEditon[InstanceType, FieldType]) {

	onAssocEdition = new(OnAssocEditon[InstanceType, FieldType])
	onAssocEdition.instance = instance
	onAssocEdition.field = field
	onAssocEdition.fieldName = fieldName
	onAssocEdition.probe = probe

	return
}

func (onAssocEditon *OnAssocEditon[InstanceType, FieldType]) OnButtonPressed() {

	tableStackName := onAssocEditon.probe.formStage.GetName() + string(form.StackNamePostFixForTableForAssociation)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEditon.probe.r, tableStackName)

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](onAssocEditon.probe.stageOfInterest)
	instanceSlice := make([]FieldType, 0)
	for instance := range instanceSet {
		instanceSlice = append(instanceSlice, instance)
	}
	sort.Slice(instanceSlice, func(i, j int) bool {
		return instanceSlice[i].GetName() < instanceSlice[j].GetName()
	})

	table := new(gongtable_models.Table).Stage(tableStageForSelection)
	table.Name = string(form.TableSelectExtraName)
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = false
	table.HasCheckableRows = true
	table.HasSaveButton = true

	// filterdInstanceSet is the set of instance that are part of the field
	filterdInstanceSet := make(map[FieldType]any, 0)
	for _, instance := range *onAssocEditon.field {
		filterdInstanceSet[instance] = nil
	}

	for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, instance := range instanceSlice {
		row := new(gongtable_models.Row).Stage(tableStageForSelection)
		row.Name = instance.GetName()
		table.Rows = append(table.Rows, row)

		if _, ok := filterdInstanceSet[instance]; ok {
			row.IsChecked = true
		}

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

	// set up control inversion for the saving of the table
	table.Impl = NewTablePickSaver[InstanceType, FieldType](
		onAssocEditon.instance,
		onAssocEditon.field,
		onAssocEditon.fieldName,
		onAssocEditon.probe)

	tableStageForSelection.Commit()
}

func NewTablePickSaver[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	fieldName string,
	probe *Probe,

) (tablePickSaver *TablePickSaver[InstanceType, FieldType]) {

	tablePickSaver = new(TablePickSaver[InstanceType, FieldType])
	tablePickSaver.instance = instance
	tablePickSaver.field = field
	tablePickSaver.fieldName = fieldName
	tablePickSaver.probe = probe

	return
}

type TablePickSaver[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	instance  InstanceType
	field     *[]FieldType
	fieldName string
	probe     *Probe
}

func (tablePickSaver *TablePickSaver[InstanceType, FieldType]) TableUpdated(
	stage *form.Stage,
	table, updatedTable *form.Table) {
	log.Println("TablePickSaver: TableUpdated")

	// checkout to the stage to get the rows that have been checked and not
	stage.Checkout()

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](tablePickSaver.probe.stageOfInterest)
	instanceSlice := make([]FieldType, 0)
	for instance := range instanceSet {
		instanceSlice = append(instanceSlice, instance)
	}
	sort.Slice(instanceSlice, func(i, j int) bool {
		return instanceSlice[i].GetName() < instanceSlice[j].GetName()
	})

	*tablePickSaver.field = make([]FieldType, 0)

	for idx, row := range table.Rows {
		if row.IsChecked {
			instance := instanceSlice[idx]
			*tablePickSaver.field = append(*tablePickSaver.field, instance)
		}
	}

	// commit the whole
	tablePickSaver.probe.stageOfInterest.Commit()

	// see the result
	fillUpTablePointerToGongstruct[InstanceType](
		tablePickSaver.probe,
	)
	tablePickSaver.probe.tableStage.Commit()
}
