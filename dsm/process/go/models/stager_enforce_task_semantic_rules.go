package models

import (
	"fmt"
	"time"
)

// enforceTaskSemanticRules checks
// - task cannot combine start, end, and fork/join roles
// - fork/join task cannot reference a process Type
// - per participant, cannot have more than one start task
func (stager *Stager) enforceTaskSemanticRules() (needCommit bool) {

	for _, task := range stager.stage.GetInstancesSorted[*Task]() {

		if task.IsStartTask && task.IsEndTask {
			needCommit = true
			task.IsEndTask = false
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" cannot be both start and end",
					task.GetName()))
			}
		}

		if task.IsStartTask && task.IsForkJoin {
			needCommit = true
			task.IsForkJoin = false
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" cannot be both start and fork/join",
					task.GetName()))
			}
		}

		if task.IsEndTask && task.IsForkJoin {
			needCommit = true
			task.IsForkJoin = false
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" cannot be both end and fork/join",
					task.GetName()))
			}
		}

		if task.IsForkJoin && task.Type != nil {
			needCommit = true
			task.Type = nil
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" is a fork/join and cannot have a process type",
					task.GetName()))
			}
		}

		if task.Type != nil && !task.IsTaskNameNotProcessName && task.GetName() != task.Type.GetName() {
			needCommit = true
			task.SetName(task.Type.GetName())
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task \"%s\" name overridden to \"%s\"",
					task.GetName(), task.GetName()))
			}
		}
	}

	for _, participant := range stager.stage.GetInstancesSorted[*Participant]() {
		nbStartTask := 0
		for _, task := range participant.Tasks {
			if nbStartTask == 1 && task.IsStartTask {
				needCommit = true
				task.IsStartTask = false
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Task \"%s\" is set to not Start since there cant be more than one start in participant \"%s\"",
							task.GetName(), participant.GetName()))
				}
			}
			if nbStartTask == 0 && task.IsStartTask {
				nbStartTask++
			}
		}
	}

	return
}
