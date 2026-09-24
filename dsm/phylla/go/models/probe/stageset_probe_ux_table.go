// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/stool"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/music"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models/abstract/clock"
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
	case "models.Angle0Shape":
		updateStageSetTable_Angle0Shape_Stage(probe)
	case "models.BottomCurvePlane1Shape":
		updateStageSetTable_BottomCurvePlane1Shape_Stage(probe)
	case "models.BottomCurvePlane2Shape":
		updateStageSetTable_BottomCurvePlane2Shape_Stage(probe)
	case "models.Circumference3DShape":
		updateStageSetTable_Circumference3DShape_Stage(probe)
	case "models.Clock2DDiagram":
		updateStageSetTable_Clock2DDiagram_Stage(probe)
	case "models.Clock3DDiagram":
		updateStageSetTable_Clock3DDiagram_Stage(probe)
	case "models.CutLine3DShape":
		updateStageSetTable_CutLine3DShape_Stage(probe)
	case "models.Leaves3DShape":
		updateStageSetTable_Leaves3DShape_Stage(probe)
	case "models.Library":
		updateStageSetTable_Library_Stage(probe)
	case "models.OriginalPoints3DShape":
		updateStageSetTable_OriginalPoints3DShape_Stage(probe)
	case "models.ParastichyMCurves3DShape":
		updateStageSetTable_ParastichyMCurves3DShape_Stage(probe)
	case "models.ParastichyNCurves3DShape":
		updateStageSetTable_ParastichyNCurves3DShape_Stage(probe)
	case "models.Plant2DDiagram":
		updateStageSetTable_Plant2DDiagram_Stage(probe)
	case "models.Plant3DDiagram":
		updateStageSetTable_Plant3DDiagram_Stage(probe)
	case "models.PlantAbstract":
		updateStageSetTable_PlantAbstract_Stage(probe)
	case "models.Rendered3DShape":
		updateStageSetTable_Rendered3DShape_Stage(probe)
	case "models.SampledPoints3DShape":
		updateStageSetTable_SampledPoints3DShape_Stage(probe)
	case "models.StackOfRotatedVaseTrapezeRingsShape":
		updateStageSetTable_StackOfRotatedVaseTrapezeRingsShape_Stage(probe)
	case "models.StackOfVaseTrapezeRingsShape":
		updateStageSetTable_StackOfVaseTrapezeRingsShape_Stage(probe)
	case "models.StemCylinder3DShape":
		updateStageSetTable_StemCylinder3DShape_Stage(probe)
	case "models.Stool2DDiagram":
		updateStageSetTable_Stool2DDiagram_Stage(probe)
	case "models.Stool3DDiagram":
		updateStageSetTable_Stool3DDiagram_Stage(probe)
	case "models.TopCurvePlane1Shape":
		updateStageSetTable_TopCurvePlane1Shape_Stage(probe)
	case "models.TopCurvePlane2Shape":
		updateStageSetTable_TopCurvePlane2Shape_Stage(probe)
	case "models.TubeVase3DDiagram":
		updateStageSetTable_TubeVase3DDiagram_Stage(probe)
	case "models.TubeVaseAbstract":
		updateStageSetTable_TubeVaseAbstract_Stage(probe)
	case "models.Vase2DDiagram":
		updateStageSetTable_Vase2DDiagram_Stage(probe)
	case "models.VaseTrapezeRingShape":
		updateStageSetTable_VaseTrapezeRingShape_Stage(probe)
	case "stool.StoolAbstract":
		updateStageSetTable_StoolAbstract_StoolStage(probe)
	case "music.MusicAbstract":
		updateStageSetTable_MusicAbstract_MusicStage(probe)
	case "clock.ClockAbstract":
		updateStageSetTable_ClockAbstract_ClockStage(probe)
	}
}

func updateStageSetTable_Angle0Shape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Angle0Shape"
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
		col.Name = "(models.TubeVase3DDiagram) -> Angle0Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Angle0Shape]()

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
				updateStageSetTable_Angle0Shape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> Angle0Shape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.Angle0Shape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_BottomCurvePlane1Shape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.BottomCurvePlane1Shape"
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
		col.Name = "(models.TubeVase3DDiagram) -> BottomCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.BottomCurvePlane1Shape]()

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
				updateStageSetTable_BottomCurvePlane1Shape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> BottomCurvePlane1Shape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.BottomCurvePlane1Shape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_BottomCurvePlane2Shape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.BottomCurvePlane2Shape"
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
		col.Name = "(models.TubeVase3DDiagram) -> BottomCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.BottomCurvePlane2Shape]()

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
				updateStageSetTable_BottomCurvePlane2Shape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> BottomCurvePlane2Shape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.BottomCurvePlane2Shape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Circumference3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Circumference3DShape"
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
		col.Name = "(models.Plant3DDiagram) -> Circumference3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Circumference3DShape]()

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
				updateStageSetTable_Circumference3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> Circumference3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Circumference3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Clock2DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Clock2DDiagram"
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
		col.Name = "Zoom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenAxesShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Clock2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Clock2DDiagram]()

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
				updateStageSetTable_Clock2DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "Zoom"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Zoom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenAxesShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenAxesShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Clock2DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Clock2DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Clock3DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Clock3DDiagram"
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
		col.Name = "IsHiddenClockTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ClockTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTorus3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Torus3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Clock3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Clock3DDiagram]()

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
				updateStageSetTable_Clock3DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "IsHiddenClockTopCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenClockTopCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ClockTopCurveShape"}
			val := ""
			if structInstance.ClockTopCurveShape != nil {
				val = structInstance.ClockTopCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTorus3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTorus3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Torus3DShape"}
			val := ""
			if structInstance.Torus3DShape != nil {
				val = structInstance.Torus3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SampledPoints3DShape"}
			val := ""
			if structInstance.SampledPoints3DShape != nil {
				val = structInstance.SampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTiledFloor3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTiledFloor3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TiledFloor3DShape"}
			val := ""
			if structInstance.TiledFloor3DShape != nil {
				val = structInstance.TiledFloor3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Rendered3DShape"}
			val := ""
			if structInstance.Rendered3DShape != nil {
				val = structInstance.Rendered3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Clock3DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Clock3DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_CutLine3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.CutLine3DShape"
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
		col.Name = "(models.Plant3DDiagram) -> CutLine3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.CutLine3DShape]()

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
				updateStageSetTable_CutLine3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> CutLine3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.CutLine3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Leaves3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Leaves3DShape"
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
		col.Name = "(models.Plant3DDiagram) -> Leaves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Leaves3DShape]()

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
				updateStageSetTable_Leaves3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> Leaves3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Leaves3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Library_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Library"
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
		col.Name = "Plants"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SubLibraries"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbPixPerCharacter"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "LogoSVGFile"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsRootLibrary"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> SubLibraries"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Library]()

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
				updateStageSetTable_Library_Stage(probe)
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
			cell := &table_models.Cell{Name: "Plants"}
			var names []string
			for _, elem := range structInstance.Plants {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SubLibraries"}
			var names []string
			for _, elem := range structInstance.SubLibraries {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbPixPerCharacter"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.NbPixPerCharacter)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "LogoSVGFile"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.LogoSVGFile)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsRootLibrary"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsRootLibrary}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> SubLibraries"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.SubLibraries {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_OriginalPoints3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.OriginalPoints3DShape"
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
		col.Name = "(models.TubeVase3DDiagram) -> OriginalPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.OriginalPoints3DShape]()

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
				updateStageSetTable_OriginalPoints3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> OriginalPoints3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.OriginalPoints3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ParastichyMCurves3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ParastichyMCurves3DShape"
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
		col.Name = "(models.Plant3DDiagram) -> ParastichyMCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ParastichyMCurves3DShape]()

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
				updateStageSetTable_ParastichyMCurves3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> ParastichyMCurves3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.ParastichyMCurves3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ParastichyNCurves3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.ParastichyNCurves3DShape"
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
		col.Name = "(models.Plant3DDiagram) -> ParastichyNCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.ParastichyNCurves3DShape]()

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
				updateStageSetTable_ParastichyNCurves3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> ParastichyNCurves3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.ParastichyNCurves3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Plant2DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Plant2DDiagram"
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
		col.Name = "OriginX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OriginY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Zoom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsRhombusNodesExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsArcNodesExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenAxesShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenReferenceRhombus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPlantCircumferenceShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenGridPathShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRhombusGridShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenExplanationTextShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedReferenceRhombus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedPlantCircumferenceShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedGridPathShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedRhombusGridShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenGrowthPathRhombusGridShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenGrowthVectorShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPerpendicularVectorGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBaseVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenArcNormalVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenMidArcVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEndArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfGrowthCurve2DByGrowthVector"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Plant2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Plant2DDiagram]()

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
				updateStageSetTable_Plant2DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "OriginX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OriginX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OriginY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OriginY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Zoom"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Zoom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsRhombusNodesExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsRhombusNodesExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsArcNodesExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsArcNodesExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenAxesShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenAxesShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenReferenceRhombus"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenReferenceRhombus}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPlantCircumferenceShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPlantCircumferenceShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenGridPathShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenGridPathShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRhombusGridShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRhombusGridShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenExplanationTextShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenExplanationTextShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedReferenceRhombus"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedReferenceRhombus}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedPlantCircumferenceShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedPlantCircumferenceShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedGridPathShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedGridPathShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedRhombusGridShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedRhombusGridShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenGrowthPathRhombusGridShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenGrowthPathRhombusGridShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenGrowthVectorShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenGrowthVectorShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPerpendicularVectorGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPerpendicularVectorGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBaseVectorShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBaseVectorShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenArcNormalVectorShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenArcNormalVectorShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStartArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStartArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenMidArcVectorShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenMidArcVectorShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEndArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEndArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenGrowthCurve2D"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenGrowthCurve2D}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfGrowthCurve2DByGrowthVector"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfGrowthCurve2DByGrowthVector}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Plant2DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Plant2DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Plant3DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Plant3DDiagram"
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
		col.Name = "IsHiddenStemCylinder3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StemCylinder3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenParastichyNCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ParastichyNCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenParastichyMCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ParastichyMCurves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenCutLine3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CutLine3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenCircumference3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Circumference3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenLeaves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Leaves3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Plant3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Plant3DDiagram]()

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
				updateStageSetTable_Plant3DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "IsHiddenStemCylinder3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStemCylinder3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StemCylinder3DShape"}
			val := ""
			if structInstance.StemCylinder3DShape != nil {
				val = structInstance.StemCylinder3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenParastichyNCurves3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenParastichyNCurves3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ParastichyNCurves3DShape"}
			val := ""
			if structInstance.ParastichyNCurves3DShape != nil {
				val = structInstance.ParastichyNCurves3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenParastichyMCurves3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenParastichyMCurves3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ParastichyMCurves3DShape"}
			val := ""
			if structInstance.ParastichyMCurves3DShape != nil {
				val = structInstance.ParastichyMCurves3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenCutLine3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenCutLine3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CutLine3DShape"}
			val := ""
			if structInstance.CutLine3DShape != nil {
				val = structInstance.CutLine3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenCircumference3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenCircumference3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Circumference3DShape"}
			val := ""
			if structInstance.Circumference3DShape != nil {
				val = structInstance.Circumference3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTiledFloor3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTiledFloor3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TiledFloor3DShape"}
			val := ""
			if structInstance.TiledFloor3DShape != nil {
				val = structInstance.TiledFloor3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenLeaves3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenLeaves3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Leaves3DShape"}
			val := ""
			if structInstance.Leaves3DShape != nil {
				val = structInstance.Leaves3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Rendered3DShape"}
			val := ""
			if structInstance.Rendered3DShape != nil {
				val = structInstance.Rendered3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Plant3DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Plant3DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_PlantAbstract_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.PlantAbstract"
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
		col.Name = "N"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "M"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RhombusInsideAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RhombusSideLength"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PlantType"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TubeVaseAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StoolAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ClockAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MusicAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "CurrentView"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsSelected"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsPlant2DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Plant2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsPlant3DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Plant3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsVase2DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Vase2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsTubeVase3DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TubeVase3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsStool2DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Stool2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsStool3DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Stool3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsClock2DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Clock2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsClock3DDiagramsNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Clock3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "AxesShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RhombusStuff"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GrowthVectorShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PerpendicularVectorGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BaseVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ArcNormalVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MidArcVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfGrowthCurve2DByGrowthVector"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Library) -> Plants"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.PlantAbstract]()

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
				updateStageSetTable_PlantAbstract_Stage(probe)
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
			cell := &table_models.Cell{Name: "N"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.N)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "M"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.M)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackHeight"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.StackHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RhombusInsideAngle"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RhombusInsideAngle)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RhombusSideLength"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RhombusSideLength)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PlantType"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.PlantType)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TubeVaseAbstract"}
			val := ""
			if structInstance.TubeVaseAbstract != nil {
				val = structInstance.TubeVaseAbstract.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StoolAbstract"}
			val := ""
			if structInstance.StoolAbstract != nil {
				val = structInstance.StoolAbstract.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ClockAbstract"}
			val := ""
			if structInstance.ClockAbstract != nil {
				val = structInstance.ClockAbstract.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MusicAbstract"}
			val := ""
			if structInstance.MusicAbstract != nil {
				val = structInstance.MusicAbstract.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "CurrentView"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.CurrentView)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsSelected"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsSelected}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsPlant2DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsPlant2DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Plant2DDiagrams"}
			var names []string
			for _, elem := range structInstance.Plant2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsPlant3DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsPlant3DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Plant3DDiagrams"}
			var names []string
			for _, elem := range structInstance.Plant3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsVase2DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsVase2DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Vase2DDiagrams"}
			var names []string
			for _, elem := range structInstance.Vase2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsTubeVase3DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsTubeVase3DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TubeVase3DDiagrams"}
			var names []string
			for _, elem := range structInstance.TubeVase3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsStool2DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsStool2DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Stool2DDiagrams"}
			var names []string
			for _, elem := range structInstance.Stool2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsStool3DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsStool3DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Stool3DDiagrams"}
			var names []string
			for _, elem := range structInstance.Stool3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsClock2DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsClock2DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Clock2DDiagrams"}
			var names []string
			for _, elem := range structInstance.Clock2DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsClock3DDiagramsNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsClock3DDiagramsNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Clock3DDiagrams"}
			var names []string
			for _, elem := range structInstance.Clock3DDiagrams {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "AxesShape"}
			val := ""
			if structInstance.AxesShape != nil {
				val = structInstance.AxesShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RhombusStuff"}
			val := ""
			if structInstance.RhombusStuff != nil {
				val = structInstance.RhombusStuff.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GrowthVectorShape"}
			val := ""
			if structInstance.GrowthVectorShape != nil {
				val = structInstance.GrowthVectorShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PerpendicularVectorGrid"}
			val := ""
			if structInstance.PerpendicularVectorGrid != nil {
				val = structInstance.PerpendicularVectorGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BaseVectorShapeGrid"}
			val := ""
			if structInstance.BaseVectorShapeGrid != nil {
				val = structInstance.BaseVectorShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ArcNormalVectorShapeGrid"}
			val := ""
			if structInstance.ArcNormalVectorShapeGrid != nil {
				val = structInstance.ArcNormalVectorShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartArcShapeGrid"}
			val := ""
			if structInstance.StartArcShapeGrid != nil {
				val = structInstance.StartArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MidArcVectorShapeGrid"}
			val := ""
			if structInstance.MidArcVectorShapeGrid != nil {
				val = structInstance.MidArcVectorShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndArcShapeGrid"}
			val := ""
			if structInstance.EndArcShapeGrid != nil {
				val = structInstance.EndArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GrowthCurve2D"}
			val := ""
			if structInstance.GrowthCurve2D != nil {
				val = structInstance.GrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfGrowthCurve2DByGrowthVector"}
			val := ""
			if structInstance.StackOfGrowthCurve2DByGrowthVector != nil {
				val = structInstance.StackOfGrowthCurve2DByGrowthVector.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Library) -> Plants"}
			var refNames []string
			for src := range probe.stageSet.Stage.Librarys {
				for _, target := range src.Plants {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Rendered3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Rendered3DShape"
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
		col.Name = "ViewX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ViewY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ViewZ"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TargetX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TargetY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TargetZ"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Fov"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Clock3DDiagram) -> Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Plant3DDiagram) -> Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Stool3DDiagram) -> Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TubeVase3DDiagram) -> Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Rendered3DShape]()

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
				updateStageSetTable_Rendered3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "ViewX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ViewX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ViewY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ViewY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ViewZ"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ViewZ)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TargetX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TargetX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TargetY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TargetY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TargetZ"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.TargetZ)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Fov"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Fov)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Clock3DDiagram) -> Rendered3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Clock3DDiagrams {
				if src.Rendered3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> Rendered3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.Rendered3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Stool3DDiagram) -> Rendered3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Stool3DDiagrams {
				if src.Rendered3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> Rendered3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.Rendered3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_SampledPoints3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.SampledPoints3DShape"
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
		col.Name = "(models.Clock3DDiagram) -> SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Stool3DDiagram) -> SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.TubeVase3DDiagram) -> SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.SampledPoints3DShape]()

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
				updateStageSetTable_SampledPoints3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.Clock3DDiagram) -> SampledPoints3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Clock3DDiagrams {
				if src.SampledPoints3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Stool3DDiagram) -> SampledPoints3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Stool3DDiagrams {
				if src.SampledPoints3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> SampledPoints3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.SampledPoints3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_StackOfRotatedVaseTrapezeRingsShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.StackOfRotatedVaseTrapezeRingsShape"
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
		col.Name = "(models.TubeVase3DDiagram) -> StackOfRotatedVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.StackOfRotatedVaseTrapezeRingsShape]()

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
				updateStageSetTable_StackOfRotatedVaseTrapezeRingsShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> StackOfRotatedVaseTrapezeRingsShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.StackOfRotatedVaseTrapezeRingsShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_StackOfVaseTrapezeRingsShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.StackOfVaseTrapezeRingsShape"
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
		col.Name = "(models.TubeVase3DDiagram) -> StackOfVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.StackOfVaseTrapezeRingsShape]()

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
				updateStageSetTable_StackOfVaseTrapezeRingsShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> StackOfVaseTrapezeRingsShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.StackOfVaseTrapezeRingsShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_StemCylinder3DShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.StemCylinder3DShape"
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
		col.Name = "Transparency"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.Plant3DDiagram) -> StemCylinder3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.StemCylinder3DShape]()

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
				updateStageSetTable_StemCylinder3DShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "Transparency"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Transparency)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.Plant3DDiagram) -> StemCylinder3DShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.Plant3DDiagrams {
				if src.StemCylinder3DShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Stool2DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Stool2DDiagram"
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
		col.Name = "Zoom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenAxesShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Stool2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Stool2DDiagram]()

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
				updateStageSetTable_Stool2DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "Zoom"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Zoom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenAxesShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenAxesShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Stool2DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Stool2DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Stool3DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Stool3DDiagram"
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
		col.Name = "IsHiddenSeatTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SeatTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedSeatTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotatedSeatTopCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedSeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotatedSeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTorus3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Torus3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotatedSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEyeSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EyeSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEyeCornersSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EyeCornersSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEye3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Eye3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEyeSeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EyeSeatBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEyeStoolBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EyeStoolBottomCurveShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSeat3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Seat3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEyeVolume3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EyeVolume3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSeatAndLegs3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SeatAndLegs3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenRotatedSeatAndLegs3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotatedSeatAndLegs3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Stool3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Stool3DDiagram]()

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
				updateStageSetTable_Stool3DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "IsHiddenSeatTopCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSeatTopCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SeatTopCurveShape"}
			val := ""
			if structInstance.SeatTopCurveShape != nil {
				val = structInstance.SeatTopCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedSeatTopCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedSeatTopCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotatedSeatTopCurveShape"}
			val := ""
			if structInstance.RotatedSeatTopCurveShape != nil {
				val = structInstance.RotatedSeatTopCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSeatBottomCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSeatBottomCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SeatBottomCurveShape"}
			val := ""
			if structInstance.SeatBottomCurveShape != nil {
				val = structInstance.SeatBottomCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedSeatBottomCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedSeatBottomCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotatedSeatBottomCurveShape"}
			val := ""
			if structInstance.RotatedSeatBottomCurveShape != nil {
				val = structInstance.RotatedSeatBottomCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTorus3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTorus3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Torus3DShape"}
			val := ""
			if structInstance.Torus3DShape != nil {
				val = structInstance.Torus3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedTorusShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedTorusShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotatedTorusShape"}
			val := ""
			if structInstance.RotatedTorusShape != nil {
				val = structInstance.RotatedTorusShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SampledPoints3DShape"}
			val := ""
			if structInstance.SampledPoints3DShape != nil {
				val = structInstance.SampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotatedSampledPoints3DShape"}
			val := ""
			if structInstance.RotatedSampledPoints3DShape != nil {
				val = structInstance.RotatedSampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEyeSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEyeSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EyeSampledPoints3DShape"}
			val := ""
			if structInstance.EyeSampledPoints3DShape != nil {
				val = structInstance.EyeSampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEyeCornersSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEyeCornersSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EyeCornersSampledPoints3DShape"}
			val := ""
			if structInstance.EyeCornersSampledPoints3DShape != nil {
				val = structInstance.EyeCornersSampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEye3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEye3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Eye3DShape"}
			val := ""
			if structInstance.Eye3DShape != nil {
				val = structInstance.Eye3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEyeSeatBottomCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEyeSeatBottomCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EyeSeatBottomCurveShape"}
			val := ""
			if structInstance.EyeSeatBottomCurveShape != nil {
				val = structInstance.EyeSeatBottomCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEyeStoolBottomCurveShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEyeStoolBottomCurveShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EyeStoolBottomCurveShape"}
			val := ""
			if structInstance.EyeStoolBottomCurveShape != nil {
				val = structInstance.EyeStoolBottomCurveShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSeat3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSeat3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Seat3DShape"}
			val := ""
			if structInstance.Seat3DShape != nil {
				val = structInstance.Seat3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEyeVolume3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEyeVolume3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EyeVolume3DShape"}
			val := ""
			if structInstance.EyeVolume3DShape != nil {
				val = structInstance.EyeVolume3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSeatAndLegs3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSeatAndLegs3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SeatAndLegs3DShape"}
			val := ""
			if structInstance.SeatAndLegs3DShape != nil {
				val = structInstance.SeatAndLegs3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenRotatedSeatAndLegs3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenRotatedSeatAndLegs3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotatedSeatAndLegs3DShape"}
			val := ""
			if structInstance.RotatedSeatAndLegs3DShape != nil {
				val = structInstance.RotatedSeatAndLegs3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTiledFloor3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTiledFloor3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TiledFloor3DShape"}
			val := ""
			if structInstance.TiledFloor3DShape != nil {
				val = structInstance.TiledFloor3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Rendered3DShape"}
			val := ""
			if structInstance.Rendered3DShape != nil {
				val = structInstance.Rendered3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Stool3DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Stool3DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TopCurvePlane1Shape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TopCurvePlane1Shape"
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
		col.Name = "(models.TubeVase3DDiagram) -> TopCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TopCurvePlane1Shape]()

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
				updateStageSetTable_TopCurvePlane1Shape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> TopCurvePlane1Shape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.TopCurvePlane1Shape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TopCurvePlane2Shape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TopCurvePlane2Shape"
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
		col.Name = "(models.TubeVase3DDiagram) -> TopCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TopCurvePlane2Shape]()

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
				updateStageSetTable_TopCurvePlane2Shape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> TopCurvePlane2Shape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.TopCurvePlane2Shape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TubeVase3DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TubeVase3DDiagram"
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
		col.Name = "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTorusStackShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenVerticalTorusStackShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPartiallyRotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfPartiallyRotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPointsAndLines3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenKeyHole3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenKey3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenVolumeKey3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTorusEdge3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenSampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenOriginalPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenAngle0Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBottomCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBottomCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenVaseTrapezeRingShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfRotatedVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Rendered3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TorusStackShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VerticalTorusStackShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PartiallyRotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfPartiallyRotatedTorusShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PointsAndLines3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "SampledPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OriginalPoints3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Angle0Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "KeyHole3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Key3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VolumeKey3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TorusEdge3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TiledFloor3DShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BottomCurvePlane1Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BottomCurvePlane2Shape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "VaseTrapezeRingShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfRotatedVaseTrapezeRingsShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> TubeVase3DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TubeVase3DDiagram]()

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
				updateStageSetTable_TubeVase3DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTorusStackShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTorusStackShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenVerticalTorusStackShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenVerticalTorusStackShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPartiallyRotatedTorusShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPartiallyRotatedTorusShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfPartiallyRotatedTorusShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfPartiallyRotatedTorusShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPointsAndLines3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPointsAndLines3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenKeyHole3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenKeyHole3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenKey3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenKey3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenVolumeKey3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenVolumeKey3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTorusEdge3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTorusEdge3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenSampledPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenSampledPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenOriginalPoints3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenOriginalPoints3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenAngle0Shape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenAngle0Shape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTiledFloor3DShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTiledFloor3DShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopCurvePlane1Shape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopCurvePlane1Shape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBottomCurvePlane1Shape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBottomCurvePlane1Shape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopCurvePlane2Shape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopCurvePlane2Shape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBottomCurvePlane2Shape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBottomCurvePlane2Shape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenVaseTrapezeRingShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenVaseTrapezeRingShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfVaseTrapezeRingsShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfVaseTrapezeRingsShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfRotatedVaseTrapezeRingsShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfRotatedVaseTrapezeRingsShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Rendered3DShape"}
			val := ""
			if structInstance.Rendered3DShape != nil {
				val = structInstance.Rendered3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TorusStackShape"}
			val := ""
			if structInstance.TorusStackShape != nil {
				val = structInstance.TorusStackShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VerticalTorusStackShape"}
			val := ""
			if structInstance.VerticalTorusStackShape != nil {
				val = structInstance.VerticalTorusStackShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PartiallyRotatedTorusShape"}
			val := ""
			if structInstance.PartiallyRotatedTorusShape != nil {
				val = structInstance.PartiallyRotatedTorusShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfPartiallyRotatedTorusShape"}
			val := ""
			if structInstance.StackOfPartiallyRotatedTorusShape != nil {
				val = structInstance.StackOfPartiallyRotatedTorusShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PointsAndLines3DShape"}
			val := ""
			if structInstance.PointsAndLines3DShape != nil {
				val = structInstance.PointsAndLines3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "SampledPoints3DShape"}
			val := ""
			if structInstance.SampledPoints3DShape != nil {
				val = structInstance.SampledPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OriginalPoints3DShape"}
			val := ""
			if structInstance.OriginalPoints3DShape != nil {
				val = structInstance.OriginalPoints3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Angle0Shape"}
			val := ""
			if structInstance.Angle0Shape != nil {
				val = structInstance.Angle0Shape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "KeyHole3DShape"}
			val := ""
			if structInstance.KeyHole3DShape != nil {
				val = structInstance.KeyHole3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Key3DShape"}
			val := ""
			if structInstance.Key3DShape != nil {
				val = structInstance.Key3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VolumeKey3DShape"}
			val := ""
			if structInstance.VolumeKey3DShape != nil {
				val = structInstance.VolumeKey3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TorusEdge3DShape"}
			val := ""
			if structInstance.TorusEdge3DShape != nil {
				val = structInstance.TorusEdge3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TiledFloor3DShape"}
			val := ""
			if structInstance.TiledFloor3DShape != nil {
				val = structInstance.TiledFloor3DShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopCurvePlane1Shape"}
			val := ""
			if structInstance.TopCurvePlane1Shape != nil {
				val = structInstance.TopCurvePlane1Shape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BottomCurvePlane1Shape"}
			val := ""
			if structInstance.BottomCurvePlane1Shape != nil {
				val = structInstance.BottomCurvePlane1Shape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopCurvePlane2Shape"}
			val := ""
			if structInstance.TopCurvePlane2Shape != nil {
				val = structInstance.TopCurvePlane2Shape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BottomCurvePlane2Shape"}
			val := ""
			if structInstance.BottomCurvePlane2Shape != nil {
				val = structInstance.BottomCurvePlane2Shape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "VaseTrapezeRingShape"}
			val := ""
			if structInstance.VaseTrapezeRingShape != nil {
				val = structInstance.VaseTrapezeRingShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfVaseTrapezeRingsShape"}
			val := ""
			if structInstance.StackOfVaseTrapezeRingsShape != nil {
				val = structInstance.StackOfVaseTrapezeRingsShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfRotatedVaseTrapezeRingsShape"}
			val := ""
			if structInstance.StackOfRotatedVaseTrapezeRingsShape != nil {
				val = structInstance.StackOfRotatedVaseTrapezeRingsShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> TubeVase3DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.TubeVase3DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_TubeVaseAbstract_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.TubeVaseAbstract"
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
		col.Name = "Z_Ribbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RibbonVerticalScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Plane1Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Plane2Height"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProjectionAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeVerticalThickness"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeRadialThickness"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeCuttedStackFloorHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeRotatedTorusSeparation"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RotationRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RadialRepetitions"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Transparency"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HasAlternatingRingColors"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeTrajectoryOffsetX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeTrajectoryOffsetY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbStepP1P2"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ChosenStep"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeHorizontalRingsHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OffsetKeyX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OffsetKeyY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "HeightKey"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "WidthKey"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeKeySize"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "MovieNbFrames"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PerpendicularVectorGridHalfway"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopEndArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShiftedBottomTopStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopMidArcVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StartHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopStartHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "EndHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopEndHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfRotatedGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopStackOfRotatedGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "TopStackOfGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StackOfRotatedGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "GrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShiftedRightGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShiftedLeftGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PartiallyGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShiftedLeftPartiallyGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PartiallyGrowthCurve2DTrajectory"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PartiallyGrowthCurve2DTrajectoryP1P2"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PxShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ChosenP1P2PairShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "KeyHoleShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> TubeVaseAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.TubeVaseAbstract]()

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
				updateStageSetTable_TubeVaseAbstract_Stage(probe)
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
			cell := &table_models.Cell{Name: "Z_Ribbon"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Z_Ribbon)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RibbonVerticalScale"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RibbonVerticalScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Plane1Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Plane1Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Plane2Height"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Plane2Height)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProjectionAngle"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ProjectionAngle)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeVerticalThickness"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeVerticalThickness)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeRadialThickness"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeRadialThickness)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeCuttedStackFloorHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeCuttedStackFloorHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeRotatedTorusSeparation"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeRotatedTorusSeparation)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RotationRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RotationRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RadialRepetitions"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.RadialRepetitions)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Transparency"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Transparency)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HasAlternatingRingColors"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.HasAlternatingRingColors}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeTrajectoryOffsetX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeTrajectoryOffsetX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeTrajectoryOffsetY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeTrajectoryOffsetY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbStepP1P2"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NbStepP1P2)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ChosenStep"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.ChosenStep)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeHorizontalRingsHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeHorizontalRingsHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OffsetKeyX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OffsetKeyX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OffsetKeyY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OffsetKeyY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "HeightKey"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.HeightKey)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "WidthKey"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.WidthKey)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeKeySize"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeKeySize)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "MovieNbFrames"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.MovieNbFrames)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PerpendicularVectorGridHalfway"}
			val := ""
			if structInstance.PerpendicularVectorGridHalfway != nil {
				val = structInstance.PerpendicularVectorGridHalfway.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopStartArcShapeGrid"}
			val := ""
			if structInstance.TopStartArcShapeGrid != nil {
				val = structInstance.TopStartArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopEndArcShapeGrid"}
			val := ""
			if structInstance.TopEndArcShapeGrid != nil {
				val = structInstance.TopEndArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShiftedBottomTopStartArcShapeGrid"}
			val := ""
			if structInstance.ShiftedBottomTopStartArcShapeGrid != nil {
				val = structInstance.ShiftedBottomTopStartArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopMidArcVectorShapeGrid"}
			val := ""
			if structInstance.TopMidArcVectorShapeGrid != nil {
				val = structInstance.TopMidArcVectorShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StartHalfwayArcShapeGrid"}
			val := ""
			if structInstance.StartHalfwayArcShapeGrid != nil {
				val = structInstance.StartHalfwayArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopStartHalfwayArcShapeGrid"}
			val := ""
			if structInstance.TopStartHalfwayArcShapeGrid != nil {
				val = structInstance.TopStartHalfwayArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "EndHalfwayArcShapeGrid"}
			val := ""
			if structInstance.EndHalfwayArcShapeGrid != nil {
				val = structInstance.EndHalfwayArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopEndHalfwayArcShapeGrid"}
			val := ""
			if structInstance.TopEndHalfwayArcShapeGrid != nil {
				val = structInstance.TopEndHalfwayArcShapeGrid.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfRotatedGrowthCurve2D"}
			val := ""
			if structInstance.StackOfRotatedGrowthCurve2D != nil {
				val = structInstance.StackOfRotatedGrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopStackOfRotatedGrowthCurve2D"}
			val := ""
			if structInstance.TopStackOfRotatedGrowthCurve2D != nil {
				val = structInstance.TopStackOfRotatedGrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopGrowthCurve2D"}
			val := ""
			if structInstance.TopGrowthCurve2D != nil {
				val = structInstance.TopGrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfGrowthCurve2D"}
			val := ""
			if structInstance.StackOfGrowthCurve2D != nil {
				val = structInstance.StackOfGrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "TopStackOfGrowthCurve2D"}
			val := ""
			if structInstance.TopStackOfGrowthCurve2D != nil {
				val = structInstance.TopStackOfGrowthCurve2D.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfGrowthCurve2DRibbon"}
			val := ""
			if structInstance.StackOfGrowthCurve2DRibbon != nil {
				val = structInstance.StackOfGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StackOfRotatedGrowthCurve2DRibbon"}
			val := ""
			if structInstance.StackOfRotatedGrowthCurve2DRibbon != nil {
				val = structInstance.StackOfRotatedGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "GrowthCurve2DRibbon"}
			val := ""
			if structInstance.GrowthCurve2DRibbon != nil {
				val = structInstance.GrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShiftedRightGrowthCurve2DRibbon"}
			val := ""
			if structInstance.ShiftedRightGrowthCurve2DRibbon != nil {
				val = structInstance.ShiftedRightGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShiftedLeftGrowthCurve2DRibbon"}
			val := ""
			if structInstance.ShiftedLeftGrowthCurve2DRibbon != nil {
				val = structInstance.ShiftedLeftGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PartiallyGrowthCurve2DRibbon"}
			val := ""
			if structInstance.PartiallyGrowthCurve2DRibbon != nil {
				val = structInstance.PartiallyGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShiftedLeftPartiallyGrowthCurve2DRibbon"}
			val := ""
			if structInstance.ShiftedLeftPartiallyGrowthCurve2DRibbon != nil {
				val = structInstance.ShiftedLeftPartiallyGrowthCurve2DRibbon.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PartiallyGrowthCurve2DTrajectory"}
			val := ""
			if structInstance.PartiallyGrowthCurve2DTrajectory != nil {
				val = structInstance.PartiallyGrowthCurve2DTrajectory.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PartiallyGrowthCurve2DTrajectoryP1P2"}
			val := ""
			if structInstance.PartiallyGrowthCurve2DTrajectoryP1P2 != nil {
				val = structInstance.PartiallyGrowthCurve2DTrajectoryP1P2.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PxShape"}
			val := ""
			if structInstance.PxShape != nil {
				val = structInstance.PxShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ChosenP1P2PairShape"}
			val := ""
			if structInstance.ChosenP1P2PairShape != nil {
				val = structInstance.ChosenP1P2PairShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "KeyHoleShape"}
			val := ""
			if structInstance.KeyHoleShape != nil {
				val = structInstance.KeyHoleShape.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> TubeVaseAbstract"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.TubeVaseAbstract == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_Vase2DDiagram_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.Vase2DDiagram"
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
		col.Name = "Zoom"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsVaseArcNodesExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsVaseClampingNodesExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenAxesShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBottomStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBottomEndArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenBottomStackOfGrowthCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedLeftStackOfGrowthCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedLeftStackOfNormalVector"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPerpendicularVectorGridHalfway"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedBottomTopStartArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopMidArcVectorShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStartHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopStartHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenEndHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopEndHalfwayArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopEndArcShapeGrid"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfGrowthCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopStackOfGrowthCurve"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenTopStackOfGrowthCurve2D"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedRightGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedLeftGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenStackOfRotatedGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPartiallyGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPartiallyGrowthCurve2DTrajectory"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenPxShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenChosenP1P2PairShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsHiddenKeyHoleShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ComputedPrefix"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> Vase2DDiagrams"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.Vase2DDiagram]()

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
				updateStageSetTable_Vase2DDiagram_Stage(probe)
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
			cell := &table_models.Cell{Name: "Zoom"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Zoom)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsVaseArcNodesExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsVaseArcNodesExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsVaseClampingNodesExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsVaseClampingNodesExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenAxesShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenAxesShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBottomStartArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBottomStartArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBottomEndArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBottomEndArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenBottomStackOfGrowthCurve"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenBottomStackOfGrowthCurve}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedLeftStackOfGrowthCurve"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedLeftStackOfGrowthCurve}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedLeftStackOfNormalVector"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedLeftStackOfNormalVector}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPerpendicularVectorGridHalfway"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPerpendicularVectorGridHalfway}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopStartArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopStartArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedBottomTopStartArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedBottomTopStartArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopMidArcVectorShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopMidArcVectorShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStartHalfwayArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStartHalfwayArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopStartHalfwayArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopStartHalfwayArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenEndHalfwayArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenEndHalfwayArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopEndHalfwayArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopEndHalfwayArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopEndArcShapeGrid"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopEndArcShapeGrid}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfGrowthCurve"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfGrowthCurve}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopStackOfGrowthCurve"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopStackOfGrowthCurve}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopGrowthCurve2D"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopGrowthCurve2D}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfGrowthCurve2D"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfGrowthCurve2D}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenTopStackOfGrowthCurve2D"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenTopStackOfGrowthCurve2D}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedRightGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedRightGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedLeftGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedLeftGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenStackOfRotatedGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenStackOfRotatedGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPartiallyGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPartiallyGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPartiallyGrowthCurve2DTrajectory"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPartiallyGrowthCurve2DTrajectory}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenPxShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenPxShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenChosenP1P2PairShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenChosenP1P2PairShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsHiddenKeyHoleShape"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsHiddenKeyHoleShape}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ComputedPrefix"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.ComputedPrefix)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> Vase2DDiagrams"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				for _, target := range src.Vase2DDiagrams {
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

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_VaseTrapezeRingShape_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "models.VaseTrapezeRingShape"
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
		col.Name = "(models.TubeVase3DDiagram) -> VaseTrapezeRingShape"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.Stage.GetInstancesByOrder[*models.VaseTrapezeRingShape]()

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
				updateStageSetTable_VaseTrapezeRingShape_Stage(probe)
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
			cell := &table_models.Cell{Name: "(models.TubeVase3DDiagram) -> VaseTrapezeRingShape"}
			var refNames []string
			for src := range probe.stageSet.Stage.TubeVase3DDiagrams {
				if src.VaseTrapezeRingShape == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_StoolAbstract_StoolStage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "stool.StoolAbstract"
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
		col.Name = "RadialRepetitions"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Transparency"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeTubeDiameter"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeHeight3DTorus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "StoolTorusVerticalScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeSeatThickness"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProjectionAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeEyeSeparationCriteria"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeEyeCornerControlVectorStrength"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> StoolAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.StoolStage.GetInstancesByOrder[*stool.StoolAbstract]()

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
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.StoolStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.StoolStage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_StoolAbstract_StoolStage(probe)
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
			cell := &table_models.Cell{Name: "RadialRepetitions"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.RadialRepetitions)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Transparency"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Transparency)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeTubeDiameter"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeTubeDiameter)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeHeight3DTorus"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeHeight3DTorus)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "StoolTorusVerticalScale"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.StoolTorusVerticalScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeSeatThickness"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeSeatThickness)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProjectionAngle"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ProjectionAngle)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeEyeSeparationCriteria"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeEyeSeparationCriteria)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeEyeCornerControlVectorStrength"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeEyeCornerControlVectorStrength)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> StoolAbstract"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.StoolAbstract == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_MusicAbstract_MusicStage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "music.MusicAbstract"
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
		col.Name = "IsChecked"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PitchHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbOfBeatsInTheme"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BeatsPerSecond"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FirstVoiceShiftX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "FirstVoiceShiftY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "PitchDifference"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Level"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ActualBeatsTemporalShift"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsMinor"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ThemeBinaryEncoding"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "BezierControlLengthRatio"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbPitchLines"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "NbBeatLines"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OriginX"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "OriginY"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ScoreScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowFirstVoice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowFirstVoiceShiftRight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowSecondVoice"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowSecondVoiceShiftRight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowFirstVoiceNotes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowFirstVoiceNotesShiftRight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowSecondVoiceNotes"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ShowSecondVoiceNotesShiftRight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "IsComposerNodeExpanded"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> MusicAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.MusicStage.GetInstancesByOrder[*music.MusicAbstract]()

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
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.MusicStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.MusicStage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_MusicAbstract_MusicStage(probe)
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
			cell := &table_models.Cell{Name: "IsChecked"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsChecked}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PitchHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.PitchHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbOfBeatsInTheme"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NbOfBeatsInTheme)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BeatsPerSecond"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.BeatsPerSecond)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FirstVoiceShiftX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.FirstVoiceShiftX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "FirstVoiceShiftY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.FirstVoiceShiftY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "PitchDifference"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.PitchDifference)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Level"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Level)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ActualBeatsTemporalShift"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.ActualBeatsTemporalShift)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsMinor"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsMinor}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ThemeBinaryEncoding"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.ThemeBinaryEncoding)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "BezierControlLengthRatio"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.BezierControlLengthRatio)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbPitchLines"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NbPitchLines)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "NbBeatLines"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NbBeatLines)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OriginX"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OriginX)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "OriginY"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.OriginY)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ScoreScale"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ScoreScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowFirstVoice"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowFirstVoice}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowFirstVoiceShiftRight"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowFirstVoiceShiftRight}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowSecondVoice"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowSecondVoice}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowSecondVoiceShiftRight"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowSecondVoiceShiftRight}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowFirstVoiceNotes"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowFirstVoiceNotes}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowFirstVoiceNotesShiftRight"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowFirstVoiceNotesShiftRight}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowSecondVoiceNotes"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowSecondVoiceNotes}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ShowSecondVoiceNotesShiftRight"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.ShowSecondVoiceNotesShiftRight}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "IsComposerNodeExpanded"}
			cell.CellBool = &table_models.CellBoolean{Value: structInstance.IsComposerNodeExpanded}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> MusicAbstract"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.MusicAbstract == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_ClockAbstract_ClockStage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "clock.ClockAbstract"
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
		col.Name = "RadialRepetitions"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Transparency"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeTubeDiameter"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeHeight3DTorus"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ClockTorusVerticalScale"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "RelativeHeight"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "ProjectionAngle"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.PlantAbstract) -> ClockAbstract"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	instances := probe.stageSet.ClockStage.GetInstancesByOrder[*clock.ClockAbstract]()

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
		cellID.CellInt = &table_models.CellInt{Value: int(probe.stageSet.ClockStage.GetOrder(structInstance))}
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
				_captured.UnstageVoid(probe.stageSet.ClockStage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_ClockAbstract_ClockStage(probe)
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
			cell := &table_models.Cell{Name: "RadialRepetitions"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.RadialRepetitions)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Transparency"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Transparency)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeTubeDiameter"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeTubeDiameter)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeHeight3DTorus"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeHeight3DTorus)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ClockTorusVerticalScale"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ClockTorusVerticalScale)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "RelativeHeight"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.RelativeHeight)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "ProjectionAngle"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.ProjectionAngle)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.PlantAbstract) -> ClockAbstract"}
			var refNames []string
			for src := range probe.stageSet.Stage.PlantAbstracts {
				if src.ClockAbstract == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

