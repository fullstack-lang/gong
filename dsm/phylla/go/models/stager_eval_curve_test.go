package models

import (
	"fmt"
	"math"
	"testing"
)

func svgCenter(x0, y0, x1, y1, R float64, sweepFlag bool) (float64, float64) {
	dx := (x0 - x1) / 2.0
	dy := (y0 - y1) / 2.0
	
	Rsq := R * R
	dsq := dx*dx + dy*dy
	if Rsq < dsq {
		R = math.Sqrt(dsq)
		Rsq = R*R
	}
	
	sign := 1.0
	if sweepFlag {
		sign = -1.0 // wait, let's test this
	}
	
	// in SVG, y goes down, but mathematically y goes up. 
	// Our evaluateCurveY is in mathematical coordinates! 
	// Wait! `computeArcV2Geometry` computes geometry for SVG!
	// Is it in mathematical coordinates?
	// `PerpendicularVector` coordinates are in MODEL coordinates (Y is up? Or Y is down?)
	// Usually Y is up in the model, and inverted in `stager_ux_svg_plant_diagram.go`.
	return 0, 0
}
