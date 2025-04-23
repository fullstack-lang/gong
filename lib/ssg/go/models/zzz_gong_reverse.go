// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Chapter:
		switch reverseField.GongstructName {
		// insertion point
		case "Content":
			switch reverseField.Fieldname {
			case "Chapters":
				if _content, ok := stage.Content_Chapters_reverseMap[inst]; ok {
					res = _content.Name
				}
			}
		}

	case *Content:
		switch reverseField.GongstructName {
		// insertion point
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
	case *Chapter:
		switch reverseField.GongstructName {
		// insertion point
		case "Content":
			switch reverseField.Fieldname {
			case "Chapters":
				res = stage.Content_Chapters_reverseMap[inst]
			}
		}

	case *Content:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
