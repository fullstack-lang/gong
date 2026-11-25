package models

import "time"

func (stager *Stager) enforceRelationDuplicates() (needCommit bool) {
	// Iterate through all diagrams in the stage
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		needCommit = needCommit || removeDuplicateRelation(stager, diagram.ControlFlow_Shapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagram.DataFlow_Shapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagram.NoteTaskShapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagram.NoteDataFlowShapes)
		needCommit = needCommit || removeDuplicateRelation(stager, diagram.NoteControlFlowShapes)
	}
	return
}

func removeDuplicateRelation[SourceAT AbstractType, TargetAT AbstractType, ACT AssociationConcreteType2[SourceAT, TargetAT]](
	stager *Stager,
	relations []ACT,
) (needCommit bool) {
	// Map to track unique source-target pairs, qualified by Gongstruct type name so
	// polymorphic relations (e.g. DataFlowShape endpoints across Task, ExternalParticipant,
	// and DataStore) never collide on per-type GongGetOrder values.
	type key struct {
		sourceType string
		source     uint
		targetType string
		target     uint
	}
	seenPairs := make(map[key]bool)

	// Iterate through relations to remove true duplicates
	for _, relation := range relations {
		start := relation.GetAbstractStartElement()
		end := relation.GetAbstractEndElement()
		if any(start) == nil || any(end) == nil {
			continue
		}

		pairKey := key{
			sourceType: start.GongGetGongstructName(),
			source:     start.GongGetOrder(stager.stage),
			targetType: end.GongGetGongstructName(),
			target:     end.GongGetOrder(stager.stage),
		}

		if seenPairs[pairKey] {
			relation.UnstageVoid(stager.stage)
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), "Unstaged duplicate relation"+relation.GetName())
			}
			needCommit = true
		} else {
			// Mark this pair as seen
			seenPairs[pairKey] = true
		}
	}

	return
}
