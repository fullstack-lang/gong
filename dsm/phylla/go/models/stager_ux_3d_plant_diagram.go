package models

import (
	"fmt"
	"math"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) ux_3d_plant_diagram() {

	if false {
		stager.threejsStage.Reset()

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
		threejs.GenerateReferenceScene(stager.threejsStage)
	}
}
