// generated code - do not edit
package models

// insertion point
func (inst *Checkbox) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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

func (inst *Layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Slider) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
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
	return
}


// insertion point
func (inst *Checkbox) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Checkboxes":
				res = stage.Group_Checkboxes_reverseMap[inst]
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

func (inst *Layout) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Slider) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Group":
			switch reverseField.Fieldname {
			case "Sliders":
				res = stage.Group_Sliders_reverseMap[inst]
			}
	}
	return res
}

