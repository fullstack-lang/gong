// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Button
	for button := range stage.Buttons {
		_ = button
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Group
	for group := range stage.Groups {
		_ = group
		// insertion point per field
		var _Buttons []*Button
		for _, _button := range group.Buttons {
			if IsStaged(stage, _button) {
			 	_Buttons = append(_Buttons, _button)
			}
		}
		group.Buttons = _Buttons
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

}
