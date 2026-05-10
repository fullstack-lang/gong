package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceTaskShapeWithinParticipant() (needCommit bool) {
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

		horizontalMargin := 10.0
		verticalTopMargin := 50.0
		verticalBottomMargin := 10.0

		participantsWidth := owningProcessShape.Width - 2*horizontalMargin
		numParticipants := len(diagramProcess.Participant_Shapes)
		var participantWidth float64
		if numParticipants > 0 {
			participantWidth = participantsWidth / float64(numParticipants)
		}

		for _, taskShape := range diagramProcess.Task_Shapes {

			// The bounding shape is ideally the participant shape (swimlane),
			// but we default to the overall process shape.
			boundingX := owningProcessShape.X
			boundingY := owningProcessShape.Y
			boundingWidth := owningProcessShape.Width
			boundingHeight := owningProcessShape.Height

			if taskShape.Task != nil && taskShape.Task.owningParticipant != nil && numParticipants > 0 {
				// Find the index of the owning participant to compute its exact SVG bounds
				idx := -1
				for i, pShape := range diagramProcess.Participant_Shapes {
					if pShape.Participant == taskShape.Task.owningParticipant {
						idx = i
						break
					}
				}

				if idx != -1 {
					boundingX = owningProcessShape.X + horizontalMargin + float64(idx)*participantWidth
					boundingWidth = participantWidth
					boundingY = owningProcessShape.Y + verticalTopMargin
					boundingHeight = owningProcessShape.Height - verticalTopMargin - verticalBottomMargin
				}
			}

			modified := false

			margin := 1.0

			if taskShape.X < boundingX+margin {
				taskShape.X = boundingX + margin
				modified = true
			}

			if taskShape.Y < boundingY+margin {
				taskShape.Y = boundingY + margin
				modified = true
			}

			if taskShape.X+taskShape.Width > boundingX+boundingWidth-margin {
				newX := boundingX + boundingWidth - taskShape.Width - margin
				if newX < boundingX+margin {
					newX = boundingX + margin
				}
				if taskShape.X != newX {
					taskShape.X = newX
					modified = true
				}
			}

			if taskShape.Y+taskShape.Height > boundingY+boundingHeight-margin {
				newY := boundingY + boundingHeight - taskShape.Height - margin
				if newY < boundingY+margin {
					newY = boundingY + margin
				}
				if taskShape.Y != newY {
					taskShape.Y = newY
					modified = true
				}
			}

			if modified {
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Task shape %s moved within bounding shape", taskShape.Name))
				}
			}
		}
	}
	return
}
