// generated code - do not edit
package probe

import (
	"fmt"
	"sort"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/pkg/docx/go/models"
)

const TableName = "Table"

// update the current table if there is one
func updateCurrentProbeTable(probe *Probe) {
	var tableName string
	for table := range probe.tableStage.Tables {
		tableName = table.Name
	}
	switch tableName {
	// insertion point
	case "Body":
		updateProbeTable[*models.Body](probe)
	case "Document":
		updateProbeTable[*models.Document](probe)
	case "Docx":
		updateProbeTable[*models.Docx](probe)
	case "File":
		updateProbeTable[*models.File](probe)
	case "Node":
		updateProbeTable[*models.Node](probe)
	case "Paragraph":
		updateProbeTable[*models.Paragraph](probe)
	case "ParagraphProperties":
		updateProbeTable[*models.ParagraphProperties](probe)
	case "ParagraphStyle":
		updateProbeTable[*models.ParagraphStyle](probe)
	case "Rune":
		updateProbeTable[*models.Rune](probe)
	case "RuneProperties":
		updateProbeTable[*models.RuneProperties](probe)
	case "Table":
		updateProbeTable[*models.Table](probe)
	case "TableColumn":
		updateProbeTable[*models.TableColumn](probe)
	case "TableProperties":
		updateProbeTable[*models.TableProperties](probe)
	case "TableRow":
		updateProbeTable[*models.TableRow](probe)
	case "TableStyle":
		updateProbeTable[*models.TableStyle](probe)
	case "Text":
		updateProbeTable[*models.Text](probe)
	}
}

func updateProbeTable[T models.PointerToGongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()

	table := new(gongtable.Table)
	table.Name = models.GetPointerToGongstructName[T]()
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFieldsFromPointer[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	setOfStructs := (*models.GetGongstructInstancesSetFromPointerType[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return models.GetOrderPointerGongstruct(probe.stageOfInterest, sliceOfGongStructsSorted[i]) <
			models.GetOrderPointerGongstruct(probe.stageOfInterest, sliceOfGongStructsSorted[j])
	})

	column := new(gongtable.DisplayedColumn)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn)
		column.Name = fieldName.Name
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row)
		value := models.GetFieldStringValueFromPointer(structInstance, "Name", probe.stageOfInterest)
		row.Name = value.GetValueString()

		updater := NewRowUpdate(structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := &gongtable.Cell{
			Name: "ID",
		}
		row.Cells = append(row.Cells, cell)
		cellInt := &gongtable.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
		}
		cell.CellInt = cellInt

		cell = &gongtable.Cell{
			Name: "Delete Icon",
		}
		row.Cells = append(row.Cells, cell)
		cellIcon := &gongtable.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm tou want to delete this instance ?",
		}
		cellIcon.Impl = &gongtable.FunctionalCellIconProxy{
			OnUpdated: func(stage *gongtable.Stage, cellIcon, updatedCellIcon *gongtable.CellIcon) {
				structInstance.UnstageVoid(probe.stageOfInterest)

				// after a delete of an instance, the stage might be dirty if a pointer or a slice of pointer
				// reference the deleted instance.
				// therefore, it is mandatory to clean the stage of interest
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()

				updateProbeTable[T](probe)
				updateAndCommitTree(probe)
			},
		}
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValueFromPointer(structInstance, fieldName.Name, probe.stageOfInterest)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &gongtable.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := &gongtable.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := &gongtable.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := &gongtable.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}
				cell.CellBool = cellBool
			default:
				cellString := &gongtable.CellString{
					Name:  name,
					Value: value.GetValueString(),
				}
				cell.CellString = cellString

			}
		}
		for _, reverseField := range reverseFields {

			value := structInstance.GongGetReverseFieldOwnerName(
				probe.stageOfInterest,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &gongtable.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			cellString := &gongtable.CellString{
				Name:  name,
				Value: value,
			}
			cell.CellString = cellString
		}
	}

	gongtable.StageBranch(probe.tableStage, table)

	probe.tableStage.Commit()
}

func NewRowUpdate[T models.PointerToGongstruct](
	Instance T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.PointerToGongstruct] struct {
	Instance T
	probe    *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.Stage, row, updatedRow *gongtable.Row) {
	// log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}

func (probe *Probe) UpdateAndCommitNotificationTable() {
	probe.notificationTableStage.Reset()

	table := new(gongtable.Table)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	column := new(gongtable.DisplayedColumn)
	column.Name = "Date"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn)
	column.Name = "Message"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	// sort notifications by date
	sort.Slice(probe.notification, func(i, j int) bool {
		return probe.notification[i].Date.After(probe.notification[j].Date)
	})

	for _, notification := range probe.notification {
		row := new(gongtable.Row)
		table.Rows = append(table.Rows, row)

		{
			cell := new(gongtable.Cell)
			cellString := new(gongtable.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Date.Format("2006-01-02 15:04:05")
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := new(gongtable.Cell)
			cellString := new(gongtable.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Message
			row.Cells = append(row.Cells, cell)
		}
	}

	gongtable.StageBranch(probe.notificationTableStage, table)

	probe.notificationTableStage.Commit()
}
