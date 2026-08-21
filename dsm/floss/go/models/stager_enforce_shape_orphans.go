package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableSystemShapes := make(map[*SystemShape]struct{})
	reachableComplexityShapes := make(map[*ComplexityShape]struct{})
	reachablePerformanceShapes := make(map[*PerformanceShape]struct{})
	reachableEffortShapes := make(map[*EffortShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteComplexityShapes := make(map[*NoteComplexityShape]struct{})
	reachableNotePerformanceShapes := make(map[*NotePerformanceShape]struct{})
	reachableNoteEffortShapes := make(map[*NoteEffortShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*DiagramFloss](stager.stage) {
		collectShapes(diagram.System_Shapes, reachableSystemShapes)
		collectShapes(diagram.Complexity_Shapes, reachableComplexityShapes)
		collectShapes(diagram.Performance_Shapes, reachablePerformanceShapes)
		collectShapes(diagram.Effort_Shapes, reachableEffortShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteComplexityShapes, reachableNoteComplexityShapes)
		collectShapes(diagram.NotePerformanceShapes, reachableNotePerformanceShapes)
		collectShapes(diagram.NoteEffortShapes, reachableNoteEffortShapes)
	}

	for _, diagram := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteComplexityShapes, reachableNoteComplexityShapes)
		collectShapes(diagram.NotePerformanceShapes, reachableNotePerformanceShapes)
		collectShapes(diagram.NoteEffortShapes, reachableNoteEffortShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableSystemShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableComplexityShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachablePerformanceShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableEffortShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteComplexityShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNotePerformanceShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteEffortShapes) || needCommit

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
