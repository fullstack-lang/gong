// generated code - do not edit
package probe

import (
	"fmt"
	"sort"
	"strings"

	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	maticons "github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
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
	case "A":
		updateStageSetTable_A_Stage(probe)
	case "B":
		updateStageSetTable_B_Stage(probe)
	case "X":
		updateStageSetTable_X_XStage(probe)
	case "Y":
		updateStageSetTable_Y_YStage(probe)
	}
}

func updateStageSetTable_A_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "A"
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
		col.Name = "NumberField"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "B"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Bs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Foo"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Bar"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "Zorgh"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	// Sort instances by name
	instances := make([]*models.A, 0, len(probe.stageSet.Stage.As))
	for inst := range probe.stageSet.Stage.As {
		instances = append(instances, inst)
	}
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].GetName() < instances[j].GetName()
	})

	for idx, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: idx}
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
				updateStageSetTable_A_Stage(probe)
				probe.ux_tree()
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
			cell := &table_models.Cell{Name: "NumberField"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.NumberField)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "B"}
			val := ""
			if structInstance.B != nil {
				val = structInstance.B.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Bs"}
			var names []string
			for _, elem := range structInstance.Bs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			cell.CellString = &table_models.CellString{Value: strings.Join(names, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "X"}
			val := ""
			if structInstance.X != nil {
				val = structInstance.X.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Foo"}
			cell.CellInt = &table_models.CellInt{Value: int(structInstance.Foo)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Bar"}
			cell.CellFloat64 = &table_models.CellFloat64{Value: float64(structInstance.Bar)}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "Zorgh"}
			cell.CellString = &table_models.CellString{Value: fmt.Sprintf("%v", structInstance.Zorgh)}
			row.Cells = append(row.Cells, cell)
		}

		table.Rows = append(table.Rows, row)
	}

	table_models.StageBranch(probe.tableStage, table)
	probe.tableStage.Commit()
}

func updateStageSetTable_B_Stage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "B"
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
		col.Name = "(models.A) -> B"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.A) -> Bs"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	// Sort instances by name
	instances := make([]*models.B, 0, len(probe.stageSet.Stage.Bs))
	for inst := range probe.stageSet.Stage.Bs {
		instances = append(instances, inst)
	}
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].GetName() < instances[j].GetName()
	})

	for idx, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: idx}
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
				updateStageSetTable_B_Stage(probe)
				probe.ux_tree()
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
			cell := &table_models.Cell{Name: "(models.A) -> B"}
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				if src.B == structInstance {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			cell.CellString = &table_models.CellString{Value: strings.Join(refNames, ", ")}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.A) -> Bs"}
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				for _, target := range src.Bs {
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

func updateStageSetTable_X_XStage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "X"
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
		col.Name = "Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}
	{
		col := new(table_models.DisplayedColumn)
		col.Name = "(models.A) -> X"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	// Sort instances by name
	instances := make([]*x.X, 0, len(probe.stageSet.XStage.Xs))
	for inst := range probe.stageSet.XStage.Xs {
		instances = append(instances, inst)
	}
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].GetName() < instances[j].GetName()
	})

	for idx, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: idx}
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
				_captured.UnstageVoid(probe.stageSet.XStage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_X_XStage(probe)
				probe.ux_tree()
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
			cell := &table_models.Cell{Name: "Y"}
			val := ""
			if structInstance.Y != nil {
				val = structInstance.Y.GetName()
			}
			cell.CellString = &table_models.CellString{Value: val}
			row.Cells = append(row.Cells, cell)
		}

		{
			cell := &table_models.Cell{Name: "(models.A) -> X"}
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				if src.X == structInstance {
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

func updateStageSetTable_Y_YStage(probe *StageSetProbe) {
	probe.tableStage.Reset()

	table := new(table_models.Table)
	table.Name = "Y"
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
		col.Name = "(x.X) -> Y"
		table.DisplayedColumns = append(table.DisplayedColumns, col)
	}

	// Sort instances by name
	instances := make([]*y.Y, 0, len(probe.stageSet.YStage.Ys))
	for inst := range probe.stageSet.YStage.Ys {
		instances = append(instances, inst)
	}
	sort.Slice(instances, func(i, j int) bool {
		return instances[i].GetName() < instances[j].GetName()
	})

	for idx, structInstance := range instances {
		row := new(table_models.Row)
		row.Name = structInstance.GetName()

		_captured := structInstance
		row.Impl = &tableRowUpdater{
			onClick: func() {
				StageSetFillUpFormFromGongstruct(_captured, probe)
			},
		}

		cellID := &table_models.Cell{Name: "ID"}
		cellID.CellInt = &table_models.CellInt{Value: idx}
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
				_captured.UnstageVoid(probe.stageSet.YStage)
				probe.stageSet.Clean()
				probe.stageSet.Commit()
				updateStageSetTable_Y_YStage(probe)
				probe.ux_tree()
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
			cell := &table_models.Cell{Name: "(x.X) -> Y"}
			var refNames []string
			for src := range probe.stageSet.XStage.Xs {
				if src.Y == structInstance {
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

