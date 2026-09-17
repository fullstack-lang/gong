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
func (checkbox *CheckBox) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckBoxCreateCallback != nil {
		stage.OnAfterCheckBoxCreateCallback.OnAfterCreate(stage, checkbox)
	}
}

func (checkbox *CheckBox) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckBoxUpdateCallback != nil {
		var frontCheckBox *CheckBox
		if front != nil {
			frontCheckBox, _ = front.(*CheckBox)
		}
		stage.OnAfterCheckBoxUpdateCallback.OnAfterUpdate(stage, checkbox, frontCheckBox)
	}
}

func (checkbox *CheckBox) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckBoxDeleteCallback != nil {
		var frontCheckBox *CheckBox
		if front != nil {
			frontCheckBox, _ = front.(*CheckBox)
		}
		stage.OnAfterCheckBoxDeleteCallback.OnAfterDelete(stage, checkbox, frontCheckBox)
	}
}

func (formdiv *FormDiv) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormDivCreateCallback != nil {
		stage.OnAfterFormDivCreateCallback.OnAfterCreate(stage, formdiv)
	}
}

func (formdiv *FormDiv) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormDivUpdateCallback != nil {
		var frontFormDiv *FormDiv
		if front != nil {
			frontFormDiv, _ = front.(*FormDiv)
		}
		stage.OnAfterFormDivUpdateCallback.OnAfterUpdate(stage, formdiv, frontFormDiv)
	}
}

func (formdiv *FormDiv) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormDivDeleteCallback != nil {
		var frontFormDiv *FormDiv
		if front != nil {
			frontFormDiv, _ = front.(*FormDiv)
		}
		stage.OnAfterFormDivDeleteCallback.OnAfterDelete(stage, formdiv, frontFormDiv)
	}
}

func (formeditassocbutton *FormEditAssocButton) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormEditAssocButtonCreateCallback != nil {
		stage.OnAfterFormEditAssocButtonCreateCallback.OnAfterCreate(stage, formeditassocbutton)
	}
}

func (formeditassocbutton *FormEditAssocButton) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormEditAssocButtonUpdateCallback != nil {
		var frontFormEditAssocButton *FormEditAssocButton
		if front != nil {
			frontFormEditAssocButton, _ = front.(*FormEditAssocButton)
		}
		stage.OnAfterFormEditAssocButtonUpdateCallback.OnAfterUpdate(stage, formeditassocbutton, frontFormEditAssocButton)
	}
}

func (formeditassocbutton *FormEditAssocButton) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormEditAssocButtonDeleteCallback != nil {
		var frontFormEditAssocButton *FormEditAssocButton
		if front != nil {
			frontFormEditAssocButton, _ = front.(*FormEditAssocButton)
		}
		stage.OnAfterFormEditAssocButtonDeleteCallback.OnAfterDelete(stage, formeditassocbutton, frontFormEditAssocButton)
	}
}

func (formfield *FormField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldCreateCallback != nil {
		stage.OnAfterFormFieldCreateCallback.OnAfterCreate(stage, formfield)
	}
}

func (formfield *FormField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldUpdateCallback != nil {
		var frontFormField *FormField
		if front != nil {
			frontFormField, _ = front.(*FormField)
		}
		stage.OnAfterFormFieldUpdateCallback.OnAfterUpdate(stage, formfield, frontFormField)
	}
}

func (formfield *FormField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldDeleteCallback != nil {
		var frontFormField *FormField
		if front != nil {
			frontFormField, _ = front.(*FormField)
		}
		stage.OnAfterFormFieldDeleteCallback.OnAfterDelete(stage, formfield, frontFormField)
	}
}

func (formfielddate *FormFieldDate) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldDateCreateCallback != nil {
		stage.OnAfterFormFieldDateCreateCallback.OnAfterCreate(stage, formfielddate)
	}
}

func (formfielddate *FormFieldDate) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldDateUpdateCallback != nil {
		var frontFormFieldDate *FormFieldDate
		if front != nil {
			frontFormFieldDate, _ = front.(*FormFieldDate)
		}
		stage.OnAfterFormFieldDateUpdateCallback.OnAfterUpdate(stage, formfielddate, frontFormFieldDate)
	}
}

func (formfielddate *FormFieldDate) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldDateDeleteCallback != nil {
		var frontFormFieldDate *FormFieldDate
		if front != nil {
			frontFormFieldDate, _ = front.(*FormFieldDate)
		}
		stage.OnAfterFormFieldDateDeleteCallback.OnAfterDelete(stage, formfielddate, frontFormFieldDate)
	}
}

func (formfielddatetime *FormFieldDateTime) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldDateTimeCreateCallback != nil {
		stage.OnAfterFormFieldDateTimeCreateCallback.OnAfterCreate(stage, formfielddatetime)
	}
}

func (formfielddatetime *FormFieldDateTime) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldDateTimeUpdateCallback != nil {
		var frontFormFieldDateTime *FormFieldDateTime
		if front != nil {
			frontFormFieldDateTime, _ = front.(*FormFieldDateTime)
		}
		stage.OnAfterFormFieldDateTimeUpdateCallback.OnAfterUpdate(stage, formfielddatetime, frontFormFieldDateTime)
	}
}

func (formfielddatetime *FormFieldDateTime) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldDateTimeDeleteCallback != nil {
		var frontFormFieldDateTime *FormFieldDateTime
		if front != nil {
			frontFormFieldDateTime, _ = front.(*FormFieldDateTime)
		}
		stage.OnAfterFormFieldDateTimeDeleteCallback.OnAfterDelete(stage, formfielddatetime, frontFormFieldDateTime)
	}
}

func (formfieldfloat64 *FormFieldFloat64) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldFloat64CreateCallback != nil {
		stage.OnAfterFormFieldFloat64CreateCallback.OnAfterCreate(stage, formfieldfloat64)
	}
}

func (formfieldfloat64 *FormFieldFloat64) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldFloat64UpdateCallback != nil {
		var frontFormFieldFloat64 *FormFieldFloat64
		if front != nil {
			frontFormFieldFloat64, _ = front.(*FormFieldFloat64)
		}
		stage.OnAfterFormFieldFloat64UpdateCallback.OnAfterUpdate(stage, formfieldfloat64, frontFormFieldFloat64)
	}
}

func (formfieldfloat64 *FormFieldFloat64) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldFloat64DeleteCallback != nil {
		var frontFormFieldFloat64 *FormFieldFloat64
		if front != nil {
			frontFormFieldFloat64, _ = front.(*FormFieldFloat64)
		}
		stage.OnAfterFormFieldFloat64DeleteCallback.OnAfterDelete(stage, formfieldfloat64, frontFormFieldFloat64)
	}
}

func (formfieldint *FormFieldInt) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldIntCreateCallback != nil {
		stage.OnAfterFormFieldIntCreateCallback.OnAfterCreate(stage, formfieldint)
	}
}

func (formfieldint *FormFieldInt) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldIntUpdateCallback != nil {
		var frontFormFieldInt *FormFieldInt
		if front != nil {
			frontFormFieldInt, _ = front.(*FormFieldInt)
		}
		stage.OnAfterFormFieldIntUpdateCallback.OnAfterUpdate(stage, formfieldint, frontFormFieldInt)
	}
}

func (formfieldint *FormFieldInt) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldIntDeleteCallback != nil {
		var frontFormFieldInt *FormFieldInt
		if front != nil {
			frontFormFieldInt, _ = front.(*FormFieldInt)
		}
		stage.OnAfterFormFieldIntDeleteCallback.OnAfterDelete(stage, formfieldint, frontFormFieldInt)
	}
}

func (formfieldselect *FormFieldSelect) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldSelectCreateCallback != nil {
		stage.OnAfterFormFieldSelectCreateCallback.OnAfterCreate(stage, formfieldselect)
	}
}

func (formfieldselect *FormFieldSelect) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldSelectUpdateCallback != nil {
		var frontFormFieldSelect *FormFieldSelect
		if front != nil {
			frontFormFieldSelect, _ = front.(*FormFieldSelect)
		}
		stage.OnAfterFormFieldSelectUpdateCallback.OnAfterUpdate(stage, formfieldselect, frontFormFieldSelect)
	}
}

func (formfieldselect *FormFieldSelect) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldSelectDeleteCallback != nil {
		var frontFormFieldSelect *FormFieldSelect
		if front != nil {
			frontFormFieldSelect, _ = front.(*FormFieldSelect)
		}
		stage.OnAfterFormFieldSelectDeleteCallback.OnAfterDelete(stage, formfieldselect, frontFormFieldSelect)
	}
}

func (formfieldstring *FormFieldString) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldStringCreateCallback != nil {
		stage.OnAfterFormFieldStringCreateCallback.OnAfterCreate(stage, formfieldstring)
	}
}

func (formfieldstring *FormFieldString) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldStringUpdateCallback != nil {
		var frontFormFieldString *FormFieldString
		if front != nil {
			frontFormFieldString, _ = front.(*FormFieldString)
		}
		stage.OnAfterFormFieldStringUpdateCallback.OnAfterUpdate(stage, formfieldstring, frontFormFieldString)
	}
}

func (formfieldstring *FormFieldString) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldStringDeleteCallback != nil {
		var frontFormFieldString *FormFieldString
		if front != nil {
			frontFormFieldString, _ = front.(*FormFieldString)
		}
		stage.OnAfterFormFieldStringDeleteCallback.OnAfterDelete(stage, formfieldstring, frontFormFieldString)
	}
}

func (formfieldtime *FormFieldTime) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormFieldTimeCreateCallback != nil {
		stage.OnAfterFormFieldTimeCreateCallback.OnAfterCreate(stage, formfieldtime)
	}
}

func (formfieldtime *FormFieldTime) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldTimeUpdateCallback != nil {
		var frontFormFieldTime *FormFieldTime
		if front != nil {
			frontFormFieldTime, _ = front.(*FormFieldTime)
		}
		stage.OnAfterFormFieldTimeUpdateCallback.OnAfterUpdate(stage, formfieldtime, frontFormFieldTime)
	}
}

func (formfieldtime *FormFieldTime) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormFieldTimeDeleteCallback != nil {
		var frontFormFieldTime *FormFieldTime
		if front != nil {
			frontFormFieldTime, _ = front.(*FormFieldTime)
		}
		stage.OnAfterFormFieldTimeDeleteCallback.OnAfterDelete(stage, formfieldtime, frontFormFieldTime)
	}
}

func (formgroup *FormGroup) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormGroupCreateCallback != nil {
		stage.OnAfterFormGroupCreateCallback.OnAfterCreate(stage, formgroup)
	}
}

func (formgroup *FormGroup) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormGroupUpdateCallback != nil {
		var frontFormGroup *FormGroup
		if front != nil {
			frontFormGroup, _ = front.(*FormGroup)
		}
		stage.OnAfterFormGroupUpdateCallback.OnAfterUpdate(stage, formgroup, frontFormGroup)
	}
}

func (formgroup *FormGroup) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormGroupDeleteCallback != nil {
		var frontFormGroup *FormGroup
		if front != nil {
			frontFormGroup, _ = front.(*FormGroup)
		}
		stage.OnAfterFormGroupDeleteCallback.OnAfterDelete(stage, formgroup, frontFormGroup)
	}
}

func (formsortassocbutton *FormSortAssocButton) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormSortAssocButtonCreateCallback != nil {
		stage.OnAfterFormSortAssocButtonCreateCallback.OnAfterCreate(stage, formsortassocbutton)
	}
}

func (formsortassocbutton *FormSortAssocButton) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormSortAssocButtonUpdateCallback != nil {
		var frontFormSortAssocButton *FormSortAssocButton
		if front != nil {
			frontFormSortAssocButton, _ = front.(*FormSortAssocButton)
		}
		stage.OnAfterFormSortAssocButtonUpdateCallback.OnAfterUpdate(stage, formsortassocbutton, frontFormSortAssocButton)
	}
}

func (formsortassocbutton *FormSortAssocButton) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormSortAssocButtonDeleteCallback != nil {
		var frontFormSortAssocButton *FormSortAssocButton
		if front != nil {
			frontFormSortAssocButton, _ = front.(*FormSortAssocButton)
		}
		stage.OnAfterFormSortAssocButtonDeleteCallback.OnAfterDelete(stage, formsortassocbutton, frontFormSortAssocButton)
	}
}

func (option *Option) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOptionCreateCallback != nil {
		stage.OnAfterOptionCreateCallback.OnAfterCreate(stage, option)
	}
}

func (option *Option) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOptionUpdateCallback != nil {
		var frontOption *Option
		if front != nil {
			frontOption, _ = front.(*Option)
		}
		stage.OnAfterOptionUpdateCallback.OnAfterUpdate(stage, option, frontOption)
	}
}

func (option *Option) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOptionDeleteCallback != nil {
		var frontOption *Option
		if front != nil {
			frontOption, _ = front.(*Option)
		}
		stage.OnAfterOptionDeleteCallback.OnAfterDelete(stage, option, frontOption)
	}
}

