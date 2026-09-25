// generated code - do not edit
package models

// insertion point
func (inst *Freqency) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Player) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
