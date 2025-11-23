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
	// clean up Checkbox
	for checkbox := range stage.Checkboxs {
		_ = checkbox
		// insertion point per field
		// insertion point per field
	}

	// clean up Group
	for group := range stage.Groups {
		_ = group
		// insertion point per field
		group.Sliders = GongCleanSlice(stage, group.Sliders)
		group.Checkboxes = GongCleanSlice(stage, group.Checkboxes)
		// insertion point per field
	}

	// clean up Layout
	for layout := range stage.Layouts {
		_ = layout
		// insertion point per field
		layout.Groups = GongCleanSlice(stage, layout.Groups)
		// insertion point per field
	}

	// clean up Slider
	for slider := range stage.Sliders {
		_ = slider
		// insertion point per field
		// insertion point per field
	}

}
