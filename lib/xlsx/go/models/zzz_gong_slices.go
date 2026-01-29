// generated code - do not edit
package models

import (
	"fmt"
	"sort"
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

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var displayselections_newInstances []*DisplaySelection
	var displayselections_deletedInstances []*DisplaySelection

	// parse all staged instances and check if they have a reference
	for displayselection := range stage.DisplaySelections {
		if ref, ok := stage.DisplaySelections_reference[displayselection]; !ok {
			displayselections_newInstances = append(displayselections_newInstances, displayselection)
			newInstancesSlice = append(newInstancesSlice, displayselection.GongMarshallIdentifier(stage))
			if stage.DisplaySelections_referenceOrder == nil {
				stage.DisplaySelections_referenceOrder = make(map[*DisplaySelection]uint)
			}
			stage.DisplaySelections_referenceOrder[displayselection] = stage.DisplaySelectionMap_Staged_Order[displayselection]
			newInstancesReverseSlice = append(newInstancesReverseSlice, displayselection.GongMarshallUnstaging(stage))
			delete(stage.DisplaySelections_referenceOrder, displayselection)
			fieldInitializers, pointersInitializations := displayselection.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DisplaySelectionMap_Staged_Order[ref] = stage.DisplaySelectionMap_Staged_Order[displayselection]
			diffs := displayselection.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, displayselection)
			delete(stage.DisplaySelectionMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", displayselection.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DisplaySelections_reference {
		if _, ok := stage.DisplaySelections[ref]; !ok {
			displayselections_deletedInstances = append(displayselections_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, xlcell.GongMarshallIdentifier(stage))
			if stage.XLCells_referenceOrder == nil {
				stage.XLCells_referenceOrder = make(map[*XLCell]uint)
			}
			stage.XLCells_referenceOrder[xlcell] = stage.XLCellMap_Staged_Order[xlcell]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlcell.GongMarshallUnstaging(stage))
			delete(stage.XLCells_referenceOrder, xlcell)
			fieldInitializers, pointersInitializations := xlcell.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLCellMap_Staged_Order[ref] = stage.XLCellMap_Staged_Order[xlcell]
			diffs := xlcell.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlcell)
			delete(stage.XLCellMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xlcell.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.XLCells_reference {
		if _, ok := stage.XLCells[ref]; !ok {
			xlcells_deletedInstances = append(xlcells_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, xlfile.GongMarshallIdentifier(stage))
			if stage.XLFiles_referenceOrder == nil {
				stage.XLFiles_referenceOrder = make(map[*XLFile]uint)
			}
			stage.XLFiles_referenceOrder[xlfile] = stage.XLFileMap_Staged_Order[xlfile]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlfile.GongMarshallUnstaging(stage))
			delete(stage.XLFiles_referenceOrder, xlfile)
			fieldInitializers, pointersInitializations := xlfile.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLFileMap_Staged_Order[ref] = stage.XLFileMap_Staged_Order[xlfile]
			diffs := xlfile.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlfile)
			delete(stage.XLFileMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xlfile.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.XLFiles_reference {
		if _, ok := stage.XLFiles[ref]; !ok {
			xlfiles_deletedInstances = append(xlfiles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, xlrow.GongMarshallIdentifier(stage))
			if stage.XLRows_referenceOrder == nil {
				stage.XLRows_referenceOrder = make(map[*XLRow]uint)
			}
			stage.XLRows_referenceOrder[xlrow] = stage.XLRowMap_Staged_Order[xlrow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlrow.GongMarshallUnstaging(stage))
			delete(stage.XLRows_referenceOrder, xlrow)
			fieldInitializers, pointersInitializations := xlrow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLRowMap_Staged_Order[ref] = stage.XLRowMap_Staged_Order[xlrow]
			diffs := xlrow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlrow)
			delete(stage.XLRowMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xlrow.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.XLRows_reference {
		if _, ok := stage.XLRows[ref]; !ok {
			xlrows_deletedInstances = append(xlrows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, xlsheet.GongMarshallIdentifier(stage))
			if stage.XLSheets_referenceOrder == nil {
				stage.XLSheets_referenceOrder = make(map[*XLSheet]uint)
			}
			stage.XLSheets_referenceOrder[xlsheet] = stage.XLSheetMap_Staged_Order[xlsheet]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlsheet.GongMarshallUnstaging(stage))
			delete(stage.XLSheets_referenceOrder, xlsheet)
			fieldInitializers, pointersInitializations := xlsheet.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLSheetMap_Staged_Order[ref] = stage.XLSheetMap_Staged_Order[xlsheet]
			diffs := xlsheet.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlsheet)
			delete(stage.XLSheetMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xlsheet.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.XLSheets_reference {
		if _, ok := stage.XLSheets[ref]; !ok {
			xlsheets_deletedInstances = append(xlsheets_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(xlsheets_newInstances)
	lenDeletedInstances += len(xlsheets_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.DisplaySelections_reference = make(map[*DisplaySelection]*DisplaySelection)
	stage.DisplaySelections_referenceOrder = make(map[*DisplaySelection]uint) // diff Unstage needs the reference order
	for instance := range stage.DisplaySelections {
		stage.DisplaySelections_reference[instance] = instance.GongCopy().(*DisplaySelection)
		stage.DisplaySelections_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.XLCells_reference = make(map[*XLCell]*XLCell)
	stage.XLCells_referenceOrder = make(map[*XLCell]uint) // diff Unstage needs the reference order
	for instance := range stage.XLCells {
		stage.XLCells_reference[instance] = instance.GongCopy().(*XLCell)
		stage.XLCells_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.XLFiles_reference = make(map[*XLFile]*XLFile)
	stage.XLFiles_referenceOrder = make(map[*XLFile]uint) // diff Unstage needs the reference order
	for instance := range stage.XLFiles {
		stage.XLFiles_reference[instance] = instance.GongCopy().(*XLFile)
		stage.XLFiles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.XLRows_reference = make(map[*XLRow]*XLRow)
	stage.XLRows_referenceOrder = make(map[*XLRow]uint) // diff Unstage needs the reference order
	for instance := range stage.XLRows {
		stage.XLRows_reference[instance] = instance.GongCopy().(*XLRow)
		stage.XLRows_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.XLSheets_reference = make(map[*XLSheet]*XLSheet)
	stage.XLSheets_referenceOrder = make(map[*XLSheet]uint) // diff Unstage needs the reference order
	for instance := range stage.XLSheets {
		stage.XLSheets_reference[instance] = instance.GongCopy().(*XLSheet)
		stage.XLSheets_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (displayselection *DisplaySelection) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DisplaySelectionMap_Staged_Order[displayselection]; ok {
		return order
	}
	return stage.DisplaySelections_referenceOrder[displayselection]
}

func (displayselection *DisplaySelection) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DisplaySelections_referenceOrder[displayselection]
}

func (xlcell *XLCell) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLCellMap_Staged_Order[xlcell]; ok {
		return order
	}
	return stage.XLCells_referenceOrder[xlcell]
}

func (xlcell *XLCell) GongGetReferenceOrder(stage *Stage) uint {
	return stage.XLCells_referenceOrder[xlcell]
}

func (xlfile *XLFile) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLFileMap_Staged_Order[xlfile]; ok {
		return order
	}
	return stage.XLFiles_referenceOrder[xlfile]
}

func (xlfile *XLFile) GongGetReferenceOrder(stage *Stage) uint {
	return stage.XLFiles_referenceOrder[xlfile]
}

func (xlrow *XLRow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLRowMap_Staged_Order[xlrow]; ok {
		return order
	}
	return stage.XLRows_referenceOrder[xlrow]
}

func (xlrow *XLRow) GongGetReferenceOrder(stage *Stage) uint {
	return stage.XLRows_referenceOrder[xlrow]
}

func (xlsheet *XLSheet) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLSheetMap_Staged_Order[xlsheet]; ok {
		return order
	}
	return stage.XLSheets_referenceOrder[xlsheet]
}

func (xlsheet *XLSheet) GongGetReferenceOrder(stage *Stage) uint {
	return stage.XLSheets_referenceOrder[xlsheet]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (displayselection *DisplaySelection) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayselection.GongGetGongstructName(), displayselection.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (displayselection *DisplaySelection) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayselection.GongGetGongstructName(), displayselection.GongGetReferenceOrder(stage))
}

func (xlcell *XLCell) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlcell.GongGetGongstructName(), xlcell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlcell *XLCell) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlcell.GongGetGongstructName(), xlcell.GongGetReferenceOrder(stage))
}

func (xlfile *XLFile) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlfile.GongGetGongstructName(), xlfile.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlfile *XLFile) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlfile.GongGetGongstructName(), xlfile.GongGetReferenceOrder(stage))
}

func (xlrow *XLRow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlrow.GongGetGongstructName(), xlrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlrow *XLRow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlrow.GongGetGongstructName(), xlrow.GongGetReferenceOrder(stage))
}

func (xlsheet *XLSheet) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsheet.GongGetGongstructName(), xlsheet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlsheet *XLSheet) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsheet.GongGetGongstructName(), xlsheet.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (displayselection *DisplaySelection) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplaySelection")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", displayselection.Name)
	return
}
func (xlcell *XLCell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLCell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlcell.Name)
	return
}
func (xlfile *XLFile) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLFile")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlfile.Name)
	return
}
func (xlrow *XLRow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLRow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlrow.Name)
	return
}
func (xlsheet *XLSheet) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLSheet")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlsheet.Name)
	return
}

// insertion point for unstaging
func (displayselection *DisplaySelection) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayselection.GongGetReferenceIdentifier(stage))
	return
}
func (xlcell *XLCell) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlcell.GongGetReferenceIdentifier(stage))
	return
}
func (xlfile *XLFile) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlfile.GongGetReferenceIdentifier(stage))
	return
}
func (xlrow *XLRow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlrow.GongGetReferenceIdentifier(stage))
	return
}
func (xlsheet *XLSheet) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsheet.GongGetReferenceIdentifier(stage))
	return
}
