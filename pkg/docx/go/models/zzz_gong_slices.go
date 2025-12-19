// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Body
	// insertion point per field
	stage.Body_Paragraphs_reverseMap = make(map[*Paragraph]*Body)
	for body := range stage.Bodys {
		_ = body
		for _, _paragraph := range body.Paragraphs {
			stage.Body_Paragraphs_reverseMap[_paragraph] = body
		}
	}
	stage.Body_Tables_reverseMap = make(map[*Table]*Body)
	for body := range stage.Bodys {
		_ = body
		for _, _table := range body.Tables {
			stage.Body_Tables_reverseMap[_table] = body
		}
	}

	// Compute reverse map for named struct Document
	// insertion point per field

	// Compute reverse map for named struct Docx
	// insertion point per field
	stage.Docx_Files_reverseMap = make(map[*File]*Docx)
	for docx := range stage.Docxs {
		_ = docx
		for _, _file := range docx.Files {
			stage.Docx_Files_reverseMap[_file] = docx
		}
	}

	// Compute reverse map for named struct File
	// insertion point per field

	// Compute reverse map for named struct Node
	// insertion point per field
	stage.Node_Nodes_reverseMap = make(map[*Node]*Node)
	for node := range stage.Nodes {
		_ = node
		for _, _node := range node.Nodes {
			stage.Node_Nodes_reverseMap[_node] = node
		}
	}

	// Compute reverse map for named struct Paragraph
	// insertion point per field
	stage.Paragraph_Runes_reverseMap = make(map[*Rune]*Paragraph)
	for paragraph := range stage.Paragraphs {
		_ = paragraph
		for _, _rune := range paragraph.Runes {
			stage.Paragraph_Runes_reverseMap[_rune] = paragraph
		}
	}

	// Compute reverse map for named struct ParagraphProperties
	// insertion point per field

	// Compute reverse map for named struct ParagraphStyle
	// insertion point per field

	// Compute reverse map for named struct Rune
	// insertion point per field

	// Compute reverse map for named struct RuneProperties
	// insertion point per field

	// Compute reverse map for named struct Table
	// insertion point per field
	stage.Table_TableRows_reverseMap = make(map[*TableRow]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _tablerow := range table.TableRows {
			stage.Table_TableRows_reverseMap[_tablerow] = table
		}
	}

	// Compute reverse map for named struct TableColumn
	// insertion point per field
	stage.TableColumn_Paragraphs_reverseMap = make(map[*Paragraph]*TableColumn)
	for tablecolumn := range stage.TableColumns {
		_ = tablecolumn
		for _, _paragraph := range tablecolumn.Paragraphs {
			stage.TableColumn_Paragraphs_reverseMap[_paragraph] = tablecolumn
		}
	}

	// Compute reverse map for named struct TableProperties
	// insertion point per field

	// Compute reverse map for named struct TableRow
	// insertion point per field
	stage.TableRow_TableColumns_reverseMap = make(map[*TableColumn]*TableRow)
	for tablerow := range stage.TableRows {
		_ = tablerow
		for _, _tablecolumn := range tablerow.TableColumns {
			stage.TableRow_TableColumns_reverseMap[_tablecolumn] = tablerow
		}
	}

	// Compute reverse map for named struct TableStyle
	// insertion point per field

	// Compute reverse map for named struct Text
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Bodys {
		res = append(res, instance)
	}

	for instance := range stage.Documents {
		res = append(res, instance)
	}

	for instance := range stage.Docxs {
		res = append(res, instance)
	}

	for instance := range stage.Files {
		res = append(res, instance)
	}

	for instance := range stage.Nodes {
		res = append(res, instance)
	}

	for instance := range stage.Paragraphs {
		res = append(res, instance)
	}

	for instance := range stage.ParagraphPropertiess {
		res = append(res, instance)
	}

	for instance := range stage.ParagraphStyles {
		res = append(res, instance)
	}

	for instance := range stage.Runes {
		res = append(res, instance)
	}

	for instance := range stage.RunePropertiess {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

	for instance := range stage.TableColumns {
		res = append(res, instance)
	}

	for instance := range stage.TablePropertiess {
		res = append(res, instance)
	}

	for instance := range stage.TableRows {
		res = append(res, instance)
	}

	for instance := range stage.TableStyles {
		res = append(res, instance)
	}

	for instance := range stage.Texts {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (body *Body) GongCopy() GongstructIF {
	newInstance := *body
	return &newInstance
}

func (document *Document) GongCopy() GongstructIF {
	newInstance := *document
	return &newInstance
}

func (docx *Docx) GongCopy() GongstructIF {
	newInstance := *docx
	return &newInstance
}

func (file *File) GongCopy() GongstructIF {
	newInstance := *file
	return &newInstance
}

func (node *Node) GongCopy() GongstructIF {
	newInstance := *node
	return &newInstance
}

func (paragraph *Paragraph) GongCopy() GongstructIF {
	newInstance := *paragraph
	return &newInstance
}

func (paragraphproperties *ParagraphProperties) GongCopy() GongstructIF {
	newInstance := *paragraphproperties
	return &newInstance
}

func (paragraphstyle *ParagraphStyle) GongCopy() GongstructIF {
	newInstance := *paragraphstyle
	return &newInstance
}

func (rune *Rune) GongCopy() GongstructIF {
	newInstance := *rune
	return &newInstance
}

func (runeproperties *RuneProperties) GongCopy() GongstructIF {
	newInstance := *runeproperties
	return &newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := *table
	return &newInstance
}

func (tablecolumn *TableColumn) GongCopy() GongstructIF {
	newInstance := *tablecolumn
	return &newInstance
}

func (tableproperties *TableProperties) GongCopy() GongstructIF {
	newInstance := *tableproperties
	return &newInstance
}

func (tablerow *TableRow) GongCopy() GongstructIF {
	newInstance := *tablerow
	return &newInstance
}

func (tablestyle *TableStyle) GongCopy() GongstructIF {
	newInstance := *tablestyle
	return &newInstance
}

func (text *Text) GongCopy() GongstructIF {
	newInstance := *text
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
}
