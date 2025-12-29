// generated code - do not edit
package models

import (
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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
	stage.XLFile_Sheets_reverseMap = make(map[*XLSheet]*XLFile)
	for xlfile := range stage.XLFiles {
		_ = xlfile
		for _, _xlsheet := range xlfile.Sheets {
			stage.XLFile_Sheets_reverseMap[_xlsheet] = xlfile
		}
	}

	// Compute reverse map for named struct XLRow
	// insertion point per field
	stage.XLRow_Cells_reverseMap = make(map[*XLCell]*XLRow)
	for xlrow := range stage.XLRows {
		_ = xlrow
		for _, _xlcell := range xlrow.Cells {
			stage.XLRow_Cells_reverseMap[_xlcell] = xlrow
		}
	}

	// Compute reverse map for named struct XLSheet
	// insertion point per field
	stage.XLSheet_Rows_reverseMap = make(map[*XLRow]*XLSheet)
	for xlsheet := range stage.XLSheets {
		_ = xlsheet
		for _, _xlrow := range xlsheet.Rows {
			stage.XLSheet_Rows_reverseMap[_xlrow] = xlsheet
		}
	}
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
	newInstance := *displayselection
	return &newInstance
}

func (xlcell *XLCell) GongCopy() GongstructIF {
	newInstance := *xlcell
	return &newInstance
}

func (xlfile *XLFile) GongCopy() GongstructIF {
	newInstance := *xlfile
	return &newInstance
}

func (xlrow *XLRow) GongCopy() GongstructIF {
	newInstance := *xlrow
	return &newInstance
}

func (xlsheet *XLSheet) GongCopy() GongstructIF {
	newInstance := *xlsheet
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var displayselections_newInstances []*DisplaySelection
	var displayselections_deletedInstances []*DisplaySelection

	// parse all staged instances and check if they have a reference
	for displayselection := range stage.DisplaySelections {
		if ref, ok := stage.DisplaySelections_reference[displayselection]; !ok {
			displayselections_newInstances = append(displayselections_newInstances, displayselection)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of DisplaySelection "+displayselection.Name,
				)
			}
		} else {
			diffs := displayselection.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of DisplaySelection \""+displayselection.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for displayselection := range stage.DisplaySelections_reference {
		if _, ok := stage.DisplaySelections[displayselection]; !ok {
			displayselections_deletedInstances = append(displayselections_deletedInstances, displayselection)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of DisplaySelection "+displayselection.Name,
				)
			}
		}
	}

	lenNewInstances += len(displayselections_newInstances)
	lenDeletedInstances += len(displayselections_deletedInstances)
	var xlcells_newInstances []*XLCell
	var xlcells_deletedInstances []*XLCell

	// parse all staged instances and check if they have a reference
	for xlcell := range stage.XLCells {
		if ref, ok := stage.XLCells_reference[xlcell]; !ok {
			xlcells_newInstances = append(xlcells_newInstances, xlcell)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of XLCell "+xlcell.Name,
				)
			}
		} else {
			diffs := xlcell.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of XLCell \""+xlcell.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for xlcell := range stage.XLCells_reference {
		if _, ok := stage.XLCells[xlcell]; !ok {
			xlcells_deletedInstances = append(xlcells_deletedInstances, xlcell)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of XLCell "+xlcell.Name,
				)
			}
		}
	}

	lenNewInstances += len(xlcells_newInstances)
	lenDeletedInstances += len(xlcells_deletedInstances)
	var xlfiles_newInstances []*XLFile
	var xlfiles_deletedInstances []*XLFile

	// parse all staged instances and check if they have a reference
	for xlfile := range stage.XLFiles {
		if ref, ok := stage.XLFiles_reference[xlfile]; !ok {
			xlfiles_newInstances = append(xlfiles_newInstances, xlfile)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of XLFile "+xlfile.Name,
				)
			}
		} else {
			diffs := xlfile.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of XLFile \""+xlfile.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for xlfile := range stage.XLFiles_reference {
		if _, ok := stage.XLFiles[xlfile]; !ok {
			xlfiles_deletedInstances = append(xlfiles_deletedInstances, xlfile)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of XLFile "+xlfile.Name,
				)
			}
		}
	}

	lenNewInstances += len(xlfiles_newInstances)
	lenDeletedInstances += len(xlfiles_deletedInstances)
	var xlrows_newInstances []*XLRow
	var xlrows_deletedInstances []*XLRow

	// parse all staged instances and check if they have a reference
	for xlrow := range stage.XLRows {
		if ref, ok := stage.XLRows_reference[xlrow]; !ok {
			xlrows_newInstances = append(xlrows_newInstances, xlrow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of XLRow "+xlrow.Name,
				)
			}
		} else {
			diffs := xlrow.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of XLRow \""+xlrow.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for xlrow := range stage.XLRows_reference {
		if _, ok := stage.XLRows[xlrow]; !ok {
			xlrows_deletedInstances = append(xlrows_deletedInstances, xlrow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of XLRow "+xlrow.Name,
				)
			}
		}
	}

	lenNewInstances += len(xlrows_newInstances)
	lenDeletedInstances += len(xlrows_deletedInstances)
	var xlsheets_newInstances []*XLSheet
	var xlsheets_deletedInstances []*XLSheet

	// parse all staged instances and check if they have a reference
	for xlsheet := range stage.XLSheets {
		if ref, ok := stage.XLSheets_reference[xlsheet]; !ok {
			xlsheets_newInstances = append(xlsheets_newInstances, xlsheet)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of XLSheet "+xlsheet.Name,
				)
			}
		} else {
			diffs := xlsheet.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of XLSheet \""+xlsheet.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for xlsheet := range stage.XLSheets_reference {
		if _, ok := stage.XLSheets[xlsheet]; !ok {
			xlsheets_deletedInstances = append(xlsheets_deletedInstances, xlsheet)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of XLSheet "+xlsheet.Name,
				)
			}
		}
	}

	lenNewInstances += len(xlsheets_newInstances)
	lenDeletedInstances += len(xlsheets_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.DisplaySelections_reference = make(map[*DisplaySelection]*DisplaySelection)
	for instance := range stage.DisplaySelections {
		stage.DisplaySelections_reference[instance] = instance.GongCopy().(*DisplaySelection)
	}

	stage.XLCells_reference = make(map[*XLCell]*XLCell)
	for instance := range stage.XLCells {
		stage.XLCells_reference[instance] = instance.GongCopy().(*XLCell)
	}

	stage.XLFiles_reference = make(map[*XLFile]*XLFile)
	for instance := range stage.XLFiles {
		stage.XLFiles_reference[instance] = instance.GongCopy().(*XLFile)
	}

	stage.XLRows_reference = make(map[*XLRow]*XLRow)
	for instance := range stage.XLRows {
		stage.XLRows_reference[instance] = instance.GongCopy().(*XLRow)
	}

	stage.XLSheets_reference = make(map[*XLSheet]*XLSheet)
	for instance := range stage.XLSheets {
		stage.XLSheets_reference[instance] = instance.GongCopy().(*XLSheet)
	}

}
