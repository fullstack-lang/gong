// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct AsSplit
	for assplit := range stage.AsSplits {
		_ = assplit
		// insertion point per field
		var _AsSplitAreas []*AsSplitArea
		for _, _assplitarea := range assplit.AsSplitAreas {
			if IsStaged(stage, _assplitarea) {
			 	_AsSplitAreas = append(_AsSplitAreas, _assplitarea)
			}
		}
		assplit.AsSplitAreas = _AsSplitAreas
		// insertion point per field
	}

	// Compute reverse map for named struct AsSplitArea
	for assplitarea := range stage.AsSplitAreas {
		_ = assplitarea
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, assplitarea.AsSplit) {
			assplitarea.AsSplit = nil
		}
		if !IsStaged(stage, assplitarea.Button) {
			assplitarea.Button = nil
		}
		if !IsStaged(stage, assplitarea.Cursor) {
			assplitarea.Cursor = nil
		}
		if !IsStaged(stage, assplitarea.Form) {
			assplitarea.Form = nil
		}
		if !IsStaged(stage, assplitarea.Load) {
			assplitarea.Load = nil
		}
		if !IsStaged(stage, assplitarea.Markdown) {
			assplitarea.Markdown = nil
		}
		if !IsStaged(stage, assplitarea.Slider) {
			assplitarea.Slider = nil
		}
		if !IsStaged(stage, assplitarea.Split) {
			assplitarea.Split = nil
		}
		if !IsStaged(stage, assplitarea.Svg) {
			assplitarea.Svg = nil
		}
		if !IsStaged(stage, assplitarea.Table) {
			assplitarea.Table = nil
		}
		if !IsStaged(stage, assplitarea.Tone) {
			assplitarea.Tone = nil
		}
		if !IsStaged(stage, assplitarea.Tree) {
			assplitarea.Tree = nil
		}
		if !IsStaged(stage, assplitarea.Xlsx) {
			assplitarea.Xlsx = nil
		}
	}

	// Compute reverse map for named struct Button
	for button := range stage.Buttons {
		_ = button
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Cursor
	for cursor := range stage.Cursors {
		_ = cursor
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FavIcon
	for favicon := range stage.FavIcons {
		_ = favicon
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Form
	for form := range stage.Forms {
		_ = form
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Load
	for load := range stage.Loads {
		_ = load
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct LogoOnTheLeft
	for logoontheleft := range stage.LogoOnTheLefts {
		_ = logoontheleft
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct LogoOnTheRight
	for logoontheright := range stage.LogoOnTheRights {
		_ = logoontheright
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Markdown
	for markdown := range stage.Markdowns {
		_ = markdown
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Slider
	for slider := range stage.Sliders {
		_ = slider
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Split
	for split := range stage.Splits {
		_ = split
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Svg
	for svg := range stage.Svgs {
		_ = svg
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Table
	for table := range stage.Tables {
		_ = table
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Title
	for title := range stage.Titles {
		_ = title
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Tone
	for tone := range stage.Tones {
		_ = tone
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Tree
	for tree := range stage.Trees {
		_ = tree
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct View
	for view := range stage.Views {
		_ = view
		// insertion point per field
		var _RootAsSplitAreas []*AsSplitArea
		for _, _assplitarea := range view.RootAsSplitAreas {
			if IsStaged(stage, _assplitarea) {
			 	_RootAsSplitAreas = append(_RootAsSplitAreas, _assplitarea)
			}
		}
		view.RootAsSplitAreas = _RootAsSplitAreas
		// insertion point per field
	}

	// Compute reverse map for named struct Xlsx
	for xlsx := range stage.Xlsxs {
		_ = xlsx
		// insertion point per field
		// insertion point per field
	}

}
