// generated code - do not edit
package models

// insertion point
func (inst *Arrow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Gantt":
		switch reverseField.Fieldname {
		case "Arrows":
			if _gantt, ok := stage.Gantt_Arrows_reverseMap[inst]; ok {
				res = _gantt.Name
			}
		}
	}
	return
}

func (inst *Bar) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Lane":
		switch reverseField.Fieldname {
		case "Bars":
			if _lane, ok := stage.Lane_Bars_reverseMap[inst]; ok {
				res = _lane.Name
			}
		}
	}
	return
}

func (inst *Gantt) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Gantt":
		switch reverseField.Fieldname {
		case "Groups":
			if _gantt, ok := stage.Gantt_Groups_reverseMap[inst]; ok {
				res = _gantt.Name
			}
		}
	}
	return
}

func (inst *Lane) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Gantt":
		switch reverseField.Fieldname {
		case "Lanes":
			if _gantt, ok := stage.Gantt_Lanes_reverseMap[inst]; ok {
				res = _gantt.Name
			}
		}
	case "Group":
		switch reverseField.Fieldname {
		case "GroupLanes":
			if _group, ok := stage.Group_GroupLanes_reverseMap[inst]; ok {
				res = _group.Name
			}
		}
	case "Milestone":
		switch reverseField.Fieldname {
		case "LanesToDisplay":
			if _milestone, ok := stage.Milestone_LanesToDisplay_reverseMap[inst]; ok {
				res = _milestone.Name
			}
		}
	}
	return
}

func (inst *LaneUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Milestone) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Gantt":
		switch reverseField.Fieldname {
		case "Milestones":
			if _gantt, ok := stage.Gantt_Milestones_reverseMap[inst]; ok {
				res = _gantt.Name
			}
		}
	}
	return
}
