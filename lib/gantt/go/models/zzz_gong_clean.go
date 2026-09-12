// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T PointerToGongstruct](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanSlice is a backward-compatible forwarder to stage.CleanSlice.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	return stage.CleanSlice(slice)
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T PointerToGongstruct](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

// GongCleanPointer is a backward-compatible forwarder to stage.CleanPointer.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	return stage.CleanPointer(element)
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
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
