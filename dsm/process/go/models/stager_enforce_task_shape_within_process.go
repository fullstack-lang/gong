package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceTaskShapeWithinProcess() (needCommit bool) {
	for _, diagramProcess := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		owningProcess := diagramProcess.owningProcess
		if owningProcess == nil {
			continue
		}

		// Find the process shape for the owning process
		var owningProcessShape *ProcessShape
		if shape, ok := diagramProcess.map_Process_ProcessShape[owningProcess]; ok {
			owningProcessShape = shape
		} else {
			continue
		}

		for _, taskShape := range diagramProcess.TaskShapes {

			// Check if task shape is within process shape
			modified := false

			margin := 20.0

			if taskShape.X < owningProcessShape.X+margin {
				taskShape.X = owningProcessShape.X + margin
				modified = true
			}

			if taskShape.Y < owningProcessShape.Y+margin {
				taskShape.Y = owningProcessShape.Y + margin
				modified = true
			}

			if taskShape.X+taskShape.Width > owningProcessShape.X+owningProcessShape.Width-margin {
				newX := owningProcessShape.X + owningProcessShape.Width - taskShape.Width - margin
				if newX < owningProcessShape.X+margin {
					newX = owningProcessShape.X + margin
				}
				if taskShape.X != newX {
					taskShape.X = newX
					modified = true
				}
			}

			if taskShape.Y+taskShape.Height > owningProcessShape.Y+owningProcessShape.Height-margin {
				newY := owningProcessShape.Y + owningProcessShape.Height - taskShape.Height - margin
				if newY < owningProcessShape.Y+margin {
					newY = owningProcessShape.Y + margin
				}
				if taskShape.Y != newY {
					taskShape.Y = newY
					modified = true
				}
			}

			if modified {
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task shape %s moved within owning process shape", taskShape.Name))
				}
			}
		}
	}
	return
}
