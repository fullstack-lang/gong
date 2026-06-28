// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *All:
		ok = stage.IsStagedAll(target)

	case *Annotation:
		ok = stage.IsStagedAnnotation(target)

	case *Attribute:
		ok = stage.IsStagedAttribute(target)

	case *AttributeGroup:
		ok = stage.IsStagedAttributeGroup(target)

	case *Choice:
		ok = stage.IsStagedChoice(target)

	case *ComplexContent:
		ok = stage.IsStagedComplexContent(target)

	case *ComplexType:
		ok = stage.IsStagedComplexType(target)

	case *Documentation:
		ok = stage.IsStagedDocumentation(target)

	case *Element:
		ok = stage.IsStagedElement(target)

	case *Enumeration:
		ok = stage.IsStagedEnumeration(target)

	case *Extension:
		ok = stage.IsStagedExtension(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Length:
		ok = stage.IsStagedLength(target)

	case *MaxInclusive:
		ok = stage.IsStagedMaxInclusive(target)

	case *MaxLength:
		ok = stage.IsStagedMaxLength(target)

	case *MinInclusive:
		ok = stage.IsStagedMinInclusive(target)

	case *MinLength:
		ok = stage.IsStagedMinLength(target)

	case *Pattern:
		ok = stage.IsStagedPattern(target)

	case *Restriction:
		ok = stage.IsStagedRestriction(target)

	case *Schema:
		ok = stage.IsStagedSchema(target)

	case *Sequence:
		ok = stage.IsStagedSequence(target)

	case *SimpleContent:
		ok = stage.IsStagedSimpleContent(target)

	case *SimpleType:
		ok = stage.IsStagedSimpleType(target)

	case *TotalDigit:
		ok = stage.IsStagedTotalDigit(target)

	case *Union:
		ok = stage.IsStagedUnion(target)

	case *WhiteSpace:
		ok = stage.IsStagedWhiteSpace(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *All:
		ok = stage.IsStagedAll(target)

	case *Annotation:
		ok = stage.IsStagedAnnotation(target)

	case *Attribute:
		ok = stage.IsStagedAttribute(target)

	case *AttributeGroup:
		ok = stage.IsStagedAttributeGroup(target)

	case *Choice:
		ok = stage.IsStagedChoice(target)

	case *ComplexContent:
		ok = stage.IsStagedComplexContent(target)

	case *ComplexType:
		ok = stage.IsStagedComplexType(target)

	case *Documentation:
		ok = stage.IsStagedDocumentation(target)

	case *Element:
		ok = stage.IsStagedElement(target)

	case *Enumeration:
		ok = stage.IsStagedEnumeration(target)

	case *Extension:
		ok = stage.IsStagedExtension(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Length:
		ok = stage.IsStagedLength(target)

	case *MaxInclusive:
		ok = stage.IsStagedMaxInclusive(target)

	case *MaxLength:
		ok = stage.IsStagedMaxLength(target)

	case *MinInclusive:
		ok = stage.IsStagedMinInclusive(target)

	case *MinLength:
		ok = stage.IsStagedMinLength(target)

	case *Pattern:
		ok = stage.IsStagedPattern(target)

	case *Restriction:
		ok = stage.IsStagedRestriction(target)

	case *Schema:
		ok = stage.IsStagedSchema(target)

	case *Sequence:
		ok = stage.IsStagedSequence(target)

	case *SimpleContent:
		ok = stage.IsStagedSimpleContent(target)

	case *SimpleType:
		ok = stage.IsStagedSimpleType(target)

	case *TotalDigit:
		ok = stage.IsStagedTotalDigit(target)

	case *Union:
		ok = stage.IsStagedUnion(target)

	case *WhiteSpace:
		ok = stage.IsStagedWhiteSpace(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAll(all *All) (ok bool) {

	_, ok = stage.Alls[all]

	return
}

func (stage *Stage) IsStagedAnnotation(annotation *Annotation) (ok bool) {

	_, ok = stage.Annotations[annotation]

	return
}

func (stage *Stage) IsStagedAttribute(attribute *Attribute) (ok bool) {

	_, ok = stage.Attributes[attribute]

	return
}

func (stage *Stage) IsStagedAttributeGroup(attributegroup *AttributeGroup) (ok bool) {

	_, ok = stage.AttributeGroups[attributegroup]

	return
}

func (stage *Stage) IsStagedChoice(choice *Choice) (ok bool) {

	_, ok = stage.Choices[choice]

	return
}

func (stage *Stage) IsStagedComplexContent(complexcontent *ComplexContent) (ok bool) {

	_, ok = stage.ComplexContents[complexcontent]

	return
}

func (stage *Stage) IsStagedComplexType(complextype *ComplexType) (ok bool) {

	_, ok = stage.ComplexTypes[complextype]

	return
}

func (stage *Stage) IsStagedDocumentation(documentation *Documentation) (ok bool) {

	_, ok = stage.Documentations[documentation]

	return
}

func (stage *Stage) IsStagedElement(element *Element) (ok bool) {

	_, ok = stage.Elements[element]

	return
}

func (stage *Stage) IsStagedEnumeration(enumeration *Enumeration) (ok bool) {

	_, ok = stage.Enumerations[enumeration]

	return
}

func (stage *Stage) IsStagedExtension(extension *Extension) (ok bool) {

	_, ok = stage.Extensions[extension]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedLength(length *Length) (ok bool) {

	_, ok = stage.Lengths[length]

	return
}

func (stage *Stage) IsStagedMaxInclusive(maxinclusive *MaxInclusive) (ok bool) {

	_, ok = stage.MaxInclusives[maxinclusive]

	return
}

func (stage *Stage) IsStagedMaxLength(maxlength *MaxLength) (ok bool) {

	_, ok = stage.MaxLengths[maxlength]

	return
}

func (stage *Stage) IsStagedMinInclusive(mininclusive *MinInclusive) (ok bool) {

	_, ok = stage.MinInclusives[mininclusive]

	return
}

func (stage *Stage) IsStagedMinLength(minlength *MinLength) (ok bool) {

	_, ok = stage.MinLengths[minlength]

	return
}

func (stage *Stage) IsStagedPattern(pattern *Pattern) (ok bool) {

	_, ok = stage.Patterns[pattern]

	return
}

func (stage *Stage) IsStagedRestriction(restriction *Restriction) (ok bool) {

	_, ok = stage.Restrictions[restriction]

	return
}

func (stage *Stage) IsStagedSchema(schema *Schema) (ok bool) {

	_, ok = stage.Schemas[schema]

	return
}

func (stage *Stage) IsStagedSequence(sequence *Sequence) (ok bool) {

	_, ok = stage.Sequences[sequence]

	return
}

func (stage *Stage) IsStagedSimpleContent(simplecontent *SimpleContent) (ok bool) {

	_, ok = stage.SimpleContents[simplecontent]

	return
}

func (stage *Stage) IsStagedSimpleType(simpletype *SimpleType) (ok bool) {

	_, ok = stage.SimpleTypes[simpletype]

	return
}

func (stage *Stage) IsStagedTotalDigit(totaldigit *TotalDigit) (ok bool) {

	_, ok = stage.TotalDigits[totaldigit]

	return
}

func (stage *Stage) IsStagedUnion(union *Union) (ok bool) {

	_, ok = stage.Unions[union]

	return
}

func (stage *Stage) IsStagedWhiteSpace(whitespace *WhiteSpace) (ok bool) {

	_, ok = stage.WhiteSpaces[whitespace]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *All:
		stage.StageBranchAll(target)

	case *Annotation:
		stage.StageBranchAnnotation(target)

	case *Attribute:
		stage.StageBranchAttribute(target)

	case *AttributeGroup:
		stage.StageBranchAttributeGroup(target)

	case *Choice:
		stage.StageBranchChoice(target)

	case *ComplexContent:
		stage.StageBranchComplexContent(target)

	case *ComplexType:
		stage.StageBranchComplexType(target)

	case *Documentation:
		stage.StageBranchDocumentation(target)

	case *Element:
		stage.StageBranchElement(target)

	case *Enumeration:
		stage.StageBranchEnumeration(target)

	case *Extension:
		stage.StageBranchExtension(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *Length:
		stage.StageBranchLength(target)

	case *MaxInclusive:
		stage.StageBranchMaxInclusive(target)

	case *MaxLength:
		stage.StageBranchMaxLength(target)

	case *MinInclusive:
		stage.StageBranchMinInclusive(target)

	case *MinLength:
		stage.StageBranchMinLength(target)

	case *Pattern:
		stage.StageBranchPattern(target)

	case *Restriction:
		stage.StageBranchRestriction(target)

	case *Schema:
		stage.StageBranchSchema(target)

	case *Sequence:
		stage.StageBranchSequence(target)

	case *SimpleContent:
		stage.StageBranchSimpleContent(target)

	case *SimpleType:
		stage.StageBranchSimpleType(target)

	case *TotalDigit:
		stage.StageBranchTotalDigit(target)

	case *Union:
		stage.StageBranchUnion(target)

	case *WhiteSpace:
		stage.StageBranchWhiteSpace(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAll(all *All) {

	// check if instance is already staged
	if IsStaged(stage, all) {
		return
	}

	all.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		StageBranch(stage, all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range all.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range all.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range all.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range all.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *Stage) StageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if IsStaged(stage, annotation) {
		return
	}

	annotation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		StageBranch(stage, _documentation)
	}

}

func (stage *Stage) StageBranchAttribute(attribute *Attribute) {

	// check if instance is already staged
	if IsStaged(stage, attribute) {
		return
	}

	attribute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		StageBranch(stage, attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchAttributeGroup(attributegroup *AttributeGroup) {

	// check if instance is already staged
	if IsStaged(stage, attributegroup) {
		return
	}

	attributegroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		StageBranch(stage, attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		StageBranch(stage, _attribute)
	}

}

func (stage *Stage) StageBranchChoice(choice *Choice) {

	// check if instance is already staged
	if IsStaged(stage, choice) {
		return
	}

	choice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		StageBranch(stage, choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range choice.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range choice.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range choice.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range choice.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *Stage) StageBranchComplexContent(complexcontent *ComplexContent) {

	// check if instance is already staged
	if IsStaged(stage, complexcontent) {
		return
	}

	complexcontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchComplexType(complextype *ComplexType) {

	// check if instance is already staged
	if IsStaged(stage, complextype) {
		return
	}

	complextype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		StageBranch(stage, complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		StageBranch(stage, complextype.Annotation)
	}
	if complextype.Extension != nil {
		StageBranch(stage, complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		StageBranch(stage, complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		StageBranch(stage, complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range complextype.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range complextype.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range complextype.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range complextype.Elements {
		StageBranch(stage, _element)
	}
	for _, _attribute := range complextype.Attributes {
		StageBranch(stage, _attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}

}

func (stage *Stage) StageBranchDocumentation(documentation *Documentation) {

	// check if instance is already staged
	if IsStaged(stage, documentation) {
		return
	}

	documentation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchElement(element *Element) {

	// check if instance is already staged
	if IsStaged(stage, element) {
		return
	}

	element.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		StageBranch(stage, element.Annotation)
	}
	if element.SimpleType != nil {
		StageBranch(stage, element.SimpleType)
	}
	if element.ComplexType != nil {
		StageBranch(stage, element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		StageBranch(stage, _group)
	}

}

func (stage *Stage) StageBranchEnumeration(enumeration *Enumeration) {

	// check if instance is already staged
	if IsStaged(stage, enumeration) {
		return
	}

	enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		StageBranch(stage, enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchExtension(extension *Extension) {

	// check if instance is already staged
	if IsStaged(stage, extension) {
		return
	}

	extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range extension.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range extension.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range extension.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range extension.Elements {
		StageBranch(stage, _element)
	}
	for _, _attribute := range extension.Attributes {
		StageBranch(stage, _attribute)
	}
	for _, _attributegroup := range extension.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}

}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		StageBranch(stage, group.Annotation)
	}
	if group.OuterElement != nil {
		StageBranch(stage, group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range group.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range group.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range group.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range group.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *Stage) StageBranchLength(length *Length) {

	// check if instance is already staged
	if IsStaged(stage, length) {
		return
	}

	length.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		StageBranch(stage, length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMaxInclusive(maxinclusive *MaxInclusive) {

	// check if instance is already staged
	if IsStaged(stage, maxinclusive) {
		return
	}

	maxinclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		StageBranch(stage, maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMaxLength(maxlength *MaxLength) {

	// check if instance is already staged
	if IsStaged(stage, maxlength) {
		return
	}

	maxlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		StageBranch(stage, maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMinInclusive(mininclusive *MinInclusive) {

	// check if instance is already staged
	if IsStaged(stage, mininclusive) {
		return
	}

	mininclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		StageBranch(stage, mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMinLength(minlength *MinLength) {

	// check if instance is already staged
	if IsStaged(stage, minlength) {
		return
	}

	minlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		StageBranch(stage, minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPattern(pattern *Pattern) {

	// check if instance is already staged
	if IsStaged(stage, pattern) {
		return
	}

	pattern.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		StageBranch(stage, pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRestriction(restriction *Restriction) {

	// check if instance is already staged
	if IsStaged(stage, restriction) {
		return
	}

	restriction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		StageBranch(stage, restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		StageBranch(stage, restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		StageBranch(stage, restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		StageBranch(stage, restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		StageBranch(stage, restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		StageBranch(stage, restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		StageBranch(stage, restriction.MaxLength)
	}
	if restriction.Length != nil {
		StageBranch(stage, restriction.Length)
	}
	if restriction.TotalDigit != nil {
		StageBranch(stage, restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		StageBranch(stage, _enumeration)
	}

}

func (stage *Stage) StageBranchSchema(schema *Schema) {

	// check if instance is already staged
	if IsStaged(stage, schema) {
		return
	}

	schema.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		StageBranch(stage, schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		StageBranch(stage, _element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		StageBranch(stage, _simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		StageBranch(stage, _complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		StageBranch(stage, _attributegroup)
	}
	for _, _group := range schema.Groups {
		StageBranch(stage, _group)
	}

}

func (stage *Stage) StageBranchSequence(sequence *Sequence) {

	// check if instance is already staged
	if IsStaged(stage, sequence) {
		return
	}

	sequence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		StageBranch(stage, sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		StageBranch(stage, _sequence)
	}
	for _, _all := range sequence.Alls {
		StageBranch(stage, _all)
	}
	for _, _choice := range sequence.Choices {
		StageBranch(stage, _choice)
	}
	for _, _group := range sequence.Groups {
		StageBranch(stage, _group)
	}
	for _, _element := range sequence.Elements {
		StageBranch(stage, _element)
	}

}

func (stage *Stage) StageBranchSimpleContent(simplecontent *SimpleContent) {

	// check if instance is already staged
	if IsStaged(stage, simplecontent) {
		return
	}

	simplecontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		StageBranch(stage, simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		StageBranch(stage, simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSimpleType(simpletype *SimpleType) {

	// check if instance is already staged
	if IsStaged(stage, simpletype) {
		return
	}

	simpletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		StageBranch(stage, simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		StageBranch(stage, simpletype.Restriction)
	}
	if simpletype.Union != nil {
		StageBranch(stage, simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTotalDigit(totaldigit *TotalDigit) {

	// check if instance is already staged
	if IsStaged(stage, totaldigit) {
		return
	}

	totaldigit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		StageBranch(stage, totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchUnion(union *Union) {

	// check if instance is already staged
	if IsStaged(stage, union) {
		return
	}

	union.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		StageBranch(stage, union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchWhiteSpace(whitespace *WhiteSpace) {

	// check if instance is already staged
	if IsStaged(stage, whitespace) {
		return
	}

	whitespace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		StageBranch(stage, whitespace.Annotation)
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
	case *All:
		toT := CopyBranchAll(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Annotation:
		toT := CopyBranchAnnotation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Attribute:
		toT := CopyBranchAttribute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AttributeGroup:
		toT := CopyBranchAttributeGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Choice:
		toT := CopyBranchChoice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexContent:
		toT := CopyBranchComplexContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexType:
		toT := CopyBranchComplexType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Documentation:
		toT := CopyBranchDocumentation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Element:
		toT := CopyBranchElement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Enumeration:
		toT := CopyBranchEnumeration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extension:
		toT := CopyBranchExtension(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Length:
		toT := CopyBranchLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxInclusive:
		toT := CopyBranchMaxInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxLength:
		toT := CopyBranchMaxLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinInclusive:
		toT := CopyBranchMinInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinLength:
		toT := CopyBranchMinLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pattern:
		toT := CopyBranchPattern(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Restriction:
		toT := CopyBranchRestriction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Schema:
		toT := CopyBranchSchema(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sequence:
		toT := CopyBranchSequence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleContent:
		toT := CopyBranchSimpleContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleType:
		toT := CopyBranchSimpleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TotalDigit:
		toT := CopyBranchTotalDigit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Union:
		toT := CopyBranchUnion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *WhiteSpace:
		toT := CopyBranchWhiteSpace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAll(mapOrigCopy map[any]any, allFrom *All) (allTo *All) {

	// allFrom has already been copied
	if _allTo, ok := mapOrigCopy[allFrom]; ok {
		allTo = _allTo.(*All)
		return
	}

	allTo = new(All)
	mapOrigCopy[allFrom] = allTo
	allFrom.CopyBasicFields(allTo)

	//insertion point for the staging of instances referenced by pointers
	if allFrom.Annotation != nil {
		allTo.Annotation = CopyBranchAnnotation(mapOrigCopy, allFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range allFrom.Sequences {
		allTo.Sequences = append(allTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range allFrom.Alls {
		allTo.Alls = append(allTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range allFrom.Choices {
		allTo.Choices = append(allTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range allFrom.Groups {
		allTo.Groups = append(allTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range allFrom.Elements {
		allTo.Elements = append(allTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchAnnotation(mapOrigCopy map[any]any, annotationFrom *Annotation) (annotationTo *Annotation) {

	// annotationFrom has already been copied
	if _annotationTo, ok := mapOrigCopy[annotationFrom]; ok {
		annotationTo = _annotationTo.(*Annotation)
		return
	}

	annotationTo = new(Annotation)
	mapOrigCopy[annotationFrom] = annotationTo
	annotationFrom.CopyBasicFields(annotationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotationFrom.Documentations {
		annotationTo.Documentations = append(annotationTo.Documentations, CopyBranchDocumentation(mapOrigCopy, _documentation))
	}

	return
}

func CopyBranchAttribute(mapOrigCopy map[any]any, attributeFrom *Attribute) (attributeTo *Attribute) {

	// attributeFrom has already been copied
	if _attributeTo, ok := mapOrigCopy[attributeFrom]; ok {
		attributeTo = _attributeTo.(*Attribute)
		return
	}

	attributeTo = new(Attribute)
	mapOrigCopy[attributeFrom] = attributeTo
	attributeFrom.CopyBasicFields(attributeTo)

	//insertion point for the staging of instances referenced by pointers
	if attributeFrom.Annotation != nil {
		attributeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, attributeFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAttributeGroup(mapOrigCopy map[any]any, attributegroupFrom *AttributeGroup) (attributegroupTo *AttributeGroup) {

	// attributegroupFrom has already been copied
	if _attributegroupTo, ok := mapOrigCopy[attributegroupFrom]; ok {
		attributegroupTo = _attributegroupTo.(*AttributeGroup)
		return
	}

	attributegroupTo = new(AttributeGroup)
	mapOrigCopy[attributegroupFrom] = attributegroupTo
	attributegroupFrom.CopyBasicFields(attributegroupTo)

	//insertion point for the staging of instances referenced by pointers
	if attributegroupFrom.Annotation != nil {
		attributegroupTo.Annotation = CopyBranchAnnotation(mapOrigCopy, attributegroupFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroupFrom.AttributeGroups {
		attributegroupTo.AttributeGroups = append(attributegroupTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _attribute := range attributegroupFrom.Attributes {
		attributegroupTo.Attributes = append(attributegroupTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}

	return
}

func CopyBranchChoice(mapOrigCopy map[any]any, choiceFrom *Choice) (choiceTo *Choice) {

	// choiceFrom has already been copied
	if _choiceTo, ok := mapOrigCopy[choiceFrom]; ok {
		choiceTo = _choiceTo.(*Choice)
		return
	}

	choiceTo = new(Choice)
	mapOrigCopy[choiceFrom] = choiceTo
	choiceFrom.CopyBasicFields(choiceTo)

	//insertion point for the staging of instances referenced by pointers
	if choiceFrom.Annotation != nil {
		choiceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, choiceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choiceFrom.Sequences {
		choiceTo.Sequences = append(choiceTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range choiceFrom.Alls {
		choiceTo.Alls = append(choiceTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range choiceFrom.Choices {
		choiceTo.Choices = append(choiceTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range choiceFrom.Groups {
		choiceTo.Groups = append(choiceTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range choiceFrom.Elements {
		choiceTo.Elements = append(choiceTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchComplexContent(mapOrigCopy map[any]any, complexcontentFrom *ComplexContent) (complexcontentTo *ComplexContent) {

	// complexcontentFrom has already been copied
	if _complexcontentTo, ok := mapOrigCopy[complexcontentFrom]; ok {
		complexcontentTo = _complexcontentTo.(*ComplexContent)
		return
	}

	complexcontentTo = new(ComplexContent)
	mapOrigCopy[complexcontentFrom] = complexcontentTo
	complexcontentFrom.CopyBasicFields(complexcontentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchComplexType(mapOrigCopy map[any]any, complextypeFrom *ComplexType) (complextypeTo *ComplexType) {

	// complextypeFrom has already been copied
	if _complextypeTo, ok := mapOrigCopy[complextypeFrom]; ok {
		complextypeTo = _complextypeTo.(*ComplexType)
		return
	}

	complextypeTo = new(ComplexType)
	mapOrigCopy[complextypeFrom] = complextypeTo
	complextypeFrom.CopyBasicFields(complextypeTo)

	//insertion point for the staging of instances referenced by pointers
	if complextypeFrom.OuterElement != nil {
		complextypeTo.OuterElement = CopyBranchElement(mapOrigCopy, complextypeFrom.OuterElement)
	}
	if complextypeFrom.Annotation != nil {
		complextypeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, complextypeFrom.Annotation)
	}
	if complextypeFrom.Extension != nil {
		complextypeTo.Extension = CopyBranchExtension(mapOrigCopy, complextypeFrom.Extension)
	}
	if complextypeFrom.SimpleContent != nil {
		complextypeTo.SimpleContent = CopyBranchSimpleContent(mapOrigCopy, complextypeFrom.SimpleContent)
	}
	if complextypeFrom.ComplexContent != nil {
		complextypeTo.ComplexContent = CopyBranchComplexContent(mapOrigCopy, complextypeFrom.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextypeFrom.Sequences {
		complextypeTo.Sequences = append(complextypeTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range complextypeFrom.Alls {
		complextypeTo.Alls = append(complextypeTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range complextypeFrom.Choices {
		complextypeTo.Choices = append(complextypeTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range complextypeFrom.Groups {
		complextypeTo.Groups = append(complextypeTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range complextypeFrom.Elements {
		complextypeTo.Elements = append(complextypeTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range complextypeFrom.Attributes {
		complextypeTo.Attributes = append(complextypeTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}
	for _, _attributegroup := range complextypeFrom.AttributeGroups {
		complextypeTo.AttributeGroups = append(complextypeTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}

	return
}

func CopyBranchDocumentation(mapOrigCopy map[any]any, documentationFrom *Documentation) (documentationTo *Documentation) {

	// documentationFrom has already been copied
	if _documentationTo, ok := mapOrigCopy[documentationFrom]; ok {
		documentationTo = _documentationTo.(*Documentation)
		return
	}

	documentationTo = new(Documentation)
	mapOrigCopy[documentationFrom] = documentationTo
	documentationFrom.CopyBasicFields(documentationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchElement(mapOrigCopy map[any]any, elementFrom *Element) (elementTo *Element) {

	// elementFrom has already been copied
	if _elementTo, ok := mapOrigCopy[elementFrom]; ok {
		elementTo = _elementTo.(*Element)
		return
	}

	elementTo = new(Element)
	mapOrigCopy[elementFrom] = elementTo
	elementFrom.CopyBasicFields(elementTo)

	//insertion point for the staging of instances referenced by pointers
	if elementFrom.Annotation != nil {
		elementTo.Annotation = CopyBranchAnnotation(mapOrigCopy, elementFrom.Annotation)
	}
	if elementFrom.SimpleType != nil {
		elementTo.SimpleType = CopyBranchSimpleType(mapOrigCopy, elementFrom.SimpleType)
	}
	if elementFrom.ComplexType != nil {
		elementTo.ComplexType = CopyBranchComplexType(mapOrigCopy, elementFrom.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range elementFrom.Groups {
		elementTo.Groups = append(elementTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func CopyBranchEnumeration(mapOrigCopy map[any]any, enumerationFrom *Enumeration) (enumerationTo *Enumeration) {

	// enumerationFrom has already been copied
	if _enumerationTo, ok := mapOrigCopy[enumerationFrom]; ok {
		enumerationTo = _enumerationTo.(*Enumeration)
		return
	}

	enumerationTo = new(Enumeration)
	mapOrigCopy[enumerationFrom] = enumerationTo
	enumerationFrom.CopyBasicFields(enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if enumerationFrom.Annotation != nil {
		enumerationTo.Annotation = CopyBranchAnnotation(mapOrigCopy, enumerationFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExtension(mapOrigCopy map[any]any, extensionFrom *Extension) (extensionTo *Extension) {

	// extensionFrom has already been copied
	if _extensionTo, ok := mapOrigCopy[extensionFrom]; ok {
		extensionTo = _extensionTo.(*Extension)
		return
	}

	extensionTo = new(Extension)
	mapOrigCopy[extensionFrom] = extensionTo
	extensionFrom.CopyBasicFields(extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extensionFrom.Sequences {
		extensionTo.Sequences = append(extensionTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range extensionFrom.Alls {
		extensionTo.Alls = append(extensionTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range extensionFrom.Choices {
		extensionTo.Choices = append(extensionTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range extensionFrom.Groups {
		extensionTo.Groups = append(extensionTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range extensionFrom.Elements {
		extensionTo.Elements = append(extensionTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range extensionFrom.Attributes {
		extensionTo.Attributes = append(extensionTo.Attributes, CopyBranchAttribute(mapOrigCopy, _attribute))
	}
	for _, _attributegroup := range extensionFrom.AttributeGroups {
		extensionTo.AttributeGroups = append(extensionTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}

	return
}

func CopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.CopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers
	if groupFrom.Annotation != nil {
		groupTo.Annotation = CopyBranchAnnotation(mapOrigCopy, groupFrom.Annotation)
	}
	if groupFrom.OuterElement != nil {
		groupTo.OuterElement = CopyBranchElement(mapOrigCopy, groupFrom.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range groupFrom.Sequences {
		groupTo.Sequences = append(groupTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range groupFrom.Alls {
		groupTo.Alls = append(groupTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range groupFrom.Choices {
		groupTo.Choices = append(groupTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range groupFrom.Groups {
		groupTo.Groups = append(groupTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range groupFrom.Elements {
		groupTo.Elements = append(groupTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchLength(mapOrigCopy map[any]any, lengthFrom *Length) (lengthTo *Length) {

	// lengthFrom has already been copied
	if _lengthTo, ok := mapOrigCopy[lengthFrom]; ok {
		lengthTo = _lengthTo.(*Length)
		return
	}

	lengthTo = new(Length)
	mapOrigCopy[lengthFrom] = lengthTo
	lengthFrom.CopyBasicFields(lengthTo)

	//insertion point for the staging of instances referenced by pointers
	if lengthFrom.Annotation != nil {
		lengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, lengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMaxInclusive(mapOrigCopy map[any]any, maxinclusiveFrom *MaxInclusive) (maxinclusiveTo *MaxInclusive) {

	// maxinclusiveFrom has already been copied
	if _maxinclusiveTo, ok := mapOrigCopy[maxinclusiveFrom]; ok {
		maxinclusiveTo = _maxinclusiveTo.(*MaxInclusive)
		return
	}

	maxinclusiveTo = new(MaxInclusive)
	mapOrigCopy[maxinclusiveFrom] = maxinclusiveTo
	maxinclusiveFrom.CopyBasicFields(maxinclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusiveFrom.Annotation != nil {
		maxinclusiveTo.Annotation = CopyBranchAnnotation(mapOrigCopy, maxinclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMaxLength(mapOrigCopy map[any]any, maxlengthFrom *MaxLength) (maxlengthTo *MaxLength) {

	// maxlengthFrom has already been copied
	if _maxlengthTo, ok := mapOrigCopy[maxlengthFrom]; ok {
		maxlengthTo = _maxlengthTo.(*MaxLength)
		return
	}

	maxlengthTo = new(MaxLength)
	mapOrigCopy[maxlengthFrom] = maxlengthTo
	maxlengthFrom.CopyBasicFields(maxlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if maxlengthFrom.Annotation != nil {
		maxlengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, maxlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMinInclusive(mapOrigCopy map[any]any, mininclusiveFrom *MinInclusive) (mininclusiveTo *MinInclusive) {

	// mininclusiveFrom has already been copied
	if _mininclusiveTo, ok := mapOrigCopy[mininclusiveFrom]; ok {
		mininclusiveTo = _mininclusiveTo.(*MinInclusive)
		return
	}

	mininclusiveTo = new(MinInclusive)
	mapOrigCopy[mininclusiveFrom] = mininclusiveTo
	mininclusiveFrom.CopyBasicFields(mininclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if mininclusiveFrom.Annotation != nil {
		mininclusiveTo.Annotation = CopyBranchAnnotation(mapOrigCopy, mininclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMinLength(mapOrigCopy map[any]any, minlengthFrom *MinLength) (minlengthTo *MinLength) {

	// minlengthFrom has already been copied
	if _minlengthTo, ok := mapOrigCopy[minlengthFrom]; ok {
		minlengthTo = _minlengthTo.(*MinLength)
		return
	}

	minlengthTo = new(MinLength)
	mapOrigCopy[minlengthFrom] = minlengthTo
	minlengthFrom.CopyBasicFields(minlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if minlengthFrom.Annotation != nil {
		minlengthTo.Annotation = CopyBranchAnnotation(mapOrigCopy, minlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPattern(mapOrigCopy map[any]any, patternFrom *Pattern) (patternTo *Pattern) {

	// patternFrom has already been copied
	if _patternTo, ok := mapOrigCopy[patternFrom]; ok {
		patternTo = _patternTo.(*Pattern)
		return
	}

	patternTo = new(Pattern)
	mapOrigCopy[patternFrom] = patternTo
	patternFrom.CopyBasicFields(patternTo)

	//insertion point for the staging of instances referenced by pointers
	if patternFrom.Annotation != nil {
		patternTo.Annotation = CopyBranchAnnotation(mapOrigCopy, patternFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRestriction(mapOrigCopy map[any]any, restrictionFrom *Restriction) (restrictionTo *Restriction) {

	// restrictionFrom has already been copied
	if _restrictionTo, ok := mapOrigCopy[restrictionFrom]; ok {
		restrictionTo = _restrictionTo.(*Restriction)
		return
	}

	restrictionTo = new(Restriction)
	mapOrigCopy[restrictionFrom] = restrictionTo
	restrictionFrom.CopyBasicFields(restrictionTo)

	//insertion point for the staging of instances referenced by pointers
	if restrictionFrom.Annotation != nil {
		restrictionTo.Annotation = CopyBranchAnnotation(mapOrigCopy, restrictionFrom.Annotation)
	}
	if restrictionFrom.MinInclusive != nil {
		restrictionTo.MinInclusive = CopyBranchMinInclusive(mapOrigCopy, restrictionFrom.MinInclusive)
	}
	if restrictionFrom.MaxInclusive != nil {
		restrictionTo.MaxInclusive = CopyBranchMaxInclusive(mapOrigCopy, restrictionFrom.MaxInclusive)
	}
	if restrictionFrom.Pattern != nil {
		restrictionTo.Pattern = CopyBranchPattern(mapOrigCopy, restrictionFrom.Pattern)
	}
	if restrictionFrom.WhiteSpace != nil {
		restrictionTo.WhiteSpace = CopyBranchWhiteSpace(mapOrigCopy, restrictionFrom.WhiteSpace)
	}
	if restrictionFrom.MinLength != nil {
		restrictionTo.MinLength = CopyBranchMinLength(mapOrigCopy, restrictionFrom.MinLength)
	}
	if restrictionFrom.MaxLength != nil {
		restrictionTo.MaxLength = CopyBranchMaxLength(mapOrigCopy, restrictionFrom.MaxLength)
	}
	if restrictionFrom.Length != nil {
		restrictionTo.Length = CopyBranchLength(mapOrigCopy, restrictionFrom.Length)
	}
	if restrictionFrom.TotalDigit != nil {
		restrictionTo.TotalDigit = CopyBranchTotalDigit(mapOrigCopy, restrictionFrom.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restrictionFrom.Enumerations {
		restrictionTo.Enumerations = append(restrictionTo.Enumerations, CopyBranchEnumeration(mapOrigCopy, _enumeration))
	}

	return
}

func CopyBranchSchema(mapOrigCopy map[any]any, schemaFrom *Schema) (schemaTo *Schema) {

	// schemaFrom has already been copied
	if _schemaTo, ok := mapOrigCopy[schemaFrom]; ok {
		schemaTo = _schemaTo.(*Schema)
		return
	}

	schemaTo = new(Schema)
	mapOrigCopy[schemaFrom] = schemaTo
	schemaFrom.CopyBasicFields(schemaTo)

	//insertion point for the staging of instances referenced by pointers
	if schemaFrom.Annotation != nil {
		schemaTo.Annotation = CopyBranchAnnotation(mapOrigCopy, schemaFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schemaFrom.Elements {
		schemaTo.Elements = append(schemaTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}
	for _, _simpletype := range schemaFrom.SimpleTypes {
		schemaTo.SimpleTypes = append(schemaTo.SimpleTypes, CopyBranchSimpleType(mapOrigCopy, _simpletype))
	}
	for _, _complextype := range schemaFrom.ComplexTypes {
		schemaTo.ComplexTypes = append(schemaTo.ComplexTypes, CopyBranchComplexType(mapOrigCopy, _complextype))
	}
	for _, _attributegroup := range schemaFrom.AttributeGroups {
		schemaTo.AttributeGroups = append(schemaTo.AttributeGroups, CopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _group := range schemaFrom.Groups {
		schemaTo.Groups = append(schemaTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func CopyBranchSequence(mapOrigCopy map[any]any, sequenceFrom *Sequence) (sequenceTo *Sequence) {

	// sequenceFrom has already been copied
	if _sequenceTo, ok := mapOrigCopy[sequenceFrom]; ok {
		sequenceTo = _sequenceTo.(*Sequence)
		return
	}

	sequenceTo = new(Sequence)
	mapOrigCopy[sequenceFrom] = sequenceTo
	sequenceFrom.CopyBasicFields(sequenceTo)

	//insertion point for the staging of instances referenced by pointers
	if sequenceFrom.Annotation != nil {
		sequenceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, sequenceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequenceFrom.Sequences {
		sequenceTo.Sequences = append(sequenceTo.Sequences, CopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range sequenceFrom.Alls {
		sequenceTo.Alls = append(sequenceTo.Alls, CopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range sequenceFrom.Choices {
		sequenceTo.Choices = append(sequenceTo.Choices, CopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range sequenceFrom.Groups {
		sequenceTo.Groups = append(sequenceTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range sequenceFrom.Elements {
		sequenceTo.Elements = append(sequenceTo.Elements, CopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func CopyBranchSimpleContent(mapOrigCopy map[any]any, simplecontentFrom *SimpleContent) (simplecontentTo *SimpleContent) {

	// simplecontentFrom has already been copied
	if _simplecontentTo, ok := mapOrigCopy[simplecontentFrom]; ok {
		simplecontentTo = _simplecontentTo.(*SimpleContent)
		return
	}

	simplecontentTo = new(SimpleContent)
	mapOrigCopy[simplecontentFrom] = simplecontentTo
	simplecontentFrom.CopyBasicFields(simplecontentTo)

	//insertion point for the staging of instances referenced by pointers
	if simplecontentFrom.Extension != nil {
		simplecontentTo.Extension = CopyBranchExtension(mapOrigCopy, simplecontentFrom.Extension)
	}
	if simplecontentFrom.Restriction != nil {
		simplecontentTo.Restriction = CopyBranchRestriction(mapOrigCopy, simplecontentFrom.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSimpleType(mapOrigCopy map[any]any, simpletypeFrom *SimpleType) (simpletypeTo *SimpleType) {

	// simpletypeFrom has already been copied
	if _simpletypeTo, ok := mapOrigCopy[simpletypeFrom]; ok {
		simpletypeTo = _simpletypeTo.(*SimpleType)
		return
	}

	simpletypeTo = new(SimpleType)
	mapOrigCopy[simpletypeFrom] = simpletypeTo
	simpletypeFrom.CopyBasicFields(simpletypeTo)

	//insertion point for the staging of instances referenced by pointers
	if simpletypeFrom.Annotation != nil {
		simpletypeTo.Annotation = CopyBranchAnnotation(mapOrigCopy, simpletypeFrom.Annotation)
	}
	if simpletypeFrom.Restriction != nil {
		simpletypeTo.Restriction = CopyBranchRestriction(mapOrigCopy, simpletypeFrom.Restriction)
	}
	if simpletypeFrom.Union != nil {
		simpletypeTo.Union = CopyBranchUnion(mapOrigCopy, simpletypeFrom.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTotalDigit(mapOrigCopy map[any]any, totaldigitFrom *TotalDigit) (totaldigitTo *TotalDigit) {

	// totaldigitFrom has already been copied
	if _totaldigitTo, ok := mapOrigCopy[totaldigitFrom]; ok {
		totaldigitTo = _totaldigitTo.(*TotalDigit)
		return
	}

	totaldigitTo = new(TotalDigit)
	mapOrigCopy[totaldigitFrom] = totaldigitTo
	totaldigitFrom.CopyBasicFields(totaldigitTo)

	//insertion point for the staging of instances referenced by pointers
	if totaldigitFrom.Annotation != nil {
		totaldigitTo.Annotation = CopyBranchAnnotation(mapOrigCopy, totaldigitFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUnion(mapOrigCopy map[any]any, unionFrom *Union) (unionTo *Union) {

	// unionFrom has already been copied
	if _unionTo, ok := mapOrigCopy[unionFrom]; ok {
		unionTo = _unionTo.(*Union)
		return
	}

	unionTo = new(Union)
	mapOrigCopy[unionFrom] = unionTo
	unionFrom.CopyBasicFields(unionTo)

	//insertion point for the staging of instances referenced by pointers
	if unionFrom.Annotation != nil {
		unionTo.Annotation = CopyBranchAnnotation(mapOrigCopy, unionFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWhiteSpace(mapOrigCopy map[any]any, whitespaceFrom *WhiteSpace) (whitespaceTo *WhiteSpace) {

	// whitespaceFrom has already been copied
	if _whitespaceTo, ok := mapOrigCopy[whitespaceFrom]; ok {
		whitespaceTo = _whitespaceTo.(*WhiteSpace)
		return
	}

	whitespaceTo = new(WhiteSpace)
	mapOrigCopy[whitespaceFrom] = whitespaceTo
	whitespaceFrom.CopyBasicFields(whitespaceTo)

	//insertion point for the staging of instances referenced by pointers
	if whitespaceFrom.Annotation != nil {
		whitespaceTo.Annotation = CopyBranchAnnotation(mapOrigCopy, whitespaceFrom.Annotation)
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
	case *All:
		stage.UnstageBranchAll(target)

	case *Annotation:
		stage.UnstageBranchAnnotation(target)

	case *Attribute:
		stage.UnstageBranchAttribute(target)

	case *AttributeGroup:
		stage.UnstageBranchAttributeGroup(target)

	case *Choice:
		stage.UnstageBranchChoice(target)

	case *ComplexContent:
		stage.UnstageBranchComplexContent(target)

	case *ComplexType:
		stage.UnstageBranchComplexType(target)

	case *Documentation:
		stage.UnstageBranchDocumentation(target)

	case *Element:
		stage.UnstageBranchElement(target)

	case *Enumeration:
		stage.UnstageBranchEnumeration(target)

	case *Extension:
		stage.UnstageBranchExtension(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *Length:
		stage.UnstageBranchLength(target)

	case *MaxInclusive:
		stage.UnstageBranchMaxInclusive(target)

	case *MaxLength:
		stage.UnstageBranchMaxLength(target)

	case *MinInclusive:
		stage.UnstageBranchMinInclusive(target)

	case *MinLength:
		stage.UnstageBranchMinLength(target)

	case *Pattern:
		stage.UnstageBranchPattern(target)

	case *Restriction:
		stage.UnstageBranchRestriction(target)

	case *Schema:
		stage.UnstageBranchSchema(target)

	case *Sequence:
		stage.UnstageBranchSequence(target)

	case *SimpleContent:
		stage.UnstageBranchSimpleContent(target)

	case *SimpleType:
		stage.UnstageBranchSimpleType(target)

	case *TotalDigit:
		stage.UnstageBranchTotalDigit(target)

	case *Union:
		stage.UnstageBranchUnion(target)

	case *WhiteSpace:
		stage.UnstageBranchWhiteSpace(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAll(all *All) {

	// check if instance is already staged
	if !IsStaged(stage, all) {
		return
	}

	all.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		UnstageBranch(stage, all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range all.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range all.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range all.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range all.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *Stage) UnstageBranchAnnotation(annotation *Annotation) {

	// check if instance is already staged
	if !IsStaged(stage, annotation) {
		return
	}

	annotation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		UnstageBranch(stage, _documentation)
	}

}

func (stage *Stage) UnstageBranchAttribute(attribute *Attribute) {

	// check if instance is already staged
	if !IsStaged(stage, attribute) {
		return
	}

	attribute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		UnstageBranch(stage, attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchAttributeGroup(attributegroup *AttributeGroup) {

	// check if instance is already staged
	if !IsStaged(stage, attributegroup) {
		return
	}

	attributegroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		UnstageBranch(stage, attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		UnstageBranch(stage, _attribute)
	}

}

func (stage *Stage) UnstageBranchChoice(choice *Choice) {

	// check if instance is already staged
	if !IsStaged(stage, choice) {
		return
	}

	choice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		UnstageBranch(stage, choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range choice.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range choice.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range choice.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range choice.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *Stage) UnstageBranchComplexContent(complexcontent *ComplexContent) {

	// check if instance is already staged
	if !IsStaged(stage, complexcontent) {
		return
	}

	complexcontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchComplexType(complextype *ComplexType) {

	// check if instance is already staged
	if !IsStaged(stage, complextype) {
		return
	}

	complextype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		UnstageBranch(stage, complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		UnstageBranch(stage, complextype.Annotation)
	}
	if complextype.Extension != nil {
		UnstageBranch(stage, complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		UnstageBranch(stage, complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		UnstageBranch(stage, complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range complextype.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range complextype.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range complextype.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range complextype.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _attribute := range complextype.Attributes {
		UnstageBranch(stage, _attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}

}

func (stage *Stage) UnstageBranchDocumentation(documentation *Documentation) {

	// check if instance is already staged
	if !IsStaged(stage, documentation) {
		return
	}

	documentation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchElement(element *Element) {

	// check if instance is already staged
	if !IsStaged(stage, element) {
		return
	}

	element.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		UnstageBranch(stage, element.Annotation)
	}
	if element.SimpleType != nil {
		UnstageBranch(stage, element.SimpleType)
	}
	if element.ComplexType != nil {
		UnstageBranch(stage, element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		UnstageBranch(stage, _group)
	}

}

func (stage *Stage) UnstageBranchEnumeration(enumeration *Enumeration) {

	// check if instance is already staged
	if !IsStaged(stage, enumeration) {
		return
	}

	enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		UnstageBranch(stage, enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchExtension(extension *Extension) {

	// check if instance is already staged
	if !IsStaged(stage, extension) {
		return
	}

	extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range extension.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range extension.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range extension.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range extension.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _attribute := range extension.Attributes {
		UnstageBranch(stage, _attribute)
	}
	for _, _attributegroup := range extension.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}

}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		UnstageBranch(stage, group.Annotation)
	}
	if group.OuterElement != nil {
		UnstageBranch(stage, group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range group.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range group.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range group.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range group.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *Stage) UnstageBranchLength(length *Length) {

	// check if instance is already staged
	if !IsStaged(stage, length) {
		return
	}

	length.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		UnstageBranch(stage, length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMaxInclusive(maxinclusive *MaxInclusive) {

	// check if instance is already staged
	if !IsStaged(stage, maxinclusive) {
		return
	}

	maxinclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		UnstageBranch(stage, maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMaxLength(maxlength *MaxLength) {

	// check if instance is already staged
	if !IsStaged(stage, maxlength) {
		return
	}

	maxlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		UnstageBranch(stage, maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMinInclusive(mininclusive *MinInclusive) {

	// check if instance is already staged
	if !IsStaged(stage, mininclusive) {
		return
	}

	mininclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		UnstageBranch(stage, mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMinLength(minlength *MinLength) {

	// check if instance is already staged
	if !IsStaged(stage, minlength) {
		return
	}

	minlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		UnstageBranch(stage, minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPattern(pattern *Pattern) {

	// check if instance is already staged
	if !IsStaged(stage, pattern) {
		return
	}

	pattern.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		UnstageBranch(stage, pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRestriction(restriction *Restriction) {

	// check if instance is already staged
	if !IsStaged(stage, restriction) {
		return
	}

	restriction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		UnstageBranch(stage, restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		UnstageBranch(stage, restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		UnstageBranch(stage, restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		UnstageBranch(stage, restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		UnstageBranch(stage, restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		UnstageBranch(stage, restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		UnstageBranch(stage, restriction.MaxLength)
	}
	if restriction.Length != nil {
		UnstageBranch(stage, restriction.Length)
	}
	if restriction.TotalDigit != nil {
		UnstageBranch(stage, restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		UnstageBranch(stage, _enumeration)
	}

}

func (stage *Stage) UnstageBranchSchema(schema *Schema) {

	// check if instance is already staged
	if !IsStaged(stage, schema) {
		return
	}

	schema.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		UnstageBranch(stage, schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		UnstageBranch(stage, _element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		UnstageBranch(stage, _simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		UnstageBranch(stage, _complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		UnstageBranch(stage, _attributegroup)
	}
	for _, _group := range schema.Groups {
		UnstageBranch(stage, _group)
	}

}

func (stage *Stage) UnstageBranchSequence(sequence *Sequence) {

	// check if instance is already staged
	if !IsStaged(stage, sequence) {
		return
	}

	sequence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		UnstageBranch(stage, sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		UnstageBranch(stage, _sequence)
	}
	for _, _all := range sequence.Alls {
		UnstageBranch(stage, _all)
	}
	for _, _choice := range sequence.Choices {
		UnstageBranch(stage, _choice)
	}
	for _, _group := range sequence.Groups {
		UnstageBranch(stage, _group)
	}
	for _, _element := range sequence.Elements {
		UnstageBranch(stage, _element)
	}

}

func (stage *Stage) UnstageBranchSimpleContent(simplecontent *SimpleContent) {

	// check if instance is already staged
	if !IsStaged(stage, simplecontent) {
		return
	}

	simplecontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		UnstageBranch(stage, simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		UnstageBranch(stage, simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSimpleType(simpletype *SimpleType) {

	// check if instance is already staged
	if !IsStaged(stage, simpletype) {
		return
	}

	simpletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		UnstageBranch(stage, simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		UnstageBranch(stage, simpletype.Restriction)
	}
	if simpletype.Union != nil {
		UnstageBranch(stage, simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTotalDigit(totaldigit *TotalDigit) {

	// check if instance is already staged
	if !IsStaged(stage, totaldigit) {
		return
	}

	totaldigit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		UnstageBranch(stage, totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchUnion(union *Union) {

	// check if instance is already staged
	if !IsStaged(stage, union) {
		return
	}

	union.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		UnstageBranch(stage, union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchWhiteSpace(whitespace *WhiteSpace) {

	// check if instance is already staged
	if !IsStaged(stage, whitespace) {
		return
	}

	whitespace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		UnstageBranch(stage, whitespace.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *All) GongReconstructPointersFromReferences(stage *Stage, instance *All) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
}

func (reference *Annotation) GongReconstructPointersFromReferences(stage *Stage, instance *Annotation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Documentations = reference.Documentations[:0]
	for _, _b := range instance.Documentations {
		reference.Documentations = append(reference.Documentations, stage.Documentations_reference[_b])
	}
}

func (reference *Attribute) GongReconstructPointersFromReferences(stage *Stage, instance *Attribute) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *AttributeGroup) GongReconstructPointersFromReferences(stage *Stage, instance *AttributeGroup) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
	reference.AttributeGroups = reference.AttributeGroups[:0]
	for _, _b := range instance.AttributeGroups {
		reference.AttributeGroups = append(reference.AttributeGroups, stage.AttributeGroups_reference[_b])
	}
	reference.Attributes = reference.Attributes[:0]
	for _, _b := range instance.Attributes {
		reference.Attributes = append(reference.Attributes, stage.Attributes_reference[_b])
	}
}

func (reference *Choice) GongReconstructPointersFromReferences(stage *Stage, instance *Choice) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
}

func (reference *ComplexContent) GongReconstructPointersFromReferences(stage *Stage, instance *ComplexContent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ComplexType) GongReconstructPointersFromReferences(stage *Stage, instance *ComplexType) {
	// insertion point for pointers field
	if instance.OuterElement != nil {
		reference.OuterElement = stage.Elements_reference[instance.OuterElement]
	}
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	if instance.Extension != nil {
		reference.Extension = stage.Extensions_reference[instance.Extension]
	}
	if instance.SimpleContent != nil {
		reference.SimpleContent = stage.SimpleContents_reference[instance.SimpleContent]
	}
	if instance.ComplexContent != nil {
		reference.ComplexContent = stage.ComplexContents_reference[instance.ComplexContent]
	}
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
	reference.Attributes = reference.Attributes[:0]
	for _, _b := range instance.Attributes {
		reference.Attributes = append(reference.Attributes, stage.Attributes_reference[_b])
	}
	reference.AttributeGroups = reference.AttributeGroups[:0]
	for _, _b := range instance.AttributeGroups {
		reference.AttributeGroups = append(reference.AttributeGroups, stage.AttributeGroups_reference[_b])
	}
}

func (reference *Documentation) GongReconstructPointersFromReferences(stage *Stage, instance *Documentation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Element) GongReconstructPointersFromReferences(stage *Stage, instance *Element) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	if instance.SimpleType != nil {
		reference.SimpleType = stage.SimpleTypes_reference[instance.SimpleType]
	}
	if instance.ComplexType != nil {
		reference.ComplexType = stage.ComplexTypes_reference[instance.ComplexType]
	}
	// insertion point for slice of pointers field
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
}

func (reference *Enumeration) GongReconstructPointersFromReferences(stage *Stage, instance *Enumeration) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *Extension) GongReconstructPointersFromReferences(stage *Stage, instance *Extension) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
	reference.Attributes = reference.Attributes[:0]
	for _, _b := range instance.Attributes {
		reference.Attributes = append(reference.Attributes, stage.Attributes_reference[_b])
	}
	reference.AttributeGroups = reference.AttributeGroups[:0]
	for _, _b := range instance.AttributeGroups {
		reference.AttributeGroups = append(reference.AttributeGroups, stage.AttributeGroups_reference[_b])
	}
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	if instance.OuterElement != nil {
		reference.OuterElement = stage.Elements_reference[instance.OuterElement]
	}
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
}

func (reference *Length) GongReconstructPointersFromReferences(stage *Stage, instance *Length) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *MaxInclusive) GongReconstructPointersFromReferences(stage *Stage, instance *MaxInclusive) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *MaxLength) GongReconstructPointersFromReferences(stage *Stage, instance *MaxLength) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *MinInclusive) GongReconstructPointersFromReferences(stage *Stage, instance *MinInclusive) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *MinLength) GongReconstructPointersFromReferences(stage *Stage, instance *MinLength) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *Pattern) GongReconstructPointersFromReferences(stage *Stage, instance *Pattern) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *Restriction) GongReconstructPointersFromReferences(stage *Stage, instance *Restriction) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	if instance.MinInclusive != nil {
		reference.MinInclusive = stage.MinInclusives_reference[instance.MinInclusive]
	}
	if instance.MaxInclusive != nil {
		reference.MaxInclusive = stage.MaxInclusives_reference[instance.MaxInclusive]
	}
	if instance.Pattern != nil {
		reference.Pattern = stage.Patterns_reference[instance.Pattern]
	}
	if instance.WhiteSpace != nil {
		reference.WhiteSpace = stage.WhiteSpaces_reference[instance.WhiteSpace]
	}
	if instance.MinLength != nil {
		reference.MinLength = stage.MinLengths_reference[instance.MinLength]
	}
	if instance.MaxLength != nil {
		reference.MaxLength = stage.MaxLengths_reference[instance.MaxLength]
	}
	if instance.Length != nil {
		reference.Length = stage.Lengths_reference[instance.Length]
	}
	if instance.TotalDigit != nil {
		reference.TotalDigit = stage.TotalDigits_reference[instance.TotalDigit]
	}
	// insertion point for slice of pointers field
	reference.Enumerations = reference.Enumerations[:0]
	for _, _b := range instance.Enumerations {
		reference.Enumerations = append(reference.Enumerations, stage.Enumerations_reference[_b])
	}
}

func (reference *Schema) GongReconstructPointersFromReferences(stage *Stage, instance *Schema) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
	reference.SimpleTypes = reference.SimpleTypes[:0]
	for _, _b := range instance.SimpleTypes {
		reference.SimpleTypes = append(reference.SimpleTypes, stage.SimpleTypes_reference[_b])
	}
	reference.ComplexTypes = reference.ComplexTypes[:0]
	for _, _b := range instance.ComplexTypes {
		reference.ComplexTypes = append(reference.ComplexTypes, stage.ComplexTypes_reference[_b])
	}
	reference.AttributeGroups = reference.AttributeGroups[:0]
	for _, _b := range instance.AttributeGroups {
		reference.AttributeGroups = append(reference.AttributeGroups, stage.AttributeGroups_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
}

func (reference *Sequence) GongReconstructPointersFromReferences(stage *Stage, instance *Sequence) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
	reference.Sequences = reference.Sequences[:0]
	for _, _b := range instance.Sequences {
		reference.Sequences = append(reference.Sequences, stage.Sequences_reference[_b])
	}
	reference.Alls = reference.Alls[:0]
	for _, _b := range instance.Alls {
		reference.Alls = append(reference.Alls, stage.Alls_reference[_b])
	}
	reference.Choices = reference.Choices[:0]
	for _, _b := range instance.Choices {
		reference.Choices = append(reference.Choices, stage.Choices_reference[_b])
	}
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.Elements = reference.Elements[:0]
	for _, _b := range instance.Elements {
		reference.Elements = append(reference.Elements, stage.Elements_reference[_b])
	}
}

func (reference *SimpleContent) GongReconstructPointersFromReferences(stage *Stage, instance *SimpleContent) {
	// insertion point for pointers field
	if instance.Extension != nil {
		reference.Extension = stage.Extensions_reference[instance.Extension]
	}
	if instance.Restriction != nil {
		reference.Restriction = stage.Restrictions_reference[instance.Restriction]
	}
	// insertion point for slice of pointers field
}

func (reference *SimpleType) GongReconstructPointersFromReferences(stage *Stage, instance *SimpleType) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	if instance.Restriction != nil {
		reference.Restriction = stage.Restrictions_reference[instance.Restriction]
	}
	if instance.Union != nil {
		reference.Union = stage.Unions_reference[instance.Union]
	}
	// insertion point for slice of pointers field
}

func (reference *TotalDigit) GongReconstructPointersFromReferences(stage *Stage, instance *TotalDigit) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *Union) GongReconstructPointersFromReferences(stage *Stage, instance *Union) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

func (reference *WhiteSpace) GongReconstructPointersFromReferences(stage *Stage, instance *WhiteSpace) {
	// insertion point for pointers field
	if instance.Annotation != nil {
		reference.Annotation = stage.Annotations_reference[instance.Annotation]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *All) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
}

func (reference *Annotation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Documentations []*Documentation
	for _, _reference := range reference.Documentations {
		if _instance, ok := stage.Documentations_instance[_reference]; ok {
			_Documentations = append(_Documentations, _instance)
		}
	}
	reference.Documentations = _Documentations
}

func (reference *Attribute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *AttributeGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _AttributeGroups []*AttributeGroup
	for _, _reference := range reference.AttributeGroups {
		if _instance, ok := stage.AttributeGroups_instance[_reference]; ok {
			_AttributeGroups = append(_AttributeGroups, _instance)
		}
	}
	reference.AttributeGroups = _AttributeGroups
	var _Attributes []*Attribute
	for _, _reference := range reference.Attributes {
		if _instance, ok := stage.Attributes_instance[_reference]; ok {
			_Attributes = append(_Attributes, _instance)
		}
	}
	reference.Attributes = _Attributes
}

func (reference *Choice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
}

func (reference *ComplexContent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ComplexType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.OuterElement; _reference != nil {
		reference.OuterElement = nil
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			reference.OuterElement = _instance
		}
	}
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	if _reference := reference.Extension; _reference != nil {
		reference.Extension = nil
		if _instance, ok := stage.Extensions_instance[_reference]; ok {
			reference.Extension = _instance
		}
	}
	if _reference := reference.SimpleContent; _reference != nil {
		reference.SimpleContent = nil
		if _instance, ok := stage.SimpleContents_instance[_reference]; ok {
			reference.SimpleContent = _instance
		}
	}
	if _reference := reference.ComplexContent; _reference != nil {
		reference.ComplexContent = nil
		if _instance, ok := stage.ComplexContents_instance[_reference]; ok {
			reference.ComplexContent = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
	var _Attributes []*Attribute
	for _, _reference := range reference.Attributes {
		if _instance, ok := stage.Attributes_instance[_reference]; ok {
			_Attributes = append(_Attributes, _instance)
		}
	}
	reference.Attributes = _Attributes
	var _AttributeGroups []*AttributeGroup
	for _, _reference := range reference.AttributeGroups {
		if _instance, ok := stage.AttributeGroups_instance[_reference]; ok {
			_AttributeGroups = append(_AttributeGroups, _instance)
		}
	}
	reference.AttributeGroups = _AttributeGroups
}

func (reference *Documentation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Element) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	if _reference := reference.SimpleType; _reference != nil {
		reference.SimpleType = nil
		if _instance, ok := stage.SimpleTypes_instance[_reference]; ok {
			reference.SimpleType = _instance
		}
	}
	if _reference := reference.ComplexType; _reference != nil {
		reference.ComplexType = nil
		if _instance, ok := stage.ComplexTypes_instance[_reference]; ok {
			reference.ComplexType = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
}

func (reference *Enumeration) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Extension) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
	var _Attributes []*Attribute
	for _, _reference := range reference.Attributes {
		if _instance, ok := stage.Attributes_instance[_reference]; ok {
			_Attributes = append(_Attributes, _instance)
		}
	}
	reference.Attributes = _Attributes
	var _AttributeGroups []*AttributeGroup
	for _, _reference := range reference.AttributeGroups {
		if _instance, ok := stage.AttributeGroups_instance[_reference]; ok {
			_AttributeGroups = append(_AttributeGroups, _instance)
		}
	}
	reference.AttributeGroups = _AttributeGroups
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	if _reference := reference.OuterElement; _reference != nil {
		reference.OuterElement = nil
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			reference.OuterElement = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
}

func (reference *Length) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MaxInclusive) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MaxLength) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MinInclusive) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MinLength) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Pattern) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Restriction) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	if _reference := reference.MinInclusive; _reference != nil {
		reference.MinInclusive = nil
		if _instance, ok := stage.MinInclusives_instance[_reference]; ok {
			reference.MinInclusive = _instance
		}
	}
	if _reference := reference.MaxInclusive; _reference != nil {
		reference.MaxInclusive = nil
		if _instance, ok := stage.MaxInclusives_instance[_reference]; ok {
			reference.MaxInclusive = _instance
		}
	}
	if _reference := reference.Pattern; _reference != nil {
		reference.Pattern = nil
		if _instance, ok := stage.Patterns_instance[_reference]; ok {
			reference.Pattern = _instance
		}
	}
	if _reference := reference.WhiteSpace; _reference != nil {
		reference.WhiteSpace = nil
		if _instance, ok := stage.WhiteSpaces_instance[_reference]; ok {
			reference.WhiteSpace = _instance
		}
	}
	if _reference := reference.MinLength; _reference != nil {
		reference.MinLength = nil
		if _instance, ok := stage.MinLengths_instance[_reference]; ok {
			reference.MinLength = _instance
		}
	}
	if _reference := reference.MaxLength; _reference != nil {
		reference.MaxLength = nil
		if _instance, ok := stage.MaxLengths_instance[_reference]; ok {
			reference.MaxLength = _instance
		}
	}
	if _reference := reference.Length; _reference != nil {
		reference.Length = nil
		if _instance, ok := stage.Lengths_instance[_reference]; ok {
			reference.Length = _instance
		}
	}
	if _reference := reference.TotalDigit; _reference != nil {
		reference.TotalDigit = nil
		if _instance, ok := stage.TotalDigits_instance[_reference]; ok {
			reference.TotalDigit = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Enumerations []*Enumeration
	for _, _reference := range reference.Enumerations {
		if _instance, ok := stage.Enumerations_instance[_reference]; ok {
			_Enumerations = append(_Enumerations, _instance)
		}
	}
	reference.Enumerations = _Enumerations
}

func (reference *Schema) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
	var _SimpleTypes []*SimpleType
	for _, _reference := range reference.SimpleTypes {
		if _instance, ok := stage.SimpleTypes_instance[_reference]; ok {
			_SimpleTypes = append(_SimpleTypes, _instance)
		}
	}
	reference.SimpleTypes = _SimpleTypes
	var _ComplexTypes []*ComplexType
	for _, _reference := range reference.ComplexTypes {
		if _instance, ok := stage.ComplexTypes_instance[_reference]; ok {
			_ComplexTypes = append(_ComplexTypes, _instance)
		}
	}
	reference.ComplexTypes = _ComplexTypes
	var _AttributeGroups []*AttributeGroup
	for _, _reference := range reference.AttributeGroups {
		if _instance, ok := stage.AttributeGroups_instance[_reference]; ok {
			_AttributeGroups = append(_AttributeGroups, _instance)
		}
	}
	reference.AttributeGroups = _AttributeGroups
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
}

func (reference *Sequence) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Sequences []*Sequence
	for _, _reference := range reference.Sequences {
		if _instance, ok := stage.Sequences_instance[_reference]; ok {
			_Sequences = append(_Sequences, _instance)
		}
	}
	reference.Sequences = _Sequences
	var _Alls []*All
	for _, _reference := range reference.Alls {
		if _instance, ok := stage.Alls_instance[_reference]; ok {
			_Alls = append(_Alls, _instance)
		}
	}
	reference.Alls = _Alls
	var _Choices []*Choice
	for _, _reference := range reference.Choices {
		if _instance, ok := stage.Choices_instance[_reference]; ok {
			_Choices = append(_Choices, _instance)
		}
	}
	reference.Choices = _Choices
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _Elements []*Element
	for _, _reference := range reference.Elements {
		if _instance, ok := stage.Elements_instance[_reference]; ok {
			_Elements = append(_Elements, _instance)
		}
	}
	reference.Elements = _Elements
}

func (reference *SimpleContent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Extension; _reference != nil {
		reference.Extension = nil
		if _instance, ok := stage.Extensions_instance[_reference]; ok {
			reference.Extension = _instance
		}
	}
	if _reference := reference.Restriction; _reference != nil {
		reference.Restriction = nil
		if _instance, ok := stage.Restrictions_instance[_reference]; ok {
			reference.Restriction = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SimpleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	if _reference := reference.Restriction; _reference != nil {
		reference.Restriction = nil
		if _instance, ok := stage.Restrictions_instance[_reference]; ok {
			reference.Restriction = _instance
		}
	}
	if _reference := reference.Union; _reference != nil {
		reference.Union = nil
		if _instance, ok := stage.Unions_instance[_reference]; ok {
			reference.Union = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TotalDigit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Union) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *WhiteSpace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Annotation; _reference != nil {
		reference.Annotation = nil
		if _instance, ok := stage.Annotations_instance[_reference]; ok {
			reference.Annotation = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (all *All) GongDiff(stage *Stage, allOther *All) (diffs []string) {
	// insertion point for field diffs
	if all.Name != allOther.Name {
		diffs = append(diffs, all.GongMarshallField(stage, "Name"))
	}
	if (all.Annotation == nil) != (allOther.Annotation == nil) {
		diffs = append(diffs, all.GongMarshallField(stage, "Annotation"))
	} else if all.Annotation != nil && allOther.Annotation != nil {
		if all.Annotation != allOther.Annotation {
			diffs = append(diffs, all.GongMarshallField(stage, "Annotation"))
		}
	}
	if all.OuterElementName != allOther.OuterElementName {
		diffs = append(diffs, all.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(all.Sequences) != len(allOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range all.Sequences {
			if (all.Sequences[i] == nil) != (allOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if all.Sequences[i] != nil && allOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if all.Sequences[i] != allOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, all, allOther, "Sequences", allOther.Sequences, all.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(all.Alls) != len(allOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range all.Alls {
			if (all.Alls[i] == nil) != (allOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if all.Alls[i] != nil && allOther.Alls[i] != nil {
				// this is a pointer comparaison
				if all.Alls[i] != allOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, all, allOther, "Alls", allOther.Alls, all.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(all.Choices) != len(allOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range all.Choices {
			if (all.Choices[i] == nil) != (allOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if all.Choices[i] != nil && allOther.Choices[i] != nil {
				// this is a pointer comparaison
				if all.Choices[i] != allOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, all, allOther, "Choices", allOther.Choices, all.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(all.Groups) != len(allOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range all.Groups {
			if (all.Groups[i] == nil) != (allOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if all.Groups[i] != nil && allOther.Groups[i] != nil {
				// this is a pointer comparaison
				if all.Groups[i] != allOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, all, allOther, "Groups", allOther.Groups, all.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(all.Elements) != len(allOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range all.Elements {
			if (all.Elements[i] == nil) != (allOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if all.Elements[i] != nil && allOther.Elements[i] != nil {
				// this is a pointer comparaison
				if all.Elements[i] != allOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, all, allOther, "Elements", allOther.Elements, all.Elements)
		diffs = append(diffs, ops)
	}
	if all.Order != allOther.Order {
		diffs = append(diffs, all.GongMarshallField(stage, "Order"))
	}
	if all.Depth != allOther.Depth {
		diffs = append(diffs, all.GongMarshallField(stage, "Depth"))
	}
	if all.MinOccurs != allOther.MinOccurs {
		diffs = append(diffs, all.GongMarshallField(stage, "MinOccurs"))
	}
	if all.MaxOccurs != allOther.MaxOccurs {
		diffs = append(diffs, all.GongMarshallField(stage, "MaxOccurs"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (annotation *Annotation) GongDiff(stage *Stage, annotationOther *Annotation) (diffs []string) {
	// insertion point for field diffs
	if annotation.Name != annotationOther.Name {
		diffs = append(diffs, annotation.GongMarshallField(stage, "Name"))
	}
	DocumentationsDifferent := false
	if len(annotation.Documentations) != len(annotationOther.Documentations) {
		DocumentationsDifferent = true
	} else {
		for i := range annotation.Documentations {
			if (annotation.Documentations[i] == nil) != (annotationOther.Documentations[i] == nil) {
				DocumentationsDifferent = true
				break
			} else if annotation.Documentations[i] != nil && annotationOther.Documentations[i] != nil {
				// this is a pointer comparaison
				if annotation.Documentations[i] != annotationOther.Documentations[i] {
					DocumentationsDifferent = true
					break
				}
			}
		}
	}
	if DocumentationsDifferent {
		ops := Diff(stage, annotation, annotationOther, "Documentations", annotationOther.Documentations, annotation.Documentations)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute *Attribute) GongDiff(stage *Stage, attributeOther *Attribute) (diffs []string) {
	// insertion point for field diffs
	if attribute.Name != attributeOther.Name {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Name"))
	}
	if attribute.NameXSD != attributeOther.NameXSD {
		diffs = append(diffs, attribute.GongMarshallField(stage, "NameXSD"))
	}
	if attribute.Type != attributeOther.Type {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Type"))
	}
	if (attribute.Annotation == nil) != (attributeOther.Annotation == nil) {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Annotation"))
	} else if attribute.Annotation != nil && attributeOther.Annotation != nil {
		if attribute.Annotation != attributeOther.Annotation {
			diffs = append(diffs, attribute.GongMarshallField(stage, "Annotation"))
		}
	}
	if attribute.HasNameConflict != attributeOther.HasNameConflict {
		diffs = append(diffs, attribute.GongMarshallField(stage, "HasNameConflict"))
	}
	if attribute.GoIdentifier != attributeOther.GoIdentifier {
		diffs = append(diffs, attribute.GongMarshallField(stage, "GoIdentifier"))
	}
	if attribute.Default != attributeOther.Default {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Default"))
	}
	if attribute.Use != attributeOther.Use {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Use"))
	}
	if attribute.Form != attributeOther.Form {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Form"))
	}
	if attribute.Fixed != attributeOther.Fixed {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Fixed"))
	}
	if attribute.Ref != attributeOther.Ref {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Ref"))
	}
	if attribute.TargetNamespace != attributeOther.TargetNamespace {
		diffs = append(diffs, attribute.GongMarshallField(stage, "TargetNamespace"))
	}
	if attribute.SimpleType != attributeOther.SimpleType {
		diffs = append(diffs, attribute.GongMarshallField(stage, "SimpleType"))
	}
	if attribute.IDXSD != attributeOther.IDXSD {
		diffs = append(diffs, attribute.GongMarshallField(stage, "IDXSD"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attributegroup *AttributeGroup) GongDiff(stage *Stage, attributegroupOther *AttributeGroup) (diffs []string) {
	// insertion point for field diffs
	if attributegroup.Name != attributegroupOther.Name {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Name"))
	}
	if attributegroup.NameXSD != attributegroupOther.NameXSD {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "NameXSD"))
	}
	if (attributegroup.Annotation == nil) != (attributegroupOther.Annotation == nil) {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Annotation"))
	} else if attributegroup.Annotation != nil && attributegroupOther.Annotation != nil {
		if attributegroup.Annotation != attributegroupOther.Annotation {
			diffs = append(diffs, attributegroup.GongMarshallField(stage, "Annotation"))
		}
	}
	if attributegroup.HasNameConflict != attributegroupOther.HasNameConflict {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "HasNameConflict"))
	}
	if attributegroup.GoIdentifier != attributegroupOther.GoIdentifier {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "GoIdentifier"))
	}
	AttributeGroupsDifferent := false
	if len(attributegroup.AttributeGroups) != len(attributegroupOther.AttributeGroups) {
		AttributeGroupsDifferent = true
	} else {
		for i := range attributegroup.AttributeGroups {
			if (attributegroup.AttributeGroups[i] == nil) != (attributegroupOther.AttributeGroups[i] == nil) {
				AttributeGroupsDifferent = true
				break
			} else if attributegroup.AttributeGroups[i] != nil && attributegroupOther.AttributeGroups[i] != nil {
				// this is a pointer comparaison
				if attributegroup.AttributeGroups[i] != attributegroupOther.AttributeGroups[i] {
					AttributeGroupsDifferent = true
					break
				}
			}
		}
	}
	if AttributeGroupsDifferent {
		ops := Diff(stage, attributegroup, attributegroupOther, "AttributeGroups", attributegroupOther.AttributeGroups, attributegroup.AttributeGroups)
		diffs = append(diffs, ops)
	}
	if attributegroup.Ref != attributegroupOther.Ref {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Ref"))
	}
	AttributesDifferent := false
	if len(attributegroup.Attributes) != len(attributegroupOther.Attributes) {
		AttributesDifferent = true
	} else {
		for i := range attributegroup.Attributes {
			if (attributegroup.Attributes[i] == nil) != (attributegroupOther.Attributes[i] == nil) {
				AttributesDifferent = true
				break
			} else if attributegroup.Attributes[i] != nil && attributegroupOther.Attributes[i] != nil {
				// this is a pointer comparaison
				if attributegroup.Attributes[i] != attributegroupOther.Attributes[i] {
					AttributesDifferent = true
					break
				}
			}
		}
	}
	if AttributesDifferent {
		ops := Diff(stage, attributegroup, attributegroupOther, "Attributes", attributegroupOther.Attributes, attributegroup.Attributes)
		diffs = append(diffs, ops)
	}
	if attributegroup.Order != attributegroupOther.Order {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Order"))
	}
	if attributegroup.Depth != attributegroupOther.Depth {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Depth"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (choice *Choice) GongDiff(stage *Stage, choiceOther *Choice) (diffs []string) {
	// insertion point for field diffs
	if choice.Name != choiceOther.Name {
		diffs = append(diffs, choice.GongMarshallField(stage, "Name"))
	}
	if (choice.Annotation == nil) != (choiceOther.Annotation == nil) {
		diffs = append(diffs, choice.GongMarshallField(stage, "Annotation"))
	} else if choice.Annotation != nil && choiceOther.Annotation != nil {
		if choice.Annotation != choiceOther.Annotation {
			diffs = append(diffs, choice.GongMarshallField(stage, "Annotation"))
		}
	}
	if choice.OuterElementName != choiceOther.OuterElementName {
		diffs = append(diffs, choice.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(choice.Sequences) != len(choiceOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range choice.Sequences {
			if (choice.Sequences[i] == nil) != (choiceOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if choice.Sequences[i] != nil && choiceOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if choice.Sequences[i] != choiceOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, choice, choiceOther, "Sequences", choiceOther.Sequences, choice.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(choice.Alls) != len(choiceOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range choice.Alls {
			if (choice.Alls[i] == nil) != (choiceOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if choice.Alls[i] != nil && choiceOther.Alls[i] != nil {
				// this is a pointer comparaison
				if choice.Alls[i] != choiceOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, choice, choiceOther, "Alls", choiceOther.Alls, choice.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(choice.Choices) != len(choiceOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range choice.Choices {
			if (choice.Choices[i] == nil) != (choiceOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if choice.Choices[i] != nil && choiceOther.Choices[i] != nil {
				// this is a pointer comparaison
				if choice.Choices[i] != choiceOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, choice, choiceOther, "Choices", choiceOther.Choices, choice.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(choice.Groups) != len(choiceOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range choice.Groups {
			if (choice.Groups[i] == nil) != (choiceOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if choice.Groups[i] != nil && choiceOther.Groups[i] != nil {
				// this is a pointer comparaison
				if choice.Groups[i] != choiceOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, choice, choiceOther, "Groups", choiceOther.Groups, choice.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(choice.Elements) != len(choiceOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range choice.Elements {
			if (choice.Elements[i] == nil) != (choiceOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if choice.Elements[i] != nil && choiceOther.Elements[i] != nil {
				// this is a pointer comparaison
				if choice.Elements[i] != choiceOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, choice, choiceOther, "Elements", choiceOther.Elements, choice.Elements)
		diffs = append(diffs, ops)
	}
	if choice.Order != choiceOther.Order {
		diffs = append(diffs, choice.GongMarshallField(stage, "Order"))
	}
	if choice.Depth != choiceOther.Depth {
		diffs = append(diffs, choice.GongMarshallField(stage, "Depth"))
	}
	if choice.MinOccurs != choiceOther.MinOccurs {
		diffs = append(diffs, choice.GongMarshallField(stage, "MinOccurs"))
	}
	if choice.MaxOccurs != choiceOther.MaxOccurs {
		diffs = append(diffs, choice.GongMarshallField(stage, "MaxOccurs"))
	}
	if choice.IsDuplicatedInXSD != choiceOther.IsDuplicatedInXSD {
		diffs = append(diffs, choice.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (complexcontent *ComplexContent) GongDiff(stage *Stage, complexcontentOther *ComplexContent) (diffs []string) {
	// insertion point for field diffs
	if complexcontent.Name != complexcontentOther.Name {
		diffs = append(diffs, complexcontent.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (complextype *ComplexType) GongDiff(stage *Stage, complextypeOther *ComplexType) (diffs []string) {
	// insertion point for field diffs
	if complextype.Name != complextypeOther.Name {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Name"))
	}
	if complextype.HasNameConflict != complextypeOther.HasNameConflict {
		diffs = append(diffs, complextype.GongMarshallField(stage, "HasNameConflict"))
	}
	if complextype.GoIdentifier != complextypeOther.GoIdentifier {
		diffs = append(diffs, complextype.GongMarshallField(stage, "GoIdentifier"))
	}
	if complextype.IsAnonymous != complextypeOther.IsAnonymous {
		diffs = append(diffs, complextype.GongMarshallField(stage, "IsAnonymous"))
	}
	if (complextype.OuterElement == nil) != (complextypeOther.OuterElement == nil) {
		diffs = append(diffs, complextype.GongMarshallField(stage, "OuterElement"))
	} else if complextype.OuterElement != nil && complextypeOther.OuterElement != nil {
		if complextype.OuterElement != complextypeOther.OuterElement {
			diffs = append(diffs, complextype.GongMarshallField(stage, "OuterElement"))
		}
	}
	if (complextype.Annotation == nil) != (complextypeOther.Annotation == nil) {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Annotation"))
	} else if complextype.Annotation != nil && complextypeOther.Annotation != nil {
		if complextype.Annotation != complextypeOther.Annotation {
			diffs = append(diffs, complextype.GongMarshallField(stage, "Annotation"))
		}
	}
	if complextype.NameXSD != complextypeOther.NameXSD {
		diffs = append(diffs, complextype.GongMarshallField(stage, "NameXSD"))
	}
	if complextype.OuterElementName != complextypeOther.OuterElementName {
		diffs = append(diffs, complextype.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(complextype.Sequences) != len(complextypeOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range complextype.Sequences {
			if (complextype.Sequences[i] == nil) != (complextypeOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if complextype.Sequences[i] != nil && complextypeOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if complextype.Sequences[i] != complextypeOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Sequences", complextypeOther.Sequences, complextype.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(complextype.Alls) != len(complextypeOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range complextype.Alls {
			if (complextype.Alls[i] == nil) != (complextypeOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if complextype.Alls[i] != nil && complextypeOther.Alls[i] != nil {
				// this is a pointer comparaison
				if complextype.Alls[i] != complextypeOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Alls", complextypeOther.Alls, complextype.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(complextype.Choices) != len(complextypeOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range complextype.Choices {
			if (complextype.Choices[i] == nil) != (complextypeOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if complextype.Choices[i] != nil && complextypeOther.Choices[i] != nil {
				// this is a pointer comparaison
				if complextype.Choices[i] != complextypeOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Choices", complextypeOther.Choices, complextype.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(complextype.Groups) != len(complextypeOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range complextype.Groups {
			if (complextype.Groups[i] == nil) != (complextypeOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if complextype.Groups[i] != nil && complextypeOther.Groups[i] != nil {
				// this is a pointer comparaison
				if complextype.Groups[i] != complextypeOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Groups", complextypeOther.Groups, complextype.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(complextype.Elements) != len(complextypeOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range complextype.Elements {
			if (complextype.Elements[i] == nil) != (complextypeOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if complextype.Elements[i] != nil && complextypeOther.Elements[i] != nil {
				// this is a pointer comparaison
				if complextype.Elements[i] != complextypeOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Elements", complextypeOther.Elements, complextype.Elements)
		diffs = append(diffs, ops)
	}
	if complextype.Order != complextypeOther.Order {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Order"))
	}
	if complextype.Depth != complextypeOther.Depth {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Depth"))
	}
	if complextype.MinOccurs != complextypeOther.MinOccurs {
		diffs = append(diffs, complextype.GongMarshallField(stage, "MinOccurs"))
	}
	if complextype.MaxOccurs != complextypeOther.MaxOccurs {
		diffs = append(diffs, complextype.GongMarshallField(stage, "MaxOccurs"))
	}
	if (complextype.Extension == nil) != (complextypeOther.Extension == nil) {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Extension"))
	} else if complextype.Extension != nil && complextypeOther.Extension != nil {
		if complextype.Extension != complextypeOther.Extension {
			diffs = append(diffs, complextype.GongMarshallField(stage, "Extension"))
		}
	}
	if (complextype.SimpleContent == nil) != (complextypeOther.SimpleContent == nil) {
		diffs = append(diffs, complextype.GongMarshallField(stage, "SimpleContent"))
	} else if complextype.SimpleContent != nil && complextypeOther.SimpleContent != nil {
		if complextype.SimpleContent != complextypeOther.SimpleContent {
			diffs = append(diffs, complextype.GongMarshallField(stage, "SimpleContent"))
		}
	}
	if (complextype.ComplexContent == nil) != (complextypeOther.ComplexContent == nil) {
		diffs = append(diffs, complextype.GongMarshallField(stage, "ComplexContent"))
	} else if complextype.ComplexContent != nil && complextypeOther.ComplexContent != nil {
		if complextype.ComplexContent != complextypeOther.ComplexContent {
			diffs = append(diffs, complextype.GongMarshallField(stage, "ComplexContent"))
		}
	}
	AttributesDifferent := false
	if len(complextype.Attributes) != len(complextypeOther.Attributes) {
		AttributesDifferent = true
	} else {
		for i := range complextype.Attributes {
			if (complextype.Attributes[i] == nil) != (complextypeOther.Attributes[i] == nil) {
				AttributesDifferent = true
				break
			} else if complextype.Attributes[i] != nil && complextypeOther.Attributes[i] != nil {
				// this is a pointer comparaison
				if complextype.Attributes[i] != complextypeOther.Attributes[i] {
					AttributesDifferent = true
					break
				}
			}
		}
	}
	if AttributesDifferent {
		ops := Diff(stage, complextype, complextypeOther, "Attributes", complextypeOther.Attributes, complextype.Attributes)
		diffs = append(diffs, ops)
	}
	AttributeGroupsDifferent := false
	if len(complextype.AttributeGroups) != len(complextypeOther.AttributeGroups) {
		AttributeGroupsDifferent = true
	} else {
		for i := range complextype.AttributeGroups {
			if (complextype.AttributeGroups[i] == nil) != (complextypeOther.AttributeGroups[i] == nil) {
				AttributeGroupsDifferent = true
				break
			} else if complextype.AttributeGroups[i] != nil && complextypeOther.AttributeGroups[i] != nil {
				// this is a pointer comparaison
				if complextype.AttributeGroups[i] != complextypeOther.AttributeGroups[i] {
					AttributeGroupsDifferent = true
					break
				}
			}
		}
	}
	if AttributeGroupsDifferent {
		ops := Diff(stage, complextype, complextypeOther, "AttributeGroups", complextypeOther.AttributeGroups, complextype.AttributeGroups)
		diffs = append(diffs, ops)
	}
	if complextype.IsDuplicatedInXSD != complextypeOther.IsDuplicatedInXSD {
		diffs = append(diffs, complextype.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (documentation *Documentation) GongDiff(stage *Stage, documentationOther *Documentation) (diffs []string) {
	// insertion point for field diffs
	if documentation.Name != documentationOther.Name {
		diffs = append(diffs, documentation.GongMarshallField(stage, "Name"))
	}
	if documentation.Text != documentationOther.Text {
		diffs = append(diffs, documentation.GongMarshallField(stage, "Text"))
	}
	if documentation.Source != documentationOther.Source {
		diffs = append(diffs, documentation.GongMarshallField(stage, "Source"))
	}
	if documentation.Lang != documentationOther.Lang {
		diffs = append(diffs, documentation.GongMarshallField(stage, "Lang"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (element *Element) GongDiff(stage *Stage, elementOther *Element) (diffs []string) {
	// insertion point for field diffs
	if element.Name != elementOther.Name {
		diffs = append(diffs, element.GongMarshallField(stage, "Name"))
	}
	if element.Order != elementOther.Order {
		diffs = append(diffs, element.GongMarshallField(stage, "Order"))
	}
	if element.Depth != elementOther.Depth {
		diffs = append(diffs, element.GongMarshallField(stage, "Depth"))
	}
	if element.HasNameConflict != elementOther.HasNameConflict {
		diffs = append(diffs, element.GongMarshallField(stage, "HasNameConflict"))
	}
	if element.GoIdentifier != elementOther.GoIdentifier {
		diffs = append(diffs, element.GongMarshallField(stage, "GoIdentifier"))
	}
	if (element.Annotation == nil) != (elementOther.Annotation == nil) {
		diffs = append(diffs, element.GongMarshallField(stage, "Annotation"))
	} else if element.Annotation != nil && elementOther.Annotation != nil {
		if element.Annotation != elementOther.Annotation {
			diffs = append(diffs, element.GongMarshallField(stage, "Annotation"))
		}
	}
	if element.NameXSD != elementOther.NameXSD {
		diffs = append(diffs, element.GongMarshallField(stage, "NameXSD"))
	}
	if element.Type != elementOther.Type {
		diffs = append(diffs, element.GongMarshallField(stage, "Type"))
	}
	if element.MinOccurs != elementOther.MinOccurs {
		diffs = append(diffs, element.GongMarshallField(stage, "MinOccurs"))
	}
	if element.MaxOccurs != elementOther.MaxOccurs {
		diffs = append(diffs, element.GongMarshallField(stage, "MaxOccurs"))
	}
	if element.Default != elementOther.Default {
		diffs = append(diffs, element.GongMarshallField(stage, "Default"))
	}
	if element.Fixed != elementOther.Fixed {
		diffs = append(diffs, element.GongMarshallField(stage, "Fixed"))
	}
	if element.Nillable != elementOther.Nillable {
		diffs = append(diffs, element.GongMarshallField(stage, "Nillable"))
	}
	if element.Ref != elementOther.Ref {
		diffs = append(diffs, element.GongMarshallField(stage, "Ref"))
	}
	if element.Abstract != elementOther.Abstract {
		diffs = append(diffs, element.GongMarshallField(stage, "Abstract"))
	}
	if element.Form != elementOther.Form {
		diffs = append(diffs, element.GongMarshallField(stage, "Form"))
	}
	if element.Block != elementOther.Block {
		diffs = append(diffs, element.GongMarshallField(stage, "Block"))
	}
	if element.Final != elementOther.Final {
		diffs = append(diffs, element.GongMarshallField(stage, "Final"))
	}
	if (element.SimpleType == nil) != (elementOther.SimpleType == nil) {
		diffs = append(diffs, element.GongMarshallField(stage, "SimpleType"))
	} else if element.SimpleType != nil && elementOther.SimpleType != nil {
		if element.SimpleType != elementOther.SimpleType {
			diffs = append(diffs, element.GongMarshallField(stage, "SimpleType"))
		}
	}
	if (element.ComplexType == nil) != (elementOther.ComplexType == nil) {
		diffs = append(diffs, element.GongMarshallField(stage, "ComplexType"))
	} else if element.ComplexType != nil && elementOther.ComplexType != nil {
		if element.ComplexType != elementOther.ComplexType {
			diffs = append(diffs, element.GongMarshallField(stage, "ComplexType"))
		}
	}
	GroupsDifferent := false
	if len(element.Groups) != len(elementOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range element.Groups {
			if (element.Groups[i] == nil) != (elementOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if element.Groups[i] != nil && elementOther.Groups[i] != nil {
				// this is a pointer comparaison
				if element.Groups[i] != elementOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, element, elementOther, "Groups", elementOther.Groups, element.Groups)
		diffs = append(diffs, ops)
	}
	if element.IsDuplicatedInXSD != elementOther.IsDuplicatedInXSD {
		diffs = append(diffs, element.GongMarshallField(stage, "IsDuplicatedInXSD"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (enumeration *Enumeration) GongDiff(stage *Stage, enumerationOther *Enumeration) (diffs []string) {
	// insertion point for field diffs
	if enumeration.Name != enumerationOther.Name {
		diffs = append(diffs, enumeration.GongMarshallField(stage, "Name"))
	}
	if (enumeration.Annotation == nil) != (enumerationOther.Annotation == nil) {
		diffs = append(diffs, enumeration.GongMarshallField(stage, "Annotation"))
	} else if enumeration.Annotation != nil && enumerationOther.Annotation != nil {
		if enumeration.Annotation != enumerationOther.Annotation {
			diffs = append(diffs, enumeration.GongMarshallField(stage, "Annotation"))
		}
	}
	if enumeration.Value != enumerationOther.Value {
		diffs = append(diffs, enumeration.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (extension *Extension) GongDiff(stage *Stage, extensionOther *Extension) (diffs []string) {
	// insertion point for field diffs
	if extension.Name != extensionOther.Name {
		diffs = append(diffs, extension.GongMarshallField(stage, "Name"))
	}
	if extension.OuterElementName != extensionOther.OuterElementName {
		diffs = append(diffs, extension.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(extension.Sequences) != len(extensionOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range extension.Sequences {
			if (extension.Sequences[i] == nil) != (extensionOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if extension.Sequences[i] != nil && extensionOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if extension.Sequences[i] != extensionOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, extension, extensionOther, "Sequences", extensionOther.Sequences, extension.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(extension.Alls) != len(extensionOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range extension.Alls {
			if (extension.Alls[i] == nil) != (extensionOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if extension.Alls[i] != nil && extensionOther.Alls[i] != nil {
				// this is a pointer comparaison
				if extension.Alls[i] != extensionOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, extension, extensionOther, "Alls", extensionOther.Alls, extension.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(extension.Choices) != len(extensionOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range extension.Choices {
			if (extension.Choices[i] == nil) != (extensionOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if extension.Choices[i] != nil && extensionOther.Choices[i] != nil {
				// this is a pointer comparaison
				if extension.Choices[i] != extensionOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, extension, extensionOther, "Choices", extensionOther.Choices, extension.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(extension.Groups) != len(extensionOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range extension.Groups {
			if (extension.Groups[i] == nil) != (extensionOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if extension.Groups[i] != nil && extensionOther.Groups[i] != nil {
				// this is a pointer comparaison
				if extension.Groups[i] != extensionOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, extension, extensionOther, "Groups", extensionOther.Groups, extension.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(extension.Elements) != len(extensionOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range extension.Elements {
			if (extension.Elements[i] == nil) != (extensionOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if extension.Elements[i] != nil && extensionOther.Elements[i] != nil {
				// this is a pointer comparaison
				if extension.Elements[i] != extensionOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, extension, extensionOther, "Elements", extensionOther.Elements, extension.Elements)
		diffs = append(diffs, ops)
	}
	if extension.Order != extensionOther.Order {
		diffs = append(diffs, extension.GongMarshallField(stage, "Order"))
	}
	if extension.Depth != extensionOther.Depth {
		diffs = append(diffs, extension.GongMarshallField(stage, "Depth"))
	}
	if extension.MinOccurs != extensionOther.MinOccurs {
		diffs = append(diffs, extension.GongMarshallField(stage, "MinOccurs"))
	}
	if extension.MaxOccurs != extensionOther.MaxOccurs {
		diffs = append(diffs, extension.GongMarshallField(stage, "MaxOccurs"))
	}
	if extension.Base != extensionOther.Base {
		diffs = append(diffs, extension.GongMarshallField(stage, "Base"))
	}
	if extension.Ref != extensionOther.Ref {
		diffs = append(diffs, extension.GongMarshallField(stage, "Ref"))
	}
	AttributesDifferent := false
	if len(extension.Attributes) != len(extensionOther.Attributes) {
		AttributesDifferent = true
	} else {
		for i := range extension.Attributes {
			if (extension.Attributes[i] == nil) != (extensionOther.Attributes[i] == nil) {
				AttributesDifferent = true
				break
			} else if extension.Attributes[i] != nil && extensionOther.Attributes[i] != nil {
				// this is a pointer comparaison
				if extension.Attributes[i] != extensionOther.Attributes[i] {
					AttributesDifferent = true
					break
				}
			}
		}
	}
	if AttributesDifferent {
		ops := Diff(stage, extension, extensionOther, "Attributes", extensionOther.Attributes, extension.Attributes)
		diffs = append(diffs, ops)
	}
	AttributeGroupsDifferent := false
	if len(extension.AttributeGroups) != len(extensionOther.AttributeGroups) {
		AttributeGroupsDifferent = true
	} else {
		for i := range extension.AttributeGroups {
			if (extension.AttributeGroups[i] == nil) != (extensionOther.AttributeGroups[i] == nil) {
				AttributeGroupsDifferent = true
				break
			} else if extension.AttributeGroups[i] != nil && extensionOther.AttributeGroups[i] != nil {
				// this is a pointer comparaison
				if extension.AttributeGroups[i] != extensionOther.AttributeGroups[i] {
					AttributeGroupsDifferent = true
					break
				}
			}
		}
	}
	if AttributeGroupsDifferent {
		ops := Diff(stage, extension, extensionOther, "AttributeGroups", extensionOther.AttributeGroups, extension.AttributeGroups)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(stage *Stage, groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, group.GongMarshallField(stage, "Name"))
	}
	if (group.Annotation == nil) != (groupOther.Annotation == nil) {
		diffs = append(diffs, group.GongMarshallField(stage, "Annotation"))
	} else if group.Annotation != nil && groupOther.Annotation != nil {
		if group.Annotation != groupOther.Annotation {
			diffs = append(diffs, group.GongMarshallField(stage, "Annotation"))
		}
	}
	if group.NameXSD != groupOther.NameXSD {
		diffs = append(diffs, group.GongMarshallField(stage, "NameXSD"))
	}
	if group.Ref != groupOther.Ref {
		diffs = append(diffs, group.GongMarshallField(stage, "Ref"))
	}
	if group.IsAnonymous != groupOther.IsAnonymous {
		diffs = append(diffs, group.GongMarshallField(stage, "IsAnonymous"))
	}
	if (group.OuterElement == nil) != (groupOther.OuterElement == nil) {
		diffs = append(diffs, group.GongMarshallField(stage, "OuterElement"))
	} else if group.OuterElement != nil && groupOther.OuterElement != nil {
		if group.OuterElement != groupOther.OuterElement {
			diffs = append(diffs, group.GongMarshallField(stage, "OuterElement"))
		}
	}
	if group.HasNameConflict != groupOther.HasNameConflict {
		diffs = append(diffs, group.GongMarshallField(stage, "HasNameConflict"))
	}
	if group.GoIdentifier != groupOther.GoIdentifier {
		diffs = append(diffs, group.GongMarshallField(stage, "GoIdentifier"))
	}
	if group.OuterElementName != groupOther.OuterElementName {
		diffs = append(diffs, group.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(group.Sequences) != len(groupOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range group.Sequences {
			if (group.Sequences[i] == nil) != (groupOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if group.Sequences[i] != nil && groupOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if group.Sequences[i] != groupOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, group, groupOther, "Sequences", groupOther.Sequences, group.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(group.Alls) != len(groupOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range group.Alls {
			if (group.Alls[i] == nil) != (groupOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if group.Alls[i] != nil && groupOther.Alls[i] != nil {
				// this is a pointer comparaison
				if group.Alls[i] != groupOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, group, groupOther, "Alls", groupOther.Alls, group.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(group.Choices) != len(groupOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range group.Choices {
			if (group.Choices[i] == nil) != (groupOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if group.Choices[i] != nil && groupOther.Choices[i] != nil {
				// this is a pointer comparaison
				if group.Choices[i] != groupOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, group, groupOther, "Choices", groupOther.Choices, group.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(group.Groups) != len(groupOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range group.Groups {
			if (group.Groups[i] == nil) != (groupOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if group.Groups[i] != nil && groupOther.Groups[i] != nil {
				// this is a pointer comparaison
				if group.Groups[i] != groupOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, group, groupOther, "Groups", groupOther.Groups, group.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(group.Elements) != len(groupOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range group.Elements {
			if (group.Elements[i] == nil) != (groupOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if group.Elements[i] != nil && groupOther.Elements[i] != nil {
				// this is a pointer comparaison
				if group.Elements[i] != groupOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, group, groupOther, "Elements", groupOther.Elements, group.Elements)
		diffs = append(diffs, ops)
	}
	if group.Order != groupOther.Order {
		diffs = append(diffs, group.GongMarshallField(stage, "Order"))
	}
	if group.Depth != groupOther.Depth {
		diffs = append(diffs, group.GongMarshallField(stage, "Depth"))
	}
	if group.MinOccurs != groupOther.MinOccurs {
		diffs = append(diffs, group.GongMarshallField(stage, "MinOccurs"))
	}
	if group.MaxOccurs != groupOther.MaxOccurs {
		diffs = append(diffs, group.GongMarshallField(stage, "MaxOccurs"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (length *Length) GongDiff(stage *Stage, lengthOther *Length) (diffs []string) {
	// insertion point for field diffs
	if length.Name != lengthOther.Name {
		diffs = append(diffs, length.GongMarshallField(stage, "Name"))
	}
	if (length.Annotation == nil) != (lengthOther.Annotation == nil) {
		diffs = append(diffs, length.GongMarshallField(stage, "Annotation"))
	} else if length.Annotation != nil && lengthOther.Annotation != nil {
		if length.Annotation != lengthOther.Annotation {
			diffs = append(diffs, length.GongMarshallField(stage, "Annotation"))
		}
	}
	if length.Value != lengthOther.Value {
		diffs = append(diffs, length.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (maxinclusive *MaxInclusive) GongDiff(stage *Stage, maxinclusiveOther *MaxInclusive) (diffs []string) {
	// insertion point for field diffs
	if maxinclusive.Name != maxinclusiveOther.Name {
		diffs = append(diffs, maxinclusive.GongMarshallField(stage, "Name"))
	}
	if (maxinclusive.Annotation == nil) != (maxinclusiveOther.Annotation == nil) {
		diffs = append(diffs, maxinclusive.GongMarshallField(stage, "Annotation"))
	} else if maxinclusive.Annotation != nil && maxinclusiveOther.Annotation != nil {
		if maxinclusive.Annotation != maxinclusiveOther.Annotation {
			diffs = append(diffs, maxinclusive.GongMarshallField(stage, "Annotation"))
		}
	}
	if maxinclusive.Value != maxinclusiveOther.Value {
		diffs = append(diffs, maxinclusive.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (maxlength *MaxLength) GongDiff(stage *Stage, maxlengthOther *MaxLength) (diffs []string) {
	// insertion point for field diffs
	if maxlength.Name != maxlengthOther.Name {
		diffs = append(diffs, maxlength.GongMarshallField(stage, "Name"))
	}
	if (maxlength.Annotation == nil) != (maxlengthOther.Annotation == nil) {
		diffs = append(diffs, maxlength.GongMarshallField(stage, "Annotation"))
	} else if maxlength.Annotation != nil && maxlengthOther.Annotation != nil {
		if maxlength.Annotation != maxlengthOther.Annotation {
			diffs = append(diffs, maxlength.GongMarshallField(stage, "Annotation"))
		}
	}
	if maxlength.Value != maxlengthOther.Value {
		diffs = append(diffs, maxlength.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (mininclusive *MinInclusive) GongDiff(stage *Stage, mininclusiveOther *MinInclusive) (diffs []string) {
	// insertion point for field diffs
	if mininclusive.Name != mininclusiveOther.Name {
		diffs = append(diffs, mininclusive.GongMarshallField(stage, "Name"))
	}
	if (mininclusive.Annotation == nil) != (mininclusiveOther.Annotation == nil) {
		diffs = append(diffs, mininclusive.GongMarshallField(stage, "Annotation"))
	} else if mininclusive.Annotation != nil && mininclusiveOther.Annotation != nil {
		if mininclusive.Annotation != mininclusiveOther.Annotation {
			diffs = append(diffs, mininclusive.GongMarshallField(stage, "Annotation"))
		}
	}
	if mininclusive.Value != mininclusiveOther.Value {
		diffs = append(diffs, mininclusive.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (minlength *MinLength) GongDiff(stage *Stage, minlengthOther *MinLength) (diffs []string) {
	// insertion point for field diffs
	if minlength.Name != minlengthOther.Name {
		diffs = append(diffs, minlength.GongMarshallField(stage, "Name"))
	}
	if (minlength.Annotation == nil) != (minlengthOther.Annotation == nil) {
		diffs = append(diffs, minlength.GongMarshallField(stage, "Annotation"))
	} else if minlength.Annotation != nil && minlengthOther.Annotation != nil {
		if minlength.Annotation != minlengthOther.Annotation {
			diffs = append(diffs, minlength.GongMarshallField(stage, "Annotation"))
		}
	}
	if minlength.Value != minlengthOther.Value {
		diffs = append(diffs, minlength.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pattern *Pattern) GongDiff(stage *Stage, patternOther *Pattern) (diffs []string) {
	// insertion point for field diffs
	if pattern.Name != patternOther.Name {
		diffs = append(diffs, pattern.GongMarshallField(stage, "Name"))
	}
	if (pattern.Annotation == nil) != (patternOther.Annotation == nil) {
		diffs = append(diffs, pattern.GongMarshallField(stage, "Annotation"))
	} else if pattern.Annotation != nil && patternOther.Annotation != nil {
		if pattern.Annotation != patternOther.Annotation {
			diffs = append(diffs, pattern.GongMarshallField(stage, "Annotation"))
		}
	}
	if pattern.Value != patternOther.Value {
		diffs = append(diffs, pattern.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (restriction *Restriction) GongDiff(stage *Stage, restrictionOther *Restriction) (diffs []string) {
	// insertion point for field diffs
	if restriction.Name != restrictionOther.Name {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Name"))
	}
	if (restriction.Annotation == nil) != (restrictionOther.Annotation == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Annotation"))
	} else if restriction.Annotation != nil && restrictionOther.Annotation != nil {
		if restriction.Annotation != restrictionOther.Annotation {
			diffs = append(diffs, restriction.GongMarshallField(stage, "Annotation"))
		}
	}
	if restriction.Base != restrictionOther.Base {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Base"))
	}
	EnumerationsDifferent := false
	if len(restriction.Enumerations) != len(restrictionOther.Enumerations) {
		EnumerationsDifferent = true
	} else {
		for i := range restriction.Enumerations {
			if (restriction.Enumerations[i] == nil) != (restrictionOther.Enumerations[i] == nil) {
				EnumerationsDifferent = true
				break
			} else if restriction.Enumerations[i] != nil && restrictionOther.Enumerations[i] != nil {
				// this is a pointer comparaison
				if restriction.Enumerations[i] != restrictionOther.Enumerations[i] {
					EnumerationsDifferent = true
					break
				}
			}
		}
	}
	if EnumerationsDifferent {
		ops := Diff(stage, restriction, restrictionOther, "Enumerations", restrictionOther.Enumerations, restriction.Enumerations)
		diffs = append(diffs, ops)
	}
	if (restriction.MinInclusive == nil) != (restrictionOther.MinInclusive == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MinInclusive"))
	} else if restriction.MinInclusive != nil && restrictionOther.MinInclusive != nil {
		if restriction.MinInclusive != restrictionOther.MinInclusive {
			diffs = append(diffs, restriction.GongMarshallField(stage, "MinInclusive"))
		}
	}
	if (restriction.MaxInclusive == nil) != (restrictionOther.MaxInclusive == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MaxInclusive"))
	} else if restriction.MaxInclusive != nil && restrictionOther.MaxInclusive != nil {
		if restriction.MaxInclusive != restrictionOther.MaxInclusive {
			diffs = append(diffs, restriction.GongMarshallField(stage, "MaxInclusive"))
		}
	}
	if (restriction.Pattern == nil) != (restrictionOther.Pattern == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Pattern"))
	} else if restriction.Pattern != nil && restrictionOther.Pattern != nil {
		if restriction.Pattern != restrictionOther.Pattern {
			diffs = append(diffs, restriction.GongMarshallField(stage, "Pattern"))
		}
	}
	if (restriction.WhiteSpace == nil) != (restrictionOther.WhiteSpace == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "WhiteSpace"))
	} else if restriction.WhiteSpace != nil && restrictionOther.WhiteSpace != nil {
		if restriction.WhiteSpace != restrictionOther.WhiteSpace {
			diffs = append(diffs, restriction.GongMarshallField(stage, "WhiteSpace"))
		}
	}
	if (restriction.MinLength == nil) != (restrictionOther.MinLength == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MinLength"))
	} else if restriction.MinLength != nil && restrictionOther.MinLength != nil {
		if restriction.MinLength != restrictionOther.MinLength {
			diffs = append(diffs, restriction.GongMarshallField(stage, "MinLength"))
		}
	}
	if (restriction.MaxLength == nil) != (restrictionOther.MaxLength == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MaxLength"))
	} else if restriction.MaxLength != nil && restrictionOther.MaxLength != nil {
		if restriction.MaxLength != restrictionOther.MaxLength {
			diffs = append(diffs, restriction.GongMarshallField(stage, "MaxLength"))
		}
	}
	if (restriction.Length == nil) != (restrictionOther.Length == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Length"))
	} else if restriction.Length != nil && restrictionOther.Length != nil {
		if restriction.Length != restrictionOther.Length {
			diffs = append(diffs, restriction.GongMarshallField(stage, "Length"))
		}
	}
	if (restriction.TotalDigit == nil) != (restrictionOther.TotalDigit == nil) {
		diffs = append(diffs, restriction.GongMarshallField(stage, "TotalDigit"))
	} else if restriction.TotalDigit != nil && restrictionOther.TotalDigit != nil {
		if restriction.TotalDigit != restrictionOther.TotalDigit {
			diffs = append(diffs, restriction.GongMarshallField(stage, "TotalDigit"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (schema *Schema) GongDiff(stage *Stage, schemaOther *Schema) (diffs []string) {
	// insertion point for field diffs
	if schema.Name != schemaOther.Name {
		diffs = append(diffs, schema.GongMarshallField(stage, "Name"))
	}
	if schema.Xs != schemaOther.Xs {
		diffs = append(diffs, schema.GongMarshallField(stage, "Xs"))
	}
	if (schema.Annotation == nil) != (schemaOther.Annotation == nil) {
		diffs = append(diffs, schema.GongMarshallField(stage, "Annotation"))
	} else if schema.Annotation != nil && schemaOther.Annotation != nil {
		if schema.Annotation != schemaOther.Annotation {
			diffs = append(diffs, schema.GongMarshallField(stage, "Annotation"))
		}
	}
	ElementsDifferent := false
	if len(schema.Elements) != len(schemaOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range schema.Elements {
			if (schema.Elements[i] == nil) != (schemaOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if schema.Elements[i] != nil && schemaOther.Elements[i] != nil {
				// this is a pointer comparaison
				if schema.Elements[i] != schemaOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, schema, schemaOther, "Elements", schemaOther.Elements, schema.Elements)
		diffs = append(diffs, ops)
	}
	SimpleTypesDifferent := false
	if len(schema.SimpleTypes) != len(schemaOther.SimpleTypes) {
		SimpleTypesDifferent = true
	} else {
		for i := range schema.SimpleTypes {
			if (schema.SimpleTypes[i] == nil) != (schemaOther.SimpleTypes[i] == nil) {
				SimpleTypesDifferent = true
				break
			} else if schema.SimpleTypes[i] != nil && schemaOther.SimpleTypes[i] != nil {
				// this is a pointer comparaison
				if schema.SimpleTypes[i] != schemaOther.SimpleTypes[i] {
					SimpleTypesDifferent = true
					break
				}
			}
		}
	}
	if SimpleTypesDifferent {
		ops := Diff(stage, schema, schemaOther, "SimpleTypes", schemaOther.SimpleTypes, schema.SimpleTypes)
		diffs = append(diffs, ops)
	}
	ComplexTypesDifferent := false
	if len(schema.ComplexTypes) != len(schemaOther.ComplexTypes) {
		ComplexTypesDifferent = true
	} else {
		for i := range schema.ComplexTypes {
			if (schema.ComplexTypes[i] == nil) != (schemaOther.ComplexTypes[i] == nil) {
				ComplexTypesDifferent = true
				break
			} else if schema.ComplexTypes[i] != nil && schemaOther.ComplexTypes[i] != nil {
				// this is a pointer comparaison
				if schema.ComplexTypes[i] != schemaOther.ComplexTypes[i] {
					ComplexTypesDifferent = true
					break
				}
			}
		}
	}
	if ComplexTypesDifferent {
		ops := Diff(stage, schema, schemaOther, "ComplexTypes", schemaOther.ComplexTypes, schema.ComplexTypes)
		diffs = append(diffs, ops)
	}
	AttributeGroupsDifferent := false
	if len(schema.AttributeGroups) != len(schemaOther.AttributeGroups) {
		AttributeGroupsDifferent = true
	} else {
		for i := range schema.AttributeGroups {
			if (schema.AttributeGroups[i] == nil) != (schemaOther.AttributeGroups[i] == nil) {
				AttributeGroupsDifferent = true
				break
			} else if schema.AttributeGroups[i] != nil && schemaOther.AttributeGroups[i] != nil {
				// this is a pointer comparaison
				if schema.AttributeGroups[i] != schemaOther.AttributeGroups[i] {
					AttributeGroupsDifferent = true
					break
				}
			}
		}
	}
	if AttributeGroupsDifferent {
		ops := Diff(stage, schema, schemaOther, "AttributeGroups", schemaOther.AttributeGroups, schema.AttributeGroups)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(schema.Groups) != len(schemaOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range schema.Groups {
			if (schema.Groups[i] == nil) != (schemaOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if schema.Groups[i] != nil && schemaOther.Groups[i] != nil {
				// this is a pointer comparaison
				if schema.Groups[i] != schemaOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, schema, schemaOther, "Groups", schemaOther.Groups, schema.Groups)
		diffs = append(diffs, ops)
	}
	if schema.Order != schemaOther.Order {
		diffs = append(diffs, schema.GongMarshallField(stage, "Order"))
	}
	if schema.Depth != schemaOther.Depth {
		diffs = append(diffs, schema.GongMarshallField(stage, "Depth"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sequence *Sequence) GongDiff(stage *Stage, sequenceOther *Sequence) (diffs []string) {
	// insertion point for field diffs
	if sequence.Name != sequenceOther.Name {
		diffs = append(diffs, sequence.GongMarshallField(stage, "Name"))
	}
	if (sequence.Annotation == nil) != (sequenceOther.Annotation == nil) {
		diffs = append(diffs, sequence.GongMarshallField(stage, "Annotation"))
	} else if sequence.Annotation != nil && sequenceOther.Annotation != nil {
		if sequence.Annotation != sequenceOther.Annotation {
			diffs = append(diffs, sequence.GongMarshallField(stage, "Annotation"))
		}
	}
	if sequence.OuterElementName != sequenceOther.OuterElementName {
		diffs = append(diffs, sequence.GongMarshallField(stage, "OuterElementName"))
	}
	SequencesDifferent := false
	if len(sequence.Sequences) != len(sequenceOther.Sequences) {
		SequencesDifferent = true
	} else {
		for i := range sequence.Sequences {
			if (sequence.Sequences[i] == nil) != (sequenceOther.Sequences[i] == nil) {
				SequencesDifferent = true
				break
			} else if sequence.Sequences[i] != nil && sequenceOther.Sequences[i] != nil {
				// this is a pointer comparaison
				if sequence.Sequences[i] != sequenceOther.Sequences[i] {
					SequencesDifferent = true
					break
				}
			}
		}
	}
	if SequencesDifferent {
		ops := Diff(stage, sequence, sequenceOther, "Sequences", sequenceOther.Sequences, sequence.Sequences)
		diffs = append(diffs, ops)
	}
	AllsDifferent := false
	if len(sequence.Alls) != len(sequenceOther.Alls) {
		AllsDifferent = true
	} else {
		for i := range sequence.Alls {
			if (sequence.Alls[i] == nil) != (sequenceOther.Alls[i] == nil) {
				AllsDifferent = true
				break
			} else if sequence.Alls[i] != nil && sequenceOther.Alls[i] != nil {
				// this is a pointer comparaison
				if sequence.Alls[i] != sequenceOther.Alls[i] {
					AllsDifferent = true
					break
				}
			}
		}
	}
	if AllsDifferent {
		ops := Diff(stage, sequence, sequenceOther, "Alls", sequenceOther.Alls, sequence.Alls)
		diffs = append(diffs, ops)
	}
	ChoicesDifferent := false
	if len(sequence.Choices) != len(sequenceOther.Choices) {
		ChoicesDifferent = true
	} else {
		for i := range sequence.Choices {
			if (sequence.Choices[i] == nil) != (sequenceOther.Choices[i] == nil) {
				ChoicesDifferent = true
				break
			} else if sequence.Choices[i] != nil && sequenceOther.Choices[i] != nil {
				// this is a pointer comparaison
				if sequence.Choices[i] != sequenceOther.Choices[i] {
					ChoicesDifferent = true
					break
				}
			}
		}
	}
	if ChoicesDifferent {
		ops := Diff(stage, sequence, sequenceOther, "Choices", sequenceOther.Choices, sequence.Choices)
		diffs = append(diffs, ops)
	}
	GroupsDifferent := false
	if len(sequence.Groups) != len(sequenceOther.Groups) {
		GroupsDifferent = true
	} else {
		for i := range sequence.Groups {
			if (sequence.Groups[i] == nil) != (sequenceOther.Groups[i] == nil) {
				GroupsDifferent = true
				break
			} else if sequence.Groups[i] != nil && sequenceOther.Groups[i] != nil {
				// this is a pointer comparaison
				if sequence.Groups[i] != sequenceOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := Diff(stage, sequence, sequenceOther, "Groups", sequenceOther.Groups, sequence.Groups)
		diffs = append(diffs, ops)
	}
	ElementsDifferent := false
	if len(sequence.Elements) != len(sequenceOther.Elements) {
		ElementsDifferent = true
	} else {
		for i := range sequence.Elements {
			if (sequence.Elements[i] == nil) != (sequenceOther.Elements[i] == nil) {
				ElementsDifferent = true
				break
			} else if sequence.Elements[i] != nil && sequenceOther.Elements[i] != nil {
				// this is a pointer comparaison
				if sequence.Elements[i] != sequenceOther.Elements[i] {
					ElementsDifferent = true
					break
				}
			}
		}
	}
	if ElementsDifferent {
		ops := Diff(stage, sequence, sequenceOther, "Elements", sequenceOther.Elements, sequence.Elements)
		diffs = append(diffs, ops)
	}
	if sequence.Order != sequenceOther.Order {
		diffs = append(diffs, sequence.GongMarshallField(stage, "Order"))
	}
	if sequence.Depth != sequenceOther.Depth {
		diffs = append(diffs, sequence.GongMarshallField(stage, "Depth"))
	}
	if sequence.MinOccurs != sequenceOther.MinOccurs {
		diffs = append(diffs, sequence.GongMarshallField(stage, "MinOccurs"))
	}
	if sequence.MaxOccurs != sequenceOther.MaxOccurs {
		diffs = append(diffs, sequence.GongMarshallField(stage, "MaxOccurs"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (simplecontent *SimpleContent) GongDiff(stage *Stage, simplecontentOther *SimpleContent) (diffs []string) {
	// insertion point for field diffs
	if simplecontent.Name != simplecontentOther.Name {
		diffs = append(diffs, simplecontent.GongMarshallField(stage, "Name"))
	}
	if (simplecontent.Extension == nil) != (simplecontentOther.Extension == nil) {
		diffs = append(diffs, simplecontent.GongMarshallField(stage, "Extension"))
	} else if simplecontent.Extension != nil && simplecontentOther.Extension != nil {
		if simplecontent.Extension != simplecontentOther.Extension {
			diffs = append(diffs, simplecontent.GongMarshallField(stage, "Extension"))
		}
	}
	if (simplecontent.Restriction == nil) != (simplecontentOther.Restriction == nil) {
		diffs = append(diffs, simplecontent.GongMarshallField(stage, "Restriction"))
	} else if simplecontent.Restriction != nil && simplecontentOther.Restriction != nil {
		if simplecontent.Restriction != simplecontentOther.Restriction {
			diffs = append(diffs, simplecontent.GongMarshallField(stage, "Restriction"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (simpletype *SimpleType) GongDiff(stage *Stage, simpletypeOther *SimpleType) (diffs []string) {
	// insertion point for field diffs
	if simpletype.Name != simpletypeOther.Name {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Name"))
	}
	if (simpletype.Annotation == nil) != (simpletypeOther.Annotation == nil) {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Annotation"))
	} else if simpletype.Annotation != nil && simpletypeOther.Annotation != nil {
		if simpletype.Annotation != simpletypeOther.Annotation {
			diffs = append(diffs, simpletype.GongMarshallField(stage, "Annotation"))
		}
	}
	if simpletype.NameXSD != simpletypeOther.NameXSD {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "NameXSD"))
	}
	if (simpletype.Restriction == nil) != (simpletypeOther.Restriction == nil) {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Restriction"))
	} else if simpletype.Restriction != nil && simpletypeOther.Restriction != nil {
		if simpletype.Restriction != simpletypeOther.Restriction {
			diffs = append(diffs, simpletype.GongMarshallField(stage, "Restriction"))
		}
	}
	if (simpletype.Union == nil) != (simpletypeOther.Union == nil) {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Union"))
	} else if simpletype.Union != nil && simpletypeOther.Union != nil {
		if simpletype.Union != simpletypeOther.Union {
			diffs = append(diffs, simpletype.GongMarshallField(stage, "Union"))
		}
	}
	if simpletype.Order != simpletypeOther.Order {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Order"))
	}
	if simpletype.Depth != simpletypeOther.Depth {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Depth"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (totaldigit *TotalDigit) GongDiff(stage *Stage, totaldigitOther *TotalDigit) (diffs []string) {
	// insertion point for field diffs
	if totaldigit.Name != totaldigitOther.Name {
		diffs = append(diffs, totaldigit.GongMarshallField(stage, "Name"))
	}
	if (totaldigit.Annotation == nil) != (totaldigitOther.Annotation == nil) {
		diffs = append(diffs, totaldigit.GongMarshallField(stage, "Annotation"))
	} else if totaldigit.Annotation != nil && totaldigitOther.Annotation != nil {
		if totaldigit.Annotation != totaldigitOther.Annotation {
			diffs = append(diffs, totaldigit.GongMarshallField(stage, "Annotation"))
		}
	}
	if totaldigit.Value != totaldigitOther.Value {
		diffs = append(diffs, totaldigit.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (union *Union) GongDiff(stage *Stage, unionOther *Union) (diffs []string) {
	// insertion point for field diffs
	if union.Name != unionOther.Name {
		diffs = append(diffs, union.GongMarshallField(stage, "Name"))
	}
	if (union.Annotation == nil) != (unionOther.Annotation == nil) {
		diffs = append(diffs, union.GongMarshallField(stage, "Annotation"))
	} else if union.Annotation != nil && unionOther.Annotation != nil {
		if union.Annotation != unionOther.Annotation {
			diffs = append(diffs, union.GongMarshallField(stage, "Annotation"))
		}
	}
	if union.MemberTypes != unionOther.MemberTypes {
		diffs = append(diffs, union.GongMarshallField(stage, "MemberTypes"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (whitespace *WhiteSpace) GongDiff(stage *Stage, whitespaceOther *WhiteSpace) (diffs []string) {
	// insertion point for field diffs
	if whitespace.Name != whitespaceOther.Name {
		diffs = append(diffs, whitespace.GongMarshallField(stage, "Name"))
	}
	if (whitespace.Annotation == nil) != (whitespaceOther.Annotation == nil) {
		diffs = append(diffs, whitespace.GongMarshallField(stage, "Annotation"))
	} else if whitespace.Annotation != nil && whitespaceOther.Annotation != nil {
		if whitespace.Annotation != whitespaceOther.Annotation {
			diffs = append(diffs, whitespace.GongMarshallField(stage, "Annotation"))
		}
	}
	if whitespace.Value != whitespaceOther.Value {
		diffs = append(diffs, whitespace.GongMarshallField(stage, "Value"))
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
