package stoolstage3d

import (
	"fmt"
	"math"
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
		// camera orientation preserved
	}

	floorMinY := math.MaxFloat64

	// Extract GrowthCurve2D arcs from PlantAbstract (StartArcShapeGrid and EndArcShapeGrid)
	if plant.StartArcShapeGrid != nil && len(plant.StartArcShapeGrid.StartArcShapes) > 0 {
		startArcs := plant.StartArcShapeGrid.StartArcShapes
		var endArcs []*models.EndArcShape
		if plant.EndArcShapeGrid != nil {
			endArcs = plant.EndArcShapeGrid.EndArcShapes
		}

		curve := (&threejs.Curve{
			Name: "Stool GrowthCurve2D Curve",
		}).Stage(stool3dStage)

		for k := 0; k < radialRepetitions; k++ {
			baseThetaOffset := float64(k) * 2.0 * math.Pi / float64(radialRepetitions)

			for i := 0; i < len(startArcs); i++ {
				sa := startArcs[i]
				u.appendArcPointsStool(stool3dStage, curve, sa.StartX, sa.StartY, sa.EndX, sa.EndY, sa.RadiusX, !sa.SweepFlag, sa.LargeArcFlag, globalR, baseThetaOffset, &floorMinY)

				if i < len(endArcs) {
					ea := endArcs[i]
					u.appendArcPointsStool(stool3dStage, curve, ea.StartX, ea.StartY, ea.EndX, ea.EndY, ea.RadiusX, !ea.SweepFlag, ea.LargeArcFlag, globalR, baseThetaOffset, &floorMinY)
				}
			}
		}

		tubeRadius := plant.RhombusSideLength * 0.02
		if tubeRadius <= 0 {
			tubeRadius = 2.0
		}

		numSegments := len(curve.Points)
		if numSegments < 2 {
			numSegments = 2
		}

		tGeom := (&threejs.TubeGeometry{
			Name:            "Stool GrowthCurve2D TubeGeom",
			Path:            curve,
			TubularSegments: numSegments,
			Radius:          tubeRadius,
			RadialSegments:  8,
			Closed:          true,
		}).Stage(stool3dStage)

		tubeMesh := (&threejs.Mesh{
			Name:         "Stool GrowthCurve2D Mesh",
			Position:     threejs.Position{X: 0, Y: 0, Z: 0},
			TubeGeometry: tGeom,
			MeshMaterialBasic: (&threejs.MeshMaterialBasic{
				Name:                 "Stool GrowthCurve2D Material",
				MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "darkgreen"},
			}).Stage(stool3dStage),
		}).Stage(stool3dStage)

		canvas.Meshs = append(canvas.Meshs, tubeMesh)
	}

	// Floor tiles that encompass the stool cylinder
	if floorMinY == math.MaxFloat64 {
		floorMinY = 0.0
	} else {
		floorMinY = floorMinY - 2.0
	}

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
			Name: fmt.Sprintf("Stool Point %d", len(targetCurve.Points)),
			X:    x3d,
			Y:    y3d,
			Z:    z3d,
		}).Stage(stool3dStage)
		targetCurve.Points = append(targetCurve.Points, vec)
	}
}
