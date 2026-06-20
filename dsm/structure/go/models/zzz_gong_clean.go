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
// Clean garbage collect unstaged instances that are referenced by DiagramStructure
func (diagramstructure *DiagramStructure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagramstructure.Part_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.PartsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.Link_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.LinksWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootSystems) || modified
	modified = GongCleanSlice(stage, &library.SystemsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Link
func (link *Link) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &link.Source) || modified
	modified = GongCleanPointer(stage, &link.Target) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinkAssociationShape
func (linkassociationshape *LinkAssociationShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &linkassociationshape.Link) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Part
func (part *Part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartShape
func (partshape *PartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &partshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &system.Parts) || modified
	modified = GongCleanSlice(stage, &system.PartsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.SubSystems) || modified
	modified = GongCleanSlice(stage, &system.SubSystemsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.Links) || modified
	modified = GongCleanSlice(stage, &system.LinksWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.DiagramStructures) || modified
	modified = GongCleanSlice(stage, &system.DiagramStructuresWhoseNodeIsExpanded) || modified
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
