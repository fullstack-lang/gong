// generated code - do not edit
package x

import (
	"fmt"
	"log"
	"sort"

	"github.com/fullstack-lang/gong/test2/go/models/x"
	orm_x "github.com/fullstack-lang/gong/test2/go/orm/x"
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"
)

func fillUpTablePointerToGongstruct[T x.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *x.A:
		fillUpTable[x.A](playground)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T x.Gongstruct](
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

	fields := x.GetFields[T]()
	reverseFields := x.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	playground.stageOfInterest.Checkout()

	setOfStructs := (*x.GetGongstructInstancesSet[T](playground.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm_x.GetID(
			playground.stageOfInterest,
			playground.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm_x.GetID(
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
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(playground.tableStage)
		row.Name = x.GetFieldStringValue[T](*structInstance, "Name")

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
			Value: orm_x.GetID(
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
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, playground)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := x.GetFieldStringValue[T](*structInstance, fieldName)
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
		for _, reverseField := range reverseFields {

			value := orm_x.GetReverseFieldOwnerName[T](
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				structInstance,
				&reverseField)
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

func NewRowUpdate[T x.Gongstruct](
	Instance *T,
	playground *Playground,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.playground = playground
	return
}

type RowUpdate[T x.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.playground)
}
