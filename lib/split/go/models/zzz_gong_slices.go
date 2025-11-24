// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AsSplit
	// insertion point per field
	clear(stage.AsSplit_AsSplitAreas_reverseMap)
	stage.AsSplit_AsSplitAreas_reverseMap = make(map[*AsSplitArea]*AsSplit)
	for assplit := range stage.AsSplits {
		_ = assplit
		for _, _assplitarea := range assplit.AsSplitAreas {
			stage.AsSplit_AsSplitAreas_reverseMap[_assplitarea] = assplit
		}
	}

	// Compute reverse map for named struct AsSplitArea
	// insertion point per field

	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Cursor
	// insertion point per field

	// Compute reverse map for named struct FavIcon
	// insertion point per field

	// Compute reverse map for named struct Form
	// insertion point per field

	// Compute reverse map for named struct Load
	// insertion point per field

	// Compute reverse map for named struct LogoOnTheLeft
	// insertion point per field

	// Compute reverse map for named struct LogoOnTheRight
	// insertion point per field

	// Compute reverse map for named struct Markdown
	// insertion point per field

	// Compute reverse map for named struct Slider
	// insertion point per field

	// Compute reverse map for named struct Split
	// insertion point per field

	// Compute reverse map for named struct Svg
	// insertion point per field

	// Compute reverse map for named struct Table
	// insertion point per field

	// Compute reverse map for named struct Title
	// insertion point per field

	// Compute reverse map for named struct Tone
	// insertion point per field

	// Compute reverse map for named struct Tree
	// insertion point per field

	// Compute reverse map for named struct View
	// insertion point per field
	clear(stage.View_RootAsSplitAreas_reverseMap)
	stage.View_RootAsSplitAreas_reverseMap = make(map[*AsSplitArea]*View)
	for view := range stage.Views {
		_ = view
		for _, _assplitarea := range view.RootAsSplitAreas {
			stage.View_RootAsSplitAreas_reverseMap[_assplitarea] = view
		}
	}

	// Compute reverse map for named struct Xlsx
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.AsSplits {
		res = append(res, instance)
	}

	for instance := range stage.AsSplitAreas {
		res = append(res, instance)
	}

	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Cursors {
		res = append(res, instance)
	}

	for instance := range stage.FavIcons {
		res = append(res, instance)
	}

	for instance := range stage.Forms {
		res = append(res, instance)
	}

	for instance := range stage.Loads {
		res = append(res, instance)
	}

	for instance := range stage.LogoOnTheLefts {
		res = append(res, instance)
	}

	for instance := range stage.LogoOnTheRights {
		res = append(res, instance)
	}

	for instance := range stage.Markdowns {
		res = append(res, instance)
	}

	for instance := range stage.Sliders {
		res = append(res, instance)
	}

	for instance := range stage.Splits {
		res = append(res, instance)
	}

	for instance := range stage.Svgs {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

	for instance := range stage.Titles {
		res = append(res, instance)
	}

	for instance := range stage.Tones {
		res = append(res, instance)
	}

	for instance := range stage.Trees {
		res = append(res, instance)
	}

	for instance := range stage.Views {
		res = append(res, instance)
	}

	for instance := range stage.Xlsxs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (assplit *AsSplit) GongCopy() GongstructIF {
	var newInstance AsSplit
	newInstance = *assplit
	return &newInstance
}

func (assplitarea *AsSplitArea) GongCopy() GongstructIF {
	var newInstance AsSplitArea
	newInstance = *assplitarea
	return &newInstance
}

func (button *Button) GongCopy() GongstructIF {
	var newInstance Button
	newInstance = *button
	return &newInstance
}

func (cursor *Cursor) GongCopy() GongstructIF {
	var newInstance Cursor
	newInstance = *cursor
	return &newInstance
}

func (favicon *FavIcon) GongCopy() GongstructIF {
	var newInstance FavIcon
	newInstance = *favicon
	return &newInstance
}

func (form *Form) GongCopy() GongstructIF {
	var newInstance Form
	newInstance = *form
	return &newInstance
}

func (load *Load) GongCopy() GongstructIF {
	var newInstance Load
	newInstance = *load
	return &newInstance
}

func (logoontheleft *LogoOnTheLeft) GongCopy() GongstructIF {
	var newInstance LogoOnTheLeft
	newInstance = *logoontheleft
	return &newInstance
}

func (logoontheright *LogoOnTheRight) GongCopy() GongstructIF {
	var newInstance LogoOnTheRight
	newInstance = *logoontheright
	return &newInstance
}

func (markdown *Markdown) GongCopy() GongstructIF {
	var newInstance Markdown
	newInstance = *markdown
	return &newInstance
}

func (slider *Slider) GongCopy() GongstructIF {
	var newInstance Slider
	newInstance = *slider
	return &newInstance
}

func (split *Split) GongCopy() GongstructIF {
	var newInstance Split
	newInstance = *split
	return &newInstance
}

func (svg *Svg) GongCopy() GongstructIF {
	var newInstance Svg
	newInstance = *svg
	return &newInstance
}

func (table *Table) GongCopy() GongstructIF {
	var newInstance Table
	newInstance = *table
	return &newInstance
}

func (title *Title) GongCopy() GongstructIF {
	var newInstance Title
	newInstance = *title
	return &newInstance
}

func (tone *Tone) GongCopy() GongstructIF {
	var newInstance Tone
	newInstance = *tone
	return &newInstance
}

func (tree *Tree) GongCopy() GongstructIF {
	var newInstance Tree
	newInstance = *tree
	return &newInstance
}

func (view *View) GongCopy() GongstructIF {
	var newInstance View
	newInstance = *view
	return &newInstance
}

func (xlsx *Xlsx) GongCopy() GongstructIF {
	var newInstance Xlsx
	newInstance = *xlsx
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
