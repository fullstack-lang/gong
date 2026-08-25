package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceStateMachineSemanticRules() (needCommit bool) {
	for _, stateMachine := range GetGongstrucsSorted[*StateMachine](stager.stage) {
		if stateMachine.InitialState == nil {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("State Machine \"%s\" has no Start State", stateMachine.Name))
			}
			continue
		}

		nbOutgoingTransitions := len(stager.map_state_nextStates[stateMachine.InitialState])
		if nbOutgoingTransitions != 1 {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("State Machine \"%s\" Start State has %d outgoing transitions (expected 1)",
						stateMachine.Name, nbOutgoingTransitions))
			}
		}
	}

	return
}
