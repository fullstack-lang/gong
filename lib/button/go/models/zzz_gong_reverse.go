// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
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

	case *ButtonToggle:
		switch reverseField.GongstructName {
		// insertion point
		case "GroupToogle":
			switch reverseField.Fieldname {
			case "ButtonToggles":
				if _grouptoogle, ok := stage.GroupToogle_ButtonToggles_reverseMap[inst]; ok {
					res = _grouptoogle.Name
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

	case *GroupToogle:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "GroupToogles":
				if _layout, ok := stage.Layout_GroupToogles_reverseMap[inst]; ok {
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
	stage *Stage,
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

	case *ButtonToggle:
		switch reverseField.GongstructName {
		// insertion point
		case "GroupToogle":
			switch reverseField.Fieldname {
			case "ButtonToggles":
				res = stage.GroupToogle_ButtonToggles_reverseMap[inst]
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

	case *GroupToogle:
		switch reverseField.GongstructName {
		// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "GroupToogles":
				res = stage.Layout_GroupToogles_reverseMap[inst]
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
