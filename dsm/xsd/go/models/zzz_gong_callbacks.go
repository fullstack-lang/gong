// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *All:
		if stage.OnAfterAllCreateCallback != nil {
			stage.OnAfterAllCreateCallback.OnAfterCreate(stage, target)
		}
	case *Annotation:
		if stage.OnAfterAnnotationCreateCallback != nil {
			stage.OnAfterAnnotationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Attribute:
		if stage.OnAfterAttributeCreateCallback != nil {
			stage.OnAfterAttributeCreateCallback.OnAfterCreate(stage, target)
		}
	case *AttributeGroup:
		if stage.OnAfterAttributeGroupCreateCallback != nil {
			stage.OnAfterAttributeGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Choice:
		if stage.OnAfterChoiceCreateCallback != nil {
			stage.OnAfterChoiceCreateCallback.OnAfterCreate(stage, target)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentCreateCallback != nil {
			stage.OnAfterComplexContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeCreateCallback != nil {
			stage.OnAfterComplexTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationCreateCallback != nil {
			stage.OnAfterDocumentationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Element:
		if stage.OnAfterElementCreateCallback != nil {
			stage.OnAfterElementCreateCallback.OnAfterCreate(stage, target)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationCreateCallback != nil {
			stage.OnAfterEnumerationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Extension:
		if stage.OnAfterExtensionCreateCallback != nil {
			stage.OnAfterExtensionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Length:
		if stage.OnAfterLengthCreateCallback != nil {
			stage.OnAfterLengthCreateCallback.OnAfterCreate(stage, target)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveCreateCallback != nil {
			stage.OnAfterMaxInclusiveCreateCallback.OnAfterCreate(stage, target)
		}
	case *MaxLength:
		if stage.OnAfterMaxLengthCreateCallback != nil {
			stage.OnAfterMaxLengthCreateCallback.OnAfterCreate(stage, target)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveCreateCallback != nil {
			stage.OnAfterMinInclusiveCreateCallback.OnAfterCreate(stage, target)
		}
	case *MinLength:
		if stage.OnAfterMinLengthCreateCallback != nil {
			stage.OnAfterMinLengthCreateCallback.OnAfterCreate(stage, target)
		}
	case *Pattern:
		if stage.OnAfterPatternCreateCallback != nil {
			stage.OnAfterPatternCreateCallback.OnAfterCreate(stage, target)
		}
	case *Restriction:
		if stage.OnAfterRestrictionCreateCallback != nil {
			stage.OnAfterRestrictionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaCreateCallback != nil {
			stage.OnAfterSchemaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Sequence:
		if stage.OnAfterSequenceCreateCallback != nil {
			stage.OnAfterSequenceCreateCallback.OnAfterCreate(stage, target)
		}
	case *SimpleContent:
		if stage.OnAfterSimpleContentCreateCallback != nil {
			stage.OnAfterSimpleContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeCreateCallback != nil {
			stage.OnAfterSimpleTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TotalDigit:
		if stage.OnAfterTotalDigitCreateCallback != nil {
			stage.OnAfterTotalDigitCreateCallback.OnAfterCreate(stage, target)
		}
	case *Union:
		if stage.OnAfterUnionCreateCallback != nil {
			stage.OnAfterUnionCreateCallback.OnAfterCreate(stage, target)
		}
	case *WhiteSpace:
		if stage.OnAfterWhiteSpaceCreateCallback != nil {
			stage.OnAfterWhiteSpaceCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *All:
		newTarget := any(new).(*All)
		if stage.OnAfterAllUpdateCallback != nil {
			stage.OnAfterAllUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Annotation:
		newTarget := any(new).(*Annotation)
		if stage.OnAfterAnnotationUpdateCallback != nil {
			stage.OnAfterAnnotationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Attribute:
		newTarget := any(new).(*Attribute)
		if stage.OnAfterAttributeUpdateCallback != nil {
			stage.OnAfterAttributeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AttributeGroup:
		newTarget := any(new).(*AttributeGroup)
		if stage.OnAfterAttributeGroupUpdateCallback != nil {
			stage.OnAfterAttributeGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Choice:
		newTarget := any(new).(*Choice)
		if stage.OnAfterChoiceUpdateCallback != nil {
			stage.OnAfterChoiceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ComplexContent:
		newTarget := any(new).(*ComplexContent)
		if stage.OnAfterComplexContentUpdateCallback != nil {
			stage.OnAfterComplexContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ComplexType:
		newTarget := any(new).(*ComplexType)
		if stage.OnAfterComplexTypeUpdateCallback != nil {
			stage.OnAfterComplexTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Documentation:
		newTarget := any(new).(*Documentation)
		if stage.OnAfterDocumentationUpdateCallback != nil {
			stage.OnAfterDocumentationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Element:
		newTarget := any(new).(*Element)
		if stage.OnAfterElementUpdateCallback != nil {
			stage.OnAfterElementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Enumeration:
		newTarget := any(new).(*Enumeration)
		if stage.OnAfterEnumerationUpdateCallback != nil {
			stage.OnAfterEnumerationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Extension:
		newTarget := any(new).(*Extension)
		if stage.OnAfterExtensionUpdateCallback != nil {
			stage.OnAfterExtensionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Length:
		newTarget := any(new).(*Length)
		if stage.OnAfterLengthUpdateCallback != nil {
			stage.OnAfterLengthUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MaxInclusive:
		newTarget := any(new).(*MaxInclusive)
		if stage.OnAfterMaxInclusiveUpdateCallback != nil {
			stage.OnAfterMaxInclusiveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MaxLength:
		newTarget := any(new).(*MaxLength)
		if stage.OnAfterMaxLengthUpdateCallback != nil {
			stage.OnAfterMaxLengthUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MinInclusive:
		newTarget := any(new).(*MinInclusive)
		if stage.OnAfterMinInclusiveUpdateCallback != nil {
			stage.OnAfterMinInclusiveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MinLength:
		newTarget := any(new).(*MinLength)
		if stage.OnAfterMinLengthUpdateCallback != nil {
			stage.OnAfterMinLengthUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Pattern:
		newTarget := any(new).(*Pattern)
		if stage.OnAfterPatternUpdateCallback != nil {
			stage.OnAfterPatternUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Restriction:
		newTarget := any(new).(*Restriction)
		if stage.OnAfterRestrictionUpdateCallback != nil {
			stage.OnAfterRestrictionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Schema:
		newTarget := any(new).(*Schema)
		if stage.OnAfterSchemaUpdateCallback != nil {
			stage.OnAfterSchemaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Sequence:
		newTarget := any(new).(*Sequence)
		if stage.OnAfterSequenceUpdateCallback != nil {
			stage.OnAfterSequenceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SimpleContent:
		newTarget := any(new).(*SimpleContent)
		if stage.OnAfterSimpleContentUpdateCallback != nil {
			stage.OnAfterSimpleContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SimpleType:
		newTarget := any(new).(*SimpleType)
		if stage.OnAfterSimpleTypeUpdateCallback != nil {
			stage.OnAfterSimpleTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TotalDigit:
		newTarget := any(new).(*TotalDigit)
		if stage.OnAfterTotalDigitUpdateCallback != nil {
			stage.OnAfterTotalDigitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Union:
		newTarget := any(new).(*Union)
		if stage.OnAfterUnionUpdateCallback != nil {
			stage.OnAfterUnionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *WhiteSpace:
		newTarget := any(new).(*WhiteSpace)
		if stage.OnAfterWhiteSpaceUpdateCallback != nil {
			stage.OnAfterWhiteSpaceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *All:
		if stage.OnAfterAllDeleteCallback != nil {
			staged := any(staged).(*All)
			stage.OnAfterAllDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Annotation:
		if stage.OnAfterAnnotationDeleteCallback != nil {
			staged := any(staged).(*Annotation)
			stage.OnAfterAnnotationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Attribute:
		if stage.OnAfterAttributeDeleteCallback != nil {
			staged := any(staged).(*Attribute)
			stage.OnAfterAttributeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AttributeGroup:
		if stage.OnAfterAttributeGroupDeleteCallback != nil {
			staged := any(staged).(*AttributeGroup)
			stage.OnAfterAttributeGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Choice:
		if stage.OnAfterChoiceDeleteCallback != nil {
			staged := any(staged).(*Choice)
			stage.OnAfterChoiceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentDeleteCallback != nil {
			staged := any(staged).(*ComplexContent)
			stage.OnAfterComplexContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeDeleteCallback != nil {
			staged := any(staged).(*ComplexType)
			stage.OnAfterComplexTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Documentation:
		if stage.OnAfterDocumentationDeleteCallback != nil {
			staged := any(staged).(*Documentation)
			stage.OnAfterDocumentationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Element:
		if stage.OnAfterElementDeleteCallback != nil {
			staged := any(staged).(*Element)
			stage.OnAfterElementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationDeleteCallback != nil {
			staged := any(staged).(*Enumeration)
			stage.OnAfterEnumerationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Extension:
		if stage.OnAfterExtensionDeleteCallback != nil {
			staged := any(staged).(*Extension)
			stage.OnAfterExtensionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Length:
		if stage.OnAfterLengthDeleteCallback != nil {
			staged := any(staged).(*Length)
			stage.OnAfterLengthDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveDeleteCallback != nil {
			staged := any(staged).(*MaxInclusive)
			stage.OnAfterMaxInclusiveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MaxLength:
		if stage.OnAfterMaxLengthDeleteCallback != nil {
			staged := any(staged).(*MaxLength)
			stage.OnAfterMaxLengthDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveDeleteCallback != nil {
			staged := any(staged).(*MinInclusive)
			stage.OnAfterMinInclusiveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MinLength:
		if stage.OnAfterMinLengthDeleteCallback != nil {
			staged := any(staged).(*MinLength)
			stage.OnAfterMinLengthDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Pattern:
		if stage.OnAfterPatternDeleteCallback != nil {
			staged := any(staged).(*Pattern)
			stage.OnAfterPatternDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Restriction:
		if stage.OnAfterRestrictionDeleteCallback != nil {
			staged := any(staged).(*Restriction)
			stage.OnAfterRestrictionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Schema:
		if stage.OnAfterSchemaDeleteCallback != nil {
			staged := any(staged).(*Schema)
			stage.OnAfterSchemaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Sequence:
		if stage.OnAfterSequenceDeleteCallback != nil {
			staged := any(staged).(*Sequence)
			stage.OnAfterSequenceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SimpleContent:
		if stage.OnAfterSimpleContentDeleteCallback != nil {
			staged := any(staged).(*SimpleContent)
			stage.OnAfterSimpleContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeDeleteCallback != nil {
			staged := any(staged).(*SimpleType)
			stage.OnAfterSimpleTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TotalDigit:
		if stage.OnAfterTotalDigitDeleteCallback != nil {
			staged := any(staged).(*TotalDigit)
			stage.OnAfterTotalDigitDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Union:
		if stage.OnAfterUnionDeleteCallback != nil {
			staged := any(staged).(*Union)
			stage.OnAfterUnionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *WhiteSpace:
		if stage.OnAfterWhiteSpaceDeleteCallback != nil {
			staged := any(staged).(*WhiteSpace)
			stage.OnAfterWhiteSpaceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *All:
		if stage.OnAfterAllReadCallback != nil {
			stage.OnAfterAllReadCallback.OnAfterRead(stage, target)
		}
	case *Annotation:
		if stage.OnAfterAnnotationReadCallback != nil {
			stage.OnAfterAnnotationReadCallback.OnAfterRead(stage, target)
		}
	case *Attribute:
		if stage.OnAfterAttributeReadCallback != nil {
			stage.OnAfterAttributeReadCallback.OnAfterRead(stage, target)
		}
	case *AttributeGroup:
		if stage.OnAfterAttributeGroupReadCallback != nil {
			stage.OnAfterAttributeGroupReadCallback.OnAfterRead(stage, target)
		}
	case *Choice:
		if stage.OnAfterChoiceReadCallback != nil {
			stage.OnAfterChoiceReadCallback.OnAfterRead(stage, target)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentReadCallback != nil {
			stage.OnAfterComplexContentReadCallback.OnAfterRead(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeReadCallback != nil {
			stage.OnAfterComplexTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationReadCallback != nil {
			stage.OnAfterDocumentationReadCallback.OnAfterRead(stage, target)
		}
	case *Element:
		if stage.OnAfterElementReadCallback != nil {
			stage.OnAfterElementReadCallback.OnAfterRead(stage, target)
		}
	case *Enumeration:
		if stage.OnAfterEnumerationReadCallback != nil {
			stage.OnAfterEnumerationReadCallback.OnAfterRead(stage, target)
		}
	case *Extension:
		if stage.OnAfterExtensionReadCallback != nil {
			stage.OnAfterExtensionReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *Length:
		if stage.OnAfterLengthReadCallback != nil {
			stage.OnAfterLengthReadCallback.OnAfterRead(stage, target)
		}
	case *MaxInclusive:
		if stage.OnAfterMaxInclusiveReadCallback != nil {
			stage.OnAfterMaxInclusiveReadCallback.OnAfterRead(stage, target)
		}
	case *MaxLength:
		if stage.OnAfterMaxLengthReadCallback != nil {
			stage.OnAfterMaxLengthReadCallback.OnAfterRead(stage, target)
		}
	case *MinInclusive:
		if stage.OnAfterMinInclusiveReadCallback != nil {
			stage.OnAfterMinInclusiveReadCallback.OnAfterRead(stage, target)
		}
	case *MinLength:
		if stage.OnAfterMinLengthReadCallback != nil {
			stage.OnAfterMinLengthReadCallback.OnAfterRead(stage, target)
		}
	case *Pattern:
		if stage.OnAfterPatternReadCallback != nil {
			stage.OnAfterPatternReadCallback.OnAfterRead(stage, target)
		}
	case *Restriction:
		if stage.OnAfterRestrictionReadCallback != nil {
			stage.OnAfterRestrictionReadCallback.OnAfterRead(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaReadCallback != nil {
			stage.OnAfterSchemaReadCallback.OnAfterRead(stage, target)
		}
	case *Sequence:
		if stage.OnAfterSequenceReadCallback != nil {
			stage.OnAfterSequenceReadCallback.OnAfterRead(stage, target)
		}
	case *SimpleContent:
		if stage.OnAfterSimpleContentReadCallback != nil {
			stage.OnAfterSimpleContentReadCallback.OnAfterRead(stage, target)
		}
	case *SimpleType:
		if stage.OnAfterSimpleTypeReadCallback != nil {
			stage.OnAfterSimpleTypeReadCallback.OnAfterRead(stage, target)
		}
	case *TotalDigit:
		if stage.OnAfterTotalDigitReadCallback != nil {
			stage.OnAfterTotalDigitReadCallback.OnAfterRead(stage, target)
		}
	case *Union:
		if stage.OnAfterUnionReadCallback != nil {
			stage.OnAfterUnionReadCallback.OnAfterRead(stage, target)
		}
	case *WhiteSpace:
		if stage.OnAfterWhiteSpaceReadCallback != nil {
			stage.OnAfterWhiteSpaceReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *All:
		stage.OnAfterAllUpdateCallback = any(callback).(OnAfterUpdateInterface[All])
	case *Annotation:
		stage.OnAfterAnnotationUpdateCallback = any(callback).(OnAfterUpdateInterface[Annotation])
	case *Attribute:
		stage.OnAfterAttributeUpdateCallback = any(callback).(OnAfterUpdateInterface[Attribute])
	case *AttributeGroup:
		stage.OnAfterAttributeGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[AttributeGroup])
	case *Choice:
		stage.OnAfterChoiceUpdateCallback = any(callback).(OnAfterUpdateInterface[Choice])
	case *ComplexContent:
		stage.OnAfterComplexContentUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexContent])
	case *ComplexType:
		stage.OnAfterComplexTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexType])
	case *Documentation:
		stage.OnAfterDocumentationUpdateCallback = any(callback).(OnAfterUpdateInterface[Documentation])
	case *Element:
		stage.OnAfterElementUpdateCallback = any(callback).(OnAfterUpdateInterface[Element])
	case *Enumeration:
		stage.OnAfterEnumerationUpdateCallback = any(callback).(OnAfterUpdateInterface[Enumeration])
	case *Extension:
		stage.OnAfterExtensionUpdateCallback = any(callback).(OnAfterUpdateInterface[Extension])
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	case *Length:
		stage.OnAfterLengthUpdateCallback = any(callback).(OnAfterUpdateInterface[Length])
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveUpdateCallback = any(callback).(OnAfterUpdateInterface[MaxInclusive])
	case *MaxLength:
		stage.OnAfterMaxLengthUpdateCallback = any(callback).(OnAfterUpdateInterface[MaxLength])
	case *MinInclusive:
		stage.OnAfterMinInclusiveUpdateCallback = any(callback).(OnAfterUpdateInterface[MinInclusive])
	case *MinLength:
		stage.OnAfterMinLengthUpdateCallback = any(callback).(OnAfterUpdateInterface[MinLength])
	case *Pattern:
		stage.OnAfterPatternUpdateCallback = any(callback).(OnAfterUpdateInterface[Pattern])
	case *Restriction:
		stage.OnAfterRestrictionUpdateCallback = any(callback).(OnAfterUpdateInterface[Restriction])
	case *Schema:
		stage.OnAfterSchemaUpdateCallback = any(callback).(OnAfterUpdateInterface[Schema])
	case *Sequence:
		stage.OnAfterSequenceUpdateCallback = any(callback).(OnAfterUpdateInterface[Sequence])
	case *SimpleContent:
		stage.OnAfterSimpleContentUpdateCallback = any(callback).(OnAfterUpdateInterface[SimpleContent])
	case *SimpleType:
		stage.OnAfterSimpleTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[SimpleType])
	case *TotalDigit:
		stage.OnAfterTotalDigitUpdateCallback = any(callback).(OnAfterUpdateInterface[TotalDigit])
	case *Union:
		stage.OnAfterUnionUpdateCallback = any(callback).(OnAfterUpdateInterface[Union])
	case *WhiteSpace:
		stage.OnAfterWhiteSpaceUpdateCallback = any(callback).(OnAfterUpdateInterface[WhiteSpace])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *All:
		stage.OnAfterAllCreateCallback = any(callback).(OnAfterCreateInterface[All])
	case *Annotation:
		stage.OnAfterAnnotationCreateCallback = any(callback).(OnAfterCreateInterface[Annotation])
	case *Attribute:
		stage.OnAfterAttributeCreateCallback = any(callback).(OnAfterCreateInterface[Attribute])
	case *AttributeGroup:
		stage.OnAfterAttributeGroupCreateCallback = any(callback).(OnAfterCreateInterface[AttributeGroup])
	case *Choice:
		stage.OnAfterChoiceCreateCallback = any(callback).(OnAfterCreateInterface[Choice])
	case *ComplexContent:
		stage.OnAfterComplexContentCreateCallback = any(callback).(OnAfterCreateInterface[ComplexContent])
	case *ComplexType:
		stage.OnAfterComplexTypeCreateCallback = any(callback).(OnAfterCreateInterface[ComplexType])
	case *Documentation:
		stage.OnAfterDocumentationCreateCallback = any(callback).(OnAfterCreateInterface[Documentation])
	case *Element:
		stage.OnAfterElementCreateCallback = any(callback).(OnAfterCreateInterface[Element])
	case *Enumeration:
		stage.OnAfterEnumerationCreateCallback = any(callback).(OnAfterCreateInterface[Enumeration])
	case *Extension:
		stage.OnAfterExtensionCreateCallback = any(callback).(OnAfterCreateInterface[Extension])
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	case *Length:
		stage.OnAfterLengthCreateCallback = any(callback).(OnAfterCreateInterface[Length])
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveCreateCallback = any(callback).(OnAfterCreateInterface[MaxInclusive])
	case *MaxLength:
		stage.OnAfterMaxLengthCreateCallback = any(callback).(OnAfterCreateInterface[MaxLength])
	case *MinInclusive:
		stage.OnAfterMinInclusiveCreateCallback = any(callback).(OnAfterCreateInterface[MinInclusive])
	case *MinLength:
		stage.OnAfterMinLengthCreateCallback = any(callback).(OnAfterCreateInterface[MinLength])
	case *Pattern:
		stage.OnAfterPatternCreateCallback = any(callback).(OnAfterCreateInterface[Pattern])
	case *Restriction:
		stage.OnAfterRestrictionCreateCallback = any(callback).(OnAfterCreateInterface[Restriction])
	case *Schema:
		stage.OnAfterSchemaCreateCallback = any(callback).(OnAfterCreateInterface[Schema])
	case *Sequence:
		stage.OnAfterSequenceCreateCallback = any(callback).(OnAfterCreateInterface[Sequence])
	case *SimpleContent:
		stage.OnAfterSimpleContentCreateCallback = any(callback).(OnAfterCreateInterface[SimpleContent])
	case *SimpleType:
		stage.OnAfterSimpleTypeCreateCallback = any(callback).(OnAfterCreateInterface[SimpleType])
	case *TotalDigit:
		stage.OnAfterTotalDigitCreateCallback = any(callback).(OnAfterCreateInterface[TotalDigit])
	case *Union:
		stage.OnAfterUnionCreateCallback = any(callback).(OnAfterCreateInterface[Union])
	case *WhiteSpace:
		stage.OnAfterWhiteSpaceCreateCallback = any(callback).(OnAfterCreateInterface[WhiteSpace])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *All:
		stage.OnAfterAllDeleteCallback = any(callback).(OnAfterDeleteInterface[All])
	case *Annotation:
		stage.OnAfterAnnotationDeleteCallback = any(callback).(OnAfterDeleteInterface[Annotation])
	case *Attribute:
		stage.OnAfterAttributeDeleteCallback = any(callback).(OnAfterDeleteInterface[Attribute])
	case *AttributeGroup:
		stage.OnAfterAttributeGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[AttributeGroup])
	case *Choice:
		stage.OnAfterChoiceDeleteCallback = any(callback).(OnAfterDeleteInterface[Choice])
	case *ComplexContent:
		stage.OnAfterComplexContentDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexContent])
	case *ComplexType:
		stage.OnAfterComplexTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexType])
	case *Documentation:
		stage.OnAfterDocumentationDeleteCallback = any(callback).(OnAfterDeleteInterface[Documentation])
	case *Element:
		stage.OnAfterElementDeleteCallback = any(callback).(OnAfterDeleteInterface[Element])
	case *Enumeration:
		stage.OnAfterEnumerationDeleteCallback = any(callback).(OnAfterDeleteInterface[Enumeration])
	case *Extension:
		stage.OnAfterExtensionDeleteCallback = any(callback).(OnAfterDeleteInterface[Extension])
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	case *Length:
		stage.OnAfterLengthDeleteCallback = any(callback).(OnAfterDeleteInterface[Length])
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveDeleteCallback = any(callback).(OnAfterDeleteInterface[MaxInclusive])
	case *MaxLength:
		stage.OnAfterMaxLengthDeleteCallback = any(callback).(OnAfterDeleteInterface[MaxLength])
	case *MinInclusive:
		stage.OnAfterMinInclusiveDeleteCallback = any(callback).(OnAfterDeleteInterface[MinInclusive])
	case *MinLength:
		stage.OnAfterMinLengthDeleteCallback = any(callback).(OnAfterDeleteInterface[MinLength])
	case *Pattern:
		stage.OnAfterPatternDeleteCallback = any(callback).(OnAfterDeleteInterface[Pattern])
	case *Restriction:
		stage.OnAfterRestrictionDeleteCallback = any(callback).(OnAfterDeleteInterface[Restriction])
	case *Schema:
		stage.OnAfterSchemaDeleteCallback = any(callback).(OnAfterDeleteInterface[Schema])
	case *Sequence:
		stage.OnAfterSequenceDeleteCallback = any(callback).(OnAfterDeleteInterface[Sequence])
	case *SimpleContent:
		stage.OnAfterSimpleContentDeleteCallback = any(callback).(OnAfterDeleteInterface[SimpleContent])
	case *SimpleType:
		stage.OnAfterSimpleTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[SimpleType])
	case *TotalDigit:
		stage.OnAfterTotalDigitDeleteCallback = any(callback).(OnAfterDeleteInterface[TotalDigit])
	case *Union:
		stage.OnAfterUnionDeleteCallback = any(callback).(OnAfterDeleteInterface[Union])
	case *WhiteSpace:
		stage.OnAfterWhiteSpaceDeleteCallback = any(callback).(OnAfterDeleteInterface[WhiteSpace])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *All:
		stage.OnAfterAllReadCallback = any(callback).(OnAfterReadInterface[All])
	case *Annotation:
		stage.OnAfterAnnotationReadCallback = any(callback).(OnAfterReadInterface[Annotation])
	case *Attribute:
		stage.OnAfterAttributeReadCallback = any(callback).(OnAfterReadInterface[Attribute])
	case *AttributeGroup:
		stage.OnAfterAttributeGroupReadCallback = any(callback).(OnAfterReadInterface[AttributeGroup])
	case *Choice:
		stage.OnAfterChoiceReadCallback = any(callback).(OnAfterReadInterface[Choice])
	case *ComplexContent:
		stage.OnAfterComplexContentReadCallback = any(callback).(OnAfterReadInterface[ComplexContent])
	case *ComplexType:
		stage.OnAfterComplexTypeReadCallback = any(callback).(OnAfterReadInterface[ComplexType])
	case *Documentation:
		stage.OnAfterDocumentationReadCallback = any(callback).(OnAfterReadInterface[Documentation])
	case *Element:
		stage.OnAfterElementReadCallback = any(callback).(OnAfterReadInterface[Element])
	case *Enumeration:
		stage.OnAfterEnumerationReadCallback = any(callback).(OnAfterReadInterface[Enumeration])
	case *Extension:
		stage.OnAfterExtensionReadCallback = any(callback).(OnAfterReadInterface[Extension])
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	case *Length:
		stage.OnAfterLengthReadCallback = any(callback).(OnAfterReadInterface[Length])
	case *MaxInclusive:
		stage.OnAfterMaxInclusiveReadCallback = any(callback).(OnAfterReadInterface[MaxInclusive])
	case *MaxLength:
		stage.OnAfterMaxLengthReadCallback = any(callback).(OnAfterReadInterface[MaxLength])
	case *MinInclusive:
		stage.OnAfterMinInclusiveReadCallback = any(callback).(OnAfterReadInterface[MinInclusive])
	case *MinLength:
		stage.OnAfterMinLengthReadCallback = any(callback).(OnAfterReadInterface[MinLength])
	case *Pattern:
		stage.OnAfterPatternReadCallback = any(callback).(OnAfterReadInterface[Pattern])
	case *Restriction:
		stage.OnAfterRestrictionReadCallback = any(callback).(OnAfterReadInterface[Restriction])
	case *Schema:
		stage.OnAfterSchemaReadCallback = any(callback).(OnAfterReadInterface[Schema])
	case *Sequence:
		stage.OnAfterSequenceReadCallback = any(callback).(OnAfterReadInterface[Sequence])
	case *SimpleContent:
		stage.OnAfterSimpleContentReadCallback = any(callback).(OnAfterReadInterface[SimpleContent])
	case *SimpleType:
		stage.OnAfterSimpleTypeReadCallback = any(callback).(OnAfterReadInterface[SimpleType])
	case *TotalDigit:
		stage.OnAfterTotalDigitReadCallback = any(callback).(OnAfterReadInterface[TotalDigit])
	case *Union:
		stage.OnAfterUnionReadCallback = any(callback).(OnAfterReadInterface[Union])
	case *WhiteSpace:
		stage.OnAfterWhiteSpaceReadCallback = any(callback).(OnAfterReadInterface[WhiteSpace])
	}
}
