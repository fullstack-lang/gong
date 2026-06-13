package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceDiagramDates() (needCommit bool) {
	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		oldStart := diagramHierarchy.ComputedStart
		oldEnd := diagramHierarchy.ComputedEnd
		oldDuration := diagramHierarchy.ComputedDuration

		diagramHierarchy.ComputeStartAndEndDate()

		if diagramHierarchy.ComputedStart != oldStart || diagramHierarchy.ComputedEnd != oldEnd || diagramHierarchy.ComputedDuration != oldDuration {
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Diagram %s dates were recomputed", diagramHierarchy.GetName()))
			}
		}
	}
	return
}
