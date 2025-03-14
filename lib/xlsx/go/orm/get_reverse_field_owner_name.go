// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.DisplaySelection:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XLCell:
		switch reverseField.GongstructName {
		// insertion point
		case "XLRow":
			switch reverseField.Fieldname {
			case "Cells":
				if _xlrow, ok := stage.XLRow_Cells_reverseMap[inst]; ok {
					res = _xlrow.Name
				}
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				if _xlsheet, ok := stage.XLSheet_SheetCells_reverseMap[inst]; ok {
					res = _xlsheet.Name
				}
			}
		}

	case *models.XLFile:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XLRow:
		switch reverseField.GongstructName {
		// insertion point
		case "XLSheet":
			switch reverseField.Fieldname {
			case "Rows":
				if _xlsheet, ok := stage.XLSheet_Rows_reverseMap[inst]; ok {
					res = _xlsheet.Name
				}
			}
		}

	case *models.XLSheet:
		switch reverseField.GongstructName {
		// insertion point
		case "XLFile":
			switch reverseField.Fieldname {
			case "Sheets":
				if _xlfile, ok := stage.XLFile_Sheets_reverseMap[inst]; ok {
					res = _xlfile.Name
				}
			}
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
	case *models.DisplaySelection:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XLCell:
		switch reverseField.GongstructName {
		// insertion point
		case "XLRow":
			switch reverseField.Fieldname {
			case "Cells":
				res = stage.XLRow_Cells_reverseMap[inst]
			}
		case "XLSheet":
			switch reverseField.Fieldname {
			case "SheetCells":
				res = stage.XLSheet_SheetCells_reverseMap[inst]
			}
		}

	case *models.XLFile:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.XLRow:
		switch reverseField.GongstructName {
		// insertion point
		case "XLSheet":
			switch reverseField.Fieldname {
			case "Rows":
				res = stage.XLSheet_Rows_reverseMap[inst]
			}
		}

	case *models.XLSheet:
		switch reverseField.GongstructName {
		// insertion point
		case "XLFile":
			switch reverseField.Fieldname {
			case "Sheets":
				res = stage.XLFile_Sheets_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
