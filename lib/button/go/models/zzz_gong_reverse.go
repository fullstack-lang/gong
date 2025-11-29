// generated code - do not edit
package models

// insertion point
func (inst *Button) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *ButtonToggle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *GroupToogle) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}

func (inst *Layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}


// insertion point
func (inst *Button) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Buttons":
				res = stage.Group_Buttons_reverseMap[inst]
			}
	}
	return res
}

func (inst *ButtonToggle) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "GroupToogle":
			switch reverseField.Fieldname {
			case "ButtonToggles":
				res = stage.GroupToogle_ButtonToggles_reverseMap[inst]
			}
	}
	return res
}

func (inst *Group) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Layout_Groups_reverseMap[inst]
			}
	}
	return res
}

func (inst *GroupToogle) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Layout":
			switch reverseField.Fieldname {
			case "GroupToogles":
				res = stage.Layout_GroupToogles_reverseMap[inst]
			}
	}
	return res
}

func (inst *Layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

