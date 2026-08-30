package models

import "math"

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagramEq := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		compareAnalysis := diagramEq.GetOwningCompareAnalysis()
		owningSystem := diagramEq.GetOwningSystem()
		if compareAnalysis == nil && owningSystem == nil {
			for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
				for _, d := range ca.DiagramFlossEquations {
					if d == diagramEq {
						compareAnalysis = ca
						break
					}
				}
			}
			if compareAnalysis == nil {
				for sys := range *GetGongstructInstancesSet[System](stager.stage) {
					for _, d := range sys.DiagramFlossEquations {
						if d == diagramEq {
							owningSystem = sys
							break
						}
					}
				}
			}
		}

		scale := diagramEq.Scale
		if scale <= 0 {
			scale = 5.0
		}

		extentAboveBaseline, extentBelowBaseline := computeFlossDiagramVerticalExtents(diagramEq, compareAnalysis, owningSystem)

		topMargin := 150.0
		bottomMargin := 120.0
		neededHeight := topMargin + extentAboveBaseline*scale + extentBelowBaseline*scale + bottomMargin
		if neededHeight < 750.0 {
			neededHeight = 750.0
		}

		boxWidth := diagramEq.GetDefaultBoxWidth()
		if boxWidth <= 0 {
			boxWidth = 250.0
		}

		xMargin := 80.0
		var columnsRight float64
		if compareAnalysis != nil && compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil && !diagramEq.IsInDelta3ColumnsMode {
			cWidth := math.Round(boxWidth * 0.65)
			if cWidth < 150.0 {
				cWidth = 150.0
			}
			pairGap := 10.0
			colSpacing := 35.0
			columnsRight = xMargin + 6*cWidth + 3*pairGap + 2*colSpacing
		} else {
			colSpacing := 40.0
			columnsRight = xMargin + 3*boxWidth + 2*colSpacing
		}
		rightIndicatorsMargin := 220.0
		neededWidth := columnsRight + rightIndicatorsMargin + 40.0

		updateDiagramSize(diagramEq.Note_Shapes, &neededWidth, &neededHeight)

		if diagramEq.Width != neededWidth {
			diagramEq.Width = neededWidth
			needCommit = true
		}
		if diagramEq.Height != neededHeight {
			diagramEq.Height = neededHeight
			needCommit = true
		}
	}

	return
}

func updateDiagramSize[T RectShapeInterface](shapes []T, width, height *float64) (needCommit bool) {
	for _, shape := range shapes {
		if shape.GetX()+shape.GetWidth() > *width {
			*width = shape.GetX() + shape.GetWidth()
			needCommit = true
		}
		if shape.GetY()+shape.GetHeight() > *height {
			*height = shape.GetY() + shape.GetHeight()
			needCommit = true
		}
	}
	return
}

func computeFlossDiagramVerticalExtents(
	diagram *DiagramFlossEquation,
	compareAnalysis *CompareAnalysis,
	owningSystem *System,
) (extentAboveBaseline, extentBelowBaseline float64) {
	if compareAnalysis != nil && compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil {
		var cTo, cFrom, pTo, pFrom, eTo, eFrom float64
		var fromComplexities, toComplexities []*Complexity
		var fromPerformances, toPerformances []*Performance
		var fromEfforts, toEfforts []*Effort

		if diagram.AreSubsystemsVisible && compareAnalysis.FromSystem.AreCPEsCompoundedFromSubSystems {
			fromComplexities, _ = compareAnalysis.FromSystem.GetEffectiveComplexities()
			fromPerformances, _ = compareAnalysis.FromSystem.GetEffectivePerformances()
			fromEfforts, _ = compareAnalysis.FromSystem.GetEffectiveEfforts()
		} else {
			fromComplexities = compareAnalysis.FromSystem.Complexities
			fromPerformances = compareAnalysis.FromSystem.Performances
			fromEfforts = compareAnalysis.FromSystem.Efforts
		}

		if diagram.AreSubsystemsVisible && compareAnalysis.ToSystem.AreCPEsCompoundedFromSubSystems {
			toComplexities, _ = compareAnalysis.ToSystem.GetEffectiveComplexities()
			toPerformances, _ = compareAnalysis.ToSystem.GetEffectivePerformances()
			toEfforts, _ = compareAnalysis.ToSystem.GetEffectiveEfforts()
		} else {
			toComplexities = compareAnalysis.ToSystem.Complexities
			toPerformances = compareAnalysis.ToSystem.Performances
			toEfforts = compareAnalysis.ToSystem.Efforts
		}

		if diagram.AreCommonElementsHidden {
			var dummyStager Stager
			fromComplexities, toComplexities, _, _ = dummyStager.filterCommonComplexities(
				fromComplexities, toComplexities, nil, nil,
			)
			fromPerformances, toPerformances, _, _ = dummyStager.filterCommonPerformances(
				fromPerformances, toPerformances, nil, nil,
			)
			fromEfforts, toEfforts, _, _ = dummyStager.filterCommonEfforts(
				fromEfforts, toEfforts, nil, nil,
			)
		}

		for _, c := range fromComplexities {
			cFrom += c.Strength
		}
		for _, c := range toComplexities {
			cTo += c.Strength
		}
		for _, p := range fromPerformances {
			pFrom += p.Strength
		}
		for _, p := range toPerformances {
			pTo += p.Strength
		}
		for _, e := range fromEfforts {
			eFrom += e.Strength
		}
		for _, e := range toEfforts {
			eTo += e.Strength
		}

		mu := compareAnalysis.Mu
		if mu == 0 {
			mu = 1.0
		}
		epsilon := compareAnalysis.Epsilon
		if epsilon == 0 {
			epsilon = 1.0
		}

		deltaC := cTo - cFrom
		deltaP := pTo - pFrom
		deltaE := eTo - eFrom
		rhs := mu*deltaP - epsilon*deltaE

		extentAboveBaseline = math.Max(cTo, mu*pTo)
		if !diagram.IsInDelta3ColumnsMode {
			extentAboveBaseline = math.Max(extentAboveBaseline, cFrom)
			extentAboveBaseline = math.Max(extentAboveBaseline, mu*pFrom)
		} else {
			if deltaP > 0 {
				extentAboveBaseline = math.Max(extentAboveBaseline, mu*deltaP)
			}
			if deltaC > 0 {
				extentAboveBaseline = math.Max(extentAboveBaseline, deltaC)
			}
		}

		extentBelowBaseline = math.Max(0, -deltaC)
		extentBelowBaseline = math.Max(extentBelowBaseline, -mu*deltaP)
		extentBelowBaseline = math.Max(extentBelowBaseline, -mu*deltaP+epsilon*eTo)
		extentBelowBaseline = math.Max(extentBelowBaseline, -rhs)
		if !diagram.IsInDelta3ColumnsMode {
			extentBelowBaseline = math.Max(extentBelowBaseline, epsilon*eTo)
			extentBelowBaseline = math.Max(extentBelowBaseline, epsilon*eFrom)
		} else {
			if deltaE > 0 {
				extentBelowBaseline = math.Max(extentBelowBaseline, -rhs)
			}
		}
	} else if owningSystem != nil {
		var cTo, pTo, eTo float64
		var toComplexities []*Complexity
		var toPerformances []*Performance
		var toEfforts []*Effort
		if diagram.AreSubsystemsVisible && owningSystem.AreCPEsCompoundedFromSubSystems {
			toComplexities, _ = owningSystem.GetEffectiveComplexities()
			toPerformances, _ = owningSystem.GetEffectivePerformances()
			toEfforts, _ = owningSystem.GetEffectiveEfforts()
		} else {
			toComplexities = owningSystem.Complexities
			toPerformances = owningSystem.Performances
			toEfforts = owningSystem.Efforts
		}
		for _, c := range toComplexities {
			cTo += c.Strength
		}
		for _, p := range toPerformances {
			pTo += p.Strength
		}
		for _, e := range toEfforts {
			eTo += e.Strength
		}
		extentAboveBaseline = math.Max(cTo, pTo)
		extentBelowBaseline = math.Max(0, eTo-pTo)
	}
	return
}
