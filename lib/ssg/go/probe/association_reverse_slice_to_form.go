// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable_fullstack "github.com/fullstack-lang/gong/lib/table/go/fullstack"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"
)

// AssociationReverseSliceToForm add a form div with 1 button
// - one for selection of A instances
func AssociationReverseSliceToForm[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	reverseFieldOwnerTypeName string,
	reverseFieldName string,
	instance FieldType,
	formGroup *form.FormGroup,
	probe *Probe,
	getSlice func(owner OwnerType) []FieldType,
) {

	formDiv := (&form.FormDiv{
		Name: reverseFieldOwnerTypeName + ":" + reverseFieldName,
	}).Stage(probe.formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	map_ID_RowID := GetMap_ID_RowID[OwnerType](probe.stageOfInterest)

	rowIDsOfInstancesInField := make([]uint, 0)
	{
		ownerSet := *models.GetGongstructInstancesSetFromPointerType[OwnerType](probe.stageOfInterest)
		for owner := range ownerSet {
			slice := getSlice(owner)
			// check if instance is in slice
			for _, item := range slice {
				if any(item) == any(instance) {
					id := models.GetOrderPointerGongstruct(probe.stageOfInterest, owner)
					rowIDsOfInstancesInField = append(rowIDsOfInstancesInField, uint(map_ID_RowID[id]))
					break
				}
			}
		}
	}

	storage, err := EncodeIntSliceToString(rowIDsOfInstancesInField)
	if err != nil {
		log.Panic("Unable to encode association")
	}

	formEditAssocButton := (&form.FormEditAssocButton{
		Name:                reverseFieldOwnerTypeName + ":" + reverseFieldName,
		Label:               reverseFieldOwnerTypeName + ":" + reverseFieldName,
		AssociationStorage:  storage,
		HasToolTip:          true,
		MatTooltipShowDelay: "1500",
		ToolTipText:         "Edit list of the instances of " + models.GetPointerToGongstructName[OwnerType]() + " associated to this field",
	}).Stage(probe.formStage)
	formDiv.FormEditAssocButton = formEditAssocButton
	onAssocEditon := NewOnReverseAssocEditon[OwnerType, FieldType](
		reverseFieldOwnerTypeName,
		reverseFieldName,
		instance,
		probe,
	)
	formEditAssocButton.OnAssocEditon = onAssocEditon
}

type OnReverseAssocEditon[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct] struct {
	reverseFieldOwnerTypeName string
	reverseFieldName          string
	instance                  FieldType
	probe                     *Probe
}

func NewOnReverseAssocEditon[OwnerType models.PointerToGongstruct, FieldType models.PointerToGongstruct](
	reverseFieldOwnerTypeName string,
	reverseFieldName          string,
	instance FieldType,
	probe *Probe,
) (onAssocEdition *OnReverseAssocEditon[OwnerType, FieldType]) {

	onAssocEdition = new(OnReverseAssocEditon[OwnerType, FieldType])
	onAssocEdition.reverseFieldOwnerTypeName = reverseFieldOwnerTypeName
	onAssocEdition.reverseFieldName = reverseFieldName
	onAssocEdition.instance = instance
	onAssocEdition.probe = probe

	return
}

func (onAssocEditon *OnReverseAssocEditon[OwnerType, FieldType]) OnButtonPressed() {

	tableStackName := onAssocEditon.probe.formStage.GetName() + string(table.StackNamePostFixForTableForAssociation)

	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEditon.probe.r, tableStackName)

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[OwnerType](onAssocEditon.probe.stageOfInterest)
	instanceSlice := make([]OwnerType, 0)
	for instance := range instanceSet {
		instanceSlice = append(instanceSlice, instance)
	}

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

	for _, fieldName := range models.GetFieldsFromPointer[OwnerType]() {
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
			Name:  "ID",
			Value: int(models.GetOrderPointerGongstruct(onAssocEditon.probe.stageOfInterest, instance)),
		}).Stage(tableStageForSelection)
		cell.CellInt = cellInt

		for _, fieldName := range models.GetFieldsFromPointer[OwnerType]() {
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
