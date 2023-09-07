// generated code - do not edit
package data

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct         *gong_models.GongStruct
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	backRepoOfInterest *orm.BackRepoStruct
	r                  *gin.Engine
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	backRepoOfInterest *orm.BackRepoStruct,
	r *gin.Engine,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.tableStage = tableStage
	nodeImplGongstruct.formStage = formStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	nodeImplGongstruct.backRepoOfInterest = backRepoOfInterest
	nodeImplGongstruct.r = r
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

	tableStage := nodeImplGongstruct.tableStage

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "GongBasicField" {
		fillUpTable[models.GongBasicField](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnum" {
		fillUpTable[models.GongEnum](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongEnumValue" {
		fillUpTable[models.GongEnumValue](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongLink" {
		fillUpTable[models.GongLink](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongNote" {
		fillUpTable[models.GongNote](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongStruct" {
		fillUpTable[models.GongStruct](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "GongTimeField" {
		fillUpTable[models.GongTimeField](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Meta" {
		fillUpTable[models.Meta](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MetaReference" {
		fillUpTable[models.MetaReference](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ModelPkg" {
		fillUpTable[models.ModelPkg](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "PointerToGongStructField" {
		fillUpTable[models.PointerToGongStructField](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SliceOfPointerToGongStructField" {
		fillUpTable[models.SliceOfPointerToGongStructField](nodeImplGongstruct.stageOfInterest, tableStage, nodeImplGongstruct.formStage, nodeImplGongstruct.r, nodeImplGongstruct.backRepoOfInterest)
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	stageOfInterest *models.StageStruct,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) {

	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			stageOfInterest,
			backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance,
			tableStage,
			formStage,
			stageOfInterest,
			r,
			backRepoOfInterest)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				stageOfInterest,
				backRepoOfInterest,
				structInstance,
			),
		}).Stage(tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	tableStage *gongtable.StageStruct,
	formStage *gongtable.StageStruct,
	stageOfInterest *models.StageStruct,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.tableStage = tableStage
	rowUpdate.formStage = formStage
	rowUpdate.stageOfInterest = stageOfInterest
	rowUpdate.r = r
	rowUpdate.backRepoOfInterest = backRepoOfInterest
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance           *T
	tableStage         *gongtable.StageStruct
	formStage          *gongtable.StageStruct
	stageOfInterest    *models.StageStruct
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.GongBasicField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongBasicFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongEnum:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongEnumValue:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongEnumValueFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongLink:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongLinkFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongNote:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongNoteFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongStruct:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongStructFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.GongTimeField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGongTimeFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.Meta:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewMetaFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.MetaReference:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewMetaReferenceFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.ModelPkg:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewModelPkgFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.PointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewPointerToGongStructFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	case *models.SliceOfPointerToGongStructField:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewSliceOfPointerToGongStructFieldFormCallback(
				rowUpdate.stageOfInterest,
				rowUpdate.tableStage,
				formStage,
				instancesTyped,
				rowUpdate.r,
				rowUpdate.backRepoOfInterest,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, rowUpdate.stageOfInterest, formStage, formGroup, rowUpdate.r)
	}
	formStage.Commit()

}
