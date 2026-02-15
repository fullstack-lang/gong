// new file
package models

import (
	"slices"
	"time"
)

func (stager *Stager) enforceVisibility() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		visibleElements := make(map[AbstractType]struct{})

		collectVisibleElements(diagram.Product_Shapes, visibleElements)
		collectVisibleElements(diagram.Task_Shapes, visibleElements)
		collectVisibleElements(diagram.Note_Shapes, visibleElements)
		collectVisibleElements(diagram.Resource_Shapes, visibleElements)

		needCommit = removeInvisibleShapes(stager, &diagram.ProductComposition_Shapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagram.TaskComposition_Shapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagram.TaskInputShapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagram.TaskOutputShapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagram.NoteProductShapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagram.NoteTaskShapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagram.ResourceComposition_Shapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagram.ResourceTaskShapes, visibleElements) || needCommit
	}
	return
}

func collectVisibleElements[T ConcreteType](shapes []T, visibleElements map[AbstractType]struct{}) {
	for _, s := range shapes {
		if e := s.GetAbstractElement(); e != nil {
			visibleElements[e] = struct{}{}
		}
	}
}

func removeInvisibleShapes[T AssociationConcreteType](stager *Stager, shapes *[]T, visibleElements map[AbstractType]struct{}) (needCommit bool) {
	for i := len(*shapes) - 1; i >= 0; i-- {
		shape := (*shapes)[i]
		start := shape.GetAbstractStartElement()
		end := shape.GetAbstractEndElement()

		_, startVisible := visibleElements[start]
		_, endVisible := visibleElements[end]

		if !startVisible || !endVisible {
			*shapes = slices.Delete(*shapes, i, i+1)
			shape.UnstageVoid(stager.stage)
			stager.probeForm.AddNotification(time.Now(), "Removed shape: "+shape.GetName())
			needCommit = true
		}
	}
	return
}
