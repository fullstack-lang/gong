// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct A
	// insertion point per field
	clear(stage.A_As_reverseMap)
	stage.A_As_reverseMap = make(map[*A]*A)
	for a := range stage.As {
		_ = a
		for _, _a := range a.As {
			stage.A_As_reverseMap[_a] = a
		}
	}

}
