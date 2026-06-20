package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforcePartSemanticRules() (needCommit bool) {
	stage := stager.stage

	for _, part := range GetGongstrucsSorted[*Part](stage) {
		if part.owningSystem == nil {
			part.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan part \"%s\" (no system)", part.GetName()))
			}
		}
	}

	return
}
