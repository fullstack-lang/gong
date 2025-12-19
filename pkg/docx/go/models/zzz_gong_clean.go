// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Body
func (body *Body) GongClean(stage *Stage) {
	// insertion point per field
	body.Paragraphs = GongCleanSlice(stage, body.Paragraphs)
	body.Tables = GongCleanSlice(stage, body.Tables)
	// insertion point per field
	body.LastParagraph = GongCleanPointer(stage, body.LastParagraph)
}

// Clean garbage collect unstaged instances that are referenced by Document
func (document *Document) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	document.File = GongCleanPointer(stage, document.File)
	document.Root = GongCleanPointer(stage, document.Root)
	document.Body = GongCleanPointer(stage, document.Body)
}

// Clean garbage collect unstaged instances that are referenced by Docx
func (docx *Docx) GongClean(stage *Stage) {
	// insertion point per field
	docx.Files = GongCleanSlice(stage, docx.Files)
	// insertion point per field
	docx.Document = GongCleanPointer(stage, docx.Document)
}

// Clean garbage collect unstaged instances that are referenced by File
func (file *File) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Node
func (node *Node) GongClean(stage *Stage) {
	// insertion point per field
	node.Nodes = GongCleanSlice(stage, node.Nodes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Paragraph
func (paragraph *Paragraph) GongClean(stage *Stage) {
	// insertion point per field
	paragraph.Runes = GongCleanSlice(stage, paragraph.Runes)
	// insertion point per field
	paragraph.Node = GongCleanPointer(stage, paragraph.Node)
	paragraph.ParagraphProperties = GongCleanPointer(stage, paragraph.ParagraphProperties)
	paragraph.Next = GongCleanPointer(stage, paragraph.Next)
	paragraph.Previous = GongCleanPointer(stage, paragraph.Previous)
	paragraph.EnclosingBody = GongCleanPointer(stage, paragraph.EnclosingBody)
	paragraph.EnclosingTableColumn = GongCleanPointer(stage, paragraph.EnclosingTableColumn)
}

// Clean garbage collect unstaged instances that are referenced by ParagraphProperties
func (paragraphproperties *ParagraphProperties) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	paragraphproperties.ParagraphStyle = GongCleanPointer(stage, paragraphproperties.ParagraphStyle)
	paragraphproperties.Node = GongCleanPointer(stage, paragraphproperties.Node)
}

// Clean garbage collect unstaged instances that are referenced by ParagraphStyle
func (paragraphstyle *ParagraphStyle) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	paragraphstyle.Node = GongCleanPointer(stage, paragraphstyle.Node)
}

// Clean garbage collect unstaged instances that are referenced by Rune
func (rune *Rune) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	rune.Node = GongCleanPointer(stage, rune.Node)
	rune.Text = GongCleanPointer(stage, rune.Text)
	rune.RuneProperties = GongCleanPointer(stage, rune.RuneProperties)
	rune.EnclosingParagraph = GongCleanPointer(stage, rune.EnclosingParagraph)
}

// Clean garbage collect unstaged instances that are referenced by RuneProperties
func (runeproperties *RuneProperties) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	runeproperties.Node = GongCleanPointer(stage, runeproperties.Node)
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) {
	// insertion point per field
	table.TableRows = GongCleanSlice(stage, table.TableRows)
	// insertion point per field
	table.Node = GongCleanPointer(stage, table.Node)
	table.TableProperties = GongCleanPointer(stage, table.TableProperties)
}

// Clean garbage collect unstaged instances that are referenced by TableColumn
func (tablecolumn *TableColumn) GongClean(stage *Stage) {
	// insertion point per field
	tablecolumn.Paragraphs = GongCleanSlice(stage, tablecolumn.Paragraphs)
	// insertion point per field
	tablecolumn.Node = GongCleanPointer(stage, tablecolumn.Node)
}

// Clean garbage collect unstaged instances that are referenced by TableProperties
func (tableproperties *TableProperties) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	tableproperties.Node = GongCleanPointer(stage, tableproperties.Node)
	tableproperties.TableStyle = GongCleanPointer(stage, tableproperties.TableStyle)
}

// Clean garbage collect unstaged instances that are referenced by TableRow
func (tablerow *TableRow) GongClean(stage *Stage) {
	// insertion point per field
	tablerow.TableColumns = GongCleanSlice(stage, tablerow.TableColumns)
	// insertion point per field
	tablerow.Node = GongCleanPointer(stage, tablerow.Node)
}

// Clean garbage collect unstaged instances that are referenced by TableStyle
func (tablestyle *TableStyle) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	tablestyle.Node = GongCleanPointer(stage, tablestyle.Node)
}

// Clean garbage collect unstaged instances that are referenced by Text
func (text *Text) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	text.Node = GongCleanPointer(stage, text.Node)
	text.EnclosingRune = GongCleanPointer(stage, text.EnclosingRune)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
