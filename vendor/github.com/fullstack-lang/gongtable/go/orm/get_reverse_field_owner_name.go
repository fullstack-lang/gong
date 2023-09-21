// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongtable/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Cell:
		tmp := GetInstanceDBFromInstance[models.Cell, CellDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			case "Cells":
				if tmp.Row_CellsDBID.Int64 != 0 {
					id := uint(tmp.Row_CellsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoRow.Map_RowDBID_RowPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CellBoolean:
		tmp := GetInstanceDBFromInstance[models.CellBoolean, CellBooleanDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CellFloat64:
		tmp := GetInstanceDBFromInstance[models.CellFloat64, CellFloat64DB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CellIcon:
		tmp := GetInstanceDBFromInstance[models.CellIcon, CellIconDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CellInt:
		tmp := GetInstanceDBFromInstance[models.CellInt, CellIntDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CellString:
		tmp := GetInstanceDBFromInstance[models.CellString, CellStringDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.CheckBox:
		tmp := GetInstanceDBFromInstance[models.CheckBox, CheckBoxDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				if tmp.FormDiv_CheckBoxsDBID.Int64 != 0 {
					id := uint(tmp.FormDiv_CheckBoxsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.DisplayedColumn:
		tmp := GetInstanceDBFromInstance[models.DisplayedColumn, DisplayedColumnDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			case "DisplayedColumns":
				if tmp.Table_DisplayedColumnsDBID.Int64 != 0 {
					id := uint(tmp.Table_DisplayedColumnsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTable.Map_TableDBID_TablePtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.FormDiv:
		tmp := GetInstanceDBFromInstance[models.FormDiv, FormDivDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				if tmp.FormGroup_FormDivsDBID.Int64 != 0 {
					id := uint(tmp.FormGroup_FormDivsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormEditAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormEditAssocButton, FormEditAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormField:
		tmp := GetInstanceDBFromInstance[models.FormField, FormFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				if tmp.FormDiv_FormFieldsDBID.Int64 != 0 {
					id := uint(tmp.FormDiv_FormFieldsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDate:
		tmp := GetInstanceDBFromInstance[models.FormFieldDate, FormFieldDateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDateTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldDateTime, FormFieldDateTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldFloat64:
		tmp := GetInstanceDBFromInstance[models.FormFieldFloat64, FormFieldFloat64DB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldInt:
		tmp := GetInstanceDBFromInstance[models.FormFieldInt, FormFieldIntDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldSelect:
		tmp := GetInstanceDBFromInstance[models.FormFieldSelect, FormFieldSelectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldString:
		tmp := GetInstanceDBFromInstance[models.FormFieldString, FormFieldStringDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldTime, FormFieldTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormGroup:
		tmp := GetInstanceDBFromInstance[models.FormGroup, FormGroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormSortAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormSortAssocButton, FormSortAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.Option:
		tmp := GetInstanceDBFromInstance[models.Option, OptionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				if tmp.FormFieldSelect_OptionsDBID.Int64 != 0 {
					id := uint(tmp.FormFieldSelect_OptionsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoFormFieldSelect.Map_FormFieldSelectDBID_FormFieldSelectPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	case *models.Row:
		tmp := GetInstanceDBFromInstance[models.Row, RowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			case "Rows":
				if tmp.Table_RowsDBID.Int64 != 0 {
					id := uint(tmp.Table_RowsDBID.Int64)
					reservePointerTarget := backRepo.BackRepoTable.Map_TableDBID_TablePtr[id]
					res = reservePointerTarget.Name
				}
			}
		}

	case *models.Table:
		tmp := GetInstanceDBFromInstance[models.Table, TableDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Cell":
			switch reverseField.Fieldname {
			}
		case "CellBoolean":
			switch reverseField.Fieldname {
			}
		case "CellFloat64":
			switch reverseField.Fieldname {
			}
		case "CellIcon":
			switch reverseField.Fieldname {
			}
		case "CellInt":
			switch reverseField.Fieldname {
			}
		case "CellString":
			switch reverseField.Fieldname {
			}
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "DisplayedColumn":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		case "Row":
			switch reverseField.Fieldname {
			}
		case "Table":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}
