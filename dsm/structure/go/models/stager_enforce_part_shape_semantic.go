package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforcePartShapeSemantic() (needCommit bool) {
	for _, diagramStructure := range stager.stage.GetInstancesSorted[*DiagramStructure]() {

		owningSystem := diagramStructure.owningSystem
		if owningSystem == nil {
			continue
		}

		// 1. Validate PartShape ownership and parent chain top-down so invalid/unchecked parents
		// cascade immediately to nested PartShapes, PortShapes, ControlFlowShapes, and DataFlowShapes.
		ordered := sortedPartShapesTopDown(diagramStructure.Part_Shapes)
		validSet := make(map[*PartShape]bool, len(ordered))
		anyRemoved := false

		for _, partShape := range ordered {
			if partShape == nil || partShape.Part == nil || partShapeDepth(partShape) < 0 {
				if partShape != nil {
					partShape.UnstageVoid(stager.stage)
					anyRemoved = true
					needCommit = true
				}
				continue
			}

			isValid := false
			if partShape.ParentPartShape == nil {
				isValid = slices.Contains(owningSystem.Parts, partShape.Part)
			} else {
				parent := partShape.ParentPartShape
				if validSet[parent] && parent.Part != nil && parent.Part.Type != nil {
					// Guard against System cycles along the ancestor chain
					ancestorSystems := map[*System]bool{owningSystem: true}
					cycleDetected := ancestorSystems[parent.Part.Type]
					if !cycleDetected {
						ancestorSystems[parent.Part.Type] = true
						for _, anc := range partShapeAncestors(parent) {
							if anc.Part != nil && anc.Part.Type != nil {
								if ancestorSystems[anc.Part.Type] {
									cycleDetected = true
									break
								}
								ancestorSystems[anc.Part.Type] = true
							}
						}
					}
					if !cycleDetected && slices.Contains(parent.Part.Type.Parts, partShape.Part) {
						isValid = true
					}
				}
			}

			if !isValid {
				partShape.UnstageVoid(stager.stage)
				anyRemoved = true
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Unstaged part shape %s (invalid parent or system ownership)", partShape.GetName()))
				}
				continue
			}

			validSet[partShape] = true
		}

		if anyRemoved {
			var kept []*PartShape
			for _, ps := range diagramStructure.Part_Shapes {
				if validSet[ps] {
					kept = append(kept, ps)
				}
			}
			diagramStructure.Part_Shapes = kept
		}

		// Also cascade unstage any concrete child shapes whose explicit ParentPartShape is not in validSet
		// before stage.Clean() clears their ParentPartShape pointer.
		for _, portShape := range diagramStructure.Port_Shapes {
			if portShape != nil && portShape.ParentPartShape != nil && !validSet[portShape.ParentPartShape] {
				portShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}
		for _, cfShape := range diagramStructure.ControlFlow_Shapes {
			if cfShape != nil && cfShape.ParentPartShape != nil && !validSet[cfShape.ParentPartShape] {
				cfShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}
		for _, dfShape := range diagramStructure.DataFlow_Shapes {
			if dfShape != nil && dfShape.ParentPartShape != nil && !validSet[dfShape.ParentPartShape] {
				for _, cp := range dfShape.ControlPointShapes {
					if cp != nil {
						cp.UnstageVoid(stager.stage)
					}
				}
				dfShape.UnstageVoid(stager.stage)
				needCommit = true
			}
		}

		// Find the system shape for the owning system
		owningSystemShape, hasOwningSystemShape := diagramStructure.map_System_SystemShape[owningSystem]

		// 2. Clamp PartShapes top-down within their enclosing shape (owningSystemShape or ParentPartShape)
		for _, partShape := range ordered {
			if !validSet[partShape] {
				continue
			}

			var boundingX, boundingY, boundingWidth, boundingHeight float64
			if partShape.ParentPartShape == nil {
				if !hasOwningSystemShape || owningSystemShape == nil {
					continue
				}
				boundingX = owningSystemShape.X
				boundingY = owningSystemShape.Y
				boundingWidth = owningSystemShape.Width
				boundingHeight = owningSystemShape.Height
			} else {
				parent := partShape.ParentPartShape
				boundingX = parent.X
				boundingY = parent.Y
				boundingWidth = parent.Width
				boundingHeight = parent.Height
			}

			modified := false
			margin := 1.0

			// Enforce maximum dimensions first
			maxWidth := boundingWidth - 2*margin
			if maxWidth < 10.0 {
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

			// Enforce positional boundaries
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
