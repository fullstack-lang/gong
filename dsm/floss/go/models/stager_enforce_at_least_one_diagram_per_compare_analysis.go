package models

func (stager *Stager) enforceAtLeastOneDiagramPerCompareAnalysis() (needCommit bool) {
	stage := stager.stage

	// enforce that there is at least one diagram per compare analysis
	for compareAnalysis := range *GetGongstructInstancesSetFromPointerType[*CompareAnalysis](stage) {
		if len(compareAnalysis.DiagramFlossEquations) == 0 {
			diagram := (&DiagramFlossEquation{
				Name:        compareAnalysis.Name + " Equation Diagram",
				IsEditable_: true,
				Scale:       5.0,
			}).Stage(stage)
			compareAnalysis.DiagramFlossEquations = append(compareAnalysis.DiagramFlossEquations, diagram)
			needCommit = true
		}

		for _, diagram := range compareAnalysis.DiagramFlossEquations {
			if diagram.Scale == 0 {
				diagram.Scale = 5.0
				needCommit = true
			}
		}
	}

	return
}
