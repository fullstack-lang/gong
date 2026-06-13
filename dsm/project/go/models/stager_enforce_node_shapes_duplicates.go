// new file
package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceNodeShapeDuplicates() (needCommit bool) {
	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		needCommit = removeDuplicateNodeShape(stager, &diagramHierarchy.Product_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagramHierarchy.Task_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagramHierarchy.Note_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagramHierarchy.Resource_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagramHierarchy.TaskGroupShapes) || needCommit
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
