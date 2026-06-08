package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceInfluenceCardinality() (needCommit bool) {
	stage := stager.stage

	for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {
		sourceCount := 0
		if influence.SourceMovement != nil {
			sourceCount++
		}
		if influence.SourceArtefactType != nil {
			sourceCount++
		}
		if influence.SourceArtist != nil {
			sourceCount++
		}

		targetCount := 0
		if influence.TargetMovement != nil {
			targetCount++
		}
		if influence.TargetArtefactType != nil {
			targetCount++
		}
		if influence.TargetArtist != nil {
			targetCount++
		}

		if sourceCount != 1 || targetCount != 1 {
			influence.Unstage(stager.stage)
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Influence %s has %d sources and %d targets (must be exactly 1 each), unstaging it", influence.GetName(), sourceCount, targetCount))
			}
		}
	}

	return
}
