package probe

const FillUpTableTemplate = `// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"{{PkgPathRoot}}/models"
	"{{PkgPathRoot}}/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point{{` + string(rune(FillUpTableCaseForCastingDown)) + `}}
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()
	probe.tableStage.Commit()

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
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

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
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
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
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
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
	Instance   *T
	probe *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
`

type FillUpTableInsertionId int

const (
	FillUpTableCase FillUpTableInsertionId = iota
	FillUpTableCaseForCastingDown
	FillUpTableCaseForDeleteIcon
)

var FillUpTableSubTemplateCode map[string]string = // new line
map[string]string{

	string(rune(FillUpTableCaseForCastingDown)): `
	case *models.{{Structname}}:
		fillUpTable[models.{{Structname}}](probe)`,
	string(rune(FillUpTableCase)): `
	case *models.{{Structname}}:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update {{Structname}} Form",
			OnSave: __gong__New__{{Structname}}FormCallback(
				instancesTyped,
				rowUpdate.probe,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.probe)`,
	string(rune(FillUpTableCaseForDeleteIcon)): `
	case *models.{{Structname}}:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)`,
}
