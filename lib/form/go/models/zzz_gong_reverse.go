// generated code - do not edit
package models

// insertion point
func (inst *CheckBox) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FormDiv":
		switch reverseField.Fieldname {
		case "CheckBoxs":
			if _formdiv, ok := stage.FormDiv_CheckBoxs_reverseMap[inst]; ok {
				res = _formdiv.Name
			}
		}
	}
	return
}

func (inst *FormDiv) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FormGroup":
		switch reverseField.Fieldname {
		case "FormDivs":
			if _formgroup, ok := stage.FormGroup_FormDivs_reverseMap[inst]; ok {
				res = _formgroup.Name
			}
		}
	}
	return
}

func (inst *FormEditAssocButton) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FormDiv":
		switch reverseField.Fieldname {
		case "FormFields":
			if _formdiv, ok := stage.FormDiv_FormFields_reverseMap[inst]; ok {
				res = _formdiv.Name
			}
		}
	}
	return
}

func (inst *FormFieldDate) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldDateTime) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldFloat64) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldInt) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldSelect) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldString) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormFieldTime) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormGroup) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormSortAssocButton) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Option) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FormFieldSelect":
		switch reverseField.Fieldname {
		case "Options":
			if _formfieldselect, ok := stage.FormFieldSelect_Options_reverseMap[inst]; ok {
				res = _formfieldselect.Name
			}
		}
	}
	return
}
