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

// insertion point per named struct
func (command *Command) GongCopy() GongstructIF {
	newInstance := *command
	return &newInstance
}

func (dummyagent *DummyAgent) GongCopy() GongstructIF {
	newInstance := *dummyagent
	return &newInstance
}

func (engine *Engine) GongCopy() GongstructIF {
	newInstance := *engine
	return &newInstance
}

func (event *Event) GongCopy() GongstructIF {
	newInstance := *event
	return &newInstance
}

func (status *Status) GongCopy() GongstructIF {
	newInstance := *status
	return &newInstance
}

func (updatestate *UpdateState) GongCopy() GongstructIF {
	newInstance := *updatestate
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
