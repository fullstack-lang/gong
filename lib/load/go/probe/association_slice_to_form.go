// generated code - do not edit
package probe

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/load/go/models"
)

// EncodeIntSliceToString encodes a slice of integers into a JSON string.
// It returns the JSON string and an error if marshalling fails.
func EncodeIntSliceToString(data []uint) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal uint slice to JSON: %w", err)
	}
	return string(jsonData), nil
}

// DecodeStringToIntSlice decodes a JSON string into a slice of integers.
// It returns the slice of integers and an error if unmarshalling fails
// or if the string is not a valid JSON representation of an int slice.
func DecodeStringToIntSlice(str string) ([]uint, error) {
	var decodedData []uint
	err := json.Unmarshal([]byte(str), &decodedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON string to uint slice: %w", err)
	}
	// Note: json.Unmarshal will also return an error if the JSON structure
	// doesn't match []int (e.g., if it's a JSON object or an array of strings).
	// So, an explicit type check like in the TypeScript example is less critical here
	// as Unmarshal handles type mismatches by returning an error.
	return decodedData, nil
}

// AssociationSliceToForm add a form div with 2 buttons
// - one for selection
// - one for sorting
func AssociationSliceToForm[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	fieldName string,
	instance InstanceType,
	field *[]FieldType,
	formGroup *table.FormGroup,
	probe *Probe,
) {

	formDiv := (&table.FormDiv{
		Name: fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	// here,  it is supposed that the table presented to the user for choosing
	// associations  of all instances is ordered by IDs
	instanceSliceID := make([]uint, 0)
	for _, instance := range *field {
		id := uint(models.GetOrderPointerGongstruct(
			probe.stageOfInterest,
			instance,
		))
		instanceSliceID = append(instanceSliceID, id)
	}
	storage, err := EncodeIntSliceToString(instanceSliceID)
	if err != nil {
		log.Panic("Unable to encode association")
	}

	formEditAssocButton := (&table.FormEditAssocButton{
		Name:                fieldName,
		Label:               fieldName,
		AssociationStorage:  storage,
		HasToolTip:          true,
		MatTooltipShowDelay: "1500",
		ToolTipText:         "Edit list of the instances of " + models.GetPointerToGongstructName[FieldType]() + " associated to this field",
	}).Stage(probe.formStage)
	formDiv.FormEditAssocButton = formEditAssocButton
	onAssocEditon := NewOnAssocEditon(instance, field, probe)
	formEditAssocButton.OnAssocEditon = onAssocEditon

	formSortAssocButton := (&table.FormSortAssocButton{
		Name:                fieldName,
		Label:               fieldName,
		HasToolTip:          true,
		MatTooltipShowDelay: "1500",
		ToolTipText:         "Edit order within the list instances of " + models.GetPointerToGongstructName[FieldType]() + " associated to this field",
		FormEditAssocButton: formEditAssocButton,
	}).Stage(probe.formStage)
	formDiv.FormSortAssocButton = formSortAssocButton
	onSortingEditon := NewOnSortingEditon(instance, field, probe)
	formSortAssocButton.OnSortEdition = onSortingEditon

}

type OnAssocEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	instance InstanceType
	field    *[]FieldType
	probe    *Probe
}

func NewOnAssocEditon[InstanceType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	instance InstanceType,
	field *[]FieldType,
	probe *Probe,
) (onAssocEdition *OnAssocEditon[InstanceType, FieldType]) {

	onAssocEdition = new(OnAssocEditon[InstanceType, FieldType])
	onAssocEdition.instance = instance
	onAssocEdition.field = field
	onAssocEdition.probe = probe

	return
}

func (onAssocEditon *OnAssocEditon[InstanceType, FieldType]) OnButtonPressed() {

	tableStackName := onAssocEditon.probe.formStage.GetName() + string(table.StackNamePostFixForTableForAssociation)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEditon.probe.r, tableStackName)

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](onAssocEditon.probe.stageOfInterest)
	instanceSlice := make([]FieldType, 0)
	for instance := range instanceSet {
		instanceSlice = append(instanceSlice, instance)
	}

	// we supposed that the table is ordered by ID
	sort.Slice(instanceSlice, func(i, j int) bool {
		idI := models.GetOrderPointerGongstruct(onAssocEditon.probe.stageOfInterest, instanceSlice[i])
		idJ := models.GetOrderPointerGongstruct(onAssocEditon.probe.stageOfInterest, instanceSlice[j])

		return idI < idJ
	})

	instancesTable := new(table.Table).Stage(tableStageForSelection)
	instancesTable.Name = string(table.TableSelectExtraName)
	instancesTable.HasColumnSorting = true
	instancesTable.HasFiltering = true
	instancesTable.HasPaginator = false
	instancesTable.HasCheckableRows = true
	instancesTable.HasSaveButton = true

	column := new(table.DisplayedColumn).Stage(tableStageForSelection)
	column.Name = "ID"
	instancesTable.DisplayedColumns = append(instancesTable.DisplayedColumns, column)

	// filterdInstanceSet is the set of instance that are part of the field
	filterdInstanceSet := make(map[FieldType]any, 0)
	for _, instance := range *onAssocEditon.field {
		filterdInstanceSet[instance] = nil
	}

	for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
		column := new(table.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fieldName.Name
		instancesTable.DisplayedColumns = append(instancesTable.DisplayedColumns, column)
	}
	for _, instance := range instanceSlice {
		row := new(table.Row).Stage(tableStageForSelection)
		row.Name = instance.GetName()
		instancesTable.Rows = append(instancesTable.Rows, row)

		cell := (&table.Cell{
			Name: "ID",
		}).Stage(tableStageForSelection)
		row.Cells = append(row.Cells, cell)
		cellInt := (&table.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				onAssocEditon.probe.stageOfInterest,
				instance,
			)),
		}).Stage(tableStageForSelection)
		cell.CellInt = cellInt

		for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
			cell := new(table.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %s - Column %s", instance.GetName(), fieldName)

			cellString := new(table.CellString).Stage(tableStageForSelection)
			value := models.GetFieldStringValueFromPointer(instance, fieldName.Name, onAssocEditon.probe.stageOfInterest)
			cellString.Name = value.GetValueString()
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	tableStageForSelection.Commit()
}
