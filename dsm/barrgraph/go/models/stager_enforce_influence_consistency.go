package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceInfluenceConsistency() (needCommit bool) {
	stage := stager.stage

	for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {
		if influence.SourceMovement != nil {
			influence.source = influence.SourceMovement
		}
		if influence.SourceArtefactType != nil {
			influence.source = influence.SourceArtefactType
		}
		if influence.SourceArtist != nil {
			influence.source = influence.SourceArtist
		}

		if influence.TargetMovement != nil {
			influence.target = influence.TargetMovement
		}
		if influence.TargetArtefactType != nil {
			influence.target = influence.TargetArtefactType
		}
		if influence.TargetArtist != nil {
			influence.target = influence.TargetArtist
		}

		if influence.source == nil || influence.target == nil {
			continue
		}

		if influence.Name != influence.source.GetName()+" to "+influence.target.GetName() {
			influence.Name = influence.source.GetName() + " to " + influence.target.GetName()
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Influence %s has a name not consistent with its source and target, setting it to %s", influence.GetName(), influence.Name))
			}
		}
	}

	return
}
