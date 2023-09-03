// generated code - do not edit
package data

import (
	"fmt"
	"log"
	"sort"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/test/go/models"
	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"
)

// AssociationSliceToForm add a form div with 2 buttons
// - one for selection
// - one for sorting
func AssociationSliceToForm[FieldType models.PointerToGongstruct](
	fieldName string, field *[]FieldType,
	stageOfInterest *models.StageStruct, formStage *form.StageStruct, formGroup *form.FormGroup,
	r *gin.Engine,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)

	formEditAssocButton := (&form.FormEditAssocButton{
		Name:  fieldName,
		Label: fieldName,
	}).Stage(formStage)
	formDiv.FormEditAssocButton = formEditAssocButton
	onAssocEditon := NewOnAssocEditon(r, formStage, field, stageOfInterest)
	formEditAssocButton.OnAssocEditon = onAssocEditon

	formSortAssocButton := (&form.FormSortAssocButton{
		Name:  fieldName,
		Label: fieldName,
	}).Stage(formStage)
	formDiv.FormSortAssocButton = formSortAssocButton
	onSortingEditon := NewOnSortingEditon(r, formStage)
	formSortAssocButton.OnSortEdition = onSortingEditon

}

func NewOnAssocEditon[FieldType models.PointerToGongstruct](
	r *gin.Engine,
	formStage *form.StageStruct,
	field *[]FieldType,
	stageOfInterest *models.StageStruct,
) (onAssocEdition *OnAssocEditon[FieldType]) {

	onAssocEdition = new(OnAssocEditon[FieldType])

	onAssocEdition.r = r
	onAssocEdition.formStage = formStage
	onAssocEdition.field = field
	onAssocEdition.stageOfInterest = stageOfInterest

	return
}

type OnAssocEditon[FieldType models.PointerToGongstruct] struct {
	r               *gin.Engine
	formStage       *form.StageStruct
	field           *[]FieldType
	stageOfInterest *models.StageStruct
}

func (onAssocEditon *OnAssocEditon[FieldType]) OnButtonPressed() {

	tableStackName := onAssocEditon.formStage.GetPath() + string(form.StackNamePostFixForTableForAssociation)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEditon.r, tableStackName)

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](onAssocEditon.stageOfInterest)
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
			cellString.Name = models.GetFieldStringValueFromPointer(instance, fieldName)
			cellString.Value = cellString.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}

	// set up control inversion for the saving of the table
	table.Impl = NewTablePickSaver[FieldType](onAssocEditon.field, onAssocEditon.stageOfInterest)

	tableStageForSelection.Commit()
}

func NewTablePickSaver[FieldType models.PointerToGongstruct](
	field *[]FieldType,
	stageOfInterest *models.StageStruct,
) (tablePickSaver *TablePickSaver[FieldType]) {

	tablePickSaver = new(TablePickSaver[FieldType])
	tablePickSaver.field = field
	tablePickSaver.stageOfInterest = stageOfInterest

	return
}

type TablePickSaver[FieldType models.PointerToGongstruct] struct {
	field           *[]FieldType
	stageOfInterest *models.StageStruct
}

func (tablePickSaver *TablePickSaver[FieldType]) TableUpdated(stage *form.StageStruct, table, updatedTable *form.Table) {
	log.Println("TablePickSaver: TableUpdated")

	// checkout to the stage to get the rows that have been checked and not
	stage.Checkout()

	instanceSet := *models.GetGongstructInstancesSetFromPointerType[FieldType](tablePickSaver.stageOfInterest)
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
}
