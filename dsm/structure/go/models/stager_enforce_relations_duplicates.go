package models

import "time"

func (stager *Stager) enforceRelationDuplicates() (needCommit bool) {
	// Iterate through all diagrams in the stage
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		_ = diagram
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
		start := relation.GetAbstractStartElement()
		end := relation.GetAbstractEndElement()
		if any(start) == nil || any(end) == nil {
			continue
		}

		pairKey := key{
			source: start.GongGetOrder(stager.stage),
			target: end.GongGetOrder(stager.stage),
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
