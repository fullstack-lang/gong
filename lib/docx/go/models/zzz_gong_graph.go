// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (body *Body) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bodys[body]
	return ok
}

func (document *Document) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Documents[document]
	return ok
}

func (docx *Docx) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Docxs[docx]
	return ok
}

func (file *File) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Files[file]
	return ok
}

func (node *Node) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Nodes[node]
	return ok
}

func (paragraph *Paragraph) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Paragraphs[paragraph]
	return ok
}

func (paragraphproperties *ParagraphProperties) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParagraphPropertiess[paragraphproperties]
	return ok
}

func (paragraphstyle *ParagraphStyle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParagraphStyles[paragraphstyle]
	return ok
}

func (rune *Rune) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Runes[rune]
	return ok
}

func (runeproperties *RuneProperties) GongIsStaged(stage *Stage) bool {
	_, ok := stage.RunePropertiess[runeproperties]
	return ok
}

func (table *Table) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tables[table]
	return ok
}

func (tablecolumn *TableColumn) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TableColumns[tablecolumn]
	return ok
}

func (tableproperties *TableProperties) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TablePropertiess[tableproperties]
	return ok
}

func (tablerow *TableRow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TableRows[tablerow]
	return ok
}

func (tablestyle *TableStyle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TableStyles[tablestyle]
	return ok
}

func (text *Text) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Texts[text]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (body *Body) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(body) {
		return
	}

	body.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		stage.StageBranch(body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		stage.StageBranch(_paragraph)
	}
	for _, _table := range body.Tables {
		stage.StageBranch(_table)
	}

}

func (document *Document) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		stage.StageBranch(document.File)
	}
	if document.Root != nil {
		stage.StageBranch(document.Root)
	}
	if document.Body != nil {
		stage.StageBranch(document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (docx *Docx) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(docx) {
		return
	}

	docx.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		stage.StageBranch(docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		stage.StageBranch(_file)
	}

}

func (file *File) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(file) {
		return
	}

	file.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (node *Node) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(node) {
		return
	}

	node.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		stage.StageBranch(_node)
	}

}

func (paragraph *Paragraph) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(paragraph) {
		return
	}

	paragraph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		stage.StageBranch(paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		stage.StageBranch(paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		stage.StageBranch(paragraph.Next)
	}
	if paragraph.Previous != nil {
		stage.StageBranch(paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		stage.StageBranch(paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		stage.StageBranch(paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		stage.StageBranch(_rune)
	}

}

func (paragraphproperties *ParagraphProperties) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(paragraphproperties) {
		return
	}

	paragraphproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		stage.StageBranch(paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		stage.StageBranch(paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (paragraphstyle *ParagraphStyle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(paragraphstyle) {
		return
	}

	paragraphstyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		stage.StageBranch(paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rune *Rune) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(rune) {
		return
	}

	rune.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		stage.StageBranch(rune.Node)
	}
	if rune.Text != nil {
		stage.StageBranch(rune.Text)
	}
	if rune.RuneProperties != nil {
		stage.StageBranch(rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		stage.StageBranch(rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (runeproperties *RuneProperties) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(runeproperties) {
		return
	}

	runeproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		stage.StageBranch(runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		stage.StageBranch(table.Node)
	}
	if table.TableProperties != nil {
		stage.StageBranch(table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		stage.StageBranch(_tablerow)
	}

}

func (tablecolumn *TableColumn) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tablecolumn) {
		return
	}

	tablecolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		stage.StageBranch(tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		stage.StageBranch(_paragraph)
	}

}

func (tableproperties *TableProperties) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tableproperties) {
		return
	}

	tableproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		stage.StageBranch(tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		stage.StageBranch(tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tablerow *TableRow) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tablerow) {
		return
	}

	tablerow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		stage.StageBranch(tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		stage.StageBranch(_tablecolumn)
	}

}

func (tablestyle *TableStyle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(tablestyle) {
		return
	}

	tablestyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		stage.StageBranch(tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		stage.StageBranch(text.Node)
	}
	if text.EnclosingRune != nil {
		stage.StageBranch(text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Body:
		toT := GongCopyBranchBody(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Document:
		toT := GongCopyBranchDocument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Docx:
		toT := GongCopyBranchDocx(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *File:
		toT := GongCopyBranchFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Node:
		toT := GongCopyBranchNode(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Paragraph:
		toT := GongCopyBranchParagraph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphProperties:
		toT := GongCopyBranchParagraphProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphStyle:
		toT := GongCopyBranchParagraphStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rune:
		toT := GongCopyBranchRune(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RuneProperties:
		toT := GongCopyBranchRuneProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := GongCopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableColumn:
		toT := GongCopyBranchTableColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableProperties:
		toT := GongCopyBranchTableProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableRow:
		toT := GongCopyBranchTableRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableStyle:
		toT := GongCopyBranchTableStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text:
		toT := GongCopyBranchText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchBody(mapOrigCopy map[any]any, bodyFrom *Body) (bodyTo *Body) {
	var alreadyCopied bool
	bodyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bodyFrom)
	if alreadyCopied {
		return
	}
	bodyFrom.GongCopyBasicFields(bodyTo)

	//insertion point for the staging of instances referenced by pointers
	if bodyFrom.LastParagraph != nil {
		bodyTo.LastParagraph = GongCopyBranchParagraph(mapOrigCopy, bodyFrom.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range bodyFrom.Paragraphs {
		bodyTo.Paragraphs = append(bodyTo.Paragraphs, GongCopyBranchParagraph(mapOrigCopy, _paragraph))
	}
	for _, _table := range bodyFrom.Tables {
		bodyTo.Tables = append(bodyTo.Tables, GongCopyBranchTable(mapOrigCopy, _table))
	}

	return
}

func GongCopyBranchDocument(mapOrigCopy map[any]any, documentFrom *Document) (documentTo *Document) {
	var alreadyCopied bool
	documentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, documentFrom)
	if alreadyCopied {
		return
	}
	documentFrom.GongCopyBasicFields(documentTo)

	//insertion point for the staging of instances referenced by pointers
	if documentFrom.File != nil {
		documentTo.File = GongCopyBranchFile(mapOrigCopy, documentFrom.File)
	}
	if documentFrom.Root != nil {
		documentTo.Root = GongCopyBranchNode(mapOrigCopy, documentFrom.Root)
	}
	if documentFrom.Body != nil {
		documentTo.Body = GongCopyBranchBody(mapOrigCopy, documentFrom.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDocx(mapOrigCopy map[any]any, docxFrom *Docx) (docxTo *Docx) {
	var alreadyCopied bool
	docxTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, docxFrom)
	if alreadyCopied {
		return
	}
	docxFrom.GongCopyBasicFields(docxTo)

	//insertion point for the staging of instances referenced by pointers
	if docxFrom.Document != nil {
		docxTo.Document = GongCopyBranchDocument(mapOrigCopy, docxFrom.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docxFrom.Files {
		docxTo.Files = append(docxTo.Files, GongCopyBranchFile(mapOrigCopy, _file))
	}

	return
}

func GongCopyBranchFile(mapOrigCopy map[any]any, fileFrom *File) (fileTo *File) {
	var alreadyCopied bool
	fileTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, fileFrom)
	if alreadyCopied {
		return
	}
	fileFrom.GongCopyBasicFields(fileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {
	var alreadyCopied bool
	nodeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, nodeFrom)
	if alreadyCopied {
		return
	}
	nodeFrom.GongCopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Nodes {
		nodeTo.Nodes = append(nodeTo.Nodes, GongCopyBranchNode(mapOrigCopy, _node))
	}

	return
}

func GongCopyBranchParagraph(mapOrigCopy map[any]any, paragraphFrom *Paragraph) (paragraphTo *Paragraph) {
	var alreadyCopied bool
	paragraphTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, paragraphFrom)
	if alreadyCopied {
		return
	}
	paragraphFrom.GongCopyBasicFields(paragraphTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphFrom.Node != nil {
		paragraphTo.Node = GongCopyBranchNode(mapOrigCopy, paragraphFrom.Node)
	}
	if paragraphFrom.ParagraphProperties != nil {
		paragraphTo.ParagraphProperties = GongCopyBranchParagraphProperties(mapOrigCopy, paragraphFrom.ParagraphProperties)
	}
	if paragraphFrom.Next != nil {
		paragraphTo.Next = GongCopyBranchParagraph(mapOrigCopy, paragraphFrom.Next)
	}
	if paragraphFrom.Previous != nil {
		paragraphTo.Previous = GongCopyBranchParagraph(mapOrigCopy, paragraphFrom.Previous)
	}
	if paragraphFrom.EnclosingBody != nil {
		paragraphTo.EnclosingBody = GongCopyBranchBody(mapOrigCopy, paragraphFrom.EnclosingBody)
	}
	if paragraphFrom.EnclosingTableColumn != nil {
		paragraphTo.EnclosingTableColumn = GongCopyBranchTableColumn(mapOrigCopy, paragraphFrom.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraphFrom.Runes {
		paragraphTo.Runes = append(paragraphTo.Runes, GongCopyBranchRune(mapOrigCopy, _rune))
	}

	return
}

func GongCopyBranchParagraphProperties(mapOrigCopy map[any]any, paragraphpropertiesFrom *ParagraphProperties) (paragraphpropertiesTo *ParagraphProperties) {
	var alreadyCopied bool
	paragraphpropertiesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, paragraphpropertiesFrom)
	if alreadyCopied {
		return
	}
	paragraphpropertiesFrom.GongCopyBasicFields(paragraphpropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphpropertiesFrom.ParagraphStyle != nil {
		paragraphpropertiesTo.ParagraphStyle = GongCopyBranchParagraphStyle(mapOrigCopy, paragraphpropertiesFrom.ParagraphStyle)
	}
	if paragraphpropertiesFrom.Node != nil {
		paragraphpropertiesTo.Node = GongCopyBranchNode(mapOrigCopy, paragraphpropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParagraphStyle(mapOrigCopy map[any]any, paragraphstyleFrom *ParagraphStyle) (paragraphstyleTo *ParagraphStyle) {
	var alreadyCopied bool
	paragraphstyleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, paragraphstyleFrom)
	if alreadyCopied {
		return
	}
	paragraphstyleFrom.GongCopyBasicFields(paragraphstyleTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyleFrom.Node != nil {
		paragraphstyleTo.Node = GongCopyBranchNode(mapOrigCopy, paragraphstyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRune(mapOrigCopy map[any]any, runeFrom *Rune) (runeTo *Rune) {
	var alreadyCopied bool
	runeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, runeFrom)
	if alreadyCopied {
		return
	}
	runeFrom.GongCopyBasicFields(runeTo)

	//insertion point for the staging of instances referenced by pointers
	if runeFrom.Node != nil {
		runeTo.Node = GongCopyBranchNode(mapOrigCopy, runeFrom.Node)
	}
	if runeFrom.Text != nil {
		runeTo.Text = GongCopyBranchText(mapOrigCopy, runeFrom.Text)
	}
	if runeFrom.RuneProperties != nil {
		runeTo.RuneProperties = GongCopyBranchRuneProperties(mapOrigCopy, runeFrom.RuneProperties)
	}
	if runeFrom.EnclosingParagraph != nil {
		runeTo.EnclosingParagraph = GongCopyBranchParagraph(mapOrigCopy, runeFrom.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRuneProperties(mapOrigCopy map[any]any, runepropertiesFrom *RuneProperties) (runepropertiesTo *RuneProperties) {
	var alreadyCopied bool
	runepropertiesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, runepropertiesFrom)
	if alreadyCopied {
		return
	}
	runepropertiesFrom.GongCopyBasicFields(runepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if runepropertiesFrom.Node != nil {
		runepropertiesTo.Node = GongCopyBranchNode(mapOrigCopy, runepropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {
	var alreadyCopied bool
	tableTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tableFrom)
	if alreadyCopied {
		return
	}
	tableFrom.GongCopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers
	if tableFrom.Node != nil {
		tableTo.Node = GongCopyBranchNode(mapOrigCopy, tableFrom.Node)
	}
	if tableFrom.TableProperties != nil {
		tableTo.TableProperties = GongCopyBranchTableProperties(mapOrigCopy, tableFrom.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range tableFrom.TableRows {
		tableTo.TableRows = append(tableTo.TableRows, GongCopyBranchTableRow(mapOrigCopy, _tablerow))
	}

	return
}

func GongCopyBranchTableColumn(mapOrigCopy map[any]any, tablecolumnFrom *TableColumn) (tablecolumnTo *TableColumn) {
	var alreadyCopied bool
	tablecolumnTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tablecolumnFrom)
	if alreadyCopied {
		return
	}
	tablecolumnFrom.GongCopyBasicFields(tablecolumnTo)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumnFrom.Node != nil {
		tablecolumnTo.Node = GongCopyBranchNode(mapOrigCopy, tablecolumnFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumnFrom.Paragraphs {
		tablecolumnTo.Paragraphs = append(tablecolumnTo.Paragraphs, GongCopyBranchParagraph(mapOrigCopy, _paragraph))
	}

	return
}

func GongCopyBranchTableProperties(mapOrigCopy map[any]any, tablepropertiesFrom *TableProperties) (tablepropertiesTo *TableProperties) {
	var alreadyCopied bool
	tablepropertiesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tablepropertiesFrom)
	if alreadyCopied {
		return
	}
	tablepropertiesFrom.GongCopyBasicFields(tablepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if tablepropertiesFrom.Node != nil {
		tablepropertiesTo.Node = GongCopyBranchNode(mapOrigCopy, tablepropertiesFrom.Node)
	}
	if tablepropertiesFrom.TableStyle != nil {
		tablepropertiesTo.TableStyle = GongCopyBranchTableStyle(mapOrigCopy, tablepropertiesFrom.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTableRow(mapOrigCopy map[any]any, tablerowFrom *TableRow) (tablerowTo *TableRow) {
	var alreadyCopied bool
	tablerowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tablerowFrom)
	if alreadyCopied {
		return
	}
	tablerowFrom.GongCopyBasicFields(tablerowTo)

	//insertion point for the staging of instances referenced by pointers
	if tablerowFrom.Node != nil {
		tablerowTo.Node = GongCopyBranchNode(mapOrigCopy, tablerowFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerowFrom.TableColumns {
		tablerowTo.TableColumns = append(tablerowTo.TableColumns, GongCopyBranchTableColumn(mapOrigCopy, _tablecolumn))
	}

	return
}

func GongCopyBranchTableStyle(mapOrigCopy map[any]any, tablestyleFrom *TableStyle) (tablestyleTo *TableStyle) {
	var alreadyCopied bool
	tablestyleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tablestyleFrom)
	if alreadyCopied {
		return
	}
	tablestyleFrom.GongCopyBasicFields(tablestyleTo)

	//insertion point for the staging of instances referenced by pointers
	if tablestyleFrom.Node != nil {
		tablestyleTo.Node = GongCopyBranchNode(mapOrigCopy, tablestyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {
	var alreadyCopied bool
	textTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textFrom)
	if alreadyCopied {
		return
	}
	textFrom.GongCopyBasicFields(textTo)

	//insertion point for the staging of instances referenced by pointers
	if textFrom.Node != nil {
		textTo.Node = GongCopyBranchNode(mapOrigCopy, textFrom.Node)
	}
	if textFrom.EnclosingRune != nil {
		textTo.EnclosingRune = GongCopyBranchRune(mapOrigCopy, textFrom.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (body *Body) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(body) {
		return
	}

	body.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		stage.UnstageBranch(body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		stage.UnstageBranch(_paragraph)
	}
	for _, _table := range body.Tables {
		stage.UnstageBranch(_table)
	}

}

func (document *Document) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		stage.UnstageBranch(document.File)
	}
	if document.Root != nil {
		stage.UnstageBranch(document.Root)
	}
	if document.Body != nil {
		stage.UnstageBranch(document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (docx *Docx) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(docx) {
		return
	}

	docx.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		stage.UnstageBranch(docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		stage.UnstageBranch(_file)
	}

}

func (file *File) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(file) {
		return
	}

	file.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (node *Node) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(node) {
		return
	}

	node.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		stage.UnstageBranch(_node)
	}

}

func (paragraph *Paragraph) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(paragraph) {
		return
	}

	paragraph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		stage.UnstageBranch(paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		stage.UnstageBranch(paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		stage.UnstageBranch(paragraph.Next)
	}
	if paragraph.Previous != nil {
		stage.UnstageBranch(paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		stage.UnstageBranch(paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		stage.UnstageBranch(paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		stage.UnstageBranch(_rune)
	}

}

func (paragraphproperties *ParagraphProperties) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(paragraphproperties) {
		return
	}

	paragraphproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		stage.UnstageBranch(paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		stage.UnstageBranch(paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (paragraphstyle *ParagraphStyle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(paragraphstyle) {
		return
	}

	paragraphstyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		stage.UnstageBranch(paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (rune *Rune) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(rune) {
		return
	}

	rune.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		stage.UnstageBranch(rune.Node)
	}
	if rune.Text != nil {
		stage.UnstageBranch(rune.Text)
	}
	if rune.RuneProperties != nil {
		stage.UnstageBranch(rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		stage.UnstageBranch(rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (runeproperties *RuneProperties) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(runeproperties) {
		return
	}

	runeproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		stage.UnstageBranch(runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		stage.UnstageBranch(table.Node)
	}
	if table.TableProperties != nil {
		stage.UnstageBranch(table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		stage.UnstageBranch(_tablerow)
	}

}

func (tablecolumn *TableColumn) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tablecolumn) {
		return
	}

	tablecolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		stage.UnstageBranch(tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		stage.UnstageBranch(_paragraph)
	}

}

func (tableproperties *TableProperties) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tableproperties) {
		return
	}

	tableproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		stage.UnstageBranch(tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		stage.UnstageBranch(tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (tablerow *TableRow) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tablerow) {
		return
	}

	tablerow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		stage.UnstageBranch(tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		stage.UnstageBranch(_tablecolumn)
	}

}

func (tablestyle *TableStyle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(tablestyle) {
		return
	}

	tablestyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		stage.UnstageBranch(tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (text *Text) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		stage.UnstageBranch(text.Node)
	}
	if text.EnclosingRune != nil {
		stage.UnstageBranch(text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Body) GongReconstructPointersFromReferences(stage *Stage, instance *Body) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LastParagraph, stage.Paragraphs_reference, instance.LastParagraph)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Paragraphs, stage.Paragraphs_reference, instance.Paragraphs)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tables, stage.Tables_reference, instance.Tables)
}

func (reference *Document) GongReconstructPointersFromReferences(stage *Stage, instance *Document) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.File, stage.Files_reference, instance.File)
	__gong__reconstructPointer(&reference.Root, stage.Nodes_reference, instance.Root)
	__gong__reconstructPointer(&reference.Body, stage.Bodys_reference, instance.Body)
	// insertion point for slice of pointers field
}

func (reference *Docx) GongReconstructPointersFromReferences(stage *Stage, instance *Docx) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Document, stage.Documents_reference, instance.Document)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Files, stage.Files_reference, instance.Files)
}

func (reference *File) GongReconstructPointersFromReferences(stage *Stage, instance *File) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Node) GongReconstructPointersFromReferences(stage *Stage, instance *Node) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Nodes, stage.Nodes_reference, instance.Nodes)
}

func (reference *Paragraph) GongReconstructPointersFromReferences(stage *Stage, instance *Paragraph) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	__gong__reconstructPointer(&reference.ParagraphProperties, stage.ParagraphPropertiess_reference, instance.ParagraphProperties)
	__gong__reconstructPointer(&reference.Next, stage.Paragraphs_reference, instance.Next)
	__gong__reconstructPointer(&reference.Previous, stage.Paragraphs_reference, instance.Previous)
	__gong__reconstructPointer(&reference.EnclosingBody, stage.Bodys_reference, instance.EnclosingBody)
	__gong__reconstructPointer(&reference.EnclosingTableColumn, stage.TableColumns_reference, instance.EnclosingTableColumn)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Runes, stage.Runes_reference, instance.Runes)
}

func (reference *ParagraphProperties) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphProperties) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ParagraphStyle, stage.ParagraphStyles_reference, instance.ParagraphStyle)
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
}

func (reference *ParagraphStyle) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphStyle) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
}

func (reference *Rune) GongReconstructPointersFromReferences(stage *Stage, instance *Rune) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	__gong__reconstructPointer(&reference.Text, stage.Texts_reference, instance.Text)
	__gong__reconstructPointer(&reference.RuneProperties, stage.RunePropertiess_reference, instance.RuneProperties)
	__gong__reconstructPointer(&reference.EnclosingParagraph, stage.Paragraphs_reference, instance.EnclosingParagraph)
	// insertion point for slice of pointers field
}

func (reference *RuneProperties) GongReconstructPointersFromReferences(stage *Stage, instance *RuneProperties) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	__gong__reconstructPointer(&reference.TableProperties, stage.TablePropertiess_reference, instance.TableProperties)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TableRows, stage.TableRows_reference, instance.TableRows)
}

func (reference *TableColumn) GongReconstructPointersFromReferences(stage *Stage, instance *TableColumn) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Paragraphs, stage.Paragraphs_reference, instance.Paragraphs)
}

func (reference *TableProperties) GongReconstructPointersFromReferences(stage *Stage, instance *TableProperties) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	__gong__reconstructPointer(&reference.TableStyle, stage.TableStyles_reference, instance.TableStyle)
	// insertion point for slice of pointers field
}

func (reference *TableRow) GongReconstructPointersFromReferences(stage *Stage, instance *TableRow) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.TableColumns, stage.TableColumns_reference, instance.TableColumns)
}

func (reference *TableStyle) GongReconstructPointersFromReferences(stage *Stage, instance *TableStyle) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	// insertion point for slice of pointers field
}

func (reference *Text) GongReconstructPointersFromReferences(stage *Stage, instance *Text) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Node, stage.Nodes_reference, instance.Node)
	__gong__reconstructPointer(&reference.EnclosingRune, stage.Runes_reference, instance.EnclosingRune)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Body) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LastParagraph, stage.Paragraphs_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Paragraphs, stage.Paragraphs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tables, stage.Tables_instance)
}

func (reference *Document) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.File, stage.Files_instance)
	__gong__reconstructPointerFromInstance(&reference.Root, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.Body, stage.Bodys_instance)
	// insertion point for slice of pointers fields
}

func (reference *Docx) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Document, stage.Documents_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Files, stage.Files_instance)
}

func (reference *File) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Node) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Nodes, stage.Nodes_instance)
}

func (reference *Paragraph) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.ParagraphProperties, stage.ParagraphPropertiess_instance)
	__gong__reconstructPointerFromInstance(&reference.Next, stage.Paragraphs_instance)
	__gong__reconstructPointerFromInstance(&reference.Previous, stage.Paragraphs_instance)
	__gong__reconstructPointerFromInstance(&reference.EnclosingBody, stage.Bodys_instance)
	__gong__reconstructPointerFromInstance(&reference.EnclosingTableColumn, stage.TableColumns_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Runes, stage.Runes_instance)
}

func (reference *ParagraphProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ParagraphStyle, stage.ParagraphStyles_instance)
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
}

func (reference *ParagraphStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Rune) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.Text, stage.Texts_instance)
	__gong__reconstructPointerFromInstance(&reference.RuneProperties, stage.RunePropertiess_instance)
	__gong__reconstructPointerFromInstance(&reference.EnclosingParagraph, stage.Paragraphs_instance)
	// insertion point for slice of pointers fields
}

func (reference *RuneProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.TableProperties, stage.TablePropertiess_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TableRows, stage.TableRows_instance)
}

func (reference *TableColumn) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Paragraphs, stage.Paragraphs_instance)
}

func (reference *TableProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.TableStyle, stage.TableStyles_instance)
	// insertion point for slice of pointers fields
}

func (reference *TableRow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.TableColumns, stage.TableColumns_instance)
}

func (reference *TableStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Node, stage.Nodes_instance)
	__gong__reconstructPointerFromInstance(&reference.EnclosingRune, stage.Runes_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (body *Body) GongDiff(stage *Stage, bodyOther *Body) (diffs []string) {
	// insertion point for field diffs
	if body.Name != bodyOther.Name {
		diffs = append(diffs, body.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, body, "Paragraphs", bodyOther.Paragraphs, body.Paragraphs); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, body, "Tables", bodyOther.Tables, body.Tables); ops != "" {
		diffs = append(diffs, ops)
	}
	if body.LastParagraph != bodyOther.LastParagraph {
		diffs = append(diffs, body.GongMarshallField(stage, "LastParagraph"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (document *Document) GongDiff(stage *Stage, documentOther *Document) (diffs []string) {
	// insertion point for field diffs
	if document.Name != documentOther.Name {
		diffs = append(diffs, document.GongMarshallField(stage, "Name"))
	}
	if document.File != documentOther.File {
		diffs = append(diffs, document.GongMarshallField(stage, "File"))
	}
	if document.Root != documentOther.Root {
		diffs = append(diffs, document.GongMarshallField(stage, "Root"))
	}
	if document.Body != documentOther.Body {
		diffs = append(diffs, document.GongMarshallField(stage, "Body"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (docx *Docx) GongDiff(stage *Stage, docxOther *Docx) (diffs []string) {
	// insertion point for field diffs
	if docx.Name != docxOther.Name {
		diffs = append(diffs, docx.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, docx, "Files", docxOther.Files, docx.Files); ops != "" {
		diffs = append(diffs, ops)
	}
	if docx.Document != docxOther.Document {
		diffs = append(diffs, docx.GongMarshallField(stage, "Document"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (file *File) GongDiff(stage *Stage, fileOther *File) (diffs []string) {
	// insertion point for field diffs
	if file.Name != fileOther.Name {
		diffs = append(diffs, file.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (node *Node) GongDiff(stage *Stage, nodeOther *Node) (diffs []string) {
	// insertion point for field diffs
	if node.Name != nodeOther.Name {
		diffs = append(diffs, node.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, node, "Nodes", nodeOther.Nodes, node.Nodes); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (paragraph *Paragraph) GongDiff(stage *Stage, paragraphOther *Paragraph) (diffs []string) {
	// insertion point for field diffs
	if paragraph.Name != paragraphOther.Name {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Name"))
	}
	if paragraph.Content != paragraphOther.Content {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Content"))
	}
	if paragraph.Node != paragraphOther.Node {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Node"))
	}
	if paragraph.ParagraphProperties != paragraphOther.ParagraphProperties {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "ParagraphProperties"))
	}
	if ops := __gong__diffSliceOfPointers(stage, paragraph, "Runes", paragraphOther.Runes, paragraph.Runes); ops != "" {
		diffs = append(diffs, ops)
	}
	if paragraph.CollatedText != paragraphOther.CollatedText {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "CollatedText"))
	}
	if paragraph.Next != paragraphOther.Next {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Next"))
	}
	if paragraph.Previous != paragraphOther.Previous {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Previous"))
	}
	if paragraph.EnclosingBody != paragraphOther.EnclosingBody {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingBody"))
	}
	if paragraph.EnclosingTableColumn != paragraphOther.EnclosingTableColumn {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingTableColumn"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (paragraphproperties *ParagraphProperties) GongDiff(stage *Stage, paragraphpropertiesOther *ParagraphProperties) (diffs []string) {
	// insertion point for field diffs
	if paragraphproperties.Name != paragraphpropertiesOther.Name {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "Name"))
	}
	if paragraphproperties.Content != paragraphpropertiesOther.Content {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "Content"))
	}
	if paragraphproperties.ParagraphStyle != paragraphpropertiesOther.ParagraphStyle {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "ParagraphStyle"))
	}
	if paragraphproperties.Node != paragraphpropertiesOther.Node {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "Node"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (paragraphstyle *ParagraphStyle) GongDiff(stage *Stage, paragraphstyleOther *ParagraphStyle) (diffs []string) {
	// insertion point for field diffs
	if paragraphstyle.Name != paragraphstyleOther.Name {
		diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "Name"))
	}
	if paragraphstyle.Node != paragraphstyleOther.Node {
		diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "Node"))
	}
	if paragraphstyle.Content != paragraphstyleOther.Content {
		diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "Content"))
	}
	if paragraphstyle.ValAttr != paragraphstyleOther.ValAttr {
		diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "ValAttr"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rune *Rune) GongDiff(stage *Stage, runeOther *Rune) (diffs []string) {
	// insertion point for field diffs
	if rune.Name != runeOther.Name {
		diffs = append(diffs, rune.GongMarshallField(stage, "Name"))
	}
	if rune.Content != runeOther.Content {
		diffs = append(diffs, rune.GongMarshallField(stage, "Content"))
	}
	if rune.Node != runeOther.Node {
		diffs = append(diffs, rune.GongMarshallField(stage, "Node"))
	}
	if rune.Text != runeOther.Text {
		diffs = append(diffs, rune.GongMarshallField(stage, "Text"))
	}
	if rune.RuneProperties != runeOther.RuneProperties {
		diffs = append(diffs, rune.GongMarshallField(stage, "RuneProperties"))
	}
	if rune.EnclosingParagraph != runeOther.EnclosingParagraph {
		diffs = append(diffs, rune.GongMarshallField(stage, "EnclosingParagraph"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (runeproperties *RuneProperties) GongDiff(stage *Stage, runepropertiesOther *RuneProperties) (diffs []string) {
	// insertion point for field diffs
	if runeproperties.Name != runepropertiesOther.Name {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "Name"))
	}
	if runeproperties.Node != runepropertiesOther.Node {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "Node"))
	}
	if runeproperties.IsBold != runepropertiesOther.IsBold {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "IsBold"))
	}
	if runeproperties.IsStrike != runepropertiesOther.IsStrike {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "IsStrike"))
	}
	if runeproperties.IsItalic != runepropertiesOther.IsItalic {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "IsItalic"))
	}
	if runeproperties.Content != runepropertiesOther.Content {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "Content"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (table *Table) GongDiff(stage *Stage, tableOther *Table) (diffs []string) {
	// insertion point for field diffs
	if table.Name != tableOther.Name {
		diffs = append(diffs, table.GongMarshallField(stage, "Name"))
	}
	if table.Node != tableOther.Node {
		diffs = append(diffs, table.GongMarshallField(stage, "Node"))
	}
	if table.Content != tableOther.Content {
		diffs = append(diffs, table.GongMarshallField(stage, "Content"))
	}
	if table.TableProperties != tableOther.TableProperties {
		diffs = append(diffs, table.GongMarshallField(stage, "TableProperties"))
	}
	if ops := __gong__diffSliceOfPointers(stage, table, "TableRows", tableOther.TableRows, table.TableRows); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tablecolumn *TableColumn) GongDiff(stage *Stage, tablecolumnOther *TableColumn) (diffs []string) {
	// insertion point for field diffs
	if tablecolumn.Name != tablecolumnOther.Name {
		diffs = append(diffs, tablecolumn.GongMarshallField(stage, "Name"))
	}
	if tablecolumn.Content != tablecolumnOther.Content {
		diffs = append(diffs, tablecolumn.GongMarshallField(stage, "Content"))
	}
	if tablecolumn.Node != tablecolumnOther.Node {
		diffs = append(diffs, tablecolumn.GongMarshallField(stage, "Node"))
	}
	if ops := __gong__diffSliceOfPointers(stage, tablecolumn, "Paragraphs", tablecolumnOther.Paragraphs, tablecolumn.Paragraphs); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tableproperties *TableProperties) GongDiff(stage *Stage, tablepropertiesOther *TableProperties) (diffs []string) {
	// insertion point for field diffs
	if tableproperties.Name != tablepropertiesOther.Name {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "Name"))
	}
	if tableproperties.Node != tablepropertiesOther.Node {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "Node"))
	}
	if tableproperties.Content != tablepropertiesOther.Content {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "Content"))
	}
	if tableproperties.TableStyle != tablepropertiesOther.TableStyle {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "TableStyle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tablerow *TableRow) GongDiff(stage *Stage, tablerowOther *TableRow) (diffs []string) {
	// insertion point for field diffs
	if tablerow.Name != tablerowOther.Name {
		diffs = append(diffs, tablerow.GongMarshallField(stage, "Name"))
	}
	if tablerow.Content != tablerowOther.Content {
		diffs = append(diffs, tablerow.GongMarshallField(stage, "Content"))
	}
	if tablerow.Node != tablerowOther.Node {
		diffs = append(diffs, tablerow.GongMarshallField(stage, "Node"))
	}
	if ops := __gong__diffSliceOfPointers(stage, tablerow, "TableColumns", tablerowOther.TableColumns, tablerow.TableColumns); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (tablestyle *TableStyle) GongDiff(stage *Stage, tablestyleOther *TableStyle) (diffs []string) {
	// insertion point for field diffs
	if tablestyle.Name != tablestyleOther.Name {
		diffs = append(diffs, tablestyle.GongMarshallField(stage, "Name"))
	}
	if tablestyle.Node != tablestyleOther.Node {
		diffs = append(diffs, tablestyle.GongMarshallField(stage, "Node"))
	}
	if tablestyle.Content != tablestyleOther.Content {
		diffs = append(diffs, tablestyle.GongMarshallField(stage, "Content"))
	}
	if tablestyle.Val != tablestyleOther.Val {
		diffs = append(diffs, tablestyle.GongMarshallField(stage, "Val"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (text *Text) GongDiff(stage *Stage, textOther *Text) (diffs []string) {
	// insertion point for field diffs
	if text.Name != textOther.Name {
		diffs = append(diffs, text.GongMarshallField(stage, "Name"))
	}
	if text.Content != textOther.Content {
		diffs = append(diffs, text.GongMarshallField(stage, "Content"))
	}
	if text.Node != textOther.Node {
		diffs = append(diffs, text.GongMarshallField(stage, "Node"))
	}
	if text.PreserveWhiteSpace != textOther.PreserveWhiteSpace {
		diffs = append(diffs, text.GongMarshallField(stage, "PreserveWhiteSpace"))
	}
	if text.EnclosingRune != textOther.EnclosingRune {
		diffs = append(diffs, text.GongMarshallField(stage, "EnclosingRune"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
