// generated code - do not edit
package models

// insertion point
func (inst *Button) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Cell) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *CellBoolean) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *CellFloat64) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *CellIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *CellInt) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *CellString) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *DisplayedColumn) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Row) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *SVGIcon) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Table) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *Button) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "Buttons":
			res = stage.Table_Buttons_reverseMap[inst]
		}
	}
	return res
}

func (inst *Cell) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Row":
		switch reverseField.Fieldname {
		case "Cells":
			res = stage.Row_Cells_reverseMap[inst]
		}
	}
	return res
}

func (inst *CellBoolean) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *CellFloat64) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *CellIcon) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *CellInt) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *CellString) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *DisplayedColumn) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "DisplayedColumns":
			res = stage.Table_DisplayedColumns_reverseMap[inst]
		}
	}
	return res
}

func (inst *Row) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "Rows":
			res = stage.Table_Rows_reverseMap[inst]
		case "RowsSelectedForBulkDelete":
			res = stage.Table_RowsSelectedForBulkDelete_reverseMap[inst]
		}
	}
	return res
}

func (inst *SVGIcon) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Table) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
