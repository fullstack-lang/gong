package models

import (
	"fmt"
	"math"
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

		for _, portShape := range diagramStructure.Port_Shapes {

			// If the port shape is attached to a part shape, enforce it on the border
			if partShape := portShape.owningPartShape; partShape != nil {
				modified := false

				// portShape center
				cx := portShape.X + portShape.Width/2.0
				cy := portShape.Y + portShape.Height/2.0

				// partShape borders
				left := partShape.X
				right := partShape.X + partShape.Width
				top := partShape.Y
				bottom := partShape.Y + partShape.Height

				distLeft := math.Abs(cx - left)
				distRight := math.Abs(cx - right)
				distTop := math.Abs(cy - top)
				distBottom := math.Abs(cy - bottom)

				minDist := distLeft
				closestBorder := "left"
				if distRight < minDist {
					minDist = distRight
					closestBorder = "right"
				}
				if distTop < minDist {
					minDist = distTop
					closestBorder = "top"
				}
				if distBottom < minDist {
					minDist = distBottom
					closestBorder = "bottom"
				}

				// check if we are on the border
				if minDist > 0.01 { // not on border
					switch closestBorder {
					case "left":
						portShape.X = left - portShape.Width/2.0
					case "right":
						portShape.X = right - portShape.Width/2.0
					case "top":
						portShape.Y = top - portShape.Height/2.0
					case "bottom":
						portShape.Y = bottom - portShape.Height/2.0
					}
					modified = true
				}

				// Also constrain the other dimension so it stays within the part segment
				// For left/right borders, constrain cy between top and bottom
				// For top/bottom borders, constrain cx between left and right
				cx = portShape.X + portShape.Width/2.0
				cy = portShape.Y + portShape.Height/2.0

				if closestBorder == "left" || closestBorder == "right" {
					if cy < top {
						portShape.Y = top - portShape.Height/2.0
						modified = true
					}
					if cy > bottom {
						portShape.Y = bottom - portShape.Height/2.0
						modified = true
					}
				} else {
					if cx < left {
						portShape.X = left - portShape.Width/2.0
						modified = true
					}
					if cx > right {
						portShape.X = right - portShape.Width/2.0
						modified = true
					}
				}

				if modified {
					needCommit = true
					if stager.probeForm != nil {
						stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port shape %s moved to border of part shape %s", portShape.Name, partShape.Name))
					}
				}
			} else {
				// The bounding shape is ideally the part shape (swimlane),
				// but we default to the overall system shape if no part shape is found.
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
						stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Port shape %s moved within bounding system shape", portShape.Name))
					}
				}
			}
		}
	}
	return
}
