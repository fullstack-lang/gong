// generated code - do not edit
package models

import (
	"fmt"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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
	newInstance := *cell
	return &newInstance
}

func (cellboolean *CellBoolean) GongCopy() GongstructIF {
	newInstance := *cellboolean
	return &newInstance
}

func (cellfloat64 *CellFloat64) GongCopy() GongstructIF {
	newInstance := *cellfloat64
	return &newInstance
}

func (cellicon *CellIcon) GongCopy() GongstructIF {
	newInstance := *cellicon
	return &newInstance
}

func (cellint *CellInt) GongCopy() GongstructIF {
	newInstance := *cellint
	return &newInstance
}

func (cellstring *CellString) GongCopy() GongstructIF {
	newInstance := *cellstring
	return &newInstance
}

func (checkbox *CheckBox) GongCopy() GongstructIF {
	newInstance := *checkbox
	return &newInstance
}

func (displayedcolumn *DisplayedColumn) GongCopy() GongstructIF {
	newInstance := *displayedcolumn
	return &newInstance
}

func (formdiv *FormDiv) GongCopy() GongstructIF {
	newInstance := *formdiv
	return &newInstance
}

func (formeditassocbutton *FormEditAssocButton) GongCopy() GongstructIF {
	newInstance := *formeditassocbutton
	return &newInstance
}

func (formfield *FormField) GongCopy() GongstructIF {
	newInstance := *formfield
	return &newInstance
}

func (formfielddate *FormFieldDate) GongCopy() GongstructIF {
	newInstance := *formfielddate
	return &newInstance
}

func (formfielddatetime *FormFieldDateTime) GongCopy() GongstructIF {
	newInstance := *formfielddatetime
	return &newInstance
}

func (formfieldfloat64 *FormFieldFloat64) GongCopy() GongstructIF {
	newInstance := *formfieldfloat64
	return &newInstance
}

func (formfieldint *FormFieldInt) GongCopy() GongstructIF {
	newInstance := *formfieldint
	return &newInstance
}

func (formfieldselect *FormFieldSelect) GongCopy() GongstructIF {
	newInstance := *formfieldselect
	return &newInstance
}

func (formfieldstring *FormFieldString) GongCopy() GongstructIF {
	newInstance := *formfieldstring
	return &newInstance
}

func (formfieldtime *FormFieldTime) GongCopy() GongstructIF {
	newInstance := *formfieldtime
	return &newInstance
}

func (formgroup *FormGroup) GongCopy() GongstructIF {
	newInstance := *formgroup
	return &newInstance
}

func (formsortassocbutton *FormSortAssocButton) GongCopy() GongstructIF {
	newInstance := *formsortassocbutton
	return &newInstance
}

func (option *Option) GongCopy() GongstructIF {
	newInstance := *option
	return &newInstance
}

func (row *Row) GongCopy() GongstructIF {
	newInstance := *row
	return &newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := *table
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

	var newInstancesReverseStmt string
	_ = newInstancesReverseStmt
	var fieldsEditReverseStmt string
	_ = fieldsEditReverseStmt
	var deletedInstancesReverseStmt string
	_ = deletedInstancesReverseStmt

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
			newInstancesStmt += cell.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cell.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cell.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cell.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cell)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cell.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Cells_reference {
		if _, ok := stage.Cells[ref]; !ok {
			cells_deletedInstances = append(cells_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += cellboolean.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cellboolean.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cellboolean.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellboolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellboolean)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cellboolean.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CellBooleans_reference {
		if _, ok := stage.CellBooleans[ref]; !ok {
			cellbooleans_deletedInstances = append(cellbooleans_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += cellfloat64.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cellfloat64.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cellfloat64.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellfloat64.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellfloat64)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cellfloat64.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CellFloat64s_reference {
		if _, ok := stage.CellFloat64s[ref]; !ok {
			cellfloat64s_deletedInstances = append(cellfloat64s_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += cellicon.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cellicon.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cellicon.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellicon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellicon)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cellicon.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CellIcons_reference {
		if _, ok := stage.CellIcons[ref]; !ok {
			cellicons_deletedInstances = append(cellicons_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += cellint.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cellint.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cellint.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellint)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cellint.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CellInts_reference {
		if _, ok := stage.CellInts[ref]; !ok {
			cellints_deletedInstances = append(cellints_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += cellstring.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += cellstring.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := cellstring.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellstring.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cellstring)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", cellstring.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CellStrings_reference {
		if _, ok := stage.CellStrings[ref]; !ok {
			cellstrings_deletedInstances = append(cellstrings_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += checkbox.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += checkbox.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := checkbox.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, checkbox)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", checkbox.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.CheckBoxs_reference {
		if _, ok := stage.CheckBoxs[ref]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += displayedcolumn.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += displayedcolumn.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := displayedcolumn.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := displayedcolumn.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, displayedcolumn)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", displayedcolumn.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DisplayedColumns_reference {
		if _, ok := stage.DisplayedColumns[ref]; !ok {
			displayedcolumns_deletedInstances = append(displayedcolumns_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formdiv.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formdiv.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formdiv.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formdiv.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formdiv)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formdiv.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormDivs_reference {
		if _, ok := stage.FormDivs[ref]; !ok {
			formdivs_deletedInstances = append(formdivs_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formeditassocbutton.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formeditassocbutton.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formeditassocbutton.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formeditassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formeditassocbutton)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formeditassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormEditAssocButtons_reference {
		if _, ok := stage.FormEditAssocButtons[ref]; !ok {
			formeditassocbuttons_deletedInstances = append(formeditassocbuttons_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfield.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfield.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfield)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFields_reference {
		if _, ok := stage.FormFields[ref]; !ok {
			formfields_deletedInstances = append(formfields_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfielddate.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfielddate.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfielddate.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfielddate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddate)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfielddate.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldDates_reference {
		if _, ok := stage.FormFieldDates[ref]; !ok {
			formfielddates_deletedInstances = append(formfielddates_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfielddatetime.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfielddatetime.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfielddatetime.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfielddatetime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddatetime)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfielddatetime.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldDateTimes_reference {
		if _, ok := stage.FormFieldDateTimes[ref]; !ok {
			formfielddatetimes_deletedInstances = append(formfielddatetimes_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfieldfloat64.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfieldfloat64.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfieldfloat64.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldfloat64.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldfloat64)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfieldfloat64.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldFloat64s_reference {
		if _, ok := stage.FormFieldFloat64s[ref]; !ok {
			formfieldfloat64s_deletedInstances = append(formfieldfloat64s_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfieldint.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfieldint.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfieldint.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldint)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfieldint.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldInts_reference {
		if _, ok := stage.FormFieldInts[ref]; !ok {
			formfieldints_deletedInstances = append(formfieldints_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfieldselect.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfieldselect.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfieldselect.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldselect.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldselect)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfieldselect.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldSelects_reference {
		if _, ok := stage.FormFieldSelects[ref]; !ok {
			formfieldselects_deletedInstances = append(formfieldselects_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfieldstring.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfieldstring.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfieldstring.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldstring.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldstring)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfieldstring.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldStrings_reference {
		if _, ok := stage.FormFieldStrings[ref]; !ok {
			formfieldstrings_deletedInstances = append(formfieldstrings_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formfieldtime.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formfieldtime.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formfieldtime.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldtime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldtime)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formfieldtime.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormFieldTimes_reference {
		if _, ok := stage.FormFieldTimes[ref]; !ok {
			formfieldtimes_deletedInstances = append(formfieldtimes_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formgroup.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formgroup.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formgroup.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formgroup.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formgroup)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formgroup.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormGroups_reference {
		if _, ok := stage.FormGroups[ref]; !ok {
			formgroups_deletedInstances = append(formgroups_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += formsortassocbutton.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += formsortassocbutton.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := formsortassocbutton.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formsortassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formsortassocbutton)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", formsortassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.FormSortAssocButtons_reference {
		if _, ok := stage.FormSortAssocButtons[ref]; !ok {
			formsortassocbuttons_deletedInstances = append(formsortassocbuttons_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += option.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += option.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := option.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := option.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, option)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", option.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Options_reference {
		if _, ok := stage.Options[ref]; !ok {
			options_deletedInstances = append(options_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += row.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += row.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := row.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := row.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, row)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", row.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Rows_reference {
		if _, ok := stage.Rows[ref]; !ok {
			rows_deletedInstances = append(rows_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
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
			newInstancesStmt += table.GongMarshallIdentifier(stage)
			newInstancesReverseStmt += table.GongMarshallUnstaging(stage)
			fieldInitializers, pointersInitializations := table.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := table.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, table)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// %s", table.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseStmt += reverseDiff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Tables_reference {
		if _, ok := stage.Tables[ref]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, ref)
			deletedInstancesStmt += ref.GongMarshallUnstaging(stage)
			deletedInstancesReverseStmt += ref.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseStmt += fieldInitializers
			fieldsEditReverseStmt += pointersInitializations
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		forwardCommit += "\n\tstage.Commit()\n"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		backwardCommit += "\n\tstage.Commit()\n"
		// append to the start of the backward commits slice
		stage.backwardCommits = append([]string{backwardCommit}, stage.backwardCommits...)

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
	stage.Cells_reference = make(map[*Cell]*Cell)
	stage.Cells_referenceOrder = make(map[*Cell]uint) // diff Unstage needs the reference order
	for instance := range stage.Cells {
		stage.Cells_reference[instance] = instance.GongCopy().(*Cell)
		stage.Cells_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CellBooleans_reference = make(map[*CellBoolean]*CellBoolean)
	stage.CellBooleans_referenceOrder = make(map[*CellBoolean]uint) // diff Unstage needs the reference order
	for instance := range stage.CellBooleans {
		stage.CellBooleans_reference[instance] = instance.GongCopy().(*CellBoolean)
		stage.CellBooleans_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CellFloat64s_reference = make(map[*CellFloat64]*CellFloat64)
	stage.CellFloat64s_referenceOrder = make(map[*CellFloat64]uint) // diff Unstage needs the reference order
	for instance := range stage.CellFloat64s {
		stage.CellFloat64s_reference[instance] = instance.GongCopy().(*CellFloat64)
		stage.CellFloat64s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CellIcons_reference = make(map[*CellIcon]*CellIcon)
	stage.CellIcons_referenceOrder = make(map[*CellIcon]uint) // diff Unstage needs the reference order
	for instance := range stage.CellIcons {
		stage.CellIcons_reference[instance] = instance.GongCopy().(*CellIcon)
		stage.CellIcons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CellInts_reference = make(map[*CellInt]*CellInt)
	stage.CellInts_referenceOrder = make(map[*CellInt]uint) // diff Unstage needs the reference order
	for instance := range stage.CellInts {
		stage.CellInts_reference[instance] = instance.GongCopy().(*CellInt)
		stage.CellInts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CellStrings_reference = make(map[*CellString]*CellString)
	stage.CellStrings_referenceOrder = make(map[*CellString]uint) // diff Unstage needs the reference order
	for instance := range stage.CellStrings {
		stage.CellStrings_reference[instance] = instance.GongCopy().(*CellString)
		stage.CellStrings_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.CheckBoxs_reference = make(map[*CheckBox]*CheckBox)
	stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint) // diff Unstage needs the reference order
	for instance := range stage.CheckBoxs {
		stage.CheckBoxs_reference[instance] = instance.GongCopy().(*CheckBox)
		stage.CheckBoxs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DisplayedColumns_reference = make(map[*DisplayedColumn]*DisplayedColumn)
	stage.DisplayedColumns_referenceOrder = make(map[*DisplayedColumn]uint) // diff Unstage needs the reference order
	for instance := range stage.DisplayedColumns {
		stage.DisplayedColumns_reference[instance] = instance.GongCopy().(*DisplayedColumn)
		stage.DisplayedColumns_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormDivs_reference = make(map[*FormDiv]*FormDiv)
	stage.FormDivs_referenceOrder = make(map[*FormDiv]uint) // diff Unstage needs the reference order
	for instance := range stage.FormDivs {
		stage.FormDivs_reference[instance] = instance.GongCopy().(*FormDiv)
		stage.FormDivs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormEditAssocButtons_reference = make(map[*FormEditAssocButton]*FormEditAssocButton)
	stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint) // diff Unstage needs the reference order
	for instance := range stage.FormEditAssocButtons {
		stage.FormEditAssocButtons_reference[instance] = instance.GongCopy().(*FormEditAssocButton)
		stage.FormEditAssocButtons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFields_reference = make(map[*FormField]*FormField)
	stage.FormFields_referenceOrder = make(map[*FormField]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFields {
		stage.FormFields_reference[instance] = instance.GongCopy().(*FormField)
		stage.FormFields_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDates_reference = make(map[*FormFieldDate]*FormFieldDate)
	stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldDates {
		stage.FormFieldDates_reference[instance] = instance.GongCopy().(*FormFieldDate)
		stage.FormFieldDates_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDateTimes_reference = make(map[*FormFieldDateTime]*FormFieldDateTime)
	stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldDateTimes {
		stage.FormFieldDateTimes_reference[instance] = instance.GongCopy().(*FormFieldDateTime)
		stage.FormFieldDateTimes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldFloat64s_reference = make(map[*FormFieldFloat64]*FormFieldFloat64)
	stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldFloat64s {
		stage.FormFieldFloat64s_reference[instance] = instance.GongCopy().(*FormFieldFloat64)
		stage.FormFieldFloat64s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldInts_reference = make(map[*FormFieldInt]*FormFieldInt)
	stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldInts {
		stage.FormFieldInts_reference[instance] = instance.GongCopy().(*FormFieldInt)
		stage.FormFieldInts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldSelects_reference = make(map[*FormFieldSelect]*FormFieldSelect)
	stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldSelects {
		stage.FormFieldSelects_reference[instance] = instance.GongCopy().(*FormFieldSelect)
		stage.FormFieldSelects_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldStrings_reference = make(map[*FormFieldString]*FormFieldString)
	stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldStrings {
		stage.FormFieldStrings_reference[instance] = instance.GongCopy().(*FormFieldString)
		stage.FormFieldStrings_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormFieldTimes_reference = make(map[*FormFieldTime]*FormFieldTime)
	stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint) // diff Unstage needs the reference order
	for instance := range stage.FormFieldTimes {
		stage.FormFieldTimes_reference[instance] = instance.GongCopy().(*FormFieldTime)
		stage.FormFieldTimes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormGroups_reference = make(map[*FormGroup]*FormGroup)
	stage.FormGroups_referenceOrder = make(map[*FormGroup]uint) // diff Unstage needs the reference order
	for instance := range stage.FormGroups {
		stage.FormGroups_reference[instance] = instance.GongCopy().(*FormGroup)
		stage.FormGroups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FormSortAssocButtons_reference = make(map[*FormSortAssocButton]*FormSortAssocButton)
	stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint) // diff Unstage needs the reference order
	for instance := range stage.FormSortAssocButtons {
		stage.FormSortAssocButtons_reference[instance] = instance.GongCopy().(*FormSortAssocButton)
		stage.FormSortAssocButtons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Options_reference = make(map[*Option]*Option)
	stage.Options_referenceOrder = make(map[*Option]uint) // diff Unstage needs the reference order
	for instance := range stage.Options {
		stage.Options_reference[instance] = instance.GongCopy().(*Option)
		stage.Options_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Rows_reference = make(map[*Row]*Row)
	stage.Rows_referenceOrder = make(map[*Row]uint) // diff Unstage needs the reference order
	for instance := range stage.Rows {
		stage.Rows_reference[instance] = instance.GongCopy().(*Row)
		stage.Rows_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
		stage.Tables_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (cell *Cell) GongGetOrder(stage *Stage) uint {
	return stage.CellMap_Staged_Order[cell]
}

func (cell *Cell) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Cells_referenceOrder[cell]
}

func (cellboolean *CellBoolean) GongGetOrder(stage *Stage) uint {
	return stage.CellBooleanMap_Staged_Order[cellboolean]
}

func (cellboolean *CellBoolean) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CellBooleans_referenceOrder[cellboolean]
}

func (cellfloat64 *CellFloat64) GongGetOrder(stage *Stage) uint {
	return stage.CellFloat64Map_Staged_Order[cellfloat64]
}

func (cellfloat64 *CellFloat64) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CellFloat64s_referenceOrder[cellfloat64]
}

func (cellicon *CellIcon) GongGetOrder(stage *Stage) uint {
	return stage.CellIconMap_Staged_Order[cellicon]
}

func (cellicon *CellIcon) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CellIcons_referenceOrder[cellicon]
}

func (cellint *CellInt) GongGetOrder(stage *Stage) uint {
	return stage.CellIntMap_Staged_Order[cellint]
}

func (cellint *CellInt) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CellInts_referenceOrder[cellint]
}

func (cellstring *CellString) GongGetOrder(stage *Stage) uint {
	return stage.CellStringMap_Staged_Order[cellstring]
}

func (cellstring *CellString) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CellStrings_referenceOrder[cellstring]
}

func (checkbox *CheckBox) GongGetOrder(stage *Stage) uint {
	return stage.CheckBoxMap_Staged_Order[checkbox]
}

func (checkbox *CheckBox) GongGetReferenceOrder(stage *Stage) uint {
	return stage.CheckBoxs_referenceOrder[checkbox]
}

func (displayedcolumn *DisplayedColumn) GongGetOrder(stage *Stage) uint {
	return stage.DisplayedColumnMap_Staged_Order[displayedcolumn]
}

func (displayedcolumn *DisplayedColumn) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DisplayedColumns_referenceOrder[displayedcolumn]
}

func (formdiv *FormDiv) GongGetOrder(stage *Stage) uint {
	return stage.FormDivMap_Staged_Order[formdiv]
}

func (formdiv *FormDiv) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormDivs_referenceOrder[formdiv]
}

func (formeditassocbutton *FormEditAssocButton) GongGetOrder(stage *Stage) uint {
	return stage.FormEditAssocButtonMap_Staged_Order[formeditassocbutton]
}

func (formeditassocbutton *FormEditAssocButton) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormEditAssocButtons_referenceOrder[formeditassocbutton]
}

func (formfield *FormField) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldMap_Staged_Order[formfield]
}

func (formfield *FormField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFields_referenceOrder[formfield]
}

func (formfielddate *FormFieldDate) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldDateMap_Staged_Order[formfielddate]
}

func (formfielddate *FormFieldDate) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldDates_referenceOrder[formfielddate]
}

func (formfielddatetime *FormFieldDateTime) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldDateTimeMap_Staged_Order[formfielddatetime]
}

func (formfielddatetime *FormFieldDateTime) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldDateTimes_referenceOrder[formfielddatetime]
}

func (formfieldfloat64 *FormFieldFloat64) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldFloat64Map_Staged_Order[formfieldfloat64]
}

func (formfieldfloat64 *FormFieldFloat64) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldFloat64s_referenceOrder[formfieldfloat64]
}

func (formfieldint *FormFieldInt) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldIntMap_Staged_Order[formfieldint]
}

func (formfieldint *FormFieldInt) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldInts_referenceOrder[formfieldint]
}

func (formfieldselect *FormFieldSelect) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldSelectMap_Staged_Order[formfieldselect]
}

func (formfieldselect *FormFieldSelect) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldSelects_referenceOrder[formfieldselect]
}

func (formfieldstring *FormFieldString) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldStringMap_Staged_Order[formfieldstring]
}

func (formfieldstring *FormFieldString) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldStrings_referenceOrder[formfieldstring]
}

func (formfieldtime *FormFieldTime) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldTimeMap_Staged_Order[formfieldtime]
}

func (formfieldtime *FormFieldTime) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormFieldTimes_referenceOrder[formfieldtime]
}

func (formgroup *FormGroup) GongGetOrder(stage *Stage) uint {
	return stage.FormGroupMap_Staged_Order[formgroup]
}

func (formgroup *FormGroup) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormGroups_referenceOrder[formgroup]
}

func (formsortassocbutton *FormSortAssocButton) GongGetOrder(stage *Stage) uint {
	return stage.FormSortAssocButtonMap_Staged_Order[formsortassocbutton]
}

func (formsortassocbutton *FormSortAssocButton) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FormSortAssocButtons_referenceOrder[formsortassocbutton]
}

func (option *Option) GongGetOrder(stage *Stage) uint {
	return stage.OptionMap_Staged_Order[option]
}

func (option *Option) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Options_referenceOrder[option]
}

func (row *Row) GongGetOrder(stage *Stage) uint {
	return stage.RowMap_Staged_Order[row]
}

func (row *Row) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Rows_referenceOrder[row]
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return stage.TableMap_Staged_Order[table]
}

func (table *Table) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tables_referenceOrder[table]
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
	return fmt.Sprintf("__%s__%08d_", cell.GongGetGongstructName(), cell.GongGetReferenceOrder(stage))
}

func (cellboolean *CellBoolean) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellboolean.GongGetGongstructName(), cellboolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellboolean *CellBoolean) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellboolean.GongGetGongstructName(), cellboolean.GongGetReferenceOrder(stage))
}

func (cellfloat64 *CellFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellfloat64.GongGetGongstructName(), cellfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellfloat64 *CellFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellfloat64.GongGetGongstructName(), cellfloat64.GongGetReferenceOrder(stage))
}

func (cellicon *CellIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellicon.GongGetGongstructName(), cellicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellicon *CellIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellicon.GongGetGongstructName(), cellicon.GongGetReferenceOrder(stage))
}

func (cellint *CellInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellint.GongGetGongstructName(), cellint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellint *CellInt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellint.GongGetGongstructName(), cellint.GongGetReferenceOrder(stage))
}

func (cellstring *CellString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellstring.GongGetGongstructName(), cellstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cellstring *CellString) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellstring.GongGetGongstructName(), cellstring.GongGetReferenceOrder(stage))
}

func (checkbox *CheckBox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkbox *CheckBox) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetReferenceOrder(stage))
}

func (displayedcolumn *DisplayedColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (displayedcolumn *DisplayedColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetReferenceOrder(stage))
}

func (formdiv *FormDiv) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formdiv *FormDiv) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetReferenceOrder(stage))
}

func (formeditassocbutton *FormEditAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formeditassocbutton *FormEditAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetReferenceOrder(stage))
}

func (formfield *FormField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfield *FormField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetReferenceOrder(stage))
}

func (formfielddate *FormFieldDate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddate *FormFieldDate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetReferenceOrder(stage))
}

func (formfielddatetime *FormFieldDateTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddatetime *FormFieldDateTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetReferenceOrder(stage))
}

func (formfieldfloat64 *FormFieldFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldfloat64 *FormFieldFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetReferenceOrder(stage))
}

func (formfieldint *FormFieldInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldint *FormFieldInt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetReferenceOrder(stage))
}

func (formfieldselect *FormFieldSelect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldselect *FormFieldSelect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetReferenceOrder(stage))
}

func (formfieldstring *FormFieldString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldstring *FormFieldString) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetReferenceOrder(stage))
}

func (formfieldtime *FormFieldTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldtime *FormFieldTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetReferenceOrder(stage))
}

func (formgroup *FormGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formgroup *FormGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetReferenceOrder(stage))
}

func (formsortassocbutton *FormSortAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formsortassocbutton *FormSortAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetReferenceOrder(stage))
}

func (option *Option) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (option *Option) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetReferenceOrder(stage))
}

func (row *Row) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (row *Row) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetReferenceOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (cell *Cell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cell.Name)
	return
}
func (cellboolean *CellBoolean) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellBoolean")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cellboolean.Name)
	return
}
func (cellfloat64 *CellFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cellfloat64.Name)
	return
}
func (cellicon *CellIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cellicon.Name)
	return
}
func (cellint *CellInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cellint.Name)
	return
}
func (cellstring *CellString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CellString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cellstring.Name)
	return
}
func (checkbox *CheckBox) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CheckBox")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", checkbox.Name)
	return
}
func (displayedcolumn *DisplayedColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplayedColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", displayedcolumn.Name)
	return
}
func (formdiv *FormDiv) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormDiv")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formdiv.Name)
	return
}
func (formeditassocbutton *FormEditAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormEditAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formeditassocbutton.Name)
	return
}
func (formfield *FormField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfield.Name)
	return
}
func (formfielddate *FormFieldDate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfielddate.Name)
	return
}
func (formfielddatetime *FormFieldDateTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDateTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfielddatetime.Name)
	return
}
func (formfieldfloat64 *FormFieldFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldfloat64.Name)
	return
}
func (formfieldint *FormFieldInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldint.Name)
	return
}
func (formfieldselect *FormFieldSelect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldSelect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldselect.Name)
	return
}
func (formfieldstring *FormFieldString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldstring.Name)
	return
}
func (formfieldtime *FormFieldTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldtime.Name)
	return
}
func (formgroup *FormGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formgroup.Name)
	return
}
func (formsortassocbutton *FormSortAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormSortAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formsortassocbutton.Name)
	return
}
func (option *Option) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Option")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", option.Name)
	return
}
func (row *Row) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Row")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", row.Name)
	return
}
func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", table.Name)
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
