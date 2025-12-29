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

	// insertion point per named struct
	var cells_newInstances []*Cell
	var cells_deletedInstances []*Cell

	// parse all staged instances and check if they have a reference
	for cell := range stage.Cells {
		if ref, ok := stage.Cells_reference[cell]; !ok {
			cells_newInstances = append(cells_newInstances, cell)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Cell "+cell.Name,
				)
			}
		} else {
			diffs := cell.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Cell \""+cell.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cell := range stage.Cells_reference {
		if _, ok := stage.Cells[cell]; !ok {
			cells_deletedInstances = append(cells_deletedInstances, cell)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Cell "+cell.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CellBoolean "+cellboolean.Name,
				)
			}
		} else {
			diffs := cellboolean.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CellBoolean \""+cellboolean.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellboolean := range stage.CellBooleans_reference {
		if _, ok := stage.CellBooleans[cellboolean]; !ok {
			cellbooleans_deletedInstances = append(cellbooleans_deletedInstances, cellboolean)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CellBoolean "+cellboolean.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CellFloat64 "+cellfloat64.Name,
				)
			}
		} else {
			diffs := cellfloat64.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CellFloat64 \""+cellfloat64.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellfloat64 := range stage.CellFloat64s_reference {
		if _, ok := stage.CellFloat64s[cellfloat64]; !ok {
			cellfloat64s_deletedInstances = append(cellfloat64s_deletedInstances, cellfloat64)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CellFloat64 "+cellfloat64.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CellIcon "+cellicon.Name,
				)
			}
		} else {
			diffs := cellicon.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CellIcon \""+cellicon.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellicon := range stage.CellIcons_reference {
		if _, ok := stage.CellIcons[cellicon]; !ok {
			cellicons_deletedInstances = append(cellicons_deletedInstances, cellicon)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CellIcon "+cellicon.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CellInt "+cellint.Name,
				)
			}
		} else {
			diffs := cellint.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CellInt \""+cellint.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellint := range stage.CellInts_reference {
		if _, ok := stage.CellInts[cellint]; !ok {
			cellints_deletedInstances = append(cellints_deletedInstances, cellint)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CellInt "+cellint.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CellString "+cellstring.Name,
				)
			}
		} else {
			diffs := cellstring.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CellString \""+cellstring.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cellstring := range stage.CellStrings_reference {
		if _, ok := stage.CellStrings[cellstring]; !ok {
			cellstrings_deletedInstances = append(cellstrings_deletedInstances, cellstring)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CellString "+cellstring.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of CheckBox "+checkbox.Name,
				)
			}
		} else {
			diffs := checkbox.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of CheckBox \""+checkbox.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for checkbox := range stage.CheckBoxs_reference {
		if _, ok := stage.CheckBoxs[checkbox]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, checkbox)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of CheckBox "+checkbox.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of DisplayedColumn "+displayedcolumn.Name,
				)
			}
		} else {
			diffs := displayedcolumn.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of DisplayedColumn \""+displayedcolumn.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for displayedcolumn := range stage.DisplayedColumns_reference {
		if _, ok := stage.DisplayedColumns[displayedcolumn]; !ok {
			displayedcolumns_deletedInstances = append(displayedcolumns_deletedInstances, displayedcolumn)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of DisplayedColumn "+displayedcolumn.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormDiv "+formdiv.Name,
				)
			}
		} else {
			diffs := formdiv.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormDiv \""+formdiv.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formdiv := range stage.FormDivs_reference {
		if _, ok := stage.FormDivs[formdiv]; !ok {
			formdivs_deletedInstances = append(formdivs_deletedInstances, formdiv)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormDiv "+formdiv.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormEditAssocButton "+formeditassocbutton.Name,
				)
			}
		} else {
			diffs := formeditassocbutton.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormEditAssocButton \""+formeditassocbutton.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formeditassocbutton := range stage.FormEditAssocButtons_reference {
		if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; !ok {
			formeditassocbuttons_deletedInstances = append(formeditassocbuttons_deletedInstances, formeditassocbutton)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormEditAssocButton "+formeditassocbutton.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormField "+formfield.Name,
				)
			}
		} else {
			diffs := formfield.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormField \""+formfield.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfield := range stage.FormFields_reference {
		if _, ok := stage.FormFields[formfield]; !ok {
			formfields_deletedInstances = append(formfields_deletedInstances, formfield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormField "+formfield.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldDate "+formfielddate.Name,
				)
			}
		} else {
			diffs := formfielddate.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldDate \""+formfielddate.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfielddate := range stage.FormFieldDates_reference {
		if _, ok := stage.FormFieldDates[formfielddate]; !ok {
			formfielddates_deletedInstances = append(formfielddates_deletedInstances, formfielddate)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldDate "+formfielddate.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldDateTime "+formfielddatetime.Name,
				)
			}
		} else {
			diffs := formfielddatetime.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldDateTime \""+formfielddatetime.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfielddatetime := range stage.FormFieldDateTimes_reference {
		if _, ok := stage.FormFieldDateTimes[formfielddatetime]; !ok {
			formfielddatetimes_deletedInstances = append(formfielddatetimes_deletedInstances, formfielddatetime)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldDateTime "+formfielddatetime.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldFloat64 "+formfieldfloat64.Name,
				)
			}
		} else {
			diffs := formfieldfloat64.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldFloat64 \""+formfieldfloat64.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldfloat64 := range stage.FormFieldFloat64s_reference {
		if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; !ok {
			formfieldfloat64s_deletedInstances = append(formfieldfloat64s_deletedInstances, formfieldfloat64)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldFloat64 "+formfieldfloat64.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldInt "+formfieldint.Name,
				)
			}
		} else {
			diffs := formfieldint.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldInt \""+formfieldint.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldint := range stage.FormFieldInts_reference {
		if _, ok := stage.FormFieldInts[formfieldint]; !ok {
			formfieldints_deletedInstances = append(formfieldints_deletedInstances, formfieldint)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldInt "+formfieldint.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldSelect "+formfieldselect.Name,
				)
			}
		} else {
			diffs := formfieldselect.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldSelect \""+formfieldselect.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldselect := range stage.FormFieldSelects_reference {
		if _, ok := stage.FormFieldSelects[formfieldselect]; !ok {
			formfieldselects_deletedInstances = append(formfieldselects_deletedInstances, formfieldselect)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldSelect "+formfieldselect.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldString "+formfieldstring.Name,
				)
			}
		} else {
			diffs := formfieldstring.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldString \""+formfieldstring.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldstring := range stage.FormFieldStrings_reference {
		if _, ok := stage.FormFieldStrings[formfieldstring]; !ok {
			formfieldstrings_deletedInstances = append(formfieldstrings_deletedInstances, formfieldstring)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldString "+formfieldstring.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormFieldTime "+formfieldtime.Name,
				)
			}
		} else {
			diffs := formfieldtime.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormFieldTime \""+formfieldtime.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formfieldtime := range stage.FormFieldTimes_reference {
		if _, ok := stage.FormFieldTimes[formfieldtime]; !ok {
			formfieldtimes_deletedInstances = append(formfieldtimes_deletedInstances, formfieldtime)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormFieldTime "+formfieldtime.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormGroup "+formgroup.Name,
				)
			}
		} else {
			diffs := formgroup.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormGroup \""+formgroup.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formgroup := range stage.FormGroups_reference {
		if _, ok := stage.FormGroups[formgroup]; !ok {
			formgroups_deletedInstances = append(formgroups_deletedInstances, formgroup)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormGroup "+formgroup.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FormSortAssocButton "+formsortassocbutton.Name,
				)
			}
		} else {
			diffs := formsortassocbutton.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FormSortAssocButton \""+formsortassocbutton.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for formsortassocbutton := range stage.FormSortAssocButtons_reference {
		if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; !ok {
			formsortassocbuttons_deletedInstances = append(formsortassocbuttons_deletedInstances, formsortassocbutton)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FormSortAssocButton "+formsortassocbutton.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Option "+option.Name,
				)
			}
		} else {
			diffs := option.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Option \""+option.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for option := range stage.Options_reference {
		if _, ok := stage.Options[option]; !ok {
			options_deletedInstances = append(options_deletedInstances, option)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Option "+option.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Row "+row.Name,
				)
			}
		} else {
			diffs := row.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Row \""+row.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for row := range stage.Rows_reference {
		if _, ok := stage.Rows[row]; !ok {
			rows_deletedInstances = append(rows_deletedInstances, row)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Row "+row.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Table "+table.Name,
				)
			}
		} else {
			diffs := table.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Table \""+table.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for table := range stage.Tables_reference {
		if _, ok := stage.Tables[table]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, table)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Table "+table.Name,
				)
			}
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
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
