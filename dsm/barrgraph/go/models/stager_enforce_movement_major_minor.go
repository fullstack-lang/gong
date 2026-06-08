package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceMovementMajorMinor() (needCommit bool) {
	stage := stager.stage
	for _, movement := range GetGongstrucsSorted[*Movement](stager.stage) {
		//  movement cannot be minor AND major
		if movement.IsMajor && movement.IsMinor {
			movement.IsMinor = false
			needCommit = true
			if stage.probeIF != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Movement %s cannot be both major and minor, setting IsMinor to false", movement.GetName()))
			}
		}
	}
	return
}
