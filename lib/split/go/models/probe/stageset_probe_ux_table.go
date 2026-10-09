// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/lib/split/go/models"
)

type tableRowUpdater struct {
	onClick func()
}

func (u *tableRowUpdater) RowUpdated(stage *table_models.Stage, row, updatedRow *table_models.Row) {
	if u.onClick != nil {
		u.onClick()
	}
}

func (probe *StageSetProbe) ux_table() {
	var tableName string
	for tbl := range probe.tableStage.Tables {
		tableName = tbl.Name
	}
	switch tableName {
	case "models.AsSplit":
		updateStageSetTable_AsSplit_Stage(probe)
	case "models.AsSplitArea":
		updateStageSetTable_AsSplitArea_Stage(probe)
	case "models.Button":
		updateStageSetTable_Button_Stage(probe)
	case "models.Cursor":
		updateStageSetTable_Cursor_Stage(probe)
	case "models.FavIcon":
		updateStageSetTable_FavIcon_Stage(probe)
	case "models.Form":
		updateStageSetTable_Form_Stage(probe)
	case "models.Load":
		updateStageSetTable_Load_Stage(probe)
	case "models.LogoOnTheLeft":
		updateStageSetTable_LogoOnTheLeft_Stage(probe)
	case "models.LogoOnTheRight":
		updateStageSetTable_LogoOnTheRight_Stage(probe)
	case "models.Markdown":
		updateStageSetTable_Markdown_Stage(probe)
	case "models.Slider":
		updateStageSetTable_Slider_Stage(probe)
	case "models.Split":
		updateStageSetTable_Split_Stage(probe)
	case "models.Svg":
		updateStageSetTable_Svg_Stage(probe)
	case "models.Table":
		updateStageSetTable_Table_Stage(probe)
	case "models.Threejs":
		updateStageSetTable_Threejs_Stage(probe)
	case "models.Title":
		updateStageSetTable_Title_Stage(probe)
	case "models.Tone":
		updateStageSetTable_Tone_Stage(probe)
	case "models.Tree":
		updateStageSetTable_Tree_Stage(probe)
	case "models.View":
		updateStageSetTable_View_Stage(probe)
	case "models.Xlsx":
		updateStageSetTable_Xlsx_Stage(probe)
	}
}

func updateStageSetTable_AsSplit_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.AsSplit"
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
		col.Name = "Direction"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AsSplitAreas"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsSizeInPixel"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsWithCustomGutterSize"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GutterSize"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> AsSplit"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.AsSplit]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_AsSplit_Stage(probe)
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
			cell := &table_models.Cell{Name: "Direction"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Direction)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AsSplitAreas"}
			var names []string
			for _, elem := range structInstance.AsSplitAreas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsSizeInPixel"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsSizeInPixel}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsWithCustomGutterSize"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsWithCustomGutterSize}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GutterSize"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.GutterSize)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> AsSplit"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.AsSplit == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_AsSplitArea_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.AsSplitArea"
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
		col.Name = "ShowNameInHeader"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Size"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsAny"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AsSplit"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Button"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Cursor"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Form"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Load"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Markdown"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Slider"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Split"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Svg"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Table"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Tone"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Tree"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Threejs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Xlsx"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HasDiv"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "DivStyle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplit) -> AsSplitAreas"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.View) -> RootAsSplitAreas"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.AsSplitArea]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_AsSplitArea_Stage(probe)
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
			cell := &table_models.Cell{Name: "ShowNameInHeader"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowNameInHeader}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Size"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Size)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsAny"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsAny}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AsSplit"}
			val := ""
			if structInstance.AsSplit != nil {
				val = structInstance.AsSplit.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Button"}
			val := ""
			if structInstance.Button != nil {
				val = structInstance.Button.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Cursor"}
			val := ""
			if structInstance.Cursor != nil {
				val = structInstance.Cursor.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Form"}
			val := ""
			if structInstance.Form != nil {
				val = structInstance.Form.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Load"}
			val := ""
			if structInstance.Load != nil {
				val = structInstance.Load.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Markdown"}
			val := ""
			if structInstance.Markdown != nil {
				val = structInstance.Markdown.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Slider"}
			val := ""
			if structInstance.Slider != nil {
				val = structInstance.Slider.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Split"}
			val := ""
			if structInstance.Split != nil {
				val = structInstance.Split.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Svg"}
			val := ""
			if structInstance.Svg != nil {
				val = structInstance.Svg.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Table"}
			val := ""
			if structInstance.Table != nil {
				val = structInstance.Table.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Tone"}
			val := ""
			if structInstance.Tone != nil {
				val = structInstance.Tone.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Tree"}
			val := ""
			if structInstance.Tree != nil {
				val = structInstance.Tree.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Threejs"}
			val := ""
			if structInstance.Threejs != nil {
				val = structInstance.Threejs.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Xlsx"}
			val := ""
			if structInstance.Xlsx != nil {
				val = structInstance.Xlsx.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HasDiv"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.HasDiv}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "DivStyle"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.DivStyle)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplit) -> AsSplitAreas"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplits {
				for _, target := range src.AsSplitAreas {
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

		{
			cell := &table_models.Cell{Name: "(models.View) -> RootAsSplitAreas"}
			var refNames []string
			for src := range probe.stageSet.Stage.Views {
				for _, target := range src.RootAsSplitAreas {
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

func updateStageSetTable_Button_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Button"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Button"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Button]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Button_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Button"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Button == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Cursor_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Cursor"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Style"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Cursor"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Cursor]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Cursor_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Style"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Style)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Cursor"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Cursor == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_FavIcon_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.FavIcon"
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
		col.Name = "SVG"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.FavIcon]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_FavIcon_Stage(probe)
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
			cell := &table_models.Cell{Name: "SVG"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SVG)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Form_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Form"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Form"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Form]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Form_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Form"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Form == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Load_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Load"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Load"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Load]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Load_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Load"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Load == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_LogoOnTheLeft_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.LogoOnTheLeft"
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
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SVG"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.LogoOnTheLeft]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_LogoOnTheLeft_Stage(probe)
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
			cell := &table_models.Cell{Name: "Width"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SVG"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SVG)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_LogoOnTheRight_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.LogoOnTheRight"
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
		col.Name = "Width"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SVG"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.LogoOnTheRight]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_LogoOnTheRight_Stage(probe)
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
			cell := &table_models.Cell{Name: "Width"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.Width)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Height"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SVG"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.SVG)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Markdown_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Markdown"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Markdown"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Markdown]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Markdown_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Markdown"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Markdown == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Slider_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Slider"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Slider"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Slider]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Slider_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Slider"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Slider == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Split_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Split"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Split"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Split]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Split_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Split"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Split == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Svg_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Svg"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Style"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Svg"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Svg]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Svg_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Style"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Style)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Svg"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Svg == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Table_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Table"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Table"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Table]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Table_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Table"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Table == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Threejs_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Threejs"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Threejs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Threejs]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Threejs_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Threejs"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Threejs == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Title_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Title"
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

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Title]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Title_Stage(probe)
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

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Tone_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Tone"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Tone"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Tone]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Tone_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Tone"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Tone == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_Tree_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Tree"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Tree"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Tree]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Tree_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Tree"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Tree == structInstance {
					refNames = append(refNames, src.GetName())
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

func updateStageSetTable_View_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.View"
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
		col.Name = "ShowViewName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RootAsSplitAreas"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsSelectedView"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Direction"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsSecondaryView"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsSizeInPixel"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsWithCustomGutterSize"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GutterSize"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.View]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_View_Stage(probe)
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
			cell := &table_models.Cell{Name: "ShowViewName"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowViewName}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RootAsSplitAreas"}
			var names []string
			for _, elem := range structInstance.RootAsSplitAreas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsSelectedView"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsSelectedView}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Direction"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Direction)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsSecondaryView"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsSecondaryView}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsSizeInPixel"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsSizeInPixel}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsWithCustomGutterSize"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsWithCustomGutterSize}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GutterSize"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.GutterSize)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	probe.tableStage.StageBranch(table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Xlsx_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Xlsx"
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
		col.Name = "StackName"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.AsSplitArea) -> Xlsx"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Xlsx]()

	for _, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.Stage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.Stage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Xlsx_Stage(probe)
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
			cell := &table_models.Cell{Name: "StackName"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.StackName)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.AsSplitArea) -> Xlsx"}
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Xlsx == structInstance {
					refNames = append(refNames, src.GetName())
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

