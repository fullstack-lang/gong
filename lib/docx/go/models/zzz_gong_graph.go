// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (body *Body) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Bodys[body]

	return
}

func (stage *Stage) IsStagedBody(body *Body) (ok bool) {

	return body.GongIsStaged(stage)
}

func (document *Document) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Documents[document]

	return
}

func (stage *Stage) IsStagedDocument(document *Document) (ok bool) {

	return document.GongIsStaged(stage)
}

func (docx *Docx) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Docxs[docx]

	return
}

func (stage *Stage) IsStagedDocx(docx *Docx) (ok bool) {

	return docx.GongIsStaged(stage)
}

func (file *File) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Files[file]

	return
}

func (stage *Stage) IsStagedFile(file *File) (ok bool) {

	return file.GongIsStaged(stage)
}

func (node *Node) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Nodes[node]

	return
}

func (stage *Stage) IsStagedNode(node *Node) (ok bool) {

	return node.GongIsStaged(stage)
}

func (paragraph *Paragraph) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Paragraphs[paragraph]

	return
}

func (stage *Stage) IsStagedParagraph(paragraph *Paragraph) (ok bool) {

	return paragraph.GongIsStaged(stage)
}

func (paragraphproperties *ParagraphProperties) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ParagraphPropertiess[paragraphproperties]

	return
}

func (stage *Stage) IsStagedParagraphProperties(paragraphproperties *ParagraphProperties) (ok bool) {

	return paragraphproperties.GongIsStaged(stage)
}

func (paragraphstyle *ParagraphStyle) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ParagraphStyles[paragraphstyle]

	return
}

func (stage *Stage) IsStagedParagraphStyle(paragraphstyle *ParagraphStyle) (ok bool) {

	return paragraphstyle.GongIsStaged(stage)
}

func (rune *Rune) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Runes[rune]

	return
}

func (stage *Stage) IsStagedRune(rune *Rune) (ok bool) {

	return rune.GongIsStaged(stage)
}

func (runeproperties *RuneProperties) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RunePropertiess[runeproperties]

	return
}

func (stage *Stage) IsStagedRuneProperties(runeproperties *RuneProperties) (ok bool) {

	return runeproperties.GongIsStaged(stage)
}

func (table *Table) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	return table.GongIsStaged(stage)
}

func (tablecolumn *TableColumn) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TableColumns[tablecolumn]

	return
}

func (stage *Stage) IsStagedTableColumn(tablecolumn *TableColumn) (ok bool) {

	return tablecolumn.GongIsStaged(stage)
}

func (tableproperties *TableProperties) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TablePropertiess[tableproperties]

	return
}

func (stage *Stage) IsStagedTableProperties(tableproperties *TableProperties) (ok bool) {

	return tableproperties.GongIsStaged(stage)
}

func (tablerow *TableRow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TableRows[tablerow]

	return
}

func (stage *Stage) IsStagedTableRow(tablerow *TableRow) (ok bool) {

	return tablerow.GongIsStaged(stage)
}

func (tablestyle *TableStyle) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TableStyles[tablestyle]

	return
}

func (stage *Stage) IsStagedTableStyle(tablestyle *TableStyle) (ok bool) {

	return tablestyle.GongIsStaged(stage)
}

func (text *Text) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

func (stage *Stage) IsStagedText(text *Text) (ok bool) {

	return text.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (body *Body) GongStageBranch(stage *Stage) {
	stage.StageBranchBody(body)
}

func (stage *Stage) StageBranchBody(body *Body) {

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
	stage.StageBranchDocument(document)
}

func (stage *Stage) StageBranchDocument(document *Document) {

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
	stage.StageBranchDocx(docx)
}

func (stage *Stage) StageBranchDocx(docx *Docx) {

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
	stage.StageBranchFile(file)
}

func (stage *Stage) StageBranchFile(file *File) {

	// check if instance is already staged
	if stage.IsStaged(file) {
		return
	}

	file.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (node *Node) GongStageBranch(stage *Stage) {
	stage.StageBranchNode(node)
}

func (stage *Stage) StageBranchNode(node *Node) {

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
	stage.StageBranchParagraph(paragraph)
}

func (stage *Stage) StageBranchParagraph(paragraph *Paragraph) {

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
	stage.StageBranchParagraphProperties(paragraphproperties)
}

func (stage *Stage) StageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

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
	stage.StageBranchParagraphStyle(paragraphstyle)
}

func (stage *Stage) StageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

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
	stage.StageBranchRune(rune)
}

func (stage *Stage) StageBranchRune(rune *Rune) {

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
	stage.StageBranchRuneProperties(runeproperties)
}

func (stage *Stage) StageBranchRuneProperties(runeproperties *RuneProperties) {

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
	stage.StageBranchTable(table)
}

func (stage *Stage) StageBranchTable(table *Table) {

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
	stage.StageBranchTableColumn(tablecolumn)
}

func (stage *Stage) StageBranchTableColumn(tablecolumn *TableColumn) {

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
	stage.StageBranchTableProperties(tableproperties)
}

func (stage *Stage) StageBranchTableProperties(tableproperties *TableProperties) {

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
	stage.StageBranchTableRow(tablerow)
}

func (stage *Stage) StageBranchTableRow(tablerow *TableRow) {

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
	stage.StageBranchTableStyle(tablestyle)
}

func (stage *Stage) StageBranchTableStyle(tablestyle *TableStyle) {

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
	stage.StageBranchText(text)
}

func (stage *Stage) StageBranchText(text *Text) {

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

	// bodyFrom has already been copied
	if _bodyTo, ok := mapOrigCopy[bodyFrom]; ok {
		bodyTo = _bodyTo.(*Body)
		return
	}

	bodyTo = new(Body)
	mapOrigCopy[bodyFrom] = bodyTo
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

	// documentFrom has already been copied
	if _documentTo, ok := mapOrigCopy[documentFrom]; ok {
		documentTo = _documentTo.(*Document)
		return
	}

	documentTo = new(Document)
	mapOrigCopy[documentFrom] = documentTo
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

	// docxFrom has already been copied
	if _docxTo, ok := mapOrigCopy[docxFrom]; ok {
		docxTo = _docxTo.(*Docx)
		return
	}

	docxTo = new(Docx)
	mapOrigCopy[docxFrom] = docxTo
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

	// fileFrom has already been copied
	if _fileTo, ok := mapOrigCopy[fileFrom]; ok {
		fileTo = _fileTo.(*File)
		return
	}

	fileTo = new(File)
	mapOrigCopy[fileFrom] = fileTo
	fileFrom.GongCopyBasicFields(fileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {

	// nodeFrom has already been copied
	if _nodeTo, ok := mapOrigCopy[nodeFrom]; ok {
		nodeTo = _nodeTo.(*Node)
		return
	}

	nodeTo = new(Node)
	mapOrigCopy[nodeFrom] = nodeTo
	nodeFrom.GongCopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Nodes {
		nodeTo.Nodes = append(nodeTo.Nodes, GongCopyBranchNode(mapOrigCopy, _node))
	}

	return
}

func GongCopyBranchParagraph(mapOrigCopy map[any]any, paragraphFrom *Paragraph) (paragraphTo *Paragraph) {

	// paragraphFrom has already been copied
	if _paragraphTo, ok := mapOrigCopy[paragraphFrom]; ok {
		paragraphTo = _paragraphTo.(*Paragraph)
		return
	}

	paragraphTo = new(Paragraph)
	mapOrigCopy[paragraphFrom] = paragraphTo
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

	// paragraphpropertiesFrom has already been copied
	if _paragraphpropertiesTo, ok := mapOrigCopy[paragraphpropertiesFrom]; ok {
		paragraphpropertiesTo = _paragraphpropertiesTo.(*ParagraphProperties)
		return
	}

	paragraphpropertiesTo = new(ParagraphProperties)
	mapOrigCopy[paragraphpropertiesFrom] = paragraphpropertiesTo
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

	// paragraphstyleFrom has already been copied
	if _paragraphstyleTo, ok := mapOrigCopy[paragraphstyleFrom]; ok {
		paragraphstyleTo = _paragraphstyleTo.(*ParagraphStyle)
		return
	}

	paragraphstyleTo = new(ParagraphStyle)
	mapOrigCopy[paragraphstyleFrom] = paragraphstyleTo
	paragraphstyleFrom.GongCopyBasicFields(paragraphstyleTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyleFrom.Node != nil {
		paragraphstyleTo.Node = GongCopyBranchNode(mapOrigCopy, paragraphstyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRune(mapOrigCopy map[any]any, runeFrom *Rune) (runeTo *Rune) {

	// runeFrom has already been copied
	if _runeTo, ok := mapOrigCopy[runeFrom]; ok {
		runeTo = _runeTo.(*Rune)
		return
	}

	runeTo = new(Rune)
	mapOrigCopy[runeFrom] = runeTo
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

	// runepropertiesFrom has already been copied
	if _runepropertiesTo, ok := mapOrigCopy[runepropertiesFrom]; ok {
		runepropertiesTo = _runepropertiesTo.(*RuneProperties)
		return
	}

	runepropertiesTo = new(RuneProperties)
	mapOrigCopy[runepropertiesFrom] = runepropertiesTo
	runepropertiesFrom.GongCopyBasicFields(runepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if runepropertiesFrom.Node != nil {
		runepropertiesTo.Node = GongCopyBranchNode(mapOrigCopy, runepropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
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

	// tablecolumnFrom has already been copied
	if _tablecolumnTo, ok := mapOrigCopy[tablecolumnFrom]; ok {
		tablecolumnTo = _tablecolumnTo.(*TableColumn)
		return
	}

	tablecolumnTo = new(TableColumn)
	mapOrigCopy[tablecolumnFrom] = tablecolumnTo
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

	// tablepropertiesFrom has already been copied
	if _tablepropertiesTo, ok := mapOrigCopy[tablepropertiesFrom]; ok {
		tablepropertiesTo = _tablepropertiesTo.(*TableProperties)
		return
	}

	tablepropertiesTo = new(TableProperties)
	mapOrigCopy[tablepropertiesFrom] = tablepropertiesTo
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

	// tablerowFrom has already been copied
	if _tablerowTo, ok := mapOrigCopy[tablerowFrom]; ok {
		tablerowTo = _tablerowTo.(*TableRow)
		return
	}

	tablerowTo = new(TableRow)
	mapOrigCopy[tablerowFrom] = tablerowTo
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

	// tablestyleFrom has already been copied
	if _tablestyleTo, ok := mapOrigCopy[tablestyleFrom]; ok {
		tablestyleTo = _tablestyleTo.(*TableStyle)
		return
	}

	tablestyleTo = new(TableStyle)
	mapOrigCopy[tablestyleFrom] = tablestyleTo
	tablestyleFrom.GongCopyBasicFields(tablestyleTo)

	//insertion point for the staging of instances referenced by pointers
	if tablestyleFrom.Node != nil {
		tablestyleTo.Node = GongCopyBranchNode(mapOrigCopy, tablestyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {

	// textFrom has already been copied
	if _textTo, ok := mapOrigCopy[textFrom]; ok {
		textTo = _textTo.(*Text)
		return
	}

	textTo = new(Text)
	mapOrigCopy[textFrom] = textTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (body *Body) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchBody(body)
}

func (stage *Stage) UnstageBranchBody(body *Body) {

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
	stage.UnstageBranchDocument(document)
}

func (stage *Stage) UnstageBranchDocument(document *Document) {

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
	stage.UnstageBranchDocx(docx)
}

func (stage *Stage) UnstageBranchDocx(docx *Docx) {

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
	stage.UnstageBranchFile(file)
}

func (stage *Stage) UnstageBranchFile(file *File) {

	// check if instance is already staged
	if !stage.IsStaged(file) {
		return
	}

	file.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (node *Node) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNode(node)
}

func (stage *Stage) UnstageBranchNode(node *Node) {

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
	stage.UnstageBranchParagraph(paragraph)
}

func (stage *Stage) UnstageBranchParagraph(paragraph *Paragraph) {

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
	stage.UnstageBranchParagraphProperties(paragraphproperties)
}

func (stage *Stage) UnstageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

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
	stage.UnstageBranchParagraphStyle(paragraphstyle)
}

func (stage *Stage) UnstageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

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
	stage.UnstageBranchRune(rune)
}

func (stage *Stage) UnstageBranchRune(rune *Rune) {

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
	stage.UnstageBranchRuneProperties(runeproperties)
}

func (stage *Stage) UnstageBranchRuneProperties(runeproperties *RuneProperties) {

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
	stage.UnstageBranchTable(table)
}

func (stage *Stage) UnstageBranchTable(table *Table) {

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
	stage.UnstageBranchTableColumn(tablecolumn)
}

func (stage *Stage) UnstageBranchTableColumn(tablecolumn *TableColumn) {

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
	stage.UnstageBranchTableProperties(tableproperties)
}

func (stage *Stage) UnstageBranchTableProperties(tableproperties *TableProperties) {

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
	stage.UnstageBranchTableRow(tablerow)
}

func (stage *Stage) UnstageBranchTableRow(tablerow *TableRow) {

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
	stage.UnstageBranchTableStyle(tablestyle)
}

func (stage *Stage) UnstageBranchTableStyle(tablestyle *TableStyle) {

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
	stage.UnstageBranchText(text)
}

func (stage *Stage) UnstageBranchText(text *Text) {

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
	if instance.LastParagraph != nil {
		reference.LastParagraph = stage.Paragraphs_reference[instance.LastParagraph]
	}
	// insertion point for slice of pointers field
	reference.Paragraphs = reference.Paragraphs[:0]
	for _, _b := range instance.Paragraphs {
		reference.Paragraphs = append(reference.Paragraphs, stage.Paragraphs_reference[_b])
	}
	reference.Tables = reference.Tables[:0]
	for _, _b := range instance.Tables {
		reference.Tables = append(reference.Tables, stage.Tables_reference[_b])
	}
}

func (reference *Document) GongReconstructPointersFromReferences(stage *Stage, instance *Document) {
	// insertion point for pointers field
	if instance.File != nil {
		reference.File = stage.Files_reference[instance.File]
	}
	if instance.Root != nil {
		reference.Root = stage.Nodes_reference[instance.Root]
	}
	if instance.Body != nil {
		reference.Body = stage.Bodys_reference[instance.Body]
	}
	// insertion point for slice of pointers field
}

func (reference *Docx) GongReconstructPointersFromReferences(stage *Stage, instance *Docx) {
	// insertion point for pointers field
	if instance.Document != nil {
		reference.Document = stage.Documents_reference[instance.Document]
	}
	// insertion point for slice of pointers field
	reference.Files = reference.Files[:0]
	for _, _b := range instance.Files {
		reference.Files = append(reference.Files, stage.Files_reference[_b])
	}
}

func (reference *File) GongReconstructPointersFromReferences(stage *Stage, instance *File) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Node) GongReconstructPointersFromReferences(stage *Stage, instance *Node) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Nodes = reference.Nodes[:0]
	for _, _b := range instance.Nodes {
		reference.Nodes = append(reference.Nodes, stage.Nodes_reference[_b])
	}
}

func (reference *Paragraph) GongReconstructPointersFromReferences(stage *Stage, instance *Paragraph) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.ParagraphProperties != nil {
		reference.ParagraphProperties = stage.ParagraphPropertiess_reference[instance.ParagraphProperties]
	}
	if instance.Next != nil {
		reference.Next = stage.Paragraphs_reference[instance.Next]
	}
	if instance.Previous != nil {
		reference.Previous = stage.Paragraphs_reference[instance.Previous]
	}
	if instance.EnclosingBody != nil {
		reference.EnclosingBody = stage.Bodys_reference[instance.EnclosingBody]
	}
	if instance.EnclosingTableColumn != nil {
		reference.EnclosingTableColumn = stage.TableColumns_reference[instance.EnclosingTableColumn]
	}
	// insertion point for slice of pointers field
	reference.Runes = reference.Runes[:0]
	for _, _b := range instance.Runes {
		reference.Runes = append(reference.Runes, stage.Runes_reference[_b])
	}
}

func (reference *ParagraphProperties) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphProperties) {
	// insertion point for pointers field
	if instance.ParagraphStyle != nil {
		reference.ParagraphStyle = stage.ParagraphStyles_reference[instance.ParagraphStyle]
	}
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
}

func (reference *ParagraphStyle) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphStyle) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
}

func (reference *Rune) GongReconstructPointersFromReferences(stage *Stage, instance *Rune) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.Text != nil {
		reference.Text = stage.Texts_reference[instance.Text]
	}
	if instance.RuneProperties != nil {
		reference.RuneProperties = stage.RunePropertiess_reference[instance.RuneProperties]
	}
	if instance.EnclosingParagraph != nil {
		reference.EnclosingParagraph = stage.Paragraphs_reference[instance.EnclosingParagraph]
	}
	// insertion point for slice of pointers field
}

func (reference *RuneProperties) GongReconstructPointersFromReferences(stage *Stage, instance *RuneProperties) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.TableProperties != nil {
		reference.TableProperties = stage.TablePropertiess_reference[instance.TableProperties]
	}
	// insertion point for slice of pointers field
	reference.TableRows = reference.TableRows[:0]
	for _, _b := range instance.TableRows {
		reference.TableRows = append(reference.TableRows, stage.TableRows_reference[_b])
	}
}

func (reference *TableColumn) GongReconstructPointersFromReferences(stage *Stage, instance *TableColumn) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
	reference.Paragraphs = reference.Paragraphs[:0]
	for _, _b := range instance.Paragraphs {
		reference.Paragraphs = append(reference.Paragraphs, stage.Paragraphs_reference[_b])
	}
}

func (reference *TableProperties) GongReconstructPointersFromReferences(stage *Stage, instance *TableProperties) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.TableStyle != nil {
		reference.TableStyle = stage.TableStyles_reference[instance.TableStyle]
	}
	// insertion point for slice of pointers field
}

func (reference *TableRow) GongReconstructPointersFromReferences(stage *Stage, instance *TableRow) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
	reference.TableColumns = reference.TableColumns[:0]
	for _, _b := range instance.TableColumns {
		reference.TableColumns = append(reference.TableColumns, stage.TableColumns_reference[_b])
	}
}

func (reference *TableStyle) GongReconstructPointersFromReferences(stage *Stage, instance *TableStyle) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
}

func (reference *Text) GongReconstructPointersFromReferences(stage *Stage, instance *Text) {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.EnclosingRune != nil {
		reference.EnclosingRune = stage.Runes_reference[instance.EnclosingRune]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Body) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.LastParagraph; _reference != nil {
		reference.LastParagraph = nil
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			reference.LastParagraph = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Paragraphs []*Paragraph
	for _, _reference := range reference.Paragraphs {
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			_Paragraphs = append(_Paragraphs, _instance)
		}
	}
	reference.Paragraphs = _Paragraphs
	var _Tables []*Table
	for _, _reference := range reference.Tables {
		if _instance, ok := stage.Tables_instance[_reference]; ok {
			_Tables = append(_Tables, _instance)
		}
	}
	reference.Tables = _Tables
}

func (reference *Document) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.File; _reference != nil {
		reference.File = nil
		if _instance, ok := stage.Files_instance[_reference]; ok {
			reference.File = _instance
		}
	}
	if _reference := reference.Root; _reference != nil {
		reference.Root = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Root = _instance
		}
	}
	if _reference := reference.Body; _reference != nil {
		reference.Body = nil
		if _instance, ok := stage.Bodys_instance[_reference]; ok {
			reference.Body = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Docx) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Document; _reference != nil {
		reference.Document = nil
		if _instance, ok := stage.Documents_instance[_reference]; ok {
			reference.Document = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Files []*File
	for _, _reference := range reference.Files {
		if _instance, ok := stage.Files_instance[_reference]; ok {
			_Files = append(_Files, _instance)
		}
	}
	reference.Files = _Files
}

func (reference *File) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Node) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Nodes []*Node
	for _, _reference := range reference.Nodes {
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			_Nodes = append(_Nodes, _instance)
		}
	}
	reference.Nodes = _Nodes
}

func (reference *Paragraph) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	if _reference := reference.ParagraphProperties; _reference != nil {
		reference.ParagraphProperties = nil
		if _instance, ok := stage.ParagraphPropertiess_instance[_reference]; ok {
			reference.ParagraphProperties = _instance
		}
	}
	if _reference := reference.Next; _reference != nil {
		reference.Next = nil
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			reference.Next = _instance
		}
	}
	if _reference := reference.Previous; _reference != nil {
		reference.Previous = nil
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			reference.Previous = _instance
		}
	}
	if _reference := reference.EnclosingBody; _reference != nil {
		reference.EnclosingBody = nil
		if _instance, ok := stage.Bodys_instance[_reference]; ok {
			reference.EnclosingBody = _instance
		}
	}
	if _reference := reference.EnclosingTableColumn; _reference != nil {
		reference.EnclosingTableColumn = nil
		if _instance, ok := stage.TableColumns_instance[_reference]; ok {
			reference.EnclosingTableColumn = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Runes []*Rune
	for _, _reference := range reference.Runes {
		if _instance, ok := stage.Runes_instance[_reference]; ok {
			_Runes = append(_Runes, _instance)
		}
	}
	reference.Runes = _Runes
}

func (reference *ParagraphProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ParagraphStyle; _reference != nil {
		reference.ParagraphStyle = nil
		if _instance, ok := stage.ParagraphStyles_instance[_reference]; ok {
			reference.ParagraphStyle = _instance
		}
	}
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ParagraphStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Rune) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	if _reference := reference.Text; _reference != nil {
		reference.Text = nil
		if _instance, ok := stage.Texts_instance[_reference]; ok {
			reference.Text = _instance
		}
	}
	if _reference := reference.RuneProperties; _reference != nil {
		reference.RuneProperties = nil
		if _instance, ok := stage.RunePropertiess_instance[_reference]; ok {
			reference.RuneProperties = _instance
		}
	}
	if _reference := reference.EnclosingParagraph; _reference != nil {
		reference.EnclosingParagraph = nil
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			reference.EnclosingParagraph = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *RuneProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	if _reference := reference.TableProperties; _reference != nil {
		reference.TableProperties = nil
		if _instance, ok := stage.TablePropertiess_instance[_reference]; ok {
			reference.TableProperties = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _TableRows []*TableRow
	for _, _reference := range reference.TableRows {
		if _instance, ok := stage.TableRows_instance[_reference]; ok {
			_TableRows = append(_TableRows, _instance)
		}
	}
	reference.TableRows = _TableRows
}

func (reference *TableColumn) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Paragraphs []*Paragraph
	for _, _reference := range reference.Paragraphs {
		if _instance, ok := stage.Paragraphs_instance[_reference]; ok {
			_Paragraphs = append(_Paragraphs, _instance)
		}
	}
	reference.Paragraphs = _Paragraphs
}

func (reference *TableProperties) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	if _reference := reference.TableStyle; _reference != nil {
		reference.TableStyle = nil
		if _instance, ok := stage.TableStyles_instance[_reference]; ok {
			reference.TableStyle = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TableRow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _TableColumns []*TableColumn
	for _, _reference := range reference.TableColumns {
		if _instance, ok := stage.TableColumns_instance[_reference]; ok {
			_TableColumns = append(_TableColumns, _instance)
		}
	}
	reference.TableColumns = _TableColumns
}

func (reference *TableStyle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Text) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	if _reference := reference.EnclosingRune; _reference != nil {
		reference.EnclosingRune = nil
		if _instance, ok := stage.Runes_instance[_reference]; ok {
			reference.EnclosingRune = _instance
		}
	}
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
	ParagraphsDifferent := false
	if len(body.Paragraphs) != len(bodyOther.Paragraphs) {
		ParagraphsDifferent = true
	} else {
		for i := range body.Paragraphs {
			if (body.Paragraphs[i] == nil) != (bodyOther.Paragraphs[i] == nil) {
				ParagraphsDifferent = true
				break
			} else if body.Paragraphs[i] != nil && bodyOther.Paragraphs[i] != nil {
				// this is a pointer comparaison
				if body.Paragraphs[i] != bodyOther.Paragraphs[i] {
					ParagraphsDifferent = true
					break
				}
			}
		}
	}
	if ParagraphsDifferent {
		ops := stage.Diff(
			body,
			"Paragraphs",
			len(bodyOther.Paragraphs),
			len(body.Paragraphs),
			func(i, j int) bool {
				return bodyOther.Paragraphs[i] == body.Paragraphs[j]
			},
			func(j int) string {
				return body.Paragraphs[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TablesDifferent := false
	if len(body.Tables) != len(bodyOther.Tables) {
		TablesDifferent = true
	} else {
		for i := range body.Tables {
			if (body.Tables[i] == nil) != (bodyOther.Tables[i] == nil) {
				TablesDifferent = true
				break
			} else if body.Tables[i] != nil && bodyOther.Tables[i] != nil {
				// this is a pointer comparaison
				if body.Tables[i] != bodyOther.Tables[i] {
					TablesDifferent = true
					break
				}
			}
		}
	}
	if TablesDifferent {
		ops := stage.Diff(
			body,
			"Tables",
			len(bodyOther.Tables),
			len(body.Tables),
			func(i, j int) bool {
				return bodyOther.Tables[i] == body.Tables[j]
			},
			func(j int) string {
				return body.Tables[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (body.LastParagraph == nil) != (bodyOther.LastParagraph == nil) {
		diffs = append(diffs, body.GongMarshallField(stage, "LastParagraph"))
	} else if body.LastParagraph != nil && bodyOther.LastParagraph != nil {
		if body.LastParagraph != bodyOther.LastParagraph {
			diffs = append(diffs, body.GongMarshallField(stage, "LastParagraph"))
		}
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
	if (document.File == nil) != (documentOther.File == nil) {
		diffs = append(diffs, document.GongMarshallField(stage, "File"))
	} else if document.File != nil && documentOther.File != nil {
		if document.File != documentOther.File {
			diffs = append(diffs, document.GongMarshallField(stage, "File"))
		}
	}
	if (document.Root == nil) != (documentOther.Root == nil) {
		diffs = append(diffs, document.GongMarshallField(stage, "Root"))
	} else if document.Root != nil && documentOther.Root != nil {
		if document.Root != documentOther.Root {
			diffs = append(diffs, document.GongMarshallField(stage, "Root"))
		}
	}
	if (document.Body == nil) != (documentOther.Body == nil) {
		diffs = append(diffs, document.GongMarshallField(stage, "Body"))
	} else if document.Body != nil && documentOther.Body != nil {
		if document.Body != documentOther.Body {
			diffs = append(diffs, document.GongMarshallField(stage, "Body"))
		}
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
	FilesDifferent := false
	if len(docx.Files) != len(docxOther.Files) {
		FilesDifferent = true
	} else {
		for i := range docx.Files {
			if (docx.Files[i] == nil) != (docxOther.Files[i] == nil) {
				FilesDifferent = true
				break
			} else if docx.Files[i] != nil && docxOther.Files[i] != nil {
				// this is a pointer comparaison
				if docx.Files[i] != docxOther.Files[i] {
					FilesDifferent = true
					break
				}
			}
		}
	}
	if FilesDifferent {
		ops := stage.Diff(
			docx,
			"Files",
			len(docxOther.Files),
			len(docx.Files),
			func(i, j int) bool {
				return docxOther.Files[i] == docx.Files[j]
			},
			func(j int) string {
				return docx.Files[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (docx.Document == nil) != (docxOther.Document == nil) {
		diffs = append(diffs, docx.GongMarshallField(stage, "Document"))
	} else if docx.Document != nil && docxOther.Document != nil {
		if docx.Document != docxOther.Document {
			diffs = append(diffs, docx.GongMarshallField(stage, "Document"))
		}
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
	NodesDifferent := false
	if len(node.Nodes) != len(nodeOther.Nodes) {
		NodesDifferent = true
	} else {
		for i := range node.Nodes {
			if (node.Nodes[i] == nil) != (nodeOther.Nodes[i] == nil) {
				NodesDifferent = true
				break
			} else if node.Nodes[i] != nil && nodeOther.Nodes[i] != nil {
				// this is a pointer comparaison
				if node.Nodes[i] != nodeOther.Nodes[i] {
					NodesDifferent = true
					break
				}
			}
		}
	}
	if NodesDifferent {
		ops := stage.Diff(
			node,
			"Nodes",
			len(nodeOther.Nodes),
			len(node.Nodes),
			func(i, j int) bool {
				return nodeOther.Nodes[i] == node.Nodes[j]
			},
			func(j int) string {
				return node.Nodes[j].GongGetIdentifier(stage)
			},
		)
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
	if (paragraph.Node == nil) != (paragraphOther.Node == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Node"))
	} else if paragraph.Node != nil && paragraphOther.Node != nil {
		if paragraph.Node != paragraphOther.Node {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "Node"))
		}
	}
	if (paragraph.ParagraphProperties == nil) != (paragraphOther.ParagraphProperties == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "ParagraphProperties"))
	} else if paragraph.ParagraphProperties != nil && paragraphOther.ParagraphProperties != nil {
		if paragraph.ParagraphProperties != paragraphOther.ParagraphProperties {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "ParagraphProperties"))
		}
	}
	RunesDifferent := false
	if len(paragraph.Runes) != len(paragraphOther.Runes) {
		RunesDifferent = true
	} else {
		for i := range paragraph.Runes {
			if (paragraph.Runes[i] == nil) != (paragraphOther.Runes[i] == nil) {
				RunesDifferent = true
				break
			} else if paragraph.Runes[i] != nil && paragraphOther.Runes[i] != nil {
				// this is a pointer comparaison
				if paragraph.Runes[i] != paragraphOther.Runes[i] {
					RunesDifferent = true
					break
				}
			}
		}
	}
	if RunesDifferent {
		ops := stage.Diff(
			paragraph,
			"Runes",
			len(paragraphOther.Runes),
			len(paragraph.Runes),
			func(i, j int) bool {
				return paragraphOther.Runes[i] == paragraph.Runes[j]
			},
			func(j int) string {
				return paragraph.Runes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if paragraph.CollatedText != paragraphOther.CollatedText {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "CollatedText"))
	}
	if (paragraph.Next == nil) != (paragraphOther.Next == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Next"))
	} else if paragraph.Next != nil && paragraphOther.Next != nil {
		if paragraph.Next != paragraphOther.Next {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "Next"))
		}
	}
	if (paragraph.Previous == nil) != (paragraphOther.Previous == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "Previous"))
	} else if paragraph.Previous != nil && paragraphOther.Previous != nil {
		if paragraph.Previous != paragraphOther.Previous {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "Previous"))
		}
	}
	if (paragraph.EnclosingBody == nil) != (paragraphOther.EnclosingBody == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingBody"))
	} else if paragraph.EnclosingBody != nil && paragraphOther.EnclosingBody != nil {
		if paragraph.EnclosingBody != paragraphOther.EnclosingBody {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingBody"))
		}
	}
	if (paragraph.EnclosingTableColumn == nil) != (paragraphOther.EnclosingTableColumn == nil) {
		diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingTableColumn"))
	} else if paragraph.EnclosingTableColumn != nil && paragraphOther.EnclosingTableColumn != nil {
		if paragraph.EnclosingTableColumn != paragraphOther.EnclosingTableColumn {
			diffs = append(diffs, paragraph.GongMarshallField(stage, "EnclosingTableColumn"))
		}
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
	if (paragraphproperties.ParagraphStyle == nil) != (paragraphpropertiesOther.ParagraphStyle == nil) {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "ParagraphStyle"))
	} else if paragraphproperties.ParagraphStyle != nil && paragraphpropertiesOther.ParagraphStyle != nil {
		if paragraphproperties.ParagraphStyle != paragraphpropertiesOther.ParagraphStyle {
			diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "ParagraphStyle"))
		}
	}
	if (paragraphproperties.Node == nil) != (paragraphpropertiesOther.Node == nil) {
		diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "Node"))
	} else if paragraphproperties.Node != nil && paragraphpropertiesOther.Node != nil {
		if paragraphproperties.Node != paragraphpropertiesOther.Node {
			diffs = append(diffs, paragraphproperties.GongMarshallField(stage, "Node"))
		}
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
	if (paragraphstyle.Node == nil) != (paragraphstyleOther.Node == nil) {
		diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "Node"))
	} else if paragraphstyle.Node != nil && paragraphstyleOther.Node != nil {
		if paragraphstyle.Node != paragraphstyleOther.Node {
			diffs = append(diffs, paragraphstyle.GongMarshallField(stage, "Node"))
		}
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
	if (rune.Node == nil) != (runeOther.Node == nil) {
		diffs = append(diffs, rune.GongMarshallField(stage, "Node"))
	} else if rune.Node != nil && runeOther.Node != nil {
		if rune.Node != runeOther.Node {
			diffs = append(diffs, rune.GongMarshallField(stage, "Node"))
		}
	}
	if (rune.Text == nil) != (runeOther.Text == nil) {
		diffs = append(diffs, rune.GongMarshallField(stage, "Text"))
	} else if rune.Text != nil && runeOther.Text != nil {
		if rune.Text != runeOther.Text {
			diffs = append(diffs, rune.GongMarshallField(stage, "Text"))
		}
	}
	if (rune.RuneProperties == nil) != (runeOther.RuneProperties == nil) {
		diffs = append(diffs, rune.GongMarshallField(stage, "RuneProperties"))
	} else if rune.RuneProperties != nil && runeOther.RuneProperties != nil {
		if rune.RuneProperties != runeOther.RuneProperties {
			diffs = append(diffs, rune.GongMarshallField(stage, "RuneProperties"))
		}
	}
	if (rune.EnclosingParagraph == nil) != (runeOther.EnclosingParagraph == nil) {
		diffs = append(diffs, rune.GongMarshallField(stage, "EnclosingParagraph"))
	} else if rune.EnclosingParagraph != nil && runeOther.EnclosingParagraph != nil {
		if rune.EnclosingParagraph != runeOther.EnclosingParagraph {
			diffs = append(diffs, rune.GongMarshallField(stage, "EnclosingParagraph"))
		}
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
	if (runeproperties.Node == nil) != (runepropertiesOther.Node == nil) {
		diffs = append(diffs, runeproperties.GongMarshallField(stage, "Node"))
	} else if runeproperties.Node != nil && runepropertiesOther.Node != nil {
		if runeproperties.Node != runepropertiesOther.Node {
			diffs = append(diffs, runeproperties.GongMarshallField(stage, "Node"))
		}
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
	if (table.Node == nil) != (tableOther.Node == nil) {
		diffs = append(diffs, table.GongMarshallField(stage, "Node"))
	} else if table.Node != nil && tableOther.Node != nil {
		if table.Node != tableOther.Node {
			diffs = append(diffs, table.GongMarshallField(stage, "Node"))
		}
	}
	if table.Content != tableOther.Content {
		diffs = append(diffs, table.GongMarshallField(stage, "Content"))
	}
	if (table.TableProperties == nil) != (tableOther.TableProperties == nil) {
		diffs = append(diffs, table.GongMarshallField(stage, "TableProperties"))
	} else if table.TableProperties != nil && tableOther.TableProperties != nil {
		if table.TableProperties != tableOther.TableProperties {
			diffs = append(diffs, table.GongMarshallField(stage, "TableProperties"))
		}
	}
	TableRowsDifferent := false
	if len(table.TableRows) != len(tableOther.TableRows) {
		TableRowsDifferent = true
	} else {
		for i := range table.TableRows {
			if (table.TableRows[i] == nil) != (tableOther.TableRows[i] == nil) {
				TableRowsDifferent = true
				break
			} else if table.TableRows[i] != nil && tableOther.TableRows[i] != nil {
				// this is a pointer comparaison
				if table.TableRows[i] != tableOther.TableRows[i] {
					TableRowsDifferent = true
					break
				}
			}
		}
	}
	if TableRowsDifferent {
		ops := stage.Diff(
			table,
			"TableRows",
			len(tableOther.TableRows),
			len(table.TableRows),
			func(i, j int) bool {
				return tableOther.TableRows[i] == table.TableRows[j]
			},
			func(j int) string {
				return table.TableRows[j].GongGetIdentifier(stage)
			},
		)
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
	if (tablecolumn.Node == nil) != (tablecolumnOther.Node == nil) {
		diffs = append(diffs, tablecolumn.GongMarshallField(stage, "Node"))
	} else if tablecolumn.Node != nil && tablecolumnOther.Node != nil {
		if tablecolumn.Node != tablecolumnOther.Node {
			diffs = append(diffs, tablecolumn.GongMarshallField(stage, "Node"))
		}
	}
	ParagraphsDifferent := false
	if len(tablecolumn.Paragraphs) != len(tablecolumnOther.Paragraphs) {
		ParagraphsDifferent = true
	} else {
		for i := range tablecolumn.Paragraphs {
			if (tablecolumn.Paragraphs[i] == nil) != (tablecolumnOther.Paragraphs[i] == nil) {
				ParagraphsDifferent = true
				break
			} else if tablecolumn.Paragraphs[i] != nil && tablecolumnOther.Paragraphs[i] != nil {
				// this is a pointer comparaison
				if tablecolumn.Paragraphs[i] != tablecolumnOther.Paragraphs[i] {
					ParagraphsDifferent = true
					break
				}
			}
		}
	}
	if ParagraphsDifferent {
		ops := stage.Diff(
			tablecolumn,
			"Paragraphs",
			len(tablecolumnOther.Paragraphs),
			len(tablecolumn.Paragraphs),
			func(i, j int) bool {
				return tablecolumnOther.Paragraphs[i] == tablecolumn.Paragraphs[j]
			},
			func(j int) string {
				return tablecolumn.Paragraphs[j].GongGetIdentifier(stage)
			},
		)
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
	if (tableproperties.Node == nil) != (tablepropertiesOther.Node == nil) {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "Node"))
	} else if tableproperties.Node != nil && tablepropertiesOther.Node != nil {
		if tableproperties.Node != tablepropertiesOther.Node {
			diffs = append(diffs, tableproperties.GongMarshallField(stage, "Node"))
		}
	}
	if tableproperties.Content != tablepropertiesOther.Content {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "Content"))
	}
	if (tableproperties.TableStyle == nil) != (tablepropertiesOther.TableStyle == nil) {
		diffs = append(diffs, tableproperties.GongMarshallField(stage, "TableStyle"))
	} else if tableproperties.TableStyle != nil && tablepropertiesOther.TableStyle != nil {
		if tableproperties.TableStyle != tablepropertiesOther.TableStyle {
			diffs = append(diffs, tableproperties.GongMarshallField(stage, "TableStyle"))
		}
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
	if (tablerow.Node == nil) != (tablerowOther.Node == nil) {
		diffs = append(diffs, tablerow.GongMarshallField(stage, "Node"))
	} else if tablerow.Node != nil && tablerowOther.Node != nil {
		if tablerow.Node != tablerowOther.Node {
			diffs = append(diffs, tablerow.GongMarshallField(stage, "Node"))
		}
	}
	TableColumnsDifferent := false
	if len(tablerow.TableColumns) != len(tablerowOther.TableColumns) {
		TableColumnsDifferent = true
	} else {
		for i := range tablerow.TableColumns {
			if (tablerow.TableColumns[i] == nil) != (tablerowOther.TableColumns[i] == nil) {
				TableColumnsDifferent = true
				break
			} else if tablerow.TableColumns[i] != nil && tablerowOther.TableColumns[i] != nil {
				// this is a pointer comparaison
				if tablerow.TableColumns[i] != tablerowOther.TableColumns[i] {
					TableColumnsDifferent = true
					break
				}
			}
		}
	}
	if TableColumnsDifferent {
		ops := stage.Diff(
			tablerow,
			"TableColumns",
			len(tablerowOther.TableColumns),
			len(tablerow.TableColumns),
			func(i, j int) bool {
				return tablerowOther.TableColumns[i] == tablerow.TableColumns[j]
			},
			func(j int) string {
				return tablerow.TableColumns[j].GongGetIdentifier(stage)
			},
		)
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
	if (tablestyle.Node == nil) != (tablestyleOther.Node == nil) {
		diffs = append(diffs, tablestyle.GongMarshallField(stage, "Node"))
	} else if tablestyle.Node != nil && tablestyleOther.Node != nil {
		if tablestyle.Node != tablestyleOther.Node {
			diffs = append(diffs, tablestyle.GongMarshallField(stage, "Node"))
		}
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
	if (text.Node == nil) != (textOther.Node == nil) {
		diffs = append(diffs, text.GongMarshallField(stage, "Node"))
	} else if text.Node != nil && textOther.Node != nil {
		if text.Node != textOther.Node {
			diffs = append(diffs, text.GongMarshallField(stage, "Node"))
		}
	}
	if text.PreserveWhiteSpace != textOther.PreserveWhiteSpace {
		diffs = append(diffs, text.GongMarshallField(stage, "PreserveWhiteSpace"))
	}
	if (text.EnclosingRune == nil) != (textOther.EnclosingRune == nil) {
		diffs = append(diffs, text.GongMarshallField(stage, "EnclosingRune"))
	} else if text.EnclosingRune != nil && textOther.EnclosingRune != nil {
		if text.EnclosingRune != textOther.EnclosingRune {
			diffs = append(diffs, text.GongMarshallField(stage, "EnclosingRune"))
		}
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

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
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
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := 0; k < n; k++ {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
