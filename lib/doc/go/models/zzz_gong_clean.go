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
// Clean garbage collect unstaged instances that are referenced by AttributeShape
func (attributeshape *AttributeShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Classdiagram
func (classdiagram *Classdiagram) GongClean(stage *Stage) {
	// insertion point per field
	classdiagram.GongStructShapes = GongCleanSlice(stage, classdiagram.GongStructShapes)
	classdiagram.GongEnumShapes = GongCleanSlice(stage, classdiagram.GongEnumShapes)
	classdiagram.GongNoteShapes = GongCleanSlice(stage, classdiagram.GongNoteShapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by DiagramPackage
func (diagrampackage *DiagramPackage) GongClean(stage *Stage) {
	// insertion point per field
	diagrampackage.Classdiagrams = GongCleanSlice(stage, diagrampackage.Classdiagrams)
	// insertion point per field
	diagrampackage.SelectedClassdiagram = GongCleanPointer(stage, diagrampackage.SelectedClassdiagram)
}

// Clean garbage collect unstaged instances that are referenced by GongEnumShape
func (gongenumshape *GongEnumShape) GongClean(stage *Stage) {
	// insertion point per field
	gongenumshape.GongEnumValueShapes = GongCleanSlice(stage, gongenumshape.GongEnumValueShapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongEnumValueShape
func (gongenumvalueshape *GongEnumValueShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongNoteLinkShape
func (gongnotelinkshape *GongNoteLinkShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongNoteShape
func (gongnoteshape *GongNoteShape) GongClean(stage *Stage) {
	// insertion point per field
	gongnoteshape.GongNoteLinkShapes = GongCleanSlice(stage, gongnoteshape.GongNoteLinkShapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongStructShape
func (gongstructshape *GongStructShape) GongClean(stage *Stage) {
	// insertion point per field
	gongstructshape.AttributeShapes = GongCleanSlice(stage, gongstructshape.AttributeShapes)
	gongstructshape.LinkShapes = GongCleanSlice(stage, gongstructshape.LinkShapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by LinkShape
func (linkshape *LinkShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
