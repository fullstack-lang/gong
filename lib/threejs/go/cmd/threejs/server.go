//go:build !js

package main

import (
	"log"
	"math"
	"strconv"

	"github.com/fullstack-lang/gong/lib/threejs/go/level1stack"
	"github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func executeServer() {

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("threejs", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()
	stack.Stage.Commit()

	var canvas *models.Canvas
	for c := range *models.GetGongstructInstancesSet[models.Canvas](stack.Stage) {
		if c.Name == "Singloton" {
			canvas = c
			break
		}
	}
	if canvas != nil {
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

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
