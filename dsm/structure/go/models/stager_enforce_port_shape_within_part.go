package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforcePortShapeWithinPart() (needCommit bool) {
	for _, diagramStructure := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		owningSystem := diagramStructure.owningSystem
		if owningSystem == nil {
			continue
		}

		// Find the system shape for the owning system
		var owningSystemShape *SystemShape
		if shape, ok := diagramStructure.map_System_SystemShape[owningSystem]; ok {
			owningSystemShape = shape
		} else {
			continue
		}

		horizontalMargin := 10.0
		verticalTopMargin := 50.0
		verticalBottomMargin := 10.0

		partsWidth := owningSystemShape.Width - 2*horizontalMargin
		var totalWeight float64
		for _, pShape := range diagramStructure.Part_Shapes {
			weight := pShape.WidthWeight
			if weight == 0 {
				weight = 1.0
			}
			totalWeight += weight
		}

		for _, portShape := range diagramStructure.Port_Shapes {

			// The bounding shape is ideally the part shape (swimlane),
			// but we default to the overall system shape.
			boundingX := owningSystemShape.X
			boundingY := owningSystemShape.Y
			boundingWidth := owningSystemShape.Width
			boundingHeight := owningSystemShape.Height

			if portShape.Port != nil && portShape.Port.owningPart != nil && totalWeight > 0 {
				// Find the index of the owning part to compute its exact SVG bounds
				currentWeight := 0.0
				var currentPartShape *PartShape
				for _, pShape := range diagramStructure.Part_Shapes {
					if pShape.Part == portShape.Port.owningPart {
						currentPartShape = pShape
						break
					}
					weight := pShape.WidthWeight
					if weight == 0 {
						weight = 1.0
					}
					currentWeight += weight
				}

				if currentPartShape != nil {
					boundingX = owningSystemShape.X + horizontalMargin + currentWeight*(partsWidth/totalWeight)
					shapeWeight := currentPartShape.WidthWeight
					if shapeWeight == 0 {
						shapeWeight = 1.0
					}
					boundingWidth = shapeWeight * (partsWidth / totalWeight)
					boundingY = owningSystemShape.Y + verticalTopMargin
					boundingHeight = owningSystemShape.Height - verticalTopMargin - verticalBottomMargin
				}
			}

			modified := false
			margin := 1.0

			// 1. Enforce maximum dimensions first
			maxWidth := boundingWidth - 2*margin
			if maxWidth < 10.0 { // arbitrary minimum width to prevent negative/invisible shapes
				maxWidth = 10.0
			}
			if portShape.Width > maxWidth {
				portShape.Width = maxWidth
				modified = true
			}

			maxHeight := boundingHeight - 2*margin
			if maxHeight < 10.0 {
				maxHeight = 10.0
			}
			if portShape.Height > maxHeight {
				portShape.Height = maxHeight
				modified = true
			}

			// 2. Enforce positional boundaries (existing logic)
			if portShape.X < boundingX+margin {
				portShape.X = boundingX + margin
				modified = true
			}

			if portShape.Y < boundingY+margin {
				portShape.Y = boundingY + margin
				modified = true
			}

			if portShape.X+portShape.Width > boundingX+boundingWidth-margin {
				newX := boundingX + boundingWidth - portShape.Width - margin
				if newX < boundingX+margin {
					newX = boundingX + margin
				}
				if portShape.X != newX {
					portShape.X = newX
					modified = true
				}
			}

			if portShape.Y+portShape.Height > boundingY+boundingHeight-margin {
				newY := boundingY + boundingHeight - portShape.Height - margin
				if newY < boundingY+margin {
					newY = boundingY + margin
				}
				if portShape.Y != newY {
					portShape.Y = newY
					modified = true
				}
			}

			if modified {
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port shape %s moved within bounding shape", portShape.Name))
				}
			}
		}
	}
	return
}
