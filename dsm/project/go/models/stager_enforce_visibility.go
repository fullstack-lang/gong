// new file
package models

import (
	"slices"
	"time"
)

func (stager *Stager) enforceVisibility() (needCommit bool) {
	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		visibleElements := make(map[AbstractType]struct{})

		collectVisibleElements(diagramHierarchy.Product_Shapes, visibleElements)
		collectVisibleElements(diagramHierarchy.Task_Shapes, visibleElements)
		collectVisibleElements(diagramHierarchy.Note_Shapes, visibleElements)
		collectVisibleElements(diagramHierarchy.Resource_Shapes, visibleElements)

		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.ProductComposition_Shapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.TaskComposition_Shapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.TaskInputShapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.TaskOutputShapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.NoteProductShapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.NoteTaskShapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.NoteResourceShapes, visibleElements) || needCommit

		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.ResourceComposition_Shapes, visibleElements) || needCommit
		needCommit = removeInvisibleShapes(stager, &diagramHierarchy.ResourceTaskShapes, visibleElements) || needCommit
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
