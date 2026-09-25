// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

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

	// Compute reverse map for named struct Docx
	// insertion point per field
	stage.Docx_Files_reverseMap = make(map[*File]*Docx)
	for docx := range stage.Docxs {
		_ = docx
		for _, _file := range docx.Files {
			stage.Docx_Files_reverseMap[_file] = docx
		}
	}

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

	// Compute reverse map for named struct TableRow
	// insertion point per field
	stage.TableRow_TableColumns_reverseMap = make(map[*TableColumn]*TableRow)
	for tablerow := range stage.TableRows {
		_ = tablerow
		for _, _tablecolumn := range tablerow.TableColumns {
			stage.TableRow_TableColumns_reverseMap[_tablecolumn] = tablerow
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Bodys)

	res = __gong__appendInstances(res, stage.Documents)

	res = __gong__appendInstances(res, stage.Docxs)

	res = __gong__appendInstances(res, stage.Files)

	res = __gong__appendInstances(res, stage.Nodes)

	res = __gong__appendInstances(res, stage.Paragraphs)

	res = __gong__appendInstances(res, stage.ParagraphPropertiess)

	res = __gong__appendInstances(res, stage.ParagraphStyles)

	res = __gong__appendInstances(res, stage.Runes)

	res = __gong__appendInstances(res, stage.RunePropertiess)

	res = __gong__appendInstances(res, stage.Tables)

	res = __gong__appendInstances(res, stage.TableColumns)

	res = __gong__appendInstances(res, stage.TablePropertiess)

	res = __gong__appendInstances(res, stage.TableRows)

	res = __gong__appendInstances(res, stage.TableStyles)

	res = __gong__appendInstances(res, stage.Texts)

	return
}

// insertion point per named struct
func (body *Body) GongCopy() GongstructIF {
	newInstance := new(Body)
	body.GongCopyBasicFields(newInstance)
	return newInstance
}

func (document *Document) GongCopy() GongstructIF {
	newInstance := new(Document)
	document.GongCopyBasicFields(newInstance)
	return newInstance
}

func (docx *Docx) GongCopy() GongstructIF {
	newInstance := new(Docx)
	docx.GongCopyBasicFields(newInstance)
	return newInstance
}

func (file *File) GongCopy() GongstructIF {
	newInstance := new(File)
	file.GongCopyBasicFields(newInstance)
	return newInstance
}

func (node *Node) GongCopy() GongstructIF {
	newInstance := new(Node)
	node.GongCopyBasicFields(newInstance)
	return newInstance
}

func (paragraph *Paragraph) GongCopy() GongstructIF {
	newInstance := new(Paragraph)
	paragraph.GongCopyBasicFields(newInstance)
	return newInstance
}

func (paragraphproperties *ParagraphProperties) GongCopy() GongstructIF {
	newInstance := new(ParagraphProperties)
	paragraphproperties.GongCopyBasicFields(newInstance)
	return newInstance
}

func (paragraphstyle *ParagraphStyle) GongCopy() GongstructIF {
	newInstance := new(ParagraphStyle)
	paragraphstyle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rune *Rune) GongCopy() GongstructIF {
	newInstance := new(Rune)
	rune.GongCopyBasicFields(newInstance)
	return newInstance
}

func (runeproperties *RuneProperties) GongCopy() GongstructIF {
	newInstance := new(RuneProperties)
	runeproperties.GongCopyBasicFields(newInstance)
	return newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := new(Table)
	table.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tablecolumn *TableColumn) GongCopy() GongstructIF {
	newInstance := new(TableColumn)
	tablecolumn.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tableproperties *TableProperties) GongCopy() GongstructIF {
	newInstance := new(TableProperties)
	tableproperties.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tablerow *TableRow) GongCopy() GongstructIF {
	newInstance := new(TableRow)
	tablerow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tablestyle *TableStyle) GongCopy() GongstructIF {
	newInstance := new(TableStyle)
	tablestyle.GongCopyBasicFields(newInstance)
	return newInstance
}

func (text *Text) GongCopy() GongstructIF {
	newInstance := new(Text)
	text.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (body *Body) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, body)
}

func (document *Document) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, document)
}

func (docx *Docx) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, docx)
}

func (file *File) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, file)
}

func (node *Node) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, node)
}

func (paragraph *Paragraph) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, paragraph)
}

func (paragraphproperties *ParagraphProperties) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, paragraphproperties)
}

func (paragraphstyle *ParagraphStyle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, paragraphstyle)
}

func (rune *Rune) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rune)
}

func (runeproperties *RuneProperties) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, runeproperties)
}

func (table *Table) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, table)
}

func (tablecolumn *TableColumn) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tablecolumn)
}

func (tableproperties *TableProperties) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tableproperties)
}

func (tablerow *TableRow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tablerow)
}

func (tablestyle *TableStyle) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tablestyle)
}

func (text *Text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, text)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.Bodys,
		stage.Body_stagedOrder,
		stage.Bodys_reference,
		&stage.Bodys_referenceOrder,
		stage.Bodys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Documents,
		stage.Document_stagedOrder,
		stage.Documents_reference,
		&stage.Documents_referenceOrder,
		stage.Documents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Docxs,
		stage.Docx_stagedOrder,
		stage.Docxs_reference,
		&stage.Docxs_referenceOrder,
		stage.Docxs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Files,
		stage.File_stagedOrder,
		stage.Files_reference,
		&stage.Files_referenceOrder,
		stage.Files_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Nodes,
		stage.Node_stagedOrder,
		stage.Nodes_reference,
		&stage.Nodes_referenceOrder,
		stage.Nodes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Paragraphs,
		stage.Paragraph_stagedOrder,
		stage.Paragraphs_reference,
		&stage.Paragraphs_referenceOrder,
		stage.Paragraphs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParagraphPropertiess,
		stage.ParagraphProperties_stagedOrder,
		stage.ParagraphPropertiess_reference,
		&stage.ParagraphPropertiess_referenceOrder,
		stage.ParagraphPropertiess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParagraphStyles,
		stage.ParagraphStyle_stagedOrder,
		stage.ParagraphStyles_reference,
		&stage.ParagraphStyles_referenceOrder,
		stage.ParagraphStyles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Runes,
		stage.Rune_stagedOrder,
		stage.Runes_reference,
		&stage.Runes_referenceOrder,
		stage.Runes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RunePropertiess,
		stage.RuneProperties_stagedOrder,
		stage.RunePropertiess_reference,
		&stage.RunePropertiess_referenceOrder,
		stage.RunePropertiess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tables,
		stage.Table_stagedOrder,
		stage.Tables_reference,
		&stage.Tables_referenceOrder,
		stage.Tables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TableColumns,
		stage.TableColumn_stagedOrder,
		stage.TableColumns_reference,
		&stage.TableColumns_referenceOrder,
		stage.TableColumns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TablePropertiess,
		stage.TableProperties_stagedOrder,
		stage.TablePropertiess_reference,
		&stage.TablePropertiess_referenceOrder,
		stage.TablePropertiess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TableRows,
		stage.TableRow_stagedOrder,
		stage.TableRows_reference,
		&stage.TableRows_referenceOrder,
		stage.TableRows_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TableStyles,
		stage.TableStyle_stagedOrder,
		stage.TableStyles_reference,
		&stage.TableStyles_referenceOrder,
		stage.TableStyles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Texts,
		stage.Text_stagedOrder,
		stage.Texts_reference,
		&stage.Texts_referenceOrder,
		stage.Texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Bodys, &stage.Bodys_reference, &stage.Bodys_referenceOrder, &stage.Bodys_instance)

	__gong__computeReferencePass1(stage, stage.Documents, &stage.Documents_reference, &stage.Documents_referenceOrder, &stage.Documents_instance)

	__gong__computeReferencePass1(stage, stage.Docxs, &stage.Docxs_reference, &stage.Docxs_referenceOrder, &stage.Docxs_instance)

	__gong__computeReferencePass1(stage, stage.Files, &stage.Files_reference, &stage.Files_referenceOrder, &stage.Files_instance)

	__gong__computeReferencePass1(stage, stage.Nodes, &stage.Nodes_reference, &stage.Nodes_referenceOrder, &stage.Nodes_instance)

	__gong__computeReferencePass1(stage, stage.Paragraphs, &stage.Paragraphs_reference, &stage.Paragraphs_referenceOrder, &stage.Paragraphs_instance)

	__gong__computeReferencePass1(stage, stage.ParagraphPropertiess, &stage.ParagraphPropertiess_reference, &stage.ParagraphPropertiess_referenceOrder, &stage.ParagraphPropertiess_instance)

	__gong__computeReferencePass1(stage, stage.ParagraphStyles, &stage.ParagraphStyles_reference, &stage.ParagraphStyles_referenceOrder, &stage.ParagraphStyles_instance)

	__gong__computeReferencePass1(stage, stage.Runes, &stage.Runes_reference, &stage.Runes_referenceOrder, &stage.Runes_instance)

	__gong__computeReferencePass1(stage, stage.RunePropertiess, &stage.RunePropertiess_reference, &stage.RunePropertiess_referenceOrder, &stage.RunePropertiess_instance)

	__gong__computeReferencePass1(stage, stage.Tables, &stage.Tables_reference, &stage.Tables_referenceOrder, &stage.Tables_instance)

	__gong__computeReferencePass1(stage, stage.TableColumns, &stage.TableColumns_reference, &stage.TableColumns_referenceOrder, &stage.TableColumns_instance)

	__gong__computeReferencePass1(stage, stage.TablePropertiess, &stage.TablePropertiess_reference, &stage.TablePropertiess_referenceOrder, &stage.TablePropertiess_instance)

	__gong__computeReferencePass1(stage, stage.TableRows, &stage.TableRows_reference, &stage.TableRows_referenceOrder, &stage.TableRows_instance)

	__gong__computeReferencePass1(stage, stage.TableStyles, &stage.TableStyles_reference, &stage.TableStyles_referenceOrder, &stage.TableStyles_instance)

	__gong__computeReferencePass1(stage, stage.Texts, &stage.Texts_reference, &stage.Texts_referenceOrder, &stage.Texts_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Bodys, stage.Bodys_reference, stage)

	__gong__computeReferencePass2(stage.Documents, stage.Documents_reference, stage)

	__gong__computeReferencePass2(stage.Docxs, stage.Docxs_reference, stage)

	__gong__computeReferencePass2(stage.Files, stage.Files_reference, stage)

	__gong__computeReferencePass2(stage.Nodes, stage.Nodes_reference, stage)

	__gong__computeReferencePass2(stage.Paragraphs, stage.Paragraphs_reference, stage)

	__gong__computeReferencePass2(stage.ParagraphPropertiess, stage.ParagraphPropertiess_reference, stage)

	__gong__computeReferencePass2(stage.ParagraphStyles, stage.ParagraphStyles_reference, stage)

	__gong__computeReferencePass2(stage.Runes, stage.Runes_reference, stage)

	__gong__computeReferencePass2(stage.RunePropertiess, stage.RunePropertiess_reference, stage)

	__gong__computeReferencePass2(stage.Tables, stage.Tables_reference, stage)

	__gong__computeReferencePass2(stage.TableColumns, stage.TableColumns_reference, stage)

	__gong__computeReferencePass2(stage.TablePropertiess, stage.TablePropertiess_reference, stage)

	__gong__computeReferencePass2(stage.TableRows, stage.TableRows_reference, stage)

	__gong__computeReferencePass2(stage.TableStyles, stage.TableStyles_reference, stage)

	__gong__computeReferencePass2(stage.Texts, stage.Texts_reference, stage)

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
	return __gong__getOrder(stage.Body_stagedOrder, stage.Bodys_referenceOrder, body, "Body")
}

func (document *Document) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Document_stagedOrder, stage.Documents_referenceOrder, document, "Document")
}

func (docx *Docx) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Docx_stagedOrder, stage.Docxs_referenceOrder, docx, "Docx")
}

func (file *File) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.File_stagedOrder, stage.Files_referenceOrder, file, "File")
}

func (node *Node) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Node_stagedOrder, stage.Nodes_referenceOrder, node, "Node")
}

func (paragraph *Paragraph) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Paragraph_stagedOrder, stage.Paragraphs_referenceOrder, paragraph, "Paragraph")
}

func (paragraphproperties *ParagraphProperties) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParagraphProperties_stagedOrder, stage.ParagraphPropertiess_referenceOrder, paragraphproperties, "ParagraphProperties")
}

func (paragraphstyle *ParagraphStyle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParagraphStyle_stagedOrder, stage.ParagraphStyles_referenceOrder, paragraphstyle, "ParagraphStyle")
}

func (rune *Rune) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Rune_stagedOrder, stage.Runes_referenceOrder, rune, "Rune")
}

func (runeproperties *RuneProperties) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RuneProperties_stagedOrder, stage.RunePropertiess_referenceOrder, runeproperties, "RuneProperties")
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Table_stagedOrder, stage.Tables_referenceOrder, table, "Table")
}

func (tablecolumn *TableColumn) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TableColumn_stagedOrder, stage.TableColumns_referenceOrder, tablecolumn, "TableColumn")
}

func (tableproperties *TableProperties) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TableProperties_stagedOrder, stage.TablePropertiess_referenceOrder, tableproperties, "TableProperties")
}

func (tablerow *TableRow) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TableRow_stagedOrder, stage.TableRows_referenceOrder, tablerow, "TableRow")
}

func (tablestyle *TableStyle) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TableStyle_stagedOrder, stage.TableStyles_referenceOrder, tablestyle, "TableStyle")
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Text_stagedOrder, stage.Texts_referenceOrder, text, "Text")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (body *Body) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(body, body.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (body *Body) GongGetReferenceIdentifier(stage *Stage) string {
	return body.GongGetIdentifier(stage)
}

func (document *Document) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(document, document.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (document *Document) GongGetReferenceIdentifier(stage *Stage) string {
	return document.GongGetIdentifier(stage)
}

func (docx *Docx) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(docx, docx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (docx *Docx) GongGetReferenceIdentifier(stage *Stage) string {
	return docx.GongGetIdentifier(stage)
}

func (file *File) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(file, file.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (file *File) GongGetReferenceIdentifier(stage *Stage) string {
	return file.GongGetIdentifier(stage)
}

func (node *Node) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(node, node.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (node *Node) GongGetReferenceIdentifier(stage *Stage) string {
	return node.GongGetIdentifier(stage)
}

func (paragraph *Paragraph) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(paragraph, paragraph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraph *Paragraph) GongGetReferenceIdentifier(stage *Stage) string {
	return paragraph.GongGetIdentifier(stage)
}

func (paragraphproperties *ParagraphProperties) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(paragraphproperties, paragraphproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphproperties *ParagraphProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return paragraphproperties.GongGetIdentifier(stage)
}

func (paragraphstyle *ParagraphStyle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(paragraphstyle, paragraphstyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphstyle *ParagraphStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return paragraphstyle.GongGetIdentifier(stage)
}

func (rune *Rune) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rune, rune.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rune *Rune) GongGetReferenceIdentifier(stage *Stage) string {
	return rune.GongGetIdentifier(stage)
}

func (runeproperties *RuneProperties) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(runeproperties, runeproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runeproperties *RuneProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return runeproperties.GongGetIdentifier(stage)
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(table, table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return table.GongGetIdentifier(stage)
}

func (tablecolumn *TableColumn) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tablecolumn, tablecolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablecolumn *TableColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return tablecolumn.GongGetIdentifier(stage)
}

func (tableproperties *TableProperties) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tableproperties, tableproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tableproperties *TableProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return tableproperties.GongGetIdentifier(stage)
}

func (tablerow *TableRow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tablerow, tablerow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablerow *TableRow) GongGetReferenceIdentifier(stage *Stage) string {
	return tablerow.GongGetIdentifier(stage)
}

func (tablestyle *TableStyle) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tablestyle, tablestyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablestyle *TableStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return tablestyle.GongGetIdentifier(stage)
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(text, text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return text.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (body *Body) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(body.GongGetIdentifier(stage), "Body", body.Name)
}

func (document *Document) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(document.GongGetIdentifier(stage), "Document", document.Name)
}

func (docx *Docx) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(docx.GongGetIdentifier(stage), "Docx", docx.Name)
}

func (file *File) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(file.GongGetIdentifier(stage), "File", file.Name)
}

func (node *Node) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(node.GongGetIdentifier(stage), "Node", node.Name)
}

func (paragraph *Paragraph) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(paragraph.GongGetIdentifier(stage), "Paragraph", paragraph.Name)
}

func (paragraphproperties *ParagraphProperties) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(paragraphproperties.GongGetIdentifier(stage), "ParagraphProperties", paragraphproperties.Name)
}

func (paragraphstyle *ParagraphStyle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(paragraphstyle.GongGetIdentifier(stage), "ParagraphStyle", paragraphstyle.Name)
}

func (rune *Rune) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rune.GongGetIdentifier(stage), "Rune", rune.Name)
}

func (runeproperties *RuneProperties) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(runeproperties.GongGetIdentifier(stage), "RuneProperties", runeproperties.Name)
}

func (table *Table) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(table.GongGetIdentifier(stage), "Table", table.Name)
}

func (tablecolumn *TableColumn) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tablecolumn.GongGetIdentifier(stage), "TableColumn", tablecolumn.Name)
}

func (tableproperties *TableProperties) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tableproperties.GongGetIdentifier(stage), "TableProperties", tableproperties.Name)
}

func (tablerow *TableRow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tablerow.GongGetIdentifier(stage), "TableRow", tablerow.Name)
}

func (tablestyle *TableStyle) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tablestyle.GongGetIdentifier(stage), "TableStyle", tablestyle.Name)
}

func (text *Text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(text.GongGetIdentifier(stage), "Text", text.Name)
}

// insertion point for unstaging
func (body *Body) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(body.GongGetReferenceIdentifier(stage))
}

func (document *Document) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(document.GongGetReferenceIdentifier(stage))
}

func (docx *Docx) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(docx.GongGetReferenceIdentifier(stage))
}

func (file *File) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(file.GongGetReferenceIdentifier(stage))
}

func (node *Node) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(node.GongGetReferenceIdentifier(stage))
}

func (paragraph *Paragraph) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(paragraph.GongGetReferenceIdentifier(stage))
}

func (paragraphproperties *ParagraphProperties) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(paragraphproperties.GongGetReferenceIdentifier(stage))
}

func (paragraphstyle *ParagraphStyle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(paragraphstyle.GongGetReferenceIdentifier(stage))
}

func (rune *Rune) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rune.GongGetReferenceIdentifier(stage))
}

func (runeproperties *RuneProperties) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(runeproperties.GongGetReferenceIdentifier(stage))
}

func (table *Table) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(table.GongGetReferenceIdentifier(stage))
}

func (tablecolumn *TableColumn) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tablecolumn.GongGetReferenceIdentifier(stage))
}

func (tableproperties *TableProperties) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tableproperties.GongGetReferenceIdentifier(stage))
}

func (tablerow *TableRow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tablerow.GongGetReferenceIdentifier(stage))
}

func (tablestyle *TableStyle) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tablestyle.GongGetReferenceIdentifier(stage))
}

func (text *Text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(text.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
