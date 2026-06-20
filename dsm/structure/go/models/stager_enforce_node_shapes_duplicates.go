package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceNodeShapeDuplicates() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		needCommit = removeDuplicateNodeShape(stager, &diagram.System_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.Part_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.Port_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.ControlFlow_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.DataFlow_Shapes) || needCommit
	}
	return
}

func removeDuplicateNodeShape[T ConcreteType](stager *Stager, shapes *[]T) (needCommit bool) {
	seen := make(map[AbstractType]bool)
	var newShapes []T

	for _, shape := range *shapes {
		abstractElement := shape.GetAbstractElement()
		if abstractElement != nil {
			if seen[abstractElement] {
				shape.UnstageVoid(stager.stage)
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate shape %s", shape.GetName()))
				needCommit = true
				continue
			}
			seen[abstractElement] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		*shapes = newShapes
	}
	return
}
