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
		}
		if controlFlow.End == nil {
			controlFlow.Unstage(stager.stage)

			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Control \"%s\" has no end, unstaging",
				controlFlow.GetName()))
		}
	}

	return
}
