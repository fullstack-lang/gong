package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceControlFlowRules() (needCommit bool) {

	for controlFlow := range *GetGongstructInstancesSetFromPointerType[*ControlFlow](stager.stage) {
		if controlFlow.Start == nil {
			controlFlow.Unstage(stager.stage)

			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control \"%s\" has no start, unstaging",
				controlFlow.GetName()))
			needCommit = true
			continue
		}
		if controlFlow.End == nil {
			controlFlow.Unstage(stager.stage)

			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control \"%s\" has no end, unstaging",
				controlFlow.GetName()))
			needCommit = true
			continue
		}
		if controlFlow.Start.owningParticipant != controlFlow.End.owningParticipant {
			controlFlow.Unstage(stager.stage)

			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control flow \"%s\" connects tasks from different participants, unstaging", controlFlow.GetName()))
			needCommit = true
			continue
		}
	}

	return
}
