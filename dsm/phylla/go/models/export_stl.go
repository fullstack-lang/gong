package models

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type vector3 struct {
	X, Y, Z float64
}

func (v vector3) sub(other vector3) vector3 {
	return vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v vector3) cross(other vector3) vector3 {
	return vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v vector3) normalize() vector3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if length == 0 {
		return vector3{0, 0, 0}
	}
	return vector3{v.X / length, v.Y / length, v.Z / length}
}

func writeFacet(sb *strings.Builder, v1, v2, v3 vector3) {
	normal := v2.sub(v1).cross(v3.sub(v1)).normalize()
	sb.WriteString(fmt.Sprintf("  facet normal %e %e %e\n", normal.X, normal.Y, normal.Z))
	sb.WriteString("    outer loop\n")
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v1.X, v1.Y, v1.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v2.X, v2.Y, v2.Z))
	sb.WriteString(fmt.Sprintf("      vertex %e %e %e\n", v3.X, v3.Y, v3.Z))
	sb.WriteString("    endloop\n")
	sb.WriteString("  endfacet\n")
}

func unwrapAnglesSTL(pts []vector3) (angles []float64, points []vector3) {
	if len(pts) == 0 {
		return nil, nil
	}
	angleToPoint := make(map[float64]vector3)

	firstP := pts[0]
	lastTheta := math.Atan2(firstP.Z, firstP.X)

	accumulated := lastTheta
	for accumulated < 0 {
		accumulated += 2 * math.Pi
	}
	for accumulated >= 2*math.Pi {
		accumulated -= 2 * math.Pi
	}

	angleToPoint[accumulated] = firstP
	lastTheta = math.Atan2(firstP.Z, firstP.X)

	for i := 1; i < len(pts); i++ {
		p := pts[i]
		theta := math.Atan2(p.Z, p.X)
		diff := theta - lastTheta

		for diff < -math.Pi {
			diff += 2 * math.Pi
		}
		for diff > math.Pi {
			diff -= 2 * math.Pi
		}

		if diff < -1e-7 {
			diff = 0
		}

		accumulated += diff
		angleToPoint[accumulated] = p
		lastTheta = theta
	}

	for a := range angleToPoint {
		angles = append(angles, a)
	}
	sort.Float64s(angles)

	for _, a := range angles {
		points = append(points, angleToPoint[a])
	}

	return angles, points
}

func getTargetAnglesSTL(
	originalBottom []vector3,
	originalTop []vector3,
	degreeInterval float64,
	radialRepetitions int,
) (targetAngles []float64, bottomAngles []float64, bottomPoints []vector3, topAngles []float64, topPoints []vector3) {
	if len(originalBottom) == 0 || len(originalTop) == 0 {
		return nil, nil, nil, nil, nil
	}

	bottomAngles, bottomPoints = unwrapAnglesSTL(originalBottom)
	topAngles, topPoints = unwrapAnglesSTL(originalTop)

	radInterval := degreeInterval * math.Pi / 180.0

	if radialRepetitions < 1 {
		radialRepetitions = 1
	}
	expectedDegrees := 360.0 / float64(radialRepetitions)

	nbPoints := int(math.Round(expectedDegrees / degreeInterval))
	for i := 0; i <= nbPoints; i++ {
		targetAngles = append(targetAngles, float64(i)*radInterval)
	}

	return targetAngles, bottomAngles, bottomPoints, topAngles, topPoints
}

func resampleCurveAtAnglesSTL(
	sortedAngles []float64,
	sortedPoints []vector3,
	targetAngles []float64,
	expectedDegrees float64,
) []vector3 {
	var resampled []vector3
	if len(sortedAngles) == 0 {
		return resampled
	}

	for _, target := range targetAngles {
		evalTarget := target
		if len(sortedAngles) > 0 && expectedDegrees > 0 {
			minA := sortedAngles[0]
			maxA := sortedAngles[len(sortedAngles)-1]
			expectedRad := expectedDegrees * math.Pi / 180.0

			for evalTarget < minA {
				evalTarget += expectedRad
			}
			for evalTarget > maxA {
				evalTarget -= expectedRad
			}
			if evalTarget < minA {
				evalTarget = minA
			}
			if evalTarget > maxA {
				evalTarget = maxA
			}
		}

		searchIdx := sort.SearchFloat64s(sortedAngles, evalTarget)
		idx := -1
		if searchIdx > 0 && searchIdx < len(sortedAngles) {
			idx = searchIdx - 1
		} else if searchIdx == 0 && sortedAngles[0] == evalTarget {
			idx = 0
		}

		isExtrapolating := false
		if idx == -1 {
			isExtrapolating = true
			if searchIdx == 0 {
				idx = 0
			} else {
				idx = len(sortedAngles) - 2
			}
		}

		a1 := sortedAngles[idx]
		a2 := sortedAngles[idx+1]
		p1 := sortedPoints[idx]
		p2 := sortedPoints[idx+1]

		t := 0.0
		if a1 != a2 {
			t = (evalTarget - a1) / (a2 - a1)
		}
		if t < 0 {
			t = 0
		}
		if t > 1 {
			t = 1
		}

		var x, y, z float64
		if isExtrapolating {
			var r float64
			if t == 0 {
				r = math.Hypot(p1.X, p1.Z)
				y = p1.Y
			} else {
				r = math.Hypot(p2.X, p2.Z)
				y = p2.Y
			}
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		} else {
			r1 := math.Hypot(p1.X, p1.Z)
			r2 := math.Hypot(p2.X, p2.Z)
			r := r1 + t*(r2-r1)
			y = p1.Y + t*(p2.Y-p1.Y)
			x = r * math.Cos(target)
			z = r * math.Sin(target)
		}

		resampled = append(resampled, vector3{X: x, Y: y, Z: z})
	}

	return resampled
}

func rotateCurveSTL(source []vector3, thetaOffset float64) []vector3 {
	rotated := make([]vector3, len(source))
	for i, p := range source {
		thetaBase := math.Atan2(p.Z, p.X)
		r := math.Hypot(p.X, p.Z)
		newTheta := thetaBase + thetaOffset
		rotated[i] = vector3{
			X: r * math.Cos(newTheta),
			Y: p.Y,
			Z: r * math.Sin(newTheta),
		}
	}
	return rotated
}

func writeRibbonLayerSTL(
	sb *strings.Builder,
	massiveBottom []vector3,
	massiveTop []vector3,
	dy float64,
	thickness float64,
	radialRepetitions int,
) {
	if len(massiveBottom) == 0 || len(massiveBottom) != len(massiveTop) {
		return
	}

	numPointsPerRep := len(massiveBottom) / radialRepetitions
	if numPointsPerRep == 0 {
		numPointsPerRep = 1
	}

	for i := 0; i < len(massiveBottom)-1; i++ {
		// Skip the bridging quad between separate radial repetitions
		if (i+1)%numPointsPerRep == 0 {
			continue
		}

		pB := massiveBottom[i]
		pB_next := massiveBottom[i+1]
		pT := massiveTop[i]
		pT_next := massiveTop[i+1]

		th_i := math.Atan2(pB.Z, pB.X)
		th_next := math.Atan2(pB_next.Z, pB_next.X)

		thTop_i := math.Atan2(pT.Z, pT.X)
		thTop_next := math.Atan2(pT_next.Z, pT_next.X)

		yB := pB.Y + dy
		yB_next := pB_next.Y + dy
		yT := pT.Y + dy
		yT_next := pT_next.Y + dy

		rB := math.Hypot(pB.X, pB.Z)
		rB_outer := rB + thickness
		rB_next := math.Hypot(pB_next.X, pB_next.Z)
		rB_next_outer := rB_next + thickness

		rT := math.Hypot(pT.X, pT.Z)
		rT_outer := rT + thickness
		rT_next := math.Hypot(pT_next.X, pT_next.Z)
		rT_next_outer := rT_next + thickness

		vIn_bottom := vector3{X: rB * math.Cos(th_i), Y: yB, Z: rB * math.Sin(th_i)}
		vOut_bottom := vector3{X: rB_outer * math.Cos(th_i), Y: yB, Z: rB_outer * math.Sin(th_i)}
		vIn_top := vector3{X: rT * math.Cos(thTop_i), Y: yT, Z: rT * math.Sin(thTop_i)}
		vOut_top := vector3{X: rT_outer * math.Cos(thTop_i), Y: yT, Z: rT_outer * math.Sin(thTop_i)}

		vIn_bottom_next := vector3{X: rB_next * math.Cos(th_next), Y: yB_next, Z: rB_next * math.Sin(th_next)}
		vOut_bottom_next := vector3{X: rB_next_outer * math.Cos(th_next), Y: yB_next, Z: rB_next_outer * math.Sin(th_next)}
		vIn_top_next := vector3{X: rT_next * math.Cos(thTop_next), Y: yT_next, Z: rT_next * math.Sin(thTop_next)}
		vOut_top_next := vector3{X: rT_next_outer * math.Cos(thTop_next), Y: yT_next, Z: rT_next_outer * math.Sin(thTop_next)}

		// 1. Outer Surface (facing outward +R)
		writeFacet(sb, vOut_bottom, vOut_top, vOut_bottom_next)
		writeFacet(sb, vOut_bottom_next, vOut_top, vOut_top_next)

		// 2. Inner Surface (facing inward -R)
		writeFacet(sb, vIn_bottom, vIn_bottom_next, vIn_top)
		writeFacet(sb, vIn_top, vIn_bottom_next, vIn_top_next)

		// 3. Bottom Face (facing downward -Y)
		writeFacet(sb, vIn_bottom, vOut_bottom, vIn_bottom_next)
		writeFacet(sb, vOut_bottom, vOut_bottom_next, vIn_bottom_next)

		// 4. Top Face (facing upward +Y)
		writeFacet(sb, vIn_top, vIn_top_next, vOut_top)
		writeFacet(sb, vOut_top, vIn_top_next, vOut_top_next)
	}
}

type stlLayerConfig struct {
	dx, dy, thetaOffset float64
}

type STLExportPart int

const (
	STLExportAll STLExportPart = iota
	STLExportTopRing
	STLExportOneRing
	STLExportBottomRing
	STLExportVaseTrapezeBaseAndFirstRing
)

func buildBaseRingSTL(resampledBaseBottom, resampledBaseTop []vector3, radialRepetitions int, thetaOffset float64) (massiveBottom, massiveTop []vector3) {
	massiveBottom = make([]vector3, 0, len(resampledBaseBottom)*radialRepetitions)
	massiveTop = make([]vector3, 0, len(resampledBaseTop)*radialRepetitions)

	for k := range radialRepetitions {
		baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
		totalThetaOffset := thetaOffset + baseThetaOffset

		localBottom := rotateCurveSTL(resampledBaseBottom, totalThetaOffset)
		localTop := rotateCurveSTL(resampledBaseTop, totalThetaOffset)

		massiveBottom = append(massiveBottom, localBottom...)
		massiveTop = append(massiveTop, localTop...)
	}
	return massiveBottom, massiveTop
}

func GenerateSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportAll)
}

func GenerateTopRingSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportTopRing)
}

func GenerateOneRingSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportOneRing)
}

func GenerateBottomRingSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportBottomRing)
}

func GenerateVaseTrapezeBaseAndFirstRingSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportVaseTrapezeBaseAndFirstRing)
}

func GenerateVaseTrapezeOneRingSTL(plant *PlantAbstract) string {
	return GenerateSTLWithPart(plant, STLExportOneRing)
}

func GenerateSTLWithPart(plant *PlantAbstract, part STLExportPart) string {
	if plant == nil {
		return ""
	}

	solidName := plant.Name
	switch part {
	case STLExportTopRing:
		solidName = plant.Name + "_top_ring"
	case STLExportOneRing:
		solidName = plant.Name + "_one_ring"
	case STLExportBottomRing:
		solidName = plant.Name + "_bottom_ring"
	case STLExportVaseTrapezeBaseAndFirstRing:
		solidName = plant.Name + "_base_first_ring"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("solid %s\n", solidName))

	sideLength := 0.0
	relativeRadialThickness := 0.0
	relativeVerticalThickness := 0.0
	relativeRotatedTorusSeparation := 0.0
	relativeCuttedStackFloorHeight := 0.0
	rotationRatio := 0.0
	relativeHorizontalRingsHeight := 0.0
	radialRepetitions := 1
	thickness := 5.0

	if (plant.PlantType == TubeVase || plant.PlantType == VaseTrapeze) && plant.TubeVaseAbstract != nil {
		vase := plant.TubeVaseAbstract
		radialRepetitions = max(vase.RadialRepetitions, 1)
		sideLength = plant.RhombusSideLength
		relativeRadialThickness = vase.RelativeRadialThickness
		relativeVerticalThickness = vase.RelativeVerticalThickness
		relativeRotatedTorusSeparation = vase.RelativeRotatedTorusSeparation
		relativeCuttedStackFloorHeight = vase.RelativeCuttedStackFloorHeight
		rotationRatio = vase.RotationRatio
		relativeHorizontalRingsHeight = vase.RelativeHorizontalRingsHeight
		if relativeRadialThickness*sideLength > 0 {
			thickness = relativeRadialThickness * sideLength
		}
	}

	var globalR float64
	{
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
		threeDModulo := max(radialRepetitions, 1)
		globalR = circumference * float64(threeDModulo) / (2 * math.Pi)
	}

	if plant.GrowthCurve2D != nil && plant.TubeVaseAbstract != nil && plant.TubeVaseAbstract.TopGrowthCurve2D != nil &&
		plant.GrowthCurve2D.StartHalfwayArcShapeGrid != nil &&
		plant.TubeVaseAbstract.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid != nil &&
		len(plant.GrowthCurve2D.StartHalfwayArcShapeGrid.StartHalfwayArcShapes) > 0 &&
		len(plant.TubeVaseAbstract.TopGrowthCurve2D.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes) > 0 {

		gc := plant.GrowthCurve2D
		tgc := plant.TubeVaseAbstract.TopGrowthCurve2D

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

		var curvePoints []vector3
		var topCurvePoints []vector3

		appendArcPointsSTL := func(targetCurve *[]vector3, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool) {
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

			steps := 50
			for i := 0; i <= steps; i++ {
				if i == 0 && len(*targetCurve) > 0 {
					continue
				}

				t := float64(i) / float64(steps)
				angle := startAngle + t*(endAngle-startAngle)
				x2d := cx + r*math.Cos(angle)
				y2d := cy + r*math.Sin(angle)

				theta := x2d / globalR
				x3d := globalR * math.Cos(theta)
				z3d := globalR * math.Sin(theta)
				y3d := y2d

				*targetCurve = append(*targetCurve, vector3{X: x3d, Y: y3d, Z: z3d})
			}
		}

		for i := range startArcs {
			sa := startArcs[i]
			appendArcPointsSTL(&curvePoints, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag)

			if i < len(endArcs) {
				ea := endArcs[i]
				appendArcPointsSTL(&curvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		for i := range topStartArcs {
			tsa := topStartArcs[i]
			appendArcPointsSTL(&topCurvePoints, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag)

			if i < len(topEndArcs) {
				ea := topEndArcs[i]
				appendArcPointsSTL(&topCurvePoints, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag)
			}
		}

		if plant.PlantType == VaseTrapeze && plant.TubeVaseAbstract != nil {
			zStart := plant.TubeVaseAbstract.Z_Ribbon
			scaleY := plant.TubeVaseAbstract.RibbonVerticalScale
			if scaleY == 0 {
				scaleY = 1.0
			}
			if len(curvePoints) > 0 {
				y0 := curvePoints[0].Y
				for i := range curvePoints {
					curvePoints[i].Y = zStart + scaleY*(curvePoints[i].Y-y0)
				}
				for i := range topCurvePoints {
					topCurvePoints[i].Y = zStart + scaleY*(topCurvePoints[i].Y-y0)
				}
			}
		}

		if len(curvePoints) > 1 && len(topCurvePoints) > 1 {
			targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := getTargetAnglesSTL(curvePoints, topCurvePoints, 0.5, radialRepetitions)

			expectedDegrees := 360.0
			if radialRepetitions > 1 {
				expectedDegrees = 360.0 / float64(radialRepetitions)
			}

			resampledBaseBottom := resampleCurveAtAnglesSTL(anglesBottom, bottomPoints, targetAngles, expectedDegrees)
			resampledBaseTop := resampleCurveAtAnglesSTL(anglesTop, topPoints, targetAngles, expectedDegrees)

			h_horiz := relativeHorizontalRingsHeight * sideLength

			if plant.PlantType == VaseTrapeze && plant.TubeVaseAbstract != nil {
				writeVaseTrapezeSTL(&sb, plant, resampledBaseBottom, resampledBaseTop, radialRepetitions, part)
			} else if part == STLExportBottomRing {
				massiveBottom, massiveTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
				writeRibbonLayerSTL(&sb, massiveBottom, massiveTop, 0.0, thickness, radialRepetitions)

				if h_horiz > 0 && len(massiveBottom) > 0 {
					minY_bottom := math.MaxFloat64
					for _, p := range massiveBottom {
						if p.Y < minY_bottom {
							minY_bottom = p.Y
						}
					}

					horizBottom := make([]vector3, len(massiveBottom))
					horizTop := make([]vector3, len(massiveBottom))
					for idx, p := range massiveBottom {
						horizBottom[idx] = vector3{X: p.X, Y: minY_bottom, Z: p.Z}
						horizTop[idx] = vector3{X: p.X, Y: minY_bottom + h_horiz, Z: p.Z}
					}
					writeRibbonLayerSTL(&sb, horizBottom, horizTop, 0.0, thickness, radialRepetitions)
				}
			} else if part == STLExportOneRing {
				massiveBottom, massiveTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
				writeRibbonLayerSTL(&sb, massiveBottom, massiveTop, 0.0, thickness, radialRepetitions)
			} else if part == STLExportTopRing {
				massiveBottom, massiveTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
				writeRibbonLayerSTL(&sb, massiveBottom, massiveTop, 0.0, thickness, radialRepetitions)

				if h_horiz > 0 && len(massiveTop) > 0 {
					maxY_top := -math.MaxFloat64
					for _, p := range massiveTop {
						if p.Y > maxY_top {
							maxY_top = p.Y
						}
					}

					horizBottom := make([]vector3, len(massiveTop))
					horizTop := make([]vector3, len(massiveTop))
					for idx, p := range massiveTop {
						horizBottom[idx] = vector3{X: p.X, Y: maxY_top - h_horiz, Z: p.Z}
						horizTop[idx] = vector3{X: p.X, Y: maxY_top, Z: p.Z}
					}
					writeRibbonLayerSTL(&sb, horizBottom, horizTop, 0.0, thickness, radialRepetitions)
				}
			} else {
				var checkedDiagram *TubeVase3DDiagram
				for _, diagram := range plant.TubeVase3DDiagrams {
					if diagram.IsChecked {
						checkedDiagram = diagram
						break
					}
				}

				stackHeight := plant.StackHeight

				var activeShapeRuns [][]stlLayerConfig

				if checkedDiagram != nil {
					if !checkedDiagram.IsHiddenTorusStackShape {
						var growthVectorX, growthVectorY float64
						if plant.GrowthVectorShape != nil {
							growthVectorX = plant.GrowthVectorShape.X
							growthVectorY = plant.GrowthVectorShape.Y
						}
						var vx, vy float64
						if plant.PerpendicularVectorGrid != nil && len(plant.PerpendicularVectorGrid.PerpendicularVectors) > 0 {
							pGrid := plant.PerpendicularVectorGrid
							vFirst := pGrid.PerpendicularVectors[0]
							vx = vFirst.EndX - vFirst.StartX
							vy = vFirst.EndY - vFirst.StartY
							vLen := math.Hypot(vx, vy)
							if vLen == 0 {
								vLen = 1
							}
							vx, vy = vx/vLen, vy/vLen
						}
						verticalThickness := relativeVerticalThickness * sideLength
						rotatedSeparation := relativeRotatedTorusSeparation * sideLength

						var run []stlLayerConfig
						for h := range stackHeight {
							dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
							dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
							thetaOffset := dx / globalR
							run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
						}
						activeShapeRuns = append(activeShapeRuns, run)
					}

					if !checkedDiagram.IsHiddenVerticalTorusStackShape {
						var run []stlLayerConfig
						for h := range stackHeight {
							dx := 0.0
							dy := float64(h) * relativeCuttedStackFloorHeight * sideLength
							thetaOffset := 0.0
							run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
						}
						activeShapeRuns = append(activeShapeRuns, run)
					}

					if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
						dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
						thetaOffset := dx / globalR
						var run []stlLayerConfig
						run = append(run, stlLayerConfig{dx: 0, dy: 0, thetaOffset: 0})
						run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
						activeShapeRuns = append(activeShapeRuns, run)
					}

					if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape && stackHeight > 0 {
						numSteps := stackHeight - 1
						dxs := make([]float64, stackHeight)
						dys := make([]float64, stackHeight)
						dxs[0] = 0.0
						dys[0] = 0.0

						if numSteps > 0 {
							totalProgress := rotationRatio * float64(numSteps)
							var cumDX, cumDY float64
							for k := 1; k <= numSteps; k++ {
								var r_k float64
								kFloat := float64(numSteps - k + 1)
								if totalProgress >= kFloat {
									r_k = 1.0
								} else if totalProgress <= kFloat-1.0 {
									r_k = 0.0
								} else {
									r_k = totalProgress - (kFloat - 1.0)
								}
								stepDX, stepDY, _ := ComputePartiallyGrowthCurveDYForRatio(plant, r_k)
								cumDX += stepDX
								cumDY += stepDY
								dxs[k] = cumDX
								dys[k] = cumDY
							}
						}

						var run []stlLayerConfig
						for h := range stackHeight {
							dx := dxs[h]
							dy := dys[h]
							thetaOffset := dx / globalR
							run = append(run, stlLayerConfig{dx: dx, dy: dy, thetaOffset: thetaOffset})
						}
						activeShapeRuns = append(activeShapeRuns, run)
					}
				}

				if len(activeShapeRuns) == 0 {
					var run []stlLayerConfig
					for h := range stackHeight {
						dy := float64(h) * relativeCuttedStackFloorHeight * sideLength
						run = append(run, stlLayerConfig{dx: 0, dy: dy, thetaOffset: 0})
					}
					activeShapeRuns = append(activeShapeRuns, run)
				}

				for _, run := range activeShapeRuns {
					for h, cfg := range run {
						massiveBottom, massiveTop := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, cfg.thetaOffset)
						writeRibbonLayerSTL(&sb, massiveBottom, massiveTop, cfg.dy, thickness, radialRepetitions)

						if h_horiz > 0 && h == 0 && len(massiveBottom) > 0 {
							minY_bottom := math.MaxFloat64
							for _, p := range massiveBottom {
								yVal := p.Y + cfg.dy
								if yVal < minY_bottom {
									minY_bottom = yVal
								}
							}

							horizBottom := make([]vector3, len(massiveBottom))
							horizTop := make([]vector3, len(massiveBottom))
							for idx, p := range massiveBottom {
								horizBottom[idx] = vector3{X: p.X, Y: minY_bottom - cfg.dy, Z: p.Z}
								horizTop[idx] = vector3{X: p.X, Y: (minY_bottom + h_horiz) - cfg.dy, Z: p.Z}
							}
							writeRibbonLayerSTL(&sb, horizBottom, horizTop, cfg.dy, thickness, radialRepetitions)
						}

						if h_horiz > 0 && h == len(run)-1 && len(massiveTop) > 0 {
							maxY_top := -math.MaxFloat64
							for _, p := range massiveTop {
								yVal := p.Y + cfg.dy
								if yVal > maxY_top {
									maxY_top = yVal
								}
							}

							horizBottom := make([]vector3, len(massiveTop))
							horizTop := make([]vector3, len(massiveTop))
							for idx, p := range massiveTop {
								horizBottom[idx] = vector3{X: p.X, Y: (maxY_top - h_horiz) - cfg.dy, Z: p.Z}
								horizTop[idx] = vector3{X: p.X, Y: maxY_top - cfg.dy, Z: p.Z}
							}
							writeRibbonLayerSTL(&sb, horizBottom, horizTop, cfg.dy, thickness, radialRepetitions)
						}
					}
				}
			}
		}
	}

	sb.WriteString(fmt.Sprintf("endsolid %s\n", solidName))

	return sb.String()
}

func writeQuadStripSTL(sb *strings.Builder, curveA, curveB []vector3) {
	M := min(len(curveA), len(curveB))
	if M < 3 {
		return
	}
	for i := range M {
		nextI := (i + 1) % M
		a0 := curveA[i]
		a1 := curveA[nextI]
		b0 := curveB[i]
		b1 := curveB[nextI]

		writeFacet(sb, a0, b0, b1)
		writeFacet(sb, a0, b1, a1)
	}
}

func writeVaseTrapezeSTL(
	sb *strings.Builder,
	plant *PlantAbstract,
	resampledBaseBottom, resampledBaseTop []vector3,
	radialRepetitions int,
	part STLExportPart,
) {
	if plant.TubeVaseAbstract == nil {
		return
	}
	vase := plant.TubeVaseAbstract

	trapezeBottomCurve, trapezeTopCurve := buildBaseRingSTL(resampledBaseBottom, resampledBaseTop, radialRepetitions, 0.0)
	if len(trapezeBottomCurve) < 3 || len(trapezeTopCurve) < 3 {
		return
	}

	scaleY := vase.RibbonVerticalScale
	if scaleY == 0 {
		scaleY = 1.0
	}

	p1H := vase.Z_Ribbon + vase.TopPlaneHeight*scaleY
	bottomPlaneHeight := vase.BottomPlaneHeight
	if part == STLExportVaseTrapezeBaseAndFirstRing || part == STLExportAll {
		bottomPlaneHeight = vase.BaseRingBottomPlaneHeight
	}
	p2H := vase.Z_Ribbon + bottomPlaneHeight*scaleY
	projAngleRad := -vase.ProjectionAngle * math.Pi / 180.0

	projectCurve := func(srcCurve []vector3, planeHeight float64, dy float64) []vector3 {
		var projCurve []vector3
		for _, pt := range srcCurve {
			origTheta := math.Atan2(pt.Z, pt.X)
			r := math.Hypot(pt.X, pt.Z)
			deltaY := (planeHeight - dy) - pt.Y
			rProj := r + deltaY*math.Tan(projAngleRad)
			newPt := vector3{
				X: rProj * math.Cos(origTheta),
				Y: planeHeight,
				Z: rProj * math.Sin(origTheta),
			}
			if len(projCurve) > 0 {
				prev := projCurve[len(projCurve)-1]
				if math.Hypot(newPt.X-prev.X, newPt.Z-prev.Z) < 1e-3 {
					continue
				}
			}
			projCurve = append(projCurve, newPt)
		}
		if len(projCurve) > 2 {
			first := projCurve[0]
			last := projCurve[len(projCurve)-1]
			if math.Hypot(last.X-first.X, last.Z-first.Z) < 1e-3 {
				projCurve = projCurve[:len(projCurve)-1]
			}
		}
		return projCurve
	}

	cTopP1 := projectCurve(trapezeTopCurve, p1H, 0)
	cBottomP1 := projectCurve(trapezeBottomCurve, p1H, 0)
	cTopP2 := projectCurve(trapezeTopCurve, p2H, 0)
	cBottomP2 := projectCurve(trapezeBottomCurve, p2H, 0)

	carvedParam := vase.CarvedOutTopRingsParameter
	if carvedParam < 0 {
		carvedParam = 0
	}
	if carvedParam > 1 {
		carvedParam = 1
	}

	computeCarvedOutCurve := func(cTop, cBottom []vector3, param float64) []vector3 {
		M := min(len(cTop), len(cBottom))
		if M == 0 {
			return nil
		}

		rMin := math.MaxFloat64
		rMax := -math.MaxFloat64
		for i := range M {
			pt := cTop[i]
			r := math.Hypot(pt.X, pt.Z)
			if r < rMin {
				rMin = r
			}
			if r > rMax {
				rMax = r
			}
		}

		rRange := rMax - rMin
		carvedCurve := make([]vector3, 0, M)
		for i := range M {
			ptTop := cTop[i]
			ptBottom := cBottom[i]
			r := math.Hypot(ptTop.X, ptTop.Z)

			coeff := 0.0
			if rRange > 1e-6 {
				coeff = param * (rMax - r) / rRange
			}

			carvedCurve = append(carvedCurve, vector3{
				X: (1.0-coeff)*ptTop.X + coeff*ptBottom.X,
				Y: (1.0-coeff)*ptTop.Y + coeff*ptBottom.Y,
				Z: (1.0-coeff)*ptTop.Z + coeff*ptBottom.Z,
			})
		}
		return carvedCurve
	}

	cCarvedTopP1 := computeCarvedOutCurve(cTopP1, cTopP2, carvedParam)
	cCarvedBottomP1 := computeCarvedOutCurve(cBottomP1, cBottomP2, carvedParam)

	M := min(len(cCarvedTopP1), len(cCarvedBottomP1), len(cTopP2), len(cBottomP2))
	if M < 3 {
		return
	}
	cCarvedTopP1 = cCarvedTopP1[:M]
	cCarvedBottomP1 = cCarvedBottomP1[:M]
	cTopP1 = cTopP1[:M]
	cBottomP1 = cBottomP1[:M]
	cTopP2 = cTopP2[:M]
	cBottomP2 = cBottomP2[:M]

	var sumRTop, sumRBottom float64
	for i := range M {
		sumRTop += math.Hypot(cTopP2[i].X, cTopP2[i].Z)
		sumRBottom += math.Hypot(cBottomP2[i].X, cBottomP2[i].Z)
	}

	cOuterP1 := cCarvedTopP1
	cInnerP1 := cCarvedBottomP1
	cOuterP2 := cTopP2
	cInnerP2 := cBottomP2
	cOuterUncarvedP1 := cTopP1
	if sumRBottom > sumRTop {
		cOuterP1 = cCarvedBottomP1
		cInnerP1 = cCarvedTopP1
		cOuterP2 = cBottomP2
		cInnerP2 = cTopP2
		cOuterUncarvedP1 = cBottomP1
	}

	pBottomH := p2H
	floorY := 0.0

	includeBasePlate := (part == STLExportVaseTrapezeBaseAndFirstRing || part == STLExportAll) && pBottomH > floorY

	// 1. Plane 1 annular face (top of ring, normal +Y)
	writeQuadStripSTL(sb, cOuterP1, cInnerP1)
	// 2. Outer wall of ring (normal +r)
	writeQuadStripSTL(sb, cOuterP2, cOuterP1)
	// 3. Inner wall of ring (normal -r)
	writeQuadStripSTL(sb, cInnerP1, cInnerP2)

	if !includeBasePlate {
		// Close Plane 2 annular face (bottom of ring, normal -Y)
		writeQuadStripSTL(sb, cInnerP2, cOuterP2)
		return
	}

	// Base plate connected seamlessly at Plane 2 (cOuterP2 and cInnerP2)
	hPlancher := vase.BasePlateHeight
	if hPlancher < 0 {
		hPlancher = 0
	}
	if hPlancher > (pBottomH - floorY) {
		hPlancher = pBottomH - floorY
	}

	startMag := vase.BulbousStartTangentMagnitude
	endAngleDeg := 30.0
	if vase.BulbousEndAngle != 0.0 {
		endAngleDeg = vase.BulbousEndAngle
	}
	endMag := vase.BulbousEndTangentMagnitude

	S := 24
	if startMag == 0 && endMag == 0 {
		S = 1
	}

	outerPts := make([][]vector3, S+1)
	innerPts := make([][]vector3, S+1)
	for s := range S + 1 {
		outerPts[s] = make([]vector3, M)
		innerPts[s] = make([]vector3, M)
	}

	var sumOuterX, sumOuterZ float64
	var sumInnerX, sumInnerZ float64

	plancherY := floorY + hPlancher
	verticalScale := 0.0
	if pBottomH > floorY {
		verticalScale = (pBottomH - plancherY) / (pBottomH - floorY)
	}

	beta := endAngleDeg * math.Pi / 180.0
	tEndR := -math.Sin(beta)
	tEndY := -math.Cos(beta)
	v1R := endMag * tEndR
	v1Y := endMag * tEndY

	for i := range M {
		ptOut := cOuterP2[i]
		ptOutP1 := cOuterUncarvedP1[i]
		ptIn := cInnerP2[i]

		sumOuterX += ptOut.X
		sumOuterZ += ptOut.Z
		sumInnerX += ptIn.X
		sumInnerZ += ptIn.Z

		r0 := math.Hypot(ptOut.X, ptOut.Z)
		rTop := math.Hypot(ptOutP1.X, ptOutP1.Z)
		theta := math.Atan2(ptOut.Z, ptOut.X)
		cosTheta := math.Cos(theta)
		sinTheta := math.Sin(theta)

		rIn0 := math.Hypot(ptIn.X, ptIn.Z)
		thetaIn := math.Atan2(ptIn.Z, ptIn.X)
		cosThetaIn := math.Cos(thetaIn)
		sinThetaIn := math.Sin(thetaIn)
		radialScale := 1.0
		if r0 > 1e-6 {
			radialScale = rIn0 / r0
		}

		// Vector from Plane 1 to Plane 2 in (r, y)
		deltaR := r0 - rTop
		deltaY := pBottomH - p1H
		lRing := math.Hypot(deltaR, deltaY)
		t0R := 0.0
		t0Y := -1.0
		if lRing > 1e-6 {
			t0R = deltaR / lRing
			t0Y = deltaY / lRing
		}

		v0R := startMag * t0R
		v0Y := startMag * t0Y

		// Cubic Bezier control points in (r, y)
		p0R, p0Y := r0, pBottomH
		p1R, p1Y := r0+v0R/3.0, pBottomH+v0Y/3.0
		p2R, p2Y := r0-v1R/3.0, floorY-v1Y/3.0
		p3R, p3Y := r0, floorY

		for s := 0; s <= S; s++ {
			t := float64(s) / float64(S)
			omt := 1.0 - t
			b0 := omt * omt * omt
			b1 := 3.0 * omt * omt * t
			b2 := 3.0 * omt * t * t
			b3 := t * t * t

			rs := b0*p0R + b1*p1R + b2*p2R + b3*p3R
			ys := b0*p0Y + b1*p1Y + b2*p2Y + b3*p3R*0 + b3*p3Y

			if s == 0 {
				rs = r0
				ys = pBottomH
				outerPts[s][i] = ptOut
			} else if s == S {
				rs = r0
				ys = floorY
				outerPts[s][i] = vector3{X: r0 * cosTheta, Y: floorY, Z: r0 * sinTheta}
			} else {
				outerPts[s][i] = vector3{X: rs * cosTheta, Y: ys, Z: rs * sinTheta}
			}

			rsIn := rs * radialScale
			ysIn := plancherY + (ys-floorY)*verticalScale

			if s == 0 {
				innerPts[s][i] = vector3{X: ptIn.X, Y: pBottomH, Z: ptIn.Z}
			} else if s == S {
				innerPts[s][i] = vector3{X: ptIn.X, Y: plancherY, Z: ptIn.Z}
			} else {
				innerPts[s][i] = vector3{X: rsIn * cosThetaIn, Y: ysIn, Z: rsIn * sinThetaIn}
			}
		}
	}

	plancherCenter := vector3{X: sumInnerX / float64(M), Y: plancherY, Z: sumInnerZ / float64(M)}
	floorCenter := vector3{X: sumOuterX / float64(M), Y: floorY, Z: sumOuterZ / float64(M)}

	for s := range S {
		for i := range M {
			nextI := (i + 1) % M
			// 1. Outer bulbous wall of base plate (normal +r)
			writeFacet(sb, outerPts[s][i], outerPts[s][nextI], outerPts[s+1][nextI])
			writeFacet(sb, outerPts[s][i], outerPts[s+1][nextI], outerPts[s+1][i])

			// 2. Inner curved wall of base plate (normal -r)
			writeFacet(sb, innerPts[s][i], innerPts[s+1][nextI], innerPts[s][nextI])
			writeFacet(sb, innerPts[s][i], innerPts[s+1][i], innerPts[s+1][nextI])
		}
	}

	for i := range M {
		nextI := (i + 1) % M

		// 3. Plancher top disk (normal +Y)
		writeFacet(sb, plancherCenter, innerPts[S][nextI], innerPts[S][i])

		// 4. Floor bottom disk (normal -Y)
		writeFacet(sb, floorCenter, outerPts[S][i], outerPts[S][nextI])
	}
}


