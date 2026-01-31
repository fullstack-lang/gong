// generated code - do not edit
package models

import (
	"fmt"
	"sort"
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

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var bodys_newInstances []*Body
	var bodys_deletedInstances []*Body

	// parse all staged instances and check if they have a reference
	for body := range stage.Bodys {
		if ref, ok := stage.Bodys_reference[body]; !ok {
			bodys_newInstances = append(bodys_newInstances, body)
			newInstancesSlice = append(newInstancesSlice, body.GongMarshallIdentifier(stage))
			if stage.Bodys_referenceOrder == nil {
				stage.Bodys_referenceOrder = make(map[*Body]uint)
			}
			stage.Bodys_referenceOrder[body] = stage.BodyMap_Staged_Order[body]
			newInstancesReverseSlice = append(newInstancesReverseSlice, body.GongMarshallUnstaging(stage))
			delete(stage.Bodys_referenceOrder, body)
			fieldInitializers, pointersInitializations := body.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BodyMap_Staged_Order[ref] = stage.BodyMap_Staged_Order[body]
			diffs := body.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, body)
			delete(stage.BodyMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", body.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Bodys_reference {
		if _, ok := stage.Bodys[ref]; !ok {
			bodys_deletedInstances = append(bodys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, document.GongMarshallIdentifier(stage))
			if stage.Documents_referenceOrder == nil {
				stage.Documents_referenceOrder = make(map[*Document]uint)
			}
			stage.Documents_referenceOrder[document] = stage.DocumentMap_Staged_Order[document]
			newInstancesReverseSlice = append(newInstancesReverseSlice, document.GongMarshallUnstaging(stage))
			delete(stage.Documents_referenceOrder, document)
			fieldInitializers, pointersInitializations := document.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DocumentMap_Staged_Order[ref] = stage.DocumentMap_Staged_Order[document]
			diffs := document.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, document)
			delete(stage.DocumentMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", document.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Documents_reference {
		if _, ok := stage.Documents[ref]; !ok {
			documents_deletedInstances = append(documents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, docx.GongMarshallIdentifier(stage))
			if stage.Docxs_referenceOrder == nil {
				stage.Docxs_referenceOrder = make(map[*Docx]uint)
			}
			stage.Docxs_referenceOrder[docx] = stage.DocxMap_Staged_Order[docx]
			newInstancesReverseSlice = append(newInstancesReverseSlice, docx.GongMarshallUnstaging(stage))
			delete(stage.Docxs_referenceOrder, docx)
			fieldInitializers, pointersInitializations := docx.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DocxMap_Staged_Order[ref] = stage.DocxMap_Staged_Order[docx]
			diffs := docx.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, docx)
			delete(stage.DocxMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", docx.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Docxs_reference {
		if _, ok := stage.Docxs[ref]; !ok {
			docxs_deletedInstances = append(docxs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, file.GongMarshallIdentifier(stage))
			if stage.Files_referenceOrder == nil {
				stage.Files_referenceOrder = make(map[*File]uint)
			}
			stage.Files_referenceOrder[file] = stage.FileMap_Staged_Order[file]
			newInstancesReverseSlice = append(newInstancesReverseSlice, file.GongMarshallUnstaging(stage))
			delete(stage.Files_referenceOrder, file)
			fieldInitializers, pointersInitializations := file.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FileMap_Staged_Order[ref] = stage.FileMap_Staged_Order[file]
			diffs := file.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, file)
			delete(stage.FileMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", file.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Files_reference {
		if _, ok := stage.Files[ref]; !ok {
			files_deletedInstances = append(files_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, node.GongMarshallIdentifier(stage))
			if stage.Nodes_referenceOrder == nil {
				stage.Nodes_referenceOrder = make(map[*Node]uint)
			}
			stage.Nodes_referenceOrder[node] = stage.NodeMap_Staged_Order[node]
			newInstancesReverseSlice = append(newInstancesReverseSlice, node.GongMarshallUnstaging(stage))
			delete(stage.Nodes_referenceOrder, node)
			fieldInitializers, pointersInitializations := node.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NodeMap_Staged_Order[ref] = stage.NodeMap_Staged_Order[node]
			diffs := node.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, node)
			delete(stage.NodeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", node.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Nodes_reference {
		if _, ok := stage.Nodes[ref]; !ok {
			nodes_deletedInstances = append(nodes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, paragraph.GongMarshallIdentifier(stage))
			if stage.Paragraphs_referenceOrder == nil {
				stage.Paragraphs_referenceOrder = make(map[*Paragraph]uint)
			}
			stage.Paragraphs_referenceOrder[paragraph] = stage.ParagraphMap_Staged_Order[paragraph]
			newInstancesReverseSlice = append(newInstancesReverseSlice, paragraph.GongMarshallUnstaging(stage))
			delete(stage.Paragraphs_referenceOrder, paragraph)
			fieldInitializers, pointersInitializations := paragraph.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParagraphMap_Staged_Order[ref] = stage.ParagraphMap_Staged_Order[paragraph]
			diffs := paragraph.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, paragraph)
			delete(stage.ParagraphMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", paragraph.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Paragraphs_reference {
		if _, ok := stage.Paragraphs[ref]; !ok {
			paragraphs_deletedInstances = append(paragraphs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, paragraphproperties.GongMarshallIdentifier(stage))
			if stage.ParagraphPropertiess_referenceOrder == nil {
				stage.ParagraphPropertiess_referenceOrder = make(map[*ParagraphProperties]uint)
			}
			stage.ParagraphPropertiess_referenceOrder[paragraphproperties] = stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties]
			newInstancesReverseSlice = append(newInstancesReverseSlice, paragraphproperties.GongMarshallUnstaging(stage))
			delete(stage.ParagraphPropertiess_referenceOrder, paragraphproperties)
			fieldInitializers, pointersInitializations := paragraphproperties.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParagraphPropertiesMap_Staged_Order[ref] = stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties]
			diffs := paragraphproperties.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, paragraphproperties)
			delete(stage.ParagraphPropertiesMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", paragraphproperties.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ParagraphPropertiess_reference {
		if _, ok := stage.ParagraphPropertiess[ref]; !ok {
			paragraphpropertiess_deletedInstances = append(paragraphpropertiess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, paragraphstyle.GongMarshallIdentifier(stage))
			if stage.ParagraphStyles_referenceOrder == nil {
				stage.ParagraphStyles_referenceOrder = make(map[*ParagraphStyle]uint)
			}
			stage.ParagraphStyles_referenceOrder[paragraphstyle] = stage.ParagraphStyleMap_Staged_Order[paragraphstyle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, paragraphstyle.GongMarshallUnstaging(stage))
			delete(stage.ParagraphStyles_referenceOrder, paragraphstyle)
			fieldInitializers, pointersInitializations := paragraphstyle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParagraphStyleMap_Staged_Order[ref] = stage.ParagraphStyleMap_Staged_Order[paragraphstyle]
			diffs := paragraphstyle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, paragraphstyle)
			delete(stage.ParagraphStyleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", paragraphstyle.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ParagraphStyles_reference {
		if _, ok := stage.ParagraphStyles[ref]; !ok {
			paragraphstyles_deletedInstances = append(paragraphstyles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, rune.GongMarshallIdentifier(stage))
			if stage.Runes_referenceOrder == nil {
				stage.Runes_referenceOrder = make(map[*Rune]uint)
			}
			stage.Runes_referenceOrder[rune] = stage.RuneMap_Staged_Order[rune]
			newInstancesReverseSlice = append(newInstancesReverseSlice, rune.GongMarshallUnstaging(stage))
			delete(stage.Runes_referenceOrder, rune)
			fieldInitializers, pointersInitializations := rune.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RuneMap_Staged_Order[ref] = stage.RuneMap_Staged_Order[rune]
			diffs := rune.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, rune)
			delete(stage.RuneMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", rune.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Runes_reference {
		if _, ok := stage.Runes[ref]; !ok {
			runes_deletedInstances = append(runes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, runeproperties.GongMarshallIdentifier(stage))
			if stage.RunePropertiess_referenceOrder == nil {
				stage.RunePropertiess_referenceOrder = make(map[*RuneProperties]uint)
			}
			stage.RunePropertiess_referenceOrder[runeproperties] = stage.RunePropertiesMap_Staged_Order[runeproperties]
			newInstancesReverseSlice = append(newInstancesReverseSlice, runeproperties.GongMarshallUnstaging(stage))
			delete(stage.RunePropertiess_referenceOrder, runeproperties)
			fieldInitializers, pointersInitializations := runeproperties.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RunePropertiesMap_Staged_Order[ref] = stage.RunePropertiesMap_Staged_Order[runeproperties]
			diffs := runeproperties.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, runeproperties)
			delete(stage.RunePropertiesMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", runeproperties.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RunePropertiess_reference {
		if _, ok := stage.RunePropertiess[ref]; !ok {
			runepropertiess_deletedInstances = append(runepropertiess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, table.GongMarshallIdentifier(stage))
			if stage.Tables_referenceOrder == nil {
				stage.Tables_referenceOrder = make(map[*Table]uint)
			}
			stage.Tables_referenceOrder[table] = stage.TableMap_Staged_Order[table]
			newInstancesReverseSlice = append(newInstancesReverseSlice, table.GongMarshallUnstaging(stage))
			delete(stage.Tables_referenceOrder, table)
			fieldInitializers, pointersInitializations := table.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TableMap_Staged_Order[ref] = stage.TableMap_Staged_Order[table]
			diffs := table.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, table)
			delete(stage.TableMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", table.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Tables_reference {
		if _, ok := stage.Tables[ref]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, tablecolumn.GongMarshallIdentifier(stage))
			if stage.TableColumns_referenceOrder == nil {
				stage.TableColumns_referenceOrder = make(map[*TableColumn]uint)
			}
			stage.TableColumns_referenceOrder[tablecolumn] = stage.TableColumnMap_Staged_Order[tablecolumn]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tablecolumn.GongMarshallUnstaging(stage))
			delete(stage.TableColumns_referenceOrder, tablecolumn)
			fieldInitializers, pointersInitializations := tablecolumn.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TableColumnMap_Staged_Order[ref] = stage.TableColumnMap_Staged_Order[tablecolumn]
			diffs := tablecolumn.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tablecolumn)
			delete(stage.TableColumnMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tablecolumn.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.TableColumns_reference {
		if _, ok := stage.TableColumns[ref]; !ok {
			tablecolumns_deletedInstances = append(tablecolumns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, tableproperties.GongMarshallIdentifier(stage))
			if stage.TablePropertiess_referenceOrder == nil {
				stage.TablePropertiess_referenceOrder = make(map[*TableProperties]uint)
			}
			stage.TablePropertiess_referenceOrder[tableproperties] = stage.TablePropertiesMap_Staged_Order[tableproperties]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tableproperties.GongMarshallUnstaging(stage))
			delete(stage.TablePropertiess_referenceOrder, tableproperties)
			fieldInitializers, pointersInitializations := tableproperties.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TablePropertiesMap_Staged_Order[ref] = stage.TablePropertiesMap_Staged_Order[tableproperties]
			diffs := tableproperties.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tableproperties)
			delete(stage.TablePropertiesMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tableproperties.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.TablePropertiess_reference {
		if _, ok := stage.TablePropertiess[ref]; !ok {
			tablepropertiess_deletedInstances = append(tablepropertiess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, tablerow.GongMarshallIdentifier(stage))
			if stage.TableRows_referenceOrder == nil {
				stage.TableRows_referenceOrder = make(map[*TableRow]uint)
			}
			stage.TableRows_referenceOrder[tablerow] = stage.TableRowMap_Staged_Order[tablerow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tablerow.GongMarshallUnstaging(stage))
			delete(stage.TableRows_referenceOrder, tablerow)
			fieldInitializers, pointersInitializations := tablerow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TableRowMap_Staged_Order[ref] = stage.TableRowMap_Staged_Order[tablerow]
			diffs := tablerow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tablerow)
			delete(stage.TableRowMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tablerow.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.TableRows_reference {
		if _, ok := stage.TableRows[ref]; !ok {
			tablerows_deletedInstances = append(tablerows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, tablestyle.GongMarshallIdentifier(stage))
			if stage.TableStyles_referenceOrder == nil {
				stage.TableStyles_referenceOrder = make(map[*TableStyle]uint)
			}
			stage.TableStyles_referenceOrder[tablestyle] = stage.TableStyleMap_Staged_Order[tablestyle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tablestyle.GongMarshallUnstaging(stage))
			delete(stage.TableStyles_referenceOrder, tablestyle)
			fieldInitializers, pointersInitializations := tablestyle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TableStyleMap_Staged_Order[ref] = stage.TableStyleMap_Staged_Order[tablestyle]
			diffs := tablestyle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tablestyle)
			delete(stage.TableStyleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tablestyle.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.TableStyles_reference {
		if _, ok := stage.TableStyles[ref]; !ok {
			tablestyles_deletedInstances = append(tablestyles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, text.GongMarshallIdentifier(stage))
			if stage.Texts_referenceOrder == nil {
				stage.Texts_referenceOrder = make(map[*Text]uint)
			}
			stage.Texts_referenceOrder[text] = stage.TextMap_Staged_Order[text]
			newInstancesReverseSlice = append(newInstancesReverseSlice, text.GongMarshallUnstaging(stage))
			delete(stage.Texts_referenceOrder, text)
			fieldInitializers, pointersInitializations := text.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TextMap_Staged_Order[ref] = stage.TextMap_Staged_Order[text]
			diffs := text.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, text)
			delete(stage.TextMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", text.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Texts_reference {
		if _, ok := stage.Texts[ref]; !ok {
			texts_deletedInstances = append(texts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(texts_newInstances)
	lenDeletedInstances += len(texts_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {

	// insertion point per named struct
	stage.Bodys_reference = make(map[*Body]*Body)
	stage.Bodys_referenceOrder = make(map[*Body]uint) // diff Unstage needs the reference order
	for instance := range stage.Bodys {
		stage.Bodys_reference[instance] = instance.GongCopy().(*Body)
		stage.Bodys_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Documents_reference = make(map[*Document]*Document)
	stage.Documents_referenceOrder = make(map[*Document]uint) // diff Unstage needs the reference order
	for instance := range stage.Documents {
		stage.Documents_reference[instance] = instance.GongCopy().(*Document)
		stage.Documents_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Docxs_reference = make(map[*Docx]*Docx)
	stage.Docxs_referenceOrder = make(map[*Docx]uint) // diff Unstage needs the reference order
	for instance := range stage.Docxs {
		stage.Docxs_reference[instance] = instance.GongCopy().(*Docx)
		stage.Docxs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Files_reference = make(map[*File]*File)
	stage.Files_referenceOrder = make(map[*File]uint) // diff Unstage needs the reference order
	for instance := range stage.Files {
		stage.Files_reference[instance] = instance.GongCopy().(*File)
		stage.Files_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	stage.Nodes_referenceOrder = make(map[*Node]uint) // diff Unstage needs the reference order
	for instance := range stage.Nodes {
		stage.Nodes_reference[instance] = instance.GongCopy().(*Node)
		stage.Nodes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Paragraphs_reference = make(map[*Paragraph]*Paragraph)
	stage.Paragraphs_referenceOrder = make(map[*Paragraph]uint) // diff Unstage needs the reference order
	for instance := range stage.Paragraphs {
		stage.Paragraphs_reference[instance] = instance.GongCopy().(*Paragraph)
		stage.Paragraphs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ParagraphPropertiess_reference = make(map[*ParagraphProperties]*ParagraphProperties)
	stage.ParagraphPropertiess_referenceOrder = make(map[*ParagraphProperties]uint) // diff Unstage needs the reference order
	for instance := range stage.ParagraphPropertiess {
		stage.ParagraphPropertiess_reference[instance] = instance.GongCopy().(*ParagraphProperties)
		stage.ParagraphPropertiess_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ParagraphStyles_reference = make(map[*ParagraphStyle]*ParagraphStyle)
	stage.ParagraphStyles_referenceOrder = make(map[*ParagraphStyle]uint) // diff Unstage needs the reference order
	for instance := range stage.ParagraphStyles {
		stage.ParagraphStyles_reference[instance] = instance.GongCopy().(*ParagraphStyle)
		stage.ParagraphStyles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Runes_reference = make(map[*Rune]*Rune)
	stage.Runes_referenceOrder = make(map[*Rune]uint) // diff Unstage needs the reference order
	for instance := range stage.Runes {
		stage.Runes_reference[instance] = instance.GongCopy().(*Rune)
		stage.Runes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RunePropertiess_reference = make(map[*RuneProperties]*RuneProperties)
	stage.RunePropertiess_referenceOrder = make(map[*RuneProperties]uint) // diff Unstage needs the reference order
	for instance := range stage.RunePropertiess {
		stage.RunePropertiess_reference[instance] = instance.GongCopy().(*RuneProperties)
		stage.RunePropertiess_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
		stage.Tables_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TableColumns_reference = make(map[*TableColumn]*TableColumn)
	stage.TableColumns_referenceOrder = make(map[*TableColumn]uint) // diff Unstage needs the reference order
	for instance := range stage.TableColumns {
		stage.TableColumns_reference[instance] = instance.GongCopy().(*TableColumn)
		stage.TableColumns_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TablePropertiess_reference = make(map[*TableProperties]*TableProperties)
	stage.TablePropertiess_referenceOrder = make(map[*TableProperties]uint) // diff Unstage needs the reference order
	for instance := range stage.TablePropertiess {
		stage.TablePropertiess_reference[instance] = instance.GongCopy().(*TableProperties)
		stage.TablePropertiess_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TableRows_reference = make(map[*TableRow]*TableRow)
	stage.TableRows_referenceOrder = make(map[*TableRow]uint) // diff Unstage needs the reference order
	for instance := range stage.TableRows {
		stage.TableRows_reference[instance] = instance.GongCopy().(*TableRow)
		stage.TableRows_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.TableStyles_reference = make(map[*TableStyle]*TableStyle)
	stage.TableStyles_referenceOrder = make(map[*TableStyle]uint) // diff Unstage needs the reference order
	for instance := range stage.TableStyles {
		stage.TableStyles_reference[instance] = instance.GongCopy().(*TableStyle)
		stage.TableStyles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Texts_reference = make(map[*Text]*Text)
	stage.Texts_referenceOrder = make(map[*Text]uint) // diff Unstage needs the reference order
	for instance := range stage.Texts {
		stage.Texts_reference[instance] = instance.GongCopy().(*Text)
		stage.Texts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (body *Body) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BodyMap_Staged_Order[body]; ok {
		return order
	}
	return stage.Bodys_referenceOrder[body]
}

func (body *Body) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Bodys_referenceOrder[body]
}

func (document *Document) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DocumentMap_Staged_Order[document]; ok {
		return order
	}
	return stage.Documents_referenceOrder[document]
}

func (document *Document) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Documents_referenceOrder[document]
}

func (docx *Docx) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DocxMap_Staged_Order[docx]; ok {
		return order
	}
	return stage.Docxs_referenceOrder[docx]
}

func (docx *Docx) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Docxs_referenceOrder[docx]
}

func (file *File) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FileMap_Staged_Order[file]; ok {
		return order
	}
	return stage.Files_referenceOrder[file]
}

func (file *File) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Files_referenceOrder[file]
}

func (node *Node) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NodeMap_Staged_Order[node]; ok {
		return order
	}
	return stage.Nodes_referenceOrder[node]
}

func (node *Node) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Nodes_referenceOrder[node]
}

func (paragraph *Paragraph) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParagraphMap_Staged_Order[paragraph]; ok {
		return order
	}
	return stage.Paragraphs_referenceOrder[paragraph]
}

func (paragraph *Paragraph) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Paragraphs_referenceOrder[paragraph]
}

func (paragraphproperties *ParagraphProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParagraphPropertiesMap_Staged_Order[paragraphproperties]; ok {
		return order
	}
	return stage.ParagraphPropertiess_referenceOrder[paragraphproperties]
}

func (paragraphproperties *ParagraphProperties) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ParagraphPropertiess_referenceOrder[paragraphproperties]
}

func (paragraphstyle *ParagraphStyle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParagraphStyleMap_Staged_Order[paragraphstyle]; ok {
		return order
	}
	return stage.ParagraphStyles_referenceOrder[paragraphstyle]
}

func (paragraphstyle *ParagraphStyle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ParagraphStyles_referenceOrder[paragraphstyle]
}

func (rune *Rune) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RuneMap_Staged_Order[rune]; ok {
		return order
	}
	return stage.Runes_referenceOrder[rune]
}

func (rune *Rune) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Runes_referenceOrder[rune]
}

func (runeproperties *RuneProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RunePropertiesMap_Staged_Order[runeproperties]; ok {
		return order
	}
	return stage.RunePropertiess_referenceOrder[runeproperties]
}

func (runeproperties *RuneProperties) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RunePropertiess_referenceOrder[runeproperties]
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableMap_Staged_Order[table]; ok {
		return order
	}
	return stage.Tables_referenceOrder[table]
}

func (table *Table) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tables_referenceOrder[table]
}

func (tablecolumn *TableColumn) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableColumnMap_Staged_Order[tablecolumn]; ok {
		return order
	}
	return stage.TableColumns_referenceOrder[tablecolumn]
}

func (tablecolumn *TableColumn) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TableColumns_referenceOrder[tablecolumn]
}

func (tableproperties *TableProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TablePropertiesMap_Staged_Order[tableproperties]; ok {
		return order
	}
	return stage.TablePropertiess_referenceOrder[tableproperties]
}

func (tableproperties *TableProperties) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TablePropertiess_referenceOrder[tableproperties]
}

func (tablerow *TableRow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableRowMap_Staged_Order[tablerow]; ok {
		return order
	}
	return stage.TableRows_referenceOrder[tablerow]
}

func (tablerow *TableRow) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TableRows_referenceOrder[tablerow]
}

func (tablestyle *TableStyle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableStyleMap_Staged_Order[tablestyle]; ok {
		return order
	}
	return stage.TableStyles_referenceOrder[tablestyle]
}

func (tablestyle *TableStyle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.TableStyles_referenceOrder[tablestyle]
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TextMap_Staged_Order[text]; ok {
		return order
	}
	return stage.Texts_referenceOrder[text]
}

func (text *Text) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Texts_referenceOrder[text]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (body *Body) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", body.GongGetGongstructName(), body.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (body *Body) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", body.GongGetGongstructName(), body.GongGetReferenceOrder(stage))
}

func (document *Document) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (document *Document) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetReferenceOrder(stage))
}

func (docx *Docx) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", docx.GongGetGongstructName(), docx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (docx *Docx) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", docx.GongGetGongstructName(), docx.GongGetReferenceOrder(stage))
}

func (file *File) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", file.GongGetGongstructName(), file.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (file *File) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", file.GongGetGongstructName(), file.GongGetReferenceOrder(stage))
}

func (node *Node) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (node *Node) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetReferenceOrder(stage))
}

func (paragraph *Paragraph) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraph.GongGetGongstructName(), paragraph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraph *Paragraph) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraph.GongGetGongstructName(), paragraph.GongGetReferenceOrder(stage))
}

func (paragraphproperties *ParagraphProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphproperties.GongGetGongstructName(), paragraphproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphproperties *ParagraphProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphproperties.GongGetGongstructName(), paragraphproperties.GongGetReferenceOrder(stage))
}

func (paragraphstyle *ParagraphStyle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphstyle.GongGetGongstructName(), paragraphstyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphstyle *ParagraphStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphstyle.GongGetGongstructName(), paragraphstyle.GongGetReferenceOrder(stage))
}

func (rune *Rune) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rune.GongGetGongstructName(), rune.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rune *Rune) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rune.GongGetGongstructName(), rune.GongGetReferenceOrder(stage))
}

func (runeproperties *RuneProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", runeproperties.GongGetGongstructName(), runeproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runeproperties *RuneProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", runeproperties.GongGetGongstructName(), runeproperties.GongGetReferenceOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetReferenceOrder(stage))
}

func (tablecolumn *TableColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablecolumn.GongGetGongstructName(), tablecolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablecolumn *TableColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablecolumn.GongGetGongstructName(), tablecolumn.GongGetReferenceOrder(stage))
}

func (tableproperties *TableProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tableproperties.GongGetGongstructName(), tableproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tableproperties *TableProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tableproperties.GongGetGongstructName(), tableproperties.GongGetReferenceOrder(stage))
}

func (tablerow *TableRow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablerow.GongGetGongstructName(), tablerow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablerow *TableRow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablerow.GongGetGongstructName(), tablerow.GongGetReferenceOrder(stage))
}

func (tablestyle *TableStyle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablestyle.GongGetGongstructName(), tablestyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablestyle *TableStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablestyle.GongGetGongstructName(), tablestyle.GongGetReferenceOrder(stage))
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (body *Body) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", body.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Body")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", body.Name)
	return
}
func (document *Document) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", document.Name)
	return
}
func (docx *Docx) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", docx.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Docx")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", docx.Name)
	return
}
func (file *File) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", file.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "File")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", file.Name)
	return
}
func (node *Node) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Node")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", node.Name)
	return
}
func (paragraph *Paragraph) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Paragraph")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraph.Name)
	return
}
func (paragraphproperties *ParagraphProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphproperties.Name)
	return
}
func (paragraphstyle *ParagraphStyle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphStyle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", paragraphstyle.Name)
	return
}
func (rune *Rune) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rune.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rune")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", rune.Name)
	return
}
func (runeproperties *RuneProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RuneProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", runeproperties.Name)
	return
}
func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", table.Name)
	return
}
func (tablecolumn *TableColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablecolumn.Name)
	return
}
func (tableproperties *TableProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tableproperties.Name)
	return
}
func (tablerow *TableRow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableRow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablerow.Name)
	return
}
func (tablestyle *TableStyle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableStyle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tablestyle.Name)
	return
}
func (text *Text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", text.Name)
	return
}

// insertion point for unstaging
func (body *Body) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", body.GongGetReferenceIdentifier(stage))
	return
}
func (document *Document) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetReferenceIdentifier(stage))
	return
}
func (docx *Docx) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", docx.GongGetReferenceIdentifier(stage))
	return
}
func (file *File) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", file.GongGetReferenceIdentifier(stage))
	return
}
func (node *Node) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetReferenceIdentifier(stage))
	return
}
func (paragraph *Paragraph) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraph.GongGetReferenceIdentifier(stage))
	return
}
func (paragraphproperties *ParagraphProperties) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphproperties.GongGetReferenceIdentifier(stage))
	return
}
func (paragraphstyle *ParagraphStyle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphstyle.GongGetReferenceIdentifier(stage))
	return
}
func (rune *Rune) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rune.GongGetReferenceIdentifier(stage))
	return
}
func (runeproperties *RuneProperties) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", runeproperties.GongGetReferenceIdentifier(stage))
	return
}
func (table *Table) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetReferenceIdentifier(stage))
	return
}
func (tablecolumn *TableColumn) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablecolumn.GongGetReferenceIdentifier(stage))
	return
}
func (tableproperties *TableProperties) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tableproperties.GongGetReferenceIdentifier(stage))
	return
}
func (tablerow *TableRow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablerow.GongGetReferenceIdentifier(stage))
	return
}
func (tablestyle *TableStyle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablestyle.GongGetReferenceIdentifier(stage))
	return
}
func (text *Text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetReferenceIdentifier(stage))
	return
}
