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

			systemShape := (&SystemShape{
				Name:    "SystemShape",
				System: system,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stage)
			diagramStructure.System_Shapes = append(diagramStructure.System_Shapes, systemShape)

			needCommit = true
		}
	}

	return
}
