// generated code - do not edit
package models

import "time"

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
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Complexity
func (complexity *Complexity) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramFloss
func (diagramfloss *DiagramFloss) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagramfloss.System_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramfloss.SystemsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Effort
func (effort *Effort) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootSystemes) || modified
	modified = GongCleanSlice(stage, &library.SystemsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootComplexitys) || modified
	modified = GongCleanSlice(stage, &library.ComplexitysWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootPerformances) || modified
	modified = GongCleanSlice(stage, &library.PerformancesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootEfforts) || modified
	modified = GongCleanSlice(stage, &library.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Performance
func (performance *Performance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &system.Complexitys) || modified
	modified = GongCleanSlice(stage, &system.Performances) || modified
	modified = GongCleanSlice(stage, &system.Efforts) || modified
	modified = GongCleanSlice(stage, &system.DiagramFlosses) || modified
	modified = GongCleanSlice(stage, &system.DiagramFlossWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.SubSystemes) || modified
	modified = GongCleanSlice(stage, &system.ComplexitysWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.PerformancesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SystemShape
func (systemshape *SystemShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &systemshape.System) || modified
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
