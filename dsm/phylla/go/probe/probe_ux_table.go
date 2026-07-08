// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"time"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
)

const TableName = "Table"

// update the current table if there is one
func (probe *Probe) ux_table() {
	var tableName string
	for table := range probe.tableStage.Tables {
		tableName = table.Name
	}
	switch tableName {
	// insertion point
	case "AxesShape":
		updateProbeTable[*models.AxesShape](probe)
	case "Library":
		updateProbeTable[*models.Library](probe)
	case "Plant":
		updateProbeTable[*models.Plant](probe)
	case "PlantDiagram":
		updateProbeTable[*models.PlantDiagram](probe)
	}
}

func updateProbeTable[T models.PointerToGongstruct](
	probe *Probe,
) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = models.GetPointerToGongstructName[T]()
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = probe.bulkDeleteMode
	table.HasBulkDeleteButton = probe.bulkDeleteMode
	table.BulkDeleteButtonTooltip = "Select rows to delete and click to delete them"
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

	// add a button for bulk delete
	bulkDeleteModeButton := &table_models.Button{
		Name: "BulkDeleteButton",
		Icon: func() string {
			if probe.bulkDeleteMode {
				return string(maticons.BUTTON_deselect)
			}
			return string(maticons.BUTTON_delete_sweep)
		}(),
		HasToolTip: true,
		ToolTipText: func() string {
			if probe.bulkDeleteMode {
				return "Cancel bulk delete"
			}
			return "Select instances to delete"
		}(),
		ToolTipPosition: table_models.Below,
	}

	bulkDeleteModeButton.OnClick = func() {
		probe.bulkDeleteMode = !probe.bulkDeleteMode
		updateProbeTable[T](probe)
	}
	table.Buttons = append(table.Buttons, bulkDeleteModeButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	if !probe.bulkDeleteMode {
		column = new(table_models.DisplayedColumn)
		column.Name = "Delete"
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	} else {
		table.OnUpdate = func(stage *table_models.Stage, updatedTable *table_models.Table) {
			// log.Println("Table updated, selectedRowIDsForBulkDelete:", updatedTable.BulkDeleteSelectedRowsIDsJson)

			for _, row := range updatedTable.RowsSelectedForBulkDelete {
				cellID := row.Cells[0]
				id := cellID.CellInt.Value
				instance := models.GongGetInstanceFromOrder[T](probe.stageOfInterest, uint(id))
				var zeroInstance T
				if instance == zeroInstance {
					continue
				}
				instance.UnstageVoid(probe.stageOfInterest)
			}
			if len(updatedTable.RowsSelectedForBulkDelete) > 0 {
				// after a delete of an instance, the stage might be dirty if a pointer or a slice of pointer
				// reference the deleted instance.
				// therefore, it is mandatory to clean the stage of interest
				probe.bulkDeleteMode = false
				probe.stageOfInterest.Commit()
				probe.Refresh()
			}

			probe.stageOfInterest.Commit()
		}
	}

	for _, fieldName := range fields {
		column := new(table_models.DisplayedColumn)
		column.Name = fieldName.Name
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(table_models.DisplayedColumn)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(table_models.Row)
		value := models.GetFieldStringValueFromPointer(structInstance, "Name", probe.stageOfInterest)
		row.Name = value.GetValueString()

		updater := NewRowUpdate(structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := &table_models.Cell{
			Name: "ID",
		}
		row.Cells = append(row.Cells, cell)
		cellInt := &table_models.CellInt{
			Name: "ID",
			Value: int(models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
		}
		cell.CellInt = cellInt

		cell = &table_models.Cell{
			Name: "Delete Icon",
		}
		row.Cells = append(row.Cells, cell)
		cellIcon := &table_models.CellIcon{
			Name: fmt.Sprintf("Delete Icon %d", models.GetOrderPointerGongstruct(
				probe.stageOfInterest,
				structInstance,
			)),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm tou want to delete this instance ?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, cellIcon, updatedCellIcon *table_models.CellIcon) {
				structInstance.UnstageVoid(probe.stageOfInterest)

				// after a delete of an instance, the stage might be dirty if a pointer or a slice of pointer
				// reference the deleted instance.
				// therefore, it is mandatory to clean the stage of interest
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()

				updateProbeTable[T](probe)
				probe.ux_tree()
			},
		}
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValueFromPointer(structInstance, fieldName.Name, probe.stageOfInterest)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value.GetValueString()
			fieldIndex++
			// log.Println(fieldName, value)
			cell := &table_models.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			switch value.GongFieldValueType {
			case models.GongFieldValueTypeInt:
				cellInt := &table_models.CellInt{
					Name:  name,
					Value: value.GetValueInt(),
				}
				cell.CellInt = cellInt
			case models.GongFieldValueTypeFloat:
				cellFloat := &table_models.CellFloat64{
					Name:  name,
					Value: value.GetValueFloat(),
				}
				cell.CellFloat64 = cellFloat
			case models.GongFieldValueTypeBool:
				cellBool := &table_models.CellBoolean{
					Name:  name,
					Value: value.GetValueBool(),
				}
				cell.CellBool = cellBool
			default:
				cellString := &table_models.CellString{
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
			cell := &table_models.Cell{
				Name: name,
			}
			row.Cells = append(row.Cells, cell)

			cellString := &table_models.CellString{
				Name:  name,
				Value: value,
			}
			cell.CellString = cellString
		}
	}

	table_models.StageBranch(probe.tableStage, table)

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

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	// log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}

func (probe *Probe) UpdateAndCommitNotificationTable() {
	probe.notificationTableStage.Reset()

	table := new(table_models.Table)
	table.Name = TableName
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	notificationsResetButton := &table_models.Button{
		Name:            "NotificationsResetButton",
		Icon:            string(tree_buttons.BUTTON_playlist_remove),
		HasToolTip:      true,
		ToolTipText:     "Reset notification table",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.ResetNotifications()
		},
	}
	table.Buttons = append(table.Buttons, notificationsResetButton)

	notificationsDownloadCSVButton := &table_models.Button{
		Name:            "NotificationsDownloadCSVButton",
		Icon:            string(tree_buttons.BUTTON_download),
		HasToolTip:      true,
		ToolTipText:     "Download notifications as CSV",
		ToolTipPosition: table_models.Below,
		OnClick: func() {
			probe.DownloadNotificationsCSV()
		},
	}
	table.Buttons = append(table.Buttons, notificationsDownloadCSVButton)

	column := new(table_models.DisplayedColumn)
	column.Name = "Date"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(table_models.DisplayedColumn)
	column.Name = "Message"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	// sort notifications by date
	sort.Slice(probe.notification, func(i, j int) bool {
		return probe.notification[i].Date.After(probe.notification[j].Date)
	})

	for _, notification := range probe.notification {
		row := new(table_models.Row)
		table.Rows = append(table.Rows, row)

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Date.Format(time.StampMicro)
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := new(table_models.Cell)
			cellString := new(table_models.CellString)
			cell.CellString = cellString
			cellString.Value = notification.Message
			row.Cells = append(row.Cells, cell)
		}
	}

	table_models.StageBranch(probe.notificationTableStage, table)

	probe.notificationTableStage.Commit()
}
