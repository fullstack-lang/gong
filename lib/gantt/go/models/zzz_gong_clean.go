// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Arrow
	for arrow := range stage.Arrows {
		_ = arrow
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, arrow.From) {
			arrow.From = nil
		}
		if !IsStaged(stage, arrow.To) {
			arrow.To = nil
		}
	}

	// Compute reverse map for named struct Bar
	for bar := range stage.Bars {
		_ = bar
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Gantt
	for gantt := range stage.Gantts {
		_ = gantt
		// insertion point per field
		var _Lanes []*Lane
		for _, _lane := range gantt.Lanes {
			if IsStaged(stage, _lane) {
			 	_Lanes = append(_Lanes, _lane)
			}
		}
		gantt.Lanes = _Lanes
		var _Milestones []*Milestone
		for _, _milestone := range gantt.Milestones {
			if IsStaged(stage, _milestone) {
			 	_Milestones = append(_Milestones, _milestone)
			}
		}
		gantt.Milestones = _Milestones
		var _Groups []*Group
		for _, _group := range gantt.Groups {
			if IsStaged(stage, _group) {
			 	_Groups = append(_Groups, _group)
			}
		}
		gantt.Groups = _Groups
		var _Arrows []*Arrow
		for _, _arrow := range gantt.Arrows {
			if IsStaged(stage, _arrow) {
			 	_Arrows = append(_Arrows, _arrow)
			}
		}
		gantt.Arrows = _Arrows
		// insertion point per field
	}

	// Compute reverse map for named struct Group
	for group := range stage.Groups {
		_ = group
		// insertion point per field
		var _GroupLanes []*Lane
		for _, _lane := range group.GroupLanes {
			if IsStaged(stage, _lane) {
			 	_GroupLanes = append(_GroupLanes, _lane)
			}
		}
		group.GroupLanes = _GroupLanes
		// insertion point per field
	}

	// Compute reverse map for named struct Lane
	for lane := range stage.Lanes {
		_ = lane
		// insertion point per field
		var _Bars []*Bar
		for _, _bar := range lane.Bars {
			if IsStaged(stage, _bar) {
			 	_Bars = append(_Bars, _bar)
			}
		}
		lane.Bars = _Bars
		// insertion point per field
	}

	// Compute reverse map for named struct LaneUse
	for laneuse := range stage.LaneUses {
		_ = laneuse
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, laneuse.Lane) {
			laneuse.Lane = nil
		}
	}

	// Compute reverse map for named struct Milestone
	for milestone := range stage.Milestones {
		_ = milestone
		// insertion point per field
		var _LanesToDisplay []*Lane
		for _, _lane := range milestone.LanesToDisplay {
			if IsStaged(stage, _lane) {
			 	_LanesToDisplay = append(_LanesToDisplay, _lane)
			}
		}
		milestone.LanesToDisplay = _LanesToDisplay
		// insertion point per field
	}

}
