package models

import (
	"math"
	"strconv"
)

func GenerateStrangeForms(stage *Stage) {
	var canvas *Canvas
	for c := range *GetGongstructInstancesSet[Canvas](stage) {
		if c.Name == "Singloton" {
			canvas = c
			break
		}
	}
	if canvas == nil {
		canvas = (&Canvas{
			Name: "Singloton",
		}).Stage(stage)
	}

	if canvas != nil {
		if canvas.AmbiantLight == nil {
			canvas.AmbiantLight = (&AmbiantLight{
				Name: "Ambiant Light",
				LightAbstract: LightAbstract{
					Intensity: 0.5,
				},
			}).Stage(stage)
		}

		if len(canvas.DirectionalLights) == 0 {
			canvas.DirectionalLights = append(canvas.DirectionalLights, (&DirectionalLight{
				Name: "Directional Light",
				Position: Position{X: 10, Y: 10, Z: 10},
				LightAbstract: LightAbstract{
					Intensity: 1.0,
				},
				IsWithCastShadow: true,
			}).Stage(stage))
		}

		var hasPlane bool
		for _, m := range canvas.Meshs {
			if m.Name == "Wall" {
				hasPlane = true
				break
			}
		}
		if !hasPlane {
			planeMesh := (&Mesh{
				Name: "Wall",
				Position: Position{X: 0, Y: 0, Z: -30},
				PlaneGeometry: (&PlaneGeometry{
					Name: "Wall Plane",
					Width: 200,
					Height: 200,
					WidthSegments: 1,
					HeightSegments: 1,
				}).Stage(stage),
				MeshMaterialBasic: (&MeshMaterialBasic{Name: "Gray", MeshMaterialAbstract: MeshMaterialAbstract{Color: "gray"}}).Stage(stage),
			}).Stage(stage)
			canvas.Meshs = append(canvas.Meshs, planeMesh)
		}

		var hasSphere bool
		for _, m := range canvas.Meshs {
			if m.Name == "Sphere" {
				hasSphere = true
				break
			}
		}
		if !hasSphere {
			sphereMesh := (&Mesh{
				Name: "Sphere",
				Position: Position{X: 5, Y: 5, Z: -5},
				SphereGeometry: (&SphereGeometry{
					Name: "Sphere Geometry",
					Radius: 2,
					WidthSegments: 32,
					HeightSegments: 16,
					PhiLength: math.Pi * 2,
					ThetaLength: math.Pi,
				}).Stage(stage),
				MeshMaterialBasic: (&MeshMaterialBasic{Name: "Cyan", MeshMaterialAbstract: MeshMaterialAbstract{Color: "cyan"}}).Stage(stage),
			}).Stage(stage)
			canvas.Meshs = append(canvas.Meshs, sphereMesh)
		}

		var hasHelix bool
		for _, m := range canvas.Meshs {
			if m.Name == "Helix Mesh" {
				hasHelix = true
				break
			}
		}
		if !hasHelix {
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
		}

		var hasExtrude bool
		for _, m := range canvas.Meshs {
			if m.Name == "Extrude Mesh" {
				hasExtrude = true
				break
			}
		}
		if !hasExtrude {
			extrudePath := (&Curve{
				Name: "Torus Like Curve",
			}).Stage(stage)

			radius := 10.0
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

			shape := (&Shape{
				Name: "Square Shape",
			}).Stage(stage)

			size := 0.75
			pts := [][2]float64{
				{-size, -size},
				{size, -size},
				{size, size},
				{-size, size},
				{-size, -size},
			}
			for i, p := range pts {
				v := (&Vector2{
					Name: "SquarePoint " + strconv.Itoa(i),
					X:    p[0],
					Y:    p[1],
				}).Stage(stage)
				shape.Points = append(shape.Points, v)
			}

			extrudeGeom := (&ExtrudeGeometry{
				Name:        "Torus Extrude Geometry",
				ExtrudePath: extrudePath,
				Shape:       shape,
				Steps:       500,
			}).Stage(stage)

			extrudeMesh := (&Mesh{
				Name:              "Extrude Mesh",
				Position:          Position{X: 0, Y: 5, Z: 0},
				ExtrudeGeometry:   extrudeGeom,
				MeshMaterialBasic: (&MeshMaterialBasic{Name: "Magenta", MeshMaterialAbstract: MeshMaterialAbstract{Color: "magenta"}}).Stage(stage),
			}).Stage(stage)

			canvas.Meshs = append(canvas.Meshs, extrudeMesh)
		}
	}
	stage.Commit()
}
