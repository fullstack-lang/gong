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

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up AttributeShape
	for attributeshape := range stage.AttributeShapes {
		_ = attributeshape
		// insertion point per field
		// insertion point per field
	}

	// clean up Classdiagram
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		// insertion point per field
		classdiagram.GongStructShapes = GongCleanSlice(stage, classdiagram.GongStructShapes)
		classdiagram.GongEnumShapes = GongCleanSlice(stage, classdiagram.GongEnumShapes)
		classdiagram.GongNoteShapes = GongCleanSlice(stage, classdiagram.GongNoteShapes)
		// insertion point per field
	}

	// clean up DiagramPackage
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		// insertion point per field
		diagrampackage.Classdiagrams = GongCleanSlice(stage, diagrampackage.Classdiagrams)
		// insertion point per field
		diagrampackage.SelectedClassdiagram = GongCleanPointer(stage, diagrampackage.SelectedClassdiagram)
	}

	// clean up GongEnumShape
	for gongenumshape := range stage.GongEnumShapes {
		_ = gongenumshape
		// insertion point per field
		gongenumshape.GongEnumValueShapes = GongCleanSlice(stage, gongenumshape.GongEnumValueShapes)
		// insertion point per field
	}

	// clean up GongEnumValueShape
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		_ = gongenumvalueshape
		// insertion point per field
		// insertion point per field
	}

	// clean up GongNoteLinkShape
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		_ = gongnotelinkshape
		// insertion point per field
		// insertion point per field
	}

	// clean up GongNoteShape
	for gongnoteshape := range stage.GongNoteShapes {
		_ = gongnoteshape
		// insertion point per field
		gongnoteshape.GongNoteLinkShapes = GongCleanSlice(stage, gongnoteshape.GongNoteLinkShapes)
		// insertion point per field
	}

	// clean up GongStructShape
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		// insertion point per field
		gongstructshape.AttributeShapes = GongCleanSlice(stage, gongstructshape.AttributeShapes)
		gongstructshape.LinkShapes = GongCleanSlice(stage, gongstructshape.LinkShapes)
		// insertion point per field
	}

	// clean up LinkShape
	for linkshape := range stage.LinkShapes {
		_ = linkshape
		// insertion point per field
		// insertion point per field
	}

}
