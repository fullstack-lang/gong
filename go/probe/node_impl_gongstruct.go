// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "GongBasicField" {
		fillUpTable[models.GongBasicField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnum" {
		fillUpTable[models.GongEnum](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValue" {
		fillUpTable[models.GongEnumValue](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongLink" {
		fillUpTable[models.GongLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNote" {
		fillUpTable[models.GongNote](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStruct" {
		fillUpTable[models.GongStruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongTimeField" {
		fillUpTable[models.GongTimeField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Meta" {
		fillUpTable[models.Meta](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MetaReference" {
		fillUpTable[models.MetaReference](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ModelPkg" {
		fillUpTable[models.ModelPkg](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "PointerToGongStructField" {
		fillUpTable[models.PointerToGongStructField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SliceOfPointerToGongStructField" {
		fillUpTable[models.SliceOfPointerToGongStructField](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.GongBasicField:
		fillUpTable[models.GongBasicField](playground)
	case *models.GongEnum:
		fillUpTable[models.GongEnum](playground)
	case *models.GongEnumValue:
		fillUpTable[models.GongEnumValue](playground)
	case *models.GongLink:
		fillUpTable[models.GongLink](playground)
	case *models.GongNote:
		fillUpTable[models.GongNote](playground)
	case *models.GongStruct:
		fillUpTable[models.GongStruct](playground)
	case *models.GongTimeField:
		fillUpTable[models.GongTimeField](playground)
	case *models.Meta:
		fillUpTable[models.Meta](playground)
	case *models.MetaReference:
		fillUpTable[models.MetaReference](playground)
	case *models.ModelPkg:
		fillUpTable[models.ModelPkg](playground)
	case *models.PointerToGongStructField:
		fillUpTable[models.PointerToGongStructField](playground)
	case *models.SliceOfPointerToGongStructField:
		fillUpTable[models.SliceOfPointerToGongStructField](playground)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	playground *Playground,
) {

	playground.tableStage.Reset()
	playground.tableStage.Commit()

	table := new(gongtable.Table).Stage(playground.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	playground.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](playground.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			playground.stageOfInterest,
			playground.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(playground.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, playground)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				structInstance,
			),
		}).Stage(playground.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(playground.tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(playground.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(playground.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.playground = playground
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongBasicFieldFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongEnum:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongEnumValue:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongLink:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongLinkFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongNote:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongNoteFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongStruct:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.GongTimeField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongTimeFieldFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Meta:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewMetaFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.MetaReference:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewMetaReferenceFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.ModelPkg:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewModelPkgFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.PointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPointerToGongStructFieldFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.SliceOfPointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	}
	formStage.Commit()

}
