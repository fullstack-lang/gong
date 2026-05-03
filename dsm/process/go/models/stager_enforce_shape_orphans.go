package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProcessShapes := make(map[*ProcessShape]struct{})
	reachableParticipantShapes := make(map[*ParticipantShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})
	reachableControlFlowShapes := make(map[*ControlFlowShape]struct{})
	reachableDataFlowShapes := make(map[*DataFlowShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		collectShapes(diagram.Process_Shapes, reachableProcessShapes)
		collectShapes(diagram.Participant_Shapes, reachableParticipantShapes)
		collectShapes(diagram.Task_Shapes, reachableTaskShapes)
		collectShapes(diagram.ControlFlow_Shapes, reachableControlFlowShapes)
		collectShapes(diagram.DataFlow_Shapes, reachableDataFlowShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableProcessShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableParticipantShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableControlFlowShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDataFlowShapes) || needCommit

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
