// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Body:
		ok = stage.IsStagedBody(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *Docx:
		ok = stage.IsStagedDocx(target)

	case *File:
		ok = stage.IsStagedFile(target)

	case *Node:
		ok = stage.IsStagedNode(target)

	case *Paragraph:
		ok = stage.IsStagedParagraph(target)

	case *ParagraphProperties:
		ok = stage.IsStagedParagraphProperties(target)

	case *ParagraphStyle:
		ok = stage.IsStagedParagraphStyle(target)

	case *Rune:
		ok = stage.IsStagedRune(target)

	case *RuneProperties:
		ok = stage.IsStagedRuneProperties(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *TableColumn:
		ok = stage.IsStagedTableColumn(target)

	case *TableProperties:
		ok = stage.IsStagedTableProperties(target)

	case *TableRow:
		ok = stage.IsStagedTableRow(target)

	case *TableStyle:
		ok = stage.IsStagedTableStyle(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Body:
		ok = stage.IsStagedBody(target)

	case *Document:
		ok = stage.IsStagedDocument(target)

	case *Docx:
		ok = stage.IsStagedDocx(target)

	case *File:
		ok = stage.IsStagedFile(target)

	case *Node:
		ok = stage.IsStagedNode(target)

	case *Paragraph:
		ok = stage.IsStagedParagraph(target)

	case *ParagraphProperties:
		ok = stage.IsStagedParagraphProperties(target)

	case *ParagraphStyle:
		ok = stage.IsStagedParagraphStyle(target)

	case *Rune:
		ok = stage.IsStagedRune(target)

	case *RuneProperties:
		ok = stage.IsStagedRuneProperties(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	case *TableColumn:
		ok = stage.IsStagedTableColumn(target)

	case *TableProperties:
		ok = stage.IsStagedTableProperties(target)

	case *TableRow:
		ok = stage.IsStagedTableRow(target)

	case *TableStyle:
		ok = stage.IsStagedTableStyle(target)

	case *Text:
		ok = stage.IsStagedText(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedBody(body *Body) (ok bool) {

	_, ok = stage.Bodys[body]

	return
}

func (stage *Stage) IsStagedDocument(document *Document) (ok bool) {

	_, ok = stage.Documents[document]

	return
}

func (stage *Stage) IsStagedDocx(docx *Docx) (ok bool) {

	_, ok = stage.Docxs[docx]

	return
}

func (stage *Stage) IsStagedFile(file *File) (ok bool) {

	_, ok = stage.Files[file]

	return
}

func (stage *Stage) IsStagedNode(node *Node) (ok bool) {

	_, ok = stage.Nodes[node]

	return
}

func (stage *Stage) IsStagedParagraph(paragraph *Paragraph) (ok bool) {

	_, ok = stage.Paragraphs[paragraph]

	return
}

func (stage *Stage) IsStagedParagraphProperties(paragraphproperties *ParagraphProperties) (ok bool) {

	_, ok = stage.ParagraphPropertiess[paragraphproperties]

	return
}

func (stage *Stage) IsStagedParagraphStyle(paragraphstyle *ParagraphStyle) (ok bool) {

	_, ok = stage.ParagraphStyles[paragraphstyle]

	return
}

func (stage *Stage) IsStagedRune(rune *Rune) (ok bool) {

	_, ok = stage.Runes[rune]

	return
}

func (stage *Stage) IsStagedRuneProperties(runeproperties *RuneProperties) (ok bool) {

	_, ok = stage.RunePropertiess[runeproperties]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *Stage) IsStagedTableColumn(tablecolumn *TableColumn) (ok bool) {

	_, ok = stage.TableColumns[tablecolumn]

	return
}

func (stage *Stage) IsStagedTableProperties(tableproperties *TableProperties) (ok bool) {

	_, ok = stage.TablePropertiess[tableproperties]

	return
}

func (stage *Stage) IsStagedTableRow(tablerow *TableRow) (ok bool) {

	_, ok = stage.TableRows[tablerow]

	return
}

func (stage *Stage) IsStagedTableStyle(tablestyle *TableStyle) (ok bool) {

	_, ok = stage.TableStyles[tablestyle]

	return
}

func (stage *Stage) IsStagedText(text *Text) (ok bool) {

	_, ok = stage.Texts[text]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Body:
		stage.StageBranchBody(target)

	case *Document:
		stage.StageBranchDocument(target)

	case *Docx:
		stage.StageBranchDocx(target)

	case *File:
		stage.StageBranchFile(target)

	case *Node:
		stage.StageBranchNode(target)

	case *Paragraph:
		stage.StageBranchParagraph(target)

	case *ParagraphProperties:
		stage.StageBranchParagraphProperties(target)

	case *ParagraphStyle:
		stage.StageBranchParagraphStyle(target)

	case *Rune:
		stage.StageBranchRune(target)

	case *RuneProperties:
		stage.StageBranchRuneProperties(target)

	case *Table:
		stage.StageBranchTable(target)

	case *TableColumn:
		stage.StageBranchTableColumn(target)

	case *TableProperties:
		stage.StageBranchTableProperties(target)

	case *TableRow:
		stage.StageBranchTableRow(target)

	case *TableStyle:
		stage.StageBranchTableStyle(target)

	case *Text:
		stage.StageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchBody(body *Body) {

	// check if instance is already staged
	if IsStaged(stage, body) {
		return
	}

	body.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		StageBranch(stage, body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		StageBranch(stage, _paragraph)
	}
	for _, _table := range body.Tables {
		StageBranch(stage, _table)
	}

}

func (stage *Stage) StageBranchDocument(document *Document) {

	// check if instance is already staged
	if IsStaged(stage, document) {
		return
	}

	document.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		StageBranch(stage, document.File)
	}
	if document.Root != nil {
		StageBranch(stage, document.Root)
	}
	if document.Body != nil {
		StageBranch(stage, document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDocx(docx *Docx) {

	// check if instance is already staged
	if IsStaged(stage, docx) {
		return
	}

	docx.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		StageBranch(stage, docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		StageBranch(stage, _file)
	}

}

func (stage *Stage) StageBranchFile(file *File) {

	// check if instance is already staged
	if IsStaged(stage, file) {
		return
	}

	file.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNode(node *Node) {

	// check if instance is already staged
	if IsStaged(stage, node) {
		return
	}

	node.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		StageBranch(stage, _node)
	}

}

func (stage *Stage) StageBranchParagraph(paragraph *Paragraph) {

	// check if instance is already staged
	if IsStaged(stage, paragraph) {
		return
	}

	paragraph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		StageBranch(stage, paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		StageBranch(stage, paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		StageBranch(stage, paragraph.Next)
	}
	if paragraph.Previous != nil {
		StageBranch(stage, paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		StageBranch(stage, paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		StageBranch(stage, paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		StageBranch(stage, _rune)
	}

}

func (stage *Stage) StageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

	// check if instance is already staged
	if IsStaged(stage, paragraphproperties) {
		return
	}

	paragraphproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		StageBranch(stage, paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		StageBranch(stage, paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

	// check if instance is already staged
	if IsStaged(stage, paragraphstyle) {
		return
	}

	paragraphstyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		StageBranch(stage, paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRune(rune *Rune) {

	// check if instance is already staged
	if IsStaged(stage, rune) {
		return
	}

	rune.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		StageBranch(stage, rune.Node)
	}
	if rune.Text != nil {
		StageBranch(stage, rune.Text)
	}
	if rune.RuneProperties != nil {
		StageBranch(stage, rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		StageBranch(stage, rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRuneProperties(runeproperties *RuneProperties) {

	// check if instance is already staged
	if IsStaged(stage, runeproperties) {
		return
	}

	runeproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		StageBranch(stage, runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTable(table *Table) {

	// check if instance is already staged
	if IsStaged(stage, table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		StageBranch(stage, table.Node)
	}
	if table.TableProperties != nil {
		StageBranch(stage, table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		StageBranch(stage, _tablerow)
	}

}

func (stage *Stage) StageBranchTableColumn(tablecolumn *TableColumn) {

	// check if instance is already staged
	if IsStaged(stage, tablecolumn) {
		return
	}

	tablecolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		StageBranch(stage, tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		StageBranch(stage, _paragraph)
	}

}

func (stage *Stage) StageBranchTableProperties(tableproperties *TableProperties) {

	// check if instance is already staged
	if IsStaged(stage, tableproperties) {
		return
	}

	tableproperties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		StageBranch(stage, tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		StageBranch(stage, tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTableRow(tablerow *TableRow) {

	// check if instance is already staged
	if IsStaged(stage, tablerow) {
		return
	}

	tablerow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		StageBranch(stage, tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		StageBranch(stage, _tablecolumn)
	}

}

func (stage *Stage) StageBranchTableStyle(tablestyle *TableStyle) {

	// check if instance is already staged
	if IsStaged(stage, tablestyle) {
		return
	}

	tablestyle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		StageBranch(stage, tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchText(text *Text) {

	// check if instance is already staged
	if IsStaged(stage, text) {
		return
	}

	text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		StageBranch(stage, text.Node)
	}
	if text.EnclosingRune != nil {
		StageBranch(stage, text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Body:
		toT := CopyBranchBody(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Document:
		toT := CopyBranchDocument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Docx:
		toT := CopyBranchDocx(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *File:
		toT := CopyBranchFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Node:
		toT := CopyBranchNode(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Paragraph:
		toT := CopyBranchParagraph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphProperties:
		toT := CopyBranchParagraphProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParagraphStyle:
		toT := CopyBranchParagraphStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rune:
		toT := CopyBranchRune(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RuneProperties:
		toT := CopyBranchRuneProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := CopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableColumn:
		toT := CopyBranchTableColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableProperties:
		toT := CopyBranchTableProperties(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableRow:
		toT := CopyBranchTableRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TableStyle:
		toT := CopyBranchTableStyle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text:
		toT := CopyBranchText(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBody(mapOrigCopy map[any]any, bodyFrom *Body) (bodyTo *Body) {

	// bodyFrom has already been copied
	if _bodyTo, ok := mapOrigCopy[bodyFrom]; ok {
		bodyTo = _bodyTo.(*Body)
		return
	}

	bodyTo = new(Body)
	mapOrigCopy[bodyFrom] = bodyTo
	bodyFrom.CopyBasicFields(bodyTo)

	//insertion point for the staging of instances referenced by pointers
	if bodyFrom.LastParagraph != nil {
		bodyTo.LastParagraph = CopyBranchParagraph(mapOrigCopy, bodyFrom.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range bodyFrom.Paragraphs {
		bodyTo.Paragraphs = append(bodyTo.Paragraphs, CopyBranchParagraph(mapOrigCopy, _paragraph))
	}
	for _, _table := range bodyFrom.Tables {
		bodyTo.Tables = append(bodyTo.Tables, CopyBranchTable(mapOrigCopy, _table))
	}

	return
}

func CopyBranchDocument(mapOrigCopy map[any]any, documentFrom *Document) (documentTo *Document) {

	// documentFrom has already been copied
	if _documentTo, ok := mapOrigCopy[documentFrom]; ok {
		documentTo = _documentTo.(*Document)
		return
	}

	documentTo = new(Document)
	mapOrigCopy[documentFrom] = documentTo
	documentFrom.CopyBasicFields(documentTo)

	//insertion point for the staging of instances referenced by pointers
	if documentFrom.File != nil {
		documentTo.File = CopyBranchFile(mapOrigCopy, documentFrom.File)
	}
	if documentFrom.Root != nil {
		documentTo.Root = CopyBranchNode(mapOrigCopy, documentFrom.Root)
	}
	if documentFrom.Body != nil {
		documentTo.Body = CopyBranchBody(mapOrigCopy, documentFrom.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDocx(mapOrigCopy map[any]any, docxFrom *Docx) (docxTo *Docx) {

	// docxFrom has already been copied
	if _docxTo, ok := mapOrigCopy[docxFrom]; ok {
		docxTo = _docxTo.(*Docx)
		return
	}

	docxTo = new(Docx)
	mapOrigCopy[docxFrom] = docxTo
	docxFrom.CopyBasicFields(docxTo)

	//insertion point for the staging of instances referenced by pointers
	if docxFrom.Document != nil {
		docxTo.Document = CopyBranchDocument(mapOrigCopy, docxFrom.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docxFrom.Files {
		docxTo.Files = append(docxTo.Files, CopyBranchFile(mapOrigCopy, _file))
	}

	return
}

func CopyBranchFile(mapOrigCopy map[any]any, fileFrom *File) (fileTo *File) {

	// fileFrom has already been copied
	if _fileTo, ok := mapOrigCopy[fileFrom]; ok {
		fileTo = _fileTo.(*File)
		return
	}

	fileTo = new(File)
	mapOrigCopy[fileFrom] = fileTo
	fileFrom.CopyBasicFields(fileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNode(mapOrigCopy map[any]any, nodeFrom *Node) (nodeTo *Node) {

	// nodeFrom has already been copied
	if _nodeTo, ok := mapOrigCopy[nodeFrom]; ok {
		nodeTo = _nodeTo.(*Node)
		return
	}

	nodeTo = new(Node)
	mapOrigCopy[nodeFrom] = nodeTo
	nodeFrom.CopyBasicFields(nodeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range nodeFrom.Nodes {
		nodeTo.Nodes = append(nodeTo.Nodes, CopyBranchNode(mapOrigCopy, _node))
	}

	return
}

func CopyBranchParagraph(mapOrigCopy map[any]any, paragraphFrom *Paragraph) (paragraphTo *Paragraph) {

	// paragraphFrom has already been copied
	if _paragraphTo, ok := mapOrigCopy[paragraphFrom]; ok {
		paragraphTo = _paragraphTo.(*Paragraph)
		return
	}

	paragraphTo = new(Paragraph)
	mapOrigCopy[paragraphFrom] = paragraphTo
	paragraphFrom.CopyBasicFields(paragraphTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphFrom.Node != nil {
		paragraphTo.Node = CopyBranchNode(mapOrigCopy, paragraphFrom.Node)
	}
	if paragraphFrom.ParagraphProperties != nil {
		paragraphTo.ParagraphProperties = CopyBranchParagraphProperties(mapOrigCopy, paragraphFrom.ParagraphProperties)
	}
	if paragraphFrom.Next != nil {
		paragraphTo.Next = CopyBranchParagraph(mapOrigCopy, paragraphFrom.Next)
	}
	if paragraphFrom.Previous != nil {
		paragraphTo.Previous = CopyBranchParagraph(mapOrigCopy, paragraphFrom.Previous)
	}
	if paragraphFrom.EnclosingBody != nil {
		paragraphTo.EnclosingBody = CopyBranchBody(mapOrigCopy, paragraphFrom.EnclosingBody)
	}
	if paragraphFrom.EnclosingTableColumn != nil {
		paragraphTo.EnclosingTableColumn = CopyBranchTableColumn(mapOrigCopy, paragraphFrom.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraphFrom.Runes {
		paragraphTo.Runes = append(paragraphTo.Runes, CopyBranchRune(mapOrigCopy, _rune))
	}

	return
}

func CopyBranchParagraphProperties(mapOrigCopy map[any]any, paragraphpropertiesFrom *ParagraphProperties) (paragraphpropertiesTo *ParagraphProperties) {

	// paragraphpropertiesFrom has already been copied
	if _paragraphpropertiesTo, ok := mapOrigCopy[paragraphpropertiesFrom]; ok {
		paragraphpropertiesTo = _paragraphpropertiesTo.(*ParagraphProperties)
		return
	}

	paragraphpropertiesTo = new(ParagraphProperties)
	mapOrigCopy[paragraphpropertiesFrom] = paragraphpropertiesTo
	paragraphpropertiesFrom.CopyBasicFields(paragraphpropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphpropertiesFrom.ParagraphStyle != nil {
		paragraphpropertiesTo.ParagraphStyle = CopyBranchParagraphStyle(mapOrigCopy, paragraphpropertiesFrom.ParagraphStyle)
	}
	if paragraphpropertiesFrom.Node != nil {
		paragraphpropertiesTo.Node = CopyBranchNode(mapOrigCopy, paragraphpropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchParagraphStyle(mapOrigCopy map[any]any, paragraphstyleFrom *ParagraphStyle) (paragraphstyleTo *ParagraphStyle) {

	// paragraphstyleFrom has already been copied
	if _paragraphstyleTo, ok := mapOrigCopy[paragraphstyleFrom]; ok {
		paragraphstyleTo = _paragraphstyleTo.(*ParagraphStyle)
		return
	}

	paragraphstyleTo = new(ParagraphStyle)
	mapOrigCopy[paragraphstyleFrom] = paragraphstyleTo
	paragraphstyleFrom.CopyBasicFields(paragraphstyleTo)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyleFrom.Node != nil {
		paragraphstyleTo.Node = CopyBranchNode(mapOrigCopy, paragraphstyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRune(mapOrigCopy map[any]any, runeFrom *Rune) (runeTo *Rune) {

	// runeFrom has already been copied
	if _runeTo, ok := mapOrigCopy[runeFrom]; ok {
		runeTo = _runeTo.(*Rune)
		return
	}

	runeTo = new(Rune)
	mapOrigCopy[runeFrom] = runeTo
	runeFrom.CopyBasicFields(runeTo)

	//insertion point for the staging of instances referenced by pointers
	if runeFrom.Node != nil {
		runeTo.Node = CopyBranchNode(mapOrigCopy, runeFrom.Node)
	}
	if runeFrom.Text != nil {
		runeTo.Text = CopyBranchText(mapOrigCopy, runeFrom.Text)
	}
	if runeFrom.RuneProperties != nil {
		runeTo.RuneProperties = CopyBranchRuneProperties(mapOrigCopy, runeFrom.RuneProperties)
	}
	if runeFrom.EnclosingParagraph != nil {
		runeTo.EnclosingParagraph = CopyBranchParagraph(mapOrigCopy, runeFrom.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRuneProperties(mapOrigCopy map[any]any, runepropertiesFrom *RuneProperties) (runepropertiesTo *RuneProperties) {

	// runepropertiesFrom has already been copied
	if _runepropertiesTo, ok := mapOrigCopy[runepropertiesFrom]; ok {
		runepropertiesTo = _runepropertiesTo.(*RuneProperties)
		return
	}

	runepropertiesTo = new(RuneProperties)
	mapOrigCopy[runepropertiesFrom] = runepropertiesTo
	runepropertiesFrom.CopyBasicFields(runepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if runepropertiesFrom.Node != nil {
		runepropertiesTo.Node = CopyBranchNode(mapOrigCopy, runepropertiesFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.CopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers
	if tableFrom.Node != nil {
		tableTo.Node = CopyBranchNode(mapOrigCopy, tableFrom.Node)
	}
	if tableFrom.TableProperties != nil {
		tableTo.TableProperties = CopyBranchTableProperties(mapOrigCopy, tableFrom.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range tableFrom.TableRows {
		tableTo.TableRows = append(tableTo.TableRows, CopyBranchTableRow(mapOrigCopy, _tablerow))
	}

	return
}

func CopyBranchTableColumn(mapOrigCopy map[any]any, tablecolumnFrom *TableColumn) (tablecolumnTo *TableColumn) {

	// tablecolumnFrom has already been copied
	if _tablecolumnTo, ok := mapOrigCopy[tablecolumnFrom]; ok {
		tablecolumnTo = _tablecolumnTo.(*TableColumn)
		return
	}

	tablecolumnTo = new(TableColumn)
	mapOrigCopy[tablecolumnFrom] = tablecolumnTo
	tablecolumnFrom.CopyBasicFields(tablecolumnTo)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumnFrom.Node != nil {
		tablecolumnTo.Node = CopyBranchNode(mapOrigCopy, tablecolumnFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumnFrom.Paragraphs {
		tablecolumnTo.Paragraphs = append(tablecolumnTo.Paragraphs, CopyBranchParagraph(mapOrigCopy, _paragraph))
	}

	return
}

func CopyBranchTableProperties(mapOrigCopy map[any]any, tablepropertiesFrom *TableProperties) (tablepropertiesTo *TableProperties) {

	// tablepropertiesFrom has already been copied
	if _tablepropertiesTo, ok := mapOrigCopy[tablepropertiesFrom]; ok {
		tablepropertiesTo = _tablepropertiesTo.(*TableProperties)
		return
	}

	tablepropertiesTo = new(TableProperties)
	mapOrigCopy[tablepropertiesFrom] = tablepropertiesTo
	tablepropertiesFrom.CopyBasicFields(tablepropertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if tablepropertiesFrom.Node != nil {
		tablepropertiesTo.Node = CopyBranchNode(mapOrigCopy, tablepropertiesFrom.Node)
	}
	if tablepropertiesFrom.TableStyle != nil {
		tablepropertiesTo.TableStyle = CopyBranchTableStyle(mapOrigCopy, tablepropertiesFrom.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTableRow(mapOrigCopy map[any]any, tablerowFrom *TableRow) (tablerowTo *TableRow) {

	// tablerowFrom has already been copied
	if _tablerowTo, ok := mapOrigCopy[tablerowFrom]; ok {
		tablerowTo = _tablerowTo.(*TableRow)
		return
	}

	tablerowTo = new(TableRow)
	mapOrigCopy[tablerowFrom] = tablerowTo
	tablerowFrom.CopyBasicFields(tablerowTo)

	//insertion point for the staging of instances referenced by pointers
	if tablerowFrom.Node != nil {
		tablerowTo.Node = CopyBranchNode(mapOrigCopy, tablerowFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerowFrom.TableColumns {
		tablerowTo.TableColumns = append(tablerowTo.TableColumns, CopyBranchTableColumn(mapOrigCopy, _tablecolumn))
	}

	return
}

func CopyBranchTableStyle(mapOrigCopy map[any]any, tablestyleFrom *TableStyle) (tablestyleTo *TableStyle) {

	// tablestyleFrom has already been copied
	if _tablestyleTo, ok := mapOrigCopy[tablestyleFrom]; ok {
		tablestyleTo = _tablestyleTo.(*TableStyle)
		return
	}

	tablestyleTo = new(TableStyle)
	mapOrigCopy[tablestyleFrom] = tablestyleTo
	tablestyleFrom.CopyBasicFields(tablestyleTo)

	//insertion point for the staging of instances referenced by pointers
	if tablestyleFrom.Node != nil {
		tablestyleTo.Node = CopyBranchNode(mapOrigCopy, tablestyleFrom.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchText(mapOrigCopy map[any]any, textFrom *Text) (textTo *Text) {

	// textFrom has already been copied
	if _textTo, ok := mapOrigCopy[textFrom]; ok {
		textTo = _textTo.(*Text)
		return
	}

	textTo = new(Text)
	mapOrigCopy[textFrom] = textTo
	textFrom.CopyBasicFields(textTo)

	//insertion point for the staging of instances referenced by pointers
	if textFrom.Node != nil {
		textTo.Node = CopyBranchNode(mapOrigCopy, textFrom.Node)
	}
	if textFrom.EnclosingRune != nil {
		textTo.EnclosingRune = CopyBranchRune(mapOrigCopy, textFrom.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Body:
		stage.UnstageBranchBody(target)

	case *Document:
		stage.UnstageBranchDocument(target)

	case *Docx:
		stage.UnstageBranchDocx(target)

	case *File:
		stage.UnstageBranchFile(target)

	case *Node:
		stage.UnstageBranchNode(target)

	case *Paragraph:
		stage.UnstageBranchParagraph(target)

	case *ParagraphProperties:
		stage.UnstageBranchParagraphProperties(target)

	case *ParagraphStyle:
		stage.UnstageBranchParagraphStyle(target)

	case *Rune:
		stage.UnstageBranchRune(target)

	case *RuneProperties:
		stage.UnstageBranchRuneProperties(target)

	case *Table:
		stage.UnstageBranchTable(target)

	case *TableColumn:
		stage.UnstageBranchTableColumn(target)

	case *TableProperties:
		stage.UnstageBranchTableProperties(target)

	case *TableRow:
		stage.UnstageBranchTableRow(target)

	case *TableStyle:
		stage.UnstageBranchTableStyle(target)

	case *Text:
		stage.UnstageBranchText(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchBody(body *Body) {

	// check if instance is already staged
	if !IsStaged(stage, body) {
		return
	}

	body.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if body.LastParagraph != nil {
		UnstageBranch(stage, body.LastParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range body.Paragraphs {
		UnstageBranch(stage, _paragraph)
	}
	for _, _table := range body.Tables {
		UnstageBranch(stage, _table)
	}

}

func (stage *Stage) UnstageBranchDocument(document *Document) {

	// check if instance is already staged
	if !IsStaged(stage, document) {
		return
	}

	document.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if document.File != nil {
		UnstageBranch(stage, document.File)
	}
	if document.Root != nil {
		UnstageBranch(stage, document.Root)
	}
	if document.Body != nil {
		UnstageBranch(stage, document.Body)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDocx(docx *Docx) {

	// check if instance is already staged
	if !IsStaged(stage, docx) {
		return
	}

	docx.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if docx.Document != nil {
		UnstageBranch(stage, docx.Document)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _file := range docx.Files {
		UnstageBranch(stage, _file)
	}

}

func (stage *Stage) UnstageBranchFile(file *File) {

	// check if instance is already staged
	if !IsStaged(stage, file) {
		return
	}

	file.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNode(node *Node) {

	// check if instance is already staged
	if !IsStaged(stage, node) {
		return
	}

	node.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _node := range node.Nodes {
		UnstageBranch(stage, _node)
	}

}

func (stage *Stage) UnstageBranchParagraph(paragraph *Paragraph) {

	// check if instance is already staged
	if !IsStaged(stage, paragraph) {
		return
	}

	paragraph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraph.Node != nil {
		UnstageBranch(stage, paragraph.Node)
	}
	if paragraph.ParagraphProperties != nil {
		UnstageBranch(stage, paragraph.ParagraphProperties)
	}
	if paragraph.Next != nil {
		UnstageBranch(stage, paragraph.Next)
	}
	if paragraph.Previous != nil {
		UnstageBranch(stage, paragraph.Previous)
	}
	if paragraph.EnclosingBody != nil {
		UnstageBranch(stage, paragraph.EnclosingBody)
	}
	if paragraph.EnclosingTableColumn != nil {
		UnstageBranch(stage, paragraph.EnclosingTableColumn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rune := range paragraph.Runes {
		UnstageBranch(stage, _rune)
	}

}

func (stage *Stage) UnstageBranchParagraphProperties(paragraphproperties *ParagraphProperties) {

	// check if instance is already staged
	if !IsStaged(stage, paragraphproperties) {
		return
	}

	paragraphproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphproperties.ParagraphStyle != nil {
		UnstageBranch(stage, paragraphproperties.ParagraphStyle)
	}
	if paragraphproperties.Node != nil {
		UnstageBranch(stage, paragraphproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchParagraphStyle(paragraphstyle *ParagraphStyle) {

	// check if instance is already staged
	if !IsStaged(stage, paragraphstyle) {
		return
	}

	paragraphstyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if paragraphstyle.Node != nil {
		UnstageBranch(stage, paragraphstyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRune(rune *Rune) {

	// check if instance is already staged
	if !IsStaged(stage, rune) {
		return
	}

	rune.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if rune.Node != nil {
		UnstageBranch(stage, rune.Node)
	}
	if rune.Text != nil {
		UnstageBranch(stage, rune.Text)
	}
	if rune.RuneProperties != nil {
		UnstageBranch(stage, rune.RuneProperties)
	}
	if rune.EnclosingParagraph != nil {
		UnstageBranch(stage, rune.EnclosingParagraph)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRuneProperties(runeproperties *RuneProperties) {

	// check if instance is already staged
	if !IsStaged(stage, runeproperties) {
		return
	}

	runeproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if runeproperties.Node != nil {
		UnstageBranch(stage, runeproperties.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !IsStaged(stage, table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if table.Node != nil {
		UnstageBranch(stage, table.Node)
	}
	if table.TableProperties != nil {
		UnstageBranch(stage, table.TableProperties)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablerow := range table.TableRows {
		UnstageBranch(stage, _tablerow)
	}

}

func (stage *Stage) UnstageBranchTableColumn(tablecolumn *TableColumn) {

	// check if instance is already staged
	if !IsStaged(stage, tablecolumn) {
		return
	}

	tablecolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablecolumn.Node != nil {
		UnstageBranch(stage, tablecolumn.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _paragraph := range tablecolumn.Paragraphs {
		UnstageBranch(stage, _paragraph)
	}

}

func (stage *Stage) UnstageBranchTableProperties(tableproperties *TableProperties) {

	// check if instance is already staged
	if !IsStaged(stage, tableproperties) {
		return
	}

	tableproperties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tableproperties.Node != nil {
		UnstageBranch(stage, tableproperties.Node)
	}
	if tableproperties.TableStyle != nil {
		UnstageBranch(stage, tableproperties.TableStyle)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTableRow(tablerow *TableRow) {

	// check if instance is already staged
	if !IsStaged(stage, tablerow) {
		return
	}

	tablerow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablerow.Node != nil {
		UnstageBranch(stage, tablerow.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tablecolumn := range tablerow.TableColumns {
		UnstageBranch(stage, _tablecolumn)
	}

}

func (stage *Stage) UnstageBranchTableStyle(tablestyle *TableStyle) {

	// check if instance is already staged
	if !IsStaged(stage, tablestyle) {
		return
	}

	tablestyle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tablestyle.Node != nil {
		UnstageBranch(stage, tablestyle.Node)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchText(text *Text) {

	// check if instance is already staged
	if !IsStaged(stage, text) {
		return
	}

	text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if text.Node != nil {
		UnstageBranch(stage, text.Node)
	}
	if text.EnclosingRune != nil {
		UnstageBranch(stage, text.EnclosingRune)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Body) GongReconstructPointersFromReferences(stage *Stage, instance *Body) () {
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

	return
}

func (reference *Document) GongReconstructPointersFromReferences(stage *Stage, instance *Document) () {
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

	return
}

func (reference *Docx) GongReconstructPointersFromReferences(stage *Stage, instance *Docx) () {
	// insertion point for pointers field
	if instance.Document != nil {
		reference.Document = stage.Documents_reference[instance.Document]
	}
	// insertion point for slice of pointers field
	reference.Files = reference.Files[:0]
	for _, _b := range instance.Files {
		reference.Files = append(reference.Files, stage.Files_reference[_b])
	}

	return
}

func (reference *File) GongReconstructPointersFromReferences(stage *Stage, instance *File) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Node) GongReconstructPointersFromReferences(stage *Stage, instance *Node) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Nodes = reference.Nodes[:0]
	for _, _b := range instance.Nodes {
		reference.Nodes = append(reference.Nodes, stage.Nodes_reference[_b])
	}

	return
}

func (reference *Paragraph) GongReconstructPointersFromReferences(stage *Stage, instance *Paragraph) () {
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

	return
}

func (reference *ParagraphProperties) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphProperties) () {
	// insertion point for pointers field
	if instance.ParagraphStyle != nil {
		reference.ParagraphStyle = stage.ParagraphStyles_reference[instance.ParagraphStyle]
	}
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *ParagraphStyle) GongReconstructPointersFromReferences(stage *Stage, instance *ParagraphStyle) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Rune) GongReconstructPointersFromReferences(stage *Stage, instance *Rune) () {
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

	return
}

func (reference *RuneProperties) GongReconstructPointersFromReferences(stage *Stage, instance *RuneProperties) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) () {
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

	return
}

func (reference *TableColumn) GongReconstructPointersFromReferences(stage *Stage, instance *TableColumn) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
	reference.Paragraphs = reference.Paragraphs[:0]
	for _, _b := range instance.Paragraphs {
		reference.Paragraphs = append(reference.Paragraphs, stage.Paragraphs_reference[_b])
	}

	return
}

func (reference *TableProperties) GongReconstructPointersFromReferences(stage *Stage, instance *TableProperties) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.TableStyle != nil {
		reference.TableStyle = stage.TableStyles_reference[instance.TableStyle]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *TableRow) GongReconstructPointersFromReferences(stage *Stage, instance *TableRow) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field
	reference.TableColumns = reference.TableColumns[:0]
	for _, _b := range instance.TableColumns {
		reference.TableColumns = append(reference.TableColumns, stage.TableColumns_reference[_b])
	}

	return
}

func (reference *TableStyle) GongReconstructPointersFromReferences(stage *Stage, instance *TableStyle) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Text) GongReconstructPointersFromReferences(stage *Stage, instance *Text) () {
	// insertion point for pointers field
	if instance.Node != nil {
		reference.Node = stage.Nodes_reference[instance.Node]
	}
	if instance.EnclosingRune != nil {
		reference.EnclosingRune = stage.Runes_reference[instance.EnclosingRune]
	}
	// insertion point for slice of pointers field

	return
}

// insertion point for pointer reconstruction from instances
func (reference *Body) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_Paragraphs = append(_Paragraphs, stage.Paragraphs_reference[_instance])
		}
	}
	reference.Paragraphs = _Paragraphs
	var _Tables []*Table
	for _, _reference := range reference.Tables {
		if _instance, ok := stage.Tables_instance[_reference]; ok {
			_Tables = append(_Tables, stage.Tables_reference[_instance])
		}
	}
	reference.Tables = _Tables

	return
}

func (reference *Document) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
}

func (reference *Docx) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_Files = append(_Files, stage.Files_reference[_instance])
		}
	}
	reference.Files = _Files

	return
}

func (reference *File) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Node) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Nodes []*Node
	for _, _reference := range reference.Nodes {
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			_Nodes = append(_Nodes, stage.Nodes_reference[_instance])
		}
	}
	reference.Nodes = _Nodes

	return
}

func (reference *Paragraph) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_Runes = append(_Runes, stage.Runes_reference[_instance])
		}
	}
	reference.Runes = _Runes

	return
}

func (reference *ParagraphProperties) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
}

func (reference *ParagraphStyle) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Rune) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
}

func (reference *RuneProperties) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_TableRows = append(_TableRows, stage.TableRows_reference[_instance])
		}
	}
	reference.TableRows = _TableRows

	return
}

func (reference *TableColumn) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_Paragraphs = append(_Paragraphs, stage.Paragraphs_reference[_instance])
		}
	}
	reference.Paragraphs = _Paragraphs

	return
}

func (reference *TableProperties) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
}

func (reference *TableRow) GongReconstructPointersFromInstances(stage *Stage) () {
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
			_TableColumns = append(_TableColumns, stage.TableColumns_reference[_instance])
		}
	}
	reference.TableColumns = _TableColumns

	return
}

func (reference *TableStyle) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Node; _reference != nil {
		reference.Node = nil
		if _instance, ok := stage.Nodes_instance[_reference]; ok {
			reference.Node = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Text) GongReconstructPointersFromInstances(stage *Stage) () {
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

	return
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
		ops := Diff(stage, body, bodyOther, "Paragraphs", bodyOther.Paragraphs, body.Paragraphs)
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
		ops := Diff(stage, body, bodyOther, "Tables", bodyOther.Tables, body.Tables)
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
		ops := Diff(stage, docx, docxOther, "Files", docxOther.Files, docx.Files)
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
		ops := Diff(stage, node, nodeOther, "Nodes", nodeOther.Nodes, node.Nodes)
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
		ops := Diff(stage, paragraph, paragraphOther, "Runes", paragraphOther.Runes, paragraph.Runes)
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
		ops := Diff(stage, table, tableOther, "TableRows", tableOther.TableRows, table.TableRows)
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
		ops := Diff(stage, tablecolumn, tablecolumnOther, "Paragraphs", tablecolumnOther.Paragraphs, tablecolumn.Paragraphs)
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
		ops := Diff(stage, tablerow, tablerowOther, "TableColumns", tablerowOther.TableColumns, tablerow.TableColumns)
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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
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
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
