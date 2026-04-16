// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Button
func (button *Button) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &button.SVGIcon) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Cell
func (cell *Cell) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &cell.CellString) || modified
	modified = GongCleanPointer(stage, &cell.CellFloat64) || modified
	modified = GongCleanPointer(stage, &cell.CellInt) || modified
	modified = GongCleanPointer(stage, &cell.CellBool) || modified
	modified = GongCleanPointer(stage, &cell.CellIcon) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CellBoolean
func (cellboolean *CellBoolean) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellFloat64
func (cellfloat64 *CellFloat64) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellIcon
func (cellicon *CellIcon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellInt
func (cellint *CellInt) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellString
func (cellstring *CellString) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DisplayedColumn
func (displayedcolumn *DisplayedColumn) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Row
func (row *Row) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &row.Cells) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SVGIcon
func (svgicon *SVGIcon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &table.DisplayedColumns) || modified
	modified = GongCleanSlice(stage, &table.Rows) || modified
	modified = GongCleanSlice(stage, &table.RowsSelectedForBulkDelete) || modified
	modified = GongCleanSlice(stage, &table.Buttons) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
