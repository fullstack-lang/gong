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

	// Compute reverse map for named struct Element
	// insertion point per field
	stage.Element_Groups_reverseMap = make(map[*Group]*Element)
	for element := range stage.Elements {
		_ = element
		for _, _group := range element.Groups {
			stage.Element_Groups_reverseMap[_group] = element
		}
	}

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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Alls)

	res = __gong__appendInstances(res, stage.Annotations)

	res = __gong__appendInstances(res, stage.Attributes)

	res = __gong__appendInstances(res, stage.AttributeGroups)

	res = __gong__appendInstances(res, stage.Choices)

	res = __gong__appendInstances(res, stage.ComplexContents)

	res = __gong__appendInstances(res, stage.ComplexTypes)

	res = __gong__appendInstances(res, stage.Documentations)

	res = __gong__appendInstances(res, stage.Elements)

	res = __gong__appendInstances(res, stage.Enumerations)

	res = __gong__appendInstances(res, stage.Extensions)

	res = __gong__appendInstances(res, stage.Groups)

	res = __gong__appendInstances(res, stage.Lengths)

	res = __gong__appendInstances(res, stage.MaxInclusives)

	res = __gong__appendInstances(res, stage.MaxLengths)

	res = __gong__appendInstances(res, stage.MinInclusives)

	res = __gong__appendInstances(res, stage.MinLengths)

	res = __gong__appendInstances(res, stage.Patterns)

	res = __gong__appendInstances(res, stage.Restrictions)

	res = __gong__appendInstances(res, stage.Schemas)

	res = __gong__appendInstances(res, stage.Sequences)

	res = __gong__appendInstances(res, stage.SimpleContents)

	res = __gong__appendInstances(res, stage.SimpleTypes)

	res = __gong__appendInstances(res, stage.TotalDigits)

	res = __gong__appendInstances(res, stage.Unions)

	res = __gong__appendInstances(res, stage.WhiteSpaces)

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
func (all *All) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, all)
}

func (annotation *Annotation) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, annotation)
}

func (attribute *Attribute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute)
}

func (attributegroup *AttributeGroup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attributegroup)
}

func (choice *Choice) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, choice)
}

func (complexcontent *ComplexContent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, complexcontent)
}

func (complextype *ComplexType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, complextype)
}

func (documentation *Documentation) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, documentation)
}

func (element *Element) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, element)
}

func (enumeration *Enumeration) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, enumeration)
}

func (extension *Extension) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, extension)
}

func (group *Group) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group)
}

func (length *Length) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, length)
}

func (maxinclusive *MaxInclusive) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, maxinclusive)
}

func (maxlength *MaxLength) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, maxlength)
}

func (mininclusive *MinInclusive) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mininclusive)
}

func (minlength *MinLength) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, minlength)
}

func (pattern *Pattern) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pattern)
}

func (restriction *Restriction) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, restriction)
}

func (schema *Schema) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, schema)
}

func (sequence *Sequence) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sequence)
}

func (simplecontent *SimpleContent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, simplecontent)
}

func (simpletype *SimpleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, simpletype)
}

func (totaldigit *TotalDigit) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, totaldigit)
}

func (union *Union) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, union)
}

func (whitespace *WhiteSpace) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, whitespace)
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
	__gong__computeReferencePass1(stage, stage.Alls, &stage.Alls_reference, &stage.Alls_referenceOrder, &stage.Alls_instance)

	__gong__computeReferencePass1(stage, stage.Annotations, &stage.Annotations_reference, &stage.Annotations_referenceOrder, &stage.Annotations_instance)

	__gong__computeReferencePass1(stage, stage.Attributes, &stage.Attributes_reference, &stage.Attributes_referenceOrder, &stage.Attributes_instance)

	__gong__computeReferencePass1(stage, stage.AttributeGroups, &stage.AttributeGroups_reference, &stage.AttributeGroups_referenceOrder, &stage.AttributeGroups_instance)

	__gong__computeReferencePass1(stage, stage.Choices, &stage.Choices_reference, &stage.Choices_referenceOrder, &stage.Choices_instance)

	__gong__computeReferencePass1(stage, stage.ComplexContents, &stage.ComplexContents_reference, &stage.ComplexContents_referenceOrder, &stage.ComplexContents_instance)

	__gong__computeReferencePass1(stage, stage.ComplexTypes, &stage.ComplexTypes_reference, &stage.ComplexTypes_referenceOrder, &stage.ComplexTypes_instance)

	__gong__computeReferencePass1(stage, stage.Documentations, &stage.Documentations_reference, &stage.Documentations_referenceOrder, &stage.Documentations_instance)

	__gong__computeReferencePass1(stage, stage.Elements, &stage.Elements_reference, &stage.Elements_referenceOrder, &stage.Elements_instance)

	__gong__computeReferencePass1(stage, stage.Enumerations, &stage.Enumerations_reference, &stage.Enumerations_referenceOrder, &stage.Enumerations_instance)

	__gong__computeReferencePass1(stage, stage.Extensions, &stage.Extensions_reference, &stage.Extensions_referenceOrder, &stage.Extensions_instance)

	__gong__computeReferencePass1(stage, stage.Groups, &stage.Groups_reference, &stage.Groups_referenceOrder, &stage.Groups_instance)

	__gong__computeReferencePass1(stage, stage.Lengths, &stage.Lengths_reference, &stage.Lengths_referenceOrder, &stage.Lengths_instance)

	__gong__computeReferencePass1(stage, stage.MaxInclusives, &stage.MaxInclusives_reference, &stage.MaxInclusives_referenceOrder, &stage.MaxInclusives_instance)

	__gong__computeReferencePass1(stage, stage.MaxLengths, &stage.MaxLengths_reference, &stage.MaxLengths_referenceOrder, &stage.MaxLengths_instance)

	__gong__computeReferencePass1(stage, stage.MinInclusives, &stage.MinInclusives_reference, &stage.MinInclusives_referenceOrder, &stage.MinInclusives_instance)

	__gong__computeReferencePass1(stage, stage.MinLengths, &stage.MinLengths_reference, &stage.MinLengths_referenceOrder, &stage.MinLengths_instance)

	__gong__computeReferencePass1(stage, stage.Patterns, &stage.Patterns_reference, &stage.Patterns_referenceOrder, &stage.Patterns_instance)

	__gong__computeReferencePass1(stage, stage.Restrictions, &stage.Restrictions_reference, &stage.Restrictions_referenceOrder, &stage.Restrictions_instance)

	__gong__computeReferencePass1(stage, stage.Schemas, &stage.Schemas_reference, &stage.Schemas_referenceOrder, &stage.Schemas_instance)

	__gong__computeReferencePass1(stage, stage.Sequences, &stage.Sequences_reference, &stage.Sequences_referenceOrder, &stage.Sequences_instance)

	__gong__computeReferencePass1(stage, stage.SimpleContents, &stage.SimpleContents_reference, &stage.SimpleContents_referenceOrder, &stage.SimpleContents_instance)

	__gong__computeReferencePass1(stage, stage.SimpleTypes, &stage.SimpleTypes_reference, &stage.SimpleTypes_referenceOrder, &stage.SimpleTypes_instance)

	__gong__computeReferencePass1(stage, stage.TotalDigits, &stage.TotalDigits_reference, &stage.TotalDigits_referenceOrder, &stage.TotalDigits_instance)

	__gong__computeReferencePass1(stage, stage.Unions, &stage.Unions_reference, &stage.Unions_referenceOrder, &stage.Unions_instance)

	__gong__computeReferencePass1(stage, stage.WhiteSpaces, &stage.WhiteSpaces_reference, &stage.WhiteSpaces_referenceOrder, &stage.WhiteSpaces_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Alls, stage.Alls_reference, stage)

	__gong__computeReferencePass2(stage.Annotations, stage.Annotations_reference, stage)

	__gong__computeReferencePass2(stage.Attributes, stage.Attributes_reference, stage)

	__gong__computeReferencePass2(stage.AttributeGroups, stage.AttributeGroups_reference, stage)

	__gong__computeReferencePass2(stage.Choices, stage.Choices_reference, stage)

	__gong__computeReferencePass2(stage.ComplexContents, stage.ComplexContents_reference, stage)

	__gong__computeReferencePass2(stage.ComplexTypes, stage.ComplexTypes_reference, stage)

	__gong__computeReferencePass2(stage.Documentations, stage.Documentations_reference, stage)

	__gong__computeReferencePass2(stage.Elements, stage.Elements_reference, stage)

	__gong__computeReferencePass2(stage.Enumerations, stage.Enumerations_reference, stage)

	__gong__computeReferencePass2(stage.Extensions, stage.Extensions_reference, stage)

	__gong__computeReferencePass2(stage.Groups, stage.Groups_reference, stage)

	__gong__computeReferencePass2(stage.Lengths, stage.Lengths_reference, stage)

	__gong__computeReferencePass2(stage.MaxInclusives, stage.MaxInclusives_reference, stage)

	__gong__computeReferencePass2(stage.MaxLengths, stage.MaxLengths_reference, stage)

	__gong__computeReferencePass2(stage.MinInclusives, stage.MinInclusives_reference, stage)

	__gong__computeReferencePass2(stage.MinLengths, stage.MinLengths_reference, stage)

	__gong__computeReferencePass2(stage.Patterns, stage.Patterns_reference, stage)

	__gong__computeReferencePass2(stage.Restrictions, stage.Restrictions_reference, stage)

	__gong__computeReferencePass2(stage.Schemas, stage.Schemas_reference, stage)

	__gong__computeReferencePass2(stage.Sequences, stage.Sequences_reference, stage)

	__gong__computeReferencePass2(stage.SimpleContents, stage.SimpleContents_reference, stage)

	__gong__computeReferencePass2(stage.SimpleTypes, stage.SimpleTypes_reference, stage)

	__gong__computeReferencePass2(stage.TotalDigits, stage.TotalDigits_reference, stage)

	__gong__computeReferencePass2(stage.Unions, stage.Unions_reference, stage)

	__gong__computeReferencePass2(stage.WhiteSpaces, stage.WhiteSpaces_reference, stage)

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
	return __gong__getOrder(stage.All_stagedOrder, stage.Alls_referenceOrder, all, "All")
}

func (annotation *Annotation) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Annotation_stagedOrder, stage.Annotations_referenceOrder, annotation, "Annotation")
}

func (attribute *Attribute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Attribute_stagedOrder, stage.Attributes_referenceOrder, attribute, "Attribute")
}

func (attributegroup *AttributeGroup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AttributeGroup_stagedOrder, stage.AttributeGroups_referenceOrder, attributegroup, "AttributeGroup")
}

func (choice *Choice) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Choice_stagedOrder, stage.Choices_referenceOrder, choice, "Choice")
}

func (complexcontent *ComplexContent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ComplexContent_stagedOrder, stage.ComplexContents_referenceOrder, complexcontent, "ComplexContent")
}

func (complextype *ComplexType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ComplexType_stagedOrder, stage.ComplexTypes_referenceOrder, complextype, "ComplexType")
}

func (documentation *Documentation) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Documentation_stagedOrder, stage.Documentations_referenceOrder, documentation, "Documentation")
}

func (element *Element) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Element_stagedOrder, stage.Elements_referenceOrder, element, "Element")
}

func (enumeration *Enumeration) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Enumeration_stagedOrder, stage.Enumerations_referenceOrder, enumeration, "Enumeration")
}

func (extension *Extension) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Extension_stagedOrder, stage.Extensions_referenceOrder, extension, "Extension")
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_stagedOrder, stage.Groups_referenceOrder, group, "Group")
}

func (length *Length) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Length_stagedOrder, stage.Lengths_referenceOrder, length, "Length")
}

func (maxinclusive *MaxInclusive) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MaxInclusive_stagedOrder, stage.MaxInclusives_referenceOrder, maxinclusive, "MaxInclusive")
}

func (maxlength *MaxLength) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MaxLength_stagedOrder, stage.MaxLengths_referenceOrder, maxlength, "MaxLength")
}

func (mininclusive *MinInclusive) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MinInclusive_stagedOrder, stage.MinInclusives_referenceOrder, mininclusive, "MinInclusive")
}

func (minlength *MinLength) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MinLength_stagedOrder, stage.MinLengths_referenceOrder, minlength, "MinLength")
}

func (pattern *Pattern) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pattern_stagedOrder, stage.Patterns_referenceOrder, pattern, "Pattern")
}

func (restriction *Restriction) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Restriction_stagedOrder, stage.Restrictions_referenceOrder, restriction, "Restriction")
}

func (schema *Schema) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Schema_stagedOrder, stage.Schemas_referenceOrder, schema, "Schema")
}

func (sequence *Sequence) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Sequence_stagedOrder, stage.Sequences_referenceOrder, sequence, "Sequence")
}

func (simplecontent *SimpleContent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SimpleContent_stagedOrder, stage.SimpleContents_referenceOrder, simplecontent, "SimpleContent")
}

func (simpletype *SimpleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SimpleType_stagedOrder, stage.SimpleTypes_referenceOrder, simpletype, "SimpleType")
}

func (totaldigit *TotalDigit) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TotalDigit_stagedOrder, stage.TotalDigits_referenceOrder, totaldigit, "TotalDigit")
}

func (union *Union) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Union_stagedOrder, stage.Unions_referenceOrder, union, "Union")
}

func (whitespace *WhiteSpace) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.WhiteSpace_stagedOrder, stage.WhiteSpaces_referenceOrder, whitespace, "WhiteSpace")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (all *All) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(all, all.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (all *All) GongGetReferenceIdentifier(stage *Stage) string {
	return all.GongGetIdentifier(stage)
}

func (annotation *Annotation) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(annotation, annotation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (annotation *Annotation) GongGetReferenceIdentifier(stage *Stage) string {
	return annotation.GongGetIdentifier(stage)
}

func (attribute *Attribute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute, attribute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute *Attribute) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute.GongGetIdentifier(stage)
}

func (attributegroup *AttributeGroup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attributegroup, attributegroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributegroup *AttributeGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return attributegroup.GongGetIdentifier(stage)
}

func (choice *Choice) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(choice, choice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (choice *Choice) GongGetReferenceIdentifier(stage *Stage) string {
	return choice.GongGetIdentifier(stage)
}

func (complexcontent *ComplexContent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(complexcontent, complexcontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complexcontent *ComplexContent) GongGetReferenceIdentifier(stage *Stage) string {
	return complexcontent.GongGetIdentifier(stage)
}

func (complextype *ComplexType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(complextype, complextype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (complextype *ComplexType) GongGetReferenceIdentifier(stage *Stage) string {
	return complextype.GongGetIdentifier(stage)
}

func (documentation *Documentation) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(documentation, documentation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (documentation *Documentation) GongGetReferenceIdentifier(stage *Stage) string {
	return documentation.GongGetIdentifier(stage)
}

func (element *Element) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(element, element.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (element *Element) GongGetReferenceIdentifier(stage *Stage) string {
	return element.GongGetIdentifier(stage)
}

func (enumeration *Enumeration) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(enumeration, enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enumeration *Enumeration) GongGetReferenceIdentifier(stage *Stage) string {
	return enumeration.GongGetIdentifier(stage)
}

func (extension *Extension) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(extension, extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extension *Extension) GongGetReferenceIdentifier(stage *Stage) string {
	return extension.GongGetIdentifier(stage)
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group, group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return group.GongGetIdentifier(stage)
}

func (length *Length) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(length, length.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (length *Length) GongGetReferenceIdentifier(stage *Stage) string {
	return length.GongGetIdentifier(stage)
}

func (maxinclusive *MaxInclusive) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(maxinclusive, maxinclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxinclusive *MaxInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return maxinclusive.GongGetIdentifier(stage)
}

func (maxlength *MaxLength) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(maxlength, maxlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (maxlength *MaxLength) GongGetReferenceIdentifier(stage *Stage) string {
	return maxlength.GongGetIdentifier(stage)
}

func (mininclusive *MinInclusive) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mininclusive, mininclusive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mininclusive *MinInclusive) GongGetReferenceIdentifier(stage *Stage) string {
	return mininclusive.GongGetIdentifier(stage)
}

func (minlength *MinLength) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(minlength, minlength.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (minlength *MinLength) GongGetReferenceIdentifier(stage *Stage) string {
	return minlength.GongGetIdentifier(stage)
}

func (pattern *Pattern) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pattern, pattern.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pattern *Pattern) GongGetReferenceIdentifier(stage *Stage) string {
	return pattern.GongGetIdentifier(stage)
}

func (restriction *Restriction) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(restriction, restriction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (restriction *Restriction) GongGetReferenceIdentifier(stage *Stage) string {
	return restriction.GongGetIdentifier(stage)
}

func (schema *Schema) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(schema, schema.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (schema *Schema) GongGetReferenceIdentifier(stage *Stage) string {
	return schema.GongGetIdentifier(stage)
}

func (sequence *Sequence) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sequence, sequence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sequence *Sequence) GongGetReferenceIdentifier(stage *Stage) string {
	return sequence.GongGetIdentifier(stage)
}

func (simplecontent *SimpleContent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(simplecontent, simplecontent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simplecontent *SimpleContent) GongGetReferenceIdentifier(stage *Stage) string {
	return simplecontent.GongGetIdentifier(stage)
}

func (simpletype *SimpleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(simpletype, simpletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (simpletype *SimpleType) GongGetReferenceIdentifier(stage *Stage) string {
	return simpletype.GongGetIdentifier(stage)
}

func (totaldigit *TotalDigit) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(totaldigit, totaldigit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (totaldigit *TotalDigit) GongGetReferenceIdentifier(stage *Stage) string {
	return totaldigit.GongGetIdentifier(stage)
}

func (union *Union) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(union, union.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (union *Union) GongGetReferenceIdentifier(stage *Stage) string {
	return union.GongGetIdentifier(stage)
}

func (whitespace *WhiteSpace) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(whitespace, whitespace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (whitespace *WhiteSpace) GongGetReferenceIdentifier(stage *Stage) string {
	return whitespace.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (all *All) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(all.GongGetIdentifier(stage), "All", all.Name)
}

func (annotation *Annotation) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(annotation.GongGetIdentifier(stage), "Annotation", annotation.Name)
}

func (attribute *Attribute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute.GongGetIdentifier(stage), "Attribute", attribute.Name)
}

func (attributegroup *AttributeGroup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attributegroup.GongGetIdentifier(stage), "AttributeGroup", attributegroup.Name)
}

func (choice *Choice) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(choice.GongGetIdentifier(stage), "Choice", choice.Name)
}

func (complexcontent *ComplexContent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(complexcontent.GongGetIdentifier(stage), "ComplexContent", complexcontent.Name)
}

func (complextype *ComplexType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(complextype.GongGetIdentifier(stage), "ComplexType", complextype.Name)
}

func (documentation *Documentation) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(documentation.GongGetIdentifier(stage), "Documentation", documentation.Name)
}

func (element *Element) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(element.GongGetIdentifier(stage), "Element", element.Name)
}

func (enumeration *Enumeration) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(enumeration.GongGetIdentifier(stage), "Enumeration", enumeration.Name)
}

func (extension *Extension) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(extension.GongGetIdentifier(stage), "Extension", extension.Name)
}

func (group *Group) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group.GongGetIdentifier(stage), "Group", group.Name)
}

func (length *Length) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(length.GongGetIdentifier(stage), "Length", length.Name)
}

func (maxinclusive *MaxInclusive) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(maxinclusive.GongGetIdentifier(stage), "MaxInclusive", maxinclusive.Name)
}

func (maxlength *MaxLength) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(maxlength.GongGetIdentifier(stage), "MaxLength", maxlength.Name)
}

func (mininclusive *MinInclusive) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mininclusive.GongGetIdentifier(stage), "MinInclusive", mininclusive.Name)
}

func (minlength *MinLength) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(minlength.GongGetIdentifier(stage), "MinLength", minlength.Name)
}

func (pattern *Pattern) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pattern.GongGetIdentifier(stage), "Pattern", pattern.Name)
}

func (restriction *Restriction) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(restriction.GongGetIdentifier(stage), "Restriction", restriction.Name)
}

func (schema *Schema) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(schema.GongGetIdentifier(stage), "Schema", schema.Name)
}

func (sequence *Sequence) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sequence.GongGetIdentifier(stage), "Sequence", sequence.Name)
}

func (simplecontent *SimpleContent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(simplecontent.GongGetIdentifier(stage), "SimpleContent", simplecontent.Name)
}

func (simpletype *SimpleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(simpletype.GongGetIdentifier(stage), "SimpleType", simpletype.Name)
}

func (totaldigit *TotalDigit) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(totaldigit.GongGetIdentifier(stage), "TotalDigit", totaldigit.Name)
}

func (union *Union) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(union.GongGetIdentifier(stage), "Union", union.Name)
}

func (whitespace *WhiteSpace) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(whitespace.GongGetIdentifier(stage), "WhiteSpace", whitespace.Name)
}

// insertion point for unstaging
func (all *All) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(all.GongGetReferenceIdentifier(stage))
}

func (annotation *Annotation) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(annotation.GongGetReferenceIdentifier(stage))
}

func (attribute *Attribute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute.GongGetReferenceIdentifier(stage))
}

func (attributegroup *AttributeGroup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attributegroup.GongGetReferenceIdentifier(stage))
}

func (choice *Choice) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(choice.GongGetReferenceIdentifier(stage))
}

func (complexcontent *ComplexContent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(complexcontent.GongGetReferenceIdentifier(stage))
}

func (complextype *ComplexType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(complextype.GongGetReferenceIdentifier(stage))
}

func (documentation *Documentation) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(documentation.GongGetReferenceIdentifier(stage))
}

func (element *Element) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(element.GongGetReferenceIdentifier(stage))
}

func (enumeration *Enumeration) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(enumeration.GongGetReferenceIdentifier(stage))
}

func (extension *Extension) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(extension.GongGetReferenceIdentifier(stage))
}

func (group *Group) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group.GongGetReferenceIdentifier(stage))
}

func (length *Length) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(length.GongGetReferenceIdentifier(stage))
}

func (maxinclusive *MaxInclusive) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(maxinclusive.GongGetReferenceIdentifier(stage))
}

func (maxlength *MaxLength) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(maxlength.GongGetReferenceIdentifier(stage))
}

func (mininclusive *MinInclusive) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mininclusive.GongGetReferenceIdentifier(stage))
}

func (minlength *MinLength) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(minlength.GongGetReferenceIdentifier(stage))
}

func (pattern *Pattern) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pattern.GongGetReferenceIdentifier(stage))
}

func (restriction *Restriction) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(restriction.GongGetReferenceIdentifier(stage))
}

func (schema *Schema) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(schema.GongGetReferenceIdentifier(stage))
}

func (sequence *Sequence) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sequence.GongGetReferenceIdentifier(stage))
}

func (simplecontent *SimpleContent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(simplecontent.GongGetReferenceIdentifier(stage))
}

func (simpletype *SimpleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(simpletype.GongGetReferenceIdentifier(stage))
}

func (totaldigit *TotalDigit) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(totaldigit.GongGetReferenceIdentifier(stage))
}

func (union *Union) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(union.GongGetReferenceIdentifier(stage))
}

func (whitespace *WhiteSpace) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(whitespace.GongGetReferenceIdentifier(stage))
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
