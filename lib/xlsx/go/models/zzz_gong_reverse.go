// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *DisplaySelection:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *XLCell:
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

	case *XLFile:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *XLRow:
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

	case *XLSheet:
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

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *DisplaySelection:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *XLCell:
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

	case *XLFile:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *XLRow:
		switch reverseField.GongstructName {
		// insertion point
		case "XLSheet":
			switch reverseField.Fieldname {
			case "Rows":
				res = stage.XLSheet_Rows_reverseMap[inst]
			}
		}

	case *XLSheet:
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
