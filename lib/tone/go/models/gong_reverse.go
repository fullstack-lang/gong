// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *StageStruct,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Freqency:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Frequencies":
				if _note, ok := stage.Note_Frequencies_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Player:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *StageStruct,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *Freqency:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Frequencies":
				res = stage.Note_Frequencies_reverseMap[inst]
			}
		}

	case *Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Player:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
