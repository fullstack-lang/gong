// generated code - do not edit
package models

// insertion point
func (inst *DisplaySelection) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *XLCell) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}

func (inst *XLFile) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *XLRow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}

func (inst *XLSheet) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}
