package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceControlFlowShapesRules() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {

		var validControlFlowShapes []*ControlFlowShape
		for _, controlFlowShape := range diagram.ControlFlow_Shapes {
			isValid := false
			if controlFlowShape.ControlFlow != nil &&
				controlFlowShape.ControlFlow.Start != nil &&
				controlFlowShape.ControlFlow.End != nil {
				parentPS := effectiveControlFlowParentPartShape(diagram, controlFlowShape)
				if parentPS != nil && slices.Contains(diagram.Part_Shapes, parentPS) && diagram.map_PortShapeKey_PortShape != nil {
					_, hasStart := diagram.map_PortShapeKey_PortShape[portShapeKey{
						port:            controlFlowShape.ControlFlow.Start,
						parentPartShape: parentPS,
					}]
					_, hasEnd := diagram.map_PortShapeKey_PortShape[portShapeKey{
						port:            controlFlowShape.ControlFlow.End,
						parentPartShape: parentPS,
					}]
					isValid = hasStart && hasEnd
				}
			}

			if isValid {
				validControlFlowShapes = append(validControlFlowShapes, controlFlowShape)
				if controlFlowShape.Name != controlFlowShape.ControlFlow.Name {
					controlFlowShape.Name = controlFlowShape.ControlFlow.Name
					needCommit = true
				}
			} else {
				controlFlowShape.UnstageVoid(stager.stage)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged control flow shape %s (missing start or end port shape)", controlFlowShape.GetName()))
				}
			}
		}
		if len(validControlFlowShapes) != len(diagram.ControlFlow_Shapes) {
			diagram.ControlFlow_Shapes = validControlFlowShapes
			needCommit = true
		}
	}

	return
}
