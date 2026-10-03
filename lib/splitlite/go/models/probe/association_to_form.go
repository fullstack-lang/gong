// generated code - do not edit
package probe

import (
	"fmt"
	"log"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/splitlite/go/models"
)

func GetAssociationOptionName[FieldType models.PointerToGongstruct](
	instance FieldType,
	stageOfInterest *models.Stage,
	nameCount map[string]int,
) string {
	name := instance.GetName()
	id := stageOfInterest.GetOrder(instance)
	if name == "" {
		return fmt.Sprintf("[Unnamed] (ID: %d)", id)
	}
	if nameCount[name] > 1 {
		return fmt.Sprintf("%s (ID: %d)", name, id)
	}
	return name
}

// AssociationFieldToForm will append a div to the form
// with a dropdown (with disambiguated options) and a table picker button
func AssociationFieldToForm[FieldType models.PointerToGongstruct](
	fieldName string, field FieldType, formGroup *form.FormGroup, probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(probe.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:                 "association",
		CanBeEmpty:           true,
		PreserveInitialOrder: true,
	}).Stage(probe.formStage)
	formField.FormFieldSelect = formFieldSelect

	orderedByIdInstance := probe.stageOfInterest.GetInstancesByOrder[FieldType]()

	nameCount := make(map[string]int)
	for _, instance := range orderedByIdInstance {
		nameCount[instance.GetName()]++
	}

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for _, instance := range orderedByIdInstance {
		optionName := GetAssociationOptionName(instance, probe.stageOfInterest, nameCount)
		option := (&form.Option{
			Name: optionName,
		}).Stage(probe.formStage)

		// set up select value if field matches the instance
		if instance == field {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}

	// table picker button for 1-1 association
	map_ID_RowID := GetMap_ID_RowID[FieldType](probe.stageOfInterest)

	rowIDsOfInstancesInField := make([]uint, 0)
	var zero FieldType
	if field != zero {
		id := probe.stageOfInterest.GetOrder(field)
		if rowID, ok := map_ID_RowID[id]; ok {
			rowIDsOfInstancesInField = append(rowIDsOfInstancesInField, uint(rowID))
		}
	}

	storage, err := EncodeIntSliceToString(rowIDsOfInstancesInField)
	if err != nil {
		log.Panic("Unable to encode association")
	}

	formEditAssocButton := (&form.FormEditAssocButton{
		Name:                fieldName,
		Label:               "",
		AssociationStorage:  storage,
		HasToolTip:          true,
		MatTooltipShowDelay: "1500",
		ToolTipText:         "Select " + models.GetPointerToGongstructName[FieldType]() + " from table",
	}).Stage(probe.formStage)
	formDiv.FormEditAssocButton = formEditAssocButton
	onAssocEdition := NewOnAssocEditonSingle(field, probe)
	formEditAssocButton.OnAssocEditon = onAssocEdition
}

type OnAssocEditonSingle[FieldType models.PointerToGongstruct] struct {
	field FieldType
	probe *Probe
}

func NewOnAssocEditonSingle[FieldType models.PointerToGongstruct](
	field FieldType,
	probe *Probe,
) (onAssocEdition *OnAssocEditonSingle[FieldType]) {

	onAssocEdition = new(OnAssocEditonSingle[FieldType])
	onAssocEdition.field = field
	onAssocEdition.probe = probe

	return
}

func (onAssocEdition *OnAssocEditonSingle[FieldType]) OnButtonPressed() {

	tableStackName := onAssocEdition.probe.formStage.GetName() + string(table.StackNamePostFixForTableForAssociation)

	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEdition.probe.r, tableStackName)
	tableStageForSelection.Reset()

	instanceSlice := onAssocEdition.probe.stageOfInterest.GetInstancesByOrder[FieldType]()

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
			Value: int(onAssocEdition.probe.stageOfInterest.GetOrder(
				instance,
			)),
		}).Stage(tableStageForSelection)
		cell.CellInt = cellInt

		for _, fieldName := range models.GetFieldsFromPointer[FieldType]() {
			cell := new(table.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %s - Column %s", instance.GetName(), fieldName)

			cellString := new(table.CellString).Stage(tableStageForSelection)
			value := models.GetFieldStringValueFromPointer(instance, fieldName.Name, onAssocEdition.probe.stageOfInterest)
			cellString.Name = value.GetValueString()
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	tableStageForSelection.Commit()
}

func AssociationReverseFieldToForm[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	owner OwnerType,
	fieldName string,
	instance FieldType,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	formDiv := (&form.FormDiv{
		Name: models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Label:       models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		Placeholder: "",
	}).Stage(probe.formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:                 models.GetPointerToGongstructName[OwnerType]() + ":" + fieldName,
		CanBeEmpty:           true,
		PreserveInitialOrder: true,
	}).Stage(probe.formStage)
	formField.FormFieldSelect = formFieldSelect

	orderedByIdInstance := probe.stageOfInterest.GetInstancesByOrder[OwnerType]()

	nameCount := make(map[string]int)
	for _, _instance := range orderedByIdInstance {
		nameCount[_instance.GetName()]++
	}

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for _, _instance := range orderedByIdInstance {
		optionName := GetAssociationOptionName(_instance, probe.stageOfInterest, nameCount)
		option := (&form.Option{
			Name: optionName,
		}).Stage(probe.formStage)

		// set up select value if field matches the instance
		var zero OwnerType
		if owner != zero && _instance == owner {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}
