package models

import "math"

// DefaultVaseTrapezeLipExtraRatio is the default LipHeight ratio (multiplied by RhombusSideLength)
// added past the nominal lower-pot contact edge (~5.04 mm for RhombusSideLength = 120 mm).
const DefaultVaseTrapezeLipExtraRatio = 0.042

func canUseVaseTrapezeConstantThickness(plant *PlantAbstract) bool {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil || !plant.TubeVaseAbstract.HasConstantThickness {
		return false
	}
	if plant.PerpendicularVectorGrid == nil || len(plant.PerpendicularVectorGrid.PerpendicularVectors) < 2 {
		return false
	}
	tanAlpha := math.Tan(-plant.TubeVaseAbstract.ProjectionAngle * math.Pi / 180.0)
	return math.Abs(tanAlpha) >= 1e-6
}

// IsVaseTrapezeLipActive reports whether the load-bearing outer stacking lip
// is enabled and has valid dimensions for a VaseTrapeze plant.
func IsVaseTrapezeLipActive(plant *PlantAbstract) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		plant.TubeVaseAbstract.HasLip &&
		plant.TubeVaseAbstract.LipHeight >= 0 &&
		plant.TubeVaseAbstract.LipWidth > 1e-9 &&
		plant.RhombusSideLength > 1e-9
}

// VaseTrapezeLipDimensionsForMode returns:
//   - lipHeight: total vertical extent of the lip measured downward from the pot top rim,
//     equal to the rotated lower-pot contact step plus LipHeight * RhombusSideLength (~5mm)
//   - lipWidth: outward horizontal perpendicular lip width (LipWidth * RhombusSideLength)
//   - lipExtra: the extra vertical clearance margin (LipHeight * RhombusSideLength)
func VaseTrapezeLipDimensionsForMode(plant *PlantAbstract, hasConstantThickness bool) (lipHeight, lipWidth, lipExtra float64, ok bool) {
	if !IsVaseTrapezeLipActive(plant) {
		return 0, 0, 0, false
	}
	side := plant.RhombusSideLength
	lipExtra = math.Max(0.0, plant.TubeVaseAbstract.LipHeight) * side
	baseContactDY := math.Max(0.0, ComputeVaseTrapezeRotatedContactStepDY(plant, hasConstantThickness))
	lipHeight = baseContactDY + lipExtra
	lipWidth = plant.TubeVaseAbstract.LipWidth * side
	if lipHeight <= 1e-9 || lipWidth <= 1e-9 {
		return 0, 0, 0, false
	}
	return lipHeight, lipWidth, lipExtra, true
}

// VaseTrapezeLipDimensions returns the total scaled lip height (from pot top down to the
// lower pot's top edge plus LipHeight * RhombusSideLength) and outward lip width when active.
func VaseTrapezeLipDimensions(plant *PlantAbstract) (lipHeight, lipWidth float64, ok bool) {
	h, w, _, active := VaseTrapezeLipDimensionsForMode(plant, canUseVaseTrapezeConstantThickness(plant))
	return h, w, active
}

// ComputeVaseTrapezeLipRotatedStepDY returns the vertical step between consecutive rotated
// VaseTrapeze pots so that when HasLip is active, the upper pot's horizontal lip underside
// rests directly on the lower pot's horizontal top rim (lifting the upper pot by lipExtra
// so its lower wall has clearance and puts no wedging stress on the lower pot).
func ComputeVaseTrapezeLipRotatedStepDY(
	plant *PlantAbstract,
	baseRotatedStepDY, ringHeight float64,
	hasConstantThickness bool,
) float64 {
	lipHeight, _, _, hasLip := VaseTrapezeLipDimensionsForMode(plant, hasConstantThickness)
	if !hasLip || ringHeight <= 1e-6 {
		return baseRotatedStepDY
	}
	rotContactDY := ComputeVaseTrapezeRotatedContactStepDY(plant, hasConstantThickness)
	sepDY := baseRotatedStepDY - rotContactDY
	actualLipHeight := ringHeight - VaseTrapezeLipBottomY(0.0, ringHeight, lipHeight)
	return math.Max(rotContactDY, actualLipHeight) + sepDY
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

func computeVaseTrapezeLipGridVec3(
	cOuterBot, cOuterTop, cOuterStraightBot, cOuterUncarvedP1 []vector3,
	yBotRef, spanH, dy, lipHeight, lipWidth float64,
	sCollar int,
	isCurvedWall bool,
	outerDeltaRAtLocalY func(float64) float64,
) (cWallLipBot []vector3, collarRows [][]vector3) {
	M := min(len(cOuterBot), len(cOuterTop), len(cOuterStraightBot), len(cOuterUncarvedP1))
	if M == 0 {
		return nil, nil
	}
	if !isCurvedWall || sCollar < 1 {
		sCollar = 1
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

	collarRows = make([][]vector3, sCollar+1)
	collarRows[0] = make([]vector3, M)
	collarRows[sCollar] = make([]vector3, M)
	for i := range M {
		bx, by, bz := OffsetRadialPoint(cWallLipBot[i].X, cWallLipBot[i].Y, cWallLipBot[i].Z, botOffsets[i])
		collarRows[0][i] = vector3{X: bx, Y: by, Z: bz}

		tx, ty, tz := OffsetRadialPoint(cOuterTop[i].X, cOuterTop[i].Y, cOuterTop[i].Z, topOffsets[i])
		collarRows[sCollar][i] = vector3{X: tx, Y: ty, Z: tz}
	}

	for s := 1; s < sCollar; s++ {
		t := float64(s) / float64(sCollar)
		row := make([]vector3, M)
		for i := range M {
			y := (1.0-t)*cWallLipBot[i].Y + t*cOuterTop[i].Y
			u := 0.0
			if math.Abs(spanH) > 1e-9 {
				u = (y - yBotRef) / spanH
			}
			sx := (1.0-u)*cOuterStraightBot[i].X + u*cOuterUncarvedP1[i].X
			sz := (1.0-u)*cOuterStraightBot[i].Z + u*cOuterUncarvedP1[i].Z
			wx, wy, wz := sx, y, sz
			if isCurvedWall && outerDeltaRAtLocalY != nil {
				wx, wy, wz = OffsetRadialPoint(sx, y, sz, outerDeltaRAtLocalY(y-dy))
			}
			off := (1.0-t)*botOffsets[i] + t*topOffsets[i]
			lx, ly, lz := OffsetRadialPoint(wx, wy, wz, off)
			row[i] = vector3{X: lx, Y: ly, Z: lz}
		}
		collarRows[s] = row
	}

	return cWallLipBot, collarRows
}

func computeVaseTrapezeLipCurvesVec3(
	cOuterBot, cOuterTop, cOuterStraightBot, cOuterUncarvedP1 []vector3,
	yBotRef, spanH, dy, lipHeight, lipWidth float64,
	isCurvedWall bool,
	outerDeltaRAtLocalY func(float64) float64,
) (cWallLipBot, cLipBot, cLipTop []vector3) {
	cWallLipBot, rows := computeVaseTrapezeLipGridVec3(
		cOuterBot, cOuterTop, cOuterStraightBot, cOuterUncarvedP1,
		yBotRef, spanH, dy, lipHeight, lipWidth, 1, isCurvedWall, outerDeltaRAtLocalY,
	)
	if len(rows) == 0 {
		return nil, nil, nil
	}
	return cWallLipBot, rows[0], rows[len(rows)-1]
}
