// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Command
	// insertion point per field

	// Compute reverse map for named struct DummyAgent
	// insertion point per field

	// Compute reverse map for named struct Engine
	// insertion point per field

	// Compute reverse map for named struct Event
	// insertion point per field

	// Compute reverse map for named struct Status
	// insertion point per field

	// Compute reverse map for named struct UpdateState
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Commands {
		res = append(res, instance)
	}

	for instance := range stage.DummyAgents {
		res = append(res, instance)
	}

	for instance := range stage.Engines {
		res = append(res, instance)
	}

	for instance := range stage.Events {
		res = append(res, instance)
	}

	for instance := range stage.Statuss {
		res = append(res, instance)
	}

	for instance := range stage.UpdateStates {
		res = append(res, instance)
	}


	return
}
