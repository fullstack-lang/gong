// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
