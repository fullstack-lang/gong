package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceInfluenceShapeConsistency() (needCommit bool) {
	stage := stager.stage

	for _, influenceShape := range GetGongstrucsSorted[*InfluenceShape](stager.stage) {
		if influenceShape.Influence == nil {
			influenceShape.Unstage(stager.stage)
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s has no influence, unstaging it", influenceShape.GetName()))
			}
			continue
		}
		if influenceShape.Name != influenceShape.Influence.Name {
			influenceShape.Name = influenceShape.Influence.Name
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("InfluenceShape %s has a name not consistent with its influence, setting it to %s", influenceShape.GetName(), influenceShape.Name))
			}
			continue
		}
	}

	return
}
