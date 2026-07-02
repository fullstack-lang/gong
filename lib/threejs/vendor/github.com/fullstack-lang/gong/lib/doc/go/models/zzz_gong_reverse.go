// generated code - do not edit
package models

// insertion point
func (inst *AttributeShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStructShape":
		switch reverseField.Fieldname {
		case "AttributeShapes":
			if _gongstructshape, ok := stage.GongStructShape_AttributeShapes_reverseMap[inst]; ok {
				res = _gongstructshape.Name
			}
		}
	}
	return
}

func (inst *Classdiagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramPackage":
		switch reverseField.Fieldname {
		case "Classdiagrams":
			if _diagrampackage, ok := stage.DiagramPackage_Classdiagrams_reverseMap[inst]; ok {
				res = _diagrampackage.Name
			}
		}
	}
	return
}

func (inst *DiagramPackage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GongEnumShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongEnumShapes":
			if _classdiagram, ok := stage.Classdiagram_GongEnumShapes_reverseMap[inst]; ok {
				res = _classdiagram.Name
			}
		}
	}
	return
}

func (inst *GongEnumValueShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongEnumShape":
		switch reverseField.Fieldname {
		case "GongEnumValueShapes":
			if _gongenumshape, ok := stage.GongEnumShape_GongEnumValueShapes_reverseMap[inst]; ok {
				res = _gongenumshape.Name
			}
		}
	}
	return
}

func (inst *GongNoteLinkShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongNoteShape":
		switch reverseField.Fieldname {
		case "GongNoteLinkShapes":
			if _gongnoteshape, ok := stage.GongNoteShape_GongNoteLinkShapes_reverseMap[inst]; ok {
				res = _gongnoteshape.Name
			}
		}
	}
	return
}

func (inst *GongNoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongNoteShapes":
			if _classdiagram, ok := stage.Classdiagram_GongNoteShapes_reverseMap[inst]; ok {
				res = _classdiagram.Name
			}
		}
	}
	return
}

func (inst *GongStructShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongStructShapes":
			if _classdiagram, ok := stage.Classdiagram_GongStructShapes_reverseMap[inst]; ok {
				res = _classdiagram.Name
			}
		}
	}
	return
}

func (inst *LinkShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GongStructShape":
		switch reverseField.Fieldname {
		case "LinkShapes":
			if _gongstructshape, ok := stage.GongStructShape_LinkShapes_reverseMap[inst]; ok {
				res = _gongstructshape.Name
			}
		}
	}
	return
}

// insertion point
func (inst *AttributeShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStructShape":
		switch reverseField.Fieldname {
		case "AttributeShapes":
			res = stage.GongStructShape_AttributeShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Classdiagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramPackage":
		switch reverseField.Fieldname {
		case "Classdiagrams":
			res = stage.DiagramPackage_Classdiagrams_reverseMap[inst]
		}
	}
	return res
}

func (inst *DiagramPackage) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GongEnumShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongEnumShapes":
			res = stage.Classdiagram_GongEnumShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongEnumValueShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongEnumShape":
		switch reverseField.Fieldname {
		case "GongEnumValueShapes":
			res = stage.GongEnumShape_GongEnumValueShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongNoteLinkShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongNoteShape":
		switch reverseField.Fieldname {
		case "GongNoteLinkShapes":
			res = stage.GongNoteShape_GongNoteLinkShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongNoteShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongNoteShapes":
			res = stage.Classdiagram_GongNoteShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GongStructShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Classdiagram":
		switch reverseField.Fieldname {
		case "GongStructShapes":
			res = stage.Classdiagram_GongStructShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *LinkShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GongStructShape":
		switch reverseField.Fieldname {
		case "LinkShapes":
			res = stage.GongStructShape_LinkShapes_reverseMap[inst]
		}
	}
	return res
}
