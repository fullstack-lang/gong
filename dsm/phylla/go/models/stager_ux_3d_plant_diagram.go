package models

import (
	"fmt"
	"math"
	"strconv"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
)

func (stager *Stager) ux_3d_plant_diagram() {

	if true {
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

		if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: checkedDiagram.Rendered3DShape.ViewX,
					Y: checkedDiagram.Rendered3DShape.ViewY,
					Z: checkedDiagram.Rendered3DShape.ViewZ,
				},
			}).Stage(stager.threejsStage)
		} else {
			canvas.Camera = (&threejs.Camera{
				Name: "Default Camera",
				Position: threejs.Position{
					X: 15,
					Y: 15,
					Z: 15,
				},
			}).Stage(stager.threejsStage)
		}

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

		if checkedDiagram != nil && checkedDiagram.Rendered3DShape != nil {
			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: checkedDiagram.Rendered3DShape.ViewX,
					Y: checkedDiagram.Rendered3DShape.ViewY,
					Z: checkedDiagram.Rendered3DShape.ViewZ,
				},
				TargetX: checkedDiagram.Rendered3DShape.TargetX,
				TargetY: checkedDiagram.Rendered3DShape.TargetY,
				TargetZ: checkedDiagram.Rendered3DShape.TargetZ,
				Fov:     checkedDiagram.Rendered3DShape.Fov,
			}).Stage(stager.threejsStage)
		} else {
			canvas.Camera = (&threejs.Camera{
				Name: "Camera",
				Position: threejs.Position{
					X: 15,
					Y: 15,
					Z: 15,
				},
				TargetX: 0,
				TargetY: 0,
				TargetZ: 0,
			}).Stage(stager.threejsStage)
		}
		tileSize := 10.0
		gridSize := 10

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
						Y: -0.05,
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
					}).Stage(stager.threejsStage),
					MeshMaterialBasic: (&threejs.MeshMaterialBasic{
						Name:                 "Tile Material " + color,
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage),
				}).Stage(stager.threejsStage)

				canvas.Meshs = append(canvas.Meshs, tileMesh)
			}
		}

		// tubes for the stack
		if plant.StackOfGrowthCurve != nil && len(plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes) > 0 {
			circumference := 10.0 // default fallback
			if plant.PlantCircumferenceShape != nil && plant.PlantCircumferenceShape.Length > 0 {
				circumference = plant.PlantCircumferenceShape.Length
			}
			R := circumference / (2 * math.Pi)

			for bezierIndex, bezier := range plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes {
				curve := (&threejs.Curve{
					Name: fmt.Sprintf("Growth Stack Curve %d", bezierIndex),
				}).Stage(stager.threejsStage)

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
						Name: fmt.Sprintf("Point %d-%d", bezierIndex, i),
						X:    x3d,
						Y:    y3d,
						Z:    z3d,
					}).Stage(stager.threejsStage)
					curve.Points = append(curve.Points, vec)
				}

				if len(curve.Points) > 1 {
					tubeGeometry := (&threejs.TubeGeometry{
						Name:            fmt.Sprintf("Growth Stack Tube Geometry %d", bezierIndex),
						Path:            curve,
						TubularSegments: len(curve.Points) * 2,
						Radius:          R * 0.05, // scale tube radius proportional to cylinder radius
						RadialSegments:  16,
						Closed:          false,
					}).Stage(stager.threejsStage)

					colors := []string{"#2E8B57", "#3CB371", "#228B22", "#006400", "#556B2F"}
					color := colors[bezierIndex%len(colors)]

					tubeMaterial := (&threejs.MeshPhysicalMaterial{
						Name:                 fmt.Sprintf("Tube Material %d", bezierIndex),
						MeshMaterialAbstract: threejs.MeshMaterialAbstract{Color: color},
					}).Stage(stager.threejsStage)

					tubeMesh := (&threejs.Mesh{
						Name:                 fmt.Sprintf("Growth Stack Tube Mesh %d", bezierIndex),
						Position:             threejs.Position{X: 0, Y: 0, Z: 0},
						TubeGeometry:         tubeGeometry,
						MeshPhysicalMaterial: tubeMaterial,
					}).Stage(stager.threejsStage)

					canvas.Meshs = append(canvas.Meshs, tubeMesh)
				}
			}
		}

		stager.threejsStage.Commit()
	} else {
		threejs.GenerateReferenceScene(stager.threejsStage)
	}
}
