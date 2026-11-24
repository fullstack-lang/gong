package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceNodeShapeDuplicates() (needCommit bool) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		needCommit = removeDuplicateNodeShape(stager, &diagram.Process_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.Participant_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.ExternalParticipant_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.Task_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.ControlFlow_Shapes) || needCommit
		needCommit = removeDuplicateNodeShape(stager, &diagram.DataFlow_Shapes) || needCommit
		needCommit = removeDuplicateShapeByKey(stager, &diagram.Data_Shapes, func(ds *DataShape) (dataShapeKey, bool) {
			if ds == nil || ds.Data == nil || ds.DataFlow == nil {
				return dataShapeKey{}, false
			}
			return dataShapeKey{dataFlow: ds.DataFlow, data: ds.Data}, true
		}) || needCommit
	}
	return
}

func removeDuplicateShapeByKey[T ConcreteType, K comparable](
	stager *Stager,
	shapes *[]T,
	keyFn func(T) (K, bool),
) (needCommit bool) {
	seen := make(map[K]bool)
	var newShapes []T

	for _, shape := range *shapes {
		if key, ok := keyFn(shape); ok {
			if seen[key] {
				shape.UnstageVoid(stager.stage)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate shape %s", shape.GetName()))
				}
				needCommit = true
				continue
			}
			seen[key] = true
		}
		newShapes = append(newShapes, shape)
	}
	if needCommit {
		*shapes = newShapes
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
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaged duplicate shape %s", shape.GetName()))
				}
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
