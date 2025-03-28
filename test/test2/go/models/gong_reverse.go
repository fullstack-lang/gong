// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *StageStruct,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *A:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *B:
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				if _a, ok := stage.A_Bs_reverseMap[inst]; ok {
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
	stage *StageStruct,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *A:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *B:
		switch reverseField.GongstructName {
		// insertion point
		case "A":
			switch reverseField.Fieldname {
			case "Bs":
				res = stage.A_Bs_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
