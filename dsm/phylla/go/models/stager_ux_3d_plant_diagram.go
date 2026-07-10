package models

import (
	"fmt"
	"math"
	"strconv"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) ux_3d_plant_diagram() {
	stager.threejsStage.Reset()

	if false {

		plant := stager.GetCurrentPlant()
		if plant == nil {
			stager.threejsStage.Commit()
			return
		}

		canvas := (&threejs.Canvas{
			Name: "Plant 3D Canvas",
		}).Stage(stager.threejsStage)

		// lights
		dirLight := (&threejs.DirectionalLight{
			Name:             "Directional Light",
			Position:         threejs.Position{X: 10, Y: 20, Z: 10},
			LightAbstract:    threejs.LightAbstract{Intensity: 1.0},
			IsWithCastShadow: true,
		}).Stage(stager.threejsStage)
		canvas.DirectionalLights = append(canvas.DirectionalLights, dirLight)

		ambiantLight := (&threejs.AmbiantLight{
			Name:          "Ambiant Light",
			LightAbstract: threejs.LightAbstract{Intensity: 0.5},
		}).Stage(stager.threejsStage)
		canvas.AmbiantLight = ambiantLight

		// floor
		floorMaterial := (&threejs.MeshMaterialBasic{
			Name:                 "Floor Material",
			MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "lightgray"},
		}).Stage(stager.threejsStage)

		floorGeometry := (&threejs.BoxGeometry{
			Name:           "Floor Geometry",
			Width:          100,
			Height:         0.1,
			Depth:          100,
			WidthSegments:  1,
			HeightSegments: 1,
			DepthSegments:  1,
		}).Stage(stager.threejsStage)

		floorMesh := (&threejs.Mesh{
			Name:              "Floor Mesh",
			Position:          threejs.Position{X: 0, Y: -0.05, Z: 0},
			BoxGeometry:       floorGeometry,
			MeshMaterialBasic: floorMaterial,
		}).Stage(stager.threejsStage)
		canvas.Meshs = append(canvas.Meshs, floorMesh)

		// tube for the stack
		if plant.StackOfGrowthCurve != nil && len(plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes) > 0 {
			circumference := 10.0 // default fallback
			if plant.PlantCircumferenceShape != nil && plant.PlantCircumferenceShape.Length > 0 {
				circumference = plant.PlantCircumferenceShape.Length
			}
			R := circumference / (2 * math.Pi)

			curve := (&threejs.Curve{
				Name: "Growth Stack Curve",
			}).Stage(stager.threejsStage)

			pointIndex := 0
			for _, bezier := range plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes {
				// interpolate 20 steps per bezier curve
				steps := 20
				for i := 0; i <= steps; i++ {
					t := float64(i) / float64(steps)

					// cubic bezier interpolation
					x2d := math.Pow(1-t, 3)*bezier.StartX +
						3*math.Pow(1-t, 2)*t*bezier.ControlPointStartX +
						3*(1-t)*math.Pow(t, 2)*bezier.ControlPointEndX +
						math.Pow(t, 3)*bezier.EndX

					y2d := math.Pow(1-t, 3)*bezier.StartY +
						3*math.Pow(1-t, 2)*t*bezier.ControlPointStartY +
						3*(1-t)*math.Pow(t, 2)*bezier.ControlPointEndY +
						math.Pow(t, 3)*bezier.EndY

					// wrap around the cylinder
					theta := x2d / R
					x3d := R * math.Cos(theta)
					z3d := R * math.Sin(theta)
					y3d := y2d

					vec := (&threejs.Vector3{
						Name: fmt.Sprintf("Point %d", pointIndex),
						X:    x3d,
						Y:    y3d,
						Z:    z3d,
					}).Stage(stager.threejsStage)
					curve.Points = append(curve.Points, vec)
					pointIndex++
				}
			}

			if len(curve.Points) > 1 {
				tubeGeometry := (&threejs.TubeGeometry{
					Name:            "Growth Stack Tube Geometry",
					Path:            curve,
					TubularSegments: len(curve.Points) * 2,
					Radius:          R * 0.05, // scale tube radius proportional to cylinder radius
					RadialSegments:  16,
					Closed:          false,
				}).Stage(stager.threejsStage)

				tubeMaterial := (&threejs.MeshMaterialBasic{
					Name:                 "Tube Material",
					MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "green"},
				}).Stage(stager.threejsStage)

				tubeMesh := (&threejs.Mesh{
					Name:              "Growth Stack Tube Mesh",
					Position:          threejs.Position{X: 0, Y: 0, Z: 0},
					TubeGeometry:      tubeGeometry,
					MeshMaterialBasic: tubeMaterial,
				}).Stage(stager.threejsStage)

				canvas.Meshs = append(canvas.Meshs, tubeMesh)
			}
		}

		stager.threejsStage.Commit()
	} else {
		generateStrangeForms(stager.threejsStage)
	}
}

func generateStrangeForms(threejsStage *threejs.Stage) {
	threejsStage.Reset()

	var canvas *threejs.Canvas
	for c := range *threejs.GetGongstructInstancesSet[threejs.Canvas](threejsStage) {
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
			curve := (&threejs.Curve{
				Name: "Helix Path",
			}).Stage(threejsStage)

			turns := 4.0
			radius := 2.0
			height := 6.0
			for i := 0; i <= 200; i++ {
				t := float64(i) / 200.0
				angle := t * 3.141592653589793 * 2 * turns
				x := math.Cos(angle) * radius
				z := math.Sin(angle) * radius
				y := t * height

				vec := (&threejs.Vector3{
					Name: "Point " + strconv.Itoa(i),
					X:    x,
					Y:    y,
					Z:    z,
				}).Stage(threejsStage)
				curve.Points = append(curve.Points, vec)
			}

			tubeGeometry := (&threejs.TubeGeometry{
				Name:            "Helix Tube Geometry",
				Path:            curve,
				TubularSegments: 200,
				Radius:          0.4,
				RadialSegments:  16,
				Closed:          false,
			}).Stage(threejsStage)

			tubeMesh := (&threejs.Mesh{
				Name:              "Helix Mesh",
				Position:          threejs.Position{X: 0, Y: 3.5, Z: -10},
				TubeGeometry:      tubeGeometry,
				MeshMaterialBasic: (&threejs.MeshMaterialBasic{Name: "Silver", MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "silver"}}).Stage(threejsStage),
			}).Stage(threejsStage)

			canvas.Meshs = append(canvas.Meshs, tubeMesh)
			threejsStage.Commit()
		}

		var hasExtrude bool
		for _, m := range canvas.Meshs {
			if m.Name == "Extrude Mesh" {
				hasExtrude = true
				break
			}
		}
		if !hasExtrude {
			extrudePath := (&threejs.Curve{
				Name: "Torus Like Curve",
			}).Stage(threejsStage)

			radius := 10.0
			waves := 5.0
			amplitude := 2.0
			for i := 0; i <= 500; i++ {
				t := float64(i) / 500.0
				angle := t * 3.141592653589793 * 2
				x := math.Cos(angle) * radius
				z := math.Sin(angle) * radius
				y := math.Sin(angle*waves) * amplitude

				vec := (&threejs.Vector3{
					Name: "TorusPoint " + strconv.Itoa(i),
					X:    x,
					Y:    y,
					Z:    z,
				}).Stage(threejsStage)
				extrudePath.Points = append(extrudePath.Points, vec)
			}

			shape := (&threejs.Shape{
				Name: "Square Shape",
			}).Stage(threejsStage)

			size := 0.75
			pts := [][2]float64{
				{-size, -size},
				{size, -size},
				{size, size},
				{-size, size},
				{-size, -size},
			}
			for i, p := range pts {
				v := (&threejs.Vector2{
					Name: "SquarePoint " + strconv.Itoa(i),
					X:    p[0],
					Y:    p[1],
				}).Stage(threejsStage)
				shape.Points = append(shape.Points, v)
			}

			extrudeGeom := (&threejs.ExtrudeGeometry{
				Name:        "Torus Extrude Geometry",
				ExtrudePath: extrudePath,
				Shape:       shape,
				Steps:       500,
			}).Stage(threejsStage)

			extrudeMesh := (&threejs.Mesh{
				Name:              "Extrude Mesh",
				Position:          threejs.Position{X: 0, Y: 5, Z: 0},
				ExtrudeGeometry:   extrudeGeom,
				MeshMaterialBasic: (&threejs.MeshMaterialBasic{Name: "Magenta", MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: "magenta"}}).Stage(threejsStage),
			}).Stage(threejsStage)

			canvas.Meshs = append(canvas.Meshs, extrudeMesh)
		}
	}
	threejsStage.Commit()
}
