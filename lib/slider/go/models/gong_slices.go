// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Checkbox
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	clear(stage.Group_Sliders_reverseMap)
	stage.Group_Sliders_reverseMap = make(map[*Slider]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _slider := range group.Sliders {
			stage.Group_Sliders_reverseMap[_slider] = group
		}
	}
	clear(stage.Group_Checkboxes_reverseMap)
	stage.Group_Checkboxes_reverseMap = make(map[*Checkbox]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _checkbox := range group.Checkboxes {
			stage.Group_Checkboxes_reverseMap[_checkbox] = group
		}
	}

	// Compute reverse map for named struct Layout
	// insertion point per field
	clear(stage.Layout_Groups_reverseMap)
	stage.Layout_Groups_reverseMap = make(map[*Group]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _group := range layout.Groups {
			stage.Layout_Groups_reverseMap[_group] = layout
		}
	}

	// Compute reverse map for named struct Slider
	// insertion point per field

}
