package models

import "math"

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramFloss](stager.stage) {

		width := 0.0
		height := 0.0

		updateDiagramSize(diagram.System_Shapes, &width, &height)
		updateDiagramSize(diagram.Complexity_Shapes, &width, &height)
		updateDiagramSize(diagram.Performance_Shapes, &width, &height)
		updateDiagramSize(diagram.Effort_Shapes, &width, &height)
		updateDiagramSize(diagram.Note_Shapes, &width, &height)

		margin := 300.0
		width += margin
		height += margin

		if width != diagram.Width {
			diagram.Width = width
			needCommit = true
		}

		if height != diagram.Height {
			diagram.Height = height
			needCommit = true
		}
	}

	for _, diagramEq := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		compareAnalysis := diagramEq.GetOwningCompareAnalysis()
		if compareAnalysis == nil {
			for ca := range *GetGongstructInstancesSet[CompareAnalysis](stager.stage) {
				for _, d := range ca.DiagramFlossEquations {
					if d == diagramEq {
						compareAnalysis = ca
						break
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

			alpha := compareAnalysis.Alpha
			if alpha == 0 {
				alpha = 1.0
			}
			beta := compareAnalysis.Beta
			_ = deltaE
			_ = beta

			maxVal = math.Max(math.Abs(alpha*deltaP), math.Abs(deltaC))
		}

		neededHeight := 180.0 + maxVal*scale + 140.0
		if neededHeight < 750.0 {
			neededHeight = 750.0
		}


		neededWidth := 1050.0
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
