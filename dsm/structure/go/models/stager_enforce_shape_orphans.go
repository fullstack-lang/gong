package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableSystemShapes := make(map[*SystemShape]struct{})
	reachablePartShapes := make(map[*PartShape]struct{})
	reachableExternalPartShapes := make(map[*ExternalPartShape]struct{})
	reachablePortShapes := make(map[*PortShape]struct{})
	reachableControlFlowShapes := make(map[*ControlFlowShape]struct{})
	reachableDataFlowShapes := make(map[*DataFlowShape]struct{})
	reachableDataShapes := make(map[*DataShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNotePortShapes := make(map[*NotePortShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		collectShapes(diagram.System_Shapes, reachableSystemShapes)
		collectShapes(diagram.Part_Shapes, reachablePartShapes)
		collectShapes(diagram.ExternalPart_Shapes, reachableExternalPartShapes)
		collectShapes(diagram.Port_Shapes, reachablePortShapes)
		collectShapes(diagram.ControlFlow_Shapes, reachableControlFlowShapes)
		collectShapes(diagram.DataFlow_Shapes, reachableDataFlowShapes)
		collectShapes(diagram.Data_Shapes, reachableDataShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NotePortShapes, reachableNotePortShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableSystemShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachablePartShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableExternalPartShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachablePortShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableControlFlowShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDataFlowShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDataShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNotePortShapes) || needCommit

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
