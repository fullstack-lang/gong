// generated code - do not edit
package models

// insertion point
func (inst *Checkbox) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Layout) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Slider) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
