package models

import (
	"fmt"
	"log"
	"time"
)

func (stager *Stager) enforceAtLeastOneDiagramPerStateMachine() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per state machine
	for _, stateMachine := range GetGongstrucsSorted[*StateMachine](stage) {
		if len(stateMachine.Diagrams) == 0 {
			msg := fmt.Sprintf("State Machine \"%s\": each state machine has to have at least one diagram, adding one", stateMachine.Name)
			log.Println(msg)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), msg)
			}

			hasCheckedDiagram := false
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stage) {
				if diagram_.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}

			newDiagram := (&Diagram{
				Name:        "New Diagram",
				IsEditable_: true,
				IsChecked:   !hasCheckedDiagram,
			}).Stage(stage)
			stateMachine.Diagrams = append(stateMachine.Diagrams, newDiagram)

			needCommit = true
		}
	}

	return
}
