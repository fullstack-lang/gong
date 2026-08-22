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

		maxVal := 50.0
		if compareAnalysis != nil && compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil {
			var cTo, cFrom, pTo, pFrom, eTo, eFrom float64
			var fromComplexities, toComplexities []*Complexity
			var fromPerformances, toPerformances []*Performance
			var fromEfforts, toEfforts []*Effort

			if diagramEq.AreSubsystemsVisible && compareAnalysis.FromSystem.AreCPEsCompoundedFromSubSystems {
				fromComplexities, _ = compareAnalysis.FromSystem.GetEffectiveComplexities()
				fromPerformances, _ = compareAnalysis.FromSystem.GetEffectivePerformances()
				fromEfforts, _ = compareAnalysis.FromSystem.GetEffectiveEfforts()
			} else {
				fromComplexities = compareAnalysis.FromSystem.Complexities
				fromPerformances = compareAnalysis.FromSystem.Performances
				fromEfforts = compareAnalysis.FromSystem.Efforts
			}

			if diagramEq.AreSubsystemsVisible && compareAnalysis.ToSystem.AreCPEsCompoundedFromSubSystems {
				toComplexities, _ = compareAnalysis.ToSystem.GetEffectiveComplexities()
				toPerformances, _ = compareAnalysis.ToSystem.GetEffectivePerformances()
				toEfforts, _ = compareAnalysis.ToSystem.GetEffectiveEfforts()
			} else {
				toComplexities = compareAnalysis.ToSystem.Complexities
				toPerformances = compareAnalysis.ToSystem.Performances
				toEfforts = compareAnalysis.ToSystem.Efforts
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

			alpha := compareAnalysis.Alpha
			if alpha == 0 {
				alpha = 1.0
			}
			beta := compareAnalysis.Beta
			if beta == 0 {
				beta = 1.0
			}

			maxVal = math.Max(math.Max(math.Max(cTo, cFrom), math.Max(alpha*pTo, alpha*pFrom)), math.Max(beta*eTo, beta*eFrom))
			maxVal = math.Max(maxVal, math.Abs(cFrom-cTo))
		} else if owningSystem != nil {
			var cTo, pTo float64
			var toComplexities []*Complexity
			var toPerformances []*Performance
			if diagramEq.AreSubsystemsVisible && owningSystem.AreCPEsCompoundedFromSubSystems {
				toComplexities, _ = owningSystem.GetEffectiveComplexities()
				toPerformances, _ = owningSystem.GetEffectivePerformances()
			} else {
				toComplexities = owningSystem.Complexities
				toPerformances = owningSystem.Performances
			}
			for _, c := range toComplexities {
				cTo += c.Strength
			}
			for _, p := range toPerformances {
				pTo += p.Strength
			}
			maxVal = math.Max(pTo, cTo)
		}

		neededHeight := 180.0 + maxVal*scale + 140.0
		if neededHeight < 800.0 {
			neededHeight = 800.0
		}

		boxWidth := diagramEq.GetDefaultBoxWidth()
		if boxWidth <= 0 {
			boxWidth = 250.0
		}

		xMargin := 80.0
		var columnsRight float64
		if compareAnalysis != nil && compareAnalysis.FromSystem != nil && compareAnalysis.ToSystem != nil && !diagramEq.IsInDelta3ColumnsMode {
			pairGap := 15.0
			colSpacing := 50.0
			columnsRight = xMargin + 6*boxWidth + 3*pairGap + 2*colSpacing
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
