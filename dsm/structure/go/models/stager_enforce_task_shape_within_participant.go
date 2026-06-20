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
		var totalWeight float64
		for _, pShape := range diagramProcess.Participant_Shapes {
			weight := pShape.WidthWeight
			if weight == 0 {
				weight = 1.0
			}
			totalWeight += weight
		}

		for _, taskShape := range diagramProcess.Task_Shapes {

			// The bounding shape is ideally the participant shape (swimlane),
			// but we default to the overall process shape.
			boundingX := owningProcessShape.X
			boundingY := owningProcessShape.Y
			boundingWidth := owningProcessShape.Width
			boundingHeight := owningProcessShape.Height

			if taskShape.Task != nil && taskShape.Task.owningParticipant != nil && totalWeight > 0 {
				// Find the index of the owning participant to compute its exact SVG bounds
				currentWeight := 0.0
				var currentParticipantShape *ParticipantShape
				for _, pShape := range diagramProcess.Participant_Shapes {
					if pShape.Participant == taskShape.Task.owningParticipant {
						currentParticipantShape = pShape
						break
					}
					weight := pShape.WidthWeight
					if weight == 0 {
						weight = 1.0
					}
					currentWeight += weight
				}

				if currentParticipantShape != nil {
					boundingX = owningProcessShape.X + horizontalMargin + currentWeight*(participantsWidth/totalWeight)
					shapeWeight := currentParticipantShape.WidthWeight
					if shapeWeight == 0 {
						shapeWeight = 1.0
					}
					boundingWidth = shapeWeight * (participantsWidth / totalWeight)
					boundingY = owningProcessShape.Y + verticalTopMargin
					boundingHeight = owningProcessShape.Height - verticalTopMargin - verticalBottomMargin
				}
			}

			modified := false
			margin := 1.0

			// 1. Enforce maximum dimensions first
			maxWidth := boundingWidth - 2*margin
			if maxWidth < 10.0 { // arbitrary minimum width to prevent negative/invisible shapes
				maxWidth = 10.0
			}
			if taskShape.Width > maxWidth {
				taskShape.Width = maxWidth
				modified = true
			}

			maxHeight := boundingHeight - 2*margin
			if maxHeight < 10.0 {
				maxHeight = 10.0
			}
			if taskShape.Height > maxHeight {
				taskShape.Height = maxHeight
				modified = true
			}

			// 2. Enforce positional boundaries (existing logic)
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
