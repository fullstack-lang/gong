// generated code - do not edit
package models

// insertion point
func (inst *Arrow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Bar) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Gantt) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Group) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Lane) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *LaneUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Milestone) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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


// insertion point
func (inst *Arrow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				res = stage.Gantt_Arrows_reverseMap[inst]
			}
	}
	return res
}

func (inst *Bar) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				res = stage.Lane_Bars_reverseMap[inst]
			}
	}
	return res
}

func (inst *Gantt) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Group) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Gantt_Groups_reverseMap[inst]
			}
	}
	return res
}

func (inst *Lane) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Lanes":
				res = stage.Gantt_Lanes_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "GroupLanes":
				res = stage.Group_GroupLanes_reverseMap[inst]
			}
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplay":
				res = stage.Milestone_LanesToDisplay_reverseMap[inst]
			}
	}
	return res
}

func (inst *LaneUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Milestone) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				res = stage.Gantt_Milestones_reverseMap[inst]
			}
	}
	return res
}

