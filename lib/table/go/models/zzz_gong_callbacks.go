// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (button *Button) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterButtonCreateCallback != nil {
		stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, button)
	}
}

func (button *Button) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonUpdateCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, button, frontButton)
	}
}

func (button *Button) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonDeleteCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, button, frontButton)
	}
}

func (cell *Cell) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellCreateCallback != nil {
		stage.OnAfterCellCreateCallback.OnAfterCreate(stage, cell)
	}
}

func (cell *Cell) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellUpdateCallback != nil {
		var frontCell *Cell
		if front != nil {
			frontCell, _ = front.(*Cell)
		}
		stage.OnAfterCellUpdateCallback.OnAfterUpdate(stage, cell, frontCell)
	}
}

func (cell *Cell) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellDeleteCallback != nil {
		var frontCell *Cell
		if front != nil {
			frontCell, _ = front.(*Cell)
		}
		stage.OnAfterCellDeleteCallback.OnAfterDelete(stage, cell, frontCell)
	}
}

func (cellboolean *CellBoolean) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellBooleanCreateCallback != nil {
		stage.OnAfterCellBooleanCreateCallback.OnAfterCreate(stage, cellboolean)
	}
}

func (cellboolean *CellBoolean) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellBooleanUpdateCallback != nil {
		var frontCellBoolean *CellBoolean
		if front != nil {
			frontCellBoolean, _ = front.(*CellBoolean)
		}
		stage.OnAfterCellBooleanUpdateCallback.OnAfterUpdate(stage, cellboolean, frontCellBoolean)
	}
}

func (cellboolean *CellBoolean) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellBooleanDeleteCallback != nil {
		var frontCellBoolean *CellBoolean
		if front != nil {
			frontCellBoolean, _ = front.(*CellBoolean)
		}
		stage.OnAfterCellBooleanDeleteCallback.OnAfterDelete(stage, cellboolean, frontCellBoolean)
	}
}

func (cellfloat64 *CellFloat64) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellFloat64CreateCallback != nil {
		stage.OnAfterCellFloat64CreateCallback.OnAfterCreate(stage, cellfloat64)
	}
}

func (cellfloat64 *CellFloat64) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellFloat64UpdateCallback != nil {
		var frontCellFloat64 *CellFloat64
		if front != nil {
			frontCellFloat64, _ = front.(*CellFloat64)
		}
		stage.OnAfterCellFloat64UpdateCallback.OnAfterUpdate(stage, cellfloat64, frontCellFloat64)
	}
}

func (cellfloat64 *CellFloat64) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellFloat64DeleteCallback != nil {
		var frontCellFloat64 *CellFloat64
		if front != nil {
			frontCellFloat64, _ = front.(*CellFloat64)
		}
		stage.OnAfterCellFloat64DeleteCallback.OnAfterDelete(stage, cellfloat64, frontCellFloat64)
	}
}

func (cellicon *CellIcon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellIconCreateCallback != nil {
		stage.OnAfterCellIconCreateCallback.OnAfterCreate(stage, cellicon)
	}
}

func (cellicon *CellIcon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellIconUpdateCallback != nil {
		var frontCellIcon *CellIcon
		if front != nil {
			frontCellIcon, _ = front.(*CellIcon)
		}
		stage.OnAfterCellIconUpdateCallback.OnAfterUpdate(stage, cellicon, frontCellIcon)
	}
}

func (cellicon *CellIcon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellIconDeleteCallback != nil {
		var frontCellIcon *CellIcon
		if front != nil {
			frontCellIcon, _ = front.(*CellIcon)
		}
		stage.OnAfterCellIconDeleteCallback.OnAfterDelete(stage, cellicon, frontCellIcon)
	}
}

func (cellint *CellInt) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellIntCreateCallback != nil {
		stage.OnAfterCellIntCreateCallback.OnAfterCreate(stage, cellint)
	}
}

func (cellint *CellInt) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellIntUpdateCallback != nil {
		var frontCellInt *CellInt
		if front != nil {
			frontCellInt, _ = front.(*CellInt)
		}
		stage.OnAfterCellIntUpdateCallback.OnAfterUpdate(stage, cellint, frontCellInt)
	}
}

func (cellint *CellInt) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellIntDeleteCallback != nil {
		var frontCellInt *CellInt
		if front != nil {
			frontCellInt, _ = front.(*CellInt)
		}
		stage.OnAfterCellIntDeleteCallback.OnAfterDelete(stage, cellint, frontCellInt)
	}
}

func (cellstring *CellString) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCellStringCreateCallback != nil {
		stage.OnAfterCellStringCreateCallback.OnAfterCreate(stage, cellstring)
	}
}

func (cellstring *CellString) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellStringUpdateCallback != nil {
		var frontCellString *CellString
		if front != nil {
			frontCellString, _ = front.(*CellString)
		}
		stage.OnAfterCellStringUpdateCallback.OnAfterUpdate(stage, cellstring, frontCellString)
	}
}

func (cellstring *CellString) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCellStringDeleteCallback != nil {
		var frontCellString *CellString
		if front != nil {
			frontCellString, _ = front.(*CellString)
		}
		stage.OnAfterCellStringDeleteCallback.OnAfterDelete(stage, cellstring, frontCellString)
	}
}

func (displayedcolumn *DisplayedColumn) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDisplayedColumnCreateCallback != nil {
		stage.OnAfterDisplayedColumnCreateCallback.OnAfterCreate(stage, displayedcolumn)
	}
}

func (displayedcolumn *DisplayedColumn) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDisplayedColumnUpdateCallback != nil {
		var frontDisplayedColumn *DisplayedColumn
		if front != nil {
			frontDisplayedColumn, _ = front.(*DisplayedColumn)
		}
		stage.OnAfterDisplayedColumnUpdateCallback.OnAfterUpdate(stage, displayedcolumn, frontDisplayedColumn)
	}
}

func (displayedcolumn *DisplayedColumn) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDisplayedColumnDeleteCallback != nil {
		var frontDisplayedColumn *DisplayedColumn
		if front != nil {
			frontDisplayedColumn, _ = front.(*DisplayedColumn)
		}
		stage.OnAfterDisplayedColumnDeleteCallback.OnAfterDelete(stage, displayedcolumn, frontDisplayedColumn)
	}
}

func (row *Row) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRowCreateCallback != nil {
		stage.OnAfterRowCreateCallback.OnAfterCreate(stage, row)
	}
}

func (row *Row) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRowUpdateCallback != nil {
		var frontRow *Row
		if front != nil {
			frontRow, _ = front.(*Row)
		}
		stage.OnAfterRowUpdateCallback.OnAfterUpdate(stage, row, frontRow)
	}
}

func (row *Row) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRowDeleteCallback != nil {
		var frontRow *Row
		if front != nil {
			frontRow, _ = front.(*Row)
		}
		stage.OnAfterRowDeleteCallback.OnAfterDelete(stage, row, frontRow)
	}
}

func (svgicon *SVGIcon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSVGIconCreateCallback != nil {
		stage.OnAfterSVGIconCreateCallback.OnAfterCreate(stage, svgicon)
	}
}

func (svgicon *SVGIcon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGIconUpdateCallback != nil {
		var frontSVGIcon *SVGIcon
		if front != nil {
			frontSVGIcon, _ = front.(*SVGIcon)
		}
		stage.OnAfterSVGIconUpdateCallback.OnAfterUpdate(stage, svgicon, frontSVGIcon)
	}
}

func (svgicon *SVGIcon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGIconDeleteCallback != nil {
		var frontSVGIcon *SVGIcon
		if front != nil {
			frontSVGIcon, _ = front.(*SVGIcon)
		}
		stage.OnAfterSVGIconDeleteCallback.OnAfterDelete(stage, svgicon, frontSVGIcon)
	}
}

func (table *Table) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableCreateCallback != nil {
		stage.OnAfterTableCreateCallback.OnAfterCreate(stage, table)
	}
}

func (table *Table) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableUpdateCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, table, frontTable)
	}
}

func (table *Table) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableDeleteCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, table, frontTable)
	}
}

