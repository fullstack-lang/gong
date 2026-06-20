package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforcePartShapeSemantic() (needCommit bool) {
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

		for _, partShape := range diagramStructure.Part_Shapes {

			boundingX := owningSystemShape.X
			boundingY := owningSystemShape.Y
			boundingWidth := owningSystemShape.Width
			boundingHeight := owningSystemShape.Height

			modified := false
			margin := 1.0

			// 1. Enforce maximum dimensions first
			maxWidth := boundingWidth - 2*margin
			if maxWidth < 10.0 { // arbitrary minimum width to prevent negative/invisible shapes
				maxWidth = 10.0
			}
			if partShape.Width > maxWidth {
				partShape.Width = maxWidth
				modified = true
			}

			maxHeight := boundingHeight - 2*margin
			if maxHeight < 10.0 {
				maxHeight = 10.0
			}
			if partShape.Height > maxHeight {
				partShape.Height = maxHeight
				modified = true
			}

			// 2. Enforce positional boundaries (existing logic)
			if partShape.X < boundingX+margin {
				partShape.X = boundingX + margin
				modified = true
			}

			if partShape.Y < boundingY+margin {
				partShape.Y = boundingY + margin
				modified = true
			}

			if partShape.X+partShape.Width > boundingX+boundingWidth-margin {
				newX := boundingX + boundingWidth - partShape.Width - margin
				if newX < boundingX+margin {
					newX = boundingX + margin
				}
				if partShape.X != newX {
					partShape.X = newX
					modified = true
				}
			}

			if partShape.Y+partShape.Height > boundingY+boundingHeight-margin {
				newY := boundingY + boundingHeight - partShape.Height - margin
				if newY < boundingY+margin {
					newY = boundingY + margin
				}
				if partShape.Y != newY {
					partShape.Y = newY
					modified = true
				}
			}

			if modified {
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Part shape %s moved within bounding shape", partShape.Name))
				}
			}
		}
	}
	return
}
