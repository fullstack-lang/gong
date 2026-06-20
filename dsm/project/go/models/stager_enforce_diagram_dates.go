package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDiagramDates() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		oldStart := diagram.ComputedStart
		oldEnd := diagram.ComputedEnd
		oldDuration := diagram.ComputedDuration

		diagram.computeStartAndEndDate()

		if diagram.ComputedStart != oldStart || diagram.ComputedEnd != oldEnd || diagram.ComputedDuration != oldDuration {
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Diagram %s dates were recomputed", diagram.GetName()))
			}
		}
	}
	return
}
