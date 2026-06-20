package models

func (stager *Stager) enforceAtLeastOneDiagramPerSystem() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per system
	for system := range *GetGongstructInstancesSetFromPointerType[*System](stage) {
		if len(system.DiagramStructures) == 0 {
			diagramStructure := (&DiagramStructure{
				Name:        "DiagramStructure",
				IsEditable_: true,
			}).Stage(stage)
			system.DiagramStructures = append(system.DiagramStructures, diagramStructure)

			needCommit = true
		}
	}

	return
}
