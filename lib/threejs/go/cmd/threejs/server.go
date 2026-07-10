//go:build !js

package main

import (
	"log"
	"math"
	"strconv"

	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	threejs_stack "github.com/fullstack-lang/gong/lib/threejs/go/stack"
	threejs_static "github.com/fullstack-lang/gong/lib/threejs/go/static"
)

func executeServer() {

	r := threejs_static.ServeStaticFiles(logGINFlag)

	// setup
	// - model stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := threejs_stack.NewStack(r, "threejs", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	if generateStrangeFormsFlag {
		generateStrangeForms(stack)
	} else {
		stack.Stage.Commit()
	}

	// initiates the UX loop
	models.NewStager(
		r,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func generateStrangeForms(stack *threejs_stack.Stack) {
	var canvas *models.Canvas
	for c := range *models.GetGongstructInstancesSet[models.Canvas](stack.Stage) {
		if c.Name == "Singloton" {
			canvas = c
			break
		}
	}
	if canvas == nil {
		canvas = (&models.Canvas{
			Name: "Singloton",
		}).Stage(stack.Stage)
	}

	if canvas != nil {
		if canvas.AmbiantLight == nil {
			canvas.AmbiantLight = (&models.AmbiantLight{
				Name: "Ambiant Light",
				LightAbstract: models.LightAbstract{
					Intensity: 0.5,
				},
			}).Stage(stack.Stage)
		}

		if len(canvas.DirectionalLights) == 0 {
			canvas.DirectionalLights = append(canvas.DirectionalLights, (&models.DirectionalLight{
				Name: "Directional Light",
				Position: models.Position{X: 10, Y: 10, Z: 10},
				LightAbstract: models.LightAbstract{
					Intensity: 1.0,
				},
				IsWithCastShadow: true,
			}).Stage(stack.Stage))
		}

		var hasPlane bool
		for _, m := range canvas.Meshs {
			if m.Name == "Wall" {
				hasPlane = true
				break
			}
		}
		if !hasPlane {
			planeMesh := (&models.Mesh{
				Name: "Wall",
				Position: models.Position{X: 0, Y: 0, Z: -30},
				PlaneGeometry: (&models.PlaneGeometry{
					Name: "Wall Plane",
					Width: 200,
					Height: 200,
					WidthSegments: 1,
					HeightSegments: 1,
				}).Stage(stack.Stage),
				MeshMaterialBasic: (&models.MeshMaterialBasic{Name: "Gray", MeshMaterialAbstract: models.MeshMaterialAbstract{Color: "gray"}}).Stage(stack.Stage),
			}).Stage(stack.Stage)
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
			sphereMesh := (&models.Mesh{
				Name: "Sphere",
				Position: models.Position{X: 5, Y: 5, Z: -5},
				SphereGeometry: (&models.SphereGeometry{
					Name: "Sphere Geometry",
					Radius: 2,
					WidthSegments: 32,
					HeightSegments: 16,
					PhiLength: math.Pi * 2,
					ThetaLength: math.Pi,
				}).Stage(stack.Stage),
				MeshMaterialBasic: (&models.MeshMaterialBasic{Name: "Cyan", MeshMaterialAbstract: models.MeshMaterialAbstract{Color: "cyan"}}).Stage(stack.Stage),
			}).Stage(stack.Stage)
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
			curve := (&models.Curve{
				Name: "Helix Path",
			}).Stage(stack.Stage)

			turns := 4.0
			radius := 2.0
			height := 6.0
			for i := 0; i <= 200; i++ {
				t := float64(i) / 200.0
				angle := t * 3.141592653589793 * 2 * turns
				x := math.Cos(angle) * radius
				z := math.Sin(angle) * radius
				y := t * height

				vec := (&models.Vector3{
					Name: "Point " + strconv.Itoa(i),
					X:    x,
					Y:    y,
					Z:    z,
				}).Stage(stack.Stage)
				curve.Points = append(curve.Points, vec)
			}

			tubeGeometry := (&models.TubeGeometry{
				Name:            "Helix Tube Geometry",
				Path:            curve,
				TubularSegments: 200,
				Radius:          0.4,
				RadialSegments:  16,
				Closed:          false,
			}).Stage(stack.Stage)

			tubeMesh := (&models.Mesh{
				Name:              "Helix Mesh",
				Position:          models.Position{X: 0, Y: 3.5, Z: -10},
				TubeGeometry:      tubeGeometry,
				MeshMaterialBasic: (&models.MeshMaterialBasic{Name: "Silver", MeshMaterialAbstract: models.MeshMaterialAbstract{Color: "silver"}}).Stage(stack.Stage),
			}).Stage(stack.Stage)

			canvas.Meshs = append(canvas.Meshs, tubeMesh)
			stack.Stage.Commit()
		}

		var hasExtrude bool
		for _, m := range canvas.Meshs {
			if m.Name == "Extrude Mesh" {
				hasExtrude = true
				break
			}
		}
		if !hasExtrude {
			extrudePath := (&models.Curve{
				Name: "Torus Like Curve",
			}).Stage(stack.Stage)

			radius := 10.0
			waves := 5.0
			amplitude := 2.0
			for i := 0; i <= 500; i++ {
				t := float64(i) / 500.0
				angle := t * 3.141592653589793 * 2
				x := math.Cos(angle) * radius
				z := math.Sin(angle) * radius
				y := math.Sin(angle*waves) * amplitude

				vec := (&models.Vector3{
					Name: "TorusPoint " + strconv.Itoa(i),
					X:    x,
					Y:    y,
					Z:    z,
				}).Stage(stack.Stage)
				extrudePath.Points = append(extrudePath.Points, vec)
			}

			shape := (&models.Shape{
				Name: "Square Shape",
			}).Stage(stack.Stage)

			size := 0.75
			pts := [][2]float64{
				{-size, -size},
				{size, -size},
				{size, size},
				{-size, size},
				{-size, -size},
			}
			for i, p := range pts {
				v := (&models.Vector2{
					Name: "SquarePoint " + strconv.Itoa(i),
					X:    p[0],
					Y:    p[1],
				}).Stage(stack.Stage)
				shape.Points = append(shape.Points, v)
			}

			extrudeGeom := (&models.ExtrudeGeometry{
				Name:        "Torus Extrude Geometry",
				ExtrudePath: extrudePath,
				Shape:       shape,
				Steps:       500,
			}).Stage(stack.Stage)

			extrudeMesh := (&models.Mesh{
				Name:              "Extrude Mesh",
				Position:          models.Position{X: 0, Y: 5, Z: 0},
				ExtrudeGeometry:   extrudeGeom,
				MeshMaterialBasic: (&models.MeshMaterialBasic{Name: "Magenta", MeshMaterialAbstract: models.MeshMaterialAbstract{Color: "magenta"}}).Stage(stack.Stage),
			}).Stage(stack.Stage)

			canvas.Meshs = append(canvas.Meshs, extrudeMesh)
			stack.Stage.Commit()
		}
	}
}
