package models

import "math"

// IsVaseTrapezeCurvedWallActive reports whether convex curved walls are enabled
// and non-zero (WallCurvature < 0) for a VaseTrapeze plant.
func IsVaseTrapezeCurvedWallActive(plant *PlantAbstract) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		plant.TubeVaseAbstract.HasCurvedWall &&
		plant.TubeVaseAbstract.WallCurvature < -1e-9
}

// VaseTrapezeWallSubdivisions returns the number of vertical subdivisions to use
// for VaseTrapeze ring walls (1 when straight, 24 when curved).
func VaseTrapezeWallSubdivisions(plant *PlantAbstract) int {
	if IsVaseTrapezeCurvedWallActive(plant) {
		return 24
	}
	return 1
}

// ComputeVaseTrapezeRotatedContactStepDY computes the zero-separation vertical step
// between consecutive rotated VaseTrapeze rings when CarvedOutTopRingsParameter == 0.
func ComputeVaseTrapezeRotatedContactStepDY(plant *PlantAbstract, hasConstantThickness bool) float64 {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return 0.0
	}
	vase := plant.TubeVaseAbstract
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	sideLength := plant.RhombusSideLength
	verticalThickness := vase.RelativeVerticalThickness * sideLength

	var growthVectorY float64
	if plant.GrowthVectorShape != nil {
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

	if hasConstantThickness && math.Abs(vy) > 1e-6 {
		deltaYV := (verticalThickness / vy) * scaleY
		return growthVectorY*scaleY + deltaYV
	}
	return (growthVectorY + verticalThickness*vy) * scaleY
}

// ComputeVaseTrapezeContactStepDY computes the active zero-separation vertical contact step
// anchoring yContact = p1H - deltaYContact on curved VaseTrapeze ring walls, adapting smoothly
// with RotationRatio from the unrotated rim contact step (at RotationRatio == 0) to the
// rotated contact step (at RotationRatio == 1).
func ComputeVaseTrapezeContactStepDY(plant *PlantAbstract, hasConstantThickness bool) float64 {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return 0.0
	}
	rotContactDY := ComputeVaseTrapezeRotatedContactStepDY(plant, hasConstantThickness)
	w := VaseTrapezeRotationBlend(plant)
	if w >= 1.0-1e-12 {
		return rotContactDY
	}

	vase := plant.TubeVaseAbstract
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	sideLength := plant.RhombusSideLength
	verticalThickness := vase.RelativeVerticalThickness * sideLength

	var vx, vy float64
	if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
		vx = vFirst.EndX - vFirst.StartX
		vy = vFirst.EndY - vFirst.StartY
		if vLen := math.Hypot(vx, vy); vLen > 0 {
			vx, vy = vx/vLen, vy/vLen
		}
	}
	if math.Abs(vy) <= 1e-6 {
		return rotContactDY
	}

	deltaYV := (verticalThickness / vy) * scaleY
	unrotContactDY := deltaYV
	if hasConstantThickness {
		radialRepetitions := max(vase.RadialRepetitions, 1)
		circumference := 0.0
		if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
			circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
		} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 1 {
			pGrid := plant.PerpendicularVectorGrid
			circumference = pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1].StartX - pGrid.PerpendicularVectors[0].StartX
		}
		globalR := circumference * float64(radialRepetitions) / (2.0 * math.Pi)
		tanAlpha := math.Tan(-vase.ProjectionAngle * math.Pi / 180.0)
		if globalR > 1e-3 && math.Abs(tanAlpha) > 1e-6 {
			sV := (-vx / vy) * globalR * scaleY
			kSlope := tanAlpha / globalR
			unrotContactDY = deltaYV / math.Hypot(1.0, sV*kSlope)
		}
	}
	return (1.0-w)*unrotContactDY + w*rotContactDY
}

// ComputeVaseTrapezeContactStepDYForCurves computes the active zero-separation vertical contact step
// anchoring yContact = p1H - deltaYContact using the resampled bottom and top cylinder curves so that
// at RotationRatio == 0, yContact matches the unrotated curve separation to machine precision.
func ComputeVaseTrapezeContactStepDYForCurves(
	plant *PlantAbstract,
	hasConstantThickness bool,
	baseBottomY, baseTopY []float64,
) float64 {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return 0.0
	}
	rotContactDY := ComputeVaseTrapezeRotatedContactStepDY(plant, hasConstantThickness)
	w := VaseTrapezeRotationBlend(plant)
	if w >= 1.0-1e-12 {
		return rotContactDY
	}
	M := min(len(baseBottomY), len(baseTopY))
	if M == 0 {
		return ComputeVaseTrapezeContactStepDY(plant, hasConstantThickness)
	}
	minDiff := math.MaxFloat64
	maxDiff := 0.0
	for i := range M {
		d := math.Abs(baseTopY[i] - baseBottomY[i])
		if d < minDiff {
			minDiff = d
		}
		if d > maxDiff {
			maxDiff = d
		}
	}
	unrotContactDY := maxDiff
	if hasConstantThickness && minDiff > 1e-6 && minDiff < math.MaxFloat64 {
		unrotContactDY = minDiff
	}
	return (1.0-w)*unrotContactDY + w*rotContactDY
}

// estimateMinRingInnerStraightRadius estimates a conservative, scale-homogeneous lower bound
// on the straight inner radius of a VaseTrapeze ring at height yEval across all azimuths.
func estimateMinRingInnerStraightRadius(plant *PlantAbstract, yEval, tanAlpha float64) float64 {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return 0.0
	}
	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)
	circumference := 0.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 1 {
		pGrid := plant.PerpendicularVectorGrid
		circumference = pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1].StartX - pGrid.PerpendicularVectors[0].StartX
	}
	if circumference <= 0 {
		return 0.0
	}
	globalR := circumference * float64(radialRepetitions) / (2.0 * math.Pi)
	if math.Abs(tanAlpha) <= 1e-6 {
		return globalR
	}

	sideLength := plant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	zStart := vase.Z_Ribbon * sideLength

	hasY0 := false
	y0 := 0.0
	minY2D := math.MaxFloat64
	maxY2D := -math.MaxFloat64
	updateArcBounds := func(x1, y1, x2, y2, r float64, largeArc bool) {
		if !hasY0 {
			y0 = y1
			hasY0 = true
		}
		sag := 0.0
		halfChord := 0.5 * math.Hypot(x2-x1, y2-y1)
		if !largeArc && r > halfChord && halfChord > 0 {
			sag = r - math.Sqrt(math.Max(0.0, r*r-halfChord*halfChord))
		}
		lo := math.Min(y1, y2) - sag
		hi := math.Max(y1, y2) + sag
		if lo < minY2D {
			minY2D = lo
		}
		if hi > maxY2D {
			maxY2D = hi
		}
	}

	if plant.GrowthCurve2D != nil && plant.GrowthCurve2D.StartHalfwayArcShapeGrid != nil {
		for _, sa := range plant.GrowthCurve2D.StartHalfwayArcShapeGrid.StartHalfwayArcShapes {
			updateArcBounds(sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, sa.LargeArcFlag)
		}
		if plant.GrowthCurve2D.EndHalfwayArcShapeGrid != nil {
			for _, ea := range plant.GrowthCurve2D.EndHalfwayArcShapeGrid.EndHalfwayArcShapes {
				updateArcBounds(ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, ea.LargeArcFlag)
			}
		}
	}
	if vase.TopGrowthCurve2D != nil && vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != nil {
		for _, tsa := range vase.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes {
			updateArcBounds(tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, tsa.LargeArcFlag)
		}
		if vase.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid != nil {
			for _, tea := range vase.TopGrowthCurve2D.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes {
				updateArcBounds(tea.StartX, tea.StartY, tea.EndX, tea.EndY, tea.RadiusX, tea.LargeArcFlag)
			}
		}
	}
	if !hasY0 && plant.PerpendicularVectorGrid != nil {
		for _, pv := range plant.PerpendicularVectorGrid.PerpendicularVectors {
			if !hasY0 {
				y0 = pv.StartY
				hasY0 = true
			}
			if pv.StartY < minY2D {
				minY2D = pv.StartY
			}
			if pv.StartY > maxY2D {
				maxY2D = pv.StartY
			}
		}
	}
	if !hasY0 {
		return globalR
	}

	minCurveY := zStart + scaleY*(minY2D-y0)
	maxCurveY := zStart + scaleY*(maxY2D-y0)
	if vase.HasConstantThickness && plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
		vFirst := plant.PerpendicularVectorGrid.PerpendicularVectors[0]
		vx := vFirst.EndX - vFirst.StartX
		vy := vFirst.EndY - vFirst.StartY
		if vLen := math.Hypot(vx, vy); vLen > 0 {
			vy /= vLen
			if math.Abs(vy) > 1e-6 {
				verticalThickness := vase.RelativeVerticalThickness * sideLength
				deltaYV := (verticalThickness / vy) * scaleY
				extraTopY := deltaYV - verticalThickness*vy*scaleY
				if extraTopY > 0 {
					maxCurveY += extraTopY
				}
			}
		}
	}

	rAtMin := globalR + (yEval-minCurveY)*tanAlpha
	rAtMax := globalR + (yEval-maxCurveY)*tanAlpha
	rMin := math.Min(rAtMin, rAtMax)
	if rMin < 0.15*globalR {
		return 0.15 * globalR
	}
	return rMin
}

// computeEuclideanConvexArcParams computes the circle parameters in the 2D (r, y) meridian
// plane for a single constant-curvature convex arc (wallCurvature < 0) passing through
// (r_straight(yRefBot), yRefBot) and (r_straight(yRefTop), yRefTop) and remaining
// well-defined across the full evaluation interval [yMinEval, yMaxEval].
func computeEuclideanConvexArcParams(
	yRefBot, yRefTop, yMinEval, yMaxEval, ampSpan, maxInwardDR, projAngleDeg, wallCurvature float64,
) (rCenterRel, yCenter, radius, tanAlpha float64, ok bool) {
	if wallCurvature >= -1e-9 {
		return 0, 0, 0, 0, false
	}
	if wallCurvature < -1.0 {
		wallCurvature = -1.0
	}
	chordDY := yRefTop - yRefBot
	if chordDY <= 1e-6 {
		return 0, 0, 0, 0, false
	}
	if ampSpan <= 1e-6 {
		ampSpan = chordDY
	}

	alpha := -projAngleDeg * math.Pi / 180.0
	if alpha > 1.2 {
		alpha = 1.2
	} else if alpha < -1.2 {
		alpha = -1.2
	}
	cosA := math.Cos(alpha)
	sinA := math.Sin(alpha)
	tanAlpha = math.Tan(alpha)

	chordLen := chordDY / cosA
	halfChord := 0.5 * chordLen
	yMid := 0.5 * (yRefBot + yRefTop)

	eta := -0.25 * wallCurvature
	refChordLen := math.Max(chordLen, ampSpan/cosA)
	sag := eta * (ampSpan / cosA)
	if sag <= 1e-9 {
		return 0, 0, 0, 0, false
	}

	radius = (refChordLen*refChordLen)/(8.0*sag) + 0.5*sag
	gamma := math.Max(0.96, math.Min(0.995, math.Abs(sinA)+0.05))

	isValid := func(rCand float64) bool {
		if rCand <= halfChord*1.001 {
			return false
		}
		d := math.Sqrt(rCand*rCand - halfChord*halfChord)
		if d*cosA <= 0.5*math.Abs(chordDY*tanAlpha)+1e-4*chordLen {
			return false
		}
		yc := yMid + d*sinA
		if math.Abs(yMinEval-yc) > gamma*rCand || math.Abs(yMaxEval-yc) > gamma*rCand {
			return false
		}
		if maxInwardDR > 0 && yMinEval < yRefBot-1e-9 {
			rCenter := 0.5*chordDY*tanAlpha - d*cosA
			dyMin := yMinEval - yc
			radMin := math.Sqrt(math.Max(0.0, rCand*rCand-dyMin*dyMin))
			drMin := rCenter + radMin - (yMinEval-yRefBot)*tanAlpha
			if -drMin > maxInwardDR {
				return false
			}
		}
		return true
	}

	if !isValid(radius) {
		lo := halfChord * 1.001
		hi := math.Max(radius, chordLen)
		for iter := 0; iter < 60 && !isValid(hi); iter++ {
			hi *= 2.0
		}
		if !isValid(hi) {
			return 0, 0, 0, 0, false
		}
		for iter := 0; iter < 40; iter++ {
			mid := 0.5 * (lo + hi)
			if isValid(mid) {
				hi = mid
			} else {
				lo = mid
			}
		}
		radius = hi
	}

	d := math.Sqrt(radius*radius - halfChord*halfChord)
	rCenterRel = 0.5*chordDY*tanAlpha - d*cosA
	yCenter = yMid + d*sinA
	return rCenterRel, yCenter, radius, tanAlpha, true
}

func evalEuclideanConvexArc(
	localY, yRefBot, yRefTop, rCenterRel, yCenter, radius, tanAlpha float64,
) (deltaR, dDeltaRDy float64) {
	dyC := localY - yCenter
	rad := math.Sqrt(math.Max(1e-12, radius*radius-dyC*dyC))
	dDeltaRDy = -dyC/rad - tanAlpha
	if math.Abs(localY-yRefBot) <= 1e-9 || math.Abs(localY-yRefTop) <= 1e-9 {
		return 0.0, dDeltaRDy
	}
	deltaR = rCenterRel + rad - (localY-yRefBot)*tanAlpha
	return deltaR, dDeltaRDy
}

// EvalVaseTrapezeRingWallDeltaR returns the radial offset deltaR(localY) for a stacked
// or standalone VaseTrapeze ring (localY in [p2H, p1H]).
//
// It evaluates a single constant-curvature circular arc in the (r, y) cross-section plane:
//   - deltaR(p1H) == 0 (conserved top contour when uncarved)
//   - deltaR(p1H - deltaYContact) == 0 (conserved contour at contact line with rotated ring below)
//   - Constant Euclidean curvature 1/R everywhere above AND below the contact line down to p2H
//     (bulging convexly > 0 above y_contact and curving smoothly inward < 0 below y_contact).
func EvalVaseTrapezeRingWallDeltaR(plant *PlantAbstract, localY, p1H, p2H, deltaYContact float64) float64 {
	if !IsVaseTrapezeCurvedWallActive(plant) {
		return 0.0
	}
	vase := plant.TubeVaseAbstract
	span := deltaYContact
	if span <= 1e-6 {
		span = math.Abs(p1H - p2H)
	}
	if span <= 1e-6 {
		return 0.0
	}
	yContact := p1H - span
	yMinEval := math.Min(p2H, yContact)
	yMaxEval := p1H
	ampSpan := math.Max(span, ComputeVaseTrapezeRotatedContactStepDY(plant, vase.HasConstantThickness))
	tanA := math.Tan(-vase.ProjectionAngle * math.Pi / 180.0)
	maxInwardDR := 0.72 * estimateMinRingInnerStraightRadius(plant, yMinEval, tanA)

	rCenterRel, yCenter, radius, tanAlpha, ok := computeEuclideanConvexArcParams(
		yContact, p1H, yMinEval, yMaxEval, ampSpan, maxInwardDR, vase.ProjectionAngle, vase.WallCurvature,
	)
	if !ok {
		return 0.0
	}
	dr, _ := evalEuclideanConvexArc(localY, yContact, p1H, rCenterRel, yCenter, radius, tanAlpha)
	return dr
}

// EvalVaseTrapezeBaseRingOuterDeltaRAndSlope evaluates the radial offset deltaR(localY)
// and vertical derivative d(deltaR)/dy for the outer wall of Ring 0 (attached to Base_Plate,
// spanning localY in [p2HBase, p1H]).
//
// It spans one constant-curvature convex circular arch in (r, y) from p2HBase to p1H,
// returning deltaR == 0 at both endpoints and providing the exact slope at p2HBase+
// to feed tangent-continuously (C1) into the bulbous base.
func EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant *PlantAbstract, localY, p1H, p2HBase, deltaYContact float64) (deltaR float64, dDeltaRDy float64) {
	if !IsVaseTrapezeCurvedWallActive(plant) {
		return 0.0, 0.0
	}
	vase := plant.TubeVaseAbstract
	spanBase := p1H - p2HBase
	if spanBase <= 1e-6 {
		return 0.0, 0.0
	}
	ampScale := spanBase
	refSpan := math.Max(deltaYContact, ComputeVaseTrapezeRotatedContactStepDY(plant, vase.HasConstantThickness))
	if refSpan > 1e-6 && refSpan < ampScale {
		ampScale = refSpan
	}
	if localY < p2HBase {
		localY = p2HBase
	} else if localY > p1H {
		localY = p1H
	}
	rCenterRel, yCenter, radius, tanAlpha, ok := computeEuclideanConvexArcParams(
		p2HBase, p1H, p2HBase, p1H, ampScale, 0.0, vase.ProjectionAngle, vase.WallCurvature,
	)
	if !ok {
		return 0.0, 0.0
	}
	return evalEuclideanConvexArc(localY, p2HBase, p1H, rCenterRel, yCenter, radius, tanAlpha)
}

// EvalVaseTrapezeBaseRingInnerDeltaR evaluates the radial offset deltaR(localY) for the
// inner wall of Ring 0 (attached to Base_Plate, spanning localY in [p2HBase, p1H]).
//
// Since WallCurvature <= 0, Ring 0's arch bulges outward (>= 0) on [p2HBase, p1H] while
// Ring 1's lower wall below its contact line curves inward (<= 0), guaranteeing full
// non-collision clearance with a single clean constant-curvature inner arch.
func EvalVaseTrapezeBaseRingInnerDeltaR(
	plant *PlantAbstract,
	localY, p1H, p2HBase, p2H, deltaYContact float64,
	bottomIsOuter bool,
) float64 {
	_ = p2H
	_ = bottomIsOuter
	dr, _ := EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant, localY, p1H, p2HBase, deltaYContact)
	return dr
}

// RingWallSampleY computes the vertical sample coordinate y for row s in [0, sWall]
// between yBot (at s == 0) and yTop (at s == sWall), placing row sSplit = sWall/3
// directly on yContactGlobal whenever yContactGlobal lies inside [yBot+minSpan, yTop-minSpan]
// so the contact line is represented without polygon chordal error while keeping rows strictly monotonic.
func RingWallSampleY(s, sWall int, yBot, yTop, yContactGlobal float64) float64 {
	if s <= 0 {
		return yBot
	}
	if s >= sWall {
		return yTop
	}
	if sWall < 3 || yTop <= yBot+1e-9 {
		return yBot + (float64(s)/float64(sWall))*(yTop-yBot)
	}
	sSplit := sWall / 3
	span := yTop - yBot
	minBot := yBot + 0.25*span*(float64(sSplit)/float64(sWall))
	maxTop := yTop - 0.25*span*(float64(sWall-sSplit)/float64(sWall))
	if yContactGlobal <= minBot || yContactGlobal >= maxTop {
		return yBot + (float64(s)/float64(sWall))*span
	}
	if s <= sSplit {
		return yBot + (float64(s)/float64(sSplit))*(yContactGlobal-yBot)
	}
	return yContactGlobal + (float64(s-sSplit)/float64(sWall-sSplit))*(yTop-yContactGlobal)
}

// OffsetRadialPoint shifts a 3D point (X, Y, Z) horizontally in cylindrical radius by deltaR.
func OffsetRadialPoint(x, y, z, deltaR float64) (nx, ny, nz float64) {
	if math.Abs(deltaR) <= 1e-12 {
		return x, y, z
	}
	r := math.Hypot(x, z)
	theta := math.Atan2(z, x)
	rNew := math.Max(0.0, r+deltaR)
	return rNew * math.Cos(theta), y, rNew * math.Sin(theta)
}

type meridianPt struct {
	r float64
	y float64
}

func polyMaxYAtR(poly []meridianPt, r float64) (float64, bool) {
	bestY := -math.MaxFloat64
	found := false
	n := len(poly)
	for e := range n {
		a := poly[e]
		b := poly[(e+1)%n]
		dr := b.r - a.r
		if math.Abs(dr) <= 1e-11 {
			if math.Abs(r-a.r) <= 1e-9 {
				if cand := math.Max(a.y, b.y); cand > bestY {
					bestY = cand
				}
				found = true
			}
			continue
		}
		rLo, rHi := a.r, b.r
		if rLo > rHi {
			rLo, rHi = rHi, rLo
		}
		if r >= rLo-1e-9 && r <= rHi+1e-9 {
			t := math.Max(0.0, math.Min(1.0, (r-a.r)/dr))
			if cand := (1.0-t)*a.y + t*b.y; cand > bestY {
				bestY = cand
			}
			found = true
		}
	}
	return bestY, found
}

func polyMinYAtR(poly []meridianPt, r float64) (float64, bool) {
	bestY := math.MaxFloat64
	found := false
	n := len(poly)
	for e := range n {
		a := poly[e]
		b := poly[(e+1)%n]
		dr := b.r - a.r
		if math.Abs(dr) <= 1e-11 {
			if math.Abs(r-a.r) <= 1e-9 {
				if cand := math.Min(a.y, b.y); cand < bestY {
					bestY = cand
				}
				found = true
			}
			continue
		}
		rLo, rHi := a.r, b.r
		if rLo > rHi {
			rLo, rHi = rHi, rLo
		}
		if r >= rLo-1e-9 && r <= rHi+1e-9 {
			t := math.Max(0.0, math.Min(1.0, (r-a.r)/dr))
			if cand := (1.0-t)*a.y + t*b.y; cand < bestY {
				bestY = cand
			}
			found = true
		}
	}
	return bestY, found
}

func verticalDropInContactStep(polyL, polyU []meridianPt) (float64, bool) {
	if len(polyL) < 2 || len(polyU) < 2 {
		return 0.0, false
	}
	rMinL, rMaxL := math.MaxFloat64, -math.MaxFloat64
	for _, p := range polyL {
		if p.r < rMinL {
			rMinL = p.r
		}
		if p.r > rMaxL {
			rMaxL = p.r
		}
	}
	rMinU, rMaxU := math.MaxFloat64, -math.MaxFloat64
	for _, p := range polyU {
		if p.r < rMinU {
			rMinU = p.r
		}
		if p.r > rMaxU {
			rMaxU = p.r
		}
	}
	if rMaxL < rMinU-1e-9 || rMaxU < rMinL-1e-9 {
		return 0.0, false
	}

	maxStep := -math.MaxFloat64
	found := false
	for _, v := range polyL {
		if v.r < rMinU-1e-9 || v.r > rMaxU+1e-9 {
			continue
		}
		if minYU, ok := polyMinYAtR(polyU, v.r); ok {
			if step := v.y - minYU; step > maxStep {
				maxStep = step
			}
			found = true
		}
	}
	for _, w := range polyU {
		if w.r < rMinL-1e-9 || w.r > rMaxL+1e-9 {
			continue
		}
		if maxYL, ok := polyMaxYAtR(polyL, w.r); ok {
			if step := maxYL - w.y; step > maxStep {
				maxStep = step
			}
			found = true
		}
	}
	return maxStep, found
}

// ComputeVaseTrapezeUnrotatedStepDYs computes the exact vertical rigid-body drop-in
// contact steps for unrotated stacking (DeltaTheta == 0) of VaseTrapeze rings:
//   - stepDYBase: vertical offset for Ring 1 resting unrotated on Base Plate + Ring 0
//   - stepDYRing: vertical offset for Ring k (k >= 2) resting unrotated on Ring k-1
//
// It accounts for straight or convex curved walls (including rim undercuts), constant
// or variable thickness, HasRingBottom (plancher), Base Plate bulbous walls, and
// CarvedOutTopRingsParameter.
func ComputeVaseTrapezeUnrotatedStepDYs(
	plant *PlantAbstract,
	globalR float64,
	baseBottomY, baseTopY []float64,
	hasConstantThickness bool,
) (stepDYBase, stepDYRing float64) {
	if plant == nil || plant.TubeVaseAbstract == nil {
		return 0.0, 0.0
	}
	vase := plant.TubeVaseAbstract
	sideLength := plant.RhombusSideLength
	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}
	sepDY := vase.RelativeRotatedTorusSeparation * sideLength * scaleY

	M := min(len(baseBottomY), len(baseTopY))
	if M == 0 || globalR <= 0 {
		return sepDY, sepDY
	}

	maxCurveDiff := 0.0
	for i := range M {
		if d := math.Abs(baseTopY[i] - baseBottomY[i]); d > maxCurveDiff {
			maxCurveDiff = d
		}
	}

	zStart := vase.Z_Ribbon * sideLength
	p1H := zStart + vase.TopPlaneHeight*sideLength*scaleY
	p2H := zStart + vase.BottomPlaneHeight*sideLength*scaleY
	p2HBase := zStart + vase.BaseRingBottomPlaneHeight*sideLength*scaleY
	ringHeight := p1H - p2H
	if ringHeight <= 1e-6 {
		return maxCurveDiff + sepDY, maxCurveDiff + sepDY
	}

	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	tanAlpha := math.Tan(projAngleRad)
	carvedParam := math.Max(0.0, math.Min(1.0, vase.CarvedOutTopRingsParameter))

	basePlateHeight := vase.BasePlateHeight * sideLength
	hasRingBottom := vase.HasRingBottom && basePlateHeight > 0 && ringHeight > 0
	hRingBottom := 0.0
	if hasRingBottom {
		hRingBottom = math.Min(basePlateHeight, ringHeight)
	}
	yFloorRing := p2H + hRingBottom
	footHRing := ComputeVaseTrapezeFootRimHeight(plant, hRingBottom, hasRingBottom)
	hasFootRim := footHRing > 1e-4
	yOuterWallBotRing := p2H + footHRing
	footInsetRing := 0.0
	if hasFootRim {
		footInsetRing = math.Min(DefaultVaseTrapezeFootRimInsetRatio*sideLength, 0.75*footHRing)
	}

	deltaYContact := ComputeVaseTrapezeContactStepDYForCurves(plant, hasConstantThickness, baseBottomY, baseTopY)
	yContact := p1H - deltaYContact
	sWall := VaseTrapezeWallSubdivisions(plant)
	isCurvedWall := IsVaseTrapezeCurvedWallActive(plant)
	lipHeight, lipWidth, _, hasLip := VaseTrapezeLipDimensionsForMode(plant, hasConstantThickness)

	var sumRBottom, sumRTop float64
	for i := range M {
		sumRBottom += math.Max(0.0, globalR+(p2H-baseBottomY[i])*tanAlpha)
		sumRTop += math.Max(0.0, globalR+(p2H-baseTopY[i])*tanAlpha)
	}
	bottomIsOuter := sumRBottom > sumRTop

	yCurveOut := baseTopY
	yCurveIn := baseBottomY
	if bottomIsOuter {
		yCurveOut = baseBottomY
		yCurveIn = baseTopY
	}

	rStOut := func(i int, y float64) float64 {
		return math.Max(0.0, globalR+(y-yCurveOut[i])*tanAlpha)
	}
	rStIn := func(i int, y float64) float64 {
		return math.Max(0.0, globalR+(y-yCurveIn[i])*tanAlpha)
	}

	rMinOutP1, rMaxOutP1 := math.MaxFloat64, -math.MaxFloat64
	rMinInP1, rMaxInP1 := math.MaxFloat64, -math.MaxFloat64
	for i := range M {
		ro := rStOut(i, p1H)
		if ro < rMinOutP1 {
			rMinOutP1 = ro
		}
		if ro > rMaxOutP1 {
			rMaxOutP1 = ro
		}
		ri := rStIn(i, p1H)
		if ri < rMinInP1 {
			rMinInP1 = ri
		}
		if ri > rMaxInP1 {
			rMaxInP1 = ri
		}
	}
	rRangeOut := rMaxOutP1 - rMinOutP1
	rRangeIn := rMaxInP1 - rMinInP1

	ringDR := func(localY float64) float64 {
		if !isCurvedWall {
			return 0.0
		}
		return EvalVaseTrapezeRingWallDeltaR(plant, localY, p1H, p2H, deltaYContact)
	}

	const floorY = 0.0
	hasBasePlate := p2HBase > floorY && p1H > p2HBase+1e-6
	hPlancherBase := 0.0
	plancherYBase := floorY
	verticalScaleBase := 0.0
	startMag := vase.BulbousStartTangentMagnitude * sideLength
	endAngleDeg := 30.0
	if vase.BulbousEndAngle != 0.0 {
		endAngleDeg = vase.BulbousEndAngle
	}
	endMag := vase.BulbousEndTangentMagnitude * sideLength
	sBulb := 24
	if startMag == 0 && endMag == 0 {
		sBulb = 1
	}
	beta := endAngleDeg * math.Pi / 180.0
	v1R := -endMag * math.Sin(beta)
	v1Y := -endMag * math.Cos(beta)
	var dBaseOutDyAtBottom float64
	if hasBasePlate {
		hPlancherBase = math.Max(0.0, math.Min(p2HBase-floorY, basePlateHeight))
		plancherYBase = floorY + hPlancherBase
		verticalScaleBase = (p2HBase - plancherYBase) / (p2HBase - floorY)
		_, dBaseOutDyAtBottom = EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant, p2HBase, p1H, p2HBase, deltaYContact)
	}

	baseOutDR := func(localY float64) float64 {
		if !isCurvedWall {
			return 0.0
		}
		dr, _ := EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant, localY, p1H, p2HBase, deltaYContact)
		return dr
	}
	baseInDR := func(localY float64) float64 {
		if !isCurvedWall {
			return 0.0
		}
		return EvalVaseTrapezeBaseRingInnerDeltaR(plant, localY, p1H, p2HBase, p2H, deltaYContact, bottomIsOuter)
	}

	polyRing := make([]meridianPt, 0, 2*(sWall+6))
	polyBase := make([]meridianPt, 0, 2*(sWall+sBulb+6))

	K := M
	if M > 3 && math.Abs(baseBottomY[M-1]-baseBottomY[0]) < 1e-6 && math.Abs(baseTopY[M-1]-baseTopY[0]) < 1e-6 {
		K = M - 1
	}
	var ringBotLipOff, ringTopLipOff, baseBotLipOff, baseTopLipOff []float64
	if hasLip && K > 0 {
		radialReps := max(vase.RadialRepetitions, 1)
		dTh2 := 2.0 * (2.0 * math.Pi / float64(radialReps)) / float64(K)
		computeSectorLipOff := func(radii []float64) []float64 {
			offs := make([]float64, K)
			if K < 3 || dTh2 <= 1e-9 {
				for idx := range K {
					offs[idx] = lipWidth
				}
				return offs
			}
			for idx := range K {
				prevIdx := (idx - 1 + K) % K
				nextIdx := (idx + 1) % K
				factor := 1.0
				if radii[idx] > 1e-6 {
					slopeNorm := (radii[nextIdx] - radii[prevIdx]) / (radii[idx] * dTh2)
					factor = math.Min(3.0, math.Hypot(1.0, slopeNorm))
				}
				offs[idx] = lipWidth * factor
			}
			return offs
		}

		rRingBotLip := make([]float64, K)
		rRingTop := make([]float64, K)
		rBaseBotLip := make([]float64, K)
		rBaseTop := make([]float64, K)
		for idx := range K {
			cOut := 0.0
			if rRangeOut > 1e-6 {
				cOut = carvedParam * (rMaxOutP1 - rStOut(idx, p1H)) / rRangeOut
			}
			ytRing := (1.0-cOut)*p1H + cOut*yFloorRing
			ylRing := VaseTrapezeLipBottomY(yOuterWallBotRing, ytRing, lipHeight)
			rRingBotLip[idx] = math.Max(0.0, rStOut(idx, ylRing)+ringDR(ylRing))
			rRingTop[idx] = math.Max(0.0, rStOut(idx, ytRing)+ringDR(ytRing))
			if hasBasePlate {
				ytBase := (1.0-cOut)*p1H + cOut*p2HBase
				ylBase := VaseTrapezeLipBottomY(p2HBase, ytBase, lipHeight)
				rBaseBotLip[idx] = math.Max(0.0, rStOut(idx, ylBase)+baseOutDR(ylBase))
				rBaseTop[idx] = math.Max(0.0, rStOut(idx, ytBase)+baseOutDR(ytBase))
			}
		}
		ringBotLipOff = computeSectorLipOff(rRingBotLip)
		ringTopLipOff = computeSectorLipOff(rRingTop)
		if hasBasePlate {
			baseBotLipOff = computeSectorLipOff(rBaseBotLip)
			baseTopLipOff = computeSectorLipOff(rBaseTop)
		}
	}

	maxStepRing := 0.0
	maxStepBase := 0.0

	for i := range M {
		iMod := i % K
		coeffOut := 0.0
		if rRangeOut > 1e-6 {
			coeffOut = carvedParam * (rMaxOutP1 - rStOut(i, p1H)) / rRangeOut
		}
		coeffIn := 0.0
		if rRangeIn > 1e-6 {
			coeffIn = carvedParam * (rMaxInP1 - rStIn(i, p1H)) / rRangeIn
		}

		yTopOutRing := (1.0-coeffOut)*p1H + coeffOut*yFloorRing
		yTopInRing := (1.0-coeffIn)*p1H + coeffIn*yFloorRing

		// 1. Build Standard Ring meridian polygon polyRing
		polyRing = polyRing[:0]
		yWallTopRing := yTopOutRing
		if hasLip {
			yWallTopRing = VaseTrapezeLipBottomY(yOuterWallBotRing, yTopOutRing, lipHeight)
		}
		if hasFootRim {
			rFootOut := math.Max(0.0, rStOut(i, yOuterWallBotRing)+ringDR(yOuterWallBotRing)-footInsetRing)
			polyRing = append(polyRing, meridianPt{r: rFootOut, y: p2H})
		}
		for s := 0; s <= sWall; s++ {
			var y float64
			if s == 0 {
				y = yOuterWallBotRing
			} else if s == sWall {
				y = yWallTopRing
			} else {
				y = RingWallSampleY(s, sWall, yOuterWallBotRing, yWallTopRing, yContact)
			}
			r := math.Max(0.0, rStOut(i, y)+ringDR(y))
			polyRing = append(polyRing, meridianPt{r: r, y: y})
		}
		if hasLip {
			for s := 0; s <= sWall; s++ {
				var y, off float64
				if s == 0 {
					y = yWallTopRing
					off = ringBotLipOff[iMod]
				} else if s == sWall {
					y = yTopOutRing
					off = ringTopLipOff[iMod]
				} else {
					t := float64(s) / float64(sWall)
					y = (1.0-t)*yWallTopRing + t*yTopOutRing
					off = (1.0-t)*ringBotLipOff[iMod] + t*ringTopLipOff[iMod]
				}
				r := math.Max(0.0, rStOut(i, y)+ringDR(y)) + off
				polyRing = append(polyRing, meridianPt{r: r, y: y})
			}
		}
		if hasRingBottom {
			for s := sWall; s >= 0; s-- {
				var y float64
				if s == sWall {
					y = yTopInRing
				} else if s == 0 {
					y = yFloorRing
				} else {
					y = RingWallSampleY(s, sWall, yFloorRing, yTopInRing, yContact)
				}
				r := math.Max(0.0, rStIn(i, y)+ringDR(y))
				polyRing = append(polyRing, meridianPt{r: r, y: y})
			}
			if hasFootRim {
				rRecOut := math.Max(0.0, rStOut(i, yOuterWallBotRing)+ringDR(yOuterWallBotRing)-footInsetRing-vase.FootRimWidth*sideLength)
				polyRing = append(polyRing,
					meridianPt{r: 0.0, y: yFloorRing},
					meridianPt{r: 0.0, y: yOuterWallBotRing},
					meridianPt{r: rRecOut, y: yOuterWallBotRing},
					meridianPt{r: rRecOut, y: p2H},
				)
			} else {
				polyRing = append(polyRing,
					meridianPt{r: 0.0, y: yFloorRing},
					meridianPt{r: 0.0, y: p2H},
				)
			}
		} else {
			for s := sWall; s >= 0; s-- {
				var y float64
				if s == sWall {
					y = yTopInRing
				} else if s == 0 {
					y = p2H
				} else {
					y = RingWallSampleY(s, sWall, p2H, yTopInRing, yContact)
				}
				r := math.Max(0.0, rStIn(i, y)+ringDR(y))
				polyRing = append(polyRing, meridianPt{r: r, y: y})
			}
		}

		if stepR, ok := verticalDropInContactStep(polyRing, polyRing); ok && stepR > maxStepRing {
			maxStepRing = stepR
		}

		if !hasBasePlate {
			if maxStepRing > maxStepBase {
				maxStepBase = maxStepRing
			}
			continue
		}

		// 2. Build Base Plate + Ring 0 meridian polygon polyBase
		yTopOutBase := (1.0-coeffOut)*p1H + coeffOut*p2HBase
		yTopInBase := (1.0-coeffIn)*p1H + coeffIn*p2HBase

		polyBase = polyBase[:0]
		yWallTopBase := yTopOutBase
		if hasLip {
			yWallTopBase = VaseTrapezeLipBottomY(p2HBase, yTopOutBase, lipHeight)
			for s := 0; s <= sWall; s++ {
				var y, off float64
				if s == 0 {
					y = yTopOutBase
					off = baseTopLipOff[iMod]
				} else if s == sWall {
					y = yWallTopBase
					off = baseBotLipOff[iMod]
				} else {
					t := float64(s) / float64(sWall)
					y = (1.0-t)*yTopOutBase + t*yWallTopBase
					off = (1.0-t)*baseTopLipOff[iMod] + t*baseBotLipOff[iMod]
				}
				r := math.Max(0.0, rStOut(i, y)+baseOutDR(y)) + off
				polyBase = append(polyBase, meridianPt{r: r, y: y})
			}
		}
		for rw := 0; rw <= sWall; rw++ {
			var oy float64
			if rw == 0 {
				oy = yWallTopBase
			} else if rw == sWall {
				oy = p2HBase
			} else {
				tOut := float64(rw) / float64(sWall)
				oy = (1.0-tOut)*yWallTopBase + tOut*p2HBase
			}
			or := math.Max(0.0, rStOut(i, oy)+baseOutDR(oy))
			polyBase = append(polyBase, meridianPt{r: or, y: oy})
		}

		r0 := rStOut(i, p2HBase)
		rTop := rStOut(i, p1H)
		rIn0 := rStIn(i, p2HBase)
		radialScale := 1.0
		if r0 > 1e-6 {
			radialScale = rIn0 / r0
		}

		deltaR := r0 - rTop
		deltaY := p2HBase - p1H
		if math.Abs(deltaY) > 1e-6 {
			deltaR += dBaseOutDyAtBottom * deltaY
		}
		lRing := math.Hypot(deltaR, deltaY)
		t0R, t0Y := 0.0, -1.0
		if lRing > 1e-6 {
			t0R = deltaR / lRing
			t0Y = deltaY / lRing
		}
		v0R := startMag * t0R
		v0Y := startMag * t0Y

		p0R, p0Y := r0, p2HBase
		p1R, p1Y := r0+v0R/3.0, p2HBase+v0Y/3.0
		p2R, p2Y := r0-v1R/3.0, floorY-v1Y/3.0
		p3R, p3Y := r0, floorY

		for s := 1; s <= sBulb; s++ {
			var rs, ys float64
			if s == sBulb {
				rs = r0
				ys = floorY
			} else {
				t := float64(s) / float64(sBulb)
				omt := 1.0 - t
				b0 := omt * omt * omt
				b1 := 3.0 * omt * omt * t
				b2 := 3.0 * omt * t * t
				b3 := t * t * t
				rs = math.Max(0.0, b0*p0R+b1*p1R+b2*p2R+b3*p3R)
				ys = b0*p0Y + b1*p1Y + b2*p2Y + b3*p3Y
			}
			polyBase = append(polyBase, meridianPt{r: rs, y: ys})
		}

		polyBase = append(polyBase,
			meridianPt{r: 0.0, y: floorY},
			meridianPt{r: 0.0, y: plancherYBase},
			meridianPt{r: rIn0, y: plancherYBase},
		)

		for s := sBulb - 1; s >= 1; s-- {
			t := float64(s) / float64(sBulb)
			omt := 1.0 - t
			b0 := omt * omt * omt
			b1 := 3.0 * omt * omt * t
			b2 := 3.0 * omt * t * t
			b3 := t * t * t
			rs := math.Max(0.0, b0*p0R+b1*p1R+b2*p2R+b3*p3R)
			ys := b0*p0Y + b1*p1Y + b2*p2Y + b3*p3Y
			polyBase = append(polyBase, meridianPt{
				r: rs * radialScale,
				y: plancherYBase + (ys-floorY)*verticalScaleBase,
			})
		}

		for rw := sWall; rw >= 0; rw-- {
			var iy float64
			if rw == sWall {
				iy = p2HBase
			} else if rw == 0 {
				iy = yTopInBase
			} else {
				tIn := float64(rw) / float64(sWall)
				iy = (1.0-tIn)*yTopInBase + tIn*p2HBase
			}
			ir := math.Max(0.0, rStIn(i, iy)+baseInDR(iy))
			polyBase = append(polyBase, meridianPt{r: ir, y: iy})
		}

		if stepB, ok := verticalDropInContactStep(polyBase, polyRing); ok && stepB > maxStepBase {
			maxStepBase = stepB
		}
	}

	return math.Max(0.0, maxStepBase) + sepDY, math.Max(0.0, maxStepRing) + sepDY
}
