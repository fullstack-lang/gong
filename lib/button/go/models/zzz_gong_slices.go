// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct ButtonToggle
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Buttons_reverseMap = make(map[*Button]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _button := range group.Buttons {
			stage.Group_Buttons_reverseMap[_button] = group
		}
	}

	// Compute reverse map for named struct GroupToogle
	// insertion point per field
	stage.GroupToogle_ButtonToggles_reverseMap = make(map[*ButtonToggle]*GroupToogle)
	for grouptoogle := range stage.GroupToogles {
		_ = grouptoogle
		for _, _buttontoggle := range grouptoogle.ButtonToggles {
			stage.GroupToogle_ButtonToggles_reverseMap[_buttontoggle] = grouptoogle
		}
	}

	// Compute reverse map for named struct Layout
	// insertion point per field
	stage.Layout_Groups_reverseMap = make(map[*Group]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _group := range layout.Groups {
			stage.Layout_Groups_reverseMap[_group] = layout
		}
	}
	stage.Layout_GroupToogles_reverseMap = make(map[*GroupToogle]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _grouptoogle := range layout.GroupToogles {
			stage.Layout_GroupToogles_reverseMap[_grouptoogle] = layout
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.ButtonToggles {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.GroupToogles {
		res = append(res, instance)
	}

	for instance := range stage.Layouts {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := *button
	return &newInstance
}

func (buttontoggle *ButtonToggle) GongCopy() GongstructIF {
	newInstance := *buttontoggle
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := *group
	return &newInstance
}

func (grouptoogle *GroupToogle) GongCopy() GongstructIF {
	newInstance := *grouptoogle
	return &newInstance
}

func (layout *Layout) GongCopy() GongstructIF {
	newInstance := *layout
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
