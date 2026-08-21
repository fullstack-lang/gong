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

	for system := range *GetGongstructInstancesSet[System](stager.stage) {
		if system.AreCPEsCompoundedFromSubSystems && len(system.SubSystemes) > 0 {
			// Enforce C1 = sum of subsystem complexities
			if len(system.Complexities) == 1 {
				effC, _ := system.GetEffectiveComplexities()
				var sumC float64
				for _, c := range effC {
					sumC += c.Strength
				}
				if system.Complexities[0].Strength != sumC {
					system.Complexities[0].Strength = sumC
					needCommit = true
				}
			}

			// Enforce P1 = sum of subsystem performances
			if len(system.Performances) == 1 {
				effP, _ := system.GetEffectivePerformances()
				var sumP float64
				for _, p := range effP {
					sumP += p.Strength
				}
				if system.Performances[0].Strength != sumP {
					system.Performances[0].Strength = sumP
					needCommit = true
				}
			}

			// Enforce E1 = sum of subsystem efforts
			if len(system.Efforts) == 1 {
				effE, _ := system.GetEffectiveEfforts()
				var sumE float64
				for _, e := range effE {
					sumE += e.Strength
				}
				if system.Efforts[0].Strength != sumE {
					system.Efforts[0].Strength = sumE
					needCommit = true
				}
			}
		}
	}

	return
}
