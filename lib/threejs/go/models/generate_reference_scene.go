package models

import (
	"math"
	"strconv"
)

func GenerateReferenceScene(stage *Stage) {
	stage.Reset()

	canvas := (&Canvas{
		Name: "Singloton",
	}).Stage(stage)

	canvas.Camera = (&Camera{
		Name:     "Camera",
		Position: Position{X: 30, Y: 30, Z: 30},
		TargetX:  0,
		TargetY:  5,
		TargetZ:  0,
		Fov:      45,
	}).Stage(stage)

	canvas.AmbiantLight = (&AmbiantLight{
		Name: "Ambiant Light",
		LightAbstract: LightAbstract{
			Intensity: 0.3,
		},
	}).Stage(stage)

	dirLight1 := (&DirectionalLight{
		Name:     "Directional Light 1 (Key)",
		Position: Position{X: 15, Y: 20, Z: 15},
		LightAbstract: LightAbstract{
			Intensity: 1.2,
		},
		IsWithCastShadow: true,
	}).Stage(stage)

	dirLight2 := (&DirectionalLight{
		Name:     "Directional Light 2 (Fill)",
		Position: Position{X: -15, Y: 10, Z: -15},
		LightAbstract: LightAbstract{
			Intensity: 0.6,
		},
		IsWithCastShadow: false,
	}).Stage(stage)

	dirLight3 := (&DirectionalLight{
		Name:     "Directional Light 3 (Rim)",
		Position: Position{X: 0, Y: 30, Z: -20},
		LightAbstract: LightAbstract{
			Intensity: 0.8,
		},
		IsWithCastShadow: false,
	}).Stage(stage)

	canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight1, dirLight2, dirLight3)

	tileSize := 10.0
	tilesCount := 10

	boxGeom := (&BoxGeometry{
		Name:           "Floor Tile Geometry",
		Width:          tileSize,
		Height:         0.1,
		Depth:          tileSize,
		WidthSegments:  1,
		HeightSegments: 1,
		DepthSegments:  1,
	}).Stage(stage)

	color1 := (&MeshMaterialBasic{Name: "TileColor1", MeshMaterialAbstract: MeshMaterialAbstract{Color: "white"}}).Stage(stage)
	color2 := (&MeshMaterialBasic{Name: "TileColor2", MeshMaterialAbstract: MeshMaterialAbstract{Color: "black"}}).Stage(stage)

	for i := 0; i < tilesCount; i++ {
		for j := 0; j < tilesCount; j++ {
			var mat *MeshMaterialBasic
			if (i+j)%2 == 0 {
				mat = color1
			} else {
				mat = color2
			}

			tileMesh := (&Mesh{
				Name: "Floor Tile " + strconv.Itoa(i) + "-" + strconv.Itoa(j),
				Position: Position{
					X: (float64(i) - float64(tilesCount)/2.0 + 0.5) * tileSize,
					Y: -0.05,
					Z: (float64(j) - float64(tilesCount)/2.0 + 0.5) * tileSize,
				},
				BoxGeometry:       boxGeom,
				MeshMaterialBasic: mat,
			}).Stage(stage)

			canvas.Meshs = append(canvas.Meshs, tileMesh)
		}
	}

	planeMesh := (&Mesh{
		Name:     "Wall",
		Position: Position{X: 0, Y: 0, Z: -30},
		PlaneGeometry: (&PlaneGeometry{
			Name:           "Wall Plane",
			Width:          200,
			Height:         200,
			WidthSegments:  1,
			HeightSegments: 1,
		}).Stage(stage),
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Gray", MeshMaterialAbstract: MeshMaterialAbstract{Color: "gray"}}).Stage(stage),
	}).Stage(stage)
	canvas.Meshs = append(canvas.Meshs, planeMesh)

	sphereMesh := (&Mesh{
		Name:     "Sphere",
		Position: Position{X: 5, Y: 5, Z: -5},
		SphereGeometry: (&SphereGeometry{
			Name:           "Sphere Geometry",
			Radius:         2,
			WidthSegments:  32,
			HeightSegments: 16,
			PhiLength:      math.Pi * 2,
			ThetaLength:    math.Pi,
		}).Stage(stage),
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Cyan", MeshMaterialAbstract: MeshMaterialAbstract{Color: "cyan"}}).Stage(stage),
	}).Stage(stage)
	canvas.Meshs = append(canvas.Meshs, sphereMesh)

	curve := (&Curve{
		Name: "Helix Path",
	}).Stage(stage)

	turns := 4.0
	radius := 2.0
	height := 6.0
	for i := 0; i <= 200; i++ {
		t := float64(i) / 200.0
		angle := t * math.Pi * 2 * turns
		x := math.Cos(angle) * radius
		z := math.Sin(angle) * radius
		y := t * height

		vec := (&Vector3{
			Name: "Point " + strconv.Itoa(i),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stage)
		curve.Points = append(curve.Points, vec)
	}

	tubeGeometry := (&TubeGeometry{
		Name:            "Helix Tube Geometry",
		Path:            curve,
		TubularSegments: 200,
		Radius:          0.4,
		RadialSegments:  16,
		Closed:          false,
	}).Stage(stage)

	tubeMesh := (&Mesh{
		Name:              "Helix Mesh",
		Position:          Position{X: 0, Y: 3.5, Z: -10},
		TubeGeometry:      tubeGeometry,
		MeshMaterialBasic: (&MeshMaterialBasic{Name: "Silver", MeshMaterialAbstract: MeshMaterialAbstract{Color: "silver"}}).Stage(stage),
	}).Stage(stage)

	canvas.Meshs = append(canvas.Meshs, tubeMesh)

	extrudePath := (&Curve{
		Name: "Torus Like Curve",
	}).Stage(stage)

	radius = 10.0
	waves := 5.0
	amplitude := 2.0
	for i := 0; i <= 500; i++ {
		t := float64(i) / 500.0
		angle := t * math.Pi * 2
		x := math.Cos(angle) * radius
		z := math.Sin(angle) * radius
		y := math.Sin(angle*waves) * amplitude

		vec := (&Vector3{
			Name: "TorusPoint " + strconv.Itoa(i),
			X:    x,
			Y:    y,
			Z:    z,
		}).Stage(stage)
		extrudePath.Points = append(extrudePath.Points, vec)
	}

	thickness := 1.5 // 2 * size (0.75)
	e := thickness * 0.05

	createFaceMesh := func(name string, color string, xMin, xMax, yMin, yMax float64) *Mesh {
		shape := (&Shape{
			Name: "Torus " + name + " Shape",
		}).Stage(stage)

		pts := [][2]float64{
			{xMin, yMin},
			{xMax, yMin},
			{xMax, yMax},
			{xMin, yMax},
			{xMin, yMin},
		}
		for i, p := range pts {
			v := (&Vector2{
				Name: "SquarePoint " + name + " " + strconv.Itoa(i),
				X:    p[0],
				Y:    p[1],
			}).Stage(stage)
			shape.Points = append(shape.Points, v)
		}

		extrudeGeom := (&ExtrudeGeometry{
			Name:        "Torus " + name + " Extrude Geometry",
			ExtrudePath: extrudePath,
			Shape:       shape,
			Steps:       500,
		}).Stage(stage)

		return (&Mesh{
			Name:            "Torus " + name + " Mesh",
			Position:        Position{X: 0, Y: 5, Z: 0},
			ExtrudeGeometry: extrudeGeom,
			MeshPhysicalMaterial: (&MeshPhysicalMaterial{
				Name:                 "Torus " + name + " Material",
				MeshMaterialAbstract: MeshMaterialAbstract{Color: color},
			}).Stage(stage),
		}).Stage(stage)
	}

	// Create 4 faces centered at 0,0
	half := thickness / 2.0
	bottomFace := createFaceMesh("Bottom", "#1f77b4", -half, half, -half, -half+e)
	topFace := createFaceMesh("Top", "#d62728", -half, half, half-e, half)
	innerFace := createFaceMesh("Inner", "#ff7f0e", -half, -half+e, -half, half)
	outerFace := createFaceMesh("Outer", "#2ca02c", half-e, half, -half, half)

	canvas.Meshs = append(canvas.Meshs, bottomFace, topFace, innerFace, outerFace)

	stage.Commit()
}
