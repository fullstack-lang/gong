// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Checkbox
	for checkbox := range stage.Checkboxs {
		_ = checkbox
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Group
	for group := range stage.Groups {
		_ = group
		// insertion point per field
		var _Sliders []*Slider
		for _, _slider := range group.Sliders {
			if IsStaged(stage, _slider) {
			 	_Sliders = append(_Sliders, _slider)
			}
		}
		group.Sliders = _Sliders
		var _Checkboxes []*Checkbox
		for _, _checkbox := range group.Checkboxes {
			if IsStaged(stage, _checkbox) {
			 	_Checkboxes = append(_Checkboxes, _checkbox)
			}
		}
		group.Checkboxes = _Checkboxes
		// insertion point per field
	}

	// Compute reverse map for named struct Layout
	for layout := range stage.Layouts {
		_ = layout
		// insertion point per field
		var _Groups []*Group
		for _, _group := range layout.Groups {
			if IsStaged(stage, _group) {
			 	_Groups = append(_Groups, _group)
			}
		}
		layout.Groups = _Groups
		// insertion point per field
	}

	// Compute reverse map for named struct Slider
	for slider := range stage.Sliders {
		_ = slider
		// insertion point per field
		// insertion point per field
	}

}
