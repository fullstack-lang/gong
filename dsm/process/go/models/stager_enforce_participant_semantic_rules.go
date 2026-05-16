package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceParticipantSemanticRules() (needCommit bool) {
	stage := stager.stage

	for _, participant := range GetGongstrucsSorted[*Participant](stage) {
		if participant.owningProcess == nil {
			participant.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan participant \"%s\" (no process)", participant.GetName()))
			}
		}
	}

	return
}
