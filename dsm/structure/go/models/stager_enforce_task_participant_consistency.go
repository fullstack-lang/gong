package models

import (
	"fmt"
	"slices"
	"time"
)

// enforceTaskParticipantConsistency unstages Tasks that do not belong to any participant
// or that belong to more than one participant.
func (stager *Stager) enforceTaskParticipantConsistency() (needCommit bool) {
	stage := stager.stage

	// 1. Build a reverse map from Task to its owning Participants
	taskToParticipants := make(map[*Task][]*Participant)
	for _, participant := range GetGongstrucsSorted[*Participant](stage) {
		for _, task := range participant.Tasks {
			taskToParticipants[task] = append(taskToParticipants[task], participant)
		}
	}

	// 2. Iterate over all tasks and check their ownership
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		owners := taskToParticipants[task]

		// Unstage tasks with no owner
		if len(owners) == 0 {
			task.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan task \"%s\" (no participant)", task.GetName()))
			}
			continue // continue to next task
		}

		// Unstage tasks with more than one owner
		if len(owners) > 1 {
			// To make the operation idempotent and clear, we remove the task from all its owners' lists
			// before unstaging it.
			for _, owner := range owners {
				if idx := slices.Index(owner.Tasks, task); idx != -1 {
					owner.Tasks = slices.Delete(owner.Tasks, idx, idx+1)
				}
			}
			task.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged task \"%s\" with multiple participants", task.GetName()))
			}
		}
	}

	return
}
