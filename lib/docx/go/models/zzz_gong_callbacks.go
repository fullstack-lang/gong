// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (body *Body) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBodyCreateCallback != nil {
		stage.OnAfterBodyCreateCallback.OnAfterCreate(stage, body)
	}
}

func (body *Body) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBodyUpdateCallback != nil {
		var frontBody *Body
		if front != nil {
			frontBody, _ = front.(*Body)
		}
		stage.OnAfterBodyUpdateCallback.OnAfterUpdate(stage, body, frontBody)
	}
}

func (body *Body) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBodyDeleteCallback != nil {
		var frontBody *Body
		if front != nil {
			frontBody, _ = front.(*Body)
		}
		stage.OnAfterBodyDeleteCallback.OnAfterDelete(stage, body, frontBody)
	}
}

func (document *Document) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDocumentCreateCallback != nil {
		stage.OnAfterDocumentCreateCallback.OnAfterCreate(stage, document)
	}
}

func (document *Document) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentUpdateCallback != nil {
		var frontDocument *Document
		if front != nil {
			frontDocument, _ = front.(*Document)
		}
		stage.OnAfterDocumentUpdateCallback.OnAfterUpdate(stage, document, frontDocument)
	}
}

func (document *Document) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentDeleteCallback != nil {
		var frontDocument *Document
		if front != nil {
			frontDocument, _ = front.(*Document)
		}
		stage.OnAfterDocumentDeleteCallback.OnAfterDelete(stage, document, frontDocument)
	}
}

func (docx *Docx) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDocxCreateCallback != nil {
		stage.OnAfterDocxCreateCallback.OnAfterCreate(stage, docx)
	}
}

func (docx *Docx) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocxUpdateCallback != nil {
		var frontDocx *Docx
		if front != nil {
			frontDocx, _ = front.(*Docx)
		}
		stage.OnAfterDocxUpdateCallback.OnAfterUpdate(stage, docx, frontDocx)
	}
}

func (docx *Docx) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocxDeleteCallback != nil {
		var frontDocx *Docx
		if front != nil {
			frontDocx, _ = front.(*Docx)
		}
		stage.OnAfterDocxDeleteCallback.OnAfterDelete(stage, docx, frontDocx)
	}
}

func (file *File) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFileCreateCallback != nil {
		stage.OnAfterFileCreateCallback.OnAfterCreate(stage, file)
	}
}

func (file *File) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileUpdateCallback != nil {
		var frontFile *File
		if front != nil {
			frontFile, _ = front.(*File)
		}
		stage.OnAfterFileUpdateCallback.OnAfterUpdate(stage, file, frontFile)
	}
}

func (file *File) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileDeleteCallback != nil {
		var frontFile *File
		if front != nil {
			frontFile, _ = front.(*File)
		}
		stage.OnAfterFileDeleteCallback.OnAfterDelete(stage, file, frontFile)
	}
}

func (node *Node) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNodeCreateCallback != nil {
		stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, node)
	}
}

func (node *Node) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeUpdateCallback != nil {
		var frontNode *Node
		if front != nil {
			frontNode, _ = front.(*Node)
		}
		stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, node, frontNode)
	}
}

func (node *Node) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeDeleteCallback != nil {
		var frontNode *Node
		if front != nil {
			frontNode, _ = front.(*Node)
		}
		stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, node, frontNode)
	}
}

func (paragraph *Paragraph) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParagraphCreateCallback != nil {
		stage.OnAfterParagraphCreateCallback.OnAfterCreate(stage, paragraph)
	}
}

func (paragraph *Paragraph) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphUpdateCallback != nil {
		var frontParagraph *Paragraph
		if front != nil {
			frontParagraph, _ = front.(*Paragraph)
		}
		stage.OnAfterParagraphUpdateCallback.OnAfterUpdate(stage, paragraph, frontParagraph)
	}
}

func (paragraph *Paragraph) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphDeleteCallback != nil {
		var frontParagraph *Paragraph
		if front != nil {
			frontParagraph, _ = front.(*Paragraph)
		}
		stage.OnAfterParagraphDeleteCallback.OnAfterDelete(stage, paragraph, frontParagraph)
	}
}

func (paragraphproperties *ParagraphProperties) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParagraphPropertiesCreateCallback != nil {
		stage.OnAfterParagraphPropertiesCreateCallback.OnAfterCreate(stage, paragraphproperties)
	}
}

func (paragraphproperties *ParagraphProperties) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphPropertiesUpdateCallback != nil {
		var frontParagraphProperties *ParagraphProperties
		if front != nil {
			frontParagraphProperties, _ = front.(*ParagraphProperties)
		}
		stage.OnAfterParagraphPropertiesUpdateCallback.OnAfterUpdate(stage, paragraphproperties, frontParagraphProperties)
	}
}

func (paragraphproperties *ParagraphProperties) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphPropertiesDeleteCallback != nil {
		var frontParagraphProperties *ParagraphProperties
		if front != nil {
			frontParagraphProperties, _ = front.(*ParagraphProperties)
		}
		stage.OnAfterParagraphPropertiesDeleteCallback.OnAfterDelete(stage, paragraphproperties, frontParagraphProperties)
	}
}

func (paragraphstyle *ParagraphStyle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParagraphStyleCreateCallback != nil {
		stage.OnAfterParagraphStyleCreateCallback.OnAfterCreate(stage, paragraphstyle)
	}
}

func (paragraphstyle *ParagraphStyle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphStyleUpdateCallback != nil {
		var frontParagraphStyle *ParagraphStyle
		if front != nil {
			frontParagraphStyle, _ = front.(*ParagraphStyle)
		}
		stage.OnAfterParagraphStyleUpdateCallback.OnAfterUpdate(stage, paragraphstyle, frontParagraphStyle)
	}
}

func (paragraphstyle *ParagraphStyle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParagraphStyleDeleteCallback != nil {
		var frontParagraphStyle *ParagraphStyle
		if front != nil {
			frontParagraphStyle, _ = front.(*ParagraphStyle)
		}
		stage.OnAfterParagraphStyleDeleteCallback.OnAfterDelete(stage, paragraphstyle, frontParagraphStyle)
	}
}

func (rune *Rune) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRuneCreateCallback != nil {
		stage.OnAfterRuneCreateCallback.OnAfterCreate(stage, rune)
	}
}

func (rune *Rune) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRuneUpdateCallback != nil {
		var frontRune *Rune
		if front != nil {
			frontRune, _ = front.(*Rune)
		}
		stage.OnAfterRuneUpdateCallback.OnAfterUpdate(stage, rune, frontRune)
	}
}

func (rune *Rune) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRuneDeleteCallback != nil {
		var frontRune *Rune
		if front != nil {
			frontRune, _ = front.(*Rune)
		}
		stage.OnAfterRuneDeleteCallback.OnAfterDelete(stage, rune, frontRune)
	}
}

func (runeproperties *RuneProperties) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunePropertiesCreateCallback != nil {
		stage.OnAfterRunePropertiesCreateCallback.OnAfterCreate(stage, runeproperties)
	}
}

func (runeproperties *RuneProperties) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunePropertiesUpdateCallback != nil {
		var frontRuneProperties *RuneProperties
		if front != nil {
			frontRuneProperties, _ = front.(*RuneProperties)
		}
		stage.OnAfterRunePropertiesUpdateCallback.OnAfterUpdate(stage, runeproperties, frontRuneProperties)
	}
}

func (runeproperties *RuneProperties) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunePropertiesDeleteCallback != nil {
		var frontRuneProperties *RuneProperties
		if front != nil {
			frontRuneProperties, _ = front.(*RuneProperties)
		}
		stage.OnAfterRunePropertiesDeleteCallback.OnAfterDelete(stage, runeproperties, frontRuneProperties)
	}
}

func (table *Table) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableCreateCallback != nil {
		stage.OnAfterTableCreateCallback.OnAfterCreate(stage, table)
	}
}

func (table *Table) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableUpdateCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, table, frontTable)
	}
}

func (table *Table) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableDeleteCallback != nil {
		var frontTable *Table
		if front != nil {
			frontTable, _ = front.(*Table)
		}
		stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, table, frontTable)
	}
}

func (tablecolumn *TableColumn) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableColumnCreateCallback != nil {
		stage.OnAfterTableColumnCreateCallback.OnAfterCreate(stage, tablecolumn)
	}
}

func (tablecolumn *TableColumn) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableColumnUpdateCallback != nil {
		var frontTableColumn *TableColumn
		if front != nil {
			frontTableColumn, _ = front.(*TableColumn)
		}
		stage.OnAfterTableColumnUpdateCallback.OnAfterUpdate(stage, tablecolumn, frontTableColumn)
	}
}

func (tablecolumn *TableColumn) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableColumnDeleteCallback != nil {
		var frontTableColumn *TableColumn
		if front != nil {
			frontTableColumn, _ = front.(*TableColumn)
		}
		stage.OnAfterTableColumnDeleteCallback.OnAfterDelete(stage, tablecolumn, frontTableColumn)
	}
}

func (tableproperties *TableProperties) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTablePropertiesCreateCallback != nil {
		stage.OnAfterTablePropertiesCreateCallback.OnAfterCreate(stage, tableproperties)
	}
}

func (tableproperties *TableProperties) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTablePropertiesUpdateCallback != nil {
		var frontTableProperties *TableProperties
		if front != nil {
			frontTableProperties, _ = front.(*TableProperties)
		}
		stage.OnAfterTablePropertiesUpdateCallback.OnAfterUpdate(stage, tableproperties, frontTableProperties)
	}
}

func (tableproperties *TableProperties) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTablePropertiesDeleteCallback != nil {
		var frontTableProperties *TableProperties
		if front != nil {
			frontTableProperties, _ = front.(*TableProperties)
		}
		stage.OnAfterTablePropertiesDeleteCallback.OnAfterDelete(stage, tableproperties, frontTableProperties)
	}
}

func (tablerow *TableRow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableRowCreateCallback != nil {
		stage.OnAfterTableRowCreateCallback.OnAfterCreate(stage, tablerow)
	}
}

func (tablerow *TableRow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableRowUpdateCallback != nil {
		var frontTableRow *TableRow
		if front != nil {
			frontTableRow, _ = front.(*TableRow)
		}
		stage.OnAfterTableRowUpdateCallback.OnAfterUpdate(stage, tablerow, frontTableRow)
	}
}

func (tablerow *TableRow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableRowDeleteCallback != nil {
		var frontTableRow *TableRow
		if front != nil {
			frontTableRow, _ = front.(*TableRow)
		}
		stage.OnAfterTableRowDeleteCallback.OnAfterDelete(stage, tablerow, frontTableRow)
	}
}

func (tablestyle *TableStyle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTableStyleCreateCallback != nil {
		stage.OnAfterTableStyleCreateCallback.OnAfterCreate(stage, tablestyle)
	}
}

func (tablestyle *TableStyle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableStyleUpdateCallback != nil {
		var frontTableStyle *TableStyle
		if front != nil {
			frontTableStyle, _ = front.(*TableStyle)
		}
		stage.OnAfterTableStyleUpdateCallback.OnAfterUpdate(stage, tablestyle, frontTableStyle)
	}
}

func (tablestyle *TableStyle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTableStyleDeleteCallback != nil {
		var frontTableStyle *TableStyle
		if front != nil {
			frontTableStyle, _ = front.(*TableStyle)
		}
		stage.OnAfterTableStyleDeleteCallback.OnAfterDelete(stage, tablestyle, frontTableStyle)
	}
}

func (text *Text) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextCreateCallback != nil {
		stage.OnAfterTextCreateCallback.OnAfterCreate(stage, text)
	}
}

func (text *Text) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextUpdateCallback != nil {
		var frontText *Text
		if front != nil {
			frontText, _ = front.(*Text)
		}
		stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, text, frontText)
	}
}

func (text *Text) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDeleteCallback != nil {
		var frontText *Text
		if front != nil {
			frontText, _ = front.(*Text)
		}
		stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, text, frontText)
	}
}

