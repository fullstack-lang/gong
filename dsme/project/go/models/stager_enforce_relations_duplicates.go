package models

func (stager *Stager) enforceRelationDuplicates() (needCommit bool) {
	// Iterate through all diagrams in the stage
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		removeDuplicateRelation(stager, diagram.ProductComposition_Shapes)
		removeDuplicateRelation(stager, diagram.TaskComposition_Shapes)
		removeDuplicateRelation(stager, diagram.TaskInputShapes)
		removeDuplicateRelation(stager, diagram.TaskOutputShapes)
		removeDuplicateRelation(stager, diagram.NoteProductShapes)
		removeDuplicateRelation(stager, diagram.NoteTaskShapes)
	}
	return
}

func removeDuplicateRelation[SourceAT AbstractType, TargetAT AbstractType, ACT AssociationConcreteType2[SourceAT, TargetAT]](
	stager *Stager,
	relations []ACT,
) (needCommit bool) {

	// Map to track unique source-target pairs
	type key struct {
		source string
		target string
	}
	seenPairs := make(map[key]bool)

	// Iterate backwards to safely remove items
	for _, relation := range relations {

		pairKey := key{source: relation.GetAbstractStartElement().GetName(), target: relation.GetAbstractEndElement().GetName()}

		if seenPairs[pairKey] {
			relation.UnstageVoid(stager.stage)
			needCommit = true
		} else {
			// Mark this pair as seen
			seenPairs[pairKey] = true
		}
	}

	return
}
