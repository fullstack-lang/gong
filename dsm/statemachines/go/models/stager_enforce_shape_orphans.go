package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableStateShapes := make(map[*StateShape]struct{})
	reachableTransitionShapes := make(map[*Transition_Shape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteStateShapes := make(map[*NoteStateShape]struct{})
	reachableNoteTransitionShapes := make(map[*NoteTransitionShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		collectShapes(diagram.State_Shapes, reachableStateShapes)
		collectShapes(diagram.Transition_Shapes, reachableTransitionShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteState_Shapes, reachableNoteStateShapes)
		collectShapes(diagram.NoteTransition_Shapes, reachableNoteTransitionShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableStateShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTransitionShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteStateShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteTransitionShapes) || needCommit

	return
}

func collectShapes[T comparable](shapes []T, reachable map[T]struct{}) {
	for _, shape := range shapes {
		reachable[shape] = struct{}{}
	}
}

func unstageUnreachableOrphans[T PointerToGongstruct](stager *Stager, reachable map[T]struct{}) (needCommit bool) {
	for _, object := range GetGongstrucsSorted[T](stager.stage) {
		if _, ok := reachable[object]; !ok {
			object.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaging orphan shape object \"%s\" of type \"%s\"",
					object.GetName(), object.GongGetGongstructName()))
			}
		}
	}
	return
}
