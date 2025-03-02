// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/lib/table/go/models"
	"github.com/fullstack-lang/gong/lib/table/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Cell:
		fillUpTable[models.Cell](probe)
	case *models.CellBoolean:
		fillUpTable[models.CellBoolean](probe)
	case *models.CellFloat64:
		fillUpTable[models.CellFloat64](probe)
	case *models.CellIcon:
		fillUpTable[models.CellIcon](probe)
	case *models.CellInt:
		fillUpTable[models.CellInt](probe)
	case *models.CellString:
		fillUpTable[models.CellString](probe)
	case *models.CheckBox:
		fillUpTable[models.CheckBox](probe)
	case *models.DisplayedColumn:
		fillUpTable[models.DisplayedColumn](probe)
	case *models.FormDiv:
		fillUpTable[models.FormDiv](probe)
	case *models.FormEditAssocButton:
		fillUpTable[models.FormEditAssocButton](probe)
	case *models.FormField:
		fillUpTable[models.FormField](probe)
	case *models.FormFieldDate:
		fillUpTable[models.FormFieldDate](probe)
	case *models.FormFieldDateTime:
		fillUpTable[models.FormFieldDateTime](probe)
	case *models.FormFieldFloat64:
		fillUpTable[models.FormFieldFloat64](probe)
	case *models.FormFieldInt:
		fillUpTable[models.FormFieldInt](probe)
	case *models.FormFieldSelect:
		fillUpTable[models.FormFieldSelect](probe)
	case *models.FormFieldString:
		fillUpTable[models.FormFieldString](probe)
	case *models.FormFieldTime:
		fillUpTable[models.FormFieldTime](probe)
	case *models.FormGroup:
		fillUpTable[models.FormGroup](probe)
	case *models.FormSortAssocButton:
		fillUpTable[models.FormSortAssocButton](probe)
	case *models.Option:
		fillUpTable[models.Option](probe)
	case *models.Row:
		fillUpTable[models.Row](probe)
	case *models.Table:
		fillUpTable[models.Table](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			probe.stageOfInterest,
			probe.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		value := models.GetFieldStringValue(*structInstance, "Name")
		row.Name = value.GetValueString()

		updater := NewRowUpdate[T](structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			)),
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue(*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := (&gongtable.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}).Stage(probe.tableStage)
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := (&gongtable.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}).Stage(probe.tableStage)
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := (&gongtable.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}).Stage(probe.tableStage)
				cell.CellBool = cellBool
			default:
				cellString := (&gongtable.CellString{
					Name:  name,
					Value: value.GetValueString(),
				}).Stage(probe.tableStage)
				cell.CellString = cellString

			}
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance *T
	probe    *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
