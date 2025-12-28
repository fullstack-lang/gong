// generated code - do not edit
package models

import "time"

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct A
	// insertion point per field
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
	newInstance := *a
	return &newInstance
}

func (b *B) GongCopy() GongstructIF {
	newInstance := *b
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeDifference() {
	var lenNewInstances int

	var As_newInstances []*A
	// parse all staged instances and check if they have a reference
	for instance := range stage.As {
		if _, ok := stage.As_reference[instance]; ok {

		} else {
			As_newInstances = append(As_newInstances, instance)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Instance \""+instance.GetName()+"\" added. "+"Type; "+instance.GongGetGongstructName(),
				)
			}
		}
	}
	lenNewInstances += len(As_newInstances)

	var Bs_newInstances []*B
	// parse all staged instances and check if they have a reference
	for instance := range stage.Bs {
		if _, ok := stage.Bs_reference[instance]; ok {

		} else {
			Bs_newInstances = append(Bs_newInstances, instance)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Instance \""+instance.GetName()+"\" added. "+"Type; "+instance.GongGetGongstructName(),
				)
			}
		}
	}
	lenNewInstances += len(Bs_newInstances)

	if lenNewInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

func (stage *Stage) ComputeReference() {
	stage.As_reference = make(map[*A]*A)
	for instance := range stage.As {
		stage.As_reference[instance] = instance
	}

	stage.Bs_reference = make(map[*B]*B)
	for instance := range stage.Bs {
		stage.Bs_reference[instance] = instance
	}
}
