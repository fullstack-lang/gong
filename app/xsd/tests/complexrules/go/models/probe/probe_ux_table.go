// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/common"
)

const TableName = "Table"

var (
	_ = fmt.Sprintf
	_ = sort.Strings
	_ = strings.Join
	_ = maticons.BUTTON_delete
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *Probe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
	case "models.Catalog", "Catalog":
		updateStageSetTable_Catalog_Stage(probe)
	case "models.DerivedItemType", "DerivedItemType":
		updateStageSetTable_DerivedItemType_Stage(probe)
	case "refs.RefTargetType":
		updateStageSetTable_RefTargetType_RefsStage(probe)
	case "common.BaseItemType":
		updateStageSetTable_BaseItemType_CommonStage(probe)
	}
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

	probe.notificationTableStage.StageBranch(table)
	probe.notificationTableStage.Commit()
}

func updateStageSetTable_Catalog_Stage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Catalog"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Target"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Item"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GetInstancesByOrder[*models.Catalog]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageOfInterest)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_Catalog_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Target"}
			var names []string
			for _, elem := range structInstance.Target {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Item"}
			var names []string
			for _, elem := range structInstance.Item {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_DerivedItemType_Stage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.DerivedItemType"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Code"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Score"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TargetRef"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OptionA"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OptionB"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Catalog) -> Item"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.GetInstancesByOrder[*models.DerivedItemType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageOfInterest)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_DerivedItemType_Stage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Code"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Code)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Score"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Score)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TargetRef"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.TargetRef)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OptionA"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.OptionA)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OptionB"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.OptionB)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Catalog) -> Item"}
			var refNames []string
			for src := range probe.stageOfInterest.Catalogs {
				for _, target := range src.Item {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_RefTargetType_RefsStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "refs.RefTargetType"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Id"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Status"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CreatedBy"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Label"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Catalog) -> Target"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.RefsStage.GetInstancesByOrder[*refs.RefTargetType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.RefsStage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageOfInterest.RefsStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_RefTargetType_RefsStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Id"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Id)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Status"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Status)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CreatedBy"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.CreatedBy)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Label"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Label)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Catalog) -> Target"}
			var refNames []string
			for src := range probe.stageOfInterest.Catalogs {
				for _, target := range src.Target {
					if target == structInstance {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_BaseItemType_CommonStage(probe *Probe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "common.BaseItemType"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true

	colID := new(table_models.DisplayedColumn)
	colID.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, colID)

	colDel := new(table_models.DisplayedColumn)
	colDel.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, colDel)

	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Name"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Code"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Score"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageOfInterest.CommonStage.GetInstancesByOrder[*common.BaseItemType]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				FillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageOfInterest.CommonStage.GetOrder(structInstance))}
		row.Cells = append(row.Cells, cellID)

		cellDel := &table_models.Cell{Name: "Delete Icon"}
		cellIcon := &table_models.CellIcon{
			Name:                fmt.Sprintf("Delete %s", structInstance.GetName()),
			Icon:                string(maticons.BUTTON_delete),
			NeedsConfirmation:   true,
			ConfirmationMessage: "Do you confirm you want to delete this instance?",
		}
		cellIcon.Impl = &table_models.FunctionalCellIconProxy{
			OnUpdated: func(stage *table_models.Stage, ci, uci *table_models.CellIcon) {
				_captured.UnstageVoid(probe.stageOfInterest.CommonStage)
				probe.stageOfInterest.Clean()
				probe.stageOfInterest.Commit()
				updateStageSetTable_BaseItemType_CommonStage(probe)
				probe.ux_tree()
				if probe.docStager != nil {
					probe.docStager.SetMap_GongStructName_InstancesNb(probe.ComputeInstancesNb())
					probe.docStager.Svg()
				}
			},
		}
		cellDel.CellIcon = cellIcon
		row.Cells = append(row.Cells, cellDel)


		{
			cell := &table_models.Cell{Name: "Name"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Name)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Code"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Code)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Score"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Score)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

