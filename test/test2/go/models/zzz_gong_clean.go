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
		var _Bs []*B
		for _, _b := range a.Bs {
			if IsStaged(stage, _b) {
			 	_Bs = append(_Bs, _b)
			}
		}
		a.Bs = _Bs
		// insertion point per field
		if !IsStaged(stage, a.B) {
			a.B = nil
		}
	}

	// Compute reverse map for named struct B
	for b := range stage.Bs {
		_ = b
		// insertion point per field
		// insertion point per field
	}

}
