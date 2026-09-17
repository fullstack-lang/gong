package models

import (
	"cmp"
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*Diagram)
	stager.map_state_nextStates = make(map[*State][]*State)
	stager.set_StartStates = make(map[*State]struct{})
	stager.map_stateMachine_objects = make(map[*StateMachine][]*Object)
	stager.map_state_stateMachine = make(map[*State]*StateMachine)
	stager.map_diagram_stateMachine = make(map[*Diagram]*StateMachine)

	// Populate transitions nextStates map
	for _, transition := range stager.stage.GetInstancesSorted[*Transition]() {
		if transition.Start != nil && transition.End != nil {
			stager.map_state_nextStates[transition.Start] = append(stager.map_state_nextStates[transition.Start], transition.End)
		}
	}

	// Map objects to state machines and start states
	map_State_Objects := stager.stage.GetPointerReverseMap[Object, State](GongGetAssociationName[Object]().State.Name)
	for _, stateMachine := range stager.stage.GetInstancesSorted[*StateMachine]() {
		stager.map_stateMachine_objects[stateMachine] = []*Object{}

		for _, state := range stateMachine.States {
			objects := map_State_Objects[state]
			stager.map_stateMachine_objects[stateMachine] = append(stager.map_stateMachine_objects[stateMachine], objects...)
		}

		if stateMachine.InitialState != nil {
			stager.set_StartStates[stateMachine.InitialState] = struct{}{}
		}
	}

	// Map states to their state machines & populate state.Diagrams
	map_State_StateMachines := stager.stage.GetSliceOfPointersReverseMap[StateMachine, State](GongGetAssociationName[StateMachine]().States[0].Name)
	for _, state := range stager.stage.GetInstancesSorted[*State]() {
		if stateMachines, ok := map_State_StateMachines[state]; ok && len(stateMachines) > 0 {
			stager.map_state_stateMachine[state] = stateMachines[0]
		}

		state.Diagrams = nil
		for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
			for _, stateShape := range diagram.State_Shapes {
				if stateShape.State == state {
					state.Diagrams = append(state.Diagrams, diagram)
					break
				}
			}
		}

		slices.SortFunc(state.Diagrams, func(a, b *Diagram) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}

	// Map diagrams to their state machines
	map_diagram_StateMachines := stager.stage.GetSliceOfPointersReverseMap[StateMachine, Diagram](GongGetAssociationName[StateMachine]().Diagrams[0].Name)
	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
		if stateMachines, ok := map_diagram_StateMachines[diagram]; ok && len(stateMachines) > 0 {
			stager.map_diagram_stateMachine[diagram] = stateMachines[0]
		}
	}

	// Populate transition.Diagrams
	for _, transition := range stager.stage.GetInstancesSorted[*Transition]() {
		transition.Diagrams = nil
		for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
			for _, transitionShape := range diagram.Transition_Shapes {
				if transitionShape.Transition == transition {
					transition.Diagrams = append(transition.Diagrams, diagram)
					break
				}
			}
		}

		slices.SortFunc(transition.Diagrams, func(a, b *Diagram) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}

	// Populate map_Element_Diagrams for abstract elements (StateMachine, Note)
	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
		if sm, ok := stager.map_diagram_stateMachine[diagram]; ok {
			stager.addDiagramToElement(sm, diagram)
		}
		for _, noteShape := range diagram.Note_Shapes {
			if noteShape.Note != nil {
				stager.addDiagramToElement(noteShape.Note, diagram)
			}
		}
	}
}

func (stager *Stager) addDiagramToElement(element AbstractType, diagram *Diagram) {
	diagrams := stager.map_Element_Diagrams[element]
	if !slices.Contains(diagrams, diagram) {
		diagrams = append(diagrams, diagram)
		stager.map_Element_Diagrams[element] = diagrams
	}
}
