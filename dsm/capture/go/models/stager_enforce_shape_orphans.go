package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableDeliverableShapes := make(map[*DeliverableShape]struct{})
	reachableDeliverableCompositionShapes := make(map[*DeliverableCompositionShape]struct{})
	reachableTaskShapes := make(map[*ConcernShape]struct{})
	reachableTaskCompositionShapes := make(map[*ConcernCompositionShape]struct{})
	reachableTaskInputShapes := make(map[*ConcernInputShape]struct{})
	reachableTaskOutputShapes := make(map[*ConcernOutputShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteDeliverableShapes := make(map[*NoteDeliverableShape]struct{})
	reachableNoteTaskShapes := make(map[*NoteTaskShape]struct{})
	reachableNoteResourceShapes := make(map[*NoteStakeholderShape]struct{})
	reachableStakeholderShapes := make(map[*StakeholderShape]struct{})
	reachableResourceCompositionShapes := make(map[*StakeholderCompositionShape]struct{})
	reachableStakeholderConcernShapes := make(map[*StakeholderConcernShape]struct{})
	reachableRequirementShapes := make(map[*RequirementShape]struct{})
	reachableViewShapes := make(map[*ViewShape]struct{})
	reachableDeliverableViewShapes := make(map[*DeliverableViewShape]struct{})
	reachableViewConceptShapes := make(map[*ViewConceptShape]struct{})
	reachableConceptShapes := make(map[*ConceptShape]struct{})
	reachableDeliverableConceptShapes := make(map[*DeliverableConceptShape]struct{})
	reachableDiagramShapes := make(map[*DiagramShape]struct{})

	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
		collectShapes(diagram.Deliverable_Shapes, reachableDeliverableShapes)
		collectShapes(diagram.DeliverableComposition_Shapes, reachableDeliverableCompositionShapes)
		collectShapes(diagram.Concern_Shapes, reachableTaskShapes)
		collectShapes(diagram.ConcernComposition_Shapes, reachableTaskCompositionShapes)
		collectShapes(diagram.ConcernInputShapes, reachableTaskInputShapes)
		collectShapes(diagram.ConcernOutputShapes, reachableTaskOutputShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteDeliverableShapes, reachableNoteDeliverableShapes)
		collectShapes(diagram.NoteTaskShapes, reachableNoteTaskShapes)
		collectShapes(diagram.NoteResourceShapes, reachableNoteResourceShapes)
		collectShapes(diagram.Stakeholder_Shapes, reachableStakeholderShapes)
		collectShapes(diagram.ResourceComposition_Shapes, reachableResourceCompositionShapes)
		collectShapes(diagram.StakeholderConcernShapes, reachableStakeholderConcernShapes)
		collectShapes(diagram.Requirement_Shapes, reachableRequirementShapes)
		collectShapes(diagram.View_Shapes, reachableViewShapes)
		collectShapes(diagram.DeliverableViewShapes, reachableDeliverableViewShapes)
		collectShapes(diagram.ViewConceptShapes, reachableViewConceptShapes)
		collectShapes(diagram.Concept_Shapes, reachableConceptShapes)
		collectShapes(diagram.DeliverableConceptShapes, reachableDeliverableConceptShapes)
		collectShapes(diagram.Diagram_Shapes, reachableDiagramShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageUnreachableOrphans(stager, reachableDeliverableShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDeliverableCompositionShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskCompositionShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskInputShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableTaskOutputShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteDeliverableShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteTaskShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableNoteResourceShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableStakeholderShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableResourceCompositionShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableStakeholderConcernShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableRequirementShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableViewShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDeliverableViewShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableViewConceptShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableConceptShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDeliverableConceptShapes) || needCommit
	needCommit = unstageUnreachableOrphans(stager, reachableDiagramShapes) || needCommit

	return
}

func collectShapes[T comparable](shapes []T, reachable map[T]struct{}) {
	for _, shape := range shapes {
		reachable[shape] = struct{}{}
	}
}

func unstageUnreachableOrphans[T PointerToGongstruct](stager *Stager, reachable map[T]struct{}) (needCommit bool) {
	for _, object := range stager.stage.GetInstancesSorted[T]() {
		if _, ok := reachable[object]; !ok {
			object.UnstageVoid(stager.stage)
			needCommit = true
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unstaging orphan shape object \"%s\" of type \"%s\"",
				object.GetName(), object.GongGetGongstructName()))
		}
	}
	return
}
