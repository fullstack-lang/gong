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

	"github.com/fullstack-lang/gong/test/go/models"
	"github.com/fullstack-lang/gong/test/go/orm"
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
	if nodeImplGongstruct.gongStruct.GetName() == "Astruct" {
		fillUpTable[models.Astruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AstructBstruct2Use" {
		fillUpTable[models.AstructBstruct2Use](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AstructBstructUse" {
		fillUpTable[models.AstructBstructUse](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bstruct" {
		fillUpTable[models.Bstruct](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Dstruct" {
		fillUpTable[models.Dstruct](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Astruct:
		fillUpTable[models.Astruct](playground)
	case *models.AstructBstruct2Use:
		fillUpTable[models.AstructBstruct2Use](playground)
	case *models.AstructBstructUse:
		fillUpTable[models.AstructBstructUse](playground)
	case *models.Bstruct:
		fillUpTable[models.Bstruct](playground)
	case *models.Dstruct:
		fillUpTable[models.Dstruct](playground)
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
	case *models.Astruct:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Astruct Form",
			OnSave: NewAstructFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.AstructBstruct2Use:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update AstructBstruct2Use Form",
			OnSave: NewAstructBstruct2UseFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.AstructBstructUse:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update AstructBstructUse Form",
			OnSave: NewAstructBstructUseFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Bstruct:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Bstruct Form",
			OnSave: NewBstructFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Dstruct:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Dstruct Form",
			OnSave: NewDstructFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	}
	formStage.Commit()

}
