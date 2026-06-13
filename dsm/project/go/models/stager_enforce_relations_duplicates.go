package models

import "time"

func (stager *Stager) enforceRelationDuplicates() (needCommit bool) {
	// Iterate through all diagrams in the stage
	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.ProductComposition_Shapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.TaskComposition_Shapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.TaskInputShapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.TaskOutputShapes)

		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.NoteProductShapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.NoteTaskShapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.NoteResourceShapes)

		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.ResourceComposition_Shapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagramHierarchy.ResourceTaskShapes)
	}
	return
}

func removeDuplicateRelation[SourceAT AbstractType, TargetAT AbstractType, ACT AssociationConcreteType2[SourceAT, TargetAT]](
	stager *Stager,
	relations []ACT,
) (needCommit bool) {

	// Map to track unique source-target pairs
	type key struct {
		source uint
		target uint
	}
	seenPairs := make(map[key]bool)

	// Iterate backwards to safely remove items
	for _, relation := range relations {

		pairKey := key{
			source: relation.GetAbstractStartElement().GongGetOrder(stager.stage),
			target: relation.GetAbstractEndElement().GongGetOrder(stager.stage)}

		if seenPairs[pairKey] {
			relation.UnstageVoid(stager.stage)
			stager.probeForm.AddNotification(time.Now(), "Unstaged duplicate relation"+relation.GetName())
			needCommit = true
		} else {
			// Mark this pair as seen
			seenPairs[pairKey] = true
		}
	}

	return
}
