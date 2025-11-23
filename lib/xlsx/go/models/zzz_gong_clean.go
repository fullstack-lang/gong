// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}
    
	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up DisplaySelection
	for displayselection := range stage.DisplaySelections {
		_ = displayselection
		// insertion point per field
		// insertion point per field
		displayselection.XLFile = GongCleanPointer(stage, displayselection.XLFile)
		displayselection.XLSheet = GongCleanPointer(stage, displayselection.XLSheet)
	}

	// clean up XLCell
	for xlcell := range stage.XLCells {
		_ = xlcell
		// insertion point per field
		// insertion point per field
	}

	// clean up XLFile
	for xlfile := range stage.XLFiles {
		_ = xlfile
		// insertion point per field
		xlfile.Sheets = GongCleanSlice(stage, xlfile.Sheets)
		// insertion point per field
	}

	// clean up XLRow
	for xlrow := range stage.XLRows {
		_ = xlrow
		// insertion point per field
		xlrow.Cells = GongCleanSlice(stage, xlrow.Cells)
		// insertion point per field
	}

	// clean up XLSheet
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		// insertion point per field
		xlsheet.Rows = GongCleanSlice(stage, xlsheet.Rows)
		xlsheet.SheetCells = GongCleanSlice(stage, xlsheet.SheetCells)
		// insertion point per field
	}

}
