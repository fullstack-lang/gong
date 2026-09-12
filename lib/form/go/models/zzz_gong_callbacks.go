// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CheckBox:
		if stage.OnAfterCheckBoxCreateCallback != nil {
			stage.OnAfterCheckBoxCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormDiv:
		if stage.OnAfterFormDivCreateCallback != nil {
			stage.OnAfterFormDivCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormEditAssocButton:
		if stage.OnAfterFormEditAssocButtonCreateCallback != nil {
			stage.OnAfterFormEditAssocButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormField:
		if stage.OnAfterFormFieldCreateCallback != nil {
			stage.OnAfterFormFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldDate:
		if stage.OnAfterFormFieldDateCreateCallback != nil {
			stage.OnAfterFormFieldDateCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldDateTime:
		if stage.OnAfterFormFieldDateTimeCreateCallback != nil {
			stage.OnAfterFormFieldDateTimeCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldFloat64:
		if stage.OnAfterFormFieldFloat64CreateCallback != nil {
			stage.OnAfterFormFieldFloat64CreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldInt:
		if stage.OnAfterFormFieldIntCreateCallback != nil {
			stage.OnAfterFormFieldIntCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldSelect:
		if stage.OnAfterFormFieldSelectCreateCallback != nil {
			stage.OnAfterFormFieldSelectCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldString:
		if stage.OnAfterFormFieldStringCreateCallback != nil {
			stage.OnAfterFormFieldStringCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldTime:
		if stage.OnAfterFormFieldTimeCreateCallback != nil {
			stage.OnAfterFormFieldTimeCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormGroup:
		if stage.OnAfterFormGroupCreateCallback != nil {
			stage.OnAfterFormGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormSortAssocButton:
		if stage.OnAfterFormSortAssocButtonCreateCallback != nil {
			stage.OnAfterFormSortAssocButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Option:
		if stage.OnAfterOptionCreateCallback != nil {
			stage.OnAfterOptionCreateCallback.OnAfterCreate(stage, target)
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
	case *CheckBox:
		newTarget := any(new).(*CheckBox)
		if stage.OnAfterCheckBoxUpdateCallback != nil {
			stage.OnAfterCheckBoxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormDiv:
		newTarget := any(new).(*FormDiv)
		if stage.OnAfterFormDivUpdateCallback != nil {
			stage.OnAfterFormDivUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormEditAssocButton:
		newTarget := any(new).(*FormEditAssocButton)
		if stage.OnAfterFormEditAssocButtonUpdateCallback != nil {
			stage.OnAfterFormEditAssocButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormField:
		newTarget := any(new).(*FormField)
		if stage.OnAfterFormFieldUpdateCallback != nil {
			stage.OnAfterFormFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldDate:
		newTarget := any(new).(*FormFieldDate)
		if stage.OnAfterFormFieldDateUpdateCallback != nil {
			stage.OnAfterFormFieldDateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldDateTime:
		newTarget := any(new).(*FormFieldDateTime)
		if stage.OnAfterFormFieldDateTimeUpdateCallback != nil {
			stage.OnAfterFormFieldDateTimeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldFloat64:
		newTarget := any(new).(*FormFieldFloat64)
		if stage.OnAfterFormFieldFloat64UpdateCallback != nil {
			stage.OnAfterFormFieldFloat64UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldInt:
		newTarget := any(new).(*FormFieldInt)
		if stage.OnAfterFormFieldIntUpdateCallback != nil {
			stage.OnAfterFormFieldIntUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldSelect:
		newTarget := any(new).(*FormFieldSelect)
		if stage.OnAfterFormFieldSelectUpdateCallback != nil {
			stage.OnAfterFormFieldSelectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldString:
		newTarget := any(new).(*FormFieldString)
		if stage.OnAfterFormFieldStringUpdateCallback != nil {
			stage.OnAfterFormFieldStringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldTime:
		newTarget := any(new).(*FormFieldTime)
		if stage.OnAfterFormFieldTimeUpdateCallback != nil {
			stage.OnAfterFormFieldTimeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormGroup:
		newTarget := any(new).(*FormGroup)
		if stage.OnAfterFormGroupUpdateCallback != nil {
			stage.OnAfterFormGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormSortAssocButton:
		newTarget := any(new).(*FormSortAssocButton)
		if stage.OnAfterFormSortAssocButtonUpdateCallback != nil {
			stage.OnAfterFormSortAssocButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Option:
		newTarget := any(new).(*Option)
		if stage.OnAfterOptionUpdateCallback != nil {
			stage.OnAfterOptionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *CheckBox:
		if stage.OnAfterCheckBoxDeleteCallback != nil {
			staged := any(staged).(*CheckBox)
			stage.OnAfterCheckBoxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormDiv:
		if stage.OnAfterFormDivDeleteCallback != nil {
			staged := any(staged).(*FormDiv)
			stage.OnAfterFormDivDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormEditAssocButton:
		if stage.OnAfterFormEditAssocButtonDeleteCallback != nil {
			staged := any(staged).(*FormEditAssocButton)
			stage.OnAfterFormEditAssocButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormField:
		if stage.OnAfterFormFieldDeleteCallback != nil {
			staged := any(staged).(*FormField)
			stage.OnAfterFormFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldDate:
		if stage.OnAfterFormFieldDateDeleteCallback != nil {
			staged := any(staged).(*FormFieldDate)
			stage.OnAfterFormFieldDateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldDateTime:
		if stage.OnAfterFormFieldDateTimeDeleteCallback != nil {
			staged := any(staged).(*FormFieldDateTime)
			stage.OnAfterFormFieldDateTimeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldFloat64:
		if stage.OnAfterFormFieldFloat64DeleteCallback != nil {
			staged := any(staged).(*FormFieldFloat64)
			stage.OnAfterFormFieldFloat64DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldInt:
		if stage.OnAfterFormFieldIntDeleteCallback != nil {
			staged := any(staged).(*FormFieldInt)
			stage.OnAfterFormFieldIntDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldSelect:
		if stage.OnAfterFormFieldSelectDeleteCallback != nil {
			staged := any(staged).(*FormFieldSelect)
			stage.OnAfterFormFieldSelectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldString:
		if stage.OnAfterFormFieldStringDeleteCallback != nil {
			staged := any(staged).(*FormFieldString)
			stage.OnAfterFormFieldStringDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldTime:
		if stage.OnAfterFormFieldTimeDeleteCallback != nil {
			staged := any(staged).(*FormFieldTime)
			stage.OnAfterFormFieldTimeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormGroup:
		if stage.OnAfterFormGroupDeleteCallback != nil {
			staged := any(staged).(*FormGroup)
			stage.OnAfterFormGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormSortAssocButton:
		if stage.OnAfterFormSortAssocButtonDeleteCallback != nil {
			staged := any(staged).(*FormSortAssocButton)
			stage.OnAfterFormSortAssocButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Option:
		if stage.OnAfterOptionDeleteCallback != nil {
			staged := any(staged).(*Option)
			stage.OnAfterOptionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
