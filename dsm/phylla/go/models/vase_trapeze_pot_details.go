package models

import "math"

const (
	// DefaultVaseTrapezeFillLineClearanceRatio is the fixed safety clearance below the seated rotated
	// upper pot's bottom (~4.8 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeFillLineClearanceRatio = 0.040

	// DefaultVaseTrapezeFillLineDepthRatio is the perpendicular radial depth of the interior recessed
	// V-groove (~1.44 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeFillLineDepthRatio = 0.012

	// DefaultVaseTrapezeFillLineHeightRatio is the total vertical height of the interior recessed
	// V-groove (~3.0 mm at RhombusSideLength = 120 mm, giving ~46° support-free bevel faces).
	DefaultVaseTrapezeFillLineHeightRatio = 0.025

	// DefaultVaseTrapezeDrainageHoleRadiusRatio is the default central drainage hole radius ratio
	// (~10.2 mm radius / ~Ø 20.4 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeDrainageHoleRadiusRatio = 0.085

	// DefaultVaseTrapezeDrainageSlopeDeg is the gentle conical funnel slope of the interior floor
	// guiding water from the wavy outer lobes toward the central drainage hole.
	DefaultVaseTrapezeDrainageSlopeDeg = 2.0

	// DefaultVaseTrapezeDrainageChamferRatio is the 45° countersink chamfer at the interior drainage
	// hole lip (~2.16 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeDrainageChamferRatio = 0.018

	// DefaultVaseTrapezeFootRimHeightRatio is the upward underside recess / raised foot rim height
	// (~4.2 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeFootRimHeightRatio = 0.035

	// DefaultVaseTrapezeFootRimWidthRatio is the perpendicular horizontal width of the standing
	// ceramic-cup foot rim (~5.4 mm at RhombusSideLength = 120 mm).
	DefaultVaseTrapezeFootRimWidthRatio = 0.045

	// DefaultVaseTrapezeFootRimInsetRatio is the subtle ceramic-cup inward bevel at the outer base
	// edge (~2.4 mm at RhombusSideLength = 120 mm), giving a visible shadow reveal while staying
	// steeper than 45° for support-free 3D printing.
	DefaultVaseTrapezeFootRimInsetRatio = 0.020

	// DefaultVaseTrapezeFootRimNotchCount is the number of evenly spaced drainage breather notches
	// cut across the thin raised foot rim so water and air escape freely on flat surfaces.
	DefaultVaseTrapezeFootRimNotchCount = 3
)

// IsVaseTrapezeFillLevelLineActive reports whether the interior soil fill level line is enabled
// for a closed-bottom VaseTrapeze ring.
func IsVaseTrapezeFillLevelLineActive(plant *PlantAbstract, hasRingBottom bool) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		hasRingBottom &&
		plant.TubeVaseAbstract.HasFillLevelLine &&
		plant.RhombusSideLength > 1e-9
}

// IsVaseTrapezeDrainageHoleActive reports whether the central drainage hole (and 2° sloped floor)
// is enabled for a closed-bottom VaseTrapeze ring.
func IsVaseTrapezeDrainageHoleActive(plant *PlantAbstract, hasRingBottom bool) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		hasRingBottom &&
		plant.TubeVaseAbstract.HasDrainageHole &&
		plant.TubeVaseAbstract.DrainageHoleRadius > 1e-9 &&
		plant.RhombusSideLength > 1e-9
}

// IsVaseTrapezeFootRimActive reports whether the ceramic-cup-style raised foot rim is enabled
// for a closed-bottom VaseTrapeze ring.
func IsVaseTrapezeFootRimActive(plant *PlantAbstract, hasRingBottom bool) bool {
	return plant != nil &&
		plant.PlantType == VaseTrapeze &&
		plant.TubeVaseAbstract != nil &&
		hasRingBottom &&
		plant.TubeVaseAbstract.HasFootRim &&
		plant.TubeVaseAbstract.FootRimHeight > 1e-9 &&
		plant.TubeVaseAbstract.FootRimWidth > 1e-9 &&
		plant.RhombusSideLength > 1e-9
}

// ComputeVaseTrapezeFillLevelLineParams returns the local vertical coordinate yFillLocal,
// perpendicular groove depth, and half vertical groove height of the interior soil fill-level
// V-groove for a closed-bottom VaseTrapeze ring.
//
// yFillLocal is derived purely from the seated rotated stacking step (RotationRatio == 1,
// zero preview separation, accounting for HasLip) minus a fixed 0.04 * RhombusSideLength (~5 mm)
// clearance margin so the groove never drifts with preview sliders.
func ComputeVaseTrapezeFillLevelLineParams(
	plant *PlantAbstract,
	p1H, p2H, hRingBottom float64,
	hasConstantThickness bool,
) (yFillLocal, grooveDepth, grooveHalfH float64, ok bool) {
	if !IsVaseTrapezeFillLevelLineActive(plant, hRingBottom > 0) {
		return 0, 0, 0, false
	}
	ringHeight := p1H - p2H
	if ringHeight <= 1e-6 {
		return 0, 0, 0, false
	}
	side := plant.RhombusSideLength
	rotContact := ComputeVaseTrapezeRotatedContactStepDY(plant, hasConstantThickness)
	rotSeated := ComputeVaseTrapezeLipRotatedStepDY(plant, rotContact, ringHeight, hasConstantThickness)
	if rotSeated <= 1e-6 {
		return 0, 0, 0, false
	}

	clearance := DefaultVaseTrapezeFillLineClearanceRatio * side
	grooveDepth = DefaultVaseTrapezeFillLineDepthRatio * side
	grooveHalfH = 0.5 * DefaultVaseTrapezeFillLineHeightRatio * side
	yFillLocal = p2H + rotSeated - clearance

	yPlancher := p2H + hRingBottom
	if yFillLocal <= yPlancher+2.0*grooveHalfH || yFillLocal >= p1H-2.0*grooveHalfH {
		return 0, 0, 0, false
	}
	return yFillLocal, grooveDepth, grooveHalfH, true
}

// ComputeVaseTrapezeFootRimHeight returns the clamped upward recess height of the raised foot rim
// (<= 45% of hRingBottom so structural floor integrity is preserved).
func ComputeVaseTrapezeFootRimHeight(plant *PlantAbstract, hRingBottom float64, hasRingBottom bool) float64 {
	if !IsVaseTrapezeFootRimActive(plant, hasRingBottom) || hRingBottom <= 1e-6 {
		return 0.0
	}
	h := plant.TubeVaseAbstract.FootRimHeight * plant.RhombusSideLength
	maxH := 0.45 * hRingBottom
	if h > maxH {
		h = maxH
	}
	if h <= 1e-4 {
		return 0.0
	}
	return h
}

// ComputeVaseTrapezeInnerFillGrooveGrid builds the subdivided inner wall geometry split into
// three contiguous bands sharing exact boundary coordinates:
//   - lowerRows (sLower+1 rows): from cInBot up to the lower groove lip yLo
//   - grooveRows (3 rows): [0] = yLo, [1] = yMid (recessed outward V-groove apex), [2] = yHi
//   - upperRows (sUpper+1 rows): from the upper groove lip yHi up to cInTop
//
// Splitting into dedicated bands lets renderers (ThreeJS/GLB) keep razor-sharp creases along the
// V-groove while preserving 100% identical shared coordinates for watertight 2-manifold STLs.
func ComputeVaseTrapezeInnerFillGrooveGrid(
	cInBotX, cInBotY, cInBotZ []float64,
	cInTopX, cInTopY, cInTopZ []float64,
	cInStBotX, cInStBotZ, cInStTopX, cInStTopZ []float64,
	cOutStBotX, cOutStBotZ, cOutStTopX, cOutStTopZ []float64,
	p2H, ringHeight, dy, yContactGlobal float64,
	yFillLocal, grooveDepth, grooveHalfH float64,
	sWall int,
	isCurvedWall bool,
	ringDeltaR func(float64) float64,
) (
	lowerX, lowerY, lowerZ [][]float64,
	grooveX, grooveY, grooveZ [][]float64,
	upperX, upperY, upperZ [][]float64,
) {
	M := min(
		len(cInBotX), len(cInBotY), len(cInBotZ),
		len(cInTopX), len(cInTopY), len(cInTopZ),
		len(cInStBotX), len(cInStBotZ), len(cInStTopX), len(cInStTopZ),
		len(cOutStBotX), len(cOutStBotZ), len(cOutStTopX), len(cOutStTopZ),
	)
	if M < 3 {
		return
	}

	sLower := 1
	sUpper := 1
	if isCurvedWall && sWall > 1 {
		sLower = max(1, sWall/2)
		sUpper = max(1, sWall-sLower)
	}

	evalInnerWall := func(i int, yGlobal float64) (float64, float64, float64) {
		u := 0.0
		if math.Abs(ringHeight) > 1e-9 {
			u = (yGlobal - (p2H + dy)) / ringHeight
		}
		sx := (1.0-u)*cInStBotX[i] + u*cInStTopX[i]
		sz := (1.0-u)*cInStBotZ[i] + u*cInStTopZ[i]
		if isCurvedWall && ringDeltaR != nil {
			return OffsetRadialPoint(sx, yGlobal, sz, ringDeltaR(yGlobal-dy))
		}
		return sx, yGlobal, sz
	}

	evalOuterWall := func(i int, yGlobal float64) (float64, float64, float64) {
		u := 0.0
		if math.Abs(ringHeight) > 1e-9 {
			u = (yGlobal - (p2H + dy)) / ringHeight
		}
		sx := (1.0-u)*cOutStBotX[i] + u*cOutStTopX[i]
		sz := (1.0-u)*cOutStBotZ[i] + u*cOutStTopZ[i]
		if isCurvedWall && ringDeltaR != nil {
			return OffsetRadialPoint(sx, yGlobal, sz, ringDeltaR(yGlobal-dy))
		}
		return sx, yGlobal, sz
	}

	yFillGlobal := yFillLocal + dy

	yLos := make([]float64, M)
	yMids := make([]float64, M)
	yHis := make([]float64, M)
	fades := make([]float64, M)
	midBaseXs := make([]float64, M)
	midBaseZs := make([]float64, M)
	wallThicks := make([]float64, M)

	for i := range M {
		yBot := cInBotY[i]
		yTop := cInTopY[i]
		avail := yTop - yBot
		if avail <= 1e-6 {
			yLos[i] = yBot
			yMids[i] = yBot
			yHis[i] = yBot
			fades[i] = 0.0
			midBaseXs[i], _, midBaseZs[i] = cInBotX[i], yBot, cInBotZ[i]
			continue
		}

		hh := math.Min(grooveHalfH, 0.12*avail)
		minMargin := math.Min(0.20*avail, 2.0*hh)
		yMid := yFillGlobal
		if yMid < yBot+minMargin {
			yMid = yBot + minMargin
		}
		if yMid > yTop-minMargin {
			yMid = yTop - minMargin
		}
		yLo := yMid - hh
		yHi := yMid + hh

		distBot := (yFillGlobal - grooveHalfH) - yBot
		distTop := yTop - (yFillGlobal + grooveHalfH)
		minDist := math.Min(distBot, distTop)
		fade := 0.0
		if minDist >= 2.0*grooveHalfH {
			fade = 1.0
		} else if minDist > 0.0 && grooveHalfH > 1e-9 {
			fade = minDist / (2.0 * grooveHalfH)
		}

		yLos[i] = yLo
		yMids[i] = yMid
		yHis[i] = yHi
		fades[i] = fade

		ix, _, iz := evalInnerWall(i, yMid)
		ox, _, oz := evalOuterWall(i, yMid)
		midBaseXs[i] = ix
		midBaseZs[i] = iz
		rIn := math.Hypot(ix, iz)
		rOut := math.Hypot(ox, oz)
		wallThicks[i] = math.Max(0.0, rOut-rIn)
	}

	radialGrooveOffs := ComputeRadialLipOffsets(midBaseXs, midBaseZs, grooveDepth)

	allocRows := func(nRows int) ([][]float64, [][]float64, [][]float64) {
		rx := make([][]float64, nRows)
		ry := make([][]float64, nRows)
		rz := make([][]float64, nRows)
		for r := range nRows {
			rx[r] = make([]float64, M)
			ry[r] = make([]float64, M)
			rz[r] = make([]float64, M)
		}
		return rx, ry, rz
	}

	lowerX, lowerY, lowerZ = allocRows(sLower + 1)
	grooveX, grooveY, grooveZ = allocRows(3)
	upperX, upperY, upperZ = allocRows(sUpper + 1)

	for i := range M {
		// Groove rows: [0] = yLo, [1] = recessed apex at yMid, [2] = yHi
		lx, ly, lz := evalInnerWall(i, yLos[i])
		grooveX[0][i], grooveY[0][i], grooveZ[0][i] = lx, ly, lz

		effDepth := math.Min(radialGrooveOffs[i], 0.25*wallThicks[i]) * fades[i]
		mx, my, mz := OffsetRadialPoint(midBaseXs[i], yMids[i], midBaseZs[i], effDepth)
		grooveX[1][i], grooveY[1][i], grooveZ[1][i] = mx, my, mz

		hx, hy, hz := evalInnerWall(i, yHis[i])
		grooveX[2][i], grooveY[2][i], grooveZ[2][i] = hx, hy, hz

		// Lower band: cInBot -> groove[0]
		lowerX[0][i], lowerY[0][i], lowerZ[0][i] = cInBotX[i], cInBotY[i], cInBotZ[i]
		lowerX[sLower][i], lowerY[sLower][i], lowerZ[sLower][i] = lx, ly, lz
		useContactLower := sLower >= 3 && yContactGlobal > cInBotY[i]+1e-3 && yContactGlobal < yLos[i]-1e-3
		for s := 1; s < sLower; s++ {
			var y float64
			if useContactLower {
				y = RingWallSampleY(s, sLower, cInBotY[i], yLos[i], yContactGlobal)
			} else {
				y = cInBotY[i] + (float64(s)/float64(sLower))*(yLos[i]-cInBotY[i])
			}
			px, py, pz := evalInnerWall(i, y)
			lowerX[s][i], lowerY[s][i], lowerZ[s][i] = px, py, pz
		}

		// Upper band: groove[2] -> cInTop
		upperX[0][i], upperY[0][i], upperZ[0][i] = hx, hy, hz
		upperX[sUpper][i], upperY[sUpper][i], upperZ[sUpper][i] = cInTopX[i], cInTopY[i], cInTopZ[i]
		useContactUpper := sUpper >= 3 && yContactGlobal > yHis[i]+1e-3 && yContactGlobal < cInTopY[i]-1e-3
		for s := 1; s < sUpper; s++ {
			var y float64
			if useContactUpper {
				y = RingWallSampleY(s, sUpper, yHis[i], cInTopY[i], yContactGlobal)
			} else {
				y = yHis[i] + (float64(s)/float64(sUpper))*(cInTopY[i]-yHis[i])
			}
			px, py, pz := evalInnerWall(i, y)
			upperX[s][i], upperY[s][i], upperZ[s][i] = px, py, pz
		}
	}

	return lowerX, lowerY, lowerZ, grooveX, grooveY, grooveZ, upperX, upperY, upperZ
}

// computeSafeFloorCenter returns the centroid (cx, cz) of the interior plancher contour, pulled
// smoothly toward (0, 0) only if needed so both inner and outer bottom contours are strictly
// star-shaped around (cx, cz) with monotonic azimuth and positive radial clearance.
func computeSafeFloorCenter(inX, inZ, outX, outZ []float64) (cx, cz float64) {
	M := min(len(inX), len(inZ), len(outX), len(outZ))
	if M < 3 {
		return 0, 0
	}
	var sumX, sumZ float64
	for i := range M {
		sumX += inX[i]
		sumZ += inZ[i]
	}
	cx0 := sumX / float64(M)
	cz0 := sumZ / float64(M)

	// Determine reference winding sign around (0, 0)
	var refCross float64
	for i := range M {
		nextI := (i + 1) % M
		refCross += inX[i]*inZ[nextI] - inZ[i]*inX[nextI]
	}
	sign := 1.0
	if refCross < 0 {
		sign = -1.0
	}

	isStarShaped := func(px, pz float64) bool {
		checkLoop := func(xs, zs []float64) bool {
			for i := range M {
				nextI := (i + 1) % M
				dx0, dz0 := xs[i]-px, zs[i]-pz
				dx1, dz1 := xs[nextI]-px, zs[nextI]-pz
				if math.Hypot(dx0, dz0) < 1e-3 {
					return false
				}
				cross := dx0*dz1 - dz0*dx1
				if sign*cross <= 1e-6 {
					return false
				}
			}
			return true
		}
		return checkLoop(inX, inZ) && checkLoop(outX, outZ)
	}

	if isStarShaped(cx0, cz0) {
		return cx0, cz0
	}
	lo, hi := 0.0, 1.0
	for range 20 {
		mid := 0.5 * (lo + hi)
		if isStarShaped(mid*cx0, mid*cz0) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo * cx0, lo * cz0
}

// ComputeVaseTrapezeRingBottomApparatus computes all nested horizontal/vertical contour loops
// for the bottom slab of a closed-bottom VaseTrapeze ring, handling all combinations of
// HasFootRim (with 45° outer ceramic inset, 3 drainage breather notches, and 45° inner bevel)
// and HasDrainageHole (with centroid placement, 2° sloped interior floor, 45° countersink lip,
// and vertical cylindrical drainage bore).
//
// Returned contour slices all have length M (or nil when inactive):
//   - cOuterWallBot: bottom edge of outer pot wall at yRecess = p2H + dy + footH
//   - cFootOut, cFootIn, cRecessOut: raised foot rim contours when activeFootRim is true
//   - cHoleChamfer, cHoleTop, cHoleBot: drainage hole contours when activeHole is true
//   - centerX, centerZ: safe floor center shared by interior plancher and underside
//   - yRecess: global Y of the flat recessed exterior underside ceiling
func ComputeVaseTrapezeRingBottomApparatus(
	plant *PlantAbstract,
	cOuterStP2X, cOuterStP2Z, cOuterStP1X, cOuterStP1Z []float64,
	cInnerPlX, cInnerPlY, cInnerPlZ []float64,
	p2H, ringHeight, hRingBottom, dy float64,
	isCurvedWall bool,
	ringDeltaR func(float64) float64,
) (
	cOuterWallBotX, cOuterWallBotY, cOuterWallBotZ []float64,
	cFootOutX, cFootOutY, cFootOutZ []float64,
	cFootInX, cFootInY, cFootInZ []float64,
	cRecessOutX, cRecessOutY, cRecessOutZ []float64,
	cHoleChamferX, cHoleChamferY, cHoleChamferZ []float64,
	cHoleTopX, cHoleTopY, cHoleTopZ []float64,
	cHoleBotX, cHoleBotY, cHoleBotZ []float64,
	centerX, centerZ, yRecess float64,
	activeFootRim, activeHole bool,
) {
	M := min(
		len(cOuterStP2X), len(cOuterStP2Z), len(cOuterStP1X), len(cOuterStP1Z),
		len(cInnerPlX), len(cInnerPlY), len(cInnerPlZ),
	)
	if M < 3 || plant == nil || plant.TubeVaseAbstract == nil || hRingBottom <= 1e-6 || ringHeight <= 1e-6 {
		return
	}

	vase := plant.TubeVaseAbstract
	side := plant.RhombusSideLength
	yFloorGlobal := p2H + dy

	footH := ComputeVaseTrapezeFootRimHeight(plant, hRingBottom, true)
	activeFootRim = footH > 1e-4
	yRecess = yFloorGlobal + footH

	// 1. Compute outer wall bottom contour cOuterWallBot at y = yRecess
	cOuterWallBotX = make([]float64, M)
	cOuterWallBotY = make([]float64, M)
	cOuterWallBotZ = make([]float64, M)
	uRecess := footH / ringHeight
	localYRecess := p2H + footH
	drRecess := 0.0
	if isCurvedWall && ringDeltaR != nil {
		drRecess = ringDeltaR(localYRecess)
	}
	for i := range M {
		sx := (1.0-uRecess)*cOuterStP2X[i] + uRecess*cOuterStP1X[i]
		sz := (1.0-uRecess)*cOuterStP2Z[i] + uRecess*cOuterStP1Z[i]
		ox, oy, oz := OffsetRadialPoint(sx, yRecess, sz, drRecess)
		cOuterWallBotX[i] = ox
		cOuterWallBotY[i] = oy
		cOuterWallBotZ[i] = oz
	}

	// 2. Safe floor centroid (cx, cz) shared by interior plancher, drainage hole, and underside
	centerX, centerZ = computeSafeFloorCenter(cInnerPlX, cInnerPlZ, cOuterWallBotX, cOuterWallBotZ)

	// Measure radial geometry around (centerX, centerZ)
	dOut := make([]float64, M)
	uxOut := make([]float64, M)
	uzOut := make([]float64, M)
	dMinOut := math.MaxFloat64

	dIn := make([]float64, M)
	uxIn := make([]float64, M)
	uzIn := make([]float64, M)
	dMinIn := math.MaxFloat64
	sumDIn := 0.0
	sumYPlancher := 0.0

	for i := range M {
		dxO, dzO := cOuterWallBotX[i]-centerX, cOuterWallBotZ[i]-centerZ
		do := math.Hypot(dxO, dzO)
		dOut[i] = do
		if do > 1e-9 {
			uxOut[i], uzOut[i] = dxO/do, dzO/do
		} else {
			uxOut[i], uzOut[i] = 1.0, 0.0
		}
		if do < dMinOut {
			dMinOut = do
		}

		dxI, dzI := cInnerPlX[i]-centerX, cInnerPlZ[i]-centerZ
		di := math.Hypot(dxI, dzI)
		dIn[i] = di
		if di > 1e-9 {
			uxIn[i], uzIn[i] = dxI/di, dzI/di
		} else {
			uxIn[i], uzIn[i] = 1.0, 0.0
		}
		if di < dMinIn {
			dMinIn = di
		}
		sumDIn += di
		sumYPlancher += cInnerPlY[i]
	}
	meanDIn := sumDIn / float64(M)
	yPlancherEdge := sumYPlancher / float64(M)

	// 3. Determine clamped drainage hole radius rHole
	footInsetNom := 0.0
	footWidthNom := 0.0
	innerBevelNom := 0.0
	if activeFootRim {
		footInsetNom = math.Min(DefaultVaseTrapezeFootRimInsetRatio*side, 0.75*footH)
		footWidthNom = vase.FootRimWidth * side
		innerBevelNom = 0.60 * footH
	}
	nomFootSpan := footInsetNom + footWidthNom + innerBevelNom
	maxOuterEncroach := math.Min(nomFootSpan, 0.45*dMinOut)

	rHole := 0.0
	if IsVaseTrapezeDrainageHoleActive(plant, true) && yPlancherEdge > yRecess+1e-3 {
		minClearRadius := math.Min(dMinIn, dMinOut-maxOuterEncroach)
		rHole = math.Min(vase.DrainageHoleRadius*side, 0.45*minClearRadius)
		if rHole >= 0.5 {
			activeHole = true
		} else {
			rHole = 0.0
		}
	}

	// 4. Build raised ceramic-cup foot rim contours with 3 subtle drainage breather notches
	if activeFootRim {
		shiftedOutX := make([]float64, M)
		shiftedOutZ := make([]float64, M)
		for i := range M {
			shiftedOutX[i] = cOuterWallBotX[i] - centerX
			shiftedOutZ[i] = cOuterWallBotZ[i] - centerZ
		}
		normFactors := ComputeRadialLipOffsets(shiftedOutX, shiftedOutZ, 1.0)

		// Cumulative arc length along cOuterWallBot for uniform physical notch sizing
		cumS := make([]float64, M)
		totalL := 0.0
		for i := range M {
			cumS[i] = totalL
			nextI := (i + 1) % M
			totalL += math.Hypot(cOuterWallBotX[nextI]-cOuterWallBotX[i], cOuterWallBotZ[nextI]-cOuterWallBotZ[i])
		}

		const nNotches = DefaultVaseTrapezeFootRimNotchCount
		wFlat := math.Min(2.4*footH, totalL/18.0)
		wRamp := math.Min(1.5*footH, totalL/32.0)
		halfFlat := 0.5 * wFlat
		halfTotal := halfFlat + wRamp

		notchFracAt := func(s float64) float64 {
			if totalL <= 1e-6 || halfTotal <= 1e-6 {
				return 0.0
			}
			best := 0.0
			for k := range nNotches {
				// Place notches at (2*k+1)/(2*nNotches) of perimeter so none straddles the i=0 seam
				centerS := (float64(2*k+1) / float64(2*nNotches)) * totalL
				dist := math.Abs(s - centerS)
				if dist > 0.5*totalL {
					dist = totalL - dist
				}
				if dist <= halfFlat {
					return 1.0
				}
				if dist < halfTotal && wRamp > 1e-9 {
					f := (halfTotal - dist) / wRamp
					if f > best {
						best = f
					}
				}
			}
			return best
		}

		cFootOutX = make([]float64, M)
		cFootOutY = make([]float64, M)
		cFootOutZ = make([]float64, M)
		cFootInX = make([]float64, M)
		cFootInY = make([]float64, M)
		cFootInZ = make([]float64, M)
		cRecessOutX = make([]float64, M)
		cRecessOutY = make([]float64, M)
		cRecessOutZ = make([]float64, M)

		for i := range M {
			nf := notchFracAt(cumS[i])
			// Lift up to 88% of footH inside each notch so water & air escape freely while
			// preserving >= 12% vertical wall height (zero degenerate triangles).
			yFoot := yFloorGlobal + 0.88*footH*nf
			localDrop := yRecess - yFoot

			f := normFactors[i]
			shiftOut := footInsetNom * f
			shiftRim := footWidthNom * f
			shiftInBevel := 0.60 * localDrop * f

			minInnerLim := 0.25 * dOut[i]
			if activeHole {
				minInnerLim = math.Max(minInnerLim, rHole+math.Max(1.5, 0.02*side))
			}
			availR := math.Max(0.3, dOut[i]-minInnerLim)
			sumShift := shiftOut + shiftRim + shiftInBevel
			scale := 1.0
			if sumShift > availR && sumShift > 1e-9 {
				scale = availR / sumShift
			}

			dFO := dOut[i] - shiftOut*scale
			dFI := dFO - shiftRim*scale
			dRO := dFI - shiftInBevel*scale

			cFootOutX[i] = centerX + dFO*uxOut[i]
			cFootOutY[i] = yFoot
			cFootOutZ[i] = centerZ + dFO*uzOut[i]

			cFootInX[i] = centerX + dFI*uxOut[i]
			cFootInY[i] = yFoot
			cFootInZ[i] = centerZ + dFI*uzOut[i]

			cRecessOutX[i] = centerX + dRO*uxOut[i]
			cRecessOutY[i] = yRecess
			cRecessOutZ[i] = centerZ + dRO*uzOut[i]
		}
	}

	// 5. Build central drainage hole contours + gentle 2° interior funnel slope & 45° countersink lip
	if activeHole {
		slabThick := yPlancherEdge - yRecess
		tanSlope := math.Tan(DefaultVaseTrapezeDrainageSlopeDeg * math.Pi / 180.0)
		funnelDrop := math.Max(0.0, (meanDIn-rHole)*tanSlope)
		if maxDrop := 0.30 * slabThick; funnelDrop > maxDrop {
			funnelDrop = maxDrop
		}
		yHoleChamferTop := yPlancherEdge - funnelDrop
		availBore := yHoleChamferTop - yRecess

		chamfer := DefaultVaseTrapezeDrainageChamferRatio * side
		if cMax := 0.25 * rHole; chamfer > cMax {
			chamfer = cMax
		}
		if cMax := 0.22 * availBore; chamfer > cMax {
			chamfer = cMax
		}
		if cMax := 0.85*dMinIn - rHole; chamfer > cMax {
			chamfer = math.Max(0.0, cMax)
		}
		yHoleTop := yHoleChamferTop - chamfer

		cHoleChamferX = make([]float64, M)
		cHoleChamferY = make([]float64, M)
		cHoleChamferZ = make([]float64, M)
		cHoleTopX = make([]float64, M)
		cHoleTopY = make([]float64, M)
		cHoleTopZ = make([]float64, M)
		cHoleBotX = make([]float64, M)
		cHoleBotY = make([]float64, M)
		cHoleBotZ = make([]float64, M)

		rCh := rHole + chamfer
		for i := range M {
			ux, uz := uxIn[i], uzIn[i]

			cHoleChamferX[i] = centerX + rCh*ux
			cHoleChamferY[i] = yHoleChamferTop
			cHoleChamferZ[i] = centerZ + rCh*uz

			cHoleTopX[i] = centerX + rHole*ux
			cHoleTopY[i] = yHoleTop
			cHoleTopZ[i] = centerZ + rHole*uz

			cHoleBotX[i] = centerX + rHole*ux
			cHoleBotY[i] = yRecess
			cHoleBotZ[i] = centerZ + rHole*uz
		}
	}

	return
}

// Unexported vector3 helpers for models-package exporters (STL, GLB, Blender Terrace).

func unpackVec3Coords(pts []vector3) (xs, ys, zs []float64) {
	xs = make([]float64, len(pts))
	ys = make([]float64, len(pts))
	zs = make([]float64, len(pts))
	for i, p := range pts {
		xs[i], ys[i], zs[i] = p.X, p.Y, p.Z
	}
	return xs, ys, zs
}

func packVec3Coords(xs, ys, zs []float64) []vector3 {
	n := min(len(xs), len(ys), len(zs))
	if n == 0 {
		return nil
	}
	out := make([]vector3, n)
	for i := range n {
		out[i] = vector3{X: xs[i], Y: ys[i], Z: zs[i]}
	}
	return out
}

func packVec3Grid(gx, gy, gz [][]float64) [][]vector3 {
	if len(gx) == 0 {
		return nil
	}
	rows := make([][]vector3, len(gx))
	for r := range gx {
		rows[r] = packVec3Coords(gx[r], gy[r], gz[r])
	}
	return rows
}

type ringBottomApparatusVec3 struct {
	cOuterWallBot []vector3
	cFootOut      []vector3
	cFootIn       []vector3
	cRecessOut    []vector3
	cHoleChamfer  []vector3
	cHoleTop      []vector3
	cHoleBot      []vector3
	centerX       float64
	centerZ       float64
	yRecess       float64
	activeFootRim bool
	activeHole    bool
}

func computeRingBottomApparatusVec3(
	plant *PlantAbstract,
	cOuterStP2, cOuterStP1, cInnerPlancher []vector3,
	p2H, ringHeight, hRingBottom, dy float64,
	isCurvedWall bool,
	ringDeltaR func(float64) float64,
) ringBottomApparatusVec3 {
	oP2X, _, oP2Z := unpackVec3Coords(cOuterStP2)
	oP1X, _, oP1Z := unpackVec3Coords(cOuterStP1)
	iPlX, iPlY, iPlZ := unpackVec3Coords(cInnerPlancher)

	owX, owY, owZ,
		foX, foY, foZ,
		fiX, fiY, fiZ,
		roX, roY, roZ,
		hcX, hcY, hcZ,
		htX, htY, htZ,
		hbX, hbY, hbZ,
		cx, cz, yRec,
		hasFoot, hasHole := ComputeVaseTrapezeRingBottomApparatus(
		plant,
		oP2X, oP2Z, oP1X, oP1Z,
		iPlX, iPlY, iPlZ,
		p2H, ringHeight, hRingBottom, dy,
		isCurvedWall, ringDeltaR,
	)

	return ringBottomApparatusVec3{
		cOuterWallBot: packVec3Coords(owX, owY, owZ),
		cFootOut:      packVec3Coords(foX, foY, foZ),
		cFootIn:       packVec3Coords(fiX, fiY, fiZ),
		cRecessOut:    packVec3Coords(roX, roY, roZ),
		cHoleChamfer:  packVec3Coords(hcX, hcY, hcZ),
		cHoleTop:      packVec3Coords(htX, htY, htZ),
		cHoleBot:      packVec3Coords(hbX, hbY, hbZ),
		centerX:       cx,
		centerZ:       cz,
		yRecess:       yRec,
		activeFootRim: hasFoot,
		activeHole:    hasHole,
	}
}

func computeInnerFillGrooveGridVec3(
	cInBot, cInTop, cInStBot, cInStTop, cOutStBot, cOutStTop []vector3,
	p2H, ringHeight, dy, yContactGlobal float64,
	yFillLocal, grooveDepth, grooveHalfH float64,
	sWall int,
	isCurvedWall bool,
	ringDeltaR func(float64) float64,
) (lowerRows, grooveRows, upperRows [][]vector3) {
	ibX, ibY, ibZ := unpackVec3Coords(cInBot)
	itX, itY, itZ := unpackVec3Coords(cInTop)
	isbX, _, isbZ := unpackVec3Coords(cInStBot)
	istX, _, istZ := unpackVec3Coords(cInStTop)
	osbX, _, osbZ := unpackVec3Coords(cOutStBot)
	ostX, _, ostZ := unpackVec3Coords(cOutStTop)

	lx, ly, lz, gx, gy, gz, ux, uy, uz := ComputeVaseTrapezeInnerFillGrooveGrid(
		ibX, ibY, ibZ,
		itX, itY, itZ,
		isbX, isbZ, istX, istZ,
		osbX, osbZ, ostX, ostZ,
		p2H, ringHeight, dy, yContactGlobal,
		yFillLocal, grooveDepth, grooveHalfH,
		sWall, isCurvedWall, ringDeltaR,
	)
	return packVec3Grid(lx, ly, lz), packVec3Grid(gx, gy, gz), packVec3Grid(ux, uy, uz)
}
