package models

import (
	"fmt"
	"log"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) ux_3d_plant_diagram() {

	var preservedX, preservedY, preservedZ float64
	var preservedTargetX, preservedTargetY, preservedTargetZ float64
	var preservedFov float64
	var hasPreservedCamera bool

	for cam := range stager.threejsStage.Cameras {
		preservedX = cam.X
		preservedY = cam.Y
		preservedZ = cam.Z
		preservedTargetX = cam.TargetX
		preservedTargetY = cam.TargetY
		preservedTargetZ = cam.TargetZ
		preservedFov = cam.Fov
		hasPreservedCamera = true
		break
	}

	stager.threejsStage.Reset()

	plant := stager.GetCurrentPlant()
	if plant == nil {
		stager.threejsStage.Commit()
		return
	}

	canvas := (&threejs.Canvas{
		Name: "Plant 3D Canvas",
	}).Stage(stager.threejsStage)

	var checkedDiagram *PlantDiagram
	for _, diagram := range plant.PlantDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}

	// lights
	stager.addLights(canvas)

	globalR := stager.computeGlobalRadius(plant)

	stager.preserveCamera(hasPreservedCamera, preservedFov, canvas, preservedX, preservedY, preservedZ, preservedTargetX, preservedTargetY, preservedTargetZ, checkedDiagram, globalR)

	floorMinY := math.MaxFloat64

	// isOne3DShapeVisible is true if any of the shapes are visible
	isOne3DShapeVisible := !checkedDiagram.IsHiddenTorusStackShape &&
		!checkedDiagram.IsHiddenVerticalTorusStackShape &&
		!checkedDiagram.IsHiddenPartiallyRotatedTorusShape &&
		!checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape &&
		!checkedDiagram.IsHiddenPointsAndLines3DShape &&
		!checkedDiagram.IsHiddenKeyHole3DShape

	// Ribbon generated from GrowthCurve2D and TopGrowthCurve2D
	if checkedDiagram == nil || isOne3DShapeVisible || plant.StackHeight == 0 {
		stager.threejsStage.Commit()
		return
	}

	gc := plant.GrowthCurve2D
	tgc := plant.TopGrowthCurve2D

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

	thickness := plant.RelativeRadialThickness * plant.RhombusSideLength
	if thickness == 0 {
		thickness = 5.0
	}

	// We will use BufferGeometry directly, no need for 2D shapes

	curve := (&threejs.Curve{
		Name: "Torus Continuous Curve Base",
	}).Stage(stager.threejsStage)

	topCurve := (&threejs.Curve{
		Name: "Torus Continuous Curve Top",
	}).Stage(stager.threejsStage)

	for i := 0; i < len(startArcs); i++ {
		sa := startArcs[i]
		// Cartesian sweep is the inverse of SVG sweep due to Y-axis mirroring
		stager.appendArcPoints(curve, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag, globalR, &floorMinY)

		if i < len(endArcs) {
			ea := endArcs[i]
			// Cartesian sweep is !ea.SweepFlag.
			// Traversing forwards, so we pass !ea.SweepFlag.
			stager.appendArcPoints(curve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, &floorMinY)
		}
	}

	for i := 0; i < len(topStartArcs); i++ {
		tsa := topStartArcs[i]
		stager.appendArcPoints(topCurve, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag, globalR, &floorMinY)

		if i < len(topEndArcs) {
			ea := topEndArcs[i]
			// Traversing forwards
			stager.appendArcPoints(topCurve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, &floorMinY)
		}
	}

	stackHeight := plant.StackHeight

	generateRibbonLayer := func(h int, dx, dy, thetaOffset float64, baseNamePrefix string) {
		threeDModulo := plant.ThreeDModulo
		if threeDModulo < 1 {
			threeDModulo = 1
		}

		for k := 0; k < threeDModulo; k++ {
			baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(threeDModulo)
			var bottomEdges, topEdges, innerEdges, outerEdges [][2]*threejs.Vector3

			for i := 0; i < len(curve.Points) && i < len(topCurve.Points); i++ {
				p := curve.Points[i]
				pTop := topCurve.Points[i]

				thetaBase := math.Atan2(p.Z, p.X)
				theta := thetaBase + thetaOffset + baseThetaOffset

				thetaBaseTop := math.Atan2(pTop.Z, pTop.X)
				thetaTop := thetaBaseTop + thetaOffset + baseThetaOffset

				yBase := p.Y + dy
				yBaseTop := pTop.Y + dy

				rBase := math.Sqrt(p.X*p.X + p.Z*p.Z)
				rOuter := rBase + thickness

				rBaseTop := math.Sqrt(pTop.X*pTop.X + pTop.Z*pTop.Z)
				rOuterTop := rBaseTop + thickness

				xBL := rBase * math.Cos(theta)
				zBL := rBase * math.Sin(theta)
				yBL := yBase

				xBR := rOuter * math.Cos(theta)
				zBR := rOuter * math.Sin(theta)
				yBR := yBase

				xTL := rBaseTop * math.Cos(thetaTop)
				zTL := rBaseTop * math.Sin(thetaTop)
				yTL := yBaseTop

				xTR := rOuterTop * math.Cos(thetaTop)
				zTR := rOuterTop * math.Sin(thetaTop)
				yTR := yBaseTop

				vBL := (&threejs.Vector3{Name: "BL", X: xBL, Y: yBL, Z: zBL}).Stage(stager.threejsStage)
				vBR := (&threejs.Vector3{Name: "BR", X: xBR, Y: yBR, Z: zBR}).Stage(stager.threejsStage)
				vTL := (&threejs.Vector3{Name: "TL", X: xTL, Y: yTL, Z: zTL}).Stage(stager.threejsStage)
				vTR := (&threejs.Vector3{Name: "TR", X: xTR, Y: yTR, Z: zTR}).Stage(stager.threejsStage)

				bottomEdges = append(bottomEdges, [2]*threejs.Vector3{vBL, vBR})
				topEdges = append(topEdges, [2]*threejs.Vector3{vTL, vTR})
				innerEdges = append(innerEdges, [2]*threejs.Vector3{vBL, vTL})
				outerEdges = append(outerEdges, [2]*threejs.Vector3{vBR, vTR})
			}

			namePrefix := fmt.Sprintf("%s Layer %d", baseNamePrefix, h)

			japanesePaperColor := "#fdf6e3" // Off-white cream color for Washi paper

			bottomFace := stager.createFaceMesh(namePrefix+" Bottom", japanesePaperColor, bottomEdges, false, plant.Transparency)
			topFace := stager.createFaceMesh(namePrefix+" Top", japanesePaperColor, topEdges, true, plant.Transparency)
			innerFace := stager.createFaceMesh(namePrefix+" Inner", japanesePaperColor, innerEdges, true, plant.Transparency)
			outerFace := stager.createFaceMesh(namePrefix+" Outer", japanesePaperColor, outerEdges, false, plant.Transparency)

			canvas.Meshs = append(canvas.Meshs, bottomFace, topFace, innerFace, outerFace)

			outerRadius := 0.1
			innerRadius := outerRadius * 0.85

			bambooColor := "#4a3623" // dark brown

			canvas.Meshs = append(canvas.Meshs,
				stager.createTube(namePrefix+" BottomInner", bambooColor, bottomEdges, true, innerRadius),
				stager.createTube(namePrefix+" BottomOuter", bambooColor, bottomEdges, false, outerRadius),
				stager.createTube(namePrefix+" TopInner", bambooColor, topEdges, true, innerRadius),
				stager.createTube(namePrefix+" TopOuter", bambooColor, topEdges, false, outerRadius),
			)

			if !checkedDiagram.IsHiddenPointsAndLines3DShape && h < stackHeight-1 && (plant.ChosenP1P2PairShape != nil || plant.PxShape != nil) {
				var p1x, p1y, p2x, p2y, pxx, pxy float64
				hasP1P2 := false
				if plant.ChosenP1P2PairShape != nil {
					p1x, p1y = plant.ChosenP1P2PairShape.P1X, plant.ChosenP1P2PairShape.P1Y
					p2x, p2y = plant.ChosenP1P2PairShape.P2X, plant.ChosenP1P2PairShape.P2Y
					pxx, pxy = plant.ChosenP1P2PairShape.PxX, plant.ChosenP1P2PairShape.PxY
					hasP1P2 = true
				} else if plant.PxShape != nil {
					pxx, pxy = plant.PxShape.X, plant.PxShape.Y
				}

				// Recompute Px for layer h based on step h+1's specific rotation ratio
				if plant.StackOfGrowthCurve2DRibbon != nil && len(plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes) > 0 && plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil {
					baseShape := plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes[0]
					circLen := plant.RhombusStuff.PlantCircumferenceShape.Length
					trajOffsetX := plant.RelativeTrajectoryOffsetX * circLen
					trajOffsetY := plant.RelativeTrajectoryOffsetY * circLen

					var r_h1 float64
					if stackHeight > 1 {
						numSteps := stackHeight - 1
						totalProgress := plant.RotationRatio * float64(numSteps)
						kStep := float64(h + 1)
						if totalProgress >= kStep {
							r_h1 = 1.0
						} else if totalProgress <= kStep-1.0 {
							r_h1 = 0.0
						} else {
							r_h1 = totalProgress - (kStep - 1.0)
						}
					} else {
						r_h1 = plant.RotationRatio
					}

					_, dyStep, currentDXStep := ComputePartiallyGrowthCurveDYForRatio(plant, r_h1)
					pxx = baseShape.BottomStartX + currentDXStep + trajOffsetX
					pxy = baseShape.BottomStartY + dyStep + trajOffsetY
				}

				rSurf := globalR

				get3DPt := func(ptX, ptY float64, ptName string) *threejs.Vector3 {
					th := (ptX+dx)/globalR + baseThetaOffset
					return (&threejs.Vector3{
						Name: fmt.Sprintf("%s %s k%d h%d", ptName, namePrefix, k, h),
						X:    rSurf * math.Cos(th),
						Y:    ptY + dy,
						Z:    rSurf * math.Sin(th),
					}).Stage(stager.threejsStage)
				}

				vPx_3d := get3DPt(pxx, pxy, "Px")

				sphereRad := thickness * 0.4
				if sphereRad < 0.3 {
					sphereRad = 0.3
				}

				createPointSphere := func(ptName string, color string, vec *threejs.Vector3) *threejs.Mesh {
					return (&threejs.Mesh{
						Name: fmt.Sprintf("Sphere %s %s k%d h%d", ptName, namePrefix, k, h),
						Position: threejs.Position{
							X: vec.X,
							Y: vec.Y,
							Z: vec.Z,
						},
						SphereGeometry: (&threejs.SphereGeometry{
							Name:           fmt.Sprintf("SphereGeom %s %s k%d h%d", ptName, namePrefix, k, h),
							Radius:         sphereRad,
							WidthSegments:  16,
							HeightSegments: 16,
						}).Stage(stager.threejsStage),
						MeshMaterialBasic: (&threejs.MeshMaterialBasic{
							Name:                 fmt.Sprintf("Material %s %s k%d h%d", ptName, namePrefix, k, h),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
						}).Stage(stager.threejsStage),
					}).Stage(stager.threejsStage)
				}

				createPairTube := func(lineName string, color string, pA, pB *threejs.Vector3) *threejs.Mesh {
					crv := (&threejs.Curve{
						Name:   fmt.Sprintf("Curve %s %s k%d h%d", lineName, namePrefix, k, h),
						Points: []*threejs.Vector3{pA, pB},
					}).Stage(stager.threejsStage)

					tGeom := (&threejs.TubeGeometry{
						Name:            fmt.Sprintf("TubeGeom %s %s k%d h%d", lineName, namePrefix, k, h),
						Path:            crv,
						TubularSegments: 8,
						Radius:          sphereRad * 0.25,
						RadialSegments:  8,
						Closed:          false,
					}).Stage(stager.threejsStage)

					return (&threejs.Mesh{
						Name:         fmt.Sprintf("TubeMesh %s %s k%d h%d", lineName, namePrefix, k, h),
						Position:     threejs.Position{X: 0, Y: 0, Z: 0},
						TubeGeometry: tGeom,
						MeshMaterialBasic: (&threejs.MeshMaterialBasic{
							Name:                 fmt.Sprintf("Material %s %s k%d h%d", lineName, namePrefix, k, h),
							MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
						}).Stage(stager.threejsStage),
					}).Stage(stager.threejsStage)
				}

				sPx := createPointSphere("Px", "purple", vPx_3d)
				canvas.Meshs = append(canvas.Meshs, sPx)

				if hasP1P2 {
					vP1_3d := get3DPt(p1x, p1y, "P1")
					vP2_3d := get3DPt(p2x, p2y, "P2")

					sP1 := createPointSphere("P1", "red", vP1_3d)
					sP2 := createPointSphere("P2", "#8d6e63", vP2_3d)

					tP1Px := createPairTube("P1-Px", "purple", vP1_3d, vPx_3d)
					tP2Px := createPairTube("P2-Px", "purple", vP2_3d, vPx_3d)

					canvas.Meshs = append(canvas.Meshs, sP1, sP2, tP1Px, tP2Px)

					dx1 := vP1_3d.X - vPx_3d.X
					dy1 := vP1_3d.Y - vPx_3d.Y
					dz1 := vP1_3d.Z - vPx_3d.Z
					distP1Px_3d := math.Sqrt(dx1*dx1 + dy1*dy1 + dz1*dz1)

					dx2 := vP2_3d.X - vPx_3d.X
					dy2 := vP2_3d.Y - vPx_3d.Y
					dz2 := vP2_3d.Z - vPx_3d.Z
					distP2Px_3d := math.Sqrt(dx2*dx2 + dy2*dy2 + dz2*dz2)

					distSum_3d := distP1Px_3d + distP2Px_3d

					if h == 0 && k == 0 {
						log.Printf("[3D Distance] %s (Layer %d, Rep %d) | P1-Px: %.4f, P2-Px: %.4f | 3D Sum: %.4f",
							baseNamePrefix, h, k, distP1Px_3d, distP2Px_3d, distSum_3d)
					}
				}
			}
		}
	}

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

		verticalThickness := plant.RelativeVerticalThickness * plant.RhombusSideLength
		rotatedSeparation := plant.RelativeRotatedTorusSeparation * plant.RhombusSideLength

		for h := 0; h < stackHeight; h++ {
			dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
			dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
			thetaOffset := dx / globalR

			generateRibbonLayer(h, dx, dy, thetaOffset, "Torus Continuous")
		}
	}

	if !checkedDiagram.IsHiddenVerticalTorusStackShape {
		for h := 0; h < stackHeight; h++ {
			dx := 0.0
			dy := float64(h) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
			thetaOffset := 0.0

			generateRibbonLayer(h, dx, dy, thetaOffset, "Vertical Torus Continuous")
		}
	}

	if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
		dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
		thetaOffset := dx / globalR

		// The Y component of GrowthVectorShape is applied for the overall stack in 2D,
		// but here the dy from ComputePartiallyGrowthCurveDY ALREADY includes the Y-shift
		// required to perfectly rest on the first ribbon (h=0).

		generateRibbonLayer(1, dx, dy, thetaOffset, "Partially Rotated Torus")
	}

	if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape && stackHeight > 0 {
		numSteps := stackHeight - 1
		dxs := make([]float64, stackHeight)
		dys := make([]float64, stackHeight)
		dxs[0] = 0.0
		dys[0] = 0.0

		if numSteps > 0 {
			totalProgress := plant.RotationRatio * float64(numSteps)
			var cumDX, cumDY float64
			for k := 1; k <= numSteps; k++ {
				var r_k float64
				kFloat := float64(k)
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

		for h := 0; h < stackHeight; h++ {
			dx := dxs[h]
			dy := dys[h]
			thetaOffset := dx / globalR

			generateRibbonLayer(h, dx, dy, thetaOffset, "Stack Of Partially Rotated Torus")
		}
	}

	if !checkedDiagram.IsHiddenKeyHole3DShape && plant.KeyHoleShape != nil && globalR > 0 {
		stackH := stackHeight
		if stackH <= 0 {
			stackH = 1
		}

		dxs3D := make([]float64, stackH)
		dys3D := make([]float64, stackH)

		numSteps3D := stackH - 1
		if numSteps3D > 0 {
			totalProgress := plant.RotationRatio * float64(numSteps3D)
			var cumDX, cumDY float64
			for k := 1; k <= numSteps3D; k++ {
				var r_k float64
				kFloat := float64(k)
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
				dxs3D[k] = cumDX
				dys3D[k] = cumDY
			}
		}

		x_left := plant.OffsetKeyX - plant.WidthKey/2.0
		x_right := plant.OffsetKeyX + plant.WidthKey/2.0
		y_bottom := plant.OffsetKeyY - plant.HeightKey/2.0
		y_top := plant.OffsetKeyY + plant.HeightKey/2.0

		tubeRadius := globalR * 0.005

		if tubeRadius < 0.2 {
			tubeRadius = 0.2
		}

		createKeyHole3DTube := func(tubeName string, pA, pB *threejs.Vector3) *threejs.Mesh {
			crv := (&threejs.Curve{
				Name:   fmt.Sprintf("Curve %s", tubeName),
				Points: []*threejs.Vector3{pA, pB},
			}).Stage(stager.threejsStage)

			tGeom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("TubeGeom %s", tubeName),
				Path:            crv,
				TubularSegments: 8,
				Radius:          tubeRadius,
				RadialSegments:  8,
				Closed:          false,
			}).Stage(stager.threejsStage)

			return (&threejs.Mesh{
				Name:         fmt.Sprintf("TubeMesh %s", tubeName),
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: tGeom,
				MeshMaterialBasic: (&threejs.MeshMaterialBasic{
					Name:                 fmt.Sprintf("Material %s", tubeName),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkred"},
				}).Stage(stager.threejsStage),
			}).Stage(stager.threejsStage)
		}

		threeDModulo := plant.ThreeDModulo
		if threeDModulo < 1 {
			threeDModulo = 1
		}

		for h := 0; h < stackH; h++ {
			dx_h := dxs3D[h]
			dy_h := dys3D[h]

			for k := 0; k < threeDModulo; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(threeDModulo)

				get3DPtHK := func(ptX, ptY float64, ptName string) *threejs.Vector3 {
					th := (ptX+dx_h)/globalR + baseThetaOffset
					return (&threejs.Vector3{
						Name: fmt.Sprintf("KeyHole3D %s h%d k%d", ptName, h, k),
						X:    globalR * math.Cos(th),
						Y:    ptY + dy_h,
						Z:    globalR * math.Sin(th),
					}).Stage(stager.threejsStage)
				}

				vBL := get3DPtHK(x_left, y_bottom, "BL")
				vBR := get3DPtHK(x_right, y_bottom, "BR")
				vTR := get3DPtHK(x_right, y_top, "TR")
				vTL := get3DPtHK(x_left, y_top, "TL")

				canvas.Meshs = append(canvas.Meshs,
					createKeyHole3DTube(fmt.Sprintf("KeyHole-BL-BR-h%d-k%d", h, k), vBL, vBR),
					createKeyHole3DTube(fmt.Sprintf("KeyHole-BR-TR-h%d-k%d", h, k), vBR, vTR),
					createKeyHole3DTube(fmt.Sprintf("KeyHole-TR-TL-h%d-k%d", h, k), vTR, vTL),
					createKeyHole3DTube(fmt.Sprintf("KeyHole-TL-BL-h%d-k%d", h, k), vTL, vBL),
				)
			}
		}
	}

	stager.addFloorTiles(floorMinY, plant, globalR, canvas)

	stager.threejsStage.Commit()

}
