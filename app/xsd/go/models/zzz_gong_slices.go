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
	// Compute reverse map for named struct All
	// insertion point per field
	stage.All_Sequences_reverseMap = make(map[*Sequence]*All)
	for all := range stage.Alls {
		_ = all
		for _, _sequence := range all.Sequences {
			stage.All_Sequences_reverseMap[_sequence] = all
		}
	}
	stage.All_Alls_reverseMap = make(map[*All]*All)
	for all := range stage.Alls {
		_ = all
		for _, _all := range all.Alls {
			stage.All_Alls_reverseMap[_all] = all
		}
	}
	stage.All_Choices_reverseMap = make(map[*Choice]*All)
	for all := range stage.Alls {
		_ = all
		for _, _choice := range all.Choices {
			stage.All_Choices_reverseMap[_choice] = all
		}
	}
	stage.All_Groups_reverseMap = make(map[*Group]*All)
	for all := range stage.Alls {
		_ = all
		for _, _group := range all.Groups {
			stage.All_Groups_reverseMap[_group] = all
		}
	}
	stage.All_Elements_reverseMap = make(map[*Element]*All)
	for all := range stage.Alls {
		_ = all
		for _, _element := range all.Elements {
			stage.All_Elements_reverseMap[_element] = all
		}
	}

	// Compute reverse map for named struct Annotation
	// insertion point per field
	stage.Annotation_Documentations_reverseMap = make(map[*Documentation]*Annotation)
	for annotation := range stage.Annotations {
		_ = annotation
		for _, _documentation := range annotation.Documentations {
			stage.Annotation_Documentations_reverseMap[_documentation] = annotation
		}
	}

	// Compute reverse map for named struct Attribute
	// insertion point per field

	// Compute reverse map for named struct AttributeGroup
	// insertion point per field
	stage.AttributeGroup_AttributeGroups_reverseMap = make(map[*AttributeGroup]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attributegroup := range attributegroup.AttributeGroups {
			stage.AttributeGroup_AttributeGroups_reverseMap[_attributegroup] = attributegroup
		}
	}
	stage.AttributeGroup_Attributes_reverseMap = make(map[*Attribute]*AttributeGroup)
	for attributegroup := range stage.AttributeGroups {
		_ = attributegroup
		for _, _attribute := range attributegroup.Attributes {
			stage.AttributeGroup_Attributes_reverseMap[_attribute] = attributegroup
		}
	}

	// Compute reverse map for named struct Choice
	// insertion point per field
	stage.Choice_Sequences_reverseMap = make(map[*Sequence]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _sequence := range choice.Sequences {
			stage.Choice_Sequences_reverseMap[_sequence] = choice
		}
	}
	stage.Choice_Alls_reverseMap = make(map[*All]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _all := range choice.Alls {
			stage.Choice_Alls_reverseMap[_all] = choice
		}
	}
	stage.Choice_Choices_reverseMap = make(map[*Choice]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _choice := range choice.Choices {
			stage.Choice_Choices_reverseMap[_choice] = choice
		}
	}
	stage.Choice_Groups_reverseMap = make(map[*Group]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _group := range choice.Groups {
			stage.Choice_Groups_reverseMap[_group] = choice
		}
	}
	stage.Choice_Elements_reverseMap = make(map[*Element]*Choice)
	for choice := range stage.Choices {
		_ = choice
		for _, _element := range choice.Elements {
			stage.Choice_Elements_reverseMap[_element] = choice
		}
	}

	// Compute reverse map for named struct ComplexContent
	// insertion point per field

	// Compute reverse map for named struct ComplexType
	// insertion point per field
	stage.ComplexType_Sequences_reverseMap = make(map[*Sequence]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _sequence := range complextype.Sequences {
			stage.ComplexType_Sequences_reverseMap[_sequence] = complextype
		}
	}
	stage.ComplexType_Alls_reverseMap = make(map[*All]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _all := range complextype.Alls {
			stage.ComplexType_Alls_reverseMap[_all] = complextype
		}
	}
	stage.ComplexType_Choices_reverseMap = make(map[*Choice]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _choice := range complextype.Choices {
			stage.ComplexType_Choices_reverseMap[_choice] = complextype
		}
	}
	stage.ComplexType_Groups_reverseMap = make(map[*Group]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _group := range complextype.Groups {
			stage.ComplexType_Groups_reverseMap[_group] = complextype
		}
	}
	stage.ComplexType_Elements_reverseMap = make(map[*Element]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _element := range complextype.Elements {
			stage.ComplexType_Elements_reverseMap[_element] = complextype
		}
	}
	stage.ComplexType_Attributes_reverseMap = make(map[*Attribute]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attribute := range complextype.Attributes {
			stage.ComplexType_Attributes_reverseMap[_attribute] = complextype
		}
	}
	stage.ComplexType_AttributeGroups_reverseMap = make(map[*AttributeGroup]*ComplexType)
	for complextype := range stage.ComplexTypes {
		_ = complextype
		for _, _attributegroup := range complextype.AttributeGroups {
			stage.ComplexType_AttributeGroups_reverseMap[_attributegroup] = complextype
		}
	}

	// Compute reverse map for named struct Documentation
	// insertion point per field

	// Compute reverse map for named struct Element
	// insertion point per field
	stage.Element_Groups_reverseMap = make(map[*Group]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _group := range element.Groups {
			stage.Element_Groups_reverseMap[_group] = element
		}
	}

	// Compute reverse map for named struct Enumeration
	// insertion point per field

	// Compute reverse map for named struct Extension
	// insertion point per field
	stage.Extension_Sequences_reverseMap = make(map[*Sequence]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _sequence := range extension.Sequences {
			stage.Extension_Sequences_reverseMap[_sequence] = extension
		}
	}
	stage.Extension_Alls_reverseMap = make(map[*All]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _all := range extension.Alls {
			stage.Extension_Alls_reverseMap[_all] = extension
		}
	}
	stage.Extension_Choices_reverseMap = make(map[*Choice]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _choice := range extension.Choices {
			stage.Extension_Choices_reverseMap[_choice] = extension
		}
	}
	stage.Extension_Groups_reverseMap = make(map[*Group]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _group := range extension.Groups {
			stage.Extension_Groups_reverseMap[_group] = extension
		}
	}
	stage.Extension_Elements_reverseMap = make(map[*Element]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _element := range extension.Elements {
			stage.Extension_Elements_reverseMap[_element] = extension
		}
	}
	stage.Extension_Attributes_reverseMap = make(map[*Attribute]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _attribute := range extension.Attributes {
			stage.Extension_Attributes_reverseMap[_attribute] = extension
		}
	}
	stage.Extension_AttributeGroups_reverseMap = make(map[*AttributeGroup]*Extension)
	for extension := range stage.Extensions {
		_ = extension
		for _, _attributegroup := range extension.AttributeGroups {
			stage.Extension_AttributeGroups_reverseMap[_attributegroup] = extension
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Sequences_reverseMap = make(map[*Sequence]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _sequence := range group.Sequences {
			stage.Group_Sequences_reverseMap[_sequence] = group
		}
	}
	stage.Group_Alls_reverseMap = make(map[*All]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _all := range group.Alls {
			stage.Group_Alls_reverseMap[_all] = group
		}
	}
	stage.Group_Choices_reverseMap = make(map[*Choice]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _choice := range group.Choices {
			stage.Group_Choices_reverseMap[_choice] = group
		}
	}
	stage.Group_Groups_reverseMap = make(map[*Group]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _group := range group.Groups {
			stage.Group_Groups_reverseMap[_group] = group
		}
	}
	stage.Group_Elements_reverseMap = make(map[*Element]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _element := range group.Elements {
			stage.Group_Elements_reverseMap[_element] = group
		}
	}

	// Compute reverse map for named struct Length
	// insertion point per field

	// Compute reverse map for named struct MaxInclusive
	// insertion point per field

	// Compute reverse map for named struct MaxLength
	// insertion point per field

	// Compute reverse map for named struct MinInclusive
	// insertion point per field

	// Compute reverse map for named struct MinLength
	// insertion point per field

	// Compute reverse map for named struct Pattern
	// insertion point per field

	// Compute reverse map for named struct Restriction
	// insertion point per field
	stage.Restriction_Enumerations_reverseMap = make(map[*Enumeration]*Restriction)
	for restriction := range stage.Restrictions {
		_ = restriction
		for _, _enumeration := range restriction.Enumerations {
			stage.Restriction_Enumerations_reverseMap[_enumeration] = restriction
		}
	}

	// Compute reverse map for named struct Schema
	// insertion point per field
	stage.Schema_Elements_reverseMap = make(map[*Element]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _element := range schema.Elements {
			stage.Schema_Elements_reverseMap[_element] = schema
		}
	}
	stage.Schema_SimpleTypes_reverseMap = make(map[*SimpleType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _simpletype := range schema.SimpleTypes {
			stage.Schema_SimpleTypes_reverseMap[_simpletype] = schema
		}
	}
	stage.Schema_ComplexTypes_reverseMap = make(map[*ComplexType]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _complextype := range schema.ComplexTypes {
			stage.Schema_ComplexTypes_reverseMap[_complextype] = schema
		}
	}
	stage.Schema_AttributeGroups_reverseMap = make(map[*AttributeGroup]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _attributegroup := range schema.AttributeGroups {
			stage.Schema_AttributeGroups_reverseMap[_attributegroup] = schema
		}
	}
	stage.Schema_Groups_reverseMap = make(map[*Group]*Schema)
	for schema := range stage.Schemas {
		_ = schema
		for _, _group := range schema.Groups {
			stage.Schema_Groups_reverseMap[_group] = schema
		}
	}

	// Compute reverse map for named struct Sequence
	// insertion point per field
	stage.Sequence_Sequences_reverseMap = make(map[*Sequence]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _sequence := range sequence.Sequences {
			stage.Sequence_Sequences_reverseMap[_sequence] = sequence
		}
	}
	stage.Sequence_Alls_reverseMap = make(map[*All]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _all := range sequence.Alls {
			stage.Sequence_Alls_reverseMap[_all] = sequence
		}
	}
	stage.Sequence_Choices_reverseMap = make(map[*Choice]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _choice := range sequence.Choices {
			stage.Sequence_Choices_reverseMap[_choice] = sequence
		}
	}
	stage.Sequence_Groups_reverseMap = make(map[*Group]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _group := range sequence.Groups {
			stage.Sequence_Groups_reverseMap[_group] = sequence
		}
	}
	stage.Sequence_Elements_reverseMap = make(map[*Element]*Sequence)
	for sequence := range stage.Sequences {
		_ = sequence
		for _, _element := range sequence.Elements {
			stage.Sequence_Elements_reverseMap[_element] = sequence
		}
	}

	// Compute reverse map for named struct SimpleContent
	// insertion point per field

	// Compute reverse map for named struct SimpleType
	// insertion point per field

	// Compute reverse map for named struct TotalDigit
	// insertion point per field

	// Compute reverse map for named struct Union
	// insertion point per field

	// Compute reverse map for named struct WhiteSpace
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Alls {
		res = append(res, instance)
	}

	for instance := range stage.Annotations {
		res = append(res, instance)
	}

	for instance := range stage.Attributes {
		res = append(res, instance)
	}

	for instance := range stage.AttributeGroups {
		res = append(res, instance)
	}

	for instance := range stage.Choices {
		res = append(res, instance)
	}

	for instance := range stage.ComplexContents {
		res = append(res, instance)
	}

	for instance := range stage.ComplexTypes {
		res = append(res, instance)
	}

	for instance := range stage.Documentations {
		res = append(res, instance)
	}

	for instance := range stage.Elements {
		res = append(res, instance)
	}

	for instance := range stage.Enumerations {
		res = append(res, instance)
	}

	for instance := range stage.Extensions {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Lengths {
		res = append(res, instance)
	}

	for instance := range stage.MaxInclusives {
		res = append(res, instance)
	}

	for instance := range stage.MaxLengths {
		res = append(res, instance)
	}

	for instance := range stage.MinInclusives {
		res = append(res, instance)
	}

	for instance := range stage.MinLengths {
		res = append(res, instance)
	}

	for instance := range stage.Patterns {
		res = append(res, instance)
	}

	for instance := range stage.Restrictions {
		res = append(res, instance)
	}

	for instance := range stage.Schemas {
		res = append(res, instance)
	}

	for instance := range stage.Sequences {
		res = append(res, instance)
	}

	for instance := range stage.SimpleContents {
		res = append(res, instance)
	}

	for instance := range stage.SimpleTypes {
		res = append(res, instance)
	}

	for instance := range stage.TotalDigits {
		res = append(res, instance)
	}

	for instance := range stage.Unions {
		res = append(res, instance)
	}

	for instance := range stage.WhiteSpaces {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (all *All) GongCopy() GongstructIF {
	newInstance := new(All)
	all.GongCopyBasicFields(newInstance)
	return newInstance
}

func (annotation *Annotation) GongCopy() GongstructIF {
	newInstance := new(Annotation)
	annotation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute *Attribute) GongCopy() GongstructIF {
	newInstance := new(Attribute)
	attribute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attributegroup *AttributeGroup) GongCopy() GongstructIF {
	newInstance := new(AttributeGroup)
	attributegroup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (choice *Choice) GongCopy() GongstructIF {
	newInstance := new(Choice)
	choice.GongCopyBasicFields(newInstance)
	return newInstance
}

func (complexcontent *ComplexContent) GongCopy() GongstructIF {
	newInstance := new(ComplexContent)
	complexcontent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (complextype *ComplexType) GongCopy() GongstructIF {
	newInstance := new(ComplexType)
	complextype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (documentation *Documentation) GongCopy() GongstructIF {
	newInstance := new(Documentation)
	documentation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (element *Element) GongCopy() GongstructIF {
	newInstance := new(Element)
	element.GongCopyBasicFields(newInstance)
	return newInstance
}

func (enumeration *Enumeration) GongCopy() GongstructIF {
	newInstance := new(Enumeration)
	enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (extension *Extension) GongCopy() GongstructIF {
	newInstance := new(Extension)
	extension.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (length *Length) GongCopy() GongstructIF {
	newInstance := new(Length)
	length.GongCopyBasicFields(newInstance)
	return newInstance
}

func (maxinclusive *MaxInclusive) GongCopy() GongstructIF {
	newInstance := new(MaxInclusive)
	maxinclusive.GongCopyBasicFields(newInstance)
	return newInstance
}

func (maxlength *MaxLength) GongCopy() GongstructIF {
	newInstance := new(MaxLength)
	maxlength.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mininclusive *MinInclusive) GongCopy() GongstructIF {
	newInstance := new(MinInclusive)
	mininclusive.GongCopyBasicFields(newInstance)
	return newInstance
}

func (minlength *MinLength) GongCopy() GongstructIF {
	newInstance := new(MinLength)
	minlength.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pattern *Pattern) GongCopy() GongstructIF {
	newInstance := new(Pattern)
	pattern.GongCopyBasicFields(newInstance)
	return newInstance
}

func (restriction *Restriction) GongCopy() GongstructIF {
	newInstance := new(Restriction)
	restriction.GongCopyBasicFields(newInstance)
	return newInstance
}

func (schema *Schema) GongCopy() GongstructIF {
	newInstance := new(Schema)
	schema.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sequence *Sequence) GongCopy() GongstructIF {
	newInstance := new(Sequence)
	sequence.GongCopyBasicFields(newInstance)
	return newInstance
}

func (simplecontent *SimpleContent) GongCopy() GongstructIF {
	newInstance := new(SimpleContent)
	simplecontent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (simpletype *SimpleType) GongCopy() GongstructIF {
	newInstance := new(SimpleType)
	simpletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (totaldigit *TotalDigit) GongCopy() GongstructIF {
	newInstance := new(TotalDigit)
	totaldigit.GongCopyBasicFields(newInstance)
	return newInstance
}

func (union *Union) GongCopy() GongstructIF {
	newInstance := new(Union)
	union.GongCopyBasicFields(newInstance)
	return newInstance
}

func (whitespace *WhiteSpace) GongCopy() GongstructIF {
	newInstance := new(WhiteSpace)
	whitespace.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (all *All) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(all).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(all), uint64(stage.GetOrder(all)))
	return
}

func (annotation *Annotation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(annotation).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(annotation), uint64(stage.GetOrder(annotation)))
	return
}

func (attribute *Attribute) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute), uint64(stage.GetOrder(attribute)))
	return
}

func (attributegroup *AttributeGroup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attributegroup).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attributegroup), uint64(stage.GetOrder(attributegroup)))
	return
}

func (choice *Choice) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(choice).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(choice), uint64(stage.GetOrder(choice)))
	return
}

func (complexcontent *ComplexContent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complexcontent).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(complexcontent), uint64(stage.GetOrder(complexcontent)))
	return
}

func (complextype *ComplexType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complextype).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(complextype), uint64(stage.GetOrder(complextype)))
	return
}

func (documentation *Documentation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(documentation).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(documentation), uint64(stage.GetOrder(documentation)))
	return
}

func (element *Element) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(element).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(element), uint64(stage.GetOrder(element)))
	return
}

func (enumeration *Enumeration) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(enumeration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(enumeration), uint64(stage.GetOrder(enumeration)))
	return
}

func (extension *Extension) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(extension).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(extension), uint64(stage.GetOrder(extension)))
	return
}

func (group *Group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group), uint64(stage.GetOrder(group)))
	return
}

func (length *Length) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(length).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(length), uint64(stage.GetOrder(length)))
	return
}

func (maxinclusive *MaxInclusive) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(maxinclusive).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(maxinclusive), uint64(stage.GetOrder(maxinclusive)))
	return
}

func (maxlength *MaxLength) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(maxlength).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(maxlength), uint64(stage.GetOrder(maxlength)))
	return
}

func (mininclusive *MinInclusive) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mininclusive).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(mininclusive), uint64(stage.GetOrder(mininclusive)))
	return
}

func (minlength *MinLength) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(minlength).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(minlength), uint64(stage.GetOrder(minlength)))
	return
}

func (pattern *Pattern) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pattern).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pattern), uint64(stage.GetOrder(pattern)))
	return
}

func (restriction *Restriction) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(restriction).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(restriction), uint64(stage.GetOrder(restriction)))
	return
}

func (schema *Schema) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(schema).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(schema), uint64(stage.GetOrder(schema)))
	return
}

func (sequence *Sequence) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(sequence).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(sequence), uint64(stage.GetOrder(sequence)))
	return
}

func (simplecontent *SimpleContent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(simplecontent).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(simplecontent), uint64(stage.GetOrder(simplecontent)))
	return
}

func (simpletype *SimpleType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(simpletype).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(simpletype), uint64(stage.GetOrder(simpletype)))
	return
}

func (totaldigit *TotalDigit) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(totaldigit).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(totaldigit), uint64(stage.GetOrder(totaldigit)))
	return
}

func (union *Union) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(union).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(union), uint64(stage.GetOrder(union)))
	return
}

func (whitespace *WhiteSpace) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(whitespace).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(whitespace), uint64(stage.GetOrder(whitespace)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
		stage.Alls,
		stage.All_stagedOrder,
		stage.Alls_reference,
		&stage.Alls_referenceOrder,
		stage.Alls_instance,
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
		stage.Annotations,
		stage.Annotation_stagedOrder,
		stage.Annotations_reference,
		&stage.Annotations_referenceOrder,
		stage.Annotations_instance,
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
		stage.Attributes,
		stage.Attribute_stagedOrder,
		stage.Attributes_reference,
		&stage.Attributes_referenceOrder,
		stage.Attributes_instance,
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
		stage.AttributeGroups,
		stage.AttributeGroup_stagedOrder,
		stage.AttributeGroups_reference,
		&stage.AttributeGroups_referenceOrder,
		stage.AttributeGroups_instance,
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
		stage.Choices,
		stage.Choice_stagedOrder,
		stage.Choices_reference,
		&stage.Choices_referenceOrder,
		stage.Choices_instance,
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
		stage.ComplexContents,
		stage.ComplexContent_stagedOrder,
		stage.ComplexContents_reference,
		&stage.ComplexContents_referenceOrder,
		stage.ComplexContents_instance,
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
		stage.ComplexTypes,
		stage.ComplexType_stagedOrder,
		stage.ComplexTypes_reference,
		&stage.ComplexTypes_referenceOrder,
		stage.ComplexTypes_instance,
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
		stage.Documentations,
		stage.Documentation_stagedOrder,
		stage.Documentations_reference,
		&stage.Documentations_referenceOrder,
		stage.Documentations_instance,
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
		stage.Elements,
		stage.Element_stagedOrder,
		stage.Elements_reference,
		&stage.Elements_referenceOrder,
		stage.Elements_instance,
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
		stage.Enumerations,
		stage.Enumeration_stagedOrder,
		stage.Enumerations_reference,
		&stage.Enumerations_referenceOrder,
		stage.Enumerations_instance,
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
		stage.Extensions,
		stage.Extension_stagedOrder,
		stage.Extensions_reference,
		&stage.Extensions_referenceOrder,
		stage.Extensions_instance,
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
		stage.Groups,
		stage.Group_stagedOrder,
		stage.Groups_reference,
		&stage.Groups_referenceOrder,
		stage.Groups_instance,
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
		stage.Lengths,
		stage.Length_stagedOrder,
		stage.Lengths_reference,
		&stage.Lengths_referenceOrder,
		stage.Lengths_instance,
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
		stage.MaxInclusives,
		stage.MaxInclusive_stagedOrder,
		stage.MaxInclusives_reference,
		&stage.MaxInclusives_referenceOrder,
		stage.MaxInclusives_instance,
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
		stage.MaxLengths,
		stage.MaxLength_stagedOrder,
		stage.MaxLengths_reference,
		&stage.MaxLengths_referenceOrder,
		stage.MaxLengths_instance,
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
		stage.MinInclusives,
		stage.MinInclusive_stagedOrder,
		stage.MinInclusives_reference,
		&stage.MinInclusives_referenceOrder,
		stage.MinInclusives_instance,
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
		stage.MinLengths,
		stage.MinLength_stagedOrder,
		stage.MinLengths_reference,
		&stage.MinLengths_referenceOrder,
		stage.MinLengths_instance,
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
		stage.Patterns,
		stage.Pattern_stagedOrder,
		stage.Patterns_reference,
		&stage.Patterns_referenceOrder,
		stage.Patterns_instance,
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
		stage.Restrictions,
		stage.Restriction_stagedOrder,
		stage.Restrictions_reference,
		&stage.Restrictions_referenceOrder,
		stage.Restrictions_instance,
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
		stage.Schemas,
		stage.Schema_stagedOrder,
		stage.Schemas_reference,
		&stage.Schemas_referenceOrder,
		stage.Schemas_instance,
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
		stage.Sequences,
		stage.Sequence_stagedOrder,
		stage.Sequences_reference,
		&stage.Sequences_referenceOrder,
		stage.Sequences_instance,
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
		stage.SimpleContents,
		stage.SimpleContent_stagedOrder,
		stage.SimpleContents_reference,
		&stage.SimpleContents_referenceOrder,
		stage.SimpleContents_instance,
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
		stage.SimpleTypes,
		stage.SimpleType_stagedOrder,
		stage.SimpleTypes_reference,
		&stage.SimpleTypes_referenceOrder,
		stage.SimpleTypes_instance,
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
		stage.TotalDigits,
		stage.TotalDigit_stagedOrder,
		stage.TotalDigits_reference,
		&stage.TotalDigits_referenceOrder,
		stage.TotalDigits_instance,
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
		stage.Unions,
		stage.Union_stagedOrder,
		stage.Unions_reference,
		&stage.Unions_referenceOrder,
		stage.Unions_instance,
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
		stage.WhiteSpaces,
		stage.WhiteSpace_stagedOrder,
		stage.WhiteSpaces_reference,
		&stage.WhiteSpaces_referenceOrder,
		stage.WhiteSpaces_instance,
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
	stage.Alls_reference = make(map[*All]*All)
	stage.Alls_referenceOrder = make(map[*All]uint) // diff Unstage needs the reference order
	stage.Alls_instance = make(map[*All]*All)
	for instance := range stage.Alls {
		_copy := instance.GongCopy().(*All)
		stage.Alls_reference[instance] = _copy
		stage.Alls_instance[_copy] = instance
		stage.Alls_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Annotations_reference = make(map[*Annotation]*Annotation)
	stage.Annotations_referenceOrder = make(map[*Annotation]uint) // diff Unstage needs the reference order
	stage.Annotations_instance = make(map[*Annotation]*Annotation)
	for instance := range stage.Annotations {
		_copy := instance.GongCopy().(*Annotation)
		stage.Annotations_reference[instance] = _copy
		stage.Annotations_instance[_copy] = instance
		stage.Annotations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Attributes_reference = make(map[*Attribute]*Attribute)
	stage.Attributes_referenceOrder = make(map[*Attribute]uint) // diff Unstage needs the reference order
	stage.Attributes_instance = make(map[*Attribute]*Attribute)
	for instance := range stage.Attributes {
		_copy := instance.GongCopy().(*Attribute)
		stage.Attributes_reference[instance] = _copy
		stage.Attributes_instance[_copy] = instance
		stage.Attributes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AttributeGroups_reference = make(map[*AttributeGroup]*AttributeGroup)
	stage.AttributeGroups_referenceOrder = make(map[*AttributeGroup]uint) // diff Unstage needs the reference order
	stage.AttributeGroups_instance = make(map[*AttributeGroup]*AttributeGroup)
	for instance := range stage.AttributeGroups {
		_copy := instance.GongCopy().(*AttributeGroup)
		stage.AttributeGroups_reference[instance] = _copy
		stage.AttributeGroups_instance[_copy] = instance
		stage.AttributeGroups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Choices_reference = make(map[*Choice]*Choice)
	stage.Choices_referenceOrder = make(map[*Choice]uint) // diff Unstage needs the reference order
	stage.Choices_instance = make(map[*Choice]*Choice)
	for instance := range stage.Choices {
		_copy := instance.GongCopy().(*Choice)
		stage.Choices_reference[instance] = _copy
		stage.Choices_instance[_copy] = instance
		stage.Choices_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ComplexContents_reference = make(map[*ComplexContent]*ComplexContent)
	stage.ComplexContents_referenceOrder = make(map[*ComplexContent]uint) // diff Unstage needs the reference order
	stage.ComplexContents_instance = make(map[*ComplexContent]*ComplexContent)
	for instance := range stage.ComplexContents {
		_copy := instance.GongCopy().(*ComplexContent)
		stage.ComplexContents_reference[instance] = _copy
		stage.ComplexContents_instance[_copy] = instance
		stage.ComplexContents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ComplexTypes_reference = make(map[*ComplexType]*ComplexType)
	stage.ComplexTypes_referenceOrder = make(map[*ComplexType]uint) // diff Unstage needs the reference order
	stage.ComplexTypes_instance = make(map[*ComplexType]*ComplexType)
	for instance := range stage.ComplexTypes {
		_copy := instance.GongCopy().(*ComplexType)
		stage.ComplexTypes_reference[instance] = _copy
		stage.ComplexTypes_instance[_copy] = instance
		stage.ComplexTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Documentations_reference = make(map[*Documentation]*Documentation)
	stage.Documentations_referenceOrder = make(map[*Documentation]uint) // diff Unstage needs the reference order
	stage.Documentations_instance = make(map[*Documentation]*Documentation)
	for instance := range stage.Documentations {
		_copy := instance.GongCopy().(*Documentation)
		stage.Documentations_reference[instance] = _copy
		stage.Documentations_instance[_copy] = instance
		stage.Documentations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Elements_reference = make(map[*Element]*Element)
	stage.Elements_referenceOrder = make(map[*Element]uint) // diff Unstage needs the reference order
	stage.Elements_instance = make(map[*Element]*Element)
	for instance := range stage.Elements {
		_copy := instance.GongCopy().(*Element)
		stage.Elements_reference[instance] = _copy
		stage.Elements_instance[_copy] = instance
		stage.Elements_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Enumerations_reference = make(map[*Enumeration]*Enumeration)
	stage.Enumerations_referenceOrder = make(map[*Enumeration]uint) // diff Unstage needs the reference order
	stage.Enumerations_instance = make(map[*Enumeration]*Enumeration)
	for instance := range stage.Enumerations {
		_copy := instance.GongCopy().(*Enumeration)
		stage.Enumerations_reference[instance] = _copy
		stage.Enumerations_instance[_copy] = instance
		stage.Enumerations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Extensions_reference = make(map[*Extension]*Extension)
	stage.Extensions_referenceOrder = make(map[*Extension]uint) // diff Unstage needs the reference order
	stage.Extensions_instance = make(map[*Extension]*Extension)
	for instance := range stage.Extensions {
		_copy := instance.GongCopy().(*Extension)
		stage.Extensions_reference[instance] = _copy
		stage.Extensions_instance[_copy] = instance
		stage.Extensions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	stage.Groups_instance = make(map[*Group]*Group)
	for instance := range stage.Groups {
		_copy := instance.GongCopy().(*Group)
		stage.Groups_reference[instance] = _copy
		stage.Groups_instance[_copy] = instance
		stage.Groups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lengths_reference = make(map[*Length]*Length)
	stage.Lengths_referenceOrder = make(map[*Length]uint) // diff Unstage needs the reference order
	stage.Lengths_instance = make(map[*Length]*Length)
	for instance := range stage.Lengths {
		_copy := instance.GongCopy().(*Length)
		stage.Lengths_reference[instance] = _copy
		stage.Lengths_instance[_copy] = instance
		stage.Lengths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MaxInclusives_reference = make(map[*MaxInclusive]*MaxInclusive)
	stage.MaxInclusives_referenceOrder = make(map[*MaxInclusive]uint) // diff Unstage needs the reference order
	stage.MaxInclusives_instance = make(map[*MaxInclusive]*MaxInclusive)
	for instance := range stage.MaxInclusives {
		_copy := instance.GongCopy().(*MaxInclusive)
		stage.MaxInclusives_reference[instance] = _copy
		stage.MaxInclusives_instance[_copy] = instance
		stage.MaxInclusives_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MaxLengths_reference = make(map[*MaxLength]*MaxLength)
	stage.MaxLengths_referenceOrder = make(map[*MaxLength]uint) // diff Unstage needs the reference order
	stage.MaxLengths_instance = make(map[*MaxLength]*MaxLength)
	for instance := range stage.MaxLengths {
		_copy := instance.GongCopy().(*MaxLength)
		stage.MaxLengths_reference[instance] = _copy
		stage.MaxLengths_instance[_copy] = instance
		stage.MaxLengths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MinInclusives_reference = make(map[*MinInclusive]*MinInclusive)
	stage.MinInclusives_referenceOrder = make(map[*MinInclusive]uint) // diff Unstage needs the reference order
	stage.MinInclusives_instance = make(map[*MinInclusive]*MinInclusive)
	for instance := range stage.MinInclusives {
		_copy := instance.GongCopy().(*MinInclusive)
		stage.MinInclusives_reference[instance] = _copy
		stage.MinInclusives_instance[_copy] = instance
		stage.MinInclusives_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MinLengths_reference = make(map[*MinLength]*MinLength)
	stage.MinLengths_referenceOrder = make(map[*MinLength]uint) // diff Unstage needs the reference order
	stage.MinLengths_instance = make(map[*MinLength]*MinLength)
	for instance := range stage.MinLengths {
		_copy := instance.GongCopy().(*MinLength)
		stage.MinLengths_reference[instance] = _copy
		stage.MinLengths_instance[_copy] = instance
		stage.MinLengths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Patterns_reference = make(map[*Pattern]*Pattern)
	stage.Patterns_referenceOrder = make(map[*Pattern]uint) // diff Unstage needs the reference order
	stage.Patterns_instance = make(map[*Pattern]*Pattern)
	for instance := range stage.Patterns {
		_copy := instance.GongCopy().(*Pattern)
		stage.Patterns_reference[instance] = _copy
		stage.Patterns_instance[_copy] = instance
		stage.Patterns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Restrictions_reference = make(map[*Restriction]*Restriction)
	stage.Restrictions_referenceOrder = make(map[*Restriction]uint) // diff Unstage needs the reference order
	stage.Restrictions_instance = make(map[*Restriction]*Restriction)
	for instance := range stage.Restrictions {
		_copy := instance.GongCopy().(*Restriction)
		stage.Restrictions_reference[instance] = _copy
		stage.Restrictions_instance[_copy] = instance
		stage.Restrictions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Schemas_reference = make(map[*Schema]*Schema)
	stage.Schemas_referenceOrder = make(map[*Schema]uint) // diff Unstage needs the reference order
	stage.Schemas_instance = make(map[*Schema]*Schema)
	for instance := range stage.Schemas {
		_copy := instance.GongCopy().(*Schema)
		stage.Schemas_reference[instance] = _copy
		stage.Schemas_instance[_copy] = instance
		stage.Schemas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Sequences_reference = make(map[*Sequence]*Sequence)
	stage.Sequences_referenceOrder = make(map[*Sequence]uint) // diff Unstage needs the reference order
	stage.Sequences_instance = make(map[*Sequence]*Sequence)
	for instance := range stage.Sequences {
		_copy := instance.GongCopy().(*Sequence)
		stage.Sequences_reference[instance] = _copy
		stage.Sequences_instance[_copy] = instance
		stage.Sequences_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SimpleContents_reference = make(map[*SimpleContent]*SimpleContent)
	stage.SimpleContents_referenceOrder = make(map[*SimpleContent]uint) // diff Unstage needs the reference order
	stage.SimpleContents_instance = make(map[*SimpleContent]*SimpleContent)
	for instance := range stage.SimpleContents {
		_copy := instance.GongCopy().(*SimpleContent)
		stage.SimpleContents_reference[instance] = _copy
		stage.SimpleContents_instance[_copy] = instance
		stage.SimpleContents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SimpleTypes_reference = make(map[*SimpleType]*SimpleType)
	stage.SimpleTypes_referenceOrder = make(map[*SimpleType]uint) // diff Unstage needs the reference order
	stage.SimpleTypes_instance = make(map[*SimpleType]*SimpleType)
	for instance := range stage.SimpleTypes {
		_copy := instance.GongCopy().(*SimpleType)
		stage.SimpleTypes_reference[instance] = _copy
		stage.SimpleTypes_instance[_copy] = instance
		stage.SimpleTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TotalDigits_reference = make(map[*TotalDigit]*TotalDigit)
	stage.TotalDigits_referenceOrder = make(map[*TotalDigit]uint) // diff Unstage needs the reference order
	stage.TotalDigits_instance = make(map[*TotalDigit]*TotalDigit)
	for instance := range stage.TotalDigits {
		_copy := instance.GongCopy().(*TotalDigit)
		stage.TotalDigits_reference[instance] = _copy
		stage.TotalDigits_instance[_copy] = instance
		stage.TotalDigits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Unions_reference = make(map[*Union]*Union)
	stage.Unions_referenceOrder = make(map[*Union]uint) // diff Unstage needs the reference order
	stage.Unions_instance = make(map[*Union]*Union)
	for instance := range stage.Unions {
		_copy := instance.GongCopy().(*Union)
		stage.Unions_reference[instance] = _copy
		stage.Unions_instance[_copy] = instance
		stage.Unions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.WhiteSpaces_reference = make(map[*WhiteSpace]*WhiteSpace)
	stage.WhiteSpaces_referenceOrder = make(map[*WhiteSpace]uint) // diff Unstage needs the reference order
	stage.WhiteSpaces_instance = make(map[*WhiteSpace]*WhiteSpace)
	for instance := range stage.WhiteSpaces {
		_copy := instance.GongCopy().(*WhiteSpace)
		stage.WhiteSpaces_reference[instance] = _copy
		stage.WhiteSpaces_instance[_copy] = instance
		stage.WhiteSpaces_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Alls {
		reference := stage.Alls_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Annotations {
		reference := stage.Annotations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Attributes {
		reference := stage.Attributes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AttributeGroups {
		reference := stage.AttributeGroups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Choices {
		reference := stage.Choices_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ComplexContents {
		reference := stage.ComplexContents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ComplexTypes {
		reference := stage.ComplexTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Documentations {
		reference := stage.Documentations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Elements {
		reference := stage.Elements_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Enumerations {
		reference := stage.Enumerations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Extensions {
		reference := stage.Extensions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Groups {
		reference := stage.Groups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lengths {
		reference := stage.Lengths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MaxInclusives {
		reference := stage.MaxInclusives_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MaxLengths {
		reference := stage.MaxLengths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MinInclusives {
		reference := stage.MinInclusives_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MinLengths {
		reference := stage.MinLengths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Patterns {
		reference := stage.Patterns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Restrictions {
		reference := stage.Restrictions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Schemas {
		reference := stage.Schemas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Sequences {
		reference := stage.Sequences_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SimpleContents {
		reference := stage.SimpleContents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SimpleTypes {
		reference := stage.SimpleTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TotalDigits {
		reference := stage.TotalDigits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Unions {
		reference := stage.Unions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.WhiteSpaces {
		reference := stage.WhiteSpaces_reference[instance]
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
func (all *All) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.All_stagedOrder[all]; ok {
		return order
	}
	if order, ok := stage.Alls_referenceOrder[all]; ok {
		return order
	} else {
		log.Printf("instance %p of type All was not staged and does not have a reference order", all)
		return 0
	}
}

func (annotation *Annotation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Annotation_stagedOrder[annotation]; ok {
		return order
	}
	if order, ok := stage.Annotations_referenceOrder[annotation]; ok {
		return order
	} else {
		log.Printf("instance %p of type Annotation was not staged and does not have a reference order", annotation)
		return 0
	}
}

func (attribute *Attribute) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Attribute_stagedOrder[attribute]; ok {
		return order
	}
	if order, ok := stage.Attributes_referenceOrder[attribute]; ok {
		return order
	} else {
		log.Printf("instance %p of type Attribute was not staged and does not have a reference order", attribute)
		return 0
	}
}

func (attributegroup *AttributeGroup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AttributeGroup_stagedOrder[attributegroup]; ok {
		return order
	}
	if order, ok := stage.AttributeGroups_referenceOrder[attributegroup]; ok {
		return order
	} else {
		log.Printf("instance %p of type AttributeGroup was not staged and does not have a reference order", attributegroup)
		return 0
	}
}

func (choice *Choice) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Choice_stagedOrder[choice]; ok {
		return order
	}
	if order, ok := stage.Choices_referenceOrder[choice]; ok {
		return order
	} else {
		log.Printf("instance %p of type Choice was not staged and does not have a reference order", choice)
		return 0
	}
}

func (complexcontent *ComplexContent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ComplexContent_stagedOrder[complexcontent]; ok {
		return order
	}
	if order, ok := stage.ComplexContents_referenceOrder[complexcontent]; ok {
		return order
	} else {
		log.Printf("instance %p of type ComplexContent was not staged and does not have a reference order", complexcontent)
		return 0
	}
}

func (complextype *ComplexType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ComplexType_stagedOrder[complextype]; ok {
		return order
	}
	if order, ok := stage.ComplexTypes_referenceOrder[complextype]; ok {
		return order
	} else {
		log.Printf("instance %p of type ComplexType was not staged and does not have a reference order", complextype)
		return 0
	}
}

func (documentation *Documentation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Documentation_stagedOrder[documentation]; ok {
		return order
	}
	if order, ok := stage.Documentations_referenceOrder[documentation]; ok {
		return order
	} else {
		log.Printf("instance %p of type Documentation was not staged and does not have a reference order", documentation)
		return 0
	}
}

func (element *Element) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Element_stagedOrder[element]; ok {
		return order
	}
	if order, ok := stage.Elements_referenceOrder[element]; ok {
		return order
	} else {
		log.Printf("instance %p of type Element was not staged and does not have a reference order", element)
		return 0
	}
}

func (enumeration *Enumeration) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Enumeration_stagedOrder[enumeration]; ok {
		return order
	}
	if order, ok := stage.Enumerations_referenceOrder[enumeration]; ok {
		return order
	} else {
		log.Printf("instance %p of type Enumeration was not staged and does not have a reference order", enumeration)
		return 0
	}
}

func (extension *Extension) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Extension_stagedOrder[extension]; ok {
		return order
	}
	if order, ok := stage.Extensions_referenceOrder[extension]; ok {
		return order
	} else {
		log.Printf("instance %p of type Extension was not staged and does not have a reference order", extension)
		return 0
	}
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_stagedOrder[group]; ok {
		return order
	}
	if order, ok := stage.Groups_referenceOrder[group]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group was not staged and does not have a reference order", group)
		return 0
	}
}

func (length *Length) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Length_stagedOrder[length]; ok {
		return order
	}
	if order, ok := stage.Lengths_referenceOrder[length]; ok {
		return order
	} else {
		log.Printf("instance %p of type Length was not staged and does not have a reference order", length)
		return 0
	}
}

func (maxinclusive *MaxInclusive) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MaxInclusive_stagedOrder[maxinclusive]; ok {
		return order
	}
	if order, ok := stage.MaxInclusives_referenceOrder[maxinclusive]; ok {
		return order
	} else {
		log.Printf("instance %p of type MaxInclusive was not staged and does not have a reference order", maxinclusive)
		return 0
	}
}

func (maxlength *MaxLength) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MaxLength_stagedOrder[maxlength]; ok {
		return order
	}
	if order, ok := stage.MaxLengths_referenceOrder[maxlength]; ok {
		return order
	} else {
		log.Printf("instance %p of type MaxLength was not staged and does not have a reference order", maxlength)
		return 0
	}
}

func (mininclusive *MinInclusive) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MinInclusive_stagedOrder[mininclusive]; ok {
		return order
	}
	if order, ok := stage.MinInclusives_referenceOrder[mininclusive]; ok {
		return order
	} else {
		log.Printf("instance %p of type MinInclusive was not staged and does not have a reference order", mininclusive)
		return 0
	}
}

func (minlength *MinLength) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MinLength_stagedOrder[minlength]; ok {
		return order
	}
	if order, ok := stage.MinLengths_referenceOrder[minlength]; ok {
		return order
	} else {
		log.Printf("instance %p of type MinLength was not staged and does not have a reference order", minlength)
		return 0
	}
}

func (pattern *Pattern) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Pattern_stagedOrder[pattern]; ok {
		return order
	}
	if order, ok := stage.Patterns_referenceOrder[pattern]; ok {
		return order
	} else {
		log.Printf("instance %p of type Pattern was not staged and does not have a reference order", pattern)
		return 0
	}
}

func (restriction *Restriction) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Restriction_stagedOrder[restriction]; ok {
		return order
	}
	if order, ok := stage.Restrictions_referenceOrder[restriction]; ok {
		return order
	} else {
		log.Printf("instance %p of type Restriction was not staged and does not have a reference order", restriction)
		return 0
	}
}

func (schema *Schema) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Schema_stagedOrder[schema]; ok {
		return order
	}
	if order, ok := stage.Schemas_referenceOrder[schema]; ok {
		return order
	} else {
		log.Printf("instance %p of type Schema was not staged and does not have a reference order", schema)
		return 0
	}
}

func (sequence *Sequence) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Sequence_stagedOrder[sequence]; ok {
		return order
	}
	if order, ok := stage.Sequences_referenceOrder[sequence]; ok {
		return order
	} else {
		log.Printf("instance %p of type Sequence was not staged and does not have a reference order", sequence)
		return 0
	}
}

func (simplecontent *SimpleContent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SimpleContent_stagedOrder[simplecontent]; ok {
		return order
	}
	if order, ok := stage.SimpleContents_referenceOrder[simplecontent]; ok {
		return order
	} else {
		log.Printf("instance %p of type SimpleContent was not staged and does not have a reference order", simplecontent)
		return 0
	}
}

func (simpletype *SimpleType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SimpleType_stagedOrder[simpletype]; ok {
		return order
	}
	if order, ok := stage.SimpleTypes_referenceOrder[simpletype]; ok {
		return order
	} else {
		log.Printf("instance %p of type SimpleType was not staged and does not have a reference order", simpletype)
		return 0
	}
}

func (totaldigit *TotalDigit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TotalDigit_stagedOrder[totaldigit]; ok {
		return order
	}
	if order, ok := stage.TotalDigits_referenceOrder[totaldigit]; ok {
		return order
	} else {
		log.Printf("instance %p of type TotalDigit was not staged and does not have a reference order", totaldigit)
		return 0
	}
}

func (union *Union) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Union_stagedOrder[union]; ok {
		return order
	}
	if order, ok := stage.Unions_referenceOrder[union]; ok {
		return order
	} else {
		log.Printf("instance %p of type Union was not staged and does not have a reference order", union)
		return 0
	}
}

func (whitespace *WhiteSpace) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.WhiteSpace_stagedOrder[whitespace]; ok {
		return order
	}
	if order, ok := stage.WhiteSpaces_referenceOrder[whitespace]; ok {
		return order
	} else {
		log.Printf("instance %p of type WhiteSpace was not staged and does not have a reference order", whitespace)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (all *All) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", all.GongGetGongstructName(), all.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (all *All) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", all.GongGetGongstructName(), all.GongGetOrder(stage))
}

func (annotation *Annotation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", annotation.GongGetGongstructName(), annotation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (annotation *Annotation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", annotation.GongGetGongstructName(), annotation.GongGetOrder(stage))
}

func (attribute *Attribute) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute.GongGetGongstructName(), attribute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute *Attribute) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute.GongGetGongstructName(), attribute.GongGetOrder(stage))
}

func (attributegroup *AttributeGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributegroup.GongGetGongstructName(), attributegroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributegroup *AttributeGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributegroup.GongGetGongstructName(), attributegroup.GongGetOrder(stage))
}

func (choice *Choice) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", choice.GongGetGongstructName(), choice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (choice *Choice) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", choice.GongGetGongstructName(), choice.GongGetOrder(stage))
}

func (complexcontent *ComplexContent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexcontent.GongGetGongstructName(), complexcontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexcontent *ComplexContent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complexcontent.GongGetGongstructName(), complexcontent.GongGetOrder(stage))
}

func (complextype *ComplexType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complextype.GongGetGongstructName(), complextype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complextype *ComplexType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", complextype.GongGetGongstructName(), complextype.GongGetOrder(stage))
}

func (documentation *Documentation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentation.GongGetGongstructName(), documentation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (documentation *Documentation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", documentation.GongGetGongstructName(), documentation.GongGetOrder(stage))
}

func (element *Element) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", element.GongGetGongstructName(), element.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (element *Element) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", element.GongGetGongstructName(), element.GongGetOrder(stage))
}

func (enumeration *Enumeration) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enumeration.GongGetGongstructName(), enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enumeration *Enumeration) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enumeration.GongGetGongstructName(), enumeration.GongGetOrder(stage))
}

func (extension *Extension) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extension.GongGetGongstructName(), extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extension *Extension) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extension.GongGetGongstructName(), extension.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (length *Length) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", length.GongGetGongstructName(), length.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (length *Length) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", length.GongGetGongstructName(), length.GongGetOrder(stage))
}

func (maxinclusive *MaxInclusive) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxinclusive.GongGetGongstructName(), maxinclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxinclusive *MaxInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxinclusive.GongGetGongstructName(), maxinclusive.GongGetOrder(stage))
}

func (maxlength *MaxLength) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxlength.GongGetGongstructName(), maxlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxlength *MaxLength) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", maxlength.GongGetGongstructName(), maxlength.GongGetOrder(stage))
}

func (mininclusive *MinInclusive) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mininclusive.GongGetGongstructName(), mininclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mininclusive *MinInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mininclusive.GongGetGongstructName(), mininclusive.GongGetOrder(stage))
}

func (minlength *MinLength) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", minlength.GongGetGongstructName(), minlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (minlength *MinLength) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", minlength.GongGetGongstructName(), minlength.GongGetOrder(stage))
}

func (pattern *Pattern) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pattern.GongGetGongstructName(), pattern.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pattern *Pattern) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pattern.GongGetGongstructName(), pattern.GongGetOrder(stage))
}

func (restriction *Restriction) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", restriction.GongGetGongstructName(), restriction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (restriction *Restriction) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", restriction.GongGetGongstructName(), restriction.GongGetOrder(stage))
}

func (schema *Schema) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", schema.GongGetGongstructName(), schema.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (schema *Schema) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", schema.GongGetGongstructName(), schema.GongGetOrder(stage))
}

func (sequence *Sequence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sequence.GongGetGongstructName(), sequence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sequence *Sequence) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sequence.GongGetGongstructName(), sequence.GongGetOrder(stage))
}

func (simplecontent *SimpleContent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simplecontent.GongGetGongstructName(), simplecontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simplecontent *SimpleContent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simplecontent.GongGetGongstructName(), simplecontent.GongGetOrder(stage))
}

func (simpletype *SimpleType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simpletype.GongGetGongstructName(), simpletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simpletype *SimpleType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", simpletype.GongGetGongstructName(), simpletype.GongGetOrder(stage))
}

func (totaldigit *TotalDigit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", totaldigit.GongGetGongstructName(), totaldigit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (totaldigit *TotalDigit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", totaldigit.GongGetGongstructName(), totaldigit.GongGetOrder(stage))
}

func (union *Union) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", union.GongGetGongstructName(), union.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (union *Union) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", union.GongGetGongstructName(), union.GongGetOrder(stage))
}

func (whitespace *WhiteSpace) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", whitespace.GongGetGongstructName(), whitespace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (whitespace *WhiteSpace) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", whitespace.GongGetGongstructName(), whitespace.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (all *All) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", all.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "All")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(all.Name))
	return
}

func (annotation *Annotation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", annotation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Annotation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(annotation.Name))
	return
}

func (attribute *Attribute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attribute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute.Name))
	return
}

func (attributegroup *AttributeGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributegroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attributegroup.Name))
	return
}

func (choice *Choice) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", choice.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Choice")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(choice.Name))
	return
}

func (complexcontent *ComplexContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexcontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(complexcontent.Name))
	return
}

func (complextype *ComplexType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complextype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(complextype.Name))
	return
}

func (documentation *Documentation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Documentation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(documentation.Name))
	return
}

func (element *Element) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", element.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(element.Name))
	return
}

func (enumeration *Enumeration) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Enumeration")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(enumeration.Name))
	return
}

func (extension *Extension) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extension.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extension")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(extension.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group.Name))
	return
}

func (length *Length) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", length.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Length")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(length.Name))
	return
}

func (maxinclusive *MaxInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxinclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(maxinclusive.Name))
	return
}

func (maxlength *MaxLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(maxlength.Name))
	return
}

func (mininclusive *MinInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mininclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(mininclusive.Name))
	return
}

func (minlength *MinLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", minlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(minlength.Name))
	return
}

func (pattern *Pattern) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pattern.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pattern")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pattern.Name))
	return
}

func (restriction *Restriction) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", restriction.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Restriction")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(restriction.Name))
	return
}

func (schema *Schema) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", schema.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Schema")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(schema.Name))
	return
}

func (sequence *Sequence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sequence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sequence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(sequence.Name))
	return
}

func (simplecontent *SimpleContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simplecontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(simplecontent.Name))
	return
}

func (simpletype *SimpleType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simpletype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(simpletype.Name))
	return
}

func (totaldigit *TotalDigit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", totaldigit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TotalDigit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(totaldigit.Name))
	return
}

func (union *Union) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", union.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Union")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(union.Name))
	return
}

func (whitespace *WhiteSpace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", whitespace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "WhiteSpace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(whitespace.Name))
	return
}

// insertion point for unstaging
func (all *All) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", all.GongGetReferenceIdentifier(stage))
	return
}

func (annotation *Annotation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", annotation.GongGetReferenceIdentifier(stage))
	return
}

func (attribute *Attribute) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute.GongGetReferenceIdentifier(stage))
	return
}

func (attributegroup *AttributeGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributegroup.GongGetReferenceIdentifier(stage))
	return
}

func (choice *Choice) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", choice.GongGetReferenceIdentifier(stage))
	return
}

func (complexcontent *ComplexContent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexcontent.GongGetReferenceIdentifier(stage))
	return
}

func (complextype *ComplexType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complextype.GongGetReferenceIdentifier(stage))
	return
}

func (documentation *Documentation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentation.GongGetReferenceIdentifier(stage))
	return
}

func (element *Element) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", element.GongGetReferenceIdentifier(stage))
	return
}

func (enumeration *Enumeration) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (extension *Extension) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extension.GongGetReferenceIdentifier(stage))
	return
}

func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}

func (length *Length) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", length.GongGetReferenceIdentifier(stage))
	return
}

func (maxinclusive *MaxInclusive) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxinclusive.GongGetReferenceIdentifier(stage))
	return
}

func (maxlength *MaxLength) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxlength.GongGetReferenceIdentifier(stage))
	return
}

func (mininclusive *MinInclusive) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mininclusive.GongGetReferenceIdentifier(stage))
	return
}

func (minlength *MinLength) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", minlength.GongGetReferenceIdentifier(stage))
	return
}

func (pattern *Pattern) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pattern.GongGetReferenceIdentifier(stage))
	return
}

func (restriction *Restriction) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", restriction.GongGetReferenceIdentifier(stage))
	return
}

func (schema *Schema) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", schema.GongGetReferenceIdentifier(stage))
	return
}

func (sequence *Sequence) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sequence.GongGetReferenceIdentifier(stage))
	return
}

func (simplecontent *SimpleContent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simplecontent.GongGetReferenceIdentifier(stage))
	return
}

func (simpletype *SimpleType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simpletype.GongGetReferenceIdentifier(stage))
	return
}

func (totaldigit *TotalDigit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", totaldigit.GongGetReferenceIdentifier(stage))
	return
}

func (union *Union) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", union.GongGetReferenceIdentifier(stage))
	return
}

func (whitespace *WhiteSpace) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", whitespace.GongGetReferenceIdentifier(stage))
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
