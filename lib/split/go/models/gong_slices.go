// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
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

	// Compute reverse map for named struct Doc
	// insertion point per field

	// Compute reverse map for named struct Form
	// insertion point per field

	// Compute reverse map for named struct Load
	// insertion point per field

	// Compute reverse map for named struct Slider
	// insertion point per field

	// Compute reverse map for named struct Split
	// insertion point per field

	// Compute reverse map for named struct Svg
	// insertion point per field

	// Compute reverse map for named struct Table
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

}
