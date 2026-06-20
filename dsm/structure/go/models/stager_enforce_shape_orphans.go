package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProcessShapes := make(map[*ProcessShape]struct{})
	reachableParticipantShapes := make(map[*ParticipantShape]struct{})
	reachableExternalParticipantShapes := make(map[*ExternalParticipantShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})
	reachableControlFlowShapes := make(map[*ControlFlowShape]struct{})
	reachableDataFlowShapes := make(map[*DataFlowShape]struct{})
	reachableDataShapes := make(map[*DataShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteTaskShapes := make(map[*NoteTaskShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		collectShapes(diagram.Process_Shapes, reachableProcessShapes)
		collectShapes(diagram.Participant_Shapes, reachableParticipantShapes)
		collectShapes(diagram.ExternalParticipant_Shapes, reachableExternalParticipantShapes)
		collectShapes(diagram.Task_Shapes, reachableTaskShapes)
		collectShapes(diagram.ControlFlow_Shapes, reachableControlFlowShapes)
		collectShapes(diagram.DataFlow_Shapes, reachableDataFlowShapes)
		collectShapes(diagram.Data_Shapes, reachableDataShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteTaskShapes, reachableNoteTaskShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableProcessShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableParticipantShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableExternalParticipantShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableControlFlowShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDataFlowShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDataShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteTaskShapes) || needCommit

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
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaging orphan shape object \"%s\" of type \"%s\"",
				object.GetName(), object.GongGetGongstructName()))
		}
	}
	return
}
