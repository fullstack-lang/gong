// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by All
func (all *All) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&all.Sequences) || modified
	modified = stage.CleanSlice(&all.Alls) || modified
	modified = stage.CleanSlice(&all.Choices) || modified
	modified = stage.CleanSlice(&all.Groups) || modified
	modified = stage.CleanSlice(&all.Elements) || modified
	// insertion point per field
	modified = stage.CleanPointer(&all.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Annotation
func (annotation *Annotation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&annotation.Documentations) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Attribute
func (attribute *Attribute) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&attribute.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AttributeGroup
func (attributegroup *AttributeGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&attributegroup.AttributeGroups) || modified
	modified = stage.CleanSlice(&attributegroup.Attributes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&attributegroup.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Choice
func (choice *Choice) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&choice.Sequences) || modified
	modified = stage.CleanSlice(&choice.Alls) || modified
	modified = stage.CleanSlice(&choice.Choices) || modified
	modified = stage.CleanSlice(&choice.Groups) || modified
	modified = stage.CleanSlice(&choice.Elements) || modified
	// insertion point per field
	modified = stage.CleanPointer(&choice.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ComplexType
func (complextype *ComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&complextype.Sequences) || modified
	modified = stage.CleanSlice(&complextype.Alls) || modified
	modified = stage.CleanSlice(&complextype.Choices) || modified
	modified = stage.CleanSlice(&complextype.Groups) || modified
	modified = stage.CleanSlice(&complextype.Elements) || modified
	modified = stage.CleanSlice(&complextype.Attributes) || modified
	modified = stage.CleanSlice(&complextype.AttributeGroups) || modified
	// insertion point per field
	modified = stage.CleanPointer(&complextype.OuterElement) || modified
	modified = stage.CleanPointer(&complextype.Annotation) || modified
	modified = stage.CleanPointer(&complextype.Extension) || modified
	modified = stage.CleanPointer(&complextype.SimpleContent) || modified
	modified = stage.CleanPointer(&complextype.ComplexContent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Element
func (element *Element) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&element.Groups) || modified
	// insertion point per field
	modified = stage.CleanPointer(&element.Annotation) || modified
	modified = stage.CleanPointer(&element.SimpleType) || modified
	modified = stage.CleanPointer(&element.ComplexType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Enumeration
func (enumeration *Enumeration) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&enumeration.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Extension
func (extension *Extension) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&extension.Sequences) || modified
	modified = stage.CleanSlice(&extension.Alls) || modified
	modified = stage.CleanSlice(&extension.Choices) || modified
	modified = stage.CleanSlice(&extension.Groups) || modified
	modified = stage.CleanSlice(&extension.Elements) || modified
	modified = stage.CleanSlice(&extension.Attributes) || modified
	modified = stage.CleanSlice(&extension.AttributeGroups) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&group.Sequences) || modified
	modified = stage.CleanSlice(&group.Alls) || modified
	modified = stage.CleanSlice(&group.Choices) || modified
	modified = stage.CleanSlice(&group.Groups) || modified
	modified = stage.CleanSlice(&group.Elements) || modified
	// insertion point per field
	modified = stage.CleanPointer(&group.Annotation) || modified
	modified = stage.CleanPointer(&group.OuterElement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Length
func (length *Length) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&length.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MaxInclusive
func (maxinclusive *MaxInclusive) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&maxinclusive.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MaxLength
func (maxlength *MaxLength) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&maxlength.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MinInclusive
func (mininclusive *MinInclusive) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&mininclusive.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MinLength
func (minlength *MinLength) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&minlength.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Pattern
func (pattern *Pattern) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pattern.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Restriction
func (restriction *Restriction) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&restriction.Enumerations) || modified
	// insertion point per field
	modified = stage.CleanPointer(&restriction.Annotation) || modified
	modified = stage.CleanPointer(&restriction.MinInclusive) || modified
	modified = stage.CleanPointer(&restriction.MaxInclusive) || modified
	modified = stage.CleanPointer(&restriction.Pattern) || modified
	modified = stage.CleanPointer(&restriction.WhiteSpace) || modified
	modified = stage.CleanPointer(&restriction.MinLength) || modified
	modified = stage.CleanPointer(&restriction.MaxLength) || modified
	modified = stage.CleanPointer(&restriction.Length) || modified
	modified = stage.CleanPointer(&restriction.TotalDigit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Schema
func (schema *Schema) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&schema.Elements) || modified
	modified = stage.CleanSlice(&schema.SimpleTypes) || modified
	modified = stage.CleanSlice(&schema.ComplexTypes) || modified
	modified = stage.CleanSlice(&schema.AttributeGroups) || modified
	modified = stage.CleanSlice(&schema.Groups) || modified
	// insertion point per field
	modified = stage.CleanPointer(&schema.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Sequence
func (sequence *Sequence) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sequence.Sequences) || modified
	modified = stage.CleanSlice(&sequence.Alls) || modified
	modified = stage.CleanSlice(&sequence.Choices) || modified
	modified = stage.CleanSlice(&sequence.Groups) || modified
	modified = stage.CleanSlice(&sequence.Elements) || modified
	// insertion point per field
	modified = stage.CleanPointer(&sequence.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SimpleContent
func (simplecontent *SimpleContent) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&simplecontent.Extension) || modified
	modified = stage.CleanPointer(&simplecontent.Restriction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SimpleType
func (simpletype *SimpleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&simpletype.Annotation) || modified
	modified = stage.CleanPointer(&simpletype.Restriction) || modified
	modified = stage.CleanPointer(&simpletype.Union) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TotalDigit
func (totaldigit *TotalDigit) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&totaldigit.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Union
func (union *Union) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&union.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by WhiteSpace
func (whitespace *WhiteSpace) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&whitespace.Annotation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
