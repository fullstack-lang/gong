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
	// clean up AsSplit
	for assplit := range stage.AsSplits {
		_ = assplit
		// insertion point per field
		assplit.AsSplitAreas = GongCleanSlice(stage, assplit.AsSplitAreas)
		// insertion point per field
	}

	// clean up AsSplitArea
	for assplitarea := range stage.AsSplitAreas {
		_ = assplitarea
		// insertion point per field
		// insertion point per field
		assplitarea.AsSplit = GongCleanPointer(stage, assplitarea.AsSplit)
		assplitarea.Button = GongCleanPointer(stage, assplitarea.Button)
		assplitarea.Cursor = GongCleanPointer(stage, assplitarea.Cursor)
		assplitarea.Form = GongCleanPointer(stage, assplitarea.Form)
		assplitarea.Load = GongCleanPointer(stage, assplitarea.Load)
		assplitarea.Markdown = GongCleanPointer(stage, assplitarea.Markdown)
		assplitarea.Slider = GongCleanPointer(stage, assplitarea.Slider)
		assplitarea.Split = GongCleanPointer(stage, assplitarea.Split)
		assplitarea.Svg = GongCleanPointer(stage, assplitarea.Svg)
		assplitarea.Table = GongCleanPointer(stage, assplitarea.Table)
		assplitarea.Tone = GongCleanPointer(stage, assplitarea.Tone)
		assplitarea.Tree = GongCleanPointer(stage, assplitarea.Tree)
		assplitarea.Xlsx = GongCleanPointer(stage, assplitarea.Xlsx)
	}

	// clean up Button
	for button := range stage.Buttons {
		_ = button
		// insertion point per field
		// insertion point per field
	}

	// clean up Cursor
	for cursor := range stage.Cursors {
		_ = cursor
		// insertion point per field
		// insertion point per field
	}

	// clean up FavIcon
	for favicon := range stage.FavIcons {
		_ = favicon
		// insertion point per field
		// insertion point per field
	}

	// clean up Form
	for form := range stage.Forms {
		_ = form
		// insertion point per field
		// insertion point per field
	}

	// clean up Load
	for load := range stage.Loads {
		_ = load
		// insertion point per field
		// insertion point per field
	}

	// clean up LogoOnTheLeft
	for logoontheleft := range stage.LogoOnTheLefts {
		_ = logoontheleft
		// insertion point per field
		// insertion point per field
	}

	// clean up LogoOnTheRight
	for logoontheright := range stage.LogoOnTheRights {
		_ = logoontheright
		// insertion point per field
		// insertion point per field
	}

	// clean up Markdown
	for markdown := range stage.Markdowns {
		_ = markdown
		// insertion point per field
		// insertion point per field
	}

	// clean up Slider
	for slider := range stage.Sliders {
		_ = slider
		// insertion point per field
		// insertion point per field
	}

	// clean up Split
	for split := range stage.Splits {
		_ = split
		// insertion point per field
		// insertion point per field
	}

	// clean up Svg
	for svg := range stage.Svgs {
		_ = svg
		// insertion point per field
		// insertion point per field
	}

	// clean up Table
	for table := range stage.Tables {
		_ = table
		// insertion point per field
		// insertion point per field
	}

	// clean up Title
	for title := range stage.Titles {
		_ = title
		// insertion point per field
		// insertion point per field
	}

	// clean up Tone
	for tone := range stage.Tones {
		_ = tone
		// insertion point per field
		// insertion point per field
	}

	// clean up Tree
	for tree := range stage.Trees {
		_ = tree
		// insertion point per field
		// insertion point per field
	}

	// clean up View
	for view := range stage.Views {
		_ = view
		// insertion point per field
		view.RootAsSplitAreas = GongCleanSlice(stage, view.RootAsSplitAreas)
		// insertion point per field
	}

	// clean up Xlsx
	for xlsx := range stage.Xlsxs {
		_ = xlsx
		// insertion point per field
		// insertion point per field
	}

}
