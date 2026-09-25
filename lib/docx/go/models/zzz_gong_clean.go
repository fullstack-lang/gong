// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Body
func (body *Body) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&body.Paragraphs) || modified
	modified = stage.CleanSlice(&body.Tables) || modified
	// insertion point per field
	modified = stage.CleanPointer(&body.LastParagraph) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Document
func (document *Document) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&document.File) || modified
	modified = stage.CleanPointer(&document.Root) || modified
	modified = stage.CleanPointer(&document.Body) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Docx
func (docx *Docx) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&docx.Files) || modified
	// insertion point per field
	modified = stage.CleanPointer(&docx.Document) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Node
func (node *Node) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&node.Nodes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Paragraph
func (paragraph *Paragraph) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&paragraph.Runes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&paragraph.Node) || modified
	modified = stage.CleanPointer(&paragraph.ParagraphProperties) || modified
	modified = stage.CleanPointer(&paragraph.Next) || modified
	modified = stage.CleanPointer(&paragraph.Previous) || modified
	modified = stage.CleanPointer(&paragraph.EnclosingBody) || modified
	modified = stage.CleanPointer(&paragraph.EnclosingTableColumn) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParagraphProperties
func (paragraphproperties *ParagraphProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&paragraphproperties.ParagraphStyle) || modified
	modified = stage.CleanPointer(&paragraphproperties.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParagraphStyle
func (paragraphstyle *ParagraphStyle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&paragraphstyle.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Rune
func (rune *Rune) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rune.Node) || modified
	modified = stage.CleanPointer(&rune.Text) || modified
	modified = stage.CleanPointer(&rune.RuneProperties) || modified
	modified = stage.CleanPointer(&rune.EnclosingParagraph) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RuneProperties
func (runeproperties *RuneProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&runeproperties.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&table.TableRows) || modified
	// insertion point per field
	modified = stage.CleanPointer(&table.Node) || modified
	modified = stage.CleanPointer(&table.TableProperties) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableColumn
func (tablecolumn *TableColumn) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tablecolumn.Paragraphs) || modified
	// insertion point per field
	modified = stage.CleanPointer(&tablecolumn.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableProperties
func (tableproperties *TableProperties) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tableproperties.Node) || modified
	modified = stage.CleanPointer(&tableproperties.TableStyle) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableRow
func (tablerow *TableRow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tablerow.TableColumns) || modified
	// insertion point per field
	modified = stage.CleanPointer(&tablerow.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TableStyle
func (tablestyle *TableStyle) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tablestyle.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Text
func (text *Text) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&text.Node) || modified
	modified = stage.CleanPointer(&text.EnclosingRune) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
