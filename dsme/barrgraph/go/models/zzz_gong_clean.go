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
	if modified {
		*slice = cleanedSlice
	}
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
// Clean garbage collect unstaged instances that are referenced by ArtefactType
func (artefacttype *ArtefactType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArtefactTypeShape
func (artefacttypeshape *ArtefactTypeShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &artefacttypeshape.ArtefactType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Artist
func (artist *Artist) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &artist.Place) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArtistShape
func (artistshape *ArtistShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &artistshape.Artist) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlPointShape
func (controlpointshape *ControlPointShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Desk
func (desk *Desk) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &desk.SelectedDiagram) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagram.MovementShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ArtefactTypeShapes) || modified
	modified = GongCleanSlice(stage, &diagram.ArtistShapes) || modified
	modified = GongCleanSlice(stage, &diagram.InfluenceShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Influence
func (influence *Influence) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &influence.SourceMovement) || modified
	modified = GongCleanPointer(stage, &influence.SourceArtefactType) || modified
	modified = GongCleanPointer(stage, &influence.SourceArtist) || modified
	modified = GongCleanPointer(stage, &influence.TargetMovement) || modified
	modified = GongCleanPointer(stage, &influence.TargetArtefactType) || modified
	modified = GongCleanPointer(stage, &influence.TargetArtist) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InfluenceShape
func (influenceshape *InfluenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &influenceshape.ControlPointShapes) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &influenceshape.Influence) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Movement
func (movement *Movement) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &movement.Places) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MovementShape
func (movementshape *MovementShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &movementshape.Movement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Place
func (place *Place) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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
