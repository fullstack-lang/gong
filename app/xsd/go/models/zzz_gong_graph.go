// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (all *All) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Alls[all]
	return ok
}

func (annotation *Annotation) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Annotations[annotation]
	return ok
}

func (attribute *Attribute) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Attributes[attribute]
	return ok
}

func (attributegroup *AttributeGroup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AttributeGroups[attributegroup]
	return ok
}

func (choice *Choice) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Choices[choice]
	return ok
}

func (complexcontent *ComplexContent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ComplexContents[complexcontent]
	return ok
}

func (complextype *ComplexType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ComplexTypes[complextype]
	return ok
}

func (documentation *Documentation) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Documentations[documentation]
	return ok
}

func (element *Element) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Elements[element]
	return ok
}

func (enumeration *Enumeration) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Enumerations[enumeration]
	return ok
}

func (extension *Extension) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Extensions[extension]
	return ok
}

func (group *Group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groups[group]
	return ok
}

func (length *Length) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Lengths[length]
	return ok
}

func (maxinclusive *MaxInclusive) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MaxInclusives[maxinclusive]
	return ok
}

func (maxlength *MaxLength) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MaxLengths[maxlength]
	return ok
}

func (mininclusive *MinInclusive) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MinInclusives[mininclusive]
	return ok
}

func (minlength *MinLength) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MinLengths[minlength]
	return ok
}

func (pattern *Pattern) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Patterns[pattern]
	return ok
}

func (restriction *Restriction) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Restrictions[restriction]
	return ok
}

func (schema *Schema) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Schemas[schema]
	return ok
}

func (sequence *Sequence) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sequences[sequence]
	return ok
}

func (simplecontent *SimpleContent) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SimpleContents[simplecontent]
	return ok
}

func (simpletype *SimpleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SimpleTypes[simpletype]
	return ok
}

func (totaldigit *TotalDigit) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TotalDigits[totaldigit]
	return ok
}

func (union *Union) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Unions[union]
	return ok
}

func (whitespace *WhiteSpace) GongIsStaged(stage *Stage) bool {
	_, ok := stage.WhiteSpaces[whitespace]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (all *All) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(all) {
		return
	}

	all.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		stage.StageBranch(all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range all.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range all.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range all.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range all.Elements {
		stage.StageBranch(_element)
	}

}

func (annotation *Annotation) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(annotation) {
		return
	}

	annotation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		stage.StageBranch(_documentation)
	}

}

func (attribute *Attribute) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attribute) {
		return
	}

	attribute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		stage.StageBranch(attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attributegroup *AttributeGroup) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attributegroup) {
		return
	}

	attributegroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		stage.StageBranch(attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		stage.StageBranch(_attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		stage.StageBranch(_attribute)
	}

}

func (choice *Choice) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(choice) {
		return
	}

	choice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		stage.StageBranch(choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range choice.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range choice.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range choice.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range choice.Elements {
		stage.StageBranch(_element)
	}

}

func (complexcontent *ComplexContent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(complexcontent) {
		return
	}

	complexcontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (complextype *ComplexType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(complextype) {
		return
	}

	complextype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		stage.StageBranch(complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		stage.StageBranch(complextype.Annotation)
	}
	if complextype.Extension != nil {
		stage.StageBranch(complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		stage.StageBranch(complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		stage.StageBranch(complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range complextype.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range complextype.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range complextype.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range complextype.Elements {
		stage.StageBranch(_element)
	}
	for _, _attribute := range complextype.Attributes {
		stage.StageBranch(_attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		stage.StageBranch(_attributegroup)
	}

}

func (documentation *Documentation) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(documentation) {
		return
	}

	documentation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (element *Element) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(element) {
		return
	}

	element.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		stage.StageBranch(element.Annotation)
	}
	if element.SimpleType != nil {
		stage.StageBranch(element.SimpleType)
	}
	if element.ComplexType != nil {
		stage.StageBranch(element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		stage.StageBranch(_group)
	}

}

func (enumeration *Enumeration) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(enumeration) {
		return
	}

	enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		stage.StageBranch(enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extension *Extension) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(extension) {
		return
	}

	extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range extension.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range extension.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range extension.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range extension.Elements {
		stage.StageBranch(_element)
	}
	for _, _attribute := range extension.Attributes {
		stage.StageBranch(_attribute)
	}
	for _, _attributegroup := range extension.AttributeGroups {
		stage.StageBranch(_attributegroup)
	}

}

func (group *Group) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		stage.StageBranch(group.Annotation)
	}
	if group.OuterElement != nil {
		stage.StageBranch(group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range group.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range group.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range group.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range group.Elements {
		stage.StageBranch(_element)
	}

}

func (length *Length) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(length) {
		return
	}

	length.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		stage.StageBranch(length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (maxinclusive *MaxInclusive) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(maxinclusive) {
		return
	}

	maxinclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		stage.StageBranch(maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (maxlength *MaxLength) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(maxlength) {
		return
	}

	maxlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		stage.StageBranch(maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mininclusive *MinInclusive) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(mininclusive) {
		return
	}

	mininclusive.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		stage.StageBranch(mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (minlength *MinLength) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(minlength) {
		return
	}

	minlength.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		stage.StageBranch(minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pattern *Pattern) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pattern) {
		return
	}

	pattern.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		stage.StageBranch(pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (restriction *Restriction) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(restriction) {
		return
	}

	restriction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		stage.StageBranch(restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		stage.StageBranch(restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		stage.StageBranch(restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		stage.StageBranch(restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		stage.StageBranch(restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		stage.StageBranch(restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		stage.StageBranch(restriction.MaxLength)
	}
	if restriction.Length != nil {
		stage.StageBranch(restriction.Length)
	}
	if restriction.TotalDigit != nil {
		stage.StageBranch(restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		stage.StageBranch(_enumeration)
	}

}

func (schema *Schema) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(schema) {
		return
	}

	schema.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		stage.StageBranch(schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		stage.StageBranch(_element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		stage.StageBranch(_simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		stage.StageBranch(_complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		stage.StageBranch(_attributegroup)
	}
	for _, _group := range schema.Groups {
		stage.StageBranch(_group)
	}

}

func (sequence *Sequence) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(sequence) {
		return
	}

	sequence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		stage.StageBranch(sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		stage.StageBranch(_sequence)
	}
	for _, _all := range sequence.Alls {
		stage.StageBranch(_all)
	}
	for _, _choice := range sequence.Choices {
		stage.StageBranch(_choice)
	}
	for _, _group := range sequence.Groups {
		stage.StageBranch(_group)
	}
	for _, _element := range sequence.Elements {
		stage.StageBranch(_element)
	}

}

func (simplecontent *SimpleContent) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(simplecontent) {
		return
	}

	simplecontent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		stage.StageBranch(simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		stage.StageBranch(simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (simpletype *SimpleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(simpletype) {
		return
	}

	simpletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		stage.StageBranch(simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		stage.StageBranch(simpletype.Restriction)
	}
	if simpletype.Union != nil {
		stage.StageBranch(simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (totaldigit *TotalDigit) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(totaldigit) {
		return
	}

	totaldigit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		stage.StageBranch(totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (union *Union) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(union) {
		return
	}

	union.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		stage.StageBranch(union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (whitespace *WhiteSpace) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(whitespace) {
		return
	}

	whitespace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		stage.StageBranch(whitespace.Annotation)
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
	case *All:
		toT := GongCopyBranchAll(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Annotation:
		toT := GongCopyBranchAnnotation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Attribute:
		toT := GongCopyBranchAttribute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AttributeGroup:
		toT := GongCopyBranchAttributeGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Choice:
		toT := GongCopyBranchChoice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexContent:
		toT := GongCopyBranchComplexContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ComplexType:
		toT := GongCopyBranchComplexType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Documentation:
		toT := GongCopyBranchDocumentation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Element:
		toT := GongCopyBranchElement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Enumeration:
		toT := GongCopyBranchEnumeration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extension:
		toT := GongCopyBranchExtension(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := GongCopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Length:
		toT := GongCopyBranchLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxInclusive:
		toT := GongCopyBranchMaxInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MaxLength:
		toT := GongCopyBranchMaxLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinInclusive:
		toT := GongCopyBranchMinInclusive(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MinLength:
		toT := GongCopyBranchMinLength(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pattern:
		toT := GongCopyBranchPattern(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Restriction:
		toT := GongCopyBranchRestriction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Schema:
		toT := GongCopyBranchSchema(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sequence:
		toT := GongCopyBranchSequence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleContent:
		toT := GongCopyBranchSimpleContent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SimpleType:
		toT := GongCopyBranchSimpleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TotalDigit:
		toT := GongCopyBranchTotalDigit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Union:
		toT := GongCopyBranchUnion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *WhiteSpace:
		toT := GongCopyBranchWhiteSpace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAll(mapOrigCopy map[any]any, allFrom *All) (allTo *All) {
	var alreadyCopied bool
	allTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, allFrom)
	if alreadyCopied {
		return
	}
	allFrom.GongCopyBasicFields(allTo)

	//insertion point for the staging of instances referenced by pointers
	if allFrom.Annotation != nil {
		allTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, allFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range allFrom.Sequences {
		allTo.Sequences = append(allTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range allFrom.Alls {
		allTo.Alls = append(allTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range allFrom.Choices {
		allTo.Choices = append(allTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range allFrom.Groups {
		allTo.Groups = append(allTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range allFrom.Elements {
		allTo.Elements = append(allTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func GongCopyBranchAnnotation(mapOrigCopy map[any]any, annotationFrom *Annotation) (annotationTo *Annotation) {
	var alreadyCopied bool
	annotationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, annotationFrom)
	if alreadyCopied {
		return
	}
	annotationFrom.GongCopyBasicFields(annotationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotationFrom.Documentations {
		annotationTo.Documentations = append(annotationTo.Documentations, GongCopyBranchDocumentation(mapOrigCopy, _documentation))
	}

	return
}

func GongCopyBranchAttribute(mapOrigCopy map[any]any, attributeFrom *Attribute) (attributeTo *Attribute) {
	var alreadyCopied bool
	attributeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributeFrom)
	if alreadyCopied {
		return
	}
	attributeFrom.GongCopyBasicFields(attributeTo)

	//insertion point for the staging of instances referenced by pointers
	if attributeFrom.Annotation != nil {
		attributeTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, attributeFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAttributeGroup(mapOrigCopy map[any]any, attributegroupFrom *AttributeGroup) (attributegroupTo *AttributeGroup) {
	var alreadyCopied bool
	attributegroupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributegroupFrom)
	if alreadyCopied {
		return
	}
	attributegroupFrom.GongCopyBasicFields(attributegroupTo)

	//insertion point for the staging of instances referenced by pointers
	if attributegroupFrom.Annotation != nil {
		attributegroupTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, attributegroupFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroupFrom.AttributeGroups {
		attributegroupTo.AttributeGroups = append(attributegroupTo.AttributeGroups, GongCopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _attribute := range attributegroupFrom.Attributes {
		attributegroupTo.Attributes = append(attributegroupTo.Attributes, GongCopyBranchAttribute(mapOrigCopy, _attribute))
	}

	return
}

func GongCopyBranchChoice(mapOrigCopy map[any]any, choiceFrom *Choice) (choiceTo *Choice) {
	var alreadyCopied bool
	choiceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, choiceFrom)
	if alreadyCopied {
		return
	}
	choiceFrom.GongCopyBasicFields(choiceTo)

	//insertion point for the staging of instances referenced by pointers
	if choiceFrom.Annotation != nil {
		choiceTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, choiceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choiceFrom.Sequences {
		choiceTo.Sequences = append(choiceTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range choiceFrom.Alls {
		choiceTo.Alls = append(choiceTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range choiceFrom.Choices {
		choiceTo.Choices = append(choiceTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range choiceFrom.Groups {
		choiceTo.Groups = append(choiceTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range choiceFrom.Elements {
		choiceTo.Elements = append(choiceTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func GongCopyBranchComplexContent(mapOrigCopy map[any]any, complexcontentFrom *ComplexContent) (complexcontentTo *ComplexContent) {
	var alreadyCopied bool
	complexcontentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, complexcontentFrom)
	if alreadyCopied {
		return
	}
	complexcontentFrom.GongCopyBasicFields(complexcontentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchComplexType(mapOrigCopy map[any]any, complextypeFrom *ComplexType) (complextypeTo *ComplexType) {
	var alreadyCopied bool
	complextypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, complextypeFrom)
	if alreadyCopied {
		return
	}
	complextypeFrom.GongCopyBasicFields(complextypeTo)

	//insertion point for the staging of instances referenced by pointers
	if complextypeFrom.OuterElement != nil {
		complextypeTo.OuterElement = GongCopyBranchElement(mapOrigCopy, complextypeFrom.OuterElement)
	}
	if complextypeFrom.Annotation != nil {
		complextypeTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, complextypeFrom.Annotation)
	}
	if complextypeFrom.Extension != nil {
		complextypeTo.Extension = GongCopyBranchExtension(mapOrigCopy, complextypeFrom.Extension)
	}
	if complextypeFrom.SimpleContent != nil {
		complextypeTo.SimpleContent = GongCopyBranchSimpleContent(mapOrigCopy, complextypeFrom.SimpleContent)
	}
	if complextypeFrom.ComplexContent != nil {
		complextypeTo.ComplexContent = GongCopyBranchComplexContent(mapOrigCopy, complextypeFrom.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextypeFrom.Sequences {
		complextypeTo.Sequences = append(complextypeTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range complextypeFrom.Alls {
		complextypeTo.Alls = append(complextypeTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range complextypeFrom.Choices {
		complextypeTo.Choices = append(complextypeTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range complextypeFrom.Groups {
		complextypeTo.Groups = append(complextypeTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range complextypeFrom.Elements {
		complextypeTo.Elements = append(complextypeTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range complextypeFrom.Attributes {
		complextypeTo.Attributes = append(complextypeTo.Attributes, GongCopyBranchAttribute(mapOrigCopy, _attribute))
	}
	for _, _attributegroup := range complextypeFrom.AttributeGroups {
		complextypeTo.AttributeGroups = append(complextypeTo.AttributeGroups, GongCopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}

	return
}

func GongCopyBranchDocumentation(mapOrigCopy map[any]any, documentationFrom *Documentation) (documentationTo *Documentation) {
	var alreadyCopied bool
	documentationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, documentationFrom)
	if alreadyCopied {
		return
	}
	documentationFrom.GongCopyBasicFields(documentationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchElement(mapOrigCopy map[any]any, elementFrom *Element) (elementTo *Element) {
	var alreadyCopied bool
	elementTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, elementFrom)
	if alreadyCopied {
		return
	}
	elementFrom.GongCopyBasicFields(elementTo)

	//insertion point for the staging of instances referenced by pointers
	if elementFrom.Annotation != nil {
		elementTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, elementFrom.Annotation)
	}
	if elementFrom.SimpleType != nil {
		elementTo.SimpleType = GongCopyBranchSimpleType(mapOrigCopy, elementFrom.SimpleType)
	}
	if elementFrom.ComplexType != nil {
		elementTo.ComplexType = GongCopyBranchComplexType(mapOrigCopy, elementFrom.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range elementFrom.Groups {
		elementTo.Groups = append(elementTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func GongCopyBranchEnumeration(mapOrigCopy map[any]any, enumerationFrom *Enumeration) (enumerationTo *Enumeration) {
	var alreadyCopied bool
	enumerationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, enumerationFrom)
	if alreadyCopied {
		return
	}
	enumerationFrom.GongCopyBasicFields(enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if enumerationFrom.Annotation != nil {
		enumerationTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, enumerationFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchExtension(mapOrigCopy map[any]any, extensionFrom *Extension) (extensionTo *Extension) {
	var alreadyCopied bool
	extensionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, extensionFrom)
	if alreadyCopied {
		return
	}
	extensionFrom.GongCopyBasicFields(extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extensionFrom.Sequences {
		extensionTo.Sequences = append(extensionTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range extensionFrom.Alls {
		extensionTo.Alls = append(extensionTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range extensionFrom.Choices {
		extensionTo.Choices = append(extensionTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range extensionFrom.Groups {
		extensionTo.Groups = append(extensionTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range extensionFrom.Elements {
		extensionTo.Elements = append(extensionTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}
	for _, _attribute := range extensionFrom.Attributes {
		extensionTo.Attributes = append(extensionTo.Attributes, GongCopyBranchAttribute(mapOrigCopy, _attribute))
	}
	for _, _attributegroup := range extensionFrom.AttributeGroups {
		extensionTo.AttributeGroups = append(extensionTo.AttributeGroups, GongCopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}

	return
}

func GongCopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {
	var alreadyCopied bool
	groupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupFrom)
	if alreadyCopied {
		return
	}
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers
	if groupFrom.Annotation != nil {
		groupTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, groupFrom.Annotation)
	}
	if groupFrom.OuterElement != nil {
		groupTo.OuterElement = GongCopyBranchElement(mapOrigCopy, groupFrom.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range groupFrom.Sequences {
		groupTo.Sequences = append(groupTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range groupFrom.Alls {
		groupTo.Alls = append(groupTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range groupFrom.Choices {
		groupTo.Choices = append(groupTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range groupFrom.Groups {
		groupTo.Groups = append(groupTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range groupFrom.Elements {
		groupTo.Elements = append(groupTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func GongCopyBranchLength(mapOrigCopy map[any]any, lengthFrom *Length) (lengthTo *Length) {
	var alreadyCopied bool
	lengthTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, lengthFrom)
	if alreadyCopied {
		return
	}
	lengthFrom.GongCopyBasicFields(lengthTo)

	//insertion point for the staging of instances referenced by pointers
	if lengthFrom.Annotation != nil {
		lengthTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, lengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMaxInclusive(mapOrigCopy map[any]any, maxinclusiveFrom *MaxInclusive) (maxinclusiveTo *MaxInclusive) {
	var alreadyCopied bool
	maxinclusiveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, maxinclusiveFrom)
	if alreadyCopied {
		return
	}
	maxinclusiveFrom.GongCopyBasicFields(maxinclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusiveFrom.Annotation != nil {
		maxinclusiveTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, maxinclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMaxLength(mapOrigCopy map[any]any, maxlengthFrom *MaxLength) (maxlengthTo *MaxLength) {
	var alreadyCopied bool
	maxlengthTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, maxlengthFrom)
	if alreadyCopied {
		return
	}
	maxlengthFrom.GongCopyBasicFields(maxlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if maxlengthFrom.Annotation != nil {
		maxlengthTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, maxlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMinInclusive(mapOrigCopy map[any]any, mininclusiveFrom *MinInclusive) (mininclusiveTo *MinInclusive) {
	var alreadyCopied bool
	mininclusiveTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, mininclusiveFrom)
	if alreadyCopied {
		return
	}
	mininclusiveFrom.GongCopyBasicFields(mininclusiveTo)

	//insertion point for the staging of instances referenced by pointers
	if mininclusiveFrom.Annotation != nil {
		mininclusiveTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, mininclusiveFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMinLength(mapOrigCopy map[any]any, minlengthFrom *MinLength) (minlengthTo *MinLength) {
	var alreadyCopied bool
	minlengthTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, minlengthFrom)
	if alreadyCopied {
		return
	}
	minlengthFrom.GongCopyBasicFields(minlengthTo)

	//insertion point for the staging of instances referenced by pointers
	if minlengthFrom.Annotation != nil {
		minlengthTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, minlengthFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPattern(mapOrigCopy map[any]any, patternFrom *Pattern) (patternTo *Pattern) {
	var alreadyCopied bool
	patternTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, patternFrom)
	if alreadyCopied {
		return
	}
	patternFrom.GongCopyBasicFields(patternTo)

	//insertion point for the staging of instances referenced by pointers
	if patternFrom.Annotation != nil {
		patternTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, patternFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRestriction(mapOrigCopy map[any]any, restrictionFrom *Restriction) (restrictionTo *Restriction) {
	var alreadyCopied bool
	restrictionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, restrictionFrom)
	if alreadyCopied {
		return
	}
	restrictionFrom.GongCopyBasicFields(restrictionTo)

	//insertion point for the staging of instances referenced by pointers
	if restrictionFrom.Annotation != nil {
		restrictionTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, restrictionFrom.Annotation)
	}
	if restrictionFrom.MinInclusive != nil {
		restrictionTo.MinInclusive = GongCopyBranchMinInclusive(mapOrigCopy, restrictionFrom.MinInclusive)
	}
	if restrictionFrom.MaxInclusive != nil {
		restrictionTo.MaxInclusive = GongCopyBranchMaxInclusive(mapOrigCopy, restrictionFrom.MaxInclusive)
	}
	if restrictionFrom.Pattern != nil {
		restrictionTo.Pattern = GongCopyBranchPattern(mapOrigCopy, restrictionFrom.Pattern)
	}
	if restrictionFrom.WhiteSpace != nil {
		restrictionTo.WhiteSpace = GongCopyBranchWhiteSpace(mapOrigCopy, restrictionFrom.WhiteSpace)
	}
	if restrictionFrom.MinLength != nil {
		restrictionTo.MinLength = GongCopyBranchMinLength(mapOrigCopy, restrictionFrom.MinLength)
	}
	if restrictionFrom.MaxLength != nil {
		restrictionTo.MaxLength = GongCopyBranchMaxLength(mapOrigCopy, restrictionFrom.MaxLength)
	}
	if restrictionFrom.Length != nil {
		restrictionTo.Length = GongCopyBranchLength(mapOrigCopy, restrictionFrom.Length)
	}
	if restrictionFrom.TotalDigit != nil {
		restrictionTo.TotalDigit = GongCopyBranchTotalDigit(mapOrigCopy, restrictionFrom.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restrictionFrom.Enumerations {
		restrictionTo.Enumerations = append(restrictionTo.Enumerations, GongCopyBranchEnumeration(mapOrigCopy, _enumeration))
	}

	return
}

func GongCopyBranchSchema(mapOrigCopy map[any]any, schemaFrom *Schema) (schemaTo *Schema) {
	var alreadyCopied bool
	schemaTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, schemaFrom)
	if alreadyCopied {
		return
	}
	schemaFrom.GongCopyBasicFields(schemaTo)

	//insertion point for the staging of instances referenced by pointers
	if schemaFrom.Annotation != nil {
		schemaTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, schemaFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schemaFrom.Elements {
		schemaTo.Elements = append(schemaTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}
	for _, _simpletype := range schemaFrom.SimpleTypes {
		schemaTo.SimpleTypes = append(schemaTo.SimpleTypes, GongCopyBranchSimpleType(mapOrigCopy, _simpletype))
	}
	for _, _complextype := range schemaFrom.ComplexTypes {
		schemaTo.ComplexTypes = append(schemaTo.ComplexTypes, GongCopyBranchComplexType(mapOrigCopy, _complextype))
	}
	for _, _attributegroup := range schemaFrom.AttributeGroups {
		schemaTo.AttributeGroups = append(schemaTo.AttributeGroups, GongCopyBranchAttributeGroup(mapOrigCopy, _attributegroup))
	}
	for _, _group := range schemaFrom.Groups {
		schemaTo.Groups = append(schemaTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func GongCopyBranchSequence(mapOrigCopy map[any]any, sequenceFrom *Sequence) (sequenceTo *Sequence) {
	var alreadyCopied bool
	sequenceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, sequenceFrom)
	if alreadyCopied {
		return
	}
	sequenceFrom.GongCopyBasicFields(sequenceTo)

	//insertion point for the staging of instances referenced by pointers
	if sequenceFrom.Annotation != nil {
		sequenceTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, sequenceFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequenceFrom.Sequences {
		sequenceTo.Sequences = append(sequenceTo.Sequences, GongCopyBranchSequence(mapOrigCopy, _sequence))
	}
	for _, _all := range sequenceFrom.Alls {
		sequenceTo.Alls = append(sequenceTo.Alls, GongCopyBranchAll(mapOrigCopy, _all))
	}
	for _, _choice := range sequenceFrom.Choices {
		sequenceTo.Choices = append(sequenceTo.Choices, GongCopyBranchChoice(mapOrigCopy, _choice))
	}
	for _, _group := range sequenceFrom.Groups {
		sequenceTo.Groups = append(sequenceTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _element := range sequenceFrom.Elements {
		sequenceTo.Elements = append(sequenceTo.Elements, GongCopyBranchElement(mapOrigCopy, _element))
	}

	return
}

func GongCopyBranchSimpleContent(mapOrigCopy map[any]any, simplecontentFrom *SimpleContent) (simplecontentTo *SimpleContent) {
	var alreadyCopied bool
	simplecontentTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, simplecontentFrom)
	if alreadyCopied {
		return
	}
	simplecontentFrom.GongCopyBasicFields(simplecontentTo)

	//insertion point for the staging of instances referenced by pointers
	if simplecontentFrom.Extension != nil {
		simplecontentTo.Extension = GongCopyBranchExtension(mapOrigCopy, simplecontentFrom.Extension)
	}
	if simplecontentFrom.Restriction != nil {
		simplecontentTo.Restriction = GongCopyBranchRestriction(mapOrigCopy, simplecontentFrom.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSimpleType(mapOrigCopy map[any]any, simpletypeFrom *SimpleType) (simpletypeTo *SimpleType) {
	var alreadyCopied bool
	simpletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, simpletypeFrom)
	if alreadyCopied {
		return
	}
	simpletypeFrom.GongCopyBasicFields(simpletypeTo)

	//insertion point for the staging of instances referenced by pointers
	if simpletypeFrom.Annotation != nil {
		simpletypeTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, simpletypeFrom.Annotation)
	}
	if simpletypeFrom.Restriction != nil {
		simpletypeTo.Restriction = GongCopyBranchRestriction(mapOrigCopy, simpletypeFrom.Restriction)
	}
	if simpletypeFrom.Union != nil {
		simpletypeTo.Union = GongCopyBranchUnion(mapOrigCopy, simpletypeFrom.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTotalDigit(mapOrigCopy map[any]any, totaldigitFrom *TotalDigit) (totaldigitTo *TotalDigit) {
	var alreadyCopied bool
	totaldigitTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, totaldigitFrom)
	if alreadyCopied {
		return
	}
	totaldigitFrom.GongCopyBasicFields(totaldigitTo)

	//insertion point for the staging of instances referenced by pointers
	if totaldigitFrom.Annotation != nil {
		totaldigitTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, totaldigitFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchUnion(mapOrigCopy map[any]any, unionFrom *Union) (unionTo *Union) {
	var alreadyCopied bool
	unionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, unionFrom)
	if alreadyCopied {
		return
	}
	unionFrom.GongCopyBasicFields(unionTo)

	//insertion point for the staging of instances referenced by pointers
	if unionFrom.Annotation != nil {
		unionTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, unionFrom.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchWhiteSpace(mapOrigCopy map[any]any, whitespaceFrom *WhiteSpace) (whitespaceTo *WhiteSpace) {
	var alreadyCopied bool
	whitespaceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, whitespaceFrom)
	if alreadyCopied {
		return
	}
	whitespaceFrom.GongCopyBasicFields(whitespaceTo)

	//insertion point for the staging of instances referenced by pointers
	if whitespaceFrom.Annotation != nil {
		whitespaceTo.Annotation = GongCopyBranchAnnotation(mapOrigCopy, whitespaceFrom.Annotation)
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

// insertion point for unstage branch per struct
func (all *All) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(all) {
		return
	}

	all.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if all.Annotation != nil {
		stage.UnstageBranch(all.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range all.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range all.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range all.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range all.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range all.Elements {
		stage.UnstageBranch(_element)
	}

}

func (annotation *Annotation) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(annotation) {
		return
	}

	annotation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _documentation := range annotation.Documentations {
		stage.UnstageBranch(_documentation)
	}

}

func (attribute *Attribute) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attribute) {
		return
	}

	attribute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute.Annotation != nil {
		stage.UnstageBranch(attribute.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attributegroup *AttributeGroup) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attributegroup) {
		return
	}

	attributegroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributegroup.Annotation != nil {
		stage.UnstageBranch(attributegroup.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributegroup := range attributegroup.AttributeGroups {
		stage.UnstageBranch(_attributegroup)
	}
	for _, _attribute := range attributegroup.Attributes {
		stage.UnstageBranch(_attribute)
	}

}

func (choice *Choice) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(choice) {
		return
	}

	choice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if choice.Annotation != nil {
		stage.UnstageBranch(choice.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range choice.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range choice.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range choice.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range choice.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range choice.Elements {
		stage.UnstageBranch(_element)
	}

}

func (complexcontent *ComplexContent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(complexcontent) {
		return
	}

	complexcontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (complextype *ComplexType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(complextype) {
		return
	}

	complextype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if complextype.OuterElement != nil {
		stage.UnstageBranch(complextype.OuterElement)
	}
	if complextype.Annotation != nil {
		stage.UnstageBranch(complextype.Annotation)
	}
	if complextype.Extension != nil {
		stage.UnstageBranch(complextype.Extension)
	}
	if complextype.SimpleContent != nil {
		stage.UnstageBranch(complextype.SimpleContent)
	}
	if complextype.ComplexContent != nil {
		stage.UnstageBranch(complextype.ComplexContent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range complextype.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range complextype.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range complextype.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range complextype.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range complextype.Elements {
		stage.UnstageBranch(_element)
	}
	for _, _attribute := range complextype.Attributes {
		stage.UnstageBranch(_attribute)
	}
	for _, _attributegroup := range complextype.AttributeGroups {
		stage.UnstageBranch(_attributegroup)
	}

}

func (documentation *Documentation) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(documentation) {
		return
	}

	documentation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (element *Element) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(element) {
		return
	}

	element.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if element.Annotation != nil {
		stage.UnstageBranch(element.Annotation)
	}
	if element.SimpleType != nil {
		stage.UnstageBranch(element.SimpleType)
	}
	if element.ComplexType != nil {
		stage.UnstageBranch(element.ComplexType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range element.Groups {
		stage.UnstageBranch(_group)
	}

}

func (enumeration *Enumeration) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(enumeration) {
		return
	}

	enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enumeration.Annotation != nil {
		stage.UnstageBranch(enumeration.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (extension *Extension) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(extension) {
		return
	}

	extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range extension.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range extension.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range extension.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range extension.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range extension.Elements {
		stage.UnstageBranch(_element)
	}
	for _, _attribute := range extension.Attributes {
		stage.UnstageBranch(_attribute)
	}
	for _, _attributegroup := range extension.AttributeGroups {
		stage.UnstageBranch(_attributegroup)
	}

}

func (group *Group) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if group.Annotation != nil {
		stage.UnstageBranch(group.Annotation)
	}
	if group.OuterElement != nil {
		stage.UnstageBranch(group.OuterElement)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range group.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range group.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range group.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range group.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range group.Elements {
		stage.UnstageBranch(_element)
	}

}

func (length *Length) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(length) {
		return
	}

	length.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if length.Annotation != nil {
		stage.UnstageBranch(length.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (maxinclusive *MaxInclusive) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(maxinclusive) {
		return
	}

	maxinclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxinclusive.Annotation != nil {
		stage.UnstageBranch(maxinclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (maxlength *MaxLength) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(maxlength) {
		return
	}

	maxlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if maxlength.Annotation != nil {
		stage.UnstageBranch(maxlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (mininclusive *MinInclusive) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(mininclusive) {
		return
	}

	mininclusive.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if mininclusive.Annotation != nil {
		stage.UnstageBranch(mininclusive.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (minlength *MinLength) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(minlength) {
		return
	}

	minlength.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if minlength.Annotation != nil {
		stage.UnstageBranch(minlength.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pattern *Pattern) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pattern) {
		return
	}

	pattern.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pattern.Annotation != nil {
		stage.UnstageBranch(pattern.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (restriction *Restriction) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(restriction) {
		return
	}

	restriction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if restriction.Annotation != nil {
		stage.UnstageBranch(restriction.Annotation)
	}
	if restriction.MinInclusive != nil {
		stage.UnstageBranch(restriction.MinInclusive)
	}
	if restriction.MaxInclusive != nil {
		stage.UnstageBranch(restriction.MaxInclusive)
	}
	if restriction.Pattern != nil {
		stage.UnstageBranch(restriction.Pattern)
	}
	if restriction.WhiteSpace != nil {
		stage.UnstageBranch(restriction.WhiteSpace)
	}
	if restriction.MinLength != nil {
		stage.UnstageBranch(restriction.MinLength)
	}
	if restriction.MaxLength != nil {
		stage.UnstageBranch(restriction.MaxLength)
	}
	if restriction.Length != nil {
		stage.UnstageBranch(restriction.Length)
	}
	if restriction.TotalDigit != nil {
		stage.UnstageBranch(restriction.TotalDigit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enumeration := range restriction.Enumerations {
		stage.UnstageBranch(_enumeration)
	}

}

func (schema *Schema) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(schema) {
		return
	}

	schema.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if schema.Annotation != nil {
		stage.UnstageBranch(schema.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _element := range schema.Elements {
		stage.UnstageBranch(_element)
	}
	for _, _simpletype := range schema.SimpleTypes {
		stage.UnstageBranch(_simpletype)
	}
	for _, _complextype := range schema.ComplexTypes {
		stage.UnstageBranch(_complextype)
	}
	for _, _attributegroup := range schema.AttributeGroups {
		stage.UnstageBranch(_attributegroup)
	}
	for _, _group := range schema.Groups {
		stage.UnstageBranch(_group)
	}

}

func (sequence *Sequence) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(sequence) {
		return
	}

	sequence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sequence.Annotation != nil {
		stage.UnstageBranch(sequence.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _sequence := range sequence.Sequences {
		stage.UnstageBranch(_sequence)
	}
	for _, _all := range sequence.Alls {
		stage.UnstageBranch(_all)
	}
	for _, _choice := range sequence.Choices {
		stage.UnstageBranch(_choice)
	}
	for _, _group := range sequence.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _element := range sequence.Elements {
		stage.UnstageBranch(_element)
	}

}

func (simplecontent *SimpleContent) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(simplecontent) {
		return
	}

	simplecontent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simplecontent.Extension != nil {
		stage.UnstageBranch(simplecontent.Extension)
	}
	if simplecontent.Restriction != nil {
		stage.UnstageBranch(simplecontent.Restriction)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (simpletype *SimpleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(simpletype) {
		return
	}

	simpletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if simpletype.Annotation != nil {
		stage.UnstageBranch(simpletype.Annotation)
	}
	if simpletype.Restriction != nil {
		stage.UnstageBranch(simpletype.Restriction)
	}
	if simpletype.Union != nil {
		stage.UnstageBranch(simpletype.Union)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (totaldigit *TotalDigit) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(totaldigit) {
		return
	}

	totaldigit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if totaldigit.Annotation != nil {
		stage.UnstageBranch(totaldigit.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (union *Union) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(union) {
		return
	}

	union.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if union.Annotation != nil {
		stage.UnstageBranch(union.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (whitespace *WhiteSpace) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(whitespace) {
		return
	}

	whitespace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if whitespace.Annotation != nil {
		stage.UnstageBranch(whitespace.Annotation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *All) GongReconstructPointersFromReferences(stage *Stage, instance *All) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
}

func (reference *Annotation) GongReconstructPointersFromReferences(stage *Stage, instance *Annotation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Documentations, stage.Documentations_reference, instance.Documentations)
}

func (reference *Attribute) GongReconstructPointersFromReferences(stage *Stage, instance *Attribute) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *AttributeGroup) GongReconstructPointersFromReferences(stage *Stage, instance *AttributeGroup) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.AttributeGroups, stage.AttributeGroups_reference, instance.AttributeGroups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Attributes, stage.Attributes_reference, instance.Attributes)
}

func (reference *Choice) GongReconstructPointersFromReferences(stage *Stage, instance *Choice) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
}

func (reference *ComplexContent) GongReconstructPointersFromReferences(stage *Stage, instance *ComplexContent) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ComplexType) GongReconstructPointersFromReferences(stage *Stage, instance *ComplexType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.OuterElement, stage.Elements_reference, instance.OuterElement)
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	__gong__reconstructPointer(&reference.Extension, stage.Extensions_reference, instance.Extension)
	__gong__reconstructPointer(&reference.SimpleContent, stage.SimpleContents_reference, instance.SimpleContent)
	__gong__reconstructPointer(&reference.ComplexContent, stage.ComplexContents_reference, instance.ComplexContent)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Attributes, stage.Attributes_reference, instance.Attributes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AttributeGroups, stage.AttributeGroups_reference, instance.AttributeGroups)
}

func (reference *Documentation) GongReconstructPointersFromReferences(stage *Stage, instance *Documentation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Element) GongReconstructPointersFromReferences(stage *Stage, instance *Element) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	__gong__reconstructPointer(&reference.SimpleType, stage.SimpleTypes_reference, instance.SimpleType)
	__gong__reconstructPointer(&reference.ComplexType, stage.ComplexTypes_reference, instance.ComplexType)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
}

func (reference *Enumeration) GongReconstructPointersFromReferences(stage *Stage, instance *Enumeration) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *Extension) GongReconstructPointersFromReferences(stage *Stage, instance *Extension) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Attributes, stage.Attributes_reference, instance.Attributes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AttributeGroups, stage.AttributeGroups_reference, instance.AttributeGroups)
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	__gong__reconstructPointer(&reference.OuterElement, stage.Elements_reference, instance.OuterElement)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
}

func (reference *Length) GongReconstructPointersFromReferences(stage *Stage, instance *Length) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *MaxInclusive) GongReconstructPointersFromReferences(stage *Stage, instance *MaxInclusive) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *MaxLength) GongReconstructPointersFromReferences(stage *Stage, instance *MaxLength) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *MinInclusive) GongReconstructPointersFromReferences(stage *Stage, instance *MinInclusive) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *MinLength) GongReconstructPointersFromReferences(stage *Stage, instance *MinLength) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *Pattern) GongReconstructPointersFromReferences(stage *Stage, instance *Pattern) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *Restriction) GongReconstructPointersFromReferences(stage *Stage, instance *Restriction) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	__gong__reconstructPointer(&reference.MinInclusive, stage.MinInclusives_reference, instance.MinInclusive)
	__gong__reconstructPointer(&reference.MaxInclusive, stage.MaxInclusives_reference, instance.MaxInclusive)
	__gong__reconstructPointer(&reference.Pattern, stage.Patterns_reference, instance.Pattern)
	__gong__reconstructPointer(&reference.WhiteSpace, stage.WhiteSpaces_reference, instance.WhiteSpace)
	__gong__reconstructPointer(&reference.MinLength, stage.MinLengths_reference, instance.MinLength)
	__gong__reconstructPointer(&reference.MaxLength, stage.MaxLengths_reference, instance.MaxLength)
	__gong__reconstructPointer(&reference.Length, stage.Lengths_reference, instance.Length)
	__gong__reconstructPointer(&reference.TotalDigit, stage.TotalDigits_reference, instance.TotalDigit)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Enumerations, stage.Enumerations_reference, instance.Enumerations)
}

func (reference *Schema) GongReconstructPointersFromReferences(stage *Stage, instance *Schema) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SimpleTypes, stage.SimpleTypes_reference, instance.SimpleTypes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComplexTypes, stage.ComplexTypes_reference, instance.ComplexTypes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AttributeGroups, stage.AttributeGroups_reference, instance.AttributeGroups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
}

func (reference *Sequence) GongReconstructPointersFromReferences(stage *Stage, instance *Sequence) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sequences, stage.Sequences_reference, instance.Sequences)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Alls, stage.Alls_reference, instance.Alls)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Choices, stage.Choices_reference, instance.Choices)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Elements, stage.Elements_reference, instance.Elements)
}

func (reference *SimpleContent) GongReconstructPointersFromReferences(stage *Stage, instance *SimpleContent) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Extension, stage.Extensions_reference, instance.Extension)
	__gong__reconstructPointer(&reference.Restriction, stage.Restrictions_reference, instance.Restriction)
	// insertion point for slice of pointers field
}

func (reference *SimpleType) GongReconstructPointersFromReferences(stage *Stage, instance *SimpleType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	__gong__reconstructPointer(&reference.Restriction, stage.Restrictions_reference, instance.Restriction)
	__gong__reconstructPointer(&reference.Union, stage.Unions_reference, instance.Union)
	// insertion point for slice of pointers field
}

func (reference *TotalDigit) GongReconstructPointersFromReferences(stage *Stage, instance *TotalDigit) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *Union) GongReconstructPointersFromReferences(stage *Stage, instance *Union) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

func (reference *WhiteSpace) GongReconstructPointersFromReferences(stage *Stage, instance *WhiteSpace) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Annotation, stage.Annotations_reference, instance.Annotation)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *All) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
}

func (reference *Annotation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Documentations, stage.Documentations_instance)
}

func (reference *Attribute) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *AttributeGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.AttributeGroups, stage.AttributeGroups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Attributes, stage.Attributes_instance)
}

func (reference *Choice) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
}

func (reference *ComplexContent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ComplexType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.OuterElement, stage.Elements_instance)
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	__gong__reconstructPointerFromInstance(&reference.Extension, stage.Extensions_instance)
	__gong__reconstructPointerFromInstance(&reference.SimpleContent, stage.SimpleContents_instance)
	__gong__reconstructPointerFromInstance(&reference.ComplexContent, stage.ComplexContents_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Attributes, stage.Attributes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AttributeGroups, stage.AttributeGroups_instance)
}

func (reference *Documentation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Element) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	__gong__reconstructPointerFromInstance(&reference.SimpleType, stage.SimpleTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ComplexType, stage.ComplexTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
}

func (reference *Enumeration) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *Extension) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Attributes, stage.Attributes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AttributeGroups, stage.AttributeGroups_instance)
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	__gong__reconstructPointerFromInstance(&reference.OuterElement, stage.Elements_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
}

func (reference *Length) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *MaxInclusive) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *MaxLength) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *MinInclusive) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *MinLength) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *Pattern) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *Restriction) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	__gong__reconstructPointerFromInstance(&reference.MinInclusive, stage.MinInclusives_instance)
	__gong__reconstructPointerFromInstance(&reference.MaxInclusive, stage.MaxInclusives_instance)
	__gong__reconstructPointerFromInstance(&reference.Pattern, stage.Patterns_instance)
	__gong__reconstructPointerFromInstance(&reference.WhiteSpace, stage.WhiteSpaces_instance)
	__gong__reconstructPointerFromInstance(&reference.MinLength, stage.MinLengths_instance)
	__gong__reconstructPointerFromInstance(&reference.MaxLength, stage.MaxLengths_instance)
	__gong__reconstructPointerFromInstance(&reference.Length, stage.Lengths_instance)
	__gong__reconstructPointerFromInstance(&reference.TotalDigit, stage.TotalDigits_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Enumerations, stage.Enumerations_instance)
}

func (reference *Schema) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SimpleTypes, stage.SimpleTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComplexTypes, stage.ComplexTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AttributeGroups, stage.AttributeGroups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
}

func (reference *Sequence) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sequences, stage.Sequences_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Alls, stage.Alls_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Choices, stage.Choices_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Elements, stage.Elements_instance)
}

func (reference *SimpleContent) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Extension, stage.Extensions_instance)
	__gong__reconstructPointerFromInstance(&reference.Restriction, stage.Restrictions_instance)
	// insertion point for slice of pointers fields
}

func (reference *SimpleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	__gong__reconstructPointerFromInstance(&reference.Restriction, stage.Restrictions_instance)
	__gong__reconstructPointerFromInstance(&reference.Union, stage.Unions_instance)
	// insertion point for slice of pointers fields
}

func (reference *TotalDigit) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *Union) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
	// insertion point for slice of pointers fields
}

func (reference *WhiteSpace) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Annotation, stage.Annotations_instance)
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
	if all.Annotation != allOther.Annotation {
		diffs = append(diffs, all.GongMarshallField(stage, "Annotation"))
	}
	if all.OuterElementName != allOther.OuterElementName {
		diffs = append(diffs, all.GongMarshallField(stage, "OuterElementName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, all, "Sequences", allOther.Sequences, all.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, all, "Alls", allOther.Alls, all.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, all, "Choices", allOther.Choices, all.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, all, "Groups", allOther.Groups, all.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, all, "Elements", allOther.Elements, all.Elements); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, annotation, "Documentations", annotationOther.Documentations, annotation.Documentations); ops != "" {
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
	if attribute.Annotation != attributeOther.Annotation {
		diffs = append(diffs, attribute.GongMarshallField(stage, "Annotation"))
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
	if attributegroup.Annotation != attributegroupOther.Annotation {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Annotation"))
	}
	if attributegroup.HasNameConflict != attributegroupOther.HasNameConflict {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "HasNameConflict"))
	}
	if attributegroup.GoIdentifier != attributegroupOther.GoIdentifier {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "GoIdentifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributegroup, "AttributeGroups", attributegroupOther.AttributeGroups, attributegroup.AttributeGroups); ops != "" {
		diffs = append(diffs, ops)
	}
	if attributegroup.Ref != attributegroupOther.Ref {
		diffs = append(diffs, attributegroup.GongMarshallField(stage, "Ref"))
	}
	if ops := __gong__diffSliceOfPointers(stage, attributegroup, "Attributes", attributegroupOther.Attributes, attributegroup.Attributes); ops != "" {
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
	if choice.Annotation != choiceOther.Annotation {
		diffs = append(diffs, choice.GongMarshallField(stage, "Annotation"))
	}
	if choice.OuterElementName != choiceOther.OuterElementName {
		diffs = append(diffs, choice.GongMarshallField(stage, "OuterElementName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, choice, "Sequences", choiceOther.Sequences, choice.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, choice, "Alls", choiceOther.Alls, choice.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, choice, "Choices", choiceOther.Choices, choice.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, choice, "Groups", choiceOther.Groups, choice.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, choice, "Elements", choiceOther.Elements, choice.Elements); ops != "" {
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
	if complextype.OuterElement != complextypeOther.OuterElement {
		diffs = append(diffs, complextype.GongMarshallField(stage, "OuterElement"))
	}
	if complextype.Annotation != complextypeOther.Annotation {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Annotation"))
	}
	if complextype.NameXSD != complextypeOther.NameXSD {
		diffs = append(diffs, complextype.GongMarshallField(stage, "NameXSD"))
	}
	if complextype.OuterElementName != complextypeOther.OuterElementName {
		diffs = append(diffs, complextype.GongMarshallField(stage, "OuterElementName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Sequences", complextypeOther.Sequences, complextype.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Alls", complextypeOther.Alls, complextype.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Choices", complextypeOther.Choices, complextype.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Groups", complextypeOther.Groups, complextype.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Elements", complextypeOther.Elements, complextype.Elements); ops != "" {
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
	if complextype.Extension != complextypeOther.Extension {
		diffs = append(diffs, complextype.GongMarshallField(stage, "Extension"))
	}
	if complextype.SimpleContent != complextypeOther.SimpleContent {
		diffs = append(diffs, complextype.GongMarshallField(stage, "SimpleContent"))
	}
	if complextype.ComplexContent != complextypeOther.ComplexContent {
		diffs = append(diffs, complextype.GongMarshallField(stage, "ComplexContent"))
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "Attributes", complextypeOther.Attributes, complextype.Attributes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, complextype, "AttributeGroups", complextypeOther.AttributeGroups, complextype.AttributeGroups); ops != "" {
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
	if element.Annotation != elementOther.Annotation {
		diffs = append(diffs, element.GongMarshallField(stage, "Annotation"))
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
	if element.SimpleType != elementOther.SimpleType {
		diffs = append(diffs, element.GongMarshallField(stage, "SimpleType"))
	}
	if element.ComplexType != elementOther.ComplexType {
		diffs = append(diffs, element.GongMarshallField(stage, "ComplexType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, element, "Groups", elementOther.Groups, element.Groups); ops != "" {
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
	if enumeration.Annotation != enumerationOther.Annotation {
		diffs = append(diffs, enumeration.GongMarshallField(stage, "Annotation"))
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
	if ops := __gong__diffSliceOfPointers(stage, extension, "Sequences", extensionOther.Sequences, extension.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extension, "Alls", extensionOther.Alls, extension.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extension, "Choices", extensionOther.Choices, extension.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extension, "Groups", extensionOther.Groups, extension.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extension, "Elements", extensionOther.Elements, extension.Elements); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, extension, "Attributes", extensionOther.Attributes, extension.Attributes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extension, "AttributeGroups", extensionOther.AttributeGroups, extension.AttributeGroups); ops != "" {
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
	if group.Annotation != groupOther.Annotation {
		diffs = append(diffs, group.GongMarshallField(stage, "Annotation"))
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
	if group.OuterElement != groupOther.OuterElement {
		diffs = append(diffs, group.GongMarshallField(stage, "OuterElement"))
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
	if ops := __gong__diffSliceOfPointers(stage, group, "Sequences", groupOther.Sequences, group.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, group, "Alls", groupOther.Alls, group.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, group, "Choices", groupOther.Choices, group.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, group, "Groups", groupOther.Groups, group.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, group, "Elements", groupOther.Elements, group.Elements); ops != "" {
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
	if length.Annotation != lengthOther.Annotation {
		diffs = append(diffs, length.GongMarshallField(stage, "Annotation"))
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
	if maxinclusive.Annotation != maxinclusiveOther.Annotation {
		diffs = append(diffs, maxinclusive.GongMarshallField(stage, "Annotation"))
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
	if maxlength.Annotation != maxlengthOther.Annotation {
		diffs = append(diffs, maxlength.GongMarshallField(stage, "Annotation"))
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
	if mininclusive.Annotation != mininclusiveOther.Annotation {
		diffs = append(diffs, mininclusive.GongMarshallField(stage, "Annotation"))
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
	if minlength.Annotation != minlengthOther.Annotation {
		diffs = append(diffs, minlength.GongMarshallField(stage, "Annotation"))
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
	if pattern.Annotation != patternOther.Annotation {
		diffs = append(diffs, pattern.GongMarshallField(stage, "Annotation"))
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
	if restriction.Annotation != restrictionOther.Annotation {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Annotation"))
	}
	if restriction.Base != restrictionOther.Base {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Base"))
	}
	if ops := __gong__diffSliceOfPointers(stage, restriction, "Enumerations", restrictionOther.Enumerations, restriction.Enumerations); ops != "" {
		diffs = append(diffs, ops)
	}
	if restriction.MinInclusive != restrictionOther.MinInclusive {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MinInclusive"))
	}
	if restriction.MaxInclusive != restrictionOther.MaxInclusive {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MaxInclusive"))
	}
	if restriction.Pattern != restrictionOther.Pattern {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Pattern"))
	}
	if restriction.WhiteSpace != restrictionOther.WhiteSpace {
		diffs = append(diffs, restriction.GongMarshallField(stage, "WhiteSpace"))
	}
	if restriction.MinLength != restrictionOther.MinLength {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MinLength"))
	}
	if restriction.MaxLength != restrictionOther.MaxLength {
		diffs = append(diffs, restriction.GongMarshallField(stage, "MaxLength"))
	}
	if restriction.Length != restrictionOther.Length {
		diffs = append(diffs, restriction.GongMarshallField(stage, "Length"))
	}
	if restriction.TotalDigit != restrictionOther.TotalDigit {
		diffs = append(diffs, restriction.GongMarshallField(stage, "TotalDigit"))
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
	if schema.Annotation != schemaOther.Annotation {
		diffs = append(diffs, schema.GongMarshallField(stage, "Annotation"))
	}
	if ops := __gong__diffSliceOfPointers(stage, schema, "Elements", schemaOther.Elements, schema.Elements); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, schema, "SimpleTypes", schemaOther.SimpleTypes, schema.SimpleTypes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, schema, "ComplexTypes", schemaOther.ComplexTypes, schema.ComplexTypes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, schema, "AttributeGroups", schemaOther.AttributeGroups, schema.AttributeGroups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, schema, "Groups", schemaOther.Groups, schema.Groups); ops != "" {
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
	if sequence.Annotation != sequenceOther.Annotation {
		diffs = append(diffs, sequence.GongMarshallField(stage, "Annotation"))
	}
	if sequence.OuterElementName != sequenceOther.OuterElementName {
		diffs = append(diffs, sequence.GongMarshallField(stage, "OuterElementName"))
	}
	if ops := __gong__diffSliceOfPointers(stage, sequence, "Sequences", sequenceOther.Sequences, sequence.Sequences); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sequence, "Alls", sequenceOther.Alls, sequence.Alls); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sequence, "Choices", sequenceOther.Choices, sequence.Choices); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sequence, "Groups", sequenceOther.Groups, sequence.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, sequence, "Elements", sequenceOther.Elements, sequence.Elements); ops != "" {
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
	if simplecontent.Extension != simplecontentOther.Extension {
		diffs = append(diffs, simplecontent.GongMarshallField(stage, "Extension"))
	}
	if simplecontent.Restriction != simplecontentOther.Restriction {
		diffs = append(diffs, simplecontent.GongMarshallField(stage, "Restriction"))
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
	if simpletype.Annotation != simpletypeOther.Annotation {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Annotation"))
	}
	if simpletype.NameXSD != simpletypeOther.NameXSD {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "NameXSD"))
	}
	if simpletype.Restriction != simpletypeOther.Restriction {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Restriction"))
	}
	if simpletype.Union != simpletypeOther.Union {
		diffs = append(diffs, simpletype.GongMarshallField(stage, "Union"))
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
	if totaldigit.Annotation != totaldigitOther.Annotation {
		diffs = append(diffs, totaldigit.GongMarshallField(stage, "Annotation"))
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
	if union.Annotation != unionOther.Annotation {
		diffs = append(diffs, union.GongMarshallField(stage, "Annotation"))
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
	if whitespace.Annotation != whitespaceOther.Annotation {
		diffs = append(diffs, whitespace.GongMarshallField(stage, "Annotation"))
	}
	if whitespace.Value != whitespaceOther.Value {
		diffs = append(diffs, whitespace.GongMarshallField(stage, "Value"))
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

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
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
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
