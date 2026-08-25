package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceTransitionSemanticRules() (needCommit bool) {
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
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Transition \"%s\" has no Start state, removing it", transition.Name))
			}
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
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Transition \"%s\" has no End state, removing it", transition.Name))
			}
			continue
		}

		if stateMachine, ok := stager.map_state_stateMachine[transition.Start]; ok {
			if stateMachine.IsWithTransitionNameAutonamticalyGenerated {
				generatedName := transition.Start.Name + " to " + transition.End.Name
				if transition.Name != generatedName {
					transition.Name = generatedName
					needCommit = true
					if stager.probeForm != nil {
						stager.probeForm.AddNotification(time.Now(),
							fmt.Sprintf("Updated Transition name to \"%s\"", generatedName))
					}
				}
			}
		}
	}

	return
}
