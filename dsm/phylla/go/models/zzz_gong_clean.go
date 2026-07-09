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

// Clean garbage collect unstaged instances that are referenced by GrowthCurveRhombusGridShape
func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &growthcurverhombusgridshape.GrowthCurveRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GrowthCurveRhombusShape
func (growthcurverhombusshape *GrowthCurveRhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialRhombusGridShape
func (initialrhombusgridshape *InitialRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &initialrhombusgridshape.InitialRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by InitialRhombusShape
func (initialrhombusshape *InitialRhombusShape) GongClean(stage *Stage) (modified bool) {
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
	modified = GongCleanPointer(stage, &plant.AxesShape) || modified
	modified = GongCleanPointer(stage, &plant.ReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plant.PlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plant.GridPathShape) || modified
	modified = GongCleanPointer(stage, &plant.InitialRhombusGridShape) || modified
	modified = GongCleanPointer(stage, &plant.ExplanationTextShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedReferenceRhombus) || modified
	modified = GongCleanPointer(stage, &plant.RotatedPlantCircumferenceShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedGridPathShape) || modified
	modified = GongCleanPointer(stage, &plant.RotatedRhombusGridShape2) || modified
	modified = GongCleanPointer(stage, &plant.GrowthCurveRhombusGridShape) || modified
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
	return
}

// Clean garbage collect unstaged instances that are referenced by RhombusShape
func (rhombusshape *RhombusShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusGridShape
func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &rotatedrhombusgridshape.RotatedRhombusShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by RotatedRhombusShape
func (rotatedrhombusshape *RotatedRhombusShape) GongClean(stage *Stage) (modified bool) {
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
