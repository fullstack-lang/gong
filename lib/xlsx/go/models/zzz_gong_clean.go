// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct DisplaySelection
	for displayselection := range stage.DisplaySelections {
		_ = displayselection
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, displayselection.XLFile) {
			displayselection.XLFile = nil
		}
		if !IsStaged(stage, displayselection.XLSheet) {
			displayselection.XLSheet = nil
		}
	}

	// Compute reverse map for named struct XLCell
	for xlcell := range stage.XLCells {
		_ = xlcell
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct XLFile
	for xlfile := range stage.XLFiles {
		_ = xlfile
		// insertion point per field
		var _Sheets []*XLSheet
		for _, _xlsheet := range xlfile.Sheets {
			if IsStaged(stage, _xlsheet) {
			 	_Sheets = append(_Sheets, _xlsheet)
			}
		}
		xlfile.Sheets = _Sheets
		// insertion point per field
	}

	// Compute reverse map for named struct XLRow
	for xlrow := range stage.XLRows {
		_ = xlrow
		// insertion point per field
		var _Cells []*XLCell
		for _, _xlcell := range xlrow.Cells {
			if IsStaged(stage, _xlcell) {
			 	_Cells = append(_Cells, _xlcell)
			}
		}
		xlrow.Cells = _Cells
		// insertion point per field
	}

	// Compute reverse map for named struct XLSheet
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		// insertion point per field
		var _Rows []*XLRow
		for _, _xlrow := range xlsheet.Rows {
			if IsStaged(stage, _xlrow) {
			 	_Rows = append(_Rows, _xlrow)
			}
		}
		xlsheet.Rows = _Rows
		var _SheetCells []*XLCell
		for _, _xlcell := range xlsheet.SheetCells {
			if IsStaged(stage, _xlcell) {
			 	_SheetCells = append(_SheetCells, _xlcell)
			}
		}
		xlsheet.SheetCells = _SheetCells
		// insertion point per field
	}

}
