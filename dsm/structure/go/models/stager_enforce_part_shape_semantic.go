package models

func (stager *Stager) enforcePartShapeSemantic() (needCommit bool) {
	for _, diagramStructure := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {

		owningSystem := diagramStructure.owningSystem
		if owningSystem == nil {
			continue
		}

	}
	return
}
