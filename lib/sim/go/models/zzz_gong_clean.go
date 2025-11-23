// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Command
	for command := range stage.Commands {
		_ = command
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, command.Engine) {
			command.Engine = nil
		}
	}

	// Compute reverse map for named struct DummyAgent
	for dummyagent := range stage.DummyAgents {
		_ = dummyagent
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Engine
	for engine := range stage.Engines {
		_ = engine
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Event
	for event := range stage.Events {
		_ = event
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Status
	for status := range stage.Statuss {
		_ = status
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct UpdateState
	for updatestate := range stage.UpdateStates {
		_ = updatestate
		// insertion point per field
		// insertion point per field
	}

}
