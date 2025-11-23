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
	// clean up Button
	for button := range stage.Buttons {
		_ = button
		// insertion point per field
		// insertion point per field
		button.SVGIcon = GongCleanPointer(stage, button.SVGIcon)
	}

	// clean up Node
	for node := range stage.Nodes {
		_ = node
		// insertion point per field
		node.Children = GongCleanSlice(stage, node.Children)
		node.Buttons = GongCleanSlice(stage, node.Buttons)
		// insertion point per field
		node.PreceedingSVGIcon = GongCleanPointer(stage, node.PreceedingSVGIcon)
	}

	// clean up SVGIcon
	for svgicon := range stage.SVGIcons {
		_ = svgicon
		// insertion point per field
		// insertion point per field
	}

	// clean up Tree
	for tree := range stage.Trees {
		_ = tree
		// insertion point per field
		tree.RootNodes = GongCleanSlice(stage, tree.RootNodes)
		// insertion point per field
	}

}
