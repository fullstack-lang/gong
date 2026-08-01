package models

import (
	"fmt"
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
		!checkedDiagram.IsHiddenKeyHole3DShape &&
		!checkedDiagram.IsHiddenVolumeKey3DShape

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

	targetAngles, anglesBottom, bottomPoints, anglesTop, topPoints := stager.getTargetAngles(curve, topCurve, 0.5, plant.RadialRepetitions)

	expectedDegrees := 360.0
	if plant.RadialRepetitions > 1 {
		expectedDegrees = 360.0 / float64(plant.RadialRepetitions)
	}

	resampledBaseBottom := stager.resampleCurveAtAngles(anglesBottom, bottomPoints, targetAngles, "Base Bottom", expectedDegrees)
	resampledBaseTop := stager.resampleCurveAtAngles(anglesTop, topPoints, targetAngles, "Base Top", expectedDegrees)

	if !checkedDiagram.IsHiddenOriginalPoints3DShape {
		stager.addPointSpheres(curve.Points, "green", canvas, plant.Name+" Original Bottom", 0, len(curve.Points))
		stager.addPointSpheres(topCurve.Points, "orange", canvas, plant.Name+" Original Top", 0, len(topCurve.Points))
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

			stager.generateLayerWithModulo(h, dx, dy, thetaOffset, "Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenVerticalTorusStackShape {
		for h := 0; h < stackHeight; h++ {
			dx := 0.0
			dy := float64(h) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
			thetaOffset := 0.0

			stager.generateLayerWithModulo(h, dx, dy, thetaOffset, "Vertical Torus Continuous", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
		dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
		thetaOffset := dx / globalR

		// The Y component of GrowthVectorShape is applied for the overall stack in 2D,
		// but here the dy from ComputePartiallyGrowthCurveDY ALREADY includes the Y-shift
		// required to perfectly rest on the first ribbon (h=0).

		stager.generateLayerWithModulo(1, dx, dy, thetaOffset, "Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
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

		for h := 0; h < stackHeight; h++ {
			dx := dxs[h]
			dy := dys[h]
			thetaOffset := dx / globalR

			stager.generateLayerWithModulo(h, dx, dy, thetaOffset, "Stack Of Partially Rotated Torus", plant, checkedDiagram, resampledBaseBottom, resampledBaseTop, thickness, globalR, canvas)
		}
	}

	if (!checkedDiagram.IsHiddenKey3DShape || !checkedDiagram.IsHiddenVolumeKey3DShape) && plant.KeyHoleShape != nil && globalR > 0 {
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
				kFloat := float64(numSteps3D - k + 1)
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

		vk_width := plant.WidthKey * plant.RelativeKeySize
		vk_height := plant.HeightKey * plant.RelativeKeySize
		vk_x_left := plant.OffsetKeyX - vk_width/2.0
		vk_x_right := plant.OffsetKeyX + vk_width/2.0
		vk_y_bottom := plant.OffsetKeyY - vk_height/2.0
		vk_y_top := plant.OffsetKeyY + vk_height/2.0

		tubeRadius := globalR * 0.005

		if tubeRadius < 0.2 {
			tubeRadius = 0.2
		}

		threeDModulo := plant.RadialRepetitions

		for h := 0; h < stackH; h++ {
			dx_h := dxs3D[h]
			dy_h := dys3D[h]

			for k := 0; k < threeDModulo; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(threeDModulo)

				if !checkedDiagram.IsHiddenKey3DShape {
					vBL := stager.get3DPtHK(x_left, y_bottom, "BL", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vBR := stager.get3DPtHK(x_right, y_bottom, "BR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vTR := stager.get3DPtHK(x_right, y_top, "TR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vTL := stager.get3DPtHK(x_left, y_top, "TL", dx_h, dy_h, globalR, baseThetaOffset, h, k)

					canvas.Meshs = append(canvas.Meshs,
						stager.createKeyHole3DTubeMesh(fmt.Sprintf("KeyHole-BL-BR-h%d-k%d", h, k), vBL, vBR, tubeRadius),
						stager.createKeyHole3DTubeMesh(fmt.Sprintf("KeyHole-BR-TR-h%d-k%d", h, k), vBR, vTR, tubeRadius),
						stager.createKeyHole3DTubeMesh(fmt.Sprintf("KeyHole-TR-TL-h%d-k%d", h, k), vTR, vTL, tubeRadius),
						stager.createKeyHole3DTubeMesh(fmt.Sprintf("KeyHole-TL-BL-h%d-k%d", h, k), vTL, vBL, tubeRadius),
					)
				}

				if !checkedDiagram.IsHiddenVolumeKey3DShape {
					// Front face at globalR - thickness
					frontR := globalR - thickness
					vF_BL := stager.get3DPtHK(vk_x_left, vk_y_bottom, "F-BL", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_BR := stager.get3DPtHK(vk_x_right, vk_y_bottom, "F-BR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_TR := stager.get3DPtHK(vk_x_right, vk_y_top, "F-TR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vF_TL := stager.get3DPtHK(vk_x_left, vk_y_top, "F-TL", dx_h, dy_h, frontR, baseThetaOffset, h, k)

					// Back face at globalR + 2*thickness
					backR := globalR + 2.0*thickness
					vB_BL := stager.get3DPtHK(vk_x_left, vk_y_bottom, "B-BL", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_BR := stager.get3DPtHK(vk_x_right, vk_y_bottom, "B-BR", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_TR := stager.get3DPtHK(vk_x_right, vk_y_top, "B-TR", dx_h, dy_h, backR, baseThetaOffset, h, k)
					vB_TL := stager.get3DPtHK(vk_x_left, vk_y_top, "B-TL", dx_h, dy_h, backR, baseThetaOffset, h, k)

					// Opaque volume key (darkgrey for a more discrete look)
					color := "darkgrey"
					canvas.Meshs = append(canvas.Meshs, stager.createVolumeKey3DBoxMesh(
						fmt.Sprintf("VolKey-h%d-k%d", h, k),
						vF_BL, vF_BR, vF_TR, vF_TL, vB_BL, vB_BR, vB_TR, vB_TL, color,
					))

					verticalThickness := plant.RelativeVerticalThickness * plant.RhombusSideLength
					append_y_top := vk_y_bottom
					append_y_bottom := vk_y_bottom - verticalThickness

					// Front face of appended volume at globalR - thickness
					vA_F_BL := stager.get3DPtHK(vk_x_left, append_y_bottom, "AF-BL", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_BR := stager.get3DPtHK(vk_x_right, append_y_bottom, "AF-BR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_TR := stager.get3DPtHK(vk_x_right, append_y_top, "AF-TR", dx_h, dy_h, frontR, baseThetaOffset, h, k)
					vA_F_TL := stager.get3DPtHK(vk_x_left, append_y_top, "AF-TL", dx_h, dy_h, frontR, baseThetaOffset, h, k)

					// Back face of appended volume at globalR (which is the end of the internal third part)
					vA_B_BL := stager.get3DPtHK(vk_x_left, append_y_bottom, "AB-BL", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_BR := stager.get3DPtHK(vk_x_right, append_y_bottom, "AB-BR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_TR := stager.get3DPtHK(vk_x_right, append_y_top, "AB-TR", dx_h, dy_h, globalR, baseThetaOffset, h, k)
					vA_B_TL := stager.get3DPtHK(vk_x_left, append_y_top, "AB-TL", dx_h, dy_h, globalR, baseThetaOffset, h, k)

					canvas.Meshs = append(canvas.Meshs, stager.createVolumeKey3DBoxMesh(
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

		topY := floorMinY + float64(plant.StackHeight)*plant.RhombusSideLength
		if plant.StackHeight == 0 {
			topY = floorMinY + 50.0
		}

		// Vertical pole at radius
		pA := (&threejs.Vector3{Name: "Angle0 Pole Base", X: globalR, Y: floorMinY, Z: 0}).Stage(stager.threejsStage)
		pB := (&threejs.Vector3{Name: "Angle0 Pole Top", X: globalR, Y: topY, Z: 0}).Stage(stager.threejsStage)

		crvPole := (&threejs.Curve{
			Name:   "Angle0 Pole Curve",
			Points: []*threejs.Vector3{pA, pB},
		}).Stage(stager.threejsStage)

		tGeomPole := (&threejs.TubeGeometry{
			Name:            "Angle0 Pole TubeGeom",
			Path:            crvPole,
			TubularSegments: 2,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          false,
		}).Stage(stager.threejsStage)

		poleMesh := (&threejs.Mesh{
			Name:         "Angle0 Pole Mesh",
			Position:     threejs.Position{X: 0, Y: 0, Z: 0},
			TubeGeometry: tGeomPole,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Angle0 Pole Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "magenta"},
				Transparent:          true,
				Opacity:              0.2,
			}).Stage(stager.threejsStage),
		}).Stage(stager.threejsStage)

		// Line from center out to 1.5x radius
		pC := (&threejs.Vector3{Name: "Angle0 Axis Start", X: 0, Y: floorMinY, Z: 0}).Stage(stager.threejsStage)
		pD := (&threejs.Vector3{Name: "Angle0 Axis End", X: globalR * 1.5, Y: floorMinY, Z: 0}).Stage(stager.threejsStage)

		crvAxis := (&threejs.Curve{
			Name:   "Angle0 Axis Curve",
			Points: []*threejs.Vector3{pC, pD},
		}).Stage(stager.threejsStage)

		tGeomAxis := (&threejs.TubeGeometry{
			Name:            "Angle0 Axis TubeGeom",
			Path:            crvAxis,
			TubularSegments: 2,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          false,
		}).Stage(stager.threejsStage)

		axisMesh := (&threejs.Mesh{
			Name:         "Angle0 Axis Mesh",
			Position:     threejs.Position{X: 0, Y: 0, Z: 0},
			TubeGeometry: tGeomAxis,
			MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
				Name:                 "Angle0 Axis Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "magenta"},
				Transparent:          true,
				Opacity:              0.2,
			}).Stage(stager.threejsStage),
		}).Stage(stager.threejsStage)

		canvas.Meshs = append(canvas.Meshs, poleMesh, axisMesh)
	}

	stager.addFloorTiles(floorMinY, plant, globalR, canvas)

	stager.threejsStage.Commit()

}
