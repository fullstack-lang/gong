package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProductShapes := make(map[*ProductShape]struct{})
	reachableProductCompositionShapes := make(map[*ProductCompositionShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})
	reachableTaskCompositionShapes := make(map[*TaskCompositionShape]struct{})
	reachableTaskInputShapes := make(map[*TaskInputShape]struct{})
	reachableTaskOutputShapes := make(map[*TaskOutputShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteProductShapes := make(map[*NoteProductShape]struct{})
	reachableNoteTaskShapes := make(map[*NoteTaskShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		collectShapes(diagram.Product_Shapes, reachableProductShapes)
		collectShapes(diagram.ProductComposition_Shapes, reachableProductCompositionShapes)
		collectShapes(diagram.Task_Shapes, reachableTaskShapes)
		collectShapes(diagram.TaskComposition_Shapes, reachableTaskCompositionShapes)
		collectShapes(diagram.TaskInputShapes, reachableTaskInputShapes)
		collectShapes(diagram.TaskOutputShapes, reachableTaskOutputShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteProductShapes, reachableNoteProductShapes)
		collectShapes(diagram.NoteTaskShapes, reachableNoteTaskShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageOrphans(stager, reachableProductShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableProductCompositionShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableTaskShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableTaskCompositionShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableTaskInputShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableTaskOutputShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableNoteProductShapes) || needCommit
	needCommit = unstageOrphans(stager, reachableNoteTaskShapes) || needCommit

	return
}

func collectShapes[T comparable](shapes []T, reachable map[T]struct{}) {
	for _, shape := range shapes {
		reachable[shape] = struct{}{}
	}
}

func unstageOrphans[T PointerToGongstruct](stager *Stager, reachable map[T]struct{}) (needCommit bool) {
	for _, shape := range GetGongstrucsSorted[T](stager.stage) {
		if _, ok := reachable[shape]; !ok {
			shape.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaging orphan shape %s", shape.GetName()))
			}
		}
	}
	return
}
