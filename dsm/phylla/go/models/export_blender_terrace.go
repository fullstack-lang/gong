package models

import (
	"fmt"
	"math"
	"strings"
)

type blenderTerraceLayer struct {
	ringIndex  int
	dyRelMM    float64
	thetaRad   float64
	paletteHex string
}

// GenerateVaseTrapezeTerraceBlenderPy generates a self-contained Blender Python (.py) script
// for a VaseTrapeze plant that stacks the upper rings (terracotta pots) directly on a sunlit
// stone terrace, fills each pot with potting earth and Mediterranean plants (lavender, rosemary,
// thyme, and phyllotactic succulents) in its open pockets, and builds a Mediterranean garden
// in the background.
func GenerateVaseTrapezeTerraceBlenderPy(plant *PlantAbstract) string {
	if plant == nil || plant.PlantType != VaseTrapeze || plant.TubeVaseAbstract == nil {
		return ""
	}

	vase := plant.TubeVaseAbstract
	radialRepetitions := max(vase.RadialRepetitions, 1)
	sideLength := plant.RhombusSideLength
	if sideLength <= 0 {
		sideLength = 120.0
	}
	relativeVerticalThickness := vase.RelativeVerticalThickness
	relativeRotatedTorusSeparation := vase.RelativeRotatedTorusSeparation
	rotRatio := vase.RotationRatio
	stackHeight := max(plant.StackHeight, 1)

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
		return ""
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
		return ""
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)
	resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
	resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)

	var trapezeUnrotatedStepDY, trapezeRotatedStepDX, trapezeRotatedStepDY float64
	var hasTrapezeConstantThickness bool
	if len(resampledBaseBottom) == len(targetAngles) && len(resampledBaseTop) == len(targetAngles) {
		baseBottomY := make([]float64, len(targetAngles))
		for i, pt := range resampledBaseBottom {
			baseBottomY[i] = pt.Y
		}
		if newBottomY, newTopY, unrotDY, rotDX, rotDY, ok := ComputeVaseTrapezeConstant3DThickness(plant, globalR, targetAngles, baseBottomY); ok {
			for i := range targetAngles {
				resampledBaseBottom[i].Y = newBottomY[i]
				resampledBaseTop[i].Y = newTopY[i]
			}
			trapezeUnrotatedStepDY = unrotDY
			trapezeRotatedStepDX = rotDX
			trapezeRotatedStepDY = rotDY
			hasTrapezeConstantThickness = true
		}
	}

	p1H := zStart + vase.TopPlaneHeight*sideLength*scaleY
	p2H := zStart + vase.BottomPlaneHeight*sideLength*scaleY
	if p1H <= p2H+1e-6 {
		return ""
	}
	ringHeight := p1H - p2H
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0
	carvedParam := math.Max(0.0, math.Min(1.0, vase.CarvedOutTopRingsParameter))

	nResampled := min(len(resampledBaseBottom), len(resampledBaseTop))
	resampledBotYs := make([]float64, nResampled)
	resampledTopYs := make([]float64, nResampled)
	for i := range nResampled {
		resampledBotYs[i] = resampledBaseBottom[i].Y
		resampledTopYs[i] = resampledBaseTop[i].Y
	}

	deltaYContact := ComputeVaseTrapezeContactStepDYForCurves(plant, hasTrapezeConstantThickness, resampledBotYs, resampledTopYs)
	sWall := VaseTrapezeWallSubdivisions(plant)
	if sWall < 6 {
		sWall = 6
	}
	isCurvedWall := IsVaseTrapezeCurvedWallActive(plant)
	lipHeight, lipWidth, _, hasLip := VaseTrapezeLipDimensionsForMode(plant, hasTrapezeConstantThickness)

	// Compute stack offsets for h = 0 .. stackHeight-1
	dxs := make([]float64, stackHeight)
	dys := make([]float64, stackHeight)
	if stackHeight > 1 {
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
		verticalThickness := relativeVerticalThickness * sideLength
		rotatedSeparation := relativeRotatedTorusSeparation * sideLength

		_, unrotatedStepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, 0.0)
		unrotatedStepDY *= scaleY
		rotatedStepDX := growthVectorX + verticalThickness*vx
		rotatedStepDY := (growthVectorY + verticalThickness*vy + rotatedSeparation) * scaleY
		if hasTrapezeConstantThickness {
			unrotatedStepDY = trapezeUnrotatedStepDY
			rotatedStepDX = trapezeRotatedStepDX
			rotatedStepDY = trapezeRotatedStepDY
		}
		rotatedStepDY = ComputeVaseTrapezeLipRotatedStepDY(plant, rotatedStepDY, ringHeight, hasTrapezeConstantThickness)

		numSteps := stackHeight - 1
		unrotatedStepDYBase := unrotatedStepDY
		unrotatedStepDYRing := unrotatedStepDY
		if rotRatio+1e-9 < float64(numSteps)/float64(stackHeight) && len(resampledBotYs) > 0 {
			unrotatedStepDYBase, unrotatedStepDYRing = ComputeVaseTrapezeUnrotatedStepDYs(
				plant, globalR, resampledBotYs, resampledTopYs, hasTrapezeConstantThickness,
			)
		}
		var cumDX, cumDY float64
		for k := 1; k <= numSteps; k++ {
			stepOrder := numSteps - k + 1
			threshold := float64(stepOrder) / float64(stackHeight)
			if rotRatio+1e-9 >= threshold {
				cumDX += rotatedStepDX
				cumDY += rotatedStepDY
			} else if k == 1 {
				cumDY += unrotatedStepDYBase
			} else {
				cumDY += unrotatedStepDYRing
			}
			dxs[k] = cumDX
			dys[k] = cumDY
		}
	}

	// Build canonical upper ring at (thetaOffset = 0, dy = 0)
	curBottom, curTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
	rawN := min(len(curBottom), len(curTop))
	if rawN < 3 {
		return ""
	}

	projectPoint := func(pt vector3, planeHeight float64) vector3 {
		origTheta := math.Atan2(pt.Z, pt.X)
		r := math.Hypot(pt.X, pt.Z)
		deltaY := planeHeight - pt.Y
		rProj := math.Max(1e-3, r+deltaY*math.Tan(projAngleRad))
		return vector3{
			X: rProj * math.Cos(origTheta),
			Y: planeHeight,
			Z: rProj * math.Sin(origTheta),
		}
	}

	var ringTopP1, ringBottomP1, ringTopP2, ringBottomP2 []vector3
	for i := range rawN {
		tp1 := projectPoint(curTop[i], p1H)
		bp1 := projectPoint(curBottom[i], p1H)
		tp2 := projectPoint(curTop[i], p2H)
		bp2 := projectPoint(curBottom[i], p2H)
		if len(ringTopP1) > 0 {
			prevT := ringTopP1[len(ringTopP1)-1]
			prevB := ringBottomP1[len(ringBottomP1)-1]
			if math.Hypot(tp1.X-prevT.X, tp1.Z-prevT.Z) < 1e-3 && math.Hypot(bp1.X-prevB.X, bp1.Z-prevB.Z) < 1e-3 {
				continue
			}
		}
		ringTopP1 = append(ringTopP1, tp1)
		ringBottomP1 = append(ringBottomP1, bp1)
		ringTopP2 = append(ringTopP2, tp2)
		ringBottomP2 = append(ringBottomP2, bp2)
	}
	if len(ringTopP1) > 2 {
		last := len(ringTopP1) - 1
		if math.Hypot(ringTopP1[last].X-ringTopP1[0].X, ringTopP1[last].Z-ringTopP1[0].Z) < 1e-3 &&
			math.Hypot(ringBottomP1[last].X-ringBottomP1[0].X, ringBottomP1[last].Z-ringBottomP1[0].Z) < 1e-3 {
			ringTopP1 = ringTopP1[:last]
			ringBottomP1 = ringBottomP1[:last]
			ringTopP2 = ringTopP2[:last]
			ringBottomP2 = ringBottomP2[:last]
		}
	}

	M := len(ringTopP1)
	if M < 3 {
		return ""
	}

	var sumRTop, sumRBottom float64
	for i := range M {
		sumRTop += math.Hypot(ringTopP2[i].X, ringTopP2[i].Z)
		sumRBottom += math.Hypot(ringBottomP2[i].X, ringBottomP2[i].Z)
	}

	cOuterUncarvedP1 := ringTopP1
	cInnerUncarvedP1 := ringBottomP1
	cOuterStraightP2 := ringTopP2
	cInnerStraightP2 := ringBottomP2
	if sumRBottom > sumRTop {
		cOuterUncarvedP1 = ringBottomP1
		cInnerUncarvedP1 = ringTopP1
		cOuterStraightP2 = ringBottomP2
		cInnerStraightP2 = ringTopP2
	}

	basePlateHeight := vase.BasePlateHeight * sideLength
	hasRingBottom := vase.HasRingBottom && basePlateHeight > 0 && ringHeight > 0
	hRingBottom := 0.0
	if hasRingBottom {
		hRingBottom = math.Min(basePlateHeight, ringHeight)
	} else {
		// On a terrace holding soil, ensure a sensible interior pot floor thickness
		hasRingBottom = true
		hRingBottom = 0.09 * ringHeight
	}

	ringDeltaR := func(localY float64) float64 {
		return EvalVaseTrapezeRingWallDeltaR(plant, localY, p1H, p2H, deltaYContact)
	}

	offsetCurve := func(src []vector3) []vector3 {
		if !isCurvedWall {
			return src
		}
		out := make([]vector3, len(src))
		for i, pt := range src {
			nx, ny, nz := OffsetRadialPoint(pt.X, pt.Y, pt.Z, ringDeltaR(pt.Y))
			out[i] = vector3{X: nx, Y: ny, Z: nz}
		}
		return out
	}

	interpolatePlancher := func(cP2, cUncarvedP1 []vector3, u float64) []vector3 {
		plancher := make([]vector3, M)
		for i := range M {
			p2 := cP2[i]
			p1 := cUncarvedP1[i]
			plancher[i] = vector3{
				X: (1.0-u)*p2.X + u*p1.X,
				Y: (1.0-u)*p2.Y + u*p1.Y,
				Z: (1.0-u)*p2.Z + u*p1.Z,
			}
		}
		return plancher
	}

	uPlancher := hRingBottom / ringHeight
	cOuterFloorStraight := interpolatePlancher(cOuterStraightP2, cOuterUncarvedP1, uPlancher)
	cInnerFloorStraight := interpolatePlancher(cInnerStraightP2, cInnerUncarvedP1, uPlancher)
	cInnerPlancher := offsetCurve(cInnerFloorStraight)

	computeCarvedOut := func(cTop, cBottom []vector3, param float64) []vector3 {
		rMin := math.MaxFloat64
		rMax := -math.MaxFloat64
		for i := range M {
			r := math.Hypot(cTop[i].X, cTop[i].Z)
			if r < rMin {
				rMin = r
			}
			if r > rMax {
				rMax = r
			}
		}
		rRange := rMax - rMin
		carved := make([]vector3, M)
		for i := range M {
			ptTop := cTop[i]
			ptBottom := cBottom[i]
			r := math.Hypot(ptTop.X, ptTop.Z)
			coeff := 0.0
			if rRange > 1e-6 {
				coeff = param * (rMax - r) / rRange
			}
			sx := (1.0-coeff)*ptTop.X + coeff*ptBottom.X
			sy := (1.0-coeff)*ptTop.Y + coeff*ptBottom.Y
			sz := (1.0-coeff)*ptTop.Z + coeff*ptBottom.Z
			if isCurvedWall {
				sx, sy, sz = OffsetRadialPoint(sx, sy, sz, ringDeltaR(sy))
			}
			carved[i] = vector3{X: sx, Y: sy, Z: sz}
		}
		return carved
	}

	cOuterP1 := computeCarvedOut(cOuterUncarvedP1, cOuterFloorStraight, carvedParam)
	cInnerP1 := computeCarvedOut(cInnerUncarvedP1, cInnerFloorStraight, carvedParam)
	cOuterP2 := offsetCurve(cOuterStraightP2)

	apparatus := computeRingBottomApparatusVec3(
		plant, cOuterStraightP2, cOuterUncarvedP1, cInnerPlancher,
		p2H, ringHeight, hRingBottom, 0.0, isCurvedWall, ringDeltaR,
	)
	cOuterWallBot := cOuterP2
	if apparatus.activeFootRim {
		cOuterWallBot = apparatus.cOuterWallBot
	}

	yFillLocal, grooveDepth, grooveHalfH, hasFillLine := ComputeVaseTrapezeFillLevelLineParams(
		plant, p1H, p2H, hRingBottom, hasTrapezeConstantThickness,
	)
	var grooveRows [][]vector3
	if hasFillLine {
		yContact := p1H - deltaYContact
		_, grooveRows, _ = computeInnerFillGrooveGridVec3(
			cInnerPlancher, cInnerP1, cInnerStraightP2, cInnerUncarvedP1, cOuterStraightP2, cOuterUncarvedP1,
			p2H, ringHeight, 0.0, yContact, yFillLocal, grooveDepth, grooveHalfH, sWall, isCurvedWall, ringDeltaR,
		)
	}

	cWallLipBot := cOuterP1
	cLipBot := cOuterP1
	cLipTop := cOuterP1
	if hasLip {
		cWallLipBot, cLipBot, cLipTop = computeVaseTrapezeLipCurvesVec3(
			cOuterWallBot, cOuterP1, cOuterStraightP2, cOuterUncarvedP1,
			p2H, ringHeight, 0.0, lipHeight, lipWidth, isCurvedWall, ringDeltaR,
		)
	}

	// Sample ringDeltaR(localY) across [p2H, p1H] into a high-precision 1D LUT
	const lutSteps = 256
	deltaRLUT := make([]float64, lutSteps+1)
	for s := 0; s <= lutSteps; s++ {
		ly := p2H + (float64(s)/float64(lutSteps))*ringHeight
		deltaRLUT[s] = ringDeltaR(ly)
	}

	// Collect upper rings (h = 1 .. stackHeight-1, or h = 0 fallback when stackHeight == 1)
	var upperLayers []blenderTerraceLayer
	if stackHeight >= 2 {
		baseDY := dys[1]
		for h := 1; h < stackHeight; h++ {
			upperLayers = append(upperLayers, blenderTerraceLayer{
				ringIndex:  h,
				dyRelMM:    dys[h] - baseDY,
				thetaRad:   dxs[h] / globalR,
				paletteHex: GetVaseRingColor(vase.ColorCombination, h),
			})
		}
	} else {
		upperLayers = append(upperLayers, blenderTerraceLayer{
			ringIndex:  1,
			dyRelMM:    0.0,
			thetaRad:   0.0,
			paletteHex: GetVaseRingColor(vase.ColorCombination, 0),
		})
	}

	// Reference camera azimuth from ThreeJS diagram
	var checkedDiagram *TubeVase3DDiagram
	for _, diagram := range plant.TubeVase3DDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}
	if checkedDiagram == nil && len(plant.TubeVase3DDiagrams) > 0 {
		checkedDiagram = plant.TubeVase3DDiagrams[0]
	}

	hasRefCamera := false
	refCamDX := 1.0
	refCamDZ := 1.0
	if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
		rs := checkedDiagram.Rendered3DShape
		dx := rs.ViewX - rs.TargetX
		dz := rs.ViewZ - rs.TargetZ
		if math.Hypot(dx, dz) > 1e-4 {
			refCamDX = dx
			refCamDZ = dz
			hasRefCamera = true
		}
	}

	slug := sanitizeBlenderSlug(plant.Name)
	var sb strings.Builder
	sb.Grow(128 * 1024)

	sb.WriteString(fmt.Sprintf(`# -*- coding: utf-8 -*-
"""
Blender Mediterranean Terrace Planter Script — Generated by Phylla (Gong DSM)
Plant: %s (N=%d, M=%d, Upper Rings=%d, Angle=%.1f°, Side=%.1f mm)

Scene overview:
  • Upper VaseTrapeze rings stacked directly on a sunlit travertine/limestone terrace
  • Handcrafted porous Mediterranean terra cotta PBR material with subtle kiln variation
  • Rich potting earth & Mediterranean plants (lavender, rosemary, thyme, phyllotactic succulents)
    fitted into each pot's open sunlit pockets + top crown bowl
  • Lush Mediterranean garden background (lavender drifts, olive tree, Italian cypresses, stone wall)

How to use in Blender (4.x / 5.x LTS):
  1. Open Blender -> switch to the 'Scripting' workspace tab.
  2. Click 'Open', select this .py file, and click 'Run Script' (▶).
  3. When AUTO_RENDER_IMAGE = True (default), running the script renders a high-resolution PNG
     next to this .py file (or ~/Downloads) and loads it into Blender's Image Editor.
"""

import math
import os
import random
import bpy
from mathutils import Vector, Matrix

# ==============================================================================
# USER CONFIGURATION
# ==============================================================================
PLANT_NAME = %q
PLANT_SLUG = %q

# Target outer diameter in meters for the widest stacked pot on the terrace
# (set to None to use raw 1 mm = 0.001 m).
PLANTER_TARGET_DIAMETER_M = 0.64
PLANTER_MAX_HEIGHT_M = 1.25

# Material & botanical toggles
USE_TERRACOTTA_MATERIAL = True
ADD_EARTH_AND_PLANTS = True
ADD_GARDEN_BACKGROUND = True
RANDOM_SEED = 42

# Render output settings
AUTO_RENDER_IMAGE = True
RENDER_RESOLUTION_X = 1600
RENDER_RESOLUTION_Y = 1600
CYCLES_SAMPLES = 192

# ==============================================================================
# PHYLLA VASE TRAPEZE UPPER-RING GEOMETRY DATA (millimeters & radians)
# ==============================================================================
GLOBAL_RADIUS_MM = %.6f
P1_H_MM = %.6f
P2_H_MM = %.6f
RING_HEIGHT_MM = %.6f
H_RING_BOTTOM_MM = %.6f
DELTA_Y_CONTACT_MM = %.6f
WALL_SUBDIVISIONS = %d
IS_CURVED_WALL = %s
HAS_LIP = %s
HAS_FILL_LEVEL_LINE = %s
Y_FILL_LOCAL_MM = %.6f
HAS_DRAINAGE_HOLE = %s
HAS_FOOT_RIM = %s
Y_RECESS_MM = %.6f
FLOOR_CENTER_X_MM = %.6f
FLOOR_CENTER_Z_MM = %.6f
HAS_REF_CAMERA = %s
REF_CAM_DX = %.6f
REF_CAM_DZ = %.6f

`,
		plant.Name, plant.N, plant.M, len(upperLayers), plant.RhombusInsideAngle, sideLength,
		plant.Name, slug,
		globalR, p1H, p2H, ringHeight, hRingBottom, deltaYContact, sWall,
		pyBool(isCurvedWall), pyBool(hasLip),
		pyBool(hasFillLine), yFillLocal,
		pyBool(apparatus.activeHole), pyBool(apparatus.activeFootRim),
		apparatus.yRecess, apparatus.centerX, apparatus.centerZ,
		pyBool(hasRefCamera), refCamDX, refCamDZ,
	))

	// Emit DELTA_R_LUT
	sb.WriteString("DELTA_R_LUT = [\n    ")
	for s, dr := range deltaRLUT {
		sb.WriteString(fmt.Sprintf("%.5f,", dr))
		if (s+1)%12 == 0 && s < len(deltaRLUT)-1 {
			sb.WriteString("\n    ")
		} else if s < len(deltaRLUT)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteString("\n]\n\n")

	// Emit UPPER_LAYERS
	sb.WriteString("UPPER_LAYERS = [\n")
	for _, l := range upperLayers {
		sb.WriteString(fmt.Sprintf(
			"    {\"ring_index\": %d, \"dy_rel_mm\": %.5f, \"theta_rad\": %.6f, \"palette_hex\": %q},\n",
			l.ringIndex, l.dyRelMM, l.thetaRad, l.paletteHex,
		))
	}
	sb.WriteString("]\n\n")

	// Emit RING_PROFILE: per azimuth sample i in 0..M-1:
	// (th_out, th_in, r_out_st_p2, r_out_st_p1, r_out_p2, r_out_top, y_out_top,
	//  r_in_st_p2, r_in_st_p1, r_in_plancher, y_in_plancher, r_in_top, y_in_top,
	//  r_wall_lip_bot, y_lip_bot, r_lip_bot, r_lip_top)
	sb.WriteString("RING_PROFILE = [\n")
	for i := range M {
		thOut := math.Atan2(cOuterStraightP2[i].Z, cOuterStraightP2[i].X)
		thIn := math.Atan2(cInnerStraightP2[i].Z, cInnerStraightP2[i].X)
		rOutStP2 := math.Hypot(cOuterStraightP2[i].X, cOuterStraightP2[i].Z)
		rOutStP1 := math.Hypot(cOuterUncarvedP1[i].X, cOuterUncarvedP1[i].Z)
		rOutP2 := math.Hypot(cOuterP2[i].X, cOuterP2[i].Z)
		rOutTop := math.Hypot(cOuterP1[i].X, cOuterP1[i].Z)
		yOutTop := cOuterP1[i].Y

		rInStP2 := math.Hypot(cInnerStraightP2[i].X, cInnerStraightP2[i].Z)
		rInStP1 := math.Hypot(cInnerUncarvedP1[i].X, cInnerUncarvedP1[i].Z)
		rInPl := math.Hypot(cInnerPlancher[i].X, cInnerPlancher[i].Z)
		yInPl := cInnerPlancher[i].Y
		rInTop := math.Hypot(cInnerP1[i].X, cInnerP1[i].Z)
		yInTop := cInnerP1[i].Y

		rWallLipBot := math.Hypot(cWallLipBot[i].X, cWallLipBot[i].Z)
		yLipBot := cWallLipBot[i].Y
		rLipBot := math.Hypot(cLipBot[i].X, cLipBot[i].Z)
		rLipTop := math.Hypot(cLipTop[i].X, cLipTop[i].Z)

		sb.WriteString(fmt.Sprintf(
			"    (%.6f, %.6f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f),\n",
			thOut, thIn,
			rOutStP2, rOutStP1, rOutP2, rOutTop, yOutTop,
			rInStP2, rInStP1, rInPl, yInPl, rInTop, yInTop,
			rWallLipBot, yLipBot, rLipBot, rLipTop,
		))
	}
	sb.WriteString("]\n\n")

	// Emit FILL_GROOVE_PROFILE: (y_lo, y_mid, r_mid, y_hi) per azimuth when HAS_FILL_LEVEL_LINE
	sb.WriteString("FILL_GROOVE_PROFILE = [\n")
	if hasFillLine && len(grooveRows) == 3 {
		for i := range M {
			rMid := math.Hypot(grooveRows[1][i].X, grooveRows[1][i].Z)
			sb.WriteString(fmt.Sprintf(
				"    (%.3f, %.3f, %.3f, %.3f),\n",
				grooveRows[0][i].Y, grooveRows[1][i].Y, rMid, grooveRows[2][i].Y,
			))
		}
	}
	sb.WriteString("]\n\n")

	// Emit FOOT_RIM_PROFILE: (ow_x, ow_y, ow_z, fo_x, fo_y, fo_z, fi_x, fi_y, fi_z, ro_x, ro_y, ro_z) when HAS_FOOT_RIM
	sb.WriteString("FOOT_RIM_PROFILE = [\n")
	if apparatus.activeFootRim {
		for i := range M {
			ow := apparatus.cOuterWallBot[i]
			fo := apparatus.cFootOut[i]
			fi := apparatus.cFootIn[i]
			ro := apparatus.cRecessOut[i]
			sb.WriteString(fmt.Sprintf(
				"    (%.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f),\n",
				ow.X, ow.Y, ow.Z, fo.X, fo.Y, fo.Z, fi.X, fi.Y, fi.Z, ro.X, ro.Y, ro.Z,
			))
		}
	}
	sb.WriteString("]\n\n")

	// Emit DRAINAGE_HOLE_PROFILE: (hc_x, hc_y, hc_z, ht_x, ht_y, ht_z, hb_x, hb_y, hb_z) when HAS_DRAINAGE_HOLE
	sb.WriteString("DRAINAGE_HOLE_PROFILE = [\n")
	if apparatus.activeHole {
		for i := range M {
			hc := apparatus.cHoleChamfer[i]
			ht := apparatus.cHoleTop[i]
			hb := apparatus.cHoleBot[i]
			sb.WriteString(fmt.Sprintf(
				"    (%.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f, %.3f),\n",
				hc.X, hc.Y, hc.Z, ht.X, ht.Y, ht.Z, hb.X, hb.Y, hb.Z,
			))
		}
	}
	sb.WriteString("]\n\n")

	sb.WriteString(blenderTerracePythonEngine)
	return sb.String()
}

const blenderTerracePythonEngine = `# ==============================================================================
# BLENDER MEDITERRANEAN TERRACE & BOTANICAL SCENE ENGINE
# ==============================================================================

GOLDEN_ANGLE = math.radians(137.507764)


def ensure_node_tree(id_block):
    if getattr(id_block, "node_tree", None) is None and hasattr(id_block, "use_nodes"):
        id_block.use_nodes = True
    return id_block.node_tree


def srgb_to_linear(c):
    if c <= 0.04045:
        return c / 12.92
    return ((c + 0.055) / 1.055) ** 2.4


def hex_to_linear_rgba(hex_str, alpha=1.0):
    s = hex_str.strip().lstrip("#")
    if len(s) != 6:
        return (0.48, 0.14, 0.07, alpha)
    r = srgb_to_linear(int(s[0:2], 16) / 255.0)
    g = srgb_to_linear(int(s[2:4], 16) / 255.0)
    b = srgb_to_linear(int(s[4:6], 16) / 255.0)
    return (r, g, b, alpha)


def set_socket(node, names, value):
    if isinstance(names, str):
        names = (names,)
    for name in names:
        if name in node.inputs:
            try:
                node.inputs[name].default_value = value
                return True
            except Exception:
                continue
    return False


def clear_all_datablocks():
    fresh_scene = bpy.data.scenes.new("__phylla_terrace_init__")
    if bpy.context.window is not None:
        bpy.context.window.scene = fresh_scene
    for sc in list(bpy.data.scenes):
        if sc != fresh_scene:
            bpy.data.scenes.remove(sc, do_unlink=True)

    for block_collection in (
        bpy.data.objects,
        bpy.data.meshes,
        bpy.data.curves,
        bpy.data.materials,
        bpy.data.lights,
        bpy.data.cameras,
        bpy.data.collections,
        bpy.data.worlds,
    ):
        for item in list(block_collection):
            block_collection.remove(item, do_unlink=True)

    return fresh_scene


def make_collection(name, parent=None):
    col = bpy.data.collections.new(name)
    if parent is not None:
        parent.children.link(col)
    return col


def add_mesh_object(name, verts, faces, collection, material=None, smooth=True):
    if not verts or not faces:
        return None
    mesh = bpy.data.meshes.new(name + "_Mesh")
    mesh.from_pydata(verts, [], faces)
    mesh.update(calc_edges=True)
    if smooth:
        if hasattr(mesh, "shade_smooth"):
            mesh.shade_smooth()
        else:
            for poly in mesh.polygons:
                poly.use_smooth = True
    else:
        if hasattr(mesh, "shade_flat"):
            mesh.shade_flat()

    obj = bpy.data.objects.new(name, mesh)
    collection.objects.link(obj)
    if material is not None:
        obj.data.materials.append(material)
    return obj


class MeshBuilder:
    """Fast batch accumulator for botanical & architectural geometry sharing one material."""

    def __init__(self):
        self.verts = []
        self.faces = []

    def add_verts_faces(self, new_verts, new_faces):
        base = len(self.verts)
        self.verts.extend(new_verts)
        for f in new_faces:
            self.faces.append(tuple(base + idx for idx in f))

    def add_tube(self, points, radii, segments=6, cap_end=True):
        n = len(points)
        if n < 2:
            return
        base = len(self.verts)
        prev_normal = None
        for idx in range(n):
            p = points[idx]
            if idx == 0:
                tang = (points[1] - points[0]).normalized()
            elif idx == n - 1:
                tang = (points[-1] - points[-2]).normalized()
            else:
                tang = (points[idx + 1] - points[idx - 1]).normalized()
            if tang.length < 1e-6:
                tang = Vector((0.0, 0.0, 1.0))

            if prev_normal is None:
                ref = Vector((0.0, 0.0, 1.0)) if abs(tang.z) < 0.92 else Vector((1.0, 0.0, 0.0))
                nrm = tang.cross(ref).normalized()
            else:
                nrm = (prev_normal - tang * prev_normal.dot(tang))
                if nrm.length < 1e-5:
                    ref = Vector((1.0, 0.0, 0.0))
                    nrm = tang.cross(ref)
                nrm.normalize()
            prev_normal = nrm
            binorm = tang.cross(nrm).normalized()
            r = radii[idx]
            for s in range(segments):
                ang = 2.0 * math.pi * s / segments
                offset = (nrm * math.cos(ang) + binorm * math.sin(ang)) * r
                v = p + offset
                self.verts.append((v.x, v.y, v.z))

        for idx in range(n - 1):
            r0 = base + idx * segments
            r1 = base + (idx + 1) * segments
            for s in range(segments):
                sn = (s + 1) % segments
                self.faces.append((r0 + s, r0 + sn, r1 + sn, r1 + s))

        if cap_end and segments >= 3:
            tip_ring = base + (n - 1) * segments
            tip_pt = points[-1]
            tip_idx = len(self.verts)
            self.verts.append((tip_pt.x, tip_pt.y, tip_pt.z))
            for s in range(segments):
                sn = (s + 1) % segments
                self.faces.append((tip_ring + s, tip_ring + sn, tip_idx))

    def add_leaf_blade(self, origin, direction, up_hint, length, width, arch=0.18, thickness=0.0012):
        d = direction.normalized()
        side = d.cross(up_hint)
        if side.length < 1e-5:
            side = d.cross(Vector((1.0, 0.0, 0.0)))
        side.normalize()
        up = side.cross(d).normalized()

        # 3D sculpted leaf blade with midrib keel and arched tip (top & bottom surfaces -> closed/solid feel)
        p0 = origin
        p1 = origin + d * (0.34 * length) + up * (0.35 * arch * length)
        p2 = origin + d * (0.72 * length) + up * (0.15 * arch * length)
        p3 = origin + d * length - up * (0.38 * arch * length)

        w1 = 0.50 * width
        w2 = 0.36 * width
        th = 0.5 * thickness

        v = [
            (p0 + up * th),
            (p1 + side * w1 + up * th * 0.4),
            (p1 - side * w1 + up * th * 0.4),
            (p2 + side * w2 + up * th * 0.3),
            (p2 - side * w2 + up * th * 0.3),
            p3,
            (p1 - up * th * 1.1),
            (p2 - up * th * 0.8),
        ]
        verts_t = [(pt.x, pt.y, pt.z) for pt in v]
        faces_t = [
            # Upper blade faces
            (0, 1, 3),
            (0, 3, 4),
            (0, 4, 2),
            (3, 5, 4),
            # Keeled underside faces
            (0, 6, 1),
            (0, 2, 6),
            (1, 6, 7),
            (1, 7, 3),
            (2, 4, 7),
            (2, 7, 6),
            (3, 7, 5),
            (4, 5, 7),
        ]
        self.add_verts_faces(verts_t, faces_t)

    def add_succulent_petal(self, origin, radial_dir, tilt_rad, length, width, thickness):
        horiz = Vector((radial_dir.x, radial_dir.y, 0.0)).normalized()
        up_z = Vector((0.0, 0.0, 1.0))
        axis = (horiz * math.sin(tilt_rad) + up_z * math.cos(tilt_rad)).normalized()
        side = axis.cross(up_z)
        if side.length < 1e-5:
            side = Vector((1.0, 0.0, 0.0))
        side.normalize()
        nrm = side.cross(axis).normalized()  # faces inward/upward toward rosette core

        # Fleshy spoon-shaped succulent leaf with upturned mucronate tip
        s0 = origin
        s1 = origin + axis * (0.32 * length) - nrm * (0.04 * length)
        s2 = origin + axis * (0.72 * length) - nrm * (0.02 * length)
        s3 = origin + axis * length + nrm * (0.10 * length)

        hw0 = 0.22 * width
        hw1 = 0.48 * width
        hw2 = 0.52 * width

        pts = [
            # 0, 1: base left/right
            s0 - side * hw0 + nrm * (0.25 * thickness),
            s0 + side * hw0 + nrm * (0.25 * thickness),
            # 2, 3, 4: section 1 left, top-center-cup, right, bottom-keel (5)
            s1 - side * hw1 + nrm * (0.35 * thickness),
            s1 + nrm * (0.22 * thickness),
            s1 + side * hw1 + nrm * (0.35 * thickness),
            s1 - nrm * (0.68 * thickness),
            # 6, 7, 8, 9: section 2 left, top-center, right, bottom-keel
            s2 - side * hw2 + nrm * (0.32 * thickness),
            s2 + nrm * (0.24 * thickness),
            s2 + side * hw2 + nrm * (0.32 * thickness),
            s2 - nrm * (0.58 * thickness),
            # 10: pointed tip
            s3,
        ]
        vt = [(p.x, p.y, p.z) for p in pts]
        fc = [
            # Upper concave face
            (0, 3, 2), (0, 1, 3), (1, 4, 3),
            (2, 3, 7), (2, 7, 6), (3, 4, 8), (3, 8, 7),
            (6, 7, 10), (7, 8, 10),
            # Lower plump keel
            (0, 2, 5), (1, 5, 4), (0, 5, 1),
            (2, 6, 9), (2, 9, 5), (4, 5, 9), (4, 9, 8),
            (6, 10, 9), (8, 9, 10),
        ]
        self.add_verts_faces(vt, fc)

    def add_ellipsoid(self, center, rx, ry, rz, rot_matrix=None, seg_u=8, seg_v=5):
        base = len(self.verts)
        cx, cy, cz = center.x, center.y, center.z
        # Bottom pole & top pole
        bot = Vector((0.0, 0.0, -rz))
        top = Vector((0.0, 0.0, rz))
        if rot_matrix is not None:
            bot = rot_matrix @ bot
            top = rot_matrix @ top
        self.verts.append((cx + bot.x, cy + bot.y, cz + bot.z))
        self.verts.append((cx + top.x, cy + top.y, cz + top.z))

        for iv in range(1, seg_v):
            phi = -0.5 * math.pi + math.pi * iv / seg_v
            cp = math.cos(phi)
            sp = math.sin(phi)
            for iu in range(seg_u):
                th = 2.0 * math.pi * iu / seg_u
                lp = Vector((rx * cp * math.cos(th), ry * cp * math.sin(th), rz * sp))
                if rot_matrix is not None:
                    lp = rot_matrix @ lp
                self.verts.append((cx + lp.x, cy + lp.y, cz + lp.z))

        # Bottom fan
        first_ring = base + 2
        for iu in range(seg_u):
            inext = (iu + 1) % seg_u
            self.faces.append((base, first_ring + inext, first_ring + iu))

        # Middle bands
        for iv in range(seg_v - 2):
            r0 = first_ring + iv * seg_u
            r1 = r0 + seg_u
            for iu in range(seg_u):
                inext = (iu + 1) % seg_u
                self.faces.append((r0 + iu, r0 + inext, r1 + inext, r1 + iu))

        # Top fan
        last_ring = first_ring + (seg_v - 2) * seg_u
        for iu in range(seg_u):
            inext = (iu + 1) % seg_u
            self.faces.append((base + 1, last_ring + iu, last_ring + inext))

    def add_foliage_spray(self, center, radius, n_leaves, leaf_len, leaf_wid, rng, upward_bias=0.35):
        """Builds an airy, organic cluster of 3D botanical leaf cards (no smooth spheres)."""
        cx, cy, cz = center.x, center.y, center.z
        for k in range(n_leaves):
            u = (k + 0.5) / max(1, n_leaves)
            az = k * GOLDEN_ANGLE + rng.uniform(-0.25, 0.25)
            el = math.asin(max(-0.85, min(0.95, 2.0 * u - 1.0))) * 0.65 + upward_bias
            dir_v = Vector((
                math.cos(el) * math.cos(az),
                math.cos(el) * math.sin(az),
                math.sin(el),
            )).normalized()
            rr = radius * (0.35 + 0.65 * math.sqrt(u)) * rng.uniform(0.75, 1.18)
            p0 = Vector((cx + dir_v.x * rr * 0.55, cy + dir_v.y * rr * 0.55, cz + dir_v.z * rr * 0.50))

            # Orient leaf outward & slightly twisted for rich specular/translucency catchlights
            up_h = Vector((rng.uniform(-0.3, 0.3), rng.uniform(-0.3, 0.3), 1.0)).normalized()
            side = dir_v.cross(up_h)
            if side.length < 1e-5:
                side = Vector((1.0, 0.0, 0.0))
            side.normalize()
            l_cur = leaf_len * rng.uniform(0.80, 1.22)
            w_cur = leaf_wid * rng.uniform(0.82, 1.18)
            mid = p0 + dir_v * (0.48 * l_cur) + Vector((0.0, 0.0, 0.08 * l_cur))
            tip = p0 + dir_v * l_cur - Vector((0.0, 0.0, 0.10 * l_cur))
            v_l = mid + side * (0.5 * w_cur)
            v_r = mid - side * (0.5 * w_cur)
            base = len(self.verts)
            self.verts.extend([
                (p0.x, p0.y, p0.z),
                (v_l.x, v_l.y, v_l.z),
                (tip.x, tip.y, tip.z),
                (v_r.x, v_r.y, v_r.z),
            ])
            self.faces.append((base, base + 1, base + 2))
            self.faces.append((base, base + 2, base + 3))

    def build(self, name, collection, material, smooth=True):
        return add_mesh_object(name, self.verts, self.faces, collection, material=material, smooth=smooth)


# ------------------------------------------------------------------------------
# MATERIALS (TERRA COTTA, POTTING EARTH, BOTANICALS, TRAVERTINE, LIMEWASH)
# ------------------------------------------------------------------------------

def create_terracotta_material(pot_index, fallback_hex="#b84a39"):
    mat = bpy.data.materials.new(f"TerraCotta_Pot_{pot_index:02d}")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (620, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (300, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-860, 0)

    if not USE_TERRACOTTA_MATERIAL:
        rgba = hex_to_linear_rgba(fallback_hex, 1.0)
        set_socket(bsdf, "Base Color", rgba)
        set_socket(bsdf, "Roughness", 0.42)
        return mat

    # Subtle kiln-firing warmth shift from pot to pot so stacked pots feel artisanal
    kiln_shifts = (-0.035, 0.025, -0.015, 0.040, -0.025, 0.015)
    k_shift = kiln_shifts[pot_index % len(kiln_shifts)]

    def warm_clay(r_lin, g_lin, b_lin):
        f = 1.0 + k_shift
        return (
            min(1.0, max(0.02, r_lin * f)),
            min(1.0, max(0.01, g_lin * (1.0 + 0.85 * k_shift))),
            min(1.0, max(0.01, b_lin * (1.0 + 0.70 * k_shift))),
            1.0,
        )

    # Macro clay mottling (fired sienna, warm ochre-terracotta, subtle pale mineral bloom)
    noise_macro = nodes.new("ShaderNodeTexNoise")
    noise_macro.location = (-580, 220)
    set_socket(noise_macro, "Scale", 9.5 + 1.3 * (pot_index % 4))
    set_socket(noise_macro, "Detail", 6.0)
    set_socket(noise_macro, "Roughness", 0.62)
    links.new(tex_coord.outputs["Object"], noise_macro.inputs["Vector"])

    # Fine mineral clay speckles & limestone efflorescence grain
    noise_grain = nodes.new("ShaderNodeTexNoise")
    noise_grain.location = (-580, -20)
    set_socket(noise_grain, "Scale", 145.0)
    set_socket(noise_grain, "Detail", 8.0)
    set_socket(noise_grain, "Roughness", 0.70)
    links.new(tex_coord.outputs["Object"], noise_grain.inputs["Vector"])

    ramp_clay = nodes.new("ShaderNodeValToRGB")
    ramp_clay.location = (-240, 220)
    elems = ramp_clay.color_ramp.elements
    elems[0].position = 0.20
    elems[0].color = warm_clay(0.235, 0.046, 0.019)   # deep fired Tuscan terracotta sienna
    elems[1].position = 0.78
    elems[1].color = warm_clay(0.395, 0.088, 0.036)   # rich sunlit Provençal terra cotta
    e_patina = ramp_clay.color_ramp.elements.new(0.95)
    e_patina.color = warm_clay(0.490, 0.155, 0.075)   # subtle dusty mineral calcaire highlight
    links.new(noise_macro.outputs["Fac"], ramp_clay.inputs["Fac"])
    links.new(ramp_clay.outputs["Color"], bsdf.inputs["Base Color"])

    # Unglazed porous clay roughness (0.68 .. 0.84)
    ramp_rough = nodes.new("ShaderNodeValToRGB")
    ramp_rough.location = (-240, -20)
    ramp_rough.color_ramp.elements[0].position = 0.20
    ramp_rough.color_ramp.elements[0].color = (0.67, 0.67, 0.67, 1.0)
    ramp_rough.color_ramp.elements[1].position = 0.85
    ramp_rough.color_ramp.elements[1].color = (0.84, 0.84, 0.84, 1.0)
    links.new(noise_grain.outputs["Fac"], ramp_rough.inputs["Fac"])
    links.new(ramp_rough.outputs["Color"], bsdf.inputs["Roughness"])

    set_socket(bsdf, "Metallic", 0.0)
    set_socket(bsdf, ("Specular IOR Level", "Specular"), 0.36)
    set_socket(bsdf, ("Subsurface Weight", "Subsurface"), 0.03)
    set_socket(bsdf, "Subsurface Scale", 0.006)
    set_socket(bsdf, "Subsurface Radius", (1.0, 0.38, 0.20))

    # Subtle throwing/molding striations + sand-grain micro-porosity bump
    mapping_stri = nodes.new("ShaderNodeMapping")
    mapping_stri.location = (-660, -260)
    set_socket(mapping_stri, "Scale", (1.0, 1.0, 42.0))
    links.new(tex_coord.outputs["Object"], mapping_stri.inputs["Vector"])

    noise_stri = nodes.new("ShaderNodeTexNoise")
    noise_stri.location = (-460, -260)
    set_socket(noise_stri, "Scale", 6.5)
    set_socket(noise_stri, "Detail", 3.0)
    links.new(mapping_stri.outputs["Vector"], noise_stri.inputs["Vector"])

    bump_stri = nodes.new("ShaderNodeBump")
    bump_stri.location = (-200, -260)
    set_socket(bump_stri, "Strength", 0.035)
    set_socket(bump_stri, "Distance", 0.0012)
    links.new(noise_stri.outputs["Fac"], bump_stri.inputs["Height"])

    bump_pore = nodes.new("ShaderNodeBump")
    bump_pore.location = (30, -220)
    set_socket(bump_pore, "Strength", 0.065)
    set_socket(bump_pore, "Distance", 0.0015)
    links.new(noise_grain.outputs["Fac"], bump_pore.inputs["Height"])
    links.new(bump_stri.outputs["Normal"], bump_pore.inputs["Normal"])
    links.new(bump_pore.outputs["Normal"], bsdf.inputs["Normal"])

    return mat


def create_earth_material():
    mat = bpy.data.materials.new("Mediterranean_Potting_Earth")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (520, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (220, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-720, 0)

    noise_clod = nodes.new("ShaderNodeTexNoise")
    noise_clod.location = (-460, 140)
    set_socket(noise_clod, "Scale", 65.0)
    set_socket(noise_clod, "Detail", 10.0)
    set_socket(noise_clod, "Roughness", 0.78)
    links.new(tex_coord.outputs["Object"], noise_clod.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-160, 140)
    ramp.color_ramp.elements[0].position = 0.25
    ramp.color_ramp.elements[0].color = (0.024, 0.015, 0.010, 1.0)  # moist rich dark humus
    ramp.color_ramp.elements[1].position = 0.80
    ramp.color_ramp.elements[1].color = (0.068, 0.042, 0.026, 1.0)  # warm sun-dried loam
    fleck = ramp.color_ramp.elements.new(0.955)
    fleck.color = (0.18, 0.13, 0.09, 1.0)                           # tiny mineral/bark flecks
    links.new(noise_clod.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", 0.94)
    set_socket(bsdf, ("Specular IOR Level", "Specular"), 0.18)

    bump = nodes.new("ShaderNodeBump")
    bump.location = (-60, -140)
    set_socket(bump, "Strength", 0.38)
    set_socket(bump, "Distance", 0.006)
    links.new(noise_clod.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])

    return mat


def create_pebble_material():
    mat = bpy.data.materials.new("Soil_Mineral_Mulch")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    noise = nodes.new("ShaderNodeTexNoise")
    set_socket(noise, "Scale", 38.0)
    set_socket(noise, "Detail", 5.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.color_ramp.elements[0].position = 0.25
    ramp.color_ramp.elements[0].color = (0.16, 0.09, 0.06, 1.0)  # crushed terracotta / bark
    ramp.color_ramp.elements[1].position = 0.78
    ramp.color_ramp.elements[1].color = (0.42, 0.36, 0.28, 1.0)  # warm garrigue limestone gravel
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])
    set_socket(bsdf, "Roughness", 0.86)
    return mat


def create_foliage_material(name, col_dark, col_light, roughness=0.48, translucency=0.20, noise_scale=18.0):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (560, 0)

    mix_sh = nodes.new("ShaderNodeMixShader")
    mix_sh.location = (340, 0)
    mix_sh.inputs[0].default_value = translucency
    links.new(mix_sh.outputs["Shader"], out.inputs["Surface"])

    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (40, 100)
    set_socket(bsdf, "Roughness", roughness)
    set_socket(bsdf, ("Specular IOR Level", "Specular"), 0.35)
    set_socket(bsdf, ("Sheen Weight", "Sheen"), 0.12)

    transl = nodes.new("ShaderNodeBsdfTranslucent")
    transl.location = (40, -140)

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-580, 0)
    noise = nodes.new("ShaderNodeTexNoise")
    noise.location = (-380, 0)
    set_socket(noise, "Scale", noise_scale)
    set_socket(noise, "Detail", 4.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-180, 80)
    ramp.color_ramp.elements[0].position = 0.22
    ramp.color_ramp.elements[0].color = (*col_dark, 1.0)
    ramp.color_ramp.elements[1].position = 0.82
    ramp.color_ramp.elements[1].color = (*col_light, 1.0)
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    # Warm sunlit chlorophyll transmission color
    tr_col = (
        min(1.0, col_light[0] * 1.25 + 0.03),
        min(1.0, col_light[1] * 1.30 + 0.04),
        min(1.0, col_light[2] * 0.85),
        1.0,
    )
    set_socket(transl, "Color", tr_col)

    links.new(bsdf.outputs["BSDF"], mix_sh.inputs[1])
    links.new(transl.outputs["BSDF"], mix_sh.inputs[2])
    return mat


def create_succulent_material():
    mat = bpy.data.materials.new("Mediterranean_Succulent_Rosette")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (480, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (160, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-560, 0)
    noise = nodes.new("ShaderNodeTexNoise")
    noise.location = (-360, 0)
    set_socket(noise, "Scale", 26.0)
    set_socket(noise, "Detail", 3.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-140, 60)
    ramp.color_ramp.elements[0].position = 0.20
    ramp.color_ramp.elements[0].color = (0.095, 0.215, 0.175, 1.0)  # glaucous sage-jade
    ramp.color_ramp.elements[1].position = 0.78
    ramp.color_ramp.elements[1].color = (0.195, 0.320, 0.265, 1.0)  # powdery pruinose blue-green
    blush = ramp.color_ramp.elements.new(0.94)
    blush.color = (0.34, 0.12, 0.10, 1.0)                           # sun-kissed terracotta rose tip
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", 0.56)
    set_socket(bsdf, ("Subsurface Weight", "Subsurface"), 0.16)
    set_socket(bsdf, "Subsurface Scale", 0.008)
    set_socket(bsdf, "Subsurface Radius", (0.6, 1.0, 0.5))
    set_socket(bsdf, ("Sheen Weight", "Sheen"), 0.22)
    return mat


def create_bark_material(name, col_a=(0.11, 0.075, 0.052), col_b=(0.22, 0.175, 0.135)):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    noise = nodes.new("ShaderNodeTexNoise")
    set_socket(noise, "Scale", 32.0)
    set_socket(noise, "Detail", 7.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.color_ramp.elements[0].position = 0.25
    ramp.color_ramp.elements[0].color = (*col_a, 1.0)
    ramp.color_ramp.elements[1].position = 0.80
    ramp.color_ramp.elements[1].color = (*col_b, 1.0)
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])
    set_socket(bsdf, "Roughness", 0.88)

    bump = nodes.new("ShaderNodeBump")
    set_socket(bump, "Strength", 0.18)
    set_socket(bump, "Distance", 0.004)
    links.new(noise.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])
    return mat


def create_travertine_stone_material():
    mat = bpy.data.materials.new("Terrace_Travertine_Limestone")
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    out.location = (560, 0)
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    bsdf.location = (260, 0)
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    tex_coord.location = (-780, 0)
    obj_info = nodes.new("ShaderNodeObjectInfo")
    obj_info.location = (-780, 220)

    noise_vein = nodes.new("ShaderNodeTexNoise")
    noise_vein.location = (-520, 40)
    set_socket(noise_vein, "Scale", 3.4)
    set_socket(noise_vein, "Detail", 6.0)
    set_socket(noise_vein, "Distortion", 0.65)
    links.new(tex_coord.outputs["Object"], noise_vein.inputs["Vector"])

    # Combine per-paver random ID with natural travertine veining
    math_mix = nodes.new("ShaderNodeMath")
    math_mix.location = (-300, 120)
    math_mix.operation = "MULTIPLY_ADD"
    math_mix.inputs[1].default_value = 0.35
    links.new(obj_info.outputs["Random"], math_mix.inputs[0])
    links.new(noise_vein.outputs["Fac"], math_mix.inputs[2])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.location = (-80, 120)
    ramp.color_ramp.elements[0].position = 0.22
    ramp.color_ramp.elements[0].color = (0.31, 0.24, 0.17, 1.0)  # warm honey travertine
    ramp.color_ramp.elements[1].position = 0.94
    ramp.color_ramp.elements[1].color = (0.47, 0.39, 0.29, 1.0)  # sunlit Mediterranean limestone
    links.new(math_mix.outputs["Value"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])

    set_socket(bsdf, "Roughness", 0.76)
    set_socket(bsdf, ("Specular IOR Level", "Specular"), 0.32)

    noise_pit = nodes.new("ShaderNodeTexNoise")
    noise_pit.location = (-520, -180)
    set_socket(noise_pit, "Scale", 68.0)
    set_socket(noise_pit, "Detail", 6.0)
    links.new(tex_coord.outputs["Object"], noise_pit.inputs["Vector"])

    bump = nodes.new("ShaderNodeBump")
    bump.location = (-80, -160)
    set_socket(bump, "Strength", 0.08)
    set_socket(bump, "Distance", 0.002)
    links.new(noise_pit.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])
    return mat


def create_limewash_wall_material(name, c_low=(0.56, 0.47, 0.35), c_high=(0.72, 0.63, 0.49)):
    mat = bpy.data.materials.new(name)
    nt = ensure_node_tree(mat)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputMaterial")
    bsdf = nodes.new("ShaderNodeBsdfPrincipled")
    links.new(bsdf.outputs["BSDF"], out.inputs["Surface"])

    tex_coord = nodes.new("ShaderNodeTexCoord")
    noise = nodes.new("ShaderNodeTexNoise")
    set_socket(noise, "Scale", 2.8)
    set_socket(noise, "Detail", 6.0)
    links.new(tex_coord.outputs["Object"], noise.inputs["Vector"])

    ramp = nodes.new("ShaderNodeValToRGB")
    ramp.color_ramp.elements[0].position = 0.22
    ramp.color_ramp.elements[0].color = (*c_low, 1.0)
    ramp.color_ramp.elements[1].position = 0.80
    ramp.color_ramp.elements[1].color = (*c_high, 1.0)
    links.new(noise.outputs["Fac"], ramp.inputs["Fac"])
    links.new(ramp.outputs["Color"], bsdf.inputs["Base Color"])
    set_socket(bsdf, "Roughness", 0.88)

    noise_fine = nodes.new("ShaderNodeTexNoise")
    set_socket(noise_fine, "Scale", 42.0)
    set_socket(noise_fine, "Detail", 5.0)
    links.new(tex_coord.outputs["Object"], noise_fine.inputs["Vector"])

    bump = nodes.new("ShaderNodeBump")
    set_socket(bump, "Strength", 0.06)
    set_socket(bump, "Distance", 0.002)
    links.new(noise_fine.outputs["Fac"], bump.inputs["Height"])
    links.new(bump.outputs["Normal"], bsdf.inputs["Normal"])
    return mat


# ------------------------------------------------------------------------------
# UPPER RINGS STACKED TERRACOTTA POTS BUILDER + POLAR CLEARANCE PROFILES
# ------------------------------------------------------------------------------

def eval_ring_delta_r_mm(local_y_mm):
    if not IS_CURVED_WALL or not DELTA_R_LUT:
        return 0.0
    t = (local_y_mm - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
    t = max(0.0, min(1.0, t))
    pos = t * (len(DELTA_R_LUT) - 1)
    idx = int(math.floor(pos))
    if idx >= len(DELTA_R_LUT) - 1:
        return DELTA_R_LUT[-1]
    frac = pos - idx
    return DELTA_R_LUT[idx] * (1.0 - frac) + DELTA_R_LUT[idx + 1] * frac


def ring_wall_sample_y_mm(s, s_wall, y_bot, y_top, y_contact):
    if s <= 0:
        return y_bot
    if s >= s_wall:
        return y_top
    if s_wall < 3 or y_top <= y_bot + 1e-6:
        return y_bot + (float(s) / float(s_wall)) * (y_top - y_bot)
    s_split = s_wall // 3
    span = y_top - y_bot
    min_bot = y_bot + 0.25 * span * (float(s_split) / float(s_wall))
    max_top = y_top - 0.25 * span * (float(s_wall - s_split) / float(s_wall))
    if y_contact <= min_bot or y_contact >= max_top:
        return y_bot + (float(s) / float(s_wall)) * span
    if s <= s_split:
        return y_bot + (float(s) / float(s_split)) * (y_contact - y_bot)
    return y_contact + (float(s - s_split) / float(s_wall - s_split)) * (y_top - y_contact)


def build_stacked_pots(col_pots, cam_xy):
    # Compute overall bounding box in mm of the stacked upper rings
    max_r_mm = 10.0
    max_y_rel_mm = RING_HEIGHT_MM
    for row in RING_PROFILE:
        r_cand = max(row[4], row[5], row[15], row[16])
        if r_cand > max_r_mm:
            max_r_mm = r_cand

    for layer in UPPER_LAYERS:
        dy_rel = layer["dy_rel_mm"]
        for row in RING_PROFILE:
            y_top_rel = (row[6] - P2_H_MM) + dy_rel
            if y_top_rel > max_y_rel_mm:
                max_y_rel_mm = y_top_rel

    raw_diam_mm = max(1.0, 2.0 * max_r_mm)
    if PLANTER_TARGET_DIAMETER_M is not None and PLANTER_TARGET_DIAMETER_M > 0:
        scale = PLANTER_TARGET_DIAMETER_M / raw_diam_mm
    else:
        scale = 0.001

    if max_y_rel_mm * scale > PLANTER_MAX_HEIGHT_M:
        scale = PLANTER_MAX_HEIGHT_M / max_y_rel_mm

    # Align ThreeJS reference camera azimuth directly toward our Blender terrace camera
    cam_azimuth = math.atan2(cam_xy[1], cam_xy[0])
    if HAS_REF_CAMERA:
        src_azimuth = math.atan2(-REF_CAM_DZ, REF_CAM_DX)
        delta_alpha = cam_azimuth - src_azimuth
    else:
        delta_alpha = 0.0

    m_pts = len(RING_PROFILE)
    s_wall = max(6, WALL_SUBDIVISIONS)
    y_contact_local = P1_H_MM - DELTA_Y_CONTACT_MM

    pot_infos = []

    for pot_idx, layer in enumerate(UPPER_LAYERS):
        dy_rel_mm = layer["dy_rel_mm"]
        theta_off = layer["theta_rad"]
        z_bot_m = dy_rel_mm * scale
        mat = create_terracotta_material(pot_idx, layer["palette_hex"])

        # Build smooth profile cross-sections at every azimuth i in 0..m_pts-1
        # Ordered continuously from outer bottom -> outer wall -> (lip) -> rounded top rim -> inner wall -> inner plancher
        profiles_xyz = []
        floor_outer_ring = []
        plancher_inner_ring = []
        outer_wall_bot_ring = []
        foot_out_ring = []
        foot_in_ring = []
        recess_out_ring = []
        hole_chamfer_ring = []
        hole_top_ring = []
        hole_bot_ring = []

        use_fill_groove = HAS_FILL_LEVEL_LINE and len(FILL_GROOVE_PROFILE) == m_pts
        use_foot_rim = HAS_FOOT_RIM and len(FOOT_RIM_PROFILE) == m_pts
        use_drainage_hole = HAS_DRAINAGE_HOLE and len(DRAINAGE_HOLE_PROFILE) == m_pts

        def go_pt_to_blender(x_mm, y_local_mm, z_mm):
            r = math.hypot(x_mm, z_mm)
            th = math.atan2(z_mm, x_mm)
            phi = -(th + theta_off) + delta_alpha
            return (
                r * math.cos(phi) * scale,
                r * math.sin(phi) * scale,
                (y_local_mm - P2_H_MM + dy_rel_mm) * scale,
            )

        # 360-bin polar lookup tables in Blender world coordinates (1-degree bins)
        bins_in_r_rim = [0.0] * 360
        bins_in_z_rim = [0.0] * 360
        bins_in_r_plancher = [0.0] * 360
        bins_out_r_bot = [0.0] * 360
        bins_out_r_top = [0.0] * 360
        bins_out_z_top = [0.0] * 360
        bins_weight = [0.0] * 360

        for i in range(m_pts):
            (
                th_out, th_in,
                r_out_st_p2, r_out_st_p1, r_out_p2, r_out_top, y_out_top,
                r_in_st_p2, r_in_st_p1, r_in_pl, y_in_pl, r_in_top, y_in_top,
                r_wall_lip_bot, y_lip_bot, r_lip_bot, r_lip_top,
            ) = RING_PROFILE[i]

            phi_out = -(th_out + theta_off) + delta_alpha
            phi_in = -(th_in + theta_off) + delta_alpha
            co_o, si_o = math.cos(phi_out), math.sin(phi_out)
            co_i, si_i = math.cos(phi_in), math.sin(phi_in)

            def pt_out(r_mm, y_local_mm):
                rm = r_mm * scale
                zm = (y_local_mm - P2_H_MM + dy_rel_mm) * scale
                return (rm * co_o, rm * si_o, zm)

            def pt_in(r_mm, y_local_mm):
                rm = r_mm * scale
                zm = (y_local_mm - P2_H_MM + dy_rel_mm) * scale
                return (rm * co_i, rm * si_i, zm)

            def pt_blend(u_in, r_mm, y_local_mm):
                # Interpolate unit direction smoothly across top rim between outer and inner rays
                dphi = (phi_in - phi_out + math.pi) % (2.0 * math.pi) - math.pi
                ph = phi_out + u_in * dphi
                rm = r_mm * scale
                zm = (y_local_mm - P2_H_MM + dy_rel_mm) * scale
                return (rm * math.cos(ph), rm * math.sin(ph), zm)

            sec = []

            # 1. Outer wall from bottom (or raised ceramic foot bevel top) up to lip underside (or outer top)
            y_wall_end = y_lip_bot if HAS_LIP else y_out_top
            r_wall_end = r_wall_lip_bot if HAS_LIP else r_out_top
            if use_foot_rim:
                ow_x, ow_y, ow_z, fo_x, fo_y, fo_z, fi_x, fi_y, fi_z, ro_x, ro_y, ro_z = FOOT_RIM_PROFILE[i]
                pt_ow = go_pt_to_blender(ow_x, ow_y, ow_z)
                outer_wall_bot_ring.append(pt_ow)
                foot_out_ring.append(go_pt_to_blender(fo_x, fo_y, fo_z))
                foot_in_ring.append(go_pt_to_blender(fi_x, fi_y, fi_z))
                recess_out_ring.append(go_pt_to_blender(ro_x, ro_y, ro_z))
                y_wall_start = ow_y
                for s in range(s_wall + 1):
                    if s == 0:
                        sec.append(pt_ow)
                        continue
                    elif s == s_wall:
                        yl = y_wall_end
                        rl = r_wall_end
                    else:
                        yl = ring_wall_sample_y_mm(s, s_wall, y_wall_start, y_wall_end, y_contact_local)
                        u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                        rl = (1.0 - u) * r_out_st_p2 + u * r_out_st_p1 + eval_ring_delta_r_mm(yl)
                    sec.append(pt_out(rl, yl))
            else:
                foot_bev_mm = min(1.6, 0.08 * max(1.0, y_wall_end - P2_H_MM))

                sec.append(pt_out(max(1.0, r_out_p2 - 0.8 * foot_bev_mm), P2_H_MM))
                for s in range(s_wall + 1):
                    if s == 0:
                        yl = P2_H_MM + foot_bev_mm
                    elif s == s_wall:
                        yl = y_wall_end
                    else:
                        yl = ring_wall_sample_y_mm(s, s_wall, P2_H_MM + foot_bev_mm, y_wall_end, y_contact_local)
                    if s == s_wall:
                        rl = r_wall_end
                    else:
                        u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                        rl = (1.0 - u) * r_out_st_p2 + u * r_out_st_p1 + eval_ring_delta_r_mm(yl)
                    sec.append(pt_out(rl, yl))

            # 2. Load-bearing outer stacking lip + sponge-rounded top outer & inner rim edges
            if HAS_LIP:
                lip_w_bot = max(0.0, r_lip_bot - r_wall_lip_bot)
                lip_w_top = max(0.0, r_lip_top - r_out_top)
                lip_h = max(0.0, y_out_top - y_lip_bot)
                b_lip = min(1.8, 0.22 * lip_w_bot, 0.22 * lip_h)
                s_lip = s_wall if IS_CURVED_WALL else 1
                if b_lip > 0.15:
                    sec.append(pt_out(r_wall_lip_bot + 0.35 * lip_w_bot, y_lip_bot))
                    sec.append(pt_out(r_lip_bot - b_lip, y_lip_bot))
                    y_c0 = y_lip_bot + b_lip
                    y_c1 = y_out_top - b_lip
                    for s in range(s_lip + 1):
                        tc = s / max(1, s_lip)
                        yl = (1.0 - tc) * y_c0 + tc * y_c1
                        if s == 0 and s_lip == 1:
                            rl = r_lip_bot
                        elif s == s_lip and s_lip == 1:
                            rl = r_lip_top
                        else:
                            u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                            toff = (yl - y_lip_bot) / max(1e-6, lip_h)
                            off = (1.0 - toff) * lip_w_bot + toff * lip_w_top
                            rl = (1.0 - u) * r_out_st_p2 + u * r_out_st_p1 + eval_ring_delta_r_mm(yl) + off
                        sec.append(pt_out(rl, yl))
                    sec.append(pt_out(r_lip_top - 0.35 * b_lip, y_out_top - 0.25 * b_lip))
                    sec.append(pt_out(r_lip_top - b_lip, y_out_top))
                else:
                    for s in range(s_lip + 1):
                        if s == 0:
                            sec.append(pt_out(r_lip_bot, y_lip_bot))
                        elif s == s_lip:
                            sec.append(pt_out(r_lip_top, y_out_top))
                        else:
                            tc = s / s_lip
                            yl = (1.0 - tc) * y_lip_bot + tc * y_out_top
                            u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                            off = (1.0 - tc) * lip_w_bot + tc * lip_w_top
                            rl = (1.0 - u) * r_out_st_p2 + u * r_out_st_p1 + eval_ring_delta_r_mm(yl) + off
                            sec.append(pt_out(rl, yl))
                r_rim_outer = r_lip_top - b_lip
            else:
                rim_span = max(0.5, r_out_top - r_in_top)
                b_out = min(1.8, 0.20 * rim_span)
                sec[-1] = pt_out(r_out_top, y_out_top - b_out)
                sec.append(pt_out(r_out_top - 0.30 * b_out, y_out_top - 0.25 * b_out))
                sec.append(pt_out(r_out_top - b_out, y_out_top))
                r_rim_outer = r_out_top - b_out

            # 3. Across top rim to inner rounded edge
            rim_span_tot = max(0.5, r_rim_outer - r_in_top)
            b_in = min(1.8, 0.22 * rim_span_tot, 0.12 * max(1.0, y_in_top - y_in_pl))
            sec.append(pt_blend(0.5, 0.5 * (r_rim_outer + r_in_top + b_in), 0.5 * (y_out_top + y_in_top)))
            sec.append(pt_in(r_in_top + b_in, y_in_top))
            sec.append(pt_in(r_in_top + 0.28 * b_in, y_in_top - 0.28 * b_in))

            # 4. Inner wall down to plancher (with optional recessed fill-level V-groove)
            if use_fill_groove:
                y_g_lo, y_g_mid, r_g_mid, y_g_hi = FILL_GROOVE_PROFILE[i]
                y_in_start = y_in_top - b_in
                y_g_hi_c = max(y_in_pl + 0.5, min(y_in_start - 0.5, y_g_hi))
                y_g_lo_c = max(y_in_pl + 0.25, min(y_g_hi_c - 0.25, y_g_lo))
                y_g_mid_c = 0.5 * (y_g_lo_c + y_g_hi_c)
                if y_g_lo_c < y_g_mid < y_g_hi_c:
                    y_g_mid_c = y_g_mid
                s_up = max(2, s_wall // 2)
                s_lo = max(2, s_wall - s_up)
                for s in range(s_up, -1, -1):
                    yl = y_g_hi_c + (s / s_up) * (y_in_start - y_g_hi_c)
                    u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                    rl = (1.0 - u) * r_in_st_p2 + u * r_in_st_p1 + eval_ring_delta_r_mm(yl)
                    sec.append(pt_in(rl, yl))
                sec.append(pt_in(r_g_mid, y_g_mid_c))
                for s in range(s_lo, -1, -1):
                    yl = y_in_pl + (s / s_lo) * (y_g_lo_c - y_in_pl)
                    if s == 0:
                        rl = r_in_pl
                    else:
                        u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                        rl = (1.0 - u) * r_in_st_p2 + u * r_in_st_p1 + eval_ring_delta_r_mm(yl)
                    sec.append(pt_in(rl, yl))
            else:
                for s in range(s_wall, -1, -1):
                    if s == s_wall:
                        yl = y_in_top - b_in
                    elif s == 0:
                        yl = y_in_pl
                    else:
                        yl = ring_wall_sample_y_mm(s, s_wall, y_in_pl, y_in_top - b_in, y_contact_local)
                    if s == 0:
                        rl = r_in_pl
                    else:
                        u = (yl - P2_H_MM) / max(1e-6, RING_HEIGHT_MM)
                        rl = (1.0 - u) * r_in_st_p2 + u * r_in_st_p1 + eval_ring_delta_r_mm(yl)
                    sec.append(pt_in(rl, yl))

            if use_drainage_hole:
                hc_x, hc_y, hc_z, ht_x, ht_y, ht_z, hb_x, hb_y, hb_z = DRAINAGE_HOLE_PROFILE[i]
                hole_chamfer_ring.append(go_pt_to_blender(hc_x, hc_y, hc_z))
                hole_top_ring.append(go_pt_to_blender(ht_x, ht_y, ht_z))
                hole_bot_ring.append(go_pt_to_blender(hb_x, hb_y, hb_z))

            profiles_xyz.append(sec)
            floor_outer_ring.append(sec[0])
            plancher_inner_ring.append(sec[-1])

            # Deposit into 360-degree world azimuth lookup table
            deg_f = (math.degrees(phi_in) % 360.0)
            b0 = int(math.floor(deg_f)) % 360
            b1 = (b0 + 1) % 360
            fr = deg_f - math.floor(deg_f)
            r_max_out_m = (r_lip_top if HAS_LIP else r_out_top) * scale
            for b_idx, w in ((b0, 1.0 - fr), (b1, fr)):
                bins_in_r_rim[b_idx] += w * (r_in_top * scale)
                bins_in_z_rim[b_idx] += w * ((y_in_top - P2_H_MM + dy_rel_mm) * scale)
                bins_in_r_plancher[b_idx] += w * (r_in_pl * scale)
                bins_out_r_bot[b_idx] += w * (r_out_p2 * scale)
                bins_out_r_top[b_idx] += w * r_max_out_m
                bins_out_z_top[b_idx] += w * ((y_out_top - P2_H_MM + dy_rel_mm) * scale)
                bins_weight[b_idx] += w

        # Fill any sparse azimuth bin by neighbor interpolation
        for _ in range(4):
            for b in range(360):
                if bins_weight[b] > 1e-5:
                    continue
                bp = (b - 1) % 360
                bn = (b + 1) % 360
                wp, wn = bins_weight[bp], bins_weight[bn]
                wt = wp + wn
                if wt > 1e-5:
                    bins_in_r_rim[b] = (bins_in_r_rim[bp] + bins_in_r_rim[bn])
                    bins_in_z_rim[b] = (bins_in_z_rim[bp] + bins_in_z_rim[bn])
                    bins_in_r_plancher[b] = (bins_in_r_plancher[bp] + bins_in_r_plancher[bn])
                    bins_out_r_bot[b] = (bins_out_r_bot[bp] + bins_out_r_bot[bn])
                    bins_out_r_top[b] = (bins_out_r_top[bp] + bins_out_r_top[bn])
                    bins_out_z_top[b] = (bins_out_z_top[bp] + bins_out_z_top[bn])
                    bins_weight[b] = wt

        for b in range(360):
            inv = 1.0 / max(1e-6, bins_weight[b])
            bins_in_r_rim[b] *= inv
            bins_in_z_rim[b] *= inv
            bins_in_r_plancher[b] *= inv
            bins_out_r_bot[b] *= inv
            bins_out_r_top[b] *= inv
            bins_out_z_top[b] *= inv

        p_len = len(profiles_xyz[0])
        verts = []
        faces = []
        for i in range(m_pts):
            verts.extend(profiles_xyz[i])

        # Smooth quad-grid sweep for outer wall, rounded lip, top rim, and inner wall
        for i in range(m_pts):
            ni = (i + 1) % m_pts
            b0 = i * p_len
            b1 = ni * p_len
            for s in range(p_len - 1):
                faces.append((b0 + s, b0 + s + 1, b1 + s + 1, b1 + s))

        def add_quad_strip(ring_a, ring_b):
            base_a = len(verts)
            verts.extend(ring_a)
            base_b = len(verts)
            verts.extend(ring_b)
            for idx in range(m_pts):
                nidx = (idx + 1) % m_pts
                faces.append((base_a + idx, base_b + idx, base_b + nidx, base_a + nidx))

        # Foot rim bevels & 3-notched standing ceramic foot ring
        if use_foot_rim:
            add_quad_strip(foot_out_ring, outer_wall_bot_ring)
            add_quad_strip(foot_in_ring, foot_out_ring)
            add_quad_strip(recess_out_ring, foot_in_ring)

        underside_outer_ring = recess_out_ring if use_foot_rim else floor_outer_ring
        cz_pl = sum(p[2] for p in plancher_inner_ring) / m_pts

        if use_drainage_hole:
            # Gently sloped (2°) interior plancher, 45° hole chamfer, cylindrical bore, and recessed underside
            add_quad_strip(plancher_inner_ring, hole_chamfer_ring)
            add_quad_strip(hole_chamfer_ring, hole_top_ring)
            add_quad_strip(hole_top_ring, hole_bot_ring)
            add_quad_strip(hole_bot_ring, underside_outer_ring)
        else:
            # Dedicated flat-normal rings for interior plancher disk (+Z) & exterior bottom foot disk (-Z)
            pl_base = len(verts)
            verts.extend(plancher_inner_ring)
            cx_pl = sum(p[0] for p in plancher_inner_ring) / m_pts
            cy_pl = sum(p[1] for p in plancher_inner_ring) / m_pts
            pl_center = len(verts)
            verts.append((cx_pl, cy_pl, cz_pl))

            fl_base = len(verts)
            verts.extend(underside_outer_ring)
            if use_foot_rim:
                fl_center_pt = go_pt_to_blender(FLOOR_CENTER_X_MM, Y_RECESS_MM, FLOOR_CENTER_Z_MM)
            else:
                fl_center_pt = (
                    sum(p[0] for p in underside_outer_ring) / m_pts,
                    sum(p[1] for p in underside_outer_ring) / m_pts,
                    sum(p[2] for p in underside_outer_ring) / m_pts,
                )
            fl_center = len(verts)
            verts.append(fl_center_pt)

            for i in range(m_pts):
                ni = (i + 1) % m_pts
                # Plancher top disk (+Z)
                faces.append((pl_center, pl_base + ni, pl_base + i))
                # Exterior bottom disk (-Z)
                faces.append((fl_center, fl_base + i, fl_base + ni))

        add_mesh_object(f"TerraCotta_Pot_{pot_idx + 1:02d}", verts, faces, col_pots, material=mat, smooth=True)

        pot_infos.append({
            "index": pot_idx,
            "z_bot": z_bot_m,
            "z_plancher": cz_pl,
            "z_fill": (Y_FILL_LOCAL_MM - P2_H_MM + dy_rel_mm) * scale if use_fill_groove else None,
            "in_r_rim": bins_in_r_rim,
            "in_z_rim": bins_in_z_rim,
            "in_r_plancher": bins_in_r_plancher,
            "out_r_bot": bins_out_r_bot,
            "out_r_top": bins_out_r_top,
            "out_z_top": bins_out_z_top,
        })

    return scale, pot_infos


# ------------------------------------------------------------------------------
# POTTING EARTH & MEDITERRANEAN PLANTS IN EACH POT
# ------------------------------------------------------------------------------

def lookup_polar(arr360, phi_rad):
    deg = math.degrees(phi_rad) % 360.0
    b0 = int(math.floor(deg)) % 360
    b1 = (b0 + 1) % 360
    fr = deg - math.floor(deg)
    return arr360[b0] * (1.0 - fr) + arr360[b1] * fr


def pot_outer_radius_at(pot_info, phi_rad, z_m):
    z0 = pot_info["z_bot"]
    z1 = lookup_polar(pot_info["out_z_top"], phi_rad)
    if z_m < z0 - 0.01 or z_m > z1 + 0.02:
        return 0.0
    u = max(0.0, min(1.0, (z_m - z0) / max(1e-4, z1 - z0)))
    r0 = lookup_polar(pot_info["out_r_bot"], phi_rad)
    r1 = lookup_polar(pot_info["out_r_top"], phi_rad)
    return (1.0 - u) * r0 + u * r1


def pot_inner_radius_at(pot_info, phi_rad, z_m):
    z0 = pot_info["z_plancher"]
    z1 = lookup_polar(pot_info["in_z_rim"], phi_rad)
    u = max(0.0, min(1.0, (z_m - z0) / max(1e-4, z1 - z0)))
    r0 = lookup_polar(pot_info["in_r_plancher"], phi_rad)
    r1 = lookup_polar(pot_info["in_r_rim"], phi_rad)
    return (1.0 - u) * r0 + u * r1


def soil_noise(x, y, seed_off=0.0):
    return (
        0.0045 * math.sin(29.0 * x + 23.0 * y + seed_off)
        + 0.0030 * math.cos(53.0 * x - 41.0 * y + 1.7 * seed_off)
        + 0.0018 * math.sin(97.0 * x + 83.0 * y - 0.9 * seed_off)
    )


def find_pot_planting_stations(pot_infos, pot_idx):
    """Finds sunlit open crescent pockets in lower pots and crown stations in the top pot."""
    p = pot_infos[pot_idx]
    num_pots = len(pot_infos)
    z_pl = p["z_plancher"]
    max_z_rim = max(p["in_z_rim"])
    min_z_rim = min(p["in_z_rim"])
    pot_depth = max(0.03, max_z_rim - z_pl)

    # Top pot: completely open skyward bowl
    if pot_idx == num_pots - 1:
        z_soil_center = z_pl + 0.84 * (min_z_rim - z_pl)
        cam_phi = math.atan2(-2.15, -0.38)
        back_dir_phi = cam_phi + math.pi
        r_hero = 0.16 * lookup_polar(p["in_r_rim"], back_dir_phi)
        stations = [
            {"phi": back_dir_phi, "r": r_hero, "z": z_soil_center + 0.012, "kind": "lavender_hero", "scale": 1.12, "lean": 0.14},
        ]
        crown_specs = (
            (cam_phi - 0.58, "rosemary", 0.98, 0.42),
            (cam_phi + 0.48, "succulent_cluster", 1.02, 0.28),
            (cam_phi - 0.06, "thyme", 0.96, 0.48),
            (cam_phi - 1.55, "rosemary", 0.92, 0.38),
            (cam_phi + 1.52, "thyme", 0.92, 0.42),
        )
        for phi_raw, ckind, cscale, clean in crown_specs:
            phi = phi_raw % (2.0 * math.pi)
            r_rim = lookup_polar(p["in_r_rim"], phi)
            z_rim = lookup_polar(p["in_z_rim"], phi)
            r_st = (0.52 if ckind == "succulent_cluster" else 0.60) * r_rim
            z_st = min(z_soil_center + 0.008, z_rim - 0.020)
            stations.append({
                "phi": phi,
                "r": r_st,
                "z": z_st,
                "kind": ckind,
                "scale": cscale,
                "lean": clean,
                "r_rim": r_rim,
            })
        return stations, z_soil_center

    # Lower / middle pot: pot_above sits stacked inside it, leaving crescent lobe pockets
    p_above = pot_infos[pot_idx + 1]
    z_soil_ref = z_pl + 0.84 * max(0.02, min_z_rim - z_pl)
    if p.get("z_fill") is not None:
        z_soil_ref = max(z_pl + 0.015, min(p["z_fill"], min_z_rim - 0.018))

    gaps = [0.0] * 360
    for b in range(360):
        phi = math.radians(b)
        z_eval = max(z_soil_ref, p["in_z_rim"][b] - 0.018)
        r_in = pot_inner_radius_at(p, phi, z_eval)
        r_ab = pot_outer_radius_at(p_above, phi, z_eval)
        gaps[b] = r_in - r_ab

    # Smooth circularly over 9 degrees
    sm_gaps = [0.0] * 360
    for b in range(360):
        acc = 0.0
        for d in range(-4, 5):
            acc += gaps[(b + d) % 360]
        sm_gaps[b] = acc / 9.0

    max_gap = max(sm_gaps)
    if max_gap <= 0.014:
        return [], z_soil_ref

    thresh = max(0.018, 0.24 * max_gap)
    open_mask = [g >= thresh for g in sm_gaps]

    # Extract contiguous angular pocket segments on the circle
    segments = []
    if all(open_mask):
        segments.append((0, 360))
    else:
        start_search = 0
        for b in range(360):
            if not open_mask[b]:
                start_search = b
                break
        in_seg = False
        seg_s = 0
        for step in range(361):
            b = (start_search + step) % 360
            if step < 360 and open_mask[b]:
                if not in_seg:
                    in_seg = True
                    seg_s = step
            else:
                if in_seg:
                    seg_len = step - seg_s
                    if seg_len >= 12:
                        segments.append(((start_search + seg_s) % 360, seg_len))
                    in_seg = False

    stations = []
    species_cycle = (
        ("lavender", "rosemary", "succulent", "thyme", "rosemary"),
        ("rosemary", "lavender", "thyme", "succulent", "lavender"),
        ("lavender", "thyme", "rosemary", "succulent", "rosemary"),
    )[pot_idx % 3]

    for seg_start, seg_len in segments:
        # Peak gap inside this pocket arc
        best_off = max(range(seg_len), key=lambda off: sm_gaps[(seg_start + off) % 360])
        peak_deg = (seg_start + best_off) % 360

        # Choose 1, 3, or 5 stations depending on angular width of the crescent pocket
        if seg_len >= 85:
            offsets_deg = [0, -23, 23, -43, 43]
        elif seg_len >= 50:
            offsets_deg = [0, -22, 22]
        else:
            offsets_deg = [0]

        for st_i, d_off in enumerate(offsets_deg):
            cand_deg = (peak_deg + d_off) % 360
            g_here = sm_gaps[cand_deg]
            if g_here < 0.020:
                continue
            phi = math.radians(cand_deg)
            z_local_soil = min(z_soil_ref + 0.008, p["in_z_rim"][cand_deg] - 0.018)
            r_in = pot_inner_radius_at(p, phi, z_local_soil)
            r_ab = pot_outer_radius_at(p_above, phi, max(z_local_soil, p_above["z_bot"] + 0.015))
            r_mid = r_ab + 0.58 * (r_in - r_ab)
            # Ensure comfortable margin inside rim and outside pot above
            r_mid = min(r_in - 0.016, max(r_ab + 0.014, r_mid))

            sp_kind = species_cycle[st_i % len(species_cycle)]
            st_scale = max(0.62, min(1.08, g_here / 0.095)) * (0.95 if pot_depth > 0.14 else 0.80)
            stations.append({
                "phi": phi,
                "r": r_mid,
                "z": z_local_soil,
                "kind": sp_kind,
                "scale": st_scale,
                "lean": 0.46 if sp_kind in ("lavender", "rosemary") else 0.54,
                "gap": g_here,
                "r_rim": r_in,
            })

    return stations, z_soil_ref


def push_point_outside_upper_pots(pt, pot_infos, from_pot_idx, clearance=0.010, keep_inside_own_pot=True):
    x, y, z = pt.x, pt.y, pt.z
    phi = math.atan2(y, x)
    r = math.hypot(x, y)
    if keep_inside_own_pot:
        own_pot = pot_infos[from_pot_idx]
        z_rim_top = max(lookup_polar(own_pot["in_z_rim"], phi), lookup_polar(own_pot["out_z_top"], phi))
        if z < z_rim_top + 0.008:
            r_in_lim = max(0.015, pot_inner_radius_at(own_pot, phi, z) - 0.009)
            r = min(r, r_in_lim)
    for upper_idx in range(from_pot_idx + 1, len(pot_infos)):
        r_req = pot_outer_radius_at(pot_infos[upper_idx], phi, z)
        if r_req > 1e-4 and r < r_req + clearance:
            r = r_req + clearance
    return Vector((r * math.cos(phi), r * math.sin(phi), z))


def safe_leaf_dir(orig, leaf_dir, length, pot_infos, pot_idx):
    """Ensures a leaf rooted below the pot rim rises above the rim before crossing the inner wall."""
    own_pot = pot_infos[pot_idx]
    tip = orig + leaf_dir * length
    phi_t = math.atan2(tip.y, tip.x)
    r_t = math.hypot(tip.x, tip.y)
    z_rim_top = max(lookup_polar(own_pot["in_z_rim"], phi_t), lookup_polar(own_pot["out_z_top"], phi_t))
    if orig.z < z_rim_top + 0.006 and tip.z < z_rim_top + 0.010:
        r_in_lim = pot_inner_radius_at(own_pot, phi_t, tip.z) - 0.006
        if r_t > r_in_lim:
            # Steer leaf upward so it clears the terracotta rim cleanly
            return (leaf_dir + Vector((0.0, 0.0, 0.95))).normalized()
    return leaf_dir


def build_soil_mesh_for_pot(pot_infos, pot_idx, z_soil_base, stations, col_earth, earth_mat, pebble_builder, rng):
    p = pot_infos[pot_idx]
    has_above = (pot_idx + 1) < len(pot_infos)
    p_above = pot_infos[pot_idx + 1] if has_above else None

    n_ang = 180
    u_rings = (0.14, 0.34, 0.56, 0.76, 0.90, 1.00)
    n_r = len(u_rings)

    verts = []
    faces = []

    # Center vertex (safely below pot_above's floor if an upper pot sits inside this pot)
    z_center = z_soil_base + 0.010
    if has_above:
        z_center = min(z_center, p_above["z_bot"] - 0.004)
    verts.append((0.0, 0.0, z_center))

    seed_off = 3.7 * pot_idx + 1.1
    for ia in range(n_ang):
        phi = 2.0 * math.pi * ia / n_ang
        co, si = math.cos(phi), math.sin(phi)
        z_rim_here = lookup_polar(p["in_z_rim"], phi)
        z_edge = min(z_soil_base, z_rim_here - 0.018)
        z_edge = max(p["z_plancher"] + 0.010, z_edge)
        r_edge = max(0.015, pot_inner_radius_at(p, phi, z_edge) - 0.0015)

        r_ab_floor = lookup_polar(p_above["in_r_plancher"], phi) if has_above else 0.0

        for u in u_rings:
            r = u * r_edge
            x = r * co
            y = r * si
            # Subtle mounding around plant root crowns + natural clod roughness near surface
            mound = 0.008 * math.sin(math.pi * min(1.0, (1.0 - u) * 1.35))
            for st in stations:
                st_x = st["r"] * math.cos(st["phi"])
                st_y = st["r"] * math.sin(st["phi"])
                d2 = (x - st_x) ** 2 + (y - st_y) ** 2
                mound += 0.009 * math.exp(-d2 / (0.045 ** 2))

            edge_atten = 0.25 if u > 0.96 else 1.0
            z = z_edge + mound + edge_atten * soil_noise(x, y, seed_off)

            # Never let lower pot soil poke into the bottom or interior bowl of the pot stacked above
            if has_above and r < r_ab_floor + 0.004:
                z = min(z, p_above["z_bot"] - 0.004)

            verts.append((x, y, z))

    # Top soil surface faces (+Z)
    for ia in range(n_ang):
        na = (ia + 1) % n_ang
        c0 = 1 + ia * n_r
        c1 = 1 + na * n_r
        faces.append((0, c0, c1))
        for ir in range(n_r - 1):
            faces.append((c0 + ir, c1 + ir, c1 + ir + 1, c0 + ir + 1))

    # Skirt & bottom cap tucked inside pot so soil is a solid 3D earth plug
    bot_base = len(verts)
    z_plug_bot = p["z_plancher"] + 0.002
    for ia in range(n_ang):
        phi = 2.0 * math.pi * ia / n_ang
        r_pl = max(0.012, lookup_polar(p["in_r_plancher"], phi) - 0.002)
        verts.append((r_pl * math.cos(phi), r_pl * math.sin(phi), z_plug_bot))
    bot_center = len(verts)
    verts.append((0.0, 0.0, z_plug_bot))

    for ia in range(n_ang):
        na = (ia + 1) % n_ang
        top_outer_0 = 1 + ia * n_r + (n_r - 1)
        top_outer_1 = 1 + na * n_r + (n_r - 1)
        faces.append((top_outer_0, top_outer_1, bot_base + na, bot_base + ia))
        faces.append((bot_center, bot_base + na, bot_base + ia))

    add_mesh_object(f"Pot_{pot_idx + 1:02d}_Earth", verts, faces, col_earth, material=earth_mat, smooth=True)

    # Scatter natural Mediterranean gravel & bark chips on exposed pocket soil
    for st in stations:
        for _ in range(14):
            ang = st["phi"] + rng.uniform(-0.32, 0.32)
            rr = st["r"] + rng.uniform(-0.032, 0.032)
            z_rim_h = lookup_polar(p["in_z_rim"], ang)
            z_s = min(z_soil_base, z_rim_h - 0.018)
            r_max = pot_inner_radius_at(p, ang, z_s) - 0.008
            r_min = pot_outer_radius_at(p_above, ang, max(z_s, p_above["z_bot"] + 0.015)) + 0.008 if has_above else 0.01
            if r_max <= r_min:
                continue
            rr = max(r_min, min(r_max, rr))
            px = rr * math.cos(ang)
            py = rr * math.sin(ang)
            pz = z_s + 0.006 + soil_noise(px, py, seed_off)
            rot = Matrix.Rotation(rng.uniform(0.0, math.tau), 3, "Z") @ Matrix.Rotation(rng.uniform(-0.4, 0.4), 3, "X")
            pebble_builder.add_ellipsoid(
                Vector((px, py, pz)),
                rng.uniform(0.0035, 0.0075),
                rng.uniform(0.0028, 0.0060),
                rng.uniform(0.0018, 0.0036),
                rot_matrix=rot,
                seg_u=6,
                seg_v=4,
            )


def add_lavender_cluster(root_pt, out_dir, lean, scale, pot_infos, pot_idx, builders, rng, is_hero=False):
    b_leaf = builders["lav_leaf"]
    b_stalk = builders["lav_stalk"]
    b_calyx = builders["lav_calyx"]
    b_petal = builders["lav_petal"]
    up = Vector((0.0, 0.0, 1.0))

    # 1. Silvery-sage basal foliage mound (phyllotactic golden-angle tuft)
    n_leaves = int((76 if is_hero else 54) * max(0.75, scale))
    for k in range(n_leaves):
        u = (k + 0.5) / n_leaves
        az = k * GOLDEN_ANGLE + rng.uniform(-0.08, 0.08)
        local_rad = Vector((math.cos(az), math.sin(az), 0.0))
        # Bias leaves gently outward toward sun/terrace
        dir_horiz = (local_rad + out_dir * (0.45 * lean)).normalized()
        elev = 0.25 + 1.05 * (1.0 - u)
        leaf_dir = (dir_horiz * math.cos(elev) + up * math.sin(elev)).normalized()
        l_len = (0.052 + 0.045 * math.sin(math.pi * u)) * scale * rng.uniform(0.85, 1.15)
        l_wid = 0.0052 * scale
        orig = root_pt + local_rad * (0.012 * u * scale) + Vector((0.0, 0.0, 0.004))
        orig = push_point_outside_upper_pots(orig, pot_infos, pot_idx, clearance=0.006)
        leaf_dir = safe_leaf_dir(orig, leaf_dir, l_len, pot_infos, pot_idx)
        b_leaf.add_leaf_blade(orig, leaf_dir, up, l_len, l_wid * 1.25, arch=0.26, thickness=0.0010)

    # 2. Slender arching flower stalks with tiered violet-purple flower spikes
    n_stalks = int((15 if is_hero else 9) * max(0.75, scale))
    for s_idx in range(n_stalks):
        az = s_idx * GOLDEN_ANGLE + rng.uniform(-0.18, 0.18)
        fan_vec = Vector((math.cos(az), math.sin(az), 0.0))
        spread = 0.28 + 0.35 * ((s_idx % 5) / 4.0)
        horiz = (out_dir * lean + fan_vec * spread)
        if horiz.length > 1e-4:
            horiz_n = horiz.normalized()
            tilt = min(0.82, horiz.length)
        else:
            horiz_n = out_dir
            tilt = 0.18

        stalk_len = rng.uniform(0.16, 0.245 if is_hero else 0.215) * scale
        pts = []
        radii = []
        n_seg = 7
        sway = Vector((-horiz_n.y, horiz_n.x, 0.0)) * rng.uniform(-0.012, 0.012)
        for j in range(n_seg + 1):
            t = j / n_seg
            # Graceful phototropic curve: leans outward near base to clear pot above, then rises toward sun
            r_reach = stalk_len * tilt * (0.65 * t + 0.35 * (t ** 1.35))
            z_reach = stalk_len * math.sqrt(max(0.15, 1.0 - 0.55 * tilt * tilt)) * (t ** 0.92)
            pt = root_pt + horiz_n * r_reach + up * z_reach + sway * math.sin(math.pi * t)
            pt = push_point_outside_upper_pots(pt, pot_infos, pot_idx, clearance=0.014)
            pts.append(pt)
            radii.append((0.0018 * (1.0 - 0.45 * t)) * max(0.8, scale))

        b_stalk.add_tube(pts, radii, segments=5, cap_end=True)

        # Tiered lavender flower whorls along top 32% of stalk
        spike_axis = (pts[-1] - pts[-3]).normalized()
        side_a = spike_axis.cross(up)
        if side_a.length < 1e-5:
            side_a = Vector((1.0, 0.0, 0.0))
        side_a.normalize()
        side_b = spike_axis.cross(side_a).normalized()

        n_whorls = 7
        for w_idx in range(n_whorls):
            uw = w_idx / (n_whorls - 1)
            # Interpolate along top two segments of stalk
            seg_f = (n_seg - 2.2) + uw * 2.15
            i0 = max(0, min(n_seg - 1, int(math.floor(seg_f))))
            frac = seg_f - i0
            center_w = pts[i0].lerp(pts[i0 + 1], frac)
            whorl_scale = (1.0 - 0.42 * (uw ** 1.6)) * max(0.80, scale)
            n_buds = 5 if w_idx < 5 else 4
            for b_i in range(n_buds):
                ang = 2.0 * math.pi * b_i / n_buds + 0.55 * w_idx
                rad_dir = side_a * math.cos(ang) + side_b * math.sin(ang)
                bud_dir = (spike_axis * 0.78 + rad_dir * 0.62).normalized()
                rot_m = bud_dir.to_track_quat("Z", "X").to_matrix()
                calyx_pos = center_w + rad_dir * (0.0036 * whorl_scale)
                b_calyx.add_ellipsoid(
                    calyx_pos,
                    0.0022 * whorl_scale,
                    0.0020 * whorl_scale,
                    0.0046 * whorl_scale,
                    rot_matrix=rot_m,
                    seg_u=6,
                    seg_v=4,
                )
                if (w_idx + b_i) % 2 == 0:
                    corolla_pos = calyx_pos + bud_dir * (0.0042 * whorl_scale)
                    b_petal.add_ellipsoid(
                        corolla_pos,
                        0.0024 * whorl_scale,
                        0.0019 * whorl_scale,
                        0.0032 * whorl_scale,
                        rot_matrix=rot_m,
                        seg_u=6,
                        seg_v=4,
                    )


def add_rosemary_bush(root_pt, out_dir, lean, scale, pot_infos, pot_idx, builders, rng):
    b_wood = builders["wood_stem"]
    b_needle = builders["rose_needle"]
    b_flower = builders["rose_flower"]
    up = Vector((0.0, 0.0, 1.0))
    own_pot = pot_infos[pot_idx]

    n_branches = int(10 * max(0.75, scale))
    for br in range(n_branches):
        az = br * GOLDEN_ANGLE + rng.uniform(-0.22, 0.22)
        fan = Vector((math.cos(az), math.sin(az), 0.0))
        horiz = (out_dir * (lean * 1.05) + fan * 0.36).normalized()
        is_trailing = (br % 3 == 0) and (horiz.dot(out_dir) > 0.45)
        br_len = rng.uniform(0.14, 0.22) * scale

        pts = []
        radii = []
        n_seg = 7
        if is_trailing:
            # 3-stage lip-clearing arc: rise inside pot -> cross high over lip -> drape outside lip
            ray_phi = math.atan2(root_pt.y + horiz.y * 0.08, root_pt.x + horiz.x * 0.08)
            z_lip = max(lookup_polar(own_pot["in_z_rim"], ray_phi), lookup_polar(own_pot["out_z_top"], ray_phi))
            r_in_safe = lookup_polar(own_pot["in_r_rim"], ray_phi) - 0.012
            r_out_safe = lookup_polar(own_pot["out_r_top"], ray_phi) + 0.022
            ray_u = Vector((math.cos(ray_phi), math.sin(ray_phi), 0.0))
            p_start = push_point_outside_upper_pots(root_pt, pot_infos, pot_idx, clearance=0.010)
            p_rise = ray_u * max(0.03, r_in_safe) + Vector((0.0, 0.0, z_lip + 0.018))
            p_crest = ray_u * r_out_safe + Vector((0.0, 0.0, z_lip + 0.016))
            p_tip = ray_u * (r_out_safe + 0.042 * scale) + Vector((0.0, 0.0, z_lip - 0.038 * scale))
            ctrl = [p_start, p_rise, p_crest, p_tip]
            for j in range(n_seg + 1):
                t = j / n_seg
                omt = 1.0 - t
                pt = (
                    ctrl[0] * (omt ** 3)
                    + ctrl[1] * (3.0 * (omt ** 2) * t)
                    + ctrl[2] * (3.0 * omt * (t ** 2))
                    + ctrl[3] * (t ** 3)
                )
                pt = push_point_outside_upper_pots(pt, pot_infos, pot_idx, clearance=0.012, keep_inside_own_pot=False)
                pts.append(pt)
                radii.append(0.0028 * (1.0 - 0.58 * t) * max(0.8, scale))
        else:
            for j in range(n_seg + 1):
                t = j / n_seg
                r_off = br_len * (0.46 * (t ** 1.25))
                z_off = br_len * (0.86 * (t ** 0.85))
                pt = root_pt + horiz * r_off + up * z_off
                pt = push_point_outside_upper_pots(pt, pot_infos, pot_idx, clearance=0.012, keep_inside_own_pot=True)
                pts.append(pt)
                radii.append(0.0028 * (1.0 - 0.58 * t) * max(0.8, scale))

        b_wood.add_tube(pts, radii, segments=5, cap_end=True)

        # Dense opposite/whorled leathery rosemary needles along each branch
        n_nodes = 14
        for nd in range(2, n_nodes + 1):
            un = nd / n_nodes
            seg_f = un * (n_seg - 0.05)
            i0 = min(n_seg - 1, int(math.floor(seg_f)))
            fr = seg_f - i0
            node_pt = pts[i0].lerp(pts[i0 + 1], fr)
            tang = (pts[i0 + 1] - pts[i0]).normalized()
            s1 = tang.cross(up)
            if s1.length < 1e-5:
                s1 = Vector((1.0, 0.0, 0.0))
            s1.normalize()
            s2 = tang.cross(s1).normalized()

            whorl_rot = 0.5 * math.pi * nd + rng.uniform(-0.15, 0.15)
            n_len = (0.025 * (1.0 - 0.25 * un)) * max(0.8, scale)
            for pair in range(4):
                ang = whorl_rot + 0.5 * math.pi * pair
                rad_v = s1 * math.cos(ang) + s2 * math.sin(ang)
                ndir = (tang * 0.58 + rad_v * 0.82).normalized()
                if not is_trailing:
                    ndir = safe_leaf_dir(node_pt, ndir, n_len, pot_infos, pot_idx)
                b_needle.add_leaf_blade(node_pt, ndir, rad_v, n_len, 0.0034 * scale, arch=0.14, thickness=0.0009)

            if nd in (6, 9, 11) and rng.random() < 0.58:
                fl_pos = node_pt + s2 * 0.004
                b_flower.add_ellipsoid(fl_pos, 0.0024, 0.0020, 0.0030, seg_u=5, seg_v=4)


def add_succulent_rosette(center_pt, out_dir, scale, builders, rng, with_pups=False):
    b_succ = builders["succulent"]
    up = Vector((0.0, 0.0, 1.0))
    # Slight tilt toward sunlight/outward
    rosette_up = (up + out_dir * 0.22).normalized()
    rot_align = up.rotation_difference(rosette_up).to_matrix()

    def build_one(origin, r_scale, n_petals):
        for k in range(n_petals):
            u = k / max(1, n_petals - 1)
            az = k * GOLDEN_ANGLE
            local_r = Vector((math.cos(az), math.sin(az), 0.0))
            world_r = rot_align @ local_r
            tilt = math.radians(14.0 + 66.0 * (u ** 0.78))
            p_len = (0.015 + 0.036 * (u ** 0.72)) * r_scale
            p_wid = (0.009 + 0.017 * (u ** 0.65)) * r_scale
            p_thk = (0.0042 + 0.0045 * u) * r_scale
            base_off = origin + world_r * (0.006 * u * r_scale) + rosette_up * (0.010 * (1.0 - u) * r_scale)
            b_succ.add_succulent_petal(base_off, world_r, tilt, p_len, p_wid, p_thk)

    build_one(center_pt + Vector((0.0, 0.0, 0.005)), scale, 36)
    if with_pups:
        side = Vector((-out_dir.y, out_dir.x, 0.0))
        for sgn in (-1.0, 1.0):
            pup_pt = center_pt + out_dir * (0.022 * scale) + side * (sgn * 0.038 * scale) + Vector((0.0, 0.0, 0.003))
            build_one(pup_pt, 0.56 * scale, 24)


def add_thyme_cushion(root_pt, out_dir, r_rim, scale, pot_infos, pot_idx, builders, rng):
    b_stem = builders["wood_stem"]
    b_leaf = builders["thyme_leaf"]
    b_fl = builders["thyme_flower"]
    up = Vector((0.0, 0.0, 1.0))

    n_sprigs = int(32 * max(0.75, scale))
    for sp in range(n_sprigs):
        az = sp * GOLDEN_ANGLE + rng.uniform(-0.18, 0.18)
        local_dir = Vector((math.cos(az), math.sin(az), 0.0))
        horiz = (local_dir * 0.72 + out_dir * 0.48).normalized()
        sprig_len = rng.uniform(0.050, 0.092) * scale

        pts = []
        radii = []
        for j in range(6):
            t = j / 5.0
            r_step = sprig_len * (0.78 * t)
            z_step = (0.042 + 0.025 * rng.random()) * scale * math.sin(0.62 * math.pi * (t ** 0.75))
            pt = root_pt + horiz * r_step + up * max(0.002, z_step)
            pt = push_point_outside_upper_pots(pt, pot_infos, pot_idx, clearance=0.008, keep_inside_own_pot=True)
            pts.append(pt)
            radii.append(0.0011 * (1.0 - 0.45 * t))

        b_stem.add_tube(pts, radii, segments=4, cap_end=False)

        for j in range(1, 6):
            node = pts[j]
            tang = (pts[j] - pts[j - 1]).normalized()
            side = tang.cross(up)
            if side.length < 1e-5:
                side = Vector((1.0, 0.0, 0.0))
            side.normalize()
            for sgn in (-1.0, 1.0):
                ld = (tang * 0.55 + side * (sgn * 0.75) + up * 0.42).normalized()
                ld = safe_leaf_dir(node, ld, 0.0115 * scale, pot_infos, pot_idx)
                b_leaf.add_leaf_blade(node, ld, up, 0.0115 * scale, 0.0062 * scale, arch=0.12, thickness=0.0007)
            if j == 5 and rng.random() < 0.58:
                for fb in range(3):
                    off = Vector((rng.uniform(-0.003, 0.003), rng.uniform(-0.003, 0.003), rng.uniform(0.001, 0.004)))
                    b_fl.add_ellipsoid(node + off, 0.0020, 0.0020, 0.0022, seg_u=5, seg_v=4)


def build_pot_earth_and_plants(pot_infos, col_earth, col_plants):
    rng = random.Random(RANDOM_SEED)
    earth_mat = create_earth_material()
    pebble_mat = create_pebble_material()

    mats = {
        "lav_leaf": create_foliage_material(
            "Plant_Lavender_Foliage", (0.115, 0.175, 0.130), (0.215, 0.285, 0.225), roughness=0.62, translucency=0.18
        ),
        "lav_stalk": create_foliage_material(
            "Plant_Lavender_Stalk", (0.105, 0.160, 0.095), (0.175, 0.240, 0.145), roughness=0.58, translucency=0.14
        ),
        "lav_calyx": create_foliage_material(
            "Plant_Lavender_Calyx", (0.115, 0.055, 0.245), (0.215, 0.110, 0.390), roughness=0.55, translucency=0.22, noise_scale=35.0
        ),
        "lav_petal": create_foliage_material(
            "Plant_Lavender_Corolla", (0.240, 0.125, 0.480), (0.390, 0.225, 0.660), roughness=0.46, translucency=0.30, noise_scale=42.0
        ),
        "wood_stem": create_bark_material("Plant_Woody_Stem", (0.10, 0.068, 0.045), (0.20, 0.150, 0.112)),
        "rose_needle": create_foliage_material(
            "Plant_Rosemary_Needles", (0.045, 0.105, 0.042), (0.105, 0.195, 0.088), roughness=0.38, translucency=0.16
        ),
        "rose_flower": create_foliage_material(
            "Plant_Rosemary_Blossom", (0.32, 0.42, 0.72), (0.52, 0.62, 0.88), roughness=0.45, translucency=0.28
        ),
        "succulent": create_succulent_material(),
        "thyme_leaf": create_foliage_material(
            "Plant_Thyme_Foliage", (0.085, 0.138, 0.068), (0.165, 0.235, 0.132), roughness=0.54, translucency=0.20
        ),
        "thyme_flower": create_foliage_material(
            "Plant_Thyme_Blossom", (0.58, 0.44, 0.50), (0.78, 0.68, 0.72), roughness=0.50, translucency=0.26
        ),
    }

    builders = {k: MeshBuilder() for k in mats}
    pebble_builder = MeshBuilder()

    for pot_idx in range(len(pot_infos)):
        stations, z_soil_base = find_pot_planting_stations(pot_infos, pot_idx)
        build_soil_mesh_for_pot(pot_infos, pot_idx, z_soil_base, stations, col_earth, earth_mat, pebble_builder, rng)

        for st in stations:
            phi = st["phi"]
            r = st["r"]
            z = st["z"]
            root_pt = Vector((r * math.cos(phi), r * math.sin(phi), z))
            out_dir = Vector((math.cos(phi), math.sin(phi), 0.0)) if r > 1e-4 else Vector((0.35, -0.93, 0.0)).normalized()
            kind = st["kind"]
            sc = st["scale"]
            lean = st["lean"]

            if kind == "lavender_hero":
                add_lavender_cluster(root_pt, out_dir, lean, sc, pot_infos, pot_idx, builders, rng, is_hero=True)
            elif kind == "lavender":
                add_lavender_cluster(root_pt, out_dir, lean, sc, pot_infos, pot_idx, builders, rng, is_hero=False)
            elif kind == "rosemary":
                add_rosemary_bush(root_pt, out_dir, lean, sc, pot_infos, pot_idx, builders, rng)
            elif kind == "succulent_cluster":
                add_succulent_rosette(root_pt, out_dir, sc, builders, rng, with_pups=True)
            elif kind == "succulent":
                add_succulent_rosette(root_pt, out_dir, sc, builders, rng, with_pups=True)
            elif kind == "thyme":
                add_thyme_cushion(root_pt, out_dir, st.get("r_rim", r + 0.04), sc, pot_infos, pot_idx, builders, rng)

    pebble_builder.build("Pots_Soil_Mulch_Pebbles", col_earth, pebble_mat, smooth=False)
    for k, b in builders.items():
        b.build(f"Planter_Botanical_{k}", col_plants, mats[k], smooth=True)


# ------------------------------------------------------------------------------
# TERRACE PAVING & LUSH MEDITERRANEAN GARDEN BACKGROUND
# ------------------------------------------------------------------------------

def add_beveled_box(name, x0, x1, y0, y1, z0, z1, bev, collection, material):
    b = min(bev, 0.25 * (x1 - x0), 0.25 * (y1 - y0), 0.35 * (z1 - z0))
    verts = [
        # Bottom rectangle (0..3)
        (x0, y0, z0), (x1, y0, z0), (x1, y1, z0), (x0, y1, z0),
        # Upper outer collar just below bevel (4..7)
        (x0, y0, z1 - b), (x1, y0, z1 - b), (x1, y1, z1 - b), (x0, y1, z1 - b),
        # Beveled top face (8..11)
        (x0 + b, y0 + b, z1), (x1 - b, y0 + b, z1), (x1 - b, y1 - b, z1), (x0 + b, y1 - b, z1),
    ]
    faces = [
        (0, 3, 2, 1),
        (0, 1, 5, 4), (1, 2, 6, 5), (2, 3, 7, 6), (3, 0, 4, 7),
        (4, 5, 9, 8), (5, 6, 10, 9), (6, 7, 11, 10), (7, 4, 8, 11),
        (8, 9, 10, 11),
    ]
    return add_mesh_object(name, verts, faces, collection, material=material, smooth=False)


def build_terrace_and_garden(col_terrace, col_garden):
    rng = random.Random(RANDOM_SEED + 101)
    stone_mat = create_travertine_stone_material()
    mortar_mat = create_limewash_wall_material("Terrace_Sandy_Mortar", (0.16, 0.125, 0.090), (0.24, 0.19, 0.14))
    wall_mat = create_limewash_wall_material("Terrace_Limewash_Parapet", (0.42, 0.34, 0.24), (0.56, 0.47, 0.35))
    coping_mat = create_limewash_wall_material("Terrace_Limestone_Coping", (0.38, 0.31, 0.22), (0.50, 0.42, 0.31))

    terrace_y_front = -2.90
    terrace_y_back = 1.48
    terrace_half_x = 3.20

    # Mortar bed under flagstones
    add_beveled_box(
        "Terrace_Mortar_Bed",
        -terrace_half_x, terrace_half_x, terrace_y_front, terrace_y_back,
        -0.060, -0.0045, 0.001, col_terrace, mortar_mat,
    )

    # Running-bond travertine limestone flagstones (top precisely at Z = 0.000)
    tile_w = 0.56
    tile_d = 0.40
    joint = 0.006
    n_rows = int(math.ceil((terrace_y_back - terrace_y_front) / tile_d))
    paver_idx = 0
    for r_i in range(n_rows):
        y0 = terrace_y_front + r_i * tile_d
        y1 = min(terrace_y_back, y0 + tile_d - joint)
        if y1 - y0 < 0.04:
            continue
        x_shift = (0.5 * tile_w) if (r_i % 2 == 1) else 0.0
        x_cur = -terrace_half_x - x_shift
        while x_cur < terrace_half_x:
            x0 = max(-terrace_half_x, x_cur)
            x1 = min(terrace_half_x, x_cur + tile_w - joint)
            if x1 - x0 > 0.05:
                paver_idx += 1
                z_jit = rng.uniform(-0.0006, 0.0004) if (abs(x0 + x1) > 0.9 or abs(y0 + y1) > 0.9) else 0.0
                add_beveled_box(
                    f"Terrace_Paver_{paver_idx:03d}",
                    x0, x1, y0, y1, -0.022, z_jit, 0.0035,
                    col_terrace, stone_mat,
                )
            x_cur += tile_w

    # Low Mediterranean terrace parapet / garden curb along Y = terrace_y_back
    curb_h = 0.28
    add_beveled_box(
        "Terrace_Parapet_Wall",
        -terrace_half_x, terrace_half_x, terrace_y_back, terrace_y_back + 0.20,
        -0.05, curb_h, 0.006, col_terrace, wall_mat,
    )
    # Segmented chiselled limestone coping stones on top of the low wall
    coping_w = 0.78
    cx = -terrace_half_x
    c_i = 0
    while cx < terrace_half_x:
        c_i += 1
        x1 = min(terrace_half_x, cx + coping_w - 0.005)
        add_beveled_box(
            f"Terrace_Coping_{c_i:02d}",
            cx, x1, terrace_y_back - 0.022, terrace_y_back + 0.222,
            curb_h, curb_h + 0.046, 0.005, col_terrace, coping_mat,
        )
        cx += coping_w

    if not ADD_GARDEN_BACKGROUND:
        return

    # --------------------------------------------------------------------------
    # MEDITERRANEAN GARDEN BACKGROUND (Y = 1.65 m .. 15.0 m)
    # --------------------------------------------------------------------------
    meadow_mat = create_foliage_material(
        "Garden_Mediterranean_Meadow", (0.075, 0.115, 0.042), (0.165, 0.205, 0.078), roughness=0.86, translucency=0.05, noise_scale=3.5
    )
    gravel_mat = create_limewash_wall_material("Garden_Crushed_Limestone_Path", (0.48, 0.42, 0.32), (0.64, 0.57, 0.46))
    olive_bark_mat = create_bark_material("Garden_Olive_Bark", (0.14, 0.125, 0.105), (0.28, 0.255, 0.220))
    olive_leaf_mat = create_foliage_material(
        "Garden_Olive_Canopy", (0.095, 0.145, 0.082), (0.225, 0.285, 0.195), roughness=0.52, translucency=0.26, noise_scale=12.0
    )
    cypress_mat = create_foliage_material(
        "Garden_Italian_Cypress", (0.024, 0.058, 0.028), (0.058, 0.115, 0.052), roughness=0.68, translucency=0.12, noise_scale=22.0
    )
    shrub_sage_mat = create_foliage_material(
        "Garden_Shrub_Sage", (0.065, 0.130, 0.062), (0.150, 0.235, 0.125), roughness=0.58, translucency=0.20, noise_scale=14.0
    )
    lav_bg_mat = create_foliage_material(
        "Garden_Lavender_Drift", (0.165, 0.095, 0.340), (0.310, 0.180, 0.540), roughness=0.52, translucency=0.26, noise_scale=28.0
    )
    oleander_blossom_mat = create_foliage_material(
        "Garden_Warm_Blossoms", (0.58, 0.20, 0.22), (0.82, 0.42, 0.38), roughness=0.48, translucency=0.32, noise_scale=30.0
    )
    grass_plume_mat = create_foliage_material(
        "Garden_Golden_Stipa_Grass", (0.32, 0.26, 0.12), (0.58, 0.48, 0.24), roughness=0.55, translucency=0.34, noise_scale=20.0
    )

    def garden_z(gx, gy):
        t = max(0.0, (gy - 1.65) / 12.0)
        return 0.06 + 0.58 * (t ** 1.15) + 0.09 * math.sin(0.55 * gx + 0.35 * gy) + 0.05 * math.cos(0.9 * gx - 0.4 * gy)

    # 1. Rolling garden ground mesh + winding crushed-stone garden path
    nx_g, ny_g = 48, 36
    gx0, gx1 = -9.5, 9.5
    gy0, gy1 = 1.62, 15.5
    g_verts = []
    g_faces = []
    p_verts = []
    p_faces = []

    for iy in range(ny_g + 1):
        vy = gy0 + (gy1 - gy0) * iy / ny_g
        path_cx = 0.55 - 1.35 * math.sin(0.26 * (vy - 1.6))
        path_hw = 0.48 + 0.03 * (vy - 1.6)
        p_verts.append((path_cx - path_hw, vy, garden_z(path_cx - path_hw, vy) + 0.014))
        p_verts.append((path_cx + path_hw, vy, garden_z(path_cx + path_hw, vy) + 0.014))
        if iy < ny_g:
            b = 2 * iy
            p_faces.append((b, b + 1, b + 3, b + 2))

        for ix in range(nx_g + 1):
            vx = gx0 + (gx1 - gx0) * ix / nx_g
            g_verts.append((vx, vy, garden_z(vx, vy)))

    stride = nx_g + 1
    for iy in range(ny_g):
        for ix in range(nx_g):
            a = iy * stride + ix
            g_faces.append((a, a + 1, a + stride + 1, a + stride))

    add_mesh_object("Garden_Meadow_Terrain", g_verts, g_faces, col_garden, material=meadow_mat, smooth=True)
    add_mesh_object("Garden_Limestone_Path", p_verts, p_faces, col_garden, material=gravel_mat, smooth=True)

    # Warm ochre dry-stone / limewashed garden wall in the far garden
    wall_y = 10.6
    wz_base = garden_z(0.0, wall_y) - 0.15
    add_beveled_box("Garden_Back_Stone_Wall", -9.0, 9.0, wall_y, wall_y + 0.32, wz_base, wz_base + 0.95, 0.015, col_garden, wall_mat)

    b_olive_bark = MeshBuilder()
    b_olive_leaf = MeshBuilder()
    b_cypress = MeshBuilder()
    b_shrub = MeshBuilder()
    b_lav_bg = MeshBuilder()
    b_blossom = MeshBuilder()
    b_grass = MeshBuilder()

    # 2. Midground Mediterranean shrub border (organic leafy sprays — lavender drifts, sage/myrtle bushes, oleander, golden grasses)
    shrub_specs = [
        (-1.85, 2.25, 0.46, "lavender"),
        (-1.12, 2.38, 0.40, "lavender"),
        (-0.48, 2.18, 0.36, "grass"),
        (0.92, 2.22, 0.44, "lavender"),
        (1.62, 2.48, 0.48, "shrub_bloom"),
        (2.35, 2.85, 0.56, "shrub"),
        (-2.45, 2.95, 0.58, "shrub_bloom"),
        (-0.85, 3.35, 0.52, "shrub"),
        (0.18, 3.60, 0.44, "lavender"),
        (1.25, 3.55, 0.54, "grass"),
        (2.05, 4.10, 0.64, "shrub_bloom"),
        (-2.10, 4.40, 0.60, "lavender"),
        (-0.45, 5.10, 0.68, "shrub"),
        (0.95, 5.50, 0.72, "shrub_bloom"),
        (-1.60, 6.60, 0.82, "shrub"),
        (2.60, 6.40, 0.78, "shrub"),
        (0.20, 7.60, 0.88, "shrub"),
        (-1.05, 8.20, 0.85, "shrub_bloom"),
    ]

    up = Vector((0.0, 0.0, 1.0))
    for sx, sy, rad, skind in shrub_specs:
        sz = garden_z(sx, sy)
        base_c = Vector((sx, sy, sz))
        if skind == "lavender":
            # Organic leafy silvery-sage mound + slender tapered violet flower spikes
            for k in range(18):
                u = (k + 0.5) / 18.0
                az = k * GOLDEN_ANGLE
                rr = rad * 0.68 * math.sqrt(u)
                zz = rad * 0.44 * math.cos(0.45 * math.pi * math.sqrt(u))
                pos = base_c + Vector((rr * math.cos(az), rr * math.sin(az), zz + 0.04))
                b_shrub.add_foliage_spray(pos, rad * 0.26, 20, 0.075, 0.020, rng, upward_bias=0.45)
            for k in range(44):
                u = (k + 0.5) / 44.0
                az = k * GOLDEN_ANGLE + 0.4
                rr = rad * 0.82 * math.sqrt(u)
                zz = rad * (0.50 + 0.20 * rng.random()) * math.cos(0.42 * math.pi * math.sqrt(u))
                p0 = base_c + Vector((rr * math.cos(az), rr * math.sin(az), zz + 0.04))
                tilt_v = Vector((math.cos(az) * 0.32, math.sin(az) * 0.32, 1.0)).normalized()
                p1 = p0 + tilt_v * rng.uniform(0.09, 0.14)
                b_lav_bg.add_tube([p0, p1], [0.014, 0.004], segments=5, cap_end=True)
        elif skind == "grass":
            # Feathery golden Mediterranean Stipa grass fountain
            for k in range(65):
                az = k * GOLDEN_ANGLE + rng.uniform(-0.1, 0.1)
                hd = Vector((math.cos(az), math.sin(az), 0.0))
                gl = rad * rng.uniform(0.95, 1.50)
                arch_dir = (hd * rng.uniform(0.45, 0.88) + up * rng.uniform(0.72, 1.15)).normalized()
                b_grass.add_leaf_blade(base_c, arch_dir, up, gl, 0.015, arch=0.48, thickness=0.002)
        else:
            n_clusters = 24
            for k in range(n_clusters):
                u = (k + 0.5) / n_clusters
                az = k * GOLDEN_ANGLE
                rr = rad * 0.78 * math.sqrt(u) * rng.uniform(0.82, 1.15)
                zz = rad * 0.62 * math.cos(0.46 * math.pi * math.sqrt(u)) * rng.uniform(0.85, 1.15)
                pos = base_c + Vector((rr * math.cos(az), rr * math.sin(az), zz + 0.08))
                b_shrub.add_foliage_spray(pos, rad * 0.30, 22, 0.095, 0.038, rng, upward_bias=0.30)
                if skind == "shrub_bloom" and k % 3 == 0:
                    for b_i in range(4):
                        fl_off = Vector((rng.uniform(-0.06, 0.06), rng.uniform(-0.06, 0.04), rad * 0.14 + rng.uniform(0.0, 0.06)))
                        b_blossom.add_ellipsoid(pos + fl_off, 0.024, 0.022, 0.018, seg_u=6, seg_v=4)

    # 3. Hero Gnarled Mediterranean Olive Tree (left-midground garden, framing upper-left with airy silver-green leaves)
    olive_base = Vector((-1.32, 4.15, garden_z(-1.32, 4.15) - 0.04))
    trunk_pts = [
        olive_base,
        olive_base + Vector((0.08, 0.02, 0.38)),
        olive_base + Vector((0.02, -0.04, 0.76)),
        olive_base + Vector((0.14, -0.02, 1.12)),
    ]
    b_olive_bark.add_tube(trunk_pts, [0.145, 0.118, 0.100, 0.084], segments=10, cap_end=True)

    crown_top = trunk_pts[-1]
    limb_specs = [
        (Vector((-0.55, -0.28, 0.42)), 0.48),
        (Vector((0.68, -0.36, 0.46)), 0.54),
        (Vector((0.28, -0.45, 0.26)), 0.42),
        (Vector((-0.20, 0.42, 0.62)), 0.48),
        (Vector((0.42, 0.30, 0.72)), 0.50),
    ]
    for limb_vec, cloud_r in limb_specs:
        mid_l = crown_top + limb_vec * 0.52 + Vector((0.0, 0.0, 0.06))
        end_l = crown_top + limb_vec
        b_olive_bark.add_tube([crown_top, mid_l, end_l], [0.060, 0.040, 0.020], segments=7, cap_end=True)
        for k in range(26):
            u = (k + 0.5) / 26.0
            az = k * GOLDEN_ANGLE
            rr = cloud_r * math.sqrt(u) * rng.uniform(0.72, 1.15)
            zz = cloud_r * 0.44 * math.sin(math.pi * u) + rng.uniform(-0.08, 0.10)
            lpos = end_l + Vector((rr * math.cos(az), rr * math.sin(az), zz))
            b_olive_leaf.add_foliage_spray(lpos, 0.18, 20, 0.085, 0.026, rng, upward_bias=0.18)

    # 4. Slender Tuscan / Provençal Cypresses rising with feathered evergreen sprays (no bubble spheres)
    cypress_specs = [
        (1.48, 5.95, 3.45, 0.34),
        (2.28, 7.55, 3.95, 0.38),
        (-2.65, 8.15, 3.70, 0.36),
        (0.58, 9.20, 4.20, 0.40),
    ]
    for cx_p, cy_p, c_height, c_rad in cypress_specs:
        cz_p = garden_z(cx_p, cy_p) - 0.05
        c_base = Vector((cx_p, cy_p, cz_p))
        b_olive_bark.add_tube([c_base, c_base + Vector((0.0, 0.0, c_height * 0.75))], [0.072, 0.025], segments=7, cap_end=True)
        n_tiers = 44
        for t_i in range(n_tiers):
            ut = t_i / (n_tiers - 1)
            env = math.sin(math.pi * (0.10 + 0.86 * (ut ** 0.78))) * (1.0 - 0.28 * ut)
            r_tier = c_rad * max(0.06, env)
            z_tier = cz_p + 0.26 + ut * (c_height - 0.26)
            n_tufts = 5 if r_tier > 0.16 else 3
            for j in range(n_tufts):
                ang = t_i * GOLDEN_ANGLE + 2.0 * math.pi * j / n_tufts
                rr = 0.54 * r_tier * rng.uniform(0.75, 1.12)
                tpos = Vector((cx_p + rr * math.cos(ang), cy_p + rr * math.sin(ang), z_tier))
                b_cypress.add_foliage_spray(tpos, r_tier * 0.55, 16, 0.095, 0.036, rng, upward_bias=0.68)

    b_olive_bark.build("Garden_Olive_And_Tree_Trunks", col_garden, olive_bark_mat, smooth=True)
    b_olive_leaf.build("Garden_Olive_Foliage", col_garden, olive_leaf_mat, smooth=True)
    b_cypress.build("Garden_Italian_Cypresses", col_garden, cypress_mat, smooth=True)
    b_shrub.build("Garden_Evergreen_Shrubs", col_garden, shrub_sage_mat, smooth=True)
    b_lav_bg.build("Garden_Lavender_Spikes", col_garden, lav_bg_mat, smooth=True)
    b_blossom.build("Garden_Mediterranean_Blossoms", col_garden, oleander_blossom_mat, smooth=True)
    b_grass.build("Garden_Ornamental_Grasses", col_garden, grass_plume_mat, smooth=True)


# ------------------------------------------------------------------------------
# LIGHTING, CAMERA & CYCLES RENDER ENGINE
# ------------------------------------------------------------------------------

def create_terrace_camera(col_cam, stack_top_z):
    # 52mm architectural/portrait lens with 3/4 elevated view into the pot earth & creamy garden bokeh
    target_z = max(0.32, 0.52 * stack_top_z)
    cam_dist = max(2.12, 2.48 * stack_top_z)
    cam_loc = Vector((-0.38, -cam_dist, target_z + 0.72))
    cam_target = Vector((0.0, 0.0, target_z))

    cam_data = bpy.data.cameras.new("Terrace_Planter_Camera")
    cam_data.lens = 52.0
    cam_data.sensor_width = 36.0
    cam_data.clip_start = 0.05
    cam_data.clip_end = 80.0
    cam_data.dof.use_dof = True
    cam_data.dof.focus_distance = (cam_target - cam_loc).length - 0.12
    cam_data.dof.aperture_fstop = 2.8

    cam_obj = bpy.data.objects.new("Terrace_Planter_Camera", cam_data)
    cam_obj.location = cam_loc
    direction = cam_target - cam_loc
    cam_obj.rotation_euler = direction.to_track_quat("-Z", "Y").to_euler()
    col_cam.objects.link(cam_obj)
    return cam_obj


def setup_mediterranean_sky_and_lights(col_lights):
    # 1. Physical Mediterranean sky world (calibrated strength for Nishita spectral radiance)
    world = bpy.data.worlds.new("Mediterranean_Daylight_Sky")
    nt = ensure_node_tree(world)
    nodes = nt.nodes
    links = nt.links
    nodes.clear()

    out = nodes.new("ShaderNodeOutputWorld")
    out.location = (360, 0)
    bg = nodes.new("ShaderNodeBackground")
    bg.location = (140, 0)
    links.new(bg.outputs["Background"], out.inputs["Surface"])

    sky = None
    try:
        sky = nodes.new("ShaderNodeTexSky")
        sky.location = (-160, 0)
        for st in ("MULTIPLE_SCATTERING", "NISHITA", "HOSEK_WILKIE", "PREETHAM"):
            try:
                sky.sky_type = st
                break
            except Exception:
                continue
        if hasattr(sky, "sun_disc"):
            sky.sun_disc = False
        if hasattr(sky, "sun_elevation"):
            sky.sun_elevation = math.radians(28.0)
        if hasattr(sky, "sun_rotation"):
            sky.sun_rotation = math.radians(125.0)
        if hasattr(sky, "air_density"):
            sky.air_density = 1.15
        if hasattr(sky, "dust_density"):
            sky.dust_density = 1.25
        if hasattr(sky, "ozone_density"):
            sky.ozone_density = 1.40
        set_socket(bg, "Strength", 0.11)
        links.new(sky.outputs["Color"], bg.inputs["Color"])
    except Exception:
        set_socket(bg, "Color", (0.28, 0.50, 0.88, 1.0))
        set_socket(bg, "Strength", 0.85)

    # 2. Golden late-afternoon raking Mediterranean Sun (sculpts terracotta undulations & backlights garden)
    sun_data = bpy.data.lights.new("Mediterranean_Golden_Sun", type="SUN")
    sun_data.energy = 3.8
    sun_data.color = (1.0, 0.89, 0.74)
    if hasattr(sun_data, "angle"):
        sun_data.angle = math.radians(1.5)
    sun_obj = bpy.data.objects.new("Mediterranean_Golden_Sun", sun_data)
    sun_from = Vector((-3.6, -2.1, 2.75))
    sun_obj.location = sun_from
    sun_obj.rotation_euler = (-sun_from).to_track_quat("-Z", "Y").to_euler()
    col_lights.objects.link(sun_obj)

    # 3. Subtle warm limestone terrace bounce fill from camera-right so shadow troughs stay luminous
    fill_data = bpy.data.lights.new("Terrace_Warm_Fill", type="AREA")
    fill_data.shape = "RECTANGLE"
    fill_data.size = 2.8
    fill_data.size_y = 2.0
    fill_data.energy = 9.5
    fill_data.color = (1.0, 0.94, 0.86)
    fill_obj = bpy.data.objects.new("Terrace_Warm_Fill", fill_data)
    fill_loc = Vector((2.4, -2.2, 1.35))
    fill_obj.location = fill_loc
    fill_obj.rotation_euler = (Vector((0.0, 0.0, 0.55)) - fill_loc).to_track_quat("-Z", "Y").to_euler()
    col_lights.objects.link(fill_obj)

    return world


def enable_cycles_gpu():
    try:
        cprefs = bpy.context.preferences.addons["cycles"].preferences
    except Exception:
        return "CPU"

    preferred_backends = ("METAL", "OPTIX", "CUDA", "HIP", "ONEAPI")
    try:
        available = [item[0] for item in cprefs.get_device_types(bpy.context)]
    except Exception:
        available = list(preferred_backends)

    for backend in preferred_backends:
        if backend not in available:
            continue
        try:
            cprefs.compute_device_type = backend
            if hasattr(cprefs, "refresh_devices"):
                cprefs.refresh_devices()
            elif hasattr(cprefs, "get_devices"):
                cprefs.get_devices()
            found_gpu = False
            for dev in cprefs.devices:
                if dev.type == backend:
                    dev.use = True
                    found_gpu = True
                elif dev.type == "CPU":
                    dev.use = False
            if found_gpu:
                return backend
        except Exception:
            continue

    return "CPU"


def configure_scene_render(scene, camera_obj, world, samples):
    scene.camera = camera_obj
    scene.world = world

    rd = scene.render
    rd.engine = "CYCLES"
    rd.resolution_x = RENDER_RESOLUTION_X
    rd.resolution_y = RENDER_RESOLUTION_Y
    rd.resolution_percentage = 100
    rd.image_settings.file_format = "PNG"
    rd.image_settings.color_mode = "RGBA"
    rd.image_settings.color_depth = "8"

    vs = scene.view_settings
    for vt in ("AgX", "Filmic", "Standard"):
        try:
            vs.view_transform = vt
            break
        except Exception:
            continue
    for lk in ("AgX - Medium High Contrast", "Medium High Contrast", "None"):
        try:
            vs.look = lk
            break
        except Exception:
            continue
    vs.exposure = 0.0

    cscene = scene.cycles
    cscene.samples = samples
    cscene.use_adaptive_sampling = True
    cscene.adaptive_threshold = 0.012
    cscene.use_denoising = True
    if hasattr(cscene, "use_preview_denoising"):
        cscene.use_preview_denoising = True
    if hasattr(cscene, "use_light_tree"):
        cscene.use_light_tree = True
    cscene.max_bounces = 8
    cscene.diffuse_bounces = 4
    cscene.glossy_bounces = 4
    cscene.transmission_bounces = 6
    cscene.transparent_max_bounces = 8
    cscene.sample_clamp_indirect = 10.0


def resolve_output_directory():
    try:
        space = getattr(bpy.context, "space_data", None)
        if space is not None and getattr(space, "text", None) is not None and space.text.filepath:
            p = bpy.path.abspath(space.text.filepath)
            if os.path.isdir(os.path.dirname(p)):
                return os.path.dirname(p)
    except Exception:
        pass

    try:
        if "__file__" in globals() and __file__:
            d = os.path.dirname(os.path.abspath(__file__))
            if os.path.isdir(d) and os.access(d, os.W_OK):
                return d
    except Exception:
        pass

    dl = os.path.expanduser("~/Downloads")
    if os.path.isdir(dl) and os.access(dl, os.W_OK):
        return dl
    return "/tmp"


def build_and_render():
    scene = clear_all_datablocks()
    scene.name = "Mediterranean_Terrace_Planter"
    gpu_backend = enable_cycles_gpu()

    col_pots = make_collection("01_Stacked_TerraCotta_Pots", scene.collection)
    col_earth = make_collection("02_Pot_Earth_And_Mulch", scene.collection)
    col_plants = make_collection("03_Mediterranean_Plants", scene.collection)
    col_terrace = make_collection("04_Terrace_Architecture", scene.collection)
    col_garden = make_collection("05_Garden_Background", scene.collection)
    col_lights = make_collection("06_Lights_And_Camera", scene.collection)

    # Provisional camera direction for aligning the iconic Phylla viewing angle
    prov_cam_xy = (-0.38, -2.15)
    _, pot_infos = build_stacked_pots(col_pots, prov_cam_xy)

    stack_top_z = max(max(p["out_z_top"]) for p in pot_infos) if pot_infos else 0.75

    if ADD_EARTH_AND_PLANTS and pot_infos:
        build_pot_earth_and_plants(pot_infos, col_earth, col_plants)

    build_terrace_and_garden(col_terrace, col_garden)

    cam_obj = create_terrace_camera(col_lights, stack_top_z)
    world = setup_mediterranean_sky_and_lights(col_lights)
    configure_scene_render(scene, cam_obj, world, CYCLES_SAMPLES)
    if gpu_backend != "CPU":
        scene.cycles.device = "GPU"

    out_dir = resolve_output_directory()
    out_path = os.path.join(out_dir, f"{PLANT_SLUG}_terrace_planter.png")
    scene.render.filepath = out_path

    print(f"[Phylla Terrace Planter] Built {len(pot_infos)} stacked terra cotta pots using Cycles ({gpu_backend}).")
    print(f"  Render output: {out_path}")

    if AUTO_RENDER_IMAGE:
        bpy.ops.render.render(write_still=True, scene=scene.name)
        if os.path.isfile(out_path):
            img = bpy.data.images.load(out_path, check_existing=True)
            img.name = "Render_Mediterranean_Terrace_Planter"
        print("[Phylla Terrace Planter] Render complete!")

    if bpy.context.window is not None:
        bpy.context.window.scene = scene
        for screen in bpy.data.screens:
            for area in screen.areas:
                if area.type == "VIEW_3D":
                    for space in area.spaces:
                        if space.type == "VIEW_3D" and space.region_3d is not None:
                            space.region_3d.view_perspective = "CAMERA"


if __name__ == "__main__":
    build_and_render()
`
