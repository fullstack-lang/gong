package models

func (stager *Stager) enforceAtLeastOneDiagramPerSystem() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per system
	for system := range *GetGongstructInstancesSetFromPointerType[*System](stage) {
		if len(system.DiagramFlossEquations) == 0 {
			diagram := (&DiagramFlossEquation{
				Name:        system.Name + " Equation Diagram",
				IsEditable_: true,
				Scale:       5.0,
			}).Stage(stage)
			system.DiagramFlossEquations = append(system.DiagramFlossEquations, diagram)
			needCommit = true
		}

		for _, diagram := range system.DiagramFlossEquations {
			if diagram.Scale == 0 {
				diagram.Scale = 5.0
				needCommit = true
			}
		}
	}

	return
}
