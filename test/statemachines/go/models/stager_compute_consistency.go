package models

import "log"

func (stager *Stager) computeConsistency() {

	// VERY important because the probe only unstages objects
	// this is the Clean that delete them from slices and pointers that reference
	// them. If the checkout is not performed, the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	stager.stage.Clean()
	needCommit := false

	// check that there is at least one architecture
	// and that one can safely access stager.architecture
	architectures := GetGongstrucsSorted[*Architecture](stager.stage)
	if len(architectures) == 0 {
		stager.architecture = (&Architecture{Name: "Architecture"}).Stage(stager.stage)
		needCommit = true
	} else {
		stager.architecture = architectures[0]
	}

	if stager.architecture.NbPixPerCharacter == 0 {
		stager.architecture.NbPixPerCharacter = 8
		needCommit = true
	}

	stager.map_state_nextStates = make(map[*State][]*State)
	rm := GetPointerReverseMap[Transition_Shape, Transition](GetAssociationName[Transition_Shape]().Transition.Name, stager.stage)

	for _, transition := range GetGongstrucsSorted[*Transition](stager.stage) {

		if transition.Start == nil {
			transition.Unstage(stager.stage)
			if transitionShapes, ok := rm[transition]; ok {
				for _, s := range transitionShapes {
					s.Unstage(stager.stage)
				}
			}
			needCommit = true
			log.Println("Transition", transition.Name, "has no Start state, removing it")
			continue
		}
		if transition.End == nil {
			transition.Unstage(stager.stage)
			if transitionShapes, ok := rm[transition]; ok {
				for _, s := range transitionShapes {
					s.Unstage(stager.stage)
				}
			}
			needCommit = true
			log.Println("Transition", transition.Name, "has no End state, removing it")
			continue
		}

		if _, ok := stager.map_state_nextStates[transition.Start]; !ok {
			stager.map_state_nextStates[transition.Start] = []*State{}
		}
		stager.map_state_nextStates[transition.Start] =
			append(stager.map_state_nextStates[transition.Start], transition.End)
	}

	stager.set_StartStates = make(map[*State]struct{})
	stager.map_stateMachine_objects = make(map[*StateMachine][]*Object)
	map_State_Objects := GetPointerReverseMap[Object, State](GetAssociationName[Object]().State.Name, stager.stage)
	for _, stateMachine := range GetGongstrucsSorted[*StateMachine](stager.stage) {

		stager.map_stateMachine_objects[stateMachine] = []*Object{}

		for _, state := range stateMachine.States {
			objects := map_State_Objects[state]
			stager.map_stateMachine_objects[stateMachine] = append(stager.map_stateMachine_objects[stateMachine], objects...)
		}

		if stateMachine.InitialState == nil {
			log.Println("State Machine", stateMachine.Name, "has no Start State")
			continue
		}

		stager.set_StartStates[stateMachine.InitialState] = struct{}{}

		nbOutgoinTransitions := len(stager.map_state_nextStates[stateMachine.InitialState])
		if nbOutgoinTransitions != 1 {
			log.Println("State Machine", stateMachine.Name, "Start State with outgoing transition diffrent of 1", nbOutgoinTransitions)
		}
	}

	for _, transition := range GetGongstrucsSorted[*Transition](stager.stage) {
		_ = transition
		// transition.Name = transition.Start.Name + " to " + transition.End.Name
		// needCommit = true
	}

	// object have a diagrams fields that is used in the XL export
	stager.map_state_stateMachine = make(map[*State]*StateMachine)
	map_State_StateMachines := GetSliceOfPointersReverseMap[StateMachine, State](GetAssociationName[StateMachine]().States[0].Name, stager.stage)
	for _, state := range GetGongstrucsSorted[*State](stager.stage) {

		if stateMachines, ok := map_State_StateMachines[state]; !ok {
			log.Println("Warn : no state machine for state", state.Name)
		} else {
			if len(stateMachines) > 1 {
				log.Println("Warn : more than 1 state machine for state", state.Name, "nb of state machines", len(stateMachines))
			} else {
				stager.map_state_stateMachine[state] = stateMachines[0]
			}
		}

		state.Diagrams = nil
		diagramSet := *GetGongstructInstancesSet[Diagram](stager.stage)
		for diagram := range diagramSet {
			for _, stateShape := range diagram.State_Shapes {
				if stateShape.State == state {
					state.Diagrams = append(state.Diagrams, diagram)
				}
			}
		}
	}

	stager.map_diagram_stateMachine = make(map[*Diagram]*StateMachine)
	map_diagram_StateMachines := GetSliceOfPointersReverseMap[StateMachine, Diagram](GetAssociationName[StateMachine]().Diagrams[0].Name, stager.stage)
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		if stateMachines, ok := map_diagram_StateMachines[diagram]; !ok {
			log.Println("Warn : no diagram machine for diagram", diagram.Name)
		} else {
			if len(stateMachines) > 1 {
				log.Println("Warn : more than 1 diagram machine for diagram", diagram.Name, "nb of diagram machines", len(stateMachines))
			} else {
				stager.map_diagram_stateMachine[diagram] = stateMachines[0]
			}
		}
	}

	transitionSet := *GetGongstructInstancesSet[Transition](stager.stage)
	for transition := range transitionSet {

		transition.Diagrams = nil
		diagramSet := *GetGongstructInstancesSet[Diagram](stager.stage)
		for diagram := range diagramSet {
			for _, transitionShape := range diagram.Transition_Shapes {
				if transitionShape.Transition == transition {
					transition.Diagrams = append(transition.Diagrams, diagram)
				}
			}
		}
	}

	{
		rm := GetSliceOfPointersReverseMap[Diagram, StateShape](GetAssociationName[Diagram]().State_Shapes[0].Name, stager.stage)
		for _, stateShape := range GetGongstrucsSorted[*StateShape](stager.stage) {
			if _, ok := rm[stateShape]; !ok {
				// state shape has no diagrams, it's an orphan
				stateShape.Unstage(stager.stage)
				needCommit = true
				continue
			}
			if stateShape.State == nil {
				stateShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			if stateShape.Name != stateShape.State.Name {
				stateShape.Name = stateShape.State.Name
				needCommit = true
			}

		}
	}

	{
		rm := GetSliceOfPointersReverseMap[Diagram, Transition_Shape](GetAssociationName[Diagram]().Transition_Shapes[0].Name, stager.stage)
		for _, transitionShape := range GetGongstrucsSorted[*Transition_Shape](stager.stage) {
			if _, ok := rm[transitionShape]; !ok {
				// transition shape has no diagrams, it's an orphan
				transitionShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			if transitionShape.Transition == nil {
				transitionShape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}

	if needCommit {
		stager.stage.Clean()
		stager.stage.CommitWithSuspendedCallbacks()
	}
}
