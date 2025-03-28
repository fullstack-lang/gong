// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *StageStruct,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Button:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Buttons":
				if _group, ok := stage.Group_Buttons_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		}

	case *Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "Groups":
				if _layout, ok := stage.Layout_Groups_reverseMap[inst]; ok {
					res = _layout.Name
				}
			}
		}

	case *Layout:
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
	case *Button:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Buttons":
				res = stage.Group_Buttons_reverseMap[inst]
			}
		}

	case *Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Layout_Groups_reverseMap[inst]
			}
		}

	case *Layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
