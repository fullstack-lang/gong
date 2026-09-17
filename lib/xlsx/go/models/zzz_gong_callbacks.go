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
func (displayselection *DisplaySelection) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDisplaySelectionCreateCallback != nil {
		stage.OnAfterDisplaySelectionCreateCallback.OnAfterCreate(stage, displayselection)
	}
}

func (displayselection *DisplaySelection) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDisplaySelectionUpdateCallback != nil {
		var frontDisplaySelection *DisplaySelection
		if front != nil {
			frontDisplaySelection, _ = front.(*DisplaySelection)
		}
		stage.OnAfterDisplaySelectionUpdateCallback.OnAfterUpdate(stage, displayselection, frontDisplaySelection)
	}
}

func (displayselection *DisplaySelection) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDisplaySelectionDeleteCallback != nil {
		var frontDisplaySelection *DisplaySelection
		if front != nil {
			frontDisplaySelection, _ = front.(*DisplaySelection)
		}
		stage.OnAfterDisplaySelectionDeleteCallback.OnAfterDelete(stage, displayselection, frontDisplaySelection)
	}
}

func (xlcell *XLCell) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXLCellCreateCallback != nil {
		stage.OnAfterXLCellCreateCallback.OnAfterCreate(stage, xlcell)
	}
}

func (xlcell *XLCell) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLCellUpdateCallback != nil {
		var frontXLCell *XLCell
		if front != nil {
			frontXLCell, _ = front.(*XLCell)
		}
		stage.OnAfterXLCellUpdateCallback.OnAfterUpdate(stage, xlcell, frontXLCell)
	}
}

func (xlcell *XLCell) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLCellDeleteCallback != nil {
		var frontXLCell *XLCell
		if front != nil {
			frontXLCell, _ = front.(*XLCell)
		}
		stage.OnAfterXLCellDeleteCallback.OnAfterDelete(stage, xlcell, frontXLCell)
	}
}

func (xlfile *XLFile) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXLFileCreateCallback != nil {
		stage.OnAfterXLFileCreateCallback.OnAfterCreate(stage, xlfile)
	}
}

func (xlfile *XLFile) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLFileUpdateCallback != nil {
		var frontXLFile *XLFile
		if front != nil {
			frontXLFile, _ = front.(*XLFile)
		}
		stage.OnAfterXLFileUpdateCallback.OnAfterUpdate(stage, xlfile, frontXLFile)
	}
}

func (xlfile *XLFile) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLFileDeleteCallback != nil {
		var frontXLFile *XLFile
		if front != nil {
			frontXLFile, _ = front.(*XLFile)
		}
		stage.OnAfterXLFileDeleteCallback.OnAfterDelete(stage, xlfile, frontXLFile)
	}
}

func (xlrow *XLRow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXLRowCreateCallback != nil {
		stage.OnAfterXLRowCreateCallback.OnAfterCreate(stage, xlrow)
	}
}

func (xlrow *XLRow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLRowUpdateCallback != nil {
		var frontXLRow *XLRow
		if front != nil {
			frontXLRow, _ = front.(*XLRow)
		}
		stage.OnAfterXLRowUpdateCallback.OnAfterUpdate(stage, xlrow, frontXLRow)
	}
}

func (xlrow *XLRow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLRowDeleteCallback != nil {
		var frontXLRow *XLRow
		if front != nil {
			frontXLRow, _ = front.(*XLRow)
		}
		stage.OnAfterXLRowDeleteCallback.OnAfterDelete(stage, xlrow, frontXLRow)
	}
}

func (xlsheet *XLSheet) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXLSheetCreateCallback != nil {
		stage.OnAfterXLSheetCreateCallback.OnAfterCreate(stage, xlsheet)
	}
}

func (xlsheet *XLSheet) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLSheetUpdateCallback != nil {
		var frontXLSheet *XLSheet
		if front != nil {
			frontXLSheet, _ = front.(*XLSheet)
		}
		stage.OnAfterXLSheetUpdateCallback.OnAfterUpdate(stage, xlsheet, frontXLSheet)
	}
}

func (xlsheet *XLSheet) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXLSheetDeleteCallback != nil {
		var frontXLSheet *XLSheet
		if front != nil {
			frontXLSheet, _ = front.(*XLSheet)
		}
		stage.OnAfterXLSheetDeleteCallback.OnAfterDelete(stage, xlsheet, frontXLSheet)
	}
}

