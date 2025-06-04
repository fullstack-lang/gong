// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Arrow:
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

	case *Bar:
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

	case *Gantt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Group:
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

	case *Lane:
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

	case *LaneUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Milestone:
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
	case *Arrow:
		switch reverseField.GongstructName {
		// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				res = stage.Gantt_Arrows_reverseMap[inst]
			}
		}

	case *Bar:
		switch reverseField.GongstructName {
		// insertion point
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				res = stage.Lane_Bars_reverseMap[inst]
			}
		}

	case *Gantt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Gantt_Groups_reverseMap[inst]
			}
		}

	case *Lane:
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

	case *LaneUse:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Milestone:
		switch reverseField.GongstructName {
		// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				res = stage.Gantt_Milestones_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
