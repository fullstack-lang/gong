// generated code - do not edit
package models

// insertion point
func (inst *A) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "As":
				if _a, ok := stage.A_As_reverseMap[inst]; ok {
					res = _a.Name
				}
			}
	}
	return
}


// insertion point
func (inst *A) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "As":
				res = stage.A_As_reverseMap[inst]
			}
	}
	return res
}

