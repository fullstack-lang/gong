package models

func (stager *Stager) enforceFlossEquation() (needCommit bool) {
	for compareAnalysis := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
		// Enforce alpha != 0
		if compareAnalysis.Alpha == 0 {
			compareAnalysis.Alpha = 1.0
			needCommit = true
		}

		// Enforce default beta if 0
		if compareAnalysis.Beta == 0 {
			compareAnalysis.Beta = 1.0
			needCommit = true
		}

		// Ensure at least one diagram per compare analysis
		if len(compareAnalysis.DiagramFlossEquations) == 0 {
			diagram := new(DiagramFlossEquation).Stage(stager.stage)
			diagram.Name = compareAnalysis.Name + " Equation Diagram"
			diagram.Scale = 5.0
			diagram.IsEditable_ = true
			compareAnalysis.DiagramFlossEquations = append(compareAnalysis.DiagramFlossEquations, diagram)
			needCommit = true
		}


		for _, diagram := range compareAnalysis.DiagramFlossEquations {
			diagram.SetOwningCompareAnalysis(compareAnalysis)
			if diagram.Scale == 0 {
				diagram.Scale = 5.0
				needCommit = true
			}
		}
	}

	return
}
