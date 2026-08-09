package models

import (
	"math"
)

func evaluateCurveY(plant *PlantAbstract, isTop bool, x float64) float64 {
	bestY := -1e9
	if !isTop {
		if plant.GrowthCurve2D != nil && plant.GrowthCurve2D.StartHalfwayArcShapeGrid != nil {
			for _, sa := range plant.GrowthCurve2D.StartHalfwayArcShapeGrid.StartHalfwayArcShapes {
				if x >= math.Min(sa.StartX, sa.EndX) && x <= math.Max(sa.StartX, sa.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
					y := evalArcY(sa.StartX, sa.StartY, sa.EndX, sa.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
		if plant.GrowthCurve2D != nil && plant.GrowthCurve2D.EndHalfwayArcShapeGrid != nil {
			for _, ea := range plant.GrowthCurve2D.EndHalfwayArcShapeGrid.EndHalfwayArcShapes {
				if x >= math.Min(ea.StartX, ea.EndX) && x <= math.Max(ea.StartX, ea.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
					y := evalArcY(ea.StartX, ea.StartY, ea.EndX, ea.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
	} else {
		if plant.VaseAbstract != nil && plant.VaseAbstract.TopGrowthCurve2D != nil && plant.VaseAbstract.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != nil {
			for _, sa := range plant.VaseAbstract.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes {
				if x >= math.Min(sa.StartX, sa.EndX) && x <= math.Max(sa.StartX, sa.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
					y := evalArcY(sa.StartX, sa.StartY, sa.EndX, sa.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
		if plant.VaseAbstract != nil && plant.VaseAbstract.TopGrowthCurve2D != nil && plant.VaseAbstract.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid != nil {
			for _, ea := range plant.VaseAbstract.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes {
				if x >= math.Min(ea.StartX, ea.EndX) && x <= math.Max(ea.StartX, ea.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
					y := evalArcY(ea.StartX, ea.StartY, ea.EndX, ea.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
	}

	return bestY
}

func computeArcCenterFromEndpoints(x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) (float64, float64, float64) {
	dx := (x1 - x2) / 2.0
	dy := (y1 - y2) / 2.0
	d2 := dx*dx + dy*dy
	var cx, cy float64
	if d2 == 0 || r*r < d2 {
		cx = (x1 + x2) / 2.0
		cy = (y1 + y2) / 2.0
		r = math.Sqrt(d2)
	} else {
		root := math.Sqrt(r*r/d2 - 1.0)
		if largeArcFlag == sweepFlag {
			root = -root
		}
		cx = (x1+x2)/2.0 + root*dy
		cy = (y1+y2)/2.0 - root*dx
	}
	return cx, cy, r
}

func evalArcY(x0, y0, x1, y1, cx, cy, R, x float64) float64 {
	lineY := y0
	if math.Abs(x1-x0) > 1e-6 {
		lineY = y0 + (y1-y0)*(x-x0)/(x1-x0)
	} else {
		lineY = (y0 + y1) / 2
	}

	bestY := lineY
	minDist := math.MaxFloat64

	val := R*R - (x-cx)*(x-cx)
	if val >= 0 {
		y_a := cy + math.Sqrt(val)
		y_b := cy - math.Sqrt(val)
		if math.Abs(y_a-lineY) < minDist {
			minDist = math.Abs(y_a - lineY)
			bestY = y_a
		}
		if math.Abs(y_b-lineY) < minDist {
			minDist = math.Abs(y_b - lineY)
			bestY = y_b
		}
	}
	return bestY
}

func ComputePartiallyGrowthCurveDYForRatio(plant *PlantAbstract, rotationRatio float64) (dx float64, dy float64, currentDX float64) {
	if plant == nil || plant.RhombusStuff == nil || plant.RhombusStuff.PlantCircumferenceShape == nil || plant.GrowthVectorShape == nil || plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) == 0 {
		return 0, 0, 0
	}

	circLen := plant.RhombusStuff.PlantCircumferenceShape.Length
	if circLen <= 0 {
		return 0, 0, 0
	}

	vThickness := 0.0
	rotatedSeparation := 0.0
	if plant.PlantType == Vase {
		vThickness = plant.VaseAbstract.RelativeVerticalThickness * plant.RhombusSideLength
		rotatedSeparation = plant.VaseAbstract.RelativeRotatedTorusSeparation * plant.RhombusSideLength
	}
	var vx, vy float64
	if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
		vx = vFirst.EndX - vFirst.StartX
		vy = vFirst.EndY - vFirst.StartY
		vLen := math.Hypot(vx, vy)
		if vLen > 0 {
			vx, vy = vx/vLen, vy/vLen
		} else {
			vx, vy = 0, 1
		}
	} else {
		vx, vy = 0, 1
	}

	perpDX := vThickness * vx

	dx = rotationRatio*plant.GrowthVectorShape.X + perpDX

	currentDX = math.Mod(dx, circLen)
	if currentDX < 0 {
		currentDX += circLen
	}

	minX := plant.PerpendicularVectorGrid.PerpendicularVectors[0].StartX
	n := len(plant.PerpendicularVectorGrid.PerpendicularVectors)
	maxX := plant.PerpendicularVectorGrid.PerpendicularVectors[n-1].StartX

	dy = -1e9
	steps := 3600
	// The 3D torus is periodic. Check the entire circumference to find the global max dy!
	for step := 0; step <= steps; step++ {
		x := minX + (float64(step)/float64(steps))*(maxX-minX)
		yTop := evaluateCurveY(plant, true, x)

		xBot := x - currentDX
		// Wrap xBot so it stays within [minX, maxX] due to the cylindrical geometry
		for xBot < minX {
			xBot += circLen
		}
		for xBot > maxX {
			xBot -= circLen
		}

		yBot := evaluateCurveY(plant, false, xBot)

		if yTop != -1e9 && yBot != -1e9 {
			if yTop-yBot > dy {
				dy = yTop - yBot
			}
		}
	}

	if dy == -1e9 {
		dy = vThickness
	}

	dy += rotatedSeparation

	return dx, dy, currentDX
}

func ComputePartiallyGrowthCurveDY(plant *PlantAbstract) (dx float64, dy float64, currentDX float64) {
	rotRatio := 0.0
	if plant.PlantType == Vase {
		rotRatio = plant.VaseAbstract.RotationRatio
	}
	return ComputePartiallyGrowthCurveDYForRatio(plant, rotRatio)
}

func ComputeStackHeightForRotationRatio(plant *PlantAbstract, rotationRatio float64) float64 {
	stackHeight := plant.StackHeight
	if stackHeight <= 0 {
		return 0.0
	}
	_, baseDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, 0.0)
	cumDY := baseDY

	numSteps := stackHeight - 1
	if numSteps > 0 {
		totalProgress := rotationRatio * float64(numSteps)
		for k := 1; k <= numSteps; k++ {
			var r_k float64
			kFloat := float64(numSteps - k + 1)
			if totalProgress >= kFloat {
				r_k = 1.0
			} else if totalProgress <= kFloat-1.0 {
				r_k = 0.0
			} else {
				r_k = totalProgress - (kFloat - 1.0)
			}
			_, stepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, r_k)
			cumDY += stepDY
		}
	}
	return cumDY
}
