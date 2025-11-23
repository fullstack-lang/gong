// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}
    
	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up Arrow
	for arrow := range stage.Arrows {
		_ = arrow
		// insertion point per field
		// insertion point per field
		arrow.From = GongCleanPointer(stage, arrow.From)
		arrow.To = GongCleanPointer(stage, arrow.To)
	}

	// clean up Bar
	for bar := range stage.Bars {
		_ = bar
		// insertion point per field
		// insertion point per field
	}

	// clean up Gantt
	for gantt := range stage.Gantts {
		_ = gantt
		// insertion point per field
		gantt.Lanes = GongCleanSlice(stage, gantt.Lanes)
		gantt.Milestones = GongCleanSlice(stage, gantt.Milestones)
		gantt.Groups = GongCleanSlice(stage, gantt.Groups)
		gantt.Arrows = GongCleanSlice(stage, gantt.Arrows)
		// insertion point per field
	}

	// clean up Group
	for group := range stage.Groups {
		_ = group
		// insertion point per field
		group.GroupLanes = GongCleanSlice(stage, group.GroupLanes)
		// insertion point per field
	}

	// clean up Lane
	for lane := range stage.Lanes {
		_ = lane
		// insertion point per field
		lane.Bars = GongCleanSlice(stage, lane.Bars)
		// insertion point per field
	}

	// clean up LaneUse
	for laneuse := range stage.LaneUses {
		_ = laneuse
		// insertion point per field
		// insertion point per field
		laneuse.Lane = GongCleanPointer(stage, laneuse.Lane)
	}

	// clean up Milestone
	for milestone := range stage.Milestones {
		_ = milestone
		// insertion point per field
		milestone.LanesToDisplay = GongCleanSlice(stage, milestone.LanesToDisplay)
		// insertion point per field
	}

}
