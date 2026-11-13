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

// ComputeVaseTrapezeContactStepDY computes the zero-separation vertical step between
// consecutive rotated VaseTrapeze rings when CarvedOutTopRingsParameter == 0.
func ComputeVaseTrapezeContactStepDY(plant *PlantAbstract, hasConstantThickness bool) float64 {
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

// computeEuclideanConvexArcParams computes the circle parameters in the 2D (r, y) meridian
// plane for a single constant-curvature convex arc (wallCurvature < 0) passing through
// (r_straight(yRefBot), yRefBot) and (r_straight(yRefTop), yRefTop) and remaining
// well-defined across the full evaluation interval [yMinEval, yMaxEval].
func computeEuclideanConvexArcParams(
	yRefBot, yRefTop, yMinEval, yMaxEval, ampSpan, projAngleDeg, wallCurvature float64,
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
	sag := eta * (ampSpan / cosA)
	if sag <= 1e-9 {
		return 0, 0, 0, 0, false
	}

	radius = (chordLen*chordLen)/(8.0*sag) + 0.5*sag
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
		return math.Abs(yMinEval-yc) <= gamma*rCand && math.Abs(yMaxEval-yc) <= gamma*rCand
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

	rCenterRel, yCenter, radius, tanAlpha, ok := computeEuclideanConvexArcParams(
		yContact, p1H, yMinEval, yMaxEval, span, vase.ProjectionAngle, vase.WallCurvature,
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
	if deltaYContact > 1e-6 && deltaYContact < ampScale {
		ampScale = deltaYContact
	}
	if localY < p2HBase {
		localY = p2HBase
	} else if localY > p1H {
		localY = p1H
	}
	rCenterRel, yCenter, radius, tanAlpha, ok := computeEuclideanConvexArcParams(
		p2HBase, p1H, p2HBase, p1H, ampScale, vase.ProjectionAngle, vase.WallCurvature,
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
	minSpan := (yTop - yBot) / float64(sWall)
	yMid := math.Max(yBot+minSpan, math.Min(yTop-minSpan, yContactGlobal))
	if s <= sSplit {
		return yBot + (float64(s)/float64(sSplit))*(yMid-yBot)
	}
	return yMid + (float64(s-sSplit)/float64(sWall-sSplit))*(yTop-yMid)
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
