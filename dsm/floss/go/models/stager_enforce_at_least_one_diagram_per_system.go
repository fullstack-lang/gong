package models

func (stager *Stager) enforceAtLeastOneDiagramPerSystem() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per system
	for system := range *GetGongstructInstancesSetFromPointerType[*System](stage) {
		if len(system.DiagramFlosses) == 0 {
			diagramFloss := (&DiagramFloss{
				Name:        "DiagramFloss",
				IsEditable_: true,
			}).Stage(stage)
			system.DiagramFlosses = append(system.DiagramFlosses, diagramFloss)

			systemShape := (&SystemShape{
				Name:   "SystemShape",
				System: system,
				RectShape: RectShape{
					X:      100,
					Y:      50,
					Width:  500,
					Height: 1000,
				},
			}).Stage(stage)
			diagramFloss.System_Shapes = append(diagramFloss.System_Shapes, systemShape)

			needCommit = true
		}
	}

	return
}
