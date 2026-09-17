// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (all *All) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAllCreateCallback != nil {
		stage.OnAfterAllCreateCallback.OnAfterCreate(stage, all)
	}
}

func (all *All) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllUpdateCallback != nil {
		var frontAll *All
		if front != nil {
			frontAll, _ = front.(*All)
		}
		stage.OnAfterAllUpdateCallback.OnAfterUpdate(stage, all, frontAll)
	}
}

func (all *All) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllDeleteCallback != nil {
		var frontAll *All
		if front != nil {
			frontAll, _ = front.(*All)
		}
		stage.OnAfterAllDeleteCallback.OnAfterDelete(stage, all, frontAll)
	}
}

func (annotation *Annotation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAnnotationCreateCallback != nil {
		stage.OnAfterAnnotationCreateCallback.OnAfterCreate(stage, annotation)
	}
}

func (annotation *Annotation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnnotationUpdateCallback != nil {
		var frontAnnotation *Annotation
		if front != nil {
			frontAnnotation, _ = front.(*Annotation)
		}
		stage.OnAfterAnnotationUpdateCallback.OnAfterUpdate(stage, annotation, frontAnnotation)
	}
}

func (annotation *Annotation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnnotationDeleteCallback != nil {
		var frontAnnotation *Annotation
		if front != nil {
			frontAnnotation, _ = front.(*Annotation)
		}
		stage.OnAfterAnnotationDeleteCallback.OnAfterDelete(stage, annotation, frontAnnotation)
	}
}

func (attribute *Attribute) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAttributeCreateCallback != nil {
		stage.OnAfterAttributeCreateCallback.OnAfterCreate(stage, attribute)
	}
}

func (attribute *Attribute) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeUpdateCallback != nil {
		var frontAttribute *Attribute
		if front != nil {
			frontAttribute, _ = front.(*Attribute)
		}
		stage.OnAfterAttributeUpdateCallback.OnAfterUpdate(stage, attribute, frontAttribute)
	}
}

func (attribute *Attribute) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeDeleteCallback != nil {
		var frontAttribute *Attribute
		if front != nil {
			frontAttribute, _ = front.(*Attribute)
		}
		stage.OnAfterAttributeDeleteCallback.OnAfterDelete(stage, attribute, frontAttribute)
	}
}

func (attributegroup *AttributeGroup) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAttributeGroupCreateCallback != nil {
		stage.OnAfterAttributeGroupCreateCallback.OnAfterCreate(stage, attributegroup)
	}
}

func (attributegroup *AttributeGroup) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeGroupUpdateCallback != nil {
		var frontAttributeGroup *AttributeGroup
		if front != nil {
			frontAttributeGroup, _ = front.(*AttributeGroup)
		}
		stage.OnAfterAttributeGroupUpdateCallback.OnAfterUpdate(stage, attributegroup, frontAttributeGroup)
	}
}

func (attributegroup *AttributeGroup) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeGroupDeleteCallback != nil {
		var frontAttributeGroup *AttributeGroup
		if front != nil {
			frontAttributeGroup, _ = front.(*AttributeGroup)
		}
		stage.OnAfterAttributeGroupDeleteCallback.OnAfterDelete(stage, attributegroup, frontAttributeGroup)
	}
}

func (choice *Choice) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChoiceCreateCallback != nil {
		stage.OnAfterChoiceCreateCallback.OnAfterCreate(stage, choice)
	}
}

func (choice *Choice) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChoiceUpdateCallback != nil {
		var frontChoice *Choice
		if front != nil {
			frontChoice, _ = front.(*Choice)
		}
		stage.OnAfterChoiceUpdateCallback.OnAfterUpdate(stage, choice, frontChoice)
	}
}

func (choice *Choice) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChoiceDeleteCallback != nil {
		var frontChoice *Choice
		if front != nil {
			frontChoice, _ = front.(*Choice)
		}
		stage.OnAfterChoiceDeleteCallback.OnAfterDelete(stage, choice, frontChoice)
	}
}

func (complexcontent *ComplexContent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterComplexContentCreateCallback != nil {
		stage.OnAfterComplexContentCreateCallback.OnAfterCreate(stage, complexcontent)
	}
}

func (complexcontent *ComplexContent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexContentUpdateCallback != nil {
		var frontComplexContent *ComplexContent
		if front != nil {
			frontComplexContent, _ = front.(*ComplexContent)
		}
		stage.OnAfterComplexContentUpdateCallback.OnAfterUpdate(stage, complexcontent, frontComplexContent)
	}
}

func (complexcontent *ComplexContent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexContentDeleteCallback != nil {
		var frontComplexContent *ComplexContent
		if front != nil {
			frontComplexContent, _ = front.(*ComplexContent)
		}
		stage.OnAfterComplexContentDeleteCallback.OnAfterDelete(stage, complexcontent, frontComplexContent)
	}
}

func (complextype *ComplexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterComplexTypeCreateCallback != nil {
		stage.OnAfterComplexTypeCreateCallback.OnAfterCreate(stage, complextype)
	}
}

func (complextype *ComplexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexTypeUpdateCallback != nil {
		var frontComplexType *ComplexType
		if front != nil {
			frontComplexType, _ = front.(*ComplexType)
		}
		stage.OnAfterComplexTypeUpdateCallback.OnAfterUpdate(stage, complextype, frontComplexType)
	}
}

func (complextype *ComplexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexTypeDeleteCallback != nil {
		var frontComplexType *ComplexType
		if front != nil {
			frontComplexType, _ = front.(*ComplexType)
		}
		stage.OnAfterComplexTypeDeleteCallback.OnAfterDelete(stage, complextype, frontComplexType)
	}
}

func (documentation *Documentation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDocumentationCreateCallback != nil {
		stage.OnAfterDocumentationCreateCallback.OnAfterCreate(stage, documentation)
	}
}

func (documentation *Documentation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentationUpdateCallback != nil {
		var frontDocumentation *Documentation
		if front != nil {
			frontDocumentation, _ = front.(*Documentation)
		}
		stage.OnAfterDocumentationUpdateCallback.OnAfterUpdate(stage, documentation, frontDocumentation)
	}
}

func (documentation *Documentation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDocumentationDeleteCallback != nil {
		var frontDocumentation *Documentation
		if front != nil {
			frontDocumentation, _ = front.(*Documentation)
		}
		stage.OnAfterDocumentationDeleteCallback.OnAfterDelete(stage, documentation, frontDocumentation)
	}
}

func (element *Element) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElementCreateCallback != nil {
		stage.OnAfterElementCreateCallback.OnAfterCreate(stage, element)
	}
}

func (element *Element) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElementUpdateCallback != nil {
		var frontElement *Element
		if front != nil {
			frontElement, _ = front.(*Element)
		}
		stage.OnAfterElementUpdateCallback.OnAfterUpdate(stage, element, frontElement)
	}
}

func (element *Element) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElementDeleteCallback != nil {
		var frontElement *Element
		if front != nil {
			frontElement, _ = front.(*Element)
		}
		stage.OnAfterElementDeleteCallback.OnAfterDelete(stage, element, frontElement)
	}
}

func (enumeration *Enumeration) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnumerationCreateCallback != nil {
		stage.OnAfterEnumerationCreateCallback.OnAfterCreate(stage, enumeration)
	}
}

func (enumeration *Enumeration) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnumerationUpdateCallback != nil {
		var frontEnumeration *Enumeration
		if front != nil {
			frontEnumeration, _ = front.(*Enumeration)
		}
		stage.OnAfterEnumerationUpdateCallback.OnAfterUpdate(stage, enumeration, frontEnumeration)
	}
}

func (enumeration *Enumeration) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnumerationDeleteCallback != nil {
		var frontEnumeration *Enumeration
		if front != nil {
			frontEnumeration, _ = front.(*Enumeration)
		}
		stage.OnAfterEnumerationDeleteCallback.OnAfterDelete(stage, enumeration, frontEnumeration)
	}
}

func (extension *Extension) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExtensionCreateCallback != nil {
		stage.OnAfterExtensionCreateCallback.OnAfterCreate(stage, extension)
	}
}

func (extension *Extension) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtensionUpdateCallback != nil {
		var frontExtension *Extension
		if front != nil {
			frontExtension, _ = front.(*Extension)
		}
		stage.OnAfterExtensionUpdateCallback.OnAfterUpdate(stage, extension, frontExtension)
	}
}

func (extension *Extension) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtensionDeleteCallback != nil {
		var frontExtension *Extension
		if front != nil {
			frontExtension, _ = front.(*Extension)
		}
		stage.OnAfterExtensionDeleteCallback.OnAfterDelete(stage, extension, frontExtension)
	}
}

func (group *Group) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupCreateCallback != nil {
		stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, group)
	}
}

func (group *Group) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUpdateCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, group, frontGroup)
	}
}

func (group *Group) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupDeleteCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, group, frontGroup)
	}
}

func (length *Length) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLengthCreateCallback != nil {
		stage.OnAfterLengthCreateCallback.OnAfterCreate(stage, length)
	}
}

func (length *Length) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthUpdateCallback != nil {
		var frontLength *Length
		if front != nil {
			frontLength, _ = front.(*Length)
		}
		stage.OnAfterLengthUpdateCallback.OnAfterUpdate(stage, length, frontLength)
	}
}

func (length *Length) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthDeleteCallback != nil {
		var frontLength *Length
		if front != nil {
			frontLength, _ = front.(*Length)
		}
		stage.OnAfterLengthDeleteCallback.OnAfterDelete(stage, length, frontLength)
	}
}

func (maxinclusive *MaxInclusive) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMaxInclusiveCreateCallback != nil {
		stage.OnAfterMaxInclusiveCreateCallback.OnAfterCreate(stage, maxinclusive)
	}
}

func (maxinclusive *MaxInclusive) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMaxInclusiveUpdateCallback != nil {
		var frontMaxInclusive *MaxInclusive
		if front != nil {
			frontMaxInclusive, _ = front.(*MaxInclusive)
		}
		stage.OnAfterMaxInclusiveUpdateCallback.OnAfterUpdate(stage, maxinclusive, frontMaxInclusive)
	}
}

func (maxinclusive *MaxInclusive) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMaxInclusiveDeleteCallback != nil {
		var frontMaxInclusive *MaxInclusive
		if front != nil {
			frontMaxInclusive, _ = front.(*MaxInclusive)
		}
		stage.OnAfterMaxInclusiveDeleteCallback.OnAfterDelete(stage, maxinclusive, frontMaxInclusive)
	}
}

func (maxlength *MaxLength) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMaxLengthCreateCallback != nil {
		stage.OnAfterMaxLengthCreateCallback.OnAfterCreate(stage, maxlength)
	}
}

func (maxlength *MaxLength) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMaxLengthUpdateCallback != nil {
		var frontMaxLength *MaxLength
		if front != nil {
			frontMaxLength, _ = front.(*MaxLength)
		}
		stage.OnAfterMaxLengthUpdateCallback.OnAfterUpdate(stage, maxlength, frontMaxLength)
	}
}

func (maxlength *MaxLength) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMaxLengthDeleteCallback != nil {
		var frontMaxLength *MaxLength
		if front != nil {
			frontMaxLength, _ = front.(*MaxLength)
		}
		stage.OnAfterMaxLengthDeleteCallback.OnAfterDelete(stage, maxlength, frontMaxLength)
	}
}

func (mininclusive *MinInclusive) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMinInclusiveCreateCallback != nil {
		stage.OnAfterMinInclusiveCreateCallback.OnAfterCreate(stage, mininclusive)
	}
}

func (mininclusive *MinInclusive) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinInclusiveUpdateCallback != nil {
		var frontMinInclusive *MinInclusive
		if front != nil {
			frontMinInclusive, _ = front.(*MinInclusive)
		}
		stage.OnAfterMinInclusiveUpdateCallback.OnAfterUpdate(stage, mininclusive, frontMinInclusive)
	}
}

func (mininclusive *MinInclusive) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinInclusiveDeleteCallback != nil {
		var frontMinInclusive *MinInclusive
		if front != nil {
			frontMinInclusive, _ = front.(*MinInclusive)
		}
		stage.OnAfterMinInclusiveDeleteCallback.OnAfterDelete(stage, mininclusive, frontMinInclusive)
	}
}

func (minlength *MinLength) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMinLengthCreateCallback != nil {
		stage.OnAfterMinLengthCreateCallback.OnAfterCreate(stage, minlength)
	}
}

func (minlength *MinLength) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinLengthUpdateCallback != nil {
		var frontMinLength *MinLength
		if front != nil {
			frontMinLength, _ = front.(*MinLength)
		}
		stage.OnAfterMinLengthUpdateCallback.OnAfterUpdate(stage, minlength, frontMinLength)
	}
}

func (minlength *MinLength) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMinLengthDeleteCallback != nil {
		var frontMinLength *MinLength
		if front != nil {
			frontMinLength, _ = front.(*MinLength)
		}
		stage.OnAfterMinLengthDeleteCallback.OnAfterDelete(stage, minlength, frontMinLength)
	}
}

func (pattern *Pattern) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPatternCreateCallback != nil {
		stage.OnAfterPatternCreateCallback.OnAfterCreate(stage, pattern)
	}
}

func (pattern *Pattern) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPatternUpdateCallback != nil {
		var frontPattern *Pattern
		if front != nil {
			frontPattern, _ = front.(*Pattern)
		}
		stage.OnAfterPatternUpdateCallback.OnAfterUpdate(stage, pattern, frontPattern)
	}
}

func (pattern *Pattern) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPatternDeleteCallback != nil {
		var frontPattern *Pattern
		if front != nil {
			frontPattern, _ = front.(*Pattern)
		}
		stage.OnAfterPatternDeleteCallback.OnAfterDelete(stage, pattern, frontPattern)
	}
}

func (restriction *Restriction) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRestrictionCreateCallback != nil {
		stage.OnAfterRestrictionCreateCallback.OnAfterCreate(stage, restriction)
	}
}

func (restriction *Restriction) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRestrictionUpdateCallback != nil {
		var frontRestriction *Restriction
		if front != nil {
			frontRestriction, _ = front.(*Restriction)
		}
		stage.OnAfterRestrictionUpdateCallback.OnAfterUpdate(stage, restriction, frontRestriction)
	}
}

func (restriction *Restriction) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRestrictionDeleteCallback != nil {
		var frontRestriction *Restriction
		if front != nil {
			frontRestriction, _ = front.(*Restriction)
		}
		stage.OnAfterRestrictionDeleteCallback.OnAfterDelete(stage, restriction, frontRestriction)
	}
}

func (schema *Schema) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSchemaCreateCallback != nil {
		stage.OnAfterSchemaCreateCallback.OnAfterCreate(stage, schema)
	}
}

func (schema *Schema) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSchemaUpdateCallback != nil {
		var frontSchema *Schema
		if front != nil {
			frontSchema, _ = front.(*Schema)
		}
		stage.OnAfterSchemaUpdateCallback.OnAfterUpdate(stage, schema, frontSchema)
	}
}

func (schema *Schema) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSchemaDeleteCallback != nil {
		var frontSchema *Schema
		if front != nil {
			frontSchema, _ = front.(*Schema)
		}
		stage.OnAfterSchemaDeleteCallback.OnAfterDelete(stage, schema, frontSchema)
	}
}

func (sequence *Sequence) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSequenceCreateCallback != nil {
		stage.OnAfterSequenceCreateCallback.OnAfterCreate(stage, sequence)
	}
}

func (sequence *Sequence) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSequenceUpdateCallback != nil {
		var frontSequence *Sequence
		if front != nil {
			frontSequence, _ = front.(*Sequence)
		}
		stage.OnAfterSequenceUpdateCallback.OnAfterUpdate(stage, sequence, frontSequence)
	}
}

func (sequence *Sequence) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSequenceDeleteCallback != nil {
		var frontSequence *Sequence
		if front != nil {
			frontSequence, _ = front.(*Sequence)
		}
		stage.OnAfterSequenceDeleteCallback.OnAfterDelete(stage, sequence, frontSequence)
	}
}

func (simplecontent *SimpleContent) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSimpleContentCreateCallback != nil {
		stage.OnAfterSimpleContentCreateCallback.OnAfterCreate(stage, simplecontent)
	}
}

func (simplecontent *SimpleContent) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleContentUpdateCallback != nil {
		var frontSimpleContent *SimpleContent
		if front != nil {
			frontSimpleContent, _ = front.(*SimpleContent)
		}
		stage.OnAfterSimpleContentUpdateCallback.OnAfterUpdate(stage, simplecontent, frontSimpleContent)
	}
}

func (simplecontent *SimpleContent) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleContentDeleteCallback != nil {
		var frontSimpleContent *SimpleContent
		if front != nil {
			frontSimpleContent, _ = front.(*SimpleContent)
		}
		stage.OnAfterSimpleContentDeleteCallback.OnAfterDelete(stage, simplecontent, frontSimpleContent)
	}
}

func (simpletype *SimpleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSimpleTypeCreateCallback != nil {
		stage.OnAfterSimpleTypeCreateCallback.OnAfterCreate(stage, simpletype)
	}
}

func (simpletype *SimpleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleTypeUpdateCallback != nil {
		var frontSimpleType *SimpleType
		if front != nil {
			frontSimpleType, _ = front.(*SimpleType)
		}
		stage.OnAfterSimpleTypeUpdateCallback.OnAfterUpdate(stage, simpletype, frontSimpleType)
	}
}

func (simpletype *SimpleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleTypeDeleteCallback != nil {
		var frontSimpleType *SimpleType
		if front != nil {
			frontSimpleType, _ = front.(*SimpleType)
		}
		stage.OnAfterSimpleTypeDeleteCallback.OnAfterDelete(stage, simpletype, frontSimpleType)
	}
}

func (totaldigit *TotalDigit) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTotalDigitCreateCallback != nil {
		stage.OnAfterTotalDigitCreateCallback.OnAfterCreate(stage, totaldigit)
	}
}

func (totaldigit *TotalDigit) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTotalDigitUpdateCallback != nil {
		var frontTotalDigit *TotalDigit
		if front != nil {
			frontTotalDigit, _ = front.(*TotalDigit)
		}
		stage.OnAfterTotalDigitUpdateCallback.OnAfterUpdate(stage, totaldigit, frontTotalDigit)
	}
}

func (totaldigit *TotalDigit) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTotalDigitDeleteCallback != nil {
		var frontTotalDigit *TotalDigit
		if front != nil {
			frontTotalDigit, _ = front.(*TotalDigit)
		}
		stage.OnAfterTotalDigitDeleteCallback.OnAfterDelete(stage, totaldigit, frontTotalDigit)
	}
}

func (union *Union) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnionCreateCallback != nil {
		stage.OnAfterUnionCreateCallback.OnAfterCreate(stage, union)
	}
}

func (union *Union) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnionUpdateCallback != nil {
		var frontUnion *Union
		if front != nil {
			frontUnion, _ = front.(*Union)
		}
		stage.OnAfterUnionUpdateCallback.OnAfterUpdate(stage, union, frontUnion)
	}
}

func (union *Union) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnionDeleteCallback != nil {
		var frontUnion *Union
		if front != nil {
			frontUnion, _ = front.(*Union)
		}
		stage.OnAfterUnionDeleteCallback.OnAfterDelete(stage, union, frontUnion)
	}
}

func (whitespace *WhiteSpace) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterWhiteSpaceCreateCallback != nil {
		stage.OnAfterWhiteSpaceCreateCallback.OnAfterCreate(stage, whitespace)
	}
}

func (whitespace *WhiteSpace) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWhiteSpaceUpdateCallback != nil {
		var frontWhiteSpace *WhiteSpace
		if front != nil {
			frontWhiteSpace, _ = front.(*WhiteSpace)
		}
		stage.OnAfterWhiteSpaceUpdateCallback.OnAfterUpdate(stage, whitespace, frontWhiteSpace)
	}
}

func (whitespace *WhiteSpace) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterWhiteSpaceDeleteCallback != nil {
		var frontWhiteSpace *WhiteSpace
		if front != nil {
			frontWhiteSpace, _ = front.(*WhiteSpace)
		}
		stage.OnAfterWhiteSpaceDeleteCallback.OnAfterDelete(stage, whitespace, frontWhiteSpace)
	}
}

