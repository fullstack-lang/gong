// generated code - do not edit
package probe

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"
	form "github.com/fullstack-lang/gong/lib/table/go/models"
	gongtable_models "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
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
	formGroup *form.FormGroup,
	probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	// filteredInstanceSet is the set of instance that are part of the field
	filteredInstanceSet := make(map[FieldType]any, 0)
	for _, instance := range *field {
		filteredInstanceSet[instance] = nil
	}

	instanceSliceID := make([]uint, 0)
	for instance := range filteredInstanceSet {
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

	formEditAssocButton := (&form.FormEditAssocButton{
		Name:               fieldName,
		Label:              fieldName,
		AssociationStorage: storage,
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

	column := new(gongtable_models.DisplayedColumn).Stage(tableStageForSelection)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

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

		cell := (&gongtable_models.Cell{
			Name: "ID",
		}).Stage(tableStageForSelection)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable_models.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				onAssocEditon.probe.stageOfInterest,
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
