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

	// end of insertion point per named struct
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
func (body *Body) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(body).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(body), uint64(stage.GetOrder(body)))
	return
}

func (document *Document) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(document).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(document), uint64(stage.GetOrder(document)))
	return
}

func (docx *Docx) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(docx).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(docx), uint64(stage.GetOrder(docx)))
	return
}

func (file *File) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(file).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(file), uint64(stage.GetOrder(file)))
	return
}

func (node *Node) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(node).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(node), uint64(stage.GetOrder(node)))
	return
}

func (paragraph *Paragraph) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(paragraph).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(paragraph), uint64(stage.GetOrder(paragraph)))
	return
}

func (paragraphproperties *ParagraphProperties) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(paragraphproperties).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(paragraphproperties), uint64(stage.GetOrder(paragraphproperties)))
	return
}

func (paragraphstyle *ParagraphStyle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(paragraphstyle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(paragraphstyle), uint64(stage.GetOrder(paragraphstyle)))
	return
}

func (rune *Rune) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rune).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rune), uint64(stage.GetOrder(rune)))
	return
}

func (runeproperties *RuneProperties) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(runeproperties).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(runeproperties), uint64(stage.GetOrder(runeproperties)))
	return
}

func (table *Table) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(table).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(table), uint64(stage.GetOrder(table)))
	return
}

func (tablecolumn *TableColumn) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tablecolumn).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tablecolumn), uint64(stage.GetOrder(tablecolumn)))
	return
}

func (tableproperties *TableProperties) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tableproperties).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tableproperties), uint64(stage.GetOrder(tableproperties)))
	return
}

func (tablerow *TableRow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tablerow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tablerow), uint64(stage.GetOrder(tablerow)))
	return
}

func (tablestyle *TableStyle) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tablestyle).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tablestyle), uint64(stage.GetOrder(tablestyle)))
	return
}

func (text *Text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(text), uint64(stage.GetOrder(text)))
	return
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
	stage.Bodys_reference = make(map[*Body]*Body)
	stage.Bodys_referenceOrder = make(map[*Body]uint) // diff Unstage needs the reference order
	stage.Bodys_instance = make(map[*Body]*Body)
	for instance := range stage.Bodys {
		_copy := instance.GongCopy().(*Body)
		stage.Bodys_reference[instance] = _copy
		stage.Bodys_instance[_copy] = instance
		stage.Bodys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Documents_reference = make(map[*Document]*Document)
	stage.Documents_referenceOrder = make(map[*Document]uint) // diff Unstage needs the reference order
	stage.Documents_instance = make(map[*Document]*Document)
	for instance := range stage.Documents {
		_copy := instance.GongCopy().(*Document)
		stage.Documents_reference[instance] = _copy
		stage.Documents_instance[_copy] = instance
		stage.Documents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Docxs_reference = make(map[*Docx]*Docx)
	stage.Docxs_referenceOrder = make(map[*Docx]uint) // diff Unstage needs the reference order
	stage.Docxs_instance = make(map[*Docx]*Docx)
	for instance := range stage.Docxs {
		_copy := instance.GongCopy().(*Docx)
		stage.Docxs_reference[instance] = _copy
		stage.Docxs_instance[_copy] = instance
		stage.Docxs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Files_reference = make(map[*File]*File)
	stage.Files_referenceOrder = make(map[*File]uint) // diff Unstage needs the reference order
	stage.Files_instance = make(map[*File]*File)
	for instance := range stage.Files {
		_copy := instance.GongCopy().(*File)
		stage.Files_reference[instance] = _copy
		stage.Files_instance[_copy] = instance
		stage.Files_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Nodes_reference = make(map[*Node]*Node)
	stage.Nodes_referenceOrder = make(map[*Node]uint) // diff Unstage needs the reference order
	stage.Nodes_instance = make(map[*Node]*Node)
	for instance := range stage.Nodes {
		_copy := instance.GongCopy().(*Node)
		stage.Nodes_reference[instance] = _copy
		stage.Nodes_instance[_copy] = instance
		stage.Nodes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Paragraphs_reference = make(map[*Paragraph]*Paragraph)
	stage.Paragraphs_referenceOrder = make(map[*Paragraph]uint) // diff Unstage needs the reference order
	stage.Paragraphs_instance = make(map[*Paragraph]*Paragraph)
	for instance := range stage.Paragraphs {
		_copy := instance.GongCopy().(*Paragraph)
		stage.Paragraphs_reference[instance] = _copy
		stage.Paragraphs_instance[_copy] = instance
		stage.Paragraphs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParagraphPropertiess_reference = make(map[*ParagraphProperties]*ParagraphProperties)
	stage.ParagraphPropertiess_referenceOrder = make(map[*ParagraphProperties]uint) // diff Unstage needs the reference order
	stage.ParagraphPropertiess_instance = make(map[*ParagraphProperties]*ParagraphProperties)
	for instance := range stage.ParagraphPropertiess {
		_copy := instance.GongCopy().(*ParagraphProperties)
		stage.ParagraphPropertiess_reference[instance] = _copy
		stage.ParagraphPropertiess_instance[_copy] = instance
		stage.ParagraphPropertiess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParagraphStyles_reference = make(map[*ParagraphStyle]*ParagraphStyle)
	stage.ParagraphStyles_referenceOrder = make(map[*ParagraphStyle]uint) // diff Unstage needs the reference order
	stage.ParagraphStyles_instance = make(map[*ParagraphStyle]*ParagraphStyle)
	for instance := range stage.ParagraphStyles {
		_copy := instance.GongCopy().(*ParagraphStyle)
		stage.ParagraphStyles_reference[instance] = _copy
		stage.ParagraphStyles_instance[_copy] = instance
		stage.ParagraphStyles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Runes_reference = make(map[*Rune]*Rune)
	stage.Runes_referenceOrder = make(map[*Rune]uint) // diff Unstage needs the reference order
	stage.Runes_instance = make(map[*Rune]*Rune)
	for instance := range stage.Runes {
		_copy := instance.GongCopy().(*Rune)
		stage.Runes_reference[instance] = _copy
		stage.Runes_instance[_copy] = instance
		stage.Runes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RunePropertiess_reference = make(map[*RuneProperties]*RuneProperties)
	stage.RunePropertiess_referenceOrder = make(map[*RuneProperties]uint) // diff Unstage needs the reference order
	stage.RunePropertiess_instance = make(map[*RuneProperties]*RuneProperties)
	for instance := range stage.RunePropertiess {
		_copy := instance.GongCopy().(*RuneProperties)
		stage.RunePropertiess_reference[instance] = _copy
		stage.RunePropertiess_instance[_copy] = instance
		stage.RunePropertiess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	stage.Tables_instance = make(map[*Table]*Table)
	for instance := range stage.Tables {
		_copy := instance.GongCopy().(*Table)
		stage.Tables_reference[instance] = _copy
		stage.Tables_instance[_copy] = instance
		stage.Tables_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TableColumns_reference = make(map[*TableColumn]*TableColumn)
	stage.TableColumns_referenceOrder = make(map[*TableColumn]uint) // diff Unstage needs the reference order
	stage.TableColumns_instance = make(map[*TableColumn]*TableColumn)
	for instance := range stage.TableColumns {
		_copy := instance.GongCopy().(*TableColumn)
		stage.TableColumns_reference[instance] = _copy
		stage.TableColumns_instance[_copy] = instance
		stage.TableColumns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TablePropertiess_reference = make(map[*TableProperties]*TableProperties)
	stage.TablePropertiess_referenceOrder = make(map[*TableProperties]uint) // diff Unstage needs the reference order
	stage.TablePropertiess_instance = make(map[*TableProperties]*TableProperties)
	for instance := range stage.TablePropertiess {
		_copy := instance.GongCopy().(*TableProperties)
		stage.TablePropertiess_reference[instance] = _copy
		stage.TablePropertiess_instance[_copy] = instance
		stage.TablePropertiess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TableRows_reference = make(map[*TableRow]*TableRow)
	stage.TableRows_referenceOrder = make(map[*TableRow]uint) // diff Unstage needs the reference order
	stage.TableRows_instance = make(map[*TableRow]*TableRow)
	for instance := range stage.TableRows {
		_copy := instance.GongCopy().(*TableRow)
		stage.TableRows_reference[instance] = _copy
		stage.TableRows_instance[_copy] = instance
		stage.TableRows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TableStyles_reference = make(map[*TableStyle]*TableStyle)
	stage.TableStyles_referenceOrder = make(map[*TableStyle]uint) // diff Unstage needs the reference order
	stage.TableStyles_instance = make(map[*TableStyle]*TableStyle)
	for instance := range stage.TableStyles {
		_copy := instance.GongCopy().(*TableStyle)
		stage.TableStyles_reference[instance] = _copy
		stage.TableStyles_instance[_copy] = instance
		stage.TableStyles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Texts_reference = make(map[*Text]*Text)
	stage.Texts_referenceOrder = make(map[*Text]uint) // diff Unstage needs the reference order
	stage.Texts_instance = make(map[*Text]*Text)
	for instance := range stage.Texts {
		_copy := instance.GongCopy().(*Text)
		stage.Texts_reference[instance] = _copy
		stage.Texts_instance[_copy] = instance
		stage.Texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Bodys {
		reference := stage.Bodys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Documents {
		reference := stage.Documents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Docxs {
		reference := stage.Docxs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Files {
		reference := stage.Files_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Nodes {
		reference := stage.Nodes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Paragraphs {
		reference := stage.Paragraphs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParagraphPropertiess {
		reference := stage.ParagraphPropertiess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParagraphStyles {
		reference := stage.ParagraphStyles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Runes {
		reference := stage.Runes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RunePropertiess {
		reference := stage.RunePropertiess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tables {
		reference := stage.Tables_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TableColumns {
		reference := stage.TableColumns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TablePropertiess {
		reference := stage.TablePropertiess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TableRows {
		reference := stage.TableRows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TableStyles {
		reference := stage.TableStyles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Texts {
		reference := stage.Texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
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
	if order, ok := stage.Body_stagedOrder[body]; ok {
		return order
	}
	if order, ok := stage.Bodys_referenceOrder[body]; ok {
		return order
	} else {
		log.Printf("instance %p of type Body was not staged and does not have a reference order", body)
		return 0
	}
}

func (document *Document) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Document_stagedOrder[document]; ok {
		return order
	}
	if order, ok := stage.Documents_referenceOrder[document]; ok {
		return order
	} else {
		log.Printf("instance %p of type Document was not staged and does not have a reference order", document)
		return 0
	}
}

func (docx *Docx) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Docx_stagedOrder[docx]; ok {
		return order
	}
	if order, ok := stage.Docxs_referenceOrder[docx]; ok {
		return order
	} else {
		log.Printf("instance %p of type Docx was not staged and does not have a reference order", docx)
		return 0
	}
}

func (file *File) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.File_stagedOrder[file]; ok {
		return order
	}
	if order, ok := stage.Files_referenceOrder[file]; ok {
		return order
	} else {
		log.Printf("instance %p of type File was not staged and does not have a reference order", file)
		return 0
	}
}

func (node *Node) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Node_stagedOrder[node]; ok {
		return order
	}
	if order, ok := stage.Nodes_referenceOrder[node]; ok {
		return order
	} else {
		log.Printf("instance %p of type Node was not staged and does not have a reference order", node)
		return 0
	}
}

func (paragraph *Paragraph) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Paragraph_stagedOrder[paragraph]; ok {
		return order
	}
	if order, ok := stage.Paragraphs_referenceOrder[paragraph]; ok {
		return order
	} else {
		log.Printf("instance %p of type Paragraph was not staged and does not have a reference order", paragraph)
		return 0
	}
}

func (paragraphproperties *ParagraphProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParagraphProperties_stagedOrder[paragraphproperties]; ok {
		return order
	}
	if order, ok := stage.ParagraphPropertiess_referenceOrder[paragraphproperties]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParagraphProperties was not staged and does not have a reference order", paragraphproperties)
		return 0
	}
}

func (paragraphstyle *ParagraphStyle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParagraphStyle_stagedOrder[paragraphstyle]; ok {
		return order
	}
	if order, ok := stage.ParagraphStyles_referenceOrder[paragraphstyle]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParagraphStyle was not staged and does not have a reference order", paragraphstyle)
		return 0
	}
}

func (rune *Rune) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Rune_stagedOrder[rune]; ok {
		return order
	}
	if order, ok := stage.Runes_referenceOrder[rune]; ok {
		return order
	} else {
		log.Printf("instance %p of type Rune was not staged and does not have a reference order", rune)
		return 0
	}
}

func (runeproperties *RuneProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RuneProperties_stagedOrder[runeproperties]; ok {
		return order
	}
	if order, ok := stage.RunePropertiess_referenceOrder[runeproperties]; ok {
		return order
	} else {
		log.Printf("instance %p of type RuneProperties was not staged and does not have a reference order", runeproperties)
		return 0
	}
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Table_stagedOrder[table]; ok {
		return order
	}
	if order, ok := stage.Tables_referenceOrder[table]; ok {
		return order
	} else {
		log.Printf("instance %p of type Table was not staged and does not have a reference order", table)
		return 0
	}
}

func (tablecolumn *TableColumn) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableColumn_stagedOrder[tablecolumn]; ok {
		return order
	}
	if order, ok := stage.TableColumns_referenceOrder[tablecolumn]; ok {
		return order
	} else {
		log.Printf("instance %p of type TableColumn was not staged and does not have a reference order", tablecolumn)
		return 0
	}
}

func (tableproperties *TableProperties) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableProperties_stagedOrder[tableproperties]; ok {
		return order
	}
	if order, ok := stage.TablePropertiess_referenceOrder[tableproperties]; ok {
		return order
	} else {
		log.Printf("instance %p of type TableProperties was not staged and does not have a reference order", tableproperties)
		return 0
	}
}

func (tablerow *TableRow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableRow_stagedOrder[tablerow]; ok {
		return order
	}
	if order, ok := stage.TableRows_referenceOrder[tablerow]; ok {
		return order
	} else {
		log.Printf("instance %p of type TableRow was not staged and does not have a reference order", tablerow)
		return 0
	}
}

func (tablestyle *TableStyle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableStyle_stagedOrder[tablestyle]; ok {
		return order
	}
	if order, ok := stage.TableStyles_referenceOrder[tablestyle]; ok {
		return order
	} else {
		log.Printf("instance %p of type TableStyle was not staged and does not have a reference order", tablestyle)
		return 0
	}
}

func (text *Text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Text_stagedOrder[text]; ok {
		return order
	}
	if order, ok := stage.Texts_referenceOrder[text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Text was not staged and does not have a reference order", text)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", body.GongGetGongstructName(), body.GongGetOrder(stage))
}

func (document *Document) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (document *Document) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", document.GongGetGongstructName(), document.GongGetOrder(stage))
}

func (docx *Docx) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", docx.GongGetGongstructName(), docx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (docx *Docx) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", docx.GongGetGongstructName(), docx.GongGetOrder(stage))
}

func (file *File) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", file.GongGetGongstructName(), file.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (file *File) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", file.GongGetGongstructName(), file.GongGetOrder(stage))
}

func (node *Node) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (node *Node) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", node.GongGetGongstructName(), node.GongGetOrder(stage))
}

func (paragraph *Paragraph) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraph.GongGetGongstructName(), paragraph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraph *Paragraph) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraph.GongGetGongstructName(), paragraph.GongGetOrder(stage))
}

func (paragraphproperties *ParagraphProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphproperties.GongGetGongstructName(), paragraphproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphproperties *ParagraphProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphproperties.GongGetGongstructName(), paragraphproperties.GongGetOrder(stage))
}

func (paragraphstyle *ParagraphStyle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphstyle.GongGetGongstructName(), paragraphstyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (paragraphstyle *ParagraphStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", paragraphstyle.GongGetGongstructName(), paragraphstyle.GongGetOrder(stage))
}

func (rune *Rune) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rune.GongGetGongstructName(), rune.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rune *Rune) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rune.GongGetGongstructName(), rune.GongGetOrder(stage))
}

func (runeproperties *RuneProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", runeproperties.GongGetGongstructName(), runeproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (runeproperties *RuneProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", runeproperties.GongGetGongstructName(), runeproperties.GongGetOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

func (tablecolumn *TableColumn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablecolumn.GongGetGongstructName(), tablecolumn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablecolumn *TableColumn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablecolumn.GongGetGongstructName(), tablecolumn.GongGetOrder(stage))
}

func (tableproperties *TableProperties) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tableproperties.GongGetGongstructName(), tableproperties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tableproperties *TableProperties) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tableproperties.GongGetGongstructName(), tableproperties.GongGetOrder(stage))
}

func (tablerow *TableRow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablerow.GongGetGongstructName(), tablerow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablerow *TableRow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablerow.GongGetGongstructName(), tablerow.GongGetOrder(stage))
}

func (tablestyle *TableStyle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablestyle.GongGetGongstructName(), tablestyle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tablestyle *TableStyle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tablestyle.GongGetGongstructName(), tablestyle.GongGetOrder(stage))
}

func (text *Text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text *Text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text.GongGetGongstructName(), text.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (body *Body) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", body.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Body")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(body.Name))
	return
}

func (document *Document) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", document.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Document")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(document.Name))
	return
}

func (docx *Docx) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", docx.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Docx")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(docx.Name))
	return
}

func (file *File) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", file.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "File")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(file.Name))
	return
}

func (node *Node) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", node.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Node")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(node.Name))
	return
}

func (paragraph *Paragraph) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraph.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Paragraph")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(paragraph.Name))
	return
}

func (paragraphproperties *ParagraphProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(paragraphproperties.Name))
	return
}

func (paragraphstyle *ParagraphStyle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", paragraphstyle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParagraphStyle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(paragraphstyle.Name))
	return
}

func (rune *Rune) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rune.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rune")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rune.Name))
	return
}

func (runeproperties *RuneProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", runeproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RuneProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(runeproperties.Name))
	return
}

func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(table.Name))
	return
}

func (tablecolumn *TableColumn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablecolumn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableColumn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tablecolumn.Name))
	return
}

func (tableproperties *TableProperties) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tableproperties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableProperties")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tableproperties.Name))
	return
}

func (tablerow *TableRow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablerow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableRow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tablerow.Name))
	return
}

func (tablestyle *TableStyle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tablestyle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TableStyle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tablestyle.Name))
	return
}

func (text *Text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(text.Name))
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

// end of template
