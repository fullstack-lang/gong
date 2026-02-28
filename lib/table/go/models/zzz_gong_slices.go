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

	// Compute reverse map for named struct CheckBox
	// insertion point per field

	// Compute reverse map for named struct DisplayedColumn
	// insertion point per field

	// Compute reverse map for named struct FormDiv
	// insertion point per field
	stage.FormDiv_FormFields_reverseMap = make(map[*FormField]*FormDiv)
	for formdiv := range stage.FormDivs {
		_ = formdiv
		for _, _formfield := range formdiv.FormFields {
			stage.FormDiv_FormFields_reverseMap[_formfield] = formdiv
		}
	}
	stage.FormDiv_CheckBoxs_reverseMap = make(map[*CheckBox]*FormDiv)
	for formdiv := range stage.FormDivs {
		_ = formdiv
		for _, _checkbox := range formdiv.CheckBoxs {
			stage.FormDiv_CheckBoxs_reverseMap[_checkbox] = formdiv
		}
	}

	// Compute reverse map for named struct FormEditAssocButton
	// insertion point per field

	// Compute reverse map for named struct FormField
	// insertion point per field

	// Compute reverse map for named struct FormFieldDate
	// insertion point per field

	// Compute reverse map for named struct FormFieldDateTime
	// insertion point per field

	// Compute reverse map for named struct FormFieldFloat64
	// insertion point per field

	// Compute reverse map for named struct FormFieldInt
	// insertion point per field

	// Compute reverse map for named struct FormFieldSelect
	// insertion point per field
	stage.FormFieldSelect_Options_reverseMap = make(map[*Option]*FormFieldSelect)
	for formfieldselect := range stage.FormFieldSelects {
		_ = formfieldselect
		for _, _option := range formfieldselect.Options {
			stage.FormFieldSelect_Options_reverseMap[_option] = formfieldselect
		}
	}

	// Compute reverse map for named struct FormFieldString
	// insertion point per field

	// Compute reverse map for named struct FormFieldTime
	// insertion point per field

	// Compute reverse map for named struct FormGroup
	// insertion point per field
	stage.FormGroup_FormDivs_reverseMap = make(map[*FormDiv]*FormGroup)
	for formgroup := range stage.FormGroups {
		_ = formgroup
		for _, _formdiv := range formgroup.FormDivs {
			stage.FormGroup_FormDivs_reverseMap[_formdiv] = formgroup
		}
	}

	// Compute reverse map for named struct FormSortAssocButton
	// insertion point per field

	// Compute reverse map for named struct Option
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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
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

	for instance := range stage.CheckBoxs {
		res = append(res, instance)
	}

	for instance := range stage.DisplayedColumns {
		res = append(res, instance)
	}

	for instance := range stage.FormDivs {
		res = append(res, instance)
	}

	for instance := range stage.FormEditAssocButtons {
		res = append(res, instance)
	}

	for instance := range stage.FormFields {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldDates {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldDateTimes {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldFloat64s {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldInts {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldSelects {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldStrings {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldTimes {
		res = append(res, instance)
	}

	for instance := range stage.FormGroups {
		res = append(res, instance)
	}

	for instance := range stage.FormSortAssocButtons {
		res = append(res, instance)
	}

	for instance := range stage.Options {
		res = append(res, instance)
	}

	for instance := range stage.Rows {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
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

func (checkbox *CheckBox) GongCopy() GongstructIF {
	newInstance := new(CheckBox)
	checkbox.CopyBasicFields(newInstance)
	return newInstance
}

func (displayedcolumn *DisplayedColumn) GongCopy() GongstructIF {
	newInstance := new(DisplayedColumn)
	displayedcolumn.CopyBasicFields(newInstance)
	return newInstance
}

func (formdiv *FormDiv) GongCopy() GongstructIF {
	newInstance := new(FormDiv)
	formdiv.CopyBasicFields(newInstance)
	return newInstance
}

func (formeditassocbutton *FormEditAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormEditAssocButton)
	formeditassocbutton.CopyBasicFields(newInstance)
	return newInstance
}

func (formfield *FormField) GongCopy() GongstructIF {
	newInstance := new(FormField)
	formfield.CopyBasicFields(newInstance)
	return newInstance
}

func (formfielddate *FormFieldDate) GongCopy() GongstructIF {
	newInstance := new(FormFieldDate)
	formfielddate.CopyBasicFields(newInstance)
	return newInstance
}

func (formfielddatetime *FormFieldDateTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldDateTime)
	formfielddatetime.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldfloat64 *FormFieldFloat64) GongCopy() GongstructIF {
	newInstance := new(FormFieldFloat64)
	formfieldfloat64.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldint *FormFieldInt) GongCopy() GongstructIF {
	newInstance := new(FormFieldInt)
	formfieldint.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldselect *FormFieldSelect) GongCopy() GongstructIF {
	newInstance := new(FormFieldSelect)
	formfieldselect.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldstring *FormFieldString) GongCopy() GongstructIF {
	newInstance := new(FormFieldString)
	formfieldstring.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldtime *FormFieldTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldTime)
	formfieldtime.CopyBasicFields(newInstance)
	return newInstance
}

func (formgroup *FormGroup) GongCopy() GongstructIF {
	newInstance := new(FormGroup)
	formgroup.CopyBasicFields(newInstance)
	return newInstance
}

func (formsortassocbutton *FormSortAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormSortAssocButton)
	formsortassocbutton.CopyBasicFields(newInstance)
	return newInstance
}

func (option *Option) GongCopy() GongstructIF {
	newInstance := new(Option)
	option.CopyBasicFields(newInstance)
	return newInstance
}

func (row *Row) GongCopy() GongstructIF {
	newInstance := new(Row)
	row.CopyBasicFields(newInstance)
	return newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := new(Table)
	table.CopyBasicFields(newInstance)
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
	var checkboxs_newInstances []*CheckBox
	var checkboxs_deletedInstances []*CheckBox

	// parse all staged instances and check if they have a reference
	for checkbox := range stage.CheckBoxs {
		if ref, ok := stage.CheckBoxs_reference[checkbox]; !ok {
			checkboxs_newInstances = append(checkboxs_newInstances, checkbox)
			newInstancesSlice = append(newInstancesSlice, checkbox.GongMarshallIdentifier(stage))
			if stage.CheckBoxs_referenceOrder == nil {
				stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint)
			}
			stage.CheckBoxs_referenceOrder[checkbox] = stage.CheckBox_stagedOrder[checkbox]
			newInstancesReverseSlice = append(newInstancesReverseSlice, checkbox.GongMarshallUnstaging(stage))
			// delete(stage.CheckBoxs_referenceOrder, checkbox)
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CheckBox_stagedOrder[ref] = stage.CheckBox_stagedOrder[checkbox]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := checkbox.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, checkbox)
			// delete(stage.CheckBox_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", checkbox.GetName())
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
	for _, ref := range stage.CheckBoxs_reference {
		instance := stage.CheckBoxs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CheckBoxs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(checkboxs_newInstances)
	lenDeletedInstances += len(checkboxs_deletedInstances)
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
	var formdivs_newInstances []*FormDiv
	var formdivs_deletedInstances []*FormDiv

	// parse all staged instances and check if they have a reference
	for formdiv := range stage.FormDivs {
		if ref, ok := stage.FormDivs_reference[formdiv]; !ok {
			formdivs_newInstances = append(formdivs_newInstances, formdiv)
			newInstancesSlice = append(newInstancesSlice, formdiv.GongMarshallIdentifier(stage))
			if stage.FormDivs_referenceOrder == nil {
				stage.FormDivs_referenceOrder = make(map[*FormDiv]uint)
			}
			stage.FormDivs_referenceOrder[formdiv] = stage.FormDiv_stagedOrder[formdiv]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formdiv.GongMarshallUnstaging(stage))
			// delete(stage.FormDivs_referenceOrder, formdiv)
			fieldInitializers, pointersInitializations := formdiv.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormDiv_stagedOrder[ref] = stage.FormDiv_stagedOrder[formdiv]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formdiv.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formdiv)
			// delete(stage.FormDiv_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formdiv.GetName())
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
	for _, ref := range stage.FormDivs_reference {
		instance := stage.FormDivs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormDivs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formdivs_deletedInstances = append(formdivs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formdivs_newInstances)
	lenDeletedInstances += len(formdivs_deletedInstances)
	var formeditassocbuttons_newInstances []*FormEditAssocButton
	var formeditassocbuttons_deletedInstances []*FormEditAssocButton

	// parse all staged instances and check if they have a reference
	for formeditassocbutton := range stage.FormEditAssocButtons {
		if ref, ok := stage.FormEditAssocButtons_reference[formeditassocbutton]; !ok {
			formeditassocbuttons_newInstances = append(formeditassocbuttons_newInstances, formeditassocbutton)
			newInstancesSlice = append(newInstancesSlice, formeditassocbutton.GongMarshallIdentifier(stage))
			if stage.FormEditAssocButtons_referenceOrder == nil {
				stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint)
			}
			stage.FormEditAssocButtons_referenceOrder[formeditassocbutton] = stage.FormEditAssocButton_stagedOrder[formeditassocbutton]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formeditassocbutton.GongMarshallUnstaging(stage))
			// delete(stage.FormEditAssocButtons_referenceOrder, formeditassocbutton)
			fieldInitializers, pointersInitializations := formeditassocbutton.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormEditAssocButton_stagedOrder[ref] = stage.FormEditAssocButton_stagedOrder[formeditassocbutton]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formeditassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formeditassocbutton)
			// delete(stage.FormEditAssocButton_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formeditassocbutton.GetName())
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
	for _, ref := range stage.FormEditAssocButtons_reference {
		instance := stage.FormEditAssocButtons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormEditAssocButtons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formeditassocbuttons_deletedInstances = append(formeditassocbuttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formeditassocbuttons_newInstances)
	lenDeletedInstances += len(formeditassocbuttons_deletedInstances)
	var formfields_newInstances []*FormField
	var formfields_deletedInstances []*FormField

	// parse all staged instances and check if they have a reference
	for formfield := range stage.FormFields {
		if ref, ok := stage.FormFields_reference[formfield]; !ok {
			formfields_newInstances = append(formfields_newInstances, formfield)
			newInstancesSlice = append(newInstancesSlice, formfield.GongMarshallIdentifier(stage))
			if stage.FormFields_referenceOrder == nil {
				stage.FormFields_referenceOrder = make(map[*FormField]uint)
			}
			stage.FormFields_referenceOrder[formfield] = stage.FormField_stagedOrder[formfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfield.GongMarshallUnstaging(stage))
			// delete(stage.FormFields_referenceOrder, formfield)
			fieldInitializers, pointersInitializations := formfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormField_stagedOrder[ref] = stage.FormField_stagedOrder[formfield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfield)
			// delete(stage.FormField_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfield.GetName())
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
	for _, ref := range stage.FormFields_reference {
		instance := stage.FormFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfields_deletedInstances = append(formfields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfields_newInstances)
	lenDeletedInstances += len(formfields_deletedInstances)
	var formfielddates_newInstances []*FormFieldDate
	var formfielddates_deletedInstances []*FormFieldDate

	// parse all staged instances and check if they have a reference
	for formfielddate := range stage.FormFieldDates {
		if ref, ok := stage.FormFieldDates_reference[formfielddate]; !ok {
			formfielddates_newInstances = append(formfielddates_newInstances, formfielddate)
			newInstancesSlice = append(newInstancesSlice, formfielddate.GongMarshallIdentifier(stage))
			if stage.FormFieldDates_referenceOrder == nil {
				stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint)
			}
			stage.FormFieldDates_referenceOrder[formfielddate] = stage.FormFieldDate_stagedOrder[formfielddate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfielddate.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldDates_referenceOrder, formfielddate)
			fieldInitializers, pointersInitializations := formfielddate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldDate_stagedOrder[ref] = stage.FormFieldDate_stagedOrder[formfielddate]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfielddate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddate)
			// delete(stage.FormFieldDate_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfielddate.GetName())
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
	for _, ref := range stage.FormFieldDates_reference {
		instance := stage.FormFieldDates_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldDates[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfielddates_deletedInstances = append(formfielddates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfielddates_newInstances)
	lenDeletedInstances += len(formfielddates_deletedInstances)
	var formfielddatetimes_newInstances []*FormFieldDateTime
	var formfielddatetimes_deletedInstances []*FormFieldDateTime

	// parse all staged instances and check if they have a reference
	for formfielddatetime := range stage.FormFieldDateTimes {
		if ref, ok := stage.FormFieldDateTimes_reference[formfielddatetime]; !ok {
			formfielddatetimes_newInstances = append(formfielddatetimes_newInstances, formfielddatetime)
			newInstancesSlice = append(newInstancesSlice, formfielddatetime.GongMarshallIdentifier(stage))
			if stage.FormFieldDateTimes_referenceOrder == nil {
				stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint)
			}
			stage.FormFieldDateTimes_referenceOrder[formfielddatetime] = stage.FormFieldDateTime_stagedOrder[formfielddatetime]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfielddatetime.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldDateTimes_referenceOrder, formfielddatetime)
			fieldInitializers, pointersInitializations := formfielddatetime.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldDateTime_stagedOrder[ref] = stage.FormFieldDateTime_stagedOrder[formfielddatetime]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfielddatetime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddatetime)
			// delete(stage.FormFieldDateTime_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfielddatetime.GetName())
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
	for _, ref := range stage.FormFieldDateTimes_reference {
		instance := stage.FormFieldDateTimes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldDateTimes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfielddatetimes_deletedInstances = append(formfielddatetimes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfielddatetimes_newInstances)
	lenDeletedInstances += len(formfielddatetimes_deletedInstances)
	var formfieldfloat64s_newInstances []*FormFieldFloat64
	var formfieldfloat64s_deletedInstances []*FormFieldFloat64

	// parse all staged instances and check if they have a reference
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		if ref, ok := stage.FormFieldFloat64s_reference[formfieldfloat64]; !ok {
			formfieldfloat64s_newInstances = append(formfieldfloat64s_newInstances, formfieldfloat64)
			newInstancesSlice = append(newInstancesSlice, formfieldfloat64.GongMarshallIdentifier(stage))
			if stage.FormFieldFloat64s_referenceOrder == nil {
				stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint)
			}
			stage.FormFieldFloat64s_referenceOrder[formfieldfloat64] = stage.FormFieldFloat64_stagedOrder[formfieldfloat64]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldfloat64.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldFloat64s_referenceOrder, formfieldfloat64)
			fieldInitializers, pointersInitializations := formfieldfloat64.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldFloat64_stagedOrder[ref] = stage.FormFieldFloat64_stagedOrder[formfieldfloat64]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldfloat64.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldfloat64)
			// delete(stage.FormFieldFloat64_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldfloat64.GetName())
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
	for _, ref := range stage.FormFieldFloat64s_reference {
		instance := stage.FormFieldFloat64s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldFloat64s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldfloat64s_deletedInstances = append(formfieldfloat64s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldfloat64s_newInstances)
	lenDeletedInstances += len(formfieldfloat64s_deletedInstances)
	var formfieldints_newInstances []*FormFieldInt
	var formfieldints_deletedInstances []*FormFieldInt

	// parse all staged instances and check if they have a reference
	for formfieldint := range stage.FormFieldInts {
		if ref, ok := stage.FormFieldInts_reference[formfieldint]; !ok {
			formfieldints_newInstances = append(formfieldints_newInstances, formfieldint)
			newInstancesSlice = append(newInstancesSlice, formfieldint.GongMarshallIdentifier(stage))
			if stage.FormFieldInts_referenceOrder == nil {
				stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint)
			}
			stage.FormFieldInts_referenceOrder[formfieldint] = stage.FormFieldInt_stagedOrder[formfieldint]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldint.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldInts_referenceOrder, formfieldint)
			fieldInitializers, pointersInitializations := formfieldint.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldInt_stagedOrder[ref] = stage.FormFieldInt_stagedOrder[formfieldint]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldint)
			// delete(stage.FormFieldInt_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldint.GetName())
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
	for _, ref := range stage.FormFieldInts_reference {
		instance := stage.FormFieldInts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldInts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldints_deletedInstances = append(formfieldints_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldints_newInstances)
	lenDeletedInstances += len(formfieldints_deletedInstances)
	var formfieldselects_newInstances []*FormFieldSelect
	var formfieldselects_deletedInstances []*FormFieldSelect

	// parse all staged instances and check if they have a reference
	for formfieldselect := range stage.FormFieldSelects {
		if ref, ok := stage.FormFieldSelects_reference[formfieldselect]; !ok {
			formfieldselects_newInstances = append(formfieldselects_newInstances, formfieldselect)
			newInstancesSlice = append(newInstancesSlice, formfieldselect.GongMarshallIdentifier(stage))
			if stage.FormFieldSelects_referenceOrder == nil {
				stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint)
			}
			stage.FormFieldSelects_referenceOrder[formfieldselect] = stage.FormFieldSelect_stagedOrder[formfieldselect]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldselect.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldSelects_referenceOrder, formfieldselect)
			fieldInitializers, pointersInitializations := formfieldselect.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldSelect_stagedOrder[ref] = stage.FormFieldSelect_stagedOrder[formfieldselect]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldselect.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldselect)
			// delete(stage.FormFieldSelect_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldselect.GetName())
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
	for _, ref := range stage.FormFieldSelects_reference {
		instance := stage.FormFieldSelects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldSelects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldselects_deletedInstances = append(formfieldselects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldselects_newInstances)
	lenDeletedInstances += len(formfieldselects_deletedInstances)
	var formfieldstrings_newInstances []*FormFieldString
	var formfieldstrings_deletedInstances []*FormFieldString

	// parse all staged instances and check if they have a reference
	for formfieldstring := range stage.FormFieldStrings {
		if ref, ok := stage.FormFieldStrings_reference[formfieldstring]; !ok {
			formfieldstrings_newInstances = append(formfieldstrings_newInstances, formfieldstring)
			newInstancesSlice = append(newInstancesSlice, formfieldstring.GongMarshallIdentifier(stage))
			if stage.FormFieldStrings_referenceOrder == nil {
				stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint)
			}
			stage.FormFieldStrings_referenceOrder[formfieldstring] = stage.FormFieldString_stagedOrder[formfieldstring]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldstring.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldStrings_referenceOrder, formfieldstring)
			fieldInitializers, pointersInitializations := formfieldstring.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldString_stagedOrder[ref] = stage.FormFieldString_stagedOrder[formfieldstring]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldstring.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldstring)
			// delete(stage.FormFieldString_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldstring.GetName())
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
	for _, ref := range stage.FormFieldStrings_reference {
		instance := stage.FormFieldStrings_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldStrings[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldstrings_deletedInstances = append(formfieldstrings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldstrings_newInstances)
	lenDeletedInstances += len(formfieldstrings_deletedInstances)
	var formfieldtimes_newInstances []*FormFieldTime
	var formfieldtimes_deletedInstances []*FormFieldTime

	// parse all staged instances and check if they have a reference
	for formfieldtime := range stage.FormFieldTimes {
		if ref, ok := stage.FormFieldTimes_reference[formfieldtime]; !ok {
			formfieldtimes_newInstances = append(formfieldtimes_newInstances, formfieldtime)
			newInstancesSlice = append(newInstancesSlice, formfieldtime.GongMarshallIdentifier(stage))
			if stage.FormFieldTimes_referenceOrder == nil {
				stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint)
			}
			stage.FormFieldTimes_referenceOrder[formfieldtime] = stage.FormFieldTime_stagedOrder[formfieldtime]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldtime.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldTimes_referenceOrder, formfieldtime)
			fieldInitializers, pointersInitializations := formfieldtime.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldTime_stagedOrder[ref] = stage.FormFieldTime_stagedOrder[formfieldtime]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldtime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldtime)
			// delete(stage.FormFieldTime_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldtime.GetName())
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
	for _, ref := range stage.FormFieldTimes_reference {
		instance := stage.FormFieldTimes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldTimes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldtimes_deletedInstances = append(formfieldtimes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldtimes_newInstances)
	lenDeletedInstances += len(formfieldtimes_deletedInstances)
	var formgroups_newInstances []*FormGroup
	var formgroups_deletedInstances []*FormGroup

	// parse all staged instances and check if they have a reference
	for formgroup := range stage.FormGroups {
		if ref, ok := stage.FormGroups_reference[formgroup]; !ok {
			formgroups_newInstances = append(formgroups_newInstances, formgroup)
			newInstancesSlice = append(newInstancesSlice, formgroup.GongMarshallIdentifier(stage))
			if stage.FormGroups_referenceOrder == nil {
				stage.FormGroups_referenceOrder = make(map[*FormGroup]uint)
			}
			stage.FormGroups_referenceOrder[formgroup] = stage.FormGroup_stagedOrder[formgroup]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formgroup.GongMarshallUnstaging(stage))
			// delete(stage.FormGroups_referenceOrder, formgroup)
			fieldInitializers, pointersInitializations := formgroup.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormGroup_stagedOrder[ref] = stage.FormGroup_stagedOrder[formgroup]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formgroup.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formgroup)
			// delete(stage.FormGroup_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formgroup.GetName())
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
	for _, ref := range stage.FormGroups_reference {
		instance := stage.FormGroups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormGroups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formgroups_deletedInstances = append(formgroups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formgroups_newInstances)
	lenDeletedInstances += len(formgroups_deletedInstances)
	var formsortassocbuttons_newInstances []*FormSortAssocButton
	var formsortassocbuttons_deletedInstances []*FormSortAssocButton

	// parse all staged instances and check if they have a reference
	for formsortassocbutton := range stage.FormSortAssocButtons {
		if ref, ok := stage.FormSortAssocButtons_reference[formsortassocbutton]; !ok {
			formsortassocbuttons_newInstances = append(formsortassocbuttons_newInstances, formsortassocbutton)
			newInstancesSlice = append(newInstancesSlice, formsortassocbutton.GongMarshallIdentifier(stage))
			if stage.FormSortAssocButtons_referenceOrder == nil {
				stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint)
			}
			stage.FormSortAssocButtons_referenceOrder[formsortassocbutton] = stage.FormSortAssocButton_stagedOrder[formsortassocbutton]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formsortassocbutton.GongMarshallUnstaging(stage))
			// delete(stage.FormSortAssocButtons_referenceOrder, formsortassocbutton)
			fieldInitializers, pointersInitializations := formsortassocbutton.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormSortAssocButton_stagedOrder[ref] = stage.FormSortAssocButton_stagedOrder[formsortassocbutton]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formsortassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formsortassocbutton)
			// delete(stage.FormSortAssocButton_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formsortassocbutton.GetName())
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
	for _, ref := range stage.FormSortAssocButtons_reference {
		instance := stage.FormSortAssocButtons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormSortAssocButtons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formsortassocbuttons_deletedInstances = append(formsortassocbuttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formsortassocbuttons_newInstances)
	lenDeletedInstances += len(formsortassocbuttons_deletedInstances)
	var options_newInstances []*Option
	var options_deletedInstances []*Option

	// parse all staged instances and check if they have a reference
	for option := range stage.Options {
		if ref, ok := stage.Options_reference[option]; !ok {
			options_newInstances = append(options_newInstances, option)
			newInstancesSlice = append(newInstancesSlice, option.GongMarshallIdentifier(stage))
			if stage.Options_referenceOrder == nil {
				stage.Options_referenceOrder = make(map[*Option]uint)
			}
			stage.Options_referenceOrder[option] = stage.Option_stagedOrder[option]
			newInstancesReverseSlice = append(newInstancesReverseSlice, option.GongMarshallUnstaging(stage))
			// delete(stage.Options_referenceOrder, option)
			fieldInitializers, pointersInitializations := option.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Option_stagedOrder[ref] = stage.Option_stagedOrder[option]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := option.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, option)
			// delete(stage.Option_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", option.GetName())
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
	for _, ref := range stage.Options_reference {
		instance := stage.Options_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Options[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			options_deletedInstances = append(options_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(options_newInstances)
	lenDeletedInstances += len(options_deletedInstances)
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
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
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

	stage.CheckBoxs_reference = make(map[*CheckBox]*CheckBox)
	stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint) // diff Unstage needs the reference order
	stage.CheckBoxs_instance = make(map[*CheckBox]*CheckBox)
	for instance := range stage.CheckBoxs {
		_copy := instance.GongCopy().(*CheckBox)
		stage.CheckBoxs_reference[instance] = _copy
		stage.CheckBoxs_instance[_copy] = instance
		stage.CheckBoxs_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.FormDivs_reference = make(map[*FormDiv]*FormDiv)
	stage.FormDivs_referenceOrder = make(map[*FormDiv]uint) // diff Unstage needs the reference order
	stage.FormDivs_instance = make(map[*FormDiv]*FormDiv)
	for instance := range stage.FormDivs {
		_copy := instance.GongCopy().(*FormDiv)
		stage.FormDivs_reference[instance] = _copy
		stage.FormDivs_instance[_copy] = instance
		stage.FormDivs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormEditAssocButtons_reference = make(map[*FormEditAssocButton]*FormEditAssocButton)
	stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint) // diff Unstage needs the reference order
	stage.FormEditAssocButtons_instance = make(map[*FormEditAssocButton]*FormEditAssocButton)
	for instance := range stage.FormEditAssocButtons {
		_copy := instance.GongCopy().(*FormEditAssocButton)
		stage.FormEditAssocButtons_reference[instance] = _copy
		stage.FormEditAssocButtons_instance[_copy] = instance
		stage.FormEditAssocButtons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFields_reference = make(map[*FormField]*FormField)
	stage.FormFields_referenceOrder = make(map[*FormField]uint) // diff Unstage needs the reference order
	stage.FormFields_instance = make(map[*FormField]*FormField)
	for instance := range stage.FormFields {
		_copy := instance.GongCopy().(*FormField)
		stage.FormFields_reference[instance] = _copy
		stage.FormFields_instance[_copy] = instance
		stage.FormFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDates_reference = make(map[*FormFieldDate]*FormFieldDate)
	stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint) // diff Unstage needs the reference order
	stage.FormFieldDates_instance = make(map[*FormFieldDate]*FormFieldDate)
	for instance := range stage.FormFieldDates {
		_copy := instance.GongCopy().(*FormFieldDate)
		stage.FormFieldDates_reference[instance] = _copy
		stage.FormFieldDates_instance[_copy] = instance
		stage.FormFieldDates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDateTimes_reference = make(map[*FormFieldDateTime]*FormFieldDateTime)
	stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint) // diff Unstage needs the reference order
	stage.FormFieldDateTimes_instance = make(map[*FormFieldDateTime]*FormFieldDateTime)
	for instance := range stage.FormFieldDateTimes {
		_copy := instance.GongCopy().(*FormFieldDateTime)
		stage.FormFieldDateTimes_reference[instance] = _copy
		stage.FormFieldDateTimes_instance[_copy] = instance
		stage.FormFieldDateTimes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldFloat64s_reference = make(map[*FormFieldFloat64]*FormFieldFloat64)
	stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint) // diff Unstage needs the reference order
	stage.FormFieldFloat64s_instance = make(map[*FormFieldFloat64]*FormFieldFloat64)
	for instance := range stage.FormFieldFloat64s {
		_copy := instance.GongCopy().(*FormFieldFloat64)
		stage.FormFieldFloat64s_reference[instance] = _copy
		stage.FormFieldFloat64s_instance[_copy] = instance
		stage.FormFieldFloat64s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldInts_reference = make(map[*FormFieldInt]*FormFieldInt)
	stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint) // diff Unstage needs the reference order
	stage.FormFieldInts_instance = make(map[*FormFieldInt]*FormFieldInt)
	for instance := range stage.FormFieldInts {
		_copy := instance.GongCopy().(*FormFieldInt)
		stage.FormFieldInts_reference[instance] = _copy
		stage.FormFieldInts_instance[_copy] = instance
		stage.FormFieldInts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldSelects_reference = make(map[*FormFieldSelect]*FormFieldSelect)
	stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint) // diff Unstage needs the reference order
	stage.FormFieldSelects_instance = make(map[*FormFieldSelect]*FormFieldSelect)
	for instance := range stage.FormFieldSelects {
		_copy := instance.GongCopy().(*FormFieldSelect)
		stage.FormFieldSelects_reference[instance] = _copy
		stage.FormFieldSelects_instance[_copy] = instance
		stage.FormFieldSelects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldStrings_reference = make(map[*FormFieldString]*FormFieldString)
	stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint) // diff Unstage needs the reference order
	stage.FormFieldStrings_instance = make(map[*FormFieldString]*FormFieldString)
	for instance := range stage.FormFieldStrings {
		_copy := instance.GongCopy().(*FormFieldString)
		stage.FormFieldStrings_reference[instance] = _copy
		stage.FormFieldStrings_instance[_copy] = instance
		stage.FormFieldStrings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldTimes_reference = make(map[*FormFieldTime]*FormFieldTime)
	stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint) // diff Unstage needs the reference order
	stage.FormFieldTimes_instance = make(map[*FormFieldTime]*FormFieldTime)
	for instance := range stage.FormFieldTimes {
		_copy := instance.GongCopy().(*FormFieldTime)
		stage.FormFieldTimes_reference[instance] = _copy
		stage.FormFieldTimes_instance[_copy] = instance
		stage.FormFieldTimes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormGroups_reference = make(map[*FormGroup]*FormGroup)
	stage.FormGroups_referenceOrder = make(map[*FormGroup]uint) // diff Unstage needs the reference order
	stage.FormGroups_instance = make(map[*FormGroup]*FormGroup)
	for instance := range stage.FormGroups {
		_copy := instance.GongCopy().(*FormGroup)
		stage.FormGroups_reference[instance] = _copy
		stage.FormGroups_instance[_copy] = instance
		stage.FormGroups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormSortAssocButtons_reference = make(map[*FormSortAssocButton]*FormSortAssocButton)
	stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint) // diff Unstage needs the reference order
	stage.FormSortAssocButtons_instance = make(map[*FormSortAssocButton]*FormSortAssocButton)
	for instance := range stage.FormSortAssocButtons {
		_copy := instance.GongCopy().(*FormSortAssocButton)
		stage.FormSortAssocButtons_reference[instance] = _copy
		stage.FormSortAssocButtons_instance[_copy] = instance
		stage.FormSortAssocButtons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Options_reference = make(map[*Option]*Option)
	stage.Options_referenceOrder = make(map[*Option]uint) // diff Unstage needs the reference order
	stage.Options_instance = make(map[*Option]*Option)
	for instance := range stage.Options {
		_copy := instance.GongCopy().(*Option)
		stage.Options_reference[instance] = _copy
		stage.Options_instance[_copy] = instance
		stage.Options_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.CheckBoxs {
		reference := stage.CheckBoxs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DisplayedColumns {
		reference := stage.DisplayedColumns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormDivs {
		reference := stage.FormDivs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormEditAssocButtons {
		reference := stage.FormEditAssocButtons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFields {
		reference := stage.FormFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldDates {
		reference := stage.FormFieldDates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldDateTimes {
		reference := stage.FormFieldDateTimes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldFloat64s {
		reference := stage.FormFieldFloat64s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldInts {
		reference := stage.FormFieldInts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldSelects {
		reference := stage.FormFieldSelects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldStrings {
		reference := stage.FormFieldStrings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldTimes {
		reference := stage.FormFieldTimes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormGroups {
		reference := stage.FormGroups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormSortAssocButtons {
		reference := stage.FormSortAssocButtons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Options {
		reference := stage.Options_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rows {
		reference := stage.Rows_reference[instance]
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

func (checkbox *CheckBox) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CheckBox_stagedOrder[checkbox]; ok {
		return order
	}
	if order, ok := stage.CheckBoxs_referenceOrder[checkbox]; ok {
		return order
	} else {
		log.Printf("instance %p of type CheckBox was not staged and does not have a reference order", checkbox)
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

func (formdiv *FormDiv) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormDiv_stagedOrder[formdiv]; ok {
		return order
	}
	if order, ok := stage.FormDivs_referenceOrder[formdiv]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormDiv was not staged and does not have a reference order", formdiv)
		return 0
	}
}

func (formeditassocbutton *FormEditAssocButton) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormEditAssocButton_stagedOrder[formeditassocbutton]; ok {
		return order
	}
	if order, ok := stage.FormEditAssocButtons_referenceOrder[formeditassocbutton]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormEditAssocButton was not staged and does not have a reference order", formeditassocbutton)
		return 0
	}
}

func (formfield *FormField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormField_stagedOrder[formfield]; ok {
		return order
	}
	if order, ok := stage.FormFields_referenceOrder[formfield]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormField was not staged and does not have a reference order", formfield)
		return 0
	}
}

func (formfielddate *FormFieldDate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldDate_stagedOrder[formfielddate]; ok {
		return order
	}
	if order, ok := stage.FormFieldDates_referenceOrder[formfielddate]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldDate was not staged and does not have a reference order", formfielddate)
		return 0
	}
}

func (formfielddatetime *FormFieldDateTime) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldDateTime_stagedOrder[formfielddatetime]; ok {
		return order
	}
	if order, ok := stage.FormFieldDateTimes_referenceOrder[formfielddatetime]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldDateTime was not staged and does not have a reference order", formfielddatetime)
		return 0
	}
}

func (formfieldfloat64 *FormFieldFloat64) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldFloat64_stagedOrder[formfieldfloat64]; ok {
		return order
	}
	if order, ok := stage.FormFieldFloat64s_referenceOrder[formfieldfloat64]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldFloat64 was not staged and does not have a reference order", formfieldfloat64)
		return 0
	}
}

func (formfieldint *FormFieldInt) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldInt_stagedOrder[formfieldint]; ok {
		return order
	}
	if order, ok := stage.FormFieldInts_referenceOrder[formfieldint]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldInt was not staged and does not have a reference order", formfieldint)
		return 0
	}
}

func (formfieldselect *FormFieldSelect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldSelect_stagedOrder[formfieldselect]; ok {
		return order
	}
	if order, ok := stage.FormFieldSelects_referenceOrder[formfieldselect]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldSelect was not staged and does not have a reference order", formfieldselect)
		return 0
	}
}

func (formfieldstring *FormFieldString) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldString_stagedOrder[formfieldstring]; ok {
		return order
	}
	if order, ok := stage.FormFieldStrings_referenceOrder[formfieldstring]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldString was not staged and does not have a reference order", formfieldstring)
		return 0
	}
}

func (formfieldtime *FormFieldTime) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldTime_stagedOrder[formfieldtime]; ok {
		return order
	}
	if order, ok := stage.FormFieldTimes_referenceOrder[formfieldtime]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldTime was not staged and does not have a reference order", formfieldtime)
		return 0
	}
}

func (formgroup *FormGroup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormGroup_stagedOrder[formgroup]; ok {
		return order
	}
	if order, ok := stage.FormGroups_referenceOrder[formgroup]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormGroup was not staged and does not have a reference order", formgroup)
		return 0
	}
}

func (formsortassocbutton *FormSortAssocButton) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormSortAssocButton_stagedOrder[formsortassocbutton]; ok {
		return order
	}
	if order, ok := stage.FormSortAssocButtons_referenceOrder[formsortassocbutton]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormSortAssocButton was not staged and does not have a reference order", formsortassocbutton)
		return 0
	}
}

func (option *Option) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Option_stagedOrder[option]; ok {
		return order
	}
	if order, ok := stage.Options_referenceOrder[option]; ok {
		return order
	} else {
		log.Printf("instance %p of type Option was not staged and does not have a reference order", option)
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

func (checkbox *CheckBox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkbox *CheckBox) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

func (displayedcolumn *DisplayedColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (displayedcolumn *DisplayedColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

func (formdiv *FormDiv) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formdiv *FormDiv) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

func (formeditassocbutton *FormEditAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formeditassocbutton *FormEditAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

func (formfield *FormField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfield *FormField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

func (formfielddate *FormFieldDate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddate *FormFieldDate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

func (formfielddatetime *FormFieldDateTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddatetime *FormFieldDateTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

func (formfieldfloat64 *FormFieldFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldfloat64 *FormFieldFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

func (formfieldint *FormFieldInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldint *FormFieldInt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

func (formfieldselect *FormFieldSelect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldselect *FormFieldSelect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

func (formfieldstring *FormFieldString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldstring *FormFieldString) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

func (formfieldtime *FormFieldTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldtime *FormFieldTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

func (formgroup *FormGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formgroup *FormGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

func (formsortassocbutton *FormSortAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formsortassocbutton *FormSortAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

func (option *Option) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (option *Option) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

func (row *Row) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (row *Row) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
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

func (checkbox *CheckBox) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CheckBox")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(checkbox.Name))
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplayedColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(displayedcolumn.Name))
	return
}

func (formdiv *FormDiv) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormDiv")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formdiv.Name))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormEditAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.Name))
	return
}

func (formfield *FormField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfield.Name))
	return
}

func (formfielddate *FormFieldDate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddate.Name))
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDateTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddatetime.Name))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldfloat64.Name))
	return
}

func (formfieldint *FormFieldInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldint.Name))
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldSelect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldselect.Name))
	return
}

func (formfieldstring *FormFieldString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldstring.Name))
	return
}

func (formfieldtime *FormFieldTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldtime.Name))
	return
}

func (formgroup *FormGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formgroup.Name))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormSortAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.Name))
	return
}

func (option *Option) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Option")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(option.Name))
	return
}

func (row *Row) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Row")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(row.Name))
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

func (checkbox *CheckBox) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetReferenceIdentifier(stage))
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetReferenceIdentifier(stage))
	return
}

func (formdiv *FormDiv) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetReferenceIdentifier(stage))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetReferenceIdentifier(stage))
	return
}

func (formfield *FormField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetReferenceIdentifier(stage))
	return
}

func (formfielddate *FormFieldDate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetReferenceIdentifier(stage))
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldint *FormFieldInt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldstring *FormFieldString) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldtime *FormFieldTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetReferenceIdentifier(stage))
	return
}

func (formgroup *FormGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetReferenceIdentifier(stage))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetReferenceIdentifier(stage))
	return
}

func (option *Option) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetReferenceIdentifier(stage))
	return
}

func (row *Row) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetReferenceIdentifier(stage))
	return
}

func (table *Table) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetReferenceIdentifier(stage))
	return
}

// end of template
