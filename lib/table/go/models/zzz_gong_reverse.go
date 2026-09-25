// generated code - do not edit
package models

// insertion point
func (inst *Button) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "Buttons":
			if _table, ok := stage.Table_Buttons_reverseMap[inst]; ok {
				res = _table.Name
			}
		}
	}
	return
}

func (inst *Cell) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}

func (inst *CellBoolean) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CellFloat64) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CellIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CellInt) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CellString) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DisplayedColumn) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Row) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "Rows":
			if _table, ok := stage.Table_Rows_reverseMap[inst]; ok {
				res = _table.Name
			}
		case "RowsSelectedForBulkDelete":
			if _table, ok := stage.Table_RowsSelectedForBulkDelete_reverseMap[inst]; ok {
				res = _table.Name
			}
		}
	}
	return
}

func (inst *SVGIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Table) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
