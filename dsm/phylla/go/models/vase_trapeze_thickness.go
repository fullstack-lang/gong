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
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil || !plant.TubeVaseAbstract.HasConstantThickness {
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

	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * sideLength
	p2H := (vase.Z_Ribbon + vase.BottomPlaneHeight*scaleY) * sideLength
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

	// Convolve the circular-arc deviation from each segment's chord with a narrow continuous
	// Gaussian on the circle u in R/Z via its Fourier sine series sin(2*pi*k*u).
	// On the gentle M-segments, no harmonic is artificially altered after Gaussian damping,
	// so the flat constant curvature of the circular arcs is 100% preserved.
	smoothBaseY := make([]float64, len(targetAngles))
	copy(smoothBaseY, baseBottomY)
	const numHarmonics = 12
	const kCutGentle = 4.5
	const kCutSteep = 3.2

	// First compute the natural endpoint slope sVEff of the Gaussian-mollified gentle M-segments
	sVEff := sV
	for _, seg := range segments {
		if !seg.isGentle {
			continue
		}
		span := seg.endTheta - seg.startTheta
		nPts := len(seg.indices)
		if span <= 1e-9 || nPts < 4 {
			continue
		}
		dySeg := seg.endY - seg.startY
		rawSlope := 0.0
		for k := 1; k <= numHarmonics; k++ {
			sum := 0.0
			for j := 0; j+1 < nPts; j++ {
				i0 := seg.indices[j]
				i1 := seg.indices[j+1]
				u0 := (targetAngles[i0] - seg.startTheta) / span
				u1 := (targetAngles[i1] - seg.startTheta) / span
				if u0 < 0 {
					u0 = 0
				}
				if u1 > 1 {
					u1 = 1
				}
				du := u1 - u0
				if du <= 0 {
					continue
				}
				phi0 := baseBottomY[i0] - (seg.startY + u0*dySeg)
				phi1 := baseBottomY[i1] - (seg.startY + u1*dySeg)
				f0 := phi0 * math.Sin(2.0*math.Pi*float64(k)*u0)
				f1 := phi1 * math.Sin(2.0*math.Pi*float64(k)*u1)
				sum += 0.5 * (f0 + f1) * du
			}
			damp := math.Exp(-0.5 * (float64(k*k) / (kCutGentle * kCutGentle)))
			bK := 2.0 * sum * damp
			rawSlope += 2.0 * math.Pi * float64(k) * bK
		}
		sVEff = (dySeg + rawSlope) / span
		break
	}

	for _, seg := range segments {
		span := seg.endTheta - seg.startTheta
		nPts := len(seg.indices)
		if span <= 1e-9 || nPts < 4 {
			continue
		}
		dySeg := seg.endY - seg.startY
		kCut := kCutGentle
		if !seg.isGentle {
			kCut = kCutSteep
		}

		bCoeffs := make([]float64, numHarmonics+1)
		rawSlope := 0.0
		for k := 1; k <= numHarmonics; k++ {
			sum := 0.0
			for j := 0; j+1 < nPts; j++ {
				i0 := seg.indices[j]
				i1 := seg.indices[j+1]
				u0 := (targetAngles[i0] - seg.startTheta) / span
				u1 := (targetAngles[i1] - seg.startTheta) / span
				if u0 < 0 {
					u0 = 0
				}
				if u1 > 1 {
					u1 = 1
				}
				du := u1 - u0
				if du <= 0 {
					continue
				}
				phi0 := baseBottomY[i0] - (seg.startY + u0*dySeg)
				phi1 := baseBottomY[i1] - (seg.startY + u1*dySeg)
				f0 := phi0 * math.Sin(2.0*math.Pi*float64(k)*u0)
				f1 := phi1 * math.Sin(2.0*math.Pi*float64(k)*u1)
				sum += 0.5 * (f0 + f1) * du
			}
			damp := math.Exp(-0.5 * (float64(k*k) / (kCut * kCut)))
			bCoeffs[k] = 2.0 * sum * damp
			rawSlope += 2.0 * math.Pi * float64(k) * bCoeffs[k]
		}

		// Match the endpoint slope sVEff across all segment junctions
		targetPhiSlope := sVEff*span - dySeg
		slopeDiff := targetPhiSlope - rawSlope
		bCoeffs[1] += slopeDiff / (2.0 * math.Pi)

		for _, idx := range seg.indices {
			u := (targetAngles[idx] - seg.startTheta) / span
			if u < 0 {
				u = 0
			} else if u > 1 {
				u = 1
			}
			phi := 0.0
			for k := 1; k <= numHarmonics; k++ {
				phi += bCoeffs[k] * math.Sin(2.0*math.Pi*float64(k)*u)
			}
			smoothBaseY[idx] = seg.startY + u*dySeg + phi
		}
	}
	smoothBaseY[K] = smoothBaseY[0]
	sV = sVEff

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

	// Build chains of gentle M-segments connected by the rotated growth shift (+rotStepTheta)
	nextGentle := make([]int, P)
	prevGentle := make([]int, P)
	for s := range P {
		nextGentle[s] = -1
		prevGentle[s] = -1
	}
	for s := range P {
		if !segments[s].isGentle {
			continue
		}
		targetStart := math.Mod(segments[s].startTheta+rotStepTheta, thetaSector)
		if targetStart < 0 {
			targetStart += thetaSector
		}
		for cand := range P {
			if cand == s || !segments[cand].isGentle {
				continue
			}
			diff := math.Abs(segments[cand].startTheta - targetStart)
			if diff > 0.5*thetaSector {
				diff = thetaSector - diff
			}
			if diff < 2.0*dTheta+1e-3 {
				nextGentle[s] = cand
				prevGentle[cand] = s
				break
			}
		}
	}

	var gentleChains [][]int
	visitedGentle := make([]bool, P)
	for s := range P {
		if segments[s].isGentle && prevGentle[s] == -1 {
			var chain []int
			cur := s
			for cur != -1 && !visitedGentle[cur] {
				visitedGentle[cur] = true
				chain = append(chain, cur)
				cur = nextGentle[cur]
			}
			if len(chain) > 0 {
				gentleChains = append(gentleChains, chain)
			}
		}
	}
	for s := range P {
		if segments[s].isGentle && !visitedGentle[s] {
			var chain []int
			cur := s
			for cur != -1 && !visitedGentle[cur] {
				visitedGentle[cur] = true
				chain = append(chain, cur)
				cur = nextGentle[cur]
			}
			if len(chain) > 0 {
				gentleChains = append(gentleChains, chain)
			}
		}
	}

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

	solveSegmentBottomFromTargetTop := func(idxs []int, targetTop []float64) {
		n := len(idxs)
		if n < 3 {
			return
		}
		a := make([]float64, n)
		b := make([]float64, n)
		c := make([]float64, n)
		rhs := make([]float64, n)
		for iter := 0; iter < 30; iter++ {
			maxRes := 0.0
			for j := 0; j < n; j++ {
				i := idxs[j] % K
				iPrev := (i - 1 + K) % K
				iNext := (i + 1) % K
				slope := (bottomY[iNext] - bottomY[iPrev]) / (2.0 * dTheta)
				dn := deltaYNormal(slope)
				res := bottomY[i] + dn - targetTop[j]
				if math.Abs(res) > maxRes {
					maxRes = math.Abs(res)
				}
				ddnDslope := (deltaYV / denomVertex) * (slope * kSlope * kSlope) / math.Hypot(1.0, slope*kSlope)
				coeff := ddnDslope / (2.0 * dTheta)
				if j > 0 {
					a[j] = -coeff
				} else {
					a[j] = 0
				}
				b[j] = 1.0
				if j+1 < n {
					c[j] = coeff
				} else {
					c[j] = 0
				}
				rhs[j] = -res
			}
			if maxRes < 1e-6 {
				break
			}
			for j := 1; j < n; j++ {
				w := a[j] / b[j-1]
				b[j] -= w * c[j-1]
				rhs[j] -= w * rhs[j-1]
			}
			delta := make([]float64, n)
			delta[n-1] = rhs[n-1] / b[n-1]
			for j := n - 2; j >= 0; j-- {
				delta[j] = (rhs[j] - c[j]*delta[j+1]) / b[j]
			}
			for j := 0; j < n; j++ {
				i := idxs[j] % K
				bottomY[i] += 0.6 * delta[j]
			}
			bottomY[K] = bottomY[0]
		}
	}

	for outer := 0; outer < 6; outer++ {
		updateAllTopY()

		for _, gentleRun := range gentleChains {
			mRef := len(gentleRun) / 2

			// Propagate backwards from mRef-1 down to 0 (e.g. M2 -> M1):
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

			// Propagate forwards from mRef+1 up to len(gentleRun)-1 (e.g. M2 -> M3):
			// topY on segment m matches bottomY on segment m-1 shifted by (+rotStepTheta, +rotStepDYContact)
			for rank := mRef + 1; rank < len(gentleRun); rank++ {
				seg := segments[gentleRun[rank]]
				targetTop := make([]float64, len(seg.indices))
				for j, idx := range seg.indices {
					theta := targetAngles[idx]
					targetTop[j] = evalPeriodicY(bottomY, theta-rotStepTheta) + rotStepDYContact
				}
				solveSegmentBottomFromTargetTop(seg.indices, targetTop)
				updateAllTopY()
			}
		}

		// Smoothly blend any endpoint value and slope offset across non-contacting steep N-segments
		// using a C2 quintic Hermite correction on top of the circular-arc smoothBaseY curve.
		for s := range P {
			if segments[s].isGentle || len(segments[s].indices) == 0 {
				continue
			}
			idxs := segments[s].indices
			iBefore := (idxs[0] - 1 + K) % K
			iAfter := (idxs[len(idxs)-1] + 1) % K
			dStart := bottomY[iBefore] - smoothBaseY[iBefore]
			dEnd := bottomY[iAfter] - smoothBaseY[iAfter]

			iBeforePrev := (iBefore - 1 + K) % K
			iAfterNext := (iAfter + 1) % K
			slopeDiffStart := ((bottomY[iBefore] - smoothBaseY[iBefore]) - (bottomY[iBeforePrev] - smoothBaseY[iBeforePrev])) / dTheta
			slopeDiffEnd := ((bottomY[iAfterNext] - smoothBaseY[iAfterNext]) - (bottomY[iAfter] - smoothBaseY[iAfter])) / dTheta

			tStart := targetAngles[iBefore]
			tEnd := targetAngles[iAfter]
			if tEnd <= tStart {
				tEnd += thetaSector
			}
			span := tEnd - tStart
			if span > 1e-9 {
				m0 := slopeDiffStart * span
				m1 := slopeDiffEnd * span
				for _, idx := range idxs {
					th := targetAngles[idx]
					if th < tStart {
						th += thetaSector
					}
					u := (th - tStart) / span
					u2 := u * u
					u3 := u2 * u
					u4 := u2 * u2
					u5 := u3 * u2
					h00 := 1.0 - 10.0*u3 + 15.0*u4 - 6.0*u5
					h10 := u - 6.0*u3 + 8.0*u4 - 3.0*u5
					h01 := 10.0*u3 - 15.0*u4 + 6.0*u5
					h11 := -4.0*u3 + 7.0*u4 - 3.0*u5
					corr := h00*dStart + h10*m0 + h01*dEnd + h11*m1
					bottomY[idx%K] = smoothBaseY[idx%K] + corr
				}
				bottomY[K] = bottomY[0]
			}
		}
	}

	updateAllTopY()

	return bottomY, topY, unrotatedStepDY, rotatedStepDX, rotatedStepDY, true
}

// ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion computes, for a VaseTrapeze plant,
// the proportion of the interior surface that is not occupied by the ring above (measured at
// the ring top) relative to the interior surface of the ring, computed as if
// CarvedOutTopRingsParameter is 0 (flat ring top).
func ComputeVaseTrapezeUnoccupiedInteriorSurfaceProportion(
	plant *PlantAbstract,
) (unoccupiedSurface, interiorSurface, proportion float64, ok bool) {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil {
		return 0, 0, 0, false
	}
	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)

	circumference := 10.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		pGrid := plant.PerpendicularVectorGrid
		first := pGrid.PerpendicularVectors[0]
		last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
		circumference = last.StartX - first.StartX
	}
	if circumference <= 0 {
		circumference = 10.0
	}
	globalR := circumference * float64(radialRepetitions) / (2 * math.Pi)

	gc := plant.GrowthCurve2D
	tgc := vase.TopGrowthCurve2D
	if gc == nil || tgc == nil ||
		gc.StartHalfwayArcShapeGrid == nil ||
		tgc.TopStartHalfwayArcShapeGrid == nil ||
		len(gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) == 0 ||
		len(tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) == 0 {
		return 0, 0, 0, false
	}

	startArcs := gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes
	var endArcs []*EndHalfwayArcShape
	if gc.EndHalfwayArcShapeGrid != nil {
		endArcs = gc.EndHalfwayArcShapeGrid.EndHalfwayArcShapes
	}
	topStartArcs := tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes
	var topEndArcs []*TopEndHalfwayArcShape
	if tgc.TopEndHalfwayArcShapeGrid != nil {
		topEndArcs = tgc.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes
	}

	var curvePoints, topCurvePoints []vector3
	appendArcPoints := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
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
		startAngle := math.Atan2(y1-cy, x1-cx)
		endAngle := math.Atan2(y2-cy, x2-cx)
		if sweepFlag {
			for endAngle < startAngle {
				endAngle += 2 * math.Pi
			}
		} else {
			for endAngle > startAngle {
				endAngle -= 2 * math.Pi
			}
		}
		const steps = 50
		for i := 0; i <= steps; i++ {
			if i == 0 && len(*targetCurve) > 0 {
				continue
			}
			t := float64(i) / float64(steps)
			angle := startAngle + t*(endAngle-startAngle)
			x2d := cx + r*math.Cos(angle)
			y2d := cy + r*math.Sin(angle)
			theta := x2d / globalR
			*targetCurve = append(*targetCurve, vector3{
				X: globalR * math.Cos(theta),
				Y: y2d,
				Z: globalR * math.Sin(theta),
			})
		}
	}

	for i := range startArcs {
		sa := startArcs[i]
		appendArcPoints(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)
		if i < len(endArcs) {
			ea := endArcs[i]
			appendArcPoints(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
		}
	}
	for i := range topStartArcs {
		tsa := topStartArcs[i]
		appendArcPoints(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)
		if i < len(topEndArcs) {
			ea := topEndArcs[i]
			appendArcPoints(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
		}
	}

	sideLength := plant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	zStart := vase.Z_Ribbon * sideLength
	if len(curvePoints) > 0 {
		y0 := curvePoints[0].Y
		for i := range curvePoints {
			curvePoints[i].Y = zStart + scaleY*(curvePoints[i].Y-y0)
		}
		for i := range topCurvePoints {
			topCurvePoints[i].Y = zStart + scaleY*(topCurvePoints[i].Y-y0)
		}
	}

	targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := getTargetAnglesSTL(curvePoints, topCurvePoints, 0.5, radialRepetitions)
	if len(targetAngles) < 4 {
		return 0, 0, 0, false
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)
	resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
	resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)
	if len(resampledBaseBottom) != len(targetAngles) || len(resampledBaseTop) != len(targetAngles) {
		return 0, 0, 0, false
	}

	baseBottomY := make([]float64, len(targetAngles))
	baseTopY := make([]float64, len(targetAngles))
	for i := range targetAngles {
		baseBottomY[i] = resampledBaseBottom[i].Y
		baseTopY[i] = resampledBaseTop[i].Y
	}

	var rotatedStepDX, rotatedStepDY float64
	if newBottomY, newTopY, _, rotDX, rotDY, constOk := ComputeVaseTrapezeConstant3DThickness(plant, globalR, targetAngles, baseBottomY); constOk {
		baseBottomY = newBottomY
		baseTopY = newTopY
		rotatedStepDX = rotDX
		rotatedStepDY = rotDY
	} else {
		var growthVectorX, growthVectorY float64
		if plant.GrowthVectorShape != nil {
			growthVectorX = plant.GrowthVectorShape.X
			growthVectorY = plant.GrowthVectorShape.Y
		}
		var vx, vy float64
		if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
			vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
			vx = vFirst.EndX - vFirst.StartX
			vy = vFirst.EndY - vFirst.StartY
			if vLen := math.Hypot(vx, vy); vLen > 0 {
				vx, vy = vx/vLen, vy/vLen
			}
		}
		verticalThickness := vase.RelativeVerticalThickness * plant.RhombusSideLength
		rotatedSeparation := vase.RelativeRotatedTorusSeparation * plant.RhombusSideLength
		rotatedStepDX = growthVectorX + verticalThickness*vx
		rotatedStepDY = (growthVectorY + verticalThickness*vy + rotatedSeparation) * scaleY
	}

	K := len(targetAngles) - 1
	thetaSector := targetAngles[K] - targetAngles[0]
	if K < 3 || thetaSector <= 0 {
		return 0, 0, 0, false
	}
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

	// Flat ring top height (as if CarvedOutTopRingsParameter == 0)
	p1H := (vase.Z_Ribbon + vase.TopPlaneHeight*scaleY) * sideLength
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	tanAlpha := math.Tan(projAngleRad)
	rotStepTheta := rotatedStepDX / globalR

	rTop0 := make([]float64, len(targetAngles))
	rBot0 := make([]float64, len(targetAngles))
	var sumRTop0, sumRBot0 float64
	for i := range targetAngles {
		rTop0[i] = globalR + (p1H-baseTopY[i])*tanAlpha
		rBot0[i] = globalR + (p1H-baseBottomY[i])*tanAlpha
		if i < K {
			sumRTop0 += rTop0[i]
			sumRBot0 += rBot0[i]
		}
	}

	rInner0 := make([]float64, len(targetAngles))
	rOuter1 := make([]float64, len(targetAngles))
	bottomIsOuter := sumRBot0 > sumRTop0
	for i, th := range targetAngles {
		if bottomIsOuter {
			rInner0[i] = math.Max(0.0, rTop0[i])
			yOuter1 := evalPeriodicY(baseBottomY, th-rotStepTheta) + rotatedStepDY
			rOuter1[i] = math.Max(0.0, globalR+(p1H-yOuter1)*tanAlpha)
		} else {
			rInner0[i] = math.Max(0.0, rBot0[i])
			yOuter1 := evalPeriodicY(baseTopY, th-rotStepTheta) + rotatedStepDY
			rOuter1[i] = math.Max(0.0, globalR+(p1H-yOuter1)*tanAlpha)
		}
	}

	var sectorInteriorArea, sectorOccupiedArea float64
	for i := range K {
		dTh := targetAngles[i+1] - targetAngles[i]
		sinDTh := math.Sin(dTh)
		rInA := rInner0[i]
		rInB := rInner0[i+1]
		rOccA := math.Min(rInA, rOuter1[i])
		rOccB := math.Min(rInB, rOuter1[i+1])

		sectorInteriorArea += 0.5 * rInA * rInB * sinDTh
		sectorOccupiedArea += 0.5 * rOccA * rOccB * sinDTh
	}

	interiorSurface = sectorInteriorArea * float64(radialRepetitions)
	occupiedSurface := sectorOccupiedArea * float64(radialRepetitions)
	if interiorSurface <= 1e-9 {
		return 0, 0, 0, false
	}
	unoccupiedSurface = math.Max(0.0, interiorSurface-occupiedSurface)
	proportion = unoccupiedSurface / interiorSurface
	return unoccupiedSurface, interiorSurface, proportion, true
}
