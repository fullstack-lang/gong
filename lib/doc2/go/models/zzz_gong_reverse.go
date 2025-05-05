// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *AttributeShape:
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

	case *Classdiagram:
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

	case *DiagramPackage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *GongEnumShape:
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

	case *GongEnumValueShape:
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

	case *GongNoteLinkShape:
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

	case *GongNoteShape:
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

	case *GongStructShape:
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

	case *LinkShape:
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

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *AttributeShape:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "AttributeShapes":
				res = stage.GongStructShape_AttributeShapes_reverseMap[inst]
			}
		}

	case *Classdiagram:
		switch reverseField.GongstructName {
		// insertion point
		case "DiagramPackage":
			switch reverseField.Fieldname {
			case "Classdiagrams":
				res = stage.DiagramPackage_Classdiagrams_reverseMap[inst]
			}
		}

	case *DiagramPackage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *GongEnumShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongEnumShapes":
				res = stage.Classdiagram_GongEnumShapes_reverseMap[inst]
			}
		}

	case *GongEnumValueShape:
		switch reverseField.GongstructName {
		// insertion point
		case "GongEnumShape":
			switch reverseField.Fieldname {
			case "GongEnumValueShapes":
				res = stage.GongEnumShape_GongEnumValueShapes_reverseMap[inst]
			}
		}

	case *GongNoteLinkShape:
		switch reverseField.GongstructName {
		// insertion point
		case "GongNoteShape":
			switch reverseField.Fieldname {
			case "GongNoteLinkShapes":
				res = stage.GongNoteShape_GongNoteLinkShapes_reverseMap[inst]
			}
		}

	case *GongNoteShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongNoteShapes":
				res = stage.Classdiagram_GongNoteShapes_reverseMap[inst]
			}
		}

	case *GongStructShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Classdiagram":
			switch reverseField.Fieldname {
			case "GongStructShapes":
				res = stage.Classdiagram_GongStructShapes_reverseMap[inst]
			}
		}

	case *LinkShape:
		switch reverseField.GongstructName {
		// insertion point
		case "GongStructShape":
			switch reverseField.Fieldname {
			case "LinkShapes":
				res = stage.GongStructShape_LinkShapes_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
