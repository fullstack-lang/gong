// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by All
func (all *All) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &all.Sequences) || modified
	modified = GongCleanSlice(stage, &all.Alls) || modified
	modified = GongCleanSlice(stage, &all.Choices) || modified
	modified = GongCleanSlice(stage, &all.Groups) || modified
	modified = GongCleanSlice(stage, &all.Elements) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &all.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Annotation
func (annotation *Annotation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &annotation.Documentations) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Attribute
func (attribute *Attribute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AttributeGroup
func (attributegroup *AttributeGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &attributegroup.AttributeGroups) || modified
	modified = GongCleanSlice(stage, &attributegroup.Attributes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &attributegroup.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Choice
func (choice *Choice) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &choice.Sequences) || modified
	modified = GongCleanSlice(stage, &choice.Alls) || modified
	modified = GongCleanSlice(stage, &choice.Choices) || modified
	modified = GongCleanSlice(stage, &choice.Groups) || modified
	modified = GongCleanSlice(stage, &choice.Elements) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &choice.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ComplexContent
func (complexcontent *ComplexContent) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ComplexType
func (complextype *ComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &complextype.Sequences) || modified
	modified = GongCleanSlice(stage, &complextype.Alls) || modified
	modified = GongCleanSlice(stage, &complextype.Choices) || modified
	modified = GongCleanSlice(stage, &complextype.Groups) || modified
	modified = GongCleanSlice(stage, &complextype.Elements) || modified
	modified = GongCleanSlice(stage, &complextype.Attributes) || modified
	modified = GongCleanSlice(stage, &complextype.AttributeGroups) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &complextype.OuterElement) || modified
	modified = GongCleanPointer(stage, &complextype.Annotation) || modified
	modified = GongCleanPointer(stage, &complextype.Extension) || modified
	modified = GongCleanPointer(stage, &complextype.SimpleContent) || modified
	modified = GongCleanPointer(stage, &complextype.ComplexContent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Documentation
func (documentation *Documentation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Element
func (element *Element) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &element.Groups) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &element.Annotation) || modified
	modified = GongCleanPointer(stage, &element.SimpleType) || modified
	modified = GongCleanPointer(stage, &element.ComplexType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Enumeration
func (enumeration *Enumeration) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &enumeration.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Extension
func (extension *Extension) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &extension.Sequences) || modified
	modified = GongCleanSlice(stage, &extension.Alls) || modified
	modified = GongCleanSlice(stage, &extension.Choices) || modified
	modified = GongCleanSlice(stage, &extension.Groups) || modified
	modified = GongCleanSlice(stage, &extension.Elements) || modified
	modified = GongCleanSlice(stage, &extension.Attributes) || modified
	modified = GongCleanSlice(stage, &extension.AttributeGroups) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &group.Sequences) || modified
	modified = GongCleanSlice(stage, &group.Alls) || modified
	modified = GongCleanSlice(stage, &group.Choices) || modified
	modified = GongCleanSlice(stage, &group.Groups) || modified
	modified = GongCleanSlice(stage, &group.Elements) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &group.Annotation) || modified
	modified = GongCleanPointer(stage, &group.OuterElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Length
func (length *Length) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &length.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MaxInclusive
func (maxinclusive *MaxInclusive) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &maxinclusive.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MaxLength
func (maxlength *MaxLength) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &maxlength.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MinInclusive
func (mininclusive *MinInclusive) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &mininclusive.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MinLength
func (minlength *MinLength) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &minlength.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Pattern
func (pattern *Pattern) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &pattern.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Restriction
func (restriction *Restriction) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &restriction.Enumerations) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &restriction.Annotation) || modified
	modified = GongCleanPointer(stage, &restriction.MinInclusive) || modified
	modified = GongCleanPointer(stage, &restriction.MaxInclusive) || modified
	modified = GongCleanPointer(stage, &restriction.Pattern) || modified
	modified = GongCleanPointer(stage, &restriction.WhiteSpace) || modified
	modified = GongCleanPointer(stage, &restriction.MinLength) || modified
	modified = GongCleanPointer(stage, &restriction.MaxLength) || modified
	modified = GongCleanPointer(stage, &restriction.Length) || modified
	modified = GongCleanPointer(stage, &restriction.TotalDigit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Schema
func (schema *Schema) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &schema.Elements) || modified
	modified = GongCleanSlice(stage, &schema.SimpleTypes) || modified
	modified = GongCleanSlice(stage, &schema.ComplexTypes) || modified
	modified = GongCleanSlice(stage, &schema.AttributeGroups) || modified
	modified = GongCleanSlice(stage, &schema.Groups) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &schema.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Sequence
func (sequence *Sequence) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &sequence.Sequences) || modified
	modified = GongCleanSlice(stage, &sequence.Alls) || modified
	modified = GongCleanSlice(stage, &sequence.Choices) || modified
	modified = GongCleanSlice(stage, &sequence.Groups) || modified
	modified = GongCleanSlice(stage, &sequence.Elements) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &sequence.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SimpleContent
func (simplecontent *SimpleContent) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &simplecontent.Extension) || modified
	modified = GongCleanPointer(stage, &simplecontent.Restriction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SimpleType
func (simpletype *SimpleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &simpletype.Annotation) || modified
	modified = GongCleanPointer(stage, &simpletype.Restriction) || modified
	modified = GongCleanPointer(stage, &simpletype.Union) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TotalDigit
func (totaldigit *TotalDigit) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &totaldigit.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Union
func (union *Union) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &union.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by WhiteSpace
func (whitespace *WhiteSpace) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &whitespace.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
