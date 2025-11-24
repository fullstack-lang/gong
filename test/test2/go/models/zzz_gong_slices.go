// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct A
	// insertion point per field
	clear(stage.A_Bs_reverseMap)
	stage.A_Bs_reverseMap = make(map[*B]*A)
	for a := range stage.As {
		_ = a
		for _, _b := range a.Bs {
			stage.A_Bs_reverseMap[_b] = a
		}
	}

	// Compute reverse map for named struct B
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.As {
		res = append(res, instance)
	}

	for instance := range stage.Bs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (a *A) GongCopy() GongstructIF {
	var newInstance A
	newInstance = *a
	return &newInstance
}

func (b *B) GongCopy() GongstructIF {
	var newInstance B
	newInstance = *b
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
