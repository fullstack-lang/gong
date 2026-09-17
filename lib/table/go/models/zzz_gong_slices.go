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
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Cell
	// insertion point per field

	// Compute reverse map for named struct CellBoolean
	// insertion point per field

	// Compute reverse map for named struct CellFloat64
	// insertion point per field

	// Compute reverse map for named struct CellIcon
	// insertion point per field

	// Compute reverse map for named struct CellInt
	// insertion point per field

	// Compute reverse map for named struct CellString
	// insertion point per field

	// Compute reverse map for named struct DisplayedColumn
	// insertion point per field

	// Compute reverse map for named struct Row
	// insertion point per field
	stage.Row_Cells_reverseMap = make(map[*Cell]*Row)
	for row := range stage.Rows {
		_ = row
		for _, _cell := range row.Cells {
			stage.Row_Cells_reverseMap[_cell] = row
		}
	}

	// Compute reverse map for named struct SVGIcon
	// insertion point per field

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
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Cells {
		res = append(res, instance)
	}

	for instance := range stage.CellBooleans {
		res = append(res, instance)
	}

	for instance := range stage.CellFloat64s {
		res = append(res, instance)
	}

	for instance := range stage.CellIcons {
		res = append(res, instance)
	}

	for instance := range stage.CellInts {
		res = append(res, instance)
	}

	for instance := range stage.CellStrings {
		res = append(res, instance)
	}

	for instance := range stage.DisplayedColumns {
		res = append(res, instance)
	}

	for instance := range stage.Rows {
		res = append(res, instance)
	}

	for instance := range stage.SVGIcons {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

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
func (button *Button) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(button).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(button), uint64(stage.GetOrder(button)))
	return
}

func (cell *Cell) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cell).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cell), uint64(stage.GetOrder(cell)))
	return
}

func (cellboolean *CellBoolean) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellboolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cellboolean), uint64(stage.GetOrder(cellboolean)))
	return
}

func (cellfloat64 *CellFloat64) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellfloat64).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cellfloat64), uint64(stage.GetOrder(cellfloat64)))
	return
}

func (cellicon *CellIcon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellicon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cellicon), uint64(stage.GetOrder(cellicon)))
	return
}

func (cellint *CellInt) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellint).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cellint), uint64(stage.GetOrder(cellint)))
	return
}

func (cellstring *CellString) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellstring).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cellstring), uint64(stage.GetOrder(cellstring)))
	return
}

func (displayedcolumn *DisplayedColumn) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(displayedcolumn).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(displayedcolumn), uint64(stage.GetOrder(displayedcolumn)))
	return
}

func (row *Row) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(row).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(row), uint64(stage.GetOrder(row)))
	return
}

func (svgicon *SVGIcon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svgicon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(svgicon), uint64(stage.GetOrder(svgicon)))
	return
}

func (table *Table) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(table).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(table), uint64(stage.GetOrder(table)))
	return
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
	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint) // diff Unstage needs the reference order
	stage.Buttons_instance = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		_copy := instance.GongCopy().(*Button)
		stage.Buttons_reference[instance] = _copy
		stage.Buttons_instance[_copy] = instance
		stage.Buttons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Cells_reference = make(map[*Cell]*Cell)
	stage.Cells_referenceOrder = make(map[*Cell]uint) // diff Unstage needs the reference order
	stage.Cells_instance = make(map[*Cell]*Cell)
	for instance := range stage.Cells {
		_copy := instance.GongCopy().(*Cell)
		stage.Cells_reference[instance] = _copy
		stage.Cells_instance[_copy] = instance
		stage.Cells_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CellBooleans_reference = make(map[*CellBoolean]*CellBoolean)
	stage.CellBooleans_referenceOrder = make(map[*CellBoolean]uint) // diff Unstage needs the reference order
	stage.CellBooleans_instance = make(map[*CellBoolean]*CellBoolean)
	for instance := range stage.CellBooleans {
		_copy := instance.GongCopy().(*CellBoolean)
		stage.CellBooleans_reference[instance] = _copy
		stage.CellBooleans_instance[_copy] = instance
		stage.CellBooleans_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CellFloat64s_reference = make(map[*CellFloat64]*CellFloat64)
	stage.CellFloat64s_referenceOrder = make(map[*CellFloat64]uint) // diff Unstage needs the reference order
	stage.CellFloat64s_instance = make(map[*CellFloat64]*CellFloat64)
	for instance := range stage.CellFloat64s {
		_copy := instance.GongCopy().(*CellFloat64)
		stage.CellFloat64s_reference[instance] = _copy
		stage.CellFloat64s_instance[_copy] = instance
		stage.CellFloat64s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CellIcons_reference = make(map[*CellIcon]*CellIcon)
	stage.CellIcons_referenceOrder = make(map[*CellIcon]uint) // diff Unstage needs the reference order
	stage.CellIcons_instance = make(map[*CellIcon]*CellIcon)
	for instance := range stage.CellIcons {
		_copy := instance.GongCopy().(*CellIcon)
		stage.CellIcons_reference[instance] = _copy
		stage.CellIcons_instance[_copy] = instance
		stage.CellIcons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CellInts_reference = make(map[*CellInt]*CellInt)
	stage.CellInts_referenceOrder = make(map[*CellInt]uint) // diff Unstage needs the reference order
	stage.CellInts_instance = make(map[*CellInt]*CellInt)
	for instance := range stage.CellInts {
		_copy := instance.GongCopy().(*CellInt)
		stage.CellInts_reference[instance] = _copy
		stage.CellInts_instance[_copy] = instance
		stage.CellInts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.CellStrings_reference = make(map[*CellString]*CellString)
	stage.CellStrings_referenceOrder = make(map[*CellString]uint) // diff Unstage needs the reference order
	stage.CellStrings_instance = make(map[*CellString]*CellString)
	for instance := range stage.CellStrings {
		_copy := instance.GongCopy().(*CellString)
		stage.CellStrings_reference[instance] = _copy
		stage.CellStrings_instance[_copy] = instance
		stage.CellStrings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DisplayedColumns_reference = make(map[*DisplayedColumn]*DisplayedColumn)
	stage.DisplayedColumns_referenceOrder = make(map[*DisplayedColumn]uint) // diff Unstage needs the reference order
	stage.DisplayedColumns_instance = make(map[*DisplayedColumn]*DisplayedColumn)
	for instance := range stage.DisplayedColumns {
		_copy := instance.GongCopy().(*DisplayedColumn)
		stage.DisplayedColumns_reference[instance] = _copy
		stage.DisplayedColumns_instance[_copy] = instance
		stage.DisplayedColumns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Rows_reference = make(map[*Row]*Row)
	stage.Rows_referenceOrder = make(map[*Row]uint) // diff Unstage needs the reference order
	stage.Rows_instance = make(map[*Row]*Row)
	for instance := range stage.Rows {
		_copy := instance.GongCopy().(*Row)
		stage.Rows_reference[instance] = _copy
		stage.Rows_instance[_copy] = instance
		stage.Rows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SVGIcons_reference = make(map[*SVGIcon]*SVGIcon)
	stage.SVGIcons_referenceOrder = make(map[*SVGIcon]uint) // diff Unstage needs the reference order
	stage.SVGIcons_instance = make(map[*SVGIcon]*SVGIcon)
	for instance := range stage.SVGIcons {
		_copy := instance.GongCopy().(*SVGIcon)
		stage.SVGIcons_reference[instance] = _copy
		stage.SVGIcons_instance[_copy] = instance
		stage.SVGIcons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	stage.Tables_instance = make(map[*Table]*Table)
	for instance := range stage.Tables {
		_copy := instance.GongCopy().(*Table)
		stage.Tables_reference[instance] = _copy
		stage.Tables_instance[_copy] = instance
		stage.Tables_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Buttons {
		reference := stage.Buttons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Cells {
		reference := stage.Cells_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CellBooleans {
		reference := stage.CellBooleans_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CellFloat64s {
		reference := stage.CellFloat64s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CellIcons {
		reference := stage.CellIcons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CellInts {
		reference := stage.CellInts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.CellStrings {
		reference := stage.CellStrings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DisplayedColumns {
		reference := stage.DisplayedColumns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rows {
		reference := stage.Rows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SVGIcons {
		reference := stage.SVGIcons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tables {
		reference := stage.Tables_reference[instance]
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
func (button *Button) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Button_stagedOrder[button]; ok {
		return order
	}
	if order, ok := stage.Buttons_referenceOrder[button]; ok {
		return order
	} else {
		log.Printf("instance %p of type Button was not staged and does not have a reference order", button)
		return 0
	}
}

func (cell *Cell) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Cell_stagedOrder[cell]; ok {
		return order
	}
	if order, ok := stage.Cells_referenceOrder[cell]; ok {
		return order
	} else {
		log.Printf("instance %p of type Cell was not staged and does not have a reference order", cell)
		return 0
	}
}

func (cellboolean *CellBoolean) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CellBoolean_stagedOrder[cellboolean]; ok {
		return order
	}
	if order, ok := stage.CellBooleans_referenceOrder[cellboolean]; ok {
		return order
	} else {
		log.Printf("instance %p of type CellBoolean was not staged and does not have a reference order", cellboolean)
		return 0
	}
}

func (cellfloat64 *CellFloat64) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CellFloat64_stagedOrder[cellfloat64]; ok {
		return order
	}
	if order, ok := stage.CellFloat64s_referenceOrder[cellfloat64]; ok {
		return order
	} else {
		log.Printf("instance %p of type CellFloat64 was not staged and does not have a reference order", cellfloat64)
		return 0
	}
}

func (cellicon *CellIcon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CellIcon_stagedOrder[cellicon]; ok {
		return order
	}
	if order, ok := stage.CellIcons_referenceOrder[cellicon]; ok {
		return order
	} else {
		log.Printf("instance %p of type CellIcon was not staged and does not have a reference order", cellicon)
		return 0
	}
}

func (cellint *CellInt) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CellInt_stagedOrder[cellint]; ok {
		return order
	}
	if order, ok := stage.CellInts_referenceOrder[cellint]; ok {
		return order
	} else {
		log.Printf("instance %p of type CellInt was not staged and does not have a reference order", cellint)
		return 0
	}
}

func (cellstring *CellString) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CellString_stagedOrder[cellstring]; ok {
		return order
	}
	if order, ok := stage.CellStrings_referenceOrder[cellstring]; ok {
		return order
	} else {
		log.Printf("instance %p of type CellString was not staged and does not have a reference order", cellstring)
		return 0
	}
}

func (displayedcolumn *DisplayedColumn) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DisplayedColumn_stagedOrder[displayedcolumn]; ok {
		return order
	}
	if order, ok := stage.DisplayedColumns_referenceOrder[displayedcolumn]; ok {
		return order
	} else {
		log.Printf("instance %p of type DisplayedColumn was not staged and does not have a reference order", displayedcolumn)
		return 0
	}
}

func (row *Row) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Row_stagedOrder[row]; ok {
		return order
	}
	if order, ok := stage.Rows_referenceOrder[row]; ok {
		return order
	} else {
		log.Printf("instance %p of type Row was not staged and does not have a reference order", row)
		return 0
	}
}

func (svgicon *SVGIcon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SVGIcon_stagedOrder[svgicon]; ok {
		return order
	}
	if order, ok := stage.SVGIcons_referenceOrder[svgicon]; ok {
		return order
	} else {
		log.Printf("instance %p of type SVGIcon was not staged and does not have a reference order", svgicon)
		return 0
	}
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Table_stagedOrder[table]; ok {
		return order
	}
	if order, ok := stage.Tables_referenceOrder[table]; ok {
		return order
	} else {
		log.Printf("instance %p of type Table was not staged and does not have a reference order", table)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

func (cell *Cell) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cell.GongGetGongstructName(), cell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cell *Cell) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cell.GongGetGongstructName(), cell.GongGetOrder(stage))
}

func (cellboolean *CellBoolean) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellboolean.GongGetGongstructName(), cellboolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellboolean *CellBoolean) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellboolean.GongGetGongstructName(), cellboolean.GongGetOrder(stage))
}

func (cellfloat64 *CellFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellfloat64.GongGetGongstructName(), cellfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellfloat64 *CellFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellfloat64.GongGetGongstructName(), cellfloat64.GongGetOrder(stage))
}

func (cellicon *CellIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellicon.GongGetGongstructName(), cellicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellicon *CellIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellicon.GongGetGongstructName(), cellicon.GongGetOrder(stage))
}

func (cellint *CellInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellint.GongGetGongstructName(), cellint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellint *CellInt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellint.GongGetGongstructName(), cellint.GongGetOrder(stage))
}

func (cellstring *CellString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellstring.GongGetGongstructName(), cellstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellstring *CellString) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellstring.GongGetGongstructName(), cellstring.GongGetOrder(stage))
}

func (displayedcolumn *DisplayedColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (displayedcolumn *DisplayedColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

func (row *Row) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (row *Row) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
}

func (svgicon *SVGIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgicon.GongGetGongstructName(), svgicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svgicon *SVGIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svgicon.GongGetGongstructName(), svgicon.GongGetOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (button *Button) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(button.Name))
	return
}

func (cell *Cell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cell.Name))
	return
}

func (cellboolean *CellBoolean) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellBoolean")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cellboolean.Name))
	return
}

func (cellfloat64 *CellFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cellfloat64.Name))
	return
}

func (cellicon *CellIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cellicon.Name))
	return
}

func (cellint *CellInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cellint.Name))
	return
}

func (cellstring *CellString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cellstring.Name))
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplayedColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(displayedcolumn.Name))
	return
}

func (row *Row) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Row")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(row.Name))
	return
}

func (svgicon *SVGIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SVGIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(svgicon.Name))
	return
}

func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(table.Name))
	return
}

// insertion point for unstaging
func (button *Button) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetReferenceIdentifier(stage))
	return
}

func (cell *Cell) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cell.GongGetReferenceIdentifier(stage))
	return
}

func (cellboolean *CellBoolean) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellboolean.GongGetReferenceIdentifier(stage))
	return
}

func (cellfloat64 *CellFloat64) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellfloat64.GongGetReferenceIdentifier(stage))
	return
}

func (cellicon *CellIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellicon.GongGetReferenceIdentifier(stage))
	return
}

func (cellint *CellInt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellint.GongGetReferenceIdentifier(stage))
	return
}

func (cellstring *CellString) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellstring.GongGetReferenceIdentifier(stage))
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetReferenceIdentifier(stage))
	return
}

func (row *Row) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetReferenceIdentifier(stage))
	return
}

func (svgicon *SVGIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetReferenceIdentifier(stage))
	return
}

func (table *Table) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetReferenceIdentifier(stage))
	return
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

// end of template
