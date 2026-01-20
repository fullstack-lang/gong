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
// Clean garbage collect unstaged instances that are referenced by AsSplit
func (assplit *AsSplit) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &assplit.AsSplitAreas)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AsSplitArea
func (assplitarea *AsSplitArea) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &assplitarea.AsSplit)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Button)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Cursor)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Form)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Load)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Markdown)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Slider)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Split)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Svg)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Table)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Tone)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Tree)  || modified
	modified = GongCleanPointer(stage, &assplitarea.Xlsx)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Button
func (button *Button) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Cursor
func (cursor *Cursor) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FavIcon
func (favicon *FavIcon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Form
func (form *Form) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Load
func (load *Load) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LogoOnTheLeft
func (logoontheleft *LogoOnTheLeft) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LogoOnTheRight
func (logoontheright *LogoOnTheRight) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Markdown
func (markdown *Markdown) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Slider
func (slider *Slider) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Split
func (split *Split) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Svg
func (svg *Svg) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Title
func (title *Title) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tone
func (tone *Tone) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Tree
func (tree *Tree) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by View
func (view *View) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &view.RootAsSplitAreas)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Xlsx
func (xlsx *Xlsx) GongClean(stage *Stage) (modified bool) {
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
