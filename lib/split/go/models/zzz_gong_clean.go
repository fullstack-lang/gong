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
// Clean garbage collect unstaged instances that are referenced by AsSplit
func (assplit *AsSplit) GongClean(stage *Stage) {
	// insertion point per field
	assplit.AsSplitAreas = GongCleanSlice(stage, assplit.AsSplitAreas)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by AsSplitArea
func (assplitarea *AsSplitArea) GongClean(stage *Stage) {
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

// Clean garbage collect unstaged instances that are referenced by Button
func (button *Button) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Cursor
func (cursor *Cursor) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FavIcon
func (favicon *FavIcon) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Form
func (form *Form) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Load
func (load *Load) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by LogoOnTheLeft
func (logoontheleft *LogoOnTheLeft) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by LogoOnTheRight
func (logoontheright *LogoOnTheRight) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Markdown
func (markdown *Markdown) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Slider
func (slider *Slider) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Split
func (split *Split) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Svg
func (svg *Svg) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Title
func (title *Title) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Tone
func (tone *Tone) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Tree
func (tree *Tree) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by View
func (view *View) GongClean(stage *Stage) {
	// insertion point per field
	view.RootAsSplitAreas = GongCleanSlice(stage, view.RootAsSplitAreas)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Xlsx
func (xlsx *Xlsx) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
