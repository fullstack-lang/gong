// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	clear(stage.Group_Buttons_reverseMap)
	stage.Group_Buttons_reverseMap = make(map[*Button]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _button := range group.Buttons {
			stage.Group_Buttons_reverseMap[_button] = group
		}
	}

	// Compute reverse map for named struct Layout
	// insertion point per field
	clear(stage.Layout_Groups_reverseMap)
	stage.Layout_Groups_reverseMap = make(map[*Group]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _group := range layout.Groups {
			stage.Layout_Groups_reverseMap[_group] = layout
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Layouts {
		res = append(res, instance)
	}


	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	var newInstance Button
	newInstance = *button
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	var newInstance Group
	newInstance = *group
	return &newInstance
}

func (layout *Layout) GongCopy() GongstructIF {
	var newInstance Layout
	newInstance = *layout
	return &newInstance
}


// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
}
