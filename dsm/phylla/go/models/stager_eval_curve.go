package models

import (
	"math"
)

func evaluateCurveY(plant *Plant, isTop bool, x float64) float64 {
	bestY := -1e9
	if !isTop {
		if plant.StartHalfwayArcShapeGrid != nil {
			for _, sa := range plant.StartHalfwayArcShapeGrid.StartHalfwayArcShapes {
				if x >= math.Min(sa.StartX, sa.EndX) && x <= math.Max(sa.StartX, sa.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
					y := evalArcY(sa.StartX, sa.StartY, sa.EndX, sa.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
		if plant.EndHalfwayArcShapeGrid != nil {
			for _, ea := range plant.EndHalfwayArcShapeGrid.EndHalfwayArcShapes {
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
		if plant.TopStartHalfwayArcShapeGrid != nil {
			for _, sa := range plant.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes {
				if x >= math.Min(sa.StartX, sa.EndX) && x <= math.Max(sa.StartX, sa.EndX) {
					cx, cy, r := computeArcCenterFromEndpoints(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
					y := evalArcY(sa.StartX, sa.StartY, sa.EndX, sa.EndY, cx, cy, r, x)
					if y > bestY {
						bestY = y
					}
				}
			}
		}
		if plant.TopEndHalfwayArcShapeGrid != nil {
			for _, ea := range plant.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes {
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

func ComputePartiallyGrowthCurveDY(plant *Plant) (dx float64, dy float64, currentDX float64) {
	circLen := plant.RhombusStuff.PlantCircumferenceShape.Length
	if circLen <= 0 {
		return 0, 0, 0
	}

	vThickness := plant.RelativeVerticalThickness * plant.RhombusSideLength
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
	perpDY := vThickness * vy

	dx = plant.RotationRatio*plant.GrowthVectorShape.X + perpDX
	dy = plant.RotationRatio*plant.GrowthVectorShape.Y + perpDY

	currentDX = math.Mod(dx, circLen)
	if currentDX < 0 {
		currentDX += circLen
	}

	return dx, dy, currentDX
}
