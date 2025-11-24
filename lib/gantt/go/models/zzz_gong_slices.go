// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Arrow
	// insertion point per field

	// Compute reverse map for named struct Bar
	// insertion point per field

	// Compute reverse map for named struct Gantt
	// insertion point per field
	clear(stage.Gantt_Lanes_reverseMap)
	stage.Gantt_Lanes_reverseMap = make(map[*Lane]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _lane := range gantt.Lanes {
			stage.Gantt_Lanes_reverseMap[_lane] = gantt
		}
	}
	clear(stage.Gantt_Milestones_reverseMap)
	stage.Gantt_Milestones_reverseMap = make(map[*Milestone]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _milestone := range gantt.Milestones {
			stage.Gantt_Milestones_reverseMap[_milestone] = gantt
		}
	}
	clear(stage.Gantt_Groups_reverseMap)
	stage.Gantt_Groups_reverseMap = make(map[*Group]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _group := range gantt.Groups {
			stage.Gantt_Groups_reverseMap[_group] = gantt
		}
	}
	clear(stage.Gantt_Arrows_reverseMap)
	stage.Gantt_Arrows_reverseMap = make(map[*Arrow]*Gantt)
	for gantt := range stage.Gantts {
		_ = gantt
		for _, _arrow := range gantt.Arrows {
			stage.Gantt_Arrows_reverseMap[_arrow] = gantt
		}
	}

	// Compute reverse map for named struct Group
	// insertion point per field
	clear(stage.Group_GroupLanes_reverseMap)
	stage.Group_GroupLanes_reverseMap = make(map[*Lane]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _lane := range group.GroupLanes {
			stage.Group_GroupLanes_reverseMap[_lane] = group
		}
	}

	// Compute reverse map for named struct Lane
	// insertion point per field
	clear(stage.Lane_Bars_reverseMap)
	stage.Lane_Bars_reverseMap = make(map[*Bar]*Lane)
	for lane := range stage.Lanes {
		_ = lane
		for _, _bar := range lane.Bars {
			stage.Lane_Bars_reverseMap[_bar] = lane
		}
	}

	// Compute reverse map for named struct LaneUse
	// insertion point per field

	// Compute reverse map for named struct Milestone
	// insertion point per field
	clear(stage.Milestone_LanesToDisplay_reverseMap)
	stage.Milestone_LanesToDisplay_reverseMap = make(map[*Lane]*Milestone)
	for milestone := range stage.Milestones {
		_ = milestone
		for _, _lane := range milestone.LanesToDisplay {
			stage.Milestone_LanesToDisplay_reverseMap[_lane] = milestone
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Arrows {
		res = append(res, instance)
	}

	for instance := range stage.Bars {
		res = append(res, instance)
	}

	for instance := range stage.Gantts {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Lanes {
		res = append(res, instance)
	}

	for instance := range stage.LaneUses {
		res = append(res, instance)
	}

	for instance := range stage.Milestones {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (arrow *Arrow) GongCopy() GongstructIF {
	var newInstance Arrow
	newInstance = *arrow
	return &newInstance
}

func (bar *Bar) GongCopy() GongstructIF {
	var newInstance Bar
	newInstance = *bar
	return &newInstance
}

func (gantt *Gantt) GongCopy() GongstructIF {
	var newInstance Gantt
	newInstance = *gantt
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	var newInstance Group
	newInstance = *group
	return &newInstance
}

func (lane *Lane) GongCopy() GongstructIF {
	var newInstance Lane
	newInstance = *lane
	return &newInstance
}

func (laneuse *LaneUse) GongCopy() GongstructIF {
	var newInstance LaneUse
	newInstance = *laneuse
	return &newInstance
}

func (milestone *Milestone) GongCopy() GongstructIF {
	var newInstance Milestone
	newInstance = *milestone
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
