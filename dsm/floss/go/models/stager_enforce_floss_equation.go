package models

func getSystemCPE(sys *System) (c, p, e float64) {
	if sys == nil {
		return
	}
	effC, _ := sys.GetEffectiveComplexities()
	for _, comp := range effC {
		c += comp.Strength
	}
	effP, _ := sys.GetEffectivePerformances()
	for _, perf := range effP {
		p += perf.Strength
	}
	effE, _ := sys.GetEffectiveEfforts()
	for _, eff := range effE {
		e += eff.Strength
	}
	return
}

func (stager *Stager) enforceFlossEquation() (needCommit bool) {
	for compareAnalysis := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
		
		if compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil {
			C1, P1, E1 := getSystemCPE(compareAnalysis.FromSystem)
			C2, P2, E2 := getSystemCPE(compareAnalysis.ToSystem)

			D := P1*E2 - P2*E1
			if D != 0 {
				newMu := (C1*E2 - C2*E1) / D
				newEpsilon := (C1*P2 - C2*P1) / D

				if compareAnalysis.Mu != newMu {
					compareAnalysis.Mu = newMu
					needCommit = true
				}
				if compareAnalysis.Epsilon != newEpsilon {
					compareAnalysis.Epsilon = newEpsilon
					needCommit = true
				}
			}
		}

		// Enforce mu != 0
		if compareAnalysis.Mu == 0 {
			compareAnalysis.Mu = 1.0
			needCommit = true
		}

		// Enforce default epsilon if 0
		if compareAnalysis.Epsilon == 0 {
			compareAnalysis.Epsilon = 1.0
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
		if system.AreCPEsCompoundedFromSubSystems && len(system.SubSystems) > 0 {
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
