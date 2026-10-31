package models

import "math"

type trapezeSegment struct {
	startTheta float64
	endTheta   float64
	startY     float64
	endY       float64
	isGentle   bool
	indices    []int
}

// ComputeVaseTrapezeConstant3DThickness adjusts the resampled bottom and top cylinder curves
// for VaseTrapeze so that:
//  1. The 3D wall thickness normal to the projected ring surfaces is constant everywhere
//     around the ring (both on the gentle M-segments and on the steep N-segments).
//  2. In ROTATED mode, adjacent rings fit surface-to-surface with zero gap along all
//     touching gentle M-segments.
//  3. All segments join with C0 value and C1 tangent continuity at every rhombus vertex.
func ComputeVaseTrapezeConstant3DThickness(
	plant *PlantAbstract,
	globalR float64,
	targetAngles []float64,
	baseBottomY []float64,
) (bottomY []float64, topY []float64, unrotatedStepDY, rotatedStepDX, rotatedStepDY float64, ok bool) {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil {
		return nil, nil, 0, 0, 0, false
	}
	K := len(targetAngles) - 1
	if K < 4 || len(baseBottomY) != len(targetAngles) || globalR <= 0 {
		return nil, nil, 0, 0, 0, false
	}
	if plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		return nil, nil, 0, 0, 0, false
	}

	vase := plant.TubeVaseAbstract
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}

	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	tanAlpha := math.Tan(projAngleRad)
	if math.Abs(tanAlpha) < 1e-6 {
		return nil, nil, 0, 0, 0, false
	}

	pGrid := plant.PerpendicularVectorGrid
	vFirst := pGrid.PerpendicularVectors[0]
	vx := vFirst.EndX - vFirst.StartX
	vy := vFirst.EndY - vFirst.StartY
	vLen := math.Hypot(vx, vy)
	if vLen == 0 {
		return nil, nil, 0, 0, 0, false
	}
	vx, vy = vx/vLen, vy/vLen
	if math.Abs(vy) < 1e-6 {
		return nil, nil, 0, 0, 0, false
	}

	sideLength := plant.RhombusSideLength
	verticalThickness := vase.RelativeVerticalThickness * sideLength
	rotatedSeparation := vase.RelativeRotatedTorusSeparation * sideLength

	var growthVectorX, growthVectorY float64
	if plant.GrowthVectorShape != nil {
		growthVectorX = plant.GrowthVectorShape.X
		growthVectorY = plant.GrowthVectorShape.Y
	}

	// Slope dY/dtheta at every rhombus vertex (where the 2D curve is orthogonal to (vx, vy))
	sV := (-vx / vy) * globalR * scaleY

	// Vertical separation at every rhombus vertex corresponding to normal 2D thickness
	deltaYV := (verticalThickness / vy) * scaleY

	unrotatedStepDY = deltaYV
	rotatedStepDX = growthVectorX
	rotatedStepDY = growthVectorY*scaleY + deltaYV + rotatedSeparation*scaleY
	rotStepDYContact := growthVectorY*scaleY + deltaYV
	rotStepTheta := growthVectorX / globalR

	p1H := vase.Z_Ribbon + vase.TopPlaneHeight*scaleY
	p2H := vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY
	yRef := 0.5 * (p1H + p2H)

	minY := math.MaxFloat64
	maxY := -math.MaxFloat64
	for _, y := range baseBottomY {
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	yMean := 0.5 * (minY + maxY)
	rRef := globalR + (yRef-yMean)*tanAlpha
	if rRef <= 1e-3 {
		rRef = globalR
	}

	kSlope := tanAlpha / rRef
	denomVertex := math.Hypot(1.0, sV*kSlope)
	deltaYNormal := func(slope float64) float64 {
		return deltaYV * math.Hypot(1.0, slope*kSlope) / denomVertex
	}

	thetaSector := targetAngles[K] - targetAngles[0]
	dTheta := thetaSector / float64(K)

	evalPeriodicY := func(arr []float64, theta float64) float64 {
		for theta < 0 {
			theta += thetaSector
		}
		for theta >= thetaSector {
			theta -= thetaSector
		}
		pos := theta / dTheta
		idx := int(math.Floor(pos))
		if idx < 0 {
			idx = 0
		}
		if idx >= K {
			idx = K - 1
		}
		frac := pos - float64(idx)
		return arr[idx]*(1.0-frac) + arr[idx+1]*frac
	}

	var segments []trapezeSegment
	x0 := vFirst.StartX
	for k := 0; k+1 < len(pGrid.PerpendicularVectors); k++ {
		v1 := pGrid.PerpendicularVectors[k]
		v2 := pGrid.PerpendicularVectors[k+1]
		t1 := (v1.StartX - x0) / globalR
		t2 := (v2.StartX - x0) / globalR
		if t1 >= thetaSector-1e-4 {
			break
		}
		if t2 > thetaSector {
			t2 = thetaSector
		}
		isGentle := v2.StartY >= v1.StartY
		var idxs []int
		for i, th := range targetAngles {
			if th >= t1-1e-7 && th <= t2+1e-7 {
				idxs = append(idxs, i)
			}
		}
		segments = append(segments, trapezeSegment{
			startTheta: t1,
			endTheta:   t2,
			startY:     evalPeriodicY(baseBottomY, t1),
			endY:       evalPeriodicY(baseBottomY, t2),
			isGentle:   isGentle,
			indices:    idxs,
		})
	}

	P := len(segments)
	if P == 0 {
		return nil, nil, 0, 0, 0, false
	}

	// Build a C3-smooth base curve across all segments that interpolates the exact
	// rhombus vertex positions and vertex slope sV with zero 2nd and 3rd derivatives
	// at every vertex, eliminating curvature-jump squirks at arc midpoints and vertices.
	smoothBaseY := make([]float64, len(targetAngles))
	copy(smoothBaseY, baseBottomY)
	for _, seg := range segments {
		span := seg.endTheta - seg.startTheta
		if span <= 1e-9 {
			continue
		}
		dySeg := seg.endY - seg.startY
		m0 := sV * span
		for _, idx := range seg.indices {
			u := (targetAngles[idx] - seg.startTheta) / span
			if u < 0 {
				u = 0
			} else if u > 1 {
				u = 1
			}
			u2 := u * u
			u3 := u2 * u
			u4 := u2 * u2
			s5 := u3 * (10.0 - 15.0*u + 6.0*u2)
			s7 := u4 * (35.0 - 84.0*u + 70.0*u2 - 20.0*u2*u)
			smoothBaseY[idx] = seg.startY + m0*u + (dySeg-m0)*(0.5*s5+0.5*s7)
		}
	}
	smoothBaseY[K] = smoothBaseY[0]

	// Update rRef using smoothBaseY mean height
	minSmoothY := math.MaxFloat64
	maxSmoothY := -math.MaxFloat64
	for _, y := range smoothBaseY {
		if y < minSmoothY {
			minSmoothY = y
		}
		if y > maxSmoothY {
			maxSmoothY = y
		}
	}
	rRef = globalR + (yRef-0.5*(minSmoothY+maxSmoothY))*tanAlpha
	if rRef <= 1e-3 {
		rRef = globalR
	}
	kSlope = tanAlpha / rRef
	denomVertex = math.Hypot(1.0, sV*kSlope)

	bottomY = make([]float64, len(targetAngles))
	topY = make([]float64, len(targetAngles))
	copy(bottomY, smoothBaseY)

	dBaseAll := make([]float64, len(targetAngles))
	for i := range K {
		iPrev := (i - 1 + K) % K
		iNext := (i + 1) % K
		dBaseAll[i] = (smoothBaseY[iNext] - smoothBaseY[iPrev]) / (2.0 * dTheta)
	}
	dBaseAll[K] = dBaseAll[0]

	// Initialize topY everywhere with exact 3D normal offset of baseBottomY
	for i := range K {
		topY[i] = bottomY[i] + deltaYNormal(dBaseAll[i])
	}
	topY[K] = topY[0]

	// Find the contiguous run of gentle M-segments following a steep N-segment
	firstGentle := -1
	for s := range P {
		prev := (s - 1 + P) % P
		if segments[s].isGentle && !segments[prev].isGentle {
			firstGentle = s
			break
		}
	}

	var gentleRun []int
	if firstGentle != -1 {
		for step := range P {
			s := (firstGentle + step) % P
			if !segments[s].isGentle {
				break
			}
			gentleRun = append(gentleRun, s)
		}
	} else {
		for s := range P {
			if segments[s].isGentle {
				gentleRun = append(gentleRun, s)
			}
		}
	}

	mRef := len(gentleRun) - 1

	updateAllTopY := func() {
		for i := range K {
			iPrev := (i - 1 + K) % K
			iNext := (i + 1) % K
			slope := (bottomY[iNext] - bottomY[iPrev]) / (2.0 * dTheta)
			topY[i] = bottomY[i] + deltaYNormal(slope)
		}
		bottomY[K] = bottomY[0]
		topY[K] = topY[0]
	}

	for outer := 0; outer < 6; outer++ {
		updateAllTopY()

		// Propagate backwards from mRef-1 down to 0 (e.g. M3 -> M2 -> M1):
		// bottomY on segment m matches topY on segment m+1 shifted by (-rotStepTheta, -rotStepDYContact)
		for rank := mRef - 1; rank >= 0; rank-- {
			seg := segments[gentleRun[rank]]
			for _, idx := range seg.indices {
				theta := targetAngles[idx]
				bottomY[idx%K] = evalPeriodicY(topY, theta+rotStepTheta) - rotStepDYContact
			}
			bottomY[K] = bottomY[0]
			updateAllTopY()
		}

		// Smoothly blend any endpoint offset across non-contacting steep N-segments
		for s := range P {
			if segments[s].isGentle || len(segments[s].indices) == 0 {
				continue
			}
			idxs := segments[s].indices
			iBefore := (idxs[0] - 1 + K) % K
			iAfter := (idxs[len(idxs)-1] + 1) % K
			dStart := bottomY[iBefore] - smoothBaseY[iBefore]
			dEnd := bottomY[iAfter] - smoothBaseY[iAfter]
			tStart := targetAngles[iBefore]
			tEnd := targetAngles[iAfter]
			if tEnd <= tStart {
				tEnd += thetaSector
			}
			span := tEnd - tStart
			if span > 1e-9 {
				for _, idx := range idxs {
					th := targetAngles[idx]
					if th < tStart {
						th += thetaSector
					}
					u := (th - tStart) / span
					bottomY[idx%K] = smoothBaseY[idx%K] + (1.0-u)*dStart + u*dEnd
				}
				bottomY[K] = bottomY[0]
			}
		}
	}

	updateAllTopY()

	return bottomY, topY, unrotatedStepDY, rotatedStepDX, rotatedStepDY, true
}
