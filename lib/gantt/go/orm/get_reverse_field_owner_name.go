// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/gantt/go/models"
)

func GetReverseFieldOwnerName(
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance any,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
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

	case *models.Bar:
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

	case *models.Gantt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group:
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

	case *models.Lane:
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
		}

	case *models.LaneUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				if _milestone, ok := stage.Milestone_LanesToDisplayMilestoneUse_reverseMap[inst]; ok {
					res = _milestone.Name
				}
			}
		}

	case *models.Milestone:
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

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		switch reverseField.GongstructName {
		// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				res = stage.Gantt_Arrows_reverseMap[inst]
			}
		}

	case *models.Bar:
		switch reverseField.GongstructName {
		// insertion point
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				res = stage.Lane_Bars_reverseMap[inst]
			}
		}

	case *models.Gantt:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group:
		switch reverseField.GongstructName {
		// insertion point
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Gantt_Groups_reverseMap[inst]
			}
		}

	case *models.Lane:
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
		}

	case *models.LaneUse:
		switch reverseField.GongstructName {
		// insertion point
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				res = stage.Milestone_LanesToDisplayMilestoneUse_reverseMap[inst]
			}
		}

	case *models.Milestone:
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
