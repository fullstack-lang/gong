// generated code - do not edit
package models

import (
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
	all.CopyBasicFields(newInstance)
	return newInstance
}

func (annotation *Annotation) GongCopy() GongstructIF {
	newInstance := new(Annotation)
	annotation.CopyBasicFields(newInstance)
	return newInstance
}

func (attribute *Attribute) GongCopy() GongstructIF {
	newInstance := new(Attribute)
	attribute.CopyBasicFields(newInstance)
	return newInstance
}

func (attributegroup *AttributeGroup) GongCopy() GongstructIF {
	newInstance := new(AttributeGroup)
	attributegroup.CopyBasicFields(newInstance)
	return newInstance
}

func (choice *Choice) GongCopy() GongstructIF {
	newInstance := new(Choice)
	choice.CopyBasicFields(newInstance)
	return newInstance
}

func (complexcontent *ComplexContent) GongCopy() GongstructIF {
	newInstance := new(ComplexContent)
	complexcontent.CopyBasicFields(newInstance)
	return newInstance
}

func (complextype *ComplexType) GongCopy() GongstructIF {
	newInstance := new(ComplexType)
	complextype.CopyBasicFields(newInstance)
	return newInstance
}

func (documentation *Documentation) GongCopy() GongstructIF {
	newInstance := new(Documentation)
	documentation.CopyBasicFields(newInstance)
	return newInstance
}

func (element *Element) GongCopy() GongstructIF {
	newInstance := new(Element)
	element.CopyBasicFields(newInstance)
	return newInstance
}

func (enumeration *Enumeration) GongCopy() GongstructIF {
	newInstance := new(Enumeration)
	enumeration.CopyBasicFields(newInstance)
	return newInstance
}

func (extension *Extension) GongCopy() GongstructIF {
	newInstance := new(Extension)
	extension.CopyBasicFields(newInstance)
	return newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := new(Group)
	group.CopyBasicFields(newInstance)
	return newInstance
}

func (length *Length) GongCopy() GongstructIF {
	newInstance := new(Length)
	length.CopyBasicFields(newInstance)
	return newInstance
}

func (maxinclusive *MaxInclusive) GongCopy() GongstructIF {
	newInstance := new(MaxInclusive)
	maxinclusive.CopyBasicFields(newInstance)
	return newInstance
}

func (maxlength *MaxLength) GongCopy() GongstructIF {
	newInstance := new(MaxLength)
	maxlength.CopyBasicFields(newInstance)
	return newInstance
}

func (mininclusive *MinInclusive) GongCopy() GongstructIF {
	newInstance := new(MinInclusive)
	mininclusive.CopyBasicFields(newInstance)
	return newInstance
}

func (minlength *MinLength) GongCopy() GongstructIF {
	newInstance := new(MinLength)
	minlength.CopyBasicFields(newInstance)
	return newInstance
}

func (pattern *Pattern) GongCopy() GongstructIF {
	newInstance := new(Pattern)
	pattern.CopyBasicFields(newInstance)
	return newInstance
}

func (restriction *Restriction) GongCopy() GongstructIF {
	newInstance := new(Restriction)
	restriction.CopyBasicFields(newInstance)
	return newInstance
}

func (schema *Schema) GongCopy() GongstructIF {
	newInstance := new(Schema)
	schema.CopyBasicFields(newInstance)
	return newInstance
}

func (sequence *Sequence) GongCopy() GongstructIF {
	newInstance := new(Sequence)
	sequence.CopyBasicFields(newInstance)
	return newInstance
}

func (simplecontent *SimpleContent) GongCopy() GongstructIF {
	newInstance := new(SimpleContent)
	simplecontent.CopyBasicFields(newInstance)
	return newInstance
}

func (simpletype *SimpleType) GongCopy() GongstructIF {
	newInstance := new(SimpleType)
	simpletype.CopyBasicFields(newInstance)
	return newInstance
}

func (totaldigit *TotalDigit) GongCopy() GongstructIF {
	newInstance := new(TotalDigit)
	totaldigit.CopyBasicFields(newInstance)
	return newInstance
}

func (union *Union) GongCopy() GongstructIF {
	newInstance := new(Union)
	union.CopyBasicFields(newInstance)
	return newInstance
}

func (whitespace *WhiteSpace) GongCopy() GongstructIF {
	newInstance := new(WhiteSpace)
	whitespace.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (all *All) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(all).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(all), uint64(GetOrderPointerGongstruct(stage, all)))
	return
}

func (annotation *Annotation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(annotation).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(annotation), uint64(GetOrderPointerGongstruct(stage, annotation)))
	return
}

func (attribute *Attribute) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(attribute), uint64(GetOrderPointerGongstruct(stage, attribute)))
	return
}

func (attributegroup *AttributeGroup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attributegroup).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(attributegroup), uint64(GetOrderPointerGongstruct(stage, attributegroup)))
	return
}

func (choice *Choice) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(choice).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(choice), uint64(GetOrderPointerGongstruct(stage, choice)))
	return
}

func (complexcontent *ComplexContent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complexcontent).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(complexcontent), uint64(GetOrderPointerGongstruct(stage, complexcontent)))
	return
}

func (complextype *ComplexType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(complextype).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(complextype), uint64(GetOrderPointerGongstruct(stage, complextype)))
	return
}

func (documentation *Documentation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(documentation).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(documentation), uint64(GetOrderPointerGongstruct(stage, documentation)))
	return
}

func (element *Element) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(element).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(element), uint64(GetOrderPointerGongstruct(stage, element)))
	return
}

func (enumeration *Enumeration) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(enumeration).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(enumeration), uint64(GetOrderPointerGongstruct(stage, enumeration)))
	return
}

func (extension *Extension) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(extension).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(extension), uint64(GetOrderPointerGongstruct(stage, extension)))
	return
}

func (group *Group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(group), uint64(GetOrderPointerGongstruct(stage, group)))
	return
}

func (length *Length) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(length).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(length), uint64(GetOrderPointerGongstruct(stage, length)))
	return
}

func (maxinclusive *MaxInclusive) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(maxinclusive).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(maxinclusive), uint64(GetOrderPointerGongstruct(stage, maxinclusive)))
	return
}

func (maxlength *MaxLength) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(maxlength).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(maxlength), uint64(GetOrderPointerGongstruct(stage, maxlength)))
	return
}

func (mininclusive *MinInclusive) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mininclusive).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(mininclusive), uint64(GetOrderPointerGongstruct(stage, mininclusive)))
	return
}

func (minlength *MinLength) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(minlength).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(minlength), uint64(GetOrderPointerGongstruct(stage, minlength)))
	return
}

func (pattern *Pattern) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pattern).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(pattern), uint64(GetOrderPointerGongstruct(stage, pattern)))
	return
}

func (restriction *Restriction) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(restriction).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(restriction), uint64(GetOrderPointerGongstruct(stage, restriction)))
	return
}

func (schema *Schema) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(schema).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(schema), uint64(GetOrderPointerGongstruct(stage, schema)))
	return
}

func (sequence *Sequence) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(sequence).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(sequence), uint64(GetOrderPointerGongstruct(stage, sequence)))
	return
}

func (simplecontent *SimpleContent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(simplecontent).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(simplecontent), uint64(GetOrderPointerGongstruct(stage, simplecontent)))
	return
}

func (simpletype *SimpleType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(simpletype).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(simpletype), uint64(GetOrderPointerGongstruct(stage, simpletype)))
	return
}

func (totaldigit *TotalDigit) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(totaldigit).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(totaldigit), uint64(GetOrderPointerGongstruct(stage, totaldigit)))
	return
}

func (union *Union) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(union).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(union), uint64(GetOrderPointerGongstruct(stage, union)))
	return
}

func (whitespace *WhiteSpace) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(whitespace).(interface{ GongUUIDCustom() string }); ok {
		return __gong__.GongUUIDCustom()
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(whitespace), uint64(GetOrderPointerGongstruct(stage, whitespace)))
	return
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
	var alls_newInstances []*All
	var alls_deletedInstances []*All

	// parse all staged instances and check if they have a reference
	for all := range stage.Alls {
		if ref, ok := stage.Alls_reference[all]; !ok {
			alls_newInstances = append(alls_newInstances, all)
			newInstancesSlice = append(newInstancesSlice, all.GongMarshallIdentifier(stage))
			if stage.Alls_referenceOrder == nil {
				stage.Alls_referenceOrder = make(map[*All]uint)
			}
			stage.Alls_referenceOrder[all] = stage.All_stagedOrder[all]
			newInstancesReverseSlice = append(newInstancesReverseSlice, all.GongMarshallUnstaging(stage))
			// delete(stage.Alls_referenceOrder, all)
			fieldInitializers, pointersInitializations := all.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.All_stagedOrder[ref] = stage.All_stagedOrder[all]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := all.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, all)
			// delete(stage.All_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", all.GetName())
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
	for _, ref := range stage.Alls_reference {
		instance := stage.Alls_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Alls[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			alls_deletedInstances = append(alls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(alls_newInstances)
	lenDeletedInstances += len(alls_deletedInstances)
	var annotations_newInstances []*Annotation
	var annotations_deletedInstances []*Annotation

	// parse all staged instances and check if they have a reference
	for annotation := range stage.Annotations {
		if ref, ok := stage.Annotations_reference[annotation]; !ok {
			annotations_newInstances = append(annotations_newInstances, annotation)
			newInstancesSlice = append(newInstancesSlice, annotation.GongMarshallIdentifier(stage))
			if stage.Annotations_referenceOrder == nil {
				stage.Annotations_referenceOrder = make(map[*Annotation]uint)
			}
			stage.Annotations_referenceOrder[annotation] = stage.Annotation_stagedOrder[annotation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, annotation.GongMarshallUnstaging(stage))
			// delete(stage.Annotations_referenceOrder, annotation)
			fieldInitializers, pointersInitializations := annotation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Annotation_stagedOrder[ref] = stage.Annotation_stagedOrder[annotation]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := annotation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, annotation)
			// delete(stage.Annotation_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", annotation.GetName())
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
	for _, ref := range stage.Annotations_reference {
		instance := stage.Annotations_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Annotations[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			annotations_deletedInstances = append(annotations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(annotations_newInstances)
	lenDeletedInstances += len(annotations_deletedInstances)
	var attributes_newInstances []*Attribute
	var attributes_deletedInstances []*Attribute

	// parse all staged instances and check if they have a reference
	for attribute := range stage.Attributes {
		if ref, ok := stage.Attributes_reference[attribute]; !ok {
			attributes_newInstances = append(attributes_newInstances, attribute)
			newInstancesSlice = append(newInstancesSlice, attribute.GongMarshallIdentifier(stage))
			if stage.Attributes_referenceOrder == nil {
				stage.Attributes_referenceOrder = make(map[*Attribute]uint)
			}
			stage.Attributes_referenceOrder[attribute] = stage.Attribute_stagedOrder[attribute]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute.GongMarshallUnstaging(stage))
			// delete(stage.Attributes_referenceOrder, attribute)
			fieldInitializers, pointersInitializations := attribute.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Attribute_stagedOrder[ref] = stage.Attribute_stagedOrder[attribute]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := attribute.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute)
			// delete(stage.Attribute_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute.GetName())
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
	for _, ref := range stage.Attributes_reference {
		instance := stage.Attributes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Attributes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			attributes_deletedInstances = append(attributes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attributes_newInstances)
	lenDeletedInstances += len(attributes_deletedInstances)
	var attributegroups_newInstances []*AttributeGroup
	var attributegroups_deletedInstances []*AttributeGroup

	// parse all staged instances and check if they have a reference
	for attributegroup := range stage.AttributeGroups {
		if ref, ok := stage.AttributeGroups_reference[attributegroup]; !ok {
			attributegroups_newInstances = append(attributegroups_newInstances, attributegroup)
			newInstancesSlice = append(newInstancesSlice, attributegroup.GongMarshallIdentifier(stage))
			if stage.AttributeGroups_referenceOrder == nil {
				stage.AttributeGroups_referenceOrder = make(map[*AttributeGroup]uint)
			}
			stage.AttributeGroups_referenceOrder[attributegroup] = stage.AttributeGroup_stagedOrder[attributegroup]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attributegroup.GongMarshallUnstaging(stage))
			// delete(stage.AttributeGroups_referenceOrder, attributegroup)
			fieldInitializers, pointersInitializations := attributegroup.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AttributeGroup_stagedOrder[ref] = stage.AttributeGroup_stagedOrder[attributegroup]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := attributegroup.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attributegroup)
			// delete(stage.AttributeGroup_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attributegroup.GetName())
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
	for _, ref := range stage.AttributeGroups_reference {
		instance := stage.AttributeGroups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AttributeGroups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			attributegroups_deletedInstances = append(attributegroups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attributegroups_newInstances)
	lenDeletedInstances += len(attributegroups_deletedInstances)
	var choices_newInstances []*Choice
	var choices_deletedInstances []*Choice

	// parse all staged instances and check if they have a reference
	for choice := range stage.Choices {
		if ref, ok := stage.Choices_reference[choice]; !ok {
			choices_newInstances = append(choices_newInstances, choice)
			newInstancesSlice = append(newInstancesSlice, choice.GongMarshallIdentifier(stage))
			if stage.Choices_referenceOrder == nil {
				stage.Choices_referenceOrder = make(map[*Choice]uint)
			}
			stage.Choices_referenceOrder[choice] = stage.Choice_stagedOrder[choice]
			newInstancesReverseSlice = append(newInstancesReverseSlice, choice.GongMarshallUnstaging(stage))
			// delete(stage.Choices_referenceOrder, choice)
			fieldInitializers, pointersInitializations := choice.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Choice_stagedOrder[ref] = stage.Choice_stagedOrder[choice]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := choice.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, choice)
			// delete(stage.Choice_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", choice.GetName())
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
	for _, ref := range stage.Choices_reference {
		instance := stage.Choices_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Choices[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			choices_deletedInstances = append(choices_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(choices_newInstances)
	lenDeletedInstances += len(choices_deletedInstances)
	var complexcontents_newInstances []*ComplexContent
	var complexcontents_deletedInstances []*ComplexContent

	// parse all staged instances and check if they have a reference
	for complexcontent := range stage.ComplexContents {
		if ref, ok := stage.ComplexContents_reference[complexcontent]; !ok {
			complexcontents_newInstances = append(complexcontents_newInstances, complexcontent)
			newInstancesSlice = append(newInstancesSlice, complexcontent.GongMarshallIdentifier(stage))
			if stage.ComplexContents_referenceOrder == nil {
				stage.ComplexContents_referenceOrder = make(map[*ComplexContent]uint)
			}
			stage.ComplexContents_referenceOrder[complexcontent] = stage.ComplexContent_stagedOrder[complexcontent]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complexcontent.GongMarshallUnstaging(stage))
			// delete(stage.ComplexContents_referenceOrder, complexcontent)
			fieldInitializers, pointersInitializations := complexcontent.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ComplexContent_stagedOrder[ref] = stage.ComplexContent_stagedOrder[complexcontent]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := complexcontent.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complexcontent)
			// delete(stage.ComplexContent_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", complexcontent.GetName())
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
	for _, ref := range stage.ComplexContents_reference {
		instance := stage.ComplexContents_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ComplexContents[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			complexcontents_deletedInstances = append(complexcontents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complexcontents_newInstances)
	lenDeletedInstances += len(complexcontents_deletedInstances)
	var complextypes_newInstances []*ComplexType
	var complextypes_deletedInstances []*ComplexType

	// parse all staged instances and check if they have a reference
	for complextype := range stage.ComplexTypes {
		if ref, ok := stage.ComplexTypes_reference[complextype]; !ok {
			complextypes_newInstances = append(complextypes_newInstances, complextype)
			newInstancesSlice = append(newInstancesSlice, complextype.GongMarshallIdentifier(stage))
			if stage.ComplexTypes_referenceOrder == nil {
				stage.ComplexTypes_referenceOrder = make(map[*ComplexType]uint)
			}
			stage.ComplexTypes_referenceOrder[complextype] = stage.ComplexType_stagedOrder[complextype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, complextype.GongMarshallUnstaging(stage))
			// delete(stage.ComplexTypes_referenceOrder, complextype)
			fieldInitializers, pointersInitializations := complextype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ComplexType_stagedOrder[ref] = stage.ComplexType_stagedOrder[complextype]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := complextype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, complextype)
			// delete(stage.ComplexType_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", complextype.GetName())
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
	for _, ref := range stage.ComplexTypes_reference {
		instance := stage.ComplexTypes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ComplexTypes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			complextypes_deletedInstances = append(complextypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(complextypes_newInstances)
	lenDeletedInstances += len(complextypes_deletedInstances)
	var documentations_newInstances []*Documentation
	var documentations_deletedInstances []*Documentation

	// parse all staged instances and check if they have a reference
	for documentation := range stage.Documentations {
		if ref, ok := stage.Documentations_reference[documentation]; !ok {
			documentations_newInstances = append(documentations_newInstances, documentation)
			newInstancesSlice = append(newInstancesSlice, documentation.GongMarshallIdentifier(stage))
			if stage.Documentations_referenceOrder == nil {
				stage.Documentations_referenceOrder = make(map[*Documentation]uint)
			}
			stage.Documentations_referenceOrder[documentation] = stage.Documentation_stagedOrder[documentation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, documentation.GongMarshallUnstaging(stage))
			// delete(stage.Documentations_referenceOrder, documentation)
			fieldInitializers, pointersInitializations := documentation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Documentation_stagedOrder[ref] = stage.Documentation_stagedOrder[documentation]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := documentation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, documentation)
			// delete(stage.Documentation_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", documentation.GetName())
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
	for _, ref := range stage.Documentations_reference {
		instance := stage.Documentations_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Documentations[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			documentations_deletedInstances = append(documentations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(documentations_newInstances)
	lenDeletedInstances += len(documentations_deletedInstances)
	var elements_newInstances []*Element
	var elements_deletedInstances []*Element

	// parse all staged instances and check if they have a reference
	for element := range stage.Elements {
		if ref, ok := stage.Elements_reference[element]; !ok {
			elements_newInstances = append(elements_newInstances, element)
			newInstancesSlice = append(newInstancesSlice, element.GongMarshallIdentifier(stage))
			if stage.Elements_referenceOrder == nil {
				stage.Elements_referenceOrder = make(map[*Element]uint)
			}
			stage.Elements_referenceOrder[element] = stage.Element_stagedOrder[element]
			newInstancesReverseSlice = append(newInstancesReverseSlice, element.GongMarshallUnstaging(stage))
			// delete(stage.Elements_referenceOrder, element)
			fieldInitializers, pointersInitializations := element.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Element_stagedOrder[ref] = stage.Element_stagedOrder[element]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := element.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, element)
			// delete(stage.Element_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", element.GetName())
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
	for _, ref := range stage.Elements_reference {
		instance := stage.Elements_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Elements[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			elements_deletedInstances = append(elements_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(elements_newInstances)
	lenDeletedInstances += len(elements_deletedInstances)
	var enumerations_newInstances []*Enumeration
	var enumerations_deletedInstances []*Enumeration

	// parse all staged instances and check if they have a reference
	for enumeration := range stage.Enumerations {
		if ref, ok := stage.Enumerations_reference[enumeration]; !ok {
			enumerations_newInstances = append(enumerations_newInstances, enumeration)
			newInstancesSlice = append(newInstancesSlice, enumeration.GongMarshallIdentifier(stage))
			if stage.Enumerations_referenceOrder == nil {
				stage.Enumerations_referenceOrder = make(map[*Enumeration]uint)
			}
			stage.Enumerations_referenceOrder[enumeration] = stage.Enumeration_stagedOrder[enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, enumeration.GongMarshallUnstaging(stage))
			// delete(stage.Enumerations_referenceOrder, enumeration)
			fieldInitializers, pointersInitializations := enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Enumeration_stagedOrder[ref] = stage.Enumeration_stagedOrder[enumeration]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, enumeration)
			// delete(stage.Enumeration_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", enumeration.GetName())
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
	for _, ref := range stage.Enumerations_reference {
		instance := stage.Enumerations_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Enumerations[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			enumerations_deletedInstances = append(enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(enumerations_newInstances)
	lenDeletedInstances += len(enumerations_deletedInstances)
	var extensions_newInstances []*Extension
	var extensions_deletedInstances []*Extension

	// parse all staged instances and check if they have a reference
	for extension := range stage.Extensions {
		if ref, ok := stage.Extensions_reference[extension]; !ok {
			extensions_newInstances = append(extensions_newInstances, extension)
			newInstancesSlice = append(newInstancesSlice, extension.GongMarshallIdentifier(stage))
			if stage.Extensions_referenceOrder == nil {
				stage.Extensions_referenceOrder = make(map[*Extension]uint)
			}
			stage.Extensions_referenceOrder[extension] = stage.Extension_stagedOrder[extension]
			newInstancesReverseSlice = append(newInstancesReverseSlice, extension.GongMarshallUnstaging(stage))
			// delete(stage.Extensions_referenceOrder, extension)
			fieldInitializers, pointersInitializations := extension.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Extension_stagedOrder[ref] = stage.Extension_stagedOrder[extension]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := extension.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, extension)
			// delete(stage.Extension_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", extension.GetName())
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
	for _, ref := range stage.Extensions_reference {
		instance := stage.Extensions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Extensions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			extensions_deletedInstances = append(extensions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(extensions_newInstances)
	lenDeletedInstances += len(extensions_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			newInstancesSlice = append(newInstancesSlice, group.GongMarshallIdentifier(stage))
			if stage.Groups_referenceOrder == nil {
				stage.Groups_referenceOrder = make(map[*Group]uint)
			}
			stage.Groups_referenceOrder[group] = stage.Group_stagedOrder[group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, group.GongMarshallUnstaging(stage))
			// delete(stage.Groups_referenceOrder, group)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Group_stagedOrder[ref] = stage.Group_stagedOrder[group]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, group)
			// delete(stage.Group_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", group.GetName())
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
	for _, ref := range stage.Groups_reference {
		instance := stage.Groups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Groups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			groups_deletedInstances = append(groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var lengths_newInstances []*Length
	var lengths_deletedInstances []*Length

	// parse all staged instances and check if they have a reference
	for length := range stage.Lengths {
		if ref, ok := stage.Lengths_reference[length]; !ok {
			lengths_newInstances = append(lengths_newInstances, length)
			newInstancesSlice = append(newInstancesSlice, length.GongMarshallIdentifier(stage))
			if stage.Lengths_referenceOrder == nil {
				stage.Lengths_referenceOrder = make(map[*Length]uint)
			}
			stage.Lengths_referenceOrder[length] = stage.Length_stagedOrder[length]
			newInstancesReverseSlice = append(newInstancesReverseSlice, length.GongMarshallUnstaging(stage))
			// delete(stage.Lengths_referenceOrder, length)
			fieldInitializers, pointersInitializations := length.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Length_stagedOrder[ref] = stage.Length_stagedOrder[length]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := length.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, length)
			// delete(stage.Length_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", length.GetName())
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
	for _, ref := range stage.Lengths_reference {
		instance := stage.Lengths_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Lengths[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			lengths_deletedInstances = append(lengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(lengths_newInstances)
	lenDeletedInstances += len(lengths_deletedInstances)
	var maxinclusives_newInstances []*MaxInclusive
	var maxinclusives_deletedInstances []*MaxInclusive

	// parse all staged instances and check if they have a reference
	for maxinclusive := range stage.MaxInclusives {
		if ref, ok := stage.MaxInclusives_reference[maxinclusive]; !ok {
			maxinclusives_newInstances = append(maxinclusives_newInstances, maxinclusive)
			newInstancesSlice = append(newInstancesSlice, maxinclusive.GongMarshallIdentifier(stage))
			if stage.MaxInclusives_referenceOrder == nil {
				stage.MaxInclusives_referenceOrder = make(map[*MaxInclusive]uint)
			}
			stage.MaxInclusives_referenceOrder[maxinclusive] = stage.MaxInclusive_stagedOrder[maxinclusive]
			newInstancesReverseSlice = append(newInstancesReverseSlice, maxinclusive.GongMarshallUnstaging(stage))
			// delete(stage.MaxInclusives_referenceOrder, maxinclusive)
			fieldInitializers, pointersInitializations := maxinclusive.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MaxInclusive_stagedOrder[ref] = stage.MaxInclusive_stagedOrder[maxinclusive]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := maxinclusive.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, maxinclusive)
			// delete(stage.MaxInclusive_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", maxinclusive.GetName())
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
	for _, ref := range stage.MaxInclusives_reference {
		instance := stage.MaxInclusives_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MaxInclusives[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			maxinclusives_deletedInstances = append(maxinclusives_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(maxinclusives_newInstances)
	lenDeletedInstances += len(maxinclusives_deletedInstances)
	var maxlengths_newInstances []*MaxLength
	var maxlengths_deletedInstances []*MaxLength

	// parse all staged instances and check if they have a reference
	for maxlength := range stage.MaxLengths {
		if ref, ok := stage.MaxLengths_reference[maxlength]; !ok {
			maxlengths_newInstances = append(maxlengths_newInstances, maxlength)
			newInstancesSlice = append(newInstancesSlice, maxlength.GongMarshallIdentifier(stage))
			if stage.MaxLengths_referenceOrder == nil {
				stage.MaxLengths_referenceOrder = make(map[*MaxLength]uint)
			}
			stage.MaxLengths_referenceOrder[maxlength] = stage.MaxLength_stagedOrder[maxlength]
			newInstancesReverseSlice = append(newInstancesReverseSlice, maxlength.GongMarshallUnstaging(stage))
			// delete(stage.MaxLengths_referenceOrder, maxlength)
			fieldInitializers, pointersInitializations := maxlength.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MaxLength_stagedOrder[ref] = stage.MaxLength_stagedOrder[maxlength]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := maxlength.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, maxlength)
			// delete(stage.MaxLength_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", maxlength.GetName())
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
	for _, ref := range stage.MaxLengths_reference {
		instance := stage.MaxLengths_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MaxLengths[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			maxlengths_deletedInstances = append(maxlengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(maxlengths_newInstances)
	lenDeletedInstances += len(maxlengths_deletedInstances)
	var mininclusives_newInstances []*MinInclusive
	var mininclusives_deletedInstances []*MinInclusive

	// parse all staged instances and check if they have a reference
	for mininclusive := range stage.MinInclusives {
		if ref, ok := stage.MinInclusives_reference[mininclusive]; !ok {
			mininclusives_newInstances = append(mininclusives_newInstances, mininclusive)
			newInstancesSlice = append(newInstancesSlice, mininclusive.GongMarshallIdentifier(stage))
			if stage.MinInclusives_referenceOrder == nil {
				stage.MinInclusives_referenceOrder = make(map[*MinInclusive]uint)
			}
			stage.MinInclusives_referenceOrder[mininclusive] = stage.MinInclusive_stagedOrder[mininclusive]
			newInstancesReverseSlice = append(newInstancesReverseSlice, mininclusive.GongMarshallUnstaging(stage))
			// delete(stage.MinInclusives_referenceOrder, mininclusive)
			fieldInitializers, pointersInitializations := mininclusive.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MinInclusive_stagedOrder[ref] = stage.MinInclusive_stagedOrder[mininclusive]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := mininclusive.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, mininclusive)
			// delete(stage.MinInclusive_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", mininclusive.GetName())
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
	for _, ref := range stage.MinInclusives_reference {
		instance := stage.MinInclusives_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MinInclusives[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			mininclusives_deletedInstances = append(mininclusives_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(mininclusives_newInstances)
	lenDeletedInstances += len(mininclusives_deletedInstances)
	var minlengths_newInstances []*MinLength
	var minlengths_deletedInstances []*MinLength

	// parse all staged instances and check if they have a reference
	for minlength := range stage.MinLengths {
		if ref, ok := stage.MinLengths_reference[minlength]; !ok {
			minlengths_newInstances = append(minlengths_newInstances, minlength)
			newInstancesSlice = append(newInstancesSlice, minlength.GongMarshallIdentifier(stage))
			if stage.MinLengths_referenceOrder == nil {
				stage.MinLengths_referenceOrder = make(map[*MinLength]uint)
			}
			stage.MinLengths_referenceOrder[minlength] = stage.MinLength_stagedOrder[minlength]
			newInstancesReverseSlice = append(newInstancesReverseSlice, minlength.GongMarshallUnstaging(stage))
			// delete(stage.MinLengths_referenceOrder, minlength)
			fieldInitializers, pointersInitializations := minlength.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MinLength_stagedOrder[ref] = stage.MinLength_stagedOrder[minlength]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := minlength.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, minlength)
			// delete(stage.MinLength_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", minlength.GetName())
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
	for _, ref := range stage.MinLengths_reference {
		instance := stage.MinLengths_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MinLengths[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			minlengths_deletedInstances = append(minlengths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(minlengths_newInstances)
	lenDeletedInstances += len(minlengths_deletedInstances)
	var patterns_newInstances []*Pattern
	var patterns_deletedInstances []*Pattern

	// parse all staged instances and check if they have a reference
	for pattern := range stage.Patterns {
		if ref, ok := stage.Patterns_reference[pattern]; !ok {
			patterns_newInstances = append(patterns_newInstances, pattern)
			newInstancesSlice = append(newInstancesSlice, pattern.GongMarshallIdentifier(stage))
			if stage.Patterns_referenceOrder == nil {
				stage.Patterns_referenceOrder = make(map[*Pattern]uint)
			}
			stage.Patterns_referenceOrder[pattern] = stage.Pattern_stagedOrder[pattern]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pattern.GongMarshallUnstaging(stage))
			// delete(stage.Patterns_referenceOrder, pattern)
			fieldInitializers, pointersInitializations := pattern.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Pattern_stagedOrder[ref] = stage.Pattern_stagedOrder[pattern]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := pattern.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pattern)
			// delete(stage.Pattern_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", pattern.GetName())
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
	for _, ref := range stage.Patterns_reference {
		instance := stage.Patterns_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Patterns[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			patterns_deletedInstances = append(patterns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(patterns_newInstances)
	lenDeletedInstances += len(patterns_deletedInstances)
	var restrictions_newInstances []*Restriction
	var restrictions_deletedInstances []*Restriction

	// parse all staged instances and check if they have a reference
	for restriction := range stage.Restrictions {
		if ref, ok := stage.Restrictions_reference[restriction]; !ok {
			restrictions_newInstances = append(restrictions_newInstances, restriction)
			newInstancesSlice = append(newInstancesSlice, restriction.GongMarshallIdentifier(stage))
			if stage.Restrictions_referenceOrder == nil {
				stage.Restrictions_referenceOrder = make(map[*Restriction]uint)
			}
			stage.Restrictions_referenceOrder[restriction] = stage.Restriction_stagedOrder[restriction]
			newInstancesReverseSlice = append(newInstancesReverseSlice, restriction.GongMarshallUnstaging(stage))
			// delete(stage.Restrictions_referenceOrder, restriction)
			fieldInitializers, pointersInitializations := restriction.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Restriction_stagedOrder[ref] = stage.Restriction_stagedOrder[restriction]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := restriction.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, restriction)
			// delete(stage.Restriction_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", restriction.GetName())
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
	for _, ref := range stage.Restrictions_reference {
		instance := stage.Restrictions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Restrictions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			restrictions_deletedInstances = append(restrictions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(restrictions_newInstances)
	lenDeletedInstances += len(restrictions_deletedInstances)
	var schemas_newInstances []*Schema
	var schemas_deletedInstances []*Schema

	// parse all staged instances and check if they have a reference
	for schema := range stage.Schemas {
		if ref, ok := stage.Schemas_reference[schema]; !ok {
			schemas_newInstances = append(schemas_newInstances, schema)
			newInstancesSlice = append(newInstancesSlice, schema.GongMarshallIdentifier(stage))
			if stage.Schemas_referenceOrder == nil {
				stage.Schemas_referenceOrder = make(map[*Schema]uint)
			}
			stage.Schemas_referenceOrder[schema] = stage.Schema_stagedOrder[schema]
			newInstancesReverseSlice = append(newInstancesReverseSlice, schema.GongMarshallUnstaging(stage))
			// delete(stage.Schemas_referenceOrder, schema)
			fieldInitializers, pointersInitializations := schema.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Schema_stagedOrder[ref] = stage.Schema_stagedOrder[schema]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := schema.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, schema)
			// delete(stage.Schema_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", schema.GetName())
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
	for _, ref := range stage.Schemas_reference {
		instance := stage.Schemas_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Schemas[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			schemas_deletedInstances = append(schemas_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(schemas_newInstances)
	lenDeletedInstances += len(schemas_deletedInstances)
	var sequences_newInstances []*Sequence
	var sequences_deletedInstances []*Sequence

	// parse all staged instances and check if they have a reference
	for sequence := range stage.Sequences {
		if ref, ok := stage.Sequences_reference[sequence]; !ok {
			sequences_newInstances = append(sequences_newInstances, sequence)
			newInstancesSlice = append(newInstancesSlice, sequence.GongMarshallIdentifier(stage))
			if stage.Sequences_referenceOrder == nil {
				stage.Sequences_referenceOrder = make(map[*Sequence]uint)
			}
			stage.Sequences_referenceOrder[sequence] = stage.Sequence_stagedOrder[sequence]
			newInstancesReverseSlice = append(newInstancesReverseSlice, sequence.GongMarshallUnstaging(stage))
			// delete(stage.Sequences_referenceOrder, sequence)
			fieldInitializers, pointersInitializations := sequence.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Sequence_stagedOrder[ref] = stage.Sequence_stagedOrder[sequence]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := sequence.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, sequence)
			// delete(stage.Sequence_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", sequence.GetName())
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
	for _, ref := range stage.Sequences_reference {
		instance := stage.Sequences_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Sequences[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			sequences_deletedInstances = append(sequences_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(sequences_newInstances)
	lenDeletedInstances += len(sequences_deletedInstances)
	var simplecontents_newInstances []*SimpleContent
	var simplecontents_deletedInstances []*SimpleContent

	// parse all staged instances and check if they have a reference
	for simplecontent := range stage.SimpleContents {
		if ref, ok := stage.SimpleContents_reference[simplecontent]; !ok {
			simplecontents_newInstances = append(simplecontents_newInstances, simplecontent)
			newInstancesSlice = append(newInstancesSlice, simplecontent.GongMarshallIdentifier(stage))
			if stage.SimpleContents_referenceOrder == nil {
				stage.SimpleContents_referenceOrder = make(map[*SimpleContent]uint)
			}
			stage.SimpleContents_referenceOrder[simplecontent] = stage.SimpleContent_stagedOrder[simplecontent]
			newInstancesReverseSlice = append(newInstancesReverseSlice, simplecontent.GongMarshallUnstaging(stage))
			// delete(stage.SimpleContents_referenceOrder, simplecontent)
			fieldInitializers, pointersInitializations := simplecontent.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SimpleContent_stagedOrder[ref] = stage.SimpleContent_stagedOrder[simplecontent]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := simplecontent.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, simplecontent)
			// delete(stage.SimpleContent_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", simplecontent.GetName())
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
	for _, ref := range stage.SimpleContents_reference {
		instance := stage.SimpleContents_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SimpleContents[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			simplecontents_deletedInstances = append(simplecontents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(simplecontents_newInstances)
	lenDeletedInstances += len(simplecontents_deletedInstances)
	var simpletypes_newInstances []*SimpleType
	var simpletypes_deletedInstances []*SimpleType

	// parse all staged instances and check if they have a reference
	for simpletype := range stage.SimpleTypes {
		if ref, ok := stage.SimpleTypes_reference[simpletype]; !ok {
			simpletypes_newInstances = append(simpletypes_newInstances, simpletype)
			newInstancesSlice = append(newInstancesSlice, simpletype.GongMarshallIdentifier(stage))
			if stage.SimpleTypes_referenceOrder == nil {
				stage.SimpleTypes_referenceOrder = make(map[*SimpleType]uint)
			}
			stage.SimpleTypes_referenceOrder[simpletype] = stage.SimpleType_stagedOrder[simpletype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, simpletype.GongMarshallUnstaging(stage))
			// delete(stage.SimpleTypes_referenceOrder, simpletype)
			fieldInitializers, pointersInitializations := simpletype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SimpleType_stagedOrder[ref] = stage.SimpleType_stagedOrder[simpletype]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := simpletype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, simpletype)
			// delete(stage.SimpleType_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", simpletype.GetName())
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
	for _, ref := range stage.SimpleTypes_reference {
		instance := stage.SimpleTypes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SimpleTypes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			simpletypes_deletedInstances = append(simpletypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(simpletypes_newInstances)
	lenDeletedInstances += len(simpletypes_deletedInstances)
	var totaldigits_newInstances []*TotalDigit
	var totaldigits_deletedInstances []*TotalDigit

	// parse all staged instances and check if they have a reference
	for totaldigit := range stage.TotalDigits {
		if ref, ok := stage.TotalDigits_reference[totaldigit]; !ok {
			totaldigits_newInstances = append(totaldigits_newInstances, totaldigit)
			newInstancesSlice = append(newInstancesSlice, totaldigit.GongMarshallIdentifier(stage))
			if stage.TotalDigits_referenceOrder == nil {
				stage.TotalDigits_referenceOrder = make(map[*TotalDigit]uint)
			}
			stage.TotalDigits_referenceOrder[totaldigit] = stage.TotalDigit_stagedOrder[totaldigit]
			newInstancesReverseSlice = append(newInstancesReverseSlice, totaldigit.GongMarshallUnstaging(stage))
			// delete(stage.TotalDigits_referenceOrder, totaldigit)
			fieldInitializers, pointersInitializations := totaldigit.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TotalDigit_stagedOrder[ref] = stage.TotalDigit_stagedOrder[totaldigit]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := totaldigit.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, totaldigit)
			// delete(stage.TotalDigit_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", totaldigit.GetName())
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
	for _, ref := range stage.TotalDigits_reference {
		instance := stage.TotalDigits_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TotalDigits[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			totaldigits_deletedInstances = append(totaldigits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(totaldigits_newInstances)
	lenDeletedInstances += len(totaldigits_deletedInstances)
	var unions_newInstances []*Union
	var unions_deletedInstances []*Union

	// parse all staged instances and check if they have a reference
	for union := range stage.Unions {
		if ref, ok := stage.Unions_reference[union]; !ok {
			unions_newInstances = append(unions_newInstances, union)
			newInstancesSlice = append(newInstancesSlice, union.GongMarshallIdentifier(stage))
			if stage.Unions_referenceOrder == nil {
				stage.Unions_referenceOrder = make(map[*Union]uint)
			}
			stage.Unions_referenceOrder[union] = stage.Union_stagedOrder[union]
			newInstancesReverseSlice = append(newInstancesReverseSlice, union.GongMarshallUnstaging(stage))
			// delete(stage.Unions_referenceOrder, union)
			fieldInitializers, pointersInitializations := union.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Union_stagedOrder[ref] = stage.Union_stagedOrder[union]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := union.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, union)
			// delete(stage.Union_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", union.GetName())
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
	for _, ref := range stage.Unions_reference {
		instance := stage.Unions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Unions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			unions_deletedInstances = append(unions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(unions_newInstances)
	lenDeletedInstances += len(unions_deletedInstances)
	var whitespaces_newInstances []*WhiteSpace
	var whitespaces_deletedInstances []*WhiteSpace

	// parse all staged instances and check if they have a reference
	for whitespace := range stage.WhiteSpaces {
		if ref, ok := stage.WhiteSpaces_reference[whitespace]; !ok {
			whitespaces_newInstances = append(whitespaces_newInstances, whitespace)
			newInstancesSlice = append(newInstancesSlice, whitespace.GongMarshallIdentifier(stage))
			if stage.WhiteSpaces_referenceOrder == nil {
				stage.WhiteSpaces_referenceOrder = make(map[*WhiteSpace]uint)
			}
			stage.WhiteSpaces_referenceOrder[whitespace] = stage.WhiteSpace_stagedOrder[whitespace]
			newInstancesReverseSlice = append(newInstancesReverseSlice, whitespace.GongMarshallUnstaging(stage))
			// delete(stage.WhiteSpaces_referenceOrder, whitespace)
			fieldInitializers, pointersInitializations := whitespace.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.WhiteSpace_stagedOrder[ref] = stage.WhiteSpace_stagedOrder[whitespace]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := whitespace.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, whitespace)
			// delete(stage.WhiteSpace_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", whitespace.GetName())
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
	for _, ref := range stage.WhiteSpaces_reference {
		instance := stage.WhiteSpaces_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.WhiteSpaces[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			whitespaces_deletedInstances = append(whitespaces_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(whitespaces_newInstances)
	lenDeletedInstances += len(whitespaces_deletedInstances)

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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(all.Name))
	return
}

func (annotation *Annotation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", annotation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Annotation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(annotation.Name))
	return
}

func (attribute *Attribute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attribute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(attribute.Name))
	return
}

func (attributegroup *AttributeGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributegroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(attributegroup.Name))
	return
}

func (choice *Choice) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", choice.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Choice")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(choice.Name))
	return
}

func (complexcontent *ComplexContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complexcontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(complexcontent.Name))
	return
}

func (complextype *ComplexType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", complextype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(complextype.Name))
	return
}

func (documentation *Documentation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", documentation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Documentation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(documentation.Name))
	return
}

func (element *Element) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", element.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(element.Name))
	return
}

func (enumeration *Enumeration) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Enumeration")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(enumeration.Name))
	return
}

func (extension *Extension) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extension.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extension")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(extension.Name))
	return
}

func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(group.Name))
	return
}

func (length *Length) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", length.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Length")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(length.Name))
	return
}

func (maxinclusive *MaxInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxinclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(maxinclusive.Name))
	return
}

func (maxlength *MaxLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", maxlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(maxlength.Name))
	return
}

func (mininclusive *MinInclusive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mininclusive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinInclusive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(mininclusive.Name))
	return
}

func (minlength *MinLength) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", minlength.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinLength")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(minlength.Name))
	return
}

func (pattern *Pattern) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pattern.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pattern")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(pattern.Name))
	return
}

func (restriction *Restriction) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", restriction.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Restriction")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(restriction.Name))
	return
}

func (schema *Schema) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", schema.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Schema")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(schema.Name))
	return
}

func (sequence *Sequence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sequence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sequence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(sequence.Name))
	return
}

func (simplecontent *SimpleContent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simplecontent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleContent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(simplecontent.Name))
	return
}

func (simpletype *SimpleType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", simpletype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(simpletype.Name))
	return
}

func (totaldigit *TotalDigit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", totaldigit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TotalDigit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(totaldigit.Name))
	return
}

func (union *Union) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", union.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Union")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(union.Name))
	return
}

func (whitespace *WhiteSpace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", whitespace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "WhiteSpace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(whitespace.Name))
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

// end of template
