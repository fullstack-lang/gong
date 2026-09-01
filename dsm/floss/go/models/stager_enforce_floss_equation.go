package models

import "math"

func roundTo1Decimal(val float64) float64 {
	return math.Round(val*10.0) / 10.0
}

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
	// Enforce 1-decimal precision on CPE items
	for c := range *GetGongstructInstancesSet[Complexity](stager.stage) {
		rounded := roundTo1Decimal(c.Strength)
		if c.Strength != rounded {
			c.Strength = rounded
			needCommit = true
		}
	}
	for p := range *GetGongstructInstancesSet[Performance](stager.stage) {
		rounded := roundTo1Decimal(p.Strength)
		if p.Strength != rounded {
			p.Strength = rounded
			needCommit = true
		}
	}
	for e := range *GetGongstructInstancesSet[Effort](stager.stage) {
		rounded := roundTo1Decimal(e.Strength)
		if e.Strength != rounded {
			e.Strength = rounded
			needCommit = true
		}
	}

	// Enforce 1-decimal precision on diagram scale
	for d := range *GetGongstructInstancesSet[DiagramFlossEquation](stager.stage) {
		if d.Scale != 0 {
			rounded := roundTo1Decimal(d.Scale)
			if d.Scale != rounded {
				d.Scale = rounded
				needCommit = true
			}
		}
	}

	for compareAnalysis := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {

		if compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil {
			C1, P1, E1 := getSystemCPE(compareAnalysis.FromSystem)
			C2, P2, E2 := getSystemCPE(compareAnalysis.ToSystem)

			D := P1*E2 - P2*E1
			if D != 0 {
				newMu := roundTo1Decimal((C1*E2 - C2*E1) / D)
				newEpsilon := roundTo1Decimal((C1*P2 - C2*P1) / D)

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

		// Enforce mu != 0 and 1-decimal precision
		if compareAnalysis.Mu == 0 {
			compareAnalysis.Mu = 1.0
			needCommit = true
		} else {
			rounded := roundTo1Decimal(compareAnalysis.Mu)
			if compareAnalysis.Mu != rounded {
				compareAnalysis.Mu = rounded
				needCommit = true
			}
		}

		// Enforce default epsilon if 0 and 1-decimal precision
		if compareAnalysis.Epsilon == 0 {
			compareAnalysis.Epsilon = 1.0
			needCommit = true
		} else {
			rounded := roundTo1Decimal(compareAnalysis.Epsilon)
			if compareAnalysis.Epsilon != rounded {
				compareAnalysis.Epsilon = rounded
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
				sumC = roundTo1Decimal(sumC)
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
				sumP = roundTo1Decimal(sumP)
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
				sumE = roundTo1Decimal(sumE)
				if system.Efforts[0].Strength != sumE {
					system.Efforts[0].Strength = sumE
					needCommit = true
				}
			}
		}
	}

	return
}
