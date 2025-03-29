// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Checkbox:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Checkboxes":
				if _group, ok := stage.Group_Checkboxes_reverseMap[inst]; ok {
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

	case *Slider:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Sliders":
				if _group, ok := stage.Group_Sliders_reverseMap[inst]; ok {
					res = _group.Name
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
	case *Checkbox:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Checkboxes":
				res = stage.Group_Checkboxes_reverseMap[inst]
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

	case *Slider:
		switch reverseField.GongstructName {
		// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Sliders":
				res = stage.Group_Sliders_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
