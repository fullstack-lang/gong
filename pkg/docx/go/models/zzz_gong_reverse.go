// generated code - do not edit
package models

// insertion point
func (inst *Body) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Document) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Docx) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *File) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Docx":
		switch reverseField.Fieldname {
		case "Files":
			if _docx, ok := stage.Docx_Files_reverseMap[inst]; ok {
				res = _docx.Name
			}
		}
	}
	return
}

func (inst *Node) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Node":
		switch reverseField.Fieldname {
		case "Nodes":
			if _node, ok := stage.Node_Nodes_reverseMap[inst]; ok {
				res = _node.Name
			}
		}
	}
	return
}

func (inst *Paragraph) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Body":
		switch reverseField.Fieldname {
		case "Paragraphs":
			if _body, ok := stage.Body_Paragraphs_reverseMap[inst]; ok {
				res = _body.Name
			}
		}
	case "TableColumn":
		switch reverseField.Fieldname {
		case "Paragraphs":
			if _tablecolumn, ok := stage.TableColumn_Paragraphs_reverseMap[inst]; ok {
				res = _tablecolumn.Name
			}
		}
	}
	return
}

func (inst *ParagraphProperties) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ParagraphStyle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Rune) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Paragraph":
		switch reverseField.Fieldname {
		case "Runes":
			if _paragraph, ok := stage.Paragraph_Runes_reverseMap[inst]; ok {
				res = _paragraph.Name
			}
		}
	}
	return
}

func (inst *RuneProperties) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
	case "Body":
		switch reverseField.Fieldname {
		case "Tables":
			if _body, ok := stage.Body_Tables_reverseMap[inst]; ok {
				res = _body.Name
			}
		}
	}
	return
}

func (inst *TableColumn) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TableRow":
		switch reverseField.Fieldname {
		case "TableColumns":
			if _tablerow, ok := stage.TableRow_TableColumns_reverseMap[inst]; ok {
				res = _tablerow.Name
			}
		}
	}
	return
}

func (inst *TableProperties) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *TableRow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "TableRows":
			if _table, ok := stage.Table_TableRows_reverseMap[inst]; ok {
				res = _table.Name
			}
		}
	}
	return
}

func (inst *TableStyle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Text) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *Body) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Document) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Docx) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *File) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Docx":
		switch reverseField.Fieldname {
		case "Files":
			res = stage.Docx_Files_reverseMap[inst]
		}
	}
	return res
}

func (inst *Node) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Node":
		switch reverseField.Fieldname {
		case "Nodes":
			res = stage.Node_Nodes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Paragraph) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Body":
		switch reverseField.Fieldname {
		case "Paragraphs":
			res = stage.Body_Paragraphs_reverseMap[inst]
		}
	case "TableColumn":
		switch reverseField.Fieldname {
		case "Paragraphs":
			res = stage.TableColumn_Paragraphs_reverseMap[inst]
		}
	}
	return res
}

func (inst *ParagraphProperties) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ParagraphStyle) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Rune) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Paragraph":
		switch reverseField.Fieldname {
		case "Runes":
			res = stage.Paragraph_Runes_reverseMap[inst]
		}
	}
	return res
}

func (inst *RuneProperties) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

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
	case "Body":
		switch reverseField.Fieldname {
		case "Tables":
			res = stage.Body_Tables_reverseMap[inst]
		}
	}
	return res
}

func (inst *TableColumn) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "TableRow":
		switch reverseField.Fieldname {
		case "TableColumns":
			res = stage.TableRow_TableColumns_reverseMap[inst]
		}
	}
	return res
}

func (inst *TableProperties) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *TableRow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Table":
		switch reverseField.Fieldname {
		case "TableRows":
			res = stage.Table_TableRows_reverseMap[inst]
		}
	}
	return res
}

func (inst *TableStyle) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Text) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
