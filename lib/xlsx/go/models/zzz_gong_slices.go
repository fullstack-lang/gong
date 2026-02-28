// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

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

	// end of insertion point per named struct
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
	newInstance := new(DisplaySelection)
	displayselection.CopyBasicFields(newInstance)
	return newInstance
}

func (xlcell *XLCell) GongCopy() GongstructIF {
	newInstance := new(XLCell)
	xlcell.CopyBasicFields(newInstance)
	return newInstance
}

func (xlfile *XLFile) GongCopy() GongstructIF {
	newInstance := new(XLFile)
	xlfile.CopyBasicFields(newInstance)
	return newInstance
}

func (xlrow *XLRow) GongCopy() GongstructIF {
	newInstance := new(XLRow)
	xlrow.CopyBasicFields(newInstance)
	return newInstance
}

func (xlsheet *XLSheet) GongCopy() GongstructIF {
	newInstance := new(XLSheet)
	xlsheet.CopyBasicFields(newInstance)
	return newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
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
			stage.DisplaySelections_referenceOrder[displayselection] = stage.DisplaySelection_stagedOrder[displayselection]
			newInstancesReverseSlice = append(newInstancesReverseSlice, displayselection.GongMarshallUnstaging(stage))
			// delete(stage.DisplaySelections_referenceOrder, displayselection)
			fieldInitializers, pointersInitializations := displayselection.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DisplaySelection_stagedOrder[ref] = stage.DisplaySelection_stagedOrder[displayselection]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := displayselection.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, displayselection)
			// delete(stage.DisplaySelection_stagedOrder, ref)
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
	for _, ref := range stage.DisplaySelections_reference {
		instance := stage.DisplaySelections_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DisplaySelections[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.XLCells_referenceOrder[xlcell] = stage.XLCell_stagedOrder[xlcell]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlcell.GongMarshallUnstaging(stage))
			// delete(stage.XLCells_referenceOrder, xlcell)
			fieldInitializers, pointersInitializations := xlcell.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLCell_stagedOrder[ref] = stage.XLCell_stagedOrder[xlcell]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := xlcell.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlcell)
			// delete(stage.XLCell_stagedOrder, ref)
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
	for _, ref := range stage.XLCells_reference {
		instance := stage.XLCells_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.XLCells[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.XLFiles_referenceOrder[xlfile] = stage.XLFile_stagedOrder[xlfile]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlfile.GongMarshallUnstaging(stage))
			// delete(stage.XLFiles_referenceOrder, xlfile)
			fieldInitializers, pointersInitializations := xlfile.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLFile_stagedOrder[ref] = stage.XLFile_stagedOrder[xlfile]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := xlfile.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlfile)
			// delete(stage.XLFile_stagedOrder, ref)
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
	for _, ref := range stage.XLFiles_reference {
		instance := stage.XLFiles_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.XLFiles[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.XLRows_referenceOrder[xlrow] = stage.XLRow_stagedOrder[xlrow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlrow.GongMarshallUnstaging(stage))
			// delete(stage.XLRows_referenceOrder, xlrow)
			fieldInitializers, pointersInitializations := xlrow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLRow_stagedOrder[ref] = stage.XLRow_stagedOrder[xlrow]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := xlrow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlrow)
			// delete(stage.XLRow_stagedOrder, ref)
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
	for _, ref := range stage.XLRows_reference {
		instance := stage.XLRows_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.XLRows[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.XLSheets_referenceOrder[xlsheet] = stage.XLSheet_stagedOrder[xlsheet]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlsheet.GongMarshallUnstaging(stage))
			// delete(stage.XLSheets_referenceOrder, xlsheet)
			fieldInitializers, pointersInitializations := xlsheet.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XLSheet_stagedOrder[ref] = stage.XLSheet_stagedOrder[xlsheet]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := xlsheet.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlsheet)
			// delete(stage.XLSheet_stagedOrder, ref)
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
	for _, ref := range stage.XLSheets_reference {
		instance := stage.XLSheets_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.XLSheets[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.DisplaySelections_reference = make(map[*DisplaySelection]*DisplaySelection)
	stage.DisplaySelections_referenceOrder = make(map[*DisplaySelection]uint) // diff Unstage needs the reference order
	stage.DisplaySelections_instance = make(map[*DisplaySelection]*DisplaySelection)
	for instance := range stage.DisplaySelections {
		_copy := instance.GongCopy().(*DisplaySelection)
		stage.DisplaySelections_reference[instance] = _copy
		stage.DisplaySelections_instance[_copy] = instance
		stage.DisplaySelections_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.XLCells_reference = make(map[*XLCell]*XLCell)
	stage.XLCells_referenceOrder = make(map[*XLCell]uint) // diff Unstage needs the reference order
	stage.XLCells_instance = make(map[*XLCell]*XLCell)
	for instance := range stage.XLCells {
		_copy := instance.GongCopy().(*XLCell)
		stage.XLCells_reference[instance] = _copy
		stage.XLCells_instance[_copy] = instance
		stage.XLCells_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.XLFiles_reference = make(map[*XLFile]*XLFile)
	stage.XLFiles_referenceOrder = make(map[*XLFile]uint) // diff Unstage needs the reference order
	stage.XLFiles_instance = make(map[*XLFile]*XLFile)
	for instance := range stage.XLFiles {
		_copy := instance.GongCopy().(*XLFile)
		stage.XLFiles_reference[instance] = _copy
		stage.XLFiles_instance[_copy] = instance
		stage.XLFiles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.XLRows_reference = make(map[*XLRow]*XLRow)
	stage.XLRows_referenceOrder = make(map[*XLRow]uint) // diff Unstage needs the reference order
	stage.XLRows_instance = make(map[*XLRow]*XLRow)
	for instance := range stage.XLRows {
		_copy := instance.GongCopy().(*XLRow)
		stage.XLRows_reference[instance] = _copy
		stage.XLRows_instance[_copy] = instance
		stage.XLRows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.XLSheets_reference = make(map[*XLSheet]*XLSheet)
	stage.XLSheets_referenceOrder = make(map[*XLSheet]uint) // diff Unstage needs the reference order
	stage.XLSheets_instance = make(map[*XLSheet]*XLSheet)
	for instance := range stage.XLSheets {
		_copy := instance.GongCopy().(*XLSheet)
		stage.XLSheets_reference[instance] = _copy
		stage.XLSheets_instance[_copy] = instance
		stage.XLSheets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.DisplaySelections {
		reference := stage.DisplaySelections_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.XLCells {
		reference := stage.XLCells_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.XLFiles {
		reference := stage.XLFiles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.XLRows {
		reference := stage.XLRows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.XLSheets {
		reference := stage.XLSheets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (displayselection *DisplaySelection) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DisplaySelection_stagedOrder[displayselection]; ok {
		return order
	}
	if order, ok := stage.DisplaySelections_referenceOrder[displayselection]; ok {
		return order
	} else {
		log.Printf("instance %p of type DisplaySelection was not staged and does not have a reference order", displayselection)
		return 0
	}
}

func (xlcell *XLCell) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLCell_stagedOrder[xlcell]; ok {
		return order
	}
	if order, ok := stage.XLCells_referenceOrder[xlcell]; ok {
		return order
	} else {
		log.Printf("instance %p of type XLCell was not staged and does not have a reference order", xlcell)
		return 0
	}
}

func (xlfile *XLFile) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLFile_stagedOrder[xlfile]; ok {
		return order
	}
	if order, ok := stage.XLFiles_referenceOrder[xlfile]; ok {
		return order
	} else {
		log.Printf("instance %p of type XLFile was not staged and does not have a reference order", xlfile)
		return 0
	}
}

func (xlrow *XLRow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLRow_stagedOrder[xlrow]; ok {
		return order
	}
	if order, ok := stage.XLRows_referenceOrder[xlrow]; ok {
		return order
	} else {
		log.Printf("instance %p of type XLRow was not staged and does not have a reference order", xlrow)
		return 0
	}
}

func (xlsheet *XLSheet) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XLSheet_stagedOrder[xlsheet]; ok {
		return order
	}
	if order, ok := stage.XLSheets_referenceOrder[xlsheet]; ok {
		return order
	} else {
		log.Printf("instance %p of type XLSheet was not staged and does not have a reference order", xlsheet)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", displayselection.GongGetGongstructName(), displayselection.GongGetOrder(stage))
}

func (xlcell *XLCell) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlcell.GongGetGongstructName(), xlcell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlcell *XLCell) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlcell.GongGetGongstructName(), xlcell.GongGetOrder(stage))
}

func (xlfile *XLFile) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlfile.GongGetGongstructName(), xlfile.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlfile *XLFile) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlfile.GongGetGongstructName(), xlfile.GongGetOrder(stage))
}

func (xlrow *XLRow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlrow.GongGetGongstructName(), xlrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlrow *XLRow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlrow.GongGetGongstructName(), xlrow.GongGetOrder(stage))
}

func (xlsheet *XLSheet) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsheet.GongGetGongstructName(), xlsheet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlsheet *XLSheet) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsheet.GongGetGongstructName(), xlsheet.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (displayselection *DisplaySelection) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplaySelection")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(displayselection.Name))
	return
}

func (xlcell *XLCell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLCell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(xlcell.Name))
	return
}

func (xlfile *XLFile) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLFile")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(xlfile.Name))
	return
}

func (xlrow *XLRow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLRow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(xlrow.Name))
	return
}

func (xlsheet *XLSheet) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLSheet")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(xlsheet.Name))
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

// end of template
