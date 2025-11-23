// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct A
	for a := range stage.As {
		_ = a
		// insertion point per field
		var _As []*A
		for _, _a := range a.As {
			if IsStaged(stage, _a) {
			 	_As = append(_As, _a)
			}
		}
		a.As = _As
		// insertion point per field
	}

}
