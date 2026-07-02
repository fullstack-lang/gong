// generated code - do not edit
package models

// insertion point
func (inst *GongBasicField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *GongEnum) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GongEnumValue) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *GongLink) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *GongNote) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GongStruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GongTimeField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *MetaReference) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ModelPkg) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *PointerToGongStructField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *SliceOfPointerToGongStructField) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

// insertion point
func (inst *GongBasicField) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "GongBasicFields":
			res = stage.GongStruct_GongBasicFields_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongEnum) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GongEnumValue) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongEnum":
		switch reverseField.Fieldname {
		case "GongEnumValues":
			res = stage.GongEnum_GongEnumValues_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongLink) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongNote":
		switch reverseField.Fieldname {
		case "Links":
			res = stage.GongNote_Links_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongNote) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GongStruct) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GongTimeField) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "GongTimeFields":
			res = stage.GongStruct_GongTimeFields_reverseMap[inst]
		}
	}
	return res
}

func (inst *MetaReference) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ModelPkg) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *PointerToGongStructField) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "PointerToGongStructFields":
			res = stage.GongStruct_PointerToGongStructFields_reverseMap[inst]
		}
	}
	return res
}

func (inst *SliceOfPointerToGongStructField) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStruct":
		switch reverseField.Fieldname {
		case "SliceOfPointerToGongStructFields":
			res = stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[inst]
		}
	}
	return res
}
