// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Body
func (body *Body) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &body.Paragraphs) || modified
	modified = GongCleanSlice(stage, &body.Tables) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &body.LastParagraph) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Document
func (document *Document) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &document.File) || modified
	modified = GongCleanPointer(stage, &document.Root) || modified
	modified = GongCleanPointer(stage, &document.Body) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Docx
func (docx *Docx) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &docx.Files) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &docx.Document) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by File
func (file *File) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Node
func (node *Node) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &node.Nodes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Paragraph
func (paragraph *Paragraph) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &paragraph.Runes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &paragraph.Node) || modified
	modified = GongCleanPointer(stage, &paragraph.ParagraphProperties) || modified
	modified = GongCleanPointer(stage, &paragraph.Next) || modified
	modified = GongCleanPointer(stage, &paragraph.Previous) || modified
	modified = GongCleanPointer(stage, &paragraph.EnclosingBody) || modified
	modified = GongCleanPointer(stage, &paragraph.EnclosingTableColumn) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParagraphProperties
func (paragraphproperties *ParagraphProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &paragraphproperties.ParagraphStyle) || modified
	modified = GongCleanPointer(stage, &paragraphproperties.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParagraphStyle
func (paragraphstyle *ParagraphStyle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &paragraphstyle.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Rune
func (rune *Rune) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &rune.Node) || modified
	modified = GongCleanPointer(stage, &rune.Text) || modified
	modified = GongCleanPointer(stage, &rune.RuneProperties) || modified
	modified = GongCleanPointer(stage, &rune.EnclosingParagraph) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RuneProperties
func (runeproperties *RuneProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &runeproperties.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &table.TableRows) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &table.Node) || modified
	modified = GongCleanPointer(stage, &table.TableProperties) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableColumn
func (tablecolumn *TableColumn) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &tablecolumn.Paragraphs) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &tablecolumn.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableProperties
func (tableproperties *TableProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &tableproperties.Node) || modified
	modified = GongCleanPointer(stage, &tableproperties.TableStyle) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableRow
func (tablerow *TableRow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &tablerow.TableColumns) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &tablerow.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableStyle
func (tablestyle *TableStyle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &tablestyle.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Text
func (text *Text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &text.Node) || modified
	modified = GongCleanPointer(stage, &text.EnclosingRune) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
