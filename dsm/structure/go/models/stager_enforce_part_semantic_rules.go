package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforcePartSemanticRules() (needCommit bool) {
	stage := stager.stage

	sortedParts := stage.GetInstancesSorted[*Part]()
	var validParts []*Part

	for _, part := range sortedParts {
		if part.owningSystem == nil {
			part.UnstageVoid(stage)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Unstaged orphan part \"%s\" (no system)", part.GetName()))
			}
			continue
		}
		validParts = append(validParts, part)
	}

	// Pass 1: check parts whose Type changed since the last settled state first,
	// so newly introduced cycles revert the offending edit rather than pre-existing links.
	for _, part := range validParts {
		if stager.map_Part_TypeTracked[part] && stager.map_Part_PrevType[part] == part.Type {
			continue
		}
		if part.Type != nil && reachesSystemViaPartTypes(part.Type, part.owningSystem) {
			typeName := part.Type.GetName()
			var fallback *System
			if stager.map_Part_TypeTracked[part] {
				prev := stager.map_Part_PrevType[part]
				if prev != nil && !reachesSystemViaPartTypes(prev, part.owningSystem) {
					fallback = prev
				}
			}
			part.Type = fallback
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Cleared cyclic Type \"%s\" on part \"%s\"", typeName, part.GetName()))
			}
		}
	}

	// Pass 2: safety fallback in case multiple cyclic links existed simultaneously on load.
	for _, part := range validParts {
		if part.Type != nil && reachesSystemViaPartTypes(part.Type, part.owningSystem) {
			typeName := part.Type.GetName()
			part.Type = nil
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Cleared cyclic Type \"%s\" on part \"%s\"", typeName, part.GetName()))
			}
		}
	}

	stager.map_Part_PrevType = make(map[*Part]*System, len(validParts))
	stager.map_Part_TypeTracked = make(map[*Part]bool, len(validParts))
	for _, part := range validParts {
		stager.map_Part_PrevType[part] = part.Type
		stager.map_Part_TypeTracked[part] = true
	}

	return
}
