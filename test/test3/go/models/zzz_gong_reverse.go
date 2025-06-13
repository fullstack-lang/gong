// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *A:
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
	case *A:
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "As":
				res = stage.A_As_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
