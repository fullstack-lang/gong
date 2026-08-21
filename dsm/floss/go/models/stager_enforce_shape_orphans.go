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

// enforceShapesNotAttachedToSystem removes shapes of elements that are not attached to any System.
// In FLOSS, Complexity, Performance, and Effort elements must be attached to a System to have a shape.
// If they are not attached to any System, their shapes should be removed.
func (stager *Stager) enforceShapesNotAttachedToSystem() (needCommit bool) {
	// 1. Collect all elements that are attached to a System
	attachedComplexities := make(map[*Complexity]struct{})
	attachedPerformances := make(map[*Performance]struct{})
	attachedEfforts := make(map[*Effort]struct{})

	for _, system := range GetGongstrucsSorted[*System](stager.stage) {
		for _, complexity := range system.Complexities {
			attachedComplexities[complexity] = struct{}{}
		}
		for _, performance := range system.Performances {
			attachedPerformances[performance] = struct{}{}
		}
		for _, effort := range system.Efforts {
			attachedEfforts[effort] = struct{}{}
		}
	}

	// 2. Unstage shapes of Complexity elements not attached to any System
	for _, shape := range GetGongstrucsSorted[*ComplexityShape](stager.stage) {
		if _, ok := attachedComplexities[shape.Complexity]; !ok {
			shape.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removing shape for Complexity \"%s\" not attached to any System", shape.GetName()))
			}
		}
	}

	// 3. Unstage shapes of Performance elements not attached to any System
	for _, shape := range GetGongstrucsSorted[*PerformanceShape](stager.stage) {
		if _, ok := attachedPerformances[shape.Performance]; !ok {
			shape.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removing shape for Performance \"%s\" not attached to any System", shape.GetName()))
			}
		}
	}

	// 4. Unstage shapes of Effort elements not attached to any System
	for _, shape := range GetGongstrucsSorted[*EffortShape](stager.stage) {
		if _, ok := attachedEfforts[shape.Effort]; !ok {
			shape.UnstageVoid(stager.stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Removing shape for Effort \"%s\" not attached to any System", shape.GetName()))
			}
		}
	}

	return
}
