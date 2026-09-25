// generated code - do not edit
package models

// insertion point
func (inst *GongBasicField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "GongBasicFields":
			if _gongstruct, ok := stage.GongStruct_GongBasicFields_reverseMap[inst]; ok {
				res = _gongstruct.Name
			}
		}
	}
	return
}

func (inst *GongEnum) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GongEnumValue) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongEnum":
		switch reverseField.Fieldname {
		case "GongEnumValues":
			if _gongenum, ok := stage.GongEnum_GongEnumValues_reverseMap[inst]; ok {
				res = _gongenum.Name
			}
		}
	}
	return
}

func (inst *GongLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongNote":
		switch reverseField.Fieldname {
		case "Links":
			if _gongnote, ok := stage.GongNote_Links_reverseMap[inst]; ok {
				res = _gongnote.Name
			}
		}
	}
	return
}

func (inst *GongNote) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GongStruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GongTimeField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "GongTimeFields":
			if _gongstruct, ok := stage.GongStruct_GongTimeFields_reverseMap[inst]; ok {
				res = _gongstruct.Name
			}
		}
	}
	return
}

func (inst *MetaReference) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ModelPkg) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PointerToGongStructField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "PointerToGongStructFields":
			if _gongstruct, ok := stage.GongStruct_PointerToGongStructFields_reverseMap[inst]; ok {
				res = _gongstruct.Name
			}
		}
	}
	return
}

func (inst *SliceOfPointerToGongStructField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "SliceOfPointerToGongStructFields":
			if _gongstruct, ok := stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]; ok {
				res = _gongstruct.Name
			}
		}
	}
	return
}

func (inst *StageSetField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "StageSetModel":
		switch reverseField.Fieldname {
		case "Fields":
			if _stagesetmodel, ok := stage.StageSetModel_Fields_reverseMap[inst]; ok {
				res = _stagesetmodel.Name
			}
		}
	}
	return
}

func (inst *StageSetModel) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
