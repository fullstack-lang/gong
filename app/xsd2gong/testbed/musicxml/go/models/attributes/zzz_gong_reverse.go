// generated code - do not edit
package attributes

// insertion point
func (inst *Attributes) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AttributesDirective) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Directive":
			if _attributes, ok := stage.Attributes_Directive_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *BeatRepeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Cancel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Clef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Clef":
			if _attributes, ok := stage.Attributes_Clef_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Double) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ForPart) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "ForPart":
			if _attributes, ok := stage.Attributes_ForPart_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Interchangeable) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Key) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Key":
			if _attributes, ok := stage.Attributes_Key_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *KeyAccidental) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Key":
		switch reverseField.Fieldname {
		case "KeyAccidental":
			if _key, ok := stage.Key_KeyAccidental_reverseMap[inst]; ok {
				res = _key.Name
			}
		}
	}
	return
}

func (inst *KeyOctave) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Key":
		switch reverseField.Fieldname {
		case "KeyOctave":
			if _key, ok := stage.Key_KeyOctave_reverseMap[inst]; ok {
				res = _key.Name
			}
		}
	}
	return
}

func (inst *LineDetail) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "StaffDetails":
		switch reverseField.Fieldname {
		case "LineDetail":
			if _staffdetails, ok := stage.StaffDetails_LineDetail_reverseMap[inst]; ok {
				res = _staffdetails.Name
			}
		}
	}
	return
}

func (inst *MeasureRepeat) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MeasureStyle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "MeasureStyle":
			if _attributes, ok := stage.Attributes_MeasureStyle_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *MultipleRest) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PartClef) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PartSymbol) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PartTranspose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Slash) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StaffDetails) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "StaffDetails":
			if _attributes, ok := stage.Attributes_StaffDetails_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *StaffSize) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *StaffTuning) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "StaffDetails":
		switch reverseField.Fieldname {
		case "StaffTuning":
			if _staffdetails, ok := stage.StaffDetails_StaffTuning_reverseMap[inst]; ok {
				res = _staffdetails.Name
			}
		}
	}
	return
}

func (inst *Time) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Time":
			if _attributes, ok := stage.Attributes_Time_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}

func (inst *Transpose) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Attributes":
		switch reverseField.Fieldname {
		case "Transpose":
			if _attributes, ok := stage.Attributes_Transpose_reverseMap[inst]; ok {
				res = _attributes.Name
			}
		}
	}
	return
}
