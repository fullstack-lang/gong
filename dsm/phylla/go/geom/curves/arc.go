package curves

import "math"

// ComputeArcCenterFromEndpoints computes the center (cx, cy) and radius r of an arc
// given its endpoints, nominal radius, sweepFlag and largeArcFlag (matching SVG arc definition).
func ComputeArcCenterFromEndpoints(x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) (float64, float64, float64) {
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

// EvalArcY evaluates the vertical Y value of an arc at a given X coordinate.
func EvalArcY(x0, y0, x1, y1, cx, cy, R, x float64) float64 {
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
