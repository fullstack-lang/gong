// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtable/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				if _row, ok := stage.Row_Cells_reverseMap[inst]; ok {
					res = _row.Name
				}
			}
		}

	case *models.CellBoolean:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CheckBox:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				if _formdiv, ok := stage.FormDiv_CheckBoxs_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		}

	case *models.DisplayedColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "DisplayedColumns":
				if _table, ok := stage.Table_DisplayedColumns_reverseMap[inst]; ok {
					res = _table.Name
				}
			}
		}

	case *models.FormDiv:
		switch reverseField.GongstructName {
		// insertion point
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				if _formgroup, ok := stage.FormGroup_FormDivs_reverseMap[inst]; ok {
					res = _formgroup.Name
				}
			}
		}

	case *models.FormEditAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormField:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				if _formdiv, ok := stage.FormDiv_FormFields_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		}

	case *models.FormFieldDate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldDateTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldSelect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormSortAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Option:
		switch reverseField.GongstructName {
		// insertion point
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				if _formfieldselect, ok := stage.FormFieldSelect_Options_reverseMap[inst]; ok {
					res = _formfieldselect.Name
				}
			}
		}

	case *models.Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "Rows":
				if _table, ok := stage.Table_Rows_reverseMap[inst]; ok {
					res = _table.Name
				}
			}
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Cell:
		switch reverseField.GongstructName {
		// insertion point
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				res = stage.Row_Cells_reverseMap[inst]
			}
		}

	case *models.CellBoolean:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellIcon:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CellString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.CheckBox:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				res = stage.FormDiv_CheckBoxs_reverseMap[inst]
			}
		}

	case *models.DisplayedColumn:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "DisplayedColumns":
				res = stage.Table_DisplayedColumns_reverseMap[inst]
			}
		}

	case *models.FormDiv:
		switch reverseField.GongstructName {
		// insertion point
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				res = stage.FormGroup_FormDivs_reverseMap[inst]
			}
		}

	case *models.FormEditAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormField:
		switch reverseField.GongstructName {
		// insertion point
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				res = stage.FormDiv_FormFields_reverseMap[inst]
			}
		}

	case *models.FormFieldDate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldDateTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldFloat64:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldInt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldSelect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldString:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormFieldTime:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormGroup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.FormSortAssocButton:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Option:
		switch reverseField.GongstructName {
		// insertion point
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				res = stage.FormFieldSelect_Options_reverseMap[inst]
			}
		}

	case *models.Row:
		switch reverseField.GongstructName {
		// insertion point
		case "Table":
			switch reverseField.Fieldname {
			case "Rows":
				res = stage.Table_Rows_reverseMap[inst]
			}
		}

	case *models.Table:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
