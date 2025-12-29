// generated code - do not edit
package models

import (
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var bodys_newInstances []*Body
	var bodys_deletedInstances []*Body

	// parse all staged instances and check if they have a reference
	for body := range stage.Bodys {
		if ref, ok := stage.Bodys_reference[body]; !ok {
			bodys_newInstances = append(bodys_newInstances, body)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Body "+body.Name,
				)
			}
		} else {
			diffs := body.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Body \""+body.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for body := range stage.Bodys_reference {
		if _, ok := stage.Bodys[body]; !ok {
			bodys_deletedInstances = append(bodys_deletedInstances, body)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Body "+body.Name,
				)
			}
		}
	}

	lenNewInstances += len(bodys_newInstances)
	lenDeletedInstances += len(bodys_deletedInstances)
	var documents_newInstances []*Document
	var documents_deletedInstances []*Document

	// parse all staged instances and check if they have a reference
	for document := range stage.Documents {
		if ref, ok := stage.Documents_reference[document]; !ok {
			documents_newInstances = append(documents_newInstances, document)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Document "+document.Name,
				)
			}
		} else {
			diffs := document.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Document \""+document.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for document := range stage.Documents_reference {
		if _, ok := stage.Documents[document]; !ok {
			documents_deletedInstances = append(documents_deletedInstances, document)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Document "+document.Name,
				)
			}
		}
	}

	lenNewInstances += len(documents_newInstances)
	lenDeletedInstances += len(documents_deletedInstances)
	var docxs_newInstances []*Docx
	var docxs_deletedInstances []*Docx

	// parse all staged instances and check if they have a reference
	for docx := range stage.Docxs {
		if ref, ok := stage.Docxs_reference[docx]; !ok {
			docxs_newInstances = append(docxs_newInstances, docx)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Docx "+docx.Name,
				)
			}
		} else {
			diffs := docx.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Docx \""+docx.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for docx := range stage.Docxs_reference {
		if _, ok := stage.Docxs[docx]; !ok {
			docxs_deletedInstances = append(docxs_deletedInstances, docx)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Docx "+docx.Name,
				)
			}
		}
	}

	lenNewInstances += len(docxs_newInstances)
	lenDeletedInstances += len(docxs_deletedInstances)
	var files_newInstances []*File
	var files_deletedInstances []*File

	// parse all staged instances and check if they have a reference
	for file := range stage.Files {
		if ref, ok := stage.Files_reference[file]; !ok {
			files_newInstances = append(files_newInstances, file)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of File "+file.Name,
				)
			}
		} else {
			diffs := file.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of File \""+file.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for file := range stage.Files_reference {
		if _, ok := stage.Files[file]; !ok {
			files_deletedInstances = append(files_deletedInstances, file)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of File "+file.Name,
				)
			}
		}
	}

	lenNewInstances += len(files_newInstances)
	lenDeletedInstances += len(files_deletedInstances)
	var nodes_newInstances []*Node
	var nodes_deletedInstances []*Node

	// parse all staged instances and check if they have a reference
	for node := range stage.Nodes {
		if ref, ok := stage.Nodes_reference[node]; !ok {
			nodes_newInstances = append(nodes_newInstances, node)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Node "+node.Name,
				)
			}
		} else {
			diffs := node.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Node \""+node.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for node := range stage.Nodes_reference {
		if _, ok := stage.Nodes[node]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, node)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Node "+node.Name,
				)
			}
		}
	}

	lenNewInstances += len(nodes_newInstances)
	lenDeletedInstances += len(nodes_deletedInstances)
	var paragraphs_newInstances []*Paragraph
	var paragraphs_deletedInstances []*Paragraph

	// parse all staged instances and check if they have a reference
	for paragraph := range stage.Paragraphs {
		if ref, ok := stage.Paragraphs_reference[paragraph]; !ok {
			paragraphs_newInstances = append(paragraphs_newInstances, paragraph)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Paragraph "+paragraph.Name,
				)
			}
		} else {
			diffs := paragraph.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Paragraph \""+paragraph.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for paragraph := range stage.Paragraphs_reference {
		if _, ok := stage.Paragraphs[paragraph]; !ok {
			paragraphs_deletedInstances = append(paragraphs_deletedInstances, paragraph)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Paragraph "+paragraph.Name,
				)
			}
		}
	}

	lenNewInstances += len(paragraphs_newInstances)
	lenDeletedInstances += len(paragraphs_deletedInstances)
	var paragraphpropertiess_newInstances []*ParagraphProperties
	var paragraphpropertiess_deletedInstances []*ParagraphProperties

	// parse all staged instances and check if they have a reference
	for paragraphproperties := range stage.ParagraphPropertiess {
		if ref, ok := stage.ParagraphPropertiess_reference[paragraphproperties]; !ok {
			paragraphpropertiess_newInstances = append(paragraphpropertiess_newInstances, paragraphproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of ParagraphProperties "+paragraphproperties.Name,
				)
			}
		} else {
			diffs := paragraphproperties.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of ParagraphProperties \""+paragraphproperties.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for paragraphproperties := range stage.ParagraphPropertiess_reference {
		if _, ok := stage.ParagraphPropertiess[paragraphproperties]; !ok {
			paragraphpropertiess_deletedInstances = append(paragraphpropertiess_deletedInstances, paragraphproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of ParagraphProperties "+paragraphproperties.Name,
				)
			}
		}
	}

	lenNewInstances += len(paragraphpropertiess_newInstances)
	lenDeletedInstances += len(paragraphpropertiess_deletedInstances)
	var paragraphstyles_newInstances []*ParagraphStyle
	var paragraphstyles_deletedInstances []*ParagraphStyle

	// parse all staged instances and check if they have a reference
	for paragraphstyle := range stage.ParagraphStyles {
		if ref, ok := stage.ParagraphStyles_reference[paragraphstyle]; !ok {
			paragraphstyles_newInstances = append(paragraphstyles_newInstances, paragraphstyle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of ParagraphStyle "+paragraphstyle.Name,
				)
			}
		} else {
			diffs := paragraphstyle.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of ParagraphStyle \""+paragraphstyle.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for paragraphstyle := range stage.ParagraphStyles_reference {
		if _, ok := stage.ParagraphStyles[paragraphstyle]; !ok {
			paragraphstyles_deletedInstances = append(paragraphstyles_deletedInstances, paragraphstyle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of ParagraphStyle "+paragraphstyle.Name,
				)
			}
		}
	}

	lenNewInstances += len(paragraphstyles_newInstances)
	lenDeletedInstances += len(paragraphstyles_deletedInstances)
	var runes_newInstances []*Rune
	var runes_deletedInstances []*Rune

	// parse all staged instances and check if they have a reference
	for rune := range stage.Runes {
		if ref, ok := stage.Runes_reference[rune]; !ok {
			runes_newInstances = append(runes_newInstances, rune)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Rune "+rune.Name,
				)
			}
		} else {
			diffs := rune.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Rune \""+rune.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for rune := range stage.Runes_reference {
		if _, ok := stage.Runes[rune]; !ok {
			runes_deletedInstances = append(runes_deletedInstances, rune)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Rune "+rune.Name,
				)
			}
		}
	}

	lenNewInstances += len(runes_newInstances)
	lenDeletedInstances += len(runes_deletedInstances)
	var runepropertiess_newInstances []*RuneProperties
	var runepropertiess_deletedInstances []*RuneProperties

	// parse all staged instances and check if they have a reference
	for runeproperties := range stage.RunePropertiess {
		if ref, ok := stage.RunePropertiess_reference[runeproperties]; !ok {
			runepropertiess_newInstances = append(runepropertiess_newInstances, runeproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of RuneProperties "+runeproperties.Name,
				)
			}
		} else {
			diffs := runeproperties.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of RuneProperties \""+runeproperties.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for runeproperties := range stage.RunePropertiess_reference {
		if _, ok := stage.RunePropertiess[runeproperties]; !ok {
			runepropertiess_deletedInstances = append(runepropertiess_deletedInstances, runeproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of RuneProperties "+runeproperties.Name,
				)
			}
		}
	}

	lenNewInstances += len(runepropertiess_newInstances)
	lenDeletedInstances += len(runepropertiess_deletedInstances)
	var tables_newInstances []*Table
	var tables_deletedInstances []*Table

	// parse all staged instances and check if they have a reference
	for table := range stage.Tables {
		if ref, ok := stage.Tables_reference[table]; !ok {
			tables_newInstances = append(tables_newInstances, table)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Table "+table.Name,
				)
			}
		} else {
			diffs := table.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Table \""+table.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for table := range stage.Tables_reference {
		if _, ok := stage.Tables[table]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, table)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Table "+table.Name,
				)
			}
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)
	var tablecolumns_newInstances []*TableColumn
	var tablecolumns_deletedInstances []*TableColumn

	// parse all staged instances and check if they have a reference
	for tablecolumn := range stage.TableColumns {
		if ref, ok := stage.TableColumns_reference[tablecolumn]; !ok {
			tablecolumns_newInstances = append(tablecolumns_newInstances, tablecolumn)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of TableColumn "+tablecolumn.Name,
				)
			}
		} else {
			diffs := tablecolumn.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of TableColumn \""+tablecolumn.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tablecolumn := range stage.TableColumns_reference {
		if _, ok := stage.TableColumns[tablecolumn]; !ok {
			tablecolumns_deletedInstances = append(tablecolumns_deletedInstances, tablecolumn)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of TableColumn "+tablecolumn.Name,
				)
			}
		}
	}

	lenNewInstances += len(tablecolumns_newInstances)
	lenDeletedInstances += len(tablecolumns_deletedInstances)
	var tablepropertiess_newInstances []*TableProperties
	var tablepropertiess_deletedInstances []*TableProperties

	// parse all staged instances and check if they have a reference
	for tableproperties := range stage.TablePropertiess {
		if ref, ok := stage.TablePropertiess_reference[tableproperties]; !ok {
			tablepropertiess_newInstances = append(tablepropertiess_newInstances, tableproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of TableProperties "+tableproperties.Name,
				)
			}
		} else {
			diffs := tableproperties.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of TableProperties \""+tableproperties.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tableproperties := range stage.TablePropertiess_reference {
		if _, ok := stage.TablePropertiess[tableproperties]; !ok {
			tablepropertiess_deletedInstances = append(tablepropertiess_deletedInstances, tableproperties)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of TableProperties "+tableproperties.Name,
				)
			}
		}
	}

	lenNewInstances += len(tablepropertiess_newInstances)
	lenDeletedInstances += len(tablepropertiess_deletedInstances)
	var tablerows_newInstances []*TableRow
	var tablerows_deletedInstances []*TableRow

	// parse all staged instances and check if they have a reference
	for tablerow := range stage.TableRows {
		if ref, ok := stage.TableRows_reference[tablerow]; !ok {
			tablerows_newInstances = append(tablerows_newInstances, tablerow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of TableRow "+tablerow.Name,
				)
			}
		} else {
			diffs := tablerow.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of TableRow \""+tablerow.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tablerow := range stage.TableRows_reference {
		if _, ok := stage.TableRows[tablerow]; !ok {
			tablerows_deletedInstances = append(tablerows_deletedInstances, tablerow)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of TableRow "+tablerow.Name,
				)
			}
		}
	}

	lenNewInstances += len(tablerows_newInstances)
	lenDeletedInstances += len(tablerows_deletedInstances)
	var tablestyles_newInstances []*TableStyle
	var tablestyles_deletedInstances []*TableStyle

	// parse all staged instances and check if they have a reference
	for tablestyle := range stage.TableStyles {
		if ref, ok := stage.TableStyles_reference[tablestyle]; !ok {
			tablestyles_newInstances = append(tablestyles_newInstances, tablestyle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of TableStyle "+tablestyle.Name,
				)
			}
		} else {
			diffs := tablestyle.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of TableStyle \""+tablestyle.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tablestyle := range stage.TableStyles_reference {
		if _, ok := stage.TableStyles[tablestyle]; !ok {
			tablestyles_deletedInstances = append(tablestyles_deletedInstances, tablestyle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of TableStyle "+tablestyle.Name,
				)
			}
		}
	}

	lenNewInstances += len(tablestyles_newInstances)
	lenDeletedInstances += len(tablestyles_deletedInstances)
	var texts_newInstances []*Text
	var texts_deletedInstances []*Text

	// parse all staged instances and check if they have a reference
	for text := range stage.Texts {
		if ref, ok := stage.Texts_reference[text]; !ok {
			texts_newInstances = append(texts_newInstances, text)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Text "+text.Name,
				)
			}
		} else {
			diffs := text.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Text \""+text.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for text := range stage.Texts_reference {
		if _, ok := stage.Texts[text]; !ok {
			texts_deletedInstances = append(texts_deletedInstances, text)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Text "+text.Name,
				)
			}
		}
	}

	lenNewInstances += len(texts_newInstances)
	lenDeletedInstances += len(texts_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Bodys_reference = make(map[*Body]*Body)
	for instance := range stage.Bodys {
		stage.Bodys_reference[instance] = instance.GongCopy().(*Body)
	}

	stage.Documents_reference = make(map[*Document]*Document)
	for instance := range stage.Documents {
		stage.Documents_reference[instance] = instance.GongCopy().(*Document)
	}

	stage.Docxs_reference = make(map[*Docx]*Docx)
	for instance := range stage.Docxs {
		stage.Docxs_reference[instance] = instance.GongCopy().(*Docx)
	}

	stage.Files_reference = make(map[*File]*File)
	for instance := range stage.Files {
		stage.Files_reference[instance] = instance.GongCopy().(*File)
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	for instance := range stage.Nodes {
		stage.Nodes_reference[instance] = instance.GongCopy().(*Node)
	}

	stage.Paragraphs_reference = make(map[*Paragraph]*Paragraph)
	for instance := range stage.Paragraphs {
		stage.Paragraphs_reference[instance] = instance.GongCopy().(*Paragraph)
	}

	stage.ParagraphPropertiess_reference = make(map[*ParagraphProperties]*ParagraphProperties)
	for instance := range stage.ParagraphPropertiess {
		stage.ParagraphPropertiess_reference[instance] = instance.GongCopy().(*ParagraphProperties)
	}

	stage.ParagraphStyles_reference = make(map[*ParagraphStyle]*ParagraphStyle)
	for instance := range stage.ParagraphStyles {
		stage.ParagraphStyles_reference[instance] = instance.GongCopy().(*ParagraphStyle)
	}

	stage.Runes_reference = make(map[*Rune]*Rune)
	for instance := range stage.Runes {
		stage.Runes_reference[instance] = instance.GongCopy().(*Rune)
	}

	stage.RunePropertiess_reference = make(map[*RuneProperties]*RuneProperties)
	for instance := range stage.RunePropertiess {
		stage.RunePropertiess_reference[instance] = instance.GongCopy().(*RuneProperties)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
	}

	stage.TableColumns_reference = make(map[*TableColumn]*TableColumn)
	for instance := range stage.TableColumns {
		stage.TableColumns_reference[instance] = instance.GongCopy().(*TableColumn)
	}

	stage.TablePropertiess_reference = make(map[*TableProperties]*TableProperties)
	for instance := range stage.TablePropertiess {
		stage.TablePropertiess_reference[instance] = instance.GongCopy().(*TableProperties)
	}

	stage.TableRows_reference = make(map[*TableRow]*TableRow)
	for instance := range stage.TableRows {
		stage.TableRows_reference[instance] = instance.GongCopy().(*TableRow)
	}

	stage.TableStyles_reference = make(map[*TableStyle]*TableStyle)
	for instance := range stage.TableStyles {
		stage.TableStyles_reference[instance] = instance.GongCopy().(*TableStyle)
	}

	stage.Texts_reference = make(map[*Text]*Text)
	for instance := range stage.Texts {
		stage.Texts_reference[instance] = instance.GongCopy().(*Text)
	}

}
