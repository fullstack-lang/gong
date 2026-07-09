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
// Clean garbage collect unstaged instances that are referenced by AxesShape
func (axesshape *AxesShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleGridShape
func (circlegridshape *CircleGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExplanationTextShape
func (explanationtextshape *ExplanationTextShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GridPathShape
func (gridpathshape *GridPathShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.Plants) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NextCircleShape
func (nextcircleshape *NextCircleShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Plant
func (plant *Plant) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &plant.PlantDiagrams) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PlantCircumferenceShape
func (plantcircumferenceshape *PlantCircumferenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PlantDiagram
func (plantdiagram *PlantDiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &plantdiagram.AxesShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.ReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plantdiagram.PlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.GridPathShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.RhombusGridShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.ExplanationTextShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.RotatedReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plantdiagram.RotatedPlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.RotatedGridPathShape) || modified
	modified = GongCleanPointer(stage, &plantdiagram.RotatedRhombusGridShape) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusGridShape
func (rhombusgridshape *RhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &rhombusgridshape.RhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusShape
func (rhombusshape *RhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
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
