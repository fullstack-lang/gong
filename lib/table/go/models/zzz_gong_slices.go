// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
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
	// Compute reverse map for named struct Row
	// insertion point per field
	stage.Row_Cells_reverseMap = make(map[*Cell]*Row)
	for row := range stage.Rows {
		_ = row
		for _, _cell := range row.Cells {
			stage.Row_Cells_reverseMap[_cell] = row
		}
	}

	// Compute reverse map for named struct Table
	// insertion point per field
	stage.Table_DisplayedColumns_reverseMap = make(map[*DisplayedColumn]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _displayedcolumn := range table.DisplayedColumns {
			stage.Table_DisplayedColumns_reverseMap[_displayedcolumn] = table
		}
	}
	stage.Table_Rows_reverseMap = make(map[*Row]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _row := range table.Rows {
			stage.Table_Rows_reverseMap[_row] = table
		}
	}
	stage.Table_RowsSelectedForBulkDelete_reverseMap = make(map[*Row]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _row := range table.RowsSelectedForBulkDelete {
			stage.Table_RowsSelectedForBulkDelete_reverseMap[_row] = table
		}
	}
	stage.Table_Buttons_reverseMap = make(map[*Button]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _button := range table.Buttons {
			stage.Table_Buttons_reverseMap[_button] = table
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Buttons)

	res = __gong__appendInstances(res, stage.Cells)

	res = __gong__appendInstances(res, stage.CellBooleans)

	res = __gong__appendInstances(res, stage.CellFloat64s)

	res = __gong__appendInstances(res, stage.CellIcons)

	res = __gong__appendInstances(res, stage.CellInts)

	res = __gong__appendInstances(res, stage.CellStrings)

	res = __gong__appendInstances(res, stage.DisplayedColumns)

	res = __gong__appendInstances(res, stage.Rows)

	res = __gong__appendInstances(res, stage.SVGIcons)

	res = __gong__appendInstances(res, stage.Tables)

	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := new(Button)
	button.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cell *Cell) GongCopy() GongstructIF {
	newInstance := new(Cell)
	cell.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cellboolean *CellBoolean) GongCopy() GongstructIF {
	newInstance := new(CellBoolean)
	cellboolean.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cellfloat64 *CellFloat64) GongCopy() GongstructIF {
	newInstance := new(CellFloat64)
	cellfloat64.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cellicon *CellIcon) GongCopy() GongstructIF {
	newInstance := new(CellIcon)
	cellicon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cellint *CellInt) GongCopy() GongstructIF {
	newInstance := new(CellInt)
	cellint.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cellstring *CellString) GongCopy() GongstructIF {
	newInstance := new(CellString)
	cellstring.GongCopyBasicFields(newInstance)
	return newInstance
}

func (displayedcolumn *DisplayedColumn) GongCopy() GongstructIF {
	newInstance := new(DisplayedColumn)
	displayedcolumn.GongCopyBasicFields(newInstance)
	return newInstance
}

func (row *Row) GongCopy() GongstructIF {
	newInstance := new(Row)
	row.GongCopyBasicFields(newInstance)
	return newInstance
}

func (svgicon *SVGIcon) GongCopy() GongstructIF {
	newInstance := new(SVGIcon)
	svgicon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := new(Table)
	table.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (button *Button) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, button)
}

func (cell *Cell) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cell)
}

func (cellboolean *CellBoolean) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cellboolean)
}

func (cellfloat64 *CellFloat64) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cellfloat64)
}

func (cellicon *CellIcon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cellicon)
}

func (cellint *CellInt) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cellint)
}

func (cellstring *CellString) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cellstring)
}

func (displayedcolumn *DisplayedColumn) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, displayedcolumn)
}

func (row *Row) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, row)
}

func (svgicon *SVGIcon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, svgicon)
}

func (table *Table) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, table)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Buttons,
		stage.Button_stagedOrder,
		stage.Buttons_reference,
		&stage.Buttons_referenceOrder,
		stage.Buttons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Cells,
		stage.Cell_stagedOrder,
		stage.Cells_reference,
		&stage.Cells_referenceOrder,
		stage.Cells_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CellBooleans,
		stage.CellBoolean_stagedOrder,
		stage.CellBooleans_reference,
		&stage.CellBooleans_referenceOrder,
		stage.CellBooleans_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CellFloat64s,
		stage.CellFloat64_stagedOrder,
		stage.CellFloat64s_reference,
		&stage.CellFloat64s_referenceOrder,
		stage.CellFloat64s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CellIcons,
		stage.CellIcon_stagedOrder,
		stage.CellIcons_reference,
		&stage.CellIcons_referenceOrder,
		stage.CellIcons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CellInts,
		stage.CellInt_stagedOrder,
		stage.CellInts_reference,
		&stage.CellInts_referenceOrder,
		stage.CellInts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CellStrings,
		stage.CellString_stagedOrder,
		stage.CellStrings_reference,
		&stage.CellStrings_referenceOrder,
		stage.CellStrings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DisplayedColumns,
		stage.DisplayedColumn_stagedOrder,
		stage.DisplayedColumns_reference,
		&stage.DisplayedColumns_referenceOrder,
		stage.DisplayedColumns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Rows,
		stage.Row_stagedOrder,
		stage.Rows_reference,
		&stage.Rows_referenceOrder,
		stage.Rows_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SVGIcons,
		stage.SVGIcon_stagedOrder,
		stage.SVGIcons_reference,
		&stage.SVGIcons_referenceOrder,
		stage.SVGIcons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tables,
		stage.Table_stagedOrder,
		stage.Tables_reference,
		&stage.Tables_referenceOrder,
		stage.Tables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Buttons, &stage.Buttons_reference, &stage.Buttons_referenceOrder, &stage.Buttons_instance)

	__gong__computeReferencePass1(stage, stage.Cells, &stage.Cells_reference, &stage.Cells_referenceOrder, &stage.Cells_instance)

	__gong__computeReferencePass1(stage, stage.CellBooleans, &stage.CellBooleans_reference, &stage.CellBooleans_referenceOrder, &stage.CellBooleans_instance)

	__gong__computeReferencePass1(stage, stage.CellFloat64s, &stage.CellFloat64s_reference, &stage.CellFloat64s_referenceOrder, &stage.CellFloat64s_instance)

	__gong__computeReferencePass1(stage, stage.CellIcons, &stage.CellIcons_reference, &stage.CellIcons_referenceOrder, &stage.CellIcons_instance)

	__gong__computeReferencePass1(stage, stage.CellInts, &stage.CellInts_reference, &stage.CellInts_referenceOrder, &stage.CellInts_instance)

	__gong__computeReferencePass1(stage, stage.CellStrings, &stage.CellStrings_reference, &stage.CellStrings_referenceOrder, &stage.CellStrings_instance)

	__gong__computeReferencePass1(stage, stage.DisplayedColumns, &stage.DisplayedColumns_reference, &stage.DisplayedColumns_referenceOrder, &stage.DisplayedColumns_instance)

	__gong__computeReferencePass1(stage, stage.Rows, &stage.Rows_reference, &stage.Rows_referenceOrder, &stage.Rows_instance)

	__gong__computeReferencePass1(stage, stage.SVGIcons, &stage.SVGIcons_reference, &stage.SVGIcons_referenceOrder, &stage.SVGIcons_instance)

	__gong__computeReferencePass1(stage, stage.Tables, &stage.Tables_reference, &stage.Tables_referenceOrder, &stage.Tables_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Buttons, stage.Buttons_reference, stage)

	__gong__computeReferencePass2(stage.Cells, stage.Cells_reference, stage)

	__gong__computeReferencePass2(stage.CellBooleans, stage.CellBooleans_reference, stage)

	__gong__computeReferencePass2(stage.CellFloat64s, stage.CellFloat64s_reference, stage)

	__gong__computeReferencePass2(stage.CellIcons, stage.CellIcons_reference, stage)

	__gong__computeReferencePass2(stage.CellInts, stage.CellInts_reference, stage)

	__gong__computeReferencePass2(stage.CellStrings, stage.CellStrings_reference, stage)

	__gong__computeReferencePass2(stage.DisplayedColumns, stage.DisplayedColumns_reference, stage)

	__gong__computeReferencePass2(stage.Rows, stage.Rows_reference, stage)

	__gong__computeReferencePass2(stage.SVGIcons, stage.SVGIcons_reference, stage)

	__gong__computeReferencePass2(stage.Tables, stage.Tables_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (button *Button) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Button_stagedOrder, stage.Buttons_referenceOrder, button, "Button")
}

func (cell *Cell) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Cell_stagedOrder, stage.Cells_referenceOrder, cell, "Cell")
}

func (cellboolean *CellBoolean) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CellBoolean_stagedOrder, stage.CellBooleans_referenceOrder, cellboolean, "CellBoolean")
}

func (cellfloat64 *CellFloat64) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CellFloat64_stagedOrder, stage.CellFloat64s_referenceOrder, cellfloat64, "CellFloat64")
}

func (cellicon *CellIcon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CellIcon_stagedOrder, stage.CellIcons_referenceOrder, cellicon, "CellIcon")
}

func (cellint *CellInt) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CellInt_stagedOrder, stage.CellInts_referenceOrder, cellint, "CellInt")
}

func (cellstring *CellString) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CellString_stagedOrder, stage.CellStrings_referenceOrder, cellstring, "CellString")
}

func (displayedcolumn *DisplayedColumn) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DisplayedColumn_stagedOrder, stage.DisplayedColumns_referenceOrder, displayedcolumn, "DisplayedColumn")
}

func (row *Row) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Row_stagedOrder, stage.Rows_referenceOrder, row, "Row")
}

func (svgicon *SVGIcon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SVGIcon_stagedOrder, stage.SVGIcons_referenceOrder, svgicon, "SVGIcon")
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Table_stagedOrder, stage.Tables_referenceOrder, table, "Table")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(button, button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return button.GongGetIdentifier(stage)
}

func (cell *Cell) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cell, cell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cell *Cell) GongGetReferenceIdentifier(stage *Stage) string {
	return cell.GongGetIdentifier(stage)
}

func (cellboolean *CellBoolean) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cellboolean, cellboolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellboolean *CellBoolean) GongGetReferenceIdentifier(stage *Stage) string {
	return cellboolean.GongGetIdentifier(stage)
}

func (cellfloat64 *CellFloat64) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cellfloat64, cellfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellfloat64 *CellFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return cellfloat64.GongGetIdentifier(stage)
}

func (cellicon *CellIcon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cellicon, cellicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellicon *CellIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return cellicon.GongGetIdentifier(stage)
}

func (cellint *CellInt) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cellint, cellint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellint *CellInt) GongGetReferenceIdentifier(stage *Stage) string {
	return cellint.GongGetIdentifier(stage)
}

func (cellstring *CellString) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cellstring, cellstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellstring *CellString) GongGetReferenceIdentifier(stage *Stage) string {
	return cellstring.GongGetIdentifier(stage)
}

func (displayedcolumn *DisplayedColumn) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(displayedcolumn, displayedcolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (displayedcolumn *DisplayedColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return displayedcolumn.GongGetIdentifier(stage)
}

func (row *Row) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(row, row.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (row *Row) GongGetReferenceIdentifier(stage *Stage) string {
	return row.GongGetIdentifier(stage)
}

func (svgicon *SVGIcon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(svgicon, svgicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgicon *SVGIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return svgicon.GongGetIdentifier(stage)
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(table, table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return table.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (button *Button) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(button.GongGetIdentifier(stage), "Button", button.Name)
}

func (cell *Cell) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cell.GongGetIdentifier(stage), "Cell", cell.Name)
}

func (cellboolean *CellBoolean) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cellboolean.GongGetIdentifier(stage), "CellBoolean", cellboolean.Name)
}

func (cellfloat64 *CellFloat64) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cellfloat64.GongGetIdentifier(stage), "CellFloat64", cellfloat64.Name)
}

func (cellicon *CellIcon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cellicon.GongGetIdentifier(stage), "CellIcon", cellicon.Name)
}

func (cellint *CellInt) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cellint.GongGetIdentifier(stage), "CellInt", cellint.Name)
}

func (cellstring *CellString) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cellstring.GongGetIdentifier(stage), "CellString", cellstring.Name)
}

func (displayedcolumn *DisplayedColumn) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(displayedcolumn.GongGetIdentifier(stage), "DisplayedColumn", displayedcolumn.Name)
}

func (row *Row) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(row.GongGetIdentifier(stage), "Row", row.Name)
}

func (svgicon *SVGIcon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(svgicon.GongGetIdentifier(stage), "SVGIcon", svgicon.Name)
}

func (table *Table) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(table.GongGetIdentifier(stage), "Table", table.Name)
}

// insertion point for unstaging
func (button *Button) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(button.GongGetReferenceIdentifier(stage))
}

func (cell *Cell) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cell.GongGetReferenceIdentifier(stage))
}

func (cellboolean *CellBoolean) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cellboolean.GongGetReferenceIdentifier(stage))
}

func (cellfloat64 *CellFloat64) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cellfloat64.GongGetReferenceIdentifier(stage))
}

func (cellicon *CellIcon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cellicon.GongGetReferenceIdentifier(stage))
}

func (cellint *CellInt) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cellint.GongGetReferenceIdentifier(stage))
}

func (cellstring *CellString) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cellstring.GongGetReferenceIdentifier(stage))
}

func (displayedcolumn *DisplayedColumn) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(displayedcolumn.GongGetReferenceIdentifier(stage))
}

func (row *Row) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(row.GongGetReferenceIdentifier(stage))
}

func (svgicon *SVGIcon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(svgicon.GongGetReferenceIdentifier(stage))
}

func (table *Table) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(table.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
