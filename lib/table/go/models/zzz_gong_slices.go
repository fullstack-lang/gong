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
			fieldInitializers, pointersInitializations := cell.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cell.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cell.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cell := range stage.Cells_reference {
		if _, ok := stage.Cells[cell]; !ok {
			cells_deletedInstances = append(cells_deletedInstances, cell)
			deletedInstancesStmt += cell.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := cellboolean.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellboolean.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cellboolean.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellboolean := range stage.CellBooleans_reference {
		if _, ok := stage.CellBooleans[cellboolean]; !ok {
			cellbooleans_deletedInstances = append(cellbooleans_deletedInstances, cellboolean)
			deletedInstancesStmt += cellboolean.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := cellfloat64.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellfloat64.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cellfloat64.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellfloat64 := range stage.CellFloat64s_reference {
		if _, ok := stage.CellFloat64s[cellfloat64]; !ok {
			cellfloat64s_deletedInstances = append(cellfloat64s_deletedInstances, cellfloat64)
			deletedInstancesStmt += cellfloat64.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := cellicon.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellicon.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cellicon.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellicon := range stage.CellIcons_reference {
		if _, ok := stage.CellIcons[cellicon]; !ok {
			cellicons_deletedInstances = append(cellicons_deletedInstances, cellicon)
			deletedInstancesStmt += cellicon.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := cellint.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellint.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cellint.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellint := range stage.CellInts_reference {
		if _, ok := stage.CellInts[cellint]; !ok {
			cellints_deletedInstances = append(cellints_deletedInstances, cellint)
			deletedInstancesStmt += cellint.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := cellstring.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := cellstring.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", cellstring.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellstring := range stage.CellStrings_reference {
		if _, ok := stage.CellStrings[cellstring]; !ok {
			cellstrings_deletedInstances = append(cellstrings_deletedInstances, cellstring)
			deletedInstancesStmt += cellstring.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := checkbox.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", checkbox.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for checkbox := range stage.CheckBoxs_reference {
		if _, ok := stage.CheckBoxs[checkbox]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, checkbox)
			deletedInstancesStmt += checkbox.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := displayedcolumn.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := displayedcolumn.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", displayedcolumn.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for displayedcolumn := range stage.DisplayedColumns_reference {
		if _, ok := stage.DisplayedColumns[displayedcolumn]; !ok {
			displayedcolumns_deletedInstances = append(displayedcolumns_deletedInstances, displayedcolumn)
			deletedInstancesStmt += displayedcolumn.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formdiv.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formdiv.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formdiv.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formdiv := range stage.FormDivs_reference {
		if _, ok := stage.FormDivs[formdiv]; !ok {
			formdivs_deletedInstances = append(formdivs_deletedInstances, formdiv)
			deletedInstancesStmt += formdiv.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formeditassocbutton.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formeditassocbutton.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formeditassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formeditassocbutton := range stage.FormEditAssocButtons_reference {
		if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; !ok {
			formeditassocbuttons_deletedInstances = append(formeditassocbuttons_deletedInstances, formeditassocbutton)
			deletedInstancesStmt += formeditassocbutton.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfield.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfield := range stage.FormFields_reference {
		if _, ok := stage.FormFields[formfield]; !ok {
			formfields_deletedInstances = append(formfields_deletedInstances, formfield)
			deletedInstancesStmt += formfield.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfielddate.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfielddate.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfielddate.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfielddate := range stage.FormFieldDates_reference {
		if _, ok := stage.FormFieldDates[formfielddate]; !ok {
			formfielddates_deletedInstances = append(formfielddates_deletedInstances, formfielddate)
			deletedInstancesStmt += formfielddate.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfielddatetime.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfielddatetime.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfielddatetime.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfielddatetime := range stage.FormFieldDateTimes_reference {
		if _, ok := stage.FormFieldDateTimes[formfielddatetime]; !ok {
			formfielddatetimes_deletedInstances = append(formfielddatetimes_deletedInstances, formfielddatetime)
			deletedInstancesStmt += formfielddatetime.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfieldfloat64.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldfloat64.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfieldfloat64.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldfloat64 := range stage.FormFieldFloat64s_reference {
		if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; !ok {
			formfieldfloat64s_deletedInstances = append(formfieldfloat64s_deletedInstances, formfieldfloat64)
			deletedInstancesStmt += formfieldfloat64.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfieldint.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldint.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfieldint.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldint := range stage.FormFieldInts_reference {
		if _, ok := stage.FormFieldInts[formfieldint]; !ok {
			formfieldints_deletedInstances = append(formfieldints_deletedInstances, formfieldint)
			deletedInstancesStmt += formfieldint.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfieldselect.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldselect.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfieldselect.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldselect := range stage.FormFieldSelects_reference {
		if _, ok := stage.FormFieldSelects[formfieldselect]; !ok {
			formfieldselects_deletedInstances = append(formfieldselects_deletedInstances, formfieldselect)
			deletedInstancesStmt += formfieldselect.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfieldstring.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldstring.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfieldstring.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldstring := range stage.FormFieldStrings_reference {
		if _, ok := stage.FormFieldStrings[formfieldstring]; !ok {
			formfieldstrings_deletedInstances = append(formfieldstrings_deletedInstances, formfieldstring)
			deletedInstancesStmt += formfieldstring.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formfieldtime.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formfieldtime.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formfieldtime.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldtime := range stage.FormFieldTimes_reference {
		if _, ok := stage.FormFieldTimes[formfieldtime]; !ok {
			formfieldtimes_deletedInstances = append(formfieldtimes_deletedInstances, formfieldtime)
			deletedInstancesStmt += formfieldtime.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formgroup.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formgroup.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formgroup.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formgroup := range stage.FormGroups_reference {
		if _, ok := stage.FormGroups[formgroup]; !ok {
			formgroups_deletedInstances = append(formgroups_deletedInstances, formgroup)
			deletedInstancesStmt += formgroup.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := formsortassocbutton.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := formsortassocbutton.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", formsortassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formsortassocbutton := range stage.FormSortAssocButtons_reference {
		if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; !ok {
			formsortassocbuttons_deletedInstances = append(formsortassocbuttons_deletedInstances, formsortassocbutton)
			deletedInstancesStmt += formsortassocbutton.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := option.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := option.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", option.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for option := range stage.Options_reference {
		if _, ok := stage.Options[option]; !ok {
			options_deletedInstances = append(options_deletedInstances, option)
			deletedInstancesStmt += option.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := row.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := row.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", row.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for row := range stage.Rows_reference {
		if _, ok := stage.Rows[row]; !ok {
			rows_deletedInstances = append(rows_deletedInstances, row)
			deletedInstancesStmt += row.GongMarshallUnstaging(stage)
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
			fieldInitializers, pointersInitializations := table.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := table.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", table.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for table := range stage.Tables_reference {
		if _, ok := stage.Tables[table]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, table)
			deletedInstancesStmt += table.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Cells_reference = make(map[*Cell]*Cell)
	for instance := range stage.Cells {
		stage.Cells_reference[instance] = instance.GongCopy().(*Cell)
	}

	stage.CellBooleans_reference = make(map[*CellBoolean]*CellBoolean)
	for instance := range stage.CellBooleans {
		stage.CellBooleans_reference[instance] = instance.GongCopy().(*CellBoolean)
	}

	stage.CellFloat64s_reference = make(map[*CellFloat64]*CellFloat64)
	for instance := range stage.CellFloat64s {
		stage.CellFloat64s_reference[instance] = instance.GongCopy().(*CellFloat64)
	}

	stage.CellIcons_reference = make(map[*CellIcon]*CellIcon)
	for instance := range stage.CellIcons {
		stage.CellIcons_reference[instance] = instance.GongCopy().(*CellIcon)
	}

	stage.CellInts_reference = make(map[*CellInt]*CellInt)
	for instance := range stage.CellInts {
		stage.CellInts_reference[instance] = instance.GongCopy().(*CellInt)
	}

	stage.CellStrings_reference = make(map[*CellString]*CellString)
	for instance := range stage.CellStrings {
		stage.CellStrings_reference[instance] = instance.GongCopy().(*CellString)
	}

	stage.CheckBoxs_reference = make(map[*CheckBox]*CheckBox)
	for instance := range stage.CheckBoxs {
		stage.CheckBoxs_reference[instance] = instance.GongCopy().(*CheckBox)
	}

	stage.DisplayedColumns_reference = make(map[*DisplayedColumn]*DisplayedColumn)
	for instance := range stage.DisplayedColumns {
		stage.DisplayedColumns_reference[instance] = instance.GongCopy().(*DisplayedColumn)
	}

	stage.FormDivs_reference = make(map[*FormDiv]*FormDiv)
	for instance := range stage.FormDivs {
		stage.FormDivs_reference[instance] = instance.GongCopy().(*FormDiv)
	}

	stage.FormEditAssocButtons_reference = make(map[*FormEditAssocButton]*FormEditAssocButton)
	for instance := range stage.FormEditAssocButtons {
		stage.FormEditAssocButtons_reference[instance] = instance.GongCopy().(*FormEditAssocButton)
	}

	stage.FormFields_reference = make(map[*FormField]*FormField)
	for instance := range stage.FormFields {
		stage.FormFields_reference[instance] = instance.GongCopy().(*FormField)
	}

	stage.FormFieldDates_reference = make(map[*FormFieldDate]*FormFieldDate)
	for instance := range stage.FormFieldDates {
		stage.FormFieldDates_reference[instance] = instance.GongCopy().(*FormFieldDate)
	}

	stage.FormFieldDateTimes_reference = make(map[*FormFieldDateTime]*FormFieldDateTime)
	for instance := range stage.FormFieldDateTimes {
		stage.FormFieldDateTimes_reference[instance] = instance.GongCopy().(*FormFieldDateTime)
	}

	stage.FormFieldFloat64s_reference = make(map[*FormFieldFloat64]*FormFieldFloat64)
	for instance := range stage.FormFieldFloat64s {
		stage.FormFieldFloat64s_reference[instance] = instance.GongCopy().(*FormFieldFloat64)
	}

	stage.FormFieldInts_reference = make(map[*FormFieldInt]*FormFieldInt)
	for instance := range stage.FormFieldInts {
		stage.FormFieldInts_reference[instance] = instance.GongCopy().(*FormFieldInt)
	}

	stage.FormFieldSelects_reference = make(map[*FormFieldSelect]*FormFieldSelect)
	for instance := range stage.FormFieldSelects {
		stage.FormFieldSelects_reference[instance] = instance.GongCopy().(*FormFieldSelect)
	}

	stage.FormFieldStrings_reference = make(map[*FormFieldString]*FormFieldString)
	for instance := range stage.FormFieldStrings {
		stage.FormFieldStrings_reference[instance] = instance.GongCopy().(*FormFieldString)
	}

	stage.FormFieldTimes_reference = make(map[*FormFieldTime]*FormFieldTime)
	for instance := range stage.FormFieldTimes {
		stage.FormFieldTimes_reference[instance] = instance.GongCopy().(*FormFieldTime)
	}

	stage.FormGroups_reference = make(map[*FormGroup]*FormGroup)
	for instance := range stage.FormGroups {
		stage.FormGroups_reference[instance] = instance.GongCopy().(*FormGroup)
	}

	stage.FormSortAssocButtons_reference = make(map[*FormSortAssocButton]*FormSortAssocButton)
	for instance := range stage.FormSortAssocButtons {
		stage.FormSortAssocButtons_reference[instance] = instance.GongCopy().(*FormSortAssocButton)
	}

	stage.Options_reference = make(map[*Option]*Option)
	for instance := range stage.Options {
		stage.Options_reference[instance] = instance.GongCopy().(*Option)
	}

	stage.Rows_reference = make(map[*Row]*Row)
	for instance := range stage.Rows {
		stage.Rows_reference[instance] = instance.GongCopy().(*Row)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
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

func (cellboolean *CellBoolean) GongGetOrder(stage *Stage) uint {
	return stage.CellBooleanMap_Staged_Order[cellboolean]
}

func (cellfloat64 *CellFloat64) GongGetOrder(stage *Stage) uint {
	return stage.CellFloat64Map_Staged_Order[cellfloat64]
}

func (cellicon *CellIcon) GongGetOrder(stage *Stage) uint {
	return stage.CellIconMap_Staged_Order[cellicon]
}

func (cellint *CellInt) GongGetOrder(stage *Stage) uint {
	return stage.CellIntMap_Staged_Order[cellint]
}

func (cellstring *CellString) GongGetOrder(stage *Stage) uint {
	return stage.CellStringMap_Staged_Order[cellstring]
}

func (checkbox *CheckBox) GongGetOrder(stage *Stage) uint {
	return stage.CheckBoxMap_Staged_Order[checkbox]
}

func (displayedcolumn *DisplayedColumn) GongGetOrder(stage *Stage) uint {
	return stage.DisplayedColumnMap_Staged_Order[displayedcolumn]
}

func (formdiv *FormDiv) GongGetOrder(stage *Stage) uint {
	return stage.FormDivMap_Staged_Order[formdiv]
}

func (formeditassocbutton *FormEditAssocButton) GongGetOrder(stage *Stage) uint {
	return stage.FormEditAssocButtonMap_Staged_Order[formeditassocbutton]
}

func (formfield *FormField) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldMap_Staged_Order[formfield]
}

func (formfielddate *FormFieldDate) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldDateMap_Staged_Order[formfielddate]
}

func (formfielddatetime *FormFieldDateTime) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldDateTimeMap_Staged_Order[formfielddatetime]
}

func (formfieldfloat64 *FormFieldFloat64) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldFloat64Map_Staged_Order[formfieldfloat64]
}

func (formfieldint *FormFieldInt) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldIntMap_Staged_Order[formfieldint]
}

func (formfieldselect *FormFieldSelect) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldSelectMap_Staged_Order[formfieldselect]
}

func (formfieldstring *FormFieldString) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldStringMap_Staged_Order[formfieldstring]
}

func (formfieldtime *FormFieldTime) GongGetOrder(stage *Stage) uint {
	return stage.FormFieldTimeMap_Staged_Order[formfieldtime]
}

func (formgroup *FormGroup) GongGetOrder(stage *Stage) uint {
	return stage.FormGroupMap_Staged_Order[formgroup]
}

func (formsortassocbutton *FormSortAssocButton) GongGetOrder(stage *Stage) uint {
	return stage.FormSortAssocButtonMap_Staged_Order[formsortassocbutton]
}

func (option *Option) GongGetOrder(stage *Stage) uint {
	return stage.OptionMap_Staged_Order[option]
}

func (row *Row) GongGetOrder(stage *Stage) uint {
	return stage.RowMap_Staged_Order[row]
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return stage.TableMap_Staged_Order[table]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (cell *Cell) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cell.GongGetGongstructName(), cell.GongGetOrder(stage))
}

func (cellboolean *CellBoolean) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellboolean.GongGetGongstructName(), cellboolean.GongGetOrder(stage))
}

func (cellfloat64 *CellFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellfloat64.GongGetGongstructName(), cellfloat64.GongGetOrder(stage))
}

func (cellicon *CellIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellicon.GongGetGongstructName(), cellicon.GongGetOrder(stage))
}

func (cellint *CellInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellint.GongGetGongstructName(), cellint.GongGetOrder(stage))
}

func (cellstring *CellString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cellstring.GongGetGongstructName(), cellstring.GongGetOrder(stage))
}

func (checkbox *CheckBox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

func (displayedcolumn *DisplayedColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", displayedcolumn.GongGetGongstructName(), displayedcolumn.GongGetOrder(stage))
}

func (formdiv *FormDiv) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

func (formeditassocbutton *FormEditAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

func (formfield *FormField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

func (formfielddate *FormFieldDate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

func (formfielddatetime *FormFieldDateTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

func (formfieldfloat64 *FormFieldFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

func (formfieldint *FormFieldInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

func (formfieldselect *FormFieldSelect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

func (formfieldstring *FormFieldString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

func (formfieldtime *FormFieldTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

func (formgroup *FormGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

func (formsortassocbutton *FormSortAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

func (option *Option) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

func (row *Row) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", row.GongGetGongstructName(), row.GongGetOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cell.GongGetIdentifier(stage))
	return
}
func (cellboolean *CellBoolean) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
	return
}
func (cellfloat64 *CellFloat64) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
	return
}
func (cellicon *CellIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
	return
}
func (cellint *CellInt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellint.GongGetIdentifier(stage))
	return
}
func (cellstring *CellString) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
	return
}
func (checkbox *CheckBox) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	return
}
func (displayedcolumn *DisplayedColumn) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
	return
}
func (formdiv *FormDiv) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
	return
}
func (formeditassocbutton *FormEditAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
	return
}
func (formfield *FormField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetIdentifier(stage))
	return
}
func (formfielddate *FormFieldDate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
	return
}
func (formfielddatetime *FormFieldDateTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
	return
}
func (formfieldfloat64 *FormFieldFloat64) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
	return
}
func (formfieldint *FormFieldInt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
	return
}
func (formfieldselect *FormFieldSelect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
	return
}
func (formfieldstring *FormFieldString) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
	return
}
func (formfieldtime *FormFieldTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
	return
}
func (formgroup *FormGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
	return
}
func (formsortassocbutton *FormSortAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
	return
}
func (option *Option) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetIdentifier(stage))
	return
}
func (row *Row) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", row.GongGetIdentifier(stage))
	return
}
func (table *Table) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	return
}
