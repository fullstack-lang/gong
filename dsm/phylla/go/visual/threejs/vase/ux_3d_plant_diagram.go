package vase

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) ux_3d_plant_diagram(stager *models.Stager) {

	threejsStage := stager.GetThreejsStage()
	threejsStage.Reset()

	plant := stager.GetCurrentPlant()
	if plant == nil {
		threejsStage.Commit()
		return
	}

	canvas := (&threejs.Canvas{
		Name: "Plant 3D Canvas",
	}).Stage(threejsStage)

	if u.isRecording {
		canvas.IsWithLastRenderingUpdate = true
		canvas.OnUpdate = func(updatedCanvas *threejs.Canvas) {
			u.onCanvasFrameCaptured(stager, updatedCanvas)
		}
	}

	var checkedDiagram *models.TubeVase3DDiagram
	for _, diagram := range plant.TubeVase3DDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}
	if checkedDiagram == nil && len(plant.TubeVase3DDiagrams) > 0 {
		checkedDiagram = plant.TubeVase3DDiagrams[0]
	}

	globalR := u.computeGlobalRadius(plant)

	u.setupCamera(stager, canvas, checkedDiagram, globalR)

	// lights (placed relative to camera so raking side-light always sculpts the visible undulations)
	u.addLights(stager, canvas)

	floorMinY := math.MaxFloat64

	if checkedDiagram == nil {
		threejsStage.Commit()
		return
	}

	isOne3DShapeVisible := !checkedDiagram.IsHiddenTorusStackShape ||
		!checkedDiagram.IsHiddenVerticalTorusStackShape ||
		!checkedDiagram.IsHiddenPartiallyRotatedTorusShape ||
		!checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape ||
		!checkedDiagram.IsHiddenPointsAndLines3DShape ||
		!checkedDiagram.IsHiddenKeyHole3DShape ||
		!checkedDiagram.IsHiddenKey3DShape ||
		!checkedDiagram.IsHiddenVolumeKey3DShape ||
		!checkedDiagram.IsHiddenTorusEdge3DShape ||
		!checkedDiagram.IsHiddenSampledPoints3DShape ||
		!checkedDiagram.IsHiddenOriginalPoints3DShape ||
		!checkedDiagram.IsHiddenAngle0Shape ||
		!checkedDiagram.IsHiddenTiledFloor3DShape ||
		!checkedDiagram.IsHiddenTopCurvePlane1Shape ||
		!checkedDiagram.IsHiddenBottomCurvePlane1Shape ||
		!checkedDiagram.IsHiddenTopCurvePlane2Shape ||
		!checkedDiagram.IsHiddenBottomCurvePlane2Shape ||
		!checkedDiagram.IsHiddenCarvedOutTopCurvePlane1Shape ||
		!checkedDiagram.IsHiddenCarvedOutBottomCurvePlane1Shape ||
		!checkedDiagram.IsHiddenVaseTrapezeRingShape ||
		!checkedDiagram.IsHiddenCarvedOutVaseTrapezeRingShape ||
		!checkedDiagram.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape ||
		!checkedDiagram.IsHiddenVaseTrapezeBasePlateShape

	// Ribbon generated from GrowthCurve2D and TopGrowthCurve2D
	if !isOne3DShapeVisible || plant.StackHeight == 0 {
		threejsStage.Commit()
		return
	}

	gc := plant.GrowthCurve2D
	var tgc *models.TopGrowthCurve2D
	if plant.TubeVaseAbstract != nil {
		tgc = plant.TubeVaseAbstract.TopGrowthCurve2D
	}

	if gc == nil || gc.StartHalfwayArcShapeGrid == nil || tgc == nil || tgc.TopStartHalfwayArcShapeGrid == nil {
		threejsStage.Commit()
		return
	}

	startArcs := gc.StartHalfwayArcShapeGrid.StartHalfwayArcShapes
	var endArcs []*models.EndHalfwayArcShape
	if gc.EndHalfwayArcShapeGrid != nil {
		endArcs = gc.EndHalfwayArcShapeGrid.EndHalfwayArcShapes
	}

	topStartArcs := tgc.TopStartHalfwayArcShapeGrid.TopStartHalfwayArcShapes

	var topEndArcs []*models.TopEndHalfwayArcShape
	if tgc.TopEndHalfwayArcShapeGrid != nil {
		topEndArcs = tgc.TopEndHalfwayArcShapeGrid.TopEndHalfwayArcShapes
	}

	sideLength := 0.0
	radialRepetitions := 1
	rotRatio := 0.0
	relativeVerticalThickness := 0.0
	relativeRotatedTorusSeparation := 0.0
	relativeCuttedStackFloorHeight := 0.0
	offsetKeyX := 0.0
	offsetKeyY := 0.0
	widthKey := 0.0
	heightKey := 0.0
	relativeKeySize := 0.0
	thickness := 5.0
	if plant.PlantType == models.TubeVase || plant.PlantType == models.VaseTrapeze {
		vase := plant.TubeVaseAbstract
		sideLength = plant.RhombusSideLength
		if vase != nil {
			if vase.RelativeRadialThickness*sideLength > 0 {
				thickness = vase.RelativeRadialThickness * sideLength
			}
			radialRepetitions = vase.RadialRepetitions
			rotRatio = vase.RotationRatio
			relativeVerticalThickness = vase.RelativeVerticalThickness
			relativeRotatedTorusSeparation = vase.RelativeRotatedTorusSeparation
			relativeCuttedStackFloorHeight = vase.RelativeCuttedStackFloorHeight
			offsetKeyX = vase.OffsetKeyX
			offsetKeyY = vase.OffsetKeyY
			widthKey = vase.WidthKey
			heightKey = vase.HeightKey
			relativeKeySize = vase.RelativeKeySize
		}
	}

	// We will use BufferGeometry directly, no need for 2D shapes

	curve := (&threejs.Curve{
		Name: "Torus Continuous Curve Base",
	}).Stage(threejsStage)

	topCurve := (&threejs.Curve{
		Name: "Torus Continuous Curve Top",
	}).Stage(threejsStage)

	for i := range startArcs {
		sa := startArcs[i]
		// Cartesian sweep is the inverse of SVG sweep due to Y-axis mirroring
		u.appendArcPoints(stager, curve, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag, globalR, &floorMinY)

		if i < len(endArcs) {
			ea := endArcs[i]
			// Cartesian sweep is !ea.SweepFlag.
			// Traversing forwards, so we pass !ea.SweepFlag.
			u.appendArcPoints(stager, curve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, &floorMinY)
		}
	}

	for i := range topStartArcs {
		tsa := topStartArcs[i]
		u.appendArcPoints(stager, topCurve, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag, globalR, &floorMinY)

		if i < len(topEndArcs) {
			ea := topEndArcs[i]
			// Traversing forwards
			u.appendArcPoints(stager, topCurve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, &floorMinY)
		}
	}

	if plant.PlantType == models.VaseTrapeze && plant.TubeVaseAbstract != nil {
		zStart := plant.TubeVaseAbstract.Z_Ribbon * sideLength
		scaleY := plant.TubeVaseAbstract.RibbonVerticalScale
		if scaleY == 0 {
			scaleY = 1.0
		}
		if len(curve.Points) > 0 {
			y0 := curve.Points[0].Y
			for _, pt := range curve.Points {
				pt.Y = zStart + scaleY*(pt.Y-y0)
			}
			for _, pt := range topCurve.Points {
				pt.Y = zStart + scaleY*(pt.Y-y0)
			}
		}
		floorMinY = 0.0
	}

	stackHeight := plant.StackHeight

	targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := u.getTargetAngles(curve, topCurve, 0.5, radialRepetitions)

	expectedDegrees := 360.0
	if radialRepetitions > 1 {
		expectedDegrees = 360.0 / float64(radialRepetitions)
	}

	resampledBaseBottom := u.resampleCurveAtAngles(stager, anglesBottom, bottomPoints, targetAngles, "Base Bottom", expectedDegrees)
	resampledBaseTop := u.resampleCurveAtAngles(stager, anglesTop, topPoints, targetAngles, "Base Top", expectedDegrees)

	var trapezeUnrotatedStepDY, trapezeRotatedStepDX, trapezeRotatedStepDY float64
	var hasTrapezeConstantThickness bool
	if plant.PlantType == models.VaseTrapeze && len(resampledBaseBottom.Points) == len(targetAngles) && len(resampledBaseTop.Points) == len(targetAngles) {
		baseBottomY := make([]float64, len(targetAngles))
		for i, pt := range resampledBaseBottom.Points {
			baseBottomY[i] = pt.Y
		}
		if newBottomY, newTopY, unrotDY, rotDX, rotDY, ok := models.ComputeVaseTrapezeConstant3DThickness(plant, globalR, targetAngles, baseBottomY); ok {
			for i := range targetAngles {
				resampledBaseBottom.Points[i].Y = newBottomY[i]
				resampledBaseTop.Points[i].Y = newTopY[i]
			}
			trapezeUnrotatedStepDY = unrotDY
			trapezeRotatedStepDX = rotDX
			trapezeRotatedStepDY = rotDY
			hasTrapezeConstantThickness = true
		}
	}

	if !checkedDiagram.IsHiddenOriginalPoints3DShape {
		u.addPointSpheres(stager, curve.Points, "green", canvas, plant.Name+" Original Bottom", 0, len(curve.Points))
		u.addPointSpheres(stager, topCurve.Points, "orange", canvas, plant.Name+" Original Top", 0, len(topCurve.Points))
	}

	var trapezeBottomCurve, trapezeTopCurve *threejs.Curve
	if !checkedDiagram.IsHiddenTorusStackShape {
		if plant.PlantType == models.VaseTrapeze {
			trapezeBottomCurve, trapezeTopCurve = u.generateLayerWithModulo(stager, 0, 1, 0, 0, 0, "Trapeze 3D Ribbon", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		} else {
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

			for h := range stackHeight {
				dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
				dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
				thetaOffset := dx / globalR

				u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
			}
		}
	}

	if plant.PlantType == models.VaseTrapeze && plant.TubeVaseAbstract != nil {
		tubeRadius := globalR * 0.008
		if tubeRadius < 0.5 {
			tubeRadius = 0.5
		}

		if trapezeBottomCurve == nil || trapezeTopCurve == nil {
			trapezeBottomCurve, trapezeTopCurve = u.computeMassiveCurves(stager, 0, "Trapeze 3D Ribbon", plant, resampledBaseBottom, resampledBaseTop)
		}

		scaleY := plant.TubeVaseAbstract.RibbonVerticalScale
		if scaleY == 0 {
			scaleY = 1.0
		}

		zStart := plant.TubeVaseAbstract.Z_Ribbon * sideLength
		p1H := zStart + plant.TubeVaseAbstract.TopPlaneHeight*sideLength*scaleY
		p2H := zStart + plant.TubeVaseAbstract.BottomPlaneHeight*sideLength*scaleY
		p2H_base := zStart + plant.TubeVaseAbstract.BaseRingBottomPlaneHeight*sideLength*scaleY
		projAngleRad := -plant.TubeVaseAbstract.ProjectionAngle * math.Pi / 180.0

		projectCurve := func(srcCurve *threejs.Curve, planeHeight float64, dy float64, curveName string) *threejs.Curve {
			projCurve := (&threejs.Curve{Name: curveName}).Stage(threejsStage)
			for _, pt := range srcCurve.Points {
				origTheta := math.Atan2(pt.Z, pt.X)
				r := math.Hypot(pt.X, pt.Z)
				deltaY := (planeHeight - dy) - pt.Y
				rProj := r + deltaY*math.Tan(projAngleRad)
				newPt := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Point %.1f", curveName, origTheta*180.0/math.Pi),
					X:    rProj * math.Cos(origTheta),
					Y:    planeHeight,
					Z:    rProj * math.Sin(origTheta),
				}).Stage(threejsStage)

				if len(projCurve.Points) > 0 {
					prev := projCurve.Points[len(projCurve.Points)-1]
					if math.Hypot(newPt.X-prev.X, newPt.Z-prev.Z) < 1e-3 {
						continue
					}
				}
				projCurve.Points = append(projCurve.Points, newPt)
			}
			if len(projCurve.Points) > 2 {
				first := projCurve.Points[0]
				last := projCurve.Points[len(projCurve.Points)-1]
				if math.Hypot(last.X-first.X, last.Z-first.Z) < 1e-3 {
					projCurve.Points = projCurve.Points[:len(projCurve.Points)-1]
				}
			}
			return projCurve
		}

		addProjectedMesh := func(projCurve *threejs.Curve, shapeName string, color string) {
			if len(projCurve.Points) < 2 {
				return
			}
			numSegments := max(len(projCurve.Points), 2)
			tGeom := (&threejs.TubeGeometry{
				Name:            shapeName + " TubeGeom",
				Path:            projCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(threejsStage)

			tMesh := (&threejs.Mesh{
				Name: shapeName + " Mesh",
				X:    0, Y: 0, Z: 0,
				TubeGeometry: tGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:        shapeName + " Material",
					Color:       color,
					Transparent: false,
					Opacity:     1.0,
				}).Stage(threejsStage),
			}).Stage(threejsStage)
			canvas.Meshs = append(canvas.Meshs, tMesh)
		}

		cTopP1 := projectCurve(trapezeTopCurve, p1H, 0, plant.Name+" Top Curve Plane 1")
		cBottomP1 := projectCurve(trapezeBottomCurve, p1H, 0, plant.Name+" Bottom Curve Plane 1")
		cTopP2 := projectCurve(trapezeTopCurve, p2H, 0, plant.Name+" Top Curve Plane 2")
		cBottomP2 := projectCurve(trapezeBottomCurve, p2H, 0, plant.Name+" Bottom Curve Plane 2")
		cTopP2_base := projectCurve(trapezeTopCurve, p2H_base, 0, plant.Name+" Top Curve Base Plane 2")
		cBottomP2_base := projectCurve(trapezeBottomCurve, p2H_base, 0, plant.Name+" Bottom Curve Base Plane 2")

		if !checkedDiagram.IsHiddenTopCurvePlane1Shape {
			addProjectedMesh(cTopP1, plant.Name+" Top Curve Plane 1", "royalblue")
		}
		if !checkedDiagram.IsHiddenBottomCurvePlane1Shape {
			addProjectedMesh(cBottomP1, plant.Name+" Bottom Curve Plane 1", "firebrick")
		}
		if !checkedDiagram.IsHiddenTopCurvePlane2Shape {
			addProjectedMesh(cTopP2, plant.Name+" Top Curve Plane 2", "mediumpurple")
		}
		if !checkedDiagram.IsHiddenBottomCurvePlane2Shape {
			addProjectedMesh(cBottomP2, plant.Name+" Bottom Curve Plane 2", "darkorange")
		}

		ringHeight := p1H - p2H
		basePlateHeight := plant.TubeVaseAbstract.BasePlateHeight * sideLength
		hasRingBottom := plant.TubeVaseAbstract.HasRingBottom && basePlateHeight > 0 && ringHeight > 0
		hRingBottom := 0.0
		if hasRingBottom {
			hRingBottom = min(basePlateHeight, ringHeight)
		}

		deltaYContact := models.ComputeVaseTrapezeContactStepDY(plant, hasTrapezeConstantThickness)
		sWall := models.VaseTrapezeWallSubdivisions(plant)
		isCurvedWall := models.IsVaseTrapezeCurvedWallActive(plant)
		ringDeltaR := func(localY float64) float64 {
			return models.EvalVaseTrapezeRingWallDeltaR(plant, localY, p1H, p2H, deltaYContact)
		}

		offsetCurveAtLocalY := func(src *threejs.Curve, dy float64, deltaRAtLocalY func(float64) float64, curveName string) *threejs.Curve {
			if !isCurvedWall || deltaRAtLocalY == nil || src == nil {
				return src
			}
			out := (&threejs.Curve{Name: curveName}).Stage(threejsStage)
			for i, pt := range src.Points {
				nx, ny, nz := models.OffsetRadialPoint(pt.X, pt.Y, pt.Z, deltaRAtLocalY(pt.Y-dy))
				out.Points = append(out.Points, (&threejs.Vector3{
					Name: fmt.Sprintf("%s Point %d", curveName, i),
					X:    nx, Y: ny, Z: nz,
				}).Stage(threejsStage))
			}
			return out
		}

		interpolateStraightPlancherCurve := func(cP2, cUncarvedP1 *threejs.Curve, u float64, curveName string) *threejs.Curve {
			plancherCurve := (&threejs.Curve{Name: curveName}).Stage(threejsStage)
			M := min(len(cP2.Points), len(cUncarvedP1.Points))
			for i := 0; i < M; i++ {
				p2 := cP2.Points[i]
				p1 := cUncarvedP1.Points[i]
				newPt := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Point %d", curveName, i),
					X:    (1.0-u)*p2.X + u*p1.X,
					Y:    (1.0-u)*p2.Y + u*p1.Y,
					Z:    (1.0-u)*p2.Z + u*p1.Z,
				}).Stage(threejsStage)
				plancherCurve.Points = append(plancherCurve.Points, newPt)
			}
			return plancherCurve
		}

		var cTopStraightPlancher0, cBottomStraightPlancher0 *threejs.Curve
		var cTopPlancher0, cBottomPlancher0 *threejs.Curve
		getStraightPlancherCurves0 := func() (*threejs.Curve, *threejs.Curve) {
			if !hasRingBottom {
				return nil, nil
			}
			if cTopStraightPlancher0 == nil {
				u := hRingBottom / ringHeight
				cTopStraightPlancher0 = interpolateStraightPlancherCurve(cTopP2, cTopP1, u, plant.Name+" Top Straight Curve Plancher")
				cBottomStraightPlancher0 = interpolateStraightPlancherCurve(cBottomP2, cBottomP1, u, plant.Name+" Bottom Straight Curve Plancher")
			}
			return cTopStraightPlancher0, cBottomStraightPlancher0
		}
		getPlancherCurves0 := func() (*threejs.Curve, *threejs.Curve) {
			if !hasRingBottom {
				return nil, nil
			}
			if cTopPlancher0 == nil {
				stTop, stBot := getStraightPlancherCurves0()
				if isCurvedWall {
					cTopPlancher0 = offsetCurveAtLocalY(stTop, 0.0, ringDeltaR, plant.Name+" Top Curve Plancher")
					cBottomPlancher0 = offsetCurveAtLocalY(stBot, 0.0, ringDeltaR, plant.Name+" Bottom Curve Plancher")
				} else {
					cTopPlancher0 = stTop
					cBottomPlancher0 = stBot
				}
			}
			return cTopPlancher0, cBottomPlancher0
		}

		buildRingMesh := func(cTopA, cBottomA, cTopUncarvedP1, cBottomUncarvedP1, cTopStraightB, cBottomStraightB, cTopPlancher, cBottomPlancher *threejs.Curve, dy float64, ringName string, h int) {
			withBottom := cTopPlancher != nil && cBottomPlancher != nil
			M := min(
				len(cTopA.Points),
				len(cBottomA.Points),
				len(cTopUncarvedP1.Points),
				len(cBottomUncarvedP1.Points),
				len(cTopStraightB.Points),
				len(cBottomStraightB.Points),
			)
			if withBottom {
				M = min(M, len(cTopPlancher.Points), len(cBottomPlancher.Points))
			}
			if M < 3 {
				return
			}
			volGeom := (&threejs.BufferGeometry{
				Name: fmt.Sprintf("%s %s BufferGeometry", plant.Name, ringName),
			}).Stage(threejsStage)

			addQuadStrip := func(curveA, curveB *threejs.Curve, stripName string, reverse bool) {
				baseIdx := len(volGeom.Vertices)
				for i := 0; i < M; i++ {
					ptA := curveA.Points[i]
					ptB := curveB.Points[i]
					vA := (&threejs.Vector3{
						Name: fmt.Sprintf("%s %s %s A %d", plant.Name, ringName, stripName, i),
						X:    ptA.X, Y: ptA.Y, Z: ptA.Z,
					}).Stage(threejsStage)
					vB := (&threejs.Vector3{
						Name: fmt.Sprintf("%s %s %s B %d", plant.Name, ringName, stripName, i),
						X:    ptB.X, Y: ptB.Y, Z: ptB.Z,
					}).Stage(threejsStage)
					volGeom.Vertices = append(volGeom.Vertices, vA, vB)
				}

				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					idxA := baseIdx + 2*i
					idxB := baseIdx + 2*i + 1
					idxANext := baseIdx + 2*nextI
					idxBNext := baseIdx + 2*nextI + 1

					v1_1, v2_1, v3_1 := idxA, idxB, idxBNext
					v1_2, v2_2, v3_2 := idxA, idxBNext, idxANext
					if reverse {
						v2_1, v3_1 = v3_1, v2_1
						v2_2, v3_2 = v3_2, v2_2
					}
					volGeom.Faces = append(volGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s %s T1 %d", plant.Name, ringName, stripName, i),
							V1:   v1_1, V2: v2_1, V3: v3_1,
						}).Stage(threejsStage),
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s %s T2 %d", plant.Name, ringName, stripName, i),
							V1:   v1_2, V2: v2_2, V3: v3_2,
						}).Stage(threejsStage),
					)
				}
			}

			addSubdividedRingWall := func(cBotCurved, cTopCurved, cStraightP2, cStraightP1 *threejs.Curve, stripName string, inward bool) {
				if sWall <= 1 {
					if inward {
						addQuadStrip(cTopCurved, cBotCurved, stripName, false)
					} else {
						addQuadStrip(cBotCurved, cTopCurved, stripName, false)
					}
					return
				}
				yContactGlobal := (p1H - deltaYContact) + dy
				idxGrid := make([][]int, sWall+1)
				for s := 0; s <= sWall; s++ {
					idxGrid[s] = make([]int, M)
					for i := 0; i < M; i++ {
						var vx, vy, vz float64
						if s == 0 {
							pt := cBotCurved.Points[i]
							vx, vy, vz = pt.X, pt.Y, pt.Z
						} else if s == sWall {
							pt := cTopCurved.Points[i]
							vx, vy, vz = pt.X, pt.Y, pt.Z
						} else {
							yBot := cBotCurved.Points[i].Y
							yTop := cTopCurved.Points[i].Y
							y := models.RingWallSampleY(s, sWall, yBot, yTop, yContactGlobal)
							u := 0.0
							if math.Abs(ringHeight) > 1e-9 {
								u = (y - (p2H + dy)) / ringHeight
							}
							p2 := cStraightP2.Points[i]
							p1 := cStraightP1.Points[i]
							sx := (1.0-u)*p2.X + u*p1.X
							sz := (1.0-u)*p2.Z + u*p1.Z
							vx, vy, vz = models.OffsetRadialPoint(sx, y, sz, ringDeltaR(y-dy))
						}
						idxGrid[s][i] = len(volGeom.Vertices)
						volGeom.Vertices = append(volGeom.Vertices, (&threejs.Vector3{
							Name: fmt.Sprintf("%s %s %s s%d %d", plant.Name, ringName, stripName, s, i),
							X:    vx, Y: vy, Z: vz,
						}).Stage(threejsStage))
					}
				}
				for s := 0; s < sWall; s++ {
					for i := 0; i < M; i++ {
						nextI := (i + 1) % M
						var v1_1, v2_1, v3_1, v1_2, v2_2, v3_2 int
						if inward {
							v1_1, v2_1, v3_1 = idxGrid[s+1][i], idxGrid[s][i], idxGrid[s][nextI]
							v1_2, v2_2, v3_2 = idxGrid[s+1][i], idxGrid[s][nextI], idxGrid[s+1][nextI]
						} else {
							v1_1, v2_1, v3_1 = idxGrid[s][i], idxGrid[s+1][i], idxGrid[s+1][nextI]
							v1_2, v2_2, v3_2 = idxGrid[s][i], idxGrid[s+1][nextI], idxGrid[s][nextI]
						}
						volGeom.Faces = append(volGeom.Faces,
							(&threejs.Triangle{
								Name: fmt.Sprintf("%s %s %s T1 s%d %d", plant.Name, ringName, stripName, s, i),
								V1:   v1_1, V2: v2_1, V3: v3_1,
							}).Stage(threejsStage),
							(&threejs.Triangle{
								Name: fmt.Sprintf("%s %s %s T2 s%d %d", plant.Name, ringName, stripName, s, i),
								V1:   v1_2, V2: v2_2, V3: v3_2,
							}).Stage(threejsStage),
						)
					}
				}
			}

			var sumRTop, sumRBottom float64
			for i := 0; i < M; i++ {
				ptTop := cTopStraightB.Points[i]
				ptBot := cBottomStraightB.Points[i]
				sumRTop += math.Hypot(ptTop.X, ptTop.Z)
				sumRBottom += math.Hypot(ptBot.X, ptBot.Z)
			}
			cOuterA := cTopA
			cInnerA := cBottomA
			cOuterUncarvedP1 := cTopUncarvedP1
			cInnerUncarvedP1 := cBottomUncarvedP1
			cOuterStraightB := cTopStraightB
			cInnerStraightB := cBottomStraightB
			cInnerPlancher := cBottomPlancher
			if sumRBottom > sumRTop {
				cOuterA = cBottomA
				cInnerA = cTopA
				cOuterUncarvedP1 = cBottomUncarvedP1
				cInnerUncarvedP1 = cTopUncarvedP1
				cOuterStraightB = cBottomStraightB
				cInnerStraightB = cTopStraightB
				cInnerPlancher = cTopPlancher
			}
			cOuterB := offsetCurveAtLocalY(cOuterStraightB, dy, ringDeltaR, fmt.Sprintf("%s %s Outer P2 Curved", plant.Name, ringName))
			cInnerB := offsetCurveAtLocalY(cInnerStraightB, dy, ringDeltaR, fmt.Sprintf("%s %s Inner P2 Curved", plant.Name, ringName))

			// 1. Plane 1 annular face (top of ring, normal +Y)
			addQuadStrip(cOuterA, cInnerA, "Plane 1 Face", false)

			// 2. Outer wall from Plane 2 to Plane 1 (normal +r)
			addSubdividedRingWall(cOuterB, cOuterA, cOuterStraightB, cOuterUncarvedP1, "Outer Wall Face", false)

			if withBottom {
				// 3. Inner wall from Plane 1 down to plancher (normal -r)
				addSubdividedRingWall(cInnerPlancher, cInnerA, cInnerStraightB, cInnerUncarvedP1, "Inner Wall Face", true)

				var sumOuterX, sumOuterY, sumOuterZ float64
				var sumInnerX, sumInnerY, sumInnerZ float64
				idxCapInnerPlancher := make([]int, M)
				idxCapOuterFloor := make([]int, M)
				for i := 0; i < M; i++ {
					ptIn := cInnerPlancher.Points[i]
					sumInnerX += ptIn.X
					sumInnerY += ptIn.Y
					sumInnerZ += ptIn.Z
					idxCapInnerPlancher[i] = len(volGeom.Vertices)
					volGeom.Vertices = append(volGeom.Vertices, (&threejs.Vector3{
						Name: fmt.Sprintf("%s %s Cap Inner Plancher %d", plant.Name, ringName, i),
						X:    ptIn.X, Y: ptIn.Y, Z: ptIn.Z,
					}).Stage(threejsStage))

					ptOut := cOuterB.Points[i]
					sumOuterX += ptOut.X
					sumOuterY += ptOut.Y
					sumOuterZ += ptOut.Z
					idxCapOuterFloor[i] = len(volGeom.Vertices)
					volGeom.Vertices = append(volGeom.Vertices, (&threejs.Vector3{
						Name: fmt.Sprintf("%s %s Cap Outer Floor %d", plant.Name, ringName, i),
						X:    ptOut.X, Y: ptOut.Y, Z: ptOut.Z,
					}).Stage(threejsStage))
				}

				idxPlancherCenter := len(volGeom.Vertices)
				volGeom.Vertices = append(volGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Plancher Center", plant.Name, ringName),
					X:    sumInnerX / float64(M), Y: sumInnerY / float64(M), Z: sumInnerZ / float64(M),
				}).Stage(threejsStage))

				idxFloorCenter := len(volGeom.Vertices)
				volGeom.Vertices = append(volGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Floor Center", plant.Name, ringName),
					X:    sumOuterX / float64(M), Y: sumOuterY / float64(M), Z: sumOuterZ / float64(M),
				}).Stage(threejsStage))

				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					// 4. Interior plancher top disk (normal +Y)
					volGeom.Faces = append(volGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Plancher Face %d", plant.Name, ringName, i),
							V1:   idxPlancherCenter, V2: idxCapInnerPlancher[nextI], V3: idxCapInnerPlancher[i],
						}).Stage(threejsStage),
					)
					// 5. Exterior bottom floor disk (normal -Y)
					volGeom.Faces = append(volGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Floor Face %d", plant.Name, ringName, i),
							V1:   idxFloorCenter, V2: idxCapOuterFloor[i], V3: idxCapOuterFloor[nextI],
						}).Stage(threejsStage),
					)
				}
			} else {
				// 3. Plane 2 annular face (bottom of ring, normal -Y)
				addQuadStrip(cInnerB, cOuterB, "Plane 2 Face", false)

				// 4. Inner wall from Plane 1 down to Plane 2 (normal -r)
				addSubdividedRingWall(cInnerB, cInnerA, cInnerStraightB, cInnerUncarvedP1, "Inner Wall Face", true)
			}

			colorComb := 1
			if plant.TubeVaseAbstract != nil {
				colorComb = plant.TubeVaseAbstract.ColorCombination
			}
			color := models.GetVaseRingColor(colorComb, h)

			transparency := 0.0
			if plant.TubeVaseAbstract != nil {
				transparency = plant.TubeVaseAbstract.Transparency
			}
			opacity := 1.0 - transparency
			if opacity < 0.0 {
				opacity = 0.0
			}
			if opacity > 1.0 {
				opacity = 1.0
			}

			volMesh := (&threejs.Mesh{
				Name: fmt.Sprintf("%s %s Mesh", plant.Name, ringName),
				X:    0, Y: 0, Z: 0,
				BufferGeometry: volGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:        fmt.Sprintf("%s %s Material", plant.Name, ringName),
					Color:       color,
					Transparent: transparency > 0,
					Opacity:     opacity,
				}).Stage(threejsStage),
			}).Stage(threejsStage)

			canvas.Meshs = append(canvas.Meshs, volMesh)
		}

		if !checkedDiagram.IsHiddenVaseTrapezeRingShape {
			cTopPl, cBottomPl := getPlancherCurves0()
			buildRingMesh(cTopP1, cBottomP1, cTopP1, cBottomP1, cTopP2, cBottomP2, cTopPl, cBottomPl, 0.0, "Vase Trapeze Ring", 0)
		}

		carvedParam := 0.0
		if plant.TubeVaseAbstract != nil {
			carvedParam = plant.TubeVaseAbstract.CarvedOutTopRingsParameter
			if carvedParam < 0 {
				carvedParam = 0
			}
			if carvedParam > 1 {
				carvedParam = 1
			}
		}

		computeCarvedOutCurve := func(cTop, cBottom *threejs.Curve, param, dy float64, deltaRAtLocalY func(float64) float64, curveName string) *threejs.Curve {
			carvedCurve := (&threejs.Curve{Name: curveName}).Stage(threejsStage)
			M := min(len(cTop.Points), len(cBottom.Points))
			if M == 0 {
				return carvedCurve
			}

			rMin := math.MaxFloat64
			rMax := -math.MaxFloat64
			for i := 0; i < M; i++ {
				pt := cTop.Points[i]
				r := math.Hypot(pt.X, pt.Z)
				if r < rMin {
					rMin = r
				}
				if r > rMax {
					rMax = r
				}
			}

			rRange := rMax - rMin
			for i := 0; i < M; i++ {
				ptTop := cTop.Points[i]
				ptBottom := cBottom.Points[i]
				r := math.Hypot(ptTop.X, ptTop.Z)

				coeff := 0.0
				if rRange > 1e-6 {
					coeff = param * (rMax - r) / rRange
				}

				sx := (1.0-coeff)*ptTop.X + coeff*ptBottom.X
				sy := (1.0-coeff)*ptTop.Y + coeff*ptBottom.Y
				sz := (1.0-coeff)*ptTop.Z + coeff*ptBottom.Z
				if isCurvedWall && deltaRAtLocalY != nil {
					sx, sy, sz = models.OffsetRadialPoint(sx, sy, sz, deltaRAtLocalY(sy-dy))
				}

				origTheta := math.Atan2(ptTop.Z, ptTop.X)
				newPt := (&threejs.Vector3{
					Name: fmt.Sprintf("%s Point %.1f", curveName, origTheta*180.0/math.Pi),
					X:    sx,
					Y:    sy,
					Z:    sz,
				}).Stage(threejsStage)
				carvedCurve.Points = append(carvedCurve.Points, newPt)
			}
			return carvedCurve
		}

		var cCarvedTopP1, cCarvedBottomP1 *threejs.Curve
		getCarvedCurves := func() (*threejs.Curve, *threejs.Curve) {
			if cCarvedTopP1 == nil {
				cTopFloor, cBottomFloor := cTopP2, cBottomP2
				if hasRingBottom {
					cTopFloor, cBottomFloor = getStraightPlancherCurves0()
				}
				cCarvedTopP1 = computeCarvedOutCurve(cTopP1, cTopFloor, carvedParam, 0.0, ringDeltaR, plant.Name+" Carved Out Top Curve Plane 1")
				cCarvedBottomP1 = computeCarvedOutCurve(cBottomP1, cBottomFloor, carvedParam, 0.0, ringDeltaR, plant.Name+" Carved Out Bottom Curve Plane 1")
			}
			return cCarvedTopP1, cCarvedBottomP1
		}

		if !checkedDiagram.IsHiddenCarvedOutTopCurvePlane1Shape {
			cTop, _ := getCarvedCurves()
			addProjectedMesh(cTop, plant.Name+" Carved Out Top Curve Plane 1", "deepskyblue")
		}
		if !checkedDiagram.IsHiddenCarvedOutBottomCurvePlane1Shape {
			_, cBottom := getCarvedCurves()
			addProjectedMesh(cBottom, plant.Name+" Carved Out Bottom Curve Plane 1", "coral")
		}
		if !checkedDiagram.IsHiddenCarvedOutVaseTrapezeRingShape {
			cTop, cBottom := getCarvedCurves()
			cTopPl, cBottomPl := getPlancherCurves0()
			buildRingMesh(cTop, cBottom, cTopP1, cBottomP1, cTopP2, cBottomP2, cTopPl, cBottomPl, 0.0, "Carved Out Vase Trapeze Ring", 0)
		}

		if !checkedDiagram.IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape && stackHeight > 0 {
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

			numSteps := stackHeight - 1
			dxs := make([]float64, stackHeight)
			dys := make([]float64, stackHeight)
			dxs[0] = 0.0
			dys[0] = 0.0

			if numSteps > 0 {
				_, unrotatedStepDY, _ := models.ComputePartiallyGrowthCurveDYForRatio(plant, 0.0)
				unrotatedStepDY *= scaleY

				rotatedStepDX := growthVectorX + verticalThickness*vx
				rotatedStepDY := (growthVectorY + verticalThickness*vy + rotatedSeparation) * scaleY
				if hasTrapezeConstantThickness {
					unrotatedStepDY = trapezeUnrotatedStepDY
					rotatedStepDX = trapezeRotatedStepDX
					rotatedStepDY = trapezeRotatedStepDY
				}

				var cumDX, cumDY float64
				for k := 1; k <= numSteps; k++ {
					stepOrder := numSteps - k + 1
					threshold := float64(stepOrder) / float64(stackHeight)
					if rotRatio+1e-9 >= threshold {
						cumDX += rotatedStepDX
						cumDY += rotatedStepDY
					} else {
						cumDY += unrotatedStepDY
					}
					dxs[k] = cumDX
					dys[k] = cumDY
				}
			}

			for h := 1; h < stackHeight; h++ {
				dx := dxs[h]
				dy := dys[h]
				thetaOffset := dx / globalR

				curBottomCurve, curTopCurve := u.computeMassiveCurves(stager, thetaOffset, fmt.Sprintf("Trapeze Rotated Carved Stack h%d", h), plant, resampledBaseBottom, resampledBaseTop)

				ringTopP1 := projectCurve(curTopCurve, p1H+dy, dy, fmt.Sprintf("%s Rotated Carved Stack Top P1 h%d", plant.Name, h))
				ringBottomP1 := projectCurve(curBottomCurve, p1H+dy, dy, fmt.Sprintf("%s Rotated Carved Stack Bottom P1 h%d", plant.Name, h))
				ringTopP2 := projectCurve(curTopCurve, p2H+dy, dy, fmt.Sprintf("%s Rotated Carved Stack Top P2 h%d", plant.Name, h))
				ringBottomP2 := projectCurve(curBottomCurve, p2H+dy, dy, fmt.Sprintf("%s Rotated Carved Stack Bottom P2 h%d", plant.Name, h))

				var ringTopPlancher, ringBottomPlancher *threejs.Curve
				ringTopFloor, ringBottomFloor := ringTopP2, ringBottomP2
				if hasRingBottom {
					uFrac := hRingBottom / ringHeight
					stTopPl := interpolateStraightPlancherCurve(ringTopP2, ringTopP1, uFrac, fmt.Sprintf("%s Rotated Carved Stack Top Straight Plancher h%d", plant.Name, h))
					stBotPl := interpolateStraightPlancherCurve(ringBottomP2, ringBottomP1, uFrac, fmt.Sprintf("%s Rotated Carved Stack Bottom Straight Plancher h%d", plant.Name, h))
					ringTopFloor = stTopPl
					ringBottomFloor = stBotPl
					ringTopPlancher = offsetCurveAtLocalY(stTopPl, dy, ringDeltaR, fmt.Sprintf("%s Rotated Carved Stack Top Plancher h%d", plant.Name, h))
					ringBottomPlancher = offsetCurveAtLocalY(stBotPl, dy, ringDeltaR, fmt.Sprintf("%s Rotated Carved Stack Bottom Plancher h%d", plant.Name, h))
				}

				ringCarvedTopP1 := computeCarvedOutCurve(ringTopP1, ringTopFloor, carvedParam, dy, ringDeltaR, fmt.Sprintf("%s Rotated Stack Carved Top P1 h%d", plant.Name, h))
				ringCarvedBottomP1 := computeCarvedOutCurve(ringBottomP1, ringBottomFloor, carvedParam, dy, ringDeltaR, fmt.Sprintf("%s Rotated Stack Carved Bottom P1 h%d", plant.Name, h))
				buildRingMesh(ringCarvedTopP1, ringCarvedBottomP1, ringTopP1, ringBottomP1, ringTopP2, ringBottomP2, ringTopPlancher, ringBottomPlancher, dy, fmt.Sprintf("Stack Of Rotated Carved Out Vase Trapeze Rings h%d", h), h)
			}
		}

		buildBasePlateMesh := func(cTopA, cBottomA, cTopB, cBottomB *threejs.Curve, shapeName string) {
			M0 := min(len(cTopA.Points), len(cBottomA.Points), len(cTopB.Points), len(cBottomB.Points))
			if M0 < 3 {
				return
			}
			pBottomH := p2H_base
			floorY := floorMinY
			if pBottomH <= floorY {
				return
			}

			var sumRTop, sumRBottom float64
			for i := 0; i < M0; i++ {
				ptTop := cTopB.Points[i]
				ptBot := cBottomB.Points[i]
				sumRTop += math.Hypot(ptTop.X, ptTop.Z)
				sumRBottom += math.Hypot(ptBot.X, ptBot.Z)
			}
			bottomIsOuter := sumRBottom > sumRTop
			cOuter := cTopB
			cInner := cBottomB
			cOuterP1 := cTopA
			cInnerP1 := cBottomA
			if bottomIsOuter {
				cOuter = cBottomB
				cInner = cTopB
				cOuterP1 = cBottomA
				cInnerP1 = cTopA
			}

			baseOuterDeltaR := func(localY float64) float64 {
				dr, _ := models.EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant, localY, p1H, pBottomH, deltaYContact)
				return dr
			}
			baseInnerDeltaR := func(localY float64) float64 {
				return models.EvalVaseTrapezeBaseRingInnerDeltaR(plant, localY, p1H, pBottomH, p2H, deltaYContact, bottomIsOuter)
			}

			cOuterCarvedP1 := computeCarvedOutCurve(cOuterP1, cOuter, carvedParam, 0.0, baseOuterDeltaR, plant.Name+" Base Plate Carved Outer P1")
			cInnerCarvedP1 := computeCarvedOutCurve(cInnerP1, cInner, carvedParam, 0.0, baseInnerDeltaR, plant.Name+" Base Plate Carved Inner P1")
			M := min(M0, len(cOuterCarvedP1.Points), len(cInnerCarvedP1.Points))
			if M < 3 {
				return
			}

			hPlancher := 0.0
			if plant.TubeVaseAbstract != nil {
				hPlancher = plant.TubeVaseAbstract.BasePlateHeight * sideLength
			}
			if hPlancher < 0 {
				hPlancher = 0
			}
			if hPlancher > (pBottomH - floorY) {
				hPlancher = pBottomH - floorY
			}

			startMag := 0.0
			endAngleDeg := 30.0
			endMag := 0.0
			if plant.TubeVaseAbstract != nil {
				startMag = plant.TubeVaseAbstract.BulbousStartTangentMagnitude * sideLength
				if plant.TubeVaseAbstract.BulbousEndAngle != 0.0 {
					endAngleDeg = plant.TubeVaseAbstract.BulbousEndAngle
				}
				endMag = plant.TubeVaseAbstract.BulbousEndTangentMagnitude * sideLength
			}

			S := 24
			if startMag == 0 && endMag == 0 {
				S = 1
			}

			basePlateGeom := (&threejs.BufferGeometry{
				Name: fmt.Sprintf("%s %s BufferGeometry", plant.Name, shapeName),
			}).Stage(threejsStage)

			idxRingOuter := make([][]int, sWall+1)
			idxRingInner := make([][]int, sWall+1)
			baseRingSpan := p1H - pBottomH
			for rw := 0; rw < sWall; rw++ {
				idxRingOuter[rw] = make([]int, M)
				idxRingInner[rw] = make([]int, M)
				for i := 0; i < M; i++ {
					var ox, oy, oz float64
					var outName string
					if rw == 0 {
						ptOutCarved := cOuterCarvedP1.Points[i]
						ox, oy, oz = ptOutCarved.X, ptOutCarved.Y, ptOutCarved.Z
						outName = fmt.Sprintf("%s %s Ring Outer Top %d", plant.Name, shapeName, i)
					} else {
						yTopOut := cOuterCarvedP1.Points[i].Y
						tOut := float64(rw) / float64(sWall)
						oy = (1.0-tOut)*yTopOut + tOut*pBottomH
						uOut := 0.0
						if math.Abs(baseRingSpan) > 1e-9 {
							uOut = (oy - pBottomH) / baseRingSpan
						}
						sxOut := (1.0-uOut)*cOuter.Points[i].X + uOut*cOuterP1.Points[i].X
						szOut := (1.0-uOut)*cOuter.Points[i].Z + uOut*cOuterP1.Points[i].Z
						ox, oy, oz = models.OffsetRadialPoint(sxOut, oy, szOut, baseOuterDeltaR(oy))
						outName = fmt.Sprintf("%s %s Ring Outer rw%d %d", plant.Name, shapeName, rw, i)
					}
					idxRingOuter[rw][i] = len(basePlateGeom.Vertices)
					basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
						Name: outName,
						X:    ox, Y: oy, Z: oz,
					}).Stage(threejsStage))

					var ix, iy, iz float64
					var inName string
					if rw == 0 {
						ptInCarved := cInnerCarvedP1.Points[i]
						ix, iy, iz = ptInCarved.X, ptInCarved.Y, ptInCarved.Z
						inName = fmt.Sprintf("%s %s Ring Inner Top %d", plant.Name, shapeName, i)
					} else {
						yTopIn := cInnerCarvedP1.Points[i].Y
						tIn := float64(rw) / float64(sWall)
						iy = (1.0-tIn)*yTopIn + tIn*pBottomH
						uIn := 0.0
						if math.Abs(baseRingSpan) > 1e-9 {
							uIn = (iy - pBottomH) / baseRingSpan
						}
						sxIn := (1.0-uIn)*cInner.Points[i].X + uIn*cInnerP1.Points[i].X
						szIn := (1.0-uIn)*cInner.Points[i].Z + uIn*cInnerP1.Points[i].Z
						ix, iy, iz = models.OffsetRadialPoint(sxIn, iy, szIn, baseInnerDeltaR(iy))
						inName = fmt.Sprintf("%s %s Ring Inner rw%d %d", plant.Name, shapeName, rw, i)
					}
					idxRingInner[rw][i] = len(basePlateGeom.Vertices)
					basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
						Name: inName,
						X:    ix, Y: iy, Z: iz,
					}).Stage(threejsStage))
				}
			}

			idxOuter := make([][]int, S+1)
			idxInner := make([][]int, S+1)
			for s := range S + 1 {
				idxOuter[s] = make([]int, M)
				idxInner[s] = make([]int, M)
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

			_, dBaseOutDyAtBottom := models.EvalVaseTrapezeBaseRingOuterDeltaRAndSlope(plant, pBottomH, p1H, pBottomH, deltaYContact)

			for i := 0; i < M; i++ {
				ptOut := cOuter.Points[i]
				ptOutP1 := cOuterP1.Points[i]
				ptIn := cInner.Points[i]

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

				// Downward tangent vector at pBottomH in (r, y), C1-continuous with Ring 0 outer wall
				deltaR := r0 - rTop
				deltaY := pBottomH - p1H
				if math.Abs(deltaY) > 1e-6 {
					deltaR += dBaseOutDyAtBottom * deltaY
				}
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
					ys := b0*p0Y + b1*p1Y + b2*p2Y + b3*p3Y

					// Exact match at boundaries
					if s == 0 {
						rs = r0
						ys = pBottomH
					} else if s == S {
						rs = r0
						ys = floorY
					}

					var vName string
					if s == 0 {
						vName = fmt.Sprintf("%s %s Outer Top %d", plant.Name, shapeName, i)
					} else if s == S {
						vName = fmt.Sprintf("%s %s Outer Floor %d", plant.Name, shapeName, i)
					} else {
						vName = fmt.Sprintf("%s %s Outer s%d %d", plant.Name, shapeName, s, i)
					}

					idxOuter[s][i] = len(basePlateGeom.Vertices)
					basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
						Name: vName,
						X:    rs * cosTheta,
						Y:    ys,
						Z:    rs * sinTheta,
					}).Stage(threejsStage))

					rsIn := rs * radialScale
					ysIn := plancherY + (ys-floorY)*verticalScale

					var vInName string
					var inX, inY, inZ float64
					if s == 0 {
						vInName = fmt.Sprintf("%s %s Inner Top %d", plant.Name, shapeName, i)
						inX, inY, inZ = ptIn.X, pBottomH, ptIn.Z
					} else if s == S {
						vInName = fmt.Sprintf("%s %s Inner Plancher %d", plant.Name, shapeName, i)
						inX, inY, inZ = ptIn.X, plancherY, ptIn.Z
					} else {
						vInName = fmt.Sprintf("%s %s Inner s%d %d", plant.Name, shapeName, s, i)
						inX, inY, inZ = rsIn*cosThetaIn, ysIn, rsIn*sinThetaIn
					}

					idxInner[s][i] = len(basePlateGeom.Vertices)
					basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
						Name: vInName,
						X:    inX,
						Y:    inY,
						Z:    inZ,
					}).Stage(threejsStage))
				}
			}

			idxRingOuter[sWall] = idxOuter[0]
			idxRingInner[sWall] = idxInner[0]

			idxPlancherCenter := len(basePlateGeom.Vertices)
			basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s %s Plancher Center", plant.Name, shapeName),
				X:    sumInnerX / float64(M), Y: plancherY, Z: sumInnerZ / float64(M),
			}).Stage(threejsStage))

			idxFloorCenter := len(basePlateGeom.Vertices)
			basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
				Name: fmt.Sprintf("%s %s Floor Center", plant.Name, shapeName),
				X:    sumOuterX / float64(M), Y: floorY, Z: sumOuterZ / float64(M),
			}).Stage(threejsStage))

			for rw := 0; rw < sWall; rw++ {
				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					// Bottom ring (h=0) outer wall (from carved Plane 1 down to pBottomH)
					basePlateGeom.Faces = append(basePlateGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Ring Outer Wall T1 rw%d %d", plant.Name, shapeName, rw, i),
							V1:   idxRingOuter[rw][i], V2: idxRingOuter[rw][nextI], V3: idxRingOuter[rw+1][nextI],
						}).Stage(threejsStage),
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Ring Outer Wall T2 rw%d %d", plant.Name, shapeName, rw, i),
							V1:   idxRingOuter[rw][i], V2: idxRingOuter[rw+1][nextI], V3: idxRingOuter[rw+1][i],
						}).Stage(threejsStage),
					)

					// Bottom ring (h=0) inner wall (from carved Plane 1 down to pBottomH)
					basePlateGeom.Faces = append(basePlateGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Ring Inner Wall T1 rw%d %d", plant.Name, shapeName, rw, i),
							V1:   idxRingInner[rw][i], V2: idxRingInner[rw+1][nextI], V3: idxRingInner[rw][nextI],
						}).Stage(threejsStage),
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Ring Inner Wall T2 rw%d %d", plant.Name, shapeName, rw, i),
							V1:   idxRingInner[rw][i], V2: idxRingInner[rw+1][i], V3: idxRingInner[rw+1][nextI],
						}).Stage(threejsStage),
					)
				}
			}

			for s := 0; s < S; s++ {
				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					// 1. Outer bulbous wall
					basePlateGeom.Faces = append(basePlateGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Outer Wall T1 s%d %d", plant.Name, shapeName, s, i),
							V1:   idxOuter[s][i], V2: idxOuter[s][nextI], V3: idxOuter[s+1][nextI],
						}).Stage(threejsStage),
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Outer Wall T2 s%d %d", plant.Name, shapeName, s, i),
							V1:   idxOuter[s][i], V2: idxOuter[s+1][nextI], V3: idxOuter[s+1][i],
						}).Stage(threejsStage),
					)

					// 2. Inner curved wall (following scaled Bezier curve, normal pointing inward -r)
					basePlateGeom.Faces = append(basePlateGeom.Faces,
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Inner Wall T1 s%d %d", plant.Name, shapeName, s, i),
							V1:   idxInner[s][i], V2: idxInner[s+1][nextI], V3: idxInner[s][nextI],
						}).Stage(threejsStage),
						(&threejs.Triangle{
							Name: fmt.Sprintf("%s %s Inner Wall T2 s%d %d", plant.Name, shapeName, s, i),
							V1:   idxInner[s][i], V2: idxInner[s+1][i], V3: idxInner[s+1][nextI],
						}).Stage(threejsStage),
					)
				}
			}

			idxCapOuterTop := make([]int, M)
			idxCapInnerTop := make([]int, M)
			idxCapInnerFloor := make([]int, M)
			idxCapOuterFloor := make([]int, M)
			for i := 0; i < M; i++ {
				vOutTop := cOuterCarvedP1.Points[i]
				idxCapOuterTop[i] = len(basePlateGeom.Vertices)
				basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Cap Outer Top %d", plant.Name, shapeName, i),
					X:    vOutTop.X, Y: vOutTop.Y, Z: vOutTop.Z,
				}).Stage(threejsStage))

				vInTop := cInnerCarvedP1.Points[i]
				idxCapInnerTop[i] = len(basePlateGeom.Vertices)
				basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Cap Inner Top %d", plant.Name, shapeName, i),
					X:    vInTop.X, Y: vInTop.Y, Z: vInTop.Z,
				}).Stage(threejsStage))

				vInFloor := basePlateGeom.Vertices[idxInner[S][i]]
				idxCapInnerFloor[i] = len(basePlateGeom.Vertices)
				basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Cap Inner Floor %d", plant.Name, shapeName, i),
					X:    vInFloor.X, Y: vInFloor.Y, Z: vInFloor.Z,
				}).Stage(threejsStage))

				vOutFloor := basePlateGeom.Vertices[idxOuter[S][i]]
				idxCapOuterFloor[i] = len(basePlateGeom.Vertices)
				basePlateGeom.Vertices = append(basePlateGeom.Vertices, (&threejs.Vector3{
					Name: fmt.Sprintf("%s %s Cap Outer Floor %d", plant.Name, shapeName, i),
					X:    vOutFloor.X, Y: vOutFloor.Y, Z: vOutFloor.Z,
				}).Stage(threejsStage))
			}

			for i := 0; i < M; i++ {
				nextI := (i + 1) % M

				// 3. Top annular face at pBottomH (normal pointing up +Y):
				basePlateGeom.Faces = append(basePlateGeom.Faces,
					(&threejs.Triangle{
						Name: fmt.Sprintf("%s %s Top Face T1 %d", plant.Name, shapeName, i),
						V1:   idxCapInnerTop[i], V2: idxCapOuterTop[nextI], V3: idxCapOuterTop[i],
					}).Stage(threejsStage),
					(&threejs.Triangle{
						Name: fmt.Sprintf("%s %s Top Face T2 %d", plant.Name, shapeName, i),
						V1:   idxCapInnerTop[i], V2: idxCapInnerTop[nextI], V3: idxCapOuterTop[nextI],
					}).Stage(threejsStage),
				)

				// 4. Plancher top face (normal pointing up +Y):
				basePlateGeom.Faces = append(basePlateGeom.Faces,
					(&threejs.Triangle{
						Name: fmt.Sprintf("%s %s Plancher Face %d", plant.Name, shapeName, i),
						V1:   idxPlancherCenter, V2: idxCapInnerFloor[nextI], V3: idxCapInnerFloor[i],
					}).Stage(threejsStage),
				)

				// 5. Floor bottom face (normal pointing down -Y):
				basePlateGeom.Faces = append(basePlateGeom.Faces,
					(&threejs.Triangle{
						Name: fmt.Sprintf("%s %s Floor Face %d", plant.Name, shapeName, i),
						V1:   idxFloorCenter, V2: idxCapOuterFloor[i], V3: idxCapOuterFloor[nextI],
					}).Stage(threejsStage),
				)
			}

			colorComb := 1
			transparency := 0.0
			if plant.TubeVaseAbstract != nil {
				colorComb = plant.TubeVaseAbstract.ColorCombination
				transparency = plant.TubeVaseAbstract.Transparency
			}
			color := models.GetVaseRingColor(colorComb, 0)
			opacity := 1.0 - transparency
			if opacity < 0.0 {
				opacity = 0.0
			}
			if opacity > 1.0 {
				opacity = 1.0
			}

			basePlateMesh := (&threejs.Mesh{
				Name: fmt.Sprintf("%s %s Mesh", plant.Name, shapeName),
				X:    0, Y: 0, Z: 0,
				BufferGeometry: basePlateGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:        fmt.Sprintf("%s %s Material", plant.Name, shapeName),
					Color:       color,
					Transparent: transparency > 0,
					Opacity:     opacity,
				}).Stage(threejsStage),
			}).Stage(threejsStage)

			canvas.Meshs = append(canvas.Meshs, basePlateMesh)
		}

		if !checkedDiagram.IsHiddenVaseTrapezeBasePlateShape {
			buildBasePlateMesh(cTopP1, cBottomP1, cTopP2_base, cBottomP2_base, "Vase Trapeze Base Plate")
		}
	}

	if !checkedDiagram.IsHiddenVerticalTorusStackShape && plant.PlantType != models.VaseTrapeze {
		for h := range stackHeight {
			dx := 0.0
			dy := float64(h) * relativeCuttedStackFloorHeight * sideLength
			thetaOffset := 0.0

			u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Vertical Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape && plant.PlantType != models.VaseTrapeze {
		dx, dy, _ := models.ComputePartiallyGrowthCurveDY(plant)
		thetaOffset := dx / globalR

		// The Y component of GrowthVectorShape is applied for the overall stack in 2D,
		// but here the dy from ComputePartiallyGrowthCurveDY ALREADY includes the Y-shift
		// required to perfectly rest on the first ribbon (h=0).

		u.generateLayerWithModulo(stager, 1, 2, dx, dy, thetaOffset, "Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
	}

	if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape && stackHeight > 0 && plant.PlantType != models.VaseTrapeze {
		numSteps := stackHeight - 1
		dxs := make([]float64, stackHeight)
		dys := make([]float64, stackHeight)
		dxs[0] = 0.0
		dys[0] = 0.0

		if numSteps > 0 {
			totalProgress := rotRatio * float64(numSteps)
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
				stepDX, stepDY, _ := models.ComputePartiallyGrowthCurveDYForRatio(plant, r_k)
				cumDX += stepDX
				cumDY += stepDY
				dxs[k] = cumDX
				dys[k] = cumDY
			}
		}

		for h := range stackHeight {
			dx := dxs[h]
			dy := dys[h]
			thetaOffset := dx / globalR

			u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Stack Of Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if (!checkedDiagram.IsHiddenKey3DShape || !checkedDiagram.IsHiddenVolumeKey3DShape) && plant.PlantType != models.VaseTrapeze && plant.TubeVaseAbstract != nil && plant.TubeVaseAbstract.KeyHoleShape != nil && globalR > 0 {
		stackH := stackHeight
		if stackH <= 0 {
			stackH = 1
		}

		dxs3D := make([]float64, stackH)
		dys3D := make([]float64, stackH)

		numSteps3D := stackH - 1
		if numSteps3D > 0 {
			totalProgress := rotRatio * float64(numSteps3D)
			var cumDX, cumDY float64
			for k := 1; k <= numSteps3D; k++ {
				var r_k float64
				kFloat := float64(numSteps3D - k + 1)
				if totalProgress >= kFloat {
					r_k = 1.0
				} else if totalProgress <= kFloat-1.0 {
					r_k = 0.0
				} else {
					r_k = totalProgress - (kFloat - 1.0)
				}
				stepDX, stepDY, _ := models.ComputePartiallyGrowthCurveDYForRatio(plant, r_k)
				cumDX += stepDX
				cumDY += stepDY
				dxs3D[k] = cumDX
				dys3D[k] = cumDY
			}
		}

		x_left := offsetKeyX - widthKey/2.0
		x_right := offsetKeyX + widthKey/2.0
		y_bottom := offsetKeyY - heightKey/2.0
		y_top := offsetKeyY + heightKey/2.0

		vk_width := widthKey * relativeKeySize
		vk_height := heightKey * relativeKeySize
		vk_x_left := offsetKeyX - vk_width/2.0
		vk_x_right := offsetKeyX + vk_width/2.0
		vk_y_bottom := offsetKeyY - vk_height/2.0
		vk_y_top := offsetKeyY + vk_height/2.0

		tubeRadius := globalR * 0.005

		if tubeRadius < 0.2 {
			tubeRadius = 0.2
		}

		threeDModulo := radialRepetitions

		for h := 0; h < stackH; h++ {
			// Do not dig holes in the first ring
			if h == 0 {
				continue
			}

			dx_h := dxs3D[h]
			dy_h := dys3D[h]

			for k := range threeDModulo {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(threeDModulo)

				if !checkedDiagram.IsHiddenKey3DShape {
					vBL := u.get3DPtHK(stager, x_left, y_bottom, "BL", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vBR := u.get3DPtHK(stager, x_right, y_bottom, "BR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vTR := u.get3DPtHK(stager, x_right, y_top, "TR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vTL := u.get3DPtHK(stager, x_left, y_top, "TL", dx_h, dy_h, globalR, baseThetaOffset, h, k)

					canvas.Meshs = append(canvas.Meshs,
						u.createKeyHole3DTubeMesh(stager, fmt.Sprintf("KeyHole-BL-BR-h%d-k%d", h, k), vBL, vBR, tubeRadius),
						u.createKeyHole3DTubeMesh(stager, fmt.Sprintf("KeyHole-BR-TR-h%d-k%d", h, k), vBR, vTR, tubeRadius),
						u.createKeyHole3DTubeMesh(stager, fmt.Sprintf("KeyHole-TR-TL-h%d-k%d", h, k), vTR, vTL, tubeRadius),
						u.createKeyHole3DTubeMesh(stager, fmt.Sprintf("KeyHole-TL-BL-h%d-k%d", h, k), vTL, vBL, tubeRadius),
					)
				}

				if !checkedDiagram.IsHiddenVolumeKey3DShape {
					// Front face at globalR - thickness
					frontR := globalR - thickness
					vF_BL := u.get3DPtHK(stager, vk_x_left, vk_y_bottom, "F-BL", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_BR := u.get3DPtHK(stager, vk_x_right, vk_y_bottom, "F-BR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_TR := u.get3DPtHK(stager, vk_x_right, vk_y_top, "F-TR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_TL := u.get3DPtHK(stager, vk_x_left, vk_y_top, "F-TL", dx_h, dy_h, frontR, baseThetaOffset, h, k)

					// Back face at globalR + 2*thickness
					backR := globalR + 2.0*thickness
					vB_BL := u.get3DPtHK(stager, vk_x_left, vk_y_bottom, "B-BL", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_BR := u.get3DPtHK(stager, vk_x_right, vk_y_bottom, "B-BR", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_TR := u.get3DPtHK(stager, vk_x_right, vk_y_top, "B-TR", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_TL := u.get3DPtHK(stager, vk_x_left, vk_y_top, "B-TL", dx_h, dy_h, backR, baseThetaOffset, h, k)

					// Opaque volume key (darkgrey for a more discrete look)
					color := "darkgrey"
					canvas.Meshs = append(canvas.Meshs, u.createVolumeKey3DBoxMesh(
						stager,
						fmt.Sprintf("VolKey-h%d-k%d", h, k),
						vF_BL, vF_BR, vF_TR, vF_TL, vB_BL, vB_BR, vB_TR, vB_TL, color,
					))

					verticalThickness := relativeVerticalThickness * sideLength
					append_y_top := vk_y_bottom
					append_y_bottom := vk_y_bottom - verticalThickness

					// Front face of appended volume at globalR - thickness
					vA_F_BL := u.get3DPtHK(stager, vk_x_left, append_y_bottom, "AF-BL", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_BR := u.get3DPtHK(stager, vk_x_right, append_y_bottom, "AF-BR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_TR := u.get3DPtHK(stager, vk_x_right, append_y_top, "AF-TR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_TL := u.get3DPtHK(stager, vk_x_left, append_y_top, "AF-TL", dx_h, dy_h, frontR, baseThetaOffset, h, k)

					// Back face of appended volume at globalR (which is the end of the internal third part)
					vA_B_BL := u.get3DPtHK(stager, vk_x_left, append_y_bottom, "AB-BL", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_BR := u.get3DPtHK(stager, vk_x_right, append_y_bottom, "AB-BR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_TR := u.get3DPtHK(stager, vk_x_right, append_y_top, "AB-TR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_TL := u.get3DPtHK(stager, vk_x_left, append_y_top, "AB-TL", dx_h, dy_h, globalR, baseThetaOffset, h, k)

					canvas.Meshs = append(canvas.Meshs, u.createVolumeKey3DBoxMesh(
						stager,
						fmt.Sprintf("VolKeyAppend-h%d-k%d", h, k),
						vA_F_BL, vA_F_BR, vA_F_TR, vA_F_TL, vA_B_BL, vA_B_BR, vA_B_TR, vA_B_TL, color,
					))
				}
			}
		}
	}

	if !checkedDiagram.IsHiddenAngle0Shape {
		tubeRadius := globalR * 0.01
		if tubeRadius < 0.5 {
			tubeRadius = 0.5
		}

		topY := floorMinY + float64(plant.StackHeight)*sideLength
		if plant.StackHeight == 0 {
			topY = floorMinY + 50.0
		}

		// Vertical pole at radius
		pA := (&threejs.Vector3{Name: "Angle0 Pole Base", X: globalR, Y: floorMinY, Z: 0}).Stage(threejsStage)
		pB := (&threejs.Vector3{Name: "Angle0 Pole Top", X: globalR, Y: topY, Z: 0}).Stage(threejsStage)

		crvPole := (&threejs.Curve{
			Name:   "Angle0 Pole Curve",
			Points: []*threejs.Vector3{pA, pB},
		}).Stage(threejsStage)

		tGeomPole := (&threejs.TubeGeometry{
			Name:            "Angle0 Pole TubeGeom",
			Path:            crvPole,
			TubularSegments: 2,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          false,
		}).Stage(threejsStage)

		poleMesh := (&threejs.Mesh{
			Name: "Angle0 Pole Mesh",
			X:    0, Y: 0, Z: 0,
			TubeGeometry: tGeomPole,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:        "Angle0 Pole Material",
				Color:       "magenta",
				Transparent: true,
				Opacity:     0.2,
			}).Stage(threejsStage),
		}).Stage(threejsStage)

		// Line from center out to 1.5x radius
		pC := (&threejs.Vector3{Name: "Angle0 Axis Start", X: 0, Y: floorMinY, Z: 0}).Stage(threejsStage)
		pD := (&threejs.Vector3{Name: "Angle0 Axis End", X: globalR * 1.5, Y: floorMinY, Z: 0}).Stage(threejsStage)

		crvAxis := (&threejs.Curve{
			Name:   "Angle0 Axis Curve",
			Points: []*threejs.Vector3{pC, pD},
		}).Stage(threejsStage)

		tGeomAxis := (&threejs.TubeGeometry{
			Name:            "Angle0 Axis TubeGeom",
			Path:            crvAxis,
			TubularSegments: 2,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          false,
		}).Stage(threejsStage)

		axisMesh := (&threejs.Mesh{
			Name: "Angle0 Axis Mesh",
			X:    0, Y: 0, Z: 0,
			TubeGeometry: tGeomAxis,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:        "Angle0 Axis Material",
				Color:       "magenta",
				Transparent: true,
				Opacity:     0.2,
			}).Stage(threejsStage),
		}).Stage(threejsStage)

		canvas.Meshs = append(canvas.Meshs, poleMesh, axisMesh)
	}

	if !checkedDiagram.IsHiddenTiledFloor3DShape {
		u.addFloorTiles(stager, floorMinY, plant, globalR, canvas)
	}

	threejsStage.Commit()

}
