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

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Arrow
func (arrow *Arrow) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	arrow.From = GongCleanPointer(stage, arrow.From)
	arrow.To = GongCleanPointer(stage, arrow.To)
}

// Clean garbage collect unstaged instances that are referenced by Bar
func (bar *Bar) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Gantt
func (gantt *Gantt) GongClean(stage *Stage) {
	// insertion point per field
	gantt.Lanes = GongCleanSlice(stage, gantt.Lanes)
	gantt.Milestones = GongCleanSlice(stage, gantt.Milestones)
	gantt.Groups = GongCleanSlice(stage, gantt.Groups)
	gantt.Arrows = GongCleanSlice(stage, gantt.Arrows)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) {
	// insertion point per field
	group.GroupLanes = GongCleanSlice(stage, group.GroupLanes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Lane
func (lane *Lane) GongClean(stage *Stage) {
	// insertion point per field
	lane.Bars = GongCleanSlice(stage, lane.Bars)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by LaneUse
func (laneuse *LaneUse) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	laneuse.Lane = GongCleanPointer(stage, laneuse.Lane)
}

// Clean garbage collect unstaged instances that are referenced by Milestone
func (milestone *Milestone) GongClean(stage *Stage) {
	// insertion point per field
	milestone.LanesToDisplay = GongCleanSlice(stage, milestone.LanesToDisplay)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
