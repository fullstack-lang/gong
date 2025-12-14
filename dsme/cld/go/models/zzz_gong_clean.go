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
// Clean garbage collect unstaged instances that are referenced by Category1
func (category1 *Category1) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Category1Shape
func (category1shape *Category1Shape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	category1shape.Category1 = GongCleanPointer(stage, category1shape.Category1)
}

// Clean garbage collect unstaged instances that are referenced by Category2
func (category2 *Category2) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Category2Shape
func (category2shape *Category2Shape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	category2shape.Category2 = GongCleanPointer(stage, category2shape.Category2)
}

// Clean garbage collect unstaged instances that are referenced by Category3
func (category3 *Category3) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Category3Shape
func (category3shape *Category3Shape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	category3shape.Category3 = GongCleanPointer(stage, category3shape.Category3)
}

// Clean garbage collect unstaged instances that are referenced by ControlPointShape
func (controlpointshape *ControlPointShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Desk
func (desk *Desk) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	desk.SelectedDiagram = GongCleanPointer(stage, desk.SelectedDiagram)
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) {
	// insertion point per field
	diagram.Category1Shapes = GongCleanSlice(stage, diagram.Category1Shapes)
	diagram.Category2Shapes = GongCleanSlice(stage, diagram.Category2Shapes)
	diagram.Category3Shapes = GongCleanSlice(stage, diagram.Category3Shapes)
	diagram.InfluenceShapes = GongCleanSlice(stage, diagram.InfluenceShapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Influence
func (influence *Influence) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	influence.SourceCategory1 = GongCleanPointer(stage, influence.SourceCategory1)
	influence.SourceCategory2 = GongCleanPointer(stage, influence.SourceCategory2)
	influence.SourceCategory3 = GongCleanPointer(stage, influence.SourceCategory3)
	influence.TargetCategory1 = GongCleanPointer(stage, influence.TargetCategory1)
	influence.TargetCategory2 = GongCleanPointer(stage, influence.TargetCategory2)
	influence.TargetCategory3 = GongCleanPointer(stage, influence.TargetCategory3)
}

// Clean garbage collect unstaged instances that are referenced by InfluenceShape
func (influenceshape *InfluenceShape) GongClean(stage *Stage) {
	// insertion point per field
	influenceshape.ControlPointShapes = GongCleanSlice(stage, influenceshape.ControlPointShapes)
	// insertion point per field
	influenceshape.Influence = GongCleanPointer(stage, influenceshape.Influence)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
