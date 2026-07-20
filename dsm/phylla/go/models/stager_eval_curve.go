package models

import (
	"math"
)

func evaluateCurveY(plant *Plant, isTop bool, x float64) float64 {
	pGrid := plant.PerpendicularVectorGrid
	n := len(pGrid.PerpendicularVectors) - 1

	offset := 0.0
	if isTop {
		offset = plant.RelativeVerticalThickness * plant.RhombusSideLength
	}

	for i := 0; i < n; i++ {
		v1 := pGrid.PerpendicularVectors[i]
		v2 := pGrid.PerpendicularVectors[i+1]

		startX, startY, endX, endY, R, _, _, _, _ := computeArcV2Geometry(v1, v2, offset, false)
		if x >= math.Min(startX, endX) && x <= math.Max(startX, endX) {
			cx, cy := computeArcV2Center(v1, v2, false)
			return evalArcY(startX, startY, endX, endY, cx, cy, R, x)
		}

		startX2, startY2, endX2, endY2, R2, _, _, _, _ := computeArcV2Geometry(v1, v2, offset, true)
		if x >= math.Min(startX2, endX2) && x <= math.Max(startX2, endX2) {
			cx, cy := computeArcV2Center(v1, v2, true)
			return evalArcY(startX2, startY2, endX2, endY2, cx, cy, R2, x)
		}
	}

	// fallback if out of bounds
	return -1e9
}

func computeArcV2Center(v1, v2 *PerpendicularVector, isEndArc bool) (float64, float64) {
	dx := v1.EndX - v1.StartX
	dy := v1.EndY - v1.StartY
	length := math.Hypot(dx, dy)
	if length == 0 {
		length = 1
	}
	ux, uy := dx/length, dy/length

	midX := (v1.StartX + v2.StartX) / 2.0
	midY := (v1.StartY + v2.StartY) / 2.0

	Vx := v1.StartX - midX
	Vy := v1.StartY - midY
	V_sq := Vx*Vx + Vy*Vy
	V_dot_u := Vx*ux + Vy*uy
	if math.Abs(V_dot_u) < 1e-6 {
		if V_dot_u >= 0 {
			V_dot_u = 1e-6
		} else {
			V_dot_u = -1e-6
		}
	}

	R_val := -V_sq / (2.0 * V_dot_u)
	cx := v1.StartX + R_val*ux
	cy := v1.StartY + R_val*uy

	if !isEndArc {
		return cx, cy
	}
	
	cx_new := 2*midX - cx
	cy_new := 2*midY - cy
	return cx_new, cy_new
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
