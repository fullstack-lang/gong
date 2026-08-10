package stoolstage3d

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (u *Stool3DStageUpdater) ux_3d_stool(stager *models.Stager) {
	stool3dStage := stager.GetStool3dStage()
	if stool3dStage == nil {
		return
	}

	var preservedX, preservedY, preservedZ float64
	var preservedTargetX, preservedTargetY, preservedTargetZ float64
	var preservedFov float64
	var hasPreservedCamera bool

	for cam := range stool3dStage.Cameras {
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

	stool3dStage.Reset()

	plant := stager.GetCurrentPlant()
	if plant == nil || plant.PlantType != models.Stool || plant.StoolAbstract == nil {
		stool3dStage.Commit()
		return
	}

	canvas := (&threejs.Canvas{
		Name: "Stool 3D Canvas",
	}).Stage(stool3dStage)

	// Circumference and cylinder radius
	circumference := 10.0
	if plant.RhombusStuff != nil && plant.RhombusStuff.PlantCircumferenceShape != nil && plant.RhombusStuff.PlantCircumferenceShape.Length > 0 {
		circumference = plant.RhombusStuff.PlantCircumferenceShape.Length
	} else if pGrid := plant.PerpendicularVectorGrid; pGrid != nil && len(pGrid.PerpendicularVectors) > 0 {
		first := pGrid.PerpendicularVectors[0]
		last := pGrid.PerpendicularVectors[len(pGrid.PerpendicularVectors)-1]
		circumference = last.StartX - first.StartX
	}
	if circumference <= 0 {
		circumference = 10.0
	}

	radialRepetitions := plant.StoolAbstract.RadialRepetitions
	if radialRepetitions < 1 {
		radialRepetitions = 1
	}

	globalR := circumference * float64(radialRepetitions) / (2 * math.Pi)

	// Directional and ambient lights positioned relative to scene scale
	lightScale := math.Max(globalR, 50.0)
	dirLight1 := (&threejs.DirectionalLight{
		Name:             "Directional Light 1 (Key)",
		Position:         threejs.Position{X: lightScale * 2.0, Y: lightScale * 2.5, Z: lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 1.2},
		IsWithCastShadow: true,
	}).Stage(stool3dStage)

	dirLight2 := (&threejs.DirectionalLight{
		Name:             "Directional Light 2 (Fill)",
		Position:         threejs.Position{X: -lightScale * 2.0, Y: lightScale * 1.5, Z: -lightScale * 2.0},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.6},
		IsWithCastShadow: false,
	}).Stage(stool3dStage)

	dirLight3 := (&threejs.DirectionalLight{
		Name:             "Directional Light 3 (Rim)",
		Position:         threejs.Position{X: 0, Y: lightScale * 3.5, Z: -lightScale * 2.5},
		LightAbstract:    threejs.LightAbstract{Intensity: 0.8},
		IsWithCastShadow: false,
	}).Stage(stool3dStage)

	canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

	ambiantLight := (&threejs.AmbiantLight{
		Name:          "Ambiant Light",
		LightAbstract: threejs.LightAbstract{Intensity: 0.4},
	}).Stage(stool3dStage)
	canvas.AmbiantLight = ambiantLight

	// Find checked diagram
	var checkedDiagram *models.PlantDiagram
	if plant != nil {
		for _, d := range plant.PlantDiagrams {
			if d.IsChecked {
				checkedDiagram = d
				break
			}
		}
	}

	// Camera setup
	if hasPreservedCamera {
		if preservedFov == 0 {
			preservedFov = 50
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: preservedX,
				Y: preservedY,
				Z: preservedZ,
			},
			TargetX: preservedTargetX,
			TargetY: preservedTargetY,
			TargetZ: preservedTargetZ,
			Fov:     preservedFov,
		}).Stage(stool3dStage)
	} else if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && checkedDiagram.StoolDiagram.Rendered3DShape != nil {
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: checkedDiagram.StoolDiagram.Rendered3DShape.ViewX,
				Y: checkedDiagram.StoolDiagram.Rendered3DShape.ViewY,
				Z: checkedDiagram.StoolDiagram.Rendered3DShape.ViewZ,
			},
			TargetX: checkedDiagram.StoolDiagram.Rendered3DShape.TargetX,
			TargetY: checkedDiagram.StoolDiagram.Rendered3DShape.TargetY,
			TargetZ: checkedDiagram.StoolDiagram.Rendered3DShape.TargetZ,
			Fov:     checkedDiagram.StoolDiagram.Rendered3DShape.Fov,
		}).Stage(stool3dStage)
	} else {
		camDist := globalR * 2.5
		if camDist < 30 {
			camDist = 30
		}
		canvas.Camera = (&threejs.Camera{
			Name: "Camera",
			Position: threejs.Position{
				X: camDist,
				Y: camDist * 0.8,
				Z: camDist,
			},
			TargetY: globalR * 0.5,
			Fov:     50,
		}).Stage(stool3dStage)
	}

	canvas.Camera.OnUpdate = func(updatedCamera *threejs.Camera) {
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && checkedDiagram.StoolDiagram.Rendered3DShape != nil {
			checkedDiagram.StoolDiagram.Rendered3DShape.ViewX = updatedCamera.X
			checkedDiagram.StoolDiagram.Rendered3DShape.ViewY = updatedCamera.Y
			checkedDiagram.StoolDiagram.Rendered3DShape.ViewZ = updatedCamera.Z
			checkedDiagram.StoolDiagram.Rendered3DShape.TargetX = updatedCamera.TargetX
			checkedDiagram.StoolDiagram.Rendered3DShape.TargetY = updatedCamera.TargetY
			checkedDiagram.StoolDiagram.Rendered3DShape.TargetZ = updatedCamera.TargetZ
			checkedDiagram.StoolDiagram.Rendered3DShape.Fov = updatedCamera.Fov
			stager.GetStage().CommitWithSuspendedCallbacks()
		}
	}

	floorMinY := math.MaxFloat64

	// Extract GrowthCurve2D arcs from PlantAbstract (StartArcShapeGrid and EndArcShapeGrid)
	if plant.StartArcShapeGrid != nil && len(plant.StartArcShapeGrid.StartArcShapes) > 0 {
		startArcs := plant.StartArcShapeGrid.StartArcShapes
		var endArcs []*models.EndArcShape
		if plant.EndArcShapeGrid != nil {
			endArcs = plant.EndArcShapeGrid.EndArcShapes
		}

		// 1. Build initial continuous 3D curve for 1 base repetition
		baseCurve := (&threejs.Curve{
			Name: "Stool Base Curve",
		}).Stage(stool3dStage)

		for i := 0; i < len(startArcs); i++ {
			sa := startArcs[i]
			u.appendArcPointsStool(stool3dStage, baseCurve, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag, globalR, 0.0, &floorMinY)

			if i < len(endArcs) {
				ea := endArcs[i]
				u.appendArcPointsStool(stool3dStage, baseCurve, ea.EndX, ea.EndY, ea.StartX, ea.StartY, ea.RadiusX, ea.SweepFlag, ea.LargeArcFlag, globalR, 0.0, &floorMinY)
			}
		}

		// 2. Order and unwrap angles
		sortedAngles, sortedPoints := unwrapAngles(baseCurve)

		// 3. Generate target angles spaced every 0.5 degrees across 1 repetition
		degreeInterval := 0.5
		expectedDegrees := 360.0 / float64(radialRepetitions)
		radInterval := degreeInterval * math.Pi / 180.0
		nbPoints := int(math.Round(expectedDegrees / degreeInterval))

		var targetAngles []float64
		for i := 0; i <= nbPoints; i++ {
			targetAngles = append(targetAngles, float64(i)*radInterval)
		}

		// 4. Resample the curve at every 0.5 degree
		resampledBaseCurve := u.resampleCurveAtAngles(stool3dStage, sortedAngles, sortedPoints, targetAngles, "Stool Resampled", expectedDegrees)

		vertScale := plant.StoolAbstract.StoolTorusVerticalScale
		if vertScale <= 0.0 {
			vertScale = 1.0
		}
		for _, pt := range resampledBaseCurve.Points {
			pt.Y = pt.Y * vertScale
		}

		// 5. Calculate tube radius from RelativeTubeDiameter
		relDiameter := plant.StoolAbstract.RelativeTubeDiameter
		if relDiameter <= 0 {
			relDiameter = 0.01
		}
		tubeRadius := (relDiameter * plant.RhombusSideLength) / 2.0
		if tubeRadius <= 0 {
			tubeRadius = 2.0
		}

		transparency := plant.StoolAbstract.Transparency
		opacity := 1.0 - transparency
		if opacity < 0.0 {
			opacity = 0.0
		}
		if opacity > 1.0 {
			opacity = 1.0
		}

		stoolTopHeight := plant.StoolAbstract.RelativeHeight * plant.RhombusSideLength
		torusHeight := plant.StoolAbstract.RelativeHeight3DTorus * plant.RhombusSideLength

		createTorusLayer := func(dx, dy float64, namePrefix string, color string) {
			thetaOffset := dx / globalR

			layerCurve := (&threejs.Curve{
				Name: fmt.Sprintf("%s Curve", namePrefix),
			}).Stage(stool3dStage)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				totalThetaOffset := baseThetaOffset + thetaOffset

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + totalThetaOffset

					layerPtY := pt.Y + dy + torusHeight
					if layerPtY < floorMinY {
						floorMinY = layerPtY
					}

					layerCurve.Points = append(layerCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("%s Point k%d %.1f", namePrefix, k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    layerPtY,
						Z:    r * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}

			numSegments := len(layerCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			tGeom := (&threejs.TubeGeometry{
				Name:            fmt.Sprintf("%s TubeGeom", namePrefix),
				Path:            layerCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			tubeMesh := (&threejs.Mesh{
				Name:         fmt.Sprintf("%s Mesh", namePrefix),
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: tGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 fmt.Sprintf("%s Material", namePrefix),
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, tubeMesh)
		}

		// 6. Base 3D Torus
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenTorus3DShape {
			createTorusLayer(0.0, 0.0, "Stool Base Torus", "darkgreen")
		}

		// 7. Single rotated and elevated torus with growth vector parameters
		var growthVectorX, growthVectorY float64
		if plant.GrowthVectorShape != nil {
			growthVectorX = plant.GrowthVectorShape.X
			growthVectorY = plant.GrowthVectorShape.Y * vertScale
		}

		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedTorusShape {
			createTorusLayer(growthVectorX, growthVectorY, "Stool Partially Rotated Torus", "darkgreen")
		}

		projAngleRad := -plant.StoolAbstract.ProjectionAngle * math.Pi / 180.0

		// 8. Seat Top 2D Projected Curve on the horizontal stool top plane (from Base Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenSeatTopCurveShape {
			seatTopCurve := (&threejs.Curve{
				Name: "Stool Seat Top Curve",
			}).Stage(stool3dStage)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset

					ptY := pt.Y + torusHeight
					deltaY := stoolTopHeight - ptY
					rProj := r + deltaY*math.Tan(projAngleRad)

					seatTopCurve.Points = append(seatTopCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Top Point k%d %.1f", k, newTheta*180.0/math.Pi),
						X:    rProj * math.Cos(newTheta),
						Y:    stoolTopHeight,
						Z:    rProj * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}

			numSegments := len(seatTopCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			stGeom := (&threejs.TubeGeometry{
				Name:            "Stool Seat Top TubeGeom",
				Path:            seatTopCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			stMesh := (&threejs.Mesh{
				Name:         "Stool Seat Top Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: stGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Seat Top Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "royalblue"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, stMesh)
		}

		seatThickness := plant.StoolAbstract.RelativeSeatThickness * plant.RhombusSideLength
		seatBottomHeight := stoolTopHeight - seatThickness

		var rotSeatTopPoints []*threejs.Vector3
		var rotSeatBottomPoints []*threejs.Vector3

		thetaOffset := growthVectorX / globalR

		for k := 0; k < radialRepetitions; k++ {
			baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
			totalThetaOffset := baseThetaOffset + thetaOffset

			for _, pt := range resampledBaseCurve.Points {
				origTheta := math.Atan2(pt.Z, pt.X)
				r := math.Hypot(pt.X, pt.Z)
				newTheta := origTheta + totalThetaOffset

				ptY := pt.Y + growthVectorY + torusHeight

				deltaYTop := stoolTopHeight - ptY
				rProjTop := r + deltaYTop*math.Tan(projAngleRad)

				rotSeatTopPoints = append(rotSeatTopPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Rotated Seat Top Point k%d %.1f", k, newTheta*180.0/math.Pi),
					X:    rProjTop * math.Cos(newTheta),
					Y:    stoolTopHeight,
					Z:    rProjTop * math.Sin(newTheta),
				}).Stage(stool3dStage))

				deltaYBottom := seatBottomHeight - ptY
				rProjBottom := r + deltaYBottom*math.Tan(projAngleRad)

				rotSeatBottomPoints = append(rotSeatBottomPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Rotated Seat Bottom Point k%d %.1f", k, newTheta*180.0/math.Pi),
					X:    rProjBottom * math.Cos(newTheta),
					Y:    seatBottomHeight,
					Z:    rProjBottom * math.Sin(newTheta),
				}).Stage(stool3dStage))
			}
		}

		// 9. Partially Rotated Seat Top 2D Projected Curve (from Partially Rotated Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedSeatTopCurveShape {
			rotSeatTopCurve := (&threejs.Curve{
				Name:   "Stool Partially Rotated Seat Top Curve",
				Points: rotSeatTopPoints,
			}).Stage(stool3dStage)

			numSegments := len(rotSeatTopCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			rotStGeom := (&threejs.TubeGeometry{
				Name:            "Stool Partially Rotated Seat Top TubeGeom",
				Path:            rotSeatTopCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			rotStMesh := (&threejs.Mesh{
				Name:         "Stool Partially Rotated Seat Top Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: rotStGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Partially Rotated Seat Top Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkorange"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, rotStMesh)
		}

		// 10. Seat Bottom 2D Projected Curve on the horizontal stool seat bottom plane (from Base Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenSeatBottomCurveShape {
			seatBottomCurve := (&threejs.Curve{
				Name: "Stool Seat Bottom Curve",
			}).Stage(stool3dStage)

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset

					ptY := pt.Y + torusHeight
					deltaY := seatBottomHeight - ptY
					rProj := r + deltaY*math.Tan(projAngleRad)

					seatBottomCurve.Points = append(seatBottomCurve.Points, (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Bottom Point k%d %.1f", k, newTheta*180.0/math.Pi),
						X:    rProj * math.Cos(newTheta),
						Y:    seatBottomHeight,
						Z:    rProj * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}

			numSegments := len(seatBottomCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			sbGeom := (&threejs.TubeGeometry{
				Name:            "Stool Seat Bottom TubeGeom",
				Path:            seatBottomCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			sbMesh := (&threejs.Mesh{
				Name:         "Stool Seat Bottom Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: sbGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Seat Bottom Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "dodgerblue"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, sbMesh)
		}

		// 11. Rotated Seat Bottom 2D Projected Curve (from Partially Rotated Torus)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedSeatBottomCurveShape {
			rotSeatBottomCurve := (&threejs.Curve{
				Name:   "Stool Rotated Seat Bottom Curve",
				Points: rotSeatBottomPoints,
			}).Stage(stool3dStage)

			numSegments := len(rotSeatBottomCurve.Points)
			if numSegments < 2 {
				numSegments = 2
			}

			rotSbGeom := (&threejs.TubeGeometry{
				Name:            "Stool Rotated Seat Bottom TubeGeom",
				Path:            rotSeatBottomCurve,
				TubularSegments: numSegments,
				Radius:          tubeRadius,
				RadialSegments:  16,
				Closed:          true,
			}).Stage(stool3dStage)

			rotSbMesh := (&threejs.Mesh{
				Name:         "Stool Rotated Seat Bottom Mesh",
				Position:     threejs.Position{X: 0, Y: 0, Z: 0},
				TubeGeometry: rotSbGeom,
				MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
					Name:                 "Stool Rotated Seat Bottom Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "coral"},
					Transparent:          true,
					Opacity:              opacity,
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, rotSbMesh)
		}

		// 11b. 3D Seat Volume Mesh between Rotated Seat Top and Rotated Seat Bottom
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenSeat3DShape {
			if len(rotSeatTopPoints) >= 3 && len(rotSeatTopPoints) == len(rotSeatBottomPoints) {
				N := len(rotSeatTopPoints)
				seatGeom := (&threejs.BufferGeometry{
					Name: "Stool Seat BufferGeometry",
				}).Stage(stool3dStage)

				var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
				for i := 0; i < N; i++ {
					topV := (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Top V %d", i),
						X:    rotSeatTopPoints[i].X,
						Y:    rotSeatTopPoints[i].Y,
						Z:    rotSeatTopPoints[i].Z,
					}).Stage(stool3dStage)
					seatGeom.Vertices = append(seatGeom.Vertices, topV)
					sumTopX += topV.X
					sumTopZ += topV.Z
				}

				for i := 0; i < N; i++ {
					botV := (&threejs.Vector3{
						Name: fmt.Sprintf("Seat Bottom V %d", i),
						X:    rotSeatBottomPoints[i].X,
						Y:    rotSeatBottomPoints[i].Y,
						Z:    rotSeatBottomPoints[i].Z,
					}).Stage(stool3dStage)
					seatGeom.Vertices = append(seatGeom.Vertices, botV)
					sumBottomX += botV.X
					sumBottomZ += botV.Z
				}

				topCenterIdx := len(seatGeom.Vertices)
				topCenterV := (&threejs.Vector3{
					Name: "Seat Top Center",
					X:    sumTopX / float64(N),
					Y:    stoolTopHeight,
					Z:    sumTopZ / float64(N),
				}).Stage(stool3dStage)
				seatGeom.Vertices = append(seatGeom.Vertices, topCenterV)

				botCenterIdx := len(seatGeom.Vertices)
				botCenterV := (&threejs.Vector3{
					Name: "Seat Bottom Center",
					X:    sumBottomX / float64(N),
					Y:    seatBottomHeight,
					Z:    sumBottomZ / float64(N),
				}).Stage(stool3dStage)
				seatGeom.Vertices = append(seatGeom.Vertices, botCenterV)

				// 1. Top face (facing +Y): (topCenter, nextI, i)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Top Face %d", i),
						V1:   topCenterIdx,
						V2:   nextI,
						V3:   i,
					}).Stage(stool3dStage))
				}

				// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					botI := N + i
					botNextI := N + nextI
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Bottom Face %d", i),
						V1:   botCenterIdx,
						V2:   botI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				// 3. Side wall quads between Top and Bottom:
				for i := 0; i < N; i++ {
					nextI := (i + 1) % N
					topI := i
					topNextI := nextI
					botI := N + i
					botNextI := N + nextI

					// Triangle 1: (botI, topI, topNextI)
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Wall T1 %d", i),
						V1:   botI,
						V2:   topI,
						V3:   topNextI,
					}).Stage(stool3dStage))

					// Triangle 2: (botI, topNextI, botNextI)
					seatGeom.Faces = append(seatGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Seat Wall T2 %d", i),
						V1:   botI,
						V2:   topNextI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				seatMesh := (&threejs.Mesh{
					Name:           "Stool Seat 3D Mesh",
					Position:       threejs.Position{X: 0, Y: 0, Z: 0},
					BufferGeometry: seatGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Seat Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "peru"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, seatMesh)
			}
		}

		// 12. Add 3D Sampled Points visualization if toggled on
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenSampledPoints3DShape {
			numPointsPerRep := len(resampledBaseCurve.Points)
			var basePoints []*threejs.Vector3
			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + baseThetaOffset
					basePoints = append(basePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Sampled Point k%d %.1f", k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    pt.Y + torusHeight,
						Z:    r * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}
			u.addPointSpheres(stool3dStage, basePoints, "red", canvas, "Stool Sampled", 0, numPointsPerRep)
		}

		// 13. Add Rotated 3D Sampled Points visualization if toggled on
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenRotatedSampledPoints3DShape {
			numPointsPerRep := len(resampledBaseCurve.Points)
			var rotPoints []*threejs.Vector3
			thetaOffset := growthVectorX / globalR

			for k := 0; k < radialRepetitions; k++ {
				baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)
				totalThetaOffset := baseThetaOffset + thetaOffset

				for _, pt := range resampledBaseCurve.Points {
					origTheta := math.Atan2(pt.Z, pt.X)
					r := math.Hypot(pt.X, pt.Z)
					newTheta := origTheta + totalThetaOffset
					rotPoints = append(rotPoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Rotated Sampled Point k%d %.1f", k, newTheta*180.0/math.Pi),
						X:    r * math.Cos(newTheta),
						Y:    pt.Y + growthVectorY + torusHeight,
						Z:    r * math.Sin(newTheta),
					}).Stage(stool3dStage))
				}
			}
			u.addPointSpheres(stool3dStage, rotPoints, "orange", canvas, "Stool Rotated Sampled", 0, numPointsPerRep)
		}

		// Compute Eye geometry, transitions, and Bézier corners for first repetition (k=0)
		thetaOffset = growthVectorX / globalR
		eyeCriteria := plant.StoolAbstract.RelativeEyeSeparationCriteria * plant.RhombusSideLength * vertScale
		expectedRad := expectedDegrees * math.Pi / 180.0
		dStep := globalR * radInterval
		if dStep <= 0 {
			dStep = 1.0
		}

		getYAtAngle := func(evalAngle float64) float64 {
			if len(resampledBaseCurve.Points) == 0 {
				return 0
			}
			for evalAngle < 0 {
				evalAngle += expectedRad
			}
			for evalAngle > expectedRad {
				evalAngle -= expectedRad
			}
			idx := int(math.Floor(evalAngle / radInterval))
			if idx < 0 {
				idx = 0
			}
			if idx >= len(resampledBaseCurve.Points)-1 {
				return resampledBaseCurve.Points[len(resampledBaseCurve.Points)-1].Y
			}
			t := (evalAngle - float64(idx)*radInterval) / radInterval
			y0 := resampledBaseCurve.Points[idx].Y
			y1 := resampledBaseCurve.Points[idx+1].Y
			return y0 + t*(y1-y0)
		}

		// Identify transitions into and out of eye regions
		numPts := len(resampledBaseCurve.Points)
		inEye := make([]bool, numPts)
		yBaseList := make([]float64, numPts)
		yRotList := make([]float64, numPts)

		for i, pt := range resampledBaseCurve.Points {
			alpha := targetAngles[i]
			yBaseList[i] = pt.Y + torusHeight
			yRotList[i] = getYAtAngle(alpha-thetaOffset) + growthVectorY + torusHeight
			dist := math.Abs(yBaseList[i] - yRotList[i])
			inEye[i] = (dist > eyeCriteria)
		}

		type point2D struct {
			U, Y float64
		}

		evalBezier := func(c0, c1, c2, c3 point2D, t float64) point2D {
			omt := 1.0 - t
			omt2 := omt * omt
			omt3 := omt2 * omt
			t2 := t * t
			t3 := t2 * t
			return point2D{
				U: omt3*c0.U + 3*omt2*t*c1.U + 3*omt*t2*c2.U + t3*c3.U,
				Y: omt3*c0.Y + 3*omt2*t*c1.Y + 3*omt*t2*c2.Y + t3*c3.Y,
			}
		}

		getBezierPoints := func(c0, c1, c2, c3 point2D, cornerName string) []*threejs.Vector3 {
			const numSub = 200
			pts := make([]point2D, numSub+1)
			cumLen := make([]float64, numSub+1)
			totalLen := 0.0
			pts[0] = c0
			cumLen[0] = 0.0

			for k := 1; k <= numSub; k++ {
				t := float64(k) / float64(numSub)
				pts[k] = evalBezier(c0, c1, c2, c3, t)
				du := pts[k].U - pts[k-1].U
				dy := pts[k].Y - pts[k-1].Y
				totalLen += math.Hypot(du, dy)
				cumLen[k] = totalLen
			}

			n := int(math.Round(totalLen / dStep))
			if n < 2 {
				n = 2
			}

			var result []*threejs.Vector3
			for j := 1; j < n; j++ {
				targetDist := float64(j) * (totalLen / float64(n))
				searchIdx := sort.SearchFloat64s(cumLen, targetDist)
				if searchIdx <= 0 {
					searchIdx = 1
				}
				if searchIdx > numSub {
					searchIdx = numSub
				}
				segL := cumLen[searchIdx] - cumLen[searchIdx-1]
				segT := 0.0
				if segL > 0 {
					segT = (targetDist - cumLen[searchIdx-1]) / segL
				}
				uVal := pts[searchIdx-1].U + segT*(pts[searchIdx].U-pts[searchIdx-1].U)
				yVal := pts[searchIdx-1].Y + segT*(pts[searchIdx].Y-pts[searchIdx-1].Y)

				theta := uVal / globalR
				x := globalR * math.Cos(theta)
				z := globalR * math.Sin(theta)

				result = append(result, (&threejs.Vector3{
					Name: fmt.Sprintf("%s Bezier Point %d/%.1f", cornerName, j, theta*180.0/math.Pi),
					X:    x,
					Y:    yVal,
					Z:    z,
				}).Stage(stool3dStage))
			}
			return result
		}

		normalize2D := func(du, dy float64, defU, defY float64) (float64, float64) {
			lenV := math.Hypot(du, dy)
			if lenV > 1e-7 {
				return du / lenV, dy / lenV
			}
			return defU, defY
		}

		controlStrength := plant.StoolAbstract.RelativeEyeCornerControlVectorStrength
		if controlStrength <= 0.0 {
			controlStrength = 0.55
		}

		// Find eye segment endpoints
		iStart := -1
		iEnd := -1
		for i := 0; i < numPts; i++ {
			if inEye[i] {
				if iStart == -1 {
					iStart = i
				}
				iEnd = i
			}
		}

		var leftCornerPts []*threejs.Vector3
		var rightCornerPts []*threejs.Vector3

		if iStart != -1 && iEnd != -1 {
			// Left corner (connects Top -> Bottom at iStart)
			{
				i0 := iStart
				i1 := iStart + 1
				if i1 >= numPts {
					i1 = i0
				}

				pTop0 := point2D{U: targetAngles[i0] * globalR, Y: yRotList[i0]}
				pTop1 := point2D{U: targetAngles[i1] * globalR, Y: yRotList[i1]}
				tTopExitU, tTopExitY := normalize2D(pTop0.U-pTop1.U, pTop0.Y-pTop1.Y, -1.0, 0.0)

				pBottom0 := point2D{U: targetAngles[i0] * globalR, Y: yBaseList[i0]}
				pBottom1 := point2D{U: targetAngles[i1] * globalR, Y: yBaseList[i1]}
				tBottomEnterU, tBottomEnterY := normalize2D(pBottom1.U-pBottom0.U, pBottom1.Y-pBottom0.Y, 1.0, 0.0)

				c0 := pTop0
				c3 := pBottom0
				chord := math.Abs(c3.Y - c0.Y)
				handleLen := controlStrength * chord

				c1 := point2D{U: c0.U + handleLen*tTopExitU, Y: c0.Y + handleLen*tTopExitY}
				c2 := point2D{U: c3.U - handleLen*tBottomEnterU, Y: c3.Y - handleLen*tBottomEnterY}

				leftCornerPts = getBezierPoints(c0, c1, c2, c3, "Left Corner")
			}

			// Right corner (connects Bottom -> Top at iEnd)
			{
				i0 := iEnd
				i1 := iEnd - 1
				if i1 < 0 {
					i1 = 0
				}

				pBottom0 := point2D{U: targetAngles[i0] * globalR, Y: yBaseList[i0]}
				pBottom1 := point2D{U: targetAngles[i1] * globalR, Y: yBaseList[i1]}
				tBottomExitU, tBottomExitY := normalize2D(pBottom0.U-pBottom1.U, pBottom0.Y-pBottom1.Y, 1.0, 0.0)

				pTop0 := point2D{U: targetAngles[i0] * globalR, Y: yRotList[i0]}
				pTop1 := point2D{U: targetAngles[i1] * globalR, Y: yRotList[i1]}
				tTopEnterU, tTopEnterY := normalize2D(pTop1.U-pTop0.U, pTop1.Y-pTop0.Y, -1.0, 0.0)

				c0 := pBottom0
				c3 := pTop0
				chord := math.Abs(c3.Y - c0.Y)
				handleLen := controlStrength * chord

				c1 := point2D{U: c0.U + handleLen*tBottomExitU, Y: c0.Y + handleLen*tBottomExitY}
				c2 := point2D{U: c3.U - handleLen*tTopEnterU, Y: c3.Y - handleLen*tTopEnterY}

				rightCornerPts = getBezierPoints(c0, c1, c2, c3, "Right Corner")
			}
		}

		// 14. Add 3D Eye Sampled Points visualization for first repetition if toggled on
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeSampledPoints3DShape {
			var eyePoints []*threejs.Vector3
			for i, inE := range inEye {
				if inE {
					alpha := targetAngles[i]
					x := globalR * math.Cos(alpha)
					z := globalR * math.Sin(alpha)

					eyePoints = append(eyePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Base Point %.1f", alpha*180.0/math.Pi),
						X:    x,
						Y:    yBaseList[i],
						Z:    z,
					}).Stage(stool3dStage))

					eyePoints = append(eyePoints, (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Rotated Point %.1f", alpha*180.0/math.Pi),
						X:    x,
						Y:    yRotList[i],
						Z:    z,
					}).Stage(stool3dStage))
				}
			}
			u.addPointSpheres(stool3dStage, eyePoints, "magenta", canvas, "Stool Eye Sampled", 0, 0)
		}

		// 15. Render 3D Eye Corners Sampled Points
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeCornersSampledPoints3DShape {
			var cornerPoints []*threejs.Vector3
			cornerPoints = append(cornerPoints, leftCornerPts...)
			cornerPoints = append(cornerPoints, rightCornerPts...)
			u.addPointSpheres(stool3dStage, cornerPoints, "cyan", canvas, "Stool Eye Corners Sampled", 0, 0)
		}

		var eye3DPoints []*threejs.Vector3
		if iStart != -1 && iEnd != -1 {
			// A. Bottom curve: from iStart to iEnd
			for i := iStart; i <= iEnd; i++ {
				alpha := targetAngles[i]
				x := globalR * math.Cos(alpha)
				z := globalR * math.Sin(alpha)
				eye3DPoints = append(eye3DPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Eye Loop Bottom %.1f", alpha*180.0/math.Pi),
					X:    x,
					Y:    yBaseList[i],
					Z:    z,
				}).Stage(stool3dStage))
			}

			// B. Right corner Bézier: Bottom -> Top
			eye3DPoints = append(eye3DPoints, rightCornerPts...)

			// C. Top curve: from iEnd down to iStart (reversed)
			for i := iEnd; i >= iStart; i-- {
				alpha := targetAngles[i]
				x := globalR * math.Cos(alpha)
				z := globalR * math.Sin(alpha)
				eye3DPoints = append(eye3DPoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Eye Loop Top %.1f", alpha*180.0/math.Pi),
					X:    x,
					Y:    yRotList[i],
					Z:    z,
				}).Stage(stool3dStage))
			}

			// D. Left corner Bézier: Top -> Bottom
			eye3DPoints = append(eye3DPoints, leftCornerPts...)
		}

		// 16. Add 3D Eye (Continuous interpolated closed tube)
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenEye3DShape {
			if len(eye3DPoints) > 0 {
				eyeLoopCurve := (&threejs.Curve{
					Name:   "Stool Eye Continuous Loop Curve",
					Points: eye3DPoints,
				}).Stage(stool3dStage)

				numSegments := len(eyeLoopCurve.Points)
				if numSegments < 2 {
					numSegments = 2
				}

				eyeGeom := (&threejs.TubeGeometry{
					Name:            "Stool Eye TubeGeom",
					Path:            eyeLoopCurve,
					TubularSegments: numSegments,
					Radius:          tubeRadius,
					RadialSegments:  16,
					Closed:          true,
				}).Stage(stool3dStage)

				eyeMesh := (&threejs.Mesh{
					Name:         "Stool Eye Mesh",
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: eyeGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Eye Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "deeppink"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, eyeMesh)
			}
		}

		var projSeatBottomEyePoints []*threejs.Vector3
		var projStoolBottomEyePoints []*threejs.Vector3

		if len(eye3DPoints) > 0 {
			stoolBottomHeight := 0.0
			for _, pt := range eye3DPoints {
				origTheta := math.Atan2(pt.Z, pt.X)
				r := math.Hypot(pt.X, pt.Z)

				// Seat Bottom projection
				deltaYSB := seatBottomHeight - pt.Y
				rProjSB := r + deltaYSB*math.Tan(projAngleRad)
				projSeatBottomEyePoints = append(projSeatBottomEyePoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Seat Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
					X:    rProjSB * math.Cos(origTheta),
					Y:    seatBottomHeight,
					Z:    rProjSB * math.Sin(origTheta),
				}).Stage(stool3dStage))

				// Stool Bottom projection
				deltaYStool := stoolBottomHeight - pt.Y
				rProjStool := r + deltaYStool*math.Tan(projAngleRad)
				projStoolBottomEyePoints = append(projStoolBottomEyePoints, (&threejs.Vector3{
					Name: fmt.Sprintf("Stool Bottom Eye Point %.1f", origTheta*180.0/math.Pi),
					X:    rProjStool * math.Cos(origTheta),
					Y:    stoolBottomHeight,
					Z:    rProjStool * math.Sin(origTheta),
				}).Stage(stool3dStage))
			}
		}

		// 17. Seat Bottom Eye 2D Projected Curve on horizontal seat bottom plane
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeSeatBottomCurveShape {
			if len(projSeatBottomEyePoints) > 0 {
				projSeatBottomCurve := (&threejs.Curve{
					Name:   "Stool Seat Bottom Eye Curve",
					Points: projSeatBottomEyePoints,
				}).Stage(stool3dStage)

				numSegments := len(projSeatBottomCurve.Points)
				if numSegments < 2 {
					numSegments = 2
				}

				sbEyeGeom := (&threejs.TubeGeometry{
					Name:            "Stool Seat Bottom Eye TubeGeom",
					Path:            projSeatBottomCurve,
					TubularSegments: numSegments,
					Radius:          tubeRadius,
					RadialSegments:  16,
					Closed:          true,
				}).Stage(stool3dStage)

				sbEyeMesh := (&threejs.Mesh{
					Name:         "Stool Seat Bottom Eye Mesh",
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: sbEyeGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Seat Bottom Eye Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "mediumvioletred"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, sbEyeMesh)
			}
		}

		// 18. Stool Bottom Eye 2D Projected Curve on horizontal stool bottom / floor plane (Y = 0)
		if checkedDiagram != nil && checkedDiagram.StoolDiagram != nil && !checkedDiagram.StoolDiagram.IsHiddenEyeStoolBottomCurveShape {
			if len(projStoolBottomEyePoints) > 0 {
				projStoolBottomCurve := (&threejs.Curve{
					Name:   "Stool Bottom Eye Curve",
					Points: projStoolBottomEyePoints,
				}).Stage(stool3dStage)

				numSegments := len(projStoolBottomCurve.Points)
				if numSegments < 2 {
					numSegments = 2
				}

				stoolBottomEyeGeom := (&threejs.TubeGeometry{
					Name:            "Stool Bottom Eye TubeGeom",
					Path:            projStoolBottomCurve,
					TubularSegments: numSegments,
					Radius:          tubeRadius,
					RadialSegments:  16,
					Closed:          true,
				}).Stage(stool3dStage)

				stoolBottomEyeMesh := (&threejs.Mesh{
					Name:         "Stool Bottom Eye Mesh",
					Position:     threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry: stoolBottomEyeGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Bottom Eye Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkviolet"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, stoolBottomEyeMesh)
			}
		}

		// 19. 3D Eye Volume Mesh between Seat Bottom Eye Curve and Stool Bottom Eye Curve
		if checkedDiagram == nil || checkedDiagram.StoolDiagram == nil || !checkedDiagram.StoolDiagram.IsHiddenEyeVolume3DShape {
			if len(projSeatBottomEyePoints) >= 3 && len(projSeatBottomEyePoints) == len(projStoolBottomEyePoints) {
				M := len(projSeatBottomEyePoints)
				eyeVolGeom := (&threejs.BufferGeometry{
					Name: "Stool Eye Volume BufferGeometry",
				}).Stage(stool3dStage)

				var sumTopX, sumTopZ, sumBottomX, sumBottomZ float64
				for i := 0; i < M; i++ {
					topV := (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Top V %d", i),
						X:    projSeatBottomEyePoints[i].X,
						Y:    projSeatBottomEyePoints[i].Y,
						Z:    projSeatBottomEyePoints[i].Z,
					}).Stage(stool3dStage)
					eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, topV)
					sumTopX += topV.X
					sumTopZ += topV.Z
				}

				for i := 0; i < M; i++ {
					botV := (&threejs.Vector3{
						Name: fmt.Sprintf("Eye Bottom V %d", i),
						X:    projStoolBottomEyePoints[i].X,
						Y:    projStoolBottomEyePoints[i].Y,
						Z:    projStoolBottomEyePoints[i].Z,
					}).Stage(stool3dStage)
					eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, botV)
					sumBottomX += botV.X
					sumBottomZ += botV.Z
				}

				topCenterIdx := len(eyeVolGeom.Vertices)
				topCenterV := (&threejs.Vector3{
					Name: "Eye Top Center",
					X:    sumTopX / float64(M),
					Y:    seatBottomHeight,
					Z:    sumTopZ / float64(M),
				}).Stage(stool3dStage)
				eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, topCenterV)

				botCenterIdx := len(eyeVolGeom.Vertices)
				botCenterV := (&threejs.Vector3{
					Name: "Eye Bottom Center",
					X:    sumBottomX / float64(M),
					Y:    0.0,
					Z:    sumBottomZ / float64(M),
				}).Stage(stool3dStage)
				eyeVolGeom.Vertices = append(eyeVolGeom.Vertices, botCenterV)

				// 1. Top face (facing +Y): (topCenter, nextI, i)
				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Eye Top Face %d", i),
						V1:   topCenterIdx,
						V2:   nextI,
						V3:   i,
					}).Stage(stool3dStage))
				}

				// 2. Bottom face (facing -Y): (botCenter, botI, botNextI)
				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					botI := M + i
					botNextI := M + nextI
					eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Eye Bottom Face %d", i),
						V1:   botCenterIdx,
						V2:   botI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				// 3. Side wall quads between Top and Bottom:
				for i := 0; i < M; i++ {
					nextI := (i + 1) % M
					topI := i
					topNextI := nextI
					botI := M + i
					botNextI := M + nextI

					// Triangle 1: (botI, topI, topNextI)
					eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Eye Wall T1 %d", i),
						V1:   botI,
						V2:   topI,
						V3:   topNextI,
					}).Stage(stool3dStage))

					// Triangle 2: (botI, topNextI, botNextI)
					eyeVolGeom.Faces = append(eyeVolGeom.Faces, (&threejs.Triangle{
						Name: fmt.Sprintf("Eye Wall T2 %d", i),
						V1:   botI,
						V2:   topNextI,
						V3:   botNextI,
					}).Stage(stool3dStage))
				}

				eyeVolMesh := (&threejs.Mesh{
					Name:           "Stool Eye Volume 3D Mesh",
					Position:       threejs.Position{X: 0, Y: 0, Z: 0},
					BufferGeometry: eyeVolGeom,
					MeshPhysicalMaterial: (&threejs.MeshPhysicalMaterial{
						Name:                 "Stool Eye Volume Material",
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "mediumorchid"},
						Transparent:          true,
						Opacity:              opacity,
					}).Stage(stool3dStage),
				}).Stage(stool3dStage)

				canvas.Meshs = append(canvas.Meshs, eyeVolMesh)
			}
		}
	}

	// Floor tiles that encompass the stool cylinder
	if floorMinY > 0.0 {
		floorMinY = 0.0
	}
	floorMinY = floorMinY - 2.0

	floorSize := globalR * 3.0
	if floorSize < 200 {
		floorSize = 200
	}
	gridSize := 20
	tileSize := floorSize / float64(gridSize)

	for i := -gridSize / 2; i < gridSize/2; i++ {
		for j := -gridSize / 2; j < gridSize/2; j++ {
			color := "white"
			if (i+j)%2 != 0 {
				color = "black"
			}

			tileMesh := (&threejs.Mesh{
				Name: "Floor Tile " + strconv.Itoa(i) + "-" + strconv.Itoa(j),
				Position: threejs.Position{
					X: float64(i)*tileSize + tileSize/2,
					Y: floorMinY - 0.05,
					Z: float64(j)*tileSize + tileSize/2,
				},
				BoxGeometry: (&threejs.BoxGeometry{
					Name:           "Tile Geometry",
					Width:          tileSize,
					Height:         0.1,
					Depth:          tileSize,
					WidthSegments:  1,
					HeightSegments: 1,
					DepthSegments:  1,
				}).Stage(stool3dStage),
				MeshMaterialBasic: (&threejs.MeshMaterialBasic{
					Name:                 "Tile Material " + color,
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
				}).Stage(stool3dStage),
			}).Stage(stool3dStage)

			canvas.Meshs = append(canvas.Meshs, tileMesh)
		}
	}

	stool3dStage.Commit()
}

func (u *Stool3DStageUpdater) appendArcPointsStool(stool3dStage *threejs.Stage, targetCurve *threejs.Curve, x1, y1, x2, y2, r float64, sweepFlag bool, largeArcFlag bool, globalR float64, baseThetaOffset float64, floorMinY *float64) {
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
		if i == 0 && len(targetCurve.Points) > 0 {
			continue
		}

		t := float64(i) / float64(steps)
		angle := startAngle + t*(endAngle-startAngle)
		x2d := cx + r*math.Cos(angle)
		y2d := cy + r*math.Sin(angle)

		theta := x2d/globalR + baseThetaOffset
		x3d := globalR * math.Cos(theta)
		z3d := globalR * math.Sin(theta)
		y3d := y2d

		if y3d < *floorMinY {
			*floorMinY = y3d
		}

		vec := (&threejs.Vector3{
			Name: fmt.Sprintf("Stool Base Point %d", len(targetCurve.Points)),
			X:    x3d,
			Y:    y3d,
			Z:    z3d,
		}).Stage(stool3dStage)
		targetCurve.Points = append(targetCurve.Points, vec)
	}
}

// unwrapAngles processes a 3D curve to extract a strictly monotonic,
// duplicate-free mapping of points to their cylindrical angle (theta).
func unwrapAngles(curve *threejs.Curve) (angles []float64, points []*threejs.Vector3) {
	angleToPoint := make(map[float64]*threejs.Vector3)
	if len(curve.Points) == 0 {
		return nil, nil
	}

	firstP := curve.Points[0]
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

	for i := 1; i < len(curve.Points); i++ {
		p := curve.Points[i]
		theta := math.Atan2(p.Z, p.X)
		diff := theta - lastTheta

		for diff < -math.Pi {
			diff += 2 * math.Pi
		}
		for diff > math.Pi {
			diff -= 2 * math.Pi
		}

		if diff < -1e-7 {
			log.Printf("overlapping segment detected: curve goes backwards at index %d (diff: %f)", i, diff)
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

// resampleCurveAtAngles forces an existing 3D curve to conform to a specific set of target angles.
func (u *Stool3DStageUpdater) resampleCurveAtAngles(
	stool3dStage *threejs.Stage,
	sortedAngles []float64,
	sortedPoints []*threejs.Vector3,
	targetAngles []float64,
	namePrefix string,
	expectedDegrees float64,
) *threejs.Curve {
	resampled := (&threejs.Curve{
		Name: fmt.Sprintf("%s Resampled", namePrefix),
	}).Stage(stool3dStage)

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

		idx := -1

		// Use binary search to find the first index where sortedAngles[i] >= evalTarget
		searchIdx := sort.SearchFloat64s(sortedAngles, evalTarget)

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

		pt := (&threejs.Vector3{
			Name: fmt.Sprintf("%s %.1f", namePrefix, target*180.0/math.Pi),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stool3dStage)

		resampled.Points = append(resampled.Points, pt)
	}

	return resampled
}

func (u *Stool3DStageUpdater) addPointSpheres(stool3dStage *threejs.Stage, points []*threejs.Vector3, color string, canvas *threejs.Canvas, namePrefix string, dy float64, numPointsPerRep int) {
	for i, pt := range points {
		sphereColor := color
		radius := 2.0

		localIdx := i
		if numPointsPerRep > 0 {
			localIdx = i % numPointsPerRep
		}

		if localIdx%20 == 0 {
			sphereColor = "yellow"
			radius = 4.0
		}

		sphere := (&threejs.Mesh{
			Name: fmt.Sprintf("%s Sphere %d", namePrefix, i),
			Position: threejs.Position{
				X: pt.X,
				Y: pt.Y + dy,
				Z: pt.Z,
			},
			SphereGeometry: (&threejs.SphereGeometry{
				Name:   fmt.Sprintf("%s SphereGeom %d", namePrefix, i),
				Radius: radius,
			}).Stage(stool3dStage),
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name: fmt.Sprintf("%s SphereMat %d", namePrefix, i),
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{
					Color: sphereColor,
				},
			}).Stage(stool3dStage),
		}).Stage(stool3dStage)
		canvas.Meshs = append(canvas.Meshs, sphere)
	}
}
