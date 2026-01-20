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
// Clean garbage collect unstaged instances that are referenced by AttributeShape
func (attributeshape *AttributeShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Classdiagram
func (classdiagram *Classdiagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &classdiagram.GongStructShapes) || modified
	modified = GongCleanSlice(stage, &classdiagram.GongEnumShapes) || modified
	modified = GongCleanSlice(stage, &classdiagram.GongNoteShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramPackage
func (diagrampackage *DiagramPackage) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagrampackage.Classdiagrams) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &diagrampackage.SelectedClassdiagram) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GongEnumShape
func (gongenumshape *GongEnumShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongenumshape.GongEnumValueShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongEnumValueShape
func (gongenumvalueshape *GongEnumValueShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongNoteLinkShape
func (gongnotelinkshape *GongNoteLinkShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongNoteShape
func (gongnoteshape *GongNoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongnoteshape.GongNoteLinkShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongStructShape
func (gongstructshape *GongStructShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongstructshape.AttributeShapes) || modified
	modified = GongCleanSlice(stage, &gongstructshape.LinkShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LinkShape
func (linkshape *LinkShape) GongClean(stage *Stage) (modified bool) {
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
