// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct DisplaySelection
	// insertion point per field

	// Compute reverse map for named struct XLCell
	// insertion point per field

	// Compute reverse map for named struct XLFile
	// insertion point per field
	clear(stage.XLFile_Sheets_reverseMap)
	stage.XLFile_Sheets_reverseMap = make(map[*XLSheet]*XLFile)
	for xlfile := range stage.XLFiles {
		_ = xlfile
		for _, _xlsheet := range xlfile.Sheets {
			stage.XLFile_Sheets_reverseMap[_xlsheet] = xlfile
		}
	}

	// Compute reverse map for named struct XLRow
	// insertion point per field
	clear(stage.XLRow_Cells_reverseMap)
	stage.XLRow_Cells_reverseMap = make(map[*XLCell]*XLRow)
	for xlrow := range stage.XLRows {
		_ = xlrow
		for _, _xlcell := range xlrow.Cells {
			stage.XLRow_Cells_reverseMap[_xlcell] = xlrow
		}
	}

	// Compute reverse map for named struct XLSheet
	// insertion point per field
	clear(stage.XLSheet_Rows_reverseMap)
	stage.XLSheet_Rows_reverseMap = make(map[*XLRow]*XLSheet)
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		for _, _xlrow := range xlsheet.Rows {
			stage.XLSheet_Rows_reverseMap[_xlrow] = xlsheet
		}
	}
	clear(stage.XLSheet_SheetCells_reverseMap)
	stage.XLSheet_SheetCells_reverseMap = make(map[*XLCell]*XLSheet)
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		for _, _xlcell := range xlsheet.SheetCells {
			stage.XLSheet_SheetCells_reverseMap[_xlcell] = xlsheet
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.DisplaySelections {
		res = append(res, instance)
	}

	for instance := range stage.XLCells {
		res = append(res, instance)
	}

	for instance := range stage.XLFiles {
		res = append(res, instance)
	}

	for instance := range stage.XLRows {
		res = append(res, instance)
	}

	for instance := range stage.XLSheets {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (displayselection *DisplaySelection) GongCopy() GongstructIF {
	var newInstance DisplaySelection
	newInstance = *displayselection
	return &newInstance
}

func (xlcell *XLCell) GongCopy() GongstructIF {
	var newInstance XLCell
	newInstance = *xlcell
	return &newInstance
}

func (xlfile *XLFile) GongCopy() GongstructIF {
	var newInstance XLFile
	newInstance = *xlfile
	return &newInstance
}

func (xlrow *XLRow) GongCopy() GongstructIF {
	var newInstance XLRow
	newInstance = *xlrow
	return &newInstance
}

func (xlsheet *XLSheet) GongCopy() GongstructIF {
	var newInstance XLSheet
	newInstance = *xlsheet
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
