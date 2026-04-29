package models

import (
	"fmt"
	"time"
)

// enforceTaskSemanticRules checks
// - task cannot be both start and end
// - per participant, cannot have more than one start task
func (stager *Stager) enforceTaskSemanticRules() (needCommit bool) {

	for task := range *GetGongstructInstancesSetFromPointerType[*Task](stager.stage) {

		if task.IsStartTask && task.IsEndTask {
			needCommit = true
			task.IsEndTask = false
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" cannot be both start and end",
				task.GetName()))
		}
	}

	for participant := range *GetGongstructInstancesSetFromPointerType[*Participant](stager.stage) {
		nbStartTask := 0
		for _, task := range participant.Tasks {
			if nbStartTask == 1 && task.IsStartTask {
				needCommit = true
				task.IsStartTask = false
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Task \"%s\" is set to not Start since there cant be more than one start in participant \"%s\"",
						task.GetName(), participant.GetName()))
			}
			if nbStartTask == 0 && task.IsStartTask {
				nbStartTask++
			}
		}
	}

	return
}
