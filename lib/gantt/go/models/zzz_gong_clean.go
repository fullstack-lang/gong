// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	*slice = cleanedSlice
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Arrow
func (arrow *Arrow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &arrow.From) || modified
	modified = GongCleanPointer(stage, &arrow.To) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bar
func (bar *Bar) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Gantt
func (gantt *Gantt) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gantt.Lanes) || modified
	modified = GongCleanSlice(stage, &gantt.Milestones) || modified
	modified = GongCleanSlice(stage, &gantt.Groups) || modified
	modified = GongCleanSlice(stage, &gantt.Arrows) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Group
func (group *Group) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &group.GroupLanes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Lane
func (lane *Lane) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &lane.Bars) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LaneUse
func (laneuse *LaneUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &laneuse.Lane) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Milestone
func (milestone *Milestone) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &milestone.LanesToDisplay) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
