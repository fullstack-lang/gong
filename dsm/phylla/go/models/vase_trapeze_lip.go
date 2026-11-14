package models

import "math"

// IsVaseTrapezeLipActive reports whether the outdoor-pot top outer lip (bordure)
// is enabled and has strictly positive height and width for a VaseTrapeze plant.
func IsVaseTrapezeLipActive(plant *PlantAbstract) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		plant.TubeVaseAbstract.HasLip &&
		plant.TubeVaseAbstract.LipHeight > 1e-9 &&
		plant.TubeVaseAbstract.LipWidth > 1e-9 &&
		plant.RhombusSideLength > 1e-9
}

// VaseTrapezeLipDimensions returns the scaled lip height and outward lip width
// (both proportional to RhombusSideLength) when IsVaseTrapezeLipActive is true.
func VaseTrapezeLipDimensions(plant *PlantAbstract) (lipHeight, lipWidth float64, ok bool) {
	if !IsVaseTrapezeLipActive(plant) {
		return 0, 0, false
	}
	side := plant.RhombusSideLength
	return plant.TubeVaseAbstract.LipHeight * side, plant.TubeVaseAbstract.LipWidth * side, true
}

// VaseTrapezeLipBottomY computes the bottom Y coordinate of the lip band at a single
// azimuth sample between outer wall bottom yBot and outer top rim yTop, preserving
// at least 2% of the outer wall span below the lip to avoid degenerate zero-height faces.
func VaseTrapezeLipBottomY(yBot, yTop, lipHeight float64) float64 {
	avail := yTop - yBot
	if avail <= 1e-9 {
		return yBot
	}
	minWallY := yBot + 0.02*avail
	yLipBot := yTop - lipHeight
	if yLipBot < minWallY {
		return minWallY
	}
	if yLipBot > yTop {
		return yTop
	}
	return yLipBot
}

// ComputeRadialLipOffsets computes the outward radial shift deltaR[i] along each point's
// cylindrical ray theta_i so that the horizontal perpendicular width of the lip normal
// to the wavy contour (xs, zs) equals lipWidth everywhere around the circumference.
func ComputeRadialLipOffsets(xs, zs []float64, lipWidth float64) []float64 {
	n := min(len(xs), len(zs))
	out := make([]float64, n)
	if n == 0 || lipWidth <= 0 {
		return out
	}
	if n < 3 {
		for i := range n {
			out[i] = lipWidth
		}
		return out
	}
	radii := make([]float64, n)
	thetas := make([]float64, n)
	for i := range n {
		radii[i] = math.Hypot(xs[i], zs[i])
		thetas[i] = math.Atan2(zs[i], xs[i])
	}
	for i := range n {
		prevI := (i - 1 + n) % n
		nextI := (i + 1) % n
		dTh := thetas[nextI] - thetas[prevI]
		for dTh <= -math.Pi {
			dTh += 2 * math.Pi
		}
		for dTh > math.Pi {
			dTh -= 2 * math.Pi
		}
		rCur := radii[i]
		factor := 1.0
		if rCur > 1e-6 && math.Abs(dTh) > 1e-6 {
			slopeNorm := (radii[nextI] - radii[prevI]) / (rCur * dTh)
			factor = math.Min(3.0, math.Hypot(1.0, slopeNorm))
		}
		out[i] = lipWidth * factor
	}
	return out
}

func computeVaseTrapezeLipCurvesVec3(
	cOuterBot, cOuterTop, cOuterStraightBot, cOuterUncarvedP1 []vector3,
	yBotRef, spanH, dy, lipHeight, lipWidth float64,
	isCurvedWall bool,
	outerDeltaRAtLocalY func(float64) float64,
) (cWallLipBot, cLipBot, cLipTop []vector3) {
	M := min(len(cOuterBot), len(cOuterTop), len(cOuterStraightBot), len(cOuterUncarvedP1))
	if M == 0 {
		return nil, nil, nil
	}
	cWallLipBot = make([]vector3, M)
	wallBotXs := make([]float64, M)
	wallBotZs := make([]float64, M)
	topXs := make([]float64, M)
	topZs := make([]float64, M)

	for i := range M {
		yLipBot := VaseTrapezeLipBottomY(cOuterBot[i].Y, cOuterTop[i].Y, lipHeight)
		u := 0.0
		if math.Abs(spanH) > 1e-9 {
			u = (yLipBot - yBotRef) / spanH
		}
		sx := (1.0-u)*cOuterStraightBot[i].X + u*cOuterUncarvedP1[i].X
		sz := (1.0-u)*cOuterStraightBot[i].Z + u*cOuterUncarvedP1[i].Z
		wx, wy, wz := sx, yLipBot, sz
		if isCurvedWall && outerDeltaRAtLocalY != nil {
			wx, wy, wz = OffsetRadialPoint(sx, yLipBot, sz, outerDeltaRAtLocalY(yLipBot-dy))
		}
		cWallLipBot[i] = vector3{X: wx, Y: wy, Z: wz}
		wallBotXs[i] = wx
		wallBotZs[i] = wz
		topXs[i] = cOuterTop[i].X
		topZs[i] = cOuterTop[i].Z
	}

	botOffsets := ComputeRadialLipOffsets(wallBotXs, wallBotZs, lipWidth)
	topOffsets := ComputeRadialLipOffsets(topXs, topZs, lipWidth)

	cLipBot = make([]vector3, M)
	cLipTop = make([]vector3, M)
	for i := range M {
		bx, by, bz := OffsetRadialPoint(cWallLipBot[i].X, cWallLipBot[i].Y, cWallLipBot[i].Z, botOffsets[i])
		cLipBot[i] = vector3{X: bx, Y: by, Z: bz}

		tx, ty, tz := OffsetRadialPoint(cOuterTop[i].X, cOuterTop[i].Y, cOuterTop[i].Z, topOffsets[i])
		cLipTop[i] = vector3{X: tx, Y: ty, Z: tz}
	}
	return cWallLipBot, cLipBot, cLipTop
}
