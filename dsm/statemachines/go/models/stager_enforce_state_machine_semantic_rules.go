package models

import (
	"fmt"
	"log"
	"time"
)

func (stager *Stager) enforceStateMachineSemanticRules() (needCommit bool) {
	for _, stateMachine := range GetGongstrucsSorted[*StateMachine](stager.stage) {
		if stateMachine.InitialState == nil {
			msg := fmt.Sprintf("State Machine \"%s\" has no Start State", stateMachine.Name)
			log.Println(msg)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), msg)
			}
		} else {
			nbOutgoingTransitions := len(stager.map_state_nextStates[stateMachine.InitialState])
			if nbOutgoingTransitions != 1 {
				msg := fmt.Sprintf("State Machine \"%s\" Start State has %d outgoing transitions (expected 1)",
					stateMachine.Name, nbOutgoingTransitions)
				log.Println(msg)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), msg)
				}
			}
		}

		// Semantic rule: "state machine graph has to be one click"
		if len(stateMachine.States) > 0 {
			stateIndex := make(map[*State]int)
			for i, state := range stateMachine.States {
				stateIndex[state] = i
			}

			parent := make([]int, len(stateMachine.States))
			for i := range parent {
				parent[i] = i
			}

			var find func(int) int
			find = func(i int) int {
				if parent[i] != i {
					parent[i] = find(parent[i])
				}
				return parent[i]
			}

			union := func(i, j int) {
				rootI := find(i)
				rootJ := find(j)
				if rootI != rootJ {
					parent[rootI] = rootJ
				}
			}

			// Add edges from transitions between states of this state machine
			for _, transition := range GetGongstrucsSorted[*Transition](stager.stage) {
				if transition.Start != nil && transition.End != nil {
					idxStart, okStart := stateIndex[transition.Start]
					idxEnd, okEnd := stateIndex[transition.End]
					if okStart && okEnd {
						union(idxStart, idxEnd)
					}
				}
			}

			// Add edges from composite state hierarchy (SubStates)
			for _, state := range stateMachine.States {
				for _, subState := range state.SubStates {
					if subState != nil {
						idxState, okState := stateIndex[state]
						idxSub, okSub := stateIndex[subState]
						if okState && okSub {
							union(idxState, idxSub)
						}
					}
				}
			}

			roots := make(map[int]bool)
			for i := 0; i < len(stateMachine.States); i++ {
				roots[find(i)] = true
			}

			nbComponents := len(roots)
			if nbComponents != 1 {
				clickStr := "clicks"
				if nbComponents <= 1 {
					clickStr = "click"
				}
				msg := fmt.Sprintf("State Machine \"%s\": state machine graph has to be one click (computed %d %s)", stateMachine.Name, nbComponents, clickStr)
				log.Println(msg)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), msg)
				}
			}
		}
	}

	return
}
