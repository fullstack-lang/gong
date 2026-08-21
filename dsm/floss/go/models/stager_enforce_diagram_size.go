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
			for _, c := range compareAnalysis.FromSystem.Complexities {
				cFrom += c.Strength
			}
			for _, c := range compareAnalysis.ToSystem.Complexities {
				cTo += c.Strength
			}
			for _, p := range compareAnalysis.FromSystem.Performances {
				pFrom += p.Strength
			}
			for _, p := range compareAnalysis.ToSystem.Performances {
				pTo += p.Strength
			}
			for _, e := range compareAnalysis.FromSystem.Efforts {
				eFrom += e.Strength
			}
			for _, e := range compareAnalysis.ToSystem.Efforts {
				eTo += e.Strength
			}

			deltaC := cTo - cFrom
			deltaP := pTo - pFrom
			deltaE := eTo
			_ = eFrom
			_ = deltaE

			alpha := compareAnalysis.Alpha
			if alpha == 0 {
				alpha = 1.0
			}

			maxVal = math.Max(math.Abs(alpha*deltaP), math.Abs(deltaC))
		} else if owningSystem != nil {
			var cTo, pTo float64
			for _, c := range owningSystem.Complexities {
				cTo += c.Strength
			}
			for _, p := range owningSystem.Performances {
				pTo += p.Strength
			}
			maxVal = math.Max(pTo, cTo)
		}

		neededHeight := 180.0 + maxVal*scale + 140.0
		if neededHeight < 750.0 {
			neededHeight = 750.0
		}

		boxWidth := diagramEq.GetDefaultBoxWidth()
		if boxWidth <= 0 {
			boxWidth = 250.0
		}
		colSpacing := 40.0
		xMargin := 80.0
		columnsRight := xMargin + 3*boxWidth + 2*colSpacing
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
