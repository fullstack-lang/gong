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
	button.CopyBasicFields(newInstance)
	return newInstance
}

func (cell *Cell) GongCopy() GongstructIF {
	newInstance := new(Cell)
	cell.CopyBasicFields(newInstance)
	return newInstance
}

func (cellboolean *CellBoolean) GongCopy() GongstructIF {
	newInstance := new(CellBoolean)
	cellboolean.CopyBasicFields(newInstance)
	return newInstance
}

func (cellfloat64 *CellFloat64) GongCopy() GongstructIF {
	newInstance := new(CellFloat64)
	cellfloat64.CopyBasicFields(newInstance)
	return newInstance
}

func (cellicon *CellIcon) GongCopy() GongstructIF {
	newInstance := new(CellIcon)
	cellicon.CopyBasicFields(newInstance)
	return newInstance
}

func (cellint *CellInt) GongCopy() GongstructIF {
	newInstance := new(CellInt)
	cellint.CopyBasicFields(newInstance)
	return newInstance
}

func (cellstring *CellString) GongCopy() GongstructIF {
	newInstance := new(CellString)
	cellstring.CopyBasicFields(newInstance)
	return newInstance
}

func (displayedcolumn *DisplayedColumn) GongCopy() GongstructIF {
	newInstance := new(DisplayedColumn)
	displayedcolumn.CopyBasicFields(newInstance)
	return newInstance
}

func (row *Row) GongCopy() GongstructIF {
	newInstance := new(Row)
	row.CopyBasicFields(newInstance)
	return newInstance
}

func (svgicon *SVGIcon) GongCopy() GongstructIF {
	newInstance := new(SVGIcon)
	svgicon.CopyBasicFields(newInstance)
	return newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := new(Table)
	table.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (button *Button) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(button).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(button), uint64(GetOrderPointerGongstruct(stage, button)))
	return
}

func (cell *Cell) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cell).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cell), uint64(GetOrderPointerGongstruct(stage, cell)))
	return
}

func (cellboolean *CellBoolean) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellboolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cellboolean), uint64(GetOrderPointerGongstruct(stage, cellboolean)))
	return
}

func (cellfloat64 *CellFloat64) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellfloat64).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cellfloat64), uint64(GetOrderPointerGongstruct(stage, cellfloat64)))
	return
}

func (cellicon *CellIcon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellicon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cellicon), uint64(GetOrderPointerGongstruct(stage, cellicon)))
	return
}

func (cellint *CellInt) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellint).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cellint), uint64(GetOrderPointerGongstruct(stage, cellint)))
	return
}

func (cellstring *CellString) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cellstring).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(cellstring), uint64(GetOrderPointerGongstruct(stage, cellstring)))
	return
}

func (displayedcolumn *DisplayedColumn) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(displayedcolumn).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(displayedcolumn), uint64(GetOrderPointerGongstruct(stage, displayedcolumn)))
	return
}

func (row *Row) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(row).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(row), uint64(GetOrderPointerGongstruct(stage, row)))
	return
}

func (svgicon *SVGIcon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svgicon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(svgicon), uint64(GetOrderPointerGongstruct(stage, svgicon)))
	return
}

func (table *Table) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(table).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(table), uint64(GetOrderPointerGongstruct(stage, table)))
	return
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
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			newInstancesSlice = append(newInstancesSlice, button.GongMarshallIdentifier(stage))
			if stage.Buttons_referenceOrder == nil {
				stage.Buttons_referenceOrder = make(map[*Button]uint)
			}
			stage.Buttons_referenceOrder[button] = stage.Button_stagedOrder[button]
			newInstancesReverseSlice = append(newInstancesReverseSlice, button.GongMarshallUnstaging(stage))
			// delete(stage.Buttons_referenceOrder, button)
			fieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Button_stagedOrder[ref] = stage.Button_stagedOrder[button]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := button.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, button)
			// delete(stage.Button_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", button.GetName())
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
	for _, ref := range stage.Buttons_reference {
		instance := stage.Buttons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Buttons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			buttons_deletedInstances = append(buttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(buttons_newInstances)
	lenDeletedInstances += len(buttons_deletedInstances)
	var cells_newInstances []*Cell
	var cells_deletedInstances []*Cell

	// parse all staged instances and check if they have a reference
	for cell := range stage.Cells {
		if ref, ok := stage.Cells_reference[cell]; !ok {
			cells_newInstances = append(cells_newInstances, cell)
			newInstancesSlice = append(newInstancesSlice, cell.GongMarshallIdentifier(stage))
			if stage.Cells_referenceOrder == nil {
				stage.Cells_referenceOrder = make(map[*Cell]uint)
			}
			stage.Cells_referenceOrder[cell] = stage.Cell_stagedOrder[cell]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cell.GongMarshallUnstaging(stage))
			// delete(stage.Cells_referenceOrder, cell)
			fieldInitializers, pointersInitializations := cell.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Cell_stagedOrder[ref] = stage.Cell_stagedOrder[cell]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cell.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cell)
			// delete(stage.Cell_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cell.GetName())
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
	for _, ref := range stage.Cells_reference {
		instance := stage.Cells_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Cells[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cells_deletedInstances = append(cells_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cells_newInstances)
	lenDeletedInstances += len(cells_deletedInstances)
	var cellbooleans_newInstances []*CellBoolean
	var cellbooleans_deletedInstances []*CellBoolean

	// parse all staged instances and check if they have a reference
	for cellboolean := range stage.CellBooleans {
		if ref, ok := stage.CellBooleans_reference[cellboolean]; !ok {
			cellbooleans_newInstances = append(cellbooleans_newInstances, cellboolean)
			newInstancesSlice = append(newInstancesSlice, cellboolean.GongMarshallIdentifier(stage))
			if stage.CellBooleans_referenceOrder == nil {
				stage.CellBooleans_referenceOrder = make(map[*CellBoolean]uint)
			}
			stage.CellBooleans_referenceOrder[cellboolean] = stage.CellBoolean_stagedOrder[cellboolean]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cellboolean.GongMarshallUnstaging(stage))
			// delete(stage.CellBooleans_referenceOrder, cellboolean)
			fieldInitializers, pointersInitializations := cellboolean.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CellBoolean_stagedOrder[ref] = stage.CellBoolean_stagedOrder[cellboolean]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cellboolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellboolean)
			// delete(stage.CellBoolean_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cellboolean.GetName())
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
	for _, ref := range stage.CellBooleans_reference {
		instance := stage.CellBooleans_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CellBooleans[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cellbooleans_deletedInstances = append(cellbooleans_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cellbooleans_newInstances)
	lenDeletedInstances += len(cellbooleans_deletedInstances)
	var cellfloat64s_newInstances []*CellFloat64
	var cellfloat64s_deletedInstances []*CellFloat64

	// parse all staged instances and check if they have a reference
	for cellfloat64 := range stage.CellFloat64s {
		if ref, ok := stage.CellFloat64s_reference[cellfloat64]; !ok {
			cellfloat64s_newInstances = append(cellfloat64s_newInstances, cellfloat64)
			newInstancesSlice = append(newInstancesSlice, cellfloat64.GongMarshallIdentifier(stage))
			if stage.CellFloat64s_referenceOrder == nil {
				stage.CellFloat64s_referenceOrder = make(map[*CellFloat64]uint)
			}
			stage.CellFloat64s_referenceOrder[cellfloat64] = stage.CellFloat64_stagedOrder[cellfloat64]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cellfloat64.GongMarshallUnstaging(stage))
			// delete(stage.CellFloat64s_referenceOrder, cellfloat64)
			fieldInitializers, pointersInitializations := cellfloat64.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CellFloat64_stagedOrder[ref] = stage.CellFloat64_stagedOrder[cellfloat64]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cellfloat64.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellfloat64)
			// delete(stage.CellFloat64_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cellfloat64.GetName())
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
	for _, ref := range stage.CellFloat64s_reference {
		instance := stage.CellFloat64s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CellFloat64s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cellfloat64s_deletedInstances = append(cellfloat64s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cellfloat64s_newInstances)
	lenDeletedInstances += len(cellfloat64s_deletedInstances)
	var cellicons_newInstances []*CellIcon
	var cellicons_deletedInstances []*CellIcon

	// parse all staged instances and check if they have a reference
	for cellicon := range stage.CellIcons {
		if ref, ok := stage.CellIcons_reference[cellicon]; !ok {
			cellicons_newInstances = append(cellicons_newInstances, cellicon)
			newInstancesSlice = append(newInstancesSlice, cellicon.GongMarshallIdentifier(stage))
			if stage.CellIcons_referenceOrder == nil {
				stage.CellIcons_referenceOrder = make(map[*CellIcon]uint)
			}
			stage.CellIcons_referenceOrder[cellicon] = stage.CellIcon_stagedOrder[cellicon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cellicon.GongMarshallUnstaging(stage))
			// delete(stage.CellIcons_referenceOrder, cellicon)
			fieldInitializers, pointersInitializations := cellicon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CellIcon_stagedOrder[ref] = stage.CellIcon_stagedOrder[cellicon]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cellicon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellicon)
			// delete(stage.CellIcon_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cellicon.GetName())
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
	for _, ref := range stage.CellIcons_reference {
		instance := stage.CellIcons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CellIcons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cellicons_deletedInstances = append(cellicons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cellicons_newInstances)
	lenDeletedInstances += len(cellicons_deletedInstances)
	var cellints_newInstances []*CellInt
	var cellints_deletedInstances []*CellInt

	// parse all staged instances and check if they have a reference
	for cellint := range stage.CellInts {
		if ref, ok := stage.CellInts_reference[cellint]; !ok {
			cellints_newInstances = append(cellints_newInstances, cellint)
			newInstancesSlice = append(newInstancesSlice, cellint.GongMarshallIdentifier(stage))
			if stage.CellInts_referenceOrder == nil {
				stage.CellInts_referenceOrder = make(map[*CellInt]uint)
			}
			stage.CellInts_referenceOrder[cellint] = stage.CellInt_stagedOrder[cellint]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cellint.GongMarshallUnstaging(stage))
			// delete(stage.CellInts_referenceOrder, cellint)
			fieldInitializers, pointersInitializations := cellint.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CellInt_stagedOrder[ref] = stage.CellInt_stagedOrder[cellint]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cellint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellint)
			// delete(stage.CellInt_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cellint.GetName())
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
	for _, ref := range stage.CellInts_reference {
		instance := stage.CellInts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CellInts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cellints_deletedInstances = append(cellints_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cellints_newInstances)
	lenDeletedInstances += len(cellints_deletedInstances)
	var cellstrings_newInstances []*CellString
	var cellstrings_deletedInstances []*CellString

	// parse all staged instances and check if they have a reference
	for cellstring := range stage.CellStrings {
		if ref, ok := stage.CellStrings_reference[cellstring]; !ok {
			cellstrings_newInstances = append(cellstrings_newInstances, cellstring)
			newInstancesSlice = append(newInstancesSlice, cellstring.GongMarshallIdentifier(stage))
			if stage.CellStrings_referenceOrder == nil {
				stage.CellStrings_referenceOrder = make(map[*CellString]uint)
			}
			stage.CellStrings_referenceOrder[cellstring] = stage.CellString_stagedOrder[cellstring]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cellstring.GongMarshallUnstaging(stage))
			// delete(stage.CellStrings_referenceOrder, cellstring)
			fieldInitializers, pointersInitializations := cellstring.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CellString_stagedOrder[ref] = stage.CellString_stagedOrder[cellstring]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := cellstring.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellstring)
			// delete(stage.CellString_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cellstring.GetName())
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
	for _, ref := range stage.CellStrings_reference {
		instance := stage.CellStrings_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CellStrings[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			cellstrings_deletedInstances = append(cellstrings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cellstrings_newInstances)
	lenDeletedInstances += len(cellstrings_deletedInstances)
	var displayedcolumns_newInstances []*DisplayedColumn
	var displayedcolumns_deletedInstances []*DisplayedColumn

	// parse all staged instances and check if they have a reference
	for displayedcolumn := range stage.DisplayedColumns {
		if ref, ok := stage.DisplayedColumns_reference[displayedcolumn]; !ok {
			displayedcolumns_newInstances = append(displayedcolumns_newInstances, displayedcolumn)
			newInstancesSlice = append(newInstancesSlice, displayedcolumn.GongMarshallIdentifier(stage))
			if stage.DisplayedColumns_referenceOrder == nil {
				stage.DisplayedColumns_referenceOrder = make(map[*DisplayedColumn]uint)
			}
			stage.DisplayedColumns_referenceOrder[displayedcolumn] = stage.DisplayedColumn_stagedOrder[displayedcolumn]
			newInstancesReverseSlice = append(newInstancesReverseSlice, displayedcolumn.GongMarshallUnstaging(stage))
			// delete(stage.DisplayedColumns_referenceOrder, displayedcolumn)
			fieldInitializers, pointersInitializations := displayedcolumn.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DisplayedColumn_stagedOrder[ref] = stage.DisplayedColumn_stagedOrder[displayedcolumn]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := displayedcolumn.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, displayedcolumn)
			// delete(stage.DisplayedColumn_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", displayedcolumn.GetName())
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
	for _, ref := range stage.DisplayedColumns_reference {
		instance := stage.DisplayedColumns_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DisplayedColumns[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			displayedcolumns_deletedInstances = append(displayedcolumns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(displayedcolumns_newInstances)
	lenDeletedInstances += len(displayedcolumns_deletedInstances)
	var rows_newInstances []*Row
	var rows_deletedInstances []*Row

	// parse all staged instances and check if they have a reference
	for row := range stage.Rows {
		if ref, ok := stage.Rows_reference[row]; !ok {
			rows_newInstances = append(rows_newInstances, row)
			newInstancesSlice = append(newInstancesSlice, row.GongMarshallIdentifier(stage))
			if stage.Rows_referenceOrder == nil {
				stage.Rows_referenceOrder = make(map[*Row]uint)
			}
			stage.Rows_referenceOrder[row] = stage.Row_stagedOrder[row]
			newInstancesReverseSlice = append(newInstancesReverseSlice, row.GongMarshallUnstaging(stage))
			// delete(stage.Rows_referenceOrder, row)
			fieldInitializers, pointersInitializations := row.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Row_stagedOrder[ref] = stage.Row_stagedOrder[row]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := row.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, row)
			// delete(stage.Row_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", row.GetName())
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
	for _, ref := range stage.Rows_reference {
		instance := stage.Rows_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Rows[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			rows_deletedInstances = append(rows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(rows_newInstances)
	lenDeletedInstances += len(rows_deletedInstances)
	var svgicons_newInstances []*SVGIcon
	var svgicons_deletedInstances []*SVGIcon

	// parse all staged instances and check if they have a reference
	for svgicon := range stage.SVGIcons {
		if ref, ok := stage.SVGIcons_reference[svgicon]; !ok {
			svgicons_newInstances = append(svgicons_newInstances, svgicon)
			newInstancesSlice = append(newInstancesSlice, svgicon.GongMarshallIdentifier(stage))
			if stage.SVGIcons_referenceOrder == nil {
				stage.SVGIcons_referenceOrder = make(map[*SVGIcon]uint)
			}
			stage.SVGIcons_referenceOrder[svgicon] = stage.SVGIcon_stagedOrder[svgicon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svgicon.GongMarshallUnstaging(stage))
			// delete(stage.SVGIcons_referenceOrder, svgicon)
			fieldInitializers, pointersInitializations := svgicon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SVGIcon_stagedOrder[ref] = stage.SVGIcon_stagedOrder[svgicon]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := svgicon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svgicon)
			// delete(stage.SVGIcon_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svgicon.GetName())
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
	for _, ref := range stage.SVGIcons_reference {
		instance := stage.SVGIcons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SVGIcons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			svgicons_deletedInstances = append(svgicons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgicons_newInstances)
	lenDeletedInstances += len(svgicons_deletedInstances)
	var tables_newInstances []*Table
	var tables_deletedInstances []*Table

	// parse all staged instances and check if they have a reference
	for table := range stage.Tables {
		if ref, ok := stage.Tables_reference[table]; !ok {
			tables_newInstances = append(tables_newInstances, table)
			newInstancesSlice = append(newInstancesSlice, table.GongMarshallIdentifier(stage))
			if stage.Tables_referenceOrder == nil {
				stage.Tables_referenceOrder = make(map[*Table]uint)
			}
			stage.Tables_referenceOrder[table] = stage.Table_stagedOrder[table]
			newInstancesReverseSlice = append(newInstancesReverseSlice, table.GongMarshallUnstaging(stage))
			// delete(stage.Tables_referenceOrder, table)
			fieldInitializers, pointersInitializations := table.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Table_stagedOrder[ref] = stage.Table_stagedOrder[table]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := table.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, table)
			// delete(stage.Table_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", table.GetName())
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
	for _, ref := range stage.Tables_reference {
		instance := stage.Tables_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Tables[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			tables_deletedInstances = append(tables_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)

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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(button.Name))
	return
}

func (cell *Cell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cell.Name))
	return
}

func (cellboolean *CellBoolean) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellBoolean")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellboolean.Name))
	return
}

func (cellfloat64 *CellFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellfloat64.Name))
	return
}

func (cellicon *CellIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellicon.Name))
	return
}

func (cellint *CellInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellint.Name))
	return
}

func (cellstring *CellString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellstring.Name))
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplayedColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(displayedcolumn.Name))
	return
}

func (row *Row) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Row")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(row.Name))
	return
}

func (svgicon *SVGIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SVGIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgicon.Name))
	return
}

func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(table.Name))
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

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
