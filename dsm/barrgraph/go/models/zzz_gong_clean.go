// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
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

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
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

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ArtefactTypeShape
func (artefacttypeshape *ArtefactTypeShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&artefacttypeshape.ArtefactType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Artist
func (artist *Artist) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&artist.Place) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArtistShape
func (artistshape *ArtistShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&artistshape.Artist) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Desk
func (desk *Desk) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&desk.SelectedDiagram) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagram.MovementShapes) || modified
	modified = stage.CleanSlice(&diagram.ArtefactTypeShapes) || modified
	modified = stage.CleanSlice(&diagram.ArtistShapes) || modified
	modified = stage.CleanSlice(&diagram.InfluenceShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Influence
func (influence *Influence) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&influence.SourceMovement) || modified
	modified = stage.CleanPointer(&influence.SourceArtefactType) || modified
	modified = stage.CleanPointer(&influence.SourceArtist) || modified
	modified = stage.CleanPointer(&influence.TargetMovement) || modified
	modified = stage.CleanPointer(&influence.TargetArtefactType) || modified
	modified = stage.CleanPointer(&influence.TargetArtist) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InfluenceShape
func (influenceshape *InfluenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&influenceshape.ControlPointShapes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&influenceshape.Influence) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Movement
func (movement *Movement) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&movement.Places) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MovementShape
func (movementshape *MovementShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&movementshape.Movement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
