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
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
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
// Clean garbage collect unstaged instances that are referenced by Category1
func (category1 *Category1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Category1Shape
func (category1shape *Category1Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &category1shape.Category1)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Category2
func (category2 *Category2) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Category2Shape
func (category2shape *Category2Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &category2shape.Category2)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Category3
func (category3 *Category3) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Category3Shape
func (category3shape *Category3Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &category3shape.Category3)  || modified
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
	modified = GongCleanPointer(stage, &desk.SelectedDiagram)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagram.Category1Shapes)  || modified
	modified = GongCleanSlice(stage, &diagram.Category2Shapes)  || modified
	modified = GongCleanSlice(stage, &diagram.Category3Shapes)  || modified
	modified = GongCleanSlice(stage, &diagram.InfluenceShapes)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Influence
func (influence *Influence) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &influence.SourceCategory1)  || modified
	modified = GongCleanPointer(stage, &influence.SourceCategory2)  || modified
	modified = GongCleanPointer(stage, &influence.SourceCategory3)  || modified
	modified = GongCleanPointer(stage, &influence.TargetCategory1)  || modified
	modified = GongCleanPointer(stage, &influence.TargetCategory2)  || modified
	modified = GongCleanPointer(stage, &influence.TargetCategory3)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by InfluenceShape
func (influenceshape *InfluenceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &influenceshape.ControlPointShapes)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &influenceshape.Influence)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
