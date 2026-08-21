package models

func (stager *Stager) enforceFlossEquation() (needCommit bool) {
	for compareAnalysis := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
		// Enforce alpha != 0
		if compareAnalysis.Alpha == 0 {
			compareAnalysis.Alpha = 1.0
			needCommit = true
		}

		// Ensure at least one diagram per compare analysis
		if len(compareAnalysis.DiagramFlossEquations) == 0 {
			diagram := new(DiagramFlossEquation).Stage(stager.stage)
			diagram.Name = compareAnalysis.Name + " Equation Diagram"
			diagram.Scale = 5.0
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

		// Compute Delta C, Delta P, Delta E and Beta
		var cFrom, cTo float64
		var pFrom, pTo float64
		var eFrom, eTo float64

		if compareAnalysis.FromSystem != nil {
			for _, c := range compareAnalysis.FromSystem.Complexities {
				cFrom += c.Strength
			}
			for _, p := range compareAnalysis.FromSystem.Performances {
				pFrom += p.Strength
			}
			for _, e := range compareAnalysis.FromSystem.Efforts {
				eFrom += e.Strength
			}
		}

		if compareAnalysis.ToSystem != nil {
			for _, c := range compareAnalysis.ToSystem.Complexities {
				cTo += c.Strength
			}
			for _, p := range compareAnalysis.ToSystem.Performances {
				pTo += p.Strength
			}
			for _, e := range compareAnalysis.ToSystem.Efforts {
				eTo += e.Strength
			}
		}

		deltaC := cTo - cFrom
		deltaP := pTo - pFrom
		deltaE := eTo - eFrom

		var computedBeta float64
		if deltaE != 0 {
			computedBeta = (compareAnalysis.Alpha*deltaP - deltaC) / deltaE
		} else {
			computedBeta = 0
		}

		if compareAnalysis.Beta == nil {
			compareAnalysis.Beta = new(float64)
			*compareAnalysis.Beta = computedBeta
			needCommit = true
		} else if *compareAnalysis.Beta != computedBeta {
			*compareAnalysis.Beta = computedBeta
			needCommit = true
		}
	}

	return
}
