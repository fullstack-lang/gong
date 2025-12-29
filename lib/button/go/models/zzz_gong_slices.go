// generated code - do not edit
package models

import (
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Button "+button.Name,
				)
			}
		} else {
			diffs := button.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Button \""+button.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Button "+button.Name,
				)
			}
		}
	}

	lenNewInstances += len(buttons_newInstances)
	lenDeletedInstances += len(buttons_deletedInstances)
	var buttontoggles_newInstances []*ButtonToggle
	var buttontoggles_deletedInstances []*ButtonToggle

	// parse all staged instances and check if they have a reference
	for buttontoggle := range stage.ButtonToggles {
		if ref, ok := stage.ButtonToggles_reference[buttontoggle]; !ok {
			buttontoggles_newInstances = append(buttontoggles_newInstances, buttontoggle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of ButtonToggle "+buttontoggle.Name,
				)
			}
		} else {
			diffs := buttontoggle.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of ButtonToggle \""+buttontoggle.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for buttontoggle := range stage.ButtonToggles_reference {
		if _, ok := stage.ButtonToggles[buttontoggle]; !ok {
			buttontoggles_deletedInstances = append(buttontoggles_deletedInstances, buttontoggle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of ButtonToggle "+buttontoggle.Name,
				)
			}
		}
	}

	lenNewInstances += len(buttontoggles_newInstances)
	lenDeletedInstances += len(buttontoggles_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Group "+group.Name,
				)
			}
		} else {
			diffs := group.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Group \""+group.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for group := range stage.Groups_reference {
		if _, ok := stage.Groups[group]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, group)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Group "+group.Name,
				)
			}
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var grouptoogles_newInstances []*GroupToogle
	var grouptoogles_deletedInstances []*GroupToogle

	// parse all staged instances and check if they have a reference
	for grouptoogle := range stage.GroupToogles {
		if ref, ok := stage.GroupToogles_reference[grouptoogle]; !ok {
			grouptoogles_newInstances = append(grouptoogles_newInstances, grouptoogle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GroupToogle "+grouptoogle.Name,
				)
			}
		} else {
			diffs := grouptoogle.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GroupToogle \""+grouptoogle.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for grouptoogle := range stage.GroupToogles_reference {
		if _, ok := stage.GroupToogles[grouptoogle]; !ok {
			grouptoogles_deletedInstances = append(grouptoogles_deletedInstances, grouptoogle)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GroupToogle "+grouptoogle.Name,
				)
			}
		}
	}

	lenNewInstances += len(grouptoogles_newInstances)
	lenDeletedInstances += len(grouptoogles_deletedInstances)
	var layouts_newInstances []*Layout
	var layouts_deletedInstances []*Layout

	// parse all staged instances and check if they have a reference
	for layout := range stage.Layouts {
		if ref, ok := stage.Layouts_reference[layout]; !ok {
			layouts_newInstances = append(layouts_newInstances, layout)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Layout "+layout.Name,
				)
			}
		} else {
			diffs := layout.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Layout \""+layout.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for layout := range stage.Layouts_reference {
		if _, ok := stage.Layouts[layout]; !ok {
			layouts_deletedInstances = append(layouts_deletedInstances, layout)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Layout "+layout.Name,
				)
			}
		}
	}

	lenNewInstances += len(layouts_newInstances)
	lenDeletedInstances += len(layouts_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Buttons_reference = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
	}

	stage.ButtonToggles_reference = make(map[*ButtonToggle]*ButtonToggle)
	for instance := range stage.ButtonToggles {
		stage.ButtonToggles_reference[instance] = instance.GongCopy().(*ButtonToggle)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
	}

	stage.GroupToogles_reference = make(map[*GroupToogle]*GroupToogle)
	for instance := range stage.GroupToogles {
		stage.GroupToogles_reference[instance] = instance.GongCopy().(*GroupToogle)
	}

	stage.Layouts_reference = make(map[*Layout]*Layout)
	for instance := range stage.Layouts {
		stage.Layouts_reference[instance] = instance.GongCopy().(*Layout)
	}

}
