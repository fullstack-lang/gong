package vasestage3d

import (
	"fmt"
	"math"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *ThreeJSStageUpdater) ux_3d_plant_diagram(stager *models.Stager) {

	threejsStage := stager.GetThreejsStage()

	var preservedX, preservedY, preservedZ float64
	var preservedTargetX, preservedTargetY, preservedTargetZ float64
	var preservedFov float64
	var hasPreservedCamera bool

	for cam := range threejsStage.Cameras {
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

	var checkedDiagram *models.Vase3DDiagram
	for _, diagram := range plant.Vase3DDiagrams {
		if diagram.IsChecked {
			checkedDiagram = diagram
			break
		}
	}

	// lights
	u.addLights(stager, canvas)

	globalR := u.computeGlobalRadius(plant)

	u.preserveCamera(stager, hasPreservedCamera, preservedFov, canvas, preservedX, preservedY, preservedZ, preservedTargetX, preservedTargetY, preservedTargetZ, checkedDiagram, globalR)

	floorMinY := math.MaxFloat64

	if checkedDiagram == nil || checkedDiagram == nil {
		threejsStage.Commit()
		return
	}

	// isOne3DShapeVisible is true if any of the shapes are visible
	isOne3DShapeVisible := !checkedDiagram.IsHiddenTorusStackShape ||
		!checkedDiagram.IsHiddenVerticalTorusStackShape ||
		!checkedDiagram.IsHiddenPartiallyRotatedTorusShape ||
		!checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape ||
		!checkedDiagram.IsHiddenPointsAndLines3DShape ||
		!checkedDiagram.IsHiddenKeyHole3DShape ||
		!checkedDiagram.IsHiddenVolumeKey3DShape

	// Ribbon generated from GrowthCurve2D and TopGrowthCurve2D
	if !isOne3DShapeVisible || plant.StackHeight == 0 {
		threejsStage.Commit()
		return
	}

	gc := plant.GrowthCurve2D
	var tgc *models.TopGrowthCurve2D
	if plant.VaseAbstract != nil {
		tgc = plant.VaseAbstract.TopGrowthCurve2D
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
	if plant.PlantType == models.Vase {
		vase := plant.VaseAbstract
		sideLength = plant.RhombusSideLength
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

	// We will use BufferGeometry directly, no need for 2D shapes

	curve := (&threejs.Curve{
		Name: "Torus Continuous Curve Base",
	}).Stage(threejsStage)

	topCurve := (&threejs.Curve{
		Name: "Torus Continuous Curve Top",
	}).Stage(threejsStage)

	for i := 0; i < len(startArcs); i++ {
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

	for i := 0; i < len(topStartArcs); i++ {
		tsa := topStartArcs[i]
		u.appendArcPoints(stager, topCurve, tsa.StartX, tsa.StartY, tsa.EndX, tsa.EndY, tsa.RadiusX, !tsa.SweepFlag, tsa.LargeArcFlag, globalR, &floorMinY)

		if i < len(topEndArcs) {
			ea := topEndArcs[i]
			// Traversing forwards
			u.appendArcPoints(stager, topCurve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, &floorMinY)
		}
	}

	stackHeight := plant.StackHeight

	targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := u.getTargetAngles(curve, topCurve, 0.5, radialRepetitions)

	expectedDegrees := 360.0
	if radialRepetitions > 1 {
		expectedDegrees = 360.0 / float64(radialRepetitions)
	}

	resampledBaseBottom := u.resampleCurveAtAngles(stager, anglesBottom, bottomPoints, targetAngles, "Base Bottom", expectedDegrees)
	resampledBaseTop := u.resampleCurveAtAngles(stager, anglesTop, topPoints, targetAngles, "Base Top", expectedDegrees)

	if !checkedDiagram.IsHiddenOriginalPoints3DShape {
		u.addPointSpheres(stager, curve.Points, "green", canvas, plant.Name+" Original Bottom", 0, len(curve.Points))
		u.addPointSpheres(stager, topCurve.Points, "orange", canvas, plant.Name+" Original Top", 0, len(topCurve.Points))
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

		verticalThickness := relativeVerticalThickness * sideLength
		rotatedSeparation := relativeRotatedTorusSeparation * sideLength

		for h := 0; h < stackHeight; h++ {
			dx := float64(h)*growthVectorX + float64(h)*verticalThickness*vx
			dy := float64(h)*growthVectorY + float64(h)*verticalThickness*vy + float64(h)*rotatedSeparation
			thetaOffset := dx / globalR

			u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenVerticalTorusStackShape {
		for h := 0; h < stackHeight; h++ {
			dx := 0.0
			dy := float64(h) * relativeCuttedStackFloorHeight * sideLength
			thetaOffset := 0.0

			u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Vertical Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
		dx, dy, _ := models.ComputePartiallyGrowthCurveDY(plant)
		thetaOffset := dx / globalR

		// The Y component of GrowthVectorShape is applied for the overall stack in 2D,
		// but here the dy from ComputePartiallyGrowthCurveDY ALREADY includes the Y-shift
		// required to perfectly rest on the first ribbon (h=0).

		u.generateLayerWithModulo(stager, 1, 2, dx, dy, thetaOffset, "Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
	}

	if !checkedDiagram.IsHiddenStackOfPartiallyRotatedTorusShape && stackHeight > 0 {
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

		for h := 0; h < stackHeight; h++ {
			dx := dxs[h]
			dy := dys[h]
			thetaOffset := dx / globalR

			u.generateLayerWithModulo(stager, h, stackHeight, dx, dy, thetaOffset, "Stack Of Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if (!checkedDiagram.IsHiddenKey3DShape || !checkedDiagram.IsHiddenVolumeKey3DShape) && plant.VaseAbstract != nil && plant.VaseAbstract.KeyHoleShape != nil && globalR > 0 {
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

			for k := 0; k < threeDModulo; k++ {
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
			Name:         "Angle0 Pole Mesh",
			Position:     threejs.Position{X: 0, Y: 0, Z: 0},
			TubeGeometry: tGeomPole,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Angle0 Pole Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "magenta"},
				Transparent:          true,
				Opacity:              0.2,
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
			Name:         "Angle0 Axis Mesh",
			Position:     threejs.Position{X: 0, Y: 0, Z: 0},
			TubeGeometry: tGeomAxis,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Angle0 Axis Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "magenta"},
				Transparent:          true,
				Opacity:              0.2,
			}).Stage(threejsStage),
		}).Stage(threejsStage)

		canvas.Meshs = append(canvas.Meshs, poleMesh, axisMesh)
	}

	if !checkedDiagram.IsHiddenTiledFloor3DShape {
		u.addFloorTiles(stager, floorMinY, plant, globalR, canvas)
	}

	threejsStage.Commit()

}
