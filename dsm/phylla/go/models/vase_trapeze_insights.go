package models

import "math"

const (
	// StandardTerracottaDensityGPerCm3 is the standard bulk density of dry fired
	// horticultural terracotta flower pots (1.85 g/cm³ = 1850 kg/m³; commercial porous
	// earthenware pots typically range from 1.80 to 2.00 g/cm³).
	StandardTerracottaDensityGPerCm3 = 1.85

	// StandardPotRimToWallRatio is the standard reinforced rim/lip total thickness ratio
	// relative to the lower pot wall on commercial outdoor terracotta planters (~1.35x wall,
	// adding a ~35% outward rim bead for hoop stiffness, handling grip, and stacking seat).
	StandardPotRimToWallRatio = 1.35
)

// StandardWorkingPotWallThickness returns the standard wall thickness (in mm) for a working
// outdoor fired-terracotta flower pot of equivalent outer-wall top diameter outerWallDiamMM.
//
// Calibrated to commercial pressed horticultural terracotta planters (e.g. Deroma):
//   - Ø 150 mm ->  7.4 mm
//   - Ø 300 mm -> 10.2 mm
//   - Ø 420 mm -> 12.5 mm
//   - Ø 600 mm -> 15.9 mm
func StandardWorkingPotWallThickness(outerWallDiamMM float64) float64 {
	if outerWallDiamMM <= 0 {
		return 5.0
	}
	return math.Max(5.0, 4.5+0.019*outerWallDiamMM)
}

// StandardWorkingPotLipThickness returns the standard total radial rim/lip thickness (in mm)
// for a working terracotta flower pot of equivalent outer-wall top diameter outerWallDiamMM.
func StandardWorkingPotLipThickness(outerWallDiamMM float64) float64 {
	return StandardPotRimToWallRatio * StandardWorkingPotWallThickness(outerWallDiamMM)
}

// StandardWorkingPotBottomThickness returns the standard net structural floor-slab thickness
// (in mm) for a working terracotta flower pot of equivalent outer-wall top diameter outerWallDiamMM,
// sized to support saturated potting soil on a perimeter foot ring without plate-bending failure
// or differential kiln-drying cracks (~1.18x wall thickness: ~14.8 mm at Ø 420 mm).
func StandardWorkingPotBottomThickness(outerWallDiamMM float64) float64 {
	if outerWallDiamMM <= 0 {
		return 6.0
	}
	return math.Max(6.0, 5.5+0.022*outerWallDiamMM)
}

type vaseTrapezePotMetrics struct {
	weightKg          float64
	solidVolumeCm3    float64
	outerWallDiamMM   float64
	outerRimDiamMM    float64
	potHeightMM       float64
	wallThickMM       float64
	lipOverhangMM     float64
	lipThickMM        float64
	lipCollarHeightMM float64
	stdWallThickMM    float64
	stdLipThickMM     float64
	lipDiffMM         float64
	lipDiffPct        float64
	bottomSlabMM      float64
	bottomDrainMM     float64
	bottomGrossMM     float64
	footRimHeightMM   float64
	stdBottomThickMM  float64
	bottomDiffMM      float64
	bottomDiffPct     float64
	hasLip            bool
	hasRingBottom     bool
}

func polygonSignedAreaXZ(pts []vector3) float64 {
	n := len(pts)
	if n < 3 {
		return 0.0
	}
	var sum float64
	for i := range n {
		nextI := (i + 1) % n
		sum += pts[i].X*pts[nextI].Z - pts[i].Z*pts[nextI].X
	}
	return 0.5 * math.Abs(sum)
}

func computeVaseTrapezePotMetrics(plant *PlantAbstract) (m vaseTrapezePotMetrics, ok bool) {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil || plant.RhombusSideLength <= 1e-9 {
		return m, false
	}
	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)
	sideLength := plant.RhombusSideLength

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
	globalR := circumference * float64(radialRepetitions) / (2.0 * math.Pi)

	gc := plant.GrowthCurve2D
	tgc := vase.TopGrowthCurve2D
	if gc == nil || tgc == nil ||
		gc.StartHalfwayArcShapeGrid == nil ||
		tgc.TopStartHalfwayArcShapeGrid == nil ||
		len(gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) == 0 ||
		len(tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) == 0 {
		return m, false
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
		return m, false
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)
	resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
	resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)
	if len(resampledBaseBottom) != len(targetAngles) || len(resampledBaseTop) != len(targetAngles) {
		return m, false
	}

	// Evaluate the physical pot at rotBlend == 1.0 (manufactured rotated-stacking ring geometry)
	// so preview sliders (RotationRatio, RelativeRotatedTorusSeparation) do not drift physical metrics.
	hasTrapezeConstantThickness := false
	baseBottomY := make([]float64, len(targetAngles))
	for i, pt := range resampledBaseBottom {
		baseBottomY[i] = pt.Y
	}
	if newBottomY, newTopY, _, _, _, constOk := computeVaseTrapezeConstant3DThicknessWithBlend(plant, globalR, targetAngles, baseBottomY, 1.0); constOk {
		for i := range targetAngles {
			resampledBaseBottom[i].Y = newBottomY[i]
			resampledBaseTop[i].Y = newTopY[i]
		}
		hasTrapezeConstantThickness = true
	}

	p1H := zStart + vase.TopPlaneHeight*sideLength*scaleY
	p2H := zStart + vase.BottomPlaneHeight*sideLength*scaleY
	ringHeight := p1H - p2H
	if ringHeight <= 1e-6 {
		return m, false
	}

	// 1. Exact enclosed solid volume via 3D divergence theorem on the watertight OneRing mesh.
	// Temporarily pin RotationRatio to 1.0 while emitting facets so ComputeVaseTrapezeContactStepDYForCurves
	// also evaluates the manufactured rotated-contact pot geometry.
	origRotRatio := vase.RotationRatio
	vase.RotationRatio = 1.0
	var signedVol6 float64
	emitVaseTrapezeFacets(func(v1, v2, v3 vector3) {
		signedVol6 += v1.X*(v2.Y*v3.Z-v2.Z*v3.Y) +
			v1.Y*(v2.Z*v3.X-v2.X*v3.Z) +
			v1.Z*(v2.X*v3.Y-v2.Y*v3.X)
	}, plant, resampledBaseBottom, resampledBaseTop, radialRepetitions, STLExportOneRing, hasTrapezeConstantThickness)
	vase.RotationRatio = origRotRatio

	volMM3 := math.Abs(signedVol6) / 6.0
	if volMM3 <= 1e-6 {
		return m, false
	}
	m.solidVolumeCm3 = volMM3 / 1000.0
	m.weightKg = (m.solidVolumeCm3 * StandardTerracottaDensityGPerCm3) / 1000.0
	m.potHeightMM = ringHeight

	// 2. Project top and bottom contours to measure exact perpendicular wall, lip, and bottom thicknesses.
	trapezeBottomCurve, trapezeTopCurve := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	tanAlpha := math.Tan(projAngleRad)

	projectLoop := func(src []vector3, planeHeight float64) []vector3 {
		var out []vector3
		for _, pt := range src {
			th := math.Atan2(pt.Z, pt.X)
			r := math.Hypot(pt.X, pt.Z)
			rProj := r + (planeHeight-pt.Y)*tanAlpha
			np := vector3{X: rProj * math.Cos(th), Y: planeHeight, Z: rProj * math.Sin(th)}
			if len(out) > 0 {
				prev := out[len(out)-1]
				if math.Hypot(np.X-prev.X, np.Z-prev.Z) < 1e-3 {
					continue
				}
			}
			out = append(out, np)
		}
		if len(out) > 2 {
			first, last := out[0], out[len(out)-1]
			if math.Hypot(last.X-first.X, last.Z-first.Z) < 1e-3 {
				out = out[:len(out)-1]
			}
		}
		return out
	}

	cTopP1 := projectLoop(trapezeTopCurve, p1H)
	cBottomP1 := projectLoop(trapezeBottomCurve, p1H)
	cTopP2 := projectLoop(trapezeTopCurve, p2H)
	cBottomP2 := projectLoop(trapezeBottomCurve, p2H)
	M := min(len(cTopP1), len(cBottomP1), len(cTopP2), len(cBottomP2))
	if M < 3 {
		return m, false
	}
	cTopP1, cBottomP1, cTopP2, cBottomP2 = cTopP1[:M], cBottomP1[:M], cTopP2[:M], cBottomP2[:M]

	var sumRTop, sumRBot float64
	for i := range M {
		sumRTop += math.Hypot(cTopP2[i].X, cTopP2[i].Z)
		sumRBot += math.Hypot(cBottomP2[i].X, cBottomP2[i].Z)
	}
	cOuterUncarvedP1, cInnerUncarvedP1 := cTopP1, cBottomP1
	cOuterStraightP2, cInnerStraightP2 := cTopP2, cBottomP2
	if sumRBot > sumRTop {
		cOuterUncarvedP1, cInnerUncarvedP1 = cBottomP1, cTopP1
		cOuterStraightP2, cInnerStraightP2 = cBottomP2, cTopP2
	}

	outerWallArea := polygonSignedAreaXZ(cOuterUncarvedP1)
	if outerWallArea <= 1e-6 {
		return m, false
	}
	m.outerWallDiamMM = 2.0 * math.Sqrt(outerWallArea/math.Pi)

	outXs := make([]float64, M)
	outZs := make([]float64, M)
	for i := range M {
		outXs[i] = cOuterUncarvedP1[i].X
		outZs[i] = cOuterUncarvedP1[i].Z
	}
	normFactors := ComputeRadialLipOffsets(outXs, outZs, 1.0)
	var sumWallPerp float64
	for i := range M {
		rOut := math.Hypot(cOuterUncarvedP1[i].X, cOuterUncarvedP1[i].Z)
		rIn := math.Hypot(cInnerUncarvedP1[i].X, cInnerUncarvedP1[i].Z)
		f := normFactors[i]
		if f < 1.0 {
			f = 1.0
		}
		sumWallPerp += math.Max(0.0, rOut-rIn) / f
	}
	m.wallThickMM = sumWallPerp / float64(M)

	lipHeight, lipWidth, _, hasLip := VaseTrapezeLipDimensionsForMode(plant, hasTrapezeConstantThickness)
	m.hasLip = hasLip
	if hasLip {
		m.lipOverhangMM = lipWidth
		m.lipThickMM = m.wallThickMM + lipWidth
		cRimOut := make([]vector3, M)
		for i := range M {
			rx, ry, rz := OffsetRadialPoint(cOuterUncarvedP1[i].X, p1H, cOuterUncarvedP1[i].Z, lipWidth*normFactors[i])
			cRimOut[i] = vector3{X: rx, Y: ry, Z: rz}
		}
		if rimArea := polygonSignedAreaXZ(cRimOut); rimArea > 1e-6 {
			m.outerRimDiamMM = 2.0 * math.Sqrt(rimArea/math.Pi)
		} else {
			m.outerRimDiamMM = m.outerWallDiamMM + 2.0*lipWidth
		}
	} else {
		m.lipOverhangMM = 0.0
		m.lipThickMM = m.wallThickMM
		m.outerRimDiamMM = m.outerWallDiamMM
	}

	// 3. Bottom slab thickness (accounting for HasRingBottom, FootRim recess, and 2° drainage funnel slope)
	basePlateHeight := vase.BasePlateHeight * sideLength
	m.hasRingBottom = vase.HasRingBottom && basePlateHeight > 1e-6 && ringHeight > 1e-6
	yOuterWallBot := p2H
	if m.hasRingBottom {
		hRingBottom := math.Min(basePlateHeight, ringHeight)
		m.bottomGrossMM = hRingBottom
		m.footRimHeightMM = ComputeVaseTrapezeFootRimHeight(plant, hRingBottom, true)
		m.bottomSlabMM = math.Max(0.0, hRingBottom-m.footRimHeightMM)
		m.bottomDrainMM = m.bottomSlabMM
		yOuterWallBot = p2H + m.footRimHeightMM

		if IsVaseTrapezeDrainageHoleActive(plant, true) {
			u := hRingBottom / ringHeight
			cInnerPlancher := make([]vector3, M)
			deltaYContact := ComputeVaseTrapezeRotatedContactStepDY(plant, hasTrapezeConstantThickness)
			isCurvedWall := IsVaseTrapezeCurvedWallActive(plant)
			for i := range M {
				sx := (1.0-u)*cInnerStraightP2[i].X + u*cInnerUncarvedP1[i].X
				sy := (1.0-u)*cInnerStraightP2[i].Y + u*cInnerUncarvedP1[i].Y
				sz := (1.0-u)*cInnerStraightP2[i].Z + u*cInnerUncarvedP1[i].Z
				if isCurvedWall {
					dr := EvalVaseTrapezeRingWallDeltaR(plant, sy, p1H, p2H, deltaYContact)
					sx, sy, sz = OffsetRadialPoint(sx, sy, sz, dr)
				}
				cInnerPlancher[i] = vector3{X: sx, Y: sy, Z: sz}
			}
			app := computeRingBottomApparatusVec3(
				plant, cOuterStraightP2, cOuterUncarvedP1, cInnerPlancher,
				p2H, ringHeight, hRingBottom, 0.0, isCurvedWall,
				func(localY float64) float64 {
					return EvalVaseTrapezeRingWallDeltaR(plant, localY, p1H, p2H, deltaYContact)
				},
			)
			if app.activeHole && len(app.cHoleChamfer) > 0 {
				m.bottomDrainMM = math.Max(0.0, app.cHoleChamfer[0].Y-app.yRecess)
			}
		}
	}

	if hasLip {
		m.lipCollarHeightMM = p1H - VaseTrapezeLipBottomY(yOuterWallBot, p1H, lipHeight)
	}

	// 4. Compare against standard working terracotta pot thicknesses for this diameter class.
	m.stdWallThickMM = StandardWorkingPotWallThickness(m.outerWallDiamMM)
	m.stdLipThickMM = StandardWorkingPotLipThickness(m.outerWallDiamMM)
	m.stdBottomThickMM = StandardWorkingPotBottomThickness(m.outerWallDiamMM)

	m.lipDiffMM = m.lipThickMM - m.stdLipThickMM
	if m.stdLipThickMM > 1e-9 {
		m.lipDiffPct = (m.lipDiffMM / m.stdLipThickMM) * 100.0
	}
	m.bottomDiffMM = m.bottomSlabMM - m.stdBottomThickMM
	if m.stdBottomThickMM > 1e-9 {
		m.bottomDiffPct = (m.bottomDiffMM / m.stdBottomThickMM) * 100.0
	}

	return m, true
}

// ComputeVaseTrapezePotWeightInsight computes the empty pot weight (in kg) and solid volume
// (in cm³) of a single VaseTrapeze terrace pot assuming standard fired terracotta (1.85 g/cm³),
// along with its area-equivalent outer rim diameter and pot height in mm.
func ComputeVaseTrapezePotWeightInsight(
	plant *PlantAbstract,
) (weightKg, solidVolumeCm3, outerRimDiamMM, potHeightMM float64, hasRingBottom, ok bool) {
	m, valid := computeVaseTrapezePotMetrics(plant)
	if !valid {
		return 0, 0, 0, 0, false, false
	}
	return m.weightKg, m.solidVolumeCm3, m.outerRimDiamMM, m.potHeightMM, m.hasRingBottom, true
}

// ComputeVaseTrapezePotThicknessInsights computes the perpendicular rim/lip thickness and net
// structural bottom-slab thickness of a VaseTrapeze terrace pot (in mm) and how far each is
// from standard working terracotta pot thicknesses for the pot's diameter class.
func ComputeVaseTrapezePotThicknessInsights(
	plant *PlantAbstract,
) (
	lipThickMM, stdLipThickMM, lipDiffMM, lipDiffPct float64,
	wallThickMM, lipOverhangMM, lipCollarHeightMM float64,
	bottomSlabMM, stdBottomThickMM, bottomDiffMM, bottomDiffPct float64,
	bottomDrainMM, bottomGrossMM float64,
	hasLip, hasRingBottom, ok bool,
) {
	m, valid := computeVaseTrapezePotMetrics(plant)
	if !valid {
		return
	}
	return m.lipThickMM, m.stdLipThickMM, m.lipDiffMM, m.lipDiffPct,
		m.wallThickMM, m.lipOverhangMM, m.lipCollarHeightMM,
		m.bottomSlabMM, m.stdBottomThickMM, m.bottomDiffMM, m.bottomDiffPct,
		m.bottomDrainMM, m.bottomGrossMM,
		m.hasLip, m.hasRingBottom, true
}
