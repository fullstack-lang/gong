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

			stager.generateRibbonLayer(h, dx, dy, thetaOffset, "Torus Continuous", plant, checkedDiagram, curve, topCurve, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenVerticalTorusStackShape {
		for h := 0; h < stackHeight; h++ {
			dx := 0.0
			dy := float64(h) * plant.RelativeCuttedStackFloorHeight * plant.RhombusSideLength
			thetaOffset := 0.0

			stager.generateRibbonLayer(h, dx, dy, thetaOffset, "Vertical Torus Continuous", plant, checkedDiagram, curve, topCurve, thickness, globalR, canvas)
		}
	}

	if !checkedDiagram.IsHiddenPartiallyRotatedTorusShape {
		dx, dy, _ := ComputePartiallyGrowthCurveDY(plant)
		thetaOffset := dx / globalR

		// The Y component of GrowthVectorShape is applied for the overall stack in 2D,
		// but here the dy from ComputePartiallyGrowthCurveDY ALREADY includes the Y-shift
		// required to perfectly rest on the first ribbon (h=0).

		stager.generateRibbonLayer(1, dx, dy, thetaOffset, "Partially Rotated Torus", plant, checkedDiagram, curve, topCurve, thickness, globalR, canvas)
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

			stager.generateRibbonLayer(h, dx, dy, thetaOffset, "Stack Of Partially Rotated Torus", plant, checkedDiagram, curve, topCurve, thickness, globalR, canvas)
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
