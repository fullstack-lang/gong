// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}
    
	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up Command
	for command := range stage.Commands {
		_ = command
		// insertion point per field
		// insertion point per field
		command.Engine = GongCleanPointer(stage, command.Engine)
	}

	// clean up DummyAgent
	for dummyagent := range stage.DummyAgents {
		_ = dummyagent
		// insertion point per field
		// insertion point per field
	}

	// clean up Engine
	for engine := range stage.Engines {
		_ = engine
		// insertion point per field
		// insertion point per field
	}

	// clean up Event
	for event := range stage.Events {
		_ = event
		// insertion point per field
		// insertion point per field
	}

	// clean up Status
	for status := range stage.Statuss {
		_ = status
		// insertion point per field
		// insertion point per field
	}

	// clean up UpdateState
	for updatestate := range stage.UpdateStates {
		_ = updatestate
		// insertion point per field
		// insertion point per field
	}

}
