package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforcePortShapePartPresence() (needCommit bool) {
	for _, diagramStructure := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		for _, portShape := range diagramStructure.Port_Shapes {
			if portShape.Port != nil && portShape.Port.owningPart != nil {
				owningPart := portShape.Port.owningPart

				if _, ok := diagramStructure.map_Part_PartShape[owningPart]; !ok {
					portShape.UnstageVoid(stager.stage)
					needCommit = true
					if stager.probeForm != nil {
						stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port shape %s unstaged because its owning part %s has no shape in the diagram", portShape.Name, owningPart.Name))
					}
				}
			}
		}
	}
	return
}
